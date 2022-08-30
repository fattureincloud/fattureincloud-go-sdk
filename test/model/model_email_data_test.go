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

var EmailDataJsonStr string = "{\"recipient_email\":\"m.rossi@exxample.com\",\"default_sender_email\":{\"id\":0,\"email\":\"no-reply@fattureincloud.it\"},\"sender_emails_list\":[{\"id\":0,\"email\":\"no-reply@fattureincloud.it\"},{\"id\":888,\"email\":\"mariorossi@fattureincloud.it\"}],\"subject\":\"Nostra pro forma nr. 1\",\"body\":\"Pro forma body\",\"document_exists\":true,\"delivery_note_exists\":false,\"attachment_exists\":false,\"accompanying_invoice_exists\":false,\"default_attach_pdf\":false}"

func TestEmailData(t *testing.T) {
	obj := NewEmailData()
	json.Unmarshal([]byte(EmailDataJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, EmailDataJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewEmailData()
	newObj.SetRecipientEmail(obj.GetRecipientEmail())
	assert.True(t, reflect.DeepEqual(obj.GetRecipientEmail(), newObj.GetRecipientEmail()))
	newObj.SetDefaultSenderEmail(obj.GetDefaultSenderEmail())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultSenderEmail(), newObj.GetDefaultSenderEmail()))
	newObj.SetSenderEmailsList(obj.GetSenderEmailsList())
	assert.True(t, reflect.DeepEqual(obj.GetSenderEmailsList(), newObj.GetSenderEmailsList()))
	newObj.SetCcEmail(obj.GetCcEmail())
	assert.True(t, reflect.DeepEqual(obj.GetCcEmail(), newObj.GetCcEmail()))
	newObj.SetSubject(obj.GetSubject())
	assert.True(t, reflect.DeepEqual(obj.GetSubject(), newObj.GetSubject()))
	newObj.SetBody(obj.GetBody())
	assert.True(t, reflect.DeepEqual(obj.GetBody(), newObj.GetBody()))
	newObj.SetDocumentExists(obj.GetDocumentExists())
	assert.True(t, reflect.DeepEqual(obj.GetDocumentExists(), newObj.GetDocumentExists()))
	newObj.SetDeliveryNoteExists(obj.GetDeliveryNoteExists())
	assert.True(t, reflect.DeepEqual(obj.GetDeliveryNoteExists(), newObj.GetDeliveryNoteExists()))
	newObj.SetAttachmentExists(obj.GetAttachmentExists())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentExists(), newObj.GetAttachmentExists()))
	newObj.SetAccompanyingInvoiceExists(obj.GetAccompanyingInvoiceExists())
	assert.True(t, reflect.DeepEqual(obj.GetAccompanyingInvoiceExists(), newObj.GetAccompanyingInvoiceExists()))
	newObj.SetDefaultAttachPdf(obj.GetDefaultAttachPdf())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultAttachPdf(), newObj.GetDefaultAttachPdf()))
}
