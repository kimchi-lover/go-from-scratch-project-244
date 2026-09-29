package code

import (
	"fmt"
	"slices"
	"strings"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	map1, err := parseFile(filepath1)
	if err != nil {
		return "", err
	}

	map2, err := parseFile(filepath2)
	if err != nil {
		return "", err
	}

	keys := getSortedKeys(map1, map2)
	result := []string{"{"}

	for _, k := range keys {
		v1, ok1 := map1[k]
		v2, ok2 := map2[k]

		switch {
		case !ok2:
			result = append(result, fmt.Sprintf("  - %s: %v", k, v1))
		case !ok1:
			result = append(result, fmt.Sprintf("  + %s: %v", k, v2))
		case v1 == v2:
			result = append(result, fmt.Sprintf("    %s: %v", k, v1))
		default:
			result = append(result, fmt.Sprintf("  - %s: %v", k, v1))
			result = append(result, fmt.Sprintf("  + %s: %v", k, v2))
		}
	}
	result = append(result, "}")
	return strings.Join(result, "\n"), nil
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
