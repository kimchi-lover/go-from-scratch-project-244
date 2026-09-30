package stylish

import (
	"code/diff"
	"fmt"
	"strconv"
	"strings"
)

func Format(tree []diff.Node) string {
	return formatNodes(tree, 1)
}

func formatNodes(nodes []diff.Node, depth int) string {
	lines := []string{"{"}
	for _, node := range nodes {
		lines = append(lines, formatNode(node, depth))
	}
	lines = append(lines, strings.Repeat(" ", (depth-1)*4)+"}")

	return strings.Join(lines, "\n")
}

func formatNode(node diff.Node, depth int) string {
	switch node.Type {
	case diff.Nested:
		return line(depth, " ", node.Key, formatNodes(node.Children, depth+1))
	case diff.Added:
		return line(depth, "+", node.Key, formatValue(node.NewValue, depth+1))
	case diff.Removed:
		return line(depth, "-", node.Key, formatValue(node.OldValue, depth+1))
	case diff.Changed:
		removed := line(depth, "-", node.Key, formatValue(node.OldValue, depth+1))
		added := line(depth, "+", node.Key, formatValue(node.NewValue, depth+1))
		return removed + "\n" + added
	default:
		return line(depth, " ", node.Key, formatValue(node.OldValue, depth+1))
	}
}

func formatValue(value any, depth int) string {
	switch v := value.(type) {
	case map[string]any:
		return formatNodes(diff.Build(v, v), depth)
	case nil:
		return "null"
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprint(v)
	}
}

func line(depth int, sign, key, value string) string {
	return strings.Repeat(" ", depth*4-2) + sign + " " + key + ": " + value
}
