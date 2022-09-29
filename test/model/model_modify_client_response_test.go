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

var ModifyClientResponseJsonStr string = "{\"data\":{\"id\":12345,\"code\":\"AE86\",\"name\":\"Avv. Maria Rossi\",\"type\":\"person\",\"first_name\":\"Maria\",\"last_name\":\"Rossi\",\"contact_person\":\"\",\"vat_number\":\"IT12345640962\",\"tax_code\":\"BLTGNI5ABCDA794E\",\"address_street\":\"Via Roma, 1\",\"address_postal_code\":\"20900\",\"address_city\":\"Milano\",\"address_province\":\"MI\",\"address_extra\":\"\",\"country\":\"Italia\",\"email\":\"maria.rossi@example.com\",\"certified_email\":\"maria.rossi@pec.example.com\",\"phone\":\"1234567890\",\"fax\":\"\",\"notes\":\"\",\"default_vat\":{\"id\":54321,\"value\":45,\"description\":\"\",\"is_disabled\":false},\"default_payment_terms\":1,\"default_payment_terms_type\":\"standard\",\"default_payment_method\":{\"id\":386092,\"name\":\"Credit card\",\"type\":\"standard\"},\"bank_name\":\"Indesa\",\"bank_iban\":\"IT40P123456781000000123456\",\"bank_swift_code\":\"AK86PCT\",\"shipping_address\":\"Corso Magellano 4\",\"e_invoice\":true,\"ei_code\":\"111111\",\"created_at\":\"2021-04-29 08:53:07\",\"updated_at\":\"2021-04-29 08:53:07\"}}"

func TestModifyClientResponse(t *testing.T) {
	obj := NewModifyClientResponse()
	json.Unmarshal([]byte(ModifyClientResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ModifyClientResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewModifyClientResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
