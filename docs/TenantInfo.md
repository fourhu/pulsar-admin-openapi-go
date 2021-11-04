# TenantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminRoles** | Pointer to **[]string** | Comma separated list of auth principal allowed to administrate the tenant. | [optional] 
**AllowedClusters** | Pointer to **[]string** | Comma separated allowed clusters. | [optional] 

## Methods

### NewTenantInfo

`func NewTenantInfo() *TenantInfo`

NewTenantInfo instantiates a new TenantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantInfoWithDefaults

`func NewTenantInfoWithDefaults() *TenantInfo`

NewTenantInfoWithDefaults instantiates a new TenantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminRoles

`func (o *TenantInfo) GetAdminRoles() []string`

GetAdminRoles returns the AdminRoles field if non-nil, zero value otherwise.

### GetAdminRolesOk

`func (o *TenantInfo) GetAdminRolesOk() (*[]string, bool)`

GetAdminRolesOk returns a tuple with the AdminRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminRoles

`func (o *TenantInfo) SetAdminRoles(v []string)`

SetAdminRoles sets AdminRoles field to given value.

### HasAdminRoles

`func (o *TenantInfo) HasAdminRoles() bool`

HasAdminRoles returns a boolean if a field has been set.

### GetAllowedClusters

`func (o *TenantInfo) GetAllowedClusters() []string`

GetAllowedClusters returns the AllowedClusters field if non-nil, zero value otherwise.

### GetAllowedClustersOk

`func (o *TenantInfo) GetAllowedClustersOk() (*[]string, bool)`

GetAllowedClustersOk returns a tuple with the AllowedClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClusters

`func (o *TenantInfo) SetAllowedClusters(v []string)`

SetAllowedClusters sets AllowedClusters field to given value.

### HasAllowedClusters

`func (o *TenantInfo) HasAllowedClusters() bool`

HasAllowedClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


