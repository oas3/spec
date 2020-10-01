package objects_test

import (
    "encoding/json"
    "github.com/oas3/spec/objects"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "testing"
)

var (
    requestBodyObj = objects.RequestBody{
        Description: "user to add to the system",
        Content: map[string]objects.MediaType{
            "application/json": {
                Schema: map[string]interface{}{
                    "$ref": "#/components/schemas/User",
                },
                Examples: map[string]objects.Example{
                    "user": {
                        Summary:       "User Example",
                        ExternalValue: "http://foo.bar/examples/user-example.json",
                    },
                },
            },
            "application/xml": {
                Schema: map[string]interface{}{
                    "$ref": "#/components/schemas/User",
                },
                Examples: map[string]objects.Example{
                    "user": {
                        Summary:       "User example in XML",
                        ExternalValue: "http://foo.bar/examples/user-example.xml",
                    },
                },
            },
            "text/plain": {
                Examples: map[string]objects.Example{
                    "user": {
                        Summary:       "User example in Plain text",
                        ExternalValue: "http://foo.bar/examples/user-example.txt",
                    },
                },
            },
            "*/*": {
                Examples: map[string]objects.Example{
                    "user": {
                        Summary:       "User example in other format",
                        ExternalValue: "http://foo.bar/examples/user-example.whatever",
                    },
                },
            },
        },
    }
    requestBodyArrObj = objects.RequestBody{
        Description: "user to add to the system",
        Required: true,
        Content: map[string]objects.MediaType{
            "text/plain": {
                Schema: map[string]interface{}{
                    "type": "array",
                    "items": map[string]interface{}{
                        "type": "string",
                    },
                },
            },
        },
    }
)

func TestRequestBody(t *testing.T) {
    t.Run("JSON", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/request_body.json")
        var body objects.RequestBody
        if err := json.Unmarshal(raw, &body); err != nil {
            t.Error(err)
        }
        eqRequestBody(t, requestBodyObj, body)
    })

    t.Run("YAML", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/request_body.yml")
        var body objects.RequestBody
        if err := yaml.Unmarshal(raw, &body); err != nil {
            t.Error(err)
        }
        eqRequestBody(t, requestBodyObj, body)
    })
}

func TestRequestBody_array(t *testing.T) {
    t.Run("JSON", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/request_body_arr.json")
        var body objects.RequestBody
        if err := json.Unmarshal(raw, &body); err != nil {
            t.Error(err)
        }
        eqRequestBody(t, requestBodyArrObj, body)
    })

    t.Run("YAML", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/request_body_arr.yml")
        var body objects.RequestBody
        if err := yaml.Unmarshal(raw, &body); err != nil {
            t.Error(err)
        }
        eqRequestBody(t, requestBodyArrObj, body)
    })
}

func eqRequestBody(t *testing.T, b1, b2 objects.RequestBody) {
    eqStr(t, b1.Description, b2.Description)
    eqInt(t, len(b1.Content), len(b2.Content))
    for k, t1 := range b1.Content {
        t2 := b2.Content[k]
        eqMediaType(t, t1, t2)
    }
    eqBool(t, b1.Required, b2.Required)
}
