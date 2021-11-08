# Package

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SpecificationTitle** | Pointer to **string** |  | [optional] 
**SpecificationVersion** | Pointer to **string** |  | [optional] 
**SpecificationVendor** | Pointer to **string** |  | [optional] 
**ImplementationTitle** | Pointer to **string** |  | [optional] 
**ImplementationVersion** | Pointer to **string** |  | [optional] 
**ImplementationVendor** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DeclaredAnnotations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Sealed** | Pointer to **bool** |  | [optional] 

## Methods

### NewPackage

`func NewPackage() *Package`

NewPackage instantiates a new Package object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageWithDefaults

`func NewPackageWithDefaults() *Package`

NewPackageWithDefaults instantiates a new Package object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Package) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Package) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Package) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Package) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpecificationTitle

`func (o *Package) GetSpecificationTitle() string`

GetSpecificationTitle returns the SpecificationTitle field if non-nil, zero value otherwise.

### GetSpecificationTitleOk

`func (o *Package) GetSpecificationTitleOk() (*string, bool)`

GetSpecificationTitleOk returns a tuple with the SpecificationTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationTitle

`func (o *Package) SetSpecificationTitle(v string)`

SetSpecificationTitle sets SpecificationTitle field to given value.

### HasSpecificationTitle

`func (o *Package) HasSpecificationTitle() bool`

HasSpecificationTitle returns a boolean if a field has been set.

### GetSpecificationVersion

`func (o *Package) GetSpecificationVersion() string`

GetSpecificationVersion returns the SpecificationVersion field if non-nil, zero value otherwise.

### GetSpecificationVersionOk

`func (o *Package) GetSpecificationVersionOk() (*string, bool)`

GetSpecificationVersionOk returns a tuple with the SpecificationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationVersion

`func (o *Package) SetSpecificationVersion(v string)`

SetSpecificationVersion sets SpecificationVersion field to given value.

### HasSpecificationVersion

`func (o *Package) HasSpecificationVersion() bool`

HasSpecificationVersion returns a boolean if a field has been set.

### GetSpecificationVendor

`func (o *Package) GetSpecificationVendor() string`

GetSpecificationVendor returns the SpecificationVendor field if non-nil, zero value otherwise.

### GetSpecificationVendorOk

`func (o *Package) GetSpecificationVendorOk() (*string, bool)`

GetSpecificationVendorOk returns a tuple with the SpecificationVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationVendor

`func (o *Package) SetSpecificationVendor(v string)`

SetSpecificationVendor sets SpecificationVendor field to given value.

### HasSpecificationVendor

`func (o *Package) HasSpecificationVendor() bool`

HasSpecificationVendor returns a boolean if a field has been set.

### GetImplementationTitle

`func (o *Package) GetImplementationTitle() string`

GetImplementationTitle returns the ImplementationTitle field if non-nil, zero value otherwise.

### GetImplementationTitleOk

`func (o *Package) GetImplementationTitleOk() (*string, bool)`

GetImplementationTitleOk returns a tuple with the ImplementationTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationTitle

`func (o *Package) SetImplementationTitle(v string)`

SetImplementationTitle sets ImplementationTitle field to given value.

### HasImplementationTitle

`func (o *Package) HasImplementationTitle() bool`

HasImplementationTitle returns a boolean if a field has been set.

### GetImplementationVersion

`func (o *Package) GetImplementationVersion() string`

GetImplementationVersion returns the ImplementationVersion field if non-nil, zero value otherwise.

### GetImplementationVersionOk

`func (o *Package) GetImplementationVersionOk() (*string, bool)`

GetImplementationVersionOk returns a tuple with the ImplementationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationVersion

`func (o *Package) SetImplementationVersion(v string)`

SetImplementationVersion sets ImplementationVersion field to given value.

### HasImplementationVersion

`func (o *Package) HasImplementationVersion() bool`

HasImplementationVersion returns a boolean if a field has been set.

### GetImplementationVendor

`func (o *Package) GetImplementationVendor() string`

GetImplementationVendor returns the ImplementationVendor field if non-nil, zero value otherwise.

### GetImplementationVendorOk

`func (o *Package) GetImplementationVendorOk() (*string, bool)`

GetImplementationVendorOk returns a tuple with the ImplementationVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationVendor

`func (o *Package) SetImplementationVendor(v string)`

SetImplementationVendor sets ImplementationVendor field to given value.

### HasImplementationVendor

`func (o *Package) HasImplementationVendor() bool`

HasImplementationVendor returns a boolean if a field has been set.

### GetAnnotations

`func (o *Package) GetAnnotations() []map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Package) GetAnnotationsOk() (*[]map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Package) SetAnnotations(v []map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Package) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetDeclaredAnnotations

`func (o *Package) GetDeclaredAnnotations() []map[string]interface{}`

GetDeclaredAnnotations returns the DeclaredAnnotations field if non-nil, zero value otherwise.

### GetDeclaredAnnotationsOk

`func (o *Package) GetDeclaredAnnotationsOk() (*[]map[string]interface{}, bool)`

GetDeclaredAnnotationsOk returns a tuple with the DeclaredAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredAnnotations

`func (o *Package) SetDeclaredAnnotations(v []map[string]interface{})`

SetDeclaredAnnotations sets DeclaredAnnotations field to given value.

### HasDeclaredAnnotations

`func (o *Package) HasDeclaredAnnotations() bool`

HasDeclaredAnnotations returns a boolean if a field has been set.

### GetSealed

`func (o *Package) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *Package) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *Package) SetSealed(v bool)`

SetSealed sets Sealed field to given value.

### HasSealed

`func (o *Package) HasSealed() bool`

HasSealed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


