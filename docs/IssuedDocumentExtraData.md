# IssuedDocumentExtraData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowSofortButton** | Pointer to **NullableBool** |  | [optional] 
**MultifattureSent** | Pointer to **NullableInt32** |  | [optional] 
**TsCommunication** | Pointer to **NullableBool** | Send issued document to \&quot;Sistema Tessera Sanitaria\&quot; | [optional] 
**TsFlagTipoSpesa** | Pointer to **NullableFloat32** | Issued document ts \&quot;tipo spesa\&quot; [TK, FC, FV, SV,SP, AD, AS, ECG, SR] | [optional] 
**TsPagamentoTracciato** | Pointer to **NullableBool** | Issued document ts traced payment | [optional] 
**TsTipoSpesa** | Pointer to **NullableString** | Can be [ &#39;TK&#39;, &#39;FC&#39;, &#39;FV&#39;, &#39;SV&#39;, &#39;SP&#39;, &#39;AD&#39;, &#39;AS&#39;, &#39;SR&#39;, &#39;CT&#39;, &#39;PI&#39;, &#39;IC&#39;, &#39;AA&#39; ]. Refer to the technical specifications to learn more. | [optional] 
**TsOpposizione** | Pointer to **NullableBool** | Issued document ts \&quot;opposizione\&quot; | [optional] 
**TsStatus** | Pointer to **NullableInt32** | Issued document ts status | [optional] 
**TsFileId** | Pointer to **NullableString** | Issued document ts file id | [optional] 
**TsSentDate** | Pointer to **NullableString** | Issued document ts sent date | [optional] 
**TsFullAmount** | Pointer to **NullableBool** | Issued document ts total amount | [optional] 
**ImportedBy** | Pointer to **NullableString** | Issued document imported by software | [optional] 

## Methods

### NewIssuedDocumentExtraData

`func NewIssuedDocumentExtraData() *IssuedDocumentExtraData`

NewIssuedDocumentExtraData instantiates a new IssuedDocumentExtraData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentExtraDataWithDefaults

`func NewIssuedDocumentExtraDataWithDefaults() *IssuedDocumentExtraData`

NewIssuedDocumentExtraDataWithDefaults instantiates a new IssuedDocumentExtraData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowSofortButton

`func (o *IssuedDocumentExtraData) GetShowSofortButton() bool`

GetShowSofortButton returns the ShowSofortButton field if non-nil, zero value otherwise.

### GetShowSofortButtonOk

`func (o *IssuedDocumentExtraData) GetShowSofortButtonOk() (*bool, bool)`

GetShowSofortButtonOk returns a tuple with the ShowSofortButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSofortButton

`func (o *IssuedDocumentExtraData) SetShowSofortButton(v bool)`

SetShowSofortButton sets ShowSofortButton field to given value.

### HasShowSofortButton

`func (o *IssuedDocumentExtraData) HasShowSofortButton() bool`

HasShowSofortButton returns a boolean if a field has been set.

### SetShowSofortButtonNil

`func (o *IssuedDocumentExtraData) SetShowSofortButtonNil(b bool)`

 SetShowSofortButtonNil sets the value for ShowSofortButton to be an explicit nil

### UnsetShowSofortButton
`func (o *IssuedDocumentExtraData) UnsetShowSofortButton()`

UnsetShowSofortButton ensures that no value is present for ShowSofortButton, not even an explicit nil
### GetMultifattureSent

`func (o *IssuedDocumentExtraData) GetMultifattureSent() int32`

GetMultifattureSent returns the MultifattureSent field if non-nil, zero value otherwise.

### GetMultifattureSentOk

`func (o *IssuedDocumentExtraData) GetMultifattureSentOk() (*int32, bool)`

GetMultifattureSentOk returns a tuple with the MultifattureSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultifattureSent

`func (o *IssuedDocumentExtraData) SetMultifattureSent(v int32)`

SetMultifattureSent sets MultifattureSent field to given value.

### HasMultifattureSent

`func (o *IssuedDocumentExtraData) HasMultifattureSent() bool`

HasMultifattureSent returns a boolean if a field has been set.

### SetMultifattureSentNil

`func (o *IssuedDocumentExtraData) SetMultifattureSentNil(b bool)`

 SetMultifattureSentNil sets the value for MultifattureSent to be an explicit nil

### UnsetMultifattureSent
`func (o *IssuedDocumentExtraData) UnsetMultifattureSent()`

UnsetMultifattureSent ensures that no value is present for MultifattureSent, not even an explicit nil
### GetTsCommunication

`func (o *IssuedDocumentExtraData) GetTsCommunication() bool`

GetTsCommunication returns the TsCommunication field if non-nil, zero value otherwise.

### GetTsCommunicationOk

`func (o *IssuedDocumentExtraData) GetTsCommunicationOk() (*bool, bool)`

GetTsCommunicationOk returns a tuple with the TsCommunication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsCommunication

`func (o *IssuedDocumentExtraData) SetTsCommunication(v bool)`

SetTsCommunication sets TsCommunication field to given value.

### HasTsCommunication

`func (o *IssuedDocumentExtraData) HasTsCommunication() bool`

HasTsCommunication returns a boolean if a field has been set.

### SetTsCommunicationNil

`func (o *IssuedDocumentExtraData) SetTsCommunicationNil(b bool)`

 SetTsCommunicationNil sets the value for TsCommunication to be an explicit nil

### UnsetTsCommunication
`func (o *IssuedDocumentExtraData) UnsetTsCommunication()`

UnsetTsCommunication ensures that no value is present for TsCommunication, not even an explicit nil
### GetTsFlagTipoSpesa

`func (o *IssuedDocumentExtraData) GetTsFlagTipoSpesa() float32`

GetTsFlagTipoSpesa returns the TsFlagTipoSpesa field if non-nil, zero value otherwise.

### GetTsFlagTipoSpesaOk

`func (o *IssuedDocumentExtraData) GetTsFlagTipoSpesaOk() (*float32, bool)`

GetTsFlagTipoSpesaOk returns a tuple with the TsFlagTipoSpesa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsFlagTipoSpesa

`func (o *IssuedDocumentExtraData) SetTsFlagTipoSpesa(v float32)`

SetTsFlagTipoSpesa sets TsFlagTipoSpesa field to given value.

### HasTsFlagTipoSpesa

`func (o *IssuedDocumentExtraData) HasTsFlagTipoSpesa() bool`

HasTsFlagTipoSpesa returns a boolean if a field has been set.

### SetTsFlagTipoSpesaNil

`func (o *IssuedDocumentExtraData) SetTsFlagTipoSpesaNil(b bool)`

 SetTsFlagTipoSpesaNil sets the value for TsFlagTipoSpesa to be an explicit nil

### UnsetTsFlagTipoSpesa
`func (o *IssuedDocumentExtraData) UnsetTsFlagTipoSpesa()`

UnsetTsFlagTipoSpesa ensures that no value is present for TsFlagTipoSpesa, not even an explicit nil
### GetTsPagamentoTracciato

`func (o *IssuedDocumentExtraData) GetTsPagamentoTracciato() bool`

GetTsPagamentoTracciato returns the TsPagamentoTracciato field if non-nil, zero value otherwise.

### GetTsPagamentoTracciatoOk

`func (o *IssuedDocumentExtraData) GetTsPagamentoTracciatoOk() (*bool, bool)`

GetTsPagamentoTracciatoOk returns a tuple with the TsPagamentoTracciato field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsPagamentoTracciato

`func (o *IssuedDocumentExtraData) SetTsPagamentoTracciato(v bool)`

SetTsPagamentoTracciato sets TsPagamentoTracciato field to given value.

### HasTsPagamentoTracciato

`func (o *IssuedDocumentExtraData) HasTsPagamentoTracciato() bool`

HasTsPagamentoTracciato returns a boolean if a field has been set.

### SetTsPagamentoTracciatoNil

`func (o *IssuedDocumentExtraData) SetTsPagamentoTracciatoNil(b bool)`

 SetTsPagamentoTracciatoNil sets the value for TsPagamentoTracciato to be an explicit nil

### UnsetTsPagamentoTracciato
`func (o *IssuedDocumentExtraData) UnsetTsPagamentoTracciato()`

UnsetTsPagamentoTracciato ensures that no value is present for TsPagamentoTracciato, not even an explicit nil
### GetTsTipoSpesa

`func (o *IssuedDocumentExtraData) GetTsTipoSpesa() string`

GetTsTipoSpesa returns the TsTipoSpesa field if non-nil, zero value otherwise.

### GetTsTipoSpesaOk

`func (o *IssuedDocumentExtraData) GetTsTipoSpesaOk() (*string, bool)`

GetTsTipoSpesaOk returns a tuple with the TsTipoSpesa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsTipoSpesa

`func (o *IssuedDocumentExtraData) SetTsTipoSpesa(v string)`

SetTsTipoSpesa sets TsTipoSpesa field to given value.

### HasTsTipoSpesa

`func (o *IssuedDocumentExtraData) HasTsTipoSpesa() bool`

HasTsTipoSpesa returns a boolean if a field has been set.

### SetTsTipoSpesaNil

`func (o *IssuedDocumentExtraData) SetTsTipoSpesaNil(b bool)`

 SetTsTipoSpesaNil sets the value for TsTipoSpesa to be an explicit nil

### UnsetTsTipoSpesa
`func (o *IssuedDocumentExtraData) UnsetTsTipoSpesa()`

UnsetTsTipoSpesa ensures that no value is present for TsTipoSpesa, not even an explicit nil
### GetTsOpposizione

`func (o *IssuedDocumentExtraData) GetTsOpposizione() bool`

GetTsOpposizione returns the TsOpposizione field if non-nil, zero value otherwise.

### GetTsOpposizioneOk

`func (o *IssuedDocumentExtraData) GetTsOpposizioneOk() (*bool, bool)`

GetTsOpposizioneOk returns a tuple with the TsOpposizione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsOpposizione

`func (o *IssuedDocumentExtraData) SetTsOpposizione(v bool)`

SetTsOpposizione sets TsOpposizione field to given value.

### HasTsOpposizione

`func (o *IssuedDocumentExtraData) HasTsOpposizione() bool`

HasTsOpposizione returns a boolean if a field has been set.

### SetTsOpposizioneNil

`func (o *IssuedDocumentExtraData) SetTsOpposizioneNil(b bool)`

 SetTsOpposizioneNil sets the value for TsOpposizione to be an explicit nil

### UnsetTsOpposizione
`func (o *IssuedDocumentExtraData) UnsetTsOpposizione()`

UnsetTsOpposizione ensures that no value is present for TsOpposizione, not even an explicit nil
### GetTsStatus

`func (o *IssuedDocumentExtraData) GetTsStatus() int32`

GetTsStatus returns the TsStatus field if non-nil, zero value otherwise.

### GetTsStatusOk

`func (o *IssuedDocumentExtraData) GetTsStatusOk() (*int32, bool)`

GetTsStatusOk returns a tuple with the TsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsStatus

`func (o *IssuedDocumentExtraData) SetTsStatus(v int32)`

SetTsStatus sets TsStatus field to given value.

### HasTsStatus

`func (o *IssuedDocumentExtraData) HasTsStatus() bool`

HasTsStatus returns a boolean if a field has been set.

### SetTsStatusNil

`func (o *IssuedDocumentExtraData) SetTsStatusNil(b bool)`

 SetTsStatusNil sets the value for TsStatus to be an explicit nil

### UnsetTsStatus
`func (o *IssuedDocumentExtraData) UnsetTsStatus()`

UnsetTsStatus ensures that no value is present for TsStatus, not even an explicit nil
### GetTsFileId

`func (o *IssuedDocumentExtraData) GetTsFileId() string`

GetTsFileId returns the TsFileId field if non-nil, zero value otherwise.

### GetTsFileIdOk

`func (o *IssuedDocumentExtraData) GetTsFileIdOk() (*string, bool)`

GetTsFileIdOk returns a tuple with the TsFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsFileId

`func (o *IssuedDocumentExtraData) SetTsFileId(v string)`

SetTsFileId sets TsFileId field to given value.

### HasTsFileId

`func (o *IssuedDocumentExtraData) HasTsFileId() bool`

HasTsFileId returns a boolean if a field has been set.

### SetTsFileIdNil

`func (o *IssuedDocumentExtraData) SetTsFileIdNil(b bool)`

 SetTsFileIdNil sets the value for TsFileId to be an explicit nil

### UnsetTsFileId
`func (o *IssuedDocumentExtraData) UnsetTsFileId()`

UnsetTsFileId ensures that no value is present for TsFileId, not even an explicit nil
### GetTsSentDate

`func (o *IssuedDocumentExtraData) GetTsSentDate() string`

GetTsSentDate returns the TsSentDate field if non-nil, zero value otherwise.

### GetTsSentDateOk

`func (o *IssuedDocumentExtraData) GetTsSentDateOk() (*string, bool)`

GetTsSentDateOk returns a tuple with the TsSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsSentDate

`func (o *IssuedDocumentExtraData) SetTsSentDate(v string)`

SetTsSentDate sets TsSentDate field to given value.

### HasTsSentDate

`func (o *IssuedDocumentExtraData) HasTsSentDate() bool`

HasTsSentDate returns a boolean if a field has been set.

### SetTsSentDateNil

`func (o *IssuedDocumentExtraData) SetTsSentDateNil(b bool)`

 SetTsSentDateNil sets the value for TsSentDate to be an explicit nil

### UnsetTsSentDate
`func (o *IssuedDocumentExtraData) UnsetTsSentDate()`

UnsetTsSentDate ensures that no value is present for TsSentDate, not even an explicit nil
### GetTsFullAmount

`func (o *IssuedDocumentExtraData) GetTsFullAmount() bool`

GetTsFullAmount returns the TsFullAmount field if non-nil, zero value otherwise.

### GetTsFullAmountOk

`func (o *IssuedDocumentExtraData) GetTsFullAmountOk() (*bool, bool)`

GetTsFullAmountOk returns a tuple with the TsFullAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsFullAmount

`func (o *IssuedDocumentExtraData) SetTsFullAmount(v bool)`

SetTsFullAmount sets TsFullAmount field to given value.

### HasTsFullAmount

`func (o *IssuedDocumentExtraData) HasTsFullAmount() bool`

HasTsFullAmount returns a boolean if a field has been set.

### SetTsFullAmountNil

`func (o *IssuedDocumentExtraData) SetTsFullAmountNil(b bool)`

 SetTsFullAmountNil sets the value for TsFullAmount to be an explicit nil

### UnsetTsFullAmount
`func (o *IssuedDocumentExtraData) UnsetTsFullAmount()`

UnsetTsFullAmount ensures that no value is present for TsFullAmount, not even an explicit nil
### GetImportedBy

`func (o *IssuedDocumentExtraData) GetImportedBy() string`

GetImportedBy returns the ImportedBy field if non-nil, zero value otherwise.

### GetImportedByOk

`func (o *IssuedDocumentExtraData) GetImportedByOk() (*string, bool)`

GetImportedByOk returns a tuple with the ImportedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedBy

`func (o *IssuedDocumentExtraData) SetImportedBy(v string)`

SetImportedBy sets ImportedBy field to given value.

### HasImportedBy

`func (o *IssuedDocumentExtraData) HasImportedBy() bool`

HasImportedBy returns a boolean if a field has been set.

### SetImportedByNil

`func (o *IssuedDocumentExtraData) SetImportedByNil(b bool)`

 SetImportedByNil sets the value for ImportedBy to be an explicit nil

### UnsetImportedBy
`func (o *IssuedDocumentExtraData) UnsetImportedBy()`

UnsetImportedBy ensures that no value is present for ImportedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


