# City

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostalCode** | Pointer to **NullableString** | City postal code. | [optional] 
**City** | Pointer to **NullableString** | City name. | [optional] 
**Province** | Pointer to **NullableString** | Province. | [optional] 

## Methods

### NewCity

`func NewCity() *City`

NewCity instantiates a new City object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCityWithDefaults

`func NewCityWithDefaults() *City`

NewCityWithDefaults instantiates a new City object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostalCode

`func (o *City) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *City) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *City) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *City) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *City) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *City) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCity

`func (o *City) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *City) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *City) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *City) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *City) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *City) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetProvince

`func (o *City) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *City) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *City) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *City) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### SetProvinceNil

`func (o *City) SetProvinceNil(b bool)`

 SetProvinceNil sets the value for Province to be an explicit nil

### UnsetProvince
`func (o *City) UnsetProvince()`

UnsetProvince ensures that no value is present for Province, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


