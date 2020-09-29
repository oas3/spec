package objects

// Server represents a Server.
type Server struct {
	// A URL to the target host.
	URL string
	// An optional string describing the host designated by the URL.
	Description string
	// A map between a variable name and its value.
	Variables map[string]ServerVariable
}

// ServerVariable represents a Server Variable for server URL template substitution.
type ServerVariable struct {
	// An enumeration of string values to be used if the substitution options are from a
	// limited set.
	Enum []string
	// The default value to use for substitution, which shall be sent if an alternate
	// value is not supplied.
	Default string `oas3:"REQUIRED"`
	// An optional description for the server variable
	Description string
}
