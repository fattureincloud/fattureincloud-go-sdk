# GetPaymentMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PaymentMethod**](PaymentMethod.md) |  | [optional] 

## Methods

### NewGetPaymentMethodResponse

`func NewGetPaymentMethodResponse() *GetPaymentMethodResponse`

NewGetPaymentMethodResponse instantiates a new GetPaymentMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentMethodResponseWithDefaults

`func NewGetPaymentMethodResponseWithDefaults() *GetPaymentMethodResponse`

NewGetPaymentMethodResponseWithDefaults instantiates a new GetPaymentMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPaymentMethodResponse) GetData() PaymentMethod`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPaymentMethodResponse) GetDataOk() (*PaymentMethod, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPaymentMethodResponse) SetData(v PaymentMethod)`

SetData sets Data field to given value.

### HasData

`func (o *GetPaymentMethodResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


