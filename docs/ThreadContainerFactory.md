# ThreadContainerFactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadGroupName** | Pointer to **string** |  | [optional] 
**PulsarClientMemoryLimit** | Pointer to [**MemoryLimit**](MemoryLimit.md) |  | [optional] 

## Methods

### NewThreadContainerFactory

`func NewThreadContainerFactory() *ThreadContainerFactory`

NewThreadContainerFactory instantiates a new ThreadContainerFactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadContainerFactoryWithDefaults

`func NewThreadContainerFactoryWithDefaults() *ThreadContainerFactory`

NewThreadContainerFactoryWithDefaults instantiates a new ThreadContainerFactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadGroupName

`func (o *ThreadContainerFactory) GetThreadGroupName() string`

GetThreadGroupName returns the ThreadGroupName field if non-nil, zero value otherwise.

### GetThreadGroupNameOk

`func (o *ThreadContainerFactory) GetThreadGroupNameOk() (*string, bool)`

GetThreadGroupNameOk returns a tuple with the ThreadGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadGroupName

`func (o *ThreadContainerFactory) SetThreadGroupName(v string)`

SetThreadGroupName sets ThreadGroupName field to given value.

### HasThreadGroupName

`func (o *ThreadContainerFactory) HasThreadGroupName() bool`

HasThreadGroupName returns a boolean if a field has been set.

### GetPulsarClientMemoryLimit

`func (o *ThreadContainerFactory) GetPulsarClientMemoryLimit() MemoryLimit`

GetPulsarClientMemoryLimit returns the PulsarClientMemoryLimit field if non-nil, zero value otherwise.

### GetPulsarClientMemoryLimitOk

`func (o *ThreadContainerFactory) GetPulsarClientMemoryLimitOk() (*MemoryLimit, bool)`

GetPulsarClientMemoryLimitOk returns a tuple with the PulsarClientMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarClientMemoryLimit

`func (o *ThreadContainerFactory) SetPulsarClientMemoryLimit(v MemoryLimit)`

SetPulsarClientMemoryLimit sets PulsarClientMemoryLimit field to given value.

### HasPulsarClientMemoryLimit

`func (o *ThreadContainerFactory) HasPulsarClientMemoryLimit() bool`

HasPulsarClientMemoryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


