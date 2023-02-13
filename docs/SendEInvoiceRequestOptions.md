# SendEInvoiceRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **NullableBool** | If set to true the e-invoice will not be sent to the SDI. | [optional] 

## Methods

### NewSendEInvoiceRequestOptions

`func NewSendEInvoiceRequestOptions() *SendEInvoiceRequestOptions`

NewSendEInvoiceRequestOptions instantiates a new SendEInvoiceRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEInvoiceRequestOptionsWithDefaults

`func NewSendEInvoiceRequestOptionsWithDefaults() *SendEInvoiceRequestOptions`

NewSendEInvoiceRequestOptionsWithDefaults instantiates a new SendEInvoiceRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *SendEInvoiceRequestOptions) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SendEInvoiceRequestOptions) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SendEInvoiceRequestOptions) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *SendEInvoiceRequestOptions) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *SendEInvoiceRequestOptions) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *SendEInvoiceRequestOptions) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


