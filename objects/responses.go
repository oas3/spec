package objects

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
)

// Responses is a container for the expected responses of an operation.
type Responses struct {
	ResponsesFields
	SpecificationExtensions
}

type ResponsesFields map[string]Response

func (r Responses) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(r.ResponsesFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range r.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (r *Responses) UnmarshalJSON(data []byte) error {
	if err := r.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &r.ResponsesFields)
}

func (r Responses) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(r.ResponsesFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range r.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (r *Responses) UnmarshalYAML(data []byte) error {
	if err := r.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields ResponsesFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	r.ResponsesFields = fields
	return nil
}

// Response describes a single response from an API Operation, including design-time,
// static links to operations based on the response.
type Response struct {
	ResponseFields
	SpecificationExtensions
}

type ResponseFields struct {
	// A short description of the response.
	Description string `oas3:"REQUIRED"`
	// Maps a header name to its definition.
	Headers map[string]Header
	// A map containing descriptions of potential response payloads.
	Content map[string]MediaType
	// A map of operations links that can be followed from the response.
	Links map[string]Link
}

func (r Response) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(r.ResponseFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range r.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (r *Response) UnmarshalJSON(data []byte) error {
	if err := r.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &r.ResponseFields)
}

func (r Response) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(r.ResponseFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range r.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (r *Response) UnmarshalYAML(data []byte) error {
	if err := r.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields ResponseFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	r.ResponseFields = fields
	return nil
}
