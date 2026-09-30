package code_test

import (
	"code"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixture(name string) string {
	return filepath.Join("testdata", "fixture", name)
}

func TestGenDiff(t *testing.T) {
	tests := []struct {
		name     string
		file1    string
		file2    string
		expected string
	}{
		{
			name:  "flat json",
			file1: "file1.json",
			file2: "file2.json",
			expected: `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`,
		},
		{
			name:  "files in reverse order",
			file1: "file2.json",
			file2: "file1.json",
			expected: `{
  + follow: false
    host: hexlet.io
  + proxy: 123.234.53.22
  - timeout: 20
  + timeout: 50
  - verbose: true
}`,
		},
		{
			name:  "same data with different formatting and key order",
			file1: "file1.json",
			file2: "file1_reformatted.json",
			expected: `{
    follow: false
    host: hexlet.io
    proxy: 123.234.53.22
    timeout: 50
}`,
		},
		{
			name:  "all keys added",
			file1: "empty.json",
			file2: "file2.json",
			expected: `{
  + host: hexlet.io
  + timeout: 20
  + verbose: true
}`,
		},
		{
			name:     "both files empty",
			file1:    "empty.json",
			file2:    "empty.json",
			expected: "{\n}",
		},
		{
			name:  "flat yaml with yml and yaml extensions",
			file1: "file1.yml",
			file2: "file2.yaml",
			expected: `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`,
		},
		{
			name:  "same data in json and yaml",
			file1: "file1.json",
			file2: "file1.yml",
			expected: `{
    follow: false
    host: hexlet.io
    proxy: 123.234.53.22
    timeout: 50
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := code.GenDiff(fixture(tt.file1), fixture(tt.file2), "stylish")
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenDiffErrors(t *testing.T) {
	tests := []struct {
		name    string
		file1   string
		file2   string
		badFile string
	}{
		{
			name:    "missing file",
			file1:   "missing.json",
			file2:   "file2.json",
			badFile: "missing.json",
		},
		{
			name:    "unsupported format",
			file1:   "file1.json",
			file2:   "unsupported.txt",
			badFile: "unsupported.txt",
		},
		{
			name:    "invalid json",
			file1:   "file1.json",
			file2:   "invalid.json",
			badFile: "invalid.json",
		},
		{
			name:    "invalid yaml",
			file1:   "file1.yml",
			file2:   "invalid.yml",
			badFile: "invalid.yml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := code.GenDiff(fixture(tt.file1), fixture(tt.file2), "stylish")
			require.ErrorContains(t, err, fixture(tt.badFile))
			assert.Empty(t, result)
		})
	}
}
