package diff_test

import (
	"code/diff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	data1 := map[string]any{
		"unchanged": "value",
		"changed":   1.0,
		"removed":   true,
		"list":      []any{1.0, 2.0},
		"nested":    map[string]any{"key": "old"},
		"replaced":  map[string]any{"key": "value"},
	}
	data2 := map[string]any{
		"unchanged": "value",
		"changed":   2.0,
		"added":     nil,
		"list":      []any{1.0, 2.0},
		"nested":    map[string]any{"key": "new"},
		"replaced":  "str",
	}

	expected := []diff.Node{
		{Key: "added", Type: diff.Added, NewValue: nil},
		{Key: "changed", Type: diff.Changed, OldValue: 1.0, NewValue: 2.0},
		{Key: "list", Type: diff.Unchanged, OldValue: []any{1.0, 2.0}, NewValue: []any{1.0, 2.0}},
		{Key: "nested", Type: diff.Nested, Children: []diff.Node{
			{Key: "key", Type: diff.Changed, OldValue: "old", NewValue: "new"},
		}},
		{Key: "removed", Type: diff.Removed, OldValue: true},
		{Key: "replaced", Type: diff.Changed, OldValue: map[string]any{"key": "value"}, NewValue: "str"},
		{Key: "unchanged", Type: diff.Unchanged, OldValue: "value", NewValue: "value"},
	}

	assert.Equal(t, expected, diff.Build(data1, data2))
}
