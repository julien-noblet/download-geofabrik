package catalog

import "slices"

// Formats represents a list of format identifiers.
type Formats []string

// Contains checks if a format identifier exists in the formats list.
func (f Formats) Contains(format string) bool {
	return slices.Contains(f, format)
}

// Element represents a geographic region, extract, or meta-container.
// Field alignment optimized.
type Element struct {
	ID      string  `json:"id"                yaml:"id"`
	File    string  `json:"file,omitempty"    yaml:"file,omitempty"`
	Name    string  `json:"name,omitempty"    yaml:"name,omitempty"`
	Parent  string  `json:"parent,omitempty"  yaml:"parent,omitempty"`
	Formats Formats `json:"formats,omitempty" yaml:"files,omitempty"`
	Meta    bool    `json:"meta,omitempty"    yaml:"meta,omitempty"`
}

// HasParent returns true if the element has a parent identifier.
func (e *Element) HasParent() bool {
	return e != nil && e.Parent != ""
}

// Filename returns the explicit filename if set, or defaults to the element ID.
func (e *Element) Filename() string {
	if e == nil {
		return ""
	}

	if e.File != "" {
		return e.File
	}

	return e.ID
}

// ContainsFormat returns true if the element supports the given format.
func (e *Element) ContainsFormat(format string) bool {
	if e == nil {
		return false
	}

	return e.Formats.Contains(format)
}

// AddFormat adds a format if it does not already exist.
func (e *Element) AddFormat(format string) {
	if e == nil || format == "" {
		return
	}

	if !e.ContainsFormat(format) {
		e.Formats = append(e.Formats, format)
	}
}

// CreateParentElement creates a synthetic meta-parent element for the current element.
func (e *Element) CreateParentElement(grandparentID string) *Element {
	if e == nil || !e.HasParent() {
		return nil
	}

	return &Element{
		ID:      e.Parent,
		Name:    e.Parent,
		Parent:  grandparentID,
		Formats: Formats{},
		Meta:    true,
	}
}
