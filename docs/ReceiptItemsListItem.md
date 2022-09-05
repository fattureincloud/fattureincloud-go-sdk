# ReceiptItemsListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Item unique identifier. | [optional] 
**AmountNet** | Pointer to **NullableFloat32** | Item total net amount. | [optional] 
**AmountGross** | Pointer to **NullableFloat32** | Item total gross amount. | [optional] 
**Category** | Pointer to **NullableString** | Item category. | [optional] 
**Vat** | Pointer to [**NullableVatType**](VatType.md) |  | [optional] 

## Methods

### NewReceiptItemsListItem

`func NewReceiptItemsListItem() *ReceiptItemsListItem`

NewReceiptItemsListItem instantiates a new ReceiptItemsListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptItemsListItemWithDefaults

`func NewReceiptItemsListItemWithDefaults() *ReceiptItemsListItem`

NewReceiptItemsListItemWithDefaults instantiates a new ReceiptItemsListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReceiptItemsListItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceiptItemsListItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceiptItemsListItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReceiptItemsListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ReceiptItemsListItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ReceiptItemsListItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAmountNet

`func (o *ReceiptItemsListItem) GetAmountNet() float32`

GetAmountNet returns the AmountNet field if non-nil, zero value otherwise.

### GetAmountNetOk

`func (o *ReceiptItemsListItem) GetAmountNetOk() (*float32, bool)`

GetAmountNetOk returns a tuple with the AmountNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountNet

`func (o *ReceiptItemsListItem) SetAmountNet(v float32)`

SetAmountNet sets AmountNet field to given value.

### HasAmountNet

`func (o *ReceiptItemsListItem) HasAmountNet() bool`

HasAmountNet returns a boolean if a field has been set.

### SetAmountNetNil

`func (o *ReceiptItemsListItem) SetAmountNetNil(b bool)`

 SetAmountNetNil sets the value for AmountNet to be an explicit nil

### UnsetAmountNet
`func (o *ReceiptItemsListItem) UnsetAmountNet()`

UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
### GetAmountGross

`func (o *ReceiptItemsListItem) GetAmountGross() float32`

GetAmountGross returns the AmountGross field if non-nil, zero value otherwise.

### GetAmountGrossOk

`func (o *ReceiptItemsListItem) GetAmountGrossOk() (*float32, bool)`

GetAmountGrossOk returns a tuple with the AmountGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGross

`func (o *ReceiptItemsListItem) SetAmountGross(v float32)`

SetAmountGross sets AmountGross field to given value.

### HasAmountGross

`func (o *ReceiptItemsListItem) HasAmountGross() bool`

HasAmountGross returns a boolean if a field has been set.

### SetAmountGrossNil

`func (o *ReceiptItemsListItem) SetAmountGrossNil(b bool)`

 SetAmountGrossNil sets the value for AmountGross to be an explicit nil

### UnsetAmountGross
`func (o *ReceiptItemsListItem) UnsetAmountGross()`

UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
### GetCategory

`func (o *ReceiptItemsListItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ReceiptItemsListItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ReceiptItemsListItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ReceiptItemsListItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ReceiptItemsListItem) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ReceiptItemsListItem) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVat

`func (o *ReceiptItemsListItem) GetVat() VatType`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *ReceiptItemsListItem) GetVatOk() (*VatType, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *ReceiptItemsListItem) SetVat(v VatType)`

SetVat sets Vat field to given value.

### HasVat

`func (o *ReceiptItemsListItem) HasVat() bool`

HasVat returns a boolean if a field has been set.

### SetVatNil

`func (o *ReceiptItemsListItem) SetVatNil(b bool)`

 SetVatNil sets the value for Vat to be an explicit nil

### UnsetVat
`func (o *ReceiptItemsListItem) UnsetVat()`

UnsetVat ensures that no value is present for Vat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


