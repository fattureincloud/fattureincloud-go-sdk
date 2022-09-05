# SendEInvoiceResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Response message. | [optional] 
**Date** | Pointer to **NullableString** | E-invoice sent date. | [optional] 

## Methods

### NewSendEInvoiceResponseData

`func NewSendEInvoiceResponseData() *SendEInvoiceResponseData`

NewSendEInvoiceResponseData instantiates a new SendEInvoiceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEInvoiceResponseDataWithDefaults

`func NewSendEInvoiceResponseDataWithDefaults() *SendEInvoiceResponseData`

NewSendEInvoiceResponseDataWithDefaults instantiates a new SendEInvoiceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SendEInvoiceResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendEInvoiceResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendEInvoiceResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SendEInvoiceResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SendEInvoiceResponseData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SendEInvoiceResponseData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDate

`func (o *SendEInvoiceResponseData) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SendEInvoiceResponseData) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SendEInvoiceResponseData) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *SendEInvoiceResponseData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *SendEInvoiceResponseData) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *SendEInvoiceResponseData) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


