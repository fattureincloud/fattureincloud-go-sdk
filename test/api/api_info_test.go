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

func TestListArchiveCategories(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":["cat5","cat6"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListArchiveCategories(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")
	expected := []string{"cat5", "cat6"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListCities(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"postal_code":"24016","city":"San Pellegrino Terme","province":"BG"},{"postal_code":"24016","city":"San Pellegrino Terme","province":"BG"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListCities(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewCity().
		SetCity("San Pellegrino Terme").
		SetProvince("BG").
		SetPostalCode("24016")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestListCostCenters(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":["bg","mi"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListCostCenters(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := []string{"bg", "mi"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListCountries(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":["italia", "albania"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListCountries(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := []string{"italia", "albania"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListCurrencies(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":"EUR","symbol":"€","exchange_rate":"1.00000","html_symbol":"\\u0026euro;"}, {"id":"USD","symbol":"$","exchange_rate":"0.84000","html_symbol":"&usd;"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListCurrencies(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewCurrency().
		SetId("USD").
		SetSymbol("$").
		SetExchangeRate("0.84000").
		SetHtmlSymbol("&usd;")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[1]))
}

func TestListDeliveryNotesDefaultCausals(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":["causal1", "causal2"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListDeliveryNotesDefaultCausals(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := []string{"causal1", "causal2"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListDetailedCountries(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"iso":"IT","fiscal_iso":"IT","name":"Italia","settings_name":"Italia"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListDetailedCountries(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewDetailedCountry().
		SetFiscalIso("IT").
		SetIso("IT").
		SetName("Italia").
		SetSettingsName("Italia")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestListLanguages(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"code":"IT","name":"Italiano"}, {"code":"ENG","name":"Inglese"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListLanguages(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewLanguage().
		SetCode("ENG").
		SetName("Inglese")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[1]))
}

func TestListPaymentAccounts(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":12,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false},{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListPaymentAccounts(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewPaymentAccount().
		SetId(21).
		SetName("Indesa - Carta conto").
		SetType(PaymentAccountTypes.STANDARD).
		SetIban("IT84Y0300203280294126225888").
		SetSia("sai").
		SetCuc("cuc").
		SetVirtual(false)

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[1]))
}

func TestListPaymentMethods(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":12345,"name":"Bonifico bancario","type":"standard","is_default":true,"default_payment_account":{"id":21,"name":"n1","type":"standard"},"details":[{"title":"t1"}],"bank_iban":"IT62W0300203280486429468578","bank_name":"Indesa","bank_beneficiary":"mamma","ei_payment_method":"2"},{"id":12346,"name":"Bonifico bancario","type":"standard","is_default":true,"default_payment_account":{"id":21,"name":"n1","type":"standard"},"details":[{"title":"t1"}],"bank_iban":"IT62W0300203280486429468578","bank_name":"Indesa","bank_beneficiary":"mamma","ei_payment_method":"2"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListPaymentMethods(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewPaymentMethod().
		SetId(12346).
		SetName("Bonifico bancario").
		SetType(PaymentMethodTypes.STANDARD).
		SetIsDefault(true).
		SetDetails([]PaymentMethodDetails{*NewPaymentMethodDetails().
			SetTitle("t1")}).
		SetDefaultPaymentAccount(*NewPaymentAccount().
			SetId(21).
			SetName("n1").
			SetType("standard")).
		SetBankIban("IT62W0300203280486429468578").
		SetBankName("Indesa").
		SetBankBeneficiary("mamma").
		SetEiPaymentMethod("2")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[1]))
}

func TestListProductCategories(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":["cat5","cat6"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListProductCategories(context.Background(), 2).Context("asd").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := []string{"cat5", "cat6"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListReceivedDocumentCategories(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":["cat5","cat6"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListReceivedDocumentCategories(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := []string{"cat5", "cat6"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListRevenueCenters(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":["bg","mi"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListRevenueCenters(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := []string{"bg", "mi"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListTemplates(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":10,"name":"New Standard S1","type":"Tipo 1"},{"id":20,"name":"New Standard S2","type":"Tipo 2"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListTemplates(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewDocumentTemplate().
		SetId(20).
		SetName("New Standard S2").
		SetType("Tipo 2")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[1]))
}

func TestListUnitsOfMeasure(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":["kg","km"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListUnitsOfMeasure(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := []string{"kg", "km"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListVatTypes(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":1,"value":22,"description":"Non imponibile art. 123","notes":"IVA non imponibile","e_invoice":true,"ei_type":"2","ei_description":"desc","is_disabled":false},{"id":2,"value":33,"description":"Non imponibile art. 123","notes":"IVA non imponibile","e_invoice":true,"ei_type":"2","ei_description":"desc","is_disabled":false}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.InfoAPI.ListVatTypes(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewVatType().
		SetId(2).
		SetValue(33).
		SetDescription("Non imponibile art. 123").
		SetNotes("IVA non imponibile").
		SetEInvoice(true).
		SetEiType("2").
		SetEiDescription("desc").
		SetIsDisabled(false)

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[1]))
}
