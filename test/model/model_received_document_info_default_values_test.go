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

var ReceivedDocumentInfoDefaultValuesJsonStr string = "{\"detailed\":true}"

func TestReceivedDocumentInfoDefaultValues(t *testing.T) {
	obj := NewReceivedDocumentInfoDefaultValues()
	json.Unmarshal([]byte(ReceivedDocumentInfoDefaultValuesJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ReceivedDocumentInfoDefaultValuesJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewReceivedDocumentInfoDefaultValues()
	newObj.SetDetailed(obj.GetDetailed())
	assert.True(t, reflect.DeepEqual(obj.GetDetailed(), newObj.GetDetailed()))
}
