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

var SupplierJsonStr string = "{\"id\":12345,\"code\":\"AE86\",\"name\":\"Mario Rossi S.R.L.\",\"type\":\"company\",\"first_name\":\"Mario\",\"last_name\":\"Rossi\",\"contact_person\":\"\",\"vat_number\":\"111222333\",\"tax_code\":\"111122233\",\"address_street\":\"Corso Magellano, 46\",\"address_postal_code\":\"20146\",\"address_city\":\"Milano\",\"address_province\":\"MI\",\"address_extra\":\"\",\"country\":\"Italia\",\"email\":\"mario.rossi@example.com\",\"certified_email\":\"mario.rossi@pec.example.com\",\"phone\":\"1234567890\",\"fax\":\"123456789\",\"notes\":\"\",\"created_at\":\"2021-15-08\",\"updated_at\":\"2021-15-08\"}"

func TestSupplier(t *testing.T) {
	obj := NewSupplier()
	json.Unmarshal([]byte(SupplierJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, SupplierJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewSupplier()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetCode(obj.GetCode())
	assert.True(t, reflect.DeepEqual(obj.GetCode(), newObj.GetCode()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetFirstName(obj.GetFirstName())
	assert.True(t, reflect.DeepEqual(obj.GetFirstName(), newObj.GetFirstName()))
	newObj.SetLastName(obj.GetLastName())
	assert.True(t, reflect.DeepEqual(obj.GetLastName(), newObj.GetLastName()))
	newObj.SetContactPerson(obj.GetContactPerson())
	assert.True(t, reflect.DeepEqual(obj.GetContactPerson(), newObj.GetContactPerson()))
	newObj.SetVatNumber(obj.GetVatNumber())
	assert.True(t, reflect.DeepEqual(obj.GetVatNumber(), newObj.GetVatNumber()))
	newObj.SetTaxCode(obj.GetTaxCode())
	assert.True(t, reflect.DeepEqual(obj.GetTaxCode(), newObj.GetTaxCode()))
	newObj.SetAddressStreet(obj.GetAddressStreet())
	assert.True(t, reflect.DeepEqual(obj.GetAddressStreet(), newObj.GetAddressStreet()))
	newObj.SetAddressPostalCode(obj.GetAddressPostalCode())
	assert.True(t, reflect.DeepEqual(obj.GetAddressPostalCode(), newObj.GetAddressPostalCode()))
	newObj.SetAddressCity(obj.GetAddressCity())
	assert.True(t, reflect.DeepEqual(obj.GetAddressCity(), newObj.GetAddressCity()))
	newObj.SetAddressProvince(obj.GetAddressProvince())
	assert.True(t, reflect.DeepEqual(obj.GetAddressProvince(), newObj.GetAddressProvince()))
	newObj.SetAddressExtra(obj.GetAddressExtra())
	assert.True(t, reflect.DeepEqual(obj.GetAddressExtra(), newObj.GetAddressExtra()))
	newObj.SetCountry(obj.GetCountry())
	assert.True(t, reflect.DeepEqual(obj.GetCountry(), newObj.GetCountry()))
	newObj.SetEmail(obj.GetEmail())
	assert.True(t, reflect.DeepEqual(obj.GetEmail(), newObj.GetEmail()))
	newObj.SetCertifiedEmail(obj.GetCertifiedEmail())
	assert.True(t, reflect.DeepEqual(obj.GetCertifiedEmail(), newObj.GetCertifiedEmail()))
	newObj.SetPhone(obj.GetPhone())
	assert.True(t, reflect.DeepEqual(obj.GetPhone(), newObj.GetPhone()))
	newObj.SetFax(obj.GetFax())
	assert.True(t, reflect.DeepEqual(obj.GetFax(), newObj.GetFax()))
	newObj.SetNotes(obj.GetNotes())
	assert.True(t, reflect.DeepEqual(obj.GetNotes(), newObj.GetNotes()))
	newObj.SetBankIban(obj.GetBankIban())
	assert.True(t, reflect.DeepEqual(obj.GetBankIban(), newObj.GetBankIban()))
	newObj.SetCreatedAt(obj.GetCreatedAt())
	assert.True(t, reflect.DeepEqual(obj.GetCreatedAt(), newObj.GetCreatedAt()))
	newObj.SetUpdatedAt(obj.GetUpdatedAt())
	assert.True(t, reflect.DeepEqual(obj.GetUpdatedAt(), newObj.GetUpdatedAt()))
}
