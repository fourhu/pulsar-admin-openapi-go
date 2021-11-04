# DateFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calendar** | Pointer to **time.Time** |  | [optional] 
**NumberFormat** | Pointer to [**NumberFormat**](NumberFormat.md) |  | [optional] 
**TimeZone** | Pointer to [**TimeZone**](TimeZone.md) |  | [optional] 
**Lenient** | Pointer to **bool** |  | [optional] 

## Methods

### NewDateFormat

`func NewDateFormat() *DateFormat`

NewDateFormat instantiates a new DateFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateFormatWithDefaults

`func NewDateFormatWithDefaults() *DateFormat`

NewDateFormatWithDefaults instantiates a new DateFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalendar

`func (o *DateFormat) GetCalendar() time.Time`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *DateFormat) GetCalendarOk() (*time.Time, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *DateFormat) SetCalendar(v time.Time)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *DateFormat) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### GetNumberFormat

`func (o *DateFormat) GetNumberFormat() NumberFormat`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *DateFormat) GetNumberFormatOk() (*NumberFormat, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFormat

`func (o *DateFormat) SetNumberFormat(v NumberFormat)`

SetNumberFormat sets NumberFormat field to given value.

### HasNumberFormat

`func (o *DateFormat) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### GetTimeZone

`func (o *DateFormat) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *DateFormat) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *DateFormat) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *DateFormat) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetLenient

`func (o *DateFormat) GetLenient() bool`

GetLenient returns the Lenient field if non-nil, zero value otherwise.

### GetLenientOk

`func (o *DateFormat) GetLenientOk() (*bool, bool)`

GetLenientOk returns a tuple with the Lenient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLenient

`func (o *DateFormat) SetLenient(v bool)`

SetLenient sets Lenient field to given value.

### HasLenient

`func (o *DateFormat) HasLenient() bool`

HasLenient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


