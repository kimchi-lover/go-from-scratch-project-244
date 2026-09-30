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
		"host":    "hexlet.io",
		"timeout": 50.0,
		"proxy":   "123.234.53.22",
		"follow":  false,
	}
	file2 := map[string]any{
		"timeout": 20.0,
		"verbose": true,
		"host":    "hexlet.io",
	}

	tests := []struct {
		file     string
		expected map[string]any
	}{
		{file: "file1.json", expected: file1},
		{file: "file1.yml", expected: file1},
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
