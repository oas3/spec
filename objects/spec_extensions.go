package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
	"strings"
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

	e_ := make(SpecificationExtensions)
	for k, v := range all {
		lower := strings.ToLower(k)
		if strings.HasPrefix(lower, "x-") {
			e_[k] = v
		}
	}
	if len(e_) != 0 {
		*e = e_
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

	e_ := make(SpecificationExtensions)
	for k, v := range all {
		lower := strings.ToLower(k)
		if strings.HasPrefix(lower, "x-") {
			e_[k] = v
		}
	}
	if len(e_) != 0 {
		*e = e_
	}
	return nil
}
