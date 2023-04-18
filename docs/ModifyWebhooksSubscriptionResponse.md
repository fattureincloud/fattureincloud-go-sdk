# ModifyWebhooksSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**WebhooksSubscription**](WebhooksSubscription.md) |  | [optional] 
**Warnings** | Pointer to **[]string** | Webhooks registration warnings | [optional] 

## Methods

### NewModifyWebhooksSubscriptionResponse

`func NewModifyWebhooksSubscriptionResponse() *ModifyWebhooksSubscriptionResponse`

NewModifyWebhooksSubscriptionResponse instantiates a new ModifyWebhooksSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyWebhooksSubscriptionResponseWithDefaults

`func NewModifyWebhooksSubscriptionResponseWithDefaults() *ModifyWebhooksSubscriptionResponse`

NewModifyWebhooksSubscriptionResponseWithDefaults instantiates a new ModifyWebhooksSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModifyWebhooksSubscriptionResponse) GetData() WebhooksSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModifyWebhooksSubscriptionResponse) GetDataOk() (*WebhooksSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModifyWebhooksSubscriptionResponse) SetData(v WebhooksSubscription)`

SetData sets Data field to given value.

### HasData

`func (o *ModifyWebhooksSubscriptionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetWarnings

`func (o *ModifyWebhooksSubscriptionResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ModifyWebhooksSubscriptionResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ModifyWebhooksSubscriptionResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ModifyWebhooksSubscriptionResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *ModifyWebhooksSubscriptionResponse) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *ModifyWebhooksSubscriptionResponse) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


