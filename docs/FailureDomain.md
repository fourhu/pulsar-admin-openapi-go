# FailureDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brokers** | Pointer to **[]string** | The collection of brokers in the same failure domain | [optional] 

## Methods

### NewFailureDomain

`func NewFailureDomain() *FailureDomain`

NewFailureDomain instantiates a new FailureDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureDomainWithDefaults

`func NewFailureDomainWithDefaults() *FailureDomain`

NewFailureDomainWithDefaults instantiates a new FailureDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokers

`func (o *FailureDomain) GetBrokers() []string`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *FailureDomain) GetBrokersOk() (*[]string, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *FailureDomain) SetBrokers(v []string)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *FailureDomain) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


