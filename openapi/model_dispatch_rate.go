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

// DispatchRate struct for DispatchRate
type DispatchRate struct {
	DispatchThrottlingRateInByte *int64 `json:"dispatchThrottlingRateInByte,omitempty"`
	DispatchThrottlingRateInMsg *int32 `json:"dispatchThrottlingRateInMsg,omitempty"`
	RatePeriodInSecond *int32 `json:"ratePeriodInSecond,omitempty"`
	RelativeToPublishRate *bool `json:"relativeToPublishRate,omitempty"`
}

// NewDispatchRate instantiates a new DispatchRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDispatchRate() *DispatchRate {
	this := DispatchRate{}
	return &this
}

// NewDispatchRateWithDefaults instantiates a new DispatchRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDispatchRateWithDefaults() *DispatchRate {
	this := DispatchRate{}
	return &this
}

// GetDispatchThrottlingRateInByte returns the DispatchThrottlingRateInByte field value if set, zero value otherwise.
func (o *DispatchRate) GetDispatchThrottlingRateInByte() int64 {
	if o == nil || o.DispatchThrottlingRateInByte == nil {
		var ret int64
		return ret
	}
	return *o.DispatchThrottlingRateInByte
}

// GetDispatchThrottlingRateInByteOk returns a tuple with the DispatchThrottlingRateInByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchRate) GetDispatchThrottlingRateInByteOk() (*int64, bool) {
	if o == nil || o.DispatchThrottlingRateInByte == nil {
		return nil, false
	}
	return o.DispatchThrottlingRateInByte, true
}

// HasDispatchThrottlingRateInByte returns a boolean if a field has been set.
func (o *DispatchRate) HasDispatchThrottlingRateInByte() bool {
	if o != nil && o.DispatchThrottlingRateInByte != nil {
		return true
	}

	return false
}

// SetDispatchThrottlingRateInByte gets a reference to the given int64 and assigns it to the DispatchThrottlingRateInByte field.
func (o *DispatchRate) SetDispatchThrottlingRateInByte(v int64) {
	o.DispatchThrottlingRateInByte = &v
}

// GetDispatchThrottlingRateInMsg returns the DispatchThrottlingRateInMsg field value if set, zero value otherwise.
func (o *DispatchRate) GetDispatchThrottlingRateInMsg() int32 {
	if o == nil || o.DispatchThrottlingRateInMsg == nil {
		var ret int32
		return ret
	}
	return *o.DispatchThrottlingRateInMsg
}

// GetDispatchThrottlingRateInMsgOk returns a tuple with the DispatchThrottlingRateInMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchRate) GetDispatchThrottlingRateInMsgOk() (*int32, bool) {
	if o == nil || o.DispatchThrottlingRateInMsg == nil {
		return nil, false
	}
	return o.DispatchThrottlingRateInMsg, true
}

// HasDispatchThrottlingRateInMsg returns a boolean if a field has been set.
func (o *DispatchRate) HasDispatchThrottlingRateInMsg() bool {
	if o != nil && o.DispatchThrottlingRateInMsg != nil {
		return true
	}

	return false
}

// SetDispatchThrottlingRateInMsg gets a reference to the given int32 and assigns it to the DispatchThrottlingRateInMsg field.
func (o *DispatchRate) SetDispatchThrottlingRateInMsg(v int32) {
	o.DispatchThrottlingRateInMsg = &v
}

// GetRatePeriodInSecond returns the RatePeriodInSecond field value if set, zero value otherwise.
func (o *DispatchRate) GetRatePeriodInSecond() int32 {
	if o == nil || o.RatePeriodInSecond == nil {
		var ret int32
		return ret
	}
	return *o.RatePeriodInSecond
}

// GetRatePeriodInSecondOk returns a tuple with the RatePeriodInSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchRate) GetRatePeriodInSecondOk() (*int32, bool) {
	if o == nil || o.RatePeriodInSecond == nil {
		return nil, false
	}
	return o.RatePeriodInSecond, true
}

// HasRatePeriodInSecond returns a boolean if a field has been set.
func (o *DispatchRate) HasRatePeriodInSecond() bool {
	if o != nil && o.RatePeriodInSecond != nil {
		return true
	}

	return false
}

// SetRatePeriodInSecond gets a reference to the given int32 and assigns it to the RatePeriodInSecond field.
func (o *DispatchRate) SetRatePeriodInSecond(v int32) {
	o.RatePeriodInSecond = &v
}

// GetRelativeToPublishRate returns the RelativeToPublishRate field value if set, zero value otherwise.
func (o *DispatchRate) GetRelativeToPublishRate() bool {
	if o == nil || o.RelativeToPublishRate == nil {
		var ret bool
		return ret
	}
	return *o.RelativeToPublishRate
}

// GetRelativeToPublishRateOk returns a tuple with the RelativeToPublishRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchRate) GetRelativeToPublishRateOk() (*bool, bool) {
	if o == nil || o.RelativeToPublishRate == nil {
		return nil, false
	}
	return o.RelativeToPublishRate, true
}

// HasRelativeToPublishRate returns a boolean if a field has been set.
func (o *DispatchRate) HasRelativeToPublishRate() bool {
	if o != nil && o.RelativeToPublishRate != nil {
		return true
	}

	return false
}

// SetRelativeToPublishRate gets a reference to the given bool and assigns it to the RelativeToPublishRate field.
func (o *DispatchRate) SetRelativeToPublishRate(v bool) {
	o.RelativeToPublishRate = &v
}

func (o DispatchRate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DispatchThrottlingRateInByte != nil {
		toSerialize["dispatchThrottlingRateInByte"] = o.DispatchThrottlingRateInByte
	}
	if o.DispatchThrottlingRateInMsg != nil {
		toSerialize["dispatchThrottlingRateInMsg"] = o.DispatchThrottlingRateInMsg
	}
	if o.RatePeriodInSecond != nil {
		toSerialize["ratePeriodInSecond"] = o.RatePeriodInSecond
	}
	if o.RelativeToPublishRate != nil {
		toSerialize["relativeToPublishRate"] = o.RelativeToPublishRate
	}
	return json.Marshal(toSerialize)
}

type NullableDispatchRate struct {
	value *DispatchRate
	isSet bool
}

func (v NullableDispatchRate) Get() *DispatchRate {
	return v.value
}

func (v *NullableDispatchRate) Set(val *DispatchRate) {
	v.value = val
	v.isSet = true
}

func (v NullableDispatchRate) IsSet() bool {
	return v.isSet
}

func (v *NullableDispatchRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispatchRate(val *DispatchRate) *NullableDispatchRate {
	return &NullableDispatchRate{value: val, isSet: true}
}

func (v NullableDispatchRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispatchRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


