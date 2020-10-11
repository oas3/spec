package objects

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
)

// Server represents a Server.
type Server struct {
	ServerFields
	SpecificationExtensions
}

type ServerFields struct {
	// A URL to the target host.
	URL string
	// An optional string describing the host designated by the URL.
	Description string
	// A map between a variable name and its value.
	Variables map[string]ServerVariable
}

func (s Server) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(s.ServerFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range s.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (s *Server) UnmarshalJSON(data []byte) error {
	if err := s.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &s.ServerFields)
}

func (s Server) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(s.ServerFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range s.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (s *Server) UnmarshalYAML(data []byte) error {
	if err := s.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields ServerFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	s.ServerFields = fields
	return nil
}

// ServerVariable represents a Server Variable for server URL template substitution.
type ServerVariable struct {
	ServerVariableFields
	SpecificationExtensions
}

type ServerVariableFields struct {
	// An enumeration of string values to be used if the substitution options are from a
	// limited set.
	Enum []string
	// The default value to use for substitution, which shall be sent if an alternate
	// value is not supplied.
	Default string `oas3:"REQUIRED"`
	// An optional description for the server variable
	Description string
}

func (v ServerVariable) MarshalJSON() ([]byte, error) {
	fields, err := json.Marshal(v.ServerVariableFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := json.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range v.SpecificationExtensions {
		fieldMap[k] = v
	}
	return json.Marshal(fieldMap)
}

func (v *ServerVariable) UnmarshalJSON(data []byte) error {
	if err := v.SpecificationExtensions.UnmarshalJSON(data); err != nil {
		return err
	}
	return json.Unmarshal(data, &v.ServerVariableFields)
}

func (v ServerVariable) MarshalYAML() (interface{}, error) {
	fields, err := yaml.Marshal(v.ServerVariableFields)
	if err != nil {
		return nil, err
	}
	var fieldMap map[string]interface{}
	if err := yaml.Unmarshal(fields, &fieldMap); err != nil {
		return nil, err
	}
	for k, v := range v.SpecificationExtensions {
		fieldMap[k] = v
	}
	return yaml.Marshal(fieldMap)
}

func (v *ServerVariable) UnmarshalYAML(data []byte) error {
	if err := v.SpecificationExtensions.UnmarshalYAML(data); err != nil {
		return err
	}
	var fields ServerVariableFields
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return err
	}
	v.ServerVariableFields = fields
	return nil
}
