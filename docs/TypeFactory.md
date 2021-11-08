# TypeFactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassLoader** | Pointer to [**ClassLoader**](ClassLoader.md) |  | [optional] 

## Methods

### NewTypeFactory

`func NewTypeFactory() *TypeFactory`

NewTypeFactory instantiates a new TypeFactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeFactoryWithDefaults

`func NewTypeFactoryWithDefaults() *TypeFactory`

NewTypeFactoryWithDefaults instantiates a new TypeFactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassLoader

`func (o *TypeFactory) GetClassLoader() ClassLoader`

GetClassLoader returns the ClassLoader field if non-nil, zero value otherwise.

### GetClassLoaderOk

`func (o *TypeFactory) GetClassLoaderOk() (*ClassLoader, bool)`

GetClassLoaderOk returns a tuple with the ClassLoader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassLoader

`func (o *TypeFactory) SetClassLoader(v ClassLoader)`

SetClassLoader sets ClassLoader field to given value.

### HasClassLoader

`func (o *TypeFactory) HasClassLoader() bool`

HasClassLoader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


