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

var FunctionStatusJsonStr string = "{\"active\":true}"

func TestFunctionStatus(t *testing.T) {
	obj := NewFunctionStatus()
	json.Unmarshal([]byte(FunctionStatusJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, FunctionStatusJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewFunctionStatus()
	newObj.SetActive(obj.GetActive())
	assert.True(t, reflect.DeepEqual(obj.GetActive(), newObj.GetActive()))
}
