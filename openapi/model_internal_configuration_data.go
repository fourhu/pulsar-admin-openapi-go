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

// InternalConfigurationData struct for InternalConfigurationData
type InternalConfigurationData struct {
	ZookeeperServers *string `json:"zookeeperServers,omitempty"`
	ConfigurationStoreServers *string `json:"configurationStoreServers,omitempty"`
	LedgersRootPath *string `json:"ledgersRootPath,omitempty"`
	BookkeeperMetadataServiceUri *string `json:"bookkeeperMetadataServiceUri,omitempty"`
	StateStorageServiceUrl *string `json:"stateStorageServiceUrl,omitempty"`
}

// NewInternalConfigurationData instantiates a new InternalConfigurationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalConfigurationData() *InternalConfigurationData {
	this := InternalConfigurationData{}
	return &this
}

// NewInternalConfigurationDataWithDefaults instantiates a new InternalConfigurationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalConfigurationDataWithDefaults() *InternalConfigurationData {
	this := InternalConfigurationData{}
	return &this
}

// GetZookeeperServers returns the ZookeeperServers field value if set, zero value otherwise.
func (o *InternalConfigurationData) GetZookeeperServers() string {
	if o == nil || o.ZookeeperServers == nil {
		var ret string
		return ret
	}
	return *o.ZookeeperServers
}

// GetZookeeperServersOk returns a tuple with the ZookeeperServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalConfigurationData) GetZookeeperServersOk() (*string, bool) {
	if o == nil || o.ZookeeperServers == nil {
		return nil, false
	}
	return o.ZookeeperServers, true
}

// HasZookeeperServers returns a boolean if a field has been set.
func (o *InternalConfigurationData) HasZookeeperServers() bool {
	if o != nil && o.ZookeeperServers != nil {
		return true
	}

	return false
}

// SetZookeeperServers gets a reference to the given string and assigns it to the ZookeeperServers field.
func (o *InternalConfigurationData) SetZookeeperServers(v string) {
	o.ZookeeperServers = &v
}

// GetConfigurationStoreServers returns the ConfigurationStoreServers field value if set, zero value otherwise.
func (o *InternalConfigurationData) GetConfigurationStoreServers() string {
	if o == nil || o.ConfigurationStoreServers == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationStoreServers
}

// GetConfigurationStoreServersOk returns a tuple with the ConfigurationStoreServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalConfigurationData) GetConfigurationStoreServersOk() (*string, bool) {
	if o == nil || o.ConfigurationStoreServers == nil {
		return nil, false
	}
	return o.ConfigurationStoreServers, true
}

// HasConfigurationStoreServers returns a boolean if a field has been set.
func (o *InternalConfigurationData) HasConfigurationStoreServers() bool {
	if o != nil && o.ConfigurationStoreServers != nil {
		return true
	}

	return false
}

// SetConfigurationStoreServers gets a reference to the given string and assigns it to the ConfigurationStoreServers field.
func (o *InternalConfigurationData) SetConfigurationStoreServers(v string) {
	o.ConfigurationStoreServers = &v
}

// GetLedgersRootPath returns the LedgersRootPath field value if set, zero value otherwise.
func (o *InternalConfigurationData) GetLedgersRootPath() string {
	if o == nil || o.LedgersRootPath == nil {
		var ret string
		return ret
	}
	return *o.LedgersRootPath
}

// GetLedgersRootPathOk returns a tuple with the LedgersRootPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalConfigurationData) GetLedgersRootPathOk() (*string, bool) {
	if o == nil || o.LedgersRootPath == nil {
		return nil, false
	}
	return o.LedgersRootPath, true
}

// HasLedgersRootPath returns a boolean if a field has been set.
func (o *InternalConfigurationData) HasLedgersRootPath() bool {
	if o != nil && o.LedgersRootPath != nil {
		return true
	}

	return false
}

// SetLedgersRootPath gets a reference to the given string and assigns it to the LedgersRootPath field.
func (o *InternalConfigurationData) SetLedgersRootPath(v string) {
	o.LedgersRootPath = &v
}

// GetBookkeeperMetadataServiceUri returns the BookkeeperMetadataServiceUri field value if set, zero value otherwise.
func (o *InternalConfigurationData) GetBookkeeperMetadataServiceUri() string {
	if o == nil || o.BookkeeperMetadataServiceUri == nil {
		var ret string
		return ret
	}
	return *o.BookkeeperMetadataServiceUri
}

// GetBookkeeperMetadataServiceUriOk returns a tuple with the BookkeeperMetadataServiceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalConfigurationData) GetBookkeeperMetadataServiceUriOk() (*string, bool) {
	if o == nil || o.BookkeeperMetadataServiceUri == nil {
		return nil, false
	}
	return o.BookkeeperMetadataServiceUri, true
}

// HasBookkeeperMetadataServiceUri returns a boolean if a field has been set.
func (o *InternalConfigurationData) HasBookkeeperMetadataServiceUri() bool {
	if o != nil && o.BookkeeperMetadataServiceUri != nil {
		return true
	}

	return false
}

// SetBookkeeperMetadataServiceUri gets a reference to the given string and assigns it to the BookkeeperMetadataServiceUri field.
func (o *InternalConfigurationData) SetBookkeeperMetadataServiceUri(v string) {
	o.BookkeeperMetadataServiceUri = &v
}

// GetStateStorageServiceUrl returns the StateStorageServiceUrl field value if set, zero value otherwise.
func (o *InternalConfigurationData) GetStateStorageServiceUrl() string {
	if o == nil || o.StateStorageServiceUrl == nil {
		var ret string
		return ret
	}
	return *o.StateStorageServiceUrl
}

// GetStateStorageServiceUrlOk returns a tuple with the StateStorageServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalConfigurationData) GetStateStorageServiceUrlOk() (*string, bool) {
	if o == nil || o.StateStorageServiceUrl == nil {
		return nil, false
	}
	return o.StateStorageServiceUrl, true
}

// HasStateStorageServiceUrl returns a boolean if a field has been set.
func (o *InternalConfigurationData) HasStateStorageServiceUrl() bool {
	if o != nil && o.StateStorageServiceUrl != nil {
		return true
	}

	return false
}

// SetStateStorageServiceUrl gets a reference to the given string and assigns it to the StateStorageServiceUrl field.
func (o *InternalConfigurationData) SetStateStorageServiceUrl(v string) {
	o.StateStorageServiceUrl = &v
}

func (o InternalConfigurationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZookeeperServers != nil {
		toSerialize["zookeeperServers"] = o.ZookeeperServers
	}
	if o.ConfigurationStoreServers != nil {
		toSerialize["configurationStoreServers"] = o.ConfigurationStoreServers
	}
	if o.LedgersRootPath != nil {
		toSerialize["ledgersRootPath"] = o.LedgersRootPath
	}
	if o.BookkeeperMetadataServiceUri != nil {
		toSerialize["bookkeeperMetadataServiceUri"] = o.BookkeeperMetadataServiceUri
	}
	if o.StateStorageServiceUrl != nil {
		toSerialize["stateStorageServiceUrl"] = o.StateStorageServiceUrl
	}
	return json.Marshal(toSerialize)
}

type NullableInternalConfigurationData struct {
	value *InternalConfigurationData
	isSet bool
}

func (v NullableInternalConfigurationData) Get() *InternalConfigurationData {
	return v.value
}

func (v *NullableInternalConfigurationData) Set(val *InternalConfigurationData) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalConfigurationData) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalConfigurationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalConfigurationData(val *InternalConfigurationData) *NullableInternalConfigurationData {
	return &NullableInternalConfigurationData{value: val, isSet: true}
}

func (v NullableInternalConfigurationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalConfigurationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


