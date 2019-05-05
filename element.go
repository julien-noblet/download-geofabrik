package main

import (
	"fmt"
	"strings"
)

type Element struct {
	ID      string         `yaml:"id"`
	File    string         `yaml:"file,omitempty"`
	Meta    bool           `yaml:"meta,omitempty"`
	Name    string         `yaml:"name,omitempty"`
	Formats elementFormats `yaml:"files,omitempty"`
	Parent  string         `yaml:"parent,omitempty"`
}

type elementFormats []string

func (e *Element) hasParent() bool {
	return len(e.Parent) != 0
}

func elem2URL(c *Config, e *Element, ext string) (string, error) {
	var res string
	var err error
	if !e.Formats.contains(ext) {
		return "", fmt.Errorf("error!!! %s format not exist", ext)
	}
	format := c.Formats[ext]
	if format.BasePath != "" {
		if format.BaseURL != "" {
			res, err = elem2preURL(c, e, format.BaseURL, format.BasePath)
		} else {
			res, err = elem2preURL(c, e, format.BasePath)
		}
	} else {
		if format.BaseURL != "" {
			res, err = elem2preURL(c, e, format.BaseURL, "")
		} else {
			res, err = elem2preURL(c, e)
		}
	}
	// TODO check if valid URL
	if err != nil {
		return "", err
	}
	res += format.Loc
	return res, nil
}

func elem2preURL(c *Config, e *Element, b ...string) (string, error) {
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
		res, err = elem2preURL(c, parent, b...)
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
	switch len(b) {
	case 1:
		return c.BaseURL + "/" + strings.Join(b, "/") + myElem.ID, nil
	case 2:
		return strings.Join(b, "/") + myElem.ID, nil
	default:
		return c.BaseURL + "/" + myElem.ID, nil
	}
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
func stringInSlice(a *string, list *elementFormats) bool {
	for _, b := range *list {
		if b == *a {
			return true
		}
	}
	return false
}

func (f *elementFormats) contains(e string) bool {
	for _, a := range *f {
		if e == a {
			return true
		}
	}
	return false
}
