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

var UserJsonStr string = "{\"id\":12345,\"name\":\"Mario Rossi\",\"first_name\":\"Mario\",\"last_name\":\"Rossi\",\"email\":\"mario.rossi@example.com\",\"hash\":\"5add29e1234532a1bf2ed7b612043029\",\"picture\":\"picture.jpg\"}"

func TestUser(t *testing.T) {
	obj := NewUser()
	json.Unmarshal([]byte(UserJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, UserJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewUser()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetFirstName(obj.GetFirstName())
	assert.True(t, reflect.DeepEqual(obj.GetFirstName(), newObj.GetFirstName()))
	newObj.SetLastName(obj.GetLastName())
	assert.True(t, reflect.DeepEqual(obj.GetLastName(), newObj.GetLastName()))
	newObj.SetEmail(obj.GetEmail())
	assert.True(t, reflect.DeepEqual(obj.GetEmail(), newObj.GetEmail()))
	newObj.SetHash(obj.GetHash())
	assert.True(t, reflect.DeepEqual(obj.GetHash(), newObj.GetHash()))
	newObj.SetPicture(obj.GetPicture())
	assert.True(t, reflect.DeepEqual(obj.GetPicture(), newObj.GetPicture()))
}
