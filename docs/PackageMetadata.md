# PackageMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**ModificationTime** | Pointer to **int64** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPackageMetadata

`func NewPackageMetadata() *PackageMetadata`

NewPackageMetadata instantiates a new PackageMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageMetadataWithDefaults

`func NewPackageMetadataWithDefaults() *PackageMetadata`

NewPackageMetadataWithDefaults instantiates a new PackageMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PackageMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContact

`func (o *PackageMetadata) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PackageMetadata) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PackageMetadata) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PackageMetadata) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCreateTime

`func (o *PackageMetadata) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PackageMetadata) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PackageMetadata) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *PackageMetadata) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetModificationTime

`func (o *PackageMetadata) GetModificationTime() int64`

GetModificationTime returns the ModificationTime field if non-nil, zero value otherwise.

### GetModificationTimeOk

`func (o *PackageMetadata) GetModificationTimeOk() (*int64, bool)`

GetModificationTimeOk returns a tuple with the ModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationTime

`func (o *PackageMetadata) SetModificationTime(v int64)`

SetModificationTime sets ModificationTime field to given value.

### HasModificationTime

`func (o *PackageMetadata) HasModificationTime() bool`

HasModificationTime returns a boolean if a field has been set.

### GetProperties

`func (o *PackageMetadata) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PackageMetadata) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PackageMetadata) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PackageMetadata) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


