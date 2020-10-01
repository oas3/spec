package objects_test

import (
	"encoding/json"
	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var encodingObjs = []objects.Encoding{
	{
		ContentType: "application/xml; charset=utf-8",
	},
	{
		ContentType: "image/png, image/jpeg",
		Headers: map[string]objects.Header{
			"X-Rate-Limit-Limit": {
				Description: "The number of allowed requests in the current period",
				Schema: map[string]interface{}{
					"type": "integer",
				},
			},
		},
	},
}


func TestEncoding(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/encodings.json")
		var encodings []objects.Encoding
		if err := json.Unmarshal(raw, &encodings); err != nil {
			t.Error(err)
		}
		eqEncodings(t, encodingObjs, encodings)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/encodings.yml")
		var encodings []objects.Encoding
		if err := yaml.Unmarshal(raw, &encodings); err != nil {
			t.Error(err)
		}
		eqEncodings(t, encodingObjs, encodings)
	})
}

func eqEncodings(t *testing.T, es1, es2 []objects.Encoding) {
	eqInt(t, len(es1), len(es2))
	for i, e1 := range es1 {
		e2 := es2[i]
		eqEncoding(t, e1, e2)
	}
}

func eqEncoding(t *testing.T, e1, e2 objects.Encoding) {
	eqStr(t, e1.ContentType, e2.ContentType)
	eqInt(t, len(e1.Headers), len(e2.Headers))
	for k, h1 := range e1.Headers {
		h2 := e2.Headers[k]
		eqHeader(t, h1, h2)
	}
	eqStr(t, e1.Style, e2.Style)
	eqBool(t, e1.Explode, e2.Explode)
	eqBool(t, e1.AllowReserved, e2.AllowReserved)
}

