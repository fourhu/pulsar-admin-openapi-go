# Module

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layer** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Descriptor** | Pointer to [**ModuleDescriptor**](ModuleDescriptor.md) |  | [optional] 
**ClassLoader** | Pointer to [**ClassLoader**](ClassLoader.md) |  | [optional] 
**Annotations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DeclaredAnnotations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Named** | Pointer to **bool** |  | [optional] 
**Packages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModule

`func NewModule() *Module`

NewModule instantiates a new Module object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleWithDefaults

`func NewModuleWithDefaults() *Module`

NewModuleWithDefaults instantiates a new Module object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayer

`func (o *Module) GetLayer() map[string]interface{}`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *Module) GetLayerOk() (*map[string]interface{}, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *Module) SetLayer(v map[string]interface{})`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *Module) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetName

`func (o *Module) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Module) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Module) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Module) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescriptor

`func (o *Module) GetDescriptor() ModuleDescriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *Module) GetDescriptorOk() (*ModuleDescriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *Module) SetDescriptor(v ModuleDescriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *Module) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetClassLoader

`func (o *Module) GetClassLoader() ClassLoader`

GetClassLoader returns the ClassLoader field if non-nil, zero value otherwise.

### GetClassLoaderOk

`func (o *Module) GetClassLoaderOk() (*ClassLoader, bool)`

GetClassLoaderOk returns a tuple with the ClassLoader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassLoader

`func (o *Module) SetClassLoader(v ClassLoader)`

SetClassLoader sets ClassLoader field to given value.

### HasClassLoader

`func (o *Module) HasClassLoader() bool`

HasClassLoader returns a boolean if a field has been set.

### GetAnnotations

`func (o *Module) GetAnnotations() []map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Module) GetAnnotationsOk() (*[]map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Module) SetAnnotations(v []map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Module) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetDeclaredAnnotations

`func (o *Module) GetDeclaredAnnotations() []map[string]interface{}`

GetDeclaredAnnotations returns the DeclaredAnnotations field if non-nil, zero value otherwise.

### GetDeclaredAnnotationsOk

`func (o *Module) GetDeclaredAnnotationsOk() (*[]map[string]interface{}, bool)`

GetDeclaredAnnotationsOk returns a tuple with the DeclaredAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredAnnotations

`func (o *Module) SetDeclaredAnnotations(v []map[string]interface{})`

SetDeclaredAnnotations sets DeclaredAnnotations field to given value.

### HasDeclaredAnnotations

`func (o *Module) HasDeclaredAnnotations() bool`

HasDeclaredAnnotations returns a boolean if a field has been set.

### GetNamed

`func (o *Module) GetNamed() bool`

GetNamed returns the Named field if non-nil, zero value otherwise.

### GetNamedOk

`func (o *Module) GetNamedOk() (*bool, bool)`

GetNamedOk returns a tuple with the Named field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamed

`func (o *Module) SetNamed(v bool)`

SetNamed sets Named field to given value.

### HasNamed

`func (o *Module) HasNamed() bool`

HasNamed returns a boolean if a field has been set.

### GetPackages

`func (o *Module) GetPackages() []string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *Module) GetPackagesOk() (*[]string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *Module) SetPackages(v []string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *Module) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


