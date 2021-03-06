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

// Sources struct for Sources
type Sources struct {
	SourceList *[]ConnectorDefinition `json:"sourceList,omitempty"`
	ListOfConnectors *[]ConnectorDefinition `json:"listOfConnectors,omitempty"`
}

// NewSources instantiates a new Sources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSources() *Sources {
	this := Sources{}
	return &this
}

// NewSourcesWithDefaults instantiates a new Sources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcesWithDefaults() *Sources {
	this := Sources{}
	return &this
}

// GetSourceList returns the SourceList field value if set, zero value otherwise.
func (o *Sources) GetSourceList() []ConnectorDefinition {
	if o == nil || o.SourceList == nil {
		var ret []ConnectorDefinition
		return ret
	}
	return *o.SourceList
}

// GetSourceListOk returns a tuple with the SourceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sources) GetSourceListOk() (*[]ConnectorDefinition, bool) {
	if o == nil || o.SourceList == nil {
		return nil, false
	}
	return o.SourceList, true
}

// HasSourceList returns a boolean if a field has been set.
func (o *Sources) HasSourceList() bool {
	if o != nil && o.SourceList != nil {
		return true
	}

	return false
}

// SetSourceList gets a reference to the given []ConnectorDefinition and assigns it to the SourceList field.
func (o *Sources) SetSourceList(v []ConnectorDefinition) {
	o.SourceList = &v
}

// GetListOfConnectors returns the ListOfConnectors field value if set, zero value otherwise.
func (o *Sources) GetListOfConnectors() []ConnectorDefinition {
	if o == nil || o.ListOfConnectors == nil {
		var ret []ConnectorDefinition
		return ret
	}
	return *o.ListOfConnectors
}

// GetListOfConnectorsOk returns a tuple with the ListOfConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sources) GetListOfConnectorsOk() (*[]ConnectorDefinition, bool) {
	if o == nil || o.ListOfConnectors == nil {
		return nil, false
	}
	return o.ListOfConnectors, true
}

// HasListOfConnectors returns a boolean if a field has been set.
func (o *Sources) HasListOfConnectors() bool {
	if o != nil && o.ListOfConnectors != nil {
		return true
	}

	return false
}

// SetListOfConnectors gets a reference to the given []ConnectorDefinition and assigns it to the ListOfConnectors field.
func (o *Sources) SetListOfConnectors(v []ConnectorDefinition) {
	o.ListOfConnectors = &v
}

func (o Sources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceList != nil {
		toSerialize["sourceList"] = o.SourceList
	}
	if o.ListOfConnectors != nil {
		toSerialize["listOfConnectors"] = o.ListOfConnectors
	}
	return json.Marshal(toSerialize)
}

type NullableSources struct {
	value *Sources
	isSet bool
}

func (v NullableSources) Get() *Sources {
	return v.value
}

func (v *NullableSources) Set(val *Sources) {
	v.value = val
	v.isSet = true
}

func (v NullableSources) IsSet() bool {
	return v.isSet
}

func (v *NullableSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSources(val *Sources) *NullableSources {
	return &NullableSources{value: val, isSet: true}
}

func (v NullableSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


