# ListProductsResponsePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Product**](Product.md) |  | [optional] 

## Methods

### NewListProductsResponsePage

`func NewListProductsResponsePage() *ListProductsResponsePage`

NewListProductsResponsePage instantiates a new ListProductsResponsePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProductsResponsePageWithDefaults

`func NewListProductsResponsePageWithDefaults() *ListProductsResponsePage`

NewListProductsResponsePageWithDefaults instantiates a new ListProductsResponsePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListProductsResponsePage) GetData() []Product`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListProductsResponsePage) GetDataOk() (*[]Product, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListProductsResponsePage) SetData(v []Product)`

SetData sets Data field to given value.

### HasData

`func (o *ListProductsResponsePage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ListProductsResponsePage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ListProductsResponsePage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


