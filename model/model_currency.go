/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.3
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the Currency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Currency{}

// Currency struct for Currency
type Currency struct {
	// Currency code
	Id NullableString `json:"id,omitempty"`
	// Currency symbol
	Symbol NullableString `json:"symbol,omitempty"`
	// Currency exchange rate (EUR to this)
	ExchangeRate NullableString `json:"exchange_rate,omitempty"`
	// Currency html code
	HtmlSymbol NullableString `json:"html_symbol,omitempty"`
}

// NewCurrency instantiates a new Currency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrency() *Currency {
	this := Currency{}
	return &this
}

// NewCurrencyWithDefaults instantiates a new Currency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyWithDefaults() *Currency {
	this := Currency{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Currency) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Currency) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Currency) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Currency) SetId(v string) *Currency {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Currency) SetIdNil() *Currency {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Currency) UnsetId() {
	o.Id.Unset()
}

// GetSymbol returns the Symbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Currency) GetSymbol() string {
	if o == nil || IsNil(o.Symbol.Get()) {
		var ret string
		return ret
	}
	return *o.Symbol.Get()
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Currency) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Symbol.Get(), o.Symbol.IsSet()
}

// HasSymbol returns a boolean if a field has been set.
func (o *Currency) HasSymbol() bool {
	if o != nil && o.Symbol.IsSet() {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given NullableString and assigns it to the Symbol field.
func (o *Currency) SetSymbol(v string) *Currency {
	o.Symbol.Set(&v)
	return o
}
// SetSymbolNil sets the value for Symbol to be an explicit nil
func (o *Currency) SetSymbolNil() *Currency {
	o.Symbol.Set(nil)
	return o
}

// UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
func (o *Currency) UnsetSymbol() {
	o.Symbol.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Currency) GetExchangeRate() string {
	if o == nil || IsNil(o.ExchangeRate.Get()) {
		var ret string
		return ret
	}
	return *o.ExchangeRate.Get()
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Currency) GetExchangeRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExchangeRate.Get(), o.ExchangeRate.IsSet()
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *Currency) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given NullableString and assigns it to the ExchangeRate field.
func (o *Currency) SetExchangeRate(v string) *Currency {
	o.ExchangeRate.Set(&v)
	return o
}
// SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil
func (o *Currency) SetExchangeRateNil() *Currency {
	o.ExchangeRate.Set(nil)
	return o
}

// UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
func (o *Currency) UnsetExchangeRate() {
	o.ExchangeRate.Unset()
}

// GetHtmlSymbol returns the HtmlSymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Currency) GetHtmlSymbol() string {
	if o == nil || IsNil(o.HtmlSymbol.Get()) {
		var ret string
		return ret
	}
	return *o.HtmlSymbol.Get()
}

// GetHtmlSymbolOk returns a tuple with the HtmlSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Currency) GetHtmlSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HtmlSymbol.Get(), o.HtmlSymbol.IsSet()
}

// HasHtmlSymbol returns a boolean if a field has been set.
func (o *Currency) HasHtmlSymbol() bool {
	if o != nil && o.HtmlSymbol.IsSet() {
		return true
	}

	return false
}

// SetHtmlSymbol gets a reference to the given NullableString and assigns it to the HtmlSymbol field.
func (o *Currency) SetHtmlSymbol(v string) *Currency {
	o.HtmlSymbol.Set(&v)
	return o
}
// SetHtmlSymbolNil sets the value for HtmlSymbol to be an explicit nil
func (o *Currency) SetHtmlSymbolNil() *Currency {
	o.HtmlSymbol.Set(nil)
	return o
}

// UnsetHtmlSymbol ensures that no value is present for HtmlSymbol, not even an explicit nil
func (o *Currency) UnsetHtmlSymbol() {
	o.HtmlSymbol.Unset()
}

func (o Currency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Currency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Symbol.IsSet() {
		toSerialize["symbol"] = o.Symbol.Get()
	}
	if o.ExchangeRate.IsSet() {
		toSerialize["exchange_rate"] = o.ExchangeRate.Get()
	}
	if o.HtmlSymbol.IsSet() {
		toSerialize["html_symbol"] = o.HtmlSymbol.Get()
	}
	return toSerialize, nil
}

type NullableCurrency struct {
	value *Currency
	isSet bool
}

func (v NullableCurrency) Get() *Currency {
	return v.value
}

func (v *NullableCurrency) Set(val *Currency) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrency(val *Currency) *NullableCurrency {
	return &NullableCurrency{value: val, isSet: true}
}

func (v NullableCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


