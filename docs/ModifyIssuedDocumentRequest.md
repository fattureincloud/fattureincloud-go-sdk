# ModifyIssuedDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IssuedDocument**](IssuedDocument.md) |  | [optional] 
**Options** | Pointer to [**IssuedDocumentOptions**](IssuedDocumentOptions.md) |  | [optional] 

## Methods

### NewModifyIssuedDocumentRequest

`func NewModifyIssuedDocumentRequest() *ModifyIssuedDocumentRequest`

NewModifyIssuedDocumentRequest instantiates a new ModifyIssuedDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyIssuedDocumentRequestWithDefaults

`func NewModifyIssuedDocumentRequestWithDefaults() *ModifyIssuedDocumentRequest`

NewModifyIssuedDocumentRequestWithDefaults instantiates a new ModifyIssuedDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModifyIssuedDocumentRequest) GetData() IssuedDocument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModifyIssuedDocumentRequest) GetDataOk() (*IssuedDocument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModifyIssuedDocumentRequest) SetData(v IssuedDocument)`

SetData sets Data field to given value.

### HasData

`func (o *ModifyIssuedDocumentRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOptions

`func (o *ModifyIssuedDocumentRequest) GetOptions() IssuedDocumentOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ModifyIssuedDocumentRequest) GetOptionsOk() (*IssuedDocumentOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ModifyIssuedDocumentRequest) SetOptions(v IssuedDocumentOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ModifyIssuedDocumentRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


