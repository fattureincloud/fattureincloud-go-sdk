/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// ArchiveDocument struct for ArchiveDocument
type ArchiveDocument struct {
	// Archive document unique identifier.
	Id NullableInt32 `json:"id,omitempty"`
	// Archive document date.
	Date NullableString `json:"date,omitempty"`
	// Archive Document description.
	Description NullableString `json:"description,omitempty"`
	// [Read Only] Absolute url of the attached file. Authomatically set if a valid attachment token is passed via POST /archive or PUT /archive/{documentId}.
	AttachmentUrl NullableString `json:"attachment_url,omitempty"`
	// Archive document category.
	Category NullableString `json:"category,omitempty"`
	// [Write Only]  [Required] Attachment token returned by POST /archive/attachment. Used to attach the file already uploaded.
	AttachmentToken NullableString `json:"attachment_token,omitempty"`
}

// NewArchiveDocument instantiates a new ArchiveDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveDocument() *ArchiveDocument {
	this := ArchiveDocument{}
	return &this
}

// NewArchiveDocumentWithDefaults instantiates a new ArchiveDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveDocumentWithDefaults() *ArchiveDocument {
	this := ArchiveDocument{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchiveDocument) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchiveDocument) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ArchiveDocument) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *ArchiveDocument) SetId(v int32) *ArchiveDocument {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ArchiveDocument) SetIdNil() *ArchiveDocument {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ArchiveDocument) UnsetId() {
	o.Id.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchiveDocument) GetDate() string {
	if o == nil || o.Date.Get() == nil {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchiveDocument) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *ArchiveDocument) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *ArchiveDocument) SetDate(v string) *ArchiveDocument {
	o.Date.Set(&v)
	return o
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *ArchiveDocument) SetDateNil() *ArchiveDocument {
	o.Date.Set(nil)
	return o
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *ArchiveDocument) UnsetDate() {
	o.Date.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchiveDocument) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchiveDocument) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ArchiveDocument) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ArchiveDocument) SetDescription(v string) *ArchiveDocument {
	o.Description.Set(&v)
	return o
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ArchiveDocument) SetDescriptionNil() *ArchiveDocument {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ArchiveDocument) UnsetDescription() {
	o.Description.Unset()
}

// GetAttachmentUrl returns the AttachmentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchiveDocument) GetAttachmentUrl() string {
	if o == nil || o.AttachmentUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.AttachmentUrl.Get()
}

// GetAttachmentUrlOk returns a tuple with the AttachmentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchiveDocument) GetAttachmentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentUrl.Get(), o.AttachmentUrl.IsSet()
}

// HasAttachmentUrl returns a boolean if a field has been set.
func (o *ArchiveDocument) HasAttachmentUrl() bool {
	if o != nil && o.AttachmentUrl.IsSet() {
		return true
	}

	return false
}

// SetAttachmentUrl gets a reference to the given NullableString and assigns it to the AttachmentUrl field.
func (o *ArchiveDocument) SetAttachmentUrl(v string) *ArchiveDocument {
	o.AttachmentUrl.Set(&v)
	return o
}
// SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil
func (o *ArchiveDocument) SetAttachmentUrlNil() *ArchiveDocument {
	o.AttachmentUrl.Set(nil)
	return o
}

// UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
func (o *ArchiveDocument) UnsetAttachmentUrl() {
	o.AttachmentUrl.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchiveDocument) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchiveDocument) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ArchiveDocument) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *ArchiveDocument) SetCategory(v string) *ArchiveDocument {
	o.Category.Set(&v)
	return o
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ArchiveDocument) SetCategoryNil() *ArchiveDocument {
	o.Category.Set(nil)
	return o
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ArchiveDocument) UnsetCategory() {
	o.Category.Unset()
}

// GetAttachmentToken returns the AttachmentToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchiveDocument) GetAttachmentToken() string {
	if o == nil || o.AttachmentToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.AttachmentToken.Get()
}

// GetAttachmentTokenOk returns a tuple with the AttachmentToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchiveDocument) GetAttachmentTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentToken.Get(), o.AttachmentToken.IsSet()
}

// HasAttachmentToken returns a boolean if a field has been set.
func (o *ArchiveDocument) HasAttachmentToken() bool {
	if o != nil && o.AttachmentToken.IsSet() {
		return true
	}

	return false
}

// SetAttachmentToken gets a reference to the given NullableString and assigns it to the AttachmentToken field.
func (o *ArchiveDocument) SetAttachmentToken(v string) *ArchiveDocument {
	o.AttachmentToken.Set(&v)
	return o
}
// SetAttachmentTokenNil sets the value for AttachmentToken to be an explicit nil
func (o *ArchiveDocument) SetAttachmentTokenNil() *ArchiveDocument {
	o.AttachmentToken.Set(nil)
	return o
}

// UnsetAttachmentToken ensures that no value is present for AttachmentToken, not even an explicit nil
func (o *ArchiveDocument) UnsetAttachmentToken() {
	o.AttachmentToken.Unset()
}

func (o ArchiveDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.AttachmentUrl.IsSet() {
		toSerialize["attachment_url"] = o.AttachmentUrl.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.AttachmentToken.IsSet() {
		toSerialize["attachment_token"] = o.AttachmentToken.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableArchiveDocument struct {
	value *ArchiveDocument
	isSet bool
}

func (v NullableArchiveDocument) Get() *ArchiveDocument {
	return v.value
}

func (v *NullableArchiveDocument) Set(val *ArchiveDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveDocument(val *ArchiveDocument) *NullableArchiveDocument {
	return &NullableArchiveDocument{value: val, isSet: true}
}

func (v NullableArchiveDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


