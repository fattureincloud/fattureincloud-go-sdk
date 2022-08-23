# GetPaymentAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullablePaymentAccount**](PaymentAccount.md) |  | [optional] 

## Methods

### NewGetPaymentAccountResponse

`func NewGetPaymentAccountResponse() *GetPaymentAccountResponse`

NewGetPaymentAccountResponse instantiates a new GetPaymentAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentAccountResponseWithDefaults

`func NewGetPaymentAccountResponseWithDefaults() *GetPaymentAccountResponse`

NewGetPaymentAccountResponseWithDefaults instantiates a new GetPaymentAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPaymentAccountResponse) GetData() PaymentAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPaymentAccountResponse) GetDataOk() (*PaymentAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPaymentAccountResponse) SetData(v PaymentAccount)`

SetData sets Data field to given value.

### HasData

`func (o *GetPaymentAccountResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GetPaymentAccountResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetPaymentAccountResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


