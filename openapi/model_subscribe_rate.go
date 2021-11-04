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

// SubscribeRate struct for SubscribeRate
type SubscribeRate struct {
	RatePeriodInSecond *int32 `json:"ratePeriodInSecond,omitempty"`
	SubscribeThrottlingRatePerConsumer *int32 `json:"subscribeThrottlingRatePerConsumer,omitempty"`
}

// NewSubscribeRate instantiates a new SubscribeRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeRate() *SubscribeRate {
	this := SubscribeRate{}
	return &this
}

// NewSubscribeRateWithDefaults instantiates a new SubscribeRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeRateWithDefaults() *SubscribeRate {
	this := SubscribeRate{}
	return &this
}

// GetRatePeriodInSecond returns the RatePeriodInSecond field value if set, zero value otherwise.
func (o *SubscribeRate) GetRatePeriodInSecond() int32 {
	if o == nil || o.RatePeriodInSecond == nil {
		var ret int32
		return ret
	}
	return *o.RatePeriodInSecond
}

// GetRatePeriodInSecondOk returns a tuple with the RatePeriodInSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeRate) GetRatePeriodInSecondOk() (*int32, bool) {
	if o == nil || o.RatePeriodInSecond == nil {
		return nil, false
	}
	return o.RatePeriodInSecond, true
}

// HasRatePeriodInSecond returns a boolean if a field has been set.
func (o *SubscribeRate) HasRatePeriodInSecond() bool {
	if o != nil && o.RatePeriodInSecond != nil {
		return true
	}

	return false
}

// SetRatePeriodInSecond gets a reference to the given int32 and assigns it to the RatePeriodInSecond field.
func (o *SubscribeRate) SetRatePeriodInSecond(v int32) {
	o.RatePeriodInSecond = &v
}

// GetSubscribeThrottlingRatePerConsumer returns the SubscribeThrottlingRatePerConsumer field value if set, zero value otherwise.
func (o *SubscribeRate) GetSubscribeThrottlingRatePerConsumer() int32 {
	if o == nil || o.SubscribeThrottlingRatePerConsumer == nil {
		var ret int32
		return ret
	}
	return *o.SubscribeThrottlingRatePerConsumer
}

// GetSubscribeThrottlingRatePerConsumerOk returns a tuple with the SubscribeThrottlingRatePerConsumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeRate) GetSubscribeThrottlingRatePerConsumerOk() (*int32, bool) {
	if o == nil || o.SubscribeThrottlingRatePerConsumer == nil {
		return nil, false
	}
	return o.SubscribeThrottlingRatePerConsumer, true
}

// HasSubscribeThrottlingRatePerConsumer returns a boolean if a field has been set.
func (o *SubscribeRate) HasSubscribeThrottlingRatePerConsumer() bool {
	if o != nil && o.SubscribeThrottlingRatePerConsumer != nil {
		return true
	}

	return false
}

// SetSubscribeThrottlingRatePerConsumer gets a reference to the given int32 and assigns it to the SubscribeThrottlingRatePerConsumer field.
func (o *SubscribeRate) SetSubscribeThrottlingRatePerConsumer(v int32) {
	o.SubscribeThrottlingRatePerConsumer = &v
}

func (o SubscribeRate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RatePeriodInSecond != nil {
		toSerialize["ratePeriodInSecond"] = o.RatePeriodInSecond
	}
	if o.SubscribeThrottlingRatePerConsumer != nil {
		toSerialize["subscribeThrottlingRatePerConsumer"] = o.SubscribeThrottlingRatePerConsumer
	}
	return json.Marshal(toSerialize)
}

type NullableSubscribeRate struct {
	value *SubscribeRate
	isSet bool
}

func (v NullableSubscribeRate) Get() *SubscribeRate {
	return v.value
}

func (v *NullableSubscribeRate) Set(val *SubscribeRate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeRate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeRate(val *SubscribeRate) *NullableSubscribeRate {
	return &NullableSubscribeRate{value: val, isSet: true}
}

func (v NullableSubscribeRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

