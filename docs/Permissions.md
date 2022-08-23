# Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FicSituation** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicClients** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicSuppliers** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicProducts** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicIssuedDocuments** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicReceivedDocuments** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicReceipts** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicCalendar** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicArchive** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicTaxes** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicStock** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicCashbook** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicSettings** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicEmails** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicExport** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicImportBankstatements** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicImportClientsSuppliers** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicImportIssuedDocuments** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicImportProducts** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicRecurring** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicRiba** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**DicEmployees** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**DicSettings** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**DicTimesheet** | Pointer to [**PermissionLevel**](PermissionLevel.md) |  | [optional] 
**FicIssuedDocumentsDetailed** | Pointer to [**NullablePermissionsFicIssuedDocumentsDetailed**](PermissionsFicIssuedDocumentsDetailed.md) |  | [optional] 

## Methods

### NewPermissions

`func NewPermissions() *Permissions`

NewPermissions instantiates a new Permissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsWithDefaults

`func NewPermissionsWithDefaults() *Permissions`

NewPermissionsWithDefaults instantiates a new Permissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFicSituation

`func (o *Permissions) GetFicSituation() PermissionLevel`

GetFicSituation returns the FicSituation field if non-nil, zero value otherwise.

### GetFicSituationOk

`func (o *Permissions) GetFicSituationOk() (*PermissionLevel, bool)`

GetFicSituationOk returns a tuple with the FicSituation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicSituation

`func (o *Permissions) SetFicSituation(v PermissionLevel)`

SetFicSituation sets FicSituation field to given value.

### HasFicSituation

`func (o *Permissions) HasFicSituation() bool`

HasFicSituation returns a boolean if a field has been set.

### GetFicClients

`func (o *Permissions) GetFicClients() PermissionLevel`

GetFicClients returns the FicClients field if non-nil, zero value otherwise.

### GetFicClientsOk

`func (o *Permissions) GetFicClientsOk() (*PermissionLevel, bool)`

GetFicClientsOk returns a tuple with the FicClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicClients

`func (o *Permissions) SetFicClients(v PermissionLevel)`

SetFicClients sets FicClients field to given value.

### HasFicClients

`func (o *Permissions) HasFicClients() bool`

HasFicClients returns a boolean if a field has been set.

### GetFicSuppliers

`func (o *Permissions) GetFicSuppliers() PermissionLevel`

GetFicSuppliers returns the FicSuppliers field if non-nil, zero value otherwise.

### GetFicSuppliersOk

`func (o *Permissions) GetFicSuppliersOk() (*PermissionLevel, bool)`

GetFicSuppliersOk returns a tuple with the FicSuppliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicSuppliers

`func (o *Permissions) SetFicSuppliers(v PermissionLevel)`

SetFicSuppliers sets FicSuppliers field to given value.

### HasFicSuppliers

`func (o *Permissions) HasFicSuppliers() bool`

HasFicSuppliers returns a boolean if a field has been set.

### GetFicProducts

`func (o *Permissions) GetFicProducts() PermissionLevel`

GetFicProducts returns the FicProducts field if non-nil, zero value otherwise.

### GetFicProductsOk

`func (o *Permissions) GetFicProductsOk() (*PermissionLevel, bool)`

GetFicProductsOk returns a tuple with the FicProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicProducts

`func (o *Permissions) SetFicProducts(v PermissionLevel)`

SetFicProducts sets FicProducts field to given value.

### HasFicProducts

`func (o *Permissions) HasFicProducts() bool`

HasFicProducts returns a boolean if a field has been set.

### GetFicIssuedDocuments

`func (o *Permissions) GetFicIssuedDocuments() PermissionLevel`

GetFicIssuedDocuments returns the FicIssuedDocuments field if non-nil, zero value otherwise.

### GetFicIssuedDocumentsOk

`func (o *Permissions) GetFicIssuedDocumentsOk() (*PermissionLevel, bool)`

GetFicIssuedDocumentsOk returns a tuple with the FicIssuedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicIssuedDocuments

`func (o *Permissions) SetFicIssuedDocuments(v PermissionLevel)`

SetFicIssuedDocuments sets FicIssuedDocuments field to given value.

### HasFicIssuedDocuments

`func (o *Permissions) HasFicIssuedDocuments() bool`

HasFicIssuedDocuments returns a boolean if a field has been set.

### GetFicReceivedDocuments

`func (o *Permissions) GetFicReceivedDocuments() PermissionLevel`

GetFicReceivedDocuments returns the FicReceivedDocuments field if non-nil, zero value otherwise.

### GetFicReceivedDocumentsOk

`func (o *Permissions) GetFicReceivedDocumentsOk() (*PermissionLevel, bool)`

GetFicReceivedDocumentsOk returns a tuple with the FicReceivedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicReceivedDocuments

`func (o *Permissions) SetFicReceivedDocuments(v PermissionLevel)`

SetFicReceivedDocuments sets FicReceivedDocuments field to given value.

### HasFicReceivedDocuments

`func (o *Permissions) HasFicReceivedDocuments() bool`

HasFicReceivedDocuments returns a boolean if a field has been set.

### GetFicReceipts

`func (o *Permissions) GetFicReceipts() PermissionLevel`

GetFicReceipts returns the FicReceipts field if non-nil, zero value otherwise.

### GetFicReceiptsOk

`func (o *Permissions) GetFicReceiptsOk() (*PermissionLevel, bool)`

GetFicReceiptsOk returns a tuple with the FicReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicReceipts

`func (o *Permissions) SetFicReceipts(v PermissionLevel)`

SetFicReceipts sets FicReceipts field to given value.

### HasFicReceipts

`func (o *Permissions) HasFicReceipts() bool`

HasFicReceipts returns a boolean if a field has been set.

### GetFicCalendar

`func (o *Permissions) GetFicCalendar() PermissionLevel`

GetFicCalendar returns the FicCalendar field if non-nil, zero value otherwise.

### GetFicCalendarOk

`func (o *Permissions) GetFicCalendarOk() (*PermissionLevel, bool)`

GetFicCalendarOk returns a tuple with the FicCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicCalendar

`func (o *Permissions) SetFicCalendar(v PermissionLevel)`

SetFicCalendar sets FicCalendar field to given value.

### HasFicCalendar

`func (o *Permissions) HasFicCalendar() bool`

HasFicCalendar returns a boolean if a field has been set.

### GetFicArchive

`func (o *Permissions) GetFicArchive() PermissionLevel`

GetFicArchive returns the FicArchive field if non-nil, zero value otherwise.

### GetFicArchiveOk

`func (o *Permissions) GetFicArchiveOk() (*PermissionLevel, bool)`

GetFicArchiveOk returns a tuple with the FicArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicArchive

`func (o *Permissions) SetFicArchive(v PermissionLevel)`

SetFicArchive sets FicArchive field to given value.

### HasFicArchive

`func (o *Permissions) HasFicArchive() bool`

HasFicArchive returns a boolean if a field has been set.

### GetFicTaxes

`func (o *Permissions) GetFicTaxes() PermissionLevel`

GetFicTaxes returns the FicTaxes field if non-nil, zero value otherwise.

### GetFicTaxesOk

`func (o *Permissions) GetFicTaxesOk() (*PermissionLevel, bool)`

GetFicTaxesOk returns a tuple with the FicTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicTaxes

`func (o *Permissions) SetFicTaxes(v PermissionLevel)`

SetFicTaxes sets FicTaxes field to given value.

### HasFicTaxes

`func (o *Permissions) HasFicTaxes() bool`

HasFicTaxes returns a boolean if a field has been set.

### GetFicStock

`func (o *Permissions) GetFicStock() PermissionLevel`

GetFicStock returns the FicStock field if non-nil, zero value otherwise.

### GetFicStockOk

`func (o *Permissions) GetFicStockOk() (*PermissionLevel, bool)`

GetFicStockOk returns a tuple with the FicStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicStock

`func (o *Permissions) SetFicStock(v PermissionLevel)`

SetFicStock sets FicStock field to given value.

### HasFicStock

`func (o *Permissions) HasFicStock() bool`

HasFicStock returns a boolean if a field has been set.

### GetFicCashbook

`func (o *Permissions) GetFicCashbook() PermissionLevel`

GetFicCashbook returns the FicCashbook field if non-nil, zero value otherwise.

### GetFicCashbookOk

`func (o *Permissions) GetFicCashbookOk() (*PermissionLevel, bool)`

GetFicCashbookOk returns a tuple with the FicCashbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicCashbook

`func (o *Permissions) SetFicCashbook(v PermissionLevel)`

SetFicCashbook sets FicCashbook field to given value.

### HasFicCashbook

`func (o *Permissions) HasFicCashbook() bool`

HasFicCashbook returns a boolean if a field has been set.

### GetFicSettings

`func (o *Permissions) GetFicSettings() PermissionLevel`

GetFicSettings returns the FicSettings field if non-nil, zero value otherwise.

### GetFicSettingsOk

`func (o *Permissions) GetFicSettingsOk() (*PermissionLevel, bool)`

GetFicSettingsOk returns a tuple with the FicSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicSettings

`func (o *Permissions) SetFicSettings(v PermissionLevel)`

SetFicSettings sets FicSettings field to given value.

### HasFicSettings

`func (o *Permissions) HasFicSettings() bool`

HasFicSettings returns a boolean if a field has been set.

### GetFicEmails

`func (o *Permissions) GetFicEmails() PermissionLevel`

GetFicEmails returns the FicEmails field if non-nil, zero value otherwise.

### GetFicEmailsOk

`func (o *Permissions) GetFicEmailsOk() (*PermissionLevel, bool)`

GetFicEmailsOk returns a tuple with the FicEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicEmails

`func (o *Permissions) SetFicEmails(v PermissionLevel)`

SetFicEmails sets FicEmails field to given value.

### HasFicEmails

`func (o *Permissions) HasFicEmails() bool`

HasFicEmails returns a boolean if a field has been set.

### GetFicExport

`func (o *Permissions) GetFicExport() PermissionLevel`

GetFicExport returns the FicExport field if non-nil, zero value otherwise.

### GetFicExportOk

`func (o *Permissions) GetFicExportOk() (*PermissionLevel, bool)`

GetFicExportOk returns a tuple with the FicExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicExport

`func (o *Permissions) SetFicExport(v PermissionLevel)`

SetFicExport sets FicExport field to given value.

### HasFicExport

`func (o *Permissions) HasFicExport() bool`

HasFicExport returns a boolean if a field has been set.

### GetFicImportBankstatements

`func (o *Permissions) GetFicImportBankstatements() PermissionLevel`

GetFicImportBankstatements returns the FicImportBankstatements field if non-nil, zero value otherwise.

### GetFicImportBankstatementsOk

`func (o *Permissions) GetFicImportBankstatementsOk() (*PermissionLevel, bool)`

GetFicImportBankstatementsOk returns a tuple with the FicImportBankstatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicImportBankstatements

`func (o *Permissions) SetFicImportBankstatements(v PermissionLevel)`

SetFicImportBankstatements sets FicImportBankstatements field to given value.

### HasFicImportBankstatements

`func (o *Permissions) HasFicImportBankstatements() bool`

HasFicImportBankstatements returns a boolean if a field has been set.

### GetFicImportClientsSuppliers

`func (o *Permissions) GetFicImportClientsSuppliers() PermissionLevel`

GetFicImportClientsSuppliers returns the FicImportClientsSuppliers field if non-nil, zero value otherwise.

### GetFicImportClientsSuppliersOk

`func (o *Permissions) GetFicImportClientsSuppliersOk() (*PermissionLevel, bool)`

GetFicImportClientsSuppliersOk returns a tuple with the FicImportClientsSuppliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicImportClientsSuppliers

`func (o *Permissions) SetFicImportClientsSuppliers(v PermissionLevel)`

SetFicImportClientsSuppliers sets FicImportClientsSuppliers field to given value.

### HasFicImportClientsSuppliers

`func (o *Permissions) HasFicImportClientsSuppliers() bool`

HasFicImportClientsSuppliers returns a boolean if a field has been set.

### GetFicImportIssuedDocuments

`func (o *Permissions) GetFicImportIssuedDocuments() PermissionLevel`

GetFicImportIssuedDocuments returns the FicImportIssuedDocuments field if non-nil, zero value otherwise.

### GetFicImportIssuedDocumentsOk

`func (o *Permissions) GetFicImportIssuedDocumentsOk() (*PermissionLevel, bool)`

GetFicImportIssuedDocumentsOk returns a tuple with the FicImportIssuedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicImportIssuedDocuments

`func (o *Permissions) SetFicImportIssuedDocuments(v PermissionLevel)`

SetFicImportIssuedDocuments sets FicImportIssuedDocuments field to given value.

### HasFicImportIssuedDocuments

`func (o *Permissions) HasFicImportIssuedDocuments() bool`

HasFicImportIssuedDocuments returns a boolean if a field has been set.

### GetFicImportProducts

`func (o *Permissions) GetFicImportProducts() PermissionLevel`

GetFicImportProducts returns the FicImportProducts field if non-nil, zero value otherwise.

### GetFicImportProductsOk

`func (o *Permissions) GetFicImportProductsOk() (*PermissionLevel, bool)`

GetFicImportProductsOk returns a tuple with the FicImportProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicImportProducts

`func (o *Permissions) SetFicImportProducts(v PermissionLevel)`

SetFicImportProducts sets FicImportProducts field to given value.

### HasFicImportProducts

`func (o *Permissions) HasFicImportProducts() bool`

HasFicImportProducts returns a boolean if a field has been set.

### GetFicRecurring

`func (o *Permissions) GetFicRecurring() PermissionLevel`

GetFicRecurring returns the FicRecurring field if non-nil, zero value otherwise.

### GetFicRecurringOk

`func (o *Permissions) GetFicRecurringOk() (*PermissionLevel, bool)`

GetFicRecurringOk returns a tuple with the FicRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicRecurring

`func (o *Permissions) SetFicRecurring(v PermissionLevel)`

SetFicRecurring sets FicRecurring field to given value.

### HasFicRecurring

`func (o *Permissions) HasFicRecurring() bool`

HasFicRecurring returns a boolean if a field has been set.

### GetFicRiba

`func (o *Permissions) GetFicRiba() PermissionLevel`

GetFicRiba returns the FicRiba field if non-nil, zero value otherwise.

### GetFicRibaOk

`func (o *Permissions) GetFicRibaOk() (*PermissionLevel, bool)`

GetFicRibaOk returns a tuple with the FicRiba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicRiba

`func (o *Permissions) SetFicRiba(v PermissionLevel)`

SetFicRiba sets FicRiba field to given value.

### HasFicRiba

`func (o *Permissions) HasFicRiba() bool`

HasFicRiba returns a boolean if a field has been set.

### GetDicEmployees

`func (o *Permissions) GetDicEmployees() PermissionLevel`

GetDicEmployees returns the DicEmployees field if non-nil, zero value otherwise.

### GetDicEmployeesOk

`func (o *Permissions) GetDicEmployeesOk() (*PermissionLevel, bool)`

GetDicEmployeesOk returns a tuple with the DicEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDicEmployees

`func (o *Permissions) SetDicEmployees(v PermissionLevel)`

SetDicEmployees sets DicEmployees field to given value.

### HasDicEmployees

`func (o *Permissions) HasDicEmployees() bool`

HasDicEmployees returns a boolean if a field has been set.

### GetDicSettings

`func (o *Permissions) GetDicSettings() PermissionLevel`

GetDicSettings returns the DicSettings field if non-nil, zero value otherwise.

### GetDicSettingsOk

`func (o *Permissions) GetDicSettingsOk() (*PermissionLevel, bool)`

GetDicSettingsOk returns a tuple with the DicSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDicSettings

`func (o *Permissions) SetDicSettings(v PermissionLevel)`

SetDicSettings sets DicSettings field to given value.

### HasDicSettings

`func (o *Permissions) HasDicSettings() bool`

HasDicSettings returns a boolean if a field has been set.

### GetDicTimesheet

`func (o *Permissions) GetDicTimesheet() PermissionLevel`

GetDicTimesheet returns the DicTimesheet field if non-nil, zero value otherwise.

### GetDicTimesheetOk

`func (o *Permissions) GetDicTimesheetOk() (*PermissionLevel, bool)`

GetDicTimesheetOk returns a tuple with the DicTimesheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDicTimesheet

`func (o *Permissions) SetDicTimesheet(v PermissionLevel)`

SetDicTimesheet sets DicTimesheet field to given value.

### HasDicTimesheet

`func (o *Permissions) HasDicTimesheet() bool`

HasDicTimesheet returns a boolean if a field has been set.

### GetFicIssuedDocumentsDetailed

`func (o *Permissions) GetFicIssuedDocumentsDetailed() PermissionsFicIssuedDocumentsDetailed`

GetFicIssuedDocumentsDetailed returns the FicIssuedDocumentsDetailed field if non-nil, zero value otherwise.

### GetFicIssuedDocumentsDetailedOk

`func (o *Permissions) GetFicIssuedDocumentsDetailedOk() (*PermissionsFicIssuedDocumentsDetailed, bool)`

GetFicIssuedDocumentsDetailedOk returns a tuple with the FicIssuedDocumentsDetailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFicIssuedDocumentsDetailed

`func (o *Permissions) SetFicIssuedDocumentsDetailed(v PermissionsFicIssuedDocumentsDetailed)`

SetFicIssuedDocumentsDetailed sets FicIssuedDocumentsDetailed field to given value.

### HasFicIssuedDocumentsDetailed

`func (o *Permissions) HasFicIssuedDocumentsDetailed() bool`

HasFicIssuedDocumentsDetailed returns a boolean if a field has been set.

### SetFicIssuedDocumentsDetailedNil

`func (o *Permissions) SetFicIssuedDocumentsDetailedNil(b bool)`

 SetFicIssuedDocumentsDetailedNil sets the value for FicIssuedDocumentsDetailed to be an explicit nil

### UnsetFicIssuedDocumentsDetailed
`func (o *Permissions) UnsetFicIssuedDocumentsDetailed()`

UnsetFicIssuedDocumentsDetailed ensures that no value is present for FicIssuedDocumentsDetailed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


