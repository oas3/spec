package objects

// Components holds a set of reusable objects for different aspects of the OAS. All
// objects defined within the components object will have no effect on the API unless they
// are explicitly referenced from properties outside the components object.
//
// Key must match `^[a-zA-Z0-9\.\-_]+$`.
type Components struct {
	// An object to hold reusable Schema Objects.
	Schemas map[string]Schema
	// An object to hold reusable Response Objects.
	Responses map[string]Response
	// An object to hold reusable Parameter Objects.
	Parameters map[string]Parameter
	// An object to hold reusable Example Objects.
	Examples map[string]Example
	// An object to hold reusable RequestBody Objects.
	RequestBodies map[string]RequestBody `yaml:"requestBodies"`
	// An object to hold reusable Header Objects.
	Headers map[string]Header
	// An object to hold reusable SecurityScheme Objects.
	SecuritySchemes map[string]SecurityScheme `yaml:"securitySchemes"`
	// An object to hold reusable Link Objects.
	Links map[string]Link
	// An object to hold reusable Callback Objects.
	Callbacks map[string]Callback
}
