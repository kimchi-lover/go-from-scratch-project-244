package formatters

import (
	"code/diff"
	"code/formatters/json"
	"code/formatters/plain"
	"code/formatters/stylish"
	"fmt"
)

const DefaultFormat = "stylish"

type formatter func(tree []diff.Node) (string, error)

var formattersByName = map[string]formatter{
	"stylish": withoutError(stylish.Format),
	"plain":   withoutError(plain.Format),
	"json":    json.Format,
}

func Format(tree []diff.Node, format string) (string, error) {
	if format == "" {
		format = DefaultFormat
	}

	formatTree, ok := formattersByName[format]
	if !ok {
		return "", fmt.Errorf("unknown output format %q", format)
	}

	return formatTree(tree)
}

func withoutError(format func(tree []diff.Node) string) formatter {
	return func(tree []diff.Node) (string, error) {
		return format(tree), nil
	}
}
