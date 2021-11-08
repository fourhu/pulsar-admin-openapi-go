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

// NumberFormat struct for NumberFormat
type NumberFormat struct {
	GroupingUsed *bool `json:"groupingUsed,omitempty"`
	ParseIntegerOnly *bool `json:"parseIntegerOnly,omitempty"`
	MaximumIntegerDigits *int32 `json:"maximumIntegerDigits,omitempty"`
	MinimumIntegerDigits *int32 `json:"minimumIntegerDigits,omitempty"`
	MaximumFractionDigits *int32 `json:"maximumFractionDigits,omitempty"`
	MinimumFractionDigits *int32 `json:"minimumFractionDigits,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
	RoundingMode *string `json:"roundingMode,omitempty"`
}

// NewNumberFormat instantiates a new NumberFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberFormat() *NumberFormat {
	this := NumberFormat{}
	return &this
}

// NewNumberFormatWithDefaults instantiates a new NumberFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberFormatWithDefaults() *NumberFormat {
	this := NumberFormat{}
	return &this
}

// GetGroupingUsed returns the GroupingUsed field value if set, zero value otherwise.
func (o *NumberFormat) GetGroupingUsed() bool {
	if o == nil || o.GroupingUsed == nil {
		var ret bool
		return ret
	}
	return *o.GroupingUsed
}

// GetGroupingUsedOk returns a tuple with the GroupingUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetGroupingUsedOk() (*bool, bool) {
	if o == nil || o.GroupingUsed == nil {
		return nil, false
	}
	return o.GroupingUsed, true
}

// HasGroupingUsed returns a boolean if a field has been set.
func (o *NumberFormat) HasGroupingUsed() bool {
	if o != nil && o.GroupingUsed != nil {
		return true
	}

	return false
}

// SetGroupingUsed gets a reference to the given bool and assigns it to the GroupingUsed field.
func (o *NumberFormat) SetGroupingUsed(v bool) {
	o.GroupingUsed = &v
}

// GetParseIntegerOnly returns the ParseIntegerOnly field value if set, zero value otherwise.
func (o *NumberFormat) GetParseIntegerOnly() bool {
	if o == nil || o.ParseIntegerOnly == nil {
		var ret bool
		return ret
	}
	return *o.ParseIntegerOnly
}

// GetParseIntegerOnlyOk returns a tuple with the ParseIntegerOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetParseIntegerOnlyOk() (*bool, bool) {
	if o == nil || o.ParseIntegerOnly == nil {
		return nil, false
	}
	return o.ParseIntegerOnly, true
}

// HasParseIntegerOnly returns a boolean if a field has been set.
func (o *NumberFormat) HasParseIntegerOnly() bool {
	if o != nil && o.ParseIntegerOnly != nil {
		return true
	}

	return false
}

// SetParseIntegerOnly gets a reference to the given bool and assigns it to the ParseIntegerOnly field.
func (o *NumberFormat) SetParseIntegerOnly(v bool) {
	o.ParseIntegerOnly = &v
}

// GetMaximumIntegerDigits returns the MaximumIntegerDigits field value if set, zero value otherwise.
func (o *NumberFormat) GetMaximumIntegerDigits() int32 {
	if o == nil || o.MaximumIntegerDigits == nil {
		var ret int32
		return ret
	}
	return *o.MaximumIntegerDigits
}

// GetMaximumIntegerDigitsOk returns a tuple with the MaximumIntegerDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetMaximumIntegerDigitsOk() (*int32, bool) {
	if o == nil || o.MaximumIntegerDigits == nil {
		return nil, false
	}
	return o.MaximumIntegerDigits, true
}

// HasMaximumIntegerDigits returns a boolean if a field has been set.
func (o *NumberFormat) HasMaximumIntegerDigits() bool {
	if o != nil && o.MaximumIntegerDigits != nil {
		return true
	}

	return false
}

// SetMaximumIntegerDigits gets a reference to the given int32 and assigns it to the MaximumIntegerDigits field.
func (o *NumberFormat) SetMaximumIntegerDigits(v int32) {
	o.MaximumIntegerDigits = &v
}

// GetMinimumIntegerDigits returns the MinimumIntegerDigits field value if set, zero value otherwise.
func (o *NumberFormat) GetMinimumIntegerDigits() int32 {
	if o == nil || o.MinimumIntegerDigits == nil {
		var ret int32
		return ret
	}
	return *o.MinimumIntegerDigits
}

// GetMinimumIntegerDigitsOk returns a tuple with the MinimumIntegerDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetMinimumIntegerDigitsOk() (*int32, bool) {
	if o == nil || o.MinimumIntegerDigits == nil {
		return nil, false
	}
	return o.MinimumIntegerDigits, true
}

// HasMinimumIntegerDigits returns a boolean if a field has been set.
func (o *NumberFormat) HasMinimumIntegerDigits() bool {
	if o != nil && o.MinimumIntegerDigits != nil {
		return true
	}

	return false
}

// SetMinimumIntegerDigits gets a reference to the given int32 and assigns it to the MinimumIntegerDigits field.
func (o *NumberFormat) SetMinimumIntegerDigits(v int32) {
	o.MinimumIntegerDigits = &v
}

// GetMaximumFractionDigits returns the MaximumFractionDigits field value if set, zero value otherwise.
func (o *NumberFormat) GetMaximumFractionDigits() int32 {
	if o == nil || o.MaximumFractionDigits == nil {
		var ret int32
		return ret
	}
	return *o.MaximumFractionDigits
}

// GetMaximumFractionDigitsOk returns a tuple with the MaximumFractionDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetMaximumFractionDigitsOk() (*int32, bool) {
	if o == nil || o.MaximumFractionDigits == nil {
		return nil, false
	}
	return o.MaximumFractionDigits, true
}

// HasMaximumFractionDigits returns a boolean if a field has been set.
func (o *NumberFormat) HasMaximumFractionDigits() bool {
	if o != nil && o.MaximumFractionDigits != nil {
		return true
	}

	return false
}

// SetMaximumFractionDigits gets a reference to the given int32 and assigns it to the MaximumFractionDigits field.
func (o *NumberFormat) SetMaximumFractionDigits(v int32) {
	o.MaximumFractionDigits = &v
}

// GetMinimumFractionDigits returns the MinimumFractionDigits field value if set, zero value otherwise.
func (o *NumberFormat) GetMinimumFractionDigits() int32 {
	if o == nil || o.MinimumFractionDigits == nil {
		var ret int32
		return ret
	}
	return *o.MinimumFractionDigits
}

// GetMinimumFractionDigitsOk returns a tuple with the MinimumFractionDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetMinimumFractionDigitsOk() (*int32, bool) {
	if o == nil || o.MinimumFractionDigits == nil {
		return nil, false
	}
	return o.MinimumFractionDigits, true
}

// HasMinimumFractionDigits returns a boolean if a field has been set.
func (o *NumberFormat) HasMinimumFractionDigits() bool {
	if o != nil && o.MinimumFractionDigits != nil {
		return true
	}

	return false
}

// SetMinimumFractionDigits gets a reference to the given int32 and assigns it to the MinimumFractionDigits field.
func (o *NumberFormat) SetMinimumFractionDigits(v int32) {
	o.MinimumFractionDigits = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *NumberFormat) GetCurrency() Currency {
	if o == nil || o.Currency == nil {
		var ret Currency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetCurrencyOk() (*Currency, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *NumberFormat) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given Currency and assigns it to the Currency field.
func (o *NumberFormat) SetCurrency(v Currency) {
	o.Currency = &v
}

// GetRoundingMode returns the RoundingMode field value if set, zero value otherwise.
func (o *NumberFormat) GetRoundingMode() string {
	if o == nil || o.RoundingMode == nil {
		var ret string
		return ret
	}
	return *o.RoundingMode
}

// GetRoundingModeOk returns a tuple with the RoundingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetRoundingModeOk() (*string, bool) {
	if o == nil || o.RoundingMode == nil {
		return nil, false
	}
	return o.RoundingMode, true
}

// HasRoundingMode returns a boolean if a field has been set.
func (o *NumberFormat) HasRoundingMode() bool {
	if o != nil && o.RoundingMode != nil {
		return true
	}

	return false
}

// SetRoundingMode gets a reference to the given string and assigns it to the RoundingMode field.
func (o *NumberFormat) SetRoundingMode(v string) {
	o.RoundingMode = &v
}

func (o NumberFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupingUsed != nil {
		toSerialize["groupingUsed"] = o.GroupingUsed
	}
	if o.ParseIntegerOnly != nil {
		toSerialize["parseIntegerOnly"] = o.ParseIntegerOnly
	}
	if o.MaximumIntegerDigits != nil {
		toSerialize["maximumIntegerDigits"] = o.MaximumIntegerDigits
	}
	if o.MinimumIntegerDigits != nil {
		toSerialize["minimumIntegerDigits"] = o.MinimumIntegerDigits
	}
	if o.MaximumFractionDigits != nil {
		toSerialize["maximumFractionDigits"] = o.MaximumFractionDigits
	}
	if o.MinimumFractionDigits != nil {
		toSerialize["minimumFractionDigits"] = o.MinimumFractionDigits
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.RoundingMode != nil {
		toSerialize["roundingMode"] = o.RoundingMode
	}
	return json.Marshal(toSerialize)
}

type NullableNumberFormat struct {
	value *NumberFormat
	isSet bool
}

func (v NullableNumberFormat) Get() *NumberFormat {
	return v.value
}

func (v *NullableNumberFormat) Set(val *NumberFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberFormat(val *NumberFormat) *NullableNumberFormat {
	return &NullableNumberFormat{value: val, isSet: true}
}

func (v NullableNumberFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


