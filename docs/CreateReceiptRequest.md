# CreateReceiptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Receipt**](Receipt.md) |  | [optional] 
**AutocompleteNumber** | Pointer to **NullableBool** | If true, the number is autocompleted progressively. | [optional] 

## Methods

### NewCreateReceiptRequest

`func NewCreateReceiptRequest() *CreateReceiptRequest`

NewCreateReceiptRequest instantiates a new CreateReceiptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReceiptRequestWithDefaults

`func NewCreateReceiptRequestWithDefaults() *CreateReceiptRequest`

NewCreateReceiptRequestWithDefaults instantiates a new CreateReceiptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateReceiptRequest) GetData() Receipt`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateReceiptRequest) GetDataOk() (*Receipt, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateReceiptRequest) SetData(v Receipt)`

SetData sets Data field to given value.

### HasData

`func (o *CreateReceiptRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetAutocompleteNumber

`func (o *CreateReceiptRequest) GetAutocompleteNumber() bool`

GetAutocompleteNumber returns the AutocompleteNumber field if non-nil, zero value otherwise.

### GetAutocompleteNumberOk

`func (o *CreateReceiptRequest) GetAutocompleteNumberOk() (*bool, bool)`

GetAutocompleteNumberOk returns a tuple with the AutocompleteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocompleteNumber

`func (o *CreateReceiptRequest) SetAutocompleteNumber(v bool)`

SetAutocompleteNumber sets AutocompleteNumber field to given value.

### HasAutocompleteNumber

`func (o *CreateReceiptRequest) HasAutocompleteNumber() bool`

HasAutocompleteNumber returns a boolean if a field has been set.

### SetAutocompleteNumberNil

`func (o *CreateReceiptRequest) SetAutocompleteNumberNil(b bool)`

 SetAutocompleteNumberNil sets the value for AutocompleteNumber to be an explicit nil

### UnsetAutocompleteNumber
`func (o *CreateReceiptRequest) UnsetAutocompleteNumber()`

UnsetAutocompleteNumber ensures that no value is present for AutocompleteNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


