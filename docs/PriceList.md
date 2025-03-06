# PriceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Price list id | [optional] 
**Name** | Pointer to **NullableString** | Price list name | [optional] 
**PricesType** | Pointer to [**PriceListPricesType**](PriceListPricesType.md) |  | [optional] 
**IsDefault** | Pointer to **NullableBool** | This entity is default | [optional] 
**ValidFrom** | Pointer to **NullableString** | Price list validity start date | [optional] 
**ValidTo** | Pointer to **NullableString** | Price list validity end date | [optional] 
**Type** | Pointer to [**PriceListType**](PriceListType.md) |  | [optional] 

## Methods

### NewPriceList

`func NewPriceList() *PriceList`

NewPriceList instantiates a new PriceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceListWithDefaults

`func NewPriceListWithDefaults() *PriceList`

NewPriceListWithDefaults instantiates a new PriceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriceList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PriceList) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PriceList) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PriceList) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *PriceList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriceList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriceList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PriceList) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PriceList) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PriceList) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPricesType

`func (o *PriceList) GetPricesType() PriceListPricesType`

GetPricesType returns the PricesType field if non-nil, zero value otherwise.

### GetPricesTypeOk

`func (o *PriceList) GetPricesTypeOk() (*PriceListPricesType, bool)`

GetPricesTypeOk returns a tuple with the PricesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesType

`func (o *PriceList) SetPricesType(v PriceListPricesType)`

SetPricesType sets PricesType field to given value.

### HasPricesType

`func (o *PriceList) HasPricesType() bool`

HasPricesType returns a boolean if a field has been set.

### GetIsDefault

`func (o *PriceList) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PriceList) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PriceList) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PriceList) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *PriceList) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *PriceList) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetValidFrom

`func (o *PriceList) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *PriceList) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *PriceList) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *PriceList) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### SetValidFromNil

`func (o *PriceList) SetValidFromNil(b bool)`

 SetValidFromNil sets the value for ValidFrom to be an explicit nil

### UnsetValidFrom
`func (o *PriceList) UnsetValidFrom()`

UnsetValidFrom ensures that no value is present for ValidFrom, not even an explicit nil
### GetValidTo

`func (o *PriceList) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *PriceList) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *PriceList) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *PriceList) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### SetValidToNil

`func (o *PriceList) SetValidToNil(b bool)`

 SetValidToNil sets the value for ValidTo to be an explicit nil

### UnsetValidTo
`func (o *PriceList) UnsetValidTo()`

UnsetValidTo ensures that no value is present for ValidTo, not even an explicit nil
### GetType

`func (o *PriceList) GetType() PriceListType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceList) GetTypeOk() (*PriceListType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceList) SetType(v PriceListType)`

SetType sets Type field to given value.

### HasType

`func (o *PriceList) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


