package catalog

import (
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

	"gopkg.in/yaml.v3"
)

const (
	maxHierarchyDepth = 30
	defaultDirPerm    = 0o750
	defaultFilePerm   = 0o600
)

var (
	ErrNilElement        = errors.New("nil element")
	ErrElementNotFound   = errors.New("element not found")
	ErrFormatNotFound    = errors.New("format not found")
	ErrParentMismatch    = errors.New("cannot merge element with conflicting parent")
	ErrMaxHierarchyDepth = errors.New("maximum hierarchy depth exceeded (possible cycle in catalog)")
)

var supportedHashes = []string{"md5"}

// Catalog represents a collection of elements, formats, and a base URL.
// It is fully thread-safe for concurrent read and write operations.
// Field alignment optimized.
type Catalog struct {
	Formats  FormatDefinitions  `json:"formats"  yaml:"formats"`
	Elements map[string]Element `json:"elements" yaml:"elements"`
	BaseURL  string             `json:"baseURL"  yaml:"baseURL"` //nolint:tagliatelle // external yaml requirement
	mu       sync.RWMutex
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
		return nil, fmt.Errorf("cannot unmarshal YAML from %s: %w", absPath, err)
	}

	if cat.Elements == nil {
		cat.Elements = make(map[string]Element)
	}

	if cat.Formats == nil {
		cat.Formats = make(FormatDefinitions)
	}

	return cat, nil
}

// SaveFile marshals the catalog to YAML and writes it to a file.
func (c *Catalog) SaveFile(filePath string) error {
	c.mu.RLock()
	data, err := yaml.Marshal(c)
	c.mu.RUnlock()

	if err != nil {
		return fmt.Errorf("cannot marshal catalog to YAML: %w", err)
	}

	dir := filepath.Dir(filePath)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, defaultDirPerm); err != nil {
			return fmt.Errorf("cannot create directory %s: %w", dir, err)
		}
	}

	if err := os.WriteFile(filePath, data, defaultFilePerm); err != nil {
		return fmt.Errorf("cannot write catalog to %s: %w", filePath, err)
	}

	return nil
}

// Save writes the catalog YAML to any io.Writer.
func (c *Catalog) Save(w io.Writer) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if err := yaml.NewEncoder(w).Encode(c); err != nil {
		return fmt.Errorf("cannot encode catalog: %w", err)
	}

	return nil
}

// Exist returns true if the element ID is present in the catalog.
func (c *Catalog) Exist(elementID string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.Elements[elementID]

	return exists
}

// Get retrieves a copy of an element by ID.
func (c *Catalog) Get(elementID string) (Element, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	elem, exists := c.Elements[elementID]

	return elem, exists
}

// Find looks up an element pointer by ID or returns ErrElementNotFound.
func (c *Catalog) Find(elementID string) (*Element, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	elem, exists := c.Elements[elementID]
	if !exists {
		return nil, fmt.Errorf("%w: %s is not in catalog", ErrElementNotFound, elementID)
	}

	elemCopy := elem

	return &elemCopy, nil
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

	if source.Parent != "" {
		target.Parent = source.Parent
	}

	if source.Name != "" {
		target.Name = source.Name
	}

	if source.File != "" {
		target.File = source.File
	}

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

// SortedKeys returns the lexicographically sorted list of all element IDs.
func (c *Catalog) SortedKeys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return slices.Sorted(maps.Keys(c.Elements))
}

// All returns a sequence iterator over all elements (Go 1.23+).
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

	baseURL := format.BaseURL
	if baseURL == "" {
		baseURL = catalogBaseURL
	}

	preURL, err := c.ResolvePreURL(elem, baseURL, format.BasePath)
	if err != nil {
		return "", err
	}

	return preURL + format.Loc, nil
}

// ResolvePreURL recursively builds the URL path prefix with cycle protection.
func (c *Catalog) ResolvePreURL(elem *Element, baseURL ...string) (string, error) {
	return c.resolvePreURLWithDepth(elem, 0, baseURL...)
}

func (c *Catalog) resolvePreURLWithDepth(elem *Element, depth int, baseURL ...string) (string, error) {
	if depth > maxHierarchyDepth {
		return "", fmt.Errorf("%w for element %s", ErrMaxHierarchyDepth, elem.ID)
	}

	currentElem, err := c.Find(elem.ID)
	if err != nil {
		return "", err
	}

	if currentElem.HasParent() {
		parentElem, err := c.Find(currentElem.Parent)
		if err != nil {
			return "", err
		}

		parentPreURL, err := c.resolvePreURLWithDepth(parentElem, depth+1, baseURL...)
		if err != nil {
			return "", err
		}

		return parentPreURL + "/" + currentElem.Filename(), nil
	}

	prefix := buildURLPrefix(baseURL, c.BaseURL)

	return prefix + currentElem.Filename(), nil
}

func buildURLPrefix(baseURL []string, defaultBaseURL string) string {
	var prefix string

	switch len(baseURL) {
	case 1:
		prefix = defaultBaseURL + "/" + strings.Join(baseURL, "/")
	case 2: //nolint:mnd // Exactly 2 components (baseURL, basePath)
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
