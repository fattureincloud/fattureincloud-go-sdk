# SendEInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableSendEInvoiceRequestData**](SendEInvoiceRequestData.md) |  | [optional] 
**Options** | Pointer to [**NullableSendEInvoiceRequestOptions**](SendEInvoiceRequestOptions.md) |  | [optional] 

## Methods

### NewSendEInvoiceRequest

`func NewSendEInvoiceRequest() *SendEInvoiceRequest`

NewSendEInvoiceRequest instantiates a new SendEInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEInvoiceRequestWithDefaults

`func NewSendEInvoiceRequestWithDefaults() *SendEInvoiceRequest`

NewSendEInvoiceRequestWithDefaults instantiates a new SendEInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SendEInvoiceRequest) GetData() SendEInvoiceRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SendEInvoiceRequest) GetDataOk() (*SendEInvoiceRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SendEInvoiceRequest) SetData(v SendEInvoiceRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *SendEInvoiceRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SendEInvoiceRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SendEInvoiceRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetOptions

`func (o *SendEInvoiceRequest) GetOptions() SendEInvoiceRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SendEInvoiceRequest) GetOptionsOk() (*SendEInvoiceRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SendEInvoiceRequest) SetOptions(v SendEInvoiceRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SendEInvoiceRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *SendEInvoiceRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *SendEInvoiceRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


