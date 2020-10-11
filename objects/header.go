package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// Header follows the structure of the Parameter Object with the following changes:
// - Name must not be specified, it is given in the corresponding headers map.
// - In must not be specified, it is implicitly in header.
// - All traits that are affected by the location MUST be applicable to a location of
//   header (for example, style).
type Header struct {
	ParameterFields
	SpecificationExtensions
}

func (h Header) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(h.ParameterFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range h.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (h *Header) UnmarshalJSON(data []byte) error {
	if err := h.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &h.ParameterFields)
}

func (h Header) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(h.ParameterFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range h.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (h *Header) UnmarshalYAML(data []byte) error {
	if err := h.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields ParameterFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	h.ParameterFields = fields
	return nil
}
