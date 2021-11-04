# AuthPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationAuth** | Pointer to [**map[string]map[string][]string**](map.md) |  | [optional] 
**NamespaceAuth** | Pointer to **map[string][]string** |  | [optional] 
**SubscriptionAuthRoles** | Pointer to **map[string][]string** |  | [optional] 

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

### GetDestinationAuth

`func (o *AuthPolicies) GetDestinationAuth() map[string]map[string][]string`

GetDestinationAuth returns the DestinationAuth field if non-nil, zero value otherwise.

### GetDestinationAuthOk

`func (o *AuthPolicies) GetDestinationAuthOk() (*map[string]map[string][]string, bool)`

GetDestinationAuthOk returns a tuple with the DestinationAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAuth

`func (o *AuthPolicies) SetDestinationAuth(v map[string]map[string][]string)`

SetDestinationAuth sets DestinationAuth field to given value.

### HasDestinationAuth

`func (o *AuthPolicies) HasDestinationAuth() bool`

HasDestinationAuth returns a boolean if a field has been set.

### GetNamespaceAuth

`func (o *AuthPolicies) GetNamespaceAuth() map[string][]string`

GetNamespaceAuth returns the NamespaceAuth field if non-nil, zero value otherwise.

### GetNamespaceAuthOk

`func (o *AuthPolicies) GetNamespaceAuthOk() (*map[string][]string, bool)`

GetNamespaceAuthOk returns a tuple with the NamespaceAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceAuth

`func (o *AuthPolicies) SetNamespaceAuth(v map[string][]string)`

SetNamespaceAuth sets NamespaceAuth field to given value.

### HasNamespaceAuth

`func (o *AuthPolicies) HasNamespaceAuth() bool`

HasNamespaceAuth returns a boolean if a field has been set.

### GetSubscriptionAuthRoles

`func (o *AuthPolicies) GetSubscriptionAuthRoles() map[string][]string`

GetSubscriptionAuthRoles returns the SubscriptionAuthRoles field if non-nil, zero value otherwise.

### GetSubscriptionAuthRolesOk

`func (o *AuthPolicies) GetSubscriptionAuthRolesOk() (*map[string][]string, bool)`

GetSubscriptionAuthRolesOk returns a tuple with the SubscriptionAuthRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAuthRoles

`func (o *AuthPolicies) SetSubscriptionAuthRoles(v map[string][]string)`

SetSubscriptionAuthRoles sets SubscriptionAuthRoles field to given value.

### HasSubscriptionAuthRoles

`func (o *AuthPolicies) HasSubscriptionAuthRoles() bool`

HasSubscriptionAuthRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


