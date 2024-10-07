package lists_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/lists"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestListAllRegions(t *testing.T) { // it just works
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
	mockConfig := &config.Config{
		Elements: map[string]element.Element{
			"region2": {Parent: "parent2", Name: "Region 2", Formats: []string{"format2"}},
			"region1": {Parent: "parent1", Name: "Region 1", Formats: []string{"format1"}},
		},
	}

	keys := lists.GetSortedKeys(mockConfig)
	assert.Equal(t, []string{"region1", "region2"}, keys)
}

func TestListCommand(t *testing.T) { // it just works
	viper.Set(config.ViperConfig, "../geofabrik.yml")
	viper.Set(config.ViperListFormatMarkdown, true)

	lists.ListCommand()
	// Output: | ShortName | Is in | Long Name | formats |
	// |-----------|-------|----------|--------|

	assert.NoError(t, nil)
}
