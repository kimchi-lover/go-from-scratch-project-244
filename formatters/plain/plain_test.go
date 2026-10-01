package plain_test

import (
	"code/diff"
	"code/formatters/plain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	tree := []diff.Node{
		{Key: "number", Type: diff.Changed, OldValue: 1234567.0, NewValue: 0.5},
		{Key: "nested", Type: diff.Nested, Children: []diff.Node{
			{Key: "flag", Type: diff.Unchanged, OldValue: true, NewValue: true},
			{Key: "null", Type: diff.Added, NewValue: nil},
			{Key: "object", Type: diff.Removed, OldValue: map[string]any{"key": "value"}},
			{Key: "deep", Type: diff.Nested, Children: []diff.Node{
				{Key: "list", Type: diff.Changed, OldValue: []any{1.0}, NewValue: ""},
			}},
		}},
		{Key: "string", Type: diff.Changed, OldValue: false, NewValue: "blah blah"},
		{Key: "object", Type: diff.Added, NewValue: map[string]any{}},
	}

	expected := `Property 'number' was updated. From 1234567 to 0.5
Property 'nested.null' was added with value: null
Property 'nested.object' was removed
Property 'nested.deep.list' was updated. From [complex value] to ''
Property 'string' was updated. From false to 'blah blah'
Property 'object' was added with value: [complex value]`

	assert.Equal(t, expected, plain.Format(tree))
}

func TestFormatNoChanges(t *testing.T) {
	tree := []diff.Node{
		{Key: "key", Type: diff.Unchanged, OldValue: "value", NewValue: "value"},
		{Key: "nested", Type: diff.Nested, Children: []diff.Node{
			{Key: "flag", Type: diff.Unchanged, OldValue: true, NewValue: true},
		}},
	}

	assert.Empty(t, plain.Format(tree))
}
