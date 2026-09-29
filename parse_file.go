package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func parseFile(path string) (map[string]any, error) {
	if ext := filepath.Ext(path); ext != ".json" {
		return nil, fmt.Errorf("unknown file format in file %s", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var object map[string]any
	if err := json.Unmarshal(data, &object); err != nil {
		return nil, fmt.Errorf("parsing file %s: %w", path, err)
	}

	return object, nil
}
