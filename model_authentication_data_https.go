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

// AuthenticationDataHttps struct for AuthenticationDataHttps
type AuthenticationDataHttps struct {
	Subscription *string `json:"subscription,omitempty"`
	TlsCertificates *[]X509Certificate `json:"tlsCertificates,omitempty"`
	HttpAuthType *string `json:"httpAuthType,omitempty"`
	PeerAddress *map[string]interface{} `json:"peerAddress,omitempty"`
	CommandData *string `json:"commandData,omitempty"`
}

// NewAuthenticationDataHttps instantiates a new AuthenticationDataHttps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationDataHttps() *AuthenticationDataHttps {
	this := AuthenticationDataHttps{}
	return &this
}

// NewAuthenticationDataHttpsWithDefaults instantiates a new AuthenticationDataHttps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationDataHttpsWithDefaults() *AuthenticationDataHttps {
	this := AuthenticationDataHttps{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *AuthenticationDataHttps) GetSubscription() string {
	if o == nil || o.Subscription == nil {
		var ret string
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationDataHttps) GetSubscriptionOk() (*string, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *AuthenticationDataHttps) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given string and assigns it to the Subscription field.
func (o *AuthenticationDataHttps) SetSubscription(v string) {
	o.Subscription = &v
}

// GetTlsCertificates returns the TlsCertificates field value if set, zero value otherwise.
func (o *AuthenticationDataHttps) GetTlsCertificates() []X509Certificate {
	if o == nil || o.TlsCertificates == nil {
		var ret []X509Certificate
		return ret
	}
	return *o.TlsCertificates
}

// GetTlsCertificatesOk returns a tuple with the TlsCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationDataHttps) GetTlsCertificatesOk() (*[]X509Certificate, bool) {
	if o == nil || o.TlsCertificates == nil {
		return nil, false
	}
	return o.TlsCertificates, true
}

// HasTlsCertificates returns a boolean if a field has been set.
func (o *AuthenticationDataHttps) HasTlsCertificates() bool {
	if o != nil && o.TlsCertificates != nil {
		return true
	}

	return false
}

// SetTlsCertificates gets a reference to the given []X509Certificate and assigns it to the TlsCertificates field.
func (o *AuthenticationDataHttps) SetTlsCertificates(v []X509Certificate) {
	o.TlsCertificates = &v
}

// GetHttpAuthType returns the HttpAuthType field value if set, zero value otherwise.
func (o *AuthenticationDataHttps) GetHttpAuthType() string {
	if o == nil || o.HttpAuthType == nil {
		var ret string
		return ret
	}
	return *o.HttpAuthType
}

// GetHttpAuthTypeOk returns a tuple with the HttpAuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationDataHttps) GetHttpAuthTypeOk() (*string, bool) {
	if o == nil || o.HttpAuthType == nil {
		return nil, false
	}
	return o.HttpAuthType, true
}

// HasHttpAuthType returns a boolean if a field has been set.
func (o *AuthenticationDataHttps) HasHttpAuthType() bool {
	if o != nil && o.HttpAuthType != nil {
		return true
	}

	return false
}

// SetHttpAuthType gets a reference to the given string and assigns it to the HttpAuthType field.
func (o *AuthenticationDataHttps) SetHttpAuthType(v string) {
	o.HttpAuthType = &v
}

// GetPeerAddress returns the PeerAddress field value if set, zero value otherwise.
func (o *AuthenticationDataHttps) GetPeerAddress() map[string]interface{} {
	if o == nil || o.PeerAddress == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.PeerAddress
}

// GetPeerAddressOk returns a tuple with the PeerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationDataHttps) GetPeerAddressOk() (*map[string]interface{}, bool) {
	if o == nil || o.PeerAddress == nil {
		return nil, false
	}
	return o.PeerAddress, true
}

// HasPeerAddress returns a boolean if a field has been set.
func (o *AuthenticationDataHttps) HasPeerAddress() bool {
	if o != nil && o.PeerAddress != nil {
		return true
	}

	return false
}

// SetPeerAddress gets a reference to the given map[string]interface{} and assigns it to the PeerAddress field.
func (o *AuthenticationDataHttps) SetPeerAddress(v map[string]interface{}) {
	o.PeerAddress = &v
}

// GetCommandData returns the CommandData field value if set, zero value otherwise.
func (o *AuthenticationDataHttps) GetCommandData() string {
	if o == nil || o.CommandData == nil {
		var ret string
		return ret
	}
	return *o.CommandData
}

// GetCommandDataOk returns a tuple with the CommandData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationDataHttps) GetCommandDataOk() (*string, bool) {
	if o == nil || o.CommandData == nil {
		return nil, false
	}
	return o.CommandData, true
}

// HasCommandData returns a boolean if a field has been set.
func (o *AuthenticationDataHttps) HasCommandData() bool {
	if o != nil && o.CommandData != nil {
		return true
	}

	return false
}

// SetCommandData gets a reference to the given string and assigns it to the CommandData field.
func (o *AuthenticationDataHttps) SetCommandData(v string) {
	o.CommandData = &v
}

func (o AuthenticationDataHttps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	if o.TlsCertificates != nil {
		toSerialize["tlsCertificates"] = o.TlsCertificates
	}
	if o.HttpAuthType != nil {
		toSerialize["httpAuthType"] = o.HttpAuthType
	}
	if o.PeerAddress != nil {
		toSerialize["peerAddress"] = o.PeerAddress
	}
	if o.CommandData != nil {
		toSerialize["commandData"] = o.CommandData
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticationDataHttps struct {
	value *AuthenticationDataHttps
	isSet bool
}

func (v NullableAuthenticationDataHttps) Get() *AuthenticationDataHttps {
	return v.value
}

func (v *NullableAuthenticationDataHttps) Set(val *AuthenticationDataHttps) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationDataHttps) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationDataHttps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationDataHttps(val *AuthenticationDataHttps) *NullableAuthenticationDataHttps {
	return &NullableAuthenticationDataHttps{value: val, isSet: true}
}

func (v NullableAuthenticationDataHttps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationDataHttps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


