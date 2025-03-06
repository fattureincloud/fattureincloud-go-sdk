# EntityClientPreCreateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountriesList** | Pointer to **[]string** | Countries list | [optional] 
**PaymentMethodsList** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | Entity payment methods list | [optional] 
**PaymentAccountsList** | Pointer to [**[]PaymentAccount**](PaymentAccount.md) | Entity payment accounts list | [optional] 
**VatTypesList** | Pointer to [**[]VatType**](VatType.md) | Vat types list | [optional] 
**PriceLists** | Pointer to [**[]PriceList**](PriceList.md) | Entity price lists | [optional] 
**Limit** | Pointer to **NullableFloat32** | Entity limit | [optional] 
**Usage** | Pointer to **NullableFloat32** | Entity usage | [optional] 

## Methods

### NewEntityClientPreCreateInfo

`func NewEntityClientPreCreateInfo() *EntityClientPreCreateInfo`

NewEntityClientPreCreateInfo instantiates a new EntityClientPreCreateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityClientPreCreateInfoWithDefaults

`func NewEntityClientPreCreateInfoWithDefaults() *EntityClientPreCreateInfo`

NewEntityClientPreCreateInfoWithDefaults instantiates a new EntityClientPreCreateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountriesList

`func (o *EntityClientPreCreateInfo) GetCountriesList() []string`

GetCountriesList returns the CountriesList field if non-nil, zero value otherwise.

### GetCountriesListOk

`func (o *EntityClientPreCreateInfo) GetCountriesListOk() (*[]string, bool)`

GetCountriesListOk returns a tuple with the CountriesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountriesList

`func (o *EntityClientPreCreateInfo) SetCountriesList(v []string)`

SetCountriesList sets CountriesList field to given value.

### HasCountriesList

`func (o *EntityClientPreCreateInfo) HasCountriesList() bool`

HasCountriesList returns a boolean if a field has been set.

### SetCountriesListNil

`func (o *EntityClientPreCreateInfo) SetCountriesListNil(b bool)`

 SetCountriesListNil sets the value for CountriesList to be an explicit nil

### UnsetCountriesList
`func (o *EntityClientPreCreateInfo) UnsetCountriesList()`

UnsetCountriesList ensures that no value is present for CountriesList, not even an explicit nil
### GetPaymentMethodsList

`func (o *EntityClientPreCreateInfo) GetPaymentMethodsList() []PaymentMethod`

GetPaymentMethodsList returns the PaymentMethodsList field if non-nil, zero value otherwise.

### GetPaymentMethodsListOk

`func (o *EntityClientPreCreateInfo) GetPaymentMethodsListOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsListOk returns a tuple with the PaymentMethodsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodsList

`func (o *EntityClientPreCreateInfo) SetPaymentMethodsList(v []PaymentMethod)`

SetPaymentMethodsList sets PaymentMethodsList field to given value.

### HasPaymentMethodsList

`func (o *EntityClientPreCreateInfo) HasPaymentMethodsList() bool`

HasPaymentMethodsList returns a boolean if a field has been set.

### SetPaymentMethodsListNil

`func (o *EntityClientPreCreateInfo) SetPaymentMethodsListNil(b bool)`

 SetPaymentMethodsListNil sets the value for PaymentMethodsList to be an explicit nil

### UnsetPaymentMethodsList
`func (o *EntityClientPreCreateInfo) UnsetPaymentMethodsList()`

UnsetPaymentMethodsList ensures that no value is present for PaymentMethodsList, not even an explicit nil
### GetPaymentAccountsList

`func (o *EntityClientPreCreateInfo) GetPaymentAccountsList() []PaymentAccount`

GetPaymentAccountsList returns the PaymentAccountsList field if non-nil, zero value otherwise.

### GetPaymentAccountsListOk

`func (o *EntityClientPreCreateInfo) GetPaymentAccountsListOk() (*[]PaymentAccount, bool)`

GetPaymentAccountsListOk returns a tuple with the PaymentAccountsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountsList

`func (o *EntityClientPreCreateInfo) SetPaymentAccountsList(v []PaymentAccount)`

SetPaymentAccountsList sets PaymentAccountsList field to given value.

### HasPaymentAccountsList

`func (o *EntityClientPreCreateInfo) HasPaymentAccountsList() bool`

HasPaymentAccountsList returns a boolean if a field has been set.

### SetPaymentAccountsListNil

`func (o *EntityClientPreCreateInfo) SetPaymentAccountsListNil(b bool)`

 SetPaymentAccountsListNil sets the value for PaymentAccountsList to be an explicit nil

### UnsetPaymentAccountsList
`func (o *EntityClientPreCreateInfo) UnsetPaymentAccountsList()`

UnsetPaymentAccountsList ensures that no value is present for PaymentAccountsList, not even an explicit nil
### GetVatTypesList

`func (o *EntityClientPreCreateInfo) GetVatTypesList() []VatType`

GetVatTypesList returns the VatTypesList field if non-nil, zero value otherwise.

### GetVatTypesListOk

`func (o *EntityClientPreCreateInfo) GetVatTypesListOk() (*[]VatType, bool)`

GetVatTypesListOk returns a tuple with the VatTypesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatTypesList

`func (o *EntityClientPreCreateInfo) SetVatTypesList(v []VatType)`

SetVatTypesList sets VatTypesList field to given value.

### HasVatTypesList

`func (o *EntityClientPreCreateInfo) HasVatTypesList() bool`

HasVatTypesList returns a boolean if a field has been set.

### SetVatTypesListNil

`func (o *EntityClientPreCreateInfo) SetVatTypesListNil(b bool)`

 SetVatTypesListNil sets the value for VatTypesList to be an explicit nil

### UnsetVatTypesList
`func (o *EntityClientPreCreateInfo) UnsetVatTypesList()`

UnsetVatTypesList ensures that no value is present for VatTypesList, not even an explicit nil
### GetPriceLists

`func (o *EntityClientPreCreateInfo) GetPriceLists() []PriceList`

GetPriceLists returns the PriceLists field if non-nil, zero value otherwise.

### GetPriceListsOk

`func (o *EntityClientPreCreateInfo) GetPriceListsOk() (*[]PriceList, bool)`

GetPriceListsOk returns a tuple with the PriceLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLists

`func (o *EntityClientPreCreateInfo) SetPriceLists(v []PriceList)`

SetPriceLists sets PriceLists field to given value.

### HasPriceLists

`func (o *EntityClientPreCreateInfo) HasPriceLists() bool`

HasPriceLists returns a boolean if a field has been set.

### SetPriceListsNil

`func (o *EntityClientPreCreateInfo) SetPriceListsNil(b bool)`

 SetPriceListsNil sets the value for PriceLists to be an explicit nil

### UnsetPriceLists
`func (o *EntityClientPreCreateInfo) UnsetPriceLists()`

UnsetPriceLists ensures that no value is present for PriceLists, not even an explicit nil
### GetLimit

`func (o *EntityClientPreCreateInfo) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EntityClientPreCreateInfo) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EntityClientPreCreateInfo) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *EntityClientPreCreateInfo) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *EntityClientPreCreateInfo) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *EntityClientPreCreateInfo) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetUsage

`func (o *EntityClientPreCreateInfo) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EntityClientPreCreateInfo) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EntityClientPreCreateInfo) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *EntityClientPreCreateInfo) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *EntityClientPreCreateInfo) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *EntityClientPreCreateInfo) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


