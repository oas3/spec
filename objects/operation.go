package objects

// Operation describes a single API operation on a path.
type Operation struct {
	// A list of tags for API documentation control.
	Tags []string
	// A short summary of what the operation does.
	Summary string
	// A verbose explanation of the operation behavior.
	Description string
	// Additional external documentation for this operation.
	ExternalDocs ExternalDocumentation `yaml:"externalDocs"`
	// Unique string used to identify the operation.
	OperationID string `yaml:"operationId"`
	// A list of parameters that are applicable for this operation. If a parameter is
	// already defined at the Path Item, the new definition will override it but can never
	// remove it.
	Parameters []Parameter
	// The request body applicable for this operation.
	RequestBody RequestBody `yaml:"requestBody"`
	// The list of possible responses as they are returned from executing this operation.
	Responses Responses `oas3:"REQUIRED"`
	// A map of possible out-of band callbacks related to the parent operation.
	Callbacks map[string]Callback
	// Declares this operation to be deprecated. Consumers should refrain from usage of
	// the declared operation. Default value is false.
	Deprecated bool
	// A declaration of which security mechanisms can be used for this operation.
	Security []SecurityRequirement
	// An alternative server array to service this operation. If an alternative server
	// object is specified at the PathItem Object or Root level, it will be overridden by
	// this value.
	Servers []Server
}
