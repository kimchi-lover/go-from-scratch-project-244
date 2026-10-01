package plain

import (
	"code/diff"
	"fmt"
	"strconv"
	"strings"
)

func Format(tree []diff.Node) string {
	return strings.Join(formatNodes(tree, ""), "\n")
}

func formatNodes(nodes []diff.Node, parentPath string) []string {
	lines := make([]string, 0, len(nodes))

	for _, node := range nodes {
		path := parentPath + node.Key

		switch node.Type {
		case diff.Nested:
			lines = append(lines, formatNodes(node.Children, path+".")...)
		case diff.Added:
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", path, formatValue(node.NewValue)))
		case diff.Removed:
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", path))
		case diff.Changed:
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s",
				path, formatValue(node.OldValue), formatValue(node.NewValue)))
		}
	}

	return lines
}

func formatValue(value any) string {
	switch v := value.(type) {
	case map[string]any, []any:
		return "[complex value]"
	case string:
		return "'" + v + "'"
	case nil:
		return "null"
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprint(v)
	}
}
