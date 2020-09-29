package objects

// RequestBody describes a single request body.
type RequestBody struct {
	// A brief description of the request body.
	Description string
	// The content of the request body. The key is a media type or media type range and
	// the value describes it.
	Content map[string]MediaType `oas3:"REQUIRED"`
	// Determines if the request body is required in the request. Defaults to false.
	Required bool
}
