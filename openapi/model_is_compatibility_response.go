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

// IsCompatibilityResponse struct for IsCompatibilityResponse
type IsCompatibilityResponse struct {
	SchemaCompatibilityStrategy *string `json:"schemaCompatibilityStrategy,omitempty"`
	Compatibility *bool `json:"compatibility,omitempty"`
}

// NewIsCompatibilityResponse instantiates a new IsCompatibilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsCompatibilityResponse() *IsCompatibilityResponse {
	this := IsCompatibilityResponse{}
	return &this
}

// NewIsCompatibilityResponseWithDefaults instantiates a new IsCompatibilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsCompatibilityResponseWithDefaults() *IsCompatibilityResponse {
	this := IsCompatibilityResponse{}
	return &this
}

// GetSchemaCompatibilityStrategy returns the SchemaCompatibilityStrategy field value if set, zero value otherwise.
func (o *IsCompatibilityResponse) GetSchemaCompatibilityStrategy() string {
	if o == nil || o.SchemaCompatibilityStrategy == nil {
		var ret string
		return ret
	}
	return *o.SchemaCompatibilityStrategy
}

// GetSchemaCompatibilityStrategyOk returns a tuple with the SchemaCompatibilityStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsCompatibilityResponse) GetSchemaCompatibilityStrategyOk() (*string, bool) {
	if o == nil || o.SchemaCompatibilityStrategy == nil {
		return nil, false
	}
	return o.SchemaCompatibilityStrategy, true
}

// HasSchemaCompatibilityStrategy returns a boolean if a field has been set.
func (o *IsCompatibilityResponse) HasSchemaCompatibilityStrategy() bool {
	if o != nil && o.SchemaCompatibilityStrategy != nil {
		return true
	}

	return false
}

// SetSchemaCompatibilityStrategy gets a reference to the given string and assigns it to the SchemaCompatibilityStrategy field.
func (o *IsCompatibilityResponse) SetSchemaCompatibilityStrategy(v string) {
	o.SchemaCompatibilityStrategy = &v
}

// GetCompatibility returns the Compatibility field value if set, zero value otherwise.
func (o *IsCompatibilityResponse) GetCompatibility() bool {
	if o == nil || o.Compatibility == nil {
		var ret bool
		return ret
	}
	return *o.Compatibility
}

// GetCompatibilityOk returns a tuple with the Compatibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsCompatibilityResponse) GetCompatibilityOk() (*bool, bool) {
	if o == nil || o.Compatibility == nil {
		return nil, false
	}
	return o.Compatibility, true
}

// HasCompatibility returns a boolean if a field has been set.
func (o *IsCompatibilityResponse) HasCompatibility() bool {
	if o != nil && o.Compatibility != nil {
		return true
	}

	return false
}

// SetCompatibility gets a reference to the given bool and assigns it to the Compatibility field.
func (o *IsCompatibilityResponse) SetCompatibility(v bool) {
	o.Compatibility = &v
}

func (o IsCompatibilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SchemaCompatibilityStrategy != nil {
		toSerialize["schemaCompatibilityStrategy"] = o.SchemaCompatibilityStrategy
	}
	if o.Compatibility != nil {
		toSerialize["compatibility"] = o.Compatibility
	}
	return json.Marshal(toSerialize)
}

type NullableIsCompatibilityResponse struct {
	value *IsCompatibilityResponse
	isSet bool
}

func (v NullableIsCompatibilityResponse) Get() *IsCompatibilityResponse {
	return v.value
}

func (v *NullableIsCompatibilityResponse) Set(val *IsCompatibilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIsCompatibilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIsCompatibilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsCompatibilityResponse(val *IsCompatibilityResponse) *NullableIsCompatibilityResponse {
	return &NullableIsCompatibilityResponse{value: val, isSet: true}
}

func (v NullableIsCompatibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsCompatibilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


