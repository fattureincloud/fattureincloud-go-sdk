# ArchiveDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Archive document unique identifier. | [optional] 
**Date** | Pointer to **NullableString** | Archive document date. | [optional] 
**Description** | Pointer to **NullableString** | Archive Document description. | [optional] 
**AttachmentUrl** | Pointer to **NullableString** | [Temporary] [Read Only] Absolute url of the attached file. Authomatically set if a valid attachment token is passed via POST /archive or PUT /archive/{documentId}. | [optional] [readonly] 
**Category** | Pointer to **NullableString** | Archive document category. | [optional] 
**AttachmentToken** | Pointer to **NullableString** | [Write Only]  [Required] Attachment token returned by POST /archive/attachment. Used to attach the file already uploaded. | [optional] 

## Methods

### NewArchiveDocument

`func NewArchiveDocument() *ArchiveDocument`

NewArchiveDocument instantiates a new ArchiveDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveDocumentWithDefaults

`func NewArchiveDocumentWithDefaults() *ArchiveDocument`

NewArchiveDocumentWithDefaults instantiates a new ArchiveDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArchiveDocument) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveDocument) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveDocument) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ArchiveDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ArchiveDocument) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ArchiveDocument) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDate

`func (o *ArchiveDocument) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ArchiveDocument) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ArchiveDocument) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ArchiveDocument) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *ArchiveDocument) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ArchiveDocument) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *ArchiveDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArchiveDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArchiveDocument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArchiveDocument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ArchiveDocument) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ArchiveDocument) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAttachmentUrl

`func (o *ArchiveDocument) GetAttachmentUrl() string`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *ArchiveDocument) GetAttachmentUrlOk() (*string, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *ArchiveDocument) SetAttachmentUrl(v string)`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *ArchiveDocument) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### SetAttachmentUrlNil

`func (o *ArchiveDocument) SetAttachmentUrlNil(b bool)`

 SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil

### UnsetAttachmentUrl
`func (o *ArchiveDocument) UnsetAttachmentUrl()`

UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
### GetCategory

`func (o *ArchiveDocument) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ArchiveDocument) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ArchiveDocument) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ArchiveDocument) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ArchiveDocument) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ArchiveDocument) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetAttachmentToken

`func (o *ArchiveDocument) GetAttachmentToken() string`

GetAttachmentToken returns the AttachmentToken field if non-nil, zero value otherwise.

### GetAttachmentTokenOk

`func (o *ArchiveDocument) GetAttachmentTokenOk() (*string, bool)`

GetAttachmentTokenOk returns a tuple with the AttachmentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentToken

`func (o *ArchiveDocument) SetAttachmentToken(v string)`

SetAttachmentToken sets AttachmentToken field to given value.

### HasAttachmentToken

`func (o *ArchiveDocument) HasAttachmentToken() bool`

HasAttachmentToken returns a boolean if a field has been set.

### SetAttachmentTokenNil

`func (o *ArchiveDocument) SetAttachmentTokenNil(b bool)`

 SetAttachmentTokenNil sets the value for AttachmentToken to be an explicit nil

### UnsetAttachmentToken
`func (o *ArchiveDocument) UnsetAttachmentToken()`

UnsetAttachmentToken ensures that no value is present for AttachmentToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


