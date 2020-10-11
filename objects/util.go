package objects

// Convert changes all the `map[interface{}]interface{}` values into
// `map[string]interface{}` and whole numbers (float64) into integers.
func Convert(i interface{}) interface{} {
	switch t := i.(type) {
	case []interface{}:
		return ConvertArray(t)
	case map[string]interface{}:
		result := make(map[string]interface{})
		for k, v := range t {
			result[k] = Convert(v)
		}
		return result
	case Schema:
		for k, v := range t.SchemaFields {
			t.SchemaFields[k] = Convert(v)
		}
		return t
	case uint64:
		return int(t)
	case float64:
		if t == float64(int64(t)) {
			return int(t)
		}
		return t
	default:
		return t
	}
}

func ConvertArray(in []interface{}) []interface{} {
	result := make([]interface{}, len(in))
	for i, v := range in {
		result[i] = Convert(v)
	}
	return result
}
