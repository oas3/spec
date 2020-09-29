package objects

// Info provides metadata about the API. The metadata MAY be used by the clients if
// needed, and MAY be presented in editing or documentation generation tools for
// convenience.
type Info struct {
	// The title of the API.
	Title string `oas3:"REQUIRED"`
	// A short description of the API.
	Description string
	// A URL to the Terms of Service for the API.
	TermsOfService string
	// The contact information for the exposed API.
	Contact Contact
	// The license information for the exposed API.
	License License
	// The version of the OpenAPI document itself.
	Version string `oas3:"REQUIRED"`
}
