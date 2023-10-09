/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package model

import (
	"testing"

	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestEventType(t *testing.T) {
	assert.Equal(t, "it.fattureincloud.webhooks.issued_documents.invoices.create", string(EventTypes.ISSUED_DOCUMENTS_INVOICES_CREATE))
	assert.Equal(t, "it.fattureincloud.webhooks.cashbook.create", string(EventTypes.CASHBOOK_CREATE))
	assert.Equal(t, "it.fattureincloud.webhooks.entities.all.update", string(EventTypes.ENTITIES_ALL_UPDATE))
	assert.Equal(t, "it.fattureincloud.webhooks.products.delete", string(EventTypes.PRODUCTS_DELETE))
}
