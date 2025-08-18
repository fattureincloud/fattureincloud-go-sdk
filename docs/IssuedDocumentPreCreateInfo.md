# IssuedDocumentPreCreateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numerations** | Pointer to **map[string]map[string]int32** |  | [optional] 
**DnNumerations** | Pointer to **map[string]map[string]int32** |  | [optional] 
**DefaultValues** | Pointer to [**NullableIssuedDocumentPreCreateInfoDefaultValues**](IssuedDocumentPreCreateInfoDefaultValues.md) |  | [optional] 
**ExtraDataDefaultValues** | Pointer to [**NullableIssuedDocumentPreCreateInfoExtraDataDefaultValues**](IssuedDocumentPreCreateInfoExtraDataDefaultValues.md) |  | [optional] 
**ItemsDefaultValues** | Pointer to [**NullableIssuedDocumentPreCreateInfoItemsDefaultValues**](IssuedDocumentPreCreateInfoItemsDefaultValues.md) |  | [optional] 
**CountriesList** | Pointer to **[]string** | Countries list | [optional] 
**CurrenciesList** | Pointer to [**[]Currency**](Currency.md) | Currencies list | [optional] 
**TemplatesList** | Pointer to [**[]DocumentTemplate**](DocumentTemplate.md) | Document templates list | [optional] 
**DnTemplatesList** | Pointer to [**[]DocumentTemplate**](DocumentTemplate.md) | Delivery note templates list | [optional] 
**AiTemplatesList** | Pointer to [**[]DocumentTemplate**](DocumentTemplate.md) | Accompanying invoice templates list | [optional] 
**PaymentMethodsList** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | Payment methods list | [optional] 
**PaymentAccountsList** | Pointer to [**[]PaymentAccount**](PaymentAccount.md) | Payment accounts list | [optional] 
**VatTypesList** | Pointer to [**[]VatType**](VatType.md) | Vat types list | [optional] 
**LanguagesList** | Pointer to [**[]Language**](Language.md) | Languages list | [optional] 
**PriceLists** | Pointer to [**[]PriceList**](PriceList.md) | Price lists | [optional] 

## Methods

### NewIssuedDocumentPreCreateInfo

`func NewIssuedDocumentPreCreateInfo() *IssuedDocumentPreCreateInfo`

NewIssuedDocumentPreCreateInfo instantiates a new IssuedDocumentPreCreateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentPreCreateInfoWithDefaults

`func NewIssuedDocumentPreCreateInfoWithDefaults() *IssuedDocumentPreCreateInfo`

NewIssuedDocumentPreCreateInfoWithDefaults instantiates a new IssuedDocumentPreCreateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumerations

`func (o *IssuedDocumentPreCreateInfo) GetNumerations() map[string]map[string]int32`

GetNumerations returns the Numerations field if non-nil, zero value otherwise.

### GetNumerationsOk

`func (o *IssuedDocumentPreCreateInfo) GetNumerationsOk() (*map[string]map[string]int32, bool)`

GetNumerationsOk returns a tuple with the Numerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerations

`func (o *IssuedDocumentPreCreateInfo) SetNumerations(v map[string]map[string]int32)`

SetNumerations sets Numerations field to given value.

### HasNumerations

`func (o *IssuedDocumentPreCreateInfo) HasNumerations() bool`

HasNumerations returns a boolean if a field has been set.

### GetDnNumerations

`func (o *IssuedDocumentPreCreateInfo) GetDnNumerations() map[string]map[string]int32`

GetDnNumerations returns the DnNumerations field if non-nil, zero value otherwise.

### GetDnNumerationsOk

`func (o *IssuedDocumentPreCreateInfo) GetDnNumerationsOk() (*map[string]map[string]int32, bool)`

GetDnNumerationsOk returns a tuple with the DnNumerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnNumerations

`func (o *IssuedDocumentPreCreateInfo) SetDnNumerations(v map[string]map[string]int32)`

SetDnNumerations sets DnNumerations field to given value.

### HasDnNumerations

`func (o *IssuedDocumentPreCreateInfo) HasDnNumerations() bool`

HasDnNumerations returns a boolean if a field has been set.

### GetDefaultValues

`func (o *IssuedDocumentPreCreateInfo) GetDefaultValues() IssuedDocumentPreCreateInfoDefaultValues`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *IssuedDocumentPreCreateInfo) GetDefaultValuesOk() (*IssuedDocumentPreCreateInfoDefaultValues, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *IssuedDocumentPreCreateInfo) SetDefaultValues(v IssuedDocumentPreCreateInfoDefaultValues)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *IssuedDocumentPreCreateInfo) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### SetDefaultValuesNil

`func (o *IssuedDocumentPreCreateInfo) SetDefaultValuesNil(b bool)`

 SetDefaultValuesNil sets the value for DefaultValues to be an explicit nil

### UnsetDefaultValues
`func (o *IssuedDocumentPreCreateInfo) UnsetDefaultValues()`

UnsetDefaultValues ensures that no value is present for DefaultValues, not even an explicit nil
### GetExtraDataDefaultValues

`func (o *IssuedDocumentPreCreateInfo) GetExtraDataDefaultValues() IssuedDocumentPreCreateInfoExtraDataDefaultValues`

GetExtraDataDefaultValues returns the ExtraDataDefaultValues field if non-nil, zero value otherwise.

### GetExtraDataDefaultValuesOk

`func (o *IssuedDocumentPreCreateInfo) GetExtraDataDefaultValuesOk() (*IssuedDocumentPreCreateInfoExtraDataDefaultValues, bool)`

GetExtraDataDefaultValuesOk returns a tuple with the ExtraDataDefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDataDefaultValues

`func (o *IssuedDocumentPreCreateInfo) SetExtraDataDefaultValues(v IssuedDocumentPreCreateInfoExtraDataDefaultValues)`

SetExtraDataDefaultValues sets ExtraDataDefaultValues field to given value.

### HasExtraDataDefaultValues

`func (o *IssuedDocumentPreCreateInfo) HasExtraDataDefaultValues() bool`

HasExtraDataDefaultValues returns a boolean if a field has been set.

### SetExtraDataDefaultValuesNil

`func (o *IssuedDocumentPreCreateInfo) SetExtraDataDefaultValuesNil(b bool)`

 SetExtraDataDefaultValuesNil sets the value for ExtraDataDefaultValues to be an explicit nil

### UnsetExtraDataDefaultValues
`func (o *IssuedDocumentPreCreateInfo) UnsetExtraDataDefaultValues()`

UnsetExtraDataDefaultValues ensures that no value is present for ExtraDataDefaultValues, not even an explicit nil
### GetItemsDefaultValues

`func (o *IssuedDocumentPreCreateInfo) GetItemsDefaultValues() IssuedDocumentPreCreateInfoItemsDefaultValues`

GetItemsDefaultValues returns the ItemsDefaultValues field if non-nil, zero value otherwise.

### GetItemsDefaultValuesOk

`func (o *IssuedDocumentPreCreateInfo) GetItemsDefaultValuesOk() (*IssuedDocumentPreCreateInfoItemsDefaultValues, bool)`

GetItemsDefaultValuesOk returns a tuple with the ItemsDefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsDefaultValues

`func (o *IssuedDocumentPreCreateInfo) SetItemsDefaultValues(v IssuedDocumentPreCreateInfoItemsDefaultValues)`

SetItemsDefaultValues sets ItemsDefaultValues field to given value.

### HasItemsDefaultValues

`func (o *IssuedDocumentPreCreateInfo) HasItemsDefaultValues() bool`

HasItemsDefaultValues returns a boolean if a field has been set.

### SetItemsDefaultValuesNil

`func (o *IssuedDocumentPreCreateInfo) SetItemsDefaultValuesNil(b bool)`

 SetItemsDefaultValuesNil sets the value for ItemsDefaultValues to be an explicit nil

### UnsetItemsDefaultValues
`func (o *IssuedDocumentPreCreateInfo) UnsetItemsDefaultValues()`

UnsetItemsDefaultValues ensures that no value is present for ItemsDefaultValues, not even an explicit nil
### GetCountriesList

`func (o *IssuedDocumentPreCreateInfo) GetCountriesList() []string`

GetCountriesList returns the CountriesList field if non-nil, zero value otherwise.

### GetCountriesListOk

`func (o *IssuedDocumentPreCreateInfo) GetCountriesListOk() (*[]string, bool)`

GetCountriesListOk returns a tuple with the CountriesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountriesList

`func (o *IssuedDocumentPreCreateInfo) SetCountriesList(v []string)`

SetCountriesList sets CountriesList field to given value.

### HasCountriesList

`func (o *IssuedDocumentPreCreateInfo) HasCountriesList() bool`

HasCountriesList returns a boolean if a field has been set.

### SetCountriesListNil

`func (o *IssuedDocumentPreCreateInfo) SetCountriesListNil(b bool)`

 SetCountriesListNil sets the value for CountriesList to be an explicit nil

### UnsetCountriesList
`func (o *IssuedDocumentPreCreateInfo) UnsetCountriesList()`

UnsetCountriesList ensures that no value is present for CountriesList, not even an explicit nil
### GetCurrenciesList

`func (o *IssuedDocumentPreCreateInfo) GetCurrenciesList() []Currency`

GetCurrenciesList returns the CurrenciesList field if non-nil, zero value otherwise.

### GetCurrenciesListOk

`func (o *IssuedDocumentPreCreateInfo) GetCurrenciesListOk() (*[]Currency, bool)`

GetCurrenciesListOk returns a tuple with the CurrenciesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrenciesList

`func (o *IssuedDocumentPreCreateInfo) SetCurrenciesList(v []Currency)`

SetCurrenciesList sets CurrenciesList field to given value.

### HasCurrenciesList

`func (o *IssuedDocumentPreCreateInfo) HasCurrenciesList() bool`

HasCurrenciesList returns a boolean if a field has been set.

### SetCurrenciesListNil

`func (o *IssuedDocumentPreCreateInfo) SetCurrenciesListNil(b bool)`

 SetCurrenciesListNil sets the value for CurrenciesList to be an explicit nil

### UnsetCurrenciesList
`func (o *IssuedDocumentPreCreateInfo) UnsetCurrenciesList()`

UnsetCurrenciesList ensures that no value is present for CurrenciesList, not even an explicit nil
### GetTemplatesList

`func (o *IssuedDocumentPreCreateInfo) GetTemplatesList() []DocumentTemplate`

GetTemplatesList returns the TemplatesList field if non-nil, zero value otherwise.

### GetTemplatesListOk

`func (o *IssuedDocumentPreCreateInfo) GetTemplatesListOk() (*[]DocumentTemplate, bool)`

GetTemplatesListOk returns a tuple with the TemplatesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatesList

`func (o *IssuedDocumentPreCreateInfo) SetTemplatesList(v []DocumentTemplate)`

SetTemplatesList sets TemplatesList field to given value.

### HasTemplatesList

`func (o *IssuedDocumentPreCreateInfo) HasTemplatesList() bool`

HasTemplatesList returns a boolean if a field has been set.

### SetTemplatesListNil

`func (o *IssuedDocumentPreCreateInfo) SetTemplatesListNil(b bool)`

 SetTemplatesListNil sets the value for TemplatesList to be an explicit nil

### UnsetTemplatesList
`func (o *IssuedDocumentPreCreateInfo) UnsetTemplatesList()`

UnsetTemplatesList ensures that no value is present for TemplatesList, not even an explicit nil
### GetDnTemplatesList

`func (o *IssuedDocumentPreCreateInfo) GetDnTemplatesList() []DocumentTemplate`

GetDnTemplatesList returns the DnTemplatesList field if non-nil, zero value otherwise.

### GetDnTemplatesListOk

`func (o *IssuedDocumentPreCreateInfo) GetDnTemplatesListOk() (*[]DocumentTemplate, bool)`

GetDnTemplatesListOk returns a tuple with the DnTemplatesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnTemplatesList

`func (o *IssuedDocumentPreCreateInfo) SetDnTemplatesList(v []DocumentTemplate)`

SetDnTemplatesList sets DnTemplatesList field to given value.

### HasDnTemplatesList

`func (o *IssuedDocumentPreCreateInfo) HasDnTemplatesList() bool`

HasDnTemplatesList returns a boolean if a field has been set.

### SetDnTemplatesListNil

`func (o *IssuedDocumentPreCreateInfo) SetDnTemplatesListNil(b bool)`

 SetDnTemplatesListNil sets the value for DnTemplatesList to be an explicit nil

### UnsetDnTemplatesList
`func (o *IssuedDocumentPreCreateInfo) UnsetDnTemplatesList()`

UnsetDnTemplatesList ensures that no value is present for DnTemplatesList, not even an explicit nil
### GetAiTemplatesList

`func (o *IssuedDocumentPreCreateInfo) GetAiTemplatesList() []DocumentTemplate`

GetAiTemplatesList returns the AiTemplatesList field if non-nil, zero value otherwise.

### GetAiTemplatesListOk

`func (o *IssuedDocumentPreCreateInfo) GetAiTemplatesListOk() (*[]DocumentTemplate, bool)`

GetAiTemplatesListOk returns a tuple with the AiTemplatesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiTemplatesList

`func (o *IssuedDocumentPreCreateInfo) SetAiTemplatesList(v []DocumentTemplate)`

SetAiTemplatesList sets AiTemplatesList field to given value.

### HasAiTemplatesList

`func (o *IssuedDocumentPreCreateInfo) HasAiTemplatesList() bool`

HasAiTemplatesList returns a boolean if a field has been set.

### SetAiTemplatesListNil

`func (o *IssuedDocumentPreCreateInfo) SetAiTemplatesListNil(b bool)`

 SetAiTemplatesListNil sets the value for AiTemplatesList to be an explicit nil

### UnsetAiTemplatesList
`func (o *IssuedDocumentPreCreateInfo) UnsetAiTemplatesList()`

UnsetAiTemplatesList ensures that no value is present for AiTemplatesList, not even an explicit nil
### GetPaymentMethodsList

`func (o *IssuedDocumentPreCreateInfo) GetPaymentMethodsList() []PaymentMethod`

GetPaymentMethodsList returns the PaymentMethodsList field if non-nil, zero value otherwise.

### GetPaymentMethodsListOk

`func (o *IssuedDocumentPreCreateInfo) GetPaymentMethodsListOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsListOk returns a tuple with the PaymentMethodsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodsList

`func (o *IssuedDocumentPreCreateInfo) SetPaymentMethodsList(v []PaymentMethod)`

SetPaymentMethodsList sets PaymentMethodsList field to given value.

### HasPaymentMethodsList

`func (o *IssuedDocumentPreCreateInfo) HasPaymentMethodsList() bool`

HasPaymentMethodsList returns a boolean if a field has been set.

### SetPaymentMethodsListNil

`func (o *IssuedDocumentPreCreateInfo) SetPaymentMethodsListNil(b bool)`

 SetPaymentMethodsListNil sets the value for PaymentMethodsList to be an explicit nil

### UnsetPaymentMethodsList
`func (o *IssuedDocumentPreCreateInfo) UnsetPaymentMethodsList()`

UnsetPaymentMethodsList ensures that no value is present for PaymentMethodsList, not even an explicit nil
### GetPaymentAccountsList

`func (o *IssuedDocumentPreCreateInfo) GetPaymentAccountsList() []PaymentAccount`

GetPaymentAccountsList returns the PaymentAccountsList field if non-nil, zero value otherwise.

### GetPaymentAccountsListOk

`func (o *IssuedDocumentPreCreateInfo) GetPaymentAccountsListOk() (*[]PaymentAccount, bool)`

GetPaymentAccountsListOk returns a tuple with the PaymentAccountsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountsList

`func (o *IssuedDocumentPreCreateInfo) SetPaymentAccountsList(v []PaymentAccount)`

SetPaymentAccountsList sets PaymentAccountsList field to given value.

### HasPaymentAccountsList

`func (o *IssuedDocumentPreCreateInfo) HasPaymentAccountsList() bool`

HasPaymentAccountsList returns a boolean if a field has been set.

### SetPaymentAccountsListNil

`func (o *IssuedDocumentPreCreateInfo) SetPaymentAccountsListNil(b bool)`

 SetPaymentAccountsListNil sets the value for PaymentAccountsList to be an explicit nil

### UnsetPaymentAccountsList
`func (o *IssuedDocumentPreCreateInfo) UnsetPaymentAccountsList()`

UnsetPaymentAccountsList ensures that no value is present for PaymentAccountsList, not even an explicit nil
### GetVatTypesList

`func (o *IssuedDocumentPreCreateInfo) GetVatTypesList() []VatType`

GetVatTypesList returns the VatTypesList field if non-nil, zero value otherwise.

### GetVatTypesListOk

`func (o *IssuedDocumentPreCreateInfo) GetVatTypesListOk() (*[]VatType, bool)`

GetVatTypesListOk returns a tuple with the VatTypesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatTypesList

`func (o *IssuedDocumentPreCreateInfo) SetVatTypesList(v []VatType)`

SetVatTypesList sets VatTypesList field to given value.

### HasVatTypesList

`func (o *IssuedDocumentPreCreateInfo) HasVatTypesList() bool`

HasVatTypesList returns a boolean if a field has been set.

### SetVatTypesListNil

`func (o *IssuedDocumentPreCreateInfo) SetVatTypesListNil(b bool)`

 SetVatTypesListNil sets the value for VatTypesList to be an explicit nil

### UnsetVatTypesList
`func (o *IssuedDocumentPreCreateInfo) UnsetVatTypesList()`

UnsetVatTypesList ensures that no value is present for VatTypesList, not even an explicit nil
### GetLanguagesList

`func (o *IssuedDocumentPreCreateInfo) GetLanguagesList() []Language`

GetLanguagesList returns the LanguagesList field if non-nil, zero value otherwise.

### GetLanguagesListOk

`func (o *IssuedDocumentPreCreateInfo) GetLanguagesListOk() (*[]Language, bool)`

GetLanguagesListOk returns a tuple with the LanguagesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesList

`func (o *IssuedDocumentPreCreateInfo) SetLanguagesList(v []Language)`

SetLanguagesList sets LanguagesList field to given value.

### HasLanguagesList

`func (o *IssuedDocumentPreCreateInfo) HasLanguagesList() bool`

HasLanguagesList returns a boolean if a field has been set.

### SetLanguagesListNil

`func (o *IssuedDocumentPreCreateInfo) SetLanguagesListNil(b bool)`

 SetLanguagesListNil sets the value for LanguagesList to be an explicit nil

### UnsetLanguagesList
`func (o *IssuedDocumentPreCreateInfo) UnsetLanguagesList()`

UnsetLanguagesList ensures that no value is present for LanguagesList, not even an explicit nil
### GetPriceLists

`func (o *IssuedDocumentPreCreateInfo) GetPriceLists() []PriceList`

GetPriceLists returns the PriceLists field if non-nil, zero value otherwise.

### GetPriceListsOk

`func (o *IssuedDocumentPreCreateInfo) GetPriceListsOk() (*[]PriceList, bool)`

GetPriceListsOk returns a tuple with the PriceLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLists

`func (o *IssuedDocumentPreCreateInfo) SetPriceLists(v []PriceList)`

SetPriceLists sets PriceLists field to given value.

### HasPriceLists

`func (o *IssuedDocumentPreCreateInfo) HasPriceLists() bool`

HasPriceLists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


