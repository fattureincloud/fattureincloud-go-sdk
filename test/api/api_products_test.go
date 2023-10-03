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

func TestCreateProduct(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":1,"name":"neim","code":"cod","net_price":10,"gross_price":10,"use_gross_price":true,"default_vat":{"id":1},"net_cost":10,"measure":"big","description":"desc","category":"cat6","notes":"nots","in_stock":true,"stock_initial":10,"average_cost":10,"average_price":10,"created_at":"2021-10-10","updated_at":"2021-10-10"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ProductsAPI.CreateProduct(context.Background(), 2).CreateProductRequest(*NewCreateProductRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewProduct().
		SetId(1).
		SetName("neim").
		SetCode("cod").
		SetNetPrice(10).
		SetGrossPrice(10).
		SetUseGrossPrice(true).
		SetDefaultVat(*NewVatType().SetId(1)).
		SetNetCost(10).
		SetMeasure("big").
		SetDescription("desc").
		SetCategory("cat6").
		SetNotes("nots").
		SetInStock(true).
		SetStockInitial(10).
		SetAverageCost(10).
		SetAveragePrice(10).
		SetCreatedAt("2021-10-10").
		SetUpdatedAt("2021-10-10")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestDeleteProduct(t *testing.T) {
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

	resp, err := apiClient.ProductsAPI.DeleteProduct(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetProduct(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":1,"name":"neim","code":"cod","net_price":10,"gross_price":10,"use_gross_price":true,"default_vat":{"id":1},"net_cost":10,"measure":"big","description":"desc","category":"cat6","notes":"nots","in_stock":true,"stock_initial":10,"average_cost":10,"average_price":10,"created_at":"2021-10-10","updated_at":"2021-10-10"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ProductsAPI.GetProduct(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewProduct().
		SetId(1).
		SetName("neim").
		SetCode("cod").
		SetNetPrice(10).
		SetGrossPrice(10).
		SetUseGrossPrice(true).
		SetDefaultVat(*NewVatType().SetId(1)).
		SetNetCost(10).
		SetMeasure("big").
		SetDescription("desc").
		SetCategory("cat6").
		SetNotes("nots").
		SetInStock(true).
		SetStockInitial(10).
		SetAverageCost(10).
		SetAveragePrice(10).
		SetCreatedAt("2021-10-10").
		SetUpdatedAt("2021-10-10")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListProducts(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":1,"name":"neim","code":"cod","net_price":10,"gross_price":10,"use_gross_price":true,"default_vat":{"id":1},"net_cost":10,"measure":"big","description":"desc","category":"cat6","notes":"nots","in_stock":true,"stock_initial":10,"average_cost":10,"average_price":10,"created_at":"2021-10-10","updated_at":"2021-10-10"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ProductsAPI.ListProducts(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewProduct().
		SetId(1).
		SetName("neim").
		SetCode("cod").
		SetNetPrice(10).
		SetGrossPrice(10).
		SetUseGrossPrice(true).
		SetDefaultVat(*NewVatType().SetId(1)).
		SetNetCost(10).
		SetMeasure("big").
		SetDescription("desc").
		SetCategory("cat6").
		SetNotes("nots").
		SetInStock(true).
		SetStockInitial(10).
		SetAverageCost(10).
		SetAveragePrice(10).
		SetCreatedAt("2021-10-10").
		SetUpdatedAt("2021-10-10")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifyProduct(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":1,"name":"neim","code":"cod","net_price":10,"gross_price":10,"use_gross_price":true,"default_vat":{"id":1},"net_cost":10,"measure":"big","description":"desc","category":"cat6","notes":"nots","in_stock":true,"stock_initial":10,"average_cost":10,"average_price":10,"created_at":"2021-10-10","updated_at":"2021-10-10"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ProductsAPI.ModifyProduct(context.Background(), 2, 2).ModifyProductRequest(*NewModifyProductRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewProduct().
		SetId(1).
		SetName("neim").
		SetCode("cod").
		SetNetPrice(10).
		SetGrossPrice(10).
		SetUseGrossPrice(true).
		SetDefaultVat(*NewVatType().SetId(1)).
		SetNetCost(10).
		SetMeasure("big").
		SetDescription("desc").
		SetCategory("cat6").
		SetNotes("nots").
		SetInStock(true).
		SetStockInitial(10).
		SetAverageCost(10).
		SetAveragePrice(10).
		SetCreatedAt("2021-10-10").
		SetUpdatedAt("2021-10-10")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}
