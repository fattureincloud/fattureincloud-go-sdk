# SenderEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Sender email id | [optional] 
**Email** | Pointer to **NullableString** | Sender email address | [optional] 

## Methods

### NewSenderEmail

`func NewSenderEmail() *SenderEmail`

NewSenderEmail instantiates a new SenderEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSenderEmailWithDefaults

`func NewSenderEmailWithDefaults() *SenderEmail`

NewSenderEmailWithDefaults instantiates a new SenderEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SenderEmail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SenderEmail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SenderEmail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SenderEmail) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SenderEmail) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SenderEmail) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetEmail

`func (o *SenderEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SenderEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SenderEmail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SenderEmail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SenderEmail) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SenderEmail) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


