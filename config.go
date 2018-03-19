// 2015-2018 copyright Julien Noblet

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Config structure handle all elements.
// It also contain the BaseURL and Formats...
type Config struct {
	BaseURL  string             `yaml:"baseURL"`
	Formats  map[string]format  `yaml:"formats"`
	Elements map[string]Element `yaml:"elements"`
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
	myConfigPtr := new(Config)
	// Charging fileContent into myConfigPtr
	err = yaml.Unmarshal(fileContent, myConfigPtr)
	if err != nil {
		return nil, err
	}
	// Everything is OK, returning myConfigPtr
	return myConfigPtr, nil
}
