package main

import (
	"fmt"
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

func elem2URL(c *Config, e *Element, ext string) (string, error) {
	res, err := elem2preURL(c, e)
	if err != nil {
		return "", err
	}
	res += c.Formats[ext].Loc
	// TODO check if valid URL
	if !stringInSlice(&ext, &e.Formats) {
		return "", fmt.Errorf("Error!!! %s format not exist", ext)
	}
	return res, nil
}

func elem2preURL(c *Config, e *Element) (string, error) {
	var res string
	myElem, err := findElem(c, e.ID)
	if err != nil {
		return "", err
	}
	if myElem.hasParent() {
		parent, err := findElem(c, myElem.Parent)
		if err != nil {
			return "", err
		}
		res, err = elem2preURL(c, parent)
		if err != nil {
			return "", err
		}
		res = res + "/"
		if myElem.File != "" { //TODO use file in config???
			res = res + myElem.File
		} else {
			res = res + myElem.ID
		}
		return res, nil
	}
	res = c.BaseURL + "/" + myElem.ID
	return res, nil
}

func findElem(c *Config, e string) (*Element, error) {
	res := c.Elements[e]
	//fmt.Println("findElem", res.ID, e)
	if res.ID == "" || res.ID != e {
		return nil, fmt.Errorf("%s is not in config\n Please use \"list\" command", e)
	}
	return &res, nil
}

// stringInSlice : Check if a sting is present in a slice
// should be more easy to access to a map!
// TODO: remove it!
func stringInSlice(a *string, list *[]string) bool {
	for _, b := range *list {
		if b == *a {
			return true
		}
	}
	return false
}
