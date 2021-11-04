# NonPersistentPublisherStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AverageMsgSize** | Pointer to **float64** |  | [optional] 
**ChunkedMessageRate** | Pointer to **float64** |  | [optional] 
**ClientVersion** | Pointer to **string** |  | [optional] 
**ConnectedSince** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MsgDropRate** | Pointer to **float64** |  | [optional] 
**MsgRateIn** | Pointer to **float64** |  | [optional] 
**MsgThroughputIn** | Pointer to **float64** |  | [optional] 
**ProducerId** | Pointer to **int64** |  | [optional] 
**ProducerName** | Pointer to **string** |  | [optional] 

## Methods

### NewNonPersistentPublisherStats

`func NewNonPersistentPublisherStats() *NonPersistentPublisherStats`

NewNonPersistentPublisherStats instantiates a new NonPersistentPublisherStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonPersistentPublisherStatsWithDefaults

`func NewNonPersistentPublisherStatsWithDefaults() *NonPersistentPublisherStats`

NewNonPersistentPublisherStatsWithDefaults instantiates a new NonPersistentPublisherStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NonPersistentPublisherStats) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NonPersistentPublisherStats) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NonPersistentPublisherStats) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NonPersistentPublisherStats) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAverageMsgSize

`func (o *NonPersistentPublisherStats) GetAverageMsgSize() float64`

GetAverageMsgSize returns the AverageMsgSize field if non-nil, zero value otherwise.

### GetAverageMsgSizeOk

`func (o *NonPersistentPublisherStats) GetAverageMsgSizeOk() (*float64, bool)`

GetAverageMsgSizeOk returns a tuple with the AverageMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageMsgSize

`func (o *NonPersistentPublisherStats) SetAverageMsgSize(v float64)`

SetAverageMsgSize sets AverageMsgSize field to given value.

### HasAverageMsgSize

`func (o *NonPersistentPublisherStats) HasAverageMsgSize() bool`

HasAverageMsgSize returns a boolean if a field has been set.

### GetChunkedMessageRate

`func (o *NonPersistentPublisherStats) GetChunkedMessageRate() float64`

GetChunkedMessageRate returns the ChunkedMessageRate field if non-nil, zero value otherwise.

### GetChunkedMessageRateOk

`func (o *NonPersistentPublisherStats) GetChunkedMessageRateOk() (*float64, bool)`

GetChunkedMessageRateOk returns a tuple with the ChunkedMessageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkedMessageRate

`func (o *NonPersistentPublisherStats) SetChunkedMessageRate(v float64)`

SetChunkedMessageRate sets ChunkedMessageRate field to given value.

### HasChunkedMessageRate

`func (o *NonPersistentPublisherStats) HasChunkedMessageRate() bool`

HasChunkedMessageRate returns a boolean if a field has been set.

### GetClientVersion

`func (o *NonPersistentPublisherStats) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *NonPersistentPublisherStats) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *NonPersistentPublisherStats) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *NonPersistentPublisherStats) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### GetConnectedSince

`func (o *NonPersistentPublisherStats) GetConnectedSince() string`

GetConnectedSince returns the ConnectedSince field if non-nil, zero value otherwise.

### GetConnectedSinceOk

`func (o *NonPersistentPublisherStats) GetConnectedSinceOk() (*string, bool)`

GetConnectedSinceOk returns a tuple with the ConnectedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSince

`func (o *NonPersistentPublisherStats) SetConnectedSince(v string)`

SetConnectedSince sets ConnectedSince field to given value.

### HasConnectedSince

`func (o *NonPersistentPublisherStats) HasConnectedSince() bool`

HasConnectedSince returns a boolean if a field has been set.

### GetMetadata

`func (o *NonPersistentPublisherStats) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NonPersistentPublisherStats) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NonPersistentPublisherStats) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NonPersistentPublisherStats) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMsgDropRate

`func (o *NonPersistentPublisherStats) GetMsgDropRate() float64`

GetMsgDropRate returns the MsgDropRate field if non-nil, zero value otherwise.

### GetMsgDropRateOk

`func (o *NonPersistentPublisherStats) GetMsgDropRateOk() (*float64, bool)`

GetMsgDropRateOk returns a tuple with the MsgDropRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgDropRate

`func (o *NonPersistentPublisherStats) SetMsgDropRate(v float64)`

SetMsgDropRate sets MsgDropRate field to given value.

### HasMsgDropRate

`func (o *NonPersistentPublisherStats) HasMsgDropRate() bool`

HasMsgDropRate returns a boolean if a field has been set.

### GetMsgRateIn

`func (o *NonPersistentPublisherStats) GetMsgRateIn() float64`

GetMsgRateIn returns the MsgRateIn field if non-nil, zero value otherwise.

### GetMsgRateInOk

`func (o *NonPersistentPublisherStats) GetMsgRateInOk() (*float64, bool)`

GetMsgRateInOk returns a tuple with the MsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateIn

`func (o *NonPersistentPublisherStats) SetMsgRateIn(v float64)`

SetMsgRateIn sets MsgRateIn field to given value.

### HasMsgRateIn

`func (o *NonPersistentPublisherStats) HasMsgRateIn() bool`

HasMsgRateIn returns a boolean if a field has been set.

### GetMsgThroughputIn

`func (o *NonPersistentPublisherStats) GetMsgThroughputIn() float64`

GetMsgThroughputIn returns the MsgThroughputIn field if non-nil, zero value otherwise.

### GetMsgThroughputInOk

`func (o *NonPersistentPublisherStats) GetMsgThroughputInOk() (*float64, bool)`

GetMsgThroughputInOk returns a tuple with the MsgThroughputIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputIn

`func (o *NonPersistentPublisherStats) SetMsgThroughputIn(v float64)`

SetMsgThroughputIn sets MsgThroughputIn field to given value.

### HasMsgThroughputIn

`func (o *NonPersistentPublisherStats) HasMsgThroughputIn() bool`

HasMsgThroughputIn returns a boolean if a field has been set.

### GetProducerId

`func (o *NonPersistentPublisherStats) GetProducerId() int64`

GetProducerId returns the ProducerId field if non-nil, zero value otherwise.

### GetProducerIdOk

`func (o *NonPersistentPublisherStats) GetProducerIdOk() (*int64, bool)`

GetProducerIdOk returns a tuple with the ProducerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerId

`func (o *NonPersistentPublisherStats) SetProducerId(v int64)`

SetProducerId sets ProducerId field to given value.

### HasProducerId

`func (o *NonPersistentPublisherStats) HasProducerId() bool`

HasProducerId returns a boolean if a field has been set.

### GetProducerName

`func (o *NonPersistentPublisherStats) GetProducerName() string`

GetProducerName returns the ProducerName field if non-nil, zero value otherwise.

### GetProducerNameOk

`func (o *NonPersistentPublisherStats) GetProducerNameOk() (*string, bool)`

GetProducerNameOk returns a tuple with the ProducerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerName

`func (o *NonPersistentPublisherStats) SetProducerName(v string)`

SetProducerName sets ProducerName field to given value.

### HasProducerName

`func (o *NonPersistentPublisherStats) HasProducerName() bool`

HasProducerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


