package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/oas3/spec/objects"
)

var (
	responsesObj = objects.Responses{
		ResponsesFields: map[string]objects.Response{
			"200": {
				ResponseFields: objects.ResponseFields{
					Description: "a pet to be returned",
					Content: map[string]objects.MediaType{
						"application/json": {
							MediaTypeFields: objects.MediaTypeFields{
								Schema: objects.Schema{
									SchemaFields: objects.SchemaFields{
										"$ref": "#/components/schemas/Pet",
									},
								},
							},
						},
					},
				},
			},
			"default": {
				ResponseFields: objects.ResponseFields{
					Description: "Unexpected error",
					Content: map[string]objects.MediaType{
						"application/json": {
							MediaTypeFields: objects.MediaTypeFields{
								Schema: objects.Schema{
									SchemaFields: objects.SchemaFields{
										"$ref": "#/components/schemas/ErrorModel",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	responseObj = objects.Response{
		ResponseFields: objects.ResponseFields{
			Description: "A complex object array response",
			Content: map[string]objects.MediaType{
				"application/json": {
					MediaTypeFields: objects.MediaTypeFields{
						Schema: objects.Schema{
							SchemaFields: objects.SchemaFields{
								"type": "array",
								"items": map[string]interface{}{
									"$ref": "#/components/schemas/VeryComplexType",
								},
							},
						},
					},
				},
			},
		},
	}
	responseEmptyObj = objects.Response{
		ResponseFields: objects.ResponseFields{
			Description: "object created",
		},
	}
	responsePlainObj = objects.Response{
		ResponseFields: objects.ResponseFields{
			Description: "A simple string response",
			Content: map[string]objects.MediaType{
				"text/plain": {
					MediaTypeFields: objects.MediaTypeFields{
						Schema: objects.Schema{
							SchemaFields: objects.SchemaFields{
								"type":    "string",
								"example": "whoa!",
							},
						},
					},
				},
			},
			Headers: map[string]objects.Header{
				"X-Rate-Limit-Limit": {
					ParameterFields: objects.ParameterFields{
						Description: "The number of allowed requests in the current period",
						Schema: objects.Schema{
							SchemaFields: objects.SchemaFields{
								"type": "integer",
							},
						},
					},
				},
				"X-Rate-Limit-Remaining": {
					ParameterFields: objects.ParameterFields{
						Description: "The number of remaining requests in the current period",
						Schema: objects.Schema{
							SchemaFields: objects.SchemaFields{
								"type": "integer",
							},
						},
					},
				},
				"X-Rate-Limit-Reset": {
					ParameterFields: objects.ParameterFields{
						Description: "The number of seconds left in the current period",
						Schema: objects.Schema{
							SchemaFields: objects.SchemaFields{
								"type": "integer",
							},
						},
					},
				},
			},
		},
	}
	responseStrObj = objects.Response{
		ResponseFields: objects.ResponseFields{
			Description: "A simple string response",
			Content: map[string]objects.MediaType{
				"text/plain": {
					MediaTypeFields: objects.MediaTypeFields{
						Schema: objects.Schema{
							SchemaFields: objects.SchemaFields{
								"type": "string",
							},
						},
					},
				},
			},
		},
	}
)

func TestResponses(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/responses.json")
		var responses objects.Responses
		if err := json.Unmarshal(raw, &responses); err != nil {
			t.Error(err)
		}
		eqResponses(t, responsesObj, responses)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/responses.json")
		var responses objects.Responses
		if err := yaml.Unmarshal(raw, &responses); err != nil {
			t.Error(err)
		}
		eqResponses(t, responsesObj, responses)
	})
}

func TestResponse(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/response.json")
		var response objects.Response
		if err := json.Unmarshal(raw, &response); err != nil {
			t.Error(err)
		}
		eqResponse(t, responseObj, response)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/response.json")
		var response objects.Response
		if err := yaml.Unmarshal(raw, &response); err != nil {
			t.Error(err)
		}
		eqResponse(t, responseObj, response)
	})
}

func TestResponse_empty(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/response_empty.json")
		var response objects.Response
		if err := json.Unmarshal(raw, &response); err != nil {
			t.Error(err)
		}
		eqResponse(t, responseEmptyObj, response)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/response_empty.json")
		var response objects.Response
		if err := yaml.Unmarshal(raw, &response); err != nil {
			t.Error(err)
		}
		eqResponse(t, responseEmptyObj, response)
	})
}

func TestResponse_plain(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/response_plain.json")
		var response objects.Response
		if err := json.Unmarshal(raw, &response); err != nil {
			t.Error(err)
		}
		eqResponse(t, responsePlainObj, response)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/response_plain.json")
		var response objects.Response
		if err := yaml.Unmarshal(raw, &response); err != nil {
			t.Error(err)
		}
		eqResponse(t, responsePlainObj, response)
	})
}

func TestResponse_string(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/response_string.json")
		var response objects.Response
		if err := json.Unmarshal(raw, &response); err != nil {
			t.Error(err)
		}
		eqResponse(t, responseStrObj, response)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/response_string.json")
		var response objects.Response
		if err := yaml.Unmarshal(raw, &response); err != nil {
			t.Error(err)
		}
		eqResponse(t, responseStrObj, response)
	})
}

func eqResponses(t *testing.T, rs1, rs2 objects.Responses) {
	eqInt(t, len(rs1.ResponsesFields), len(rs2.ResponsesFields))
	for k, r1 := range rs1.ResponsesFields {
		r2 := rs2.ResponsesFields[k]
		eqResponse(t, r1, r2)
	}
}

func eqResponse(t *testing.T, r1, r2 objects.Response) {
	eqStr(t, r1.Description, r2.Description)
	eqInt(t, len(r1.Headers), len(r2.Headers))
	for k, h1 := range r1.Headers {
		h2 := r2.Headers[k]
		eqHeader(t, h1, h2)
	}
	eqInt(t, len(r1.Content), len(r2.Content))
	for k, t1 := range r1.Content {
		t2 := r2.Content[k]
		eqMediaType(t, t1, t2)
	}
	eqInt(t, len(r1.Links), len(r2.Links))
	for k, l1 := range r1.Links {
		l2 := r2.Links[k]
		eqLink(t, l1, l2)
	}
}
