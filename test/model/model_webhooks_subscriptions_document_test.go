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

var WebhooksSubscriptionJsonStr string = "{\"id\":\"SUB123\",\"sink\":\"https://endpoint.test\",\"verified\":true,\"types\":[\"it.fattureincloud.cashbook.create\"]}"

func TestWebhooksSubscription(t *testing.T) {
	obj := NewWebhooksSubscription()
	json.Unmarshal([]byte(WebhooksSubscriptionJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, WebhooksSubscriptionJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewWebhooksSubscription()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetSink(obj.GetSink())
	assert.True(t, reflect.DeepEqual(obj.GetSink(), newObj.GetSink()))
	newObj.SetVerified(obj.GetVerified())
	assert.True(t, reflect.DeepEqual(obj.GetVerified(), newObj.GetVerified()))
	newObj.SetTypes(obj.GetTypes())
	assert.True(t, reflect.DeepEqual(obj.GetTypes(), newObj.GetTypes()))
}
