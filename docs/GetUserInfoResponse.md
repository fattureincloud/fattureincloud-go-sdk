# GetUserInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**User**](User.md) |  | [optional] 
**Info** | Pointer to [**NullableGetUserInfoResponseInfo**](GetUserInfoResponseInfo.md) |  | [optional] 
**EmailConfirmationState** | Pointer to [**NullableGetUserInfoResponseEmailConfirmationState**](GetUserInfoResponseEmailConfirmationState.md) |  | [optional] 

## Methods

### NewGetUserInfoResponse

`func NewGetUserInfoResponse() *GetUserInfoResponse`

NewGetUserInfoResponse instantiates a new GetUserInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserInfoResponseWithDefaults

`func NewGetUserInfoResponseWithDefaults() *GetUserInfoResponse`

NewGetUserInfoResponseWithDefaults instantiates a new GetUserInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUserInfoResponse) GetData() User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserInfoResponse) GetDataOk() (*User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserInfoResponse) SetData(v User)`

SetData sets Data field to given value.

### HasData

`func (o *GetUserInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetInfo

`func (o *GetUserInfoResponse) GetInfo() GetUserInfoResponseInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GetUserInfoResponse) GetInfoOk() (*GetUserInfoResponseInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GetUserInfoResponse) SetInfo(v GetUserInfoResponseInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GetUserInfoResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *GetUserInfoResponse) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *GetUserInfoResponse) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetEmailConfirmationState

`func (o *GetUserInfoResponse) GetEmailConfirmationState() GetUserInfoResponseEmailConfirmationState`

GetEmailConfirmationState returns the EmailConfirmationState field if non-nil, zero value otherwise.

### GetEmailConfirmationStateOk

`func (o *GetUserInfoResponse) GetEmailConfirmationStateOk() (*GetUserInfoResponseEmailConfirmationState, bool)`

GetEmailConfirmationStateOk returns a tuple with the EmailConfirmationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmationState

`func (o *GetUserInfoResponse) SetEmailConfirmationState(v GetUserInfoResponseEmailConfirmationState)`

SetEmailConfirmationState sets EmailConfirmationState field to given value.

### HasEmailConfirmationState

`func (o *GetUserInfoResponse) HasEmailConfirmationState() bool`

HasEmailConfirmationState returns a boolean if a field has been set.

### SetEmailConfirmationStateNil

`func (o *GetUserInfoResponse) SetEmailConfirmationStateNil(b bool)`

 SetEmailConfirmationStateNil sets the value for EmailConfirmationState to be an explicit nil

### UnsetEmailConfirmationState
`func (o *GetUserInfoResponse) UnsetEmailConfirmationState()`

UnsetEmailConfirmationState ensures that no value is present for EmailConfirmationState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


