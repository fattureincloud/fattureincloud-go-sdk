# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Company id | [optional] 
**Name** | Pointer to **NullableString** | Company name | [optional] 
**Type** | Pointer to [**CompanyType**](CompanyType.md) |  | [optional] 
**AccessToken** | Pointer to **NullableString** | Company authentication token for this company. [Only if type&#x3D;company] | [optional] 
**ControlledCompanies** | Pointer to [**[]ControlledCompany**](ControlledCompany.md) | Company list of controlled companies [Only if type&#x3D;accountant] | [optional] 
**FicLicenseExpire** | Pointer to **NullableString** |  | [optional] 
**FicPlan** | Pointer to [**FattureInCloudPlanType**](FattureInCloudPlanType.md) |  | [optional] 
**ConnectionId** | Pointer to **NullableInt32** | Company connection id | [optional] 
**TaxCode** | Pointer to **NullableString** | Company tax code | [optional] 
**VatNumber** | Pointer to **NullableString** | Company vat number | [optional] 

## Methods

### NewCompany

`func NewCompany() *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Company) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Company) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Company) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Company) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Company) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Company) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Company) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *Company) GetType() CompanyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Company) GetTypeOk() (*CompanyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Company) SetType(v CompanyType)`

SetType sets Type field to given value.

### HasType

`func (o *Company) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessToken

`func (o *Company) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Company) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Company) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Company) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *Company) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *Company) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetControlledCompanies

`func (o *Company) GetControlledCompanies() []ControlledCompany`

GetControlledCompanies returns the ControlledCompanies field if non-nil, zero value otherwise.

### GetControlledCompaniesOk

`func (o *Company) GetControlledCompaniesOk() (*[]ControlledCompany, bool)`

GetControlledCompaniesOk returns a tuple with the ControlledCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlledCompanies

`func (o *Company) SetControlledCompanies(v []ControlledCompany)`

SetControlledCompanies sets ControlledCompanies field to given value.

### HasControlledCompanies

`func (o *Company) HasControlledCompanies() bool`

HasControlledCompanies returns a boolean if a field has been set.

### SetControlledCompaniesNil

`func (o *Company) SetControlledCompaniesNil(b bool)`

 SetControlledCompaniesNil sets the value for ControlledCompanies to be an explicit nil

### UnsetControlledCompanies
`func (o *Company) UnsetControlledCompanies()`

UnsetControlledCompanies ensures that no value is present for ControlledCompanies, not even an explicit nil
### GetFicLicenseExpire

`func (o *Company) GetFicLicenseExpire() string`

GetFicLicenseExpire returns the FicLicenseExpire field if non-nil, zero value otherwise.

### GetFicLicenseExpireOk

`func (o *Company) GetFicLicenseExpireOk() (*string, bool)`

GetFicLicenseExpireOk returns a tuple with the FicLicenseExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicLicenseExpire

`func (o *Company) SetFicLicenseExpire(v string)`

SetFicLicenseExpire sets FicLicenseExpire field to given value.

### HasFicLicenseExpire

`func (o *Company) HasFicLicenseExpire() bool`

HasFicLicenseExpire returns a boolean if a field has been set.

### SetFicLicenseExpireNil

`func (o *Company) SetFicLicenseExpireNil(b bool)`

 SetFicLicenseExpireNil sets the value for FicLicenseExpire to be an explicit nil

### UnsetFicLicenseExpire
`func (o *Company) UnsetFicLicenseExpire()`

UnsetFicLicenseExpire ensures that no value is present for FicLicenseExpire, not even an explicit nil
### GetFicPlan

`func (o *Company) GetFicPlan() FattureInCloudPlanType`

GetFicPlan returns the FicPlan field if non-nil, zero value otherwise.

### GetFicPlanOk

`func (o *Company) GetFicPlanOk() (*FattureInCloudPlanType, bool)`

GetFicPlanOk returns a tuple with the FicPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicPlan

`func (o *Company) SetFicPlan(v FattureInCloudPlanType)`

SetFicPlan sets FicPlan field to given value.

### HasFicPlan

`func (o *Company) HasFicPlan() bool`

HasFicPlan returns a boolean if a field has been set.

### GetConnectionId

`func (o *Company) GetConnectionId() int32`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Company) GetConnectionIdOk() (*int32, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Company) SetConnectionId(v int32)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Company) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *Company) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *Company) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetTaxCode

`func (o *Company) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Company) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Company) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Company) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### SetTaxCodeNil

`func (o *Company) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *Company) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetVatNumber

`func (o *Company) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Company) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Company) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Company) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *Company) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *Company) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


