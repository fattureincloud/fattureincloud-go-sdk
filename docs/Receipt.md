# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Receipt id | [optional] 
**Date** | Pointer to **NullableString** | Receipt date | [optional] 
**Number** | Pointer to **NullableFloat32** | Receipt number | [optional] 
**Numeration** | Pointer to **NullableString** | Receipt numeration | [optional] 
**AmountNet** | Pointer to **NullableFloat32** | Receipt total net amount | [optional] 
**AmountVat** | Pointer to **NullableFloat32** | Receipt total vat amount | [optional] 
**AmountGross** | Pointer to **NullableFloat32** | Receipt total gross amount | [optional] 
**UseGrossPrices** | Pointer to **NullableBool** | Receipt uses gross prices | [optional] 
**Type** | Pointer to [**ReceiptType**](ReceiptType.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Receipt description | [optional] 
**RcCenter** | Pointer to **NullableString** | Receipt revenue center | [optional] 
**CreatedAt** | Pointer to **NullableString** | Receipt creation date | [optional] 
**UpdatedAt** | Pointer to **NullableString** | Receipt last update date | [optional] 
**PaymentAccount** | Pointer to [**NullablePaymentAccount**](PaymentAccount.md) |  | [optional] 
**ItemsList** | Pointer to [**[]ReceiptItemsListItem**](ReceiptItemsListItem.md) |  | [optional] 

## Methods

### NewReceipt

`func NewReceipt() *Receipt`

NewReceipt instantiates a new Receipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptWithDefaults

`func NewReceiptWithDefaults() *Receipt`

NewReceiptWithDefaults instantiates a new Receipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Receipt) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Receipt) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Receipt) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Receipt) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Receipt) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Receipt) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDate

`func (o *Receipt) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Receipt) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Receipt) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Receipt) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *Receipt) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *Receipt) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetNumber

`func (o *Receipt) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Receipt) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Receipt) SetNumber(v float32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Receipt) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *Receipt) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *Receipt) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetNumeration

`func (o *Receipt) GetNumeration() string`

GetNumeration returns the Numeration field if non-nil, zero value otherwise.

### GetNumerationOk

`func (o *Receipt) GetNumerationOk() (*string, bool)`

GetNumerationOk returns a tuple with the Numeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeration

`func (o *Receipt) SetNumeration(v string)`

SetNumeration sets Numeration field to given value.

### HasNumeration

`func (o *Receipt) HasNumeration() bool`

HasNumeration returns a boolean if a field has been set.

### SetNumerationNil

`func (o *Receipt) SetNumerationNil(b bool)`

 SetNumerationNil sets the value for Numeration to be an explicit nil

### UnsetNumeration
`func (o *Receipt) UnsetNumeration()`

UnsetNumeration ensures that no value is present for Numeration, not even an explicit nil
### GetAmountNet

`func (o *Receipt) GetAmountNet() float32`

GetAmountNet returns the AmountNet field if non-nil, zero value otherwise.

### GetAmountNetOk

`func (o *Receipt) GetAmountNetOk() (*float32, bool)`

GetAmountNetOk returns a tuple with the AmountNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountNet

`func (o *Receipt) SetAmountNet(v float32)`

SetAmountNet sets AmountNet field to given value.

### HasAmountNet

`func (o *Receipt) HasAmountNet() bool`

HasAmountNet returns a boolean if a field has been set.

### SetAmountNetNil

`func (o *Receipt) SetAmountNetNil(b bool)`

 SetAmountNetNil sets the value for AmountNet to be an explicit nil

### UnsetAmountNet
`func (o *Receipt) UnsetAmountNet()`

UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
### GetAmountVat

`func (o *Receipt) GetAmountVat() float32`

GetAmountVat returns the AmountVat field if non-nil, zero value otherwise.

### GetAmountVatOk

`func (o *Receipt) GetAmountVatOk() (*float32, bool)`

GetAmountVatOk returns a tuple with the AmountVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountVat

`func (o *Receipt) SetAmountVat(v float32)`

SetAmountVat sets AmountVat field to given value.

### HasAmountVat

`func (o *Receipt) HasAmountVat() bool`

HasAmountVat returns a boolean if a field has been set.

### SetAmountVatNil

`func (o *Receipt) SetAmountVatNil(b bool)`

 SetAmountVatNil sets the value for AmountVat to be an explicit nil

### UnsetAmountVat
`func (o *Receipt) UnsetAmountVat()`

UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
### GetAmountGross

`func (o *Receipt) GetAmountGross() float32`

GetAmountGross returns the AmountGross field if non-nil, zero value otherwise.

### GetAmountGrossOk

`func (o *Receipt) GetAmountGrossOk() (*float32, bool)`

GetAmountGrossOk returns a tuple with the AmountGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGross

`func (o *Receipt) SetAmountGross(v float32)`

SetAmountGross sets AmountGross field to given value.

### HasAmountGross

`func (o *Receipt) HasAmountGross() bool`

HasAmountGross returns a boolean if a field has been set.

### SetAmountGrossNil

`func (o *Receipt) SetAmountGrossNil(b bool)`

 SetAmountGrossNil sets the value for AmountGross to be an explicit nil

### UnsetAmountGross
`func (o *Receipt) UnsetAmountGross()`

UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
### GetUseGrossPrices

`func (o *Receipt) GetUseGrossPrices() bool`

GetUseGrossPrices returns the UseGrossPrices field if non-nil, zero value otherwise.

### GetUseGrossPricesOk

`func (o *Receipt) GetUseGrossPricesOk() (*bool, bool)`

GetUseGrossPricesOk returns a tuple with the UseGrossPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGrossPrices

`func (o *Receipt) SetUseGrossPrices(v bool)`

SetUseGrossPrices sets UseGrossPrices field to given value.

### HasUseGrossPrices

`func (o *Receipt) HasUseGrossPrices() bool`

HasUseGrossPrices returns a boolean if a field has been set.

### SetUseGrossPricesNil

`func (o *Receipt) SetUseGrossPricesNil(b bool)`

 SetUseGrossPricesNil sets the value for UseGrossPrices to be an explicit nil

### UnsetUseGrossPrices
`func (o *Receipt) UnsetUseGrossPrices()`

UnsetUseGrossPrices ensures that no value is present for UseGrossPrices, not even an explicit nil
### GetType

`func (o *Receipt) GetType() ReceiptType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Receipt) GetTypeOk() (*ReceiptType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Receipt) SetType(v ReceiptType)`

SetType sets Type field to given value.

### HasType

`func (o *Receipt) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *Receipt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Receipt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Receipt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Receipt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Receipt) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Receipt) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRcCenter

`func (o *Receipt) GetRcCenter() string`

GetRcCenter returns the RcCenter field if non-nil, zero value otherwise.

### GetRcCenterOk

`func (o *Receipt) GetRcCenterOk() (*string, bool)`

GetRcCenterOk returns a tuple with the RcCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcCenter

`func (o *Receipt) SetRcCenter(v string)`

SetRcCenter sets RcCenter field to given value.

### HasRcCenter

`func (o *Receipt) HasRcCenter() bool`

HasRcCenter returns a boolean if a field has been set.

### SetRcCenterNil

`func (o *Receipt) SetRcCenterNil(b bool)`

 SetRcCenterNil sets the value for RcCenter to be an explicit nil

### UnsetRcCenter
`func (o *Receipt) UnsetRcCenter()`

UnsetRcCenter ensures that no value is present for RcCenter, not even an explicit nil
### GetCreatedAt

`func (o *Receipt) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Receipt) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Receipt) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Receipt) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Receipt) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Receipt) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Receipt) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Receipt) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Receipt) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Receipt) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Receipt) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Receipt) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetPaymentAccount

`func (o *Receipt) GetPaymentAccount() PaymentAccount`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *Receipt) GetPaymentAccountOk() (*PaymentAccount, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *Receipt) SetPaymentAccount(v PaymentAccount)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *Receipt) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### SetPaymentAccountNil

`func (o *Receipt) SetPaymentAccountNil(b bool)`

 SetPaymentAccountNil sets the value for PaymentAccount to be an explicit nil

### UnsetPaymentAccount
`func (o *Receipt) UnsetPaymentAccount()`

UnsetPaymentAccount ensures that no value is present for PaymentAccount, not even an explicit nil
### GetItemsList

`func (o *Receipt) GetItemsList() []ReceiptItemsListItem`

GetItemsList returns the ItemsList field if non-nil, zero value otherwise.

### GetItemsListOk

`func (o *Receipt) GetItemsListOk() (*[]ReceiptItemsListItem, bool)`

GetItemsListOk returns a tuple with the ItemsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsList

`func (o *Receipt) SetItemsList(v []ReceiptItemsListItem)`

SetItemsList sets ItemsList field to given value.

### HasItemsList

`func (o *Receipt) HasItemsList() bool`

HasItemsList returns a boolean if a field has been set.

### SetItemsListNil

`func (o *Receipt) SetItemsListNil(b bool)`

 SetItemsListNil sets the value for ItemsList to be an explicit nil

### UnsetItemsList
`func (o *Receipt) UnsetItemsList()`

UnsetItemsList ensures that no value is present for ItemsList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


