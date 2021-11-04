/*
Pulsar Admin REST API

This provides the REST API for admin operations

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ResourceDescription struct for ResourceDescription
type ResourceDescription struct {
	ResourceUsage *map[string]ResourceUsage `json:"resourceUsage,omitempty"`
	UsagePct *int32 `json:"usagePct,omitempty"`
}

// NewResourceDescription instantiates a new ResourceDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceDescription() *ResourceDescription {
	this := ResourceDescription{}
	return &this
}

// NewResourceDescriptionWithDefaults instantiates a new ResourceDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceDescriptionWithDefaults() *ResourceDescription {
	this := ResourceDescription{}
	return &this
}

// GetResourceUsage returns the ResourceUsage field value if set, zero value otherwise.
func (o *ResourceDescription) GetResourceUsage() map[string]ResourceUsage {
	if o == nil || o.ResourceUsage == nil {
		var ret map[string]ResourceUsage
		return ret
	}
	return *o.ResourceUsage
}

// GetResourceUsageOk returns a tuple with the ResourceUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceDescription) GetResourceUsageOk() (*map[string]ResourceUsage, bool) {
	if o == nil || o.ResourceUsage == nil {
		return nil, false
	}
	return o.ResourceUsage, true
}

// HasResourceUsage returns a boolean if a field has been set.
func (o *ResourceDescription) HasResourceUsage() bool {
	if o != nil && o.ResourceUsage != nil {
		return true
	}

	return false
}

// SetResourceUsage gets a reference to the given map[string]ResourceUsage and assigns it to the ResourceUsage field.
func (o *ResourceDescription) SetResourceUsage(v map[string]ResourceUsage) {
	o.ResourceUsage = &v
}

// GetUsagePct returns the UsagePct field value if set, zero value otherwise.
func (o *ResourceDescription) GetUsagePct() int32 {
	if o == nil || o.UsagePct == nil {
		var ret int32
		return ret
	}
	return *o.UsagePct
}

// GetUsagePctOk returns a tuple with the UsagePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceDescription) GetUsagePctOk() (*int32, bool) {
	if o == nil || o.UsagePct == nil {
		return nil, false
	}
	return o.UsagePct, true
}

// HasUsagePct returns a boolean if a field has been set.
func (o *ResourceDescription) HasUsagePct() bool {
	if o != nil && o.UsagePct != nil {
		return true
	}

	return false
}

// SetUsagePct gets a reference to the given int32 and assigns it to the UsagePct field.
func (o *ResourceDescription) SetUsagePct(v int32) {
	o.UsagePct = &v
}

func (o ResourceDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceUsage != nil {
		toSerialize["resourceUsage"] = o.ResourceUsage
	}
	if o.UsagePct != nil {
		toSerialize["usagePct"] = o.UsagePct
	}
	return json.Marshal(toSerialize)
}

type NullableResourceDescription struct {
	value *ResourceDescription
	isSet bool
}

func (v NullableResourceDescription) Get() *ResourceDescription {
	return v.value
}

func (v *NullableResourceDescription) Set(val *ResourceDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceDescription(val *ResourceDescription) *NullableResourceDescription {
	return &NullableResourceDescription{value: val, isSet: true}
}

func (v NullableResourceDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


