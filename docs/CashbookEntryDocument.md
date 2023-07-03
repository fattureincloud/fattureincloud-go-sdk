# CashbookEntryDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Cashbook related document id | [optional] 
**Type** | Pointer to **NullableString** | Cashbook related document type | [optional] 
**Path** | Pointer to **NullableString** | Cashbook related document path | [optional] 

## Methods

### NewCashbookEntryDocument

`func NewCashbookEntryDocument() *CashbookEntryDocument`

NewCashbookEntryDocument instantiates a new CashbookEntryDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashbookEntryDocumentWithDefaults

`func NewCashbookEntryDocumentWithDefaults() *CashbookEntryDocument`

NewCashbookEntryDocumentWithDefaults instantiates a new CashbookEntryDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CashbookEntryDocument) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CashbookEntryDocument) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CashbookEntryDocument) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CashbookEntryDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CashbookEntryDocument) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CashbookEntryDocument) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *CashbookEntryDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CashbookEntryDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CashbookEntryDocument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CashbookEntryDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CashbookEntryDocument) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CashbookEntryDocument) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPath

`func (o *CashbookEntryDocument) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CashbookEntryDocument) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CashbookEntryDocument) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CashbookEntryDocument) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *CashbookEntryDocument) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CashbookEntryDocument) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


