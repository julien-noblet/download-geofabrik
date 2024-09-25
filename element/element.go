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

// MapElement contain all Elements.
type MapElement map[string]Element

func (e *Element) HasParent() bool {
	return e.Parent != ""
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
