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

var ListSuppliersResponsePageJsonStr string = "{\"data\":[{\"id\":12345,\"code\":\"AE86\",\"name\":\"Mario Rossi S.R.L.\",\"type\":\"company\",\"first_name\":\"Mario\",\"last_name\":\"Rossi\",\"contact_person\":\"\",\"vat_number\":\"111222333\",\"tax_code\":\"111122233\",\"address_street\":\"Corso Magellano, 46\",\"address_postal_code\":\"20146\",\"address_city\":\"Milano\",\"address_province\":\"MI\",\"address_extra\":\"\",\"country\":\"Italia\",\"email\":\"mario.rossi@example.com\",\"certified_email\":\"mario.rossi@pec.example.com\",\"phone\":\"1234567890\",\"fax\":\"123456789\",\"notes\":\"\",\"created_at\":\"2021-15-08\",\"updated_at\":\"2021-15-08\"}]}"

func TestListSuppliersResponsePage(t *testing.T) {
	obj := NewListSuppliersResponsePage()
	json.Unmarshal([]byte(ListSuppliersResponsePageJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListSuppliersResponsePageJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListSuppliersResponsePage()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
