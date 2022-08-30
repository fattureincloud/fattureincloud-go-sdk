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

var VatTypeJsonStr string = "{\"id\":0,\"value\":22,\"description\":\"Non imponibile art. 123\",\"notes\":\"IVA non imponibile\",\"e_invoice\":true,\"ei_type\":\"2\",\"ei_description\":\"desc\",\"is_disabled\":false}"

func TestVatType(t *testing.T) {
	obj := NewVatType()
	json.Unmarshal([]byte(VatTypeJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, VatTypeJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewVatType()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetValue(obj.GetValue())
	assert.True(t, reflect.DeepEqual(obj.GetValue(), newObj.GetValue()))
	newObj.SetDescription(obj.GetDescription())
	assert.True(t, reflect.DeepEqual(obj.GetDescription(), newObj.GetDescription()))
	newObj.SetNotes(obj.GetNotes())
	assert.True(t, reflect.DeepEqual(obj.GetNotes(), newObj.GetNotes()))
	newObj.SetEInvoice(obj.GetEInvoice())
	assert.True(t, reflect.DeepEqual(obj.GetEInvoice(), newObj.GetEInvoice()))
	newObj.SetEiType(obj.GetEiType())
	assert.True(t, reflect.DeepEqual(obj.GetEiType(), newObj.GetEiType()))
	newObj.SetEiDescription(obj.GetEiDescription())
	assert.True(t, reflect.DeepEqual(obj.GetEiDescription(), newObj.GetEiDescription()))
	newObj.SetEditable(obj.GetEditable())
	assert.True(t, reflect.DeepEqual(obj.GetEditable(), newObj.GetEditable()))
	newObj.SetIsDisabled(obj.GetIsDisabled())
	assert.True(t, reflect.DeepEqual(obj.GetIsDisabled(), newObj.GetIsDisabled()))
}
