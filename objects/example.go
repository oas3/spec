package objects

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
)

type Example struct {
	ExampleFields
	SpecificationExtensions
}

type ExampleFields struct {
	// Short description for the example.
	Summary string
	// Long description for the example.
	Description string
	// Embedded literal example.
	Value interface{}
	// A URL that points to the literal example.
	ExternalValue string `yaml:"externalValue"`
}

func (e Example) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(e.ExampleFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range e.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (e *Example) UnmarshalJSON(data []byte) error {
	if err := e.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &e.ExampleFields)
}

func (e Example) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(e.ExampleFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range e.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (e *Example) UnmarshalYAML(data []byte) error {
	if err := e.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields ExampleFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	e.ExampleFields = fields
	return nil
}
