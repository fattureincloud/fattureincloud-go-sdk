# PendingReceivedDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Pending received document id | [optional] 
**Date** | Pointer to **NullableString** | Pending received document date | [optional] 
**Subject** | Pointer to **NullableString** | Pending received document subject | [optional] 
**Filename** | Pointer to **NullableString** | Pending received document filename | [optional] 
**Type** | Pointer to [**PendingReceivedDocumentType**](PendingReceivedDocumentType.md) |  | [optional] [default to AGYO]
**AttachmentUrl** | Pointer to **NullableString** | [Temporary] [Read Only] Pending received document url of the attached file | [optional] [readonly] 
**AmountGross** | Pointer to **NullableFloat32** | [Read Only] Pending received document total gross amount | [optional] [readonly] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**DocumentType** | Pointer to [**ReceivedDocumentType**](ReceivedDocumentType.md) |  | [optional] [default to EXPENSE]
**SupplierName** | Pointer to **NullableString** | Pending received document supplier name | [optional] 
**CostCenter** | Pointer to **NullableString** | Pending received document cost center | [optional] 
**Category** | Pointer to **NullableString** | Pending received document category | [optional] 
**OtherAttachments** | Pointer to [**[]Attachment**](Attachment.md) | Pending received document other attachments | [optional] 
**EmssionDate** | Pointer to **NullableString** | Pending received document emission date | [optional] 
**PaymentsList** | Pointer to [**[]PendingReceivedDocumentPaymentsListItem**](PendingReceivedDocumentPaymentsListItem.md) |  | [optional] 
**AmountNet** | Pointer to **NullableFloat32** | Pending received document total net amount | [optional] 
**AmountVat** | Pointer to **NullableFloat32** | Pending received document total vat amount | [optional] 
**ImportError** | Pointer to **NullableString** | Pending received document import error | [optional] 
**ExtractedData** | Pointer to [**NullablePendingReceivedDocumentExtractedData**](PendingReceivedDocumentExtractedData.md) |  | [optional] 

## Methods

### NewPendingReceivedDocument

`func NewPendingReceivedDocument() *PendingReceivedDocument`

NewPendingReceivedDocument instantiates a new PendingReceivedDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingReceivedDocumentWithDefaults

`func NewPendingReceivedDocumentWithDefaults() *PendingReceivedDocument`

NewPendingReceivedDocumentWithDefaults instantiates a new PendingReceivedDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PendingReceivedDocument) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PendingReceivedDocument) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PendingReceivedDocument) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PendingReceivedDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PendingReceivedDocument) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PendingReceivedDocument) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDate

`func (o *PendingReceivedDocument) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PendingReceivedDocument) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PendingReceivedDocument) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PendingReceivedDocument) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *PendingReceivedDocument) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *PendingReceivedDocument) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetSubject

`func (o *PendingReceivedDocument) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PendingReceivedDocument) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PendingReceivedDocument) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PendingReceivedDocument) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *PendingReceivedDocument) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *PendingReceivedDocument) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetFilename

`func (o *PendingReceivedDocument) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PendingReceivedDocument) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PendingReceivedDocument) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PendingReceivedDocument) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *PendingReceivedDocument) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *PendingReceivedDocument) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetType

`func (o *PendingReceivedDocument) GetType() PendingReceivedDocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PendingReceivedDocument) GetTypeOk() (*PendingReceivedDocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PendingReceivedDocument) SetType(v PendingReceivedDocumentType)`

SetType sets Type field to given value.

### HasType

`func (o *PendingReceivedDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttachmentUrl

`func (o *PendingReceivedDocument) GetAttachmentUrl() string`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *PendingReceivedDocument) GetAttachmentUrlOk() (*string, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *PendingReceivedDocument) SetAttachmentUrl(v string)`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *PendingReceivedDocument) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### SetAttachmentUrlNil

`func (o *PendingReceivedDocument) SetAttachmentUrlNil(b bool)`

 SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil

### UnsetAttachmentUrl
`func (o *PendingReceivedDocument) UnsetAttachmentUrl()`

UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
### GetAmountGross

`func (o *PendingReceivedDocument) GetAmountGross() float32`

GetAmountGross returns the AmountGross field if non-nil, zero value otherwise.

### GetAmountGrossOk

`func (o *PendingReceivedDocument) GetAmountGrossOk() (*float32, bool)`

GetAmountGrossOk returns a tuple with the AmountGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGross

`func (o *PendingReceivedDocument) SetAmountGross(v float32)`

SetAmountGross sets AmountGross field to given value.

### HasAmountGross

`func (o *PendingReceivedDocument) HasAmountGross() bool`

HasAmountGross returns a boolean if a field has been set.

### SetAmountGrossNil

`func (o *PendingReceivedDocument) SetAmountGrossNil(b bool)`

 SetAmountGrossNil sets the value for AmountGross to be an explicit nil

### UnsetAmountGross
`func (o *PendingReceivedDocument) UnsetAmountGross()`

UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
### GetCurrency

`func (o *PendingReceivedDocument) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PendingReceivedDocument) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PendingReceivedDocument) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PendingReceivedDocument) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDocumentType

`func (o *PendingReceivedDocument) GetDocumentType() ReceivedDocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *PendingReceivedDocument) GetDocumentTypeOk() (*ReceivedDocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *PendingReceivedDocument) SetDocumentType(v ReceivedDocumentType)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *PendingReceivedDocument) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetSupplierName

`func (o *PendingReceivedDocument) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *PendingReceivedDocument) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *PendingReceivedDocument) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.

### HasSupplierName

`func (o *PendingReceivedDocument) HasSupplierName() bool`

HasSupplierName returns a boolean if a field has been set.

### SetSupplierNameNil

`func (o *PendingReceivedDocument) SetSupplierNameNil(b bool)`

 SetSupplierNameNil sets the value for SupplierName to be an explicit nil

### UnsetSupplierName
`func (o *PendingReceivedDocument) UnsetSupplierName()`

UnsetSupplierName ensures that no value is present for SupplierName, not even an explicit nil
### GetCostCenter

`func (o *PendingReceivedDocument) GetCostCenter() string`

GetCostCenter returns the CostCenter field if non-nil, zero value otherwise.

### GetCostCenterOk

`func (o *PendingReceivedDocument) GetCostCenterOk() (*string, bool)`

GetCostCenterOk returns a tuple with the CostCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCenter

`func (o *PendingReceivedDocument) SetCostCenter(v string)`

SetCostCenter sets CostCenter field to given value.

### HasCostCenter

`func (o *PendingReceivedDocument) HasCostCenter() bool`

HasCostCenter returns a boolean if a field has been set.

### SetCostCenterNil

`func (o *PendingReceivedDocument) SetCostCenterNil(b bool)`

 SetCostCenterNil sets the value for CostCenter to be an explicit nil

### UnsetCostCenter
`func (o *PendingReceivedDocument) UnsetCostCenter()`

UnsetCostCenter ensures that no value is present for CostCenter, not even an explicit nil
### GetCategory

`func (o *PendingReceivedDocument) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PendingReceivedDocument) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PendingReceivedDocument) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PendingReceivedDocument) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *PendingReceivedDocument) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *PendingReceivedDocument) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetOtherAttachments

`func (o *PendingReceivedDocument) GetOtherAttachments() []Attachment`

GetOtherAttachments returns the OtherAttachments field if non-nil, zero value otherwise.

### GetOtherAttachmentsOk

`func (o *PendingReceivedDocument) GetOtherAttachmentsOk() (*[]Attachment, bool)`

GetOtherAttachmentsOk returns a tuple with the OtherAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherAttachments

`func (o *PendingReceivedDocument) SetOtherAttachments(v []Attachment)`

SetOtherAttachments sets OtherAttachments field to given value.

### HasOtherAttachments

`func (o *PendingReceivedDocument) HasOtherAttachments() bool`

HasOtherAttachments returns a boolean if a field has been set.

### SetOtherAttachmentsNil

`func (o *PendingReceivedDocument) SetOtherAttachmentsNil(b bool)`

 SetOtherAttachmentsNil sets the value for OtherAttachments to be an explicit nil

### UnsetOtherAttachments
`func (o *PendingReceivedDocument) UnsetOtherAttachments()`

UnsetOtherAttachments ensures that no value is present for OtherAttachments, not even an explicit nil
### GetEmssionDate

`func (o *PendingReceivedDocument) GetEmssionDate() string`

GetEmssionDate returns the EmssionDate field if non-nil, zero value otherwise.

### GetEmssionDateOk

`func (o *PendingReceivedDocument) GetEmssionDateOk() (*string, bool)`

GetEmssionDateOk returns a tuple with the EmssionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmssionDate

`func (o *PendingReceivedDocument) SetEmssionDate(v string)`

SetEmssionDate sets EmssionDate field to given value.

### HasEmssionDate

`func (o *PendingReceivedDocument) HasEmssionDate() bool`

HasEmssionDate returns a boolean if a field has been set.

### SetEmssionDateNil

`func (o *PendingReceivedDocument) SetEmssionDateNil(b bool)`

 SetEmssionDateNil sets the value for EmssionDate to be an explicit nil

### UnsetEmssionDate
`func (o *PendingReceivedDocument) UnsetEmssionDate()`

UnsetEmssionDate ensures that no value is present for EmssionDate, not even an explicit nil
### GetPaymentsList

`func (o *PendingReceivedDocument) GetPaymentsList() []PendingReceivedDocumentPaymentsListItem`

GetPaymentsList returns the PaymentsList field if non-nil, zero value otherwise.

### GetPaymentsListOk

`func (o *PendingReceivedDocument) GetPaymentsListOk() (*[]PendingReceivedDocumentPaymentsListItem, bool)`

GetPaymentsListOk returns a tuple with the PaymentsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsList

`func (o *PendingReceivedDocument) SetPaymentsList(v []PendingReceivedDocumentPaymentsListItem)`

SetPaymentsList sets PaymentsList field to given value.

### HasPaymentsList

`func (o *PendingReceivedDocument) HasPaymentsList() bool`

HasPaymentsList returns a boolean if a field has been set.

### SetPaymentsListNil

`func (o *PendingReceivedDocument) SetPaymentsListNil(b bool)`

 SetPaymentsListNil sets the value for PaymentsList to be an explicit nil

### UnsetPaymentsList
`func (o *PendingReceivedDocument) UnsetPaymentsList()`

UnsetPaymentsList ensures that no value is present for PaymentsList, not even an explicit nil
### GetAmountNet

`func (o *PendingReceivedDocument) GetAmountNet() float32`

GetAmountNet returns the AmountNet field if non-nil, zero value otherwise.

### GetAmountNetOk

`func (o *PendingReceivedDocument) GetAmountNetOk() (*float32, bool)`

GetAmountNetOk returns a tuple with the AmountNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountNet

`func (o *PendingReceivedDocument) SetAmountNet(v float32)`

SetAmountNet sets AmountNet field to given value.

### HasAmountNet

`func (o *PendingReceivedDocument) HasAmountNet() bool`

HasAmountNet returns a boolean if a field has been set.

### SetAmountNetNil

`func (o *PendingReceivedDocument) SetAmountNetNil(b bool)`

 SetAmountNetNil sets the value for AmountNet to be an explicit nil

### UnsetAmountNet
`func (o *PendingReceivedDocument) UnsetAmountNet()`

UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
### GetAmountVat

`func (o *PendingReceivedDocument) GetAmountVat() float32`

GetAmountVat returns the AmountVat field if non-nil, zero value otherwise.

### GetAmountVatOk

`func (o *PendingReceivedDocument) GetAmountVatOk() (*float32, bool)`

GetAmountVatOk returns a tuple with the AmountVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountVat

`func (o *PendingReceivedDocument) SetAmountVat(v float32)`

SetAmountVat sets AmountVat field to given value.

### HasAmountVat

`func (o *PendingReceivedDocument) HasAmountVat() bool`

HasAmountVat returns a boolean if a field has been set.

### SetAmountVatNil

`func (o *PendingReceivedDocument) SetAmountVatNil(b bool)`

 SetAmountVatNil sets the value for AmountVat to be an explicit nil

### UnsetAmountVat
`func (o *PendingReceivedDocument) UnsetAmountVat()`

UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
### GetImportError

`func (o *PendingReceivedDocument) GetImportError() string`

GetImportError returns the ImportError field if non-nil, zero value otherwise.

### GetImportErrorOk

`func (o *PendingReceivedDocument) GetImportErrorOk() (*string, bool)`

GetImportErrorOk returns a tuple with the ImportError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportError

`func (o *PendingReceivedDocument) SetImportError(v string)`

SetImportError sets ImportError field to given value.

### HasImportError

`func (o *PendingReceivedDocument) HasImportError() bool`

HasImportError returns a boolean if a field has been set.

### SetImportErrorNil

`func (o *PendingReceivedDocument) SetImportErrorNil(b bool)`

 SetImportErrorNil sets the value for ImportError to be an explicit nil

### UnsetImportError
`func (o *PendingReceivedDocument) UnsetImportError()`

UnsetImportError ensures that no value is present for ImportError, not even an explicit nil
### GetExtractedData

`func (o *PendingReceivedDocument) GetExtractedData() PendingReceivedDocumentExtractedData`

GetExtractedData returns the ExtractedData field if non-nil, zero value otherwise.

### GetExtractedDataOk

`func (o *PendingReceivedDocument) GetExtractedDataOk() (*PendingReceivedDocumentExtractedData, bool)`

GetExtractedDataOk returns a tuple with the ExtractedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedData

`func (o *PendingReceivedDocument) SetExtractedData(v PendingReceivedDocumentExtractedData)`

SetExtractedData sets ExtractedData field to given value.

### HasExtractedData

`func (o *PendingReceivedDocument) HasExtractedData() bool`

HasExtractedData returns a boolean if a field has been set.

### SetExtractedDataNil

`func (o *PendingReceivedDocument) SetExtractedDataNil(b bool)`

 SetExtractedDataNil sets the value for ExtractedData to be an explicit nil

### UnsetExtractedData
`func (o *PendingReceivedDocument) UnsetExtractedData()`

UnsetExtractedData ensures that no value is present for ExtractedData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


