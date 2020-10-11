package objects

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
)

// RequestBody describes a single request body.
type RequestBody struct {
	RequestBodyFields
	SpecificationExtensions
}

type RequestBodyFields struct {
	// A brief description of the request body.
	Description string
	// The content of the request body. The key is a media type or media type range and
	// the value describes it.
	Content map[string]MediaType `oas3:"REQUIRED"`
	// Determines if the request body is required in the request. Defaults to false.
	Required bool
}

func (f RequestBody) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(f.RequestBodyFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range f.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (f *RequestBody) UnmarshalJSON(data []byte) error {
	if err := f.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &f.RequestBodyFields)
}

func (f RequestBody) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(f.RequestBodyFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range f.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (f *RequestBody) UnmarshalYAML(data []byte) error {
	if err := f.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields RequestBodyFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	f.RequestBodyFields = fields
	return nil
}
