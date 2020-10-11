package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/oas3/spec/objects"
)

var tagObj = objects.Tag{
	TagFields: objects.TagFields{
		Name:        "pet",
		Description: "Pets operations",
	},
}

func TestTag(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/tag.json")
		var tag objects.Tag
		if err := json.Unmarshal(raw, &tag); err != nil {
			t.Error(err)
		}
		eqTag(t, tagObj, tag)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/tag.yml")
		var tag objects.Tag
		if err := yaml.Unmarshal(raw, &tag); err != nil {
			t.Error(err)
		}
		eqTag(t, tagObj, tag)
	})
}

func eqTag(t *testing.T, t1, t2 objects.Tag) {
	eqStr(t, t1.Name, t2.Name)
	eqStr(t, t1.Description, t2.Description)
	eqExternalDoc(t, t1.ExternalDocs, t2.ExternalDocs)
}
