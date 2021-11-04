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

// DeleteSchemaResponse struct for DeleteSchemaResponse
type DeleteSchemaResponse struct {
	Version *int64 `json:"version,omitempty"`
}

// NewDeleteSchemaResponse instantiates a new DeleteSchemaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSchemaResponse() *DeleteSchemaResponse {
	this := DeleteSchemaResponse{}
	return &this
}

// NewDeleteSchemaResponseWithDefaults instantiates a new DeleteSchemaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSchemaResponseWithDefaults() *DeleteSchemaResponse {
	this := DeleteSchemaResponse{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeleteSchemaResponse) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSchemaResponse) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeleteSchemaResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *DeleteSchemaResponse) SetVersion(v int64) {
	o.Version = &v
}

func (o DeleteSchemaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteSchemaResponse struct {
	value *DeleteSchemaResponse
	isSet bool
}

func (v NullableDeleteSchemaResponse) Get() *DeleteSchemaResponse {
	return v.value
}

func (v *NullableDeleteSchemaResponse) Set(val *DeleteSchemaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSchemaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSchemaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSchemaResponse(val *DeleteSchemaResponse) *NullableDeleteSchemaResponse {
	return &NullableDeleteSchemaResponse{value: val, isSet: true}
}

func (v NullableDeleteSchemaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSchemaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


