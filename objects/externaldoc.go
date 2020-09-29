package objects

// ExternalDocumentation allows referencing an external resource for extended
// documentation.
type ExternalDocumentation struct {
	// A short description of the target documentation.
	Description string
	// The URL for the target documentation.
	URL string `oas3:"REQUIRED"`
}
