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

	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/api"
	. "github.com/fattureincloud/fattureincloud-go-sdk/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateCashbookEntry(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":"1","date":"2021-12-19","description":"desc","kind":"issued_document","type":"in","entity_name":"Rossi S.r.l.","document":{"id":12345,"type":"issued_document","path":"/doc1.pdf"},"amount_in":10,"payment_account_in":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false},"amount_out":10,"payment_account_out":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false}}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.CashbookApi.CreateCashbookEntry(context.Background(), 2).CreateCashbookEntryRequest(*NewCreateCashbookEntryRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewCashbookEntry().
		SetId("1").
		SetDate("2021-12-19").
		SetDescription("desc").
		SetKind(CashbookEntryKinds.ISSUED_DOCUMENT).
		SetType(CashbookEntryTypes.IN).
		SetEntityName("Rossi S.r.l.").
		SetDocument(*NewCashbookEntryDocument().
			SetId(12345).
			SetType("issued_document").
			SetPath("/doc1.pdf")).
		SetAmountOut(10).
		SetPaymentAccountOut(*NewPaymentAccount().
			SetId(21).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD).
			SetIban("IT84Y0300203280294126225888").
			SetSia("sai").
			SetCuc("cuc").
			SetVirtual(false)).
		SetAmountIn(10).
		SetPaymentAccountIn(*NewPaymentAccount().
			SetId(21).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD).
			SetIban("IT84Y0300203280294126225888").
			SetSia("sai").
			SetCuc("cuc").
			SetVirtual(false))

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestDeleteCashbookEntry(t *testing.T) {
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

	resp, err := apiClient.CashbookApi.DeleteCashbookEntry(context.Background(), 2, "2").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetCashbookEntry(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":"1","date":"2021-12-19","description":"desc","kind":"issued_document","type":"in","entity_name":"Rossi S.r.l.","document":{"id":12345,"type":"issued_document","path":"/doc1.pdf"},"amount_in":10,"payment_account_in":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false},"amount_out":10,"payment_account_out":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false}}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.CashbookApi.GetCashbookEntry(context.Background(), 2, "2").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewCashbookEntry().
		SetId("1").
		SetDate("2021-12-19").
		SetDescription("desc").
		SetKind(CashbookEntryKinds.ISSUED_DOCUMENT).
		SetType(CashbookEntryTypes.IN).
		SetEntityName("Rossi S.r.l.").
		SetDocument(*NewCashbookEntryDocument().
			SetId(12345).
			SetType("issued_document").
			SetPath("/doc1.pdf")).
		SetAmountOut(10).
		SetPaymentAccountOut(*NewPaymentAccount().
			SetId(21).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD).
			SetIban("IT84Y0300203280294126225888").
			SetSia("sai").
			SetCuc("cuc").
			SetVirtual(false)).
		SetAmountIn(10).
		SetPaymentAccountIn(*NewPaymentAccount().
			SetId(21).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD).
			SetIban("IT84Y0300203280294126225888").
			SetSia("sai").
			SetCuc("cuc").
			SetVirtual(false))

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListCashbookEntries(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":"1","date":"2021-12-19","description":"desc","kind":"issued_document","type":"in","entity_name":"Rossi S.r.l.","document":{"id":12345,"type":"issued_document","path":"/doc1.pdf"},"amount_in":10,"payment_account_in":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false},"amount_out":10,"payment_account_out":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false}}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.CashbookApi.ListCashbookEntries(context.Background(), 2).DateFrom("2020-10-10").DateTo("2022-10-10").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewCashbookEntry().
		SetId("1").
		SetDate("2021-12-19").
		SetDescription("desc").
		SetKind(CashbookEntryKinds.ISSUED_DOCUMENT).
		SetType(CashbookEntryTypes.IN).
		SetEntityName("Rossi S.r.l.").
		SetDocument(*NewCashbookEntryDocument().
			SetId(12345).
			SetType("issued_document").
			SetPath("/doc1.pdf")).
		SetAmountOut(10).
		SetPaymentAccountOut(*NewPaymentAccount().
			SetId(21).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD).
			SetIban("IT84Y0300203280294126225888").
			SetSia("sai").
			SetCuc("cuc").
			SetVirtual(false)).
		SetAmountIn(10).
		SetPaymentAccountIn(*NewPaymentAccount().
			SetId(21).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD).
			SetIban("IT84Y0300203280294126225888").
			SetSia("sai").
			SetCuc("cuc").
			SetVirtual(false))

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifyCashbookEntry(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":"1","date":"2021-12-19","description":"desc","kind":"issued_document","type":"in","entity_name":"Rossi S.r.l.","document":{"id":12345,"type":"issued_document","path":"/doc1.pdf"},"amount_in":10,"payment_account_in":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false},"amount_out":10,"payment_account_out":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false}}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.CashbookApi.ModifyCashbookEntry(context.Background(), 2, "2").ModifyCashbookEntryRequest(*NewModifyCashbookEntryRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewCashbookEntry().
		SetId("1").
		SetDate("2021-12-19").
		SetDescription("desc").
		SetKind(CashbookEntryKinds.ISSUED_DOCUMENT).
		SetType(CashbookEntryTypes.IN).
		SetEntityName("Rossi S.r.l.").
		SetDocument(*NewCashbookEntryDocument().
			SetId(12345).
			SetType("issued_document").
			SetPath("/doc1.pdf")).
		SetAmountOut(10).
		SetPaymentAccountOut(*NewPaymentAccount().
			SetId(21).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD).
			SetIban("IT84Y0300203280294126225888").
			SetSia("sai").
			SetCuc("cuc").
			SetVirtual(false)).
		SetAmountIn(10).
		SetPaymentAccountIn(*NewPaymentAccount().
			SetId(21).
			SetName("Indesa - Carta conto").
			SetType(PaymentAccountTypes.STANDARD).
			SetIban("IT84Y0300203280294126225888").
			SetSia("sai").
			SetCuc("cuc").
			SetVirtual(false))

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}
