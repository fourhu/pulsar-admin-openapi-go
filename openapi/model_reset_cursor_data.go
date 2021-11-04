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

// ResetCursorData struct for ResetCursorData
type ResetCursorData struct {
	BatchIndex *int32 `json:"batchIndex,omitempty"`
	EntryId *int64 `json:"entryId,omitempty"`
	Excluded *bool `json:"excluded,omitempty"`
	LedgerId *int64 `json:"ledgerId,omitempty"`
	PartitionIndex *int32 `json:"partitionIndex,omitempty"`
}

// NewResetCursorData instantiates a new ResetCursorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetCursorData() *ResetCursorData {
	this := ResetCursorData{}
	return &this
}

// NewResetCursorDataWithDefaults instantiates a new ResetCursorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetCursorDataWithDefaults() *ResetCursorData {
	this := ResetCursorData{}
	return &this
}

// GetBatchIndex returns the BatchIndex field value if set, zero value otherwise.
func (o *ResetCursorData) GetBatchIndex() int32 {
	if o == nil || o.BatchIndex == nil {
		var ret int32
		return ret
	}
	return *o.BatchIndex
}

// GetBatchIndexOk returns a tuple with the BatchIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetCursorData) GetBatchIndexOk() (*int32, bool) {
	if o == nil || o.BatchIndex == nil {
		return nil, false
	}
	return o.BatchIndex, true
}

// HasBatchIndex returns a boolean if a field has been set.
func (o *ResetCursorData) HasBatchIndex() bool {
	if o != nil && o.BatchIndex != nil {
		return true
	}

	return false
}

// SetBatchIndex gets a reference to the given int32 and assigns it to the BatchIndex field.
func (o *ResetCursorData) SetBatchIndex(v int32) {
	o.BatchIndex = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *ResetCursorData) GetEntryId() int64 {
	if o == nil || o.EntryId == nil {
		var ret int64
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetCursorData) GetEntryIdOk() (*int64, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *ResetCursorData) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given int64 and assigns it to the EntryId field.
func (o *ResetCursorData) SetEntryId(v int64) {
	o.EntryId = &v
}

// GetExcluded returns the Excluded field value if set, zero value otherwise.
func (o *ResetCursorData) GetExcluded() bool {
	if o == nil || o.Excluded == nil {
		var ret bool
		return ret
	}
	return *o.Excluded
}

// GetExcludedOk returns a tuple with the Excluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetCursorData) GetExcludedOk() (*bool, bool) {
	if o == nil || o.Excluded == nil {
		return nil, false
	}
	return o.Excluded, true
}

// HasExcluded returns a boolean if a field has been set.
func (o *ResetCursorData) HasExcluded() bool {
	if o != nil && o.Excluded != nil {
		return true
	}

	return false
}

// SetExcluded gets a reference to the given bool and assigns it to the Excluded field.
func (o *ResetCursorData) SetExcluded(v bool) {
	o.Excluded = &v
}

// GetLedgerId returns the LedgerId field value if set, zero value otherwise.
func (o *ResetCursorData) GetLedgerId() int64 {
	if o == nil || o.LedgerId == nil {
		var ret int64
		return ret
	}
	return *o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetCursorData) GetLedgerIdOk() (*int64, bool) {
	if o == nil || o.LedgerId == nil {
		return nil, false
	}
	return o.LedgerId, true
}

// HasLedgerId returns a boolean if a field has been set.
func (o *ResetCursorData) HasLedgerId() bool {
	if o != nil && o.LedgerId != nil {
		return true
	}

	return false
}

// SetLedgerId gets a reference to the given int64 and assigns it to the LedgerId field.
func (o *ResetCursorData) SetLedgerId(v int64) {
	o.LedgerId = &v
}

// GetPartitionIndex returns the PartitionIndex field value if set, zero value otherwise.
func (o *ResetCursorData) GetPartitionIndex() int32 {
	if o == nil || o.PartitionIndex == nil {
		var ret int32
		return ret
	}
	return *o.PartitionIndex
}

// GetPartitionIndexOk returns a tuple with the PartitionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetCursorData) GetPartitionIndexOk() (*int32, bool) {
	if o == nil || o.PartitionIndex == nil {
		return nil, false
	}
	return o.PartitionIndex, true
}

// HasPartitionIndex returns a boolean if a field has been set.
func (o *ResetCursorData) HasPartitionIndex() bool {
	if o != nil && o.PartitionIndex != nil {
		return true
	}

	return false
}

// SetPartitionIndex gets a reference to the given int32 and assigns it to the PartitionIndex field.
func (o *ResetCursorData) SetPartitionIndex(v int32) {
	o.PartitionIndex = &v
}

func (o ResetCursorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BatchIndex != nil {
		toSerialize["batchIndex"] = o.BatchIndex
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.Excluded != nil {
		toSerialize["excluded"] = o.Excluded
	}
	if o.LedgerId != nil {
		toSerialize["ledgerId"] = o.LedgerId
	}
	if o.PartitionIndex != nil {
		toSerialize["partitionIndex"] = o.PartitionIndex
	}
	return json.Marshal(toSerialize)
}

type NullableResetCursorData struct {
	value *ResetCursorData
	isSet bool
}

func (v NullableResetCursorData) Get() *ResetCursorData {
	return v.value
}

func (v *NullableResetCursorData) Set(val *ResetCursorData) {
	v.value = val
	v.isSet = true
}

func (v NullableResetCursorData) IsSet() bool {
	return v.isSet
}

func (v *NullableResetCursorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetCursorData(val *ResetCursorData) *NullableResetCursorData {
	return &NullableResetCursorData{value: val, isSet: true}
}

func (v NullableResetCursorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetCursorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


