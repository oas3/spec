package objects

// Parameter describes a single operation parameter.
// A unique parameter is defined by a combination of a name and location.
//
// Parameter Locations
// There are four possible parameter locations specified by the in field:
// - path:    Used together with Path Templating, where the parameter value is actually
//            part of the operation's URL. This does not include the host or base path of
//            the API. For example, in /items/{itemId}, the path parameter is itemId.
// - query:   Parameters that are appended to the URL. For example, in /items?id=###, the
//            query parameter is id.
// - header:  Custom headers that are expected as part of the request. Note that RFC7230
//            states header names are case insensitive.
// - cookie:  Used to pass a specific cookie value to the API.
type Parameter struct {
	// The name of the parameter. Parameter names are case sensitive.
	Name string `oas3:"REQUIRED"`
	// The location of the parameter.
	// Possible values are "query", "header", "path" or "cookie".
	In string `oas3:"REQUIRED"`
	// A brief description of the parameter.
	Description string
	// Determines whether this parameter is mandatory.
	Required bool `oas3:"REQUIRED"`
	// Specifies that a parameter is deprecated and should be transitioned out of usage.
	// Default value is false.
	Deprecated bool
	// Sets the ability to pass empty-valued parameters.
	AllowEmptyValue bool `yaml:"allowEmptyValue"`

	// The rules for serialization of the parameter are specified in one of two ways.
	// For simpler scenarios, a schema and style can describe the structure and syntax of
	// the parameter.

	// Describes how the parameter value will be serialized depending on the type of the
	// parameter value.
	Style string
	// When this is true, parameter values of type array or object generate separate
	// parameters for each value of the array or key-value pair of the map.
	Explode bool
	// Determines whether the parameter value SHOULD allow reserved characters, as defined
	// by RFC3986 `:/?#[]@!$&'()*+,;=` to be included without percent-encoding.
	AllowReserved bool `yaml:"allowReserved"`
	// The schema defining the type used for the parameter.
	Schema Schema

	// Example of the parameter's potential value.
	Example interface{}
	// Examples of the parameter's potential value.
	Examples map[string]Example

	// For more complex scenarios, the content property can define the media type and
	// schema of the parameter.

	// A map containing the representations for the parameter.
	Content map[string]MediaType
}
