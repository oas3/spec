package objects

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
)

// ExternalDocumentation allows referencing an external resource for extended
// documentation.
type ExternalDocumentation struct {
	ExternalDocumentationFields
	SpecificationExtensions
}

type ExternalDocumentationFields struct {
	// A short description of the target documentation.
	Description string
	// The URL for the target documentation.
	URL string `oas3:"REQUIRED"`
}

func (d ExternalDocumentation) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(d.ExternalDocumentationFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range d.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (d *ExternalDocumentation) UnmarshalJSON(data []byte) error {
	if err := d.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &d.ExternalDocumentationFields)
}

func (d ExternalDocumentation) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(d.ExternalDocumentationFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range d.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (d *ExternalDocumentation) UnmarshalYAML(data []byte) error {
	if err := d.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields ExternalDocumentationFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	d.ExternalDocumentationFields = fields
	return nil
}
