# WebhooksSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Unique identifier | [optional] 
**Sink** | Pointer to **NullableString** | Webhooks callback uri. | [optional] 
**Verified** | Pointer to **NullableBool** | [Read Only] True if the webhooks subscription has been verified. | [optional] 
**Types** | Pointer to [**[]EventType**](EventType.md) | Webhooks events types. | [optional] 

## Methods

### NewWebhooksSubscription

`func NewWebhooksSubscription() *WebhooksSubscription`

NewWebhooksSubscription instantiates a new WebhooksSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksSubscriptionWithDefaults

`func NewWebhooksSubscriptionWithDefaults() *WebhooksSubscription`

NewWebhooksSubscriptionWithDefaults instantiates a new WebhooksSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhooksSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhooksSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhooksSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhooksSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WebhooksSubscription) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WebhooksSubscription) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSink

`func (o *WebhooksSubscription) GetSink() string`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *WebhooksSubscription) GetSinkOk() (*string, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *WebhooksSubscription) SetSink(v string)`

SetSink sets Sink field to given value.

### HasSink

`func (o *WebhooksSubscription) HasSink() bool`

HasSink returns a boolean if a field has been set.

### SetSinkNil

`func (o *WebhooksSubscription) SetSinkNil(b bool)`

 SetSinkNil sets the value for Sink to be an explicit nil

### UnsetSink
`func (o *WebhooksSubscription) UnsetSink()`

UnsetSink ensures that no value is present for Sink, not even an explicit nil
### GetVerified

`func (o *WebhooksSubscription) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *WebhooksSubscription) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *WebhooksSubscription) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *WebhooksSubscription) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### SetVerifiedNil

`func (o *WebhooksSubscription) SetVerifiedNil(b bool)`

 SetVerifiedNil sets the value for Verified to be an explicit nil

### UnsetVerified
`func (o *WebhooksSubscription) UnsetVerified()`

UnsetVerified ensures that no value is present for Verified, not even an explicit nil
### GetTypes

`func (o *WebhooksSubscription) GetTypes() []EventType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *WebhooksSubscription) GetTypesOk() (*[]EventType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *WebhooksSubscription) SetTypes(v []EventType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *WebhooksSubscription) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *WebhooksSubscription) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *WebhooksSubscription) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


