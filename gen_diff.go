package code

import (
	"code/diff"
	"code/formatters"
	"code/parsers"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	map1, err := parsers.ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	map2, err := parsers.ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	tree := diff.Build(map1, map2)

	return formatters.Format(tree, format)
}
