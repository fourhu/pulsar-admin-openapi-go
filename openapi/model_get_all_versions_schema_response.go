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

// GetAllVersionsSchemaResponse struct for GetAllVersionsSchemaResponse
type GetAllVersionsSchemaResponse struct {
	GetSchemaResponses *[]GetSchemaResponse `json:"getSchemaResponses,omitempty"`
}

// NewGetAllVersionsSchemaResponse instantiates a new GetAllVersionsSchemaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllVersionsSchemaResponse() *GetAllVersionsSchemaResponse {
	this := GetAllVersionsSchemaResponse{}
	return &this
}

// NewGetAllVersionsSchemaResponseWithDefaults instantiates a new GetAllVersionsSchemaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllVersionsSchemaResponseWithDefaults() *GetAllVersionsSchemaResponse {
	this := GetAllVersionsSchemaResponse{}
	return &this
}

// GetGetSchemaResponses returns the GetSchemaResponses field value if set, zero value otherwise.
func (o *GetAllVersionsSchemaResponse) GetGetSchemaResponses() []GetSchemaResponse {
	if o == nil || o.GetSchemaResponses == nil {
		var ret []GetSchemaResponse
		return ret
	}
	return *o.GetSchemaResponses
}

// GetGetSchemaResponsesOk returns a tuple with the GetSchemaResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllVersionsSchemaResponse) GetGetSchemaResponsesOk() (*[]GetSchemaResponse, bool) {
	if o == nil || o.GetSchemaResponses == nil {
		return nil, false
	}
	return o.GetSchemaResponses, true
}

// HasGetSchemaResponses returns a boolean if a field has been set.
func (o *GetAllVersionsSchemaResponse) HasGetSchemaResponses() bool {
	if o != nil && o.GetSchemaResponses != nil {
		return true
	}

	return false
}

// SetGetSchemaResponses gets a reference to the given []GetSchemaResponse and assigns it to the GetSchemaResponses field.
func (o *GetAllVersionsSchemaResponse) SetGetSchemaResponses(v []GetSchemaResponse) {
	o.GetSchemaResponses = &v
}

func (o GetAllVersionsSchemaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GetSchemaResponses != nil {
		toSerialize["getSchemaResponses"] = o.GetSchemaResponses
	}
	return json.Marshal(toSerialize)
}

type NullableGetAllVersionsSchemaResponse struct {
	value *GetAllVersionsSchemaResponse
	isSet bool
}

func (v NullableGetAllVersionsSchemaResponse) Get() *GetAllVersionsSchemaResponse {
	return v.value
}

func (v *NullableGetAllVersionsSchemaResponse) Set(val *GetAllVersionsSchemaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllVersionsSchemaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllVersionsSchemaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllVersionsSchemaResponse(val *GetAllVersionsSchemaResponse) *NullableGetAllVersionsSchemaResponse {
	return &NullableGetAllVersionsSchemaResponse{value: val, isSet: true}
}

func (v NullableGetAllVersionsSchemaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllVersionsSchemaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

