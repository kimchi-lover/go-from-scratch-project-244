package json_test

import (
	"code/diff"
	"code/formatters/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	tree := []diff.Node{
		{Key: "number", Type: diff.Changed, OldValue: 1234567.0, NewValue: 0.5},
		{Key: "nested", Type: diff.Nested, Children: []diff.Node{
			{Key: "flag", Type: diff.Unchanged, OldValue: true, NewValue: true},
			{Key: "object", Type: diff.Removed, OldValue: map[string]any{"b": []any{1.0}, "a": ""}},
		}},
	}

	expected := `{
  "diff": [
    {
      "key": "number",
      "type": "changed",
      "oldValue": 1234567,
      "newValue": 0.5
    },
    {
      "key": "nested",
      "type": "nested",
      "oldValue": null,
      "newValue": null,
      "children": [
        {
          "key": "flag",
          "type": "unchanged",
          "oldValue": true,
          "newValue": true
        },
        {
          "key": "object",
          "type": "removed",
          "oldValue": {
            "a": "",
            "b": [
              1
            ]
          },
          "newValue": null
        }
      ]
    }
  ]
}`

	result, err := json.Format(tree)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestFormatUnsupportedValue(t *testing.T) {
	tree := []diff.Node{{Key: "inf", Type: diff.Added, NewValue: math.Inf(1)}}

	result, err := json.Format(tree)
	require.Error(t, err)
	assert.Empty(t, result)
}
