# GetVatTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableVatType**](VatType.md) |  | [optional] 

## Methods

### NewGetVatTypeResponse

`func NewGetVatTypeResponse() *GetVatTypeResponse`

NewGetVatTypeResponse instantiates a new GetVatTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVatTypeResponseWithDefaults

`func NewGetVatTypeResponseWithDefaults() *GetVatTypeResponse`

NewGetVatTypeResponseWithDefaults instantiates a new GetVatTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetVatTypeResponse) GetData() VatType`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVatTypeResponse) GetDataOk() (*VatType, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVatTypeResponse) SetData(v VatType)`

SetData sets Data field to given value.

### HasData

`func (o *GetVatTypeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GetVatTypeResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetVatTypeResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


