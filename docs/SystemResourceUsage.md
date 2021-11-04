# SystemResourceUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthIn** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**BandwidthOut** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**Cpu** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**DirectMemory** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**Memory** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 

## Methods

### NewSystemResourceUsage

`func NewSystemResourceUsage() *SystemResourceUsage`

NewSystemResourceUsage instantiates a new SystemResourceUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemResourceUsageWithDefaults

`func NewSystemResourceUsageWithDefaults() *SystemResourceUsage`

NewSystemResourceUsageWithDefaults instantiates a new SystemResourceUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthIn

`func (o *SystemResourceUsage) GetBandwidthIn() ResourceUsage`

GetBandwidthIn returns the BandwidthIn field if non-nil, zero value otherwise.

### GetBandwidthInOk

`func (o *SystemResourceUsage) GetBandwidthInOk() (*ResourceUsage, bool)`

GetBandwidthInOk returns a tuple with the BandwidthIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthIn

`func (o *SystemResourceUsage) SetBandwidthIn(v ResourceUsage)`

SetBandwidthIn sets BandwidthIn field to given value.

### HasBandwidthIn

`func (o *SystemResourceUsage) HasBandwidthIn() bool`

HasBandwidthIn returns a boolean if a field has been set.

### GetBandwidthOut

`func (o *SystemResourceUsage) GetBandwidthOut() ResourceUsage`

GetBandwidthOut returns the BandwidthOut field if non-nil, zero value otherwise.

### GetBandwidthOutOk

`func (o *SystemResourceUsage) GetBandwidthOutOk() (*ResourceUsage, bool)`

GetBandwidthOutOk returns a tuple with the BandwidthOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthOut

`func (o *SystemResourceUsage) SetBandwidthOut(v ResourceUsage)`

SetBandwidthOut sets BandwidthOut field to given value.

### HasBandwidthOut

`func (o *SystemResourceUsage) HasBandwidthOut() bool`

HasBandwidthOut returns a boolean if a field has been set.

### GetCpu

`func (o *SystemResourceUsage) GetCpu() ResourceUsage`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *SystemResourceUsage) GetCpuOk() (*ResourceUsage, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *SystemResourceUsage) SetCpu(v ResourceUsage)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *SystemResourceUsage) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDirectMemory

`func (o *SystemResourceUsage) GetDirectMemory() ResourceUsage`

GetDirectMemory returns the DirectMemory field if non-nil, zero value otherwise.

### GetDirectMemoryOk

`func (o *SystemResourceUsage) GetDirectMemoryOk() (*ResourceUsage, bool)`

GetDirectMemoryOk returns a tuple with the DirectMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMemory

`func (o *SystemResourceUsage) SetDirectMemory(v ResourceUsage)`

SetDirectMemory sets DirectMemory field to given value.

### HasDirectMemory

`func (o *SystemResourceUsage) HasDirectMemory() bool`

HasDirectMemory returns a boolean if a field has been set.

### GetMemory

`func (o *SystemResourceUsage) GetMemory() ResourceUsage`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *SystemResourceUsage) GetMemoryOk() (*ResourceUsage, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *SystemResourceUsage) SetMemory(v ResourceUsage)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *SystemResourceUsage) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


