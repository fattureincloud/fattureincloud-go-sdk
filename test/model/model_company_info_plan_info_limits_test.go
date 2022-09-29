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

var CompanyInfoPlanInfoLimitsJsonStr string = "{\"clients\":5000,\"suppliers\":5000,\"products\":5000,\"documents\":3000}"

func TestCompanyInfoPlanInfoLimits(t *testing.T) {
	obj := NewCompanyInfoPlanInfoLimits()
	json.Unmarshal([]byte(CompanyInfoPlanInfoLimitsJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CompanyInfoPlanInfoLimitsJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCompanyInfoPlanInfoLimits()
	newObj.SetClients(obj.GetClients())
	assert.True(t, reflect.DeepEqual(obj.GetClients(), newObj.GetClients()))
	newObj.SetSuppliers(obj.GetSuppliers())
	assert.True(t, reflect.DeepEqual(obj.GetSuppliers(), newObj.GetSuppliers()))
	newObj.SetProducts(obj.GetProducts())
	assert.True(t, reflect.DeepEqual(obj.GetProducts(), newObj.GetProducts()))
	newObj.SetDocuments(obj.GetDocuments())
	assert.True(t, reflect.DeepEqual(obj.GetDocuments(), newObj.GetDocuments()))
}
