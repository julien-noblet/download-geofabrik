// Package catalog defines the thread-safe OpenStreetMap data catalog domain model.
//
// It manages geographic elements (continents, countries, regions, cities), supported
// extract file formats, and URL resolution rules across multiple extract providers.
// It provides loading and saving of YAML configuration files, hierarchical element
// lookups with automatic name normalization and cycle protection, and format resolution.
package catalog
