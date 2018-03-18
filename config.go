package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	BaseURL  string             `yaml:"baseURL"`
	Formats  map[string]format  `yaml:"formats"`
	Elements map[string]Element `yaml:"elements"`
}

func loadConfig(configFile string) *Config {
	filename, _ := filepath.Abs(configFile)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(fmt.Errorf("File error: %v ", err))
		os.Exit(1)
	}
	var myConfig Config
	err = yaml.Unmarshal(file, &myConfig)
	if err != nil {
		panic(err)
	}
	return &myConfig

}
