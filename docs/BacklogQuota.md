# BacklogQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 

## Methods

### NewBacklogQuota

`func NewBacklogQuota() *BacklogQuota`

NewBacklogQuota instantiates a new BacklogQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBacklogQuotaWithDefaults

`func NewBacklogQuotaWithDefaults() *BacklogQuota`

NewBacklogQuotaWithDefaults instantiates a new BacklogQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *BacklogQuota) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BacklogQuota) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BacklogQuota) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BacklogQuota) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPolicy

`func (o *BacklogQuota) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BacklogQuota) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *BacklogQuota) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *BacklogQuota) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


