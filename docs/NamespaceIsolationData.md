# NamespaceIsolationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespaces** | Pointer to **[]string** | The list of namespaces to apply this namespace isolation data | [optional] 
**Primary** | Pointer to **[]string** | The list of secondary brokers for serving the list of namespaces in this isolation policy | [optional] 
**AutoFailoverPolicy** | Pointer to [**AutoFailoverPolicyData**](AutoFailoverPolicyData.md) |  | [optional] 

## Methods

### NewNamespaceIsolationData

`func NewNamespaceIsolationData() *NamespaceIsolationData`

NewNamespaceIsolationData instantiates a new NamespaceIsolationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceIsolationDataWithDefaults

`func NewNamespaceIsolationDataWithDefaults() *NamespaceIsolationData`

NewNamespaceIsolationDataWithDefaults instantiates a new NamespaceIsolationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaces

`func (o *NamespaceIsolationData) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *NamespaceIsolationData) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *NamespaceIsolationData) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *NamespaceIsolationData) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetPrimary

`func (o *NamespaceIsolationData) GetPrimary() []string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *NamespaceIsolationData) GetPrimaryOk() (*[]string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *NamespaceIsolationData) SetPrimary(v []string)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *NamespaceIsolationData) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetAutoFailoverPolicy

`func (o *NamespaceIsolationData) GetAutoFailoverPolicy() AutoFailoverPolicyData`

GetAutoFailoverPolicy returns the AutoFailoverPolicy field if non-nil, zero value otherwise.

### GetAutoFailoverPolicyOk

`func (o *NamespaceIsolationData) GetAutoFailoverPolicyOk() (*AutoFailoverPolicyData, bool)`

GetAutoFailoverPolicyOk returns a tuple with the AutoFailoverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFailoverPolicy

`func (o *NamespaceIsolationData) SetAutoFailoverPolicy(v AutoFailoverPolicyData)`

SetAutoFailoverPolicy sets AutoFailoverPolicy field to given value.

### HasAutoFailoverPolicy

`func (o *NamespaceIsolationData) HasAutoFailoverPolicy() bool`

HasAutoFailoverPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


