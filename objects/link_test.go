package objects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/oas3/spec/objects"
	"gopkg.in/yaml.v2"
)

var linkObjs = []objects.Link{
	{
		OperationID: "getUserAddress",
		Parameters: map[string]interface{}{
			"userId": "$request.path.id",
		},
	},
	{
		OperationID: "getUserAddressByUUID",
		Parameters: map[string]interface{}{
			"userUuid": "$response.body#/uuid",
		},
	},
	{
		OperationRef: "#/paths/~12.0~1repositories~1{username}/get",
		Parameters: map[string]interface{}{
			"username": "$response.body#/username",
		},
	},
	{
		OperationRef: "https://na2.gigantic-server.com/#/paths/~12.0~1repositories~1{username}/get",
		Parameters: map[string]interface{}{
			"username": "$response.body#/username",
		},
	},
}

func TestLink(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/links.json")
		var links []objects.Link
		if err := json.Unmarshal(raw, &links); err != nil {
			t.Error(err)
		}
		eqLinks(t, linkObjs, links)
	})

	t.Run("YAML", func(t *testing.T) {
		raw, _ := ioutil.ReadFile("testdata/links.yml")
		var links []objects.Link
		if err := yaml.Unmarshal(raw, &links); err != nil {
			t.Error(err)
		}
		eqLinks(t, linkObjs, links)
	})
}

func eqLinks(t *testing.T, ls1, ls2 []objects.Link) {
	eqInt(t, len(ls1), len(ls2))
	for i, l1 := range ls1 {
		l2 := ls2[i]
		eqLink(t, l1, l2)
	}
}

func eqLink(t *testing.T, l1, l2 objects.Link) {
	eqStr(t, l1.OperationRef, l2.OperationRef)
	eqStr(t, l1.OperationID, l2.OperationID)
	eq(t, l1.Parameters, l2.Parameters)
	eq(t, l1.RequestBody, l2.RequestBody)
	eqStr(t, l1.Description, l2.Description)
	eqServer(t, l1.Server, l2.Server)
}
