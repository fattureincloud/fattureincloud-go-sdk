# CompanyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Company id | [optional] 
**Name** | Pointer to **NullableString** | Company name | [optional] 
**Email** | Pointer to **NullableString** | Company email | [optional] 
**Type** | Pointer to [**CompanyType**](CompanyType.md) |  | [optional] 
**AccessInfo** | Pointer to [**NullableCompanyInfoAccessInfo**](CompanyInfoAccessInfo.md) |  | [optional] 
**PlanInfo** | Pointer to [**NullableCompanyInfoPlanInfo**](CompanyInfoPlanInfo.md) |  | [optional] 
**AccountantId** | Pointer to **NullableInt32** | Company accountant id | [optional] 
**IsAccountant** | Pointer to **NullableBool** | Is the logged account an accountant. | [optional] 

## Methods

### NewCompanyInfo

`func NewCompanyInfo() *CompanyInfo`

NewCompanyInfo instantiates a new CompanyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyInfoWithDefaults

`func NewCompanyInfoWithDefaults() *CompanyInfo`

NewCompanyInfoWithDefaults instantiates a new CompanyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CompanyInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CompanyInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CompanyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CompanyInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CompanyInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *CompanyInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CompanyInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CompanyInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CompanyInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CompanyInfo) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CompanyInfo) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetType

`func (o *CompanyInfo) GetType() CompanyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyInfo) GetTypeOk() (*CompanyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyInfo) SetType(v CompanyType)`

SetType sets Type field to given value.

### HasType

`func (o *CompanyInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessInfo

`func (o *CompanyInfo) GetAccessInfo() CompanyInfoAccessInfo`

GetAccessInfo returns the AccessInfo field if non-nil, zero value otherwise.

### GetAccessInfoOk

`func (o *CompanyInfo) GetAccessInfoOk() (*CompanyInfoAccessInfo, bool)`

GetAccessInfoOk returns a tuple with the AccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessInfo

`func (o *CompanyInfo) SetAccessInfo(v CompanyInfoAccessInfo)`

SetAccessInfo sets AccessInfo field to given value.

### HasAccessInfo

`func (o *CompanyInfo) HasAccessInfo() bool`

HasAccessInfo returns a boolean if a field has been set.

### SetAccessInfoNil

`func (o *CompanyInfo) SetAccessInfoNil(b bool)`

 SetAccessInfoNil sets the value for AccessInfo to be an explicit nil

### UnsetAccessInfo
`func (o *CompanyInfo) UnsetAccessInfo()`

UnsetAccessInfo ensures that no value is present for AccessInfo, not even an explicit nil
### GetPlanInfo

`func (o *CompanyInfo) GetPlanInfo() CompanyInfoPlanInfo`

GetPlanInfo returns the PlanInfo field if non-nil, zero value otherwise.

### GetPlanInfoOk

`func (o *CompanyInfo) GetPlanInfoOk() (*CompanyInfoPlanInfo, bool)`

GetPlanInfoOk returns a tuple with the PlanInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanInfo

`func (o *CompanyInfo) SetPlanInfo(v CompanyInfoPlanInfo)`

SetPlanInfo sets PlanInfo field to given value.

### HasPlanInfo

`func (o *CompanyInfo) HasPlanInfo() bool`

HasPlanInfo returns a boolean if a field has been set.

### SetPlanInfoNil

`func (o *CompanyInfo) SetPlanInfoNil(b bool)`

 SetPlanInfoNil sets the value for PlanInfo to be an explicit nil

### UnsetPlanInfo
`func (o *CompanyInfo) UnsetPlanInfo()`

UnsetPlanInfo ensures that no value is present for PlanInfo, not even an explicit nil
### GetAccountantId

`func (o *CompanyInfo) GetAccountantId() int32`

GetAccountantId returns the AccountantId field if non-nil, zero value otherwise.

### GetAccountantIdOk

`func (o *CompanyInfo) GetAccountantIdOk() (*int32, bool)`

GetAccountantIdOk returns a tuple with the AccountantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountantId

`func (o *CompanyInfo) SetAccountantId(v int32)`

SetAccountantId sets AccountantId field to given value.

### HasAccountantId

`func (o *CompanyInfo) HasAccountantId() bool`

HasAccountantId returns a boolean if a field has been set.

### SetAccountantIdNil

`func (o *CompanyInfo) SetAccountantIdNil(b bool)`

 SetAccountantIdNil sets the value for AccountantId to be an explicit nil

### UnsetAccountantId
`func (o *CompanyInfo) UnsetAccountantId()`

UnsetAccountantId ensures that no value is present for AccountantId, not even an explicit nil
### GetIsAccountant

`func (o *CompanyInfo) GetIsAccountant() bool`

GetIsAccountant returns the IsAccountant field if non-nil, zero value otherwise.

### GetIsAccountantOk

`func (o *CompanyInfo) GetIsAccountantOk() (*bool, bool)`

GetIsAccountantOk returns a tuple with the IsAccountant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountant

`func (o *CompanyInfo) SetIsAccountant(v bool)`

SetIsAccountant sets IsAccountant field to given value.

### HasIsAccountant

`func (o *CompanyInfo) HasIsAccountant() bool`

HasIsAccountant returns a boolean if a field has been set.

### SetIsAccountantNil

`func (o *CompanyInfo) SetIsAccountantNil(b bool)`

 SetIsAccountantNil sets the value for IsAccountant to be an explicit nil

### UnsetIsAccountant
`func (o *CompanyInfo) UnsetIsAccountant()`

UnsetIsAccountant ensures that no value is present for IsAccountant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


