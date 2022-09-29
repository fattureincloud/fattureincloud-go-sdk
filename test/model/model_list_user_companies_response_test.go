/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package model

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

var ListUserCompaniesResponseJsonStr string = "{\"data\":{\"companies\":[{\"id\":12345,\"name\":\"Studio Commercialista\",\"type\":\"accountant\",\"access_token\":\"fc3c89ba29a926f9ef20203577c00ada\",\"controlled_companies\":[],\"connection_id\":94566,\"tax_code\":\"ABCSFN94T17A794K\"}]}}"

func TestListUserCompaniesResponse(t *testing.T) {
	obj := NewListUserCompaniesResponse()
	json.Unmarshal([]byte(ListUserCompaniesResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListUserCompaniesResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListUserCompaniesResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
