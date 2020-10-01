package objects_test

import (
	"encoding/json"
	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var headerObj = objects.Header{
	Description: "The number of allowed requests in the current period",
	Schema: map[string]interface{}{
		"type": "integer",
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
