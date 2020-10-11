package objects

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
)

// License information for the exposed API.
type License struct {
	LicenseFields
	SpecificationExtensions
}

type LicenseFields struct {
	// The license name used for the API.
	Name string `oas3:"REQUIRED"`
	// A URL to the license used for the API.
	URL string
}

func (l License) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(l.LicenseFields)
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

func (l *License) UnmarshalJSON(data []byte) error {
	if err := l.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &l.LicenseFields)
}

func (l License) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(l.LicenseFields)
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

func (l *License) UnmarshalYAML(data []byte) error {
	if err := l.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields LicenseFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	l.LicenseFields = fields
	return nil
}
