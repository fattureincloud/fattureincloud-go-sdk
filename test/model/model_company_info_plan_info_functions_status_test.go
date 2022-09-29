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

var CompanyInfoPlanInfoFunctionsStatusJsonStr string = "{\"ts_digital\":{\"active\":true},\"ts_pay\":{\"active\":false}}"

func TestCompanyInfoPlanInfoFunctionsStatus(t *testing.T) {
	obj := NewCompanyInfoPlanInfoFunctionsStatus()
	json.Unmarshal([]byte(CompanyInfoPlanInfoFunctionsStatusJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CompanyInfoPlanInfoFunctionsStatusJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCompanyInfoPlanInfoFunctionsStatus()
	newObj.SetTsDigital(obj.GetTsDigital())
	assert.True(t, reflect.DeepEqual(obj.GetTsDigital(), newObj.GetTsDigital()))
	newObj.SetTsPay(obj.GetTsPay())
	assert.True(t, reflect.DeepEqual(obj.GetTsPay(), newObj.GetTsPay()))
}
