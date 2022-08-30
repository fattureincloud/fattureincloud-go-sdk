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

var ReceivedDocumentPaymentsListItemJsonStr string = "{\"id\":1,\"amount\":10,\"due_date\":\"2021-12-25\",\"paid_date\":\"2021-12-25\",\"payment_terms\":{\"days\":10},\"status\":\"oks\",\"payment_account\":{\"id\":1,\"type\":\"standard\"}}"

func TestReceivedDocumentPaymentsListItem(t *testing.T) {
	obj := NewReceivedDocumentPaymentsListItem()
	json.Unmarshal([]byte(ReceivedDocumentPaymentsListItemJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ReceivedDocumentPaymentsListItemJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewReceivedDocumentPaymentsListItem()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetAmount(obj.GetAmount())
	assert.True(t, reflect.DeepEqual(obj.GetAmount(), newObj.GetAmount()))
	newObj.SetDueDate(obj.GetDueDate())
	assert.True(t, reflect.DeepEqual(obj.GetDueDate(), newObj.GetDueDate()))
	newObj.SetPaidDate(obj.GetPaidDate())
	assert.True(t, reflect.DeepEqual(obj.GetPaidDate(), newObj.GetPaidDate()))
	newObj.SetPaymentTerms(obj.GetPaymentTerms())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentTerms(), newObj.GetPaymentTerms()))
	newObj.SetStatus(obj.GetStatus())
	assert.True(t, reflect.DeepEqual(obj.GetStatus(), newObj.GetStatus()))
	newObj.SetPaymentAccount(obj.GetPaymentAccount())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccount(), newObj.GetPaymentAccount()))
}
