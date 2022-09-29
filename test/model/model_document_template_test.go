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

var DocumentTemplateJsonStr string = "{\"id\":10,\"name\":\"New Standard S1\",\"type\":\"Tipo 1\"}"

func TestDocumentTemplate(t *testing.T) {
	obj := NewDocumentTemplate()
	json.Unmarshal([]byte(DocumentTemplateJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, DocumentTemplateJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewDocumentTemplate()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
}
