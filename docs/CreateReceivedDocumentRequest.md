# CreateReceivedDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingId** | Pointer to **NullableInt32** | Pending received document id of the document from which the new document is created. | [optional] 
**Data** | Pointer to [**ReceivedDocument**](ReceivedDocument.md) |  | [optional] 

## Methods

### NewCreateReceivedDocumentRequest

`func NewCreateReceivedDocumentRequest() *CreateReceivedDocumentRequest`

NewCreateReceivedDocumentRequest instantiates a new CreateReceivedDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReceivedDocumentRequestWithDefaults

`func NewCreateReceivedDocumentRequestWithDefaults() *CreateReceivedDocumentRequest`

NewCreateReceivedDocumentRequestWithDefaults instantiates a new CreateReceivedDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingId

`func (o *CreateReceivedDocumentRequest) GetPendingId() int32`

GetPendingId returns the PendingId field if non-nil, zero value otherwise.

### GetPendingIdOk

`func (o *CreateReceivedDocumentRequest) GetPendingIdOk() (*int32, bool)`

GetPendingIdOk returns a tuple with the PendingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingId

`func (o *CreateReceivedDocumentRequest) SetPendingId(v int32)`

SetPendingId sets PendingId field to given value.

### HasPendingId

`func (o *CreateReceivedDocumentRequest) HasPendingId() bool`

HasPendingId returns a boolean if a field has been set.

### SetPendingIdNil

`func (o *CreateReceivedDocumentRequest) SetPendingIdNil(b bool)`

 SetPendingIdNil sets the value for PendingId to be an explicit nil

### UnsetPendingId
`func (o *CreateReceivedDocumentRequest) UnsetPendingId()`

UnsetPendingId ensures that no value is present for PendingId, not even an explicit nil
### GetData

`func (o *CreateReceivedDocumentRequest) GetData() ReceivedDocument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateReceivedDocumentRequest) GetDataOk() (*ReceivedDocument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateReceivedDocumentRequest) SetData(v ReceivedDocument)`

SetData sets Data field to given value.

### HasData

`func (o *CreateReceivedDocumentRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


