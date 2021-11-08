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

// Functions struct for Functions
type Functions struct {
	ListOfConnectors *[]ConnectorDefinition `json:"listOfConnectors,omitempty"`
}

// NewFunctions instantiates a new Functions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctions() *Functions {
	this := Functions{}
	return &this
}

// NewFunctionsWithDefaults instantiates a new Functions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionsWithDefaults() *Functions {
	this := Functions{}
	return &this
}

// GetListOfConnectors returns the ListOfConnectors field value if set, zero value otherwise.
func (o *Functions) GetListOfConnectors() []ConnectorDefinition {
	if o == nil || o.ListOfConnectors == nil {
		var ret []ConnectorDefinition
		return ret
	}
	return *o.ListOfConnectors
}

// GetListOfConnectorsOk returns a tuple with the ListOfConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Functions) GetListOfConnectorsOk() (*[]ConnectorDefinition, bool) {
	if o == nil || o.ListOfConnectors == nil {
		return nil, false
	}
	return o.ListOfConnectors, true
}

// HasListOfConnectors returns a boolean if a field has been set.
func (o *Functions) HasListOfConnectors() bool {
	if o != nil && o.ListOfConnectors != nil {
		return true
	}

	return false
}

// SetListOfConnectors gets a reference to the given []ConnectorDefinition and assigns it to the ListOfConnectors field.
func (o *Functions) SetListOfConnectors(v []ConnectorDefinition) {
	o.ListOfConnectors = &v
}

func (o Functions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ListOfConnectors != nil {
		toSerialize["listOfConnectors"] = o.ListOfConnectors
	}
	return json.Marshal(toSerialize)
}

type NullableFunctions struct {
	value *Functions
	isSet bool
}

func (v NullableFunctions) Get() *Functions {
	return v.value
}

func (v *NullableFunctions) Set(val *Functions) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctions) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctions(val *Functions) *NullableFunctions {
	return &NullableFunctions{value: val, isSet: true}
}

func (v NullableFunctions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


