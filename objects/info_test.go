package objects_test

import (
	"encoding/json"
	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var infoObj = objects.Info{
	Title:          "Sample Pet Store App",
	Description:    "This is a sample server for a pet store.",
	TermsOfService: "http://example.com/terms/",
	Contact:        contactObj,
	License:        licenseObj,
	Version:        "1.0.1",
}

func TestInfo(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/info.json")
		var info objects.Info
		if err := json.Unmarshal(raw, &info); err != nil {
			t.Error(err)
		}
		eqInfo(t, infoObj, info)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/info.yml")
		var info objects.Info
		if err := yaml.Unmarshal(raw, &info); err != nil {
			t.Error(err)
		}
		eqInfo(t, infoObj, info)
	})
}

func eqInfo(t *testing.T, i1, i2 objects.Info) {
	eqStr(t, i1.Title, i2.Title)
	eqStr(t, i1.Description, i2.Description)
	eqStr(t, i1.TermsOfService, i2.TermsOfService)
	eqContact(t, i1.Contact, i2.Contact)
	eqLicense(t, i1.License, i2.License)
	eqStr(t, i1.Version, i2.Version)
}
