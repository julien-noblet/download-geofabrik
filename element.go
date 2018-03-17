package main

import (
	"fmt"
	"log"
)

type Element struct {
	ID      string   `yaml:"id"`
	File    string   `yaml:"file,omitempty"`
	Meta    bool     `yaml:"meta,omitempty"`
	Name    string   `yaml:"name,omitempty"`
	Formats []string `yaml:"files,omitempty"`
	Parent  string   `yaml:"parent,omitempty"`
}

func (e *Element) hasParent() bool {
	return len(e.Parent) != 0
}

func elem2URL(c *Config, e *Element, ext string) *string {
	res := elem2preURL(c, e)
	res += c.Formats[ext].Loc
	// TODO check if valid URL
	if !stringInSlice(&ext, &e.Formats) {
		log.Fatalln(fmt.Errorf(" Error!!! %s not exist", res))
	}
	return &res
}

func elem2preURL(c *Config, e *Element) string {
	var res string
	if e.hasParent() {
		res = elem2preURL(c, findElem(c, e.Parent)) + "/"
		if e.File != "" { //TODO use file in config???
			res = res + e.File
		} else {
			res = res + e.ID
		}
	} else {
		res = c.BaseURL + "/" + e.ID
	}
	return res
}

func findElem(c *Config, e string) *Element {
	res := c.Elements[e]
	if res.ID == "" {
		log.Fatalln(fmt.Errorf(" %s is not in config\n Please use \"list\" command", e))
	}
	return &res
}

func stringInSlice(a *string, list *[]string) bool {
	for _, b := range *list {
		if b == *a {
			return true
		}
	}
	return false
}
