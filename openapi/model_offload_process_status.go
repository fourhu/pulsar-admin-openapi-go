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

// OffloadProcessStatus struct for OffloadProcessStatus
type OffloadProcessStatus struct {
	FirstUnoffloadedMessage *map[string]interface{} `json:"firstUnoffloadedMessage,omitempty"`
	LastError *string `json:"lastError,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewOffloadProcessStatus instantiates a new OffloadProcessStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffloadProcessStatus() *OffloadProcessStatus {
	this := OffloadProcessStatus{}
	return &this
}

// NewOffloadProcessStatusWithDefaults instantiates a new OffloadProcessStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffloadProcessStatusWithDefaults() *OffloadProcessStatus {
	this := OffloadProcessStatus{}
	return &this
}

// GetFirstUnoffloadedMessage returns the FirstUnoffloadedMessage field value if set, zero value otherwise.
func (o *OffloadProcessStatus) GetFirstUnoffloadedMessage() map[string]interface{} {
	if o == nil || o.FirstUnoffloadedMessage == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FirstUnoffloadedMessage
}

// GetFirstUnoffloadedMessageOk returns a tuple with the FirstUnoffloadedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OffloadProcessStatus) GetFirstUnoffloadedMessageOk() (*map[string]interface{}, bool) {
	if o == nil || o.FirstUnoffloadedMessage == nil {
		return nil, false
	}
	return o.FirstUnoffloadedMessage, true
}

// HasFirstUnoffloadedMessage returns a boolean if a field has been set.
func (o *OffloadProcessStatus) HasFirstUnoffloadedMessage() bool {
	if o != nil && o.FirstUnoffloadedMessage != nil {
		return true
	}

	return false
}

// SetFirstUnoffloadedMessage gets a reference to the given map[string]interface{} and assigns it to the FirstUnoffloadedMessage field.
func (o *OffloadProcessStatus) SetFirstUnoffloadedMessage(v map[string]interface{}) {
	o.FirstUnoffloadedMessage = &v
}

// GetLastError returns the LastError field value if set, zero value otherwise.
func (o *OffloadProcessStatus) GetLastError() string {
	if o == nil || o.LastError == nil {
		var ret string
		return ret
	}
	return *o.LastError
}

// GetLastErrorOk returns a tuple with the LastError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OffloadProcessStatus) GetLastErrorOk() (*string, bool) {
	if o == nil || o.LastError == nil {
		return nil, false
	}
	return o.LastError, true
}

// HasLastError returns a boolean if a field has been set.
func (o *OffloadProcessStatus) HasLastError() bool {
	if o != nil && o.LastError != nil {
		return true
	}

	return false
}

// SetLastError gets a reference to the given string and assigns it to the LastError field.
func (o *OffloadProcessStatus) SetLastError(v string) {
	o.LastError = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OffloadProcessStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OffloadProcessStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OffloadProcessStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OffloadProcessStatus) SetStatus(v string) {
	o.Status = &v
}

func (o OffloadProcessStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstUnoffloadedMessage != nil {
		toSerialize["firstUnoffloadedMessage"] = o.FirstUnoffloadedMessage
	}
	if o.LastError != nil {
		toSerialize["lastError"] = o.LastError
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableOffloadProcessStatus struct {
	value *OffloadProcessStatus
	isSet bool
}

func (v NullableOffloadProcessStatus) Get() *OffloadProcessStatus {
	return v.value
}

func (v *NullableOffloadProcessStatus) Set(val *OffloadProcessStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOffloadProcessStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOffloadProcessStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffloadProcessStatus(val *OffloadProcessStatus) *NullableOffloadProcessStatus {
	return &NullableOffloadProcessStatus{value: val, isSet: true}
}

func (v NullableOffloadProcessStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffloadProcessStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


