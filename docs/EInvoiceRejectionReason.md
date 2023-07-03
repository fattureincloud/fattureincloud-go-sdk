# EInvoiceRejectionReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **NullableString** | E-invoice rejection reason | [optional] 
**EiStatus** | Pointer to **NullableString** | E-invoice status | [optional] 
**Solution** | Pointer to **NullableString** | Error solution. | [optional] 
**Code** | Pointer to **NullableString** | E-invoice rejection error code | [optional] 
**Date** | Pointer to **NullableTime** | E-invoice rejection date | [optional] 

## Methods

### NewEInvoiceRejectionReason

`func NewEInvoiceRejectionReason() *EInvoiceRejectionReason`

NewEInvoiceRejectionReason instantiates a new EInvoiceRejectionReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEInvoiceRejectionReasonWithDefaults

`func NewEInvoiceRejectionReasonWithDefaults() *EInvoiceRejectionReason`

NewEInvoiceRejectionReasonWithDefaults instantiates a new EInvoiceRejectionReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *EInvoiceRejectionReason) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EInvoiceRejectionReason) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EInvoiceRejectionReason) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EInvoiceRejectionReason) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *EInvoiceRejectionReason) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *EInvoiceRejectionReason) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetEiStatus

`func (o *EInvoiceRejectionReason) GetEiStatus() string`

GetEiStatus returns the EiStatus field if non-nil, zero value otherwise.

### GetEiStatusOk

`func (o *EInvoiceRejectionReason) GetEiStatusOk() (*string, bool)`

GetEiStatusOk returns a tuple with the EiStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiStatus

`func (o *EInvoiceRejectionReason) SetEiStatus(v string)`

SetEiStatus sets EiStatus field to given value.

### HasEiStatus

`func (o *EInvoiceRejectionReason) HasEiStatus() bool`

HasEiStatus returns a boolean if a field has been set.

### SetEiStatusNil

`func (o *EInvoiceRejectionReason) SetEiStatusNil(b bool)`

 SetEiStatusNil sets the value for EiStatus to be an explicit nil

### UnsetEiStatus
`func (o *EInvoiceRejectionReason) UnsetEiStatus()`

UnsetEiStatus ensures that no value is present for EiStatus, not even an explicit nil
### GetSolution

`func (o *EInvoiceRejectionReason) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *EInvoiceRejectionReason) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *EInvoiceRejectionReason) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *EInvoiceRejectionReason) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### SetSolutionNil

`func (o *EInvoiceRejectionReason) SetSolutionNil(b bool)`

 SetSolutionNil sets the value for Solution to be an explicit nil

### UnsetSolution
`func (o *EInvoiceRejectionReason) UnsetSolution()`

UnsetSolution ensures that no value is present for Solution, not even an explicit nil
### GetCode

`func (o *EInvoiceRejectionReason) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EInvoiceRejectionReason) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EInvoiceRejectionReason) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EInvoiceRejectionReason) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *EInvoiceRejectionReason) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *EInvoiceRejectionReason) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDate

`func (o *EInvoiceRejectionReason) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EInvoiceRejectionReason) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EInvoiceRejectionReason) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *EInvoiceRejectionReason) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *EInvoiceRejectionReason) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *EInvoiceRejectionReason) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


