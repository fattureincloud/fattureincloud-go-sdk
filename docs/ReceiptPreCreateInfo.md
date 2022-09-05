# ReceiptPreCreateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numerations** | Pointer to **map[string]map[string]int32** |  | [optional] 
**NumerationsList** | Pointer to **[]string** | List of series used in the past. | [optional] 
**RcCentersList** | Pointer to **[]string** | List of revenue centers used in the past. | [optional] 
**PaymentAccountsList** | Pointer to [**[]PaymentAccount**](PaymentAccount.md) | User payment accounts list. | [optional] 
**CategoriesList** | Pointer to **[]string** | List of categories used in the past. | [optional] 
**VatTypesList** | Pointer to [**[]VatType**](VatType.md) | List of user vat types with the default 22%, 10%, 4% and 0% vats. | [optional] 

## Methods

### NewReceiptPreCreateInfo

`func NewReceiptPreCreateInfo() *ReceiptPreCreateInfo`

NewReceiptPreCreateInfo instantiates a new ReceiptPreCreateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptPreCreateInfoWithDefaults

`func NewReceiptPreCreateInfoWithDefaults() *ReceiptPreCreateInfo`

NewReceiptPreCreateInfoWithDefaults instantiates a new ReceiptPreCreateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumerations

`func (o *ReceiptPreCreateInfo) GetNumerations() map[string]map[string]int32`

GetNumerations returns the Numerations field if non-nil, zero value otherwise.

### GetNumerationsOk

`func (o *ReceiptPreCreateInfo) GetNumerationsOk() (*map[string]map[string]int32, bool)`

GetNumerationsOk returns a tuple with the Numerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerations

`func (o *ReceiptPreCreateInfo) SetNumerations(v map[string]map[string]int32)`

SetNumerations sets Numerations field to given value.

### HasNumerations

`func (o *ReceiptPreCreateInfo) HasNumerations() bool`

HasNumerations returns a boolean if a field has been set.

### GetNumerationsList

`func (o *ReceiptPreCreateInfo) GetNumerationsList() []string`

GetNumerationsList returns the NumerationsList field if non-nil, zero value otherwise.

### GetNumerationsListOk

`func (o *ReceiptPreCreateInfo) GetNumerationsListOk() (*[]string, bool)`

GetNumerationsListOk returns a tuple with the NumerationsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerationsList

`func (o *ReceiptPreCreateInfo) SetNumerationsList(v []string)`

SetNumerationsList sets NumerationsList field to given value.

### HasNumerationsList

`func (o *ReceiptPreCreateInfo) HasNumerationsList() bool`

HasNumerationsList returns a boolean if a field has been set.

### SetNumerationsListNil

`func (o *ReceiptPreCreateInfo) SetNumerationsListNil(b bool)`

 SetNumerationsListNil sets the value for NumerationsList to be an explicit nil

### UnsetNumerationsList
`func (o *ReceiptPreCreateInfo) UnsetNumerationsList()`

UnsetNumerationsList ensures that no value is present for NumerationsList, not even an explicit nil
### GetRcCentersList

`func (o *ReceiptPreCreateInfo) GetRcCentersList() []string`

GetRcCentersList returns the RcCentersList field if non-nil, zero value otherwise.

### GetRcCentersListOk

`func (o *ReceiptPreCreateInfo) GetRcCentersListOk() (*[]string, bool)`

GetRcCentersListOk returns a tuple with the RcCentersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcCentersList

`func (o *ReceiptPreCreateInfo) SetRcCentersList(v []string)`

SetRcCentersList sets RcCentersList field to given value.

### HasRcCentersList

`func (o *ReceiptPreCreateInfo) HasRcCentersList() bool`

HasRcCentersList returns a boolean if a field has been set.

### SetRcCentersListNil

`func (o *ReceiptPreCreateInfo) SetRcCentersListNil(b bool)`

 SetRcCentersListNil sets the value for RcCentersList to be an explicit nil

### UnsetRcCentersList
`func (o *ReceiptPreCreateInfo) UnsetRcCentersList()`

UnsetRcCentersList ensures that no value is present for RcCentersList, not even an explicit nil
### GetPaymentAccountsList

`func (o *ReceiptPreCreateInfo) GetPaymentAccountsList() []PaymentAccount`

GetPaymentAccountsList returns the PaymentAccountsList field if non-nil, zero value otherwise.

### GetPaymentAccountsListOk

`func (o *ReceiptPreCreateInfo) GetPaymentAccountsListOk() (*[]PaymentAccount, bool)`

GetPaymentAccountsListOk returns a tuple with the PaymentAccountsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountsList

`func (o *ReceiptPreCreateInfo) SetPaymentAccountsList(v []PaymentAccount)`

SetPaymentAccountsList sets PaymentAccountsList field to given value.

### HasPaymentAccountsList

`func (o *ReceiptPreCreateInfo) HasPaymentAccountsList() bool`

HasPaymentAccountsList returns a boolean if a field has been set.

### SetPaymentAccountsListNil

`func (o *ReceiptPreCreateInfo) SetPaymentAccountsListNil(b bool)`

 SetPaymentAccountsListNil sets the value for PaymentAccountsList to be an explicit nil

### UnsetPaymentAccountsList
`func (o *ReceiptPreCreateInfo) UnsetPaymentAccountsList()`

UnsetPaymentAccountsList ensures that no value is present for PaymentAccountsList, not even an explicit nil
### GetCategoriesList

`func (o *ReceiptPreCreateInfo) GetCategoriesList() []string`

GetCategoriesList returns the CategoriesList field if non-nil, zero value otherwise.

### GetCategoriesListOk

`func (o *ReceiptPreCreateInfo) GetCategoriesListOk() (*[]string, bool)`

GetCategoriesListOk returns a tuple with the CategoriesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesList

`func (o *ReceiptPreCreateInfo) SetCategoriesList(v []string)`

SetCategoriesList sets CategoriesList field to given value.

### HasCategoriesList

`func (o *ReceiptPreCreateInfo) HasCategoriesList() bool`

HasCategoriesList returns a boolean if a field has been set.

### SetCategoriesListNil

`func (o *ReceiptPreCreateInfo) SetCategoriesListNil(b bool)`

 SetCategoriesListNil sets the value for CategoriesList to be an explicit nil

### UnsetCategoriesList
`func (o *ReceiptPreCreateInfo) UnsetCategoriesList()`

UnsetCategoriesList ensures that no value is present for CategoriesList, not even an explicit nil
### GetVatTypesList

`func (o *ReceiptPreCreateInfo) GetVatTypesList() []VatType`

GetVatTypesList returns the VatTypesList field if non-nil, zero value otherwise.

### GetVatTypesListOk

`func (o *ReceiptPreCreateInfo) GetVatTypesListOk() (*[]VatType, bool)`

GetVatTypesListOk returns a tuple with the VatTypesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatTypesList

`func (o *ReceiptPreCreateInfo) SetVatTypesList(v []VatType)`

SetVatTypesList sets VatTypesList field to given value.

### HasVatTypesList

`func (o *ReceiptPreCreateInfo) HasVatTypesList() bool`

HasVatTypesList returns a boolean if a field has been set.

### SetVatTypesListNil

`func (o *ReceiptPreCreateInfo) SetVatTypesListNil(b bool)`

 SetVatTypesListNil sets the value for VatTypesList to be an explicit nil

### UnsetVatTypesList
`func (o *ReceiptPreCreateInfo) UnsetVatTypesList()`

UnsetVatTypesList ensures that no value is present for VatTypesList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


