# GetSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGetSchemaResponse

`func NewGetSchemaResponse() *GetSchemaResponse`

NewGetSchemaResponse instantiates a new GetSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSchemaResponseWithDefaults

`func NewGetSchemaResponseWithDefaults() *GetSchemaResponse`

NewGetSchemaResponseWithDefaults instantiates a new GetSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GetSchemaResponse) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetSchemaResponse) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetSchemaResponse) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetSchemaResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *GetSchemaResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSchemaResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSchemaResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSchemaResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetSchemaResponse) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetSchemaResponse) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetSchemaResponse) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetSchemaResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetData

`func (o *GetSchemaResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSchemaResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSchemaResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *GetSchemaResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetProperties

`func (o *GetSchemaResponse) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GetSchemaResponse) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GetSchemaResponse) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GetSchemaResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


