package stylish_test

import (
	"code/diff"
	"code/formatters/stylish"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	tree := []diff.Node{
		{Key: "number", Type: diff.Changed, OldValue: 1234567.0, NewValue: 0.5},
		{Key: "nested", Type: diff.Nested, Children: []diff.Node{
			{Key: "flag", Type: diff.Unchanged, OldValue: true, NewValue: true},
			{Key: "null", Type: diff.Added, NewValue: nil},
			{Key: "object", Type: diff.Removed, OldValue: map[string]any{
				"c": false,
				"a": "text",
				"b": map[string]any{},
			}},
		}},
	}

	expected := `{
  - number: 1234567
  + number: 0.5
    nested: {
        flag: true
      + null: null
      - object: {
            a: text
            b: {
            }
            c: false
        }
    }
}`

	assert.Equal(t, expected, stylish.Format(tree))
}
