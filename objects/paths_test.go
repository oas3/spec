package objects_test

import (
    "encoding/json"
    "github.com/oas3/spec/objects"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "testing"
)

var (
    pathsObj = objects.Paths{
        "/pets": {
            Get: objects.Operation{
                Description: "Returns all pets from the system that the user has access to",
                Responses: map[string]objects.Response{
                    "200": {
                        Description: "A list of pets.",
                        Content: map[string]objects.MediaType{
                            "application/json": {
                                Schema: map[string]interface{}{
                                    "type": "array",
                                    "items": map[string]interface{}{
                                        "$ref": "#/components/schemas/pet",
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
    }

    pathItemObj = objects.PathItem{
        Get: objects.Operation{
            Description: "Returns pets based on ID",
            Summary:     "Find pets by ID",
            OperationID: "getPetsById",
            Responses: objects.Responses{
                "200": {
                    Description: "pet response",
                    Content: map[string]objects.MediaType{
                        "*/*": {
                            Schema: map[string]interface{}{
                                "type": "array",
                                "items": map[string]interface{}{
                                    "$ref": "#/components/schemas/Pet",
                                },
                            },
                        },
                    },
                },
                "default": {
                    Description: "error payload",
                    Content: map[string]objects.MediaType{
                        "text/html": {
                            Schema: map[string]interface{}{
                                "$ref": "#/components/schemas/ErrorModel",
                            },
                        },
                    },
                },
            },
        },
        Parameters: []objects.Parameter{
            {
                Name:        "id",
                In:          "path",
                Description: "ID of pet to use",
                Required:    true,
                Schema: map[string]interface{}{
                    "type": "array",
                    "items": map[string]interface{}{
                        "type": "string",
                    },
                },
                Style: "simple",
            },
        },
    }
)

func TestPaths(t *testing.T) {
    t.Run("JSON", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/paths.json")
        var paths objects.Paths
        if err := json.Unmarshal(raw, &paths); err != nil {
            t.Error(err)
        }
        eqPaths(t, pathsObj, paths)
    })

    t.Run("YAML", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/paths.yml")
        var paths objects.Paths
        if err := yaml.Unmarshal(raw, &paths); err != nil {
            t.Error(err)
        }
        eqPaths(t, pathsObj, paths)
    })
}

func TestPathItem(t *testing.T) {
    t.Run("JSON", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/path_item.json")
        var item objects.PathItem
        if err := json.Unmarshal(raw, &item); err != nil {
            t.Error(err)
        }
        eqPathItem(t, pathItemObj, item)
    })

    t.Run("YAML", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/path_item.yml")
        var item objects.PathItem
        if err := yaml.Unmarshal(raw, &item); err != nil {
            t.Error(err)
        }
        eqPathItem(t, pathItemObj, item)
    })
}

func eqPaths(t *testing.T, ps1, ps2 objects.Paths) {
    eqInt(t, len(ps1), len(ps2))
    for k, i1 := range ps1 {
        i2 := ps2[k]
        eqPathItem(t, i1, i2)
    }
}

func eqPathItem(t *testing.T, i1, i2 objects.PathItem) {
    eqStr(t, i1.Ref, i2.Ref)
    eqStr(t, i1.Summary, i2.Summary)
    eqStr(t, i1.Description, i2.Description)
    eqOperation(t, i1.Get, i2.Get)
    eqOperation(t, i1.Put, i2.Put)
    eqOperation(t, i1.Post, i2.Post)
    eqOperation(t, i1.Delete, i2.Delete)
    eqOperation(t, i1.Options, i2.Options)
    eqOperation(t, i1.Head, i2.Head)
    eqOperation(t, i1.Patch, i2.Patch)
    eqOperation(t, i1.Trace, i2.Trace)
    eqInt(t, len(i1.Servers), len(i2.Servers))
    for i, s1 := range i1.Servers {
        s2 := i2.Servers[i]
        eqServer(t, s1, s2)
    }
    eqInt(t, len(i1.Parameters), len(i2.Parameters))
    for i, p1 := range i1.Parameters {
        p2 := i2.Parameters[i]
        eqParameter(t, p1, p2)
    }
}
