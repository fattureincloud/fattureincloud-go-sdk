/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.21
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// Product struct for Product
type Product struct {
	// Unique identifier.
	Id NullableInt32 `json:"id,omitempty"`
	// Product name.
	Name NullableString `json:"name,omitempty"`
	// Product code.
	Code NullableString `json:"code,omitempty"`
	// Net sale price (used if use_gross_price is false, otherwise it's competed automatically).
	NetPrice NullableFloat32 `json:"net_price,omitempty"`
	// Gross sale price (used if use_gross_price is false, otherwise it's competed automatically).
	GrossPrice NullableFloat32 `json:"gross_price,omitempty"`
	// Determine which price to use for calculations.
	UseGrossPrice NullableBool `json:"use_gross_price,omitempty"`
	DefaultVat NullableVatType `json:"default_vat,omitempty"`
	// Net cost of the product (used for received documents).
	NetCost NullableFloat32 `json:"net_cost,omitempty"`
	// Unit of measure.
	Measure NullableString `json:"measure,omitempty"`
	// Product description.
	Description NullableString `json:"description,omitempty"`
	// Product category.
	Category NullableString `json:"category,omitempty"`
	// Extra notes.
	Notes NullableString `json:"notes,omitempty"`
	// Determine if the product is in stock.
	InStock NullableBool `json:"in_stock,omitempty"`
	// Product initial stock.
	StockInitial NullableFloat32 `json:"stock_initial,omitempty"`
	// [Read Only] Product current stock.
	StockCurrent NullableFloat32 `json:"stock_current,omitempty"`
	// Product average cost.
	AverageCost NullableFloat32 `json:"average_cost,omitempty"`
	// Product average price.
	AveragePrice NullableFloat32 `json:"average_price,omitempty"`
	CreatedAt NullableString `json:"created_at,omitempty"`
	UpdatedAt NullableString `json:"updated_at,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct() *Product {
	this := Product{}
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Product) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Product) SetId(v int32) *Product {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Product) SetIdNil() *Product {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Product) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Product) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Product) SetName(v string) *Product {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Product) SetNameNil() *Product {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Product) UnsetName() {
	o.Name.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *Product) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *Product) SetCode(v string) *Product {
	o.Code.Set(&v)
	return o
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *Product) SetCodeNil() *Product {
	o.Code.Set(nil)
	return o
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *Product) UnsetCode() {
	o.Code.Unset()
}

// GetNetPrice returns the NetPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetNetPrice() float32 {
	if o == nil || o.NetPrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.NetPrice.Get()
}

// GetNetPriceOk returns a tuple with the NetPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetNetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetPrice.Get(), o.NetPrice.IsSet()
}

// HasNetPrice returns a boolean if a field has been set.
func (o *Product) HasNetPrice() bool {
	if o != nil && o.NetPrice.IsSet() {
		return true
	}

	return false
}

// SetNetPrice gets a reference to the given NullableFloat32 and assigns it to the NetPrice field.
func (o *Product) SetNetPrice(v float32) *Product {
	o.NetPrice.Set(&v)
	return o
}
// SetNetPriceNil sets the value for NetPrice to be an explicit nil
func (o *Product) SetNetPriceNil() *Product {
	o.NetPrice.Set(nil)
	return o
}

// UnsetNetPrice ensures that no value is present for NetPrice, not even an explicit nil
func (o *Product) UnsetNetPrice() {
	o.NetPrice.Unset()
}

// GetGrossPrice returns the GrossPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetGrossPrice() float32 {
	if o == nil || o.GrossPrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.GrossPrice.Get()
}

// GetGrossPriceOk returns a tuple with the GrossPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetGrossPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrossPrice.Get(), o.GrossPrice.IsSet()
}

// HasGrossPrice returns a boolean if a field has been set.
func (o *Product) HasGrossPrice() bool {
	if o != nil && o.GrossPrice.IsSet() {
		return true
	}

	return false
}

// SetGrossPrice gets a reference to the given NullableFloat32 and assigns it to the GrossPrice field.
func (o *Product) SetGrossPrice(v float32) *Product {
	o.GrossPrice.Set(&v)
	return o
}
// SetGrossPriceNil sets the value for GrossPrice to be an explicit nil
func (o *Product) SetGrossPriceNil() *Product {
	o.GrossPrice.Set(nil)
	return o
}

// UnsetGrossPrice ensures that no value is present for GrossPrice, not even an explicit nil
func (o *Product) UnsetGrossPrice() {
	o.GrossPrice.Unset()
}

// GetUseGrossPrice returns the UseGrossPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetUseGrossPrice() bool {
	if o == nil || o.UseGrossPrice.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UseGrossPrice.Get()
}

// GetUseGrossPriceOk returns a tuple with the UseGrossPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetUseGrossPriceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseGrossPrice.Get(), o.UseGrossPrice.IsSet()
}

// HasUseGrossPrice returns a boolean if a field has been set.
func (o *Product) HasUseGrossPrice() bool {
	if o != nil && o.UseGrossPrice.IsSet() {
		return true
	}

	return false
}

// SetUseGrossPrice gets a reference to the given NullableBool and assigns it to the UseGrossPrice field.
func (o *Product) SetUseGrossPrice(v bool) *Product {
	o.UseGrossPrice.Set(&v)
	return o
}
// SetUseGrossPriceNil sets the value for UseGrossPrice to be an explicit nil
func (o *Product) SetUseGrossPriceNil() *Product {
	o.UseGrossPrice.Set(nil)
	return o
}

// UnsetUseGrossPrice ensures that no value is present for UseGrossPrice, not even an explicit nil
func (o *Product) UnsetUseGrossPrice() {
	o.UseGrossPrice.Unset()
}

// GetDefaultVat returns the DefaultVat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetDefaultVat() VatType {
	if o == nil || o.DefaultVat.Get() == nil {
		var ret VatType
		return ret
	}
	return *o.DefaultVat.Get()
}

// GetDefaultVatOk returns a tuple with the DefaultVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetDefaultVatOk() (*VatType, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultVat.Get(), o.DefaultVat.IsSet()
}

// HasDefaultVat returns a boolean if a field has been set.
func (o *Product) HasDefaultVat() bool {
	if o != nil && o.DefaultVat.IsSet() {
		return true
	}

	return false
}

// SetDefaultVat gets a reference to the given NullableVatType and assigns it to the DefaultVat field.
func (o *Product) SetDefaultVat(v VatType) *Product {
	o.DefaultVat.Set(&v)
	return o
}
// SetDefaultVatNil sets the value for DefaultVat to be an explicit nil
func (o *Product) SetDefaultVatNil() *Product {
	o.DefaultVat.Set(nil)
	return o
}

// UnsetDefaultVat ensures that no value is present for DefaultVat, not even an explicit nil
func (o *Product) UnsetDefaultVat() {
	o.DefaultVat.Unset()
}

// GetNetCost returns the NetCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetNetCost() float32 {
	if o == nil || o.NetCost.Get() == nil {
		var ret float32
		return ret
	}
	return *o.NetCost.Get()
}

// GetNetCostOk returns a tuple with the NetCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetNetCostOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetCost.Get(), o.NetCost.IsSet()
}

// HasNetCost returns a boolean if a field has been set.
func (o *Product) HasNetCost() bool {
	if o != nil && o.NetCost.IsSet() {
		return true
	}

	return false
}

// SetNetCost gets a reference to the given NullableFloat32 and assigns it to the NetCost field.
func (o *Product) SetNetCost(v float32) *Product {
	o.NetCost.Set(&v)
	return o
}
// SetNetCostNil sets the value for NetCost to be an explicit nil
func (o *Product) SetNetCostNil() *Product {
	o.NetCost.Set(nil)
	return o
}

// UnsetNetCost ensures that no value is present for NetCost, not even an explicit nil
func (o *Product) UnsetNetCost() {
	o.NetCost.Unset()
}

// GetMeasure returns the Measure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetMeasure() string {
	if o == nil || o.Measure.Get() == nil {
		var ret string
		return ret
	}
	return *o.Measure.Get()
}

// GetMeasureOk returns a tuple with the Measure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Measure.Get(), o.Measure.IsSet()
}

// HasMeasure returns a boolean if a field has been set.
func (o *Product) HasMeasure() bool {
	if o != nil && o.Measure.IsSet() {
		return true
	}

	return false
}

// SetMeasure gets a reference to the given NullableString and assigns it to the Measure field.
func (o *Product) SetMeasure(v string) *Product {
	o.Measure.Set(&v)
	return o
}
// SetMeasureNil sets the value for Measure to be an explicit nil
func (o *Product) SetMeasureNil() *Product {
	o.Measure.Set(nil)
	return o
}

// UnsetMeasure ensures that no value is present for Measure, not even an explicit nil
func (o *Product) UnsetMeasure() {
	o.Measure.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Product) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Product) SetDescription(v string) *Product {
	o.Description.Set(&v)
	return o
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Product) SetDescriptionNil() *Product {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Product) UnsetDescription() {
	o.Description.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *Product) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *Product) SetCategory(v string) *Product {
	o.Category.Set(&v)
	return o
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *Product) SetCategoryNil() *Product {
	o.Category.Set(nil)
	return o
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *Product) UnsetCategory() {
	o.Category.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *Product) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *Product) SetNotes(v string) *Product {
	o.Notes.Set(&v)
	return o
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *Product) SetNotesNil() *Product {
	o.Notes.Set(nil)
	return o
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *Product) UnsetNotes() {
	o.Notes.Unset()
}

// GetInStock returns the InStock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetInStock() bool {
	if o == nil || o.InStock.Get() == nil {
		var ret bool
		return ret
	}
	return *o.InStock.Get()
}

// GetInStockOk returns a tuple with the InStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetInStockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InStock.Get(), o.InStock.IsSet()
}

// HasInStock returns a boolean if a field has been set.
func (o *Product) HasInStock() bool {
	if o != nil && o.InStock.IsSet() {
		return true
	}

	return false
}

// SetInStock gets a reference to the given NullableBool and assigns it to the InStock field.
func (o *Product) SetInStock(v bool) *Product {
	o.InStock.Set(&v)
	return o
}
// SetInStockNil sets the value for InStock to be an explicit nil
func (o *Product) SetInStockNil() *Product {
	o.InStock.Set(nil)
	return o
}

// UnsetInStock ensures that no value is present for InStock, not even an explicit nil
func (o *Product) UnsetInStock() {
	o.InStock.Unset()
}

// GetStockInitial returns the StockInitial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetStockInitial() float32 {
	if o == nil || o.StockInitial.Get() == nil {
		var ret float32
		return ret
	}
	return *o.StockInitial.Get()
}

// GetStockInitialOk returns a tuple with the StockInitial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetStockInitialOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StockInitial.Get(), o.StockInitial.IsSet()
}

// HasStockInitial returns a boolean if a field has been set.
func (o *Product) HasStockInitial() bool {
	if o != nil && o.StockInitial.IsSet() {
		return true
	}

	return false
}

// SetStockInitial gets a reference to the given NullableFloat32 and assigns it to the StockInitial field.
func (o *Product) SetStockInitial(v float32) *Product {
	o.StockInitial.Set(&v)
	return o
}
// SetStockInitialNil sets the value for StockInitial to be an explicit nil
func (o *Product) SetStockInitialNil() *Product {
	o.StockInitial.Set(nil)
	return o
}

// UnsetStockInitial ensures that no value is present for StockInitial, not even an explicit nil
func (o *Product) UnsetStockInitial() {
	o.StockInitial.Unset()
}

// GetStockCurrent returns the StockCurrent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetStockCurrent() float32 {
	if o == nil || o.StockCurrent.Get() == nil {
		var ret float32
		return ret
	}
	return *o.StockCurrent.Get()
}

// GetStockCurrentOk returns a tuple with the StockCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetStockCurrentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StockCurrent.Get(), o.StockCurrent.IsSet()
}

// HasStockCurrent returns a boolean if a field has been set.
func (o *Product) HasStockCurrent() bool {
	if o != nil && o.StockCurrent.IsSet() {
		return true
	}

	return false
}

// SetStockCurrent gets a reference to the given NullableFloat32 and assigns it to the StockCurrent field.
func (o *Product) SetStockCurrent(v float32) *Product {
	o.StockCurrent.Set(&v)
	return o
}
// SetStockCurrentNil sets the value for StockCurrent to be an explicit nil
func (o *Product) SetStockCurrentNil() *Product {
	o.StockCurrent.Set(nil)
	return o
}

// UnsetStockCurrent ensures that no value is present for StockCurrent, not even an explicit nil
func (o *Product) UnsetStockCurrent() {
	o.StockCurrent.Unset()
}

// GetAverageCost returns the AverageCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetAverageCost() float32 {
	if o == nil || o.AverageCost.Get() == nil {
		var ret float32
		return ret
	}
	return *o.AverageCost.Get()
}

// GetAverageCostOk returns a tuple with the AverageCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetAverageCostOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AverageCost.Get(), o.AverageCost.IsSet()
}

// HasAverageCost returns a boolean if a field has been set.
func (o *Product) HasAverageCost() bool {
	if o != nil && o.AverageCost.IsSet() {
		return true
	}

	return false
}

// SetAverageCost gets a reference to the given NullableFloat32 and assigns it to the AverageCost field.
func (o *Product) SetAverageCost(v float32) *Product {
	o.AverageCost.Set(&v)
	return o
}
// SetAverageCostNil sets the value for AverageCost to be an explicit nil
func (o *Product) SetAverageCostNil() *Product {
	o.AverageCost.Set(nil)
	return o
}

// UnsetAverageCost ensures that no value is present for AverageCost, not even an explicit nil
func (o *Product) UnsetAverageCost() {
	o.AverageCost.Unset()
}

// GetAveragePrice returns the AveragePrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetAveragePrice() float32 {
	if o == nil || o.AveragePrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.AveragePrice.Get()
}

// GetAveragePriceOk returns a tuple with the AveragePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetAveragePriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AveragePrice.Get(), o.AveragePrice.IsSet()
}

// HasAveragePrice returns a boolean if a field has been set.
func (o *Product) HasAveragePrice() bool {
	if o != nil && o.AveragePrice.IsSet() {
		return true
	}

	return false
}

// SetAveragePrice gets a reference to the given NullableFloat32 and assigns it to the AveragePrice field.
func (o *Product) SetAveragePrice(v float32) *Product {
	o.AveragePrice.Set(&v)
	return o
}
// SetAveragePriceNil sets the value for AveragePrice to be an explicit nil
func (o *Product) SetAveragePriceNil() *Product {
	o.AveragePrice.Set(nil)
	return o
}

// UnsetAveragePrice ensures that no value is present for AveragePrice, not even an explicit nil
func (o *Product) UnsetAveragePrice() {
	o.AveragePrice.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Product) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *Product) SetCreatedAt(v string) *Product {
	o.CreatedAt.Set(&v)
	return o
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Product) SetCreatedAtNil() *Product {
	o.CreatedAt.Set(nil)
	return o
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Product) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Product) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *Product) SetUpdatedAt(v string) *Product {
	o.UpdatedAt.Set(&v)
	return o
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Product) SetUpdatedAtNil() *Product {
	o.UpdatedAt.Set(nil)
	return o
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Product) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.NetPrice.IsSet() {
		toSerialize["net_price"] = o.NetPrice.Get()
	}
	if o.GrossPrice.IsSet() {
		toSerialize["gross_price"] = o.GrossPrice.Get()
	}
	if o.UseGrossPrice.IsSet() {
		toSerialize["use_gross_price"] = o.UseGrossPrice.Get()
	}
	if o.DefaultVat.IsSet() {
		toSerialize["default_vat"] = o.DefaultVat.Get()
	}
	if o.NetCost.IsSet() {
		toSerialize["net_cost"] = o.NetCost.Get()
	}
	if o.Measure.IsSet() {
		toSerialize["measure"] = o.Measure.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.InStock.IsSet() {
		toSerialize["in_stock"] = o.InStock.Get()
	}
	if o.StockInitial.IsSet() {
		toSerialize["stock_initial"] = o.StockInitial.Get()
	}
	if o.StockCurrent.IsSet() {
		toSerialize["stock_current"] = o.StockCurrent.Get()
	}
	if o.AverageCost.IsSet() {
		toSerialize["average_cost"] = o.AverageCost.Get()
	}
	if o.AveragePrice.IsSet() {
		toSerialize["average_price"] = o.AveragePrice.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


