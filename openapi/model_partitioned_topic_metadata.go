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

// PartitionedTopicMetadata struct for PartitionedTopicMetadata
type PartitionedTopicMetadata struct {
	Partitions *int32 `json:"partitions,omitempty"`
}

// NewPartitionedTopicMetadata instantiates a new PartitionedTopicMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartitionedTopicMetadata() *PartitionedTopicMetadata {
	this := PartitionedTopicMetadata{}
	return &this
}

// NewPartitionedTopicMetadataWithDefaults instantiates a new PartitionedTopicMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionedTopicMetadataWithDefaults() *PartitionedTopicMetadata {
	this := PartitionedTopicMetadata{}
	return &this
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *PartitionedTopicMetadata) GetPartitions() int32 {
	if o == nil || o.Partitions == nil {
		var ret int32
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionedTopicMetadata) GetPartitionsOk() (*int32, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *PartitionedTopicMetadata) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given int32 and assigns it to the Partitions field.
func (o *PartitionedTopicMetadata) SetPartitions(v int32) {
	o.Partitions = &v
}

func (o PartitionedTopicMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
	return json.Marshal(toSerialize)
}

type NullablePartitionedTopicMetadata struct {
	value *PartitionedTopicMetadata
	isSet bool
}

func (v NullablePartitionedTopicMetadata) Get() *PartitionedTopicMetadata {
	return v.value
}

func (v *NullablePartitionedTopicMetadata) Set(val *PartitionedTopicMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitionedTopicMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitionedTopicMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitionedTopicMetadata(val *PartitionedTopicMetadata) *NullablePartitionedTopicMetadata {
	return &NullablePartitionedTopicMetadata{value: val, isSet: true}
}

func (v NullablePartitionedTopicMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitionedTopicMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


