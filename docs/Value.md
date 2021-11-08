# Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentInclusion** | Pointer to **string** |  | [optional] 
**ValueInclusion** | Pointer to **string** |  | [optional] 

## Methods

### NewValue

`func NewValue() *Value`

NewValue instantiates a new Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueWithDefaults

`func NewValueWithDefaults() *Value`

NewValueWithDefaults instantiates a new Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentInclusion

`func (o *Value) GetContentInclusion() string`

GetContentInclusion returns the ContentInclusion field if non-nil, zero value otherwise.

### GetContentInclusionOk

`func (o *Value) GetContentInclusionOk() (*string, bool)`

GetContentInclusionOk returns a tuple with the ContentInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentInclusion

`func (o *Value) SetContentInclusion(v string)`

SetContentInclusion sets ContentInclusion field to given value.

### HasContentInclusion

`func (o *Value) HasContentInclusion() bool`

HasContentInclusion returns a boolean if a field has been set.

### GetValueInclusion

`func (o *Value) GetValueInclusion() string`

GetValueInclusion returns the ValueInclusion field if non-nil, zero value otherwise.

### GetValueInclusionOk

`func (o *Value) GetValueInclusionOk() (*string, bool)`

GetValueInclusionOk returns a tuple with the ValueInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInclusion

`func (o *Value) SetValueInclusion(v string)`

SetValueInclusion sets ValueInclusion field to given value.

### HasValueInclusion

`func (o *Value) HasValueInclusion() bool`

HasValueInclusion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


