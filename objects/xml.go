package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// XML allows for more fine-tuned XML model definitions.
type XML struct {
	XMLFields
	SpecificationExtensions
}

type XMLFields struct {
	// Replaces the name of the element/attribute used for the described schema property.
	Name string
	// The URI of the namespace definition. Value MUST be in the form of an absolute URI.
	Namespace string
	// The prefix to be used for the Name.
	Prefix string
	// Declares whether the property definition translates to an attribute instead of an
	// element. Default value is false.
	Attribute bool
	// Signifies whether the array is wrapped, may be used only for an array definition.
	Wrapped bool
}

func (x XML) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(x.XMLFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range x.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (x *XML) UnmarshalJSON(data []byte) error {
	if err := x.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &x.XMLFields)
}

func (x XML) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(x.XMLFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range x.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (x *XML) UnmarshalYAML(data []byte) error {
	if err := x.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields XMLFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	x.XMLFields = fields
	return nil
}
