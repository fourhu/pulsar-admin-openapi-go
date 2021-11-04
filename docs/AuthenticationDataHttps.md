# AuthenticationDataHttps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | Pointer to **string** |  | [optional] 
**TlsCertificates** | Pointer to [**[]X509Certificate**](X509Certificate.md) |  | [optional] 
**HttpAuthType** | Pointer to **string** |  | [optional] 
**PeerAddress** | Pointer to **map[string]interface{}** |  | [optional] 
**CommandData** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticationDataHttps

`func NewAuthenticationDataHttps() *AuthenticationDataHttps`

NewAuthenticationDataHttps instantiates a new AuthenticationDataHttps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationDataHttpsWithDefaults

`func NewAuthenticationDataHttpsWithDefaults() *AuthenticationDataHttps`

NewAuthenticationDataHttpsWithDefaults instantiates a new AuthenticationDataHttps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *AuthenticationDataHttps) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AuthenticationDataHttps) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AuthenticationDataHttps) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *AuthenticationDataHttps) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTlsCertificates

`func (o *AuthenticationDataHttps) GetTlsCertificates() []X509Certificate`

GetTlsCertificates returns the TlsCertificates field if non-nil, zero value otherwise.

### GetTlsCertificatesOk

`func (o *AuthenticationDataHttps) GetTlsCertificatesOk() (*[]X509Certificate, bool)`

GetTlsCertificatesOk returns a tuple with the TlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificates

`func (o *AuthenticationDataHttps) SetTlsCertificates(v []X509Certificate)`

SetTlsCertificates sets TlsCertificates field to given value.

### HasTlsCertificates

`func (o *AuthenticationDataHttps) HasTlsCertificates() bool`

HasTlsCertificates returns a boolean if a field has been set.

### GetHttpAuthType

`func (o *AuthenticationDataHttps) GetHttpAuthType() string`

GetHttpAuthType returns the HttpAuthType field if non-nil, zero value otherwise.

### GetHttpAuthTypeOk

`func (o *AuthenticationDataHttps) GetHttpAuthTypeOk() (*string, bool)`

GetHttpAuthTypeOk returns a tuple with the HttpAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAuthType

`func (o *AuthenticationDataHttps) SetHttpAuthType(v string)`

SetHttpAuthType sets HttpAuthType field to given value.

### HasHttpAuthType

`func (o *AuthenticationDataHttps) HasHttpAuthType() bool`

HasHttpAuthType returns a boolean if a field has been set.

### GetPeerAddress

`func (o *AuthenticationDataHttps) GetPeerAddress() map[string]interface{}`

GetPeerAddress returns the PeerAddress field if non-nil, zero value otherwise.

### GetPeerAddressOk

`func (o *AuthenticationDataHttps) GetPeerAddressOk() (*map[string]interface{}, bool)`

GetPeerAddressOk returns a tuple with the PeerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAddress

`func (o *AuthenticationDataHttps) SetPeerAddress(v map[string]interface{})`

SetPeerAddress sets PeerAddress field to given value.

### HasPeerAddress

`func (o *AuthenticationDataHttps) HasPeerAddress() bool`

HasPeerAddress returns a boolean if a field has been set.

### GetCommandData

`func (o *AuthenticationDataHttps) GetCommandData() string`

GetCommandData returns the CommandData field if non-nil, zero value otherwise.

### GetCommandDataOk

`func (o *AuthenticationDataHttps) GetCommandDataOk() (*string, bool)`

GetCommandDataOk returns a tuple with the CommandData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandData

`func (o *AuthenticationDataHttps) SetCommandData(v string)`

SetCommandData sets CommandData field to given value.

### HasCommandData

`func (o *AuthenticationDataHttps) HasCommandData() bool`

HasCommandData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


