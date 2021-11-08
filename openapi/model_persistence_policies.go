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

// PersistencePolicies struct for PersistencePolicies
type PersistencePolicies struct {
	BookkeeperEnsemble *int32 `json:"bookkeeperEnsemble,omitempty"`
	BookkeeperWriteQuorum *int32 `json:"bookkeeperWriteQuorum,omitempty"`
	BookkeeperAckQuorum *int32 `json:"bookkeeperAckQuorum,omitempty"`
	ManagedLedgerMaxMarkDeleteRate *float64 `json:"managedLedgerMaxMarkDeleteRate,omitempty"`
}

// NewPersistencePolicies instantiates a new PersistencePolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersistencePolicies() *PersistencePolicies {
	this := PersistencePolicies{}
	return &this
}

// NewPersistencePoliciesWithDefaults instantiates a new PersistencePolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersistencePoliciesWithDefaults() *PersistencePolicies {
	this := PersistencePolicies{}
	return &this
}

// GetBookkeeperEnsemble returns the BookkeeperEnsemble field value if set, zero value otherwise.
func (o *PersistencePolicies) GetBookkeeperEnsemble() int32 {
	if o == nil || o.BookkeeperEnsemble == nil {
		var ret int32
		return ret
	}
	return *o.BookkeeperEnsemble
}

// GetBookkeeperEnsembleOk returns a tuple with the BookkeeperEnsemble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistencePolicies) GetBookkeeperEnsembleOk() (*int32, bool) {
	if o == nil || o.BookkeeperEnsemble == nil {
		return nil, false
	}
	return o.BookkeeperEnsemble, true
}

// HasBookkeeperEnsemble returns a boolean if a field has been set.
func (o *PersistencePolicies) HasBookkeeperEnsemble() bool {
	if o != nil && o.BookkeeperEnsemble != nil {
		return true
	}

	return false
}

// SetBookkeeperEnsemble gets a reference to the given int32 and assigns it to the BookkeeperEnsemble field.
func (o *PersistencePolicies) SetBookkeeperEnsemble(v int32) {
	o.BookkeeperEnsemble = &v
}

// GetBookkeeperWriteQuorum returns the BookkeeperWriteQuorum field value if set, zero value otherwise.
func (o *PersistencePolicies) GetBookkeeperWriteQuorum() int32 {
	if o == nil || o.BookkeeperWriteQuorum == nil {
		var ret int32
		return ret
	}
	return *o.BookkeeperWriteQuorum
}

// GetBookkeeperWriteQuorumOk returns a tuple with the BookkeeperWriteQuorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistencePolicies) GetBookkeeperWriteQuorumOk() (*int32, bool) {
	if o == nil || o.BookkeeperWriteQuorum == nil {
		return nil, false
	}
	return o.BookkeeperWriteQuorum, true
}

// HasBookkeeperWriteQuorum returns a boolean if a field has been set.
func (o *PersistencePolicies) HasBookkeeperWriteQuorum() bool {
	if o != nil && o.BookkeeperWriteQuorum != nil {
		return true
	}

	return false
}

// SetBookkeeperWriteQuorum gets a reference to the given int32 and assigns it to the BookkeeperWriteQuorum field.
func (o *PersistencePolicies) SetBookkeeperWriteQuorum(v int32) {
	o.BookkeeperWriteQuorum = &v
}

// GetBookkeeperAckQuorum returns the BookkeeperAckQuorum field value if set, zero value otherwise.
func (o *PersistencePolicies) GetBookkeeperAckQuorum() int32 {
	if o == nil || o.BookkeeperAckQuorum == nil {
		var ret int32
		return ret
	}
	return *o.BookkeeperAckQuorum
}

// GetBookkeeperAckQuorumOk returns a tuple with the BookkeeperAckQuorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistencePolicies) GetBookkeeperAckQuorumOk() (*int32, bool) {
	if o == nil || o.BookkeeperAckQuorum == nil {
		return nil, false
	}
	return o.BookkeeperAckQuorum, true
}

// HasBookkeeperAckQuorum returns a boolean if a field has been set.
func (o *PersistencePolicies) HasBookkeeperAckQuorum() bool {
	if o != nil && o.BookkeeperAckQuorum != nil {
		return true
	}

	return false
}

// SetBookkeeperAckQuorum gets a reference to the given int32 and assigns it to the BookkeeperAckQuorum field.
func (o *PersistencePolicies) SetBookkeeperAckQuorum(v int32) {
	o.BookkeeperAckQuorum = &v
}

// GetManagedLedgerMaxMarkDeleteRate returns the ManagedLedgerMaxMarkDeleteRate field value if set, zero value otherwise.
func (o *PersistencePolicies) GetManagedLedgerMaxMarkDeleteRate() float64 {
	if o == nil || o.ManagedLedgerMaxMarkDeleteRate == nil {
		var ret float64
		return ret
	}
	return *o.ManagedLedgerMaxMarkDeleteRate
}

// GetManagedLedgerMaxMarkDeleteRateOk returns a tuple with the ManagedLedgerMaxMarkDeleteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistencePolicies) GetManagedLedgerMaxMarkDeleteRateOk() (*float64, bool) {
	if o == nil || o.ManagedLedgerMaxMarkDeleteRate == nil {
		return nil, false
	}
	return o.ManagedLedgerMaxMarkDeleteRate, true
}

// HasManagedLedgerMaxMarkDeleteRate returns a boolean if a field has been set.
func (o *PersistencePolicies) HasManagedLedgerMaxMarkDeleteRate() bool {
	if o != nil && o.ManagedLedgerMaxMarkDeleteRate != nil {
		return true
	}

	return false
}

// SetManagedLedgerMaxMarkDeleteRate gets a reference to the given float64 and assigns it to the ManagedLedgerMaxMarkDeleteRate field.
func (o *PersistencePolicies) SetManagedLedgerMaxMarkDeleteRate(v float64) {
	o.ManagedLedgerMaxMarkDeleteRate = &v
}

func (o PersistencePolicies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BookkeeperEnsemble != nil {
		toSerialize["bookkeeperEnsemble"] = o.BookkeeperEnsemble
	}
	if o.BookkeeperWriteQuorum != nil {
		toSerialize["bookkeeperWriteQuorum"] = o.BookkeeperWriteQuorum
	}
	if o.BookkeeperAckQuorum != nil {
		toSerialize["bookkeeperAckQuorum"] = o.BookkeeperAckQuorum
	}
	if o.ManagedLedgerMaxMarkDeleteRate != nil {
		toSerialize["managedLedgerMaxMarkDeleteRate"] = o.ManagedLedgerMaxMarkDeleteRate
	}
	return json.Marshal(toSerialize)
}

type NullablePersistencePolicies struct {
	value *PersistencePolicies
	isSet bool
}

func (v NullablePersistencePolicies) Get() *PersistencePolicies {
	return v.value
}

func (v *NullablePersistencePolicies) Set(val *PersistencePolicies) {
	v.value = val
	v.isSet = true
}

func (v NullablePersistencePolicies) IsSet() bool {
	return v.isSet
}

func (v *NullablePersistencePolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersistencePolicies(val *PersistencePolicies) *NullablePersistencePolicies {
	return &NullablePersistencePolicies{value: val, isSet: true}
}

func (v NullablePersistencePolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersistencePolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


