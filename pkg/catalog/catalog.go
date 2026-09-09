package catalog

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"iter"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	// DefaultConfigFile is the default YAML catalog configuration filename.
	DefaultConfigFile = "geofabrik.yml"

	// DefaultService is the default provider service identifier.
	DefaultService = "geofabrik"

	maxHierarchyDepth  = 30
	defaultDirPerm     = 0o750
	defaultFilePerm    = 0o600
	defaultBuilderSize = 64
	twoBaseURLParts    = 2
)

var (
	// ErrNilElement is returned when an operation receives an unexpected nil Element pointer.
	ErrNilElement = errors.New("nil element")

	// ErrElementNotFound is returned when an element ID cannot be resolved in the catalog.
	ErrElementNotFound = errors.New("element not found")

	// ErrFormatNotFound is returned when a requested file format does not exist for an element.
	ErrFormatNotFound = errors.New("format not found")

	// ErrParentMismatch is returned when merging an element with a conflicting parent identifier.
	ErrParentMismatch = errors.New("cannot merge element with conflicting parent")

	// ErrMaxHierarchyDepth is returned when resolving hierarchical URLs exceeds the maximum depth limit, indicating a cycle.
	ErrMaxHierarchyDepth = errors.New("maximum hierarchy depth exceeded (possible cycle in catalog)")

	// ErrResolveURL is returned when an element's download URL cannot be constructed.
	ErrResolveURL = errors.New("can't find url")

	// ErrElem2URL is an alias for ErrResolveURL kept for backward compatibility.
	ErrElem2URL = ErrResolveURL

	// ErrFormatNotExist is an alias for ErrFormatNotFound kept for backward compatibility.
	ErrFormatNotExist = ErrFormatNotFound

	// ErrFindElem is an alias for ErrElementNotFound kept for backward compatibility.
	ErrFindElem = ErrElementNotFound
)

var supportedHashes = []string{"md5"}

// Catalog represents a collection of elements, formats, and a base URL.
// It is fully thread-safe for concurrent read and write operations.
// Field alignment optimized.
type Catalog struct {
	// Formats maps format identifiers to their URL patterns and file templates.
	Formats FormatDefinitions `json:"formats" yaml:"formats"`

	// Elements maps element IDs to geographic element metadata.
	Elements map[string]Element `json:"elements" yaml:"elements"`

	// BaseURL is the default root download URL for the provider.
	BaseURL string `json:"baseURL" yaml:"baseURL"` //nolint:tagliatelle // external yaml requirement

	mu sync.RWMutex
}

// New creates an empty, initialized Catalog.
func New() *Catalog {
	return &Catalog{
		Formats:  make(FormatDefinitions),
		Elements: make(map[string]Element),
	}
}

// LoadFile reads and parses a YAML catalog file from disk.
func LoadFile(filePath string) (*Catalog, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("invalid path %s: %w", filePath, err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s: %w", absPath, err)
	}

	cat := New()
	if err := yaml.Unmarshal(data, cat); err != nil {
		return nil, fmt.Errorf("cannot unmarshal yaml from %s: %w", absPath, err)
	}

	if cat.Elements == nil {
		cat.Elements = make(map[string]Element)
	}

	if cat.Formats == nil {
		cat.Formats = make(FormatDefinitions)
	}

	return cat, nil
}

// SaveFile marshals the catalog to YAML and writes it atomically to a file.
func (c *Catalog) SaveFile(filePath string) error {
	c.mu.RLock()
	data, err := yaml.Marshal(c)
	c.mu.RUnlock()

	if err != nil {
		return fmt.Errorf("cannot marshal catalog to yaml: %w", err)
	}

	dir := filepath.Dir(filePath)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, defaultDirPerm); err != nil {
			return fmt.Errorf("cannot create directory %s: %w", dir, err)
		}
	}

	tmpPath := filePath + ".tmp"
	if err := os.WriteFile(tmpPath, data, defaultFilePerm); err != nil {
		return fmt.Errorf("cannot write catalog to %s: %w", tmpPath, err)
	}

	if err := os.Rename(tmpPath, filePath); err != nil {
		_ = os.Remove(tmpPath)

		return fmt.Errorf("cannot rename %s to %s: %w", tmpPath, filePath, err)
	}

	return nil
}

// Save writes the catalog YAML to any io.Writer.
func (c *Catalog) Save(writer io.Writer) error {
	c.mu.RLock()
	data, err := yaml.Marshal(c)
	c.mu.RUnlock()

	if err != nil {
		return fmt.Errorf("cannot marshal catalog: %w", err)
	}

	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("cannot write catalog: %w", err)
	}

	return nil
}

// Len returns the number of elements in the catalog thread-safely.
func (c *Catalog) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.Elements)
}

// GetFormat retrieves a copy of the format definition by ID thread-safely.
func (c *Catalog) GetFormat(formatID string) (Format, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	format, exists := c.Formats[formatID]

	return format, exists
}

// Exists returns true if the element ID is present in the catalog.
func (c *Catalog) Exists(elementID string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.getElementLocked(elementID)

	return exists
}

// Exist is an alias for Exists kept for backward compatibility.
func (c *Catalog) Exist(elementID string) bool {
	return c.Exists(elementID)
}

// Get retrieves a copy of an element by ID.
func (c *Catalog) Get(elementID string) (Element, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	elem, exists := c.getElementLocked(elementID)
	if !exists {
		return Element{}, false
	}

	elem.Formats = slices.Clone(elem.Formats)

	return elem, true
}

// Find looks up an element pointer by ID or returns ErrElementNotFound.
func (c *Catalog) Find(elementID string) (*Element, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if elem, exists := c.getElementLocked(elementID); exists {
		elemCopy := elem
		elemCopy.Formats = slices.Clone(elem.Formats)

		return &elemCopy, nil
	}

	return nil, fmt.Errorf("%w: %s is not in catalog", ErrElementNotFound, elementID)
}

func (c *Catalog) getElementLocked(elementID string) (Element, bool) {
	if elem, exists := c.Elements[elementID]; exists {
		return elem, true
	}

	if altID := strings.ReplaceAll(elementID, "-", "_"); altID != elementID {
		if elem, exists := c.Elements[altID]; exists {
			return elem, true
		}
	}

	if altID := strings.ReplaceAll(elementID, "_", "-"); altID != elementID {
		if elem, exists := c.Elements[altID]; exists {
			return elem, true
		}
	}

	if elem, ok := c.resolveDateElement(elementID); ok {
		return elem, true
	}

	return Element{}, false
}

func (c *Catalog) resolveDateElement(dateStr string) (Element, bool) {
	if _, err := time.Parse("2006-01-02", dateStr); err != nil {
		return Element{}, false
	}

	if baseElem, ok := c.Elements["czech_republic"]; ok {
		return Element{
			ID:      dateStr,
			Name:    baseElem.Name + " " + dateStr,
			File:    "czech_republic-" + dateStr,
			Formats: Formats{FormatOsmPbf, FormatOsmBz2},
		}, true
	}

	if latestElem, ok := c.Elements["latest"]; ok {
		prefix := "czech_republic"
		if idx := strings.Index(latestElem.File, "-"); idx != -1 {
			prefix = latestElem.File[:idx]
		}

		return Element{
			ID:      dateStr,
			Name:    latestElem.Name + " " + dateStr,
			File:    prefix + "-" + dateStr,
			Formats: Formats{FormatOsmPbf, FormatOsmBz2},
		}, true
	}

	return Element{}, false
}

// AddElement inserts or replaces an element in the catalog.
func (c *Catalog) AddElement(elem *Element) {
	if elem == nil || elem.ID == "" {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Elements == nil {
		c.Elements = make(map[string]Element)
	}

	c.Elements[elem.ID] = *elem
}

// MergeElement adds a new element or merges formats if it already exists.
func (c *Catalog) MergeElement(elem *Element) error {
	if elem == nil || elem.ID == "" {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Elements == nil {
		c.Elements = make(map[string]Element)
	}

	existing, exists := c.Elements[elem.ID]
	if !exists {
		c.Elements[elem.ID] = *elem

		return nil
	}

	if err := validateParentMerge(&existing, elem); err != nil {
		return err
	}

	applyElementMerge(&existing, elem)
	c.Elements[elem.ID] = existing

	return nil
}

func validateParentMerge(existing, incoming *Element) error {
	if existing.Parent != incoming.Parent && existing.Parent != "" && incoming.Parent != "" {
		return fmt.Errorf("%w: %s has parent %s, incoming has %s", ErrParentMismatch, incoming.ID, existing.Parent, incoming.Parent)
	}

	return nil
}

func applyElementMerge(target, source *Element) {
	for _, format := range source.Formats {
		target.AddFormat(format)
	}

	target.Parent = cmp.Or(source.Parent, target.Parent)
	target.Name = cmp.Or(source.Name, target.Name)
	target.File = cmp.Or(source.File, target.File)

	if source.Meta {
		target.Meta = source.Meta
	}
}

// AddExtension adds a format to an existing element if present.
func (c *Catalog) AddExtension(elementID, formatID string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, exists := c.Elements[elementID]; exists {
		elem.AddFormat(formatID)
		c.Elements[elementID] = elem
	}
}

// SortedKeys returns the lexicographically sorted list of all element IDs with single pre-allocation.
func (c *Catalog) SortedKeys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.Elements) == 0 {
		return nil
	}

	return slices.Sorted(maps.Keys(c.Elements))
}

// All returns a sequence iterator over all elements (Go 1.23+).
// Note: The catalog read lock is held for the duration of the iteration.
// Callers must not invoke mutating methods on this Catalog within the iteration loop.
func (c *Catalog) All() iter.Seq2[string, Element] {
	return func(yield func(string, Element) bool) {
		c.mu.RLock()
		defer c.mu.RUnlock()

		for k, v := range c.Elements {
			if !yield(k, v) {
				return
			}
		}
	}
}

// ResolveURL constructs the absolute download URL for an element and format.
func (c *Catalog) ResolveURL(elem *Element, formatID string) (string, error) {
	if elem == nil {
		return "", ErrNilElement
	}

	if !elem.ContainsFormat(formatID) {
		return "", fmt.Errorf("%w: %s on element %s", ErrFormatNotFound, formatID, elem.ID)
	}

	c.mu.RLock()
	format, formatExists := c.Formats[formatID]
	catalogBaseURL := c.BaseURL
	c.mu.RUnlock()

	if !formatExists {
		return "", fmt.Errorf("%w: %s definition missing from catalog", ErrFormatNotFound, formatID)
	}

	baseURL := cmp.Or(format.BaseURL, catalogBaseURL)

	preURL, err := c.ResolvePreURL(elem, baseURL, format.BasePath)
	if err != nil {
		return "", err
	}

	return preURL + format.Loc, nil
}

func (c *Catalog) collectHierarchySegments(startID string) ([]string, error) {
	var segments [maxHierarchyDepth]string

	count := 0
	currID := startID

	for count < maxHierarchyDepth {
		currentElem, exists := c.getElementLocked(currID)
		if !exists {
			return nil, fmt.Errorf("%w: %s is not in catalog", ErrElementNotFound, currID)
		}

		segments[count] = currentElem.Filename()
		count++

		if !currentElem.HasParent() {
			break
		}

		currID = currentElem.Parent
	}

	if count >= maxHierarchyDepth {
		return nil, fmt.Errorf("%w for element %s", ErrMaxHierarchyDepth, startID)
	}

	return segments[:count], nil
}

// ResolvePreURL iteratively builds the URL path prefix with cycle protection and minimal allocations.
func (c *Catalog) ResolvePreURL(elem *Element, baseURL ...string) (string, error) {
	if elem == nil {
		return "", ErrNilElement
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	segments, err := c.collectHierarchySegments(elem.ID)
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.Grow(defaultBuilderSize)

	builder.WriteString(buildURLPrefix(baseURL, c.BaseURL))

	for i, seg := range slices.Backward(segments) {
		builder.WriteString(seg)

		if i > 0 {
			builder.WriteByte('/')
		}
	}

	return builder.String(), nil
}

func buildURLPrefix(baseURL []string, defaultBaseURL string) string {
	var prefix string

	switch len(baseURL) {
	case 1:
		prefix = defaultBaseURL + "/" + strings.Join(baseURL, "/")
	case twoBaseURLParts:
		prefix = strings.Join(baseURL, "/")
	default:
		prefix = defaultBaseURL
	}

	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	return prefix
}

// IsHashable checks if a given format has a corresponding hash definition (e.g. .md5).
func (c *Catalog) IsHashable(formatID string) (ok bool, hashExt, hashType string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if _, exists := c.Formats[formatID]; exists {
		for _, h := range supportedHashes {
			candidate := formatID + "." + h
			if _, exists := c.Formats[candidate]; exists {
				return true, candidate, h
			}
		}
	}

	return false, "", ""
}
