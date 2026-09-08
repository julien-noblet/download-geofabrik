package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	"gopkg.in/yaml.v3"
)

const (
	DefaultConfigFile  = "geofabrik.yml"
	DefaultService     = "geofabrik"
	maxHierarchyDepth  = 30
	defaultBuilderSize = 64
	twoBaseURLParts    = 2
)

var (
	ErrElem2URL          = errors.New("can't find url")
	ErrLoadConfig        = errors.New("can't load config")
	ErrFindElem          = errors.New("element not found")
	ErrParentMismatch    = errors.New("can't merge")
	ErrFormatNotExist    = errors.New("format not exist")
	ErrMaxHierarchyDepth = errors.New("maximum hierarchy depth exceeded")

	hashes = []string{"md5"}
)

// Config structure handles all elements and formats from the YAML database.
type Config struct {
	Formats       formats.FormatDefinitions `yaml:"formats"`
	Elements      element.MapElement        `yaml:"elements"`
	ElementsMutex *sync.RWMutex             `yaml:"-"`       // unexported
	BaseURL       string                    `yaml:"baseURL"` //nolint:tagliatelle // external yaml requirement
}

// Options holds runtime configuration (flags).
// Field alignment optimized.
type Options struct {
	FormatFlags     map[string]bool
	ConfigFile      string
	Service         string
	OutputDirectory string
	Check           bool
	Verbose         bool
	Quiet           bool
	NoDownload      bool
	Progress        bool
}

// Generate Yaml config.
func (config *Config) Generate() ([]byte, error) {
	yml, err := yaml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to Marshal: %w", err)
	}

	return yml, nil
}

// MergeElement merges a new element into the config safely using a write lock.
func (config *Config) MergeElement(elementPtr *element.Element) error {
	config.ElementsMutex.Lock()
	defer config.ElementsMutex.Unlock()

	newElement, ok := config.Elements[elementPtr.ID]
	if ok {
		if newElement.Parent != elementPtr.Parent {
			return fmt.Errorf("%w: Parent mismatch %s != %s (%s)", ErrParentMismatch, newElement.Parent, elementPtr.Parent, elementPtr.ID)
		}

		for _, f := range elementPtr.Formats {
			if !newElement.Formats.Contains(f) {
				newElement.Formats = append(newElement.Formats, f)
			}
		}

		newElement.Meta = len(newElement.Formats) == 0
		config.Elements[elementPtr.ID] = newElement
	} else {
		config.Elements[elementPtr.ID] = *elementPtr
	}

	return nil
}

// Exist checks if an element with the given ID exists in the config.
func (config *Config) Exist(elementID string) bool {
	config.ElementsMutex.RLock()
	defer config.ElementsMutex.RUnlock()

	if _, ok := findInElements(config.Elements, elementID); ok {
		return true
	}

	if _, ok := resolveDateElement(config, elementID); ok {
		return true
	}

	return false
}

func findInElements(elements element.MapElement, elementID string) (*element.Element, bool) {
	if res, ok := elements[elementID]; ok && res.ID != "" {
		elemCopy := res

		return &elemCopy, true
	}

	if altID := strings.ReplaceAll(elementID, "-", "_"); altID != elementID {
		if res, ok := elements[altID]; ok && res.ID != "" {
			elemCopy := res

			return &elemCopy, true
		}
	}

	if altID := strings.ReplaceAll(elementID, "_", "-"); altID != elementID {
		if res, ok := elements[altID]; ok && res.ID != "" {
			elemCopy := res

			return &elemCopy, true
		}
	}

	return nil, false
}

// AddExtension adds an extension to an element, creating it if not already present.
func (config *Config) AddExtension(elementID, format string) {
	config.ElementsMutex.Lock()
	defer config.ElementsMutex.Unlock()

	elem, ok := config.Elements[elementID]
	if !ok {
		elem = element.Element{
			ID:      elementID,
			Formats: element.Formats{},
		}
	}

	if !elem.Formats.Contains(format) {
		slog.Info("Add extension to element", "format", format, "id", elem.ID)

		elem.Formats = append(elem.Formats, format)
		elem.Meta = len(elem.Formats) == 0
		config.Elements[elementID] = elem
	}
}

// GetElement gets an element by ID or returns an error if not found.
func (config *Config) GetElement(elementID string) (*element.Element, error) {
	config.ElementsMutex.RLock()
	defer config.ElementsMutex.RUnlock()

	return FindElem(config, elementID)
}

// FindElem finds an element in the config by ID, with fallback for normalized hyphen/underscore variations
// and dynamic date-based elements (YYYY-MM-DD).
func FindElem(config *Config, elementID string) (*element.Element, error) {
	if config == nil {
		return nil, fmt.Errorf("%w: %s is not in config. Please use \"list\" command", ErrFindElem, elementID)
	}

	if elem, ok := findInElements(config.Elements, elementID); ok {
		return elem, nil
	}

	if dateElem, ok := resolveDateElement(config, elementID); ok {
		return dateElem, nil
	}

	return nil, fmt.Errorf("%w: %s is not in config. Please use \"list\" command", ErrFindElem, elementID)
}

func resolveDateElement(config *Config, dateStr string) (*element.Element, bool) {
	if _, err := time.Parse("2006-01-02", dateStr); err != nil {
		return nil, false
	}

	if baseElem, ok := config.Elements["czech_republic"]; ok {
		return &element.Element{
			ID:      dateStr,
			Name:    baseElem.Name + " " + dateStr,
			File:    "czech_republic-" + dateStr,
			Formats: element.Formats{formats.FormatOsmPbf, formats.FormatOsmBz2},
		}, true
	}

	if latestElem, ok := config.Elements["latest"]; ok {
		prefix := "czech_republic"
		if idx := strings.Index(latestElem.File, "-"); idx != -1 {
			prefix = latestElem.File[:idx]
		}

		return &element.Element{
			ID:      dateStr,
			Name:    latestElem.Name + " " + dateStr,
			File:    prefix + "-" + dateStr,
			Formats: element.Formats{formats.FormatOsmPbf, formats.FormatOsmBz2},
		}, true
	}

	return nil, false
}

// GetFile gets the file name of an element.
func GetFile(myElement *element.Element) string {
	if myElement.File != "" {
		return myElement.File
	}

	return myElement.ID
}

func buildPrefix(baseURL []string, defaultBaseURL string) string {
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

func collectElementSegments(config *Config, startID string) ([]string, error) {
	var segments [maxHierarchyDepth]string

	count := 0
	currID := startID

	for count < maxHierarchyDepth {
		elemPtr, err := FindElem(config, currID)
		if err != nil {
			return nil, err
		}

		file := elemPtr.File
		if file == "" {
			file = elemPtr.ID
		}

		segments[count] = file
		count++

		if elemPtr.Parent == "" {
			break
		}

		currID = elemPtr.Parent
	}

	if count >= maxHierarchyDepth {
		return nil, fmt.Errorf("%w for element %s (possible cycle in config)", ErrMaxHierarchyDepth, startID)
	}

	return segments[:count], nil
}

// Elem2preURL generates a pre-URL for an element iteratively with single-allocation buffer.
func Elem2preURL(config *Config, elementPtr *element.Element, baseURL ...string) (string, error) {
	if config == nil || elementPtr == nil {
		return "", fmt.Errorf("%w: invalid nil argument", ErrFindElem)
	}

	segments, err := collectElementSegments(config, elementPtr.ID)
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.Grow(defaultBuilderSize)

	builder.WriteString(buildPrefix(baseURL, config.BaseURL))

	for i := len(segments) - 1; i >= 0; i-- {
		builder.WriteString(segments[i])

		if i > 0 {
			builder.WriteByte('/')
		}
	}

	return builder.String(), nil
}

// Elem2URL generates a URL for an element with the given extension.
func Elem2URL(config *Config, elementPtr *element.Element, ext string) (string, error) {
	if !elementPtr.Formats.Contains(ext) {
		return "", fmt.Errorf("%w: %s", ErrFormatNotExist, ext)
	}

	format, ok := config.Formats[ext]
	if !ok {
		return "", fmt.Errorf("%w: %s", ErrFormatNotExist, ext)
	}

	baseURL, basePath := format.BaseURL, format.BasePath
	if baseURL == "" {
		baseURL = config.BaseURL
	}

	res, err := Elem2preURL(config, elementPtr, baseURL, basePath)
	if err != nil {
		return "", err
	}

	return res + format.Loc, nil
}

// LoadConfig loads the configuration from the specified file.
func LoadConfig(configFile string) (*Config, error) {
	filename, _ := filepath.Abs(configFile)

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't open %s: %w", filename, err)
	}

	myConfigPtr := &Config{
		Formats:       formats.FormatDefinitions{},
		Elements:      element.MapElement{},
		ElementsMutex: &sync.RWMutex{},
		BaseURL:       "",
	}

	if err := yaml.Unmarshal(fileContent, myConfigPtr); err != nil {
		return nil, fmt.Errorf("can't unmarshal %s: %w", filename, err)
	}

	return myConfigPtr, nil
}

// IsHashable checks if a format is hashable.
func IsHashable(config *Config, format string) (isHashable bool, hash, extension string) {
	if _, ok := config.Formats[format]; ok {
		for _, h := range hashes {
			hashKey := format + "." + h
			if _, ok := config.Formats[hashKey]; ok {
				return true, hashKey, h
			}
		}
	}

	return false, "", ""
}
