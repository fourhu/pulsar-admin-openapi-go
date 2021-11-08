# Sinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SinkList** | Pointer to [**[]ConnectorDefinition**](ConnectorDefinition.md) |  | [optional] 
**ListOfConnectors** | Pointer to [**[]ConnectorDefinition**](ConnectorDefinition.md) |  | [optional] 

## Methods

### NewSinks

`func NewSinks() *Sinks`

NewSinks instantiates a new Sinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinksWithDefaults

`func NewSinksWithDefaults() *Sinks`

NewSinksWithDefaults instantiates a new Sinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSinkList

`func (o *Sinks) GetSinkList() []ConnectorDefinition`

GetSinkList returns the SinkList field if non-nil, zero value otherwise.

### GetSinkListOk

`func (o *Sinks) GetSinkListOk() (*[]ConnectorDefinition, bool)`

GetSinkListOk returns a tuple with the SinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinkList

`func (o *Sinks) SetSinkList(v []ConnectorDefinition)`

SetSinkList sets SinkList field to given value.

### HasSinkList

`func (o *Sinks) HasSinkList() bool`

HasSinkList returns a boolean if a field has been set.

### GetListOfConnectors

`func (o *Sinks) GetListOfConnectors() []ConnectorDefinition`

GetListOfConnectors returns the ListOfConnectors field if non-nil, zero value otherwise.

### GetListOfConnectorsOk

`func (o *Sinks) GetListOfConnectorsOk() (*[]ConnectorDefinition, bool)`

GetListOfConnectorsOk returns a tuple with the ListOfConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfConnectors

`func (o *Sinks) SetListOfConnectors(v []ConnectorDefinition)`

SetListOfConnectors sets ListOfConnectors field to given value.

### HasListOfConnectors

`func (o *Sinks) HasListOfConnectors() bool`

HasListOfConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


