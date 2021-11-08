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

// InactiveTopicPolicies struct for InactiveTopicPolicies
type InactiveTopicPolicies struct {
	InactiveTopicDeleteMode *string `json:"inactiveTopicDeleteMode,omitempty"`
	MaxInactiveDurationSeconds *int32 `json:"maxInactiveDurationSeconds,omitempty"`
	DeleteWhileInactive *bool `json:"deleteWhileInactive,omitempty"`
}

// NewInactiveTopicPolicies instantiates a new InactiveTopicPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInactiveTopicPolicies() *InactiveTopicPolicies {
	this := InactiveTopicPolicies{}
	return &this
}

// NewInactiveTopicPoliciesWithDefaults instantiates a new InactiveTopicPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInactiveTopicPoliciesWithDefaults() *InactiveTopicPolicies {
	this := InactiveTopicPolicies{}
	return &this
}

// GetInactiveTopicDeleteMode returns the InactiveTopicDeleteMode field value if set, zero value otherwise.
func (o *InactiveTopicPolicies) GetInactiveTopicDeleteMode() string {
	if o == nil || o.InactiveTopicDeleteMode == nil {
		var ret string
		return ret
	}
	return *o.InactiveTopicDeleteMode
}

// GetInactiveTopicDeleteModeOk returns a tuple with the InactiveTopicDeleteMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InactiveTopicPolicies) GetInactiveTopicDeleteModeOk() (*string, bool) {
	if o == nil || o.InactiveTopicDeleteMode == nil {
		return nil, false
	}
	return o.InactiveTopicDeleteMode, true
}

// HasInactiveTopicDeleteMode returns a boolean if a field has been set.
func (o *InactiveTopicPolicies) HasInactiveTopicDeleteMode() bool {
	if o != nil && o.InactiveTopicDeleteMode != nil {
		return true
	}

	return false
}

// SetInactiveTopicDeleteMode gets a reference to the given string and assigns it to the InactiveTopicDeleteMode field.
func (o *InactiveTopicPolicies) SetInactiveTopicDeleteMode(v string) {
	o.InactiveTopicDeleteMode = &v
}

// GetMaxInactiveDurationSeconds returns the MaxInactiveDurationSeconds field value if set, zero value otherwise.
func (o *InactiveTopicPolicies) GetMaxInactiveDurationSeconds() int32 {
	if o == nil || o.MaxInactiveDurationSeconds == nil {
		var ret int32
		return ret
	}
	return *o.MaxInactiveDurationSeconds
}

// GetMaxInactiveDurationSecondsOk returns a tuple with the MaxInactiveDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InactiveTopicPolicies) GetMaxInactiveDurationSecondsOk() (*int32, bool) {
	if o == nil || o.MaxInactiveDurationSeconds == nil {
		return nil, false
	}
	return o.MaxInactiveDurationSeconds, true
}

// HasMaxInactiveDurationSeconds returns a boolean if a field has been set.
func (o *InactiveTopicPolicies) HasMaxInactiveDurationSeconds() bool {
	if o != nil && o.MaxInactiveDurationSeconds != nil {
		return true
	}

	return false
}

// SetMaxInactiveDurationSeconds gets a reference to the given int32 and assigns it to the MaxInactiveDurationSeconds field.
func (o *InactiveTopicPolicies) SetMaxInactiveDurationSeconds(v int32) {
	o.MaxInactiveDurationSeconds = &v
}

// GetDeleteWhileInactive returns the DeleteWhileInactive field value if set, zero value otherwise.
func (o *InactiveTopicPolicies) GetDeleteWhileInactive() bool {
	if o == nil || o.DeleteWhileInactive == nil {
		var ret bool
		return ret
	}
	return *o.DeleteWhileInactive
}

// GetDeleteWhileInactiveOk returns a tuple with the DeleteWhileInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InactiveTopicPolicies) GetDeleteWhileInactiveOk() (*bool, bool) {
	if o == nil || o.DeleteWhileInactive == nil {
		return nil, false
	}
	return o.DeleteWhileInactive, true
}

// HasDeleteWhileInactive returns a boolean if a field has been set.
func (o *InactiveTopicPolicies) HasDeleteWhileInactive() bool {
	if o != nil && o.DeleteWhileInactive != nil {
		return true
	}

	return false
}

// SetDeleteWhileInactive gets a reference to the given bool and assigns it to the DeleteWhileInactive field.
func (o *InactiveTopicPolicies) SetDeleteWhileInactive(v bool) {
	o.DeleteWhileInactive = &v
}

func (o InactiveTopicPolicies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InactiveTopicDeleteMode != nil {
		toSerialize["inactiveTopicDeleteMode"] = o.InactiveTopicDeleteMode
	}
	if o.MaxInactiveDurationSeconds != nil {
		toSerialize["maxInactiveDurationSeconds"] = o.MaxInactiveDurationSeconds
	}
	if o.DeleteWhileInactive != nil {
		toSerialize["deleteWhileInactive"] = o.DeleteWhileInactive
	}
	return json.Marshal(toSerialize)
}

type NullableInactiveTopicPolicies struct {
	value *InactiveTopicPolicies
	isSet bool
}

func (v NullableInactiveTopicPolicies) Get() *InactiveTopicPolicies {
	return v.value
}

func (v *NullableInactiveTopicPolicies) Set(val *InactiveTopicPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableInactiveTopicPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableInactiveTopicPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInactiveTopicPolicies(val *InactiveTopicPolicies) *NullableInactiveTopicPolicies {
	return &NullableInactiveTopicPolicies{value: val, isSet: true}
}

func (v NullableInactiveTopicPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInactiveTopicPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


