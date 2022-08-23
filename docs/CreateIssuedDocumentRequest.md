# CreateIssuedDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IssuedDocument**](IssuedDocument.md) |  | [optional] 
**Options** | Pointer to [**IssuedDocumentOptions**](IssuedDocumentOptions.md) |  | [optional] 

## Methods

### NewCreateIssuedDocumentRequest

`func NewCreateIssuedDocumentRequest() *CreateIssuedDocumentRequest`

NewCreateIssuedDocumentRequest instantiates a new CreateIssuedDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssuedDocumentRequestWithDefaults

`func NewCreateIssuedDocumentRequestWithDefaults() *CreateIssuedDocumentRequest`

NewCreateIssuedDocumentRequestWithDefaults instantiates a new CreateIssuedDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateIssuedDocumentRequest) GetData() IssuedDocument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateIssuedDocumentRequest) GetDataOk() (*IssuedDocument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateIssuedDocumentRequest) SetData(v IssuedDocument)`

SetData sets Data field to given value.

### HasData

`func (o *CreateIssuedDocumentRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOptions

`func (o *CreateIssuedDocumentRequest) GetOptions() IssuedDocumentOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateIssuedDocumentRequest) GetOptionsOk() (*IssuedDocumentOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateIssuedDocumentRequest) SetOptions(v IssuedDocumentOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateIssuedDocumentRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


