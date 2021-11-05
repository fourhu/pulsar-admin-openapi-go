# BacklogQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitSize** | Pointer to **int64** |  | [optional] 
**LimitTime** | Pointer to **int32** |  | [optional] 
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

### GetLimitSize

`func (o *BacklogQuota) GetLimitSize() int64`

GetLimitSize returns the LimitSize field if non-nil, zero value otherwise.

### GetLimitSizeOk

`func (o *BacklogQuota) GetLimitSizeOk() (*int64, bool)`

GetLimitSizeOk returns a tuple with the LimitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSize

`func (o *BacklogQuota) SetLimitSize(v int64)`

SetLimitSize sets LimitSize field to given value.

### HasLimitSize

`func (o *BacklogQuota) HasLimitSize() bool`

HasLimitSize returns a boolean if a field has been set.

### GetLimitTime

`func (o *BacklogQuota) GetLimitTime() int32`

GetLimitTime returns the LimitTime field if non-nil, zero value otherwise.

### GetLimitTimeOk

`func (o *BacklogQuota) GetLimitTimeOk() (*int32, bool)`

GetLimitTimeOk returns a tuple with the LimitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTime

`func (o *BacklogQuota) SetLimitTime(v int32)`

SetLimitTime sets LimitTime field to given value.

### HasLimitTime

`func (o *BacklogQuota) HasLimitTime() bool`

HasLimitTime returns a boolean if a field has been set.

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


