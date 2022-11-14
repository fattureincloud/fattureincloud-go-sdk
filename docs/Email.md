# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Email unique identifier. | [optional] 
**Status** | Pointer to [**EmailStatus**](EmailStatus.md) |  | [optional] 
**SentDate** | Pointer to **time.Time** | Email sent date. | [optional] 
**ErrorsCount** | Pointer to **int32** | Errors count. | [optional] 
**ErrorLog** | Pointer to **string** | Error log. | [optional] 
**FromEmail** | Pointer to **string** | Sender email. | [optional] 
**FromName** | Pointer to **string** | Sender name. | [optional] 
**ToEmail** | Pointer to **string** | Recipient email. | [optional] 
**ToName** | Pointer to **string** | Receipient email. | [optional] 
**Subject** | Pointer to **string** | Email subject. | [optional] 
**Content** | Pointer to **string** | Email content. | [optional] 
**CopyTo** | Pointer to **string** |  | [optional] 
**RecipientStatus** | Pointer to [**EmailRecipientStatus**](EmailRecipientStatus.md) |  | [optional] 
**RecipientDate** | Pointer to **time.Time** |  | [optional] 
**Kind** | Pointer to **string** | Email kind. | [optional] 
**Attachments** | Pointer to [**[]EmailAttachment**](EmailAttachment.md) | Email attachments. | [optional] 

## Methods

### NewEmail

`func NewEmail() *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Email) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Email) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Email) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Email) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Email) GetStatus() EmailStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Email) GetStatusOk() (*EmailStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Email) SetStatus(v EmailStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Email) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSentDate

`func (o *Email) GetSentDate() time.Time`

GetSentDate returns the SentDate field if non-nil, zero value otherwise.

### GetSentDateOk

`func (o *Email) GetSentDateOk() (*time.Time, bool)`

GetSentDateOk returns a tuple with the SentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentDate

`func (o *Email) SetSentDate(v time.Time)`

SetSentDate sets SentDate field to given value.

### HasSentDate

`func (o *Email) HasSentDate() bool`

HasSentDate returns a boolean if a field has been set.

### GetErrorsCount

`func (o *Email) GetErrorsCount() int32`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *Email) GetErrorsCountOk() (*int32, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *Email) SetErrorsCount(v int32)`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *Email) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### GetErrorLog

`func (o *Email) GetErrorLog() string`

GetErrorLog returns the ErrorLog field if non-nil, zero value otherwise.

### GetErrorLogOk

`func (o *Email) GetErrorLogOk() (*string, bool)`

GetErrorLogOk returns a tuple with the ErrorLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLog

`func (o *Email) SetErrorLog(v string)`

SetErrorLog sets ErrorLog field to given value.

### HasErrorLog

`func (o *Email) HasErrorLog() bool`

HasErrorLog returns a boolean if a field has been set.

### GetFromEmail

`func (o *Email) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *Email) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *Email) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *Email) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetFromName

`func (o *Email) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *Email) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *Email) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *Email) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetToEmail

`func (o *Email) GetToEmail() string`

GetToEmail returns the ToEmail field if non-nil, zero value otherwise.

### GetToEmailOk

`func (o *Email) GetToEmailOk() (*string, bool)`

GetToEmailOk returns a tuple with the ToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEmail

`func (o *Email) SetToEmail(v string)`

SetToEmail sets ToEmail field to given value.

### HasToEmail

`func (o *Email) HasToEmail() bool`

HasToEmail returns a boolean if a field has been set.

### GetToName

`func (o *Email) GetToName() string`

GetToName returns the ToName field if non-nil, zero value otherwise.

### GetToNameOk

`func (o *Email) GetToNameOk() (*string, bool)`

GetToNameOk returns a tuple with the ToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToName

`func (o *Email) SetToName(v string)`

SetToName sets ToName field to given value.

### HasToName

`func (o *Email) HasToName() bool`

HasToName returns a boolean if a field has been set.

### GetSubject

`func (o *Email) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Email) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Email) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Email) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContent

`func (o *Email) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Email) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Email) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Email) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCopyTo

`func (o *Email) GetCopyTo() string`

GetCopyTo returns the CopyTo field if non-nil, zero value otherwise.

### GetCopyToOk

`func (o *Email) GetCopyToOk() (*string, bool)`

GetCopyToOk returns a tuple with the CopyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTo

`func (o *Email) SetCopyTo(v string)`

SetCopyTo sets CopyTo field to given value.

### HasCopyTo

`func (o *Email) HasCopyTo() bool`

HasCopyTo returns a boolean if a field has been set.

### GetRecipientStatus

`func (o *Email) GetRecipientStatus() EmailRecipientStatus`

GetRecipientStatus returns the RecipientStatus field if non-nil, zero value otherwise.

### GetRecipientStatusOk

`func (o *Email) GetRecipientStatusOk() (*EmailRecipientStatus, bool)`

GetRecipientStatusOk returns a tuple with the RecipientStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientStatus

`func (o *Email) SetRecipientStatus(v EmailRecipientStatus)`

SetRecipientStatus sets RecipientStatus field to given value.

### HasRecipientStatus

`func (o *Email) HasRecipientStatus() bool`

HasRecipientStatus returns a boolean if a field has been set.

### GetRecipientDate

`func (o *Email) GetRecipientDate() time.Time`

GetRecipientDate returns the RecipientDate field if non-nil, zero value otherwise.

### GetRecipientDateOk

`func (o *Email) GetRecipientDateOk() (*time.Time, bool)`

GetRecipientDateOk returns a tuple with the RecipientDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientDate

`func (o *Email) SetRecipientDate(v time.Time)`

SetRecipientDate sets RecipientDate field to given value.

### HasRecipientDate

`func (o *Email) HasRecipientDate() bool`

HasRecipientDate returns a boolean if a field has been set.

### GetKind

`func (o *Email) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Email) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Email) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Email) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAttachments

`func (o *Email) GetAttachments() []EmailAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Email) GetAttachmentsOk() (*[]EmailAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Email) SetAttachments(v []EmailAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Email) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


