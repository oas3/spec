package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// Info provides metadata about the API. The metadata MAY be used by the clients if
// needed, and MAY be presented in editing or documentation generation tools for
// convenience.
type Info struct {
	InfoFields
	SpecificationExtensions
}

type InfoFields struct {
	// The title of the API.
	Title string `oas3:"REQUIRED"`
	// A short description of the API.
	Description string
	// A URL to the Terms of Service for the API.
	TermsOfService string `yaml:"termsOfService"`
	// The contact information for the exposed API.
	Contact Contact
	// The license information for the exposed API.
	License License
	// The version of the OpenAPI document itself.
	Version string `oas3:"REQUIRED"`
}

func (a Info) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(a.InfoFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range a.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (a *Info) UnmarshalJSON(data []byte) error {
	if err := a.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &a.InfoFields)
}

func (a Info) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(a.InfoFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range a.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (a *Info) UnmarshalYAML(data []byte) error {
	if err := a.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields InfoFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	a.InfoFields = fields
	return nil
}
