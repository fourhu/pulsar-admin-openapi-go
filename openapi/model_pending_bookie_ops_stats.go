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

// PendingBookieOpsStats struct for PendingBookieOpsStats
type PendingBookieOpsStats struct {
	DataLedgerOpenOp *int64 `json:"dataLedgerOpenOp,omitempty"`
	DataLedgerCloseOp *int64 `json:"dataLedgerCloseOp,omitempty"`
	DataLedgerCreateOp *int64 `json:"dataLedgerCreateOp,omitempty"`
	DataLedgerDeleteOp *int64 `json:"dataLedgerDeleteOp,omitempty"`
	CursorLedgerOpenOp *int64 `json:"cursorLedgerOpenOp,omitempty"`
	CursorLedgerCloseOp *int64 `json:"cursorLedgerCloseOp,omitempty"`
	CursorLedgerCreateOp *int64 `json:"cursorLedgerCreateOp,omitempty"`
	CursorLedgerDeleteOp *int64 `json:"cursorLedgerDeleteOp,omitempty"`
}

// NewPendingBookieOpsStats instantiates a new PendingBookieOpsStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingBookieOpsStats() *PendingBookieOpsStats {
	this := PendingBookieOpsStats{}
	return &this
}

// NewPendingBookieOpsStatsWithDefaults instantiates a new PendingBookieOpsStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingBookieOpsStatsWithDefaults() *PendingBookieOpsStats {
	this := PendingBookieOpsStats{}
	return &this
}

// GetDataLedgerOpenOp returns the DataLedgerOpenOp field value if set, zero value otherwise.
func (o *PendingBookieOpsStats) GetDataLedgerOpenOp() int64 {
	if o == nil || o.DataLedgerOpenOp == nil {
		var ret int64
		return ret
	}
	return *o.DataLedgerOpenOp
}

// GetDataLedgerOpenOpOk returns a tuple with the DataLedgerOpenOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingBookieOpsStats) GetDataLedgerOpenOpOk() (*int64, bool) {
	if o == nil || o.DataLedgerOpenOp == nil {
		return nil, false
	}
	return o.DataLedgerOpenOp, true
}

// HasDataLedgerOpenOp returns a boolean if a field has been set.
func (o *PendingBookieOpsStats) HasDataLedgerOpenOp() bool {
	if o != nil && o.DataLedgerOpenOp != nil {
		return true
	}

	return false
}

// SetDataLedgerOpenOp gets a reference to the given int64 and assigns it to the DataLedgerOpenOp field.
func (o *PendingBookieOpsStats) SetDataLedgerOpenOp(v int64) {
	o.DataLedgerOpenOp = &v
}

// GetDataLedgerCloseOp returns the DataLedgerCloseOp field value if set, zero value otherwise.
func (o *PendingBookieOpsStats) GetDataLedgerCloseOp() int64 {
	if o == nil || o.DataLedgerCloseOp == nil {
		var ret int64
		return ret
	}
	return *o.DataLedgerCloseOp
}

// GetDataLedgerCloseOpOk returns a tuple with the DataLedgerCloseOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingBookieOpsStats) GetDataLedgerCloseOpOk() (*int64, bool) {
	if o == nil || o.DataLedgerCloseOp == nil {
		return nil, false
	}
	return o.DataLedgerCloseOp, true
}

// HasDataLedgerCloseOp returns a boolean if a field has been set.
func (o *PendingBookieOpsStats) HasDataLedgerCloseOp() bool {
	if o != nil && o.DataLedgerCloseOp != nil {
		return true
	}

	return false
}

// SetDataLedgerCloseOp gets a reference to the given int64 and assigns it to the DataLedgerCloseOp field.
func (o *PendingBookieOpsStats) SetDataLedgerCloseOp(v int64) {
	o.DataLedgerCloseOp = &v
}

// GetDataLedgerCreateOp returns the DataLedgerCreateOp field value if set, zero value otherwise.
func (o *PendingBookieOpsStats) GetDataLedgerCreateOp() int64 {
	if o == nil || o.DataLedgerCreateOp == nil {
		var ret int64
		return ret
	}
	return *o.DataLedgerCreateOp
}

// GetDataLedgerCreateOpOk returns a tuple with the DataLedgerCreateOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingBookieOpsStats) GetDataLedgerCreateOpOk() (*int64, bool) {
	if o == nil || o.DataLedgerCreateOp == nil {
		return nil, false
	}
	return o.DataLedgerCreateOp, true
}

// HasDataLedgerCreateOp returns a boolean if a field has been set.
func (o *PendingBookieOpsStats) HasDataLedgerCreateOp() bool {
	if o != nil && o.DataLedgerCreateOp != nil {
		return true
	}

	return false
}

// SetDataLedgerCreateOp gets a reference to the given int64 and assigns it to the DataLedgerCreateOp field.
func (o *PendingBookieOpsStats) SetDataLedgerCreateOp(v int64) {
	o.DataLedgerCreateOp = &v
}

// GetDataLedgerDeleteOp returns the DataLedgerDeleteOp field value if set, zero value otherwise.
func (o *PendingBookieOpsStats) GetDataLedgerDeleteOp() int64 {
	if o == nil || o.DataLedgerDeleteOp == nil {
		var ret int64
		return ret
	}
	return *o.DataLedgerDeleteOp
}

// GetDataLedgerDeleteOpOk returns a tuple with the DataLedgerDeleteOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingBookieOpsStats) GetDataLedgerDeleteOpOk() (*int64, bool) {
	if o == nil || o.DataLedgerDeleteOp == nil {
		return nil, false
	}
	return o.DataLedgerDeleteOp, true
}

// HasDataLedgerDeleteOp returns a boolean if a field has been set.
func (o *PendingBookieOpsStats) HasDataLedgerDeleteOp() bool {
	if o != nil && o.DataLedgerDeleteOp != nil {
		return true
	}

	return false
}

// SetDataLedgerDeleteOp gets a reference to the given int64 and assigns it to the DataLedgerDeleteOp field.
func (o *PendingBookieOpsStats) SetDataLedgerDeleteOp(v int64) {
	o.DataLedgerDeleteOp = &v
}

// GetCursorLedgerOpenOp returns the CursorLedgerOpenOp field value if set, zero value otherwise.
func (o *PendingBookieOpsStats) GetCursorLedgerOpenOp() int64 {
	if o == nil || o.CursorLedgerOpenOp == nil {
		var ret int64
		return ret
	}
	return *o.CursorLedgerOpenOp
}

// GetCursorLedgerOpenOpOk returns a tuple with the CursorLedgerOpenOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingBookieOpsStats) GetCursorLedgerOpenOpOk() (*int64, bool) {
	if o == nil || o.CursorLedgerOpenOp == nil {
		return nil, false
	}
	return o.CursorLedgerOpenOp, true
}

// HasCursorLedgerOpenOp returns a boolean if a field has been set.
func (o *PendingBookieOpsStats) HasCursorLedgerOpenOp() bool {
	if o != nil && o.CursorLedgerOpenOp != nil {
		return true
	}

	return false
}

// SetCursorLedgerOpenOp gets a reference to the given int64 and assigns it to the CursorLedgerOpenOp field.
func (o *PendingBookieOpsStats) SetCursorLedgerOpenOp(v int64) {
	o.CursorLedgerOpenOp = &v
}

// GetCursorLedgerCloseOp returns the CursorLedgerCloseOp field value if set, zero value otherwise.
func (o *PendingBookieOpsStats) GetCursorLedgerCloseOp() int64 {
	if o == nil || o.CursorLedgerCloseOp == nil {
		var ret int64
		return ret
	}
	return *o.CursorLedgerCloseOp
}

// GetCursorLedgerCloseOpOk returns a tuple with the CursorLedgerCloseOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingBookieOpsStats) GetCursorLedgerCloseOpOk() (*int64, bool) {
	if o == nil || o.CursorLedgerCloseOp == nil {
		return nil, false
	}
	return o.CursorLedgerCloseOp, true
}

// HasCursorLedgerCloseOp returns a boolean if a field has been set.
func (o *PendingBookieOpsStats) HasCursorLedgerCloseOp() bool {
	if o != nil && o.CursorLedgerCloseOp != nil {
		return true
	}

	return false
}

// SetCursorLedgerCloseOp gets a reference to the given int64 and assigns it to the CursorLedgerCloseOp field.
func (o *PendingBookieOpsStats) SetCursorLedgerCloseOp(v int64) {
	o.CursorLedgerCloseOp = &v
}

// GetCursorLedgerCreateOp returns the CursorLedgerCreateOp field value if set, zero value otherwise.
func (o *PendingBookieOpsStats) GetCursorLedgerCreateOp() int64 {
	if o == nil || o.CursorLedgerCreateOp == nil {
		var ret int64
		return ret
	}
	return *o.CursorLedgerCreateOp
}

// GetCursorLedgerCreateOpOk returns a tuple with the CursorLedgerCreateOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingBookieOpsStats) GetCursorLedgerCreateOpOk() (*int64, bool) {
	if o == nil || o.CursorLedgerCreateOp == nil {
		return nil, false
	}
	return o.CursorLedgerCreateOp, true
}

// HasCursorLedgerCreateOp returns a boolean if a field has been set.
func (o *PendingBookieOpsStats) HasCursorLedgerCreateOp() bool {
	if o != nil && o.CursorLedgerCreateOp != nil {
		return true
	}

	return false
}

// SetCursorLedgerCreateOp gets a reference to the given int64 and assigns it to the CursorLedgerCreateOp field.
func (o *PendingBookieOpsStats) SetCursorLedgerCreateOp(v int64) {
	o.CursorLedgerCreateOp = &v
}

// GetCursorLedgerDeleteOp returns the CursorLedgerDeleteOp field value if set, zero value otherwise.
func (o *PendingBookieOpsStats) GetCursorLedgerDeleteOp() int64 {
	if o == nil || o.CursorLedgerDeleteOp == nil {
		var ret int64
		return ret
	}
	return *o.CursorLedgerDeleteOp
}

// GetCursorLedgerDeleteOpOk returns a tuple with the CursorLedgerDeleteOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingBookieOpsStats) GetCursorLedgerDeleteOpOk() (*int64, bool) {
	if o == nil || o.CursorLedgerDeleteOp == nil {
		return nil, false
	}
	return o.CursorLedgerDeleteOp, true
}

// HasCursorLedgerDeleteOp returns a boolean if a field has been set.
func (o *PendingBookieOpsStats) HasCursorLedgerDeleteOp() bool {
	if o != nil && o.CursorLedgerDeleteOp != nil {
		return true
	}

	return false
}

// SetCursorLedgerDeleteOp gets a reference to the given int64 and assigns it to the CursorLedgerDeleteOp field.
func (o *PendingBookieOpsStats) SetCursorLedgerDeleteOp(v int64) {
	o.CursorLedgerDeleteOp = &v
}

func (o PendingBookieOpsStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataLedgerOpenOp != nil {
		toSerialize["dataLedgerOpenOp"] = o.DataLedgerOpenOp
	}
	if o.DataLedgerCloseOp != nil {
		toSerialize["dataLedgerCloseOp"] = o.DataLedgerCloseOp
	}
	if o.DataLedgerCreateOp != nil {
		toSerialize["dataLedgerCreateOp"] = o.DataLedgerCreateOp
	}
	if o.DataLedgerDeleteOp != nil {
		toSerialize["dataLedgerDeleteOp"] = o.DataLedgerDeleteOp
	}
	if o.CursorLedgerOpenOp != nil {
		toSerialize["cursorLedgerOpenOp"] = o.CursorLedgerOpenOp
	}
	if o.CursorLedgerCloseOp != nil {
		toSerialize["cursorLedgerCloseOp"] = o.CursorLedgerCloseOp
	}
	if o.CursorLedgerCreateOp != nil {
		toSerialize["cursorLedgerCreateOp"] = o.CursorLedgerCreateOp
	}
	if o.CursorLedgerDeleteOp != nil {
		toSerialize["cursorLedgerDeleteOp"] = o.CursorLedgerDeleteOp
	}
	return json.Marshal(toSerialize)
}

type NullablePendingBookieOpsStats struct {
	value *PendingBookieOpsStats
	isSet bool
}

func (v NullablePendingBookieOpsStats) Get() *PendingBookieOpsStats {
	return v.value
}

func (v *NullablePendingBookieOpsStats) Set(val *PendingBookieOpsStats) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingBookieOpsStats) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingBookieOpsStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingBookieOpsStats(val *PendingBookieOpsStats) *NullablePendingBookieOpsStats {
	return &NullablePendingBookieOpsStats{value: val, isSet: true}
}

func (v NullablePendingBookieOpsStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingBookieOpsStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


