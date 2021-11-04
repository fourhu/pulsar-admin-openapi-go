# JsonLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByteOffset** | Pointer to **int64** |  | [optional] 
**SourceRef** | Pointer to **map[string]interface{}** |  | [optional] 
**LineNr** | Pointer to **int32** |  | [optional] 
**ColumnNr** | Pointer to **int32** |  | [optional] 
**CharOffset** | Pointer to **int64** |  | [optional] 

## Methods

### NewJsonLocation

`func NewJsonLocation() *JsonLocation`

NewJsonLocation instantiates a new JsonLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonLocationWithDefaults

`func NewJsonLocationWithDefaults() *JsonLocation`

NewJsonLocationWithDefaults instantiates a new JsonLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByteOffset

`func (o *JsonLocation) GetByteOffset() int64`

GetByteOffset returns the ByteOffset field if non-nil, zero value otherwise.

### GetByteOffsetOk

`func (o *JsonLocation) GetByteOffsetOk() (*int64, bool)`

GetByteOffsetOk returns a tuple with the ByteOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteOffset

`func (o *JsonLocation) SetByteOffset(v int64)`

SetByteOffset sets ByteOffset field to given value.

### HasByteOffset

`func (o *JsonLocation) HasByteOffset() bool`

HasByteOffset returns a boolean if a field has been set.

### GetSourceRef

`func (o *JsonLocation) GetSourceRef() map[string]interface{}`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *JsonLocation) GetSourceRefOk() (*map[string]interface{}, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *JsonLocation) SetSourceRef(v map[string]interface{})`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *JsonLocation) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetLineNr

`func (o *JsonLocation) GetLineNr() int32`

GetLineNr returns the LineNr field if non-nil, zero value otherwise.

### GetLineNrOk

`func (o *JsonLocation) GetLineNrOk() (*int32, bool)`

GetLineNrOk returns a tuple with the LineNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNr

`func (o *JsonLocation) SetLineNr(v int32)`

SetLineNr sets LineNr field to given value.

### HasLineNr

`func (o *JsonLocation) HasLineNr() bool`

HasLineNr returns a boolean if a field has been set.

### GetColumnNr

`func (o *JsonLocation) GetColumnNr() int32`

GetColumnNr returns the ColumnNr field if non-nil, zero value otherwise.

### GetColumnNrOk

`func (o *JsonLocation) GetColumnNrOk() (*int32, bool)`

GetColumnNrOk returns a tuple with the ColumnNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNr

`func (o *JsonLocation) SetColumnNr(v int32)`

SetColumnNr sets ColumnNr field to given value.

### HasColumnNr

`func (o *JsonLocation) HasColumnNr() bool`

HasColumnNr returns a boolean if a field has been set.

### GetCharOffset

`func (o *JsonLocation) GetCharOffset() int64`

GetCharOffset returns the CharOffset field if non-nil, zero value otherwise.

### GetCharOffsetOk

`func (o *JsonLocation) GetCharOffsetOk() (*int64, bool)`

GetCharOffsetOk returns a tuple with the CharOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharOffset

`func (o *JsonLocation) SetCharOffset(v int64)`

SetCharOffset sets CharOffset field to given value.

### HasCharOffset

`func (o *JsonLocation) HasCharOffset() bool`

HasCharOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


