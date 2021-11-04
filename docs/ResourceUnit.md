# ResourceUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableResource** | Pointer to [**ResourceDescription**](ResourceDescription.md) |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceUnit

`func NewResourceUnit() *ResourceUnit`

NewResourceUnit instantiates a new ResourceUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUnitWithDefaults

`func NewResourceUnitWithDefaults() *ResourceUnit`

NewResourceUnitWithDefaults instantiates a new ResourceUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableResource

`func (o *ResourceUnit) GetAvailableResource() ResourceDescription`

GetAvailableResource returns the AvailableResource field if non-nil, zero value otherwise.

### GetAvailableResourceOk

`func (o *ResourceUnit) GetAvailableResourceOk() (*ResourceDescription, bool)`

GetAvailableResourceOk returns a tuple with the AvailableResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableResource

`func (o *ResourceUnit) SetAvailableResource(v ResourceDescription)`

SetAvailableResource sets AvailableResource field to given value.

### HasAvailableResource

`func (o *ResourceUnit) HasAvailableResource() bool`

HasAvailableResource returns a boolean if a field has been set.

### GetResourceId

`func (o *ResourceUnit) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceUnit) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceUnit) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ResourceUnit) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


