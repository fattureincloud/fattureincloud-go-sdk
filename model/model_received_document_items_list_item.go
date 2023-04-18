/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.27
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ReceivedDocumentItemsListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceivedDocumentItemsListItem{}

// ReceivedDocumentItemsListItem struct for ReceivedDocumentItemsListItem
type ReceivedDocumentItemsListItem struct {
	// Unique identifier.
	Id NullableInt32 `json:"id,omitempty"`
	// Unique identifier of the product
	ProductId NullableInt32 `json:"product_id,omitempty"`
	// Product code.
	Code NullableString `json:"code,omitempty"`
	// Product name.
	Name NullableString `json:"name,omitempty"`
	// Product measure.
	Measure NullableString `json:"measure,omitempty"`
	// Product net price.
	NetPrice NullableFloat32 `json:"net_price,omitempty"`
	// Product category.
	Category NullableString `json:"category,omitempty"`
	// Product quantity.
	Qty NullableFloat32 `json:"qty,omitempty"`
	Vat NullableVatType `json:"vat,omitempty"`
	// Number of items in stock
	Stock NullableFloat32 `json:"stock,omitempty"`
}

// NewReceivedDocumentItemsListItem instantiates a new ReceivedDocumentItemsListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivedDocumentItemsListItem() *ReceivedDocumentItemsListItem {
	this := ReceivedDocumentItemsListItem{}
	return &this
}

// NewReceivedDocumentItemsListItemWithDefaults instantiates a new ReceivedDocumentItemsListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivedDocumentItemsListItemWithDefaults() *ReceivedDocumentItemsListItem {
	this := ReceivedDocumentItemsListItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *ReceivedDocumentItemsListItem) SetId(v int32) *ReceivedDocumentItemsListItem {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetIdNil() *ReceivedDocumentItemsListItem {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetId() {
	o.Id.Unset()
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetProductId() int32 {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetProductIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableInt32 and assigns it to the ProductId field.
func (o *ReceivedDocumentItemsListItem) SetProductId(v int32) *ReceivedDocumentItemsListItem {
	o.ProductId.Set(&v)
	return o
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetProductIdNil() *ReceivedDocumentItemsListItem {
	o.ProductId.Set(nil)
	return o
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetProductId() {
	o.ProductId.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *ReceivedDocumentItemsListItem) SetCode(v string) *ReceivedDocumentItemsListItem {
	o.Code.Set(&v)
	return o
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetCodeNil() *ReceivedDocumentItemsListItem {
	o.Code.Set(nil)
	return o
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetCode() {
	o.Code.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ReceivedDocumentItemsListItem) SetName(v string) *ReceivedDocumentItemsListItem {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetNameNil() *ReceivedDocumentItemsListItem {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetName() {
	o.Name.Unset()
}

// GetMeasure returns the Measure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetMeasure() string {
	if o == nil || IsNil(o.Measure.Get()) {
		var ret string
		return ret
	}
	return *o.Measure.Get()
}

// GetMeasureOk returns a tuple with the Measure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Measure.Get(), o.Measure.IsSet()
}

// HasMeasure returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasMeasure() bool {
	if o != nil && o.Measure.IsSet() {
		return true
	}

	return false
}

// SetMeasure gets a reference to the given NullableString and assigns it to the Measure field.
func (o *ReceivedDocumentItemsListItem) SetMeasure(v string) *ReceivedDocumentItemsListItem {
	o.Measure.Set(&v)
	return o
}
// SetMeasureNil sets the value for Measure to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetMeasureNil() *ReceivedDocumentItemsListItem {
	o.Measure.Set(nil)
	return o
}

// UnsetMeasure ensures that no value is present for Measure, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetMeasure() {
	o.Measure.Unset()
}

// GetNetPrice returns the NetPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetNetPrice() float32 {
	if o == nil || IsNil(o.NetPrice.Get()) {
		var ret float32
		return ret
	}
	return *o.NetPrice.Get()
}

// GetNetPriceOk returns a tuple with the NetPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetNetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetPrice.Get(), o.NetPrice.IsSet()
}

// HasNetPrice returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasNetPrice() bool {
	if o != nil && o.NetPrice.IsSet() {
		return true
	}

	return false
}

// SetNetPrice gets a reference to the given NullableFloat32 and assigns it to the NetPrice field.
func (o *ReceivedDocumentItemsListItem) SetNetPrice(v float32) *ReceivedDocumentItemsListItem {
	o.NetPrice.Set(&v)
	return o
}
// SetNetPriceNil sets the value for NetPrice to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetNetPriceNil() *ReceivedDocumentItemsListItem {
	o.NetPrice.Set(nil)
	return o
}

// UnsetNetPrice ensures that no value is present for NetPrice, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetNetPrice() {
	o.NetPrice.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetCategory() string {
	if o == nil || IsNil(o.Category.Get()) {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *ReceivedDocumentItemsListItem) SetCategory(v string) *ReceivedDocumentItemsListItem {
	o.Category.Set(&v)
	return o
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetCategoryNil() *ReceivedDocumentItemsListItem {
	o.Category.Set(nil)
	return o
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetCategory() {
	o.Category.Unset()
}

// GetQty returns the Qty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetQty() float32 {
	if o == nil || IsNil(o.Qty.Get()) {
		var ret float32
		return ret
	}
	return *o.Qty.Get()
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Qty.Get(), o.Qty.IsSet()
}

// HasQty returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasQty() bool {
	if o != nil && o.Qty.IsSet() {
		return true
	}

	return false
}

// SetQty gets a reference to the given NullableFloat32 and assigns it to the Qty field.
func (o *ReceivedDocumentItemsListItem) SetQty(v float32) *ReceivedDocumentItemsListItem {
	o.Qty.Set(&v)
	return o
}
// SetQtyNil sets the value for Qty to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetQtyNil() *ReceivedDocumentItemsListItem {
	o.Qty.Set(nil)
	return o
}

// UnsetQty ensures that no value is present for Qty, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetQty() {
	o.Qty.Unset()
}

// GetVat returns the Vat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetVat() VatType {
	if o == nil || IsNil(o.Vat.Get()) {
		var ret VatType
		return ret
	}
	return *o.Vat.Get()
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetVatOk() (*VatType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vat.Get(), o.Vat.IsSet()
}

// HasVat returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasVat() bool {
	if o != nil && o.Vat.IsSet() {
		return true
	}

	return false
}

// SetVat gets a reference to the given NullableVatType and assigns it to the Vat field.
func (o *ReceivedDocumentItemsListItem) SetVat(v VatType) *ReceivedDocumentItemsListItem {
	o.Vat.Set(&v)
	return o
}
// SetVatNil sets the value for Vat to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetVatNil() *ReceivedDocumentItemsListItem {
	o.Vat.Set(nil)
	return o
}

// UnsetVat ensures that no value is present for Vat, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetVat() {
	o.Vat.Unset()
}

// GetStock returns the Stock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedDocumentItemsListItem) GetStock() float32 {
	if o == nil || IsNil(o.Stock.Get()) {
		var ret float32
		return ret
	}
	return *o.Stock.Get()
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedDocumentItemsListItem) GetStockOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stock.Get(), o.Stock.IsSet()
}

// HasStock returns a boolean if a field has been set.
func (o *ReceivedDocumentItemsListItem) HasStock() bool {
	if o != nil && o.Stock.IsSet() {
		return true
	}

	return false
}

// SetStock gets a reference to the given NullableFloat32 and assigns it to the Stock field.
func (o *ReceivedDocumentItemsListItem) SetStock(v float32) *ReceivedDocumentItemsListItem {
	o.Stock.Set(&v)
	return o
}
// SetStockNil sets the value for Stock to be an explicit nil
func (o *ReceivedDocumentItemsListItem) SetStockNil() *ReceivedDocumentItemsListItem {
	o.Stock.Set(nil)
	return o
}

// UnsetStock ensures that no value is present for Stock, not even an explicit nil
func (o *ReceivedDocumentItemsListItem) UnsetStock() {
	o.Stock.Unset()
}

func (o ReceivedDocumentItemsListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceivedDocumentItemsListItem) ToMap() (map[string]interface{}, error) {
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
	if o.Measure.IsSet() {
		toSerialize["measure"] = o.Measure.Get()
	}
	if o.NetPrice.IsSet() {
		toSerialize["net_price"] = o.NetPrice.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Qty.IsSet() {
		toSerialize["qty"] = o.Qty.Get()
	}
	if o.Vat.IsSet() {
		toSerialize["vat"] = o.Vat.Get()
	}
	if o.Stock.IsSet() {
		toSerialize["stock"] = o.Stock.Get()
	}
	return toSerialize, nil
}

type NullableReceivedDocumentItemsListItem struct {
	value *ReceivedDocumentItemsListItem
	isSet bool
}

func (v NullableReceivedDocumentItemsListItem) Get() *ReceivedDocumentItemsListItem {
	return v.value
}

func (v *NullableReceivedDocumentItemsListItem) Set(val *ReceivedDocumentItemsListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivedDocumentItemsListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivedDocumentItemsListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivedDocumentItemsListItem(val *ReceivedDocumentItemsListItem) *NullableReceivedDocumentItemsListItem {
	return &NullableReceivedDocumentItemsListItem{value: val, isSet: true}
}

func (v NullableReceivedDocumentItemsListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivedDocumentItemsListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


