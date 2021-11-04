# TimeZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Dstsavings** | Pointer to **int32** |  | [optional] 
**RawOffset** | Pointer to **int32** |  | [optional] 

## Methods

### NewTimeZone

`func NewTimeZone() *TimeZone`

NewTimeZone instantiates a new TimeZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeZoneWithDefaults

`func NewTimeZoneWithDefaults() *TimeZone`

NewTimeZoneWithDefaults instantiates a new TimeZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TimeZone) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TimeZone) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TimeZone) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TimeZone) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *TimeZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TimeZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDstsavings

`func (o *TimeZone) GetDstsavings() int32`

GetDstsavings returns the Dstsavings field if non-nil, zero value otherwise.

### GetDstsavingsOk

`func (o *TimeZone) GetDstsavingsOk() (*int32, bool)`

GetDstsavingsOk returns a tuple with the Dstsavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstsavings

`func (o *TimeZone) SetDstsavings(v int32)`

SetDstsavings sets Dstsavings field to given value.

### HasDstsavings

`func (o *TimeZone) HasDstsavings() bool`

HasDstsavings returns a boolean if a field has been set.

### GetRawOffset

`func (o *TimeZone) GetRawOffset() int32`

GetRawOffset returns the RawOffset field if non-nil, zero value otherwise.

### GetRawOffsetOk

`func (o *TimeZone) GetRawOffsetOk() (*int32, bool)`

GetRawOffsetOk returns a tuple with the RawOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawOffset

`func (o *TimeZone) SetRawOffset(v int32)`

SetRawOffset sets RawOffset field to given value.

### HasRawOffset

`func (o *TimeZone) HasRawOffset() bool`

HasRawOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


