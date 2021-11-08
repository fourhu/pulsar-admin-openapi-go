# IsCompatibilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaCompatibilityStrategy** | Pointer to **string** |  | [optional] 
**Compatibility** | Pointer to **bool** |  | [optional] 

## Methods

### NewIsCompatibilityResponse

`func NewIsCompatibilityResponse() *IsCompatibilityResponse`

NewIsCompatibilityResponse instantiates a new IsCompatibilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsCompatibilityResponseWithDefaults

`func NewIsCompatibilityResponseWithDefaults() *IsCompatibilityResponse`

NewIsCompatibilityResponseWithDefaults instantiates a new IsCompatibilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaCompatibilityStrategy

`func (o *IsCompatibilityResponse) GetSchemaCompatibilityStrategy() string`

GetSchemaCompatibilityStrategy returns the SchemaCompatibilityStrategy field if non-nil, zero value otherwise.

### GetSchemaCompatibilityStrategyOk

`func (o *IsCompatibilityResponse) GetSchemaCompatibilityStrategyOk() (*string, bool)`

GetSchemaCompatibilityStrategyOk returns a tuple with the SchemaCompatibilityStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaCompatibilityStrategy

`func (o *IsCompatibilityResponse) SetSchemaCompatibilityStrategy(v string)`

SetSchemaCompatibilityStrategy sets SchemaCompatibilityStrategy field to given value.

### HasSchemaCompatibilityStrategy

`func (o *IsCompatibilityResponse) HasSchemaCompatibilityStrategy() bool`

HasSchemaCompatibilityStrategy returns a boolean if a field has been set.

### GetCompatibility

`func (o *IsCompatibilityResponse) GetCompatibility() bool`

GetCompatibility returns the Compatibility field if non-nil, zero value otherwise.

### GetCompatibilityOk

`func (o *IsCompatibilityResponse) GetCompatibilityOk() (*bool, bool)`

GetCompatibilityOk returns a tuple with the Compatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibility

`func (o *IsCompatibilityResponse) SetCompatibility(v bool)`

SetCompatibility sets Compatibility field to given value.

### HasCompatibility

`func (o *IsCompatibilityResponse) HasCompatibility() bool`

HasCompatibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


