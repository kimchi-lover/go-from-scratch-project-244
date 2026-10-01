package formatters_test

import (
	"code/diff"
	"code/formatters"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	tree := []diff.Node{{Key: "key", Type: diff.Added, NewValue: "value"}}

	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{name: "stylish", format: "stylish", expected: "{\n  + key: value\n}"},
		{name: "plain", format: "plain", expected: "Property 'key' was added with value: 'value'"},
		{name: "default format", format: "", expected: "{\n  + key: value\n}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := formatters.Format(tree, tt.format)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatUnknown(t *testing.T) {
	result, err := formatters.Format(nil, "xml")
	require.ErrorContains(t, err, `"xml"`)
	assert.Empty(t, result)
}
