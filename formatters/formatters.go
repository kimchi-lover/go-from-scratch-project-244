package formatters

import (
	"code/diff"
	"code/formatters/plain"
	"code/formatters/stylish"
	"fmt"
)

const DefaultFormat = "stylish"

type formatter func(tree []diff.Node) string

var formattersByName = map[string]formatter{
	"stylish": stylish.Format,
	"plain":   plain.Format,
}

func Format(tree []diff.Node, format string) (string, error) {
	if format == "" {
		format = DefaultFormat
	}

	formatTree, ok := formattersByName[format]
	if !ok {
		return "", fmt.Errorf("unknown output format %q", format)
	}

	return formatTree(tree), nil
}
