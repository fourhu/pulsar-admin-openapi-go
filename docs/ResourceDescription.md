# ResourceDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsagePct** | Pointer to **int32** |  | [optional] 
**ResourceUsage** | Pointer to [**map[string]ResourceUsage**](ResourceUsage.md) |  | [optional] 

## Methods

### NewResourceDescription

`func NewResourceDescription() *ResourceDescription`

NewResourceDescription instantiates a new ResourceDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDescriptionWithDefaults

`func NewResourceDescriptionWithDefaults() *ResourceDescription`

NewResourceDescriptionWithDefaults instantiates a new ResourceDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsagePct

`func (o *ResourceDescription) GetUsagePct() int32`

GetUsagePct returns the UsagePct field if non-nil, zero value otherwise.

### GetUsagePctOk

`func (o *ResourceDescription) GetUsagePctOk() (*int32, bool)`

GetUsagePctOk returns a tuple with the UsagePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePct

`func (o *ResourceDescription) SetUsagePct(v int32)`

SetUsagePct sets UsagePct field to given value.

### HasUsagePct

`func (o *ResourceDescription) HasUsagePct() bool`

HasUsagePct returns a boolean if a field has been set.

### GetResourceUsage

`func (o *ResourceDescription) GetResourceUsage() map[string]ResourceUsage`

GetResourceUsage returns the ResourceUsage field if non-nil, zero value otherwise.

### GetResourceUsageOk

`func (o *ResourceDescription) GetResourceUsageOk() (*map[string]ResourceUsage, bool)`

GetResourceUsageOk returns a tuple with the ResourceUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUsage

`func (o *ResourceDescription) SetResourceUsage(v map[string]ResourceUsage)`

SetResourceUsage sets ResourceUsage field to given value.

### HasResourceUsage

`func (o *ResourceDescription) HasResourceUsage() bool`

HasResourceUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


