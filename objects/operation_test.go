package objects_test

import (
    "encoding/json"
    "github.com/oas3/spec/objects"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "testing"
)

var operationObj = objects.Operation{
    Tags:        []string{"pet"},
    Summary:     "Updates a pet in the store with form data",
    OperationID: "updatePetWithForm",
    Parameters: []objects.Parameter{
        {
            Name:        "petId",
            In:          "path",
            Description: "ID of pet that needs to be updated",
            Required:    true,
            Schema: map[string]interface{}{
                "type": "string",
            },
        },
    },
    RequestBody: objects.RequestBody{
        Content: map[string]objects.MediaType{
            "application/x-www-form-urlencoded": {
                Schema: map[string]interface{}{
                    "type": "object",
                    "properties": map[string]interface{}{
                        "name": map[string]interface{}{
                            "description": "Updated name of the pet",
                            "type":        "string",
                        },
                        "status": map[string]interface{}{
                            "description": "Updated status of the pet",
                            "type":        "string",
                        },
                    },
                    "required": []interface{}{"status"},
                },
            },
        },
    },
    Responses: map[string]objects.Response{
        "200": {
            Description: "Pet updated.",
            Content: map[string]objects.MediaType{
                "application/json": {},
                "application/xml":  {},
            },
        },
        "405": {
            Description: "Method Not Allowed",
            Content: map[string]objects.MediaType{
                "application/json": {},
                "application/xml":  {},
            },
        },
    },
    Security: []objects.SecurityRequirement{
        {
            "petstore_auth": []string{
                "write:pets",
                "read:pets",
            },
        },
    },
}

func TestOperation(t *testing.T) {
    t.Run("JSON", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/operation.json")
        var operation objects.Operation
        if err := json.Unmarshal(raw, &operation); err != nil {
            t.Error(err)
        }
        eqOperation(t, operationObj, operation)
    })

    t.Run("YAML", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/operation.yml")
        var operation objects.Operation
        if err := yaml.Unmarshal(raw, &operation); err != nil {
            t.Error(err)
        }
        eqOperation(t, operationObj, operation)
    })
}

func eqOperation(t *testing.T, o1, o2 objects.Operation) {
    eqInt(t, len(o1.Tags), len(o2.Tags))
    for i, t1 := range o1.Tags {
        t2 := o2.Tags[i]
        eqStr(t, t1, t2)
    }
    eqStr(t, o1.Summary, o2.Summary)
    eqStr(t, o1.Description, o2.Description)
    eqExternalDoc(t, o1.ExternalDocs, o2.ExternalDocs)
    eqStr(t, o1.OperationID, o2.OperationID)
    eqInt(t, len(o1.Parameters), len(o2.Parameters))
    for i, p1 := range o1.Parameters {
        p2 := o2.Parameters[i]
        eqParameter(t, p1, p2)
    }
    eqRequestBody(t, o1.RequestBody, o2.RequestBody)
    eqResponses(t, o1.Responses, o2.Responses)
    eqInt(t, len(o1.Callbacks), len(o2.Callbacks))
    for k, c1 := range o1.Callbacks {
        c2 := o2.Callbacks[k]
        eqCallback(t, c1, c2)
    }
    eqBool(t, o1.Deprecated, o2.Deprecated)
    eqInt(t, len(o1.Security), len(o2.Security))
    for i, sr1 := range o1.Security {
        sr2 := o2.Security[i]
        eqSecurityRequirement(t, sr1, sr2)
    }
    eqInt(t, len(o1.Servers), len(o2.Servers))
    for i, s1 := range o1.Servers {
        s2 := o2.Servers[i]
        eqServer(t, s1, s2)
    }
}
