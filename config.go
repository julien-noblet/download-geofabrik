// 2015-2018 copyright Julien Noblet

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

// Config structure handle all elements.
// It also contain the BaseURL and Formats...
type Config struct {
	BaseURL       string            `yaml:"baseURL"`
	Formats       formatDefinitions `yaml:"formats"`
	Elements      ElementSlice      `yaml:"elements"`
	ElementsMutex *sync.RWMutex     `yaml:"-"` //unexported
}

// ElementSlice contain all Elements
// TODO: It's not a slice but a MAP!!!!
type ElementSlice map[string]Element

// Generate Yaml config
func (c *Config) Generate() ([]byte, error) {
	return yaml.Marshal(c)
}

func (c *Config) mergeElement(element *Element) error {
	c.ElementsMutex.RLock()
	cE, ok := c.Elements[element.ID]
	c.ElementsMutex.RUnlock()
	if ok {
		if cE.Parent != element.Parent {
			return fmt.Errorf("Cant merge : Parent mismatch")
		}
		c.ElementsMutex.Lock()
		for _, f := range element.Formats {
			if !cE.Formats.contains(f) {
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
		c.Elements[element.ID] = cE
		c.ElementsMutex.Unlock()
	} else {
		c.ElementsMutex.Lock()
		c.Elements[element.ID] = *element
		c.ElementsMutex.Unlock()
	}
	return nil
}

//Exist check if id is in e.Elements
func (c *Config) Exist(id string) bool {
	c.ElementsMutex.RLock()
	r := reflect.DeepEqual(c.Elements[id], Element{})
	c.ElementsMutex.RUnlock()
	return !r
}

//AddExtension Add an extension to an element
func (c *Config) AddExtension(id, format string) {
	c.ElementsMutex.RLock()
	element := c.Elements[id]
	c.ElementsMutex.RUnlock()
	if !element.Formats.contains(format) {
		if *fVerbose && !*fQuiet && !*fProgress {
			log.Println("Add", format, "to", element.ID)
		}
		c.ElementsMutex.Lock()
		element.Formats = append(element.Formats, format)
		c.ElementsMutex.Unlock()
		c.mergeElement(&element)
	}
}

//GetElement get an element by id or error if not found
func (c *Config) GetElement(id string) (*Element, error) {
	if c.Exist(id) {
		c.ElementsMutex.RLock()
		r := c.Elements[id]
		c.ElementsMutex.RUnlock()
		return &r, nil
	}
	return nil, fmt.Errorf("element %s not found", id)
}

// loadConfig loading configFile and send *Config.
// If there is an error, return it also.
func loadConfig(configFile string) (*Config, error) {
	filename, _ := filepath.Abs(configFile)       // Get absolute path
	fileContent, err := ioutil.ReadFile(filename) // Open file as string
	if err != nil {
		return nil, err
	}
	// Create a Config ptr
	myConfigPtr := &Config{ElementsMutex: &sync.RWMutex{}}
	// Charging fileContent into myConfigPtr
	err = yaml.Unmarshal(fileContent, myConfigPtr)
	if err != nil {
		return nil, err
	}
	// Everything is OK, returning myConfigPtr
	return myConfigPtr, nil
}
