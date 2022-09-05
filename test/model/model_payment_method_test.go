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

var PaymentMethodJsonStr string = "{\"id\":12346,\"name\":\"Bonifico bancario\",\"type\":\"standard\",\"is_default\":true,\"default_payment_account\":{\"id\":21,\"name\":\"n1\",\"type\":\"standard\"},\"details\":[{\"title\":\"t1\"}],\"bank_iban\":\"IT62W0300203280486429468578\",\"bank_name\":\"Indesa\",\"bank_beneficiary\":\"mamma\",\"ei_payment_method\":\"2\"}"

func TestPaymentMethod(t *testing.T) {
	obj := NewPaymentMethod()
	json.Unmarshal([]byte(PaymentMethodJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, PaymentMethodJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewPaymentMethod()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetIsDefault(obj.GetIsDefault())
	assert.True(t, reflect.DeepEqual(obj.GetIsDefault(), newObj.GetIsDefault()))
	newObj.SetDefaultPaymentAccount(obj.GetDefaultPaymentAccount())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultPaymentAccount(), newObj.GetDefaultPaymentAccount()))
	newObj.SetDetails(obj.GetDetails())
	assert.True(t, reflect.DeepEqual(obj.GetDetails(), newObj.GetDetails()))
	newObj.SetBankIban(obj.GetBankIban())
	assert.True(t, reflect.DeepEqual(obj.GetBankIban(), newObj.GetBankIban()))
	newObj.SetBankName(obj.GetBankName())
	assert.True(t, reflect.DeepEqual(obj.GetBankName(), newObj.GetBankName()))
	newObj.SetBankBeneficiary(obj.GetBankBeneficiary())
	assert.True(t, reflect.DeepEqual(obj.GetBankBeneficiary(), newObj.GetBankBeneficiary()))
	newObj.SetEiPaymentMethod(obj.GetEiPaymentMethod())
	assert.True(t, reflect.DeepEqual(obj.GetEiPaymentMethod(), newObj.GetEiPaymentMethod()))
}
