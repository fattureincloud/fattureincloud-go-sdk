# TaxProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyType** | Pointer to **NullableString** | The company type | [optional] 
**CompanySubtype** | Pointer to **NullableString** | The company subtype | [optional] 
**Profession** | Pointer to **NullableString** | The profession | [optional] 
**Regime** | Pointer to **NullableString** | The applied regime | [optional] 
**RivalsaName** | Pointer to **NullableString** | The name of the rivalsa | [optional] 
**DefaultRivalsa** | Pointer to **NullableFloat32** | The default rivalsa amount | [optional] 
**CassaName** | Pointer to **NullableString** | The name of the cassa | [optional] 
**DefaultCassa** | Pointer to **NullableFloat32** | The default cassa amount | [optional] 
**DefaultCassaTaxable** | Pointer to **NullableFloat32** | The default taxable amount for the cassa | [optional] 
**Cassa2Name** | Pointer to **NullableString** | The name of the second cassa | [optional] 
**DefaultCassa2** | Pointer to **NullableFloat32** | The default second cassa amount | [optional] 
**DefaultCassa2Taxable** | Pointer to **NullableFloat32** | The default taxable amount for the second cassa | [optional] 
**DefaultWithholdingTax** | Pointer to **NullableFloat32** | The default withholding tax | [optional] 
**DefaultWithholdingTaxTaxable** | Pointer to **NullableFloat32** | The default taxable amount for the withholding tax | [optional] 
**DefaultOtherWithholdingTax** | Pointer to **NullableFloat32** | The default other withholding tax | [optional] 
**Enasarco** | Pointer to **NullableBool** | If it has enasarco | [optional] 
**EnasarcoType** | Pointer to **NullableString** | The enasarco type | [optional] 
**ContributionsPercentage** | Pointer to **NullableFloat32** | The contributions percentage | [optional] 
**ProfitCoefficient** | Pointer to **NullableFloat32** | The profit coefficient | [optional] 
**Med** | Pointer to **NullableBool** | If the health card system is active | [optional] 
**DefaultVat** | Pointer to [**NullableVatType**](VatType.md) |  | [optional] 

## Methods

### NewTaxProfile

`func NewTaxProfile() *TaxProfile`

NewTaxProfile instantiates a new TaxProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxProfileWithDefaults

`func NewTaxProfileWithDefaults() *TaxProfile`

NewTaxProfileWithDefaults instantiates a new TaxProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyType

`func (o *TaxProfile) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *TaxProfile) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *TaxProfile) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *TaxProfile) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### SetCompanyTypeNil

`func (o *TaxProfile) SetCompanyTypeNil(b bool)`

 SetCompanyTypeNil sets the value for CompanyType to be an explicit nil

### UnsetCompanyType
`func (o *TaxProfile) UnsetCompanyType()`

UnsetCompanyType ensures that no value is present for CompanyType, not even an explicit nil
### GetCompanySubtype

`func (o *TaxProfile) GetCompanySubtype() string`

GetCompanySubtype returns the CompanySubtype field if non-nil, zero value otherwise.

### GetCompanySubtypeOk

`func (o *TaxProfile) GetCompanySubtypeOk() (*string, bool)`

GetCompanySubtypeOk returns a tuple with the CompanySubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySubtype

`func (o *TaxProfile) SetCompanySubtype(v string)`

SetCompanySubtype sets CompanySubtype field to given value.

### HasCompanySubtype

`func (o *TaxProfile) HasCompanySubtype() bool`

HasCompanySubtype returns a boolean if a field has been set.

### SetCompanySubtypeNil

`func (o *TaxProfile) SetCompanySubtypeNil(b bool)`

 SetCompanySubtypeNil sets the value for CompanySubtype to be an explicit nil

### UnsetCompanySubtype
`func (o *TaxProfile) UnsetCompanySubtype()`

UnsetCompanySubtype ensures that no value is present for CompanySubtype, not even an explicit nil
### GetProfession

`func (o *TaxProfile) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *TaxProfile) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *TaxProfile) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *TaxProfile) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### SetProfessionNil

`func (o *TaxProfile) SetProfessionNil(b bool)`

 SetProfessionNil sets the value for Profession to be an explicit nil

### UnsetProfession
`func (o *TaxProfile) UnsetProfession()`

UnsetProfession ensures that no value is present for Profession, not even an explicit nil
### GetRegime

`func (o *TaxProfile) GetRegime() string`

GetRegime returns the Regime field if non-nil, zero value otherwise.

### GetRegimeOk

`func (o *TaxProfile) GetRegimeOk() (*string, bool)`

GetRegimeOk returns a tuple with the Regime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegime

`func (o *TaxProfile) SetRegime(v string)`

SetRegime sets Regime field to given value.

### HasRegime

`func (o *TaxProfile) HasRegime() bool`

HasRegime returns a boolean if a field has been set.

### SetRegimeNil

`func (o *TaxProfile) SetRegimeNil(b bool)`

 SetRegimeNil sets the value for Regime to be an explicit nil

### UnsetRegime
`func (o *TaxProfile) UnsetRegime()`

UnsetRegime ensures that no value is present for Regime, not even an explicit nil
### GetRivalsaName

`func (o *TaxProfile) GetRivalsaName() string`

GetRivalsaName returns the RivalsaName field if non-nil, zero value otherwise.

### GetRivalsaNameOk

`func (o *TaxProfile) GetRivalsaNameOk() (*string, bool)`

GetRivalsaNameOk returns a tuple with the RivalsaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRivalsaName

`func (o *TaxProfile) SetRivalsaName(v string)`

SetRivalsaName sets RivalsaName field to given value.

### HasRivalsaName

`func (o *TaxProfile) HasRivalsaName() bool`

HasRivalsaName returns a boolean if a field has been set.

### SetRivalsaNameNil

`func (o *TaxProfile) SetRivalsaNameNil(b bool)`

 SetRivalsaNameNil sets the value for RivalsaName to be an explicit nil

### UnsetRivalsaName
`func (o *TaxProfile) UnsetRivalsaName()`

UnsetRivalsaName ensures that no value is present for RivalsaName, not even an explicit nil
### GetDefaultRivalsa

`func (o *TaxProfile) GetDefaultRivalsa() float32`

GetDefaultRivalsa returns the DefaultRivalsa field if non-nil, zero value otherwise.

### GetDefaultRivalsaOk

`func (o *TaxProfile) GetDefaultRivalsaOk() (*float32, bool)`

GetDefaultRivalsaOk returns a tuple with the DefaultRivalsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRivalsa

`func (o *TaxProfile) SetDefaultRivalsa(v float32)`

SetDefaultRivalsa sets DefaultRivalsa field to given value.

### HasDefaultRivalsa

`func (o *TaxProfile) HasDefaultRivalsa() bool`

HasDefaultRivalsa returns a boolean if a field has been set.

### SetDefaultRivalsaNil

`func (o *TaxProfile) SetDefaultRivalsaNil(b bool)`

 SetDefaultRivalsaNil sets the value for DefaultRivalsa to be an explicit nil

### UnsetDefaultRivalsa
`func (o *TaxProfile) UnsetDefaultRivalsa()`

UnsetDefaultRivalsa ensures that no value is present for DefaultRivalsa, not even an explicit nil
### GetCassaName

`func (o *TaxProfile) GetCassaName() string`

GetCassaName returns the CassaName field if non-nil, zero value otherwise.

### GetCassaNameOk

`func (o *TaxProfile) GetCassaNameOk() (*string, bool)`

GetCassaNameOk returns a tuple with the CassaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassaName

`func (o *TaxProfile) SetCassaName(v string)`

SetCassaName sets CassaName field to given value.

### HasCassaName

`func (o *TaxProfile) HasCassaName() bool`

HasCassaName returns a boolean if a field has been set.

### SetCassaNameNil

`func (o *TaxProfile) SetCassaNameNil(b bool)`

 SetCassaNameNil sets the value for CassaName to be an explicit nil

### UnsetCassaName
`func (o *TaxProfile) UnsetCassaName()`

UnsetCassaName ensures that no value is present for CassaName, not even an explicit nil
### GetDefaultCassa

`func (o *TaxProfile) GetDefaultCassa() float32`

GetDefaultCassa returns the DefaultCassa field if non-nil, zero value otherwise.

### GetDefaultCassaOk

`func (o *TaxProfile) GetDefaultCassaOk() (*float32, bool)`

GetDefaultCassaOk returns a tuple with the DefaultCassa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCassa

`func (o *TaxProfile) SetDefaultCassa(v float32)`

SetDefaultCassa sets DefaultCassa field to given value.

### HasDefaultCassa

`func (o *TaxProfile) HasDefaultCassa() bool`

HasDefaultCassa returns a boolean if a field has been set.

### SetDefaultCassaNil

`func (o *TaxProfile) SetDefaultCassaNil(b bool)`

 SetDefaultCassaNil sets the value for DefaultCassa to be an explicit nil

### UnsetDefaultCassa
`func (o *TaxProfile) UnsetDefaultCassa()`

UnsetDefaultCassa ensures that no value is present for DefaultCassa, not even an explicit nil
### GetDefaultCassaTaxable

`func (o *TaxProfile) GetDefaultCassaTaxable() float32`

GetDefaultCassaTaxable returns the DefaultCassaTaxable field if non-nil, zero value otherwise.

### GetDefaultCassaTaxableOk

`func (o *TaxProfile) GetDefaultCassaTaxableOk() (*float32, bool)`

GetDefaultCassaTaxableOk returns a tuple with the DefaultCassaTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCassaTaxable

`func (o *TaxProfile) SetDefaultCassaTaxable(v float32)`

SetDefaultCassaTaxable sets DefaultCassaTaxable field to given value.

### HasDefaultCassaTaxable

`func (o *TaxProfile) HasDefaultCassaTaxable() bool`

HasDefaultCassaTaxable returns a boolean if a field has been set.

### SetDefaultCassaTaxableNil

`func (o *TaxProfile) SetDefaultCassaTaxableNil(b bool)`

 SetDefaultCassaTaxableNil sets the value for DefaultCassaTaxable to be an explicit nil

### UnsetDefaultCassaTaxable
`func (o *TaxProfile) UnsetDefaultCassaTaxable()`

UnsetDefaultCassaTaxable ensures that no value is present for DefaultCassaTaxable, not even an explicit nil
### GetCassa2Name

`func (o *TaxProfile) GetCassa2Name() string`

GetCassa2Name returns the Cassa2Name field if non-nil, zero value otherwise.

### GetCassa2NameOk

`func (o *TaxProfile) GetCassa2NameOk() (*string, bool)`

GetCassa2NameOk returns a tuple with the Cassa2Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassa2Name

`func (o *TaxProfile) SetCassa2Name(v string)`

SetCassa2Name sets Cassa2Name field to given value.

### HasCassa2Name

`func (o *TaxProfile) HasCassa2Name() bool`

HasCassa2Name returns a boolean if a field has been set.

### SetCassa2NameNil

`func (o *TaxProfile) SetCassa2NameNil(b bool)`

 SetCassa2NameNil sets the value for Cassa2Name to be an explicit nil

### UnsetCassa2Name
`func (o *TaxProfile) UnsetCassa2Name()`

UnsetCassa2Name ensures that no value is present for Cassa2Name, not even an explicit nil
### GetDefaultCassa2

`func (o *TaxProfile) GetDefaultCassa2() float32`

GetDefaultCassa2 returns the DefaultCassa2 field if non-nil, zero value otherwise.

### GetDefaultCassa2Ok

`func (o *TaxProfile) GetDefaultCassa2Ok() (*float32, bool)`

GetDefaultCassa2Ok returns a tuple with the DefaultCassa2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCassa2

`func (o *TaxProfile) SetDefaultCassa2(v float32)`

SetDefaultCassa2 sets DefaultCassa2 field to given value.

### HasDefaultCassa2

`func (o *TaxProfile) HasDefaultCassa2() bool`

HasDefaultCassa2 returns a boolean if a field has been set.

### SetDefaultCassa2Nil

`func (o *TaxProfile) SetDefaultCassa2Nil(b bool)`

 SetDefaultCassa2Nil sets the value for DefaultCassa2 to be an explicit nil

### UnsetDefaultCassa2
`func (o *TaxProfile) UnsetDefaultCassa2()`

UnsetDefaultCassa2 ensures that no value is present for DefaultCassa2, not even an explicit nil
### GetDefaultCassa2Taxable

`func (o *TaxProfile) GetDefaultCassa2Taxable() float32`

GetDefaultCassa2Taxable returns the DefaultCassa2Taxable field if non-nil, zero value otherwise.

### GetDefaultCassa2TaxableOk

`func (o *TaxProfile) GetDefaultCassa2TaxableOk() (*float32, bool)`

GetDefaultCassa2TaxableOk returns a tuple with the DefaultCassa2Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCassa2Taxable

`func (o *TaxProfile) SetDefaultCassa2Taxable(v float32)`

SetDefaultCassa2Taxable sets DefaultCassa2Taxable field to given value.

### HasDefaultCassa2Taxable

`func (o *TaxProfile) HasDefaultCassa2Taxable() bool`

HasDefaultCassa2Taxable returns a boolean if a field has been set.

### SetDefaultCassa2TaxableNil

`func (o *TaxProfile) SetDefaultCassa2TaxableNil(b bool)`

 SetDefaultCassa2TaxableNil sets the value for DefaultCassa2Taxable to be an explicit nil

### UnsetDefaultCassa2Taxable
`func (o *TaxProfile) UnsetDefaultCassa2Taxable()`

UnsetDefaultCassa2Taxable ensures that no value is present for DefaultCassa2Taxable, not even an explicit nil
### GetDefaultWithholdingTax

`func (o *TaxProfile) GetDefaultWithholdingTax() float32`

GetDefaultWithholdingTax returns the DefaultWithholdingTax field if non-nil, zero value otherwise.

### GetDefaultWithholdingTaxOk

`func (o *TaxProfile) GetDefaultWithholdingTaxOk() (*float32, bool)`

GetDefaultWithholdingTaxOk returns a tuple with the DefaultWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWithholdingTax

`func (o *TaxProfile) SetDefaultWithholdingTax(v float32)`

SetDefaultWithholdingTax sets DefaultWithholdingTax field to given value.

### HasDefaultWithholdingTax

`func (o *TaxProfile) HasDefaultWithholdingTax() bool`

HasDefaultWithholdingTax returns a boolean if a field has been set.

### SetDefaultWithholdingTaxNil

`func (o *TaxProfile) SetDefaultWithholdingTaxNil(b bool)`

 SetDefaultWithholdingTaxNil sets the value for DefaultWithholdingTax to be an explicit nil

### UnsetDefaultWithholdingTax
`func (o *TaxProfile) UnsetDefaultWithholdingTax()`

UnsetDefaultWithholdingTax ensures that no value is present for DefaultWithholdingTax, not even an explicit nil
### GetDefaultWithholdingTaxTaxable

`func (o *TaxProfile) GetDefaultWithholdingTaxTaxable() float32`

GetDefaultWithholdingTaxTaxable returns the DefaultWithholdingTaxTaxable field if non-nil, zero value otherwise.

### GetDefaultWithholdingTaxTaxableOk

`func (o *TaxProfile) GetDefaultWithholdingTaxTaxableOk() (*float32, bool)`

GetDefaultWithholdingTaxTaxableOk returns a tuple with the DefaultWithholdingTaxTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWithholdingTaxTaxable

`func (o *TaxProfile) SetDefaultWithholdingTaxTaxable(v float32)`

SetDefaultWithholdingTaxTaxable sets DefaultWithholdingTaxTaxable field to given value.

### HasDefaultWithholdingTaxTaxable

`func (o *TaxProfile) HasDefaultWithholdingTaxTaxable() bool`

HasDefaultWithholdingTaxTaxable returns a boolean if a field has been set.

### SetDefaultWithholdingTaxTaxableNil

`func (o *TaxProfile) SetDefaultWithholdingTaxTaxableNil(b bool)`

 SetDefaultWithholdingTaxTaxableNil sets the value for DefaultWithholdingTaxTaxable to be an explicit nil

### UnsetDefaultWithholdingTaxTaxable
`func (o *TaxProfile) UnsetDefaultWithholdingTaxTaxable()`

UnsetDefaultWithholdingTaxTaxable ensures that no value is present for DefaultWithholdingTaxTaxable, not even an explicit nil
### GetDefaultOtherWithholdingTax

`func (o *TaxProfile) GetDefaultOtherWithholdingTax() float32`

GetDefaultOtherWithholdingTax returns the DefaultOtherWithholdingTax field if non-nil, zero value otherwise.

### GetDefaultOtherWithholdingTaxOk

`func (o *TaxProfile) GetDefaultOtherWithholdingTaxOk() (*float32, bool)`

GetDefaultOtherWithholdingTaxOk returns a tuple with the DefaultOtherWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOtherWithholdingTax

`func (o *TaxProfile) SetDefaultOtherWithholdingTax(v float32)`

SetDefaultOtherWithholdingTax sets DefaultOtherWithholdingTax field to given value.

### HasDefaultOtherWithholdingTax

`func (o *TaxProfile) HasDefaultOtherWithholdingTax() bool`

HasDefaultOtherWithholdingTax returns a boolean if a field has been set.

### SetDefaultOtherWithholdingTaxNil

`func (o *TaxProfile) SetDefaultOtherWithholdingTaxNil(b bool)`

 SetDefaultOtherWithholdingTaxNil sets the value for DefaultOtherWithholdingTax to be an explicit nil

### UnsetDefaultOtherWithholdingTax
`func (o *TaxProfile) UnsetDefaultOtherWithholdingTax()`

UnsetDefaultOtherWithholdingTax ensures that no value is present for DefaultOtherWithholdingTax, not even an explicit nil
### GetEnasarco

`func (o *TaxProfile) GetEnasarco() bool`

GetEnasarco returns the Enasarco field if non-nil, zero value otherwise.

### GetEnasarcoOk

`func (o *TaxProfile) GetEnasarcoOk() (*bool, bool)`

GetEnasarcoOk returns a tuple with the Enasarco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnasarco

`func (o *TaxProfile) SetEnasarco(v bool)`

SetEnasarco sets Enasarco field to given value.

### HasEnasarco

`func (o *TaxProfile) HasEnasarco() bool`

HasEnasarco returns a boolean if a field has been set.

### SetEnasarcoNil

`func (o *TaxProfile) SetEnasarcoNil(b bool)`

 SetEnasarcoNil sets the value for Enasarco to be an explicit nil

### UnsetEnasarco
`func (o *TaxProfile) UnsetEnasarco()`

UnsetEnasarco ensures that no value is present for Enasarco, not even an explicit nil
### GetEnasarcoType

`func (o *TaxProfile) GetEnasarcoType() string`

GetEnasarcoType returns the EnasarcoType field if non-nil, zero value otherwise.

### GetEnasarcoTypeOk

`func (o *TaxProfile) GetEnasarcoTypeOk() (*string, bool)`

GetEnasarcoTypeOk returns a tuple with the EnasarcoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnasarcoType

`func (o *TaxProfile) SetEnasarcoType(v string)`

SetEnasarcoType sets EnasarcoType field to given value.

### HasEnasarcoType

`func (o *TaxProfile) HasEnasarcoType() bool`

HasEnasarcoType returns a boolean if a field has been set.

### SetEnasarcoTypeNil

`func (o *TaxProfile) SetEnasarcoTypeNil(b bool)`

 SetEnasarcoTypeNil sets the value for EnasarcoType to be an explicit nil

### UnsetEnasarcoType
`func (o *TaxProfile) UnsetEnasarcoType()`

UnsetEnasarcoType ensures that no value is present for EnasarcoType, not even an explicit nil
### GetContributionsPercentage

`func (o *TaxProfile) GetContributionsPercentage() float32`

GetContributionsPercentage returns the ContributionsPercentage field if non-nil, zero value otherwise.

### GetContributionsPercentageOk

`func (o *TaxProfile) GetContributionsPercentageOk() (*float32, bool)`

GetContributionsPercentageOk returns a tuple with the ContributionsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributionsPercentage

`func (o *TaxProfile) SetContributionsPercentage(v float32)`

SetContributionsPercentage sets ContributionsPercentage field to given value.

### HasContributionsPercentage

`func (o *TaxProfile) HasContributionsPercentage() bool`

HasContributionsPercentage returns a boolean if a field has been set.

### SetContributionsPercentageNil

`func (o *TaxProfile) SetContributionsPercentageNil(b bool)`

 SetContributionsPercentageNil sets the value for ContributionsPercentage to be an explicit nil

### UnsetContributionsPercentage
`func (o *TaxProfile) UnsetContributionsPercentage()`

UnsetContributionsPercentage ensures that no value is present for ContributionsPercentage, not even an explicit nil
### GetProfitCoefficient

`func (o *TaxProfile) GetProfitCoefficient() float32`

GetProfitCoefficient returns the ProfitCoefficient field if non-nil, zero value otherwise.

### GetProfitCoefficientOk

`func (o *TaxProfile) GetProfitCoefficientOk() (*float32, bool)`

GetProfitCoefficientOk returns a tuple with the ProfitCoefficient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitCoefficient

`func (o *TaxProfile) SetProfitCoefficient(v float32)`

SetProfitCoefficient sets ProfitCoefficient field to given value.

### HasProfitCoefficient

`func (o *TaxProfile) HasProfitCoefficient() bool`

HasProfitCoefficient returns a boolean if a field has been set.

### SetProfitCoefficientNil

`func (o *TaxProfile) SetProfitCoefficientNil(b bool)`

 SetProfitCoefficientNil sets the value for ProfitCoefficient to be an explicit nil

### UnsetProfitCoefficient
`func (o *TaxProfile) UnsetProfitCoefficient()`

UnsetProfitCoefficient ensures that no value is present for ProfitCoefficient, not even an explicit nil
### GetMed

`func (o *TaxProfile) GetMed() bool`

GetMed returns the Med field if non-nil, zero value otherwise.

### GetMedOk

`func (o *TaxProfile) GetMedOk() (*bool, bool)`

GetMedOk returns a tuple with the Med field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMed

`func (o *TaxProfile) SetMed(v bool)`

SetMed sets Med field to given value.

### HasMed

`func (o *TaxProfile) HasMed() bool`

HasMed returns a boolean if a field has been set.

### SetMedNil

`func (o *TaxProfile) SetMedNil(b bool)`

 SetMedNil sets the value for Med to be an explicit nil

### UnsetMed
`func (o *TaxProfile) UnsetMed()`

UnsetMed ensures that no value is present for Med, not even an explicit nil
### GetDefaultVat

`func (o *TaxProfile) GetDefaultVat() VatType`

GetDefaultVat returns the DefaultVat field if non-nil, zero value otherwise.

### GetDefaultVatOk

`func (o *TaxProfile) GetDefaultVatOk() (*VatType, bool)`

GetDefaultVatOk returns a tuple with the DefaultVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVat

`func (o *TaxProfile) SetDefaultVat(v VatType)`

SetDefaultVat sets DefaultVat field to given value.

### HasDefaultVat

`func (o *TaxProfile) HasDefaultVat() bool`

HasDefaultVat returns a boolean if a field has been set.

### SetDefaultVatNil

`func (o *TaxProfile) SetDefaultVatNil(b bool)`

 SetDefaultVatNil sets the value for DefaultVat to be an explicit nil

### UnsetDefaultVat
`func (o *TaxProfile) UnsetDefaultVat()`

UnsetDefaultVat ensures that no value is present for DefaultVat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


