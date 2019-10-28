// 2015-2018 copyright Julien Noblet

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// Config structure handle all elements.
// It also contain the BaseURL and Formats...
type Config struct {
	BaseURL       string                    `yaml:"baseURL"`
	Formats       formats.FormatDefinitions `yaml:"formats"`
	Elements      element.Slice             `yaml:"elements"`
	ElementsMutex *sync.RWMutex             `yaml:"-"` // unexported
}

// Generate Yaml config
func (c *Config) Generate() ([]byte, error) {
	return yaml.Marshal(c)
}

func (c *Config) MergeElement(el *element.Element) error {
	c.ElementsMutex.RLock()
	cE, ok := c.Elements[el.ID]
	c.ElementsMutex.RUnlock()

	if ok {
		if cE.Parent != el.Parent {
			return fmt.Errorf("cant merge : Parent mismatch")
		}

		c.ElementsMutex.Lock()
		for _, f := range el.Formats {
			if !cE.Formats.Contains(f) {
				cE.Formats = append(cE.Formats, f)
			}
		}
		c.ElementsMutex.Unlock()

		if len(cE.Formats) == 0 {
			cE.Meta = true
		} else {
			cE.Meta = false
		}

		c.ElementsMutex.Lock()
		c.Elements[el.ID] = cE
		c.ElementsMutex.Unlock()
	} else {
		c.ElementsMutex.Lock()
		c.Elements[el.ID] = *el
		c.ElementsMutex.Unlock()
	}

	return nil
}

// Exist check if id is in e.Elements
func (c *Config) Exist(id string) bool {
	c.ElementsMutex.RLock()
	r := reflect.DeepEqual(c.Elements[id], element.Element{})
	c.ElementsMutex.RUnlock()

	return !r
}

// AddExtension Add an extension to an element
func (c *Config) AddExtension(id, format string) {
	c.ElementsMutex.RLock()
	elem := c.Elements[id]
	c.ElementsMutex.RUnlock()

	if !elem.Formats.Contains(format) {
		if viper.GetBool("verbose") && !viper.GetBool("quiet") && !viper.GetBool("progress") {
			log.Println("Add", format, "to", elem.ID)
		}

		c.ElementsMutex.Lock()
		elem.Formats = append(elem.Formats, format)
		c.ElementsMutex.Unlock()

		if err := c.MergeElement(&elem); err != nil {
			log.Panic(err)
		}
	}
}

// GetElement get an element by id or error if not found
func (c *Config) GetElement(id string) (*element.Element, error) {
	if c.Exist(id) {
		c.ElementsMutex.RLock()
		r := c.Elements[id]
		c.ElementsMutex.RUnlock()

		return &r, nil
	}

	return nil, fmt.Errorf("element %s not found", id)
}

func FindElem(c *Config, e string) (*element.Element, error) {
	res := c.Elements[e]
	if res.ID == "" || res.ID != e {
		return nil, fmt.Errorf("%s is not in config\n Please use \"list\" command", e)
	}

	return &res, nil
}

func Elem2preURL(c *Config, e *element.Element, b ...string) (string, error) {
	myElem, err := FindElem(c, e.ID)
	if err != nil {
		return "", err
	}

	if myElem.HasParent() {
		parent, err := FindElem(c, myElem.Parent)
		if err != nil {
			return "", err
		}

		res, err := Elem2preURL(c, parent, b...)
		if err != nil {
			return "", err
		}

		res += "/"
		if myElem.File != "" { // TODO use file in config???
			res += myElem.File
		} else {
			res += myElem.ID
		}

		return res, nil
	}

	switch len(b) {
	case 1:
		return c.BaseURL + "/" + strings.Join(b, "/") + myElem.ID, nil
	case 2:
		return strings.Join(b, "/") + myElem.ID, nil
	default:
		return c.BaseURL + "/" + myElem.ID, nil
	}
}

func Elem2URL(c *Config, e *element.Element, ext string) (string, error) {
	var (
		res string
		err error
	)

	if !e.Formats.Contains(ext) {
		return "", fmt.Errorf("error!!! %s format not exist", ext)
	}

	format := c.Formats[ext]
	if format.BasePath != "" {
		if format.BaseURL != "" {
			res, err = Elem2preURL(c, e, format.BaseURL, format.BasePath)
		} else {
			res, err = Elem2preURL(c, e, format.BasePath)
		}
	} else {
		if format.BaseURL != "" {
			res, err = Elem2preURL(c, e, format.BaseURL, "")
		} else {
			res, err = Elem2preURL(c, e)
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

	fileContent, err := ioutil.ReadFile(filename) // Open file as string
	if err != nil {
		return nil, err
	}
	// Create a Config ptr
	myConfigPtr := &Config{ElementsMutex: &sync.RWMutex{}}
	// Charging fileContent into myConfigPtr
	if err := yaml.Unmarshal(fileContent, myConfigPtr); err != nil {
		return nil, err
	}
	// Everything is OK, returning myConfigPtr
	return myConfigPtr, nil
}

func IsHashable(c *Config, format string) (hashable bool, hash, hashType string) {
	hashs := []string{"md5"} // had to be globalized?

	if _, ok := c.Formats[format]; ok {
		for _, h := range hashs {
			hash := format + "." + h
			if _, ok := c.Formats[hash]; ok {
				return true, hash, h
			}
		}
	}

	return false, "", ""
}
