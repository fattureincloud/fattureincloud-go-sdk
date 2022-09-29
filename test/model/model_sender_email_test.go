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

var SenderEmailJsonStr string = "{\"id\":0,\"email\":\"no-reply@fattureincloud.it\"}"

func TestSenderEmail(t *testing.T) {
	obj := NewSenderEmail()
	json.Unmarshal([]byte(SenderEmailJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, SenderEmailJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewSenderEmail()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetEmail(obj.GetEmail())
	assert.True(t, reflect.DeepEqual(obj.GetEmail(), newObj.GetEmail()))
}
