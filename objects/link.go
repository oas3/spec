package objects

// Link represents a possible design-time link for a response.
type Link struct {
	// A relative or absolute URI reference to an OAS operation.
	OperationRef string
	// The name of an existing, resolvable OAS operation, as defined with a unique
	// OperationID.
	OperationID string
	// A map representing parameters to pass to an operation as specified with OperationID
	// or identified via OperationRef.
	Parameters map[string]interface{}
	// A literal value or expression to use as a request body when calling the target
	// operation.
	RequestBody interface{}
	// A description of the link.
	Description string
	// A server object to be used by the target operation.
	Server Server
}
