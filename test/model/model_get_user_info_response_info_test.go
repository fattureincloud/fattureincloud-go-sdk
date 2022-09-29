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

var GetUserInfoResponseInfoJsonStr string = "{\"need_marketing_consents_confirmation\":false,\"need_password_change\":true,\"need_terms_of_service_confirmation\":false}"

func TestGetUserInfoResponseInfo(t *testing.T) {
	obj := NewGetUserInfoResponseInfo()
	json.Unmarshal([]byte(GetUserInfoResponseInfoJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetUserInfoResponseInfoJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetUserInfoResponseInfo()
	newObj.SetNeedMarketingConsentsConfirmation(obj.GetNeedMarketingConsentsConfirmation())
	assert.True(t, reflect.DeepEqual(obj.GetNeedMarketingConsentsConfirmation(), newObj.GetNeedMarketingConsentsConfirmation()))
	newObj.SetNeedPasswordChange(obj.GetNeedPasswordChange())
	assert.True(t, reflect.DeepEqual(obj.GetNeedPasswordChange(), newObj.GetNeedPasswordChange()))
	newObj.SetNeedTermsOfServiceConfirmation(obj.GetNeedTermsOfServiceConfirmation())
	assert.True(t, reflect.DeepEqual(obj.GetNeedTermsOfServiceConfirmation(), newObj.GetNeedTermsOfServiceConfirmation()))
}
