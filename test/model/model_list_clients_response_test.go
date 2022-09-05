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

var ListClientsResponseJsonStr string = "{\"current_page\":10,\"first_page_url\":\"https://www.page.url/\",\"from\":10,\"last_page\":10,\"last_page_url\":\"https://www.page.url/\",\"next_page_url\":\"https://www.page.url/\",\"path\":\"https://www.page.url/\",\"per_page\":10,\"prev_page_url\":\"https://www.page.url/\",\"to\":10,\"total\":10,\"data\":[{\"id\":12345,\"code\":\"AE86\",\"name\":\"Avv. Maria Rossi\",\"type\":\"person\",\"first_name\":\"Maria\",\"last_name\":\"Rossi\",\"contact_person\":\"\",\"vat_number\":\"IT12345640962\",\"tax_code\":\"BLTGNI5ABCDA794E\",\"address_street\":\"Via Roma, 1\",\"address_postal_code\":\"20900\",\"address_city\":\"Milano\",\"address_province\":\"MI\",\"address_extra\":\"\",\"country\":\"Italia\",\"email\":\"maria.rossi@example.com\",\"certified_email\":\"maria.rossi@pec.example.com\",\"phone\":\"1234567890\",\"fax\":\"\",\"notes\":\"\",\"default_vat\":{\"id\":54321,\"value\":45,\"description\":\"\",\"is_disabled\":false},\"default_payment_terms\":1,\"default_payment_terms_type\":\"standard\",\"default_payment_method\":{\"id\":386092,\"name\":\"Credit card\",\"type\":\"standard\"},\"bank_name\":\"Indesa\",\"bank_iban\":\"IT40P123456781000000123456\",\"bank_swift_code\":\"AK86PCT\",\"shipping_address\":\"Corso Magellano 4\",\"e_invoice\":true,\"ei_code\":\"111111\",\"created_at\":\"2021-04-29 08:53:07\",\"updated_at\":\"2021-04-29 08:53:07\"}]}"

func TestListClientsResponse(t *testing.T) {
	obj := NewListClientsResponse()
	json.Unmarshal([]byte(ListClientsResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListClientsResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListClientsResponse()
	newObj.SetCurrentPage(obj.GetCurrentPage())
	assert.True(t, reflect.DeepEqual(obj.GetCurrentPage(), newObj.GetCurrentPage()))
	newObj.SetFirstPageUrl(obj.GetFirstPageUrl())
	assert.True(t, reflect.DeepEqual(obj.GetFirstPageUrl(), newObj.GetFirstPageUrl()))
	newObj.SetFrom(obj.GetFrom())
	assert.True(t, reflect.DeepEqual(obj.GetFrom(), newObj.GetFrom()))
	newObj.SetLastPage(obj.GetLastPage())
	assert.True(t, reflect.DeepEqual(obj.GetLastPage(), newObj.GetLastPage()))
	newObj.SetLastPageUrl(obj.GetLastPageUrl())
	assert.True(t, reflect.DeepEqual(obj.GetLastPageUrl(), newObj.GetLastPageUrl()))
	newObj.SetNextPageUrl(obj.GetNextPageUrl())
	assert.True(t, reflect.DeepEqual(obj.GetNextPageUrl(), newObj.GetNextPageUrl()))
	newObj.SetPath(obj.GetPath())
	assert.True(t, reflect.DeepEqual(obj.GetPath(), newObj.GetPath()))
	newObj.SetPerPage(obj.GetPerPage())
	assert.True(t, reflect.DeepEqual(obj.GetPerPage(), newObj.GetPerPage()))
	newObj.SetPrevPageUrl(obj.GetPrevPageUrl())
	assert.True(t, reflect.DeepEqual(obj.GetPrevPageUrl(), newObj.GetPrevPageUrl()))
	newObj.SetTo(obj.GetTo())
	assert.True(t, reflect.DeepEqual(obj.GetTo(), newObj.GetTo()))
	newObj.SetTotal(obj.GetTotal())
	assert.True(t, reflect.DeepEqual(obj.GetTotal(), newObj.GetTotal()))
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
