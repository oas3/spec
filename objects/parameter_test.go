package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/oas3/spec/objects"
)

var parameterObjs = []objects.Parameter{
	{
		ParameterFields: objects.ParameterFields{
			Name:        "token",
			In:          "header",
			Description: "token to be passed as a header",
			Required:    true,
			Style:       "simple",
			Schema: objects.Schema{
				SchemaFields: objects.SchemaFields{
					"type": "array",
					"items": map[string]interface{}{
						"type":   "integer",
						"format": "int64",
					},
				},
			},
		},
	},
	{
		ParameterFields: objects.ParameterFields{
			Name:        "username",
			In:          "path",
			Description: "username to fetch",
			Required:    true,
			Schema: objects.Schema{
				SchemaFields: objects.SchemaFields{
					"type": "string",
				},
			},
		},
	},
	{
		ParameterFields: objects.ParameterFields{
			Name:        "id",
			In:          "query",
			Description: "ID of the object to fetch",
			Required:    false,
			Schema: objects.Schema{
				SchemaFields: objects.SchemaFields{
					"type": "array",
					"items": map[string]interface{}{
						"type": "string",
					},
				},
			},
			Style:   "form",
			Explode: true,
		},
	},
	{
		ParameterFields: objects.ParameterFields{
			Name: "freeForm",
			In:   "query",
			Schema: objects.Schema{
				SchemaFields: objects.SchemaFields{
					"type": "object",
					"additionalProperties": map[string]interface{}{
						"type": "integer",
					},
				},
			},
			Style: "form",
		},
	},
	{
		ParameterFields: objects.ParameterFields{
			Name: "coordinates",
			In:   "query",
			Content: map[string]objects.MediaType{
				"application/json": {
					MediaTypeFields: objects.MediaTypeFields{
						Schema: objects.Schema{
							SchemaFields: objects.SchemaFields{
								"type": "object",
								"required": []interface{}{
									"lat", "long",
								},
								"properties": map[string]interface{}{
									"lat": map[string]interface{}{
										"type": "number",
									},
									"long": map[string]interface{}{
										"type": "number",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestParameter(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/parameters.json")
		var parameters []objects.Parameter
		if err := json.Unmarshal(raw, &parameters); err != nil {
			t.Error(err)
		}
		eqParameters(t, parameterObjs, parameters)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/parameters.yml")
		var parameters []objects.Parameter
		if err := yaml.Unmarshal(raw, &parameters); err != nil {
			t.Error(err)
		}
		eqParameters(t, parameterObjs, parameters)
	})
}

func eqParameters(t *testing.T, ps1, ps2 []objects.Parameter) {
	eqInt(t, len(ps1), len(ps2))
	for i, p1 := range ps1 {
		p2 := ps2[i]
		eqParameter(t, p1, p2)
	}
}

func eqParameter(t *testing.T, p1, p2 objects.Parameter) {
	eqStr(t, p1.Name, p2.Name)
	eqStr(t, p1.In, p2.In)
	eqStr(t, p1.Description, p2.Description)
	eqBool(t, p1.Required, p2.Required)
	eqBool(t, p1.Deprecated, p2.Deprecated)
	eqBool(t, p1.AllowEmptyValue, p2.AllowEmptyValue)
	eqStr(t, p1.Style, p2.Style)
	eqBool(t, p1.Explode, p2.Explode)
	eqBool(t, p1.AllowReserved, p2.AllowReserved)
	eqSchema(t, p1.Schema, p2.Schema)
	eq(t, p1.Example, p2.Example)
	eqInt(t, len(p1.Examples), len(p2.Examples))
	for k, e1 := range p1.Examples {
		e2 := p2.Examples[k]
		eqExample(t, e1, e2)
	}
	eqInt(t, len(p1.Content), len(p2.Content))
	for k, t1 := range p1.Content {
		t2 := p2.Content[k]
		eqMediaType(t, t1, t2)
	}
}
