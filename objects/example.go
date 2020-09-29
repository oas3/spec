package objects

type Example struct {
	// Short description for the example.
	Summary string
	// Long description for the example.
	Description string
	// Embedded literal example.
	Value interface{}
	// A URL that points to the literal example.
	ExternalValue string
}
