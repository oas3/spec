package objects_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/oas3/spec/objects"
)

func TestSchemaSimplePrimitive(t *testing.T) {
	var obj = objects.Schema{
		SchemaFields: objects.SchemaFields{
			"type":   "string",
			"format": "email",
		},
	}

	const (
		JSON = `
{
  "type": "string",
  "format": "email"
}
`
		YAML = `
type: string
format: email
`
	)

	t.Run("JSON", func(t *testing.T) {
		var schema objects.Schema
		if err := json.Unmarshal([]byte(JSON), &schema); err != nil {
			t.Error(err)
		}
		eq(t, obj, schema)
	})

	t.Run("YAML", func(t *testing.T) {
		var schema objects.Schema
		if err := yaml.Unmarshal([]byte(YAML), &schema); err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(obj, schema) {
			t.Errorf("expected\n%#v\ngot\n%#v", obj, schema)
		}
	})
}

const (
	simpleModelJSON = `
{
  "type": "object",
  "required": [
    "name"
  ],
  "properties": {
    "name": {
      "type": "string"
    },
    "address": {
      "$ref": "#/components/schemas/Address"
    },
    "age": {
      "type": "integer",
      "format": "int32",
      "minimum": 0
    }
  }
}
`
	simpleModelYAML = `
type: object
required:
  - name
properties:
  name:
    type: string
  address:
    $ref: '#/components/schemas/Address'
  age:
    type: integer
    format: int32
    minimum: 0
`
)

func TestSchemaSimpleModel(t *testing.T) {
	var obj = objects.Schema{
		SchemaFields: objects.SchemaFields{
			"type":     "object",
			"required": []interface{}{"name"},
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type": "string",
				},
				"address": map[string]interface{}{
					"$ref": "#/components/schemas/Address",
				},
				"age": map[string]interface{}{
					"type":    "integer",
					"format":  "int32",
					"minimum": 0,
				},
			},
		},
	}

	t.Run("JSON", func(t *testing.T) {
		var schema objects.Schema
		if err := json.Unmarshal([]byte(simpleModelJSON), &schema); err != nil {
			t.Error(err)
		}

		eqSchema(t, obj, schema)
	})

	t.Run("YAML", func(t *testing.T) {
		var schema objects.Schema
		if err := yaml.Unmarshal([]byte(simpleModelYAML), &schema); err != nil {
			t.Error(err)
		}

		eqSchema(t, obj, schema)
	})
}

func eqSchema(t *testing.T, s1, s2 objects.Schema) {
	objects.Convert(s2)
	eq(t, s1, s2)
}
