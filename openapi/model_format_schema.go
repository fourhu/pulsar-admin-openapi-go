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

// FormatSchema struct for FormatSchema
type FormatSchema struct {
	SchemaType *string `json:"schemaType,omitempty"`
}

// NewFormatSchema instantiates a new FormatSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormatSchema() *FormatSchema {
	this := FormatSchema{}
	return &this
}

// NewFormatSchemaWithDefaults instantiates a new FormatSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatSchemaWithDefaults() *FormatSchema {
	this := FormatSchema{}
	return &this
}

// GetSchemaType returns the SchemaType field value if set, zero value otherwise.
func (o *FormatSchema) GetSchemaType() string {
	if o == nil || o.SchemaType == nil {
		var ret string
		return ret
	}
	return *o.SchemaType
}

// GetSchemaTypeOk returns a tuple with the SchemaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatSchema) GetSchemaTypeOk() (*string, bool) {
	if o == nil || o.SchemaType == nil {
		return nil, false
	}
	return o.SchemaType, true
}

// HasSchemaType returns a boolean if a field has been set.
func (o *FormatSchema) HasSchemaType() bool {
	if o != nil && o.SchemaType != nil {
		return true
	}

	return false
}

// SetSchemaType gets a reference to the given string and assigns it to the SchemaType field.
func (o *FormatSchema) SetSchemaType(v string) {
	o.SchemaType = &v
}

func (o FormatSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SchemaType != nil {
		toSerialize["schemaType"] = o.SchemaType
	}
	return json.Marshal(toSerialize)
}

type NullableFormatSchema struct {
	value *FormatSchema
	isSet bool
}

func (v NullableFormatSchema) Get() *FormatSchema {
	return v.value
}

func (v *NullableFormatSchema) Set(val *FormatSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFormatSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFormatSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormatSchema(val *FormatSchema) *NullableFormatSchema {
	return &NullableFormatSchema{value: val, isSet: true}
}

func (v NullableFormatSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormatSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


