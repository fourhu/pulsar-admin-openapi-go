# InactiveTopicPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteWhileInactive** | Pointer to **bool** |  | [optional] 
**InactiveTopicDeleteMode** | Pointer to **string** |  | [optional] 
**MaxInactiveDurationSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewInactiveTopicPolicies

`func NewInactiveTopicPolicies() *InactiveTopicPolicies`

NewInactiveTopicPolicies instantiates a new InactiveTopicPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInactiveTopicPoliciesWithDefaults

`func NewInactiveTopicPoliciesWithDefaults() *InactiveTopicPolicies`

NewInactiveTopicPoliciesWithDefaults instantiates a new InactiveTopicPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteWhileInactive

`func (o *InactiveTopicPolicies) GetDeleteWhileInactive() bool`

GetDeleteWhileInactive returns the DeleteWhileInactive field if non-nil, zero value otherwise.

### GetDeleteWhileInactiveOk

`func (o *InactiveTopicPolicies) GetDeleteWhileInactiveOk() (*bool, bool)`

GetDeleteWhileInactiveOk returns a tuple with the DeleteWhileInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteWhileInactive

`func (o *InactiveTopicPolicies) SetDeleteWhileInactive(v bool)`

SetDeleteWhileInactive sets DeleteWhileInactive field to given value.

### HasDeleteWhileInactive

`func (o *InactiveTopicPolicies) HasDeleteWhileInactive() bool`

HasDeleteWhileInactive returns a boolean if a field has been set.

### GetInactiveTopicDeleteMode

`func (o *InactiveTopicPolicies) GetInactiveTopicDeleteMode() string`

GetInactiveTopicDeleteMode returns the InactiveTopicDeleteMode field if non-nil, zero value otherwise.

### GetInactiveTopicDeleteModeOk

`func (o *InactiveTopicPolicies) GetInactiveTopicDeleteModeOk() (*string, bool)`

GetInactiveTopicDeleteModeOk returns a tuple with the InactiveTopicDeleteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveTopicDeleteMode

`func (o *InactiveTopicPolicies) SetInactiveTopicDeleteMode(v string)`

SetInactiveTopicDeleteMode sets InactiveTopicDeleteMode field to given value.

### HasInactiveTopicDeleteMode

`func (o *InactiveTopicPolicies) HasInactiveTopicDeleteMode() bool`

HasInactiveTopicDeleteMode returns a boolean if a field has been set.

### GetMaxInactiveDurationSeconds

`func (o *InactiveTopicPolicies) GetMaxInactiveDurationSeconds() int32`

GetMaxInactiveDurationSeconds returns the MaxInactiveDurationSeconds field if non-nil, zero value otherwise.

### GetMaxInactiveDurationSecondsOk

`func (o *InactiveTopicPolicies) GetMaxInactiveDurationSecondsOk() (*int32, bool)`

GetMaxInactiveDurationSecondsOk returns a tuple with the MaxInactiveDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInactiveDurationSeconds

`func (o *InactiveTopicPolicies) SetMaxInactiveDurationSeconds(v int32)`

SetMaxInactiveDurationSeconds sets MaxInactiveDurationSeconds field to given value.

### HasMaxInactiveDurationSeconds

`func (o *InactiveTopicPolicies) HasMaxInactiveDurationSeconds() bool`

HasMaxInactiveDurationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


