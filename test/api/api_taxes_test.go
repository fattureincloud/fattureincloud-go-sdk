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

func TestCreateF24(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"due_date":"2021-12-31","status":"paid","payment_account":{"id":111,"name":"Indesa - Carta conto","type":"standard"},"amount":840.36,"attachment_token":"Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879","description":"PAGAMENTO IVA 2021"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.TaxesApi.CreateF24(context.Background(), 2).CreateF24Request(*NewCreateF24Request()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewF24().
		SetId(12345).
		SetAmount(840.36).
		SetDescription("PAGAMENTO IVA 2021").
		SetDueDate("2021-12-31").
		SetStatus(F24Statuses.PAID).
		SetPaymentAccount(*NewPaymentAccount().
			SetId(111).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD)).
		SetAttachmentToken("Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestDeleteF24(t *testing.T) {
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

	resp, err := apiClient.TaxesApi.DeleteF24(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteF24Attachment(t *testing.T) {
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

	resp, err := apiClient.TaxesApi.DeleteF24Attachment(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetF24(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"due_date":"2021-12-31","status":"paid","payment_account":{"id":111,"name":"Indesa - Carta conto","type":"standard"},"amount":840.36,"attachment_token":"Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879","description":"PAGAMENTO IVA 2021"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.TaxesApi.GetF24(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewF24().
		SetId(12345).
		SetAmount(840.36).
		SetDescription("PAGAMENTO IVA 2021").
		SetDueDate("2021-12-31").
		SetStatus(F24Statuses.PAID).
		SetPaymentAccount(*NewPaymentAccount().
			SetId(111).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD)).
		SetAttachmentToken("Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListF24(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":12345,"due_date":"2021-12-31","status":"paid","payment_account":{"id":111,"name":"Indesa - Carta conto","type":"standard"},"amount":840.36,"attachment_token":"Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879","description":"PAGAMENTO IVA 2021"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.TaxesApi.ListF24(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewF24().
		SetId(12345).
		SetAmount(840.36).
		SetDescription("PAGAMENTO IVA 2021").
		SetDueDate("2021-12-31").
		SetStatus(F24Statuses.PAID).
		SetPaymentAccount(*NewPaymentAccount().
			SetId(111).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD)).
		SetAttachmentToken("Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifyF24(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"due_date":"2021-12-31","status":"paid","payment_account":{"id":111,"name":"Indesa - Carta conto","type":"standard"},"amount":840.36,"attachment_token":"Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879","description":"PAGAMENTO IVA 2021"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.TaxesApi.ModifyF24(context.Background(), 2, 2).ModifyF24Request(*NewModifyF24Request()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewF24().
		SetId(12345).
		SetAmount(840.36).
		SetDescription("PAGAMENTO IVA 2021").
		SetDueDate("2021-12-31").
		SetStatus(F24Statuses.PAID).
		SetPaymentAccount(*NewPaymentAccount().
			SetId(111).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD)).
		SetAttachmentToken("Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestUploadF24Attachment(t *testing.T) {
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

	_, resp, err := apiClient.TaxesApi.UploadF24Attachment(context.Background(), 2).Filename("filename").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}
