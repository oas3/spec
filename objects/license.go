package objects

// License information for the exposed API.
type License struct {
	// The license name used for the API.
	Name string `oas3:"REQUIRED"`
	// A URL to the license used for the API.
	URL string
}
