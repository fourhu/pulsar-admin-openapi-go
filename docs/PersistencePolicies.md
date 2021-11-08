# PersistencePolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookkeeperEnsemble** | Pointer to **int32** |  | [optional] 
**BookkeeperWriteQuorum** | Pointer to **int32** |  | [optional] 
**BookkeeperAckQuorum** | Pointer to **int32** |  | [optional] 
**ManagedLedgerMaxMarkDeleteRate** | Pointer to **float64** |  | [optional] 

## Methods

### NewPersistencePolicies

`func NewPersistencePolicies() *PersistencePolicies`

NewPersistencePolicies instantiates a new PersistencePolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistencePoliciesWithDefaults

`func NewPersistencePoliciesWithDefaults() *PersistencePolicies`

NewPersistencePoliciesWithDefaults instantiates a new PersistencePolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookkeeperEnsemble

`func (o *PersistencePolicies) GetBookkeeperEnsemble() int32`

GetBookkeeperEnsemble returns the BookkeeperEnsemble field if non-nil, zero value otherwise.

### GetBookkeeperEnsembleOk

`func (o *PersistencePolicies) GetBookkeeperEnsembleOk() (*int32, bool)`

GetBookkeeperEnsembleOk returns a tuple with the BookkeeperEnsemble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookkeeperEnsemble

`func (o *PersistencePolicies) SetBookkeeperEnsemble(v int32)`

SetBookkeeperEnsemble sets BookkeeperEnsemble field to given value.

### HasBookkeeperEnsemble

`func (o *PersistencePolicies) HasBookkeeperEnsemble() bool`

HasBookkeeperEnsemble returns a boolean if a field has been set.

### GetBookkeeperWriteQuorum

`func (o *PersistencePolicies) GetBookkeeperWriteQuorum() int32`

GetBookkeeperWriteQuorum returns the BookkeeperWriteQuorum field if non-nil, zero value otherwise.

### GetBookkeeperWriteQuorumOk

`func (o *PersistencePolicies) GetBookkeeperWriteQuorumOk() (*int32, bool)`

GetBookkeeperWriteQuorumOk returns a tuple with the BookkeeperWriteQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookkeeperWriteQuorum

`func (o *PersistencePolicies) SetBookkeeperWriteQuorum(v int32)`

SetBookkeeperWriteQuorum sets BookkeeperWriteQuorum field to given value.

### HasBookkeeperWriteQuorum

`func (o *PersistencePolicies) HasBookkeeperWriteQuorum() bool`

HasBookkeeperWriteQuorum returns a boolean if a field has been set.

### GetBookkeeperAckQuorum

`func (o *PersistencePolicies) GetBookkeeperAckQuorum() int32`

GetBookkeeperAckQuorum returns the BookkeeperAckQuorum field if non-nil, zero value otherwise.

### GetBookkeeperAckQuorumOk

`func (o *PersistencePolicies) GetBookkeeperAckQuorumOk() (*int32, bool)`

GetBookkeeperAckQuorumOk returns a tuple with the BookkeeperAckQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookkeeperAckQuorum

`func (o *PersistencePolicies) SetBookkeeperAckQuorum(v int32)`

SetBookkeeperAckQuorum sets BookkeeperAckQuorum field to given value.

### HasBookkeeperAckQuorum

`func (o *PersistencePolicies) HasBookkeeperAckQuorum() bool`

HasBookkeeperAckQuorum returns a boolean if a field has been set.

### GetManagedLedgerMaxMarkDeleteRate

`func (o *PersistencePolicies) GetManagedLedgerMaxMarkDeleteRate() float64`

GetManagedLedgerMaxMarkDeleteRate returns the ManagedLedgerMaxMarkDeleteRate field if non-nil, zero value otherwise.

### GetManagedLedgerMaxMarkDeleteRateOk

`func (o *PersistencePolicies) GetManagedLedgerMaxMarkDeleteRateOk() (*float64, bool)`

GetManagedLedgerMaxMarkDeleteRateOk returns a tuple with the ManagedLedgerMaxMarkDeleteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerMaxMarkDeleteRate

`func (o *PersistencePolicies) SetManagedLedgerMaxMarkDeleteRate(v float64)`

SetManagedLedgerMaxMarkDeleteRate sets ManagedLedgerMaxMarkDeleteRate field to given value.

### HasManagedLedgerMaxMarkDeleteRate

`func (o *PersistencePolicies) HasManagedLedgerMaxMarkDeleteRate() bool`

HasManagedLedgerMaxMarkDeleteRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


