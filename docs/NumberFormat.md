# NumberFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupingUsed** | Pointer to **bool** |  | [optional] 
**ParseIntegerOnly** | Pointer to **bool** |  | [optional] 
**MaximumIntegerDigits** | Pointer to **int32** |  | [optional] 
**MinimumIntegerDigits** | Pointer to **int32** |  | [optional] 
**MaximumFractionDigits** | Pointer to **int32** |  | [optional] 
**MinimumFractionDigits** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**RoundingMode** | Pointer to **string** |  | [optional] 

## Methods

### NewNumberFormat

`func NewNumberFormat() *NumberFormat`

NewNumberFormat instantiates a new NumberFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberFormatWithDefaults

`func NewNumberFormatWithDefaults() *NumberFormat`

NewNumberFormatWithDefaults instantiates a new NumberFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupingUsed

`func (o *NumberFormat) GetGroupingUsed() bool`

GetGroupingUsed returns the GroupingUsed field if non-nil, zero value otherwise.

### GetGroupingUsedOk

`func (o *NumberFormat) GetGroupingUsedOk() (*bool, bool)`

GetGroupingUsedOk returns a tuple with the GroupingUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingUsed

`func (o *NumberFormat) SetGroupingUsed(v bool)`

SetGroupingUsed sets GroupingUsed field to given value.

### HasGroupingUsed

`func (o *NumberFormat) HasGroupingUsed() bool`

HasGroupingUsed returns a boolean if a field has been set.

### GetParseIntegerOnly

`func (o *NumberFormat) GetParseIntegerOnly() bool`

GetParseIntegerOnly returns the ParseIntegerOnly field if non-nil, zero value otherwise.

### GetParseIntegerOnlyOk

`func (o *NumberFormat) GetParseIntegerOnlyOk() (*bool, bool)`

GetParseIntegerOnlyOk returns a tuple with the ParseIntegerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseIntegerOnly

`func (o *NumberFormat) SetParseIntegerOnly(v bool)`

SetParseIntegerOnly sets ParseIntegerOnly field to given value.

### HasParseIntegerOnly

`func (o *NumberFormat) HasParseIntegerOnly() bool`

HasParseIntegerOnly returns a boolean if a field has been set.

### GetMaximumIntegerDigits

`func (o *NumberFormat) GetMaximumIntegerDigits() int32`

GetMaximumIntegerDigits returns the MaximumIntegerDigits field if non-nil, zero value otherwise.

### GetMaximumIntegerDigitsOk

`func (o *NumberFormat) GetMaximumIntegerDigitsOk() (*int32, bool)`

GetMaximumIntegerDigitsOk returns a tuple with the MaximumIntegerDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumIntegerDigits

`func (o *NumberFormat) SetMaximumIntegerDigits(v int32)`

SetMaximumIntegerDigits sets MaximumIntegerDigits field to given value.

### HasMaximumIntegerDigits

`func (o *NumberFormat) HasMaximumIntegerDigits() bool`

HasMaximumIntegerDigits returns a boolean if a field has been set.

### GetMinimumIntegerDigits

`func (o *NumberFormat) GetMinimumIntegerDigits() int32`

GetMinimumIntegerDigits returns the MinimumIntegerDigits field if non-nil, zero value otherwise.

### GetMinimumIntegerDigitsOk

`func (o *NumberFormat) GetMinimumIntegerDigitsOk() (*int32, bool)`

GetMinimumIntegerDigitsOk returns a tuple with the MinimumIntegerDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumIntegerDigits

`func (o *NumberFormat) SetMinimumIntegerDigits(v int32)`

SetMinimumIntegerDigits sets MinimumIntegerDigits field to given value.

### HasMinimumIntegerDigits

`func (o *NumberFormat) HasMinimumIntegerDigits() bool`

HasMinimumIntegerDigits returns a boolean if a field has been set.

### GetMaximumFractionDigits

`func (o *NumberFormat) GetMaximumFractionDigits() int32`

GetMaximumFractionDigits returns the MaximumFractionDigits field if non-nil, zero value otherwise.

### GetMaximumFractionDigitsOk

`func (o *NumberFormat) GetMaximumFractionDigitsOk() (*int32, bool)`

GetMaximumFractionDigitsOk returns a tuple with the MaximumFractionDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFractionDigits

`func (o *NumberFormat) SetMaximumFractionDigits(v int32)`

SetMaximumFractionDigits sets MaximumFractionDigits field to given value.

### HasMaximumFractionDigits

`func (o *NumberFormat) HasMaximumFractionDigits() bool`

HasMaximumFractionDigits returns a boolean if a field has been set.

### GetMinimumFractionDigits

`func (o *NumberFormat) GetMinimumFractionDigits() int32`

GetMinimumFractionDigits returns the MinimumFractionDigits field if non-nil, zero value otherwise.

### GetMinimumFractionDigitsOk

`func (o *NumberFormat) GetMinimumFractionDigitsOk() (*int32, bool)`

GetMinimumFractionDigitsOk returns a tuple with the MinimumFractionDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFractionDigits

`func (o *NumberFormat) SetMinimumFractionDigits(v int32)`

SetMinimumFractionDigits sets MinimumFractionDigits field to given value.

### HasMinimumFractionDigits

`func (o *NumberFormat) HasMinimumFractionDigits() bool`

HasMinimumFractionDigits returns a boolean if a field has been set.

### GetCurrency

`func (o *NumberFormat) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *NumberFormat) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *NumberFormat) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *NumberFormat) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRoundingMode

`func (o *NumberFormat) GetRoundingMode() string`

GetRoundingMode returns the RoundingMode field if non-nil, zero value otherwise.

### GetRoundingModeOk

`func (o *NumberFormat) GetRoundingModeOk() (*string, bool)`

GetRoundingModeOk returns a tuple with the RoundingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingMode

`func (o *NumberFormat) SetRoundingMode(v string)`

SetRoundingMode sets RoundingMode field to given value.

### HasRoundingMode

`func (o *NumberFormat) HasRoundingMode() bool`

HasRoundingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


