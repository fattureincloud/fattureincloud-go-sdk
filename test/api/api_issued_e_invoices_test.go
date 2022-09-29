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
	"time"

	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestGetEInvoiceRejectionReason(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"reason":"errore","date":"2021-12-31T00:00:00Z"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedEInvoicesApi.GetEInvoiceRejectionReason(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewEInvoiceRejectionReason().
		SetDate(time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC)).
		SetReason("errore")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetEInvoiceXml(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`<xml-fattura>fields</xml-fattura>`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedEInvoicesApi.GetEInvoiceXml(context.Background(), 2, 2).IncludeAttachment(true).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.True(t, reflect.DeepEqual("<xml-fattura>fields</xml-fattura>", actual))
}

func TestSendEInvoice(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"name":"neim","date":"2021-12-31"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedEInvoicesApi.SendEInvoice(context.Background(), 2, 2).SendEInvoiceRequest(*NewSendEInvoiceRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewSendEInvoiceResponseData().SetName("neim").SetDate("2021-12-31")

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestVerifyEInvoiceXml(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"success":true}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedEInvoicesApi.VerifyEInvoiceXml(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewVerifyEInvoiceXmlResponseData().SetSuccess(true)

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}
