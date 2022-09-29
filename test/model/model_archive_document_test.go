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

var ArchiveDocumentJsonStr string = "{\"id\":12345,\"date\":\"2021-08-20\",\"description\":\"spesa 2\",\"category\":\"Altri documenti\",\"attachment_token\":\"jwfbaiuwbfoiewfoa8weohafw7gefa9we\"}"

func TestArchiveDocument(t *testing.T) {
	obj := NewArchiveDocument()
	json.Unmarshal([]byte(ArchiveDocumentJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ArchiveDocumentJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewArchiveDocument()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetDate(obj.GetDate())
	assert.True(t, reflect.DeepEqual(obj.GetDate(), newObj.GetDate()))
	newObj.SetDescription(obj.GetDescription())
	assert.True(t, reflect.DeepEqual(obj.GetDescription(), newObj.GetDescription()))
	newObj.SetAttachmentUrl(obj.GetAttachmentUrl())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentUrl(), newObj.GetAttachmentUrl()))
	newObj.SetCategory(obj.GetCategory())
	assert.True(t, reflect.DeepEqual(obj.GetCategory(), newObj.GetCategory()))
	newObj.SetAttachmentToken(obj.GetAttachmentToken())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentToken(), newObj.GetAttachmentToken()))
}
