# EmailAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Email attachment filename | [optional] 
**Url** | Pointer to **string** | Email attachment url | [optional] 

## Methods

### NewEmailAttachment

`func NewEmailAttachment() *EmailAttachment`

NewEmailAttachment instantiates a new EmailAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAttachmentWithDefaults

`func NewEmailAttachmentWithDefaults() *EmailAttachment`

NewEmailAttachmentWithDefaults instantiates a new EmailAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *EmailAttachment) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *EmailAttachment) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *EmailAttachment) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *EmailAttachment) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetUrl

`func (o *EmailAttachment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmailAttachment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmailAttachment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EmailAttachment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


