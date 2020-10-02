package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
)

var externalDocObj = objects.ExternalDocumentation{
	Description: "Find more info here",
	URL:         "https://example.com",
}

func TestExternalDocumentation(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/external_doc.json")
		var externalDoc objects.ExternalDocumentation
		if err := json.Unmarshal(raw, &externalDoc); err != nil {
			t.Error(err)
		}
		eqExternalDoc(t, externalDocObj, externalDoc)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/external_doc.yml")
		var externalDoc objects.ExternalDocumentation
		if err := yaml.Unmarshal(raw, &externalDoc); err != nil {
			t.Error(err)
		}
		eqExternalDoc(t, externalDocObj, externalDoc)
	})
}

func eqExternalDoc(t *testing.T, d1, d2 objects.ExternalDocumentation) {
	eqStr(t, d1.Description, d2.Description)
	eqStr(t, d1.URL, d2.URL)
}
