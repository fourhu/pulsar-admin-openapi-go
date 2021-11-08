# AuthPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicAuthentication** | Pointer to [**map[string]map[string][]string**](map.md) |  | [optional] 
**SubscriptionAuthentication** | Pointer to **map[string][]string** |  | [optional] 
**NamespaceAuthentication** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewAuthPolicies

`func NewAuthPolicies() *AuthPolicies`

NewAuthPolicies instantiates a new AuthPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthPoliciesWithDefaults

`func NewAuthPoliciesWithDefaults() *AuthPolicies`

NewAuthPoliciesWithDefaults instantiates a new AuthPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicAuthentication

`func (o *AuthPolicies) GetTopicAuthentication() map[string]map[string][]string`

GetTopicAuthentication returns the TopicAuthentication field if non-nil, zero value otherwise.

### GetTopicAuthenticationOk

`func (o *AuthPolicies) GetTopicAuthenticationOk() (*map[string]map[string][]string, bool)`

GetTopicAuthenticationOk returns a tuple with the TopicAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicAuthentication

`func (o *AuthPolicies) SetTopicAuthentication(v map[string]map[string][]string)`

SetTopicAuthentication sets TopicAuthentication field to given value.

### HasTopicAuthentication

`func (o *AuthPolicies) HasTopicAuthentication() bool`

HasTopicAuthentication returns a boolean if a field has been set.

### GetSubscriptionAuthentication

`func (o *AuthPolicies) GetSubscriptionAuthentication() map[string][]string`

GetSubscriptionAuthentication returns the SubscriptionAuthentication field if non-nil, zero value otherwise.

### GetSubscriptionAuthenticationOk

`func (o *AuthPolicies) GetSubscriptionAuthenticationOk() (*map[string][]string, bool)`

GetSubscriptionAuthenticationOk returns a tuple with the SubscriptionAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAuthentication

`func (o *AuthPolicies) SetSubscriptionAuthentication(v map[string][]string)`

SetSubscriptionAuthentication sets SubscriptionAuthentication field to given value.

### HasSubscriptionAuthentication

`func (o *AuthPolicies) HasSubscriptionAuthentication() bool`

HasSubscriptionAuthentication returns a boolean if a field has been set.

### GetNamespaceAuthentication

`func (o *AuthPolicies) GetNamespaceAuthentication() map[string][]string`

GetNamespaceAuthentication returns the NamespaceAuthentication field if non-nil, zero value otherwise.

### GetNamespaceAuthenticationOk

`func (o *AuthPolicies) GetNamespaceAuthenticationOk() (*map[string][]string, bool)`

GetNamespaceAuthenticationOk returns a tuple with the NamespaceAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceAuthentication

`func (o *AuthPolicies) SetNamespaceAuthentication(v map[string][]string)`

SetNamespaceAuthentication sets NamespaceAuthentication field to given value.

### HasNamespaceAuthentication

`func (o *AuthPolicies) HasNamespaceAuthentication() bool`

HasNamespaceAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


