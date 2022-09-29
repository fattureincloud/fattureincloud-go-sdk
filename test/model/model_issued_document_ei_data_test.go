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

var IssuedDocumentEiDataJsonStr string = "{\"vat_kind\":\"D\",\"original_document_type\":\"contratto\",\"od_number\":\"odnum\",\"od_date\":\"2021-12-25\",\"cig\":\"string\",\"cup\":\"string\",\"payment_method\":\"EC21\",\"bank_name\":\"nome bancala\",\"bank_iban\":\"ITP00001231248137y41\",\"bank_beneficiary\":\"bank beneficiary\",\"invoice_number\":\"123\",\"invoice_date\":\"2021-12-25\"}"

func TestIssuedDocumentEiData(t *testing.T) {
	obj := NewIssuedDocumentEiData()
	json.Unmarshal([]byte(IssuedDocumentEiDataJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentEiDataJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocumentEiData()
	newObj.SetVatKind(obj.GetVatKind())
	assert.True(t, reflect.DeepEqual(obj.GetVatKind(), newObj.GetVatKind()))
	newObj.SetOriginalDocumentType(obj.GetOriginalDocumentType())
	assert.True(t, reflect.DeepEqual(obj.GetOriginalDocumentType(), newObj.GetOriginalDocumentType()))
	newObj.SetOdNumber(obj.GetOdNumber())
	assert.True(t, reflect.DeepEqual(obj.GetOdNumber(), newObj.GetOdNumber()))
	newObj.SetOdDate(obj.GetOdDate())
	assert.True(t, reflect.DeepEqual(obj.GetOdDate(), newObj.GetOdDate()))
	newObj.SetCig(obj.GetCig())
	assert.True(t, reflect.DeepEqual(obj.GetCig(), newObj.GetCig()))
	newObj.SetCup(obj.GetCup())
	assert.True(t, reflect.DeepEqual(obj.GetCup(), newObj.GetCup()))
	newObj.SetPaymentMethod(obj.GetPaymentMethod())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentMethod(), newObj.GetPaymentMethod()))
	newObj.SetBankName(obj.GetBankName())
	assert.True(t, reflect.DeepEqual(obj.GetBankName(), newObj.GetBankName()))
	newObj.SetBankIban(obj.GetBankIban())
	assert.True(t, reflect.DeepEqual(obj.GetBankIban(), newObj.GetBankIban()))
	newObj.SetBankBeneficiary(obj.GetBankBeneficiary())
	assert.True(t, reflect.DeepEqual(obj.GetBankBeneficiary(), newObj.GetBankBeneficiary()))
	newObj.SetInvoiceNumber(obj.GetInvoiceNumber())
	assert.True(t, reflect.DeepEqual(obj.GetInvoiceNumber(), newObj.GetInvoiceNumber()))
	newObj.SetInvoiceDate(obj.GetInvoiceDate())
	assert.True(t, reflect.DeepEqual(obj.GetInvoiceDate(), newObj.GetInvoiceDate()))
}
