package objects

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
)

// Components holds a set of reusable objects for different aspects of the OAS. All
// objects defined within the components object will have no effect on the API unless they
// are explicitly referenced from properties outside the components object.
//
// Key must match `^[a-zA-Z0-9\.\-_]+$`.
type Components struct {
	ComponentsFields
	SpecificationExtensions
}

type ComponentsFields struct {
	// An object to hold reusable Schema Objects.
	Schemas map[string]Schema
	// An object to hold reusable Response Objects.
	Responses map[string]Response
	// An object to hold reusable Parameter Objects.
	Parameters map[string]Parameter
	// An object to hold reusable Example Objects.
	Examples map[string]Example
	// An object to hold reusable RequestBody Objects.
	RequestBodies map[string]RequestBody `yaml:"requestBodies"`
	// An object to hold reusable Header Objects.
	Headers map[string]Header
	// An object to hold reusable SecurityScheme Objects.
	SecuritySchemes map[string]SecurityScheme `yaml:"securitySchemes"`
	// An object to hold reusable Link Objects.
	Links map[string]Link
	// An object to hold reusable Callback Objects.
	Callbacks map[string]Callback
}

func (c Components) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(c.ComponentsFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range c.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (c *Components) UnmarshalJSON(data []byte) error {
	if err := c.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &c.ComponentsFields)
}

func (c Components) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(c.ComponentsFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range c.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (c *Components) UnmarshalYAML(data []byte) error {
	if err := c.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields ComponentsFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	c.ComponentsFields = fields
	return nil
}
