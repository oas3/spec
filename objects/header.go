package objects

// Header follows the structure of the Parameter Object with the following changes:
// - Name must not be specified, it is given in the corresponding headers map.
// - In must not be specified, it is implicitly in header.
// - All traits that are affected by the location MUST be applicable to a location of
//   header (for example, style).
type Header Parameter
