# WorkerService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionsV2** | Pointer to [**FunctionsV2WorkerService**](FunctionsV2WorkerService.md) |  | [optional] 
**WorkerConfig** | Pointer to [**WorkerConfig**](WorkerConfig.md) |  | [optional] 
**Sinks** | Pointer to [**SinksWorkerService**](SinksWorkerService.md) |  | [optional] 
**Functions** | Pointer to [**FunctionsWorkerService**](FunctionsWorkerService.md) |  | [optional] 
**Workers** | Pointer to **map[string]interface{}** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**Sources** | Pointer to [**SourcesWorkerService**](SourcesWorkerService.md) |  | [optional] 

## Methods

### NewWorkerService

`func NewWorkerService() *WorkerService`

NewWorkerService instantiates a new WorkerService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerServiceWithDefaults

`func NewWorkerServiceWithDefaults() *WorkerService`

NewWorkerServiceWithDefaults instantiates a new WorkerService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionsV2

`func (o *WorkerService) GetFunctionsV2() FunctionsV2WorkerService`

GetFunctionsV2 returns the FunctionsV2 field if non-nil, zero value otherwise.

### GetFunctionsV2Ok

`func (o *WorkerService) GetFunctionsV2Ok() (*FunctionsV2WorkerService, bool)`

GetFunctionsV2Ok returns a tuple with the FunctionsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsV2

`func (o *WorkerService) SetFunctionsV2(v FunctionsV2WorkerService)`

SetFunctionsV2 sets FunctionsV2 field to given value.

### HasFunctionsV2

`func (o *WorkerService) HasFunctionsV2() bool`

HasFunctionsV2 returns a boolean if a field has been set.

### GetWorkerConfig

`func (o *WorkerService) GetWorkerConfig() WorkerConfig`

GetWorkerConfig returns the WorkerConfig field if non-nil, zero value otherwise.

### GetWorkerConfigOk

`func (o *WorkerService) GetWorkerConfigOk() (*WorkerConfig, bool)`

GetWorkerConfigOk returns a tuple with the WorkerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerConfig

`func (o *WorkerService) SetWorkerConfig(v WorkerConfig)`

SetWorkerConfig sets WorkerConfig field to given value.

### HasWorkerConfig

`func (o *WorkerService) HasWorkerConfig() bool`

HasWorkerConfig returns a boolean if a field has been set.

### GetSinks

`func (o *WorkerService) GetSinks() SinksWorkerService`

GetSinks returns the Sinks field if non-nil, zero value otherwise.

### GetSinksOk

`func (o *WorkerService) GetSinksOk() (*SinksWorkerService, bool)`

GetSinksOk returns a tuple with the Sinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinks

`func (o *WorkerService) SetSinks(v SinksWorkerService)`

SetSinks sets Sinks field to given value.

### HasSinks

`func (o *WorkerService) HasSinks() bool`

HasSinks returns a boolean if a field has been set.

### GetFunctions

`func (o *WorkerService) GetFunctions() FunctionsWorkerService`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *WorkerService) GetFunctionsOk() (*FunctionsWorkerService, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *WorkerService) SetFunctions(v FunctionsWorkerService)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *WorkerService) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetWorkers

`func (o *WorkerService) GetWorkers() map[string]interface{}`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *WorkerService) GetWorkersOk() (*map[string]interface{}, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *WorkerService) SetWorkers(v map[string]interface{})`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *WorkerService) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetInitialized

`func (o *WorkerService) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *WorkerService) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *WorkerService) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *WorkerService) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetSources

`func (o *WorkerService) GetSources() SourcesWorkerService`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *WorkerService) GetSourcesOk() (*SourcesWorkerService, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *WorkerService) SetSources(v SourcesWorkerService)`

SetSources sets Sources field to given value.

### HasSources

`func (o *WorkerService) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


