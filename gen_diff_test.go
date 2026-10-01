package code_test

import (
	"code"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixture(name string) string {
	return filepath.Join("testdata", "fixture", name)
}

func readFixture(t *testing.T, name string) string {
	t.Helper()

	data, err := os.ReadFile(fixture(name))
	require.NoError(t, err)

	return strings.TrimSuffix(string(data), "\n")
}

func TestGenDiff(t *testing.T) {
	stylish := readFixture(t, "result_stylish.txt")
	unchanged := readFixture(t, "result_stylish_unchanged.txt")
	plain := readFixture(t, "result_plain.txt")

	tests := []struct {
		name     string
		file1    string
		file2    string
		format   string
		expected string
	}{
		{
			name:     "nested json",
			file1:    "file1.json",
			file2:    "file2.json",
			format:   "stylish",
			expected: stylish,
		},
		{
			name:     "nested yaml with yml and yaml extensions",
			file1:    "file1.yml",
			file2:    "file2.yaml",
			format:   "stylish",
			expected: stylish,
		},
		{
			name:     "json and yaml",
			file1:    "file1.json",
			file2:    "file2.yaml",
			format:   "stylish",
			expected: stylish,
		},
		{
			name:     "stylish is default format",
			file1:    "file1.json",
			file2:    "file2.json",
			format:   "",
			expected: stylish,
		},
		{
			name:     "same data in json and yaml",
			file1:    "file1.json",
			file2:    "file1.yml",
			format:   "stylish",
			expected: unchanged,
		},
		{
			name:     "same data with different formatting and key order",
			file1:    "file1.json",
			file2:    "file1_reformatted.json",
			format:   "stylish",
			expected: unchanged,
		},
		{
			name:     "plain json",
			file1:    "file1.json",
			file2:    "file2.json",
			format:   "plain",
			expected: plain,
		},
		{
			name:     "plain yaml",
			file1:    "file1.yml",
			file2:    "file2.yaml",
			format:   "plain",
			expected: plain,
		},
		{
			name:     "plain without changes",
			file1:    "file1.json",
			file2:    "file1.yml",
			format:   "plain",
			expected: "",
		},
		{
			name:     "both files empty",
			file1:    "empty.json",
			file2:    "empty.json",
			format:   "stylish",
			expected: "{\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := code.GenDiff(fixture(tt.file1), fixture(tt.file2), tt.format)
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

func TestGenDiffUnknownOutputFormat(t *testing.T) {
	result, err := code.GenDiff(fixture("file1.json"), fixture("file2.json"), "xml")
	require.ErrorContains(t, err, `"xml"`)
	assert.Empty(t, result)
}
