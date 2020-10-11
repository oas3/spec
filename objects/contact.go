package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// Contact information for the exposed API.
type Contact struct {
	ContactFields
	SpecificationExtensions
}

type ContactFields struct {
	// The identifying name of the contact person/organization.
	Name string
	// The URL pointing to the contact information.
	URL string
	// The email address of the contact person/organization.
	Email string
}

func (c Contact) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(c.ContactFields)
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

func (c *Contact) UnmarshalJSON(data []byte) error {
	if err := c.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &c.ContactFields)
}

func (c Contact) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(c.ContactFields)
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

func (c *Contact) UnmarshalYAML(data []byte) error {
	if err := c.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields ContactFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	c.ContactFields = fields
	return nil
}
