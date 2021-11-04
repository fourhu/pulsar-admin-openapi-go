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

// BrokerNamespaceIsolationData The namespace isolation data for a given broker
type BrokerNamespaceIsolationData struct {
	// The broker name
	BrokerName *string `json:"brokerName,omitempty"`
	// Is Primary broker
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// The namespace-isolation policies attached to this broker
	NamespaceRegex *[]string `json:"namespaceRegex,omitempty"`
	// Policy name
	PolicyName *string `json:"policyName,omitempty"`
}

// NewBrokerNamespaceIsolationData instantiates a new BrokerNamespaceIsolationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerNamespaceIsolationData() *BrokerNamespaceIsolationData {
	this := BrokerNamespaceIsolationData{}
	return &this
}

// NewBrokerNamespaceIsolationDataWithDefaults instantiates a new BrokerNamespaceIsolationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerNamespaceIsolationDataWithDefaults() *BrokerNamespaceIsolationData {
	this := BrokerNamespaceIsolationData{}
	return &this
}

// GetBrokerName returns the BrokerName field value if set, zero value otherwise.
func (o *BrokerNamespaceIsolationData) GetBrokerName() string {
	if o == nil || o.BrokerName == nil {
		var ret string
		return ret
	}
	return *o.BrokerName
}

// GetBrokerNameOk returns a tuple with the BrokerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerNamespaceIsolationData) GetBrokerNameOk() (*string, bool) {
	if o == nil || o.BrokerName == nil {
		return nil, false
	}
	return o.BrokerName, true
}

// HasBrokerName returns a boolean if a field has been set.
func (o *BrokerNamespaceIsolationData) HasBrokerName() bool {
	if o != nil && o.BrokerName != nil {
		return true
	}

	return false
}

// SetBrokerName gets a reference to the given string and assigns it to the BrokerName field.
func (o *BrokerNamespaceIsolationData) SetBrokerName(v string) {
	o.BrokerName = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *BrokerNamespaceIsolationData) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerNamespaceIsolationData) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *BrokerNamespaceIsolationData) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *BrokerNamespaceIsolationData) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetNamespaceRegex returns the NamespaceRegex field value if set, zero value otherwise.
func (o *BrokerNamespaceIsolationData) GetNamespaceRegex() []string {
	if o == nil || o.NamespaceRegex == nil {
		var ret []string
		return ret
	}
	return *o.NamespaceRegex
}

// GetNamespaceRegexOk returns a tuple with the NamespaceRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerNamespaceIsolationData) GetNamespaceRegexOk() (*[]string, bool) {
	if o == nil || o.NamespaceRegex == nil {
		return nil, false
	}
	return o.NamespaceRegex, true
}

// HasNamespaceRegex returns a boolean if a field has been set.
func (o *BrokerNamespaceIsolationData) HasNamespaceRegex() bool {
	if o != nil && o.NamespaceRegex != nil {
		return true
	}

	return false
}

// SetNamespaceRegex gets a reference to the given []string and assigns it to the NamespaceRegex field.
func (o *BrokerNamespaceIsolationData) SetNamespaceRegex(v []string) {
	o.NamespaceRegex = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *BrokerNamespaceIsolationData) GetPolicyName() string {
	if o == nil || o.PolicyName == nil {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerNamespaceIsolationData) GetPolicyNameOk() (*string, bool) {
	if o == nil || o.PolicyName == nil {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *BrokerNamespaceIsolationData) HasPolicyName() bool {
	if o != nil && o.PolicyName != nil {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *BrokerNamespaceIsolationData) SetPolicyName(v string) {
	o.PolicyName = &v
}

func (o BrokerNamespaceIsolationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BrokerName != nil {
		toSerialize["brokerName"] = o.BrokerName
	}
	if o.IsPrimary != nil {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if o.NamespaceRegex != nil {
		toSerialize["namespaceRegex"] = o.NamespaceRegex
	}
	if o.PolicyName != nil {
		toSerialize["policyName"] = o.PolicyName
	}
	return json.Marshal(toSerialize)
}

type NullableBrokerNamespaceIsolationData struct {
	value *BrokerNamespaceIsolationData
	isSet bool
}

func (v NullableBrokerNamespaceIsolationData) Get() *BrokerNamespaceIsolationData {
	return v.value
}

func (v *NullableBrokerNamespaceIsolationData) Set(val *BrokerNamespaceIsolationData) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerNamespaceIsolationData) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerNamespaceIsolationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerNamespaceIsolationData(val *BrokerNamespaceIsolationData) *NullableBrokerNamespaceIsolationData {
	return &NullableBrokerNamespaceIsolationData{value: val, isSet: true}
}

func (v NullableBrokerNamespaceIsolationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerNamespaceIsolationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


