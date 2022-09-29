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

var GetUserInfoResponseJsonStr string = "{\"data\":{\"id\":12345,\"name\":\"Mario Rossi\",\"first_name\":\"Mario\",\"last_name\":\"Rossi\",\"email\":\"mario.rossi@example.com\",\"hash\":\"5add29e1234532a1bf2ed7b612043029\",\"picture\":\"picture.jpg\"},\"info\":{\"need_marketing_consents_confirmation\":false,\"need_password_change\":false,\"need_terms_of_service_confirmation\":false},\"email_confirmation_state\":{\"need_confirmation\":true}}"

func TestGetUserInfoResponse(t *testing.T) {
	obj := NewGetUserInfoResponse()
	json.Unmarshal([]byte(GetUserInfoResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetUserInfoResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetUserInfoResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
	newObj.SetInfo(obj.GetInfo())
	assert.True(t, reflect.DeepEqual(obj.GetInfo(), newObj.GetInfo()))
	newObj.SetEmailConfirmationState(obj.GetEmailConfirmationState())
	assert.True(t, reflect.DeepEqual(obj.GetEmailConfirmationState(), newObj.GetEmailConfirmationState()))
}
