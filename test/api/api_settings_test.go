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

func TestCreatePaymentAccount(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.CreatePaymentAccount(context.Background(), 2).CreatePaymentAccountRequest(*NewCreatePaymentAccountRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewPaymentAccount().
		SetId(21).
		SetName("Indesa - Carta conto").
		SetType(PaymentAccountTypes.STANDARD).
		SetIban("IT84Y0300203280294126225888").
		SetSia("sai").
		SetCuc("cuc").
		SetVirtual(false)

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestCreatePaymentMethod(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12346,"name":"Bonifico bancario","type":"standard","is_default":true,"default_payment_account":{"id":21,"name":"n1","type":"standard"},"details":[{"title":"t1"}],"bank_iban":"IT62W0300203280486429468578","bank_name":"Indesa","bank_beneficiary":"mamma","ei_payment_method":"2"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.CreatePaymentMethod(context.Background(), 2).CreatePaymentMethodRequest(*NewCreatePaymentMethodRequest()).Execute()
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
			SetType(PaymentAccountTypes.STANDARD)).
		SetBankIban("IT62W0300203280486429468578").
		SetBankName("Indesa").
		SetBankBeneficiary("mamma").
		SetEiPaymentMethod("2")

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestCreateVatType(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":0,"value":22,"description":"Non imponibile art. 123","notes":"IVA non imponibile","e_invoice":true,"ei_type":"2","ei_description":"desc","is_disabled":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.CreateVatType(context.Background(), 2).CreateVatTypeRequest(*NewCreateVatTypeRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewVatType().
		SetId(0).
		SetValue(22).
		SetDescription("Non imponibile art. 123").
		SetNotes("IVA non imponibile").
		SetEInvoice(true).
		SetEiType("2").
		SetEiDescription("desc").
		SetIsDisabled(false)

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestDeletePaymentAccount(t *testing.T) {
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

	resp, err := apiClient.SettingsAPI.DeletePaymentAccount(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeletePaymentMethod(t *testing.T) {
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

	resp, err := apiClient.SettingsAPI.DeletePaymentMethod(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteVatType(t *testing.T) {
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

	resp, err := apiClient.SettingsAPI.DeleteVatType(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetPaymentAccount(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.GetPaymentAccount(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewPaymentAccount().
		SetId(21).
		SetName("Indesa - Carta conto").
		SetType(PaymentAccountTypes.STANDARD).
		SetIban("IT84Y0300203280294126225888").
		SetSia("sai").
		SetCuc("cuc").
		SetVirtual(false)

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestGetPaymentMethod(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12346,"name":"Bonifico bancario","type":"standard","is_default":true,"default_payment_account":{"id":21,"name":"n1","type":"standard"},"details":[{"title":"t1"}],"bank_iban":"IT62W0300203280486429468578","bank_name":"Indesa","bank_beneficiary":"mamma","ei_payment_method":"2"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.GetPaymentMethod(context.Background(), 2, 2).Execute()
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
			SetType(PaymentAccountTypes.STANDARD)).
		SetBankIban("IT62W0300203280486429468578").
		SetBankName("Indesa").
		SetBankBeneficiary("mamma").
		SetEiPaymentMethod("2")

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestGetVatType(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":0,"value":22,"description":"Non imponibile art. 123","notes":"IVA non imponibile","e_invoice":true,"ei_type":"2","ei_description":"desc","is_disabled":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.GetVatType(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewVatType().
		SetId(0).
		SetValue(22).
		SetDescription("Non imponibile art. 123").
		SetNotes("IVA non imponibile").
		SetEInvoice(true).
		SetEiType("2").
		SetEiDescription("desc").
		SetIsDisabled(false)

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestModifyPaymentAccount(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":21,"name":"Indesa - Carta conto","type":"standard","iban":"IT84Y0300203280294126225888","sia":"sai","cuc":"cuc","virtual":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.ModifyPaymentAccount(context.Background(), 2, 2).ModifyPaymentAccountRequest(*NewModifyPaymentAccountRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewPaymentAccount().
		SetId(21).
		SetName("Indesa - Carta conto").
		SetType(PaymentAccountTypes.STANDARD).
		SetIban("IT84Y0300203280294126225888").
		SetSia("sai").
		SetCuc("cuc").
		SetVirtual(false)

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestModifyPaymentMethod(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12346,"name":"Bonifico bancario","type":"standard","is_default":true,"default_payment_account":{"id":21,"name":"n1","type":"standard"},"details":[{"title":"t1"}],"bank_iban":"IT62W0300203280486429468578","bank_name":"Indesa","bank_beneficiary":"mamma","ei_payment_method":"2"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.ModifyPaymentMethod(context.Background(), 2, 2).ModifyPaymentMethodRequest(*NewModifyPaymentMethodRequest()).Execute()
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
			SetType(PaymentAccountTypes.STANDARD)).
		SetBankIban("IT62W0300203280486429468578").
		SetBankName("Indesa").
		SetBankBeneficiary("mamma").
		SetEiPaymentMethod("2")

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestModifyVatType(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":0,"value":22,"description":"Non imponibile art. 123","notes":"IVA non imponibile","e_invoice":true,"ei_type":"2","ei_description":"desc","is_disabled":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.ModifyVatType(context.Background(), 2, 2).ModifyVatTypeRequest(*NewModifyVatTypeRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewVatType().
		SetId(0).
		SetValue(22).
		SetDescription("Non imponibile art. 123").
		SetNotes("IVA non imponibile").
		SetEInvoice(true).
		SetEiType("2").
		SetEiDescription("desc").
		SetIsDisabled(false)

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}

func TestGetTaxProfile(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"company_type":"individual","company_subtype":"artigiani","regime":"forfettario_5","rivalsa_name":"","default_rivalsa":0,"cassa_name":"","default_cassa":0,"default_cassa_taxable":100,"cassa2_name":"","default_cassa2":0,"default_cassa2_taxable":0,"default_withholding_tax":0,"default_withholding_tax_taxable":100,"default_other_withholding_tax":0,"enasarco":false,"contributions_percentage":0,"med":false,"default_vat":{"id":66,"value":0,"description":"Contribuenti forfettari","notes":"Operazione non soggetta a IVA ai sensi dell'art. 1, commi 54-89, Legge n. 190/2014 e succ. modifiche/integrazioni","e_invoice":true,"ei_type":"2.2","ei_description":"Non soggetta art. 1/54-89 L. 190/2014 e succ. modifiche/integrazioni","editable":false,"is_disabled":false,"default":true}}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SettingsAPI.GetTaxProfile(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewTaxProfile().
		SetCompanyType("individual").
		SetCompanySubtype("artigiani").
		SetRegime("forfettario_5").
		SetRivalsaName("").
		SetDefaultRivalsa(0).
		SetCassaName("").
		SetDefaultCassa(0).
		SetDefaultCassaTaxable(100).
		SetCassa2Name("").
		SetDefaultCassa2(0).
		SetDefaultCassa2Taxable(0).
		SetDefaultWithholdingTax(0).
		SetDefaultWithholdingTaxTaxable(100).
		SetDefaultOtherWithholdingTax(0).
		SetEnasarco(false).
		SetContributionsPercentage(0).
		SetMed(false).
		SetDefaultVat(*NewVatType().
			SetId(66).
			SetValue(0).
			SetDescription("Contribuenti forfettari").
			SetNotes("Operazione non soggetta a IVA ai sensi dell'art. 1, commi 54-89, Legge n. 190/2014 e succ. modifiche/integrazioni").
			SetEInvoice(true).
			SetEiType("2.2").
			SetEiDescription("Non soggetta art. 1/54-89 L. 190/2014 e succ. modifiche/integrazioni").
			SetEditable(false).
			SetIsDisabled(false).
			SetDefault(true))

	assert.True(t, reflect.DeepEqual(*expected, actual.GetData()))
}
