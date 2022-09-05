# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique identifier. | [optional] 
**Name** | Pointer to **NullableString** | Product name. | [optional] 
**Code** | Pointer to **NullableString** | Product code. | [optional] 
**NetPrice** | Pointer to **NullableFloat32** | Net sale price (used if use_gross_price is false, otherwise it&#39;s competed automatically). | [optional] 
**GrossPrice** | Pointer to **NullableFloat32** | Gross sale price (used if use_gross_price is false, otherwise it&#39;s competed automatically). | [optional] 
**UseGrossPrice** | Pointer to **NullableBool** | Determine which price to use for calculations. | [optional] 
**DefaultVat** | Pointer to [**NullableVatType**](VatType.md) |  | [optional] 
**NetCost** | Pointer to **NullableFloat32** | Net cost of the product (used for received documents). | [optional] 
**Measure** | Pointer to **NullableString** | Unit of measure. | [optional] 
**Description** | Pointer to **NullableString** | Product description. | [optional] 
**Category** | Pointer to **NullableString** | Product category. | [optional] 
**Notes** | Pointer to **NullableString** | Extra notes. | [optional] 
**InStock** | Pointer to **NullableBool** | Determine if the product is in stock. | [optional] 
**StockInitial** | Pointer to **NullableFloat32** | Product initial stock. | [optional] 
**StockCurrent** | Pointer to **NullableFloat32** | [Read Only] Product current stock. | [optional] [readonly] 
**AverageCost** | Pointer to **NullableFloat32** | Product average cost. | [optional] 
**AveragePrice** | Pointer to **NullableFloat32** | Product average price. | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Product) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Product) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Product) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Product) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Product) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Product) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *Product) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Product) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Product) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Product) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Product) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Product) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetNetPrice

`func (o *Product) GetNetPrice() float32`

GetNetPrice returns the NetPrice field if non-nil, zero value otherwise.

### GetNetPriceOk

`func (o *Product) GetNetPriceOk() (*float32, bool)`

GetNetPriceOk returns a tuple with the NetPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPrice

`func (o *Product) SetNetPrice(v float32)`

SetNetPrice sets NetPrice field to given value.

### HasNetPrice

`func (o *Product) HasNetPrice() bool`

HasNetPrice returns a boolean if a field has been set.

### SetNetPriceNil

`func (o *Product) SetNetPriceNil(b bool)`

 SetNetPriceNil sets the value for NetPrice to be an explicit nil

### UnsetNetPrice
`func (o *Product) UnsetNetPrice()`

UnsetNetPrice ensures that no value is present for NetPrice, not even an explicit nil
### GetGrossPrice

`func (o *Product) GetGrossPrice() float32`

GetGrossPrice returns the GrossPrice field if non-nil, zero value otherwise.

### GetGrossPriceOk

`func (o *Product) GetGrossPriceOk() (*float32, bool)`

GetGrossPriceOk returns a tuple with the GrossPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossPrice

`func (o *Product) SetGrossPrice(v float32)`

SetGrossPrice sets GrossPrice field to given value.

### HasGrossPrice

`func (o *Product) HasGrossPrice() bool`

HasGrossPrice returns a boolean if a field has been set.

### SetGrossPriceNil

`func (o *Product) SetGrossPriceNil(b bool)`

 SetGrossPriceNil sets the value for GrossPrice to be an explicit nil

### UnsetGrossPrice
`func (o *Product) UnsetGrossPrice()`

UnsetGrossPrice ensures that no value is present for GrossPrice, not even an explicit nil
### GetUseGrossPrice

`func (o *Product) GetUseGrossPrice() bool`

GetUseGrossPrice returns the UseGrossPrice field if non-nil, zero value otherwise.

### GetUseGrossPriceOk

`func (o *Product) GetUseGrossPriceOk() (*bool, bool)`

GetUseGrossPriceOk returns a tuple with the UseGrossPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGrossPrice

`func (o *Product) SetUseGrossPrice(v bool)`

SetUseGrossPrice sets UseGrossPrice field to given value.

### HasUseGrossPrice

`func (o *Product) HasUseGrossPrice() bool`

HasUseGrossPrice returns a boolean if a field has been set.

### SetUseGrossPriceNil

`func (o *Product) SetUseGrossPriceNil(b bool)`

 SetUseGrossPriceNil sets the value for UseGrossPrice to be an explicit nil

### UnsetUseGrossPrice
`func (o *Product) UnsetUseGrossPrice()`

UnsetUseGrossPrice ensures that no value is present for UseGrossPrice, not even an explicit nil
### GetDefaultVat

`func (o *Product) GetDefaultVat() VatType`

GetDefaultVat returns the DefaultVat field if non-nil, zero value otherwise.

### GetDefaultVatOk

`func (o *Product) GetDefaultVatOk() (*VatType, bool)`

GetDefaultVatOk returns a tuple with the DefaultVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVat

`func (o *Product) SetDefaultVat(v VatType)`

SetDefaultVat sets DefaultVat field to given value.

### HasDefaultVat

`func (o *Product) HasDefaultVat() bool`

HasDefaultVat returns a boolean if a field has been set.

### SetDefaultVatNil

`func (o *Product) SetDefaultVatNil(b bool)`

 SetDefaultVatNil sets the value for DefaultVat to be an explicit nil

### UnsetDefaultVat
`func (o *Product) UnsetDefaultVat()`

UnsetDefaultVat ensures that no value is present for DefaultVat, not even an explicit nil
### GetNetCost

`func (o *Product) GetNetCost() float32`

GetNetCost returns the NetCost field if non-nil, zero value otherwise.

### GetNetCostOk

`func (o *Product) GetNetCostOk() (*float32, bool)`

GetNetCostOk returns a tuple with the NetCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCost

`func (o *Product) SetNetCost(v float32)`

SetNetCost sets NetCost field to given value.

### HasNetCost

`func (o *Product) HasNetCost() bool`

HasNetCost returns a boolean if a field has been set.

### SetNetCostNil

`func (o *Product) SetNetCostNil(b bool)`

 SetNetCostNil sets the value for NetCost to be an explicit nil

### UnsetNetCost
`func (o *Product) UnsetNetCost()`

UnsetNetCost ensures that no value is present for NetCost, not even an explicit nil
### GetMeasure

`func (o *Product) GetMeasure() string`

GetMeasure returns the Measure field if non-nil, zero value otherwise.

### GetMeasureOk

`func (o *Product) GetMeasureOk() (*string, bool)`

GetMeasureOk returns a tuple with the Measure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasure

`func (o *Product) SetMeasure(v string)`

SetMeasure sets Measure field to given value.

### HasMeasure

`func (o *Product) HasMeasure() bool`

HasMeasure returns a boolean if a field has been set.

### SetMeasureNil

`func (o *Product) SetMeasureNil(b bool)`

 SetMeasureNil sets the value for Measure to be an explicit nil

### UnsetMeasure
`func (o *Product) UnsetMeasure()`

UnsetMeasure ensures that no value is present for Measure, not even an explicit nil
### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Product) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Product) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *Product) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Product) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Product) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Product) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *Product) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Product) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetNotes

`func (o *Product) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Product) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Product) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Product) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Product) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Product) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetInStock

`func (o *Product) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *Product) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *Product) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *Product) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### SetInStockNil

`func (o *Product) SetInStockNil(b bool)`

 SetInStockNil sets the value for InStock to be an explicit nil

### UnsetInStock
`func (o *Product) UnsetInStock()`

UnsetInStock ensures that no value is present for InStock, not even an explicit nil
### GetStockInitial

`func (o *Product) GetStockInitial() float32`

GetStockInitial returns the StockInitial field if non-nil, zero value otherwise.

### GetStockInitialOk

`func (o *Product) GetStockInitialOk() (*float32, bool)`

GetStockInitialOk returns a tuple with the StockInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockInitial

`func (o *Product) SetStockInitial(v float32)`

SetStockInitial sets StockInitial field to given value.

### HasStockInitial

`func (o *Product) HasStockInitial() bool`

HasStockInitial returns a boolean if a field has been set.

### SetStockInitialNil

`func (o *Product) SetStockInitialNil(b bool)`

 SetStockInitialNil sets the value for StockInitial to be an explicit nil

### UnsetStockInitial
`func (o *Product) UnsetStockInitial()`

UnsetStockInitial ensures that no value is present for StockInitial, not even an explicit nil
### GetStockCurrent

`func (o *Product) GetStockCurrent() float32`

GetStockCurrent returns the StockCurrent field if non-nil, zero value otherwise.

### GetStockCurrentOk

`func (o *Product) GetStockCurrentOk() (*float32, bool)`

GetStockCurrentOk returns a tuple with the StockCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockCurrent

`func (o *Product) SetStockCurrent(v float32)`

SetStockCurrent sets StockCurrent field to given value.

### HasStockCurrent

`func (o *Product) HasStockCurrent() bool`

HasStockCurrent returns a boolean if a field has been set.

### SetStockCurrentNil

`func (o *Product) SetStockCurrentNil(b bool)`

 SetStockCurrentNil sets the value for StockCurrent to be an explicit nil

### UnsetStockCurrent
`func (o *Product) UnsetStockCurrent()`

UnsetStockCurrent ensures that no value is present for StockCurrent, not even an explicit nil
### GetAverageCost

`func (o *Product) GetAverageCost() float32`

GetAverageCost returns the AverageCost field if non-nil, zero value otherwise.

### GetAverageCostOk

`func (o *Product) GetAverageCostOk() (*float32, bool)`

GetAverageCostOk returns a tuple with the AverageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCost

`func (o *Product) SetAverageCost(v float32)`

SetAverageCost sets AverageCost field to given value.

### HasAverageCost

`func (o *Product) HasAverageCost() bool`

HasAverageCost returns a boolean if a field has been set.

### SetAverageCostNil

`func (o *Product) SetAverageCostNil(b bool)`

 SetAverageCostNil sets the value for AverageCost to be an explicit nil

### UnsetAverageCost
`func (o *Product) UnsetAverageCost()`

UnsetAverageCost ensures that no value is present for AverageCost, not even an explicit nil
### GetAveragePrice

`func (o *Product) GetAveragePrice() float32`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *Product) GetAveragePriceOk() (*float32, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *Product) SetAveragePrice(v float32)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *Product) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### SetAveragePriceNil

`func (o *Product) SetAveragePriceNil(b bool)`

 SetAveragePriceNil sets the value for AveragePrice to be an explicit nil

### UnsetAveragePrice
`func (o *Product) UnsetAveragePrice()`

UnsetAveragePrice ensures that no value is present for AveragePrice, not even an explicit nil
### GetCreatedAt

`func (o *Product) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Product) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Product) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Product) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Product) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Product) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Product) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Product) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Product) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Product) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Product) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Product) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


