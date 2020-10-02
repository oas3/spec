package objects

// OpenAPI is the root document object of the OpenAPI document.
type OpenAPI struct {
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
