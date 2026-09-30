package parsers

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixture(name string) string {
	return filepath.Join("..", "testdata", "fixture", name)
}

func TestParseFile(t *testing.T) {
	file1 := map[string]any{
		"common": map[string]any{
			"setting1": "Value 1",
			"setting2": 200.0,
			"setting3": true,
			"setting6": map[string]any{
				"key":  "value",
				"doge": map[string]any{"wow": ""},
			},
		},
		"group1": map[string]any{
			"baz":  "bas",
			"foo":  "bar",
			"nest": map[string]any{"key": "value"},
		},
		"group2": map[string]any{
			"abc":  12345.0,
			"deep": map[string]any{"id": 45.0},
		},
	}
	file2 := map[string]any{
		"common": map[string]any{
			"follow":   false,
			"setting1": "Value 1",
			"setting3": nil,
			"setting4": "blah blah",
			"setting5": map[string]any{"key5": "value5"},
			"setting6": map[string]any{
				"key":  "value",
				"ops":  "vops",
				"doge": map[string]any{"wow": "so much"},
			},
		},
		"group1": map[string]any{
			"foo":  "bar",
			"baz":  "bars",
			"nest": "str",
		},
		"group3": map[string]any{
			"deep": map[string]any{
				"id": map[string]any{"number": 45.0},
			},
			"fee": 100500.0,
		},
	}

	tests := []struct {
		file     string
		expected map[string]any
	}{
		{file: "file1.json", expected: file1},
		{file: "file1.yml", expected: file1},
		{file: "file2.json", expected: file2},
		{file: "file2.yaml", expected: file2},
	}

	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			result, err := ParseFile(fixture(tt.file))
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseFileErrors(t *testing.T) {
	tests := []struct {
		name string
		file string
	}{
		{name: "missing file", file: "missing.json"},
		{name: "unknown format", file: "unsupported.txt"},
		{name: "invalid yaml", file: "invalid.yml"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseFile(fixture(tt.file))
			require.ErrorContains(t, err, fixture(tt.file))
			assert.Nil(t, result)
		})
	}
}

func TestParseYAMLEmpty(t *testing.T) {
	result, err := parseYAML([]byte(""))
	require.NoError(t, err)
	assert.Empty(t, result)
}

func TestParseYAMLNumbersInLists(t *testing.T) {
	result, err := parseYAML([]byte("ports: [80, 443]\n"))
	require.NoError(t, err)
	assert.Equal(t, map[string]any{"ports": []any{80.0, 443.0}}, result)
}

func TestParserErrors(t *testing.T) {
	tests := []struct {
		name  string
		parse parser
		data  string
	}{
		{
			name:  "invalid json",
			parse: parseJSON,
			data:  `{"host": "hexlet.io",`,
		},
		{
			name:  "invalid yaml",
			parse: parseYAML,
			data:  "host: hexlet.io\ntimeout: [50\n",
		},
		{
			name:  "yaml list instead of mapping",
			parse: parseYAML,
			data:  "- host\n- timeout\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.parse([]byte(tt.data))
			require.Error(t, err)
			assert.Nil(t, result)
		})
	}
}
