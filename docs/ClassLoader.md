# ClassLoader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | Pointer to [**ClassLoader**](ClassLoader.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UnnamedModule** | Pointer to [**Module**](Module.md) |  | [optional] 
**RegisteredAsParallelCapable** | Pointer to **bool** |  | [optional] 
**DefinedPackages** | Pointer to [**[]Package**](Package.md) |  | [optional] 

## Methods

### NewClassLoader

`func NewClassLoader() *ClassLoader`

NewClassLoader instantiates a new ClassLoader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassLoaderWithDefaults

`func NewClassLoaderWithDefaults() *ClassLoader`

NewClassLoaderWithDefaults instantiates a new ClassLoader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *ClassLoader) GetParent() ClassLoader`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ClassLoader) GetParentOk() (*ClassLoader, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ClassLoader) SetParent(v ClassLoader)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ClassLoader) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetName

`func (o *ClassLoader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClassLoader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClassLoader) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClassLoader) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnnamedModule

`func (o *ClassLoader) GetUnnamedModule() Module`

GetUnnamedModule returns the UnnamedModule field if non-nil, zero value otherwise.

### GetUnnamedModuleOk

`func (o *ClassLoader) GetUnnamedModuleOk() (*Module, bool)`

GetUnnamedModuleOk returns a tuple with the UnnamedModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnnamedModule

`func (o *ClassLoader) SetUnnamedModule(v Module)`

SetUnnamedModule sets UnnamedModule field to given value.

### HasUnnamedModule

`func (o *ClassLoader) HasUnnamedModule() bool`

HasUnnamedModule returns a boolean if a field has been set.

### GetRegisteredAsParallelCapable

`func (o *ClassLoader) GetRegisteredAsParallelCapable() bool`

GetRegisteredAsParallelCapable returns the RegisteredAsParallelCapable field if non-nil, zero value otherwise.

### GetRegisteredAsParallelCapableOk

`func (o *ClassLoader) GetRegisteredAsParallelCapableOk() (*bool, bool)`

GetRegisteredAsParallelCapableOk returns a tuple with the RegisteredAsParallelCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAsParallelCapable

`func (o *ClassLoader) SetRegisteredAsParallelCapable(v bool)`

SetRegisteredAsParallelCapable sets RegisteredAsParallelCapable field to given value.

### HasRegisteredAsParallelCapable

`func (o *ClassLoader) HasRegisteredAsParallelCapable() bool`

HasRegisteredAsParallelCapable returns a boolean if a field has been set.

### GetDefinedPackages

`func (o *ClassLoader) GetDefinedPackages() []Package`

GetDefinedPackages returns the DefinedPackages field if non-nil, zero value otherwise.

### GetDefinedPackagesOk

`func (o *ClassLoader) GetDefinedPackagesOk() (*[]Package, bool)`

GetDefinedPackagesOk returns a tuple with the DefinedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedPackages

`func (o *ClassLoader) SetDefinedPackages(v []Package)`

SetDefinedPackages sets DefinedPackages field to given value.

### HasDefinedPackages

`func (o *ClassLoader) HasDefinedPackages() bool`

HasDefinedPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


