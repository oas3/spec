package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// Schema allows the definition of input and output data types.
type Schema struct {
	SchemaFields
	SpecificationExtensions
}

type SchemaFields map[string]interface{}

func (s Schema) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(s.SchemaFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range s.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (s *Schema) UnmarshalJSON(data []byte) error {
	if err := s.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &s.SchemaFields)
}

func (s Schema) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(s.SchemaFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range s.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (s *Schema) UnmarshalYAML(data []byte) error {
	if err := s.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields SchemaFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	s.SchemaFields = fields
	return nil
}
