package element

// Element represent a part to download with formats, name, parent...
type Element struct {
	ID      string  `yaml:"id"`
	File    string  `yaml:"file,omitempty"`
	Name    string  `yaml:"name,omitempty"`
	Parent  string  `yaml:"parent,omitempty"`
	Formats Formats `yaml:"files,omitempty"`
	Meta    bool    `yaml:"meta,omitempty"`
}

type Formats []string

// Slice contain all Elements
// TODO: It's not a slice but a MAP!!!!
type Slice map[string]Element

func (e *Element) HasParent() bool {
	return e.Parent != ""
}

// StringInSlice : Check if a sting is present in a slice
// should be more easy to access to a map!
// TODO: remove it!
func StringInSlice(a *string, list *Formats) bool {
	for _, b := range *list {
		if b == *a {
			return true
		}
	}

	return false
}

func (f *Formats) Contains(e string) bool {
	for _, a := range *f {
		if e == a {
			return true
		}
	}

	return false
}

// MakeParent make e parent(id=name=gparent)
// Useful for meta parents.
func MakeParent(e *Element, gparent string) *Element {
	if e.HasParent() {
		return &Element{
			ID:      e.Parent,
			File:    "",
			Name:    e.Parent,
			Parent:  gparent,
			Formats: Formats{},
			Meta:    true,
		}
	}

	return nil
}
