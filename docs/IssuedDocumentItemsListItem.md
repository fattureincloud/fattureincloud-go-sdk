# IssuedDocumentItemsListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Issued document item id | [optional] 
**ProductId** | Pointer to **NullableInt32** | Issued document item product id | [optional] 
**Code** | Pointer to **NullableString** | Issued document item product code | [optional] 
**Name** | Pointer to **NullableString** | Issued document item product name | [optional] 
**Category** | Pointer to **NullableString** | Issued document item product category | [optional] 
**Description** | Pointer to **NullableString** | Issued document product description | [optional] 
**Qty** | Pointer to **NullableFloat32** | Issued document item quantity | [optional] 
**Measure** | Pointer to **NullableString** | Issued document item measure | [optional] 
**NetPrice** | Pointer to **NullableFloat32** | Issued document item net price | [optional] 
**GrossPrice** | Pointer to **NullableFloat32** | Issued document item gross price | [optional] 
**Vat** | Pointer to [**NullableVatType**](VatType.md) |  | [optional] 
**NotTaxable** | Pointer to **NullableBool** | Issued document item is not taxable | [optional] 
**ApplyWithholdingTaxes** | Pointer to **NullableBool** | Issued document item apply withholding taxes, rivalsa and cassa | [optional] 
**Discount** | Pointer to **NullableFloat32** | Issued document item discount percentual value | [optional] 
**DiscountHighlight** | Pointer to **NullableBool** | Issued document item highlight discount | [optional] 
**InDn** | Pointer to **NullableBool** | Issued document item add in delivery note | [optional] 
**Stock** | Pointer to **NullableBool** | Issued document item move stock | [optional] 
**EiRaw** | Pointer to **map[string]interface{}** | Issued document advanced raw attributes for e-invoices | [optional] 

## Methods

### NewIssuedDocumentItemsListItem

`func NewIssuedDocumentItemsListItem() *IssuedDocumentItemsListItem`

NewIssuedDocumentItemsListItem instantiates a new IssuedDocumentItemsListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentItemsListItemWithDefaults

`func NewIssuedDocumentItemsListItemWithDefaults() *IssuedDocumentItemsListItem`

NewIssuedDocumentItemsListItemWithDefaults instantiates a new IssuedDocumentItemsListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssuedDocumentItemsListItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssuedDocumentItemsListItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssuedDocumentItemsListItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IssuedDocumentItemsListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IssuedDocumentItemsListItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IssuedDocumentItemsListItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProductId

`func (o *IssuedDocumentItemsListItem) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *IssuedDocumentItemsListItem) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *IssuedDocumentItemsListItem) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *IssuedDocumentItemsListItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *IssuedDocumentItemsListItem) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *IssuedDocumentItemsListItem) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetCode

`func (o *IssuedDocumentItemsListItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IssuedDocumentItemsListItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IssuedDocumentItemsListItem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IssuedDocumentItemsListItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *IssuedDocumentItemsListItem) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *IssuedDocumentItemsListItem) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *IssuedDocumentItemsListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssuedDocumentItemsListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssuedDocumentItemsListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssuedDocumentItemsListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IssuedDocumentItemsListItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IssuedDocumentItemsListItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCategory

`func (o *IssuedDocumentItemsListItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IssuedDocumentItemsListItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IssuedDocumentItemsListItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IssuedDocumentItemsListItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *IssuedDocumentItemsListItem) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *IssuedDocumentItemsListItem) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *IssuedDocumentItemsListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssuedDocumentItemsListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssuedDocumentItemsListItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssuedDocumentItemsListItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IssuedDocumentItemsListItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IssuedDocumentItemsListItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQty

`func (o *IssuedDocumentItemsListItem) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *IssuedDocumentItemsListItem) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *IssuedDocumentItemsListItem) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *IssuedDocumentItemsListItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### SetQtyNil

`func (o *IssuedDocumentItemsListItem) SetQtyNil(b bool)`

 SetQtyNil sets the value for Qty to be an explicit nil

### UnsetQty
`func (o *IssuedDocumentItemsListItem) UnsetQty()`

UnsetQty ensures that no value is present for Qty, not even an explicit nil
### GetMeasure

`func (o *IssuedDocumentItemsListItem) GetMeasure() string`

GetMeasure returns the Measure field if non-nil, zero value otherwise.

### GetMeasureOk

`func (o *IssuedDocumentItemsListItem) GetMeasureOk() (*string, bool)`

GetMeasureOk returns a tuple with the Measure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasure

`func (o *IssuedDocumentItemsListItem) SetMeasure(v string)`

SetMeasure sets Measure field to given value.

### HasMeasure

`func (o *IssuedDocumentItemsListItem) HasMeasure() bool`

HasMeasure returns a boolean if a field has been set.

### SetMeasureNil

`func (o *IssuedDocumentItemsListItem) SetMeasureNil(b bool)`

 SetMeasureNil sets the value for Measure to be an explicit nil

### UnsetMeasure
`func (o *IssuedDocumentItemsListItem) UnsetMeasure()`

UnsetMeasure ensures that no value is present for Measure, not even an explicit nil
### GetNetPrice

`func (o *IssuedDocumentItemsListItem) GetNetPrice() float32`

GetNetPrice returns the NetPrice field if non-nil, zero value otherwise.

### GetNetPriceOk

`func (o *IssuedDocumentItemsListItem) GetNetPriceOk() (*float32, bool)`

GetNetPriceOk returns a tuple with the NetPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPrice

`func (o *IssuedDocumentItemsListItem) SetNetPrice(v float32)`

SetNetPrice sets NetPrice field to given value.

### HasNetPrice

`func (o *IssuedDocumentItemsListItem) HasNetPrice() bool`

HasNetPrice returns a boolean if a field has been set.

### SetNetPriceNil

`func (o *IssuedDocumentItemsListItem) SetNetPriceNil(b bool)`

 SetNetPriceNil sets the value for NetPrice to be an explicit nil

### UnsetNetPrice
`func (o *IssuedDocumentItemsListItem) UnsetNetPrice()`

UnsetNetPrice ensures that no value is present for NetPrice, not even an explicit nil
### GetGrossPrice

`func (o *IssuedDocumentItemsListItem) GetGrossPrice() float32`

GetGrossPrice returns the GrossPrice field if non-nil, zero value otherwise.

### GetGrossPriceOk

`func (o *IssuedDocumentItemsListItem) GetGrossPriceOk() (*float32, bool)`

GetGrossPriceOk returns a tuple with the GrossPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossPrice

`func (o *IssuedDocumentItemsListItem) SetGrossPrice(v float32)`

SetGrossPrice sets GrossPrice field to given value.

### HasGrossPrice

`func (o *IssuedDocumentItemsListItem) HasGrossPrice() bool`

HasGrossPrice returns a boolean if a field has been set.

### SetGrossPriceNil

`func (o *IssuedDocumentItemsListItem) SetGrossPriceNil(b bool)`

 SetGrossPriceNil sets the value for GrossPrice to be an explicit nil

### UnsetGrossPrice
`func (o *IssuedDocumentItemsListItem) UnsetGrossPrice()`

UnsetGrossPrice ensures that no value is present for GrossPrice, not even an explicit nil
### GetVat

`func (o *IssuedDocumentItemsListItem) GetVat() VatType`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *IssuedDocumentItemsListItem) GetVatOk() (*VatType, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *IssuedDocumentItemsListItem) SetVat(v VatType)`

SetVat sets Vat field to given value.

### HasVat

`func (o *IssuedDocumentItemsListItem) HasVat() bool`

HasVat returns a boolean if a field has been set.

### SetVatNil

`func (o *IssuedDocumentItemsListItem) SetVatNil(b bool)`

 SetVatNil sets the value for Vat to be an explicit nil

### UnsetVat
`func (o *IssuedDocumentItemsListItem) UnsetVat()`

UnsetVat ensures that no value is present for Vat, not even an explicit nil
### GetNotTaxable

`func (o *IssuedDocumentItemsListItem) GetNotTaxable() bool`

GetNotTaxable returns the NotTaxable field if non-nil, zero value otherwise.

### GetNotTaxableOk

`func (o *IssuedDocumentItemsListItem) GetNotTaxableOk() (*bool, bool)`

GetNotTaxableOk returns a tuple with the NotTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotTaxable

`func (o *IssuedDocumentItemsListItem) SetNotTaxable(v bool)`

SetNotTaxable sets NotTaxable field to given value.

### HasNotTaxable

`func (o *IssuedDocumentItemsListItem) HasNotTaxable() bool`

HasNotTaxable returns a boolean if a field has been set.

### SetNotTaxableNil

`func (o *IssuedDocumentItemsListItem) SetNotTaxableNil(b bool)`

 SetNotTaxableNil sets the value for NotTaxable to be an explicit nil

### UnsetNotTaxable
`func (o *IssuedDocumentItemsListItem) UnsetNotTaxable()`

UnsetNotTaxable ensures that no value is present for NotTaxable, not even an explicit nil
### GetApplyWithholdingTaxes

`func (o *IssuedDocumentItemsListItem) GetApplyWithholdingTaxes() bool`

GetApplyWithholdingTaxes returns the ApplyWithholdingTaxes field if non-nil, zero value otherwise.

### GetApplyWithholdingTaxesOk

`func (o *IssuedDocumentItemsListItem) GetApplyWithholdingTaxesOk() (*bool, bool)`

GetApplyWithholdingTaxesOk returns a tuple with the ApplyWithholdingTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyWithholdingTaxes

`func (o *IssuedDocumentItemsListItem) SetApplyWithholdingTaxes(v bool)`

SetApplyWithholdingTaxes sets ApplyWithholdingTaxes field to given value.

### HasApplyWithholdingTaxes

`func (o *IssuedDocumentItemsListItem) HasApplyWithholdingTaxes() bool`

HasApplyWithholdingTaxes returns a boolean if a field has been set.

### SetApplyWithholdingTaxesNil

`func (o *IssuedDocumentItemsListItem) SetApplyWithholdingTaxesNil(b bool)`

 SetApplyWithholdingTaxesNil sets the value for ApplyWithholdingTaxes to be an explicit nil

### UnsetApplyWithholdingTaxes
`func (o *IssuedDocumentItemsListItem) UnsetApplyWithholdingTaxes()`

UnsetApplyWithholdingTaxes ensures that no value is present for ApplyWithholdingTaxes, not even an explicit nil
### GetDiscount

`func (o *IssuedDocumentItemsListItem) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *IssuedDocumentItemsListItem) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *IssuedDocumentItemsListItem) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *IssuedDocumentItemsListItem) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *IssuedDocumentItemsListItem) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *IssuedDocumentItemsListItem) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetDiscountHighlight

`func (o *IssuedDocumentItemsListItem) GetDiscountHighlight() bool`

GetDiscountHighlight returns the DiscountHighlight field if non-nil, zero value otherwise.

### GetDiscountHighlightOk

`func (o *IssuedDocumentItemsListItem) GetDiscountHighlightOk() (*bool, bool)`

GetDiscountHighlightOk returns a tuple with the DiscountHighlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountHighlight

`func (o *IssuedDocumentItemsListItem) SetDiscountHighlight(v bool)`

SetDiscountHighlight sets DiscountHighlight field to given value.

### HasDiscountHighlight

`func (o *IssuedDocumentItemsListItem) HasDiscountHighlight() bool`

HasDiscountHighlight returns a boolean if a field has been set.

### SetDiscountHighlightNil

`func (o *IssuedDocumentItemsListItem) SetDiscountHighlightNil(b bool)`

 SetDiscountHighlightNil sets the value for DiscountHighlight to be an explicit nil

### UnsetDiscountHighlight
`func (o *IssuedDocumentItemsListItem) UnsetDiscountHighlight()`

UnsetDiscountHighlight ensures that no value is present for DiscountHighlight, not even an explicit nil
### GetInDn

`func (o *IssuedDocumentItemsListItem) GetInDn() bool`

GetInDn returns the InDn field if non-nil, zero value otherwise.

### GetInDnOk

`func (o *IssuedDocumentItemsListItem) GetInDnOk() (*bool, bool)`

GetInDnOk returns a tuple with the InDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInDn

`func (o *IssuedDocumentItemsListItem) SetInDn(v bool)`

SetInDn sets InDn field to given value.

### HasInDn

`func (o *IssuedDocumentItemsListItem) HasInDn() bool`

HasInDn returns a boolean if a field has been set.

### SetInDnNil

`func (o *IssuedDocumentItemsListItem) SetInDnNil(b bool)`

 SetInDnNil sets the value for InDn to be an explicit nil

### UnsetInDn
`func (o *IssuedDocumentItemsListItem) UnsetInDn()`

UnsetInDn ensures that no value is present for InDn, not even an explicit nil
### GetStock

`func (o *IssuedDocumentItemsListItem) GetStock() bool`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *IssuedDocumentItemsListItem) GetStockOk() (*bool, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *IssuedDocumentItemsListItem) SetStock(v bool)`

SetStock sets Stock field to given value.

### HasStock

`func (o *IssuedDocumentItemsListItem) HasStock() bool`

HasStock returns a boolean if a field has been set.

### SetStockNil

`func (o *IssuedDocumentItemsListItem) SetStockNil(b bool)`

 SetStockNil sets the value for Stock to be an explicit nil

### UnsetStock
`func (o *IssuedDocumentItemsListItem) UnsetStock()`

UnsetStock ensures that no value is present for Stock, not even an explicit nil
### GetEiRaw

`func (o *IssuedDocumentItemsListItem) GetEiRaw() map[string]interface{}`

GetEiRaw returns the EiRaw field if non-nil, zero value otherwise.

### GetEiRawOk

`func (o *IssuedDocumentItemsListItem) GetEiRawOk() (*map[string]interface{}, bool)`

GetEiRawOk returns a tuple with the EiRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiRaw

`func (o *IssuedDocumentItemsListItem) SetEiRaw(v map[string]interface{})`

SetEiRaw sets EiRaw field to given value.

### HasEiRaw

`func (o *IssuedDocumentItemsListItem) HasEiRaw() bool`

HasEiRaw returns a boolean if a field has been set.

### SetEiRawNil

`func (o *IssuedDocumentItemsListItem) SetEiRawNil(b bool)`

 SetEiRawNil sets the value for EiRaw to be an explicit nil

### UnsetEiRaw
`func (o *IssuedDocumentItemsListItem) UnsetEiRaw()`

UnsetEiRaw ensures that no value is present for EiRaw, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


