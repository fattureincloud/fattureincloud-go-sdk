# ControlledCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Controlled company id | [optional] 
**Name** | Pointer to **NullableString** | Controlled company id | [optional] 
**Type** | Pointer to [**CompanyType**](CompanyType.md) |  | [optional] 
**AccessToken** | Pointer to **NullableString** | Controlled company access token Only if type&#x3D;company] | [optional] 
**FicLicenseExpire** | Pointer to **NullableString** |  | [optional] 
**FicPlan** | Pointer to [**FattureInCloudPlanType**](FattureInCloudPlanType.md) |  | [optional] 
**ConnectionId** | Pointer to **NullableFloat32** | Controlled company connection id | [optional] 
**TaxCode** | Pointer to **NullableString** | Controlled company tax code | [optional] 

## Methods

### NewControlledCompany

`func NewControlledCompany() *ControlledCompany`

NewControlledCompany instantiates a new ControlledCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlledCompanyWithDefaults

`func NewControlledCompanyWithDefaults() *ControlledCompany`

NewControlledCompanyWithDefaults instantiates a new ControlledCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControlledCompany) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControlledCompany) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControlledCompany) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ControlledCompany) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ControlledCompany) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ControlledCompany) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ControlledCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControlledCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControlledCompany) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControlledCompany) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ControlledCompany) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ControlledCompany) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *ControlledCompany) GetType() CompanyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ControlledCompany) GetTypeOk() (*CompanyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ControlledCompany) SetType(v CompanyType)`

SetType sets Type field to given value.

### HasType

`func (o *ControlledCompany) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessToken

`func (o *ControlledCompany) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ControlledCompany) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ControlledCompany) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *ControlledCompany) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *ControlledCompany) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *ControlledCompany) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetFicLicenseExpire

`func (o *ControlledCompany) GetFicLicenseExpire() string`

GetFicLicenseExpire returns the FicLicenseExpire field if non-nil, zero value otherwise.

### GetFicLicenseExpireOk

`func (o *ControlledCompany) GetFicLicenseExpireOk() (*string, bool)`

GetFicLicenseExpireOk returns a tuple with the FicLicenseExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicLicenseExpire

`func (o *ControlledCompany) SetFicLicenseExpire(v string)`

SetFicLicenseExpire sets FicLicenseExpire field to given value.

### HasFicLicenseExpire

`func (o *ControlledCompany) HasFicLicenseExpire() bool`

HasFicLicenseExpire returns a boolean if a field has been set.

### SetFicLicenseExpireNil

`func (o *ControlledCompany) SetFicLicenseExpireNil(b bool)`

 SetFicLicenseExpireNil sets the value for FicLicenseExpire to be an explicit nil

### UnsetFicLicenseExpire
`func (o *ControlledCompany) UnsetFicLicenseExpire()`

UnsetFicLicenseExpire ensures that no value is present for FicLicenseExpire, not even an explicit nil
### GetFicPlan

`func (o *ControlledCompany) GetFicPlan() FattureInCloudPlanType`

GetFicPlan returns the FicPlan field if non-nil, zero value otherwise.

### GetFicPlanOk

`func (o *ControlledCompany) GetFicPlanOk() (*FattureInCloudPlanType, bool)`

GetFicPlanOk returns a tuple with the FicPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicPlan

`func (o *ControlledCompany) SetFicPlan(v FattureInCloudPlanType)`

SetFicPlan sets FicPlan field to given value.

### HasFicPlan

`func (o *ControlledCompany) HasFicPlan() bool`

HasFicPlan returns a boolean if a field has been set.

### GetConnectionId

`func (o *ControlledCompany) GetConnectionId() float32`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ControlledCompany) GetConnectionIdOk() (*float32, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ControlledCompany) SetConnectionId(v float32)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ControlledCompany) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *ControlledCompany) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *ControlledCompany) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetTaxCode

`func (o *ControlledCompany) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ControlledCompany) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ControlledCompany) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *ControlledCompany) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### SetTaxCodeNil

`func (o *ControlledCompany) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *ControlledCompany) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


