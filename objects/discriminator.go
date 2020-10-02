package objects

// Discriminator can be used to aid in serialization, deserialization, and validation when
// request bodies or response payloads may be one of a number of different schemas.
type Discriminator struct {
	// The name of the property in the payload that will hold the discriminator value.
	PropertyName string `yaml:"propertyName" oas3:"REQUIRED"`
	// An object to hold mappings between payload values and schema names or references.
	Mapping      map[string]string
}
