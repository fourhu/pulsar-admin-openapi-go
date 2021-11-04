# TypeBindings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Empty** | Pointer to **bool** |  | [optional] 
**TypeParameters** | Pointer to [**[]JavaType**](JavaType.md) |  | [optional] 

## Methods

### NewTypeBindings

`func NewTypeBindings() *TypeBindings`

NewTypeBindings instantiates a new TypeBindings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeBindingsWithDefaults

`func NewTypeBindingsWithDefaults() *TypeBindings`

NewTypeBindingsWithDefaults instantiates a new TypeBindings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmpty

`func (o *TypeBindings) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *TypeBindings) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *TypeBindings) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *TypeBindings) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetTypeParameters

`func (o *TypeBindings) GetTypeParameters() []JavaType`

GetTypeParameters returns the TypeParameters field if non-nil, zero value otherwise.

### GetTypeParametersOk

`func (o *TypeBindings) GetTypeParametersOk() (*[]JavaType, bool)`

GetTypeParametersOk returns a tuple with the TypeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeParameters

`func (o *TypeBindings) SetTypeParameters(v []JavaType)`

SetTypeParameters sets TypeParameters field to given value.

### HasTypeParameters

`func (o *TypeBindings) HasTypeParameters() bool`

HasTypeParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


