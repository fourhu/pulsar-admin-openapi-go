/*
Pulsar Packages REST API

This provides the REST API for Pulsar Packages operations

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TimeZone struct for TimeZone
type TimeZone struct {
	DisplayName *string `json:"displayName,omitempty"`
	Id *string `json:"id,omitempty"`
	Dstsavings *int32 `json:"dstsavings,omitempty"`
	RawOffset *int32 `json:"rawOffset,omitempty"`
}

// NewTimeZone instantiates a new TimeZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeZone() *TimeZone {
	this := TimeZone{}
	return &this
}

// NewTimeZoneWithDefaults instantiates a new TimeZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeZoneWithDefaults() *TimeZone {
	this := TimeZone{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TimeZone) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeZone) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TimeZone) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TimeZone) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TimeZone) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeZone) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TimeZone) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TimeZone) SetId(v string) {
	o.Id = &v
}

// GetDstsavings returns the Dstsavings field value if set, zero value otherwise.
func (o *TimeZone) GetDstsavings() int32 {
	if o == nil || o.Dstsavings == nil {
		var ret int32
		return ret
	}
	return *o.Dstsavings
}

// GetDstsavingsOk returns a tuple with the Dstsavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeZone) GetDstsavingsOk() (*int32, bool) {
	if o == nil || o.Dstsavings == nil {
		return nil, false
	}
	return o.Dstsavings, true
}

// HasDstsavings returns a boolean if a field has been set.
func (o *TimeZone) HasDstsavings() bool {
	if o != nil && o.Dstsavings != nil {
		return true
	}

	return false
}

// SetDstsavings gets a reference to the given int32 and assigns it to the Dstsavings field.
func (o *TimeZone) SetDstsavings(v int32) {
	o.Dstsavings = &v
}

// GetRawOffset returns the RawOffset field value if set, zero value otherwise.
func (o *TimeZone) GetRawOffset() int32 {
	if o == nil || o.RawOffset == nil {
		var ret int32
		return ret
	}
	return *o.RawOffset
}

// GetRawOffsetOk returns a tuple with the RawOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeZone) GetRawOffsetOk() (*int32, bool) {
	if o == nil || o.RawOffset == nil {
		return nil, false
	}
	return o.RawOffset, true
}

// HasRawOffset returns a boolean if a field has been set.
func (o *TimeZone) HasRawOffset() bool {
	if o != nil && o.RawOffset != nil {
		return true
	}

	return false
}

// SetRawOffset gets a reference to the given int32 and assigns it to the RawOffset field.
func (o *TimeZone) SetRawOffset(v int32) {
	o.RawOffset = &v
}

func (o TimeZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Dstsavings != nil {
		toSerialize["dstsavings"] = o.Dstsavings
	}
	if o.RawOffset != nil {
		toSerialize["rawOffset"] = o.RawOffset
	}
	return json.Marshal(toSerialize)
}

type NullableTimeZone struct {
	value *TimeZone
	isSet bool
}

func (v NullableTimeZone) Get() *TimeZone {
	return v.value
}

func (v *NullableTimeZone) Set(val *TimeZone) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeZone) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeZone(val *TimeZone) *NullableTimeZone {
	return &NullableTimeZone{value: val, isSet: true}
}

func (v NullableTimeZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


