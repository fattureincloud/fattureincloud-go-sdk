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

var IssuedDocumentPaymentsListItemJsonStr string = "{\"due_date\":\"2021-12-25\",\"amount\":10,\"status\":\"paid\",\"payment_account\":{\"id\":1,\"type\":\"standard\"},\"paid_date\":\"2021-12-25\"}"

func TestIssuedDocumentPaymentsListItem(t *testing.T) {
	obj := NewIssuedDocumentPaymentsListItem()
	json.Unmarshal([]byte(IssuedDocumentPaymentsListItemJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentPaymentsListItemJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocumentPaymentsListItem()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetDueDate(obj.GetDueDate())
	assert.True(t, reflect.DeepEqual(obj.GetDueDate(), newObj.GetDueDate()))
	newObj.SetAmount(obj.GetAmount())
	assert.True(t, reflect.DeepEqual(obj.GetAmount(), newObj.GetAmount()))
	newObj.SetStatus(obj.GetStatus())
	assert.True(t, reflect.DeepEqual(obj.GetStatus(), newObj.GetStatus()))
	newObj.SetPaymentAccount(obj.GetPaymentAccount())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccount(), newObj.GetPaymentAccount()))
	newObj.SetPaidDate(obj.GetPaidDate())
	assert.True(t, reflect.DeepEqual(obj.GetPaidDate(), newObj.GetPaidDate()))
	newObj.SetEiRaw(obj.GetEiRaw())
	assert.True(t, reflect.DeepEqual(obj.GetEiRaw(), newObj.GetEiRaw()))
	newObj.SetPaymentTerms(obj.GetPaymentTerms())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentTerms(), newObj.GetPaymentTerms()))
}
