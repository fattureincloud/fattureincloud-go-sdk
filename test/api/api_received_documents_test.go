/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateReceivedDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":1,"type":"expense","entity":{"id":1,"name":"neim"},"date":"2021-12-25","category":"cat6","description":"Ricarica","amount_net":10,"amount_vat":0,"amount_withholding_tax":0,"amount_other_withholding_tax":0,"amortization":10,"rc_center":"bg","invoice_number":"in","is_marked":false,"is_detailed":false,"e_invoice":false,"currency":{"id":"EUR","exchange_rate":"1"},"tax_deductibility":50,"vat_deductibility":100,"items_list":[{"id":1}],"payments_list":[{"amount":592,"due_date":"2021-12-25","paid_date":"2021-12-25","payment_terms":{"days":0},"status":"paid","payment_account":{"id":21}}],"attachment_token":"dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ=="}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceivedDocumentsApi.CreateReceivedDocument(context.Background(), 2).CreateReceivedDocumentRequest(*NewCreateReceivedDocumentRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceivedDocument().
		SetId(1).
		SetType(ReceivedDocumentTypes.EXPENSE).
		SetEntity(*NewEntity().
			SetId(1).
			SetName("neim")).
		SetDate("2021-12-25").
		SetCategory("cat6").
		SetDescription("Ricarica").
		SetAmountNet(10).
		SetAmortization(10).
		SetRcCenter("bg").
		SetAmountVat(0).
		SetAmountWithholdingTax(0).
		SetAmountOtherWithholdingTax(0).
		SetTaxDeductibility(50).
		SetVatDeductibility(100).
		SetIsMarked(false).
		SetInvoiceNumber("in").
		SetIsMarked(false).
		SetIsDetailed(false).
		SetEInvoice(false).
		SetAttachmentToken("dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ==").
		SetCurrency(*NewCurrency().
			SetId("EUR").
			SetExchangeRate("1")).
		SetPaymentsList([]ReceivedDocumentPaymentsListItem{
			*NewReceivedDocumentPaymentsListItem().
				SetAmount(592).
				SetDueDate("2021-12-25").
				SetPaidDate("2021-12-25").
				SetPaymentTerms(*NewReceivedDocumentPaymentsListItemPaymentTerms().
					SetDays(0)).
				SetStatus("paid").
				SetPaymentAccount(*NewPaymentAccount().
					SetId(21))}).
		SetItemsList([]ReceivedDocumentItemsListItem{*NewReceivedDocumentItemsListItem().
			SetId(1)}).
		SetIsDetailed(false)

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestDeleteReceivedDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(``))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	resp, err := apiClient.ReceivedDocumentsApi.DeleteReceivedDocument(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteReceivedDocumentAttachment(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(``))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	resp, err := apiClient.ReceivedDocumentsApi.DeleteReceivedDocumentAttachment(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetExistingReceivedDocumentTotals(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"amount_net":10,"amount_vat":10,"amount_gross":10,"amount_withholding_tax":10,"amount_other_withholding_tax":10,"amount_due":10,"payments_sum":10}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceivedDocumentsApi.GetExistingReceivedDocumentTotals(context.Background(), 2, 2).GetExistingReceivedDocumentTotalsRequest(*NewGetExistingReceivedDocumentTotalsRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceivedDocumentTotals().
		SetAmountNet(10).
		SetAmountVat(10).
		SetAmountGross(10).
		SetAmountWithholdingTax(10).
		SetAmountOtherWithholdingTax(10).
		SetAmountDue((10)).
		SetPaymentsSum((10))

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetNewReceivedDocumentTotals(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"amount_net":10,"amount_vat":10,"amount_gross":10,"amount_withholding_tax":10,"amount_other_withholding_tax":10,"amount_due":10,"payments_sum":10}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceivedDocumentsApi.GetNewReceivedDocumentTotals(context.Background(), 2).GetNewReceivedDocumentTotalsRequest(*NewGetNewReceivedDocumentTotalsRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceivedDocumentTotals().
		SetAmountNet(10).
		SetAmountVat(10).
		SetAmountGross(10).
		SetAmountWithholdingTax(10).
		SetAmountOtherWithholdingTax(10).
		SetAmountDue((10)).
		SetPaymentsSum((10))

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetReceivedDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":1,"type":"expense","entity":{"id":1,"name":"neim"},"date":"2021-12-25","category":"cat6","description":"Ricarica","amount_net":10,"amount_vat":0,"amount_withholding_tax":0,"amount_other_withholding_tax":0,"amortization":10,"rc_center":"bg","invoice_number":"in","is_marked":false,"is_detailed":false,"e_invoice":false,"currency":{"id":"EUR","exchange_rate":"1"},"tax_deductibility":50,"vat_deductibility":100,"items_list":[{"id":1}],"payments_list":[{"amount":592,"due_date":"2021-12-25","paid_date":"2021-12-25","payment_terms":{"days":0},"status":"paid","payment_account":{"id":21}}],"attachment_token":"dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ=="}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceivedDocumentsApi.GetReceivedDocument(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceivedDocument().
		SetId(1).
		SetType(ReceivedDocumentTypes.EXPENSE).
		SetEntity(*NewEntity().
			SetId(1).
			SetName("neim")).
		SetDate("2021-12-25").
		SetCategory("cat6").
		SetDescription("Ricarica").
		SetAmountNet(10).
		SetAmortization(10).
		SetRcCenter("bg").
		SetAmountVat(0).
		SetAmountWithholdingTax(0).
		SetAmountOtherWithholdingTax(0).
		SetTaxDeductibility(50).
		SetVatDeductibility(100).
		SetIsMarked(false).
		SetInvoiceNumber("in").
		SetIsMarked(false).
		SetIsDetailed(false).
		SetEInvoice(false).
		SetAttachmentToken("dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ==").
		SetCurrency(*NewCurrency().
			SetId("EUR").
			SetExchangeRate("1")).
		SetPaymentsList([]ReceivedDocumentPaymentsListItem{
			*NewReceivedDocumentPaymentsListItem().
				SetAmount(592).
				SetDueDate("2021-12-25").
				SetPaidDate("2021-12-25").
				SetPaymentTerms(*NewReceivedDocumentPaymentsListItemPaymentTerms().
					SetDays(0)).
				SetStatus("paid").
				SetPaymentAccount(*NewPaymentAccount().
					SetId(21))}).
		SetItemsList([]ReceivedDocumentItemsListItem{*NewReceivedDocumentItemsListItem().
			SetId(1)}).
		SetIsDetailed(false)

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetReceivedDocumentPreCreateInfo(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"default_values":{"detailed":false},"items_default_values":{"vat":10},"countries_list":["Italia"],"currencies_list":[{"id":"EUR"}],"categories_list":["cat6"],"payment_accounts_list":[{"id":1}],"vat_types_list":[{"id":1}]}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceivedDocumentsApi.GetReceivedDocumentPreCreateInfo(context.Background(), 2).Type_("expense").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceivedDocumentInfo().
		SetDefaultValues(*NewReceivedDocumentInfoDefaultValues().SetDetailed(false)).
		SetItemsDefaultValues(*NewReceivedDocumentInfoItemsDefaultValues().SetVat(10)).
		SetCountriesList([]string{"Italia"}).
		SetCurrenciesList([]Currency{*NewCurrency().SetId("EUR")}).
		SetCategoriesList([]string{"cat6"}).
		SetPaymentAccountsList([]PaymentAccount{*NewPaymentAccount().SetId(1)}).
		SetVatTypesList([]VatType{*NewVatType().SetId(1)})

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListReceivedDocuments(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":1,"type":"expense","entity":{"id":1,"name":"neim"},"date":"2021-12-25","category":"cat6","description":"Ricarica","amount_net":10,"amount_vat":0,"amount_withholding_tax":0,"amount_other_withholding_tax":0,"amortization":10,"rc_center":"bg","invoice_number":"in","is_marked":false,"is_detailed":false,"e_invoice":false,"currency":{"id":"EUR","exchange_rate":"1"},"tax_deductibility":50,"vat_deductibility":100,"items_list":[{"id":1}],"payments_list":[{"amount":592,"due_date":"2021-12-25","paid_date":"2021-12-25","payment_terms":{"days":0},"status":"paid","payment_account":{"id":21}}],"attachment_token":"dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ=="}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceivedDocumentsApi.ListReceivedDocuments(context.Background(), 2).Type_("expense").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceivedDocument().
		SetId(1).
		SetType(ReceivedDocumentTypes.EXPENSE).
		SetEntity(*NewEntity().
			SetId(1).
			SetName("neim")).
		SetDate("2021-12-25").
		SetCategory("cat6").
		SetDescription("Ricarica").
		SetAmountNet(10).
		SetAmortization(10).
		SetRcCenter("bg").
		SetAmountVat(0).
		SetAmountWithholdingTax(0).
		SetAmountOtherWithholdingTax(0).
		SetTaxDeductibility(50).
		SetVatDeductibility(100).
		SetIsMarked(false).
		SetInvoiceNumber("in").
		SetIsMarked(false).
		SetIsDetailed(false).
		SetEInvoice(false).
		SetAttachmentToken("dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ==").
		SetCurrency(*NewCurrency().
			SetId("EUR").
			SetExchangeRate("1")).
		SetPaymentsList([]ReceivedDocumentPaymentsListItem{
			*NewReceivedDocumentPaymentsListItem().
				SetAmount(592).
				SetDueDate("2021-12-25").
				SetPaidDate("2021-12-25").
				SetPaymentTerms(*NewReceivedDocumentPaymentsListItemPaymentTerms().
					SetDays(0)).
				SetStatus("paid").
				SetPaymentAccount(*NewPaymentAccount().
					SetId(21))}).
		SetItemsList([]ReceivedDocumentItemsListItem{*NewReceivedDocumentItemsListItem().
			SetId(1)}).
		SetIsDetailed(false)

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifyReceivedDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":1,"type":"expense","entity":{"id":1,"name":"neim"},"date":"2021-12-25","category":"cat6","description":"Ricarica","amount_net":10,"amount_vat":0,"amount_withholding_tax":0,"amount_other_withholding_tax":0,"amortization":10,"rc_center":"bg","invoice_number":"in","is_marked":false,"is_detailed":false,"e_invoice":false,"currency":{"id":"EUR","exchange_rate":"1"},"tax_deductibility":50,"vat_deductibility":100,"items_list":[{"id":1}],"payments_list":[{"amount":592,"due_date":"2021-12-25","paid_date":"2021-12-25","payment_terms":{"days":0},"status":"paid","payment_account":{"id":21}}],"attachment_token":"dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ=="}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceivedDocumentsApi.ModifyReceivedDocument(context.Background(), 2, 2).ModifyReceivedDocumentRequest(*NewModifyReceivedDocumentRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceivedDocument().
		SetId(1).
		SetType(ReceivedDocumentTypes.EXPENSE).
		SetEntity(*NewEntity().
			SetId(1).
			SetName("neim")).
		SetDate("2021-12-25").
		SetCategory("cat6").
		SetDescription("Ricarica").
		SetAmountNet(10).
		SetAmortization(10).
		SetRcCenter("bg").
		SetAmountVat(0).
		SetAmountWithholdingTax(0).
		SetAmountOtherWithholdingTax(0).
		SetTaxDeductibility(50).
		SetVatDeductibility(100).
		SetIsMarked(false).
		SetInvoiceNumber("in").
		SetIsMarked(false).
		SetIsDetailed(false).
		SetEInvoice(false).
		SetAttachmentToken("dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ==").
		SetCurrency(*NewCurrency().
			SetId("EUR").
			SetExchangeRate("1")).
		SetPaymentsList([]ReceivedDocumentPaymentsListItem{
			*NewReceivedDocumentPaymentsListItem().
				SetAmount(592).
				SetDueDate("2021-12-25").
				SetPaidDate("2021-12-25").
				SetPaymentTerms(*NewReceivedDocumentPaymentsListItemPaymentTerms().
					SetDays(0)).
				SetStatus("paid").
				SetPaymentAccount(*NewPaymentAccount().
					SetId(21))}).
		SetItemsList([]ReceivedDocumentItemsListItem{*NewReceivedDocumentItemsListItem().
			SetId(1)}).
		SetIsDetailed(false)

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestUploadReceivedDocumentAttachment(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(``))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	_, resp, err := apiClient.ReceivedDocumentsApi.UploadReceivedDocumentAttachment(context.Background(), 2).Filename("filename").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}
