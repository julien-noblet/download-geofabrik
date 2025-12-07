package lists_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/internal/lists"
	"github.com/stretchr/testify/assert"
)

func TestListAllRegions(t *testing.T) { // it just works
	t.Parallel()

	mockConfig := &config.Config{
		Elements: map[string]element.Element{
			"region1": {Parent: "parent1", Name: "Region 1", Formats: []string{"format1"}},
			"region2": {Parent: "parent2", Name: "Region 2", Formats: []string{"format2"}},
		},
	}

	tests := []struct {
		name   string
		format string
	}{
		{name: "Default format", format: ""},
		{name: "Markdown format", format: "Markdown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := lists.ListAllRegions(mockConfig, tt.format)
			// Output: | ShortName | Is in | Long Name | formats |
			// |-----------|-------|----------|--------|
			// | region1   | parent1 | Region 1 | format1 |
			// | region2   | parent2 | Region 2 | format2 |
			// Total elements: 2

			assert.NoError(t, err)
		})
	}
}

func TestGetSortedKeys(t *testing.T) {
	t.Parallel()

	mockConfig := &config.Config{
		Elements: map[string]element.Element{
			"region2": {Parent: "parent2", Name: "Region 2", Formats: []string{"format2"}},
			"region1": {Parent: "parent1", Name: "Region 1", Formats: []string{"format1"}},
		},
	}

	keys := lists.GetSortedKeys(mockConfig)
	assert.Equal(t, []string{"region1", "region2"}, keys)
}
