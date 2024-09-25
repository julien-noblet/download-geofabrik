package element

// Element represents a part to download with formats, name, parent, etc.
type Element struct {
	ID      string  `yaml:"id"`
	File    string  `yaml:"file,omitempty"`
	Name    string  `yaml:"name,omitempty"`
	Parent  string  `yaml:"parent,omitempty"`
	Formats Formats `yaml:"files,omitempty"`
	Meta    bool    `yaml:"meta,omitempty"`
}

type Formats []string

// MapElement contains all Elements.
type MapElement map[string]Element

// HasParent checks if the element has a parent.
func (e *Element) HasParent() bool {
	return e.Parent != ""
}

// Contains checks if the format list contains a specific format.
func (f *Formats) Contains(format string) bool {
	for _, existingFormat := range *f {
		if format == existingFormat {
			return true
		}
	}
	return false
}

// CreateParentElement creates a parent element for the given element.
// Useful for meta parents.
func CreateParentElement(e *Element, grandparentID string) *Element {
	if e.HasParent() {
		return &Element{
			ID:      e.Parent,
			File:    "",
			Name:    e.Parent,
			Parent:  grandparentID,
			Formats: Formats{},
			Meta:    true,
		}
	}
	return nil
}
