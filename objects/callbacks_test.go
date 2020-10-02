package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
)

var callbackObj = objects.Callback{
	"{$request.query.queryUrl}": objects.PathItem{
		Post: objects.Operation{
			RequestBody: objects.RequestBody{
				Description: "Callback payload",
				Content: map[string]objects.MediaType{
					"application/json": {
						Schema: map[string]interface{}{
							"$ref": "#/components/schemas/SomePayload",
						},
					},
				},
			},
			Responses: map[string]objects.Response{
				"200": {
					Description: "callback successfully processed",
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
	eqInt(t, len(c1), len(c2))
	for k, i1 := range c1 {
		i2 := c2[k]
		eqPathItem(t, i1, i2)
	}
}
