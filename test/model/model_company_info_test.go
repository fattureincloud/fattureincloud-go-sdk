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

	. "github.com/fattureincloud/fattureincloud-go-sdk/model"
	"github.com/stretchr/testify/assert"
)

var CompanyInfoJsonStr string = "{\"id\":12346,\"name\":\"Studio Commercialista\",\"email\":\"mario.rossi@example.com\",\"type\":\"accountant\",\"access_info\":{\"role\":\"master\",\"through_accountant\":false},\"plan_info\":{\"limits\":{\"clients\":5000,\"suppliers\":5000,\"products\":5000,\"documents\":3000},\"functions\":{\"archive\":true,\"cerved\":true,\"document_attachments\":true,\"e_invoice\":true,\"genius\":true,\"mail_tracking\":true,\"payment_notifications\":true,\"paypal\":true,\"receipts\":true,\"recurring\":true,\"smtp\":true,\"sofort\":false,\"stock\":true,\"subaccounts\":true,\"tessera_sanitaria\":true,\"ts_digital\":true,\"ts_invoice_trading\":true,\"ts_pay\":true},\"functions_status\":{\"ts_digital\":{\"active\":true},\"ts_pay\":{\"active\":false}}},\"accountant_id\":12345,\"is_accountant\":true}"

func TestCompanyInfo(t *testing.T) {
	obj := NewCompanyInfo()
	json.Unmarshal([]byte(CompanyInfoJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CompanyInfoJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCompanyInfo()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetEmail(obj.GetEmail())
	assert.True(t, reflect.DeepEqual(obj.GetEmail(), newObj.GetEmail()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetAccessInfo(obj.GetAccessInfo())
	assert.True(t, reflect.DeepEqual(obj.GetAccessInfo(), newObj.GetAccessInfo()))
	newObj.SetPlanInfo(obj.GetPlanInfo())
	assert.True(t, reflect.DeepEqual(obj.GetPlanInfo(), newObj.GetPlanInfo()))
	newObj.SetAccountantId(obj.GetAccountantId())
	assert.True(t, reflect.DeepEqual(obj.GetAccountantId(), newObj.GetAccountantId()))
	newObj.SetIsAccountant(obj.GetIsAccountant())
	assert.True(t, reflect.DeepEqual(obj.GetIsAccountant(), newObj.GetIsAccountant()))
}
