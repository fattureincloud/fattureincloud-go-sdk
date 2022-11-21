/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.22
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// IssuedDocumentItemsListItem struct for IssuedDocumentItemsListItem
type IssuedDocumentItemsListItem struct {
	// Unique identifier.
	Id NullableInt32 `json:"id,omitempty"`
	// Unique identifier of the product.
	ProductId NullableInt32 `json:"product_id,omitempty"`
	// Product code.
	Code NullableString `json:"code,omitempty"`
	// Product name.
	Name NullableString `json:"name,omitempty"`
	// Product category
	Category NullableString `json:"category,omitempty"`
	// Product description.
	Description NullableString `json:"description,omitempty"`
	// Items quantity,
	Qty NullableFloat32 `json:"qty,omitempty"`
	// Item measure.
	Measure NullableString `json:"measure,omitempty"`
	// Net price.
	NetPrice NullableFloat32 `json:"net_price,omitempty"`
	// Gross price.
	GrossPrice NullableFloat32 `json:"gross_price,omitempty"`
	Vat NullableVatType `json:"vat,omitempty"`
	NotTaxable NullableBool `json:"not_taxable,omitempty"`
	// Apply withholding taxes, rivalsa and cassa.
	ApplyWithholdingTaxes NullableBool `json:"apply_withholding_taxes,omitempty"`
	// Discount percentual value.
	Discount NullableFloat32 `json:"discount,omitempty"`
	DiscountHighlight NullableBool `json:"discount_highlight,omitempty"`
	InDdt NullableBool `json:"in_ddt,omitempty"`
	Stock NullableBool `json:"stock,omitempty"`
	// Advanced raw attributes for e-invoices.
	EiRaw map[string]interface{} `json:"ei_raw,omitempty"`
}

// NewIssuedDocumentItemsListItem instantiates a new IssuedDocumentItemsListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentItemsListItem() *IssuedDocumentItemsListItem {
	this := IssuedDocumentItemsListItem{}
	return &this
}

// NewIssuedDocumentItemsListItemWithDefaults instantiates a new IssuedDocumentItemsListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentItemsListItemWithDefaults() *IssuedDocumentItemsListItem {
	this := IssuedDocumentItemsListItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *IssuedDocumentItemsListItem) SetId(v int32) *IssuedDocumentItemsListItem {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetIdNil() *IssuedDocumentItemsListItem {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetId() {
	o.Id.Unset()
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetProductId() int32 {
	if o == nil || o.ProductId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetProductIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableInt32 and assigns it to the ProductId field.
func (o *IssuedDocumentItemsListItem) SetProductId(v int32) *IssuedDocumentItemsListItem {
	o.ProductId.Set(&v)
	return o
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetProductIdNil() *IssuedDocumentItemsListItem {
	o.ProductId.Set(nil)
	return o
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetProductId() {
	o.ProductId.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *IssuedDocumentItemsListItem) SetCode(v string) *IssuedDocumentItemsListItem {
	o.Code.Set(&v)
	return o
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetCodeNil() *IssuedDocumentItemsListItem {
	o.Code.Set(nil)
	return o
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetCode() {
	o.Code.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IssuedDocumentItemsListItem) SetName(v string) *IssuedDocumentItemsListItem {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetNameNil() *IssuedDocumentItemsListItem {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetName() {
	o.Name.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *IssuedDocumentItemsListItem) SetCategory(v string) *IssuedDocumentItemsListItem {
	o.Category.Set(&v)
	return o
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetCategoryNil() *IssuedDocumentItemsListItem {
	o.Category.Set(nil)
	return o
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetCategory() {
	o.Category.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IssuedDocumentItemsListItem) SetDescription(v string) *IssuedDocumentItemsListItem {
	o.Description.Set(&v)
	return o
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetDescriptionNil() *IssuedDocumentItemsListItem {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetDescription() {
	o.Description.Unset()
}

// GetQty returns the Qty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetQty() float32 {
	if o == nil || o.Qty.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Qty.Get()
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Qty.Get(), o.Qty.IsSet()
}

// HasQty returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasQty() bool {
	if o != nil && o.Qty.IsSet() {
		return true
	}

	return false
}

// SetQty gets a reference to the given NullableFloat32 and assigns it to the Qty field.
func (o *IssuedDocumentItemsListItem) SetQty(v float32) *IssuedDocumentItemsListItem {
	o.Qty.Set(&v)
	return o
}
// SetQtyNil sets the value for Qty to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetQtyNil() *IssuedDocumentItemsListItem {
	o.Qty.Set(nil)
	return o
}

// UnsetQty ensures that no value is present for Qty, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetQty() {
	o.Qty.Unset()
}

// GetMeasure returns the Measure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetMeasure() string {
	if o == nil || o.Measure.Get() == nil {
		var ret string
		return ret
	}
	return *o.Measure.Get()
}

// GetMeasureOk returns a tuple with the Measure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Measure.Get(), o.Measure.IsSet()
}

// HasMeasure returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasMeasure() bool {
	if o != nil && o.Measure.IsSet() {
		return true
	}

	return false
}

// SetMeasure gets a reference to the given NullableString and assigns it to the Measure field.
func (o *IssuedDocumentItemsListItem) SetMeasure(v string) *IssuedDocumentItemsListItem {
	o.Measure.Set(&v)
	return o
}
// SetMeasureNil sets the value for Measure to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetMeasureNil() *IssuedDocumentItemsListItem {
	o.Measure.Set(nil)
	return o
}

// UnsetMeasure ensures that no value is present for Measure, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetMeasure() {
	o.Measure.Unset()
}

// GetNetPrice returns the NetPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetNetPrice() float32 {
	if o == nil || o.NetPrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.NetPrice.Get()
}

// GetNetPriceOk returns a tuple with the NetPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetNetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetPrice.Get(), o.NetPrice.IsSet()
}

// HasNetPrice returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasNetPrice() bool {
	if o != nil && o.NetPrice.IsSet() {
		return true
	}

	return false
}

// SetNetPrice gets a reference to the given NullableFloat32 and assigns it to the NetPrice field.
func (o *IssuedDocumentItemsListItem) SetNetPrice(v float32) *IssuedDocumentItemsListItem {
	o.NetPrice.Set(&v)
	return o
}
// SetNetPriceNil sets the value for NetPrice to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetNetPriceNil() *IssuedDocumentItemsListItem {
	o.NetPrice.Set(nil)
	return o
}

// UnsetNetPrice ensures that no value is present for NetPrice, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetNetPrice() {
	o.NetPrice.Unset()
}

// GetGrossPrice returns the GrossPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetGrossPrice() float32 {
	if o == nil || o.GrossPrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.GrossPrice.Get()
}

// GetGrossPriceOk returns a tuple with the GrossPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetGrossPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrossPrice.Get(), o.GrossPrice.IsSet()
}

// HasGrossPrice returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasGrossPrice() bool {
	if o != nil && o.GrossPrice.IsSet() {
		return true
	}

	return false
}

// SetGrossPrice gets a reference to the given NullableFloat32 and assigns it to the GrossPrice field.
func (o *IssuedDocumentItemsListItem) SetGrossPrice(v float32) *IssuedDocumentItemsListItem {
	o.GrossPrice.Set(&v)
	return o
}
// SetGrossPriceNil sets the value for GrossPrice to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetGrossPriceNil() *IssuedDocumentItemsListItem {
	o.GrossPrice.Set(nil)
	return o
}

// UnsetGrossPrice ensures that no value is present for GrossPrice, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetGrossPrice() {
	o.GrossPrice.Unset()
}

// GetVat returns the Vat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetVat() VatType {
	if o == nil || o.Vat.Get() == nil {
		var ret VatType
		return ret
	}
	return *o.Vat.Get()
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetVatOk() (*VatType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vat.Get(), o.Vat.IsSet()
}

// HasVat returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasVat() bool {
	if o != nil && o.Vat.IsSet() {
		return true
	}

	return false
}

// SetVat gets a reference to the given NullableVatType and assigns it to the Vat field.
func (o *IssuedDocumentItemsListItem) SetVat(v VatType) *IssuedDocumentItemsListItem {
	o.Vat.Set(&v)
	return o
}
// SetVatNil sets the value for Vat to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetVatNil() *IssuedDocumentItemsListItem {
	o.Vat.Set(nil)
	return o
}

// UnsetVat ensures that no value is present for Vat, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetVat() {
	o.Vat.Unset()
}

// GetNotTaxable returns the NotTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetNotTaxable() bool {
	if o == nil || o.NotTaxable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NotTaxable.Get()
}

// GetNotTaxableOk returns a tuple with the NotTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetNotTaxableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotTaxable.Get(), o.NotTaxable.IsSet()
}

// HasNotTaxable returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasNotTaxable() bool {
	if o != nil && o.NotTaxable.IsSet() {
		return true
	}

	return false
}

// SetNotTaxable gets a reference to the given NullableBool and assigns it to the NotTaxable field.
func (o *IssuedDocumentItemsListItem) SetNotTaxable(v bool) *IssuedDocumentItemsListItem {
	o.NotTaxable.Set(&v)
	return o
}
// SetNotTaxableNil sets the value for NotTaxable to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetNotTaxableNil() *IssuedDocumentItemsListItem {
	o.NotTaxable.Set(nil)
	return o
}

// UnsetNotTaxable ensures that no value is present for NotTaxable, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetNotTaxable() {
	o.NotTaxable.Unset()
}

// GetApplyWithholdingTaxes returns the ApplyWithholdingTaxes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetApplyWithholdingTaxes() bool {
	if o == nil || o.ApplyWithholdingTaxes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ApplyWithholdingTaxes.Get()
}

// GetApplyWithholdingTaxesOk returns a tuple with the ApplyWithholdingTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetApplyWithholdingTaxesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplyWithholdingTaxes.Get(), o.ApplyWithholdingTaxes.IsSet()
}

// HasApplyWithholdingTaxes returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasApplyWithholdingTaxes() bool {
	if o != nil && o.ApplyWithholdingTaxes.IsSet() {
		return true
	}

	return false
}

// SetApplyWithholdingTaxes gets a reference to the given NullableBool and assigns it to the ApplyWithholdingTaxes field.
func (o *IssuedDocumentItemsListItem) SetApplyWithholdingTaxes(v bool) *IssuedDocumentItemsListItem {
	o.ApplyWithholdingTaxes.Set(&v)
	return o
}
// SetApplyWithholdingTaxesNil sets the value for ApplyWithholdingTaxes to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetApplyWithholdingTaxesNil() *IssuedDocumentItemsListItem {
	o.ApplyWithholdingTaxes.Set(nil)
	return o
}

// UnsetApplyWithholdingTaxes ensures that no value is present for ApplyWithholdingTaxes, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetApplyWithholdingTaxes() {
	o.ApplyWithholdingTaxes.Unset()
}

// GetDiscount returns the Discount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetDiscount() float32 {
	if o == nil || o.Discount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Discount.Get()
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetDiscountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discount.Get(), o.Discount.IsSet()
}

// HasDiscount returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasDiscount() bool {
	if o != nil && o.Discount.IsSet() {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given NullableFloat32 and assigns it to the Discount field.
func (o *IssuedDocumentItemsListItem) SetDiscount(v float32) *IssuedDocumentItemsListItem {
	o.Discount.Set(&v)
	return o
}
// SetDiscountNil sets the value for Discount to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetDiscountNil() *IssuedDocumentItemsListItem {
	o.Discount.Set(nil)
	return o
}

// UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetDiscount() {
	o.Discount.Unset()
}

// GetDiscountHighlight returns the DiscountHighlight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetDiscountHighlight() bool {
	if o == nil || o.DiscountHighlight.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DiscountHighlight.Get()
}

// GetDiscountHighlightOk returns a tuple with the DiscountHighlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetDiscountHighlightOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountHighlight.Get(), o.DiscountHighlight.IsSet()
}

// HasDiscountHighlight returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasDiscountHighlight() bool {
	if o != nil && o.DiscountHighlight.IsSet() {
		return true
	}

	return false
}

// SetDiscountHighlight gets a reference to the given NullableBool and assigns it to the DiscountHighlight field.
func (o *IssuedDocumentItemsListItem) SetDiscountHighlight(v bool) *IssuedDocumentItemsListItem {
	o.DiscountHighlight.Set(&v)
	return o
}
// SetDiscountHighlightNil sets the value for DiscountHighlight to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetDiscountHighlightNil() *IssuedDocumentItemsListItem {
	o.DiscountHighlight.Set(nil)
	return o
}

// UnsetDiscountHighlight ensures that no value is present for DiscountHighlight, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetDiscountHighlight() {
	o.DiscountHighlight.Unset()
}

// GetInDdt returns the InDdt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetInDdt() bool {
	if o == nil || o.InDdt.Get() == nil {
		var ret bool
		return ret
	}
	return *o.InDdt.Get()
}

// GetInDdtOk returns a tuple with the InDdt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetInDdtOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InDdt.Get(), o.InDdt.IsSet()
}

// HasInDdt returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasInDdt() bool {
	if o != nil && o.InDdt.IsSet() {
		return true
	}

	return false
}

// SetInDdt gets a reference to the given NullableBool and assigns it to the InDdt field.
func (o *IssuedDocumentItemsListItem) SetInDdt(v bool) *IssuedDocumentItemsListItem {
	o.InDdt.Set(&v)
	return o
}
// SetInDdtNil sets the value for InDdt to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetInDdtNil() *IssuedDocumentItemsListItem {
	o.InDdt.Set(nil)
	return o
}

// UnsetInDdt ensures that no value is present for InDdt, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetInDdt() {
	o.InDdt.Unset()
}

// GetStock returns the Stock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetStock() bool {
	if o == nil || o.Stock.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Stock.Get()
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetStockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stock.Get(), o.Stock.IsSet()
}

// HasStock returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasStock() bool {
	if o != nil && o.Stock.IsSet() {
		return true
	}

	return false
}

// SetStock gets a reference to the given NullableBool and assigns it to the Stock field.
func (o *IssuedDocumentItemsListItem) SetStock(v bool) *IssuedDocumentItemsListItem {
	o.Stock.Set(&v)
	return o
}
// SetStockNil sets the value for Stock to be an explicit nil
func (o *IssuedDocumentItemsListItem) SetStockNil() *IssuedDocumentItemsListItem {
	o.Stock.Set(nil)
	return o
}

// UnsetStock ensures that no value is present for Stock, not even an explicit nil
func (o *IssuedDocumentItemsListItem) UnsetStock() {
	o.Stock.Unset()
}

// GetEiRaw returns the EiRaw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentItemsListItem) GetEiRaw() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EiRaw
}

// GetEiRawOk returns a tuple with the EiRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentItemsListItem) GetEiRawOk() (map[string]interface{}, bool) {
	if o == nil || o.EiRaw == nil {
		return nil, false
	}
	return o.EiRaw, true
}

// HasEiRaw returns a boolean if a field has been set.
func (o *IssuedDocumentItemsListItem) HasEiRaw() bool {
	if o != nil && o.EiRaw != nil {
		return true
	}

	return false
}

// SetEiRaw gets a reference to the given map[string]interface{} and assigns it to the EiRaw field.
func (o *IssuedDocumentItemsListItem) SetEiRaw(v map[string]interface{}) *IssuedDocumentItemsListItem {
	o.EiRaw = v
	return o
}

func (o IssuedDocumentItemsListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ProductId.IsSet() {
		toSerialize["product_id"] = o.ProductId.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Qty.IsSet() {
		toSerialize["qty"] = o.Qty.Get()
	}
	if o.Measure.IsSet() {
		toSerialize["measure"] = o.Measure.Get()
	}
	if o.NetPrice.IsSet() {
		toSerialize["net_price"] = o.NetPrice.Get()
	}
	if o.GrossPrice.IsSet() {
		toSerialize["gross_price"] = o.GrossPrice.Get()
	}
	if o.Vat.IsSet() {
		toSerialize["vat"] = o.Vat.Get()
	}
	if o.NotTaxable.IsSet() {
		toSerialize["not_taxable"] = o.NotTaxable.Get()
	}
	if o.ApplyWithholdingTaxes.IsSet() {
		toSerialize["apply_withholding_taxes"] = o.ApplyWithholdingTaxes.Get()
	}
	if o.Discount.IsSet() {
		toSerialize["discount"] = o.Discount.Get()
	}
	if o.DiscountHighlight.IsSet() {
		toSerialize["discount_highlight"] = o.DiscountHighlight.Get()
	}
	if o.InDdt.IsSet() {
		toSerialize["in_ddt"] = o.InDdt.Get()
	}
	if o.Stock.IsSet() {
		toSerialize["stock"] = o.Stock.Get()
	}
	if o.EiRaw != nil {
		toSerialize["ei_raw"] = o.EiRaw
	}
	return json.Marshal(toSerialize)
}

type NullableIssuedDocumentItemsListItem struct {
	value *IssuedDocumentItemsListItem
	isSet bool
}

func (v NullableIssuedDocumentItemsListItem) Get() *IssuedDocumentItemsListItem {
	return v.value
}

func (v *NullableIssuedDocumentItemsListItem) Set(val *IssuedDocumentItemsListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentItemsListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentItemsListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentItemsListItem(val *IssuedDocumentItemsListItem) *NullableIssuedDocumentItemsListItem {
	return &NullableIssuedDocumentItemsListItem{value: val, isSet: true}
}

func (v NullableIssuedDocumentItemsListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentItemsListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


