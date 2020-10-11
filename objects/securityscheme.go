package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// SecurityScheme defines a security scheme that can be used by the operations.
type SecurityScheme struct {
	SecuritySchemeFields
	SpecificationExtensions
}

type SecuritySchemeFields struct {
	// The type of the security scheme.
	Type string `oas3:"REQUIRED"`
	// A short description for security scheme.
	Description string
	// The name of the header, query or cookie parameter to be used.
	Name string `oas3:"REQUIRED"`
	// The location of the API key. Valid values are "query", "header" or "cookie".
	In string `oas3:"REQUIRED"`
	// The name of the HTTP Authorization scheme to be used in the Authorization header
	// as defined in RFC7235.
	Scheme string `oas3:"REQUIRED"`
	// A hint to the client to identify how the bearer token is formatted.
	BearerFormat string `yaml:"bearerFormat"`
	// An object containing configuration information for the flow types supported.
	Flows OAuthFlows `oas3:"REQUIRED"`
	// OpenId Connect URL to discover OAuth2 configuration values.
	OpenIDConnectURL string `yaml:"openIdConnectUrl" oas3:"REQUIRED"`
}

func (x SecurityScheme) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(x.SecuritySchemeFields)
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

func (x *SecurityScheme) UnmarshalJSON(data []byte) error {
	if err := x.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &x.SecuritySchemeFields)
}

func (x SecurityScheme) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(x.SecuritySchemeFields)
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

func (x *SecurityScheme) UnmarshalYAML(data []byte) error {
	if err := x.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields SecuritySchemeFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	x.SecuritySchemeFields = fields
	return nil
}

// OAuthFlows are configuration details for a supported OAuth Flow.
type OAuthFlows struct {
	OAuthFlowsFields
	SpecificationExtensions
}

type OAuthFlowsFields struct {
	// Configuration for the OAuth Implicit flow
	Implicit OAuthFlow
	// Configuration for the OAuth Resource Owner Password flow.
	Password OAuthFlow
	// Configuration for the OAuth Client Credentials flow. Previously called
	// `application` in OpenAPI 2.0.
	ClientCredentials OAuthFlow `yaml:"clientCredentials"`
	// Configuration for the OAuth Authorization Code flow. Previously called `accessCode`
	// in OpenAPI 2.0.
	AuthorizationCode OAuthFlow `yaml:"authorizationCode"`
}

func (f OAuthFlows) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(f.OAuthFlowsFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range f.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (f *OAuthFlows) UnmarshalJSON(data []byte) error {
	if err := f.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &f.OAuthFlowsFields)
}

func (f OAuthFlows) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(f.OAuthFlowsFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range f.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (f *OAuthFlows) UnmarshalYAML(data []byte) error {
	if err := f.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields OAuthFlowsFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	f.OAuthFlowsFields = fields
	return nil
}

// OAuthFlow allows configuration of the supported OAuth Flows.
type OAuthFlow struct {
	OAuthFlowFields
	SpecificationExtensions
}

type OAuthFlowFields struct {
	// The authorization URL to be used for this flow.
	AuthorizationURL string `yaml:"authorizationUrl" oas3:"REQUIRED"`
	// The token URL to be used for this flow.
	TokenURL string `yaml:"tokenUrl" oas3:"REQUIRED"`
	// The URL to be used for obtaining refresh tokens.
	RefreshURL string `yaml:"refreshUrl"`
	// The available scopes for the OAuth2 security scheme. A map between the scope name
	// and a short description for it. The map may be empty.
	Scopes map[string]string `oas3:"REQUIRED"`
}

func (f OAuthFlow) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(f.OAuthFlowFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range f.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (f *OAuthFlow) UnmarshalJSON(data []byte) error {
	if err := f.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &f.OAuthFlowFields)
}

func (f OAuthFlow) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(f.OAuthFlowFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range f.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (f *OAuthFlow) UnmarshalYAML(data []byte) error {
	if err := f.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields OAuthFlowFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	f.OAuthFlowFields = fields
	return nil
}
