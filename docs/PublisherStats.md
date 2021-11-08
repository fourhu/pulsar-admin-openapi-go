# PublisherStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProducerName** | Pointer to **string** |  | [optional] 
**MsgThroughputIn** | Pointer to **float64** |  | [optional] 
**AccessMode** | Pointer to **string** |  | [optional] 
**AverageMsgSize** | Pointer to **float64** |  | [optional] 
**ConnectedSince** | Pointer to **string** |  | [optional] 
**MsgRateIn** | Pointer to **float64** |  | [optional] 
**ChunkedMessageRate** | Pointer to **float64** |  | [optional] 
**ClientVersion** | Pointer to **string** |  | [optional] 
**ProducerId** | Pointer to **int64** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 

## Methods

### NewPublisherStats

`func NewPublisherStats() *PublisherStats`

NewPublisherStats instantiates a new PublisherStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublisherStatsWithDefaults

`func NewPublisherStatsWithDefaults() *PublisherStats`

NewPublisherStatsWithDefaults instantiates a new PublisherStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducerName

`func (o *PublisherStats) GetProducerName() string`

GetProducerName returns the ProducerName field if non-nil, zero value otherwise.

### GetProducerNameOk

`func (o *PublisherStats) GetProducerNameOk() (*string, bool)`

GetProducerNameOk returns a tuple with the ProducerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerName

`func (o *PublisherStats) SetProducerName(v string)`

SetProducerName sets ProducerName field to given value.

### HasProducerName

`func (o *PublisherStats) HasProducerName() bool`

HasProducerName returns a boolean if a field has been set.

### GetMsgThroughputIn

`func (o *PublisherStats) GetMsgThroughputIn() float64`

GetMsgThroughputIn returns the MsgThroughputIn field if non-nil, zero value otherwise.

### GetMsgThroughputInOk

`func (o *PublisherStats) GetMsgThroughputInOk() (*float64, bool)`

GetMsgThroughputInOk returns a tuple with the MsgThroughputIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputIn

`func (o *PublisherStats) SetMsgThroughputIn(v float64)`

SetMsgThroughputIn sets MsgThroughputIn field to given value.

### HasMsgThroughputIn

`func (o *PublisherStats) HasMsgThroughputIn() bool`

HasMsgThroughputIn returns a boolean if a field has been set.

### GetAccessMode

`func (o *PublisherStats) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *PublisherStats) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *PublisherStats) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *PublisherStats) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetAverageMsgSize

`func (o *PublisherStats) GetAverageMsgSize() float64`

GetAverageMsgSize returns the AverageMsgSize field if non-nil, zero value otherwise.

### GetAverageMsgSizeOk

`func (o *PublisherStats) GetAverageMsgSizeOk() (*float64, bool)`

GetAverageMsgSizeOk returns a tuple with the AverageMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageMsgSize

`func (o *PublisherStats) SetAverageMsgSize(v float64)`

SetAverageMsgSize sets AverageMsgSize field to given value.

### HasAverageMsgSize

`func (o *PublisherStats) HasAverageMsgSize() bool`

HasAverageMsgSize returns a boolean if a field has been set.

### GetConnectedSince

`func (o *PublisherStats) GetConnectedSince() string`

GetConnectedSince returns the ConnectedSince field if non-nil, zero value otherwise.

### GetConnectedSinceOk

`func (o *PublisherStats) GetConnectedSinceOk() (*string, bool)`

GetConnectedSinceOk returns a tuple with the ConnectedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSince

`func (o *PublisherStats) SetConnectedSince(v string)`

SetConnectedSince sets ConnectedSince field to given value.

### HasConnectedSince

`func (o *PublisherStats) HasConnectedSince() bool`

HasConnectedSince returns a boolean if a field has been set.

### GetMsgRateIn

`func (o *PublisherStats) GetMsgRateIn() float64`

GetMsgRateIn returns the MsgRateIn field if non-nil, zero value otherwise.

### GetMsgRateInOk

`func (o *PublisherStats) GetMsgRateInOk() (*float64, bool)`

GetMsgRateInOk returns a tuple with the MsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateIn

`func (o *PublisherStats) SetMsgRateIn(v float64)`

SetMsgRateIn sets MsgRateIn field to given value.

### HasMsgRateIn

`func (o *PublisherStats) HasMsgRateIn() bool`

HasMsgRateIn returns a boolean if a field has been set.

### GetChunkedMessageRate

`func (o *PublisherStats) GetChunkedMessageRate() float64`

GetChunkedMessageRate returns the ChunkedMessageRate field if non-nil, zero value otherwise.

### GetChunkedMessageRateOk

`func (o *PublisherStats) GetChunkedMessageRateOk() (*float64, bool)`

GetChunkedMessageRateOk returns a tuple with the ChunkedMessageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkedMessageRate

`func (o *PublisherStats) SetChunkedMessageRate(v float64)`

SetChunkedMessageRate sets ChunkedMessageRate field to given value.

### HasChunkedMessageRate

`func (o *PublisherStats) HasChunkedMessageRate() bool`

HasChunkedMessageRate returns a boolean if a field has been set.

### GetClientVersion

`func (o *PublisherStats) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *PublisherStats) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *PublisherStats) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *PublisherStats) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### GetProducerId

`func (o *PublisherStats) GetProducerId() int64`

GetProducerId returns the ProducerId field if non-nil, zero value otherwise.

### GetProducerIdOk

`func (o *PublisherStats) GetProducerIdOk() (*int64, bool)`

GetProducerIdOk returns a tuple with the ProducerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerId

`func (o *PublisherStats) SetProducerId(v int64)`

SetProducerId sets ProducerId field to given value.

### HasProducerId

`func (o *PublisherStats) HasProducerId() bool`

HasProducerId returns a boolean if a field has been set.

### GetMetadata

`func (o *PublisherStats) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublisherStats) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublisherStats) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublisherStats) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAddress

`func (o *PublisherStats) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PublisherStats) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PublisherStats) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PublisherStats) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


