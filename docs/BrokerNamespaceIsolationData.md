# BrokerNamespaceIsolationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokerName** | Pointer to **string** | The broker name | [optional] 
**IsPrimary** | Pointer to **bool** | Is Primary broker | [optional] 
**NamespaceRegex** | Pointer to **[]string** | The namespace-isolation policies attached to this broker | [optional] 
**PolicyName** | Pointer to **string** | Policy name | [optional] 

## Methods

### NewBrokerNamespaceIsolationData

`func NewBrokerNamespaceIsolationData() *BrokerNamespaceIsolationData`

NewBrokerNamespaceIsolationData instantiates a new BrokerNamespaceIsolationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerNamespaceIsolationDataWithDefaults

`func NewBrokerNamespaceIsolationDataWithDefaults() *BrokerNamespaceIsolationData`

NewBrokerNamespaceIsolationDataWithDefaults instantiates a new BrokerNamespaceIsolationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokerName

`func (o *BrokerNamespaceIsolationData) GetBrokerName() string`

GetBrokerName returns the BrokerName field if non-nil, zero value otherwise.

### GetBrokerNameOk

`func (o *BrokerNamespaceIsolationData) GetBrokerNameOk() (*string, bool)`

GetBrokerNameOk returns a tuple with the BrokerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerName

`func (o *BrokerNamespaceIsolationData) SetBrokerName(v string)`

SetBrokerName sets BrokerName field to given value.

### HasBrokerName

`func (o *BrokerNamespaceIsolationData) HasBrokerName() bool`

HasBrokerName returns a boolean if a field has been set.

### GetIsPrimary

`func (o *BrokerNamespaceIsolationData) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *BrokerNamespaceIsolationData) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *BrokerNamespaceIsolationData) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *BrokerNamespaceIsolationData) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetNamespaceRegex

`func (o *BrokerNamespaceIsolationData) GetNamespaceRegex() []string`

GetNamespaceRegex returns the NamespaceRegex field if non-nil, zero value otherwise.

### GetNamespaceRegexOk

`func (o *BrokerNamespaceIsolationData) GetNamespaceRegexOk() (*[]string, bool)`

GetNamespaceRegexOk returns a tuple with the NamespaceRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceRegex

`func (o *BrokerNamespaceIsolationData) SetNamespaceRegex(v []string)`

SetNamespaceRegex sets NamespaceRegex field to given value.

### HasNamespaceRegex

`func (o *BrokerNamespaceIsolationData) HasNamespaceRegex() bool`

HasNamespaceRegex returns a boolean if a field has been set.

### GetPolicyName

`func (o *BrokerNamespaceIsolationData) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *BrokerNamespaceIsolationData) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *BrokerNamespaceIsolationData) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *BrokerNamespaceIsolationData) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


