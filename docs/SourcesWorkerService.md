# SourcesWorkerService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceList** | Pointer to [**[]ConnectorDefinition**](ConnectorDefinition.md) |  | [optional] 
**ListOfConnectors** | Pointer to [**[]ConnectorDefinition**](ConnectorDefinition.md) |  | [optional] 

## Methods

### NewSourcesWorkerService

`func NewSourcesWorkerService() *SourcesWorkerService`

NewSourcesWorkerService instantiates a new SourcesWorkerService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesWorkerServiceWithDefaults

`func NewSourcesWorkerServiceWithDefaults() *SourcesWorkerService`

NewSourcesWorkerServiceWithDefaults instantiates a new SourcesWorkerService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceList

`func (o *SourcesWorkerService) GetSourceList() []ConnectorDefinition`

GetSourceList returns the SourceList field if non-nil, zero value otherwise.

### GetSourceListOk

`func (o *SourcesWorkerService) GetSourceListOk() (*[]ConnectorDefinition, bool)`

GetSourceListOk returns a tuple with the SourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceList

`func (o *SourcesWorkerService) SetSourceList(v []ConnectorDefinition)`

SetSourceList sets SourceList field to given value.

### HasSourceList

`func (o *SourcesWorkerService) HasSourceList() bool`

HasSourceList returns a boolean if a field has been set.

### GetListOfConnectors

`func (o *SourcesWorkerService) GetListOfConnectors() []ConnectorDefinition`

GetListOfConnectors returns the ListOfConnectors field if non-nil, zero value otherwise.

### GetListOfConnectorsOk

`func (o *SourcesWorkerService) GetListOfConnectorsOk() (*[]ConnectorDefinition, bool)`

GetListOfConnectorsOk returns a tuple with the ListOfConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfConnectors

`func (o *SourcesWorkerService) SetListOfConnectors(v []ConnectorDefinition)`

SetListOfConnectors sets ListOfConnectors field to given value.

### HasListOfConnectors

`func (o *SourcesWorkerService) HasListOfConnectors() bool`

HasListOfConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


