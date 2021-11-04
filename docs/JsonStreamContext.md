# JsonStreamContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentValue** | Pointer to **map[string]interface{}** |  | [optional] 
**CurrentName** | Pointer to **string** |  | [optional] 
**TypeDesc** | Pointer to **string** |  | [optional] 
**EntryCount** | Pointer to **int32** |  | [optional] 
**Parent** | Pointer to [**JsonStreamContext**](JsonStreamContext.md) |  | [optional] 
**CurrentIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewJsonStreamContext

`func NewJsonStreamContext() *JsonStreamContext`

NewJsonStreamContext instantiates a new JsonStreamContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonStreamContextWithDefaults

`func NewJsonStreamContextWithDefaults() *JsonStreamContext`

NewJsonStreamContextWithDefaults instantiates a new JsonStreamContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentValue

`func (o *JsonStreamContext) GetCurrentValue() map[string]interface{}`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *JsonStreamContext) GetCurrentValueOk() (*map[string]interface{}, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *JsonStreamContext) SetCurrentValue(v map[string]interface{})`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *JsonStreamContext) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetCurrentName

`func (o *JsonStreamContext) GetCurrentName() string`

GetCurrentName returns the CurrentName field if non-nil, zero value otherwise.

### GetCurrentNameOk

`func (o *JsonStreamContext) GetCurrentNameOk() (*string, bool)`

GetCurrentNameOk returns a tuple with the CurrentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentName

`func (o *JsonStreamContext) SetCurrentName(v string)`

SetCurrentName sets CurrentName field to given value.

### HasCurrentName

`func (o *JsonStreamContext) HasCurrentName() bool`

HasCurrentName returns a boolean if a field has been set.

### GetTypeDesc

`func (o *JsonStreamContext) GetTypeDesc() string`

GetTypeDesc returns the TypeDesc field if non-nil, zero value otherwise.

### GetTypeDescOk

`func (o *JsonStreamContext) GetTypeDescOk() (*string, bool)`

GetTypeDescOk returns a tuple with the TypeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDesc

`func (o *JsonStreamContext) SetTypeDesc(v string)`

SetTypeDesc sets TypeDesc field to given value.

### HasTypeDesc

`func (o *JsonStreamContext) HasTypeDesc() bool`

HasTypeDesc returns a boolean if a field has been set.

### GetEntryCount

`func (o *JsonStreamContext) GetEntryCount() int32`

GetEntryCount returns the EntryCount field if non-nil, zero value otherwise.

### GetEntryCountOk

`func (o *JsonStreamContext) GetEntryCountOk() (*int32, bool)`

GetEntryCountOk returns a tuple with the EntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCount

`func (o *JsonStreamContext) SetEntryCount(v int32)`

SetEntryCount sets EntryCount field to given value.

### HasEntryCount

`func (o *JsonStreamContext) HasEntryCount() bool`

HasEntryCount returns a boolean if a field has been set.

### GetParent

`func (o *JsonStreamContext) GetParent() JsonStreamContext`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *JsonStreamContext) GetParentOk() (*JsonStreamContext, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *JsonStreamContext) SetParent(v JsonStreamContext)`

SetParent sets Parent field to given value.

### HasParent

`func (o *JsonStreamContext) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetCurrentIndex

`func (o *JsonStreamContext) GetCurrentIndex() int32`

GetCurrentIndex returns the CurrentIndex field if non-nil, zero value otherwise.

### GetCurrentIndexOk

`func (o *JsonStreamContext) GetCurrentIndexOk() (*int32, bool)`

GetCurrentIndexOk returns a tuple with the CurrentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIndex

`func (o *JsonStreamContext) SetCurrentIndex(v int32)`

SetCurrentIndex sets CurrentIndex field to given value.

### HasCurrentIndex

`func (o *JsonStreamContext) HasCurrentIndex() bool`

HasCurrentIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


