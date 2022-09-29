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

var CompanyInfoAccessInfoJsonStr string = "{\"role\":\"master\",\"permissions\":{\"fic_situation\":\"read\",\"fic_clients\":\"write\",\"fic_suppliers\":\"write\",\"fic_products\":\"write\",\"fic_issued_documents\":\"detailed\"},\"through_accountant\":false}"

func TestCompanyInfoAccessInfo(t *testing.T) {
	obj := NewCompanyInfoAccessInfo()
	json.Unmarshal([]byte(CompanyInfoAccessInfoJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CompanyInfoAccessInfoJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCompanyInfoAccessInfo()
	newObj.SetRole(obj.GetRole())
	assert.True(t, reflect.DeepEqual(obj.GetRole(), newObj.GetRole()))
	newObj.SetPermissions(obj.GetPermissions())
	assert.True(t, reflect.DeepEqual(obj.GetPermissions(), newObj.GetPermissions()))
	newObj.SetThroughAccountant(obj.GetThroughAccountant())
	assert.True(t, reflect.DeepEqual(obj.GetThroughAccountant(), newObj.GetThroughAccountant()))
}
