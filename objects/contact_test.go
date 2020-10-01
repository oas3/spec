package objects_test

import (
	"encoding/json"
	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var contactObj = objects.Contact{
	Name:  "API Support",
	URL:   "http://www.example.com/support",
	Email: "support@example.com",
}

func TestContact(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/contact.json")
		var contact objects.Contact
		if err := json.Unmarshal(raw, &contact); err != nil {
			t.Error(err)
		}
		eqContact(t, contactObj, contact)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/contact.yml")
		var contact objects.Contact
		if err := yaml.Unmarshal(raw, &contact); err != nil {
			t.Error(err)
		}
		eqContact(t, contactObj, contact)
	})
}

func eqContact(t *testing.T, c1, c2 objects.Contact) {
	eqStr(t, c1.Name, c2.Name)
	eqStr(t, c1.URL, c2.URL)
	eqStr(t, c1.Email, c2.Email)
}
