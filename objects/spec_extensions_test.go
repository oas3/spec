package objects_test

import (
	"encoding/json"
	"github.com/oas3/spec/objects"
	"testing"
)

func TestSpecificationExtensions_MarshalJSON(t *testing.T) {
	extensions := objects.SpecificationExtensions{
		"x-test": "test",
		"test":   "test",
	}

	raw, err := json.Marshal(extensions)
	if err != nil {
		t.Error(err)
	}
	if string(raw) != "{\"x-test\":\"test\"}" {
		t.Error("values did not match.")
	}
}

func TestSpecificationExtensions_UnmarshalJSON(t *testing.T) {
	var extensions objects.SpecificationExtensions
	if err := json.Unmarshal([]byte("{\"x-test\":\"test\",\"test\":\"test\"}"), &extensions); err != nil {
		t.Error(err)
	}
	if len(extensions) != 1 {
		t.Error("too much values.")
	}
}
