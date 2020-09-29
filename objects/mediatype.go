package objects

// MediaType provides a schema and examples for the media type identified by its key.
type MediaType struct {
	// The schema defining the content of the request, response, or parameter.
	Schema Schema
	// Example of the media type.
	Example interface{}
	// Examples of the media type.
	Examples map[string]Example
	// A map between a property name and its encoding information.
	Encoding map[string]Encoding
}
