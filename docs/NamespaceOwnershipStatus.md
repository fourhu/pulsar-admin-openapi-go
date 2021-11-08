# NamespaceOwnershipStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokerAssignment** | Pointer to **string** |  | [optional] 
**IsControlled** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 

## Methods

### NewNamespaceOwnershipStatus

`func NewNamespaceOwnershipStatus() *NamespaceOwnershipStatus`

NewNamespaceOwnershipStatus instantiates a new NamespaceOwnershipStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceOwnershipStatusWithDefaults

`func NewNamespaceOwnershipStatusWithDefaults() *NamespaceOwnershipStatus`

NewNamespaceOwnershipStatusWithDefaults instantiates a new NamespaceOwnershipStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokerAssignment

`func (o *NamespaceOwnershipStatus) GetBrokerAssignment() string`

GetBrokerAssignment returns the BrokerAssignment field if non-nil, zero value otherwise.

### GetBrokerAssignmentOk

`func (o *NamespaceOwnershipStatus) GetBrokerAssignmentOk() (*string, bool)`

GetBrokerAssignmentOk returns a tuple with the BrokerAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAssignment

`func (o *NamespaceOwnershipStatus) SetBrokerAssignment(v string)`

SetBrokerAssignment sets BrokerAssignment field to given value.

### HasBrokerAssignment

`func (o *NamespaceOwnershipStatus) HasBrokerAssignment() bool`

HasBrokerAssignment returns a boolean if a field has been set.

### GetIsControlled

`func (o *NamespaceOwnershipStatus) GetIsControlled() bool`

GetIsControlled returns the IsControlled field if non-nil, zero value otherwise.

### GetIsControlledOk

`func (o *NamespaceOwnershipStatus) GetIsControlledOk() (*bool, bool)`

GetIsControlledOk returns a tuple with the IsControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsControlled

`func (o *NamespaceOwnershipStatus) SetIsControlled(v bool)`

SetIsControlled sets IsControlled field to given value.

### HasIsControlled

`func (o *NamespaceOwnershipStatus) HasIsControlled() bool`

HasIsControlled returns a boolean if a field has been set.

### GetIsActive

`func (o *NamespaceOwnershipStatus) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *NamespaceOwnershipStatus) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *NamespaceOwnershipStatus) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *NamespaceOwnershipStatus) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


