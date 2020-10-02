package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
)

var licenseObj = objects.License{
	Name: "Apache 2.0",
	URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
}

func TestLicense(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/license.json")
		var license objects.License
		if err := json.Unmarshal(raw, &license); err != nil {
			t.Error(err)
		}
		eqLicense(t, licenseObj, license)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/license.yml")
		var license objects.License
		if err := yaml.Unmarshal(raw, &license); err != nil {
			t.Error(err)
		}
		eqLicense(t, licenseObj, license)
	})
}

func eqLicense(t *testing.T, l1, l2 objects.License) {
	eqStr(t, l1.Name, l2.Name)
	eqStr(t, l1.URL, l2.URL)
}
