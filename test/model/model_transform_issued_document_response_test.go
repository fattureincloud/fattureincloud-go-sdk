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

var TransformIssuedDocumentResponseJsonStr string = "{\"data\":{\"id\":12345,\"type\":\"invoice\",\"show_totals\":\"all\"},\"options\":{\"create_from\":[\"82112399\"],\"transform\":true,\"keep_copy\":true}}"

func TestTransformIssuedDocumentResponse(t *testing.T) {
	obj := NewTransformIssuedDocumentResponse()
	json.Unmarshal([]byte(TransformIssuedDocumentResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, TransformIssuedDocumentResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewTransformIssuedDocumentResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
	newObj.SetOptions(obj.GetOptions())
	assert.True(t, reflect.DeepEqual(obj.GetOptions(), newObj.GetOptions()))
}
