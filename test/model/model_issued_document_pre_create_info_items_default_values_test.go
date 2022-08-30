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

var IssuedDocumentPreCreateInfoItemsDefaultValuesJsonStr string = "{\"vat\":{\"id\":1}}"

func TestIssuedDocumentPreCreateInfoItemsDefaultValues(t *testing.T) {
	obj := NewIssuedDocumentPreCreateInfoItemsDefaultValues()
	json.Unmarshal([]byte(IssuedDocumentPreCreateInfoItemsDefaultValuesJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentPreCreateInfoItemsDefaultValuesJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocumentPreCreateInfoItemsDefaultValues()
	newObj.SetVat(obj.GetVat())
	assert.True(t, reflect.DeepEqual(obj.GetVat(), newObj.GetVat()))
}
