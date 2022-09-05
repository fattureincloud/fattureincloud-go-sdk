# ReceivedDocumentItemsListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique identifier. | [optional] 
**ProductId** | Pointer to **NullableInt32** | Unique identifier of the product | [optional] 
**Code** | Pointer to **NullableString** | Product code. | [optional] 
**Name** | Pointer to **NullableString** | Product name. | [optional] 
**Measure** | Pointer to **NullableString** | Product measure. | [optional] 
**NetPrice** | Pointer to **NullableFloat32** | Product net price. | [optional] 
**Category** | Pointer to **NullableString** | Product category. | [optional] 
**Qty** | Pointer to **NullableFloat32** | Product quantity. | [optional] 
**Vat** | Pointer to [**NullableVatType**](VatType.md) |  | [optional] 
**Stock** | Pointer to **NullableFloat32** | Number of items in stock | [optional] 

## Methods

### NewReceivedDocumentItemsListItem

`func NewReceivedDocumentItemsListItem() *ReceivedDocumentItemsListItem`

NewReceivedDocumentItemsListItem instantiates a new ReceivedDocumentItemsListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivedDocumentItemsListItemWithDefaults

`func NewReceivedDocumentItemsListItemWithDefaults() *ReceivedDocumentItemsListItem`

NewReceivedDocumentItemsListItemWithDefaults instantiates a new ReceivedDocumentItemsListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReceivedDocumentItemsListItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceivedDocumentItemsListItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceivedDocumentItemsListItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReceivedDocumentItemsListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ReceivedDocumentItemsListItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ReceivedDocumentItemsListItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProductId

`func (o *ReceivedDocumentItemsListItem) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ReceivedDocumentItemsListItem) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ReceivedDocumentItemsListItem) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ReceivedDocumentItemsListItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *ReceivedDocumentItemsListItem) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *ReceivedDocumentItemsListItem) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetCode

`func (o *ReceivedDocumentItemsListItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReceivedDocumentItemsListItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReceivedDocumentItemsListItem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReceivedDocumentItemsListItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ReceivedDocumentItemsListItem) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ReceivedDocumentItemsListItem) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *ReceivedDocumentItemsListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReceivedDocumentItemsListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReceivedDocumentItemsListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReceivedDocumentItemsListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ReceivedDocumentItemsListItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ReceivedDocumentItemsListItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMeasure

`func (o *ReceivedDocumentItemsListItem) GetMeasure() string`

GetMeasure returns the Measure field if non-nil, zero value otherwise.

### GetMeasureOk

`func (o *ReceivedDocumentItemsListItem) GetMeasureOk() (*string, bool)`

GetMeasureOk returns a tuple with the Measure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasure

`func (o *ReceivedDocumentItemsListItem) SetMeasure(v string)`

SetMeasure sets Measure field to given value.

### HasMeasure

`func (o *ReceivedDocumentItemsListItem) HasMeasure() bool`

HasMeasure returns a boolean if a field has been set.

### SetMeasureNil

`func (o *ReceivedDocumentItemsListItem) SetMeasureNil(b bool)`

 SetMeasureNil sets the value for Measure to be an explicit nil

### UnsetMeasure
`func (o *ReceivedDocumentItemsListItem) UnsetMeasure()`

UnsetMeasure ensures that no value is present for Measure, not even an explicit nil
### GetNetPrice

`func (o *ReceivedDocumentItemsListItem) GetNetPrice() float32`

GetNetPrice returns the NetPrice field if non-nil, zero value otherwise.

### GetNetPriceOk

`func (o *ReceivedDocumentItemsListItem) GetNetPriceOk() (*float32, bool)`

GetNetPriceOk returns a tuple with the NetPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPrice

`func (o *ReceivedDocumentItemsListItem) SetNetPrice(v float32)`

SetNetPrice sets NetPrice field to given value.

### HasNetPrice

`func (o *ReceivedDocumentItemsListItem) HasNetPrice() bool`

HasNetPrice returns a boolean if a field has been set.

### SetNetPriceNil

`func (o *ReceivedDocumentItemsListItem) SetNetPriceNil(b bool)`

 SetNetPriceNil sets the value for NetPrice to be an explicit nil

### UnsetNetPrice
`func (o *ReceivedDocumentItemsListItem) UnsetNetPrice()`

UnsetNetPrice ensures that no value is present for NetPrice, not even an explicit nil
### GetCategory

`func (o *ReceivedDocumentItemsListItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ReceivedDocumentItemsListItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ReceivedDocumentItemsListItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ReceivedDocumentItemsListItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ReceivedDocumentItemsListItem) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ReceivedDocumentItemsListItem) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetQty

`func (o *ReceivedDocumentItemsListItem) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *ReceivedDocumentItemsListItem) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *ReceivedDocumentItemsListItem) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *ReceivedDocumentItemsListItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### SetQtyNil

`func (o *ReceivedDocumentItemsListItem) SetQtyNil(b bool)`

 SetQtyNil sets the value for Qty to be an explicit nil

### UnsetQty
`func (o *ReceivedDocumentItemsListItem) UnsetQty()`

UnsetQty ensures that no value is present for Qty, not even an explicit nil
### GetVat

`func (o *ReceivedDocumentItemsListItem) GetVat() VatType`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *ReceivedDocumentItemsListItem) GetVatOk() (*VatType, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *ReceivedDocumentItemsListItem) SetVat(v VatType)`

SetVat sets Vat field to given value.

### HasVat

`func (o *ReceivedDocumentItemsListItem) HasVat() bool`

HasVat returns a boolean if a field has been set.

### SetVatNil

`func (o *ReceivedDocumentItemsListItem) SetVatNil(b bool)`

 SetVatNil sets the value for Vat to be an explicit nil

### UnsetVat
`func (o *ReceivedDocumentItemsListItem) UnsetVat()`

UnsetVat ensures that no value is present for Vat, not even an explicit nil
### GetStock

`func (o *ReceivedDocumentItemsListItem) GetStock() float32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ReceivedDocumentItemsListItem) GetStockOk() (*float32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ReceivedDocumentItemsListItem) SetStock(v float32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ReceivedDocumentItemsListItem) HasStock() bool`

HasStock returns a boolean if a field has been set.

### SetStockNil

`func (o *ReceivedDocumentItemsListItem) SetStockNil(b bool)`

 SetStockNil sets the value for Stock to be an explicit nil

### UnsetStock
`func (o *ReceivedDocumentItemsListItem) UnsetStock()`

UnsetStock ensures that no value is present for Stock, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


