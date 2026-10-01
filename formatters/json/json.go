package json

import (
	"code/diff"
	"encoding/json"
)

func Format(tree []diff.Node) (string, error) {
	result, err := json.MarshalIndent(map[string]any{"diff": tree}, "", "  ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}
