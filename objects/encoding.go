package objects

// Encoding is an encoding definition applied to a single schema property.
type Encoding struct {
	// The Content-Type for encoding a specific property.
	ContentType string `yaml:"contentType"`
	// A map allowing additional information to be provided as headers.
	Headers map[string]Header
	// Describes how a specific property value will be serialized depending on its type.
	Style string
	// When this is true, property values of type array or object generate separate
	// parameters for each value of the array, or key-value-pair of the map.
	Explode bool
	// Determines whether the parameter value SHOULD allow reserved characters, as defined
	// by RFC3986 `:/?#[]@!$&'()*+,;=` to be included without percent-encoding.
	AllowReserved bool `yaml:"allowReserved"`
}
