package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// Link represents a possible design-time link for a response.
type Link struct {
	LinkFields
	SpecificationExtensions
}

type LinkFields struct {
	// A relative or absolute URI reference to an OAS operation.
	OperationRef string `yaml:"operationRef"`
	// The name of an existing, resolvable OAS operation, as defined with a unique
	// OperationID.
	OperationID string `yaml:"operationId"`
	// A map representing parameters to pass to an operation as specified with OperationID
	// or identified via OperationRef.
	Parameters map[string]interface{}
	// A literal value or expression to use as a request body when calling the target
	// operation.
	RequestBody interface{} `yaml:"requestBody"`
	// A description of the link.
	Description string
	// A server object to be used by the target operation.
	Server Server
}

func (l Link) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(l.LinkFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range l.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (l *Link) UnmarshalJSON(data []byte) error {
	if err := l.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &l.LinkFields)
}

func (l Link) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(l.LinkFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range l.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (l *Link) UnmarshalYAML(data []byte) error {
	if err := l.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields LinkFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	l.LinkFields = fields
	return nil
}
