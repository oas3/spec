package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// OpenAPI is the root document object of the OpenAPI document.
type OpenAPI struct {
	OpenAPIFields
	SpecificationExtensions
}

type OpenAPIFields struct {
	// This string MUST be the semantic version number of the OpenAPI Specification
	// version that the OpenAPI document uses.
	OpenAPI string `yaml:"openapi" oas3:"REQUIRED"`
	// Provides metadata about the API.
	Info Info `oas3:"REQUIRED"`
	// An array of Server Objects, which provide connectivity information to a target
	// server. The default value would be a Server Object with a url value of `/`.
	Servers []Server
	// The available paths and operations for the API.
	Paths Paths `oas3:"REQUIRED"`
	// An element to hold various schemas for the specification.
	Components Components
	// A declaration of which security mechanisms can be used across the API.
	Security []SecurityRequirement
	// A list of tags used by the specification with additional metadata.
	Tags []Tag
	// Additional external documentation.
	ExternalDocs ExternalDocumentation `yaml:"externalDocs"`
}

func (a OpenAPI) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(a.OpenAPIFields)
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

func (a *OpenAPI) UnmarshalJSON(data []byte) error {
	if err := a.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &a.OpenAPIFields)
}

func (a OpenAPI) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(a.OpenAPIFields)
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

func (a *OpenAPI) UnmarshalYAML(data []byte) error {
	if err := a.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields OpenAPIFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	a.OpenAPIFields = fields
	return nil
}
