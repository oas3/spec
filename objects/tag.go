package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// Tag adds metadata to a single tag that is used by the Operation Object.
type Tag struct {
	TagFields
	SpecificationExtensions
}

type TagFields struct {
	// The name of the tag.
	Name string `oas3:"REQUIRED"`
	// A short description for the tag.
	Description string
	// Additional external documentation for this tag.
	ExternalDocs ExternalDocumentation `yaml:"externalDocs"`
}

func (t Tag) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(t.TagFields)
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

func (t *Tag) UnmarshalJSON(data []byte) error {
	if err := t.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &t.TagFields)
}

func (t Tag) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(t.TagFields)
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

func (t *Tag) UnmarshalYAML(data []byte) error {
	if err := t.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields TagFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	t.TagFields = fields
	return nil
}
