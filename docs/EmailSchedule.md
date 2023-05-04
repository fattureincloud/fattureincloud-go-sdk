# EmailSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderId** | Pointer to **NullableInt32** | Sender id. Required if &#x60;sender_email&#x60; is not specified | [optional] 
**SenderEmail** | Pointer to **NullableString** | Sender email. Required if &#x60;sender_id&#x60; is not specified | [optional] 
**RecipientEmail** | Pointer to **NullableString** | One or more comma separated recipient emails | [optional] 
**Subject** | Pointer to **NullableString** | Email subject | [optional] 
**Body** | Pointer to **NullableString** | Email body [HTML Escaped] [max size 50KiB] | [optional] 
**Include** | Pointer to [**NullableEmailScheduleInclude**](EmailScheduleInclude.md) |  | [optional] 
**AttachPdf** | Pointer to **NullableBool** | If set to true, documents will be sent as PDF attachments too | [optional] 
**SendCopy** | Pointer to **NullableBool** | If set to true, a copy of the email will be sent to the &#x60;cc_email&#x60; specified by &#x60;Get email data&#x60; | [optional] 

## Methods

### NewEmailSchedule

`func NewEmailSchedule() *EmailSchedule`

NewEmailSchedule instantiates a new EmailSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailScheduleWithDefaults

`func NewEmailScheduleWithDefaults() *EmailSchedule`

NewEmailScheduleWithDefaults instantiates a new EmailSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenderId

`func (o *EmailSchedule) GetSenderId() int32`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *EmailSchedule) GetSenderIdOk() (*int32, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *EmailSchedule) SetSenderId(v int32)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *EmailSchedule) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### SetSenderIdNil

`func (o *EmailSchedule) SetSenderIdNil(b bool)`

 SetSenderIdNil sets the value for SenderId to be an explicit nil

### UnsetSenderId
`func (o *EmailSchedule) UnsetSenderId()`

UnsetSenderId ensures that no value is present for SenderId, not even an explicit nil
### GetSenderEmail

`func (o *EmailSchedule) GetSenderEmail() string`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *EmailSchedule) GetSenderEmailOk() (*string, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *EmailSchedule) SetSenderEmail(v string)`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *EmailSchedule) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.

### SetSenderEmailNil

`func (o *EmailSchedule) SetSenderEmailNil(b bool)`

 SetSenderEmailNil sets the value for SenderEmail to be an explicit nil

### UnsetSenderEmail
`func (o *EmailSchedule) UnsetSenderEmail()`

UnsetSenderEmail ensures that no value is present for SenderEmail, not even an explicit nil
### GetRecipientEmail

`func (o *EmailSchedule) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *EmailSchedule) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *EmailSchedule) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *EmailSchedule) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *EmailSchedule) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *EmailSchedule) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetSubject

`func (o *EmailSchedule) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailSchedule) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailSchedule) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailSchedule) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *EmailSchedule) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *EmailSchedule) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetBody

`func (o *EmailSchedule) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailSchedule) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailSchedule) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EmailSchedule) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *EmailSchedule) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *EmailSchedule) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetInclude

`func (o *EmailSchedule) GetInclude() EmailScheduleInclude`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *EmailSchedule) GetIncludeOk() (*EmailScheduleInclude, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *EmailSchedule) SetInclude(v EmailScheduleInclude)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *EmailSchedule) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### SetIncludeNil

`func (o *EmailSchedule) SetIncludeNil(b bool)`

 SetIncludeNil sets the value for Include to be an explicit nil

### UnsetInclude
`func (o *EmailSchedule) UnsetInclude()`

UnsetInclude ensures that no value is present for Include, not even an explicit nil
### GetAttachPdf

`func (o *EmailSchedule) GetAttachPdf() bool`

GetAttachPdf returns the AttachPdf field if non-nil, zero value otherwise.

### GetAttachPdfOk

`func (o *EmailSchedule) GetAttachPdfOk() (*bool, bool)`

GetAttachPdfOk returns a tuple with the AttachPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPdf

`func (o *EmailSchedule) SetAttachPdf(v bool)`

SetAttachPdf sets AttachPdf field to given value.

### HasAttachPdf

`func (o *EmailSchedule) HasAttachPdf() bool`

HasAttachPdf returns a boolean if a field has been set.

### SetAttachPdfNil

`func (o *EmailSchedule) SetAttachPdfNil(b bool)`

 SetAttachPdfNil sets the value for AttachPdf to be an explicit nil

### UnsetAttachPdf
`func (o *EmailSchedule) UnsetAttachPdf()`

UnsetAttachPdf ensures that no value is present for AttachPdf, not even an explicit nil
### GetSendCopy

`func (o *EmailSchedule) GetSendCopy() bool`

GetSendCopy returns the SendCopy field if non-nil, zero value otherwise.

### GetSendCopyOk

`func (o *EmailSchedule) GetSendCopyOk() (*bool, bool)`

GetSendCopyOk returns a tuple with the SendCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCopy

`func (o *EmailSchedule) SetSendCopy(v bool)`

SetSendCopy sets SendCopy field to given value.

### HasSendCopy

`func (o *EmailSchedule) HasSendCopy() bool`

HasSendCopy returns a boolean if a field has been set.

### SetSendCopyNil

`func (o *EmailSchedule) SetSendCopyNil(b bool)`

 SetSendCopyNil sets the value for SendCopy to be an explicit nil

### UnsetSendCopy
`func (o *EmailSchedule) UnsetSendCopy()`

UnsetSendCopy ensures that no value is present for SendCopy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


