package objects

import (
	"encoding/json"
	"strings"

	"github.com/goccy/go-yaml"
)

type SpecificationExtensions map[string]interface{}

func (e SpecificationExtensions) MarshalJSON() ([]byte, error) {
	extensions := make(map[string]interface{})
	for k, v := range e {
		lower := strings.ToLower(k)
		if strings.HasPrefix(lower, "x-") {
			extensions[k] = v
		}
	}
	return json.Marshal(extensions)
}

func (e *SpecificationExtensions) UnmarshalJSON(data []byte) error {
	all := make(map[string]interface{})
	if err := json.Unmarshal(data, &all); err != nil {
		return err
	}

	extensions := make(SpecificationExtensions)
	for k, v := range all {
		lower := strings.ToLower(k)
		if strings.HasPrefix(lower, "x-") {
			extensions[k] = v
		}
	}
	if len(extensions) != 0 {
		*e = extensions
	}
	return nil
}

func (e SpecificationExtensions) MarshalYAML() (interface{}, error) {
	extensions := make(map[string]interface{})
	for k, v := range e {
		lower := strings.ToLower(k)
		if strings.HasPrefix(lower, "x-") {
			extensions[k] = v
		}
	}
	return yaml.Marshal(extensions)
}

func (e *SpecificationExtensions) UnmarshalYAML(data []byte) error {
	all := make(map[string]interface{})
	if err := yaml.Unmarshal(data, &all); err != nil {
		return err
	}

	extensions := make(SpecificationExtensions)
	for k, v := range all {
		lower := strings.ToLower(k)
		if strings.HasPrefix(lower, "x-") {
			extensions[k] = v
		}
	}
	if len(extensions) != 0 {
		*e = extensions
	}
	return nil
}
