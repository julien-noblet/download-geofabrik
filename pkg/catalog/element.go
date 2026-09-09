package catalog

import (
	"cmp"
	"slices"
)

// Formats represents a list of format identifiers.
type Formats []string

// Contains checks if a format identifier exists in the formats list.
func (f Formats) Contains(format string) bool {
	return slices.Contains(f, format)
}

// Element represents a geographic region, extract, or meta-container.
// Field alignment optimized.
type Element struct {
	// ID is the unique slug or identifier for this geographic extract (e.g., "europe/france").
	ID string `json:"id" yaml:"id"`

	// File is the extract filename template (or override) if different from ID.
	File string `json:"file,omitempty" yaml:"file,omitempty"`

	// Name is the human-readable display name of the geographic area.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Parent is the identifier of the parent container element (e.g., continent or country).
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// Formats is the list of format identifiers available for this extract.
	Formats Formats `json:"formats,omitempty" yaml:"files,omitempty"`

	// Meta indicates whether this element is an organizational category rather than a downloadable extract.
	Meta bool `json:"meta,omitempty" yaml:"meta,omitempty"`
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

	return cmp.Or(e.File, e.ID)
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

// Clone returns a deep copy of the Element with an independent Formats slice.
func (e *Element) Clone() Element {
	if e == nil {
		return Element{}
	}

	clone := *e
	clone.Formats = slices.Clone(e.Formats)

	return clone
}
