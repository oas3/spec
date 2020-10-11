package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/oas3/spec/objects"
)

var headerObj = objects.Header{
	ParameterFields: objects.ParameterFields{
		Description: "The number of allowed requests in the current period",
		Schema: objects.Schema{
			SchemaFields: objects.SchemaFields{
				"type": "integer",
			},
		},
	},
}

func TestHeader(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/header.json")
		var header objects.Header
		if err := json.Unmarshal(raw, &header); err != nil {
			t.Error(err)
		}
		eqHeader(t, headerObj, header)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/header.yml")
		var header objects.Header
		if err := yaml.Unmarshal(raw, &header); err != nil {
			t.Error(err)
		}
		eqHeader(t, headerObj, header)
	})
}

func eqHeader(t *testing.T, h1, h2 objects.Header) {
	eqParameter(t, objects.Parameter(h1), objects.Parameter(h2))
}
