package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
)

var componentsObj = objects.Components{
	Schemas: map[string]objects.Schema{
		"GeneralError": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"code": map[string]interface{}{
					"type":   "integer",
					"format": "int32",
				},
				"message": map[string]interface{}{
					"type": "string",
				},
			},
		},
		"Category": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]interface{}{
					"type":   "integer",
					"format": "int64",
				},
				"name": map[string]interface{}{
					"type": "string",
				},
			},
		},
		"Tag": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]interface{}{
					"type":   "integer",
					"format": "int64",
				},
				"name": map[string]interface{}{
					"type": "string",
				},
			},
		},
	},
	Parameters: map[string]objects.Parameter{
		"skipParam": {
			Name:        "skip",
			In:          "query",
			Description: "number of items to skip",
			Required:    true,
			Schema: map[string]interface{}{
				"type":   "integer",
				"format": "int32",
			},
		},
		"limitParam": {
			Name:        "limit",
			In:          "query",
			Description: "max records to return",
			Required:    true,
			Schema: map[string]interface{}{
				"type":   "integer",
				"format": "int32",
			},
		},
	},
	Responses: map[string]objects.Response{
		"NotFound": {
			Description: "Entity not found.",
		},
		"IllegalInput": {
			Description: "Illegal input for operation.",
		},
		"GeneralError": {
			Description: "General Error",
			Content: map[string]objects.MediaType{
				"application/json": {
					Schema: map[string]interface{}{
						"$ref": "#/components/schemas/GeneralError",
					},
				},
			},
		},
	},
	SecuritySchemes: map[string]objects.SecurityScheme{
		"api_key": {
			Type: "apiKey",
			Name: "api_key",
			In:   "header",
		},
		"petstore_auth": {
			Type: "oauth2",
			Flows: objects.OAuthFlows{
				Implicit: objects.OAuthFlow{
					AuthorizationURL: "http://example.org/api/oauth/dialog",
					Scopes: map[string]string{
						"write:pets": "modify pets in your account",
						"read:pets":  "read your pets",
					},
				},
			},
		},
	},
}

func TestComponents(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/components.json")
		var components objects.Components
		if err := json.Unmarshal(raw, &components); err != nil {
			t.Error(err)
		}
		eqComponents(t, componentsObj, components)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/components.yml")
		var components objects.Components
		if err := yaml.Unmarshal(raw, &components); err != nil {
			t.Error(err)
		}
		eqComponents(t, componentsObj, components)
	})
}

func eqComponents(t *testing.T, c1, c2 objects.Components) {
	eqInt(t, len(c1.Schemas), len(c2.Schemas))
	for k, s1 := range c1.Schemas {
		s2 := c2.Schemas[k]
		eqSchema(t, s1, s2)
	}
	eqInt(t, len(c1.Responses), len(c2.Responses))
	for k, r1 := range c1.Responses {
		r2 := c2.Responses[k]
		eqResponse(t, r1, r2)
	}
	eqInt(t, len(c1.Parameters), len(c2.Parameters))
	for k, p1 := range c1.Parameters {
		p2 := c2.Parameters[k]
		eqParameter(t, p1, p2)
	}
	eqInt(t, len(c1.Examples), len(c2.Examples))
	for k, e1 := range c1.Examples {
		e2 := c2.Examples[k]
		eqExample(t, e1, e2)
	}
	eqInt(t, len(c1.RequestBodies), len(c2.RequestBodies))
	for k, b1 := range c1.RequestBodies {
		b2 := c2.RequestBodies[k]
		eqRequestBody(t, b1, b2)
	}
	eqInt(t, len(c1.Headers), len(c2.Headers))
	for k, h1 := range c1.Headers {
		h2 := c2.Headers[k]
		eqHeader(t, h1, h2)
	}
	eqInt(t, len(c1.SecuritySchemes), len(c2.SecuritySchemes))
	for k, s1 := range c1.SecuritySchemes {
		s2 := c2.SecuritySchemes[k]
		eqSecurityScheme(t, s1, s2)
	}
	eqInt(t, len(c1.Links), len(c2.Links))
	for k, l1 := range c1.Links {
		l2 := c2.Links[k]
		eqLink(t, l1, l2)
	}
	eqInt(t, len(c1.Callbacks), len(c2.Callbacks))
	for k, b1 := range c1.Callbacks {
		b2 := c2.Callbacks[k]
		eqCallback(t, b1, b2)
	}
}
