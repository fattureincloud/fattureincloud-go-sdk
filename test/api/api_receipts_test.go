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

func TestCreateReceipt(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":10,"date":"2021-12-25","number":10,"numeration":"num","amount_net":10,"amount_vat":10,"amount_gross":10,"use_gross_prices":true,"type":"till_receipt","description":"descr","rc_center":"bg","created_at":"2021-10-10","updated_at":"2021-10-10","payment_account":{"id":1,"type":"standard"},"items_list":[{"id":1}]}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceiptsApi.CreateReceipt(context.Background(), 2).CreateReceiptRequest(*NewCreateReceiptRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceipt().
		SetId(10).
		SetDate("2021-12-25").
		SetNumber(10).
		SetNumeration("num").
		SetAmountNet(10).
		SetAmountVat(10).
		SetAmountGross(10).
		SetUseGrossPrices(true).
		SetType(ReceiptTypes.TILL_RECEIPT).
		SetDescription("descr").
		SetRcCenter("bg").
		SetCreatedAt("2021-10-10").
		SetUpdatedAt("2021-10-10").
		SetPaymentAccount(*NewPaymentAccount().SetId(1).SetType(PaymentAccountTypes.STANDARD)).
		SetItemsList([]ReceiptItemsListItem{*NewReceiptItemsListItem().SetId(1)})

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestDeleteReceipt(t *testing.T) {
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

	resp, err := apiClient.ReceiptsApi.DeleteReceipt(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetReceipt(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":10,"date":"2021-12-25","number":10,"numeration":"num","amount_net":10,"amount_vat":10,"amount_gross":10,"use_gross_prices":true,"type":"till_receipt","description":"descr","rc_center":"bg","created_at":"2021-10-10","updated_at":"2021-10-10","payment_account":{"id":1,"type":"standard"},"items_list":[{"id":1}]}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceiptsApi.GetReceipt(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceipt().
		SetId(10).
		SetDate("2021-12-25").
		SetNumber(10).
		SetNumeration("num").
		SetAmountNet(10).
		SetAmountVat(10).
		SetAmountGross(10).
		SetUseGrossPrices(true).
		SetType(ReceiptTypes.TILL_RECEIPT).
		SetDescription("descr").
		SetRcCenter("bg").
		SetCreatedAt("2021-10-10").
		SetUpdatedAt("2021-10-10").
		SetPaymentAccount(*NewPaymentAccount().SetId(1).SetType(PaymentAccountTypes.STANDARD)).
		SetItemsList([]ReceiptItemsListItem{*NewReceiptItemsListItem().SetId(1)})

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetReceiptPreCreateInfo(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"numerations_list":["num"],"rc_centers_list":["bg"],"payment_accounts_list":[{"id":1,"type":"standard"}],"categories_list":["cat6"],"vat_types_list":[{"id":1}]}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceiptsApi.GetReceiptPreCreateInfo(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceiptPreCreateInfo().
		SetNumerations(nil).
		SetNumerationsList([]string{"num"}).
		SetRcCentersList([]string{"bg"}).
		SetPaymentAccountsList([]PaymentAccount{*NewPaymentAccount().SetId(1).SetType(PaymentAccountTypes.STANDARD)}).
		SetCategoriesList([]string{"cat6"}).
		SetVatTypesList([]VatType{*NewVatType().SetId(1)})

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetReceiptsMonthlyTotals(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"net":10,"gross":10,"count":10}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceiptsApi.GetReceiptsMonthlyTotals(context.Background(), 2).Type_("issued_document").Year("2022").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewMonthlyTotal().
		SetNet(10).
		SetGross(10).
		SetCount(10)

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestListReceipts(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":10,"date":"2021-12-25","number":10,"numeration":"num","amount_net":10,"amount_vat":10,"amount_gross":10,"use_gross_prices":true,"type":"till_receipt","description":"descr","rc_center":"bg","created_at":"2021-10-10","updated_at":"2021-10-10","payment_account":{"id":1,"type":"standard"},"items_list":[{"id":1}]}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceiptsApi.ListReceipts(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceipt().
		SetId(10).
		SetDate("2021-12-25").
		SetNumber(10).
		SetNumeration("num").
		SetAmountNet(10).
		SetAmountVat(10).
		SetAmountGross(10).
		SetUseGrossPrices(true).
		SetType(ReceiptTypes.TILL_RECEIPT).
		SetDescription("descr").
		SetRcCenter("bg").
		SetCreatedAt("2021-10-10").
		SetUpdatedAt("2021-10-10").
		SetPaymentAccount(*NewPaymentAccount().SetId(1).SetType(PaymentAccountTypes.STANDARD)).
		SetItemsList([]ReceiptItemsListItem{*NewReceiptItemsListItem().SetId(1)})

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifyReceipt(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":10,"date":"2021-12-25","number":10,"numeration":"num","amount_net":10,"amount_vat":10,"amount_gross":10,"use_gross_prices":true,"type":"till_receipt","description":"descr","rc_center":"bg","created_at":"2021-10-10","updated_at":"2021-10-10","payment_account":{"id":1,"type":"standard"},"items_list":[{"id":1}]}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ReceiptsApi.ModifyReceipt(context.Background(), 2, 2).ModifyReceiptRequest(*NewModifyReceiptRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewReceipt().
		SetId(10).
		SetDate("2021-12-25").
		SetNumber(10).
		SetNumeration("num").
		SetAmountNet(10).
		SetAmountVat(10).
		SetAmountGross(10).
		SetUseGrossPrices(true).
		SetType(ReceiptTypes.TILL_RECEIPT).
		SetDescription("descr").
		SetRcCenter("bg").
		SetCreatedAt("2021-10-10").
		SetUpdatedAt("2021-10-10").
		SetPaymentAccount(*NewPaymentAccount().SetId(1).SetType(PaymentAccountTypes.STANDARD)).
		SetItemsList([]ReceiptItemsListItem{*NewReceiptItemsListItem().SetId(1)})

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}
