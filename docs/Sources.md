# Sources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceList** | Pointer to [**[]ConnectorDefinition**](ConnectorDefinition.md) |  | [optional] 
**ListOfConnectors** | Pointer to [**[]ConnectorDefinition**](ConnectorDefinition.md) |  | [optional] 

## Methods

### NewSources

`func NewSources() *Sources`

NewSources instantiates a new Sources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesWithDefaults

`func NewSourcesWithDefaults() *Sources`

NewSourcesWithDefaults instantiates a new Sources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceList

`func (o *Sources) GetSourceList() []ConnectorDefinition`

GetSourceList returns the SourceList field if non-nil, zero value otherwise.

### GetSourceListOk

`func (o *Sources) GetSourceListOk() (*[]ConnectorDefinition, bool)`

GetSourceListOk returns a tuple with the SourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceList

`func (o *Sources) SetSourceList(v []ConnectorDefinition)`

SetSourceList sets SourceList field to given value.

### HasSourceList

`func (o *Sources) HasSourceList() bool`

HasSourceList returns a boolean if a field has been set.

### GetListOfConnectors

`func (o *Sources) GetListOfConnectors() []ConnectorDefinition`

GetListOfConnectors returns the ListOfConnectors field if non-nil, zero value otherwise.

### GetListOfConnectorsOk

`func (o *Sources) GetListOfConnectorsOk() (*[]ConnectorDefinition, bool)`

GetListOfConnectorsOk returns a tuple with the ListOfConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfConnectors

`func (o *Sources) SetListOfConnectors(v []ConnectorDefinition)`

SetListOfConnectors sets ListOfConnectors field to given value.

### HasListOfConnectors

`func (o *Sources) HasListOfConnectors() bool`

HasListOfConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


