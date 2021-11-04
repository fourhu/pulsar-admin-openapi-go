/*
Pulsar Packages REST API

This provides the REST API for Pulsar Packages operations

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PackageMetadata struct for PackageMetadata
type PackageMetadata struct {
	Description *string `json:"description,omitempty"`
	Contact *string `json:"contact,omitempty"`
	CreateTime *int64 `json:"createTime,omitempty"`
	ModificationTime *int64 `json:"modificationTime,omitempty"`
	Properties *map[string]string `json:"properties,omitempty"`
}

// NewPackageMetadata instantiates a new PackageMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageMetadata() *PackageMetadata {
	this := PackageMetadata{}
	return &this
}

// NewPackageMetadataWithDefaults instantiates a new PackageMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageMetadataWithDefaults() *PackageMetadata {
	this := PackageMetadata{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PackageMetadata) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PackageMetadata) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PackageMetadata) SetDescription(v string) {
	o.Description = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *PackageMetadata) GetContact() string {
	if o == nil || o.Contact == nil {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetContactOk() (*string, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *PackageMetadata) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *PackageMetadata) SetContact(v string) {
	o.Contact = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *PackageMetadata) GetCreateTime() int64 {
	if o == nil || o.CreateTime == nil {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetCreateTimeOk() (*int64, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *PackageMetadata) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *PackageMetadata) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetModificationTime returns the ModificationTime field value if set, zero value otherwise.
func (o *PackageMetadata) GetModificationTime() int64 {
	if o == nil || o.ModificationTime == nil {
		var ret int64
		return ret
	}
	return *o.ModificationTime
}

// GetModificationTimeOk returns a tuple with the ModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetModificationTimeOk() (*int64, bool) {
	if o == nil || o.ModificationTime == nil {
		return nil, false
	}
	return o.ModificationTime, true
}

// HasModificationTime returns a boolean if a field has been set.
func (o *PackageMetadata) HasModificationTime() bool {
	if o != nil && o.ModificationTime != nil {
		return true
	}

	return false
}

// SetModificationTime gets a reference to the given int64 and assigns it to the ModificationTime field.
func (o *PackageMetadata) SetModificationTime(v int64) {
	o.ModificationTime = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PackageMetadata) GetProperties() map[string]string {
	if o == nil || o.Properties == nil {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PackageMetadata) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *PackageMetadata) SetProperties(v map[string]string) {
	o.Properties = &v
}

func (o PackageMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.ModificationTime != nil {
		toSerialize["modificationTime"] = o.ModificationTime
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullablePackageMetadata struct {
	value *PackageMetadata
	isSet bool
}

func (v NullablePackageMetadata) Get() *PackageMetadata {
	return v.value
}

func (v *NullablePackageMetadata) Set(val *PackageMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageMetadata(val *PackageMetadata) *NullablePackageMetadata {
	return &NullablePackageMetadata{value: val, isSet: true}
}

func (v NullablePackageMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


