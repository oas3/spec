package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// A map of possible out-of band callbacks related to the parent operation. Each value in
// the map is a Path Item Object that describes a set of requests that may be initiated by
// the API provider and the expected responses.
type Callback struct {
	CallbackFields
	SpecificationExtensions
}

type CallbackFields map[string]PathItem

func (c Callback) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(c.CallbackFields)
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

func (c *Callback) UnmarshalJSON(data []byte) error {
	if err := c.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &c.CallbackFields)
}

func (c Callback) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(c.CallbackFields)
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

func (c *Callback) UnmarshalYAML(data []byte) error {
	if err := c.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields CallbackFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	c.CallbackFields = fields
	return nil
}
