# EmailData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientEmail** | Pointer to **NullableString** | Recipient&#39;s email | [optional] 
**DefaultSenderEmail** | Pointer to [**NullableEmailDataDefaultSenderEmail**](EmailDataDefaultSenderEmail.md) |  | [optional] 
**SenderEmailsList** | Pointer to [**[]SenderEmail**](SenderEmail.md) | List of all emails from which the document can be sent | [optional] 
**CcEmail** | Pointer to **NullableString** | By default is the logged company email. This is the email address to which a copy will be sent. | [optional] 
**Subject** | Pointer to **NullableString** | Email subject | [optional] 
**Body** | Pointer to **NullableString** | Email body | [optional] 
**DocumentExists** | Pointer to **NullableBool** | If the document is not a delivery note, this flag will be set to true | [optional] 
**DeliveryNoteExists** | Pointer to **NullableBool** | If the document is a delivery note, this flag will be set to true | [optional] 
**AttachmentExists** | Pointer to **NullableBool** | If the document has one or more attachments, this flag will be set to true | [optional] 
**AccompanyingInvoiceExists** | Pointer to **NullableBool** | If an accompanying invoice exists, this flag will be set to true | [optional] 
**DefaultAttachPdf** | Pointer to **NullableBool** | If a pdf is attached, this flag will be set to true | [optional] 

## Methods

### NewEmailData

`func NewEmailData() *EmailData`

NewEmailData instantiates a new EmailData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDataWithDefaults

`func NewEmailDataWithDefaults() *EmailData`

NewEmailDataWithDefaults instantiates a new EmailData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientEmail

`func (o *EmailData) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *EmailData) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *EmailData) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *EmailData) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *EmailData) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *EmailData) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetDefaultSenderEmail

`func (o *EmailData) GetDefaultSenderEmail() EmailDataDefaultSenderEmail`

GetDefaultSenderEmail returns the DefaultSenderEmail field if non-nil, zero value otherwise.

### GetDefaultSenderEmailOk

`func (o *EmailData) GetDefaultSenderEmailOk() (*EmailDataDefaultSenderEmail, bool)`

GetDefaultSenderEmailOk returns a tuple with the DefaultSenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSenderEmail

`func (o *EmailData) SetDefaultSenderEmail(v EmailDataDefaultSenderEmail)`

SetDefaultSenderEmail sets DefaultSenderEmail field to given value.

### HasDefaultSenderEmail

`func (o *EmailData) HasDefaultSenderEmail() bool`

HasDefaultSenderEmail returns a boolean if a field has been set.

### SetDefaultSenderEmailNil

`func (o *EmailData) SetDefaultSenderEmailNil(b bool)`

 SetDefaultSenderEmailNil sets the value for DefaultSenderEmail to be an explicit nil

### UnsetDefaultSenderEmail
`func (o *EmailData) UnsetDefaultSenderEmail()`

UnsetDefaultSenderEmail ensures that no value is present for DefaultSenderEmail, not even an explicit nil
### GetSenderEmailsList

`func (o *EmailData) GetSenderEmailsList() []SenderEmail`

GetSenderEmailsList returns the SenderEmailsList field if non-nil, zero value otherwise.

### GetSenderEmailsListOk

`func (o *EmailData) GetSenderEmailsListOk() (*[]SenderEmail, bool)`

GetSenderEmailsListOk returns a tuple with the SenderEmailsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmailsList

`func (o *EmailData) SetSenderEmailsList(v []SenderEmail)`

SetSenderEmailsList sets SenderEmailsList field to given value.

### HasSenderEmailsList

`func (o *EmailData) HasSenderEmailsList() bool`

HasSenderEmailsList returns a boolean if a field has been set.

### SetSenderEmailsListNil

`func (o *EmailData) SetSenderEmailsListNil(b bool)`

 SetSenderEmailsListNil sets the value for SenderEmailsList to be an explicit nil

### UnsetSenderEmailsList
`func (o *EmailData) UnsetSenderEmailsList()`

UnsetSenderEmailsList ensures that no value is present for SenderEmailsList, not even an explicit nil
### GetCcEmail

`func (o *EmailData) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *EmailData) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *EmailData) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *EmailData) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### SetCcEmailNil

`func (o *EmailData) SetCcEmailNil(b bool)`

 SetCcEmailNil sets the value for CcEmail to be an explicit nil

### UnsetCcEmail
`func (o *EmailData) UnsetCcEmail()`

UnsetCcEmail ensures that no value is present for CcEmail, not even an explicit nil
### GetSubject

`func (o *EmailData) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailData) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailData) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailData) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *EmailData) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *EmailData) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetBody

`func (o *EmailData) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailData) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailData) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EmailData) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *EmailData) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *EmailData) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetDocumentExists

`func (o *EmailData) GetDocumentExists() bool`

GetDocumentExists returns the DocumentExists field if non-nil, zero value otherwise.

### GetDocumentExistsOk

`func (o *EmailData) GetDocumentExistsOk() (*bool, bool)`

GetDocumentExistsOk returns a tuple with the DocumentExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentExists

`func (o *EmailData) SetDocumentExists(v bool)`

SetDocumentExists sets DocumentExists field to given value.

### HasDocumentExists

`func (o *EmailData) HasDocumentExists() bool`

HasDocumentExists returns a boolean if a field has been set.

### SetDocumentExistsNil

`func (o *EmailData) SetDocumentExistsNil(b bool)`

 SetDocumentExistsNil sets the value for DocumentExists to be an explicit nil

### UnsetDocumentExists
`func (o *EmailData) UnsetDocumentExists()`

UnsetDocumentExists ensures that no value is present for DocumentExists, not even an explicit nil
### GetDeliveryNoteExists

`func (o *EmailData) GetDeliveryNoteExists() bool`

GetDeliveryNoteExists returns the DeliveryNoteExists field if non-nil, zero value otherwise.

### GetDeliveryNoteExistsOk

`func (o *EmailData) GetDeliveryNoteExistsOk() (*bool, bool)`

GetDeliveryNoteExistsOk returns a tuple with the DeliveryNoteExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNoteExists

`func (o *EmailData) SetDeliveryNoteExists(v bool)`

SetDeliveryNoteExists sets DeliveryNoteExists field to given value.

### HasDeliveryNoteExists

`func (o *EmailData) HasDeliveryNoteExists() bool`

HasDeliveryNoteExists returns a boolean if a field has been set.

### SetDeliveryNoteExistsNil

`func (o *EmailData) SetDeliveryNoteExistsNil(b bool)`

 SetDeliveryNoteExistsNil sets the value for DeliveryNoteExists to be an explicit nil

### UnsetDeliveryNoteExists
`func (o *EmailData) UnsetDeliveryNoteExists()`

UnsetDeliveryNoteExists ensures that no value is present for DeliveryNoteExists, not even an explicit nil
### GetAttachmentExists

`func (o *EmailData) GetAttachmentExists() bool`

GetAttachmentExists returns the AttachmentExists field if non-nil, zero value otherwise.

### GetAttachmentExistsOk

`func (o *EmailData) GetAttachmentExistsOk() (*bool, bool)`

GetAttachmentExistsOk returns a tuple with the AttachmentExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentExists

`func (o *EmailData) SetAttachmentExists(v bool)`

SetAttachmentExists sets AttachmentExists field to given value.

### HasAttachmentExists

`func (o *EmailData) HasAttachmentExists() bool`

HasAttachmentExists returns a boolean if a field has been set.

### SetAttachmentExistsNil

`func (o *EmailData) SetAttachmentExistsNil(b bool)`

 SetAttachmentExistsNil sets the value for AttachmentExists to be an explicit nil

### UnsetAttachmentExists
`func (o *EmailData) UnsetAttachmentExists()`

UnsetAttachmentExists ensures that no value is present for AttachmentExists, not even an explicit nil
### GetAccompanyingInvoiceExists

`func (o *EmailData) GetAccompanyingInvoiceExists() bool`

GetAccompanyingInvoiceExists returns the AccompanyingInvoiceExists field if non-nil, zero value otherwise.

### GetAccompanyingInvoiceExistsOk

`func (o *EmailData) GetAccompanyingInvoiceExistsOk() (*bool, bool)`

GetAccompanyingInvoiceExistsOk returns a tuple with the AccompanyingInvoiceExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccompanyingInvoiceExists

`func (o *EmailData) SetAccompanyingInvoiceExists(v bool)`

SetAccompanyingInvoiceExists sets AccompanyingInvoiceExists field to given value.

### HasAccompanyingInvoiceExists

`func (o *EmailData) HasAccompanyingInvoiceExists() bool`

HasAccompanyingInvoiceExists returns a boolean if a field has been set.

### SetAccompanyingInvoiceExistsNil

`func (o *EmailData) SetAccompanyingInvoiceExistsNil(b bool)`

 SetAccompanyingInvoiceExistsNil sets the value for AccompanyingInvoiceExists to be an explicit nil

### UnsetAccompanyingInvoiceExists
`func (o *EmailData) UnsetAccompanyingInvoiceExists()`

UnsetAccompanyingInvoiceExists ensures that no value is present for AccompanyingInvoiceExists, not even an explicit nil
### GetDefaultAttachPdf

`func (o *EmailData) GetDefaultAttachPdf() bool`

GetDefaultAttachPdf returns the DefaultAttachPdf field if non-nil, zero value otherwise.

### GetDefaultAttachPdfOk

`func (o *EmailData) GetDefaultAttachPdfOk() (*bool, bool)`

GetDefaultAttachPdfOk returns a tuple with the DefaultAttachPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAttachPdf

`func (o *EmailData) SetDefaultAttachPdf(v bool)`

SetDefaultAttachPdf sets DefaultAttachPdf field to given value.

### HasDefaultAttachPdf

`func (o *EmailData) HasDefaultAttachPdf() bool`

HasDefaultAttachPdf returns a boolean if a field has been set.

### SetDefaultAttachPdfNil

`func (o *EmailData) SetDefaultAttachPdfNil(b bool)`

 SetDefaultAttachPdfNil sets the value for DefaultAttachPdf to be an explicit nil

### UnsetDefaultAttachPdf
`func (o *EmailData) UnsetDefaultAttachPdf()`

UnsetDefaultAttachPdf ensures that no value is present for DefaultAttachPdf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


