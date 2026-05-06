# DocumentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Template id | [optional] 
**Privacy** | Pointer to **string** | Template privacy | [optional] 
**Type** | Pointer to [**TemplateType**](TemplateType.md) |  | [optional] 
**Name** | Pointer to **string** | Template name | [optional] 
**CanDisableWatermark** | Pointer to **bool** | Can disable watermark | [optional] 
**Author** | Pointer to **string** | Template author | [optional] 
**Content** | Pointer to **string** | Template definition content | [optional] 
**SupportsCustomTaxable** | Pointer to **bool** | Supports custom taxable | [optional] 

## Methods

### NewDocumentTemplate

`func NewDocumentTemplate() *DocumentTemplate`

NewDocumentTemplate instantiates a new DocumentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentTemplateWithDefaults

`func NewDocumentTemplateWithDefaults() *DocumentTemplate`

NewDocumentTemplateWithDefaults instantiates a new DocumentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrivacy

`func (o *DocumentTemplate) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *DocumentTemplate) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *DocumentTemplate) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *DocumentTemplate) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetType

`func (o *DocumentTemplate) GetType() TemplateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentTemplate) GetTypeOk() (*TemplateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentTemplate) SetType(v TemplateType)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DocumentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCanDisableWatermark

`func (o *DocumentTemplate) GetCanDisableWatermark() bool`

GetCanDisableWatermark returns the CanDisableWatermark field if non-nil, zero value otherwise.

### GetCanDisableWatermarkOk

`func (o *DocumentTemplate) GetCanDisableWatermarkOk() (*bool, bool)`

GetCanDisableWatermarkOk returns a tuple with the CanDisableWatermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDisableWatermark

`func (o *DocumentTemplate) SetCanDisableWatermark(v bool)`

SetCanDisableWatermark sets CanDisableWatermark field to given value.

### HasCanDisableWatermark

`func (o *DocumentTemplate) HasCanDisableWatermark() bool`

HasCanDisableWatermark returns a boolean if a field has been set.

### GetAuthor

`func (o *DocumentTemplate) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *DocumentTemplate) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *DocumentTemplate) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *DocumentTemplate) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetContent

`func (o *DocumentTemplate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DocumentTemplate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DocumentTemplate) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DocumentTemplate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSupportsCustomTaxable

`func (o *DocumentTemplate) GetSupportsCustomTaxable() bool`

GetSupportsCustomTaxable returns the SupportsCustomTaxable field if non-nil, zero value otherwise.

### GetSupportsCustomTaxableOk

`func (o *DocumentTemplate) GetSupportsCustomTaxableOk() (*bool, bool)`

GetSupportsCustomTaxableOk returns a tuple with the SupportsCustomTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCustomTaxable

`func (o *DocumentTemplate) SetSupportsCustomTaxable(v bool)`

SetSupportsCustomTaxable sets SupportsCustomTaxable field to given value.

### HasSupportsCustomTaxable

`func (o *DocumentTemplate) HasSupportsCustomTaxable() bool`

HasSupportsCustomTaxable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


