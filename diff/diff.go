package diff

import (
	"reflect"
	"slices"
)

type NodeType string

const (
	Added     NodeType = "added"
	Removed   NodeType = "removed"
	Unchanged NodeType = "unchanged"
	Changed   NodeType = "changed"
	Nested    NodeType = "nested"
)

type Node struct {
	Key      string
	Type     NodeType
	OldValue any
	NewValue any
	Children []Node
}

func Build(data1, data2 map[string]any) []Node {
	keys := getSortedKeys(data1, data2)
	tree := make([]Node, 0, len(keys))

	for _, key := range keys {
		tree = append(tree, buildNode(key, data1, data2))
	}

	return tree
}

func buildNode(key string, data1, data2 map[string]any) Node {
	value1, ok1 := data1[key]
	value2, ok2 := data2[key]
	map1, isMap1 := value1.(map[string]any)
	map2, isMap2 := value2.(map[string]any)

	switch {
	case !ok2:
		return Node{Key: key, Type: Removed, OldValue: value1}
	case !ok1:
		return Node{Key: key, Type: Added, NewValue: value2}
	case isMap1 && isMap2:
		return Node{Key: key, Type: Nested, Children: Build(map1, map2)}
	case reflect.DeepEqual(value1, value2):
		return Node{Key: key, Type: Unchanged, OldValue: value1, NewValue: value2}
	default:
		return Node{Key: key, Type: Changed, OldValue: value1, NewValue: value2}
	}
}

func getSortedKeys(map1, map2 map[string]any) []string {
	result := make([]string, 0, len(map1)+len(map2))

	for key := range map1 {
		result = append(result, key)
	}

	for key := range map2 {
		if _, ok := map1[key]; !ok {
			result = append(result, key)
		}
	}

	slices.Sort(result)

	return result
}
