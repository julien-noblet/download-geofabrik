// 2015-2018 copyright Julien Noblet

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package config

import (
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
	ErrFindElem   = "can't find %s"
	ErrLoadConfig = "can't load config"
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
		return nil, fmt.Errorf("failed to Marshal : %w", err)
	}

	return yml, nil
}

func (config *Config) MergeElement(elementPtr *element.Element) error {
	config.ElementsMutex.RLock()
	newElement, ok := config.Elements[elementPtr.ID]
	config.ElementsMutex.RUnlock()

	if ok { //nolint:nestif // TODO : Refactor?
		if newElement.Parent != elementPtr.Parent {
			return fmt.Errorf("can't merge : Parent mismatch %s != %s (%s)", newElement.Parent, elementPtr.Parent, elementPtr.ID)
		}

		config.ElementsMutex.Lock()
		for _, f := range elementPtr.Formats {
			if !newElement.Formats.Contains(f) {
				newElement.Formats = append(newElement.Formats, f)
			}
		}
		config.ElementsMutex.Unlock()

		if len(newElement.Formats) == 0 {
			newElement.Meta = true
		} else {
			newElement.Meta = false
		}

		config.ElementsMutex.Lock()
		config.Elements[elementPtr.ID] = newElement
		config.ElementsMutex.Unlock()
	} else {
		config.ElementsMutex.Lock()
		config.Elements[elementPtr.ID] = *elementPtr
		config.ElementsMutex.Unlock()
	}

	return nil
}

// Exist check if id is in e.Elements.
func (config *Config) Exist(id string) bool {
	config.ElementsMutex.RLock()
	result := reflect.DeepEqual(config.Elements[id], element.Element{}) //nolint:exhaustruct,lll // TODO : Move config.Elements map[string]Element to maps[string]*Element
	config.ElementsMutex.RUnlock()

	return !result
}

// AddExtension Add an extension to an element.
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

// GetElement get an element by id or error if not found.
func (config *Config) GetElement(id string) (*element.Element, error) { //nolint:varnamelen // id is ok
	if config.Exist(id) {
		config.ElementsMutex.RLock()
		r := config.Elements[id]
		config.ElementsMutex.RUnlock()

		return &r, nil
	}

	return nil, fmt.Errorf("element %s not found", id)
}

func FindElem(config *Config, e string) (*element.Element, error) {
	res := config.Elements[e]
	if res.ID == "" || res.ID != e {
		return nil, fmt.Errorf("%s is not in config. Please use \"list\" command", e)
	}

	return &res, nil
}

func GetFile(myElement *element.Element) string {
	if myElement.File != "" {
		return myElement.File
	}

	return myElement.ID
}

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

		res += "/"
		res += GetFile(myElement)

		return res, nil
	}

	switch len(baseURL) {
	case 1:
		return config.BaseURL + "/" + strings.Join(baseURL, "/") + GetFile(myElement), nil
	case 2: //nolint:gomnd // return without c.BaseURL
		return strings.Join(baseURL, "/") + GetFile(myElement), nil
	default: // len(b)==0 or >2
		return config.BaseURL + "/" + GetFile(myElement), nil
	}
}

func Elem2URL(config *Config, elementPtr *element.Element, ext string) (string, error) {
	var (
		res string
		err error
	)

	if !elementPtr.Formats.Contains(ext) {
		return "", fmt.Errorf("error!!! %s format not exist", ext)
	}

	format := config.Formats[ext]
	if format.BasePath != "" { //nolint:nestif // TODO : Refactor?
		if format.BaseURL != "" {
			res, err = Elem2preURL(config, elementPtr, format.BaseURL, format.BasePath)
		} else {
			res, err = Elem2preURL(config, elementPtr, format.BasePath)
		}
	} else {
		if format.BaseURL != "" {
			res, err = Elem2preURL(config, elementPtr, format.BaseURL, "")
		} else {
			res, err = Elem2preURL(config, elementPtr)
		}
	}
	// TODO check if valid URL
	if err != nil {
		return "", err
	}

	res += format.Loc

	return res, nil
}

// LoadConfig loading configFile and send *Config.
// If there is an error, return it also.
func LoadConfig(configFile string) (*Config, error) {
	filename, _ := filepath.Abs(configFile) // Get absolute path

	fileContent, err := os.ReadFile(filename) // Open file as string
	if err != nil {
		return nil, fmt.Errorf("can't open %s : %w", filename, err)
	}
	// Create a Config ptr
	myConfigPtr := &Config{
		Formats:       formats.FormatDefinitions{},
		Elements:      element.Slice{},
		ElementsMutex: &sync.RWMutex{},
		BaseURL:       "",
	}
	// Charging fileContent into myConfigPtr
	if err := yaml.Unmarshal(fileContent, myConfigPtr); err != nil {
		return nil, fmt.Errorf("can't unmarshal %s : %w", filename, err)
	}
	// Everything is OK, returning myConfigPtr
	return myConfigPtr, nil
}

func IsHashable(config *Config, format string) (hashable bool, hashFilename, hashType string) { //nolint:nonamedreturns // better for documentation
	hashs := []string{"md5"} // had to be globalized?

	if _, ok := config.Formats[format]; ok {
		for _, h := range hashs {
			hash := format + "." + h
			if _, ok := config.Formats[hash]; ok {
				return true, hash, h
			}
		}
	}

	return false, "", ""
}
