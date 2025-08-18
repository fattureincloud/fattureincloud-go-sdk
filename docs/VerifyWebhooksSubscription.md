# VerifyWebhooksSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Webhooks subscription id | [optional] 
**VerificationMethod** | Pointer to [**WebhooksSubscriptionVerificationMethod**](WebhooksSubscriptionVerificationMethod.md) |  | [optional] 

## Methods

### NewVerifyWebhooksSubscription

`func NewVerifyWebhooksSubscription() *VerifyWebhooksSubscription`

NewVerifyWebhooksSubscription instantiates a new VerifyWebhooksSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyWebhooksSubscriptionWithDefaults

`func NewVerifyWebhooksSubscriptionWithDefaults() *VerifyWebhooksSubscription`

NewVerifyWebhooksSubscriptionWithDefaults instantiates a new VerifyWebhooksSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerifyWebhooksSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifyWebhooksSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifyWebhooksSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerifyWebhooksSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VerifyWebhooksSubscription) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VerifyWebhooksSubscription) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetVerificationMethod

`func (o *VerifyWebhooksSubscription) GetVerificationMethod() WebhooksSubscriptionVerificationMethod`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *VerifyWebhooksSubscription) GetVerificationMethodOk() (*WebhooksSubscriptionVerificationMethod, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *VerifyWebhooksSubscription) SetVerificationMethod(v WebhooksSubscriptionVerificationMethod)`

SetVerificationMethod sets VerificationMethod field to given value.

### HasVerificationMethod

`func (o *VerifyWebhooksSubscription) HasVerificationMethod() bool`

HasVerificationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


