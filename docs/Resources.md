# Resources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **float64** |  | [optional] 
**Ram** | Pointer to **int64** |  | [optional] 
**Disk** | Pointer to **int64** |  | [optional] 

## Methods

### NewResources

`func NewResources() *Resources`

NewResources instantiates a new Resources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesWithDefaults

`func NewResourcesWithDefaults() *Resources`

NewResourcesWithDefaults instantiates a new Resources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *Resources) GetCpu() float64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Resources) GetCpuOk() (*float64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Resources) SetCpu(v float64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Resources) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *Resources) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *Resources) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *Resources) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *Resources) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetDisk

`func (o *Resources) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Resources) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Resources) SetDisk(v int64)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *Resources) HasDisk() bool`

HasDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


