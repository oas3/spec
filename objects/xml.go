package objects

// XML allows for more fine-tuned XML model definitions.
type XML struct {
	// Replaces the name of the element/attribute used for the described schema property.
	Name string
	// The URI of the namespace definition. Value MUST be in the form of an absolute URI.
	Namespace string
	// The prefix to be used for the Name.
	Prefix string
	// Declares whether the property definition translates to an attribute instead of an
	// element. Default value is false.
	Attribute bool
	// Signifies whether the array is wrapped, may be used only for an array definition.
	Wrapped bool
}
