package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// Operation describes a single API operation on a path.
type Operation struct {
	OperationFields
	SpecificationExtensions
}

type OperationFields struct {
	// A list of tags for API documentation control.
	Tags []string
	// A short summary of what the operation does.
	Summary string
	// A verbose explanation of the operation behavior.
	Description string
	// Additional external documentation for this operation.
	ExternalDocs ExternalDocumentation `yaml:"externalDocs"`
	// Unique string used to identify the operation.
	OperationID string `yaml:"operationId"`
	// A list of parameters that are applicable for this operation. If a parameter is
	// already defined at the Path Item, the new definition will override it but can never
	// remove it.
	Parameters []Parameter
	// The request body applicable for this operation.
	RequestBody RequestBody `yaml:"requestBody"`
	// The list of possible responses as they are returned from executing this operation.
	Responses Responses `oas3:"REQUIRED"`
	// A map of possible out-of band callbacks related to the parent operation.
	Callbacks map[string]Callback
	// Declares this operation to be deprecated. Consumers should refrain from usage of
	// the declared operation. Default value is false.
	Deprecated bool
	// A declaration of which security mechanisms can be used for this operation.
	Security []SecurityRequirement
	// An alternative server array to service this operation. If an alternative server
	// object is specified at the PathItem Object or Root level, it will be overridden by
	// this value.
	Servers []Server
}

func (o Operation) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(o.OperationFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range o.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (o *Operation) UnmarshalJSON(data []byte) error {
	if err := o.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &o.OperationFields)
}

func (o Operation) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(o.OperationFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range o.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (o *Operation) UnmarshalYAML(data []byte) error {
	if err := o.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields OperationFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	o.OperationFields = fields
	return nil
}
