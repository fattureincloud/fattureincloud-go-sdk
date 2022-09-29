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

var EmailScheduleIncludeJsonStr string = "{\"document\":true,\"delivery_note\":false,\"attachment\":true,\"accompanying_invoice\":false}"

func TestEmailScheduleInclude(t *testing.T) {
	obj := NewEmailScheduleInclude()
	json.Unmarshal([]byte(EmailScheduleIncludeJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, EmailScheduleIncludeJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewEmailScheduleInclude()
	newObj.SetDocument(obj.GetDocument())
	assert.True(t, reflect.DeepEqual(obj.GetDocument(), newObj.GetDocument()))
	newObj.SetDeliveryNote(obj.GetDeliveryNote())
	assert.True(t, reflect.DeepEqual(obj.GetDeliveryNote(), newObj.GetDeliveryNote()))
	newObj.SetAttachment(obj.GetAttachment())
	assert.True(t, reflect.DeepEqual(obj.GetAttachment(), newObj.GetAttachment()))
	newObj.SetAccompanyingInvoice(obj.GetAccompanyingInvoice())
	assert.True(t, reflect.DeepEqual(obj.GetAccompanyingInvoice(), newObj.GetAccompanyingInvoice()))
}
