# ReceivedDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique identifier of the document. | [optional] 
**Type** | Pointer to [**ReceivedDocumentType**](ReceivedDocumentType.md) |  | [optional] [default to EXPENSE]
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 
**Date** | Pointer to **NullableString** | Date of the document [If not specified, today date is used]. | [optional] 
**Category** | Pointer to **NullableString** | Document category. | [optional] 
**Description** | Pointer to **NullableString** | Document description. | [optional] 
**AmountNet** | Pointer to **NullableFloat32** | Total net amount. | [optional] 
**AmountVat** | Pointer to **NullableFloat32** | Total vat amount. | [optional] 
**AmountWithholdingTax** | Pointer to **NullableFloat32** | Withholding tax amount. | [optional] 
**AmountOtherWithholdingTax** | Pointer to **NullableFloat32** | Other withholding tax amount. | [optional] 
**AmountGross** | Pointer to **NullableFloat32** | [Read Only] Total gross amount. | [optional] [readonly] 
**Amortization** | Pointer to **NullableFloat32** | Amortization value | [optional] 
**RcCenter** | Pointer to **NullableString** | Revenue center. | [optional] 
**InvoiceNumber** | Pointer to **NullableString** | Invoice number | [optional] 
**IsMarked** | Pointer to **NullableBool** |  | [optional] 
**IsDetailed** | Pointer to **NullableBool** |  | [optional] 
**EInvoice** | Pointer to **NullableBool** | [Read Only] Indicates if this is an e-invoice. | [optional] 
**NextDueDate** | Pointer to **NullableString** | [Read Only] Next due date. | [optional] [readonly] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**TaxDeductibility** | Pointer to **NullableFloat32** | Tax deducibility percentage. | [optional] 
**VatDeductibility** | Pointer to **NullableFloat32** | Vat deducibility percentage. | [optional] 
**ItemsList** | Pointer to [**[]ReceivedDocumentItemsListItem**](ReceivedDocumentItemsListItem.md) |  | [optional] 
**PaymentsList** | Pointer to [**[]ReceivedDocumentPaymentsListItem**](ReceivedDocumentPaymentsListItem.md) |  | [optional] 
**AttachmentUrl** | Pointer to **NullableString** | [Temporary] [Read Only]  Public url of the attached file. Authomatically set if a valid attachment token is passed via POST /received_documents or PUT /received_documents/{documentId}. | [optional] [readonly] 
**AttachmentPreviewUrl** | Pointer to **NullableString** | [Temporary] [Read Only]  Attachment preview url. | [optional] [readonly] 
**AutoCalculate** | Pointer to **NullableBool** | If set to false total items amount and total payments amount can be different. | [optional] 
**AttachmentToken** | Pointer to **NullableString** | Uploaded attachement token. | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReceivedDocument

`func NewReceivedDocument() *ReceivedDocument`

NewReceivedDocument instantiates a new ReceivedDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivedDocumentWithDefaults

`func NewReceivedDocumentWithDefaults() *ReceivedDocument`

NewReceivedDocumentWithDefaults instantiates a new ReceivedDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReceivedDocument) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceivedDocument) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceivedDocument) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReceivedDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ReceivedDocument) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ReceivedDocument) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *ReceivedDocument) GetType() ReceivedDocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReceivedDocument) GetTypeOk() (*ReceivedDocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReceivedDocument) SetType(v ReceivedDocumentType)`

SetType sets Type field to given value.

### HasType

`func (o *ReceivedDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEntity

`func (o *ReceivedDocument) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ReceivedDocument) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ReceivedDocument) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ReceivedDocument) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetDate

`func (o *ReceivedDocument) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ReceivedDocument) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ReceivedDocument) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ReceivedDocument) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *ReceivedDocument) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ReceivedDocument) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetCategory

`func (o *ReceivedDocument) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ReceivedDocument) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ReceivedDocument) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ReceivedDocument) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ReceivedDocument) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ReceivedDocument) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *ReceivedDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReceivedDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReceivedDocument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReceivedDocument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ReceivedDocument) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ReceivedDocument) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAmountNet

`func (o *ReceivedDocument) GetAmountNet() float32`

GetAmountNet returns the AmountNet field if non-nil, zero value otherwise.

### GetAmountNetOk

`func (o *ReceivedDocument) GetAmountNetOk() (*float32, bool)`

GetAmountNetOk returns a tuple with the AmountNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountNet

`func (o *ReceivedDocument) SetAmountNet(v float32)`

SetAmountNet sets AmountNet field to given value.

### HasAmountNet

`func (o *ReceivedDocument) HasAmountNet() bool`

HasAmountNet returns a boolean if a field has been set.

### SetAmountNetNil

`func (o *ReceivedDocument) SetAmountNetNil(b bool)`

 SetAmountNetNil sets the value for AmountNet to be an explicit nil

### UnsetAmountNet
`func (o *ReceivedDocument) UnsetAmountNet()`

UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
### GetAmountVat

`func (o *ReceivedDocument) GetAmountVat() float32`

GetAmountVat returns the AmountVat field if non-nil, zero value otherwise.

### GetAmountVatOk

`func (o *ReceivedDocument) GetAmountVatOk() (*float32, bool)`

GetAmountVatOk returns a tuple with the AmountVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountVat

`func (o *ReceivedDocument) SetAmountVat(v float32)`

SetAmountVat sets AmountVat field to given value.

### HasAmountVat

`func (o *ReceivedDocument) HasAmountVat() bool`

HasAmountVat returns a boolean if a field has been set.

### SetAmountVatNil

`func (o *ReceivedDocument) SetAmountVatNil(b bool)`

 SetAmountVatNil sets the value for AmountVat to be an explicit nil

### UnsetAmountVat
`func (o *ReceivedDocument) UnsetAmountVat()`

UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
### GetAmountWithholdingTax

`func (o *ReceivedDocument) GetAmountWithholdingTax() float32`

GetAmountWithholdingTax returns the AmountWithholdingTax field if non-nil, zero value otherwise.

### GetAmountWithholdingTaxOk

`func (o *ReceivedDocument) GetAmountWithholdingTaxOk() (*float32, bool)`

GetAmountWithholdingTaxOk returns a tuple with the AmountWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountWithholdingTax

`func (o *ReceivedDocument) SetAmountWithholdingTax(v float32)`

SetAmountWithholdingTax sets AmountWithholdingTax field to given value.

### HasAmountWithholdingTax

`func (o *ReceivedDocument) HasAmountWithholdingTax() bool`

HasAmountWithholdingTax returns a boolean if a field has been set.

### SetAmountWithholdingTaxNil

`func (o *ReceivedDocument) SetAmountWithholdingTaxNil(b bool)`

 SetAmountWithholdingTaxNil sets the value for AmountWithholdingTax to be an explicit nil

### UnsetAmountWithholdingTax
`func (o *ReceivedDocument) UnsetAmountWithholdingTax()`

UnsetAmountWithholdingTax ensures that no value is present for AmountWithholdingTax, not even an explicit nil
### GetAmountOtherWithholdingTax

`func (o *ReceivedDocument) GetAmountOtherWithholdingTax() float32`

GetAmountOtherWithholdingTax returns the AmountOtherWithholdingTax field if non-nil, zero value otherwise.

### GetAmountOtherWithholdingTaxOk

`func (o *ReceivedDocument) GetAmountOtherWithholdingTaxOk() (*float32, bool)`

GetAmountOtherWithholdingTaxOk returns a tuple with the AmountOtherWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOtherWithholdingTax

`func (o *ReceivedDocument) SetAmountOtherWithholdingTax(v float32)`

SetAmountOtherWithholdingTax sets AmountOtherWithholdingTax field to given value.

### HasAmountOtherWithholdingTax

`func (o *ReceivedDocument) HasAmountOtherWithholdingTax() bool`

HasAmountOtherWithholdingTax returns a boolean if a field has been set.

### SetAmountOtherWithholdingTaxNil

`func (o *ReceivedDocument) SetAmountOtherWithholdingTaxNil(b bool)`

 SetAmountOtherWithholdingTaxNil sets the value for AmountOtherWithholdingTax to be an explicit nil

### UnsetAmountOtherWithholdingTax
`func (o *ReceivedDocument) UnsetAmountOtherWithholdingTax()`

UnsetAmountOtherWithholdingTax ensures that no value is present for AmountOtherWithholdingTax, not even an explicit nil
### GetAmountGross

`func (o *ReceivedDocument) GetAmountGross() float32`

GetAmountGross returns the AmountGross field if non-nil, zero value otherwise.

### GetAmountGrossOk

`func (o *ReceivedDocument) GetAmountGrossOk() (*float32, bool)`

GetAmountGrossOk returns a tuple with the AmountGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGross

`func (o *ReceivedDocument) SetAmountGross(v float32)`

SetAmountGross sets AmountGross field to given value.

### HasAmountGross

`func (o *ReceivedDocument) HasAmountGross() bool`

HasAmountGross returns a boolean if a field has been set.

### SetAmountGrossNil

`func (o *ReceivedDocument) SetAmountGrossNil(b bool)`

 SetAmountGrossNil sets the value for AmountGross to be an explicit nil

### UnsetAmountGross
`func (o *ReceivedDocument) UnsetAmountGross()`

UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
### GetAmortization

`func (o *ReceivedDocument) GetAmortization() float32`

GetAmortization returns the Amortization field if non-nil, zero value otherwise.

### GetAmortizationOk

`func (o *ReceivedDocument) GetAmortizationOk() (*float32, bool)`

GetAmortizationOk returns a tuple with the Amortization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmortization

`func (o *ReceivedDocument) SetAmortization(v float32)`

SetAmortization sets Amortization field to given value.

### HasAmortization

`func (o *ReceivedDocument) HasAmortization() bool`

HasAmortization returns a boolean if a field has been set.

### SetAmortizationNil

`func (o *ReceivedDocument) SetAmortizationNil(b bool)`

 SetAmortizationNil sets the value for Amortization to be an explicit nil

### UnsetAmortization
`func (o *ReceivedDocument) UnsetAmortization()`

UnsetAmortization ensures that no value is present for Amortization, not even an explicit nil
### GetRcCenter

`func (o *ReceivedDocument) GetRcCenter() string`

GetRcCenter returns the RcCenter field if non-nil, zero value otherwise.

### GetRcCenterOk

`func (o *ReceivedDocument) GetRcCenterOk() (*string, bool)`

GetRcCenterOk returns a tuple with the RcCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcCenter

`func (o *ReceivedDocument) SetRcCenter(v string)`

SetRcCenter sets RcCenter field to given value.

### HasRcCenter

`func (o *ReceivedDocument) HasRcCenter() bool`

HasRcCenter returns a boolean if a field has been set.

### SetRcCenterNil

`func (o *ReceivedDocument) SetRcCenterNil(b bool)`

 SetRcCenterNil sets the value for RcCenter to be an explicit nil

### UnsetRcCenter
`func (o *ReceivedDocument) UnsetRcCenter()`

UnsetRcCenter ensures that no value is present for RcCenter, not even an explicit nil
### GetInvoiceNumber

`func (o *ReceivedDocument) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *ReceivedDocument) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *ReceivedDocument) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *ReceivedDocument) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### SetInvoiceNumberNil

`func (o *ReceivedDocument) SetInvoiceNumberNil(b bool)`

 SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil

### UnsetInvoiceNumber
`func (o *ReceivedDocument) UnsetInvoiceNumber()`

UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
### GetIsMarked

`func (o *ReceivedDocument) GetIsMarked() bool`

GetIsMarked returns the IsMarked field if non-nil, zero value otherwise.

### GetIsMarkedOk

`func (o *ReceivedDocument) GetIsMarkedOk() (*bool, bool)`

GetIsMarkedOk returns a tuple with the IsMarked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarked

`func (o *ReceivedDocument) SetIsMarked(v bool)`

SetIsMarked sets IsMarked field to given value.

### HasIsMarked

`func (o *ReceivedDocument) HasIsMarked() bool`

HasIsMarked returns a boolean if a field has been set.

### SetIsMarkedNil

`func (o *ReceivedDocument) SetIsMarkedNil(b bool)`

 SetIsMarkedNil sets the value for IsMarked to be an explicit nil

### UnsetIsMarked
`func (o *ReceivedDocument) UnsetIsMarked()`

UnsetIsMarked ensures that no value is present for IsMarked, not even an explicit nil
### GetIsDetailed

`func (o *ReceivedDocument) GetIsDetailed() bool`

GetIsDetailed returns the IsDetailed field if non-nil, zero value otherwise.

### GetIsDetailedOk

`func (o *ReceivedDocument) GetIsDetailedOk() (*bool, bool)`

GetIsDetailedOk returns a tuple with the IsDetailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDetailed

`func (o *ReceivedDocument) SetIsDetailed(v bool)`

SetIsDetailed sets IsDetailed field to given value.

### HasIsDetailed

`func (o *ReceivedDocument) HasIsDetailed() bool`

HasIsDetailed returns a boolean if a field has been set.

### SetIsDetailedNil

`func (o *ReceivedDocument) SetIsDetailedNil(b bool)`

 SetIsDetailedNil sets the value for IsDetailed to be an explicit nil

### UnsetIsDetailed
`func (o *ReceivedDocument) UnsetIsDetailed()`

UnsetIsDetailed ensures that no value is present for IsDetailed, not even an explicit nil
### GetEInvoice

`func (o *ReceivedDocument) GetEInvoice() bool`

GetEInvoice returns the EInvoice field if non-nil, zero value otherwise.

### GetEInvoiceOk

`func (o *ReceivedDocument) GetEInvoiceOk() (*bool, bool)`

GetEInvoiceOk returns a tuple with the EInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoice

`func (o *ReceivedDocument) SetEInvoice(v bool)`

SetEInvoice sets EInvoice field to given value.

### HasEInvoice

`func (o *ReceivedDocument) HasEInvoice() bool`

HasEInvoice returns a boolean if a field has been set.

### SetEInvoiceNil

`func (o *ReceivedDocument) SetEInvoiceNil(b bool)`

 SetEInvoiceNil sets the value for EInvoice to be an explicit nil

### UnsetEInvoice
`func (o *ReceivedDocument) UnsetEInvoice()`

UnsetEInvoice ensures that no value is present for EInvoice, not even an explicit nil
### GetNextDueDate

`func (o *ReceivedDocument) GetNextDueDate() string`

GetNextDueDate returns the NextDueDate field if non-nil, zero value otherwise.

### GetNextDueDateOk

`func (o *ReceivedDocument) GetNextDueDateOk() (*string, bool)`

GetNextDueDateOk returns a tuple with the NextDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDueDate

`func (o *ReceivedDocument) SetNextDueDate(v string)`

SetNextDueDate sets NextDueDate field to given value.

### HasNextDueDate

`func (o *ReceivedDocument) HasNextDueDate() bool`

HasNextDueDate returns a boolean if a field has been set.

### SetNextDueDateNil

`func (o *ReceivedDocument) SetNextDueDateNil(b bool)`

 SetNextDueDateNil sets the value for NextDueDate to be an explicit nil

### UnsetNextDueDate
`func (o *ReceivedDocument) UnsetNextDueDate()`

UnsetNextDueDate ensures that no value is present for NextDueDate, not even an explicit nil
### GetCurrency

`func (o *ReceivedDocument) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReceivedDocument) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReceivedDocument) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ReceivedDocument) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTaxDeductibility

`func (o *ReceivedDocument) GetTaxDeductibility() float32`

GetTaxDeductibility returns the TaxDeductibility field if non-nil, zero value otherwise.

### GetTaxDeductibilityOk

`func (o *ReceivedDocument) GetTaxDeductibilityOk() (*float32, bool)`

GetTaxDeductibilityOk returns a tuple with the TaxDeductibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDeductibility

`func (o *ReceivedDocument) SetTaxDeductibility(v float32)`

SetTaxDeductibility sets TaxDeductibility field to given value.

### HasTaxDeductibility

`func (o *ReceivedDocument) HasTaxDeductibility() bool`

HasTaxDeductibility returns a boolean if a field has been set.

### SetTaxDeductibilityNil

`func (o *ReceivedDocument) SetTaxDeductibilityNil(b bool)`

 SetTaxDeductibilityNil sets the value for TaxDeductibility to be an explicit nil

### UnsetTaxDeductibility
`func (o *ReceivedDocument) UnsetTaxDeductibility()`

UnsetTaxDeductibility ensures that no value is present for TaxDeductibility, not even an explicit nil
### GetVatDeductibility

`func (o *ReceivedDocument) GetVatDeductibility() float32`

GetVatDeductibility returns the VatDeductibility field if non-nil, zero value otherwise.

### GetVatDeductibilityOk

`func (o *ReceivedDocument) GetVatDeductibilityOk() (*float32, bool)`

GetVatDeductibilityOk returns a tuple with the VatDeductibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatDeductibility

`func (o *ReceivedDocument) SetVatDeductibility(v float32)`

SetVatDeductibility sets VatDeductibility field to given value.

### HasVatDeductibility

`func (o *ReceivedDocument) HasVatDeductibility() bool`

HasVatDeductibility returns a boolean if a field has been set.

### SetVatDeductibilityNil

`func (o *ReceivedDocument) SetVatDeductibilityNil(b bool)`

 SetVatDeductibilityNil sets the value for VatDeductibility to be an explicit nil

### UnsetVatDeductibility
`func (o *ReceivedDocument) UnsetVatDeductibility()`

UnsetVatDeductibility ensures that no value is present for VatDeductibility, not even an explicit nil
### GetItemsList

`func (o *ReceivedDocument) GetItemsList() []ReceivedDocumentItemsListItem`

GetItemsList returns the ItemsList field if non-nil, zero value otherwise.

### GetItemsListOk

`func (o *ReceivedDocument) GetItemsListOk() (*[]ReceivedDocumentItemsListItem, bool)`

GetItemsListOk returns a tuple with the ItemsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsList

`func (o *ReceivedDocument) SetItemsList(v []ReceivedDocumentItemsListItem)`

SetItemsList sets ItemsList field to given value.

### HasItemsList

`func (o *ReceivedDocument) HasItemsList() bool`

HasItemsList returns a boolean if a field has been set.

### SetItemsListNil

`func (o *ReceivedDocument) SetItemsListNil(b bool)`

 SetItemsListNil sets the value for ItemsList to be an explicit nil

### UnsetItemsList
`func (o *ReceivedDocument) UnsetItemsList()`

UnsetItemsList ensures that no value is present for ItemsList, not even an explicit nil
### GetPaymentsList

`func (o *ReceivedDocument) GetPaymentsList() []ReceivedDocumentPaymentsListItem`

GetPaymentsList returns the PaymentsList field if non-nil, zero value otherwise.

### GetPaymentsListOk

`func (o *ReceivedDocument) GetPaymentsListOk() (*[]ReceivedDocumentPaymentsListItem, bool)`

GetPaymentsListOk returns a tuple with the PaymentsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsList

`func (o *ReceivedDocument) SetPaymentsList(v []ReceivedDocumentPaymentsListItem)`

SetPaymentsList sets PaymentsList field to given value.

### HasPaymentsList

`func (o *ReceivedDocument) HasPaymentsList() bool`

HasPaymentsList returns a boolean if a field has been set.

### SetPaymentsListNil

`func (o *ReceivedDocument) SetPaymentsListNil(b bool)`

 SetPaymentsListNil sets the value for PaymentsList to be an explicit nil

### UnsetPaymentsList
`func (o *ReceivedDocument) UnsetPaymentsList()`

UnsetPaymentsList ensures that no value is present for PaymentsList, not even an explicit nil
### GetAttachmentUrl

`func (o *ReceivedDocument) GetAttachmentUrl() string`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *ReceivedDocument) GetAttachmentUrlOk() (*string, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *ReceivedDocument) SetAttachmentUrl(v string)`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *ReceivedDocument) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### SetAttachmentUrlNil

`func (o *ReceivedDocument) SetAttachmentUrlNil(b bool)`

 SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil

### UnsetAttachmentUrl
`func (o *ReceivedDocument) UnsetAttachmentUrl()`

UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
### GetAttachmentPreviewUrl

`func (o *ReceivedDocument) GetAttachmentPreviewUrl() string`

GetAttachmentPreviewUrl returns the AttachmentPreviewUrl field if non-nil, zero value otherwise.

### GetAttachmentPreviewUrlOk

`func (o *ReceivedDocument) GetAttachmentPreviewUrlOk() (*string, bool)`

GetAttachmentPreviewUrlOk returns a tuple with the AttachmentPreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentPreviewUrl

`func (o *ReceivedDocument) SetAttachmentPreviewUrl(v string)`

SetAttachmentPreviewUrl sets AttachmentPreviewUrl field to given value.

### HasAttachmentPreviewUrl

`func (o *ReceivedDocument) HasAttachmentPreviewUrl() bool`

HasAttachmentPreviewUrl returns a boolean if a field has been set.

### SetAttachmentPreviewUrlNil

`func (o *ReceivedDocument) SetAttachmentPreviewUrlNil(b bool)`

 SetAttachmentPreviewUrlNil sets the value for AttachmentPreviewUrl to be an explicit nil

### UnsetAttachmentPreviewUrl
`func (o *ReceivedDocument) UnsetAttachmentPreviewUrl()`

UnsetAttachmentPreviewUrl ensures that no value is present for AttachmentPreviewUrl, not even an explicit nil
### GetAutoCalculate

`func (o *ReceivedDocument) GetAutoCalculate() bool`

GetAutoCalculate returns the AutoCalculate field if non-nil, zero value otherwise.

### GetAutoCalculateOk

`func (o *ReceivedDocument) GetAutoCalculateOk() (*bool, bool)`

GetAutoCalculateOk returns a tuple with the AutoCalculate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCalculate

`func (o *ReceivedDocument) SetAutoCalculate(v bool)`

SetAutoCalculate sets AutoCalculate field to given value.

### HasAutoCalculate

`func (o *ReceivedDocument) HasAutoCalculate() bool`

HasAutoCalculate returns a boolean if a field has been set.

### SetAutoCalculateNil

`func (o *ReceivedDocument) SetAutoCalculateNil(b bool)`

 SetAutoCalculateNil sets the value for AutoCalculate to be an explicit nil

### UnsetAutoCalculate
`func (o *ReceivedDocument) UnsetAutoCalculate()`

UnsetAutoCalculate ensures that no value is present for AutoCalculate, not even an explicit nil
### GetAttachmentToken

`func (o *ReceivedDocument) GetAttachmentToken() string`

GetAttachmentToken returns the AttachmentToken field if non-nil, zero value otherwise.

### GetAttachmentTokenOk

`func (o *ReceivedDocument) GetAttachmentTokenOk() (*string, bool)`

GetAttachmentTokenOk returns a tuple with the AttachmentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentToken

`func (o *ReceivedDocument) SetAttachmentToken(v string)`

SetAttachmentToken sets AttachmentToken field to given value.

### HasAttachmentToken

`func (o *ReceivedDocument) HasAttachmentToken() bool`

HasAttachmentToken returns a boolean if a field has been set.

### SetAttachmentTokenNil

`func (o *ReceivedDocument) SetAttachmentTokenNil(b bool)`

 SetAttachmentTokenNil sets the value for AttachmentToken to be an explicit nil

### UnsetAttachmentToken
`func (o *ReceivedDocument) UnsetAttachmentToken()`

UnsetAttachmentToken ensures that no value is present for AttachmentToken, not even an explicit nil
### GetCreatedAt

`func (o *ReceivedDocument) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReceivedDocument) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReceivedDocument) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReceivedDocument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ReceivedDocument) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ReceivedDocument) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ReceivedDocument) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReceivedDocument) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReceivedDocument) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReceivedDocument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ReceivedDocument) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ReceivedDocument) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


