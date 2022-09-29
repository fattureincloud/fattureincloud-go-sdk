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

var CompanyInfoPlanInfoJsonStr string = "{\"limits\":{\"clients\":5000,\"suppliers\":5000,\"products\":5000,\"documents\":3000},\"functions\":{\"archive\":true,\"cerved\":true,\"document_attachments\":true,\"e_invoice\":true,\"genius\":true,\"mail_tracking\":true,\"payment_notifications\":true,\"paypal\":true,\"receipts\":true,\"recurring\":true,\"smtp\":true,\"sofort\":false,\"stock\":true,\"subaccounts\":true,\"tessera_sanitaria\":true,\"ts_digital\":true,\"ts_invoice_trading\":true,\"ts_pay\":true},\"functions_status\":{\"ts_digital\":{\"active\":true},\"ts_pay\":{\"active\":true}}}"

func TestCompanyInfoPlanInfo(t *testing.T) {
	obj := NewCompanyInfoPlanInfo()
	json.Unmarshal([]byte(CompanyInfoPlanInfoJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CompanyInfoPlanInfoJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCompanyInfoPlanInfo()
	newObj.SetLimits(obj.GetLimits())
	assert.True(t, reflect.DeepEqual(obj.GetLimits(), newObj.GetLimits()))
	newObj.SetFunctions(obj.GetFunctions())
	assert.True(t, reflect.DeepEqual(obj.GetFunctions(), newObj.GetFunctions()))
	newObj.SetFunctionsStatus(obj.GetFunctionsStatus())
	assert.True(t, reflect.DeepEqual(obj.GetFunctionsStatus(), newObj.GetFunctionsStatus()))
}
