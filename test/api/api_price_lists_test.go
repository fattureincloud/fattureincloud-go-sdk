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

func TestGetPriceListItems(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"1":{"price":3.5}}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.PriceListsAPI.GetPriceListItems(context.Background(), 2, "1").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewGetPriceListItemsResponse().
		SetData(map[string]PriceListItem{
			"1": *NewPriceListItem().SetPrice(3.5),
		})

	assert.True(t, reflect.DeepEqual(expected.GetData(), *actual.Data))
}

func TestGetPriceLists(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":"10","name":"Listino 1","prices_type":"net","is_default":true,"valid_from":"2023-01-01","valid_to":"2023-12-31","type":"sell"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.PriceListsAPI.GetPriceLists(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewPriceList().
		SetId("10").
		SetName("Listino 1").
		SetPricesType("net").
		SetIsDefault(true).
		SetValidFrom("2023-01-01").
		SetValidTo("2023-12-31").
		SetType("sell")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}
