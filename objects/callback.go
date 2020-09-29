package objects

// A map of possible out-of band callbacks related to the parent operation. Each value in
// the map is a Path Item Object that describes a set of requests that may be initiated by
// the API provider and the expected responses.
type Callback map[string]PathItem
