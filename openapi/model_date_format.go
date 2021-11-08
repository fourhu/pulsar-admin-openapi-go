/*
Pulsar Admin REST API

This provides the REST API for admin operations

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// DateFormat struct for DateFormat
type DateFormat struct {
	Calendar *time.Time `json:"calendar,omitempty"`
	NumberFormat *NumberFormat `json:"numberFormat,omitempty"`
	TimeZone *TimeZone `json:"timeZone,omitempty"`
	Lenient *bool `json:"lenient,omitempty"`
}

// NewDateFormat instantiates a new DateFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateFormat() *DateFormat {
	this := DateFormat{}
	return &this
}

// NewDateFormatWithDefaults instantiates a new DateFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateFormatWithDefaults() *DateFormat {
	this := DateFormat{}
	return &this
}

// GetCalendar returns the Calendar field value if set, zero value otherwise.
func (o *DateFormat) GetCalendar() time.Time {
	if o == nil || o.Calendar == nil {
		var ret time.Time
		return ret
	}
	return *o.Calendar
}

// GetCalendarOk returns a tuple with the Calendar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetCalendarOk() (*time.Time, bool) {
	if o == nil || o.Calendar == nil {
		return nil, false
	}
	return o.Calendar, true
}

// HasCalendar returns a boolean if a field has been set.
func (o *DateFormat) HasCalendar() bool {
	if o != nil && o.Calendar != nil {
		return true
	}

	return false
}

// SetCalendar gets a reference to the given time.Time and assigns it to the Calendar field.
func (o *DateFormat) SetCalendar(v time.Time) {
	o.Calendar = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *DateFormat) GetNumberFormat() NumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret NumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetNumberFormatOk() (*NumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *DateFormat) HasNumberFormat() bool {
	if o != nil && o.NumberFormat != nil {
		return true
	}

	return false
}

// SetNumberFormat gets a reference to the given NumberFormat and assigns it to the NumberFormat field.
func (o *DateFormat) SetNumberFormat(v NumberFormat) {
	o.NumberFormat = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *DateFormat) GetTimeZone() TimeZone {
	if o == nil || o.TimeZone == nil {
		var ret TimeZone
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetTimeZoneOk() (*TimeZone, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *DateFormat) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given TimeZone and assigns it to the TimeZone field.
func (o *DateFormat) SetTimeZone(v TimeZone) {
	o.TimeZone = &v
}

// GetLenient returns the Lenient field value if set, zero value otherwise.
func (o *DateFormat) GetLenient() bool {
	if o == nil || o.Lenient == nil {
		var ret bool
		return ret
	}
	return *o.Lenient
}

// GetLenientOk returns a tuple with the Lenient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetLenientOk() (*bool, bool) {
	if o == nil || o.Lenient == nil {
		return nil, false
	}
	return o.Lenient, true
}

// HasLenient returns a boolean if a field has been set.
func (o *DateFormat) HasLenient() bool {
	if o != nil && o.Lenient != nil {
		return true
	}

	return false
}

// SetLenient gets a reference to the given bool and assigns it to the Lenient field.
func (o *DateFormat) SetLenient(v bool) {
	o.Lenient = &v
}

func (o DateFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Calendar != nil {
		toSerialize["calendar"] = o.Calendar
	}
	if o.NumberFormat != nil {
		toSerialize["numberFormat"] = o.NumberFormat
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.Lenient != nil {
		toSerialize["lenient"] = o.Lenient
	}
	return json.Marshal(toSerialize)
}

type NullableDateFormat struct {
	value *DateFormat
	isSet bool
}

func (v NullableDateFormat) Get() *DateFormat {
	return v.value
}

func (v *NullableDateFormat) Set(val *DateFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableDateFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableDateFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateFormat(val *DateFormat) *NullableDateFormat {
	return &NullableDateFormat{value: val, isSet: true}
}

func (v NullableDateFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


