package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
)

var mediaTypeObj = objects.MediaType{
	Schema: map[string]interface{}{
		"$ref": "#/components/schemas/Pet",
	},
	Examples: map[string]objects.Example{
		"cat": {
			Summary: "An example of a cat",
			Value: map[string]interface{}{
				"name":    "Fluffy",
				"petType": "Cat",
				"color":   "White",
				"gender":  "male",
				"breed":   "Persian",
			},
		},
		"dog": {
			Summary: "An example of a dog with a cat's name",
			Value: map[string]interface{}{
				"name":    "Puma",
				"petType": "Dog",
				"color":   "Black",
				"gender":  "Female",
				"breed":   "Mixed",
			},
		},
	},
}

func TestMediaType(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/media_type.json")
		var mediaType objects.MediaType
		if err := json.Unmarshal(raw, &mediaType); err != nil {
			t.Error(err)
		}
		eqMediaType(t, mediaTypeObj, mediaType)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/media_type.yml")
		var mediaType objects.MediaType
		if err := yaml.Unmarshal(raw, &mediaType); err != nil {
			t.Error(err)
		}
		eqMediaType(t, mediaTypeObj, mediaType)
	})
}

func eqMediaType(t *testing.T, t1, t2 objects.MediaType) {
	eqSchema(t, t1.Schema, t2.Schema)
	eq(t, t1.Example, t2.Example)
	eqInt(t, len(t1.Examples), len(t2.Examples))
	for k, e1 := range t1.Examples {
		e2 := t2.Examples[k]
		eqExample(t, e1, e2)
	}
	eqInt(t, len(t1.Encoding), len(t2.Encoding))
	for k, e1 := range t1.Encoding {
		e2 := t2.Encoding[k]
		eqEncoding(t, e1, e2)
	}
}
