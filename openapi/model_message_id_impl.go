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

// MessageIdImpl struct for MessageIdImpl
type MessageIdImpl struct {
	LedgerId *int64 `json:"ledgerId,omitempty"`
	EntryId *int64 `json:"entryId,omitempty"`
	PartitionIndex *int32 `json:"partitionIndex,omitempty"`
}

// NewMessageIdImpl instantiates a new MessageIdImpl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageIdImpl() *MessageIdImpl {
	this := MessageIdImpl{}
	return &this
}

// NewMessageIdImplWithDefaults instantiates a new MessageIdImpl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageIdImplWithDefaults() *MessageIdImpl {
	this := MessageIdImpl{}
	return &this
}

// GetLedgerId returns the LedgerId field value if set, zero value otherwise.
func (o *MessageIdImpl) GetLedgerId() int64 {
	if o == nil || o.LedgerId == nil {
		var ret int64
		return ret
	}
	return *o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageIdImpl) GetLedgerIdOk() (*int64, bool) {
	if o == nil || o.LedgerId == nil {
		return nil, false
	}
	return o.LedgerId, true
}

// HasLedgerId returns a boolean if a field has been set.
func (o *MessageIdImpl) HasLedgerId() bool {
	if o != nil && o.LedgerId != nil {
		return true
	}

	return false
}

// SetLedgerId gets a reference to the given int64 and assigns it to the LedgerId field.
func (o *MessageIdImpl) SetLedgerId(v int64) {
	o.LedgerId = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *MessageIdImpl) GetEntryId() int64 {
	if o == nil || o.EntryId == nil {
		var ret int64
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageIdImpl) GetEntryIdOk() (*int64, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *MessageIdImpl) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given int64 and assigns it to the EntryId field.
func (o *MessageIdImpl) SetEntryId(v int64) {
	o.EntryId = &v
}

// GetPartitionIndex returns the PartitionIndex field value if set, zero value otherwise.
func (o *MessageIdImpl) GetPartitionIndex() int32 {
	if o == nil || o.PartitionIndex == nil {
		var ret int32
		return ret
	}
	return *o.PartitionIndex
}

// GetPartitionIndexOk returns a tuple with the PartitionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageIdImpl) GetPartitionIndexOk() (*int32, bool) {
	if o == nil || o.PartitionIndex == nil {
		return nil, false
	}
	return o.PartitionIndex, true
}

// HasPartitionIndex returns a boolean if a field has been set.
func (o *MessageIdImpl) HasPartitionIndex() bool {
	if o != nil && o.PartitionIndex != nil {
		return true
	}

	return false
}

// SetPartitionIndex gets a reference to the given int32 and assigns it to the PartitionIndex field.
func (o *MessageIdImpl) SetPartitionIndex(v int32) {
	o.PartitionIndex = &v
}

func (o MessageIdImpl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LedgerId != nil {
		toSerialize["ledgerId"] = o.LedgerId
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.PartitionIndex != nil {
		toSerialize["partitionIndex"] = o.PartitionIndex
	}
	return json.Marshal(toSerialize)
}

type NullableMessageIdImpl struct {
	value *MessageIdImpl
	isSet bool
}

func (v NullableMessageIdImpl) Get() *MessageIdImpl {
	return v.value
}

func (v *NullableMessageIdImpl) Set(val *MessageIdImpl) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageIdImpl) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageIdImpl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageIdImpl(val *MessageIdImpl) *NullableMessageIdImpl {
	return &NullableMessageIdImpl{value: val, isSet: true}
}

func (v NullableMessageIdImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageIdImpl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


