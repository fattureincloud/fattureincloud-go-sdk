/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.29
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the Permissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Permissions{}

// Permissions 
type Permissions struct {
	FicSituation *PermissionLevel `json:"fic_situation,omitempty"`
	FicClients *PermissionLevel `json:"fic_clients,omitempty"`
	FicSuppliers *PermissionLevel `json:"fic_suppliers,omitempty"`
	FicProducts *PermissionLevel `json:"fic_products,omitempty"`
	FicIssuedDocuments *PermissionLevel `json:"fic_issued_documents,omitempty"`
	FicReceivedDocuments *PermissionLevel `json:"fic_received_documents,omitempty"`
	FicReceipts *PermissionLevel `json:"fic_receipts,omitempty"`
	FicCalendar *PermissionLevel `json:"fic_calendar,omitempty"`
	FicArchive *PermissionLevel `json:"fic_archive,omitempty"`
	FicTaxes *PermissionLevel `json:"fic_taxes,omitempty"`
	FicStock *PermissionLevel `json:"fic_stock,omitempty"`
	FicCashbook *PermissionLevel `json:"fic_cashbook,omitempty"`
	FicSettings *PermissionLevel `json:"fic_settings,omitempty"`
	FicEmails *PermissionLevel `json:"fic_emails,omitempty"`
	FicExport *PermissionLevel `json:"fic_export,omitempty"`
	FicImportBankstatements *PermissionLevel `json:"fic_import_bankstatements,omitempty"`
	FicImportClientsSuppliers *PermissionLevel `json:"fic_import_clients_suppliers,omitempty"`
	FicImportIssuedDocuments *PermissionLevel `json:"fic_import_issued_documents,omitempty"`
	FicImportProducts *PermissionLevel `json:"fic_import_products,omitempty"`
	FicRecurring *PermissionLevel `json:"fic_recurring,omitempty"`
	FicRiba *PermissionLevel `json:"fic_riba,omitempty"`
	DicEmployees *PermissionLevel `json:"dic_employees,omitempty"`
	DicSettings *PermissionLevel `json:"dic_settings,omitempty"`
	DicTimesheet *PermissionLevel `json:"dic_timesheet,omitempty"`
	FicIssuedDocumentsDetailed NullablePermissionsFicIssuedDocumentsDetailed `json:"fic_issued_documents_detailed,omitempty"`
}

// NewPermissions instantiates a new Permissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissions() *Permissions {
	this := Permissions{}
	return &this
}

// NewPermissionsWithDefaults instantiates a new Permissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsWithDefaults() *Permissions {
	this := Permissions{}
	return &this
}

// GetFicSituation returns the FicSituation field value if set, zero value otherwise.
func (o *Permissions) GetFicSituation() PermissionLevel {
	if o == nil || IsNil(o.FicSituation) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicSituation
}

// GetFicSituationOk returns a tuple with the FicSituation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicSituationOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicSituation) {
		return nil, false
	}
	return o.FicSituation, true
}

// HasFicSituation returns a boolean if a field has been set.
func (o *Permissions) HasFicSituation() bool {
	if o != nil && !IsNil(o.FicSituation) {
		return true
	}

	return false
}

// SetFicSituation gets a reference to the given PermissionLevel and assigns it to the FicSituation field.
func (o *Permissions) SetFicSituation(v PermissionLevel) *Permissions {
	o.FicSituation = &v
	return o
}

// GetFicClients returns the FicClients field value if set, zero value otherwise.
func (o *Permissions) GetFicClients() PermissionLevel {
	if o == nil || IsNil(o.FicClients) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicClients
}

// GetFicClientsOk returns a tuple with the FicClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicClientsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicClients) {
		return nil, false
	}
	return o.FicClients, true
}

// HasFicClients returns a boolean if a field has been set.
func (o *Permissions) HasFicClients() bool {
	if o != nil && !IsNil(o.FicClients) {
		return true
	}

	return false
}

// SetFicClients gets a reference to the given PermissionLevel and assigns it to the FicClients field.
func (o *Permissions) SetFicClients(v PermissionLevel) *Permissions {
	o.FicClients = &v
	return o
}

// GetFicSuppliers returns the FicSuppliers field value if set, zero value otherwise.
func (o *Permissions) GetFicSuppliers() PermissionLevel {
	if o == nil || IsNil(o.FicSuppliers) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicSuppliers
}

// GetFicSuppliersOk returns a tuple with the FicSuppliers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicSuppliersOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicSuppliers) {
		return nil, false
	}
	return o.FicSuppliers, true
}

// HasFicSuppliers returns a boolean if a field has been set.
func (o *Permissions) HasFicSuppliers() bool {
	if o != nil && !IsNil(o.FicSuppliers) {
		return true
	}

	return false
}

// SetFicSuppliers gets a reference to the given PermissionLevel and assigns it to the FicSuppliers field.
func (o *Permissions) SetFicSuppliers(v PermissionLevel) *Permissions {
	o.FicSuppliers = &v
	return o
}

// GetFicProducts returns the FicProducts field value if set, zero value otherwise.
func (o *Permissions) GetFicProducts() PermissionLevel {
	if o == nil || IsNil(o.FicProducts) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicProducts
}

// GetFicProductsOk returns a tuple with the FicProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicProductsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicProducts) {
		return nil, false
	}
	return o.FicProducts, true
}

// HasFicProducts returns a boolean if a field has been set.
func (o *Permissions) HasFicProducts() bool {
	if o != nil && !IsNil(o.FicProducts) {
		return true
	}

	return false
}

// SetFicProducts gets a reference to the given PermissionLevel and assigns it to the FicProducts field.
func (o *Permissions) SetFicProducts(v PermissionLevel) *Permissions {
	o.FicProducts = &v
	return o
}

// GetFicIssuedDocuments returns the FicIssuedDocuments field value if set, zero value otherwise.
func (o *Permissions) GetFicIssuedDocuments() PermissionLevel {
	if o == nil || IsNil(o.FicIssuedDocuments) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicIssuedDocuments
}

// GetFicIssuedDocumentsOk returns a tuple with the FicIssuedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicIssuedDocumentsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicIssuedDocuments) {
		return nil, false
	}
	return o.FicIssuedDocuments, true
}

// HasFicIssuedDocuments returns a boolean if a field has been set.
func (o *Permissions) HasFicIssuedDocuments() bool {
	if o != nil && !IsNil(o.FicIssuedDocuments) {
		return true
	}

	return false
}

// SetFicIssuedDocuments gets a reference to the given PermissionLevel and assigns it to the FicIssuedDocuments field.
func (o *Permissions) SetFicIssuedDocuments(v PermissionLevel) *Permissions {
	o.FicIssuedDocuments = &v
	return o
}

// GetFicReceivedDocuments returns the FicReceivedDocuments field value if set, zero value otherwise.
func (o *Permissions) GetFicReceivedDocuments() PermissionLevel {
	if o == nil || IsNil(o.FicReceivedDocuments) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicReceivedDocuments
}

// GetFicReceivedDocumentsOk returns a tuple with the FicReceivedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicReceivedDocumentsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicReceivedDocuments) {
		return nil, false
	}
	return o.FicReceivedDocuments, true
}

// HasFicReceivedDocuments returns a boolean if a field has been set.
func (o *Permissions) HasFicReceivedDocuments() bool {
	if o != nil && !IsNil(o.FicReceivedDocuments) {
		return true
	}

	return false
}

// SetFicReceivedDocuments gets a reference to the given PermissionLevel and assigns it to the FicReceivedDocuments field.
func (o *Permissions) SetFicReceivedDocuments(v PermissionLevel) *Permissions {
	o.FicReceivedDocuments = &v
	return o
}

// GetFicReceipts returns the FicReceipts field value if set, zero value otherwise.
func (o *Permissions) GetFicReceipts() PermissionLevel {
	if o == nil || IsNil(o.FicReceipts) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicReceipts
}

// GetFicReceiptsOk returns a tuple with the FicReceipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicReceiptsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicReceipts) {
		return nil, false
	}
	return o.FicReceipts, true
}

// HasFicReceipts returns a boolean if a field has been set.
func (o *Permissions) HasFicReceipts() bool {
	if o != nil && !IsNil(o.FicReceipts) {
		return true
	}

	return false
}

// SetFicReceipts gets a reference to the given PermissionLevel and assigns it to the FicReceipts field.
func (o *Permissions) SetFicReceipts(v PermissionLevel) *Permissions {
	o.FicReceipts = &v
	return o
}

// GetFicCalendar returns the FicCalendar field value if set, zero value otherwise.
func (o *Permissions) GetFicCalendar() PermissionLevel {
	if o == nil || IsNil(o.FicCalendar) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicCalendar
}

// GetFicCalendarOk returns a tuple with the FicCalendar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicCalendarOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicCalendar) {
		return nil, false
	}
	return o.FicCalendar, true
}

// HasFicCalendar returns a boolean if a field has been set.
func (o *Permissions) HasFicCalendar() bool {
	if o != nil && !IsNil(o.FicCalendar) {
		return true
	}

	return false
}

// SetFicCalendar gets a reference to the given PermissionLevel and assigns it to the FicCalendar field.
func (o *Permissions) SetFicCalendar(v PermissionLevel) *Permissions {
	o.FicCalendar = &v
	return o
}

// GetFicArchive returns the FicArchive field value if set, zero value otherwise.
func (o *Permissions) GetFicArchive() PermissionLevel {
	if o == nil || IsNil(o.FicArchive) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicArchive
}

// GetFicArchiveOk returns a tuple with the FicArchive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicArchiveOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicArchive) {
		return nil, false
	}
	return o.FicArchive, true
}

// HasFicArchive returns a boolean if a field has been set.
func (o *Permissions) HasFicArchive() bool {
	if o != nil && !IsNil(o.FicArchive) {
		return true
	}

	return false
}

// SetFicArchive gets a reference to the given PermissionLevel and assigns it to the FicArchive field.
func (o *Permissions) SetFicArchive(v PermissionLevel) *Permissions {
	o.FicArchive = &v
	return o
}

// GetFicTaxes returns the FicTaxes field value if set, zero value otherwise.
func (o *Permissions) GetFicTaxes() PermissionLevel {
	if o == nil || IsNil(o.FicTaxes) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicTaxes
}

// GetFicTaxesOk returns a tuple with the FicTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicTaxesOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicTaxes) {
		return nil, false
	}
	return o.FicTaxes, true
}

// HasFicTaxes returns a boolean if a field has been set.
func (o *Permissions) HasFicTaxes() bool {
	if o != nil && !IsNil(o.FicTaxes) {
		return true
	}

	return false
}

// SetFicTaxes gets a reference to the given PermissionLevel and assigns it to the FicTaxes field.
func (o *Permissions) SetFicTaxes(v PermissionLevel) *Permissions {
	o.FicTaxes = &v
	return o
}

// GetFicStock returns the FicStock field value if set, zero value otherwise.
func (o *Permissions) GetFicStock() PermissionLevel {
	if o == nil || IsNil(o.FicStock) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicStock
}

// GetFicStockOk returns a tuple with the FicStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicStockOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicStock) {
		return nil, false
	}
	return o.FicStock, true
}

// HasFicStock returns a boolean if a field has been set.
func (o *Permissions) HasFicStock() bool {
	if o != nil && !IsNil(o.FicStock) {
		return true
	}

	return false
}

// SetFicStock gets a reference to the given PermissionLevel and assigns it to the FicStock field.
func (o *Permissions) SetFicStock(v PermissionLevel) *Permissions {
	o.FicStock = &v
	return o
}

// GetFicCashbook returns the FicCashbook field value if set, zero value otherwise.
func (o *Permissions) GetFicCashbook() PermissionLevel {
	if o == nil || IsNil(o.FicCashbook) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicCashbook
}

// GetFicCashbookOk returns a tuple with the FicCashbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicCashbookOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicCashbook) {
		return nil, false
	}
	return o.FicCashbook, true
}

// HasFicCashbook returns a boolean if a field has been set.
func (o *Permissions) HasFicCashbook() bool {
	if o != nil && !IsNil(o.FicCashbook) {
		return true
	}

	return false
}

// SetFicCashbook gets a reference to the given PermissionLevel and assigns it to the FicCashbook field.
func (o *Permissions) SetFicCashbook(v PermissionLevel) *Permissions {
	o.FicCashbook = &v
	return o
}

// GetFicSettings returns the FicSettings field value if set, zero value otherwise.
func (o *Permissions) GetFicSettings() PermissionLevel {
	if o == nil || IsNil(o.FicSettings) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicSettings
}

// GetFicSettingsOk returns a tuple with the FicSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicSettingsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicSettings) {
		return nil, false
	}
	return o.FicSettings, true
}

// HasFicSettings returns a boolean if a field has been set.
func (o *Permissions) HasFicSettings() bool {
	if o != nil && !IsNil(o.FicSettings) {
		return true
	}

	return false
}

// SetFicSettings gets a reference to the given PermissionLevel and assigns it to the FicSettings field.
func (o *Permissions) SetFicSettings(v PermissionLevel) *Permissions {
	o.FicSettings = &v
	return o
}

// GetFicEmails returns the FicEmails field value if set, zero value otherwise.
func (o *Permissions) GetFicEmails() PermissionLevel {
	if o == nil || IsNil(o.FicEmails) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicEmails
}

// GetFicEmailsOk returns a tuple with the FicEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicEmailsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicEmails) {
		return nil, false
	}
	return o.FicEmails, true
}

// HasFicEmails returns a boolean if a field has been set.
func (o *Permissions) HasFicEmails() bool {
	if o != nil && !IsNil(o.FicEmails) {
		return true
	}

	return false
}

// SetFicEmails gets a reference to the given PermissionLevel and assigns it to the FicEmails field.
func (o *Permissions) SetFicEmails(v PermissionLevel) *Permissions {
	o.FicEmails = &v
	return o
}

// GetFicExport returns the FicExport field value if set, zero value otherwise.
func (o *Permissions) GetFicExport() PermissionLevel {
	if o == nil || IsNil(o.FicExport) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicExport
}

// GetFicExportOk returns a tuple with the FicExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicExportOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicExport) {
		return nil, false
	}
	return o.FicExport, true
}

// HasFicExport returns a boolean if a field has been set.
func (o *Permissions) HasFicExport() bool {
	if o != nil && !IsNil(o.FicExport) {
		return true
	}

	return false
}

// SetFicExport gets a reference to the given PermissionLevel and assigns it to the FicExport field.
func (o *Permissions) SetFicExport(v PermissionLevel) *Permissions {
	o.FicExport = &v
	return o
}

// GetFicImportBankstatements returns the FicImportBankstatements field value if set, zero value otherwise.
func (o *Permissions) GetFicImportBankstatements() PermissionLevel {
	if o == nil || IsNil(o.FicImportBankstatements) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicImportBankstatements
}

// GetFicImportBankstatementsOk returns a tuple with the FicImportBankstatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicImportBankstatementsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicImportBankstatements) {
		return nil, false
	}
	return o.FicImportBankstatements, true
}

// HasFicImportBankstatements returns a boolean if a field has been set.
func (o *Permissions) HasFicImportBankstatements() bool {
	if o != nil && !IsNil(o.FicImportBankstatements) {
		return true
	}

	return false
}

// SetFicImportBankstatements gets a reference to the given PermissionLevel and assigns it to the FicImportBankstatements field.
func (o *Permissions) SetFicImportBankstatements(v PermissionLevel) *Permissions {
	o.FicImportBankstatements = &v
	return o
}

// GetFicImportClientsSuppliers returns the FicImportClientsSuppliers field value if set, zero value otherwise.
func (o *Permissions) GetFicImportClientsSuppliers() PermissionLevel {
	if o == nil || IsNil(o.FicImportClientsSuppliers) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicImportClientsSuppliers
}

// GetFicImportClientsSuppliersOk returns a tuple with the FicImportClientsSuppliers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicImportClientsSuppliersOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicImportClientsSuppliers) {
		return nil, false
	}
	return o.FicImportClientsSuppliers, true
}

// HasFicImportClientsSuppliers returns a boolean if a field has been set.
func (o *Permissions) HasFicImportClientsSuppliers() bool {
	if o != nil && !IsNil(o.FicImportClientsSuppliers) {
		return true
	}

	return false
}

// SetFicImportClientsSuppliers gets a reference to the given PermissionLevel and assigns it to the FicImportClientsSuppliers field.
func (o *Permissions) SetFicImportClientsSuppliers(v PermissionLevel) *Permissions {
	o.FicImportClientsSuppliers = &v
	return o
}

// GetFicImportIssuedDocuments returns the FicImportIssuedDocuments field value if set, zero value otherwise.
func (o *Permissions) GetFicImportIssuedDocuments() PermissionLevel {
	if o == nil || IsNil(o.FicImportIssuedDocuments) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicImportIssuedDocuments
}

// GetFicImportIssuedDocumentsOk returns a tuple with the FicImportIssuedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicImportIssuedDocumentsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicImportIssuedDocuments) {
		return nil, false
	}
	return o.FicImportIssuedDocuments, true
}

// HasFicImportIssuedDocuments returns a boolean if a field has been set.
func (o *Permissions) HasFicImportIssuedDocuments() bool {
	if o != nil && !IsNil(o.FicImportIssuedDocuments) {
		return true
	}

	return false
}

// SetFicImportIssuedDocuments gets a reference to the given PermissionLevel and assigns it to the FicImportIssuedDocuments field.
func (o *Permissions) SetFicImportIssuedDocuments(v PermissionLevel) *Permissions {
	o.FicImportIssuedDocuments = &v
	return o
}

// GetFicImportProducts returns the FicImportProducts field value if set, zero value otherwise.
func (o *Permissions) GetFicImportProducts() PermissionLevel {
	if o == nil || IsNil(o.FicImportProducts) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicImportProducts
}

// GetFicImportProductsOk returns a tuple with the FicImportProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicImportProductsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicImportProducts) {
		return nil, false
	}
	return o.FicImportProducts, true
}

// HasFicImportProducts returns a boolean if a field has been set.
func (o *Permissions) HasFicImportProducts() bool {
	if o != nil && !IsNil(o.FicImportProducts) {
		return true
	}

	return false
}

// SetFicImportProducts gets a reference to the given PermissionLevel and assigns it to the FicImportProducts field.
func (o *Permissions) SetFicImportProducts(v PermissionLevel) *Permissions {
	o.FicImportProducts = &v
	return o
}

// GetFicRecurring returns the FicRecurring field value if set, zero value otherwise.
func (o *Permissions) GetFicRecurring() PermissionLevel {
	if o == nil || IsNil(o.FicRecurring) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicRecurring
}

// GetFicRecurringOk returns a tuple with the FicRecurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicRecurringOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicRecurring) {
		return nil, false
	}
	return o.FicRecurring, true
}

// HasFicRecurring returns a boolean if a field has been set.
func (o *Permissions) HasFicRecurring() bool {
	if o != nil && !IsNil(o.FicRecurring) {
		return true
	}

	return false
}

// SetFicRecurring gets a reference to the given PermissionLevel and assigns it to the FicRecurring field.
func (o *Permissions) SetFicRecurring(v PermissionLevel) *Permissions {
	o.FicRecurring = &v
	return o
}

// GetFicRiba returns the FicRiba field value if set, zero value otherwise.
func (o *Permissions) GetFicRiba() PermissionLevel {
	if o == nil || IsNil(o.FicRiba) {
		var ret PermissionLevel
		return ret
	}
	return *o.FicRiba
}

// GetFicRibaOk returns a tuple with the FicRiba field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetFicRibaOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.FicRiba) {
		return nil, false
	}
	return o.FicRiba, true
}

// HasFicRiba returns a boolean if a field has been set.
func (o *Permissions) HasFicRiba() bool {
	if o != nil && !IsNil(o.FicRiba) {
		return true
	}

	return false
}

// SetFicRiba gets a reference to the given PermissionLevel and assigns it to the FicRiba field.
func (o *Permissions) SetFicRiba(v PermissionLevel) *Permissions {
	o.FicRiba = &v
	return o
}

// GetDicEmployees returns the DicEmployees field value if set, zero value otherwise.
func (o *Permissions) GetDicEmployees() PermissionLevel {
	if o == nil || IsNil(o.DicEmployees) {
		var ret PermissionLevel
		return ret
	}
	return *o.DicEmployees
}

// GetDicEmployeesOk returns a tuple with the DicEmployees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetDicEmployeesOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.DicEmployees) {
		return nil, false
	}
	return o.DicEmployees, true
}

// HasDicEmployees returns a boolean if a field has been set.
func (o *Permissions) HasDicEmployees() bool {
	if o != nil && !IsNil(o.DicEmployees) {
		return true
	}

	return false
}

// SetDicEmployees gets a reference to the given PermissionLevel and assigns it to the DicEmployees field.
func (o *Permissions) SetDicEmployees(v PermissionLevel) *Permissions {
	o.DicEmployees = &v
	return o
}

// GetDicSettings returns the DicSettings field value if set, zero value otherwise.
func (o *Permissions) GetDicSettings() PermissionLevel {
	if o == nil || IsNil(o.DicSettings) {
		var ret PermissionLevel
		return ret
	}
	return *o.DicSettings
}

// GetDicSettingsOk returns a tuple with the DicSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetDicSettingsOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.DicSettings) {
		return nil, false
	}
	return o.DicSettings, true
}

// HasDicSettings returns a boolean if a field has been set.
func (o *Permissions) HasDicSettings() bool {
	if o != nil && !IsNil(o.DicSettings) {
		return true
	}

	return false
}

// SetDicSettings gets a reference to the given PermissionLevel and assigns it to the DicSettings field.
func (o *Permissions) SetDicSettings(v PermissionLevel) *Permissions {
	o.DicSettings = &v
	return o
}

// GetDicTimesheet returns the DicTimesheet field value if set, zero value otherwise.
func (o *Permissions) GetDicTimesheet() PermissionLevel {
	if o == nil || IsNil(o.DicTimesheet) {
		var ret PermissionLevel
		return ret
	}
	return *o.DicTimesheet
}

// GetDicTimesheetOk returns a tuple with the DicTimesheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetDicTimesheetOk() (*PermissionLevel, bool) {
	if o == nil || IsNil(o.DicTimesheet) {
		return nil, false
	}
	return o.DicTimesheet, true
}

// HasDicTimesheet returns a boolean if a field has been set.
func (o *Permissions) HasDicTimesheet() bool {
	if o != nil && !IsNil(o.DicTimesheet) {
		return true
	}

	return false
}

// SetDicTimesheet gets a reference to the given PermissionLevel and assigns it to the DicTimesheet field.
func (o *Permissions) SetDicTimesheet(v PermissionLevel) *Permissions {
	o.DicTimesheet = &v
	return o
}

// GetFicIssuedDocumentsDetailed returns the FicIssuedDocumentsDetailed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permissions) GetFicIssuedDocumentsDetailed() PermissionsFicIssuedDocumentsDetailed {
	if o == nil || IsNil(o.FicIssuedDocumentsDetailed.Get()) {
		var ret PermissionsFicIssuedDocumentsDetailed
		return ret
	}
	return *o.FicIssuedDocumentsDetailed.Get()
}

// GetFicIssuedDocumentsDetailedOk returns a tuple with the FicIssuedDocumentsDetailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permissions) GetFicIssuedDocumentsDetailedOk() (*PermissionsFicIssuedDocumentsDetailed, bool) {
	if o == nil {
		return nil, false
	}
	return o.FicIssuedDocumentsDetailed.Get(), o.FicIssuedDocumentsDetailed.IsSet()
}

// HasFicIssuedDocumentsDetailed returns a boolean if a field has been set.
func (o *Permissions) HasFicIssuedDocumentsDetailed() bool {
	if o != nil && o.FicIssuedDocumentsDetailed.IsSet() {
		return true
	}

	return false
}

// SetFicIssuedDocumentsDetailed gets a reference to the given NullablePermissionsFicIssuedDocumentsDetailed and assigns it to the FicIssuedDocumentsDetailed field.
func (o *Permissions) SetFicIssuedDocumentsDetailed(v PermissionsFicIssuedDocumentsDetailed) *Permissions {
	o.FicIssuedDocumentsDetailed.Set(&v)
	return o
}
// SetFicIssuedDocumentsDetailedNil sets the value for FicIssuedDocumentsDetailed to be an explicit nil
func (o *Permissions) SetFicIssuedDocumentsDetailedNil() *Permissions {
	o.FicIssuedDocumentsDetailed.Set(nil)
	return o
}

// UnsetFicIssuedDocumentsDetailed ensures that no value is present for FicIssuedDocumentsDetailed, not even an explicit nil
func (o *Permissions) UnsetFicIssuedDocumentsDetailed() {
	o.FicIssuedDocumentsDetailed.Unset()
}

func (o Permissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Permissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FicSituation) {
		toSerialize["fic_situation"] = o.FicSituation
	}
	if !IsNil(o.FicClients) {
		toSerialize["fic_clients"] = o.FicClients
	}
	if !IsNil(o.FicSuppliers) {
		toSerialize["fic_suppliers"] = o.FicSuppliers
	}
	if !IsNil(o.FicProducts) {
		toSerialize["fic_products"] = o.FicProducts
	}
	if !IsNil(o.FicIssuedDocuments) {
		toSerialize["fic_issued_documents"] = o.FicIssuedDocuments
	}
	if !IsNil(o.FicReceivedDocuments) {
		toSerialize["fic_received_documents"] = o.FicReceivedDocuments
	}
	if !IsNil(o.FicReceipts) {
		toSerialize["fic_receipts"] = o.FicReceipts
	}
	if !IsNil(o.FicCalendar) {
		toSerialize["fic_calendar"] = o.FicCalendar
	}
	if !IsNil(o.FicArchive) {
		toSerialize["fic_archive"] = o.FicArchive
	}
	if !IsNil(o.FicTaxes) {
		toSerialize["fic_taxes"] = o.FicTaxes
	}
	if !IsNil(o.FicStock) {
		toSerialize["fic_stock"] = o.FicStock
	}
	if !IsNil(o.FicCashbook) {
		toSerialize["fic_cashbook"] = o.FicCashbook
	}
	if !IsNil(o.FicSettings) {
		toSerialize["fic_settings"] = o.FicSettings
	}
	if !IsNil(o.FicEmails) {
		toSerialize["fic_emails"] = o.FicEmails
	}
	if !IsNil(o.FicExport) {
		toSerialize["fic_export"] = o.FicExport
	}
	if !IsNil(o.FicImportBankstatements) {
		toSerialize["fic_import_bankstatements"] = o.FicImportBankstatements
	}
	if !IsNil(o.FicImportClientsSuppliers) {
		toSerialize["fic_import_clients_suppliers"] = o.FicImportClientsSuppliers
	}
	if !IsNil(o.FicImportIssuedDocuments) {
		toSerialize["fic_import_issued_documents"] = o.FicImportIssuedDocuments
	}
	if !IsNil(o.FicImportProducts) {
		toSerialize["fic_import_products"] = o.FicImportProducts
	}
	if !IsNil(o.FicRecurring) {
		toSerialize["fic_recurring"] = o.FicRecurring
	}
	if !IsNil(o.FicRiba) {
		toSerialize["fic_riba"] = o.FicRiba
	}
	if !IsNil(o.DicEmployees) {
		toSerialize["dic_employees"] = o.DicEmployees
	}
	if !IsNil(o.DicSettings) {
		toSerialize["dic_settings"] = o.DicSettings
	}
	if !IsNil(o.DicTimesheet) {
		toSerialize["dic_timesheet"] = o.DicTimesheet
	}
	if o.FicIssuedDocumentsDetailed.IsSet() {
		toSerialize["fic_issued_documents_detailed"] = o.FicIssuedDocumentsDetailed.Get()
	}
	return toSerialize, nil
}

type NullablePermissions struct {
	value *Permissions
	isSet bool
}

func (v NullablePermissions) Get() *Permissions {
	return v.value
}

func (v *NullablePermissions) Set(val *Permissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissions(val *Permissions) *NullablePermissions {
	return &NullablePermissions{value: val, isSet: true}
}

func (v NullablePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


