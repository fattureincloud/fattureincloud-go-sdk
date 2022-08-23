# ReceivedDocumentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValues** | Pointer to [**NullableReceivedDocumentInfoDefaultValues**](ReceivedDocumentInfoDefaultValues.md) |  | [optional] 
**ItemsDefaultValues** | Pointer to [**NullableReceivedDocumentInfoItemsDefaultValues**](ReceivedDocumentInfoItemsDefaultValues.md) |  | [optional] 
**CountriesList** | Pointer to **[]string** |  | [optional] 
**CurrenciesList** | Pointer to [**[]Currency**](Currency.md) |  | [optional] 
**CategoriesList** | Pointer to **[]string** |  | [optional] 
**PaymentAccountsList** | Pointer to [**[]PaymentAccount**](PaymentAccount.md) |  | [optional] 
**VatTypesList** | Pointer to [**[]VatType**](VatType.md) |  | [optional] 

## Methods

### NewReceivedDocumentInfo

`func NewReceivedDocumentInfo() *ReceivedDocumentInfo`

NewReceivedDocumentInfo instantiates a new ReceivedDocumentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivedDocumentInfoWithDefaults

`func NewReceivedDocumentInfoWithDefaults() *ReceivedDocumentInfo`

NewReceivedDocumentInfoWithDefaults instantiates a new ReceivedDocumentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValues

`func (o *ReceivedDocumentInfo) GetDefaultValues() ReceivedDocumentInfoDefaultValues`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *ReceivedDocumentInfo) GetDefaultValuesOk() (*ReceivedDocumentInfoDefaultValues, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *ReceivedDocumentInfo) SetDefaultValues(v ReceivedDocumentInfoDefaultValues)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *ReceivedDocumentInfo) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### SetDefaultValuesNil

`func (o *ReceivedDocumentInfo) SetDefaultValuesNil(b bool)`

 SetDefaultValuesNil sets the value for DefaultValues to be an explicit nil

### UnsetDefaultValues
`func (o *ReceivedDocumentInfo) UnsetDefaultValues()`

UnsetDefaultValues ensures that no value is present for DefaultValues, not even an explicit nil
### GetItemsDefaultValues

`func (o *ReceivedDocumentInfo) GetItemsDefaultValues() ReceivedDocumentInfoItemsDefaultValues`

GetItemsDefaultValues returns the ItemsDefaultValues field if non-nil, zero value otherwise.

### GetItemsDefaultValuesOk

`func (o *ReceivedDocumentInfo) GetItemsDefaultValuesOk() (*ReceivedDocumentInfoItemsDefaultValues, bool)`

GetItemsDefaultValuesOk returns a tuple with the ItemsDefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsDefaultValues

`func (o *ReceivedDocumentInfo) SetItemsDefaultValues(v ReceivedDocumentInfoItemsDefaultValues)`

SetItemsDefaultValues sets ItemsDefaultValues field to given value.

### HasItemsDefaultValues

`func (o *ReceivedDocumentInfo) HasItemsDefaultValues() bool`

HasItemsDefaultValues returns a boolean if a field has been set.

### SetItemsDefaultValuesNil

`func (o *ReceivedDocumentInfo) SetItemsDefaultValuesNil(b bool)`

 SetItemsDefaultValuesNil sets the value for ItemsDefaultValues to be an explicit nil

### UnsetItemsDefaultValues
`func (o *ReceivedDocumentInfo) UnsetItemsDefaultValues()`

UnsetItemsDefaultValues ensures that no value is present for ItemsDefaultValues, not even an explicit nil
### GetCountriesList

`func (o *ReceivedDocumentInfo) GetCountriesList() []string`

GetCountriesList returns the CountriesList field if non-nil, zero value otherwise.

### GetCountriesListOk

`func (o *ReceivedDocumentInfo) GetCountriesListOk() (*[]string, bool)`

GetCountriesListOk returns a tuple with the CountriesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountriesList

`func (o *ReceivedDocumentInfo) SetCountriesList(v []string)`

SetCountriesList sets CountriesList field to given value.

### HasCountriesList

`func (o *ReceivedDocumentInfo) HasCountriesList() bool`

HasCountriesList returns a boolean if a field has been set.

### SetCountriesListNil

`func (o *ReceivedDocumentInfo) SetCountriesListNil(b bool)`

 SetCountriesListNil sets the value for CountriesList to be an explicit nil

### UnsetCountriesList
`func (o *ReceivedDocumentInfo) UnsetCountriesList()`

UnsetCountriesList ensures that no value is present for CountriesList, not even an explicit nil
### GetCurrenciesList

`func (o *ReceivedDocumentInfo) GetCurrenciesList() []Currency`

GetCurrenciesList returns the CurrenciesList field if non-nil, zero value otherwise.

### GetCurrenciesListOk

`func (o *ReceivedDocumentInfo) GetCurrenciesListOk() (*[]Currency, bool)`

GetCurrenciesListOk returns a tuple with the CurrenciesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrenciesList

`func (o *ReceivedDocumentInfo) SetCurrenciesList(v []Currency)`

SetCurrenciesList sets CurrenciesList field to given value.

### HasCurrenciesList

`func (o *ReceivedDocumentInfo) HasCurrenciesList() bool`

HasCurrenciesList returns a boolean if a field has been set.

### SetCurrenciesListNil

`func (o *ReceivedDocumentInfo) SetCurrenciesListNil(b bool)`

 SetCurrenciesListNil sets the value for CurrenciesList to be an explicit nil

### UnsetCurrenciesList
`func (o *ReceivedDocumentInfo) UnsetCurrenciesList()`

UnsetCurrenciesList ensures that no value is present for CurrenciesList, not even an explicit nil
### GetCategoriesList

`func (o *ReceivedDocumentInfo) GetCategoriesList() []string`

GetCategoriesList returns the CategoriesList field if non-nil, zero value otherwise.

### GetCategoriesListOk

`func (o *ReceivedDocumentInfo) GetCategoriesListOk() (*[]string, bool)`

GetCategoriesListOk returns a tuple with the CategoriesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesList

`func (o *ReceivedDocumentInfo) SetCategoriesList(v []string)`

SetCategoriesList sets CategoriesList field to given value.

### HasCategoriesList

`func (o *ReceivedDocumentInfo) HasCategoriesList() bool`

HasCategoriesList returns a boolean if a field has been set.

### SetCategoriesListNil

`func (o *ReceivedDocumentInfo) SetCategoriesListNil(b bool)`

 SetCategoriesListNil sets the value for CategoriesList to be an explicit nil

### UnsetCategoriesList
`func (o *ReceivedDocumentInfo) UnsetCategoriesList()`

UnsetCategoriesList ensures that no value is present for CategoriesList, not even an explicit nil
### GetPaymentAccountsList

`func (o *ReceivedDocumentInfo) GetPaymentAccountsList() []PaymentAccount`

GetPaymentAccountsList returns the PaymentAccountsList field if non-nil, zero value otherwise.

### GetPaymentAccountsListOk

`func (o *ReceivedDocumentInfo) GetPaymentAccountsListOk() (*[]PaymentAccount, bool)`

GetPaymentAccountsListOk returns a tuple with the PaymentAccountsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountsList

`func (o *ReceivedDocumentInfo) SetPaymentAccountsList(v []PaymentAccount)`

SetPaymentAccountsList sets PaymentAccountsList field to given value.

### HasPaymentAccountsList

`func (o *ReceivedDocumentInfo) HasPaymentAccountsList() bool`

HasPaymentAccountsList returns a boolean if a field has been set.

### SetPaymentAccountsListNil

`func (o *ReceivedDocumentInfo) SetPaymentAccountsListNil(b bool)`

 SetPaymentAccountsListNil sets the value for PaymentAccountsList to be an explicit nil

### UnsetPaymentAccountsList
`func (o *ReceivedDocumentInfo) UnsetPaymentAccountsList()`

UnsetPaymentAccountsList ensures that no value is present for PaymentAccountsList, not even an explicit nil
### GetVatTypesList

`func (o *ReceivedDocumentInfo) GetVatTypesList() []VatType`

GetVatTypesList returns the VatTypesList field if non-nil, zero value otherwise.

### GetVatTypesListOk

`func (o *ReceivedDocumentInfo) GetVatTypesListOk() (*[]VatType, bool)`

GetVatTypesListOk returns a tuple with the VatTypesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatTypesList

`func (o *ReceivedDocumentInfo) SetVatTypesList(v []VatType)`

SetVatTypesList sets VatTypesList field to given value.

### HasVatTypesList

`func (o *ReceivedDocumentInfo) HasVatTypesList() bool`

HasVatTypesList returns a boolean if a field has been set.

### SetVatTypesListNil

`func (o *ReceivedDocumentInfo) SetVatTypesListNil(b bool)`

 SetVatTypesListNil sets the value for VatTypesList to be an explicit nil

### UnsetVatTypesList
`func (o *ReceivedDocumentInfo) UnsetVatTypesList()`

UnsetVatTypesList ensures that no value is present for VatTypesList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


