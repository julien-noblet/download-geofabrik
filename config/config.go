package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/apex/log"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	yaml "gopkg.in/yaml.v3"
)

const (
	ErrElem2URL   = "can't find url"
	ErrLoadConfig = "can't load config"
)

var (
	ErrFindElem       = errors.New("element not found")
	ErrParentMismatch = errors.New("can't merge")
	ErrFormatNotExist = errors.New("format not exist")

	hashes = []string{"md5"}
)

// Config structure handle all elements.
// It also contain the BaseURL and Formats...
type Config struct {
	Formats       formats.FormatDefinitions `yaml:"formats"`
	Elements      element.Slice             `yaml:"elements"`
	ElementsMutex *sync.RWMutex             `yaml:"-"`       // unexported
	BaseURL       string                    `yaml:"baseURL"` //nolint:tagliatelle // revive prefer this form
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
func (config *Config) Exist(id string) bool {
	config.ElementsMutex.RLock()
	defer config.ElementsMutex.RUnlock()
	result := reflect.DeepEqual(config.Elements[id], element.Element{})

	return !result
}

// AddExtension adds an extension to an element.
func (config *Config) AddExtension(id, format string) {
	config.ElementsMutex.RLock()
	elem := config.Elements[id]
	config.ElementsMutex.RUnlock()

	if !elem.Formats.Contains(format) {
		log.Infof("Add %s to %s", format, elem.ID)

		config.ElementsMutex.Lock()
		elem.Formats = append(elem.Formats, format)
		config.ElementsMutex.Unlock()

		if err := config.MergeElement(&elem); err != nil {
			log.WithError(err).Fatalf("can't merge %s", elem.Name)
		}
	}
}

// GetElement gets an element by ID or returns an error if not found.
func (config *Config) GetElement(id string) (*element.Element, error) { //nolint:varnamelen // id is ok
	if config.Exist(id) {
		config.ElementsMutex.RLock()
		r := config.Elements[id]
		config.ElementsMutex.RUnlock()

		return &r, nil
	}

	return nil, fmt.Errorf("%w: %s", ErrFindElem, id)
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
		return config.BaseURL + "/" + strings.Join(baseURL, "/") + GetFile(myElement), nil
	case 2: //nolint:mnd // return without c.BaseURL
		return strings.Join(baseURL, "/") + GetFile(myElement), nil
	default: // len(b)==0 or >2
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
	if baseURL == "" && basePath == "" {
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
		Elements:      element.Slice{},
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
