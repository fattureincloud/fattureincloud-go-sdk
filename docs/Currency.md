# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Currency code | [optional] 
**Symbol** | Pointer to **NullableString** | Currency symbol | [optional] 
**ExchangeRate** | Pointer to **NullableString** | Currency exchange rate (EUR to this) | [optional] 
**HtmlSymbol** | Pointer to **NullableString** | Currency html code | [optional] 

## Methods

### NewCurrency

`func NewCurrency() *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Currency) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Currency) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Currency) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Currency) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Currency) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Currency) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSymbol

`func (o *Currency) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Currency) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Currency) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Currency) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### SetSymbolNil

`func (o *Currency) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *Currency) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
### GetExchangeRate

`func (o *Currency) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *Currency) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *Currency) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *Currency) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *Currency) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *Currency) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetHtmlSymbol

`func (o *Currency) GetHtmlSymbol() string`

GetHtmlSymbol returns the HtmlSymbol field if non-nil, zero value otherwise.

### GetHtmlSymbolOk

`func (o *Currency) GetHtmlSymbolOk() (*string, bool)`

GetHtmlSymbolOk returns a tuple with the HtmlSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlSymbol

`func (o *Currency) SetHtmlSymbol(v string)`

SetHtmlSymbol sets HtmlSymbol field to given value.

### HasHtmlSymbol

`func (o *Currency) HasHtmlSymbol() bool`

HasHtmlSymbol returns a boolean if a field has been set.

### SetHtmlSymbolNil

`func (o *Currency) SetHtmlSymbolNil(b bool)`

 SetHtmlSymbolNil sets the value for HtmlSymbol to be an explicit nil

### UnsetHtmlSymbol
`func (o *Currency) UnsetHtmlSymbol()`

UnsetHtmlSymbol ensures that no value is present for HtmlSymbol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


