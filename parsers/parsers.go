package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.yaml.in/yaml/v3"
)

type parser func(data []byte) (map[string]any, error)

var parsersByFormat = map[string]parser{
	"json": parseJSON,
	"yml":  parseYAML,
	"yaml": parseYAML,
}

func ParseFile(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	format := strings.TrimPrefix(filepath.Ext(path), ".")

	parseFormat, ok := parsersByFormat[format]
	if !ok {
		return nil, fmt.Errorf("unknown format %q in file %s", format, path)
	}

	result, err := parseFormat(data)
	if err != nil {
		return nil, fmt.Errorf("parsing file %s: %w", path, err)
	}

	return result, nil
}

func parseJSON(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func parseYAML(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	for key, value := range result {
		result[key] = normalize(value)
	}

	return result, nil
}

func normalize(value any) any {
	switch v := value.(type) {
	case int:
		return float64(v)
	case map[string]any:
		for key, item := range v {
			v[key] = normalize(item)
		}
	case []any:
		for i, item := range v {
			v[i] = normalize(item)
		}
	}

	return value
}
