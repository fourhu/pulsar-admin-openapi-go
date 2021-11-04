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

// TypeBindings struct for TypeBindings
type TypeBindings struct {
	Empty *bool `json:"empty,omitempty"`
	TypeParameters *[]JavaType `json:"typeParameters,omitempty"`
}

// NewTypeBindings instantiates a new TypeBindings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypeBindings() *TypeBindings {
	this := TypeBindings{}
	return &this
}

// NewTypeBindingsWithDefaults instantiates a new TypeBindings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypeBindingsWithDefaults() *TypeBindings {
	this := TypeBindings{}
	return &this
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *TypeBindings) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeBindings) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *TypeBindings) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *TypeBindings) SetEmpty(v bool) {
	o.Empty = &v
}

// GetTypeParameters returns the TypeParameters field value if set, zero value otherwise.
func (o *TypeBindings) GetTypeParameters() []JavaType {
	if o == nil || o.TypeParameters == nil {
		var ret []JavaType
		return ret
	}
	return *o.TypeParameters
}

// GetTypeParametersOk returns a tuple with the TypeParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeBindings) GetTypeParametersOk() (*[]JavaType, bool) {
	if o == nil || o.TypeParameters == nil {
		return nil, false
	}
	return o.TypeParameters, true
}

// HasTypeParameters returns a boolean if a field has been set.
func (o *TypeBindings) HasTypeParameters() bool {
	if o != nil && o.TypeParameters != nil {
		return true
	}

	return false
}

// SetTypeParameters gets a reference to the given []JavaType and assigns it to the TypeParameters field.
func (o *TypeBindings) SetTypeParameters(v []JavaType) {
	o.TypeParameters = &v
}

func (o TypeBindings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	if o.TypeParameters != nil {
		toSerialize["typeParameters"] = o.TypeParameters
	}
	return json.Marshal(toSerialize)
}

type NullableTypeBindings struct {
	value *TypeBindings
	isSet bool
}

func (v NullableTypeBindings) Get() *TypeBindings {
	return v.value
}

func (v *NullableTypeBindings) Set(val *TypeBindings) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeBindings) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeBindings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeBindings(val *TypeBindings) *NullableTypeBindings {
	return &NullableTypeBindings{value: val, isSet: true}
}

func (v NullableTypeBindings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeBindings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


