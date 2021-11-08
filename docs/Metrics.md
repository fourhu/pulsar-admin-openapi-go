# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Dimensions** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMetrics

`func NewMetrics() *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *Metrics) GetMetrics() map[string]map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Metrics) GetMetricsOk() (*map[string]map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Metrics) SetMetrics(v map[string]map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Metrics) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetDimensions

`func (o *Metrics) GetDimensions() map[string]string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *Metrics) GetDimensionsOk() (*map[string]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *Metrics) SetDimensions(v map[string]string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *Metrics) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


