package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	"gopkg.in/yaml.v3"
)

const (
	DefaultConfigFile = "geofabrik.yml"
	DefaultService    = "geofabrik"
)

var (
	ErrElem2URL       = errors.New("can't find url")
	ErrLoadConfig     = errors.New("can't load config")
	ErrFindElem       = errors.New("element not found")
	ErrParentMismatch = errors.New("can't merge")
	ErrFormatNotExist = errors.New("format not exist")

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

// MergeElement merges a new element into the config.
func (config *Config) MergeElement(elementPtr *element.Element) error {
	config.ElementsMutex.RLock()
	newElement, ok := config.Elements[elementPtr.ID]
	config.ElementsMutex.RUnlock()

	if ok {
		if newElement.Parent != elementPtr.Parent {
			return fmt.Errorf("%w: Parent mismatch %s != %s (%s)", ErrParentMismatch, newElement.Parent, elementPtr.Parent, elementPtr.ID)
		}

		config.ElementsMutex.Lock()
		defer config.ElementsMutex.Unlock()

		for _, f := range elementPtr.Formats {
			if !newElement.Formats.Contains(f) {
				newElement.Formats = append(newElement.Formats, f)
			}
		}

		newElement.Meta = len(newElement.Formats) == 0
		config.Elements[elementPtr.ID] = newElement
	} else {
		config.ElementsMutex.Lock()
		defer config.ElementsMutex.Unlock()

		config.Elements[elementPtr.ID] = *elementPtr
	}

	return nil
}

// Exist checks if an element with the given ID exists in the config.
func (config *Config) Exist(elementID string) bool {
	config.ElementsMutex.RLock()
	defer config.ElementsMutex.RUnlock()

	result := reflect.DeepEqual(config.Elements[elementID], element.Element{})

	return !result
}

// AddExtension adds an extension to an element.
func (config *Config) AddExtension(elementID, format string) {
	config.ElementsMutex.RLock()
	elem := config.Elements[elementID]
	config.ElementsMutex.RUnlock()

	if !elem.Formats.Contains(format) {
		slog.Info("Add extension to element", "format", format, "id", elem.ID)

		config.ElementsMutex.Lock()

		elem.Formats = append(elem.Formats, format)

		config.ElementsMutex.Unlock()

		if err := config.MergeElement(&elem); err != nil {
			slog.Error("can't merge element", "error", err, "name", elem.Name)
			os.Exit(1) // Or handle better
		}
	}
}

// GetElement gets an element by ID or returns an error if not found.
func (config *Config) GetElement(elementID string) (*element.Element, error) {
	if config.Exist(elementID) {
		config.ElementsMutex.RLock()
		r := config.Elements[elementID]
		config.ElementsMutex.RUnlock()

		return &r, nil
	}

	return nil, fmt.Errorf("%w: %s", ErrFindElem, elementID)
}

// FindElem finds an element in the config by ID.
func FindElem(config *Config, e string) (*element.Element, error) {
	res := config.Elements[e]
	if res.ID == "" || res.ID != e {
		return nil, fmt.Errorf("%w: %s is not in config. Please use \"list\" command", ErrFindElem, e)
	}

	return &res, nil
}

// GetFile gets the file name of an element.
func GetFile(myElement *element.Element) string {
	if myElement.File != "" {
		return myElement.File
	}

	return myElement.ID
}

// Elem2preURL generates a pre-URL for an element.
func Elem2preURL(config *Config, elementPtr *element.Element, baseURL ...string) (string, error) {
	myElement, err := FindElem(config, elementPtr.ID)
	if err != nil {
		return "", err
	}

	if myElement.HasParent() {
		parent, err := FindElem(config, myElement.Parent)
		if err != nil {
			return "", err
		}

		res, err := Elem2preURL(config, parent, baseURL...)
		if err != nil {
			return "", err
		}

		res += "/" + GetFile(myElement)

		return res, nil
	}

	switch len(baseURL) {
	case 1:
		prefix := config.BaseURL + "/" + strings.Join(baseURL, "/")
		if !strings.HasSuffix(prefix, "/") {
			prefix += "/"
		}

		return prefix + GetFile(myElement), nil

	case 2: //nolint:mnd // This case handles exactly 2 base URL components
		prefix := strings.Join(baseURL, "/")
		if !strings.HasSuffix(prefix, "/") {
			prefix += "/"
		}

		return prefix + GetFile(myElement), nil

	default:
		return config.BaseURL + "/" + GetFile(myElement), nil
	}
}

// Elem2URL generates a URL for an element with the given extension.
func Elem2URL(config *Config, elementPtr *element.Element, ext string) (string, error) {
	if !elementPtr.Formats.Contains(ext) {
		return "", fmt.Errorf("%w: %s", ErrFormatNotExist, ext)
	}

	format := config.Formats[ext]
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
			hash := format + "." + h
			if _, ok := config.Formats[hash]; ok {
				return true, hash, h
			}
		}
	}

	return false, "", ""
}
