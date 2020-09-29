package objects

// Tag adds metadata to a single tag that is used by the Operation Object.
type Tag struct {
	// The name of the tag.
	Name string `oas3:"REQUIRED"`
	// A short description for the tag.
	Description string
	// Additional external documentation for this tag.
	ExternalDocs ExternalDocumentation
}
