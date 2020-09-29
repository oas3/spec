package objects

// Responses is a container for the expected responses of an operation.
type Responses map[string]Response

// Response describes a single response from an API Operation, including design-time,
// static links to operations based on the response.
type Response struct {
	// A short description of the response.
	Description string `oas3:"REQUIRED"`
	// Maps a header name to its definition.
	Headers     map[string]Header
	// A map containing descriptions of potential response payloads.
	Content     map[string]MediaType
	// A map of operations links that can be followed from the response.
	Links       map[string]Link
}
