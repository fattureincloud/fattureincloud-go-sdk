/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package model

import (
	"testing"

	. "github.com/fattureincloud/fattureincloud-go-sdk/model"
	"github.com/stretchr/testify/assert"
)

func TestIssuedDocumentTypeResponse(t *testing.T) {
	assert.Equal(t, "invoice", string(IssuedDocumentTypes.INVOICE))
	assert.Equal(t, "quote", string(IssuedDocumentTypes.QUOTE))
	assert.Equal(t, "proforma", string(IssuedDocumentTypes.PROFORMA))
	assert.Equal(t, "receipt", string(IssuedDocumentTypes.RECEIPT))
	assert.Equal(t, "delivery_note", string(IssuedDocumentTypes.DELIVERY_NOTE))
	assert.Equal(t, "credit_note", string(IssuedDocumentTypes.CREDIT_NOTE))
	assert.Equal(t, "order", string(IssuedDocumentTypes.ORDER))
	assert.Equal(t, "work_report", string(IssuedDocumentTypes.WORK_REPORT))
	assert.Equal(t, "supplier_order", string(IssuedDocumentTypes.SUPPLIER_ORDER))
	assert.Equal(t, "self_own_invoice", string(IssuedDocumentTypes.SELF_OWN_INVOICE))
	assert.Equal(t, "self_supplier_invoice", string(IssuedDocumentTypes.SELF_SUPPLIER_INVOICE))
}
