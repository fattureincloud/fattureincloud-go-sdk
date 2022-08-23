# PaymentAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique identifier | [optional] 
**Name** | Pointer to **NullableString** | Payment account name. | [optional] 
**Type** | Pointer to [**PaymentAccountType**](PaymentAccountType.md) |  | [optional] [default to STANDARD]
**Iban** | Pointer to **NullableString** | Payment account iban. | [optional] 
**Sia** | Pointer to **NullableString** | Payment account sia. | [optional] 
**Cuc** | Pointer to **NullableString** | Payment account cuc. | [optional] 
**Virtual** | Pointer to **NullableBool** | Determine if the payment method is virtual. | [optional] 

## Methods

### NewPaymentAccount

`func NewPaymentAccount() *PaymentAccount`

NewPaymentAccount instantiates a new PaymentAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAccountWithDefaults

`func NewPaymentAccountWithDefaults() *PaymentAccount`

NewPaymentAccountWithDefaults instantiates a new PaymentAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentAccount) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentAccount) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *PaymentAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PaymentAccount) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PaymentAccount) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *PaymentAccount) GetType() PaymentAccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentAccount) GetTypeOk() (*PaymentAccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentAccount) SetType(v PaymentAccountType)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIban

`func (o *PaymentAccount) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentAccount) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentAccount) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PaymentAccount) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *PaymentAccount) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *PaymentAccount) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetSia

`func (o *PaymentAccount) GetSia() string`

GetSia returns the Sia field if non-nil, zero value otherwise.

### GetSiaOk

`func (o *PaymentAccount) GetSiaOk() (*string, bool)`

GetSiaOk returns a tuple with the Sia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSia

`func (o *PaymentAccount) SetSia(v string)`

SetSia sets Sia field to given value.

### HasSia

`func (o *PaymentAccount) HasSia() bool`

HasSia returns a boolean if a field has been set.

### SetSiaNil

`func (o *PaymentAccount) SetSiaNil(b bool)`

 SetSiaNil sets the value for Sia to be an explicit nil

### UnsetSia
`func (o *PaymentAccount) UnsetSia()`

UnsetSia ensures that no value is present for Sia, not even an explicit nil
### GetCuc

`func (o *PaymentAccount) GetCuc() string`

GetCuc returns the Cuc field if non-nil, zero value otherwise.

### GetCucOk

`func (o *PaymentAccount) GetCucOk() (*string, bool)`

GetCucOk returns a tuple with the Cuc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuc

`func (o *PaymentAccount) SetCuc(v string)`

SetCuc sets Cuc field to given value.

### HasCuc

`func (o *PaymentAccount) HasCuc() bool`

HasCuc returns a boolean if a field has been set.

### SetCucNil

`func (o *PaymentAccount) SetCucNil(b bool)`

 SetCucNil sets the value for Cuc to be an explicit nil

### UnsetCuc
`func (o *PaymentAccount) UnsetCuc()`

UnsetCuc ensures that no value is present for Cuc, not even an explicit nil
### GetVirtual

`func (o *PaymentAccount) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *PaymentAccount) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *PaymentAccount) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *PaymentAccount) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.

### SetVirtualNil

`func (o *PaymentAccount) SetVirtualNil(b bool)`

 SetVirtualNil sets the value for Virtual to be an explicit nil

### UnsetVirtual
`func (o *PaymentAccount) UnsetVirtual()`

UnsetVirtual ensures that no value is present for Virtual, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


