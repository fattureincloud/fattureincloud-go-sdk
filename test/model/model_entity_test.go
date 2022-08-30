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

var EntityJsonStr string = "{\"id\":16451,\"code\":\"AE86\",\"name\":\"Avv. Maria Rossi\",\"type\":\"person\",\"first_name\":\"Maria\",\"last_name\":\"Rossi\",\"contact_person\":\"\",\"vat_number\":\"IT12345640962\",\"tax_code\":\"BLTGNI5ABCDA794E\",\"address_street\":\"Via Roma, 1\",\"address_postal_code\":\"20900\",\"address_city\":\"Milano\",\"address_province\":\"MI\",\"address_extra\":\"\",\"country\":\"Italia\",\"email\":\"maria.rossi@example.com\",\"certified_email\":\"maria.rossi@pec.example.com\",\"phone\":\"1234567890\",\"fax\":\"\",\"notes\":\"\",\"default_vat\":{\"id\":54321,\"value\":45,\"description\":\"\",\"is_disabled\":false},\"default_payment_terms\":1,\"default_payment_terms_type\":\"standard\",\"default_payment_method\":{\"id\":386092,\"name\":\"Credit card\",\"type\":\"standard\"},\"bank_name\":\"Indesa\",\"bank_iban\":\"IT40P123456781000000123456\",\"bank_swift_code\":\"AK86PCT\",\"shipping_address\":\"Corso Magellano 4\",\"e_invoice\":true,\"ei_code\":\"111111\",\"created_at\":\"2021-04-29 08:53:07\",\"updated_at\":\"2021-04-29 08:53:07\"}"

func TestEntity(t *testing.T) {
	obj := NewEntity()
	json.Unmarshal([]byte(EntityJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, EntityJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewEntity()
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
	newObj.SetDefaultVat(obj.GetDefaultVat())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultVat(), newObj.GetDefaultVat()))
	newObj.SetDefaultPaymentTerms(obj.GetDefaultPaymentTerms())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultPaymentTerms(), newObj.GetDefaultPaymentTerms()))
	newObj.SetDefaultPaymentTermsType(obj.GetDefaultPaymentTermsType())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultPaymentTermsType(), newObj.GetDefaultPaymentTermsType()))
	newObj.SetDefaultPaymentMethod(obj.GetDefaultPaymentMethod())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultPaymentMethod(), newObj.GetDefaultPaymentMethod()))
	newObj.SetBankName(obj.GetBankName())
	assert.True(t, reflect.DeepEqual(obj.GetBankName(), newObj.GetBankName()))
	newObj.SetBankIban(obj.GetBankIban())
	assert.True(t, reflect.DeepEqual(obj.GetBankIban(), newObj.GetBankIban()))
	newObj.SetBankSwiftCode(obj.GetBankSwiftCode())
	assert.True(t, reflect.DeepEqual(obj.GetBankSwiftCode(), newObj.GetBankSwiftCode()))
	newObj.SetShippingAddress(obj.GetShippingAddress())
	assert.True(t, reflect.DeepEqual(obj.GetShippingAddress(), newObj.GetShippingAddress()))
	newObj.SetEInvoice(obj.GetEInvoice())
	assert.True(t, reflect.DeepEqual(obj.GetEInvoice(), newObj.GetEInvoice()))
	newObj.SetEiCode(obj.GetEiCode())
	assert.True(t, reflect.DeepEqual(obj.GetEiCode(), newObj.GetEiCode()))
	newObj.SetHasIntentDeclaration(obj.GetHasIntentDeclaration())
	assert.True(t, reflect.DeepEqual(obj.GetHasIntentDeclaration(), newObj.GetHasIntentDeclaration()))
	newObj.SetIntentDeclarationProtocolNumber(obj.GetIntentDeclarationProtocolNumber())
	assert.True(t, reflect.DeepEqual(obj.GetIntentDeclarationProtocolNumber(), newObj.GetIntentDeclarationProtocolNumber()))
	newObj.SetIntentDeclarationProtocolDate(obj.GetIntentDeclarationProtocolDate())
	assert.True(t, reflect.DeepEqual(obj.GetIntentDeclarationProtocolDate(), newObj.GetIntentDeclarationProtocolDate()))
	newObj.SetCreatedAt(obj.GetCreatedAt())
	assert.True(t, reflect.DeepEqual(obj.GetCreatedAt(), newObj.GetCreatedAt()))
	newObj.SetUpdatedAt(obj.GetUpdatedAt())
	assert.True(t, reflect.DeepEqual(obj.GetUpdatedAt(), newObj.GetUpdatedAt()))
}
