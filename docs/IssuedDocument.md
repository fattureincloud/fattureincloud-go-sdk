# IssuedDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Issued document id | [optional] 
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 
**Type** | Pointer to [**IssuedDocumentType**](IssuedDocumentType.md) |  | [optional] [default to INVOICE]
**Number** | Pointer to **NullableInt32** | Issued document number [If not specified, next number is used] | [optional] 
**Numeration** | Pointer to **NullableString** | Issued document numeration [Not available if type&#x3D;delivery_note] | [optional] 
**Date** | Pointer to **NullableString** | Issued document date [defaults to today&#39;s date] | [optional] 
**Year** | Pointer to **NullableInt32** | Issued document year | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 
**Subject** | Pointer to **NullableString** | Issued document subject [not shown on the PDF] | [optional] 
**VisibleSubject** | Pointer to **NullableString** | Issued document visible subject | [optional] 
**RcCenter** | Pointer to **NullableString** | Issued document revenue center [or cost center if type&#x3D;supplier_order]. | [optional] 
**Notes** | Pointer to **NullableString** | Issued document extra notes | [optional] 
**Rivalsa** | Pointer to **NullableFloat32** | Issued document \&quot;Rivalsa INPS\&quot; percentual value | [optional] 
**Cassa** | Pointer to **NullableFloat32** | Issued document \&quot;Cassa previdenziale\&quot; percentual value | [optional] 
**AmountCassa** | Pointer to **NullableFloat32** | [Read Only] Issued document cassa amount. | [optional] [readonly] 
**CassaTaxable** | Pointer to **NullableFloat32** | Issued document cassa taxable percentage | [optional] 
**AmountCassaTaxable** | Pointer to **NullableFloat32** | [Can be set only if cassa_taxable is NULL] Issued document cassa taxable amount | [optional] 
**Cassa2** | Pointer to **NullableFloat32** | Issued document \&quot;Cassa previdenziale 2\&quot; percentual value | [optional] 
**AmountCassa2** | Pointer to **NullableFloat32** | [Read Only] Issued document cassa2 amount | [optional] [readonly] 
**Cassa2Taxable** | Pointer to **NullableFloat32** | Issued document cassa2 taxable percentage | [optional] 
**AmountCassa2Taxable** | Pointer to **NullableFloat32** | [Can be set only if cassa2_taxable is NULL] Issued document cassa2 taxable amount | [optional] 
**GlobalCassaTaxable** | Pointer to **NullableFloat32** | Issued document global cassa taxable percentage | [optional] 
**AmountGlobalCassaTaxable** | Pointer to **NullableFloat32** | [Can be set only if global_cassa_taxable is NULL] Issued document global cassa taxable amount | [optional] 
**WithholdingTax** | Pointer to **NullableFloat32** | Issued document withholding tax (ritenuta d&#39;acconto) percentual value | [optional] 
**WithholdingTaxTaxable** | Pointer to **NullableFloat32** | Issued document withholding tax taxable (imponibile) percentual value | [optional] 
**OtherWithholdingTax** | Pointer to **NullableFloat32** | Issued document other withholding tax (altra ritenuta) percentual value | [optional] 
**StampDuty** | Pointer to **NullableFloat32** | Issued document stamp duty value [0 if not present] | [optional] 
**PaymentMethod** | Pointer to [**PaymentMethod**](PaymentMethod.md) |  | [optional] 
**UseSplitPayment** | Pointer to **NullableBool** | Issued document uses split payment | [optional] 
**UseGrossPrices** | Pointer to **NullableBool** | Issued document uses gross prices | [optional] 
**EInvoice** | Pointer to **NullableBool** | Issued document is an e-invoice. | [optional] 
**EiData** | Pointer to [**NullableIssuedDocumentEiData**](IssuedDocumentEiData.md) |  | [optional] 
**EiCassaType** | Pointer to **NullableString** | E-invoice cassa type | [optional] 
**EiCassa2Type** | Pointer to **NullableString** | E-invoice cassa2 type | [optional] 
**EiWithholdingTaxCausal** | Pointer to **NullableString** | E-invoice withholding tax causal | [optional] 
**EiOtherWithholdingTaxType** | Pointer to **NullableString** | E-invoice other withholding tax type | [optional] 
**EiOtherWithholdingTaxCausal** | Pointer to **NullableString** | E-invoice other withholding tax causal | [optional] 
**ItemsList** | Pointer to [**[]IssuedDocumentItemsListItem**](IssuedDocumentItemsListItem.md) |  | [optional] 
**PaymentsList** | Pointer to [**[]IssuedDocumentPaymentsListItem**](IssuedDocumentPaymentsListItem.md) |  | [optional] 
**Template** | Pointer to [**DocumentTemplate**](DocumentTemplate.md) |  | [optional] 
**DeliveryNoteTemplate** | Pointer to [**DocumentTemplate**](DocumentTemplate.md) |  | [optional] 
**AccInvTemplate** | Pointer to [**DocumentTemplate**](DocumentTemplate.md) |  | [optional] 
**HMargins** | Pointer to **NullableInt32** | Issued document PDF horizontal margins | [optional] 
**VMargins** | Pointer to **NullableInt32** | Issued document PDF vertical margins | [optional] 
**ShowPayments** | Pointer to **NullableBool** | Show the expiration dates of the payments on the document | [optional] 
**ShowPaymentMethod** | Pointer to **NullableBool** | Show the payment method details on the document | [optional] 
**ShowTotals** | Pointer to [**ShowTotalsMode**](ShowTotalsMode.md) |  | [optional] [default to ALL]
**ShowNotificationButton** | Pointer to **NullableBool** | Show notification button in the PDF | [optional] 
**ShowTspayButton** | Pointer to **NullableBool** | Show ts pay button in the PDF | [optional] 
**DeliveryNote** | Pointer to **NullableBool** | Issued document has delivery note | [optional] 
**AccompanyingInvoice** | Pointer to **NullableBool** | Issued document has an accompanying invoice | [optional] 
**DnNumber** | Pointer to **NullableInt32** | Issued document attached delivery note number | [optional] 
**DnDate** | Pointer to **NullableString** | Issued document attached delivery note date | [optional] 
**DnAiPackagesNumber** | Pointer to **NullableString** | Issued document attached delivery note number of packages | [optional] 
**DnAiWeight** | Pointer to **NullableString** | Issued document attached delivery note package weight | [optional] 
**DnAiCausal** | Pointer to **NullableString** | Issued document attached delivery note causal | [optional] 
**DnAiDestination** | Pointer to **NullableString** | Issued document attached delivery note destination | [optional] 
**DnAiTransporter** | Pointer to **NullableString** | Issued document attached delivery note transporter | [optional] 
**DnAiNotes** | Pointer to **NullableString** | Issued document attached delivery note notes | [optional] 
**IsMarked** | Pointer to **NullableBool** | Issued document is marked | [optional] 
**AmountNet** | Pointer to **NullableFloat32** | [Read only] Issued document total net amount | [optional] [readonly] 
**AmountVat** | Pointer to **NullableFloat32** | [Read Only] Issued document total vat amount | [optional] [readonly] 
**AmountGross** | Pointer to **NullableFloat32** | [Read Only] Issued document total gross amount | [optional] [readonly] 
**AmountDueDiscount** | Pointer to **NullableFloat32** | Issued document amount due discount | [optional] 
**AmountRivalsa** | Pointer to **NullableFloat32** | [Read Only] Issued document rivalsa amount | [optional] [readonly] 
**AmountRivalsaTaxable** | Pointer to **NullableFloat32** | Issued document taxable rivalsa amount | [optional] 
**AmountWithholdingTax** | Pointer to **NullableFloat32** | [Read Only] Issued document withholding tax amount (ritenuta d&#39;acconto). | [optional] [readonly] 
**AmountWithholdingTaxTaxable** | Pointer to **NullableFloat32** | Issued document taxable withholding tax amount | [optional] 
**AmountOtherWithholdingTax** | Pointer to **NullableFloat32** | [Read Only] Issued document other withholding tax amount (altra ritenuta) | [optional] [readonly] 
**AmountOtherWithholdingTaxTaxable** | Pointer to **NullableFloat32** | Issued document taxable other withholding tax amount | [optional] 
**AmountEnasarcoTaxable** | Pointer to **NullableFloat32** | Issued document taxable enasarco amount | [optional] 
**ExtraData** | Pointer to [**NullableIssuedDocumentExtraData**](IssuedDocumentExtraData.md) |  | [optional] 
**SeenDate** | Pointer to **NullableString** | Issued document seen date | [optional] 
**NextDueDate** | Pointer to **NullableString** | Issued document date of the next not paid payment | [optional] 
**Url** | Pointer to **NullableString** | [Temporary] [Read Only] Issued document url of the document PDF file | [optional] 
**DnUrl** | Pointer to **NullableString** | [Temporary] [Read Only] Issued document url of the attached delivery note PDF file | [optional] 
**AiUrl** | Pointer to **NullableString** | [Temporary] [Read Only] Issued document url of the accompanying invoice PDF file | [optional] 
**AttachmentUrl** | Pointer to **NullableString** | [Temporary] [Read Only] Issued document url of the attached file | [optional] [readonly] 
**AttachmentToken** | Pointer to **NullableString** | [Write Only] Issued document attachment token returned by POST /issued_documents/attachment | [optional] 
**EiRaw** | Pointer to **map[string]interface{}** | Issued document advanced raw attributes for e-invoices | [optional] 
**EiStatus** | Pointer to **NullableString** | [Read only] Status of the e-invoice.   * **attempt** - We are trying to send the invoice, please wait up to 2 hours   * **missing** - The invoice is missing   * **not_sent** - The invoice has yet to be sent   * **sent** - The invoice was sent   * **pending** - The checks for the digital signature and sending are in progress   * **processing** - The SDI is delivering the invoice to the customer   * **error** - An error occurred while handling the invoice, please try to resend it or contact support   * **discarded** - The invoice has been rejected by the SDI, so it must be corrected and re-sent   * **not_delivered** - The SDI was unable to deliver the invoice   * **accepted** - The customer accepted the invoice   * **rejected** - The customer rejected the invoice, so it must be corrected   * **no_response** - A response has not yet been received whithin the deadline, contact the customer to ascertain the status of the invoice   * **manual_accepted** - The customer accepted the invoice   * **manual_rejected** - The customer rejected the invoice  | [optional] 
**Locked** | Pointer to **NullableBool** | Issued Document can&#39;t be edited | [optional] 
**CreatedAt** | Pointer to **NullableString** | Issued document creation date | [optional] 
**UpdatedAt** | Pointer to **NullableString** | Issued document last update date | [optional] 

## Methods

### NewIssuedDocument

`func NewIssuedDocument() *IssuedDocument`

NewIssuedDocument instantiates a new IssuedDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentWithDefaults

`func NewIssuedDocumentWithDefaults() *IssuedDocument`

NewIssuedDocumentWithDefaults instantiates a new IssuedDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssuedDocument) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssuedDocument) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssuedDocument) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IssuedDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IssuedDocument) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IssuedDocument) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetEntity

`func (o *IssuedDocument) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *IssuedDocument) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *IssuedDocument) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *IssuedDocument) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetType

`func (o *IssuedDocument) GetType() IssuedDocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssuedDocument) GetTypeOk() (*IssuedDocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssuedDocument) SetType(v IssuedDocumentType)`

SetType sets Type field to given value.

### HasType

`func (o *IssuedDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumber

`func (o *IssuedDocument) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IssuedDocument) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IssuedDocument) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *IssuedDocument) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *IssuedDocument) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *IssuedDocument) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetNumeration

`func (o *IssuedDocument) GetNumeration() string`

GetNumeration returns the Numeration field if non-nil, zero value otherwise.

### GetNumerationOk

`func (o *IssuedDocument) GetNumerationOk() (*string, bool)`

GetNumerationOk returns a tuple with the Numeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeration

`func (o *IssuedDocument) SetNumeration(v string)`

SetNumeration sets Numeration field to given value.

### HasNumeration

`func (o *IssuedDocument) HasNumeration() bool`

HasNumeration returns a boolean if a field has been set.

### SetNumerationNil

`func (o *IssuedDocument) SetNumerationNil(b bool)`

 SetNumerationNil sets the value for Numeration to be an explicit nil

### UnsetNumeration
`func (o *IssuedDocument) UnsetNumeration()`

UnsetNumeration ensures that no value is present for Numeration, not even an explicit nil
### GetDate

`func (o *IssuedDocument) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *IssuedDocument) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *IssuedDocument) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *IssuedDocument) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *IssuedDocument) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *IssuedDocument) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetYear

`func (o *IssuedDocument) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *IssuedDocument) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *IssuedDocument) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *IssuedDocument) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *IssuedDocument) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *IssuedDocument) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetCurrency

`func (o *IssuedDocument) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IssuedDocument) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IssuedDocument) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IssuedDocument) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLanguage

`func (o *IssuedDocument) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *IssuedDocument) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *IssuedDocument) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *IssuedDocument) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetSubject

`func (o *IssuedDocument) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IssuedDocument) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IssuedDocument) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IssuedDocument) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *IssuedDocument) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *IssuedDocument) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetVisibleSubject

`func (o *IssuedDocument) GetVisibleSubject() string`

GetVisibleSubject returns the VisibleSubject field if non-nil, zero value otherwise.

### GetVisibleSubjectOk

`func (o *IssuedDocument) GetVisibleSubjectOk() (*string, bool)`

GetVisibleSubjectOk returns a tuple with the VisibleSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSubject

`func (o *IssuedDocument) SetVisibleSubject(v string)`

SetVisibleSubject sets VisibleSubject field to given value.

### HasVisibleSubject

`func (o *IssuedDocument) HasVisibleSubject() bool`

HasVisibleSubject returns a boolean if a field has been set.

### SetVisibleSubjectNil

`func (o *IssuedDocument) SetVisibleSubjectNil(b bool)`

 SetVisibleSubjectNil sets the value for VisibleSubject to be an explicit nil

### UnsetVisibleSubject
`func (o *IssuedDocument) UnsetVisibleSubject()`

UnsetVisibleSubject ensures that no value is present for VisibleSubject, not even an explicit nil
### GetRcCenter

`func (o *IssuedDocument) GetRcCenter() string`

GetRcCenter returns the RcCenter field if non-nil, zero value otherwise.

### GetRcCenterOk

`func (o *IssuedDocument) GetRcCenterOk() (*string, bool)`

GetRcCenterOk returns a tuple with the RcCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcCenter

`func (o *IssuedDocument) SetRcCenter(v string)`

SetRcCenter sets RcCenter field to given value.

### HasRcCenter

`func (o *IssuedDocument) HasRcCenter() bool`

HasRcCenter returns a boolean if a field has been set.

### SetRcCenterNil

`func (o *IssuedDocument) SetRcCenterNil(b bool)`

 SetRcCenterNil sets the value for RcCenter to be an explicit nil

### UnsetRcCenter
`func (o *IssuedDocument) UnsetRcCenter()`

UnsetRcCenter ensures that no value is present for RcCenter, not even an explicit nil
### GetNotes

`func (o *IssuedDocument) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *IssuedDocument) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *IssuedDocument) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *IssuedDocument) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *IssuedDocument) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *IssuedDocument) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetRivalsa

`func (o *IssuedDocument) GetRivalsa() float32`

GetRivalsa returns the Rivalsa field if non-nil, zero value otherwise.

### GetRivalsaOk

`func (o *IssuedDocument) GetRivalsaOk() (*float32, bool)`

GetRivalsaOk returns a tuple with the Rivalsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRivalsa

`func (o *IssuedDocument) SetRivalsa(v float32)`

SetRivalsa sets Rivalsa field to given value.

### HasRivalsa

`func (o *IssuedDocument) HasRivalsa() bool`

HasRivalsa returns a boolean if a field has been set.

### SetRivalsaNil

`func (o *IssuedDocument) SetRivalsaNil(b bool)`

 SetRivalsaNil sets the value for Rivalsa to be an explicit nil

### UnsetRivalsa
`func (o *IssuedDocument) UnsetRivalsa()`

UnsetRivalsa ensures that no value is present for Rivalsa, not even an explicit nil
### GetCassa

`func (o *IssuedDocument) GetCassa() float32`

GetCassa returns the Cassa field if non-nil, zero value otherwise.

### GetCassaOk

`func (o *IssuedDocument) GetCassaOk() (*float32, bool)`

GetCassaOk returns a tuple with the Cassa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassa

`func (o *IssuedDocument) SetCassa(v float32)`

SetCassa sets Cassa field to given value.

### HasCassa

`func (o *IssuedDocument) HasCassa() bool`

HasCassa returns a boolean if a field has been set.

### SetCassaNil

`func (o *IssuedDocument) SetCassaNil(b bool)`

 SetCassaNil sets the value for Cassa to be an explicit nil

### UnsetCassa
`func (o *IssuedDocument) UnsetCassa()`

UnsetCassa ensures that no value is present for Cassa, not even an explicit nil
### GetAmountCassa

`func (o *IssuedDocument) GetAmountCassa() float32`

GetAmountCassa returns the AmountCassa field if non-nil, zero value otherwise.

### GetAmountCassaOk

`func (o *IssuedDocument) GetAmountCassaOk() (*float32, bool)`

GetAmountCassaOk returns a tuple with the AmountCassa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCassa

`func (o *IssuedDocument) SetAmountCassa(v float32)`

SetAmountCassa sets AmountCassa field to given value.

### HasAmountCassa

`func (o *IssuedDocument) HasAmountCassa() bool`

HasAmountCassa returns a boolean if a field has been set.

### SetAmountCassaNil

`func (o *IssuedDocument) SetAmountCassaNil(b bool)`

 SetAmountCassaNil sets the value for AmountCassa to be an explicit nil

### UnsetAmountCassa
`func (o *IssuedDocument) UnsetAmountCassa()`

UnsetAmountCassa ensures that no value is present for AmountCassa, not even an explicit nil
### GetCassaTaxable

`func (o *IssuedDocument) GetCassaTaxable() float32`

GetCassaTaxable returns the CassaTaxable field if non-nil, zero value otherwise.

### GetCassaTaxableOk

`func (o *IssuedDocument) GetCassaTaxableOk() (*float32, bool)`

GetCassaTaxableOk returns a tuple with the CassaTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassaTaxable

`func (o *IssuedDocument) SetCassaTaxable(v float32)`

SetCassaTaxable sets CassaTaxable field to given value.

### HasCassaTaxable

`func (o *IssuedDocument) HasCassaTaxable() bool`

HasCassaTaxable returns a boolean if a field has been set.

### SetCassaTaxableNil

`func (o *IssuedDocument) SetCassaTaxableNil(b bool)`

 SetCassaTaxableNil sets the value for CassaTaxable to be an explicit nil

### UnsetCassaTaxable
`func (o *IssuedDocument) UnsetCassaTaxable()`

UnsetCassaTaxable ensures that no value is present for CassaTaxable, not even an explicit nil
### GetAmountCassaTaxable

`func (o *IssuedDocument) GetAmountCassaTaxable() float32`

GetAmountCassaTaxable returns the AmountCassaTaxable field if non-nil, zero value otherwise.

### GetAmountCassaTaxableOk

`func (o *IssuedDocument) GetAmountCassaTaxableOk() (*float32, bool)`

GetAmountCassaTaxableOk returns a tuple with the AmountCassaTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCassaTaxable

`func (o *IssuedDocument) SetAmountCassaTaxable(v float32)`

SetAmountCassaTaxable sets AmountCassaTaxable field to given value.

### HasAmountCassaTaxable

`func (o *IssuedDocument) HasAmountCassaTaxable() bool`

HasAmountCassaTaxable returns a boolean if a field has been set.

### SetAmountCassaTaxableNil

`func (o *IssuedDocument) SetAmountCassaTaxableNil(b bool)`

 SetAmountCassaTaxableNil sets the value for AmountCassaTaxable to be an explicit nil

### UnsetAmountCassaTaxable
`func (o *IssuedDocument) UnsetAmountCassaTaxable()`

UnsetAmountCassaTaxable ensures that no value is present for AmountCassaTaxable, not even an explicit nil
### GetCassa2

`func (o *IssuedDocument) GetCassa2() float32`

GetCassa2 returns the Cassa2 field if non-nil, zero value otherwise.

### GetCassa2Ok

`func (o *IssuedDocument) GetCassa2Ok() (*float32, bool)`

GetCassa2Ok returns a tuple with the Cassa2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassa2

`func (o *IssuedDocument) SetCassa2(v float32)`

SetCassa2 sets Cassa2 field to given value.

### HasCassa2

`func (o *IssuedDocument) HasCassa2() bool`

HasCassa2 returns a boolean if a field has been set.

### SetCassa2Nil

`func (o *IssuedDocument) SetCassa2Nil(b bool)`

 SetCassa2Nil sets the value for Cassa2 to be an explicit nil

### UnsetCassa2
`func (o *IssuedDocument) UnsetCassa2()`

UnsetCassa2 ensures that no value is present for Cassa2, not even an explicit nil
### GetAmountCassa2

`func (o *IssuedDocument) GetAmountCassa2() float32`

GetAmountCassa2 returns the AmountCassa2 field if non-nil, zero value otherwise.

### GetAmountCassa2Ok

`func (o *IssuedDocument) GetAmountCassa2Ok() (*float32, bool)`

GetAmountCassa2Ok returns a tuple with the AmountCassa2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCassa2

`func (o *IssuedDocument) SetAmountCassa2(v float32)`

SetAmountCassa2 sets AmountCassa2 field to given value.

### HasAmountCassa2

`func (o *IssuedDocument) HasAmountCassa2() bool`

HasAmountCassa2 returns a boolean if a field has been set.

### SetAmountCassa2Nil

`func (o *IssuedDocument) SetAmountCassa2Nil(b bool)`

 SetAmountCassa2Nil sets the value for AmountCassa2 to be an explicit nil

### UnsetAmountCassa2
`func (o *IssuedDocument) UnsetAmountCassa2()`

UnsetAmountCassa2 ensures that no value is present for AmountCassa2, not even an explicit nil
### GetCassa2Taxable

`func (o *IssuedDocument) GetCassa2Taxable() float32`

GetCassa2Taxable returns the Cassa2Taxable field if non-nil, zero value otherwise.

### GetCassa2TaxableOk

`func (o *IssuedDocument) GetCassa2TaxableOk() (*float32, bool)`

GetCassa2TaxableOk returns a tuple with the Cassa2Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassa2Taxable

`func (o *IssuedDocument) SetCassa2Taxable(v float32)`

SetCassa2Taxable sets Cassa2Taxable field to given value.

### HasCassa2Taxable

`func (o *IssuedDocument) HasCassa2Taxable() bool`

HasCassa2Taxable returns a boolean if a field has been set.

### SetCassa2TaxableNil

`func (o *IssuedDocument) SetCassa2TaxableNil(b bool)`

 SetCassa2TaxableNil sets the value for Cassa2Taxable to be an explicit nil

### UnsetCassa2Taxable
`func (o *IssuedDocument) UnsetCassa2Taxable()`

UnsetCassa2Taxable ensures that no value is present for Cassa2Taxable, not even an explicit nil
### GetAmountCassa2Taxable

`func (o *IssuedDocument) GetAmountCassa2Taxable() float32`

GetAmountCassa2Taxable returns the AmountCassa2Taxable field if non-nil, zero value otherwise.

### GetAmountCassa2TaxableOk

`func (o *IssuedDocument) GetAmountCassa2TaxableOk() (*float32, bool)`

GetAmountCassa2TaxableOk returns a tuple with the AmountCassa2Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCassa2Taxable

`func (o *IssuedDocument) SetAmountCassa2Taxable(v float32)`

SetAmountCassa2Taxable sets AmountCassa2Taxable field to given value.

### HasAmountCassa2Taxable

`func (o *IssuedDocument) HasAmountCassa2Taxable() bool`

HasAmountCassa2Taxable returns a boolean if a field has been set.

### SetAmountCassa2TaxableNil

`func (o *IssuedDocument) SetAmountCassa2TaxableNil(b bool)`

 SetAmountCassa2TaxableNil sets the value for AmountCassa2Taxable to be an explicit nil

### UnsetAmountCassa2Taxable
`func (o *IssuedDocument) UnsetAmountCassa2Taxable()`

UnsetAmountCassa2Taxable ensures that no value is present for AmountCassa2Taxable, not even an explicit nil
### GetGlobalCassaTaxable

`func (o *IssuedDocument) GetGlobalCassaTaxable() float32`

GetGlobalCassaTaxable returns the GlobalCassaTaxable field if non-nil, zero value otherwise.

### GetGlobalCassaTaxableOk

`func (o *IssuedDocument) GetGlobalCassaTaxableOk() (*float32, bool)`

GetGlobalCassaTaxableOk returns a tuple with the GlobalCassaTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCassaTaxable

`func (o *IssuedDocument) SetGlobalCassaTaxable(v float32)`

SetGlobalCassaTaxable sets GlobalCassaTaxable field to given value.

### HasGlobalCassaTaxable

`func (o *IssuedDocument) HasGlobalCassaTaxable() bool`

HasGlobalCassaTaxable returns a boolean if a field has been set.

### SetGlobalCassaTaxableNil

`func (o *IssuedDocument) SetGlobalCassaTaxableNil(b bool)`

 SetGlobalCassaTaxableNil sets the value for GlobalCassaTaxable to be an explicit nil

### UnsetGlobalCassaTaxable
`func (o *IssuedDocument) UnsetGlobalCassaTaxable()`

UnsetGlobalCassaTaxable ensures that no value is present for GlobalCassaTaxable, not even an explicit nil
### GetAmountGlobalCassaTaxable

`func (o *IssuedDocument) GetAmountGlobalCassaTaxable() float32`

GetAmountGlobalCassaTaxable returns the AmountGlobalCassaTaxable field if non-nil, zero value otherwise.

### GetAmountGlobalCassaTaxableOk

`func (o *IssuedDocument) GetAmountGlobalCassaTaxableOk() (*float32, bool)`

GetAmountGlobalCassaTaxableOk returns a tuple with the AmountGlobalCassaTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGlobalCassaTaxable

`func (o *IssuedDocument) SetAmountGlobalCassaTaxable(v float32)`

SetAmountGlobalCassaTaxable sets AmountGlobalCassaTaxable field to given value.

### HasAmountGlobalCassaTaxable

`func (o *IssuedDocument) HasAmountGlobalCassaTaxable() bool`

HasAmountGlobalCassaTaxable returns a boolean if a field has been set.

### SetAmountGlobalCassaTaxableNil

`func (o *IssuedDocument) SetAmountGlobalCassaTaxableNil(b bool)`

 SetAmountGlobalCassaTaxableNil sets the value for AmountGlobalCassaTaxable to be an explicit nil

### UnsetAmountGlobalCassaTaxable
`func (o *IssuedDocument) UnsetAmountGlobalCassaTaxable()`

UnsetAmountGlobalCassaTaxable ensures that no value is present for AmountGlobalCassaTaxable, not even an explicit nil
### GetWithholdingTax

`func (o *IssuedDocument) GetWithholdingTax() float32`

GetWithholdingTax returns the WithholdingTax field if non-nil, zero value otherwise.

### GetWithholdingTaxOk

`func (o *IssuedDocument) GetWithholdingTaxOk() (*float32, bool)`

GetWithholdingTaxOk returns a tuple with the WithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholdingTax

`func (o *IssuedDocument) SetWithholdingTax(v float32)`

SetWithholdingTax sets WithholdingTax field to given value.

### HasWithholdingTax

`func (o *IssuedDocument) HasWithholdingTax() bool`

HasWithholdingTax returns a boolean if a field has been set.

### SetWithholdingTaxNil

`func (o *IssuedDocument) SetWithholdingTaxNil(b bool)`

 SetWithholdingTaxNil sets the value for WithholdingTax to be an explicit nil

### UnsetWithholdingTax
`func (o *IssuedDocument) UnsetWithholdingTax()`

UnsetWithholdingTax ensures that no value is present for WithholdingTax, not even an explicit nil
### GetWithholdingTaxTaxable

`func (o *IssuedDocument) GetWithholdingTaxTaxable() float32`

GetWithholdingTaxTaxable returns the WithholdingTaxTaxable field if non-nil, zero value otherwise.

### GetWithholdingTaxTaxableOk

`func (o *IssuedDocument) GetWithholdingTaxTaxableOk() (*float32, bool)`

GetWithholdingTaxTaxableOk returns a tuple with the WithholdingTaxTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholdingTaxTaxable

`func (o *IssuedDocument) SetWithholdingTaxTaxable(v float32)`

SetWithholdingTaxTaxable sets WithholdingTaxTaxable field to given value.

### HasWithholdingTaxTaxable

`func (o *IssuedDocument) HasWithholdingTaxTaxable() bool`

HasWithholdingTaxTaxable returns a boolean if a field has been set.

### SetWithholdingTaxTaxableNil

`func (o *IssuedDocument) SetWithholdingTaxTaxableNil(b bool)`

 SetWithholdingTaxTaxableNil sets the value for WithholdingTaxTaxable to be an explicit nil

### UnsetWithholdingTaxTaxable
`func (o *IssuedDocument) UnsetWithholdingTaxTaxable()`

UnsetWithholdingTaxTaxable ensures that no value is present for WithholdingTaxTaxable, not even an explicit nil
### GetOtherWithholdingTax

`func (o *IssuedDocument) GetOtherWithholdingTax() float32`

GetOtherWithholdingTax returns the OtherWithholdingTax field if non-nil, zero value otherwise.

### GetOtherWithholdingTaxOk

`func (o *IssuedDocument) GetOtherWithholdingTaxOk() (*float32, bool)`

GetOtherWithholdingTaxOk returns a tuple with the OtherWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherWithholdingTax

`func (o *IssuedDocument) SetOtherWithholdingTax(v float32)`

SetOtherWithholdingTax sets OtherWithholdingTax field to given value.

### HasOtherWithholdingTax

`func (o *IssuedDocument) HasOtherWithholdingTax() bool`

HasOtherWithholdingTax returns a boolean if a field has been set.

### SetOtherWithholdingTaxNil

`func (o *IssuedDocument) SetOtherWithholdingTaxNil(b bool)`

 SetOtherWithholdingTaxNil sets the value for OtherWithholdingTax to be an explicit nil

### UnsetOtherWithholdingTax
`func (o *IssuedDocument) UnsetOtherWithholdingTax()`

UnsetOtherWithholdingTax ensures that no value is present for OtherWithholdingTax, not even an explicit nil
### GetStampDuty

`func (o *IssuedDocument) GetStampDuty() float32`

GetStampDuty returns the StampDuty field if non-nil, zero value otherwise.

### GetStampDutyOk

`func (o *IssuedDocument) GetStampDutyOk() (*float32, bool)`

GetStampDutyOk returns a tuple with the StampDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDuty

`func (o *IssuedDocument) SetStampDuty(v float32)`

SetStampDuty sets StampDuty field to given value.

### HasStampDuty

`func (o *IssuedDocument) HasStampDuty() bool`

HasStampDuty returns a boolean if a field has been set.

### SetStampDutyNil

`func (o *IssuedDocument) SetStampDutyNil(b bool)`

 SetStampDutyNil sets the value for StampDuty to be an explicit nil

### UnsetStampDuty
`func (o *IssuedDocument) UnsetStampDuty()`

UnsetStampDuty ensures that no value is present for StampDuty, not even an explicit nil
### GetPaymentMethod

`func (o *IssuedDocument) GetPaymentMethod() PaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *IssuedDocument) GetPaymentMethodOk() (*PaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *IssuedDocument) SetPaymentMethod(v PaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *IssuedDocument) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetUseSplitPayment

`func (o *IssuedDocument) GetUseSplitPayment() bool`

GetUseSplitPayment returns the UseSplitPayment field if non-nil, zero value otherwise.

### GetUseSplitPaymentOk

`func (o *IssuedDocument) GetUseSplitPaymentOk() (*bool, bool)`

GetUseSplitPaymentOk returns a tuple with the UseSplitPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSplitPayment

`func (o *IssuedDocument) SetUseSplitPayment(v bool)`

SetUseSplitPayment sets UseSplitPayment field to given value.

### HasUseSplitPayment

`func (o *IssuedDocument) HasUseSplitPayment() bool`

HasUseSplitPayment returns a boolean if a field has been set.

### SetUseSplitPaymentNil

`func (o *IssuedDocument) SetUseSplitPaymentNil(b bool)`

 SetUseSplitPaymentNil sets the value for UseSplitPayment to be an explicit nil

### UnsetUseSplitPayment
`func (o *IssuedDocument) UnsetUseSplitPayment()`

UnsetUseSplitPayment ensures that no value is present for UseSplitPayment, not even an explicit nil
### GetUseGrossPrices

`func (o *IssuedDocument) GetUseGrossPrices() bool`

GetUseGrossPrices returns the UseGrossPrices field if non-nil, zero value otherwise.

### GetUseGrossPricesOk

`func (o *IssuedDocument) GetUseGrossPricesOk() (*bool, bool)`

GetUseGrossPricesOk returns a tuple with the UseGrossPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGrossPrices

`func (o *IssuedDocument) SetUseGrossPrices(v bool)`

SetUseGrossPrices sets UseGrossPrices field to given value.

### HasUseGrossPrices

`func (o *IssuedDocument) HasUseGrossPrices() bool`

HasUseGrossPrices returns a boolean if a field has been set.

### SetUseGrossPricesNil

`func (o *IssuedDocument) SetUseGrossPricesNil(b bool)`

 SetUseGrossPricesNil sets the value for UseGrossPrices to be an explicit nil

### UnsetUseGrossPrices
`func (o *IssuedDocument) UnsetUseGrossPrices()`

UnsetUseGrossPrices ensures that no value is present for UseGrossPrices, not even an explicit nil
### GetEInvoice

`func (o *IssuedDocument) GetEInvoice() bool`

GetEInvoice returns the EInvoice field if non-nil, zero value otherwise.

### GetEInvoiceOk

`func (o *IssuedDocument) GetEInvoiceOk() (*bool, bool)`

GetEInvoiceOk returns a tuple with the EInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoice

`func (o *IssuedDocument) SetEInvoice(v bool)`

SetEInvoice sets EInvoice field to given value.

### HasEInvoice

`func (o *IssuedDocument) HasEInvoice() bool`

HasEInvoice returns a boolean if a field has been set.

### SetEInvoiceNil

`func (o *IssuedDocument) SetEInvoiceNil(b bool)`

 SetEInvoiceNil sets the value for EInvoice to be an explicit nil

### UnsetEInvoice
`func (o *IssuedDocument) UnsetEInvoice()`

UnsetEInvoice ensures that no value is present for EInvoice, not even an explicit nil
### GetEiData

`func (o *IssuedDocument) GetEiData() IssuedDocumentEiData`

GetEiData returns the EiData field if non-nil, zero value otherwise.

### GetEiDataOk

`func (o *IssuedDocument) GetEiDataOk() (*IssuedDocumentEiData, bool)`

GetEiDataOk returns a tuple with the EiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiData

`func (o *IssuedDocument) SetEiData(v IssuedDocumentEiData)`

SetEiData sets EiData field to given value.

### HasEiData

`func (o *IssuedDocument) HasEiData() bool`

HasEiData returns a boolean if a field has been set.

### SetEiDataNil

`func (o *IssuedDocument) SetEiDataNil(b bool)`

 SetEiDataNil sets the value for EiData to be an explicit nil

### UnsetEiData
`func (o *IssuedDocument) UnsetEiData()`

UnsetEiData ensures that no value is present for EiData, not even an explicit nil
### GetEiCassaType

`func (o *IssuedDocument) GetEiCassaType() string`

GetEiCassaType returns the EiCassaType field if non-nil, zero value otherwise.

### GetEiCassaTypeOk

`func (o *IssuedDocument) GetEiCassaTypeOk() (*string, bool)`

GetEiCassaTypeOk returns a tuple with the EiCassaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiCassaType

`func (o *IssuedDocument) SetEiCassaType(v string)`

SetEiCassaType sets EiCassaType field to given value.

### HasEiCassaType

`func (o *IssuedDocument) HasEiCassaType() bool`

HasEiCassaType returns a boolean if a field has been set.

### SetEiCassaTypeNil

`func (o *IssuedDocument) SetEiCassaTypeNil(b bool)`

 SetEiCassaTypeNil sets the value for EiCassaType to be an explicit nil

### UnsetEiCassaType
`func (o *IssuedDocument) UnsetEiCassaType()`

UnsetEiCassaType ensures that no value is present for EiCassaType, not even an explicit nil
### GetEiCassa2Type

`func (o *IssuedDocument) GetEiCassa2Type() string`

GetEiCassa2Type returns the EiCassa2Type field if non-nil, zero value otherwise.

### GetEiCassa2TypeOk

`func (o *IssuedDocument) GetEiCassa2TypeOk() (*string, bool)`

GetEiCassa2TypeOk returns a tuple with the EiCassa2Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiCassa2Type

`func (o *IssuedDocument) SetEiCassa2Type(v string)`

SetEiCassa2Type sets EiCassa2Type field to given value.

### HasEiCassa2Type

`func (o *IssuedDocument) HasEiCassa2Type() bool`

HasEiCassa2Type returns a boolean if a field has been set.

### SetEiCassa2TypeNil

`func (o *IssuedDocument) SetEiCassa2TypeNil(b bool)`

 SetEiCassa2TypeNil sets the value for EiCassa2Type to be an explicit nil

### UnsetEiCassa2Type
`func (o *IssuedDocument) UnsetEiCassa2Type()`

UnsetEiCassa2Type ensures that no value is present for EiCassa2Type, not even an explicit nil
### GetEiWithholdingTaxCausal

`func (o *IssuedDocument) GetEiWithholdingTaxCausal() string`

GetEiWithholdingTaxCausal returns the EiWithholdingTaxCausal field if non-nil, zero value otherwise.

### GetEiWithholdingTaxCausalOk

`func (o *IssuedDocument) GetEiWithholdingTaxCausalOk() (*string, bool)`

GetEiWithholdingTaxCausalOk returns a tuple with the EiWithholdingTaxCausal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiWithholdingTaxCausal

`func (o *IssuedDocument) SetEiWithholdingTaxCausal(v string)`

SetEiWithholdingTaxCausal sets EiWithholdingTaxCausal field to given value.

### HasEiWithholdingTaxCausal

`func (o *IssuedDocument) HasEiWithholdingTaxCausal() bool`

HasEiWithholdingTaxCausal returns a boolean if a field has been set.

### SetEiWithholdingTaxCausalNil

`func (o *IssuedDocument) SetEiWithholdingTaxCausalNil(b bool)`

 SetEiWithholdingTaxCausalNil sets the value for EiWithholdingTaxCausal to be an explicit nil

### UnsetEiWithholdingTaxCausal
`func (o *IssuedDocument) UnsetEiWithholdingTaxCausal()`

UnsetEiWithholdingTaxCausal ensures that no value is present for EiWithholdingTaxCausal, not even an explicit nil
### GetEiOtherWithholdingTaxType

`func (o *IssuedDocument) GetEiOtherWithholdingTaxType() string`

GetEiOtherWithholdingTaxType returns the EiOtherWithholdingTaxType field if non-nil, zero value otherwise.

### GetEiOtherWithholdingTaxTypeOk

`func (o *IssuedDocument) GetEiOtherWithholdingTaxTypeOk() (*string, bool)`

GetEiOtherWithholdingTaxTypeOk returns a tuple with the EiOtherWithholdingTaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiOtherWithholdingTaxType

`func (o *IssuedDocument) SetEiOtherWithholdingTaxType(v string)`

SetEiOtherWithholdingTaxType sets EiOtherWithholdingTaxType field to given value.

### HasEiOtherWithholdingTaxType

`func (o *IssuedDocument) HasEiOtherWithholdingTaxType() bool`

HasEiOtherWithholdingTaxType returns a boolean if a field has been set.

### SetEiOtherWithholdingTaxTypeNil

`func (o *IssuedDocument) SetEiOtherWithholdingTaxTypeNil(b bool)`

 SetEiOtherWithholdingTaxTypeNil sets the value for EiOtherWithholdingTaxType to be an explicit nil

### UnsetEiOtherWithholdingTaxType
`func (o *IssuedDocument) UnsetEiOtherWithholdingTaxType()`

UnsetEiOtherWithholdingTaxType ensures that no value is present for EiOtherWithholdingTaxType, not even an explicit nil
### GetEiOtherWithholdingTaxCausal

`func (o *IssuedDocument) GetEiOtherWithholdingTaxCausal() string`

GetEiOtherWithholdingTaxCausal returns the EiOtherWithholdingTaxCausal field if non-nil, zero value otherwise.

### GetEiOtherWithholdingTaxCausalOk

`func (o *IssuedDocument) GetEiOtherWithholdingTaxCausalOk() (*string, bool)`

GetEiOtherWithholdingTaxCausalOk returns a tuple with the EiOtherWithholdingTaxCausal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiOtherWithholdingTaxCausal

`func (o *IssuedDocument) SetEiOtherWithholdingTaxCausal(v string)`

SetEiOtherWithholdingTaxCausal sets EiOtherWithholdingTaxCausal field to given value.

### HasEiOtherWithholdingTaxCausal

`func (o *IssuedDocument) HasEiOtherWithholdingTaxCausal() bool`

HasEiOtherWithholdingTaxCausal returns a boolean if a field has been set.

### SetEiOtherWithholdingTaxCausalNil

`func (o *IssuedDocument) SetEiOtherWithholdingTaxCausalNil(b bool)`

 SetEiOtherWithholdingTaxCausalNil sets the value for EiOtherWithholdingTaxCausal to be an explicit nil

### UnsetEiOtherWithholdingTaxCausal
`func (o *IssuedDocument) UnsetEiOtherWithholdingTaxCausal()`

UnsetEiOtherWithholdingTaxCausal ensures that no value is present for EiOtherWithholdingTaxCausal, not even an explicit nil
### GetItemsList

`func (o *IssuedDocument) GetItemsList() []IssuedDocumentItemsListItem`

GetItemsList returns the ItemsList field if non-nil, zero value otherwise.

### GetItemsListOk

`func (o *IssuedDocument) GetItemsListOk() (*[]IssuedDocumentItemsListItem, bool)`

GetItemsListOk returns a tuple with the ItemsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsList

`func (o *IssuedDocument) SetItemsList(v []IssuedDocumentItemsListItem)`

SetItemsList sets ItemsList field to given value.

### HasItemsList

`func (o *IssuedDocument) HasItemsList() bool`

HasItemsList returns a boolean if a field has been set.

### SetItemsListNil

`func (o *IssuedDocument) SetItemsListNil(b bool)`

 SetItemsListNil sets the value for ItemsList to be an explicit nil

### UnsetItemsList
`func (o *IssuedDocument) UnsetItemsList()`

UnsetItemsList ensures that no value is present for ItemsList, not even an explicit nil
### GetPaymentsList

`func (o *IssuedDocument) GetPaymentsList() []IssuedDocumentPaymentsListItem`

GetPaymentsList returns the PaymentsList field if non-nil, zero value otherwise.

### GetPaymentsListOk

`func (o *IssuedDocument) GetPaymentsListOk() (*[]IssuedDocumentPaymentsListItem, bool)`

GetPaymentsListOk returns a tuple with the PaymentsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsList

`func (o *IssuedDocument) SetPaymentsList(v []IssuedDocumentPaymentsListItem)`

SetPaymentsList sets PaymentsList field to given value.

### HasPaymentsList

`func (o *IssuedDocument) HasPaymentsList() bool`

HasPaymentsList returns a boolean if a field has been set.

### SetPaymentsListNil

`func (o *IssuedDocument) SetPaymentsListNil(b bool)`

 SetPaymentsListNil sets the value for PaymentsList to be an explicit nil

### UnsetPaymentsList
`func (o *IssuedDocument) UnsetPaymentsList()`

UnsetPaymentsList ensures that no value is present for PaymentsList, not even an explicit nil
### GetTemplate

`func (o *IssuedDocument) GetTemplate() DocumentTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IssuedDocument) GetTemplateOk() (*DocumentTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IssuedDocument) SetTemplate(v DocumentTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IssuedDocument) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDeliveryNoteTemplate

`func (o *IssuedDocument) GetDeliveryNoteTemplate() DocumentTemplate`

GetDeliveryNoteTemplate returns the DeliveryNoteTemplate field if non-nil, zero value otherwise.

### GetDeliveryNoteTemplateOk

`func (o *IssuedDocument) GetDeliveryNoteTemplateOk() (*DocumentTemplate, bool)`

GetDeliveryNoteTemplateOk returns a tuple with the DeliveryNoteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNoteTemplate

`func (o *IssuedDocument) SetDeliveryNoteTemplate(v DocumentTemplate)`

SetDeliveryNoteTemplate sets DeliveryNoteTemplate field to given value.

### HasDeliveryNoteTemplate

`func (o *IssuedDocument) HasDeliveryNoteTemplate() bool`

HasDeliveryNoteTemplate returns a boolean if a field has been set.

### GetAccInvTemplate

`func (o *IssuedDocument) GetAccInvTemplate() DocumentTemplate`

GetAccInvTemplate returns the AccInvTemplate field if non-nil, zero value otherwise.

### GetAccInvTemplateOk

`func (o *IssuedDocument) GetAccInvTemplateOk() (*DocumentTemplate, bool)`

GetAccInvTemplateOk returns a tuple with the AccInvTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccInvTemplate

`func (o *IssuedDocument) SetAccInvTemplate(v DocumentTemplate)`

SetAccInvTemplate sets AccInvTemplate field to given value.

### HasAccInvTemplate

`func (o *IssuedDocument) HasAccInvTemplate() bool`

HasAccInvTemplate returns a boolean if a field has been set.

### GetHMargins

`func (o *IssuedDocument) GetHMargins() int32`

GetHMargins returns the HMargins field if non-nil, zero value otherwise.

### GetHMarginsOk

`func (o *IssuedDocument) GetHMarginsOk() (*int32, bool)`

GetHMarginsOk returns a tuple with the HMargins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHMargins

`func (o *IssuedDocument) SetHMargins(v int32)`

SetHMargins sets HMargins field to given value.

### HasHMargins

`func (o *IssuedDocument) HasHMargins() bool`

HasHMargins returns a boolean if a field has been set.

### SetHMarginsNil

`func (o *IssuedDocument) SetHMarginsNil(b bool)`

 SetHMarginsNil sets the value for HMargins to be an explicit nil

### UnsetHMargins
`func (o *IssuedDocument) UnsetHMargins()`

UnsetHMargins ensures that no value is present for HMargins, not even an explicit nil
### GetVMargins

`func (o *IssuedDocument) GetVMargins() int32`

GetVMargins returns the VMargins field if non-nil, zero value otherwise.

### GetVMarginsOk

`func (o *IssuedDocument) GetVMarginsOk() (*int32, bool)`

GetVMarginsOk returns a tuple with the VMargins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMargins

`func (o *IssuedDocument) SetVMargins(v int32)`

SetVMargins sets VMargins field to given value.

### HasVMargins

`func (o *IssuedDocument) HasVMargins() bool`

HasVMargins returns a boolean if a field has been set.

### SetVMarginsNil

`func (o *IssuedDocument) SetVMarginsNil(b bool)`

 SetVMarginsNil sets the value for VMargins to be an explicit nil

### UnsetVMargins
`func (o *IssuedDocument) UnsetVMargins()`

UnsetVMargins ensures that no value is present for VMargins, not even an explicit nil
### GetShowPayments

`func (o *IssuedDocument) GetShowPayments() bool`

GetShowPayments returns the ShowPayments field if non-nil, zero value otherwise.

### GetShowPaymentsOk

`func (o *IssuedDocument) GetShowPaymentsOk() (*bool, bool)`

GetShowPaymentsOk returns a tuple with the ShowPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPayments

`func (o *IssuedDocument) SetShowPayments(v bool)`

SetShowPayments sets ShowPayments field to given value.

### HasShowPayments

`func (o *IssuedDocument) HasShowPayments() bool`

HasShowPayments returns a boolean if a field has been set.

### SetShowPaymentsNil

`func (o *IssuedDocument) SetShowPaymentsNil(b bool)`

 SetShowPaymentsNil sets the value for ShowPayments to be an explicit nil

### UnsetShowPayments
`func (o *IssuedDocument) UnsetShowPayments()`

UnsetShowPayments ensures that no value is present for ShowPayments, not even an explicit nil
### GetShowPaymentMethod

`func (o *IssuedDocument) GetShowPaymentMethod() bool`

GetShowPaymentMethod returns the ShowPaymentMethod field if non-nil, zero value otherwise.

### GetShowPaymentMethodOk

`func (o *IssuedDocument) GetShowPaymentMethodOk() (*bool, bool)`

GetShowPaymentMethodOk returns a tuple with the ShowPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPaymentMethod

`func (o *IssuedDocument) SetShowPaymentMethod(v bool)`

SetShowPaymentMethod sets ShowPaymentMethod field to given value.

### HasShowPaymentMethod

`func (o *IssuedDocument) HasShowPaymentMethod() bool`

HasShowPaymentMethod returns a boolean if a field has been set.

### SetShowPaymentMethodNil

`func (o *IssuedDocument) SetShowPaymentMethodNil(b bool)`

 SetShowPaymentMethodNil sets the value for ShowPaymentMethod to be an explicit nil

### UnsetShowPaymentMethod
`func (o *IssuedDocument) UnsetShowPaymentMethod()`

UnsetShowPaymentMethod ensures that no value is present for ShowPaymentMethod, not even an explicit nil
### GetShowTotals

`func (o *IssuedDocument) GetShowTotals() ShowTotalsMode`

GetShowTotals returns the ShowTotals field if non-nil, zero value otherwise.

### GetShowTotalsOk

`func (o *IssuedDocument) GetShowTotalsOk() (*ShowTotalsMode, bool)`

GetShowTotalsOk returns a tuple with the ShowTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTotals

`func (o *IssuedDocument) SetShowTotals(v ShowTotalsMode)`

SetShowTotals sets ShowTotals field to given value.

### HasShowTotals

`func (o *IssuedDocument) HasShowTotals() bool`

HasShowTotals returns a boolean if a field has been set.

### GetShowNotificationButton

`func (o *IssuedDocument) GetShowNotificationButton() bool`

GetShowNotificationButton returns the ShowNotificationButton field if non-nil, zero value otherwise.

### GetShowNotificationButtonOk

`func (o *IssuedDocument) GetShowNotificationButtonOk() (*bool, bool)`

GetShowNotificationButtonOk returns a tuple with the ShowNotificationButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNotificationButton

`func (o *IssuedDocument) SetShowNotificationButton(v bool)`

SetShowNotificationButton sets ShowNotificationButton field to given value.

### HasShowNotificationButton

`func (o *IssuedDocument) HasShowNotificationButton() bool`

HasShowNotificationButton returns a boolean if a field has been set.

### SetShowNotificationButtonNil

`func (o *IssuedDocument) SetShowNotificationButtonNil(b bool)`

 SetShowNotificationButtonNil sets the value for ShowNotificationButton to be an explicit nil

### UnsetShowNotificationButton
`func (o *IssuedDocument) UnsetShowNotificationButton()`

UnsetShowNotificationButton ensures that no value is present for ShowNotificationButton, not even an explicit nil
### GetShowTspayButton

`func (o *IssuedDocument) GetShowTspayButton() bool`

GetShowTspayButton returns the ShowTspayButton field if non-nil, zero value otherwise.

### GetShowTspayButtonOk

`func (o *IssuedDocument) GetShowTspayButtonOk() (*bool, bool)`

GetShowTspayButtonOk returns a tuple with the ShowTspayButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTspayButton

`func (o *IssuedDocument) SetShowTspayButton(v bool)`

SetShowTspayButton sets ShowTspayButton field to given value.

### HasShowTspayButton

`func (o *IssuedDocument) HasShowTspayButton() bool`

HasShowTspayButton returns a boolean if a field has been set.

### SetShowTspayButtonNil

`func (o *IssuedDocument) SetShowTspayButtonNil(b bool)`

 SetShowTspayButtonNil sets the value for ShowTspayButton to be an explicit nil

### UnsetShowTspayButton
`func (o *IssuedDocument) UnsetShowTspayButton()`

UnsetShowTspayButton ensures that no value is present for ShowTspayButton, not even an explicit nil
### GetDeliveryNote

`func (o *IssuedDocument) GetDeliveryNote() bool`

GetDeliveryNote returns the DeliveryNote field if non-nil, zero value otherwise.

### GetDeliveryNoteOk

`func (o *IssuedDocument) GetDeliveryNoteOk() (*bool, bool)`

GetDeliveryNoteOk returns a tuple with the DeliveryNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNote

`func (o *IssuedDocument) SetDeliveryNote(v bool)`

SetDeliveryNote sets DeliveryNote field to given value.

### HasDeliveryNote

`func (o *IssuedDocument) HasDeliveryNote() bool`

HasDeliveryNote returns a boolean if a field has been set.

### SetDeliveryNoteNil

`func (o *IssuedDocument) SetDeliveryNoteNil(b bool)`

 SetDeliveryNoteNil sets the value for DeliveryNote to be an explicit nil

### UnsetDeliveryNote
`func (o *IssuedDocument) UnsetDeliveryNote()`

UnsetDeliveryNote ensures that no value is present for DeliveryNote, not even an explicit nil
### GetAccompanyingInvoice

`func (o *IssuedDocument) GetAccompanyingInvoice() bool`

GetAccompanyingInvoice returns the AccompanyingInvoice field if non-nil, zero value otherwise.

### GetAccompanyingInvoiceOk

`func (o *IssuedDocument) GetAccompanyingInvoiceOk() (*bool, bool)`

GetAccompanyingInvoiceOk returns a tuple with the AccompanyingInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccompanyingInvoice

`func (o *IssuedDocument) SetAccompanyingInvoice(v bool)`

SetAccompanyingInvoice sets AccompanyingInvoice field to given value.

### HasAccompanyingInvoice

`func (o *IssuedDocument) HasAccompanyingInvoice() bool`

HasAccompanyingInvoice returns a boolean if a field has been set.

### SetAccompanyingInvoiceNil

`func (o *IssuedDocument) SetAccompanyingInvoiceNil(b bool)`

 SetAccompanyingInvoiceNil sets the value for AccompanyingInvoice to be an explicit nil

### UnsetAccompanyingInvoice
`func (o *IssuedDocument) UnsetAccompanyingInvoice()`

UnsetAccompanyingInvoice ensures that no value is present for AccompanyingInvoice, not even an explicit nil
### GetDnNumber

`func (o *IssuedDocument) GetDnNumber() int32`

GetDnNumber returns the DnNumber field if non-nil, zero value otherwise.

### GetDnNumberOk

`func (o *IssuedDocument) GetDnNumberOk() (*int32, bool)`

GetDnNumberOk returns a tuple with the DnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnNumber

`func (o *IssuedDocument) SetDnNumber(v int32)`

SetDnNumber sets DnNumber field to given value.

### HasDnNumber

`func (o *IssuedDocument) HasDnNumber() bool`

HasDnNumber returns a boolean if a field has been set.

### SetDnNumberNil

`func (o *IssuedDocument) SetDnNumberNil(b bool)`

 SetDnNumberNil sets the value for DnNumber to be an explicit nil

### UnsetDnNumber
`func (o *IssuedDocument) UnsetDnNumber()`

UnsetDnNumber ensures that no value is present for DnNumber, not even an explicit nil
### GetDnDate

`func (o *IssuedDocument) GetDnDate() string`

GetDnDate returns the DnDate field if non-nil, zero value otherwise.

### GetDnDateOk

`func (o *IssuedDocument) GetDnDateOk() (*string, bool)`

GetDnDateOk returns a tuple with the DnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnDate

`func (o *IssuedDocument) SetDnDate(v string)`

SetDnDate sets DnDate field to given value.

### HasDnDate

`func (o *IssuedDocument) HasDnDate() bool`

HasDnDate returns a boolean if a field has been set.

### SetDnDateNil

`func (o *IssuedDocument) SetDnDateNil(b bool)`

 SetDnDateNil sets the value for DnDate to be an explicit nil

### UnsetDnDate
`func (o *IssuedDocument) UnsetDnDate()`

UnsetDnDate ensures that no value is present for DnDate, not even an explicit nil
### GetDnAiPackagesNumber

`func (o *IssuedDocument) GetDnAiPackagesNumber() string`

GetDnAiPackagesNumber returns the DnAiPackagesNumber field if non-nil, zero value otherwise.

### GetDnAiPackagesNumberOk

`func (o *IssuedDocument) GetDnAiPackagesNumberOk() (*string, bool)`

GetDnAiPackagesNumberOk returns a tuple with the DnAiPackagesNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAiPackagesNumber

`func (o *IssuedDocument) SetDnAiPackagesNumber(v string)`

SetDnAiPackagesNumber sets DnAiPackagesNumber field to given value.

### HasDnAiPackagesNumber

`func (o *IssuedDocument) HasDnAiPackagesNumber() bool`

HasDnAiPackagesNumber returns a boolean if a field has been set.

### SetDnAiPackagesNumberNil

`func (o *IssuedDocument) SetDnAiPackagesNumberNil(b bool)`

 SetDnAiPackagesNumberNil sets the value for DnAiPackagesNumber to be an explicit nil

### UnsetDnAiPackagesNumber
`func (o *IssuedDocument) UnsetDnAiPackagesNumber()`

UnsetDnAiPackagesNumber ensures that no value is present for DnAiPackagesNumber, not even an explicit nil
### GetDnAiWeight

`func (o *IssuedDocument) GetDnAiWeight() string`

GetDnAiWeight returns the DnAiWeight field if non-nil, zero value otherwise.

### GetDnAiWeightOk

`func (o *IssuedDocument) GetDnAiWeightOk() (*string, bool)`

GetDnAiWeightOk returns a tuple with the DnAiWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAiWeight

`func (o *IssuedDocument) SetDnAiWeight(v string)`

SetDnAiWeight sets DnAiWeight field to given value.

### HasDnAiWeight

`func (o *IssuedDocument) HasDnAiWeight() bool`

HasDnAiWeight returns a boolean if a field has been set.

### SetDnAiWeightNil

`func (o *IssuedDocument) SetDnAiWeightNil(b bool)`

 SetDnAiWeightNil sets the value for DnAiWeight to be an explicit nil

### UnsetDnAiWeight
`func (o *IssuedDocument) UnsetDnAiWeight()`

UnsetDnAiWeight ensures that no value is present for DnAiWeight, not even an explicit nil
### GetDnAiCausal

`func (o *IssuedDocument) GetDnAiCausal() string`

GetDnAiCausal returns the DnAiCausal field if non-nil, zero value otherwise.

### GetDnAiCausalOk

`func (o *IssuedDocument) GetDnAiCausalOk() (*string, bool)`

GetDnAiCausalOk returns a tuple with the DnAiCausal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAiCausal

`func (o *IssuedDocument) SetDnAiCausal(v string)`

SetDnAiCausal sets DnAiCausal field to given value.

### HasDnAiCausal

`func (o *IssuedDocument) HasDnAiCausal() bool`

HasDnAiCausal returns a boolean if a field has been set.

### SetDnAiCausalNil

`func (o *IssuedDocument) SetDnAiCausalNil(b bool)`

 SetDnAiCausalNil sets the value for DnAiCausal to be an explicit nil

### UnsetDnAiCausal
`func (o *IssuedDocument) UnsetDnAiCausal()`

UnsetDnAiCausal ensures that no value is present for DnAiCausal, not even an explicit nil
### GetDnAiDestination

`func (o *IssuedDocument) GetDnAiDestination() string`

GetDnAiDestination returns the DnAiDestination field if non-nil, zero value otherwise.

### GetDnAiDestinationOk

`func (o *IssuedDocument) GetDnAiDestinationOk() (*string, bool)`

GetDnAiDestinationOk returns a tuple with the DnAiDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAiDestination

`func (o *IssuedDocument) SetDnAiDestination(v string)`

SetDnAiDestination sets DnAiDestination field to given value.

### HasDnAiDestination

`func (o *IssuedDocument) HasDnAiDestination() bool`

HasDnAiDestination returns a boolean if a field has been set.

### SetDnAiDestinationNil

`func (o *IssuedDocument) SetDnAiDestinationNil(b bool)`

 SetDnAiDestinationNil sets the value for DnAiDestination to be an explicit nil

### UnsetDnAiDestination
`func (o *IssuedDocument) UnsetDnAiDestination()`

UnsetDnAiDestination ensures that no value is present for DnAiDestination, not even an explicit nil
### GetDnAiTransporter

`func (o *IssuedDocument) GetDnAiTransporter() string`

GetDnAiTransporter returns the DnAiTransporter field if non-nil, zero value otherwise.

### GetDnAiTransporterOk

`func (o *IssuedDocument) GetDnAiTransporterOk() (*string, bool)`

GetDnAiTransporterOk returns a tuple with the DnAiTransporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAiTransporter

`func (o *IssuedDocument) SetDnAiTransporter(v string)`

SetDnAiTransporter sets DnAiTransporter field to given value.

### HasDnAiTransporter

`func (o *IssuedDocument) HasDnAiTransporter() bool`

HasDnAiTransporter returns a boolean if a field has been set.

### SetDnAiTransporterNil

`func (o *IssuedDocument) SetDnAiTransporterNil(b bool)`

 SetDnAiTransporterNil sets the value for DnAiTransporter to be an explicit nil

### UnsetDnAiTransporter
`func (o *IssuedDocument) UnsetDnAiTransporter()`

UnsetDnAiTransporter ensures that no value is present for DnAiTransporter, not even an explicit nil
### GetDnAiNotes

`func (o *IssuedDocument) GetDnAiNotes() string`

GetDnAiNotes returns the DnAiNotes field if non-nil, zero value otherwise.

### GetDnAiNotesOk

`func (o *IssuedDocument) GetDnAiNotesOk() (*string, bool)`

GetDnAiNotesOk returns a tuple with the DnAiNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAiNotes

`func (o *IssuedDocument) SetDnAiNotes(v string)`

SetDnAiNotes sets DnAiNotes field to given value.

### HasDnAiNotes

`func (o *IssuedDocument) HasDnAiNotes() bool`

HasDnAiNotes returns a boolean if a field has been set.

### SetDnAiNotesNil

`func (o *IssuedDocument) SetDnAiNotesNil(b bool)`

 SetDnAiNotesNil sets the value for DnAiNotes to be an explicit nil

### UnsetDnAiNotes
`func (o *IssuedDocument) UnsetDnAiNotes()`

UnsetDnAiNotes ensures that no value is present for DnAiNotes, not even an explicit nil
### GetIsMarked

`func (o *IssuedDocument) GetIsMarked() bool`

GetIsMarked returns the IsMarked field if non-nil, zero value otherwise.

### GetIsMarkedOk

`func (o *IssuedDocument) GetIsMarkedOk() (*bool, bool)`

GetIsMarkedOk returns a tuple with the IsMarked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarked

`func (o *IssuedDocument) SetIsMarked(v bool)`

SetIsMarked sets IsMarked field to given value.

### HasIsMarked

`func (o *IssuedDocument) HasIsMarked() bool`

HasIsMarked returns a boolean if a field has been set.

### SetIsMarkedNil

`func (o *IssuedDocument) SetIsMarkedNil(b bool)`

 SetIsMarkedNil sets the value for IsMarked to be an explicit nil

### UnsetIsMarked
`func (o *IssuedDocument) UnsetIsMarked()`

UnsetIsMarked ensures that no value is present for IsMarked, not even an explicit nil
### GetAmountNet

`func (o *IssuedDocument) GetAmountNet() float32`

GetAmountNet returns the AmountNet field if non-nil, zero value otherwise.

### GetAmountNetOk

`func (o *IssuedDocument) GetAmountNetOk() (*float32, bool)`

GetAmountNetOk returns a tuple with the AmountNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountNet

`func (o *IssuedDocument) SetAmountNet(v float32)`

SetAmountNet sets AmountNet field to given value.

### HasAmountNet

`func (o *IssuedDocument) HasAmountNet() bool`

HasAmountNet returns a boolean if a field has been set.

### SetAmountNetNil

`func (o *IssuedDocument) SetAmountNetNil(b bool)`

 SetAmountNetNil sets the value for AmountNet to be an explicit nil

### UnsetAmountNet
`func (o *IssuedDocument) UnsetAmountNet()`

UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
### GetAmountVat

`func (o *IssuedDocument) GetAmountVat() float32`

GetAmountVat returns the AmountVat field if non-nil, zero value otherwise.

### GetAmountVatOk

`func (o *IssuedDocument) GetAmountVatOk() (*float32, bool)`

GetAmountVatOk returns a tuple with the AmountVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountVat

`func (o *IssuedDocument) SetAmountVat(v float32)`

SetAmountVat sets AmountVat field to given value.

### HasAmountVat

`func (o *IssuedDocument) HasAmountVat() bool`

HasAmountVat returns a boolean if a field has been set.

### SetAmountVatNil

`func (o *IssuedDocument) SetAmountVatNil(b bool)`

 SetAmountVatNil sets the value for AmountVat to be an explicit nil

### UnsetAmountVat
`func (o *IssuedDocument) UnsetAmountVat()`

UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
### GetAmountGross

`func (o *IssuedDocument) GetAmountGross() float32`

GetAmountGross returns the AmountGross field if non-nil, zero value otherwise.

### GetAmountGrossOk

`func (o *IssuedDocument) GetAmountGrossOk() (*float32, bool)`

GetAmountGrossOk returns a tuple with the AmountGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGross

`func (o *IssuedDocument) SetAmountGross(v float32)`

SetAmountGross sets AmountGross field to given value.

### HasAmountGross

`func (o *IssuedDocument) HasAmountGross() bool`

HasAmountGross returns a boolean if a field has been set.

### SetAmountGrossNil

`func (o *IssuedDocument) SetAmountGrossNil(b bool)`

 SetAmountGrossNil sets the value for AmountGross to be an explicit nil

### UnsetAmountGross
`func (o *IssuedDocument) UnsetAmountGross()`

UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
### GetAmountDueDiscount

`func (o *IssuedDocument) GetAmountDueDiscount() float32`

GetAmountDueDiscount returns the AmountDueDiscount field if non-nil, zero value otherwise.

### GetAmountDueDiscountOk

`func (o *IssuedDocument) GetAmountDueDiscountOk() (*float32, bool)`

GetAmountDueDiscountOk returns a tuple with the AmountDueDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDueDiscount

`func (o *IssuedDocument) SetAmountDueDiscount(v float32)`

SetAmountDueDiscount sets AmountDueDiscount field to given value.

### HasAmountDueDiscount

`func (o *IssuedDocument) HasAmountDueDiscount() bool`

HasAmountDueDiscount returns a boolean if a field has been set.

### SetAmountDueDiscountNil

`func (o *IssuedDocument) SetAmountDueDiscountNil(b bool)`

 SetAmountDueDiscountNil sets the value for AmountDueDiscount to be an explicit nil

### UnsetAmountDueDiscount
`func (o *IssuedDocument) UnsetAmountDueDiscount()`

UnsetAmountDueDiscount ensures that no value is present for AmountDueDiscount, not even an explicit nil
### GetAmountRivalsa

`func (o *IssuedDocument) GetAmountRivalsa() float32`

GetAmountRivalsa returns the AmountRivalsa field if non-nil, zero value otherwise.

### GetAmountRivalsaOk

`func (o *IssuedDocument) GetAmountRivalsaOk() (*float32, bool)`

GetAmountRivalsaOk returns a tuple with the AmountRivalsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRivalsa

`func (o *IssuedDocument) SetAmountRivalsa(v float32)`

SetAmountRivalsa sets AmountRivalsa field to given value.

### HasAmountRivalsa

`func (o *IssuedDocument) HasAmountRivalsa() bool`

HasAmountRivalsa returns a boolean if a field has been set.

### SetAmountRivalsaNil

`func (o *IssuedDocument) SetAmountRivalsaNil(b bool)`

 SetAmountRivalsaNil sets the value for AmountRivalsa to be an explicit nil

### UnsetAmountRivalsa
`func (o *IssuedDocument) UnsetAmountRivalsa()`

UnsetAmountRivalsa ensures that no value is present for AmountRivalsa, not even an explicit nil
### GetAmountRivalsaTaxable

`func (o *IssuedDocument) GetAmountRivalsaTaxable() float32`

GetAmountRivalsaTaxable returns the AmountRivalsaTaxable field if non-nil, zero value otherwise.

### GetAmountRivalsaTaxableOk

`func (o *IssuedDocument) GetAmountRivalsaTaxableOk() (*float32, bool)`

GetAmountRivalsaTaxableOk returns a tuple with the AmountRivalsaTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRivalsaTaxable

`func (o *IssuedDocument) SetAmountRivalsaTaxable(v float32)`

SetAmountRivalsaTaxable sets AmountRivalsaTaxable field to given value.

### HasAmountRivalsaTaxable

`func (o *IssuedDocument) HasAmountRivalsaTaxable() bool`

HasAmountRivalsaTaxable returns a boolean if a field has been set.

### SetAmountRivalsaTaxableNil

`func (o *IssuedDocument) SetAmountRivalsaTaxableNil(b bool)`

 SetAmountRivalsaTaxableNil sets the value for AmountRivalsaTaxable to be an explicit nil

### UnsetAmountRivalsaTaxable
`func (o *IssuedDocument) UnsetAmountRivalsaTaxable()`

UnsetAmountRivalsaTaxable ensures that no value is present for AmountRivalsaTaxable, not even an explicit nil
### GetAmountWithholdingTax

`func (o *IssuedDocument) GetAmountWithholdingTax() float32`

GetAmountWithholdingTax returns the AmountWithholdingTax field if non-nil, zero value otherwise.

### GetAmountWithholdingTaxOk

`func (o *IssuedDocument) GetAmountWithholdingTaxOk() (*float32, bool)`

GetAmountWithholdingTaxOk returns a tuple with the AmountWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountWithholdingTax

`func (o *IssuedDocument) SetAmountWithholdingTax(v float32)`

SetAmountWithholdingTax sets AmountWithholdingTax field to given value.

### HasAmountWithholdingTax

`func (o *IssuedDocument) HasAmountWithholdingTax() bool`

HasAmountWithholdingTax returns a boolean if a field has been set.

### SetAmountWithholdingTaxNil

`func (o *IssuedDocument) SetAmountWithholdingTaxNil(b bool)`

 SetAmountWithholdingTaxNil sets the value for AmountWithholdingTax to be an explicit nil

### UnsetAmountWithholdingTax
`func (o *IssuedDocument) UnsetAmountWithholdingTax()`

UnsetAmountWithholdingTax ensures that no value is present for AmountWithholdingTax, not even an explicit nil
### GetAmountWithholdingTaxTaxable

`func (o *IssuedDocument) GetAmountWithholdingTaxTaxable() float32`

GetAmountWithholdingTaxTaxable returns the AmountWithholdingTaxTaxable field if non-nil, zero value otherwise.

### GetAmountWithholdingTaxTaxableOk

`func (o *IssuedDocument) GetAmountWithholdingTaxTaxableOk() (*float32, bool)`

GetAmountWithholdingTaxTaxableOk returns a tuple with the AmountWithholdingTaxTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountWithholdingTaxTaxable

`func (o *IssuedDocument) SetAmountWithholdingTaxTaxable(v float32)`

SetAmountWithholdingTaxTaxable sets AmountWithholdingTaxTaxable field to given value.

### HasAmountWithholdingTaxTaxable

`func (o *IssuedDocument) HasAmountWithholdingTaxTaxable() bool`

HasAmountWithholdingTaxTaxable returns a boolean if a field has been set.

### SetAmountWithholdingTaxTaxableNil

`func (o *IssuedDocument) SetAmountWithholdingTaxTaxableNil(b bool)`

 SetAmountWithholdingTaxTaxableNil sets the value for AmountWithholdingTaxTaxable to be an explicit nil

### UnsetAmountWithholdingTaxTaxable
`func (o *IssuedDocument) UnsetAmountWithholdingTaxTaxable()`

UnsetAmountWithholdingTaxTaxable ensures that no value is present for AmountWithholdingTaxTaxable, not even an explicit nil
### GetAmountOtherWithholdingTax

`func (o *IssuedDocument) GetAmountOtherWithholdingTax() float32`

GetAmountOtherWithholdingTax returns the AmountOtherWithholdingTax field if non-nil, zero value otherwise.

### GetAmountOtherWithholdingTaxOk

`func (o *IssuedDocument) GetAmountOtherWithholdingTaxOk() (*float32, bool)`

GetAmountOtherWithholdingTaxOk returns a tuple with the AmountOtherWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOtherWithholdingTax

`func (o *IssuedDocument) SetAmountOtherWithholdingTax(v float32)`

SetAmountOtherWithholdingTax sets AmountOtherWithholdingTax field to given value.

### HasAmountOtherWithholdingTax

`func (o *IssuedDocument) HasAmountOtherWithholdingTax() bool`

HasAmountOtherWithholdingTax returns a boolean if a field has been set.

### SetAmountOtherWithholdingTaxNil

`func (o *IssuedDocument) SetAmountOtherWithholdingTaxNil(b bool)`

 SetAmountOtherWithholdingTaxNil sets the value for AmountOtherWithholdingTax to be an explicit nil

### UnsetAmountOtherWithholdingTax
`func (o *IssuedDocument) UnsetAmountOtherWithholdingTax()`

UnsetAmountOtherWithholdingTax ensures that no value is present for AmountOtherWithholdingTax, not even an explicit nil
### GetAmountOtherWithholdingTaxTaxable

`func (o *IssuedDocument) GetAmountOtherWithholdingTaxTaxable() float32`

GetAmountOtherWithholdingTaxTaxable returns the AmountOtherWithholdingTaxTaxable field if non-nil, zero value otherwise.

### GetAmountOtherWithholdingTaxTaxableOk

`func (o *IssuedDocument) GetAmountOtherWithholdingTaxTaxableOk() (*float32, bool)`

GetAmountOtherWithholdingTaxTaxableOk returns a tuple with the AmountOtherWithholdingTaxTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOtherWithholdingTaxTaxable

`func (o *IssuedDocument) SetAmountOtherWithholdingTaxTaxable(v float32)`

SetAmountOtherWithholdingTaxTaxable sets AmountOtherWithholdingTaxTaxable field to given value.

### HasAmountOtherWithholdingTaxTaxable

`func (o *IssuedDocument) HasAmountOtherWithholdingTaxTaxable() bool`

HasAmountOtherWithholdingTaxTaxable returns a boolean if a field has been set.

### SetAmountOtherWithholdingTaxTaxableNil

`func (o *IssuedDocument) SetAmountOtherWithholdingTaxTaxableNil(b bool)`

 SetAmountOtherWithholdingTaxTaxableNil sets the value for AmountOtherWithholdingTaxTaxable to be an explicit nil

### UnsetAmountOtherWithholdingTaxTaxable
`func (o *IssuedDocument) UnsetAmountOtherWithholdingTaxTaxable()`

UnsetAmountOtherWithholdingTaxTaxable ensures that no value is present for AmountOtherWithholdingTaxTaxable, not even an explicit nil
### GetAmountEnasarcoTaxable

`func (o *IssuedDocument) GetAmountEnasarcoTaxable() float32`

GetAmountEnasarcoTaxable returns the AmountEnasarcoTaxable field if non-nil, zero value otherwise.

### GetAmountEnasarcoTaxableOk

`func (o *IssuedDocument) GetAmountEnasarcoTaxableOk() (*float32, bool)`

GetAmountEnasarcoTaxableOk returns a tuple with the AmountEnasarcoTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountEnasarcoTaxable

`func (o *IssuedDocument) SetAmountEnasarcoTaxable(v float32)`

SetAmountEnasarcoTaxable sets AmountEnasarcoTaxable field to given value.

### HasAmountEnasarcoTaxable

`func (o *IssuedDocument) HasAmountEnasarcoTaxable() bool`

HasAmountEnasarcoTaxable returns a boolean if a field has been set.

### SetAmountEnasarcoTaxableNil

`func (o *IssuedDocument) SetAmountEnasarcoTaxableNil(b bool)`

 SetAmountEnasarcoTaxableNil sets the value for AmountEnasarcoTaxable to be an explicit nil

### UnsetAmountEnasarcoTaxable
`func (o *IssuedDocument) UnsetAmountEnasarcoTaxable()`

UnsetAmountEnasarcoTaxable ensures that no value is present for AmountEnasarcoTaxable, not even an explicit nil
### GetExtraData

`func (o *IssuedDocument) GetExtraData() IssuedDocumentExtraData`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *IssuedDocument) GetExtraDataOk() (*IssuedDocumentExtraData, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *IssuedDocument) SetExtraData(v IssuedDocumentExtraData)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *IssuedDocument) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### SetExtraDataNil

`func (o *IssuedDocument) SetExtraDataNil(b bool)`

 SetExtraDataNil sets the value for ExtraData to be an explicit nil

### UnsetExtraData
`func (o *IssuedDocument) UnsetExtraData()`

UnsetExtraData ensures that no value is present for ExtraData, not even an explicit nil
### GetSeenDate

`func (o *IssuedDocument) GetSeenDate() string`

GetSeenDate returns the SeenDate field if non-nil, zero value otherwise.

### GetSeenDateOk

`func (o *IssuedDocument) GetSeenDateOk() (*string, bool)`

GetSeenDateOk returns a tuple with the SeenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenDate

`func (o *IssuedDocument) SetSeenDate(v string)`

SetSeenDate sets SeenDate field to given value.

### HasSeenDate

`func (o *IssuedDocument) HasSeenDate() bool`

HasSeenDate returns a boolean if a field has been set.

### SetSeenDateNil

`func (o *IssuedDocument) SetSeenDateNil(b bool)`

 SetSeenDateNil sets the value for SeenDate to be an explicit nil

### UnsetSeenDate
`func (o *IssuedDocument) UnsetSeenDate()`

UnsetSeenDate ensures that no value is present for SeenDate, not even an explicit nil
### GetNextDueDate

`func (o *IssuedDocument) GetNextDueDate() string`

GetNextDueDate returns the NextDueDate field if non-nil, zero value otherwise.

### GetNextDueDateOk

`func (o *IssuedDocument) GetNextDueDateOk() (*string, bool)`

GetNextDueDateOk returns a tuple with the NextDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDueDate

`func (o *IssuedDocument) SetNextDueDate(v string)`

SetNextDueDate sets NextDueDate field to given value.

### HasNextDueDate

`func (o *IssuedDocument) HasNextDueDate() bool`

HasNextDueDate returns a boolean if a field has been set.

### SetNextDueDateNil

`func (o *IssuedDocument) SetNextDueDateNil(b bool)`

 SetNextDueDateNil sets the value for NextDueDate to be an explicit nil

### UnsetNextDueDate
`func (o *IssuedDocument) UnsetNextDueDate()`

UnsetNextDueDate ensures that no value is present for NextDueDate, not even an explicit nil
### GetUrl

`func (o *IssuedDocument) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IssuedDocument) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IssuedDocument) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IssuedDocument) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IssuedDocument) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IssuedDocument) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetDnUrl

`func (o *IssuedDocument) GetDnUrl() string`

GetDnUrl returns the DnUrl field if non-nil, zero value otherwise.

### GetDnUrlOk

`func (o *IssuedDocument) GetDnUrlOk() (*string, bool)`

GetDnUrlOk returns a tuple with the DnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnUrl

`func (o *IssuedDocument) SetDnUrl(v string)`

SetDnUrl sets DnUrl field to given value.

### HasDnUrl

`func (o *IssuedDocument) HasDnUrl() bool`

HasDnUrl returns a boolean if a field has been set.

### SetDnUrlNil

`func (o *IssuedDocument) SetDnUrlNil(b bool)`

 SetDnUrlNil sets the value for DnUrl to be an explicit nil

### UnsetDnUrl
`func (o *IssuedDocument) UnsetDnUrl()`

UnsetDnUrl ensures that no value is present for DnUrl, not even an explicit nil
### GetAiUrl

`func (o *IssuedDocument) GetAiUrl() string`

GetAiUrl returns the AiUrl field if non-nil, zero value otherwise.

### GetAiUrlOk

`func (o *IssuedDocument) GetAiUrlOk() (*string, bool)`

GetAiUrlOk returns a tuple with the AiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiUrl

`func (o *IssuedDocument) SetAiUrl(v string)`

SetAiUrl sets AiUrl field to given value.

### HasAiUrl

`func (o *IssuedDocument) HasAiUrl() bool`

HasAiUrl returns a boolean if a field has been set.

### SetAiUrlNil

`func (o *IssuedDocument) SetAiUrlNil(b bool)`

 SetAiUrlNil sets the value for AiUrl to be an explicit nil

### UnsetAiUrl
`func (o *IssuedDocument) UnsetAiUrl()`

UnsetAiUrl ensures that no value is present for AiUrl, not even an explicit nil
### GetAttachmentUrl

`func (o *IssuedDocument) GetAttachmentUrl() string`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *IssuedDocument) GetAttachmentUrlOk() (*string, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *IssuedDocument) SetAttachmentUrl(v string)`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *IssuedDocument) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### SetAttachmentUrlNil

`func (o *IssuedDocument) SetAttachmentUrlNil(b bool)`

 SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil

### UnsetAttachmentUrl
`func (o *IssuedDocument) UnsetAttachmentUrl()`

UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
### GetAttachmentToken

`func (o *IssuedDocument) GetAttachmentToken() string`

GetAttachmentToken returns the AttachmentToken field if non-nil, zero value otherwise.

### GetAttachmentTokenOk

`func (o *IssuedDocument) GetAttachmentTokenOk() (*string, bool)`

GetAttachmentTokenOk returns a tuple with the AttachmentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentToken

`func (o *IssuedDocument) SetAttachmentToken(v string)`

SetAttachmentToken sets AttachmentToken field to given value.

### HasAttachmentToken

`func (o *IssuedDocument) HasAttachmentToken() bool`

HasAttachmentToken returns a boolean if a field has been set.

### SetAttachmentTokenNil

`func (o *IssuedDocument) SetAttachmentTokenNil(b bool)`

 SetAttachmentTokenNil sets the value for AttachmentToken to be an explicit nil

### UnsetAttachmentToken
`func (o *IssuedDocument) UnsetAttachmentToken()`

UnsetAttachmentToken ensures that no value is present for AttachmentToken, not even an explicit nil
### GetEiRaw

`func (o *IssuedDocument) GetEiRaw() map[string]interface{}`

GetEiRaw returns the EiRaw field if non-nil, zero value otherwise.

### GetEiRawOk

`func (o *IssuedDocument) GetEiRawOk() (*map[string]interface{}, bool)`

GetEiRawOk returns a tuple with the EiRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiRaw

`func (o *IssuedDocument) SetEiRaw(v map[string]interface{})`

SetEiRaw sets EiRaw field to given value.

### HasEiRaw

`func (o *IssuedDocument) HasEiRaw() bool`

HasEiRaw returns a boolean if a field has been set.

### SetEiRawNil

`func (o *IssuedDocument) SetEiRawNil(b bool)`

 SetEiRawNil sets the value for EiRaw to be an explicit nil

### UnsetEiRaw
`func (o *IssuedDocument) UnsetEiRaw()`

UnsetEiRaw ensures that no value is present for EiRaw, not even an explicit nil
### GetEiStatus

`func (o *IssuedDocument) GetEiStatus() string`

GetEiStatus returns the EiStatus field if non-nil, zero value otherwise.

### GetEiStatusOk

`func (o *IssuedDocument) GetEiStatusOk() (*string, bool)`

GetEiStatusOk returns a tuple with the EiStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiStatus

`func (o *IssuedDocument) SetEiStatus(v string)`

SetEiStatus sets EiStatus field to given value.

### HasEiStatus

`func (o *IssuedDocument) HasEiStatus() bool`

HasEiStatus returns a boolean if a field has been set.

### SetEiStatusNil

`func (o *IssuedDocument) SetEiStatusNil(b bool)`

 SetEiStatusNil sets the value for EiStatus to be an explicit nil

### UnsetEiStatus
`func (o *IssuedDocument) UnsetEiStatus()`

UnsetEiStatus ensures that no value is present for EiStatus, not even an explicit nil
### GetLocked

`func (o *IssuedDocument) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *IssuedDocument) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *IssuedDocument) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *IssuedDocument) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *IssuedDocument) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *IssuedDocument) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetCreatedAt

`func (o *IssuedDocument) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssuedDocument) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssuedDocument) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssuedDocument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IssuedDocument) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IssuedDocument) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *IssuedDocument) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssuedDocument) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssuedDocument) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssuedDocument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IssuedDocument) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IssuedDocument) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


