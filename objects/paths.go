package objects

// Paths holds the relative paths to the individual endpoints and their operations. The
// path is appended to the URL from the Server Object in order to construct the full URL.
//
// Key is a relative path to an individual endpoint and must begin with a forward slash.
type Paths map[string]PathItem

// Describes the operations available on a single path. The path itself is still exposed
// to the documentation viewer but they will not know which operations and parameters are
// available.
type PathItem struct {
	// Allows for an external definition of this path item.
	Ref string
	// An optional, string summary, intended to apply to all operations in this path.
	Summary string
	// An optional, string description, intended to apply to all operations in this path.
	Description string
	// A definition of a GET operation on this path.
	Get Operation
	// A definition of a PUT operation on this path.
	Put Operation
	// A definition of a POST operation on this path.
	Post Operation
	// A definition of a DELETE operation on this path.
	Delete Operation
	// A definition of a OPTIONS operation on this path.
	Options Operation
	// A definition of a HEAD operation on this path.
	Head Operation
	// A definition of a PATCH operation on this path.
	Patch Operation
	// A definition of a TRACE operation on this path.
	Trace Operation
	// An alternative server array to service all operations in this path.
	Servers []Server
	// A list of parameters that are applicable for all the operations described under
	// this path.
	Parameters []Parameter
}
