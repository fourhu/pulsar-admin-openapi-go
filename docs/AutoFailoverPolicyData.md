# AutoFailoverPolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyType** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAutoFailoverPolicyData

`func NewAutoFailoverPolicyData() *AutoFailoverPolicyData`

NewAutoFailoverPolicyData instantiates a new AutoFailoverPolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoFailoverPolicyDataWithDefaults

`func NewAutoFailoverPolicyDataWithDefaults() *AutoFailoverPolicyData`

NewAutoFailoverPolicyDataWithDefaults instantiates a new AutoFailoverPolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyType

`func (o *AutoFailoverPolicyData) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *AutoFailoverPolicyData) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *AutoFailoverPolicyData) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *AutoFailoverPolicyData) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetParameters

`func (o *AutoFailoverPolicyData) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AutoFailoverPolicyData) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AutoFailoverPolicyData) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AutoFailoverPolicyData) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


