# SendEInvoiceRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassaType** | Pointer to **NullableString** | Value of TipoCassa used (optional, override the company default value). | [optional] 
**WithholdingTaxCausal** | Pointer to **NullableString** | Value of CausalePagamento used (optional, override the company default value). | [optional] 

## Methods

### NewSendEInvoiceRequestData

`func NewSendEInvoiceRequestData() *SendEInvoiceRequestData`

NewSendEInvoiceRequestData instantiates a new SendEInvoiceRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEInvoiceRequestDataWithDefaults

`func NewSendEInvoiceRequestDataWithDefaults() *SendEInvoiceRequestData`

NewSendEInvoiceRequestDataWithDefaults instantiates a new SendEInvoiceRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassaType

`func (o *SendEInvoiceRequestData) GetCassaType() string`

GetCassaType returns the CassaType field if non-nil, zero value otherwise.

### GetCassaTypeOk

`func (o *SendEInvoiceRequestData) GetCassaTypeOk() (*string, bool)`

GetCassaTypeOk returns a tuple with the CassaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassaType

`func (o *SendEInvoiceRequestData) SetCassaType(v string)`

SetCassaType sets CassaType field to given value.

### HasCassaType

`func (o *SendEInvoiceRequestData) HasCassaType() bool`

HasCassaType returns a boolean if a field has been set.

### SetCassaTypeNil

`func (o *SendEInvoiceRequestData) SetCassaTypeNil(b bool)`

 SetCassaTypeNil sets the value for CassaType to be an explicit nil

### UnsetCassaType
`func (o *SendEInvoiceRequestData) UnsetCassaType()`

UnsetCassaType ensures that no value is present for CassaType, not even an explicit nil
### GetWithholdingTaxCausal

`func (o *SendEInvoiceRequestData) GetWithholdingTaxCausal() string`

GetWithholdingTaxCausal returns the WithholdingTaxCausal field if non-nil, zero value otherwise.

### GetWithholdingTaxCausalOk

`func (o *SendEInvoiceRequestData) GetWithholdingTaxCausalOk() (*string, bool)`

GetWithholdingTaxCausalOk returns a tuple with the WithholdingTaxCausal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholdingTaxCausal

`func (o *SendEInvoiceRequestData) SetWithholdingTaxCausal(v string)`

SetWithholdingTaxCausal sets WithholdingTaxCausal field to given value.

### HasWithholdingTaxCausal

`func (o *SendEInvoiceRequestData) HasWithholdingTaxCausal() bool`

HasWithholdingTaxCausal returns a boolean if a field has been set.

### SetWithholdingTaxCausalNil

`func (o *SendEInvoiceRequestData) SetWithholdingTaxCausalNil(b bool)`

 SetWithholdingTaxCausalNil sets the value for WithholdingTaxCausal to be an explicit nil

### UnsetWithholdingTaxCausal
`func (o *SendEInvoiceRequestData) UnsetWithholdingTaxCausal()`

UnsetWithholdingTaxCausal ensures that no value is present for WithholdingTaxCausal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


