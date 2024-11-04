/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.3
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the ListProductsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListProductsResponse{}

// ListProductsResponse struct for ListProductsResponse
type ListProductsResponse struct {
	// Current page number.
	CurrentPage NullableInt32 `json:"current_page,omitempty"`
	// First page url.
	FirstPageUrl NullableString `json:"first_page_url,omitempty"`
	// First result of the page.
	From NullableInt32 `json:"from,omitempty"`
	// Last page number.
	LastPage NullableInt32 `json:"last_page,omitempty"`
	// Last page url.
	LastPageUrl NullableString `json:"last_page_url,omitempty"`
	// Next page url
	NextPageUrl NullableString `json:"next_page_url,omitempty"`
	// Request path.
	Path NullableString `json:"path,omitempty"`
	// Number of result per page.
	PerPage NullableInt32 `json:"per_page,omitempty"`
	// Previous page url.
	PrevPageUrl NullableString `json:"prev_page_url,omitempty"`
	// Last result of the page.
	To NullableInt32 `json:"to,omitempty"`
	// Total number of results
	Total NullableInt32 `json:"total,omitempty"`
	Data []Product `json:"data,omitempty"`
}

// NewListProductsResponse instantiates a new ListProductsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProductsResponse() *ListProductsResponse {
	this := ListProductsResponse{}
	return &this
}

// NewListProductsResponseWithDefaults instantiates a new ListProductsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProductsResponseWithDefaults() *ListProductsResponse {
	this := ListProductsResponse{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetCurrentPage() int32 {
	if o == nil || IsNil(o.CurrentPage.Get()) {
		var ret int32
		return ret
	}
	return *o.CurrentPage.Get()
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetCurrentPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPage.Get(), o.CurrentPage.IsSet()
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *ListProductsResponse) HasCurrentPage() bool {
	if o != nil && o.CurrentPage.IsSet() {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given NullableInt32 and assigns it to the CurrentPage field.
func (o *ListProductsResponse) SetCurrentPage(v int32) *ListProductsResponse {
	o.CurrentPage.Set(&v)
		return o
}
// SetCurrentPageNil sets the value for CurrentPage to be an explicit nil
func (o *ListProductsResponse) SetCurrentPageNil() *ListProductsResponse {
	o.CurrentPage.Set(nil)
	return o
}

// UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
func (o *ListProductsResponse) UnsetCurrentPage() {
	o.CurrentPage.Unset()
}

// GetFirstPageUrl returns the FirstPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetFirstPageUrl() string {
	if o == nil || IsNil(o.FirstPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.FirstPageUrl.Get()
}

// GetFirstPageUrlOk returns a tuple with the FirstPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetFirstPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstPageUrl.Get(), o.FirstPageUrl.IsSet()
}

// HasFirstPageUrl returns a boolean if a field has been set.
func (o *ListProductsResponse) HasFirstPageUrl() bool {
	if o != nil && o.FirstPageUrl.IsSet() {
		return true
	}

	return false
}

// SetFirstPageUrl gets a reference to the given NullableString and assigns it to the FirstPageUrl field.
func (o *ListProductsResponse) SetFirstPageUrl(v string) *ListProductsResponse {
	o.FirstPageUrl.Set(&v)
		return o
}
// SetFirstPageUrlNil sets the value for FirstPageUrl to be an explicit nil
func (o *ListProductsResponse) SetFirstPageUrlNil() *ListProductsResponse {
	o.FirstPageUrl.Set(nil)
	return o
}

// UnsetFirstPageUrl ensures that no value is present for FirstPageUrl, not even an explicit nil
func (o *ListProductsResponse) UnsetFirstPageUrl() {
	o.FirstPageUrl.Unset()
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetFrom() int32 {
	if o == nil || IsNil(o.From.Get()) {
		var ret int32
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetFromOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *ListProductsResponse) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableInt32 and assigns it to the From field.
func (o *ListProductsResponse) SetFrom(v int32) *ListProductsResponse {
	o.From.Set(&v)
		return o
}
// SetFromNil sets the value for From to be an explicit nil
func (o *ListProductsResponse) SetFromNil() *ListProductsResponse {
	o.From.Set(nil)
	return o
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *ListProductsResponse) UnsetFrom() {
	o.From.Unset()
}

// GetLastPage returns the LastPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetLastPage() int32 {
	if o == nil || IsNil(o.LastPage.Get()) {
		var ret int32
		return ret
	}
	return *o.LastPage.Get()
}

// GetLastPageOk returns a tuple with the LastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetLastPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPage.Get(), o.LastPage.IsSet()
}

// HasLastPage returns a boolean if a field has been set.
func (o *ListProductsResponse) HasLastPage() bool {
	if o != nil && o.LastPage.IsSet() {
		return true
	}

	return false
}

// SetLastPage gets a reference to the given NullableInt32 and assigns it to the LastPage field.
func (o *ListProductsResponse) SetLastPage(v int32) *ListProductsResponse {
	o.LastPage.Set(&v)
		return o
}
// SetLastPageNil sets the value for LastPage to be an explicit nil
func (o *ListProductsResponse) SetLastPageNil() *ListProductsResponse {
	o.LastPage.Set(nil)
	return o
}

// UnsetLastPage ensures that no value is present for LastPage, not even an explicit nil
func (o *ListProductsResponse) UnsetLastPage() {
	o.LastPage.Unset()
}

// GetLastPageUrl returns the LastPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetLastPageUrl() string {
	if o == nil || IsNil(o.LastPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LastPageUrl.Get()
}

// GetLastPageUrlOk returns a tuple with the LastPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetLastPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPageUrl.Get(), o.LastPageUrl.IsSet()
}

// HasLastPageUrl returns a boolean if a field has been set.
func (o *ListProductsResponse) HasLastPageUrl() bool {
	if o != nil && o.LastPageUrl.IsSet() {
		return true
	}

	return false
}

// SetLastPageUrl gets a reference to the given NullableString and assigns it to the LastPageUrl field.
func (o *ListProductsResponse) SetLastPageUrl(v string) *ListProductsResponse {
	o.LastPageUrl.Set(&v)
		return o
}
// SetLastPageUrlNil sets the value for LastPageUrl to be an explicit nil
func (o *ListProductsResponse) SetLastPageUrlNil() *ListProductsResponse {
	o.LastPageUrl.Set(nil)
	return o
}

// UnsetLastPageUrl ensures that no value is present for LastPageUrl, not even an explicit nil
func (o *ListProductsResponse) UnsetLastPageUrl() {
	o.LastPageUrl.Unset()
}

// GetNextPageUrl returns the NextPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetNextPageUrl() string {
	if o == nil || IsNil(o.NextPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageUrl.Get()
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetNextPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageUrl.Get(), o.NextPageUrl.IsSet()
}

// HasNextPageUrl returns a boolean if a field has been set.
func (o *ListProductsResponse) HasNextPageUrl() bool {
	if o != nil && o.NextPageUrl.IsSet() {
		return true
	}

	return false
}

// SetNextPageUrl gets a reference to the given NullableString and assigns it to the NextPageUrl field.
func (o *ListProductsResponse) SetNextPageUrl(v string) *ListProductsResponse {
	o.NextPageUrl.Set(&v)
		return o
}
// SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil
func (o *ListProductsResponse) SetNextPageUrlNil() *ListProductsResponse {
	o.NextPageUrl.Set(nil)
	return o
}

// UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
func (o *ListProductsResponse) UnsetNextPageUrl() {
	o.NextPageUrl.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *ListProductsResponse) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *ListProductsResponse) SetPath(v string) *ListProductsResponse {
	o.Path.Set(&v)
		return o
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *ListProductsResponse) SetPathNil() *ListProductsResponse {
	o.Path.Set(nil)
	return o
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *ListProductsResponse) UnsetPath() {
	o.Path.Unset()
}

// GetPerPage returns the PerPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetPerPage() int32 {
	if o == nil || IsNil(o.PerPage.Get()) {
		var ret int32
		return ret
	}
	return *o.PerPage.Get()
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerPage.Get(), o.PerPage.IsSet()
}

// HasPerPage returns a boolean if a field has been set.
func (o *ListProductsResponse) HasPerPage() bool {
	if o != nil && o.PerPage.IsSet() {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given NullableInt32 and assigns it to the PerPage field.
func (o *ListProductsResponse) SetPerPage(v int32) *ListProductsResponse {
	o.PerPage.Set(&v)
		return o
}
// SetPerPageNil sets the value for PerPage to be an explicit nil
func (o *ListProductsResponse) SetPerPageNil() *ListProductsResponse {
	o.PerPage.Set(nil)
	return o
}

// UnsetPerPage ensures that no value is present for PerPage, not even an explicit nil
func (o *ListProductsResponse) UnsetPerPage() {
	o.PerPage.Unset()
}

// GetPrevPageUrl returns the PrevPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetPrevPageUrl() string {
	if o == nil || IsNil(o.PrevPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PrevPageUrl.Get()
}

// GetPrevPageUrlOk returns a tuple with the PrevPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetPrevPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevPageUrl.Get(), o.PrevPageUrl.IsSet()
}

// HasPrevPageUrl returns a boolean if a field has been set.
func (o *ListProductsResponse) HasPrevPageUrl() bool {
	if o != nil && o.PrevPageUrl.IsSet() {
		return true
	}

	return false
}

// SetPrevPageUrl gets a reference to the given NullableString and assigns it to the PrevPageUrl field.
func (o *ListProductsResponse) SetPrevPageUrl(v string) *ListProductsResponse {
	o.PrevPageUrl.Set(&v)
		return o
}
// SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil
func (o *ListProductsResponse) SetPrevPageUrlNil() *ListProductsResponse {
	o.PrevPageUrl.Set(nil)
	return o
}

// UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
func (o *ListProductsResponse) UnsetPrevPageUrl() {
	o.PrevPageUrl.Unset()
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetTo() int32 {
	if o == nil || IsNil(o.To.Get()) {
		var ret int32
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *ListProductsResponse) HasTo() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableInt32 and assigns it to the To field.
func (o *ListProductsResponse) SetTo(v int32) *ListProductsResponse {
	o.To.Set(&v)
		return o
}
// SetToNil sets the value for To to be an explicit nil
func (o *ListProductsResponse) SetToNil() *ListProductsResponse {
	o.To.Set(nil)
	return o
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *ListProductsResponse) UnsetTo() {
	o.To.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetTotal() int32 {
	if o == nil || IsNil(o.Total.Get()) {
		var ret int32
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *ListProductsResponse) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableInt32 and assigns it to the Total field.
func (o *ListProductsResponse) SetTotal(v int32) *ListProductsResponse {
	o.Total.Set(&v)
		return o
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *ListProductsResponse) SetTotalNil() *ListProductsResponse {
	o.Total.Set(nil)
	return o
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *ListProductsResponse) UnsetTotal() {
	o.Total.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListProductsResponse) GetData() []Product {
	if o == nil {
		var ret []Product
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListProductsResponse) GetDataOk() ([]Product, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListProductsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Product and assigns it to the Data field.
func (o *ListProductsResponse) SetData(v []Product) *ListProductsResponse {
	o.Data = v
		return o
}

func (o ListProductsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListProductsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentPage.IsSet() {
		toSerialize["current_page"] = o.CurrentPage.Get()
	}
	if o.FirstPageUrl.IsSet() {
		toSerialize["first_page_url"] = o.FirstPageUrl.Get()
	}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.LastPage.IsSet() {
		toSerialize["last_page"] = o.LastPage.Get()
	}
	if o.LastPageUrl.IsSet() {
		toSerialize["last_page_url"] = o.LastPageUrl.Get()
	}
	if o.NextPageUrl.IsSet() {
		toSerialize["next_page_url"] = o.NextPageUrl.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.PerPage.IsSet() {
		toSerialize["per_page"] = o.PerPage.Get()
	}
	if o.PrevPageUrl.IsSet() {
		toSerialize["prev_page_url"] = o.PrevPageUrl.Get()
	}
	if o.To.IsSet() {
		toSerialize["to"] = o.To.Get()
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListProductsResponse struct {
	value *ListProductsResponse
	isSet bool
}

func (v NullableListProductsResponse) Get() *ListProductsResponse {
	return v.value
}

func (v *NullableListProductsResponse) Set(val *ListProductsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListProductsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListProductsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProductsResponse(val *ListProductsResponse) *NullableListProductsResponse {
	return &NullableListProductsResponse{value: val, isSet: true}
}

func (v NullableListProductsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProductsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


