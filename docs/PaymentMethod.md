# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique identifier | [optional] 
**Name** | Pointer to **NullableString** | Name of the payment method | [optional] 
**Type** | Pointer to [**PaymentMethodType**](PaymentMethodType.md) |  | [optional] [default to STANDARD]
**IsDefault** | Pointer to **NullableBool** | Determines if this is the default payment method. | [optional] 
**DefaultPaymentAccount** | Pointer to [**NullablePaymentAccount**](PaymentAccount.md) |  | [optional] 
**Details** | Pointer to [**[]PaymentMethodDetails**](PaymentMethodDetails.md) | Method details rows | [optional] 
**BankIban** | Pointer to **NullableString** | Bank iban | [optional] 
**BankName** | Pointer to **NullableString** | Bank name | [optional] 
**BankBeneficiary** | Pointer to **NullableString** | Bank beneficiary | [optional] 
**EiPaymentMethod** | Pointer to **NullableString** | E-invoice payment method | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentMethod) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentMethod) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *PaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PaymentMethod) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PaymentMethod) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *PaymentMethod) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v PaymentMethodType)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsDefault

`func (o *PaymentMethod) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PaymentMethod) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PaymentMethod) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PaymentMethod) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *PaymentMethod) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *PaymentMethod) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetDefaultPaymentAccount

`func (o *PaymentMethod) GetDefaultPaymentAccount() PaymentAccount`

GetDefaultPaymentAccount returns the DefaultPaymentAccount field if non-nil, zero value otherwise.

### GetDefaultPaymentAccountOk

`func (o *PaymentMethod) GetDefaultPaymentAccountOk() (*PaymentAccount, bool)`

GetDefaultPaymentAccountOk returns a tuple with the DefaultPaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentAccount

`func (o *PaymentMethod) SetDefaultPaymentAccount(v PaymentAccount)`

SetDefaultPaymentAccount sets DefaultPaymentAccount field to given value.

### HasDefaultPaymentAccount

`func (o *PaymentMethod) HasDefaultPaymentAccount() bool`

HasDefaultPaymentAccount returns a boolean if a field has been set.

### SetDefaultPaymentAccountNil

`func (o *PaymentMethod) SetDefaultPaymentAccountNil(b bool)`

 SetDefaultPaymentAccountNil sets the value for DefaultPaymentAccount to be an explicit nil

### UnsetDefaultPaymentAccount
`func (o *PaymentMethod) UnsetDefaultPaymentAccount()`

UnsetDefaultPaymentAccount ensures that no value is present for DefaultPaymentAccount, not even an explicit nil
### GetDetails

`func (o *PaymentMethod) GetDetails() []PaymentMethodDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PaymentMethod) GetDetailsOk() (*[]PaymentMethodDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PaymentMethod) SetDetails(v []PaymentMethodDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PaymentMethod) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PaymentMethod) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PaymentMethod) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetBankIban

`func (o *PaymentMethod) GetBankIban() string`

GetBankIban returns the BankIban field if non-nil, zero value otherwise.

### GetBankIbanOk

`func (o *PaymentMethod) GetBankIbanOk() (*string, bool)`

GetBankIbanOk returns a tuple with the BankIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankIban

`func (o *PaymentMethod) SetBankIban(v string)`

SetBankIban sets BankIban field to given value.

### HasBankIban

`func (o *PaymentMethod) HasBankIban() bool`

HasBankIban returns a boolean if a field has been set.

### SetBankIbanNil

`func (o *PaymentMethod) SetBankIbanNil(b bool)`

 SetBankIbanNil sets the value for BankIban to be an explicit nil

### UnsetBankIban
`func (o *PaymentMethod) UnsetBankIban()`

UnsetBankIban ensures that no value is present for BankIban, not even an explicit nil
### GetBankName

`func (o *PaymentMethod) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PaymentMethod) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PaymentMethod) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PaymentMethod) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### SetBankNameNil

`func (o *PaymentMethod) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *PaymentMethod) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetBankBeneficiary

`func (o *PaymentMethod) GetBankBeneficiary() string`

GetBankBeneficiary returns the BankBeneficiary field if non-nil, zero value otherwise.

### GetBankBeneficiaryOk

`func (o *PaymentMethod) GetBankBeneficiaryOk() (*string, bool)`

GetBankBeneficiaryOk returns a tuple with the BankBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBeneficiary

`func (o *PaymentMethod) SetBankBeneficiary(v string)`

SetBankBeneficiary sets BankBeneficiary field to given value.

### HasBankBeneficiary

`func (o *PaymentMethod) HasBankBeneficiary() bool`

HasBankBeneficiary returns a boolean if a field has been set.

### SetBankBeneficiaryNil

`func (o *PaymentMethod) SetBankBeneficiaryNil(b bool)`

 SetBankBeneficiaryNil sets the value for BankBeneficiary to be an explicit nil

### UnsetBankBeneficiary
`func (o *PaymentMethod) UnsetBankBeneficiary()`

UnsetBankBeneficiary ensures that no value is present for BankBeneficiary, not even an explicit nil
### GetEiPaymentMethod

`func (o *PaymentMethod) GetEiPaymentMethod() string`

GetEiPaymentMethod returns the EiPaymentMethod field if non-nil, zero value otherwise.

### GetEiPaymentMethodOk

`func (o *PaymentMethod) GetEiPaymentMethodOk() (*string, bool)`

GetEiPaymentMethodOk returns a tuple with the EiPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiPaymentMethod

`func (o *PaymentMethod) SetEiPaymentMethod(v string)`

SetEiPaymentMethod sets EiPaymentMethod field to given value.

### HasEiPaymentMethod

`func (o *PaymentMethod) HasEiPaymentMethod() bool`

HasEiPaymentMethod returns a boolean if a field has been set.

### SetEiPaymentMethodNil

`func (o *PaymentMethod) SetEiPaymentMethodNil(b bool)`

 SetEiPaymentMethodNil sets the value for EiPaymentMethod to be an explicit nil

### UnsetEiPaymentMethod
`func (o *PaymentMethod) UnsetEiPaymentMethod()`

UnsetEiPaymentMethod ensures that no value is present for EiPaymentMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


