# CashbookEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Cashbook id | [optional] 
**Date** | Pointer to **NullableString** | Cashbook date | [optional] 
**Description** | Pointer to **NullableString** | Cashbook description | [optional] 
**Kind** | Pointer to [**CashbookEntryKind**](CashbookEntryKind.md) |  | [optional] 
**Type** | Pointer to [**NullableCashbookEntryType**](CashbookEntryType.md) |  | [optional] 
**EntityName** | Pointer to **NullableString** | Cashbook entity name | [optional] 
**Document** | Pointer to [**NullableCashbookEntryDocument**](CashbookEntryDocument.md) |  | [optional] 
**AmountIn** | Pointer to **NullableFloat32** | [Only for cashbook entry in] Cashbook total amount in | [optional] 
**PaymentAccountIn** | Pointer to [**NullablePaymentAccount**](PaymentAccount.md) |  | [optional] 
**AmountOut** | Pointer to **NullableFloat32** | [Only for cashbook entry out] Cashbook total amount out | [optional] 
**PaymentAccountOut** | Pointer to [**NullablePaymentAccount**](PaymentAccount.md) |  | [optional] 

## Methods

### NewCashbookEntry

`func NewCashbookEntry() *CashbookEntry`

NewCashbookEntry instantiates a new CashbookEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashbookEntryWithDefaults

`func NewCashbookEntryWithDefaults() *CashbookEntry`

NewCashbookEntryWithDefaults instantiates a new CashbookEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CashbookEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CashbookEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CashbookEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CashbookEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CashbookEntry) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CashbookEntry) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDate

`func (o *CashbookEntry) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CashbookEntry) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CashbookEntry) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CashbookEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *CashbookEntry) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *CashbookEntry) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *CashbookEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CashbookEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CashbookEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CashbookEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CashbookEntry) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CashbookEntry) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKind

`func (o *CashbookEntry) GetKind() CashbookEntryKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CashbookEntry) GetKindOk() (*CashbookEntryKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CashbookEntry) SetKind(v CashbookEntryKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CashbookEntry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetType

`func (o *CashbookEntry) GetType() CashbookEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CashbookEntry) GetTypeOk() (*CashbookEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CashbookEntry) SetType(v CashbookEntryType)`

SetType sets Type field to given value.

### HasType

`func (o *CashbookEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CashbookEntry) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CashbookEntry) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEntityName

`func (o *CashbookEntry) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *CashbookEntry) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *CashbookEntry) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *CashbookEntry) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### SetEntityNameNil

`func (o *CashbookEntry) SetEntityNameNil(b bool)`

 SetEntityNameNil sets the value for EntityName to be an explicit nil

### UnsetEntityName
`func (o *CashbookEntry) UnsetEntityName()`

UnsetEntityName ensures that no value is present for EntityName, not even an explicit nil
### GetDocument

`func (o *CashbookEntry) GetDocument() CashbookEntryDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *CashbookEntry) GetDocumentOk() (*CashbookEntryDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *CashbookEntry) SetDocument(v CashbookEntryDocument)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *CashbookEntry) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *CashbookEntry) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *CashbookEntry) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetAmountIn

`func (o *CashbookEntry) GetAmountIn() float32`

GetAmountIn returns the AmountIn field if non-nil, zero value otherwise.

### GetAmountInOk

`func (o *CashbookEntry) GetAmountInOk() (*float32, bool)`

GetAmountInOk returns a tuple with the AmountIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIn

`func (o *CashbookEntry) SetAmountIn(v float32)`

SetAmountIn sets AmountIn field to given value.

### HasAmountIn

`func (o *CashbookEntry) HasAmountIn() bool`

HasAmountIn returns a boolean if a field has been set.

### SetAmountInNil

`func (o *CashbookEntry) SetAmountInNil(b bool)`

 SetAmountInNil sets the value for AmountIn to be an explicit nil

### UnsetAmountIn
`func (o *CashbookEntry) UnsetAmountIn()`

UnsetAmountIn ensures that no value is present for AmountIn, not even an explicit nil
### GetPaymentAccountIn

`func (o *CashbookEntry) GetPaymentAccountIn() PaymentAccount`

GetPaymentAccountIn returns the PaymentAccountIn field if non-nil, zero value otherwise.

### GetPaymentAccountInOk

`func (o *CashbookEntry) GetPaymentAccountInOk() (*PaymentAccount, bool)`

GetPaymentAccountInOk returns a tuple with the PaymentAccountIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountIn

`func (o *CashbookEntry) SetPaymentAccountIn(v PaymentAccount)`

SetPaymentAccountIn sets PaymentAccountIn field to given value.

### HasPaymentAccountIn

`func (o *CashbookEntry) HasPaymentAccountIn() bool`

HasPaymentAccountIn returns a boolean if a field has been set.

### SetPaymentAccountInNil

`func (o *CashbookEntry) SetPaymentAccountInNil(b bool)`

 SetPaymentAccountInNil sets the value for PaymentAccountIn to be an explicit nil

### UnsetPaymentAccountIn
`func (o *CashbookEntry) UnsetPaymentAccountIn()`

UnsetPaymentAccountIn ensures that no value is present for PaymentAccountIn, not even an explicit nil
### GetAmountOut

`func (o *CashbookEntry) GetAmountOut() float32`

GetAmountOut returns the AmountOut field if non-nil, zero value otherwise.

### GetAmountOutOk

`func (o *CashbookEntry) GetAmountOutOk() (*float32, bool)`

GetAmountOutOk returns a tuple with the AmountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOut

`func (o *CashbookEntry) SetAmountOut(v float32)`

SetAmountOut sets AmountOut field to given value.

### HasAmountOut

`func (o *CashbookEntry) HasAmountOut() bool`

HasAmountOut returns a boolean if a field has been set.

### SetAmountOutNil

`func (o *CashbookEntry) SetAmountOutNil(b bool)`

 SetAmountOutNil sets the value for AmountOut to be an explicit nil

### UnsetAmountOut
`func (o *CashbookEntry) UnsetAmountOut()`

UnsetAmountOut ensures that no value is present for AmountOut, not even an explicit nil
### GetPaymentAccountOut

`func (o *CashbookEntry) GetPaymentAccountOut() PaymentAccount`

GetPaymentAccountOut returns the PaymentAccountOut field if non-nil, zero value otherwise.

### GetPaymentAccountOutOk

`func (o *CashbookEntry) GetPaymentAccountOutOk() (*PaymentAccount, bool)`

GetPaymentAccountOutOk returns a tuple with the PaymentAccountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountOut

`func (o *CashbookEntry) SetPaymentAccountOut(v PaymentAccount)`

SetPaymentAccountOut sets PaymentAccountOut field to given value.

### HasPaymentAccountOut

`func (o *CashbookEntry) HasPaymentAccountOut() bool`

HasPaymentAccountOut returns a boolean if a field has been set.

### SetPaymentAccountOutNil

`func (o *CashbookEntry) SetPaymentAccountOutNil(b bool)`

 SetPaymentAccountOutNil sets the value for PaymentAccountOut to be an explicit nil

### UnsetPaymentAccountOut
`func (o *CashbookEntry) UnsetPaymentAccountOut()`

UnsetPaymentAccountOut ensures that no value is present for PaymentAccountOut, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


