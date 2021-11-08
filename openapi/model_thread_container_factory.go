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

// ThreadContainerFactory struct for ThreadContainerFactory
type ThreadContainerFactory struct {
	ThreadGroupName *string `json:"threadGroupName,omitempty"`
	PulsarClientMemoryLimit *MemoryLimit `json:"pulsarClientMemoryLimit,omitempty"`
}

// NewThreadContainerFactory instantiates a new ThreadContainerFactory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadContainerFactory() *ThreadContainerFactory {
	this := ThreadContainerFactory{}
	return &this
}

// NewThreadContainerFactoryWithDefaults instantiates a new ThreadContainerFactory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadContainerFactoryWithDefaults() *ThreadContainerFactory {
	this := ThreadContainerFactory{}
	return &this
}

// GetThreadGroupName returns the ThreadGroupName field value if set, zero value otherwise.
func (o *ThreadContainerFactory) GetThreadGroupName() string {
	if o == nil || o.ThreadGroupName == nil {
		var ret string
		return ret
	}
	return *o.ThreadGroupName
}

// GetThreadGroupNameOk returns a tuple with the ThreadGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadContainerFactory) GetThreadGroupNameOk() (*string, bool) {
	if o == nil || o.ThreadGroupName == nil {
		return nil, false
	}
	return o.ThreadGroupName, true
}

// HasThreadGroupName returns a boolean if a field has been set.
func (o *ThreadContainerFactory) HasThreadGroupName() bool {
	if o != nil && o.ThreadGroupName != nil {
		return true
	}

	return false
}

// SetThreadGroupName gets a reference to the given string and assigns it to the ThreadGroupName field.
func (o *ThreadContainerFactory) SetThreadGroupName(v string) {
	o.ThreadGroupName = &v
}

// GetPulsarClientMemoryLimit returns the PulsarClientMemoryLimit field value if set, zero value otherwise.
func (o *ThreadContainerFactory) GetPulsarClientMemoryLimit() MemoryLimit {
	if o == nil || o.PulsarClientMemoryLimit == nil {
		var ret MemoryLimit
		return ret
	}
	return *o.PulsarClientMemoryLimit
}

// GetPulsarClientMemoryLimitOk returns a tuple with the PulsarClientMemoryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadContainerFactory) GetPulsarClientMemoryLimitOk() (*MemoryLimit, bool) {
	if o == nil || o.PulsarClientMemoryLimit == nil {
		return nil, false
	}
	return o.PulsarClientMemoryLimit, true
}

// HasPulsarClientMemoryLimit returns a boolean if a field has been set.
func (o *ThreadContainerFactory) HasPulsarClientMemoryLimit() bool {
	if o != nil && o.PulsarClientMemoryLimit != nil {
		return true
	}

	return false
}

// SetPulsarClientMemoryLimit gets a reference to the given MemoryLimit and assigns it to the PulsarClientMemoryLimit field.
func (o *ThreadContainerFactory) SetPulsarClientMemoryLimit(v MemoryLimit) {
	o.PulsarClientMemoryLimit = &v
}

func (o ThreadContainerFactory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ThreadGroupName != nil {
		toSerialize["threadGroupName"] = o.ThreadGroupName
	}
	if o.PulsarClientMemoryLimit != nil {
		toSerialize["pulsarClientMemoryLimit"] = o.PulsarClientMemoryLimit
	}
	return json.Marshal(toSerialize)
}

type NullableThreadContainerFactory struct {
	value *ThreadContainerFactory
	isSet bool
}

func (v NullableThreadContainerFactory) Get() *ThreadContainerFactory {
	return v.value
}

func (v *NullableThreadContainerFactory) Set(val *ThreadContainerFactory) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadContainerFactory) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadContainerFactory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadContainerFactory(val *ThreadContainerFactory) *NullableThreadContainerFactory {
	return &NullableThreadContainerFactory{value: val, isSet: true}
}

func (v NullableThreadContainerFactory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadContainerFactory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


