# ResourceUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | Pointer to **float64** |  | [optional] 
**Limit** | Pointer to **float64** |  | [optional] 

## Methods

### NewResourceUsage

`func NewResourceUsage() *ResourceUsage`

NewResourceUsage instantiates a new ResourceUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsageWithDefaults

`func NewResourceUsageWithDefaults() *ResourceUsage`

NewResourceUsageWithDefaults instantiates a new ResourceUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *ResourceUsage) GetUsage() float64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ResourceUsage) GetUsageOk() (*float64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ResourceUsage) SetUsage(v float64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ResourceUsage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLimit

`func (o *ResourceUsage) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResourceUsage) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResourceUsage) SetLimit(v float64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResourceUsage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


