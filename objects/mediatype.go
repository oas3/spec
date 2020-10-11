package objects

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
)

// MediaType provides a schema and examples for the media type identified by its key.
type MediaType struct {
	MediaTypeFields
	SpecificationExtensions
}

type MediaTypeFields struct {
	// The schema defining the content of the request, response, or parameter.
	Schema Schema
	// Example of the media type.
	Example interface{}
	// Examples of the media type.
	Examples map[string]Example
	// A map between a property name and its encoding information.
	Encoding map[string]Encoding
}

func (t MediaType) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(t.MediaTypeFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range t.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (t *MediaType) UnmarshalJSON(data []byte) error {
	if err := t.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &t.MediaTypeFields)
}

func (t MediaType) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(t.MediaTypeFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range t.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (t *MediaType) UnmarshalYAML(data []byte) error {
	if err := t.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields MediaTypeFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	t.MediaTypeFields = fields
	return nil
}
