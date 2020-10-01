package objects_test

import (
	"encoding/json"
	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var exampleObjs = []objects.Example{
	{
		Summary: "A foo example",
		Value: map[string]interface{}{
			"foo": "bar",
		},
	},
	{
		Summary: "A bar example",
		Value: map[string]interface{}{
			"bar": "baz",
		},
	},
	{
		Summary:       "This is an example in XML",
		ExternalValue: "http://example.org/examples/address-example.xml",
	},
	{
		Summary:       "This is a text example",
		ExternalValue: "http://foo.bar/examples/address-example.txt",
	},
}

func TestExample(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/examples.json")
		var examples []objects.Example
		if err := json.Unmarshal(raw, &examples); err != nil {
			t.Error(err)
		}
		eqExamples(t, exampleObjs, examples)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/examples.yml")
		var examples []objects.Example
		if err := yaml.Unmarshal(raw, &examples); err != nil {
			t.Error(err)
		}
		eqExamples(t, exampleObjs, examples)
	})
}

func eqExamples(t *testing.T, is1, is2 []objects.Example) {
	eqInt(t, len(is1), len(is2))
	for i, i1 := range is1 {
		i2 := is2[i]
		eqExample(t, i1, i2)
	}
}

func eqExample(t *testing.T, i1, i2 objects.Example) {
	eqStr(t, i1.Summary, i2.Summary)
	eqStr(t, i1.Description, i2.Description)

	switch i1Value := i1.Value.(type) {
	case map[string]interface{}:
		var i2Value map[string]interface{}
		switch v := i2.Value.(type) {
		case map[string]interface{}:
			i2Value = v
		case map[interface{}]interface{}:
			i2Value = objects.ConvertMap(v)
		}
		eq(t, i1Value, i2Value)
	}

	eqStr(t, i1.ExternalValue, i2.ExternalValue)
}
