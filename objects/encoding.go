package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// Encoding is an encoding definition applied to a single schema property.
type Encoding struct {
	EncodingFields
	SpecificationExtensions
}

type EncodingFields struct {
	// The Content-Type for encoding a specific property.
	ContentType string `yaml:"contentType"`
	// A map allowing additional information to be provided as headers.
	Headers map[string]Header
	// Describes how a specific property value will be serialized depending on its type.
	Style string
	// When this is true, property values of type array or object generate separate
	// parameters for each value of the array, or key-value-pair of the map.
	Explode bool
	// Determines whether the parameter value SHOULD allow reserved characters, as defined
	// by RFC3986 `:/?#[]@!$&'()*+,;=` to be included without percent-encoding.
	AllowReserved bool `yaml:"allowReserved"`
}

func (e Encoding) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(e.EncodingFields)
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

func (e *Encoding) UnmarshalJSON(data []byte) error {
	if err := e.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &e.EncodingFields)
}

func (e Encoding) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(e.EncodingFields)
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

func (e *Encoding) UnmarshalYAML(data []byte) error {
	if err := e.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields EncodingFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	e.EncodingFields = fields
	return nil
}
