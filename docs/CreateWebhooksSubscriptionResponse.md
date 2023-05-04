# CreateWebhooksSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**WebhooksSubscription**](WebhooksSubscription.md) |  | [optional] 
**Warnings** | Pointer to **[]string** | Webhooks registration warnings | [optional] 

## Methods

### NewCreateWebhooksSubscriptionResponse

`func NewCreateWebhooksSubscriptionResponse() *CreateWebhooksSubscriptionResponse`

NewCreateWebhooksSubscriptionResponse instantiates a new CreateWebhooksSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhooksSubscriptionResponseWithDefaults

`func NewCreateWebhooksSubscriptionResponseWithDefaults() *CreateWebhooksSubscriptionResponse`

NewCreateWebhooksSubscriptionResponseWithDefaults instantiates a new CreateWebhooksSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateWebhooksSubscriptionResponse) GetData() WebhooksSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateWebhooksSubscriptionResponse) GetDataOk() (*WebhooksSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateWebhooksSubscriptionResponse) SetData(v WebhooksSubscription)`

SetData sets Data field to given value.

### HasData

`func (o *CreateWebhooksSubscriptionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetWarnings

`func (o *CreateWebhooksSubscriptionResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CreateWebhooksSubscriptionResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CreateWebhooksSubscriptionResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CreateWebhooksSubscriptionResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *CreateWebhooksSubscriptionResponse) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *CreateWebhooksSubscriptionResponse) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


