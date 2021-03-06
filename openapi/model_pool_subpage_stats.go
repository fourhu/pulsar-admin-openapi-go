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

// PoolSubpageStats struct for PoolSubpageStats
type PoolSubpageStats struct {
	MaxNumElements *int32 `json:"maxNumElements,omitempty"`
	NumAvailable *int32 `json:"numAvailable,omitempty"`
	ElementSize *int32 `json:"elementSize,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
}

// NewPoolSubpageStats instantiates a new PoolSubpageStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolSubpageStats() *PoolSubpageStats {
	this := PoolSubpageStats{}
	return &this
}

// NewPoolSubpageStatsWithDefaults instantiates a new PoolSubpageStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolSubpageStatsWithDefaults() *PoolSubpageStats {
	this := PoolSubpageStats{}
	return &this
}

// GetMaxNumElements returns the MaxNumElements field value if set, zero value otherwise.
func (o *PoolSubpageStats) GetMaxNumElements() int32 {
	if o == nil || o.MaxNumElements == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumElements
}

// GetMaxNumElementsOk returns a tuple with the MaxNumElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolSubpageStats) GetMaxNumElementsOk() (*int32, bool) {
	if o == nil || o.MaxNumElements == nil {
		return nil, false
	}
	return o.MaxNumElements, true
}

// HasMaxNumElements returns a boolean if a field has been set.
func (o *PoolSubpageStats) HasMaxNumElements() bool {
	if o != nil && o.MaxNumElements != nil {
		return true
	}

	return false
}

// SetMaxNumElements gets a reference to the given int32 and assigns it to the MaxNumElements field.
func (o *PoolSubpageStats) SetMaxNumElements(v int32) {
	o.MaxNumElements = &v
}

// GetNumAvailable returns the NumAvailable field value if set, zero value otherwise.
func (o *PoolSubpageStats) GetNumAvailable() int32 {
	if o == nil || o.NumAvailable == nil {
		var ret int32
		return ret
	}
	return *o.NumAvailable
}

// GetNumAvailableOk returns a tuple with the NumAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolSubpageStats) GetNumAvailableOk() (*int32, bool) {
	if o == nil || o.NumAvailable == nil {
		return nil, false
	}
	return o.NumAvailable, true
}

// HasNumAvailable returns a boolean if a field has been set.
func (o *PoolSubpageStats) HasNumAvailable() bool {
	if o != nil && o.NumAvailable != nil {
		return true
	}

	return false
}

// SetNumAvailable gets a reference to the given int32 and assigns it to the NumAvailable field.
func (o *PoolSubpageStats) SetNumAvailable(v int32) {
	o.NumAvailable = &v
}

// GetElementSize returns the ElementSize field value if set, zero value otherwise.
func (o *PoolSubpageStats) GetElementSize() int32 {
	if o == nil || o.ElementSize == nil {
		var ret int32
		return ret
	}
	return *o.ElementSize
}

// GetElementSizeOk returns a tuple with the ElementSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolSubpageStats) GetElementSizeOk() (*int32, bool) {
	if o == nil || o.ElementSize == nil {
		return nil, false
	}
	return o.ElementSize, true
}

// HasElementSize returns a boolean if a field has been set.
func (o *PoolSubpageStats) HasElementSize() bool {
	if o != nil && o.ElementSize != nil {
		return true
	}

	return false
}

// SetElementSize gets a reference to the given int32 and assigns it to the ElementSize field.
func (o *PoolSubpageStats) SetElementSize(v int32) {
	o.ElementSize = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PoolSubpageStats) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolSubpageStats) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PoolSubpageStats) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *PoolSubpageStats) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o PoolSubpageStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxNumElements != nil {
		toSerialize["maxNumElements"] = o.MaxNumElements
	}
	if o.NumAvailable != nil {
		toSerialize["numAvailable"] = o.NumAvailable
	}
	if o.ElementSize != nil {
		toSerialize["elementSize"] = o.ElementSize
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullablePoolSubpageStats struct {
	value *PoolSubpageStats
	isSet bool
}

func (v NullablePoolSubpageStats) Get() *PoolSubpageStats {
	return v.value
}

func (v *NullablePoolSubpageStats) Set(val *PoolSubpageStats) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolSubpageStats) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolSubpageStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolSubpageStats(val *PoolSubpageStats) *NullablePoolSubpageStats {
	return &NullablePoolSubpageStats{value: val, isSet: true}
}

func (v NullablePoolSubpageStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolSubpageStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


