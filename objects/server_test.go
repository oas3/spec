package objects_test

import (
	"encoding/json"
	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var (
	serverObj = objects.Server{
		URL:         "https://development.gigantic-server.com/v1",
		Description: "Development server",
	}
	serverVariablesObj = objects.Server{
		URL:         "https://{username}.gigantic-server.com:{port}/{basePath}",
		Description: "The production API server",
		Variables: map[string]objects.ServerVariable{
			"username": {
				Default:     "demo",
				Description: "this value is assigned by the service provider, in this example `gigantic-server.com`",
			},
			"port": {
				Enum: []string{
					"8443",
					"443",
				},
				Default: "8443",
			},
			"basePath": {
				Default: "v2",
			},
		},
	}
)

func TestServer(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/server.json")
		var server objects.Server
		if err := json.Unmarshal(raw, &server); err != nil {
			t.Error(err)
		}
		eqServer(t, serverObj, server)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/server.yml")
		var server objects.Server
		if err := yaml.Unmarshal(raw, &server); err != nil {
			t.Error(err)
		}
		eqServer(t, serverObj, server)
	})
}

func TestServerVariables(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/server_var.json")
		var server objects.Server
		if err := json.Unmarshal(raw, &server); err != nil {
			t.Error(err)
		}
		eqServer(t, serverVariablesObj, server)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/server_var.yml")
		var server objects.Server
		if err := yaml.Unmarshal(raw, &server); err != nil {
			t.Error(err)
		}
		eqServer(t, serverVariablesObj, server)
	})
}

func eqServer(t *testing.T, s1, s2 objects.Server) {
	eqStr(t, s1.URL, s2.URL)
	eqStr(t, s1.URL, s2.URL)

	l1 := len(s1.Variables)
	l2 := len(s2.Variables)
	if l1 != l2 {
		t.Errorf("expected %d vars, got %d", l1, l2)
	}
	for k, v1 := range s1.Variables {
		v2, ok := s2.Variables[k]
		if !ok {
			t.Errorf("did not find key %s", k)
		}
		eqServerVariable(t, v1, v2)
	}
}

func eqServerVariable(t *testing.T, v1, v2 objects.ServerVariable) {
	l1 := len(v1.Enum)
	l2 := len(v2.Enum)
	if l1 != l2 {
		t.Errorf("expected %d values, got %d", l1, l2)
	}
	for i, v1 := range v1.Enum {
		eqStr(t, v1, v2.Enum[i])
	}
	eqStr(t, v1.Default, v2.Default)
	eqStr(t, v1.Description, v2.Description)
}
