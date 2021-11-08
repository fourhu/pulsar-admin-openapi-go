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

// FunctionsV2WorkerService struct for FunctionsV2WorkerService
type FunctionsV2WorkerService struct {
	ListOfConnectors *[]ConnectorDefinition `json:"listOfConnectors,omitempty"`
}

// NewFunctionsV2WorkerService instantiates a new FunctionsV2WorkerService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionsV2WorkerService() *FunctionsV2WorkerService {
	this := FunctionsV2WorkerService{}
	return &this
}

// NewFunctionsV2WorkerServiceWithDefaults instantiates a new FunctionsV2WorkerService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionsV2WorkerServiceWithDefaults() *FunctionsV2WorkerService {
	this := FunctionsV2WorkerService{}
	return &this
}

// GetListOfConnectors returns the ListOfConnectors field value if set, zero value otherwise.
func (o *FunctionsV2WorkerService) GetListOfConnectors() []ConnectorDefinition {
	if o == nil || o.ListOfConnectors == nil {
		var ret []ConnectorDefinition
		return ret
	}
	return *o.ListOfConnectors
}

// GetListOfConnectorsOk returns a tuple with the ListOfConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionsV2WorkerService) GetListOfConnectorsOk() (*[]ConnectorDefinition, bool) {
	if o == nil || o.ListOfConnectors == nil {
		return nil, false
	}
	return o.ListOfConnectors, true
}

// HasListOfConnectors returns a boolean if a field has been set.
func (o *FunctionsV2WorkerService) HasListOfConnectors() bool {
	if o != nil && o.ListOfConnectors != nil {
		return true
	}

	return false
}

// SetListOfConnectors gets a reference to the given []ConnectorDefinition and assigns it to the ListOfConnectors field.
func (o *FunctionsV2WorkerService) SetListOfConnectors(v []ConnectorDefinition) {
	o.ListOfConnectors = &v
}

func (o FunctionsV2WorkerService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ListOfConnectors != nil {
		toSerialize["listOfConnectors"] = o.ListOfConnectors
	}
	return json.Marshal(toSerialize)
}

type NullableFunctionsV2WorkerService struct {
	value *FunctionsV2WorkerService
	isSet bool
}

func (v NullableFunctionsV2WorkerService) Get() *FunctionsV2WorkerService {
	return v.value
}

func (v *NullableFunctionsV2WorkerService) Set(val *FunctionsV2WorkerService) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionsV2WorkerService) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionsV2WorkerService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionsV2WorkerService(val *FunctionsV2WorkerService) *NullableFunctionsV2WorkerService {
	return &NullableFunctionsV2WorkerService{value: val, isSet: true}
}

func (v NullableFunctionsV2WorkerService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionsV2WorkerService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


