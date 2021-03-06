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

// PublisherStats struct for PublisherStats
type PublisherStats struct {
	ProducerName *string `json:"producerName,omitempty"`
	MsgThroughputIn *float64 `json:"msgThroughputIn,omitempty"`
	AccessMode *string `json:"accessMode,omitempty"`
	AverageMsgSize *float64 `json:"averageMsgSize,omitempty"`
	ConnectedSince *string `json:"connectedSince,omitempty"`
	MsgRateIn *float64 `json:"msgRateIn,omitempty"`
	ChunkedMessageRate *float64 `json:"chunkedMessageRate,omitempty"`
	ClientVersion *string `json:"clientVersion,omitempty"`
	ProducerId *int64 `json:"producerId,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Address *string `json:"address,omitempty"`
}

// NewPublisherStats instantiates a new PublisherStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublisherStats() *PublisherStats {
	this := PublisherStats{}
	return &this
}

// NewPublisherStatsWithDefaults instantiates a new PublisherStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublisherStatsWithDefaults() *PublisherStats {
	this := PublisherStats{}
	return &this
}

// GetProducerName returns the ProducerName field value if set, zero value otherwise.
func (o *PublisherStats) GetProducerName() string {
	if o == nil || o.ProducerName == nil {
		var ret string
		return ret
	}
	return *o.ProducerName
}

// GetProducerNameOk returns a tuple with the ProducerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetProducerNameOk() (*string, bool) {
	if o == nil || o.ProducerName == nil {
		return nil, false
	}
	return o.ProducerName, true
}

// HasProducerName returns a boolean if a field has been set.
func (o *PublisherStats) HasProducerName() bool {
	if o != nil && o.ProducerName != nil {
		return true
	}

	return false
}

// SetProducerName gets a reference to the given string and assigns it to the ProducerName field.
func (o *PublisherStats) SetProducerName(v string) {
	o.ProducerName = &v
}

// GetMsgThroughputIn returns the MsgThroughputIn field value if set, zero value otherwise.
func (o *PublisherStats) GetMsgThroughputIn() float64 {
	if o == nil || o.MsgThroughputIn == nil {
		var ret float64
		return ret
	}
	return *o.MsgThroughputIn
}

// GetMsgThroughputInOk returns a tuple with the MsgThroughputIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetMsgThroughputInOk() (*float64, bool) {
	if o == nil || o.MsgThroughputIn == nil {
		return nil, false
	}
	return o.MsgThroughputIn, true
}

// HasMsgThroughputIn returns a boolean if a field has been set.
func (o *PublisherStats) HasMsgThroughputIn() bool {
	if o != nil && o.MsgThroughputIn != nil {
		return true
	}

	return false
}

// SetMsgThroughputIn gets a reference to the given float64 and assigns it to the MsgThroughputIn field.
func (o *PublisherStats) SetMsgThroughputIn(v float64) {
	o.MsgThroughputIn = &v
}

// GetAccessMode returns the AccessMode field value if set, zero value otherwise.
func (o *PublisherStats) GetAccessMode() string {
	if o == nil || o.AccessMode == nil {
		var ret string
		return ret
	}
	return *o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetAccessModeOk() (*string, bool) {
	if o == nil || o.AccessMode == nil {
		return nil, false
	}
	return o.AccessMode, true
}

// HasAccessMode returns a boolean if a field has been set.
func (o *PublisherStats) HasAccessMode() bool {
	if o != nil && o.AccessMode != nil {
		return true
	}

	return false
}

// SetAccessMode gets a reference to the given string and assigns it to the AccessMode field.
func (o *PublisherStats) SetAccessMode(v string) {
	o.AccessMode = &v
}

// GetAverageMsgSize returns the AverageMsgSize field value if set, zero value otherwise.
func (o *PublisherStats) GetAverageMsgSize() float64 {
	if o == nil || o.AverageMsgSize == nil {
		var ret float64
		return ret
	}
	return *o.AverageMsgSize
}

// GetAverageMsgSizeOk returns a tuple with the AverageMsgSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetAverageMsgSizeOk() (*float64, bool) {
	if o == nil || o.AverageMsgSize == nil {
		return nil, false
	}
	return o.AverageMsgSize, true
}

// HasAverageMsgSize returns a boolean if a field has been set.
func (o *PublisherStats) HasAverageMsgSize() bool {
	if o != nil && o.AverageMsgSize != nil {
		return true
	}

	return false
}

// SetAverageMsgSize gets a reference to the given float64 and assigns it to the AverageMsgSize field.
func (o *PublisherStats) SetAverageMsgSize(v float64) {
	o.AverageMsgSize = &v
}

// GetConnectedSince returns the ConnectedSince field value if set, zero value otherwise.
func (o *PublisherStats) GetConnectedSince() string {
	if o == nil || o.ConnectedSince == nil {
		var ret string
		return ret
	}
	return *o.ConnectedSince
}

// GetConnectedSinceOk returns a tuple with the ConnectedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetConnectedSinceOk() (*string, bool) {
	if o == nil || o.ConnectedSince == nil {
		return nil, false
	}
	return o.ConnectedSince, true
}

// HasConnectedSince returns a boolean if a field has been set.
func (o *PublisherStats) HasConnectedSince() bool {
	if o != nil && o.ConnectedSince != nil {
		return true
	}

	return false
}

// SetConnectedSince gets a reference to the given string and assigns it to the ConnectedSince field.
func (o *PublisherStats) SetConnectedSince(v string) {
	o.ConnectedSince = &v
}

// GetMsgRateIn returns the MsgRateIn field value if set, zero value otherwise.
func (o *PublisherStats) GetMsgRateIn() float64 {
	if o == nil || o.MsgRateIn == nil {
		var ret float64
		return ret
	}
	return *o.MsgRateIn
}

// GetMsgRateInOk returns a tuple with the MsgRateIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetMsgRateInOk() (*float64, bool) {
	if o == nil || o.MsgRateIn == nil {
		return nil, false
	}
	return o.MsgRateIn, true
}

// HasMsgRateIn returns a boolean if a field has been set.
func (o *PublisherStats) HasMsgRateIn() bool {
	if o != nil && o.MsgRateIn != nil {
		return true
	}

	return false
}

// SetMsgRateIn gets a reference to the given float64 and assigns it to the MsgRateIn field.
func (o *PublisherStats) SetMsgRateIn(v float64) {
	o.MsgRateIn = &v
}

// GetChunkedMessageRate returns the ChunkedMessageRate field value if set, zero value otherwise.
func (o *PublisherStats) GetChunkedMessageRate() float64 {
	if o == nil || o.ChunkedMessageRate == nil {
		var ret float64
		return ret
	}
	return *o.ChunkedMessageRate
}

// GetChunkedMessageRateOk returns a tuple with the ChunkedMessageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetChunkedMessageRateOk() (*float64, bool) {
	if o == nil || o.ChunkedMessageRate == nil {
		return nil, false
	}
	return o.ChunkedMessageRate, true
}

// HasChunkedMessageRate returns a boolean if a field has been set.
func (o *PublisherStats) HasChunkedMessageRate() bool {
	if o != nil && o.ChunkedMessageRate != nil {
		return true
	}

	return false
}

// SetChunkedMessageRate gets a reference to the given float64 and assigns it to the ChunkedMessageRate field.
func (o *PublisherStats) SetChunkedMessageRate(v float64) {
	o.ChunkedMessageRate = &v
}

// GetClientVersion returns the ClientVersion field value if set, zero value otherwise.
func (o *PublisherStats) GetClientVersion() string {
	if o == nil || o.ClientVersion == nil {
		var ret string
		return ret
	}
	return *o.ClientVersion
}

// GetClientVersionOk returns a tuple with the ClientVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetClientVersionOk() (*string, bool) {
	if o == nil || o.ClientVersion == nil {
		return nil, false
	}
	return o.ClientVersion, true
}

// HasClientVersion returns a boolean if a field has been set.
func (o *PublisherStats) HasClientVersion() bool {
	if o != nil && o.ClientVersion != nil {
		return true
	}

	return false
}

// SetClientVersion gets a reference to the given string and assigns it to the ClientVersion field.
func (o *PublisherStats) SetClientVersion(v string) {
	o.ClientVersion = &v
}

// GetProducerId returns the ProducerId field value if set, zero value otherwise.
func (o *PublisherStats) GetProducerId() int64 {
	if o == nil || o.ProducerId == nil {
		var ret int64
		return ret
	}
	return *o.ProducerId
}

// GetProducerIdOk returns a tuple with the ProducerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetProducerIdOk() (*int64, bool) {
	if o == nil || o.ProducerId == nil {
		return nil, false
	}
	return o.ProducerId, true
}

// HasProducerId returns a boolean if a field has been set.
func (o *PublisherStats) HasProducerId() bool {
	if o != nil && o.ProducerId != nil {
		return true
	}

	return false
}

// SetProducerId gets a reference to the given int64 and assigns it to the ProducerId field.
func (o *PublisherStats) SetProducerId(v int64) {
	o.ProducerId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PublisherStats) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PublisherStats) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *PublisherStats) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PublisherStats) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublisherStats) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PublisherStats) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *PublisherStats) SetAddress(v string) {
	o.Address = &v
}

func (o PublisherStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerName != nil {
		toSerialize["producerName"] = o.ProducerName
	}
	if o.MsgThroughputIn != nil {
		toSerialize["msgThroughputIn"] = o.MsgThroughputIn
	}
	if o.AccessMode != nil {
		toSerialize["accessMode"] = o.AccessMode
	}
	if o.AverageMsgSize != nil {
		toSerialize["averageMsgSize"] = o.AverageMsgSize
	}
	if o.ConnectedSince != nil {
		toSerialize["connectedSince"] = o.ConnectedSince
	}
	if o.MsgRateIn != nil {
		toSerialize["msgRateIn"] = o.MsgRateIn
	}
	if o.ChunkedMessageRate != nil {
		toSerialize["chunkedMessageRate"] = o.ChunkedMessageRate
	}
	if o.ClientVersion != nil {
		toSerialize["clientVersion"] = o.ClientVersion
	}
	if o.ProducerId != nil {
		toSerialize["producerId"] = o.ProducerId
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullablePublisherStats struct {
	value *PublisherStats
	isSet bool
}

func (v NullablePublisherStats) Get() *PublisherStats {
	return v.value
}

func (v *NullablePublisherStats) Set(val *PublisherStats) {
	v.value = val
	v.isSet = true
}

func (v NullablePublisherStats) IsSet() bool {
	return v.isSet
}

func (v *NullablePublisherStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublisherStats(val *PublisherStats) *NullablePublisherStats {
	return &NullablePublisherStats{value: val, isSet: true}
}

func (v NullablePublisherStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublisherStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


