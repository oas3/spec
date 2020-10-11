package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/oas3/spec/objects"
)

var callbackObj = objects.Callback{
	CallbackFields: objects.CallbackFields{
		"{$request.query.queryUrl}": objects.PathItem{
			PathItemFields: objects.PathItemFields{
				Post: objects.Operation{
					OperationFields: objects.OperationFields{
						RequestBody: objects.RequestBody{
							RequestBodyFields: objects.RequestBodyFields{
								Description: "Callback payload",
								Content: map[string]objects.MediaType{
									"application/json": {
										MediaTypeFields: objects.MediaTypeFields{
											Schema: objects.Schema{
												SchemaFields: objects.SchemaFields{
													"$ref": "#/components/schemas/SomePayload",
												},
											},
										},
									},
								},
							},
						},
						Responses: objects.Responses{
							ResponsesFields: map[string]objects.Response{
								"200": {
									ResponseFields: objects.ResponseFields{
										Description: "callback successfully processed",
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

func TestCallback(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/callback.json")
		var callback objects.Callback
		if err := json.Unmarshal(raw, &callback); err != nil {
			t.Error(err)
		}
		eqCallback(t, callbackObj, callback)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/callback.yml")
		var callback objects.Callback
		if err := yaml.Unmarshal(raw, &callback); err != nil {
			t.Error(err)
		}
		eqCallback(t, callbackObj, callback)
	})
}

func eqCallback(t *testing.T, c1, c2 objects.Callback) {
	eqInt(t, len(c1.CallbackFields), len(c2.CallbackFields))
	for k, i1 := range c1.CallbackFields {
		i2 := c2.CallbackFields[k]
		eqPathItem(t, i1, i2)
	}
}
