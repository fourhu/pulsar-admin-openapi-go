# LoadReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**BrokerVersionString** | Pointer to **string** |  | [optional] 
**WebServiceUrl** | Pointer to **string** |  | [optional] 
**WebServiceUrlTls** | Pointer to **string** |  | [optional] 
**PulsarServiceUrl** | Pointer to **string** |  | [optional] 
**PulsarServiceUrlTls** | Pointer to **string** |  | [optional] 
**PersistentTopicsEnabled** | Pointer to **bool** |  | [optional] 
**NonPersistentTopicsEnabled** | Pointer to **bool** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**MsgRateIn** | Pointer to **float64** |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 
**NumTopics** | Pointer to **int32** |  | [optional] 
**NumConsumers** | Pointer to **int32** |  | [optional] 
**NumProducers** | Pointer to **int32** |  | [optional] 
**NumBundles** | Pointer to **int32** |  | [optional] 
**Protocols** | Pointer to **map[string]string** |  | [optional] 
**SystemResourceUsage** | Pointer to [**SystemResourceUsage**](SystemResourceUsage.md) |  | [optional] 
**BundleStats** | Pointer to [**map[string]NamespaceBundleStats**](NamespaceBundleStats.md) |  | [optional] 
**BundleGains** | Pointer to **[]string** |  | [optional] 
**BundleLosses** | Pointer to **[]string** |  | [optional] 
**AllocatedCPU** | Pointer to **float64** |  | [optional] 
**AllocatedMemory** | Pointer to **float64** |  | [optional] 
**AllocatedBandwidthIn** | Pointer to **float64** |  | [optional] 
**AllocatedBandwidthOut** | Pointer to **float64** |  | [optional] 
**AllocatedMsgRateIn** | Pointer to **float64** |  | [optional] 
**AllocatedMsgRateOut** | Pointer to **float64** |  | [optional] 
**PreAllocatedCPU** | Pointer to **float64** |  | [optional] 
**PreAllocatedMemory** | Pointer to **float64** |  | [optional] 
**PreAllocatedBandwidthIn** | Pointer to **float64** |  | [optional] 
**PreAllocatedBandwidthOut** | Pointer to **float64** |  | [optional] 
**PreAllocatedMsgRateIn** | Pointer to **float64** |  | [optional] 
**PreAllocatedMsgRateOut** | Pointer to **float64** |  | [optional] 
**UnderLoaded** | Pointer to **bool** |  | [optional] 
**OverLoaded** | Pointer to **bool** |  | [optional] 
**LoadReportType** | Pointer to **string** |  | [optional] 
**MsgThroughputIn** | Pointer to **float64** |  | [optional] 
**LastUpdate** | Pointer to **int64** |  | [optional] 
**DirectMemory** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**Cpu** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**Memory** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**BandwidthIn** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**BandwidthOut** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 

## Methods

### NewLoadReport

`func NewLoadReport() *LoadReport`

NewLoadReport instantiates a new LoadReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadReportWithDefaults

`func NewLoadReportWithDefaults() *LoadReport`

NewLoadReportWithDefaults instantiates a new LoadReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoadReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadReport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadReport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBrokerVersionString

`func (o *LoadReport) GetBrokerVersionString() string`

GetBrokerVersionString returns the BrokerVersionString field if non-nil, zero value otherwise.

### GetBrokerVersionStringOk

`func (o *LoadReport) GetBrokerVersionStringOk() (*string, bool)`

GetBrokerVersionStringOk returns a tuple with the BrokerVersionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerVersionString

`func (o *LoadReport) SetBrokerVersionString(v string)`

SetBrokerVersionString sets BrokerVersionString field to given value.

### HasBrokerVersionString

`func (o *LoadReport) HasBrokerVersionString() bool`

HasBrokerVersionString returns a boolean if a field has been set.

### GetWebServiceUrl

`func (o *LoadReport) GetWebServiceUrl() string`

GetWebServiceUrl returns the WebServiceUrl field if non-nil, zero value otherwise.

### GetWebServiceUrlOk

`func (o *LoadReport) GetWebServiceUrlOk() (*string, bool)`

GetWebServiceUrlOk returns a tuple with the WebServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceUrl

`func (o *LoadReport) SetWebServiceUrl(v string)`

SetWebServiceUrl sets WebServiceUrl field to given value.

### HasWebServiceUrl

`func (o *LoadReport) HasWebServiceUrl() bool`

HasWebServiceUrl returns a boolean if a field has been set.

### GetWebServiceUrlTls

`func (o *LoadReport) GetWebServiceUrlTls() string`

GetWebServiceUrlTls returns the WebServiceUrlTls field if non-nil, zero value otherwise.

### GetWebServiceUrlTlsOk

`func (o *LoadReport) GetWebServiceUrlTlsOk() (*string, bool)`

GetWebServiceUrlTlsOk returns a tuple with the WebServiceUrlTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceUrlTls

`func (o *LoadReport) SetWebServiceUrlTls(v string)`

SetWebServiceUrlTls sets WebServiceUrlTls field to given value.

### HasWebServiceUrlTls

`func (o *LoadReport) HasWebServiceUrlTls() bool`

HasWebServiceUrlTls returns a boolean if a field has been set.

### GetPulsarServiceUrl

`func (o *LoadReport) GetPulsarServiceUrl() string`

GetPulsarServiceUrl returns the PulsarServiceUrl field if non-nil, zero value otherwise.

### GetPulsarServiceUrlOk

`func (o *LoadReport) GetPulsarServiceUrlOk() (*string, bool)`

GetPulsarServiceUrlOk returns a tuple with the PulsarServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarServiceUrl

`func (o *LoadReport) SetPulsarServiceUrl(v string)`

SetPulsarServiceUrl sets PulsarServiceUrl field to given value.

### HasPulsarServiceUrl

`func (o *LoadReport) HasPulsarServiceUrl() bool`

HasPulsarServiceUrl returns a boolean if a field has been set.

### GetPulsarServiceUrlTls

`func (o *LoadReport) GetPulsarServiceUrlTls() string`

GetPulsarServiceUrlTls returns the PulsarServiceUrlTls field if non-nil, zero value otherwise.

### GetPulsarServiceUrlTlsOk

`func (o *LoadReport) GetPulsarServiceUrlTlsOk() (*string, bool)`

GetPulsarServiceUrlTlsOk returns a tuple with the PulsarServiceUrlTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarServiceUrlTls

`func (o *LoadReport) SetPulsarServiceUrlTls(v string)`

SetPulsarServiceUrlTls sets PulsarServiceUrlTls field to given value.

### HasPulsarServiceUrlTls

`func (o *LoadReport) HasPulsarServiceUrlTls() bool`

HasPulsarServiceUrlTls returns a boolean if a field has been set.

### GetPersistentTopicsEnabled

`func (o *LoadReport) GetPersistentTopicsEnabled() bool`

GetPersistentTopicsEnabled returns the PersistentTopicsEnabled field if non-nil, zero value otherwise.

### GetPersistentTopicsEnabledOk

`func (o *LoadReport) GetPersistentTopicsEnabledOk() (*bool, bool)`

GetPersistentTopicsEnabledOk returns a tuple with the PersistentTopicsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentTopicsEnabled

`func (o *LoadReport) SetPersistentTopicsEnabled(v bool)`

SetPersistentTopicsEnabled sets PersistentTopicsEnabled field to given value.

### HasPersistentTopicsEnabled

`func (o *LoadReport) HasPersistentTopicsEnabled() bool`

HasPersistentTopicsEnabled returns a boolean if a field has been set.

### GetNonPersistentTopicsEnabled

`func (o *LoadReport) GetNonPersistentTopicsEnabled() bool`

GetNonPersistentTopicsEnabled returns the NonPersistentTopicsEnabled field if non-nil, zero value otherwise.

### GetNonPersistentTopicsEnabledOk

`func (o *LoadReport) GetNonPersistentTopicsEnabledOk() (*bool, bool)`

GetNonPersistentTopicsEnabledOk returns a tuple with the NonPersistentTopicsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPersistentTopicsEnabled

`func (o *LoadReport) SetNonPersistentTopicsEnabled(v bool)`

SetNonPersistentTopicsEnabled sets NonPersistentTopicsEnabled field to given value.

### HasNonPersistentTopicsEnabled

`func (o *LoadReport) HasNonPersistentTopicsEnabled() bool`

HasNonPersistentTopicsEnabled returns a boolean if a field has been set.

### GetTimestamp

`func (o *LoadReport) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LoadReport) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LoadReport) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LoadReport) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMsgRateIn

`func (o *LoadReport) GetMsgRateIn() float64`

GetMsgRateIn returns the MsgRateIn field if non-nil, zero value otherwise.

### GetMsgRateInOk

`func (o *LoadReport) GetMsgRateInOk() (*float64, bool)`

GetMsgRateInOk returns a tuple with the MsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateIn

`func (o *LoadReport) SetMsgRateIn(v float64)`

SetMsgRateIn sets MsgRateIn field to given value.

### HasMsgRateIn

`func (o *LoadReport) HasMsgRateIn() bool`

HasMsgRateIn returns a boolean if a field has been set.

### GetMsgRateOut

`func (o *LoadReport) GetMsgRateOut() float64`

GetMsgRateOut returns the MsgRateOut field if non-nil, zero value otherwise.

### GetMsgRateOutOk

`func (o *LoadReport) GetMsgRateOutOk() (*float64, bool)`

GetMsgRateOutOk returns a tuple with the MsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateOut

`func (o *LoadReport) SetMsgRateOut(v float64)`

SetMsgRateOut sets MsgRateOut field to given value.

### HasMsgRateOut

`func (o *LoadReport) HasMsgRateOut() bool`

HasMsgRateOut returns a boolean if a field has been set.

### GetNumTopics

`func (o *LoadReport) GetNumTopics() int32`

GetNumTopics returns the NumTopics field if non-nil, zero value otherwise.

### GetNumTopicsOk

`func (o *LoadReport) GetNumTopicsOk() (*int32, bool)`

GetNumTopicsOk returns a tuple with the NumTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTopics

`func (o *LoadReport) SetNumTopics(v int32)`

SetNumTopics sets NumTopics field to given value.

### HasNumTopics

`func (o *LoadReport) HasNumTopics() bool`

HasNumTopics returns a boolean if a field has been set.

### GetNumConsumers

`func (o *LoadReport) GetNumConsumers() int32`

GetNumConsumers returns the NumConsumers field if non-nil, zero value otherwise.

### GetNumConsumersOk

`func (o *LoadReport) GetNumConsumersOk() (*int32, bool)`

GetNumConsumersOk returns a tuple with the NumConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumConsumers

`func (o *LoadReport) SetNumConsumers(v int32)`

SetNumConsumers sets NumConsumers field to given value.

### HasNumConsumers

`func (o *LoadReport) HasNumConsumers() bool`

HasNumConsumers returns a boolean if a field has been set.

### GetNumProducers

`func (o *LoadReport) GetNumProducers() int32`

GetNumProducers returns the NumProducers field if non-nil, zero value otherwise.

### GetNumProducersOk

`func (o *LoadReport) GetNumProducersOk() (*int32, bool)`

GetNumProducersOk returns a tuple with the NumProducers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProducers

`func (o *LoadReport) SetNumProducers(v int32)`

SetNumProducers sets NumProducers field to given value.

### HasNumProducers

`func (o *LoadReport) HasNumProducers() bool`

HasNumProducers returns a boolean if a field has been set.

### GetNumBundles

`func (o *LoadReport) GetNumBundles() int32`

GetNumBundles returns the NumBundles field if non-nil, zero value otherwise.

### GetNumBundlesOk

`func (o *LoadReport) GetNumBundlesOk() (*int32, bool)`

GetNumBundlesOk returns a tuple with the NumBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBundles

`func (o *LoadReport) SetNumBundles(v int32)`

SetNumBundles sets NumBundles field to given value.

### HasNumBundles

`func (o *LoadReport) HasNumBundles() bool`

HasNumBundles returns a boolean if a field has been set.

### GetProtocols

`func (o *LoadReport) GetProtocols() map[string]string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *LoadReport) GetProtocolsOk() (*map[string]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *LoadReport) SetProtocols(v map[string]string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *LoadReport) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetSystemResourceUsage

`func (o *LoadReport) GetSystemResourceUsage() SystemResourceUsage`

GetSystemResourceUsage returns the SystemResourceUsage field if non-nil, zero value otherwise.

### GetSystemResourceUsageOk

`func (o *LoadReport) GetSystemResourceUsageOk() (*SystemResourceUsage, bool)`

GetSystemResourceUsageOk returns a tuple with the SystemResourceUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemResourceUsage

`func (o *LoadReport) SetSystemResourceUsage(v SystemResourceUsage)`

SetSystemResourceUsage sets SystemResourceUsage field to given value.

### HasSystemResourceUsage

`func (o *LoadReport) HasSystemResourceUsage() bool`

HasSystemResourceUsage returns a boolean if a field has been set.

### GetBundleStats

`func (o *LoadReport) GetBundleStats() map[string]NamespaceBundleStats`

GetBundleStats returns the BundleStats field if non-nil, zero value otherwise.

### GetBundleStatsOk

`func (o *LoadReport) GetBundleStatsOk() (*map[string]NamespaceBundleStats, bool)`

GetBundleStatsOk returns a tuple with the BundleStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleStats

`func (o *LoadReport) SetBundleStats(v map[string]NamespaceBundleStats)`

SetBundleStats sets BundleStats field to given value.

### HasBundleStats

`func (o *LoadReport) HasBundleStats() bool`

HasBundleStats returns a boolean if a field has been set.

### GetBundleGains

`func (o *LoadReport) GetBundleGains() []string`

GetBundleGains returns the BundleGains field if non-nil, zero value otherwise.

### GetBundleGainsOk

`func (o *LoadReport) GetBundleGainsOk() (*[]string, bool)`

GetBundleGainsOk returns a tuple with the BundleGains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleGains

`func (o *LoadReport) SetBundleGains(v []string)`

SetBundleGains sets BundleGains field to given value.

### HasBundleGains

`func (o *LoadReport) HasBundleGains() bool`

HasBundleGains returns a boolean if a field has been set.

### GetBundleLosses

`func (o *LoadReport) GetBundleLosses() []string`

GetBundleLosses returns the BundleLosses field if non-nil, zero value otherwise.

### GetBundleLossesOk

`func (o *LoadReport) GetBundleLossesOk() (*[]string, bool)`

GetBundleLossesOk returns a tuple with the BundleLosses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleLosses

`func (o *LoadReport) SetBundleLosses(v []string)`

SetBundleLosses sets BundleLosses field to given value.

### HasBundleLosses

`func (o *LoadReport) HasBundleLosses() bool`

HasBundleLosses returns a boolean if a field has been set.

### GetAllocatedCPU

`func (o *LoadReport) GetAllocatedCPU() float64`

GetAllocatedCPU returns the AllocatedCPU field if non-nil, zero value otherwise.

### GetAllocatedCPUOk

`func (o *LoadReport) GetAllocatedCPUOk() (*float64, bool)`

GetAllocatedCPUOk returns a tuple with the AllocatedCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedCPU

`func (o *LoadReport) SetAllocatedCPU(v float64)`

SetAllocatedCPU sets AllocatedCPU field to given value.

### HasAllocatedCPU

`func (o *LoadReport) HasAllocatedCPU() bool`

HasAllocatedCPU returns a boolean if a field has been set.

### GetAllocatedMemory

`func (o *LoadReport) GetAllocatedMemory() float64`

GetAllocatedMemory returns the AllocatedMemory field if non-nil, zero value otherwise.

### GetAllocatedMemoryOk

`func (o *LoadReport) GetAllocatedMemoryOk() (*float64, bool)`

GetAllocatedMemoryOk returns a tuple with the AllocatedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedMemory

`func (o *LoadReport) SetAllocatedMemory(v float64)`

SetAllocatedMemory sets AllocatedMemory field to given value.

### HasAllocatedMemory

`func (o *LoadReport) HasAllocatedMemory() bool`

HasAllocatedMemory returns a boolean if a field has been set.

### GetAllocatedBandwidthIn

`func (o *LoadReport) GetAllocatedBandwidthIn() float64`

GetAllocatedBandwidthIn returns the AllocatedBandwidthIn field if non-nil, zero value otherwise.

### GetAllocatedBandwidthInOk

`func (o *LoadReport) GetAllocatedBandwidthInOk() (*float64, bool)`

GetAllocatedBandwidthInOk returns a tuple with the AllocatedBandwidthIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedBandwidthIn

`func (o *LoadReport) SetAllocatedBandwidthIn(v float64)`

SetAllocatedBandwidthIn sets AllocatedBandwidthIn field to given value.

### HasAllocatedBandwidthIn

`func (o *LoadReport) HasAllocatedBandwidthIn() bool`

HasAllocatedBandwidthIn returns a boolean if a field has been set.

### GetAllocatedBandwidthOut

`func (o *LoadReport) GetAllocatedBandwidthOut() float64`

GetAllocatedBandwidthOut returns the AllocatedBandwidthOut field if non-nil, zero value otherwise.

### GetAllocatedBandwidthOutOk

`func (o *LoadReport) GetAllocatedBandwidthOutOk() (*float64, bool)`

GetAllocatedBandwidthOutOk returns a tuple with the AllocatedBandwidthOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedBandwidthOut

`func (o *LoadReport) SetAllocatedBandwidthOut(v float64)`

SetAllocatedBandwidthOut sets AllocatedBandwidthOut field to given value.

### HasAllocatedBandwidthOut

`func (o *LoadReport) HasAllocatedBandwidthOut() bool`

HasAllocatedBandwidthOut returns a boolean if a field has been set.

### GetAllocatedMsgRateIn

`func (o *LoadReport) GetAllocatedMsgRateIn() float64`

GetAllocatedMsgRateIn returns the AllocatedMsgRateIn field if non-nil, zero value otherwise.

### GetAllocatedMsgRateInOk

`func (o *LoadReport) GetAllocatedMsgRateInOk() (*float64, bool)`

GetAllocatedMsgRateInOk returns a tuple with the AllocatedMsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedMsgRateIn

`func (o *LoadReport) SetAllocatedMsgRateIn(v float64)`

SetAllocatedMsgRateIn sets AllocatedMsgRateIn field to given value.

### HasAllocatedMsgRateIn

`func (o *LoadReport) HasAllocatedMsgRateIn() bool`

HasAllocatedMsgRateIn returns a boolean if a field has been set.

### GetAllocatedMsgRateOut

`func (o *LoadReport) GetAllocatedMsgRateOut() float64`

GetAllocatedMsgRateOut returns the AllocatedMsgRateOut field if non-nil, zero value otherwise.

### GetAllocatedMsgRateOutOk

`func (o *LoadReport) GetAllocatedMsgRateOutOk() (*float64, bool)`

GetAllocatedMsgRateOutOk returns a tuple with the AllocatedMsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedMsgRateOut

`func (o *LoadReport) SetAllocatedMsgRateOut(v float64)`

SetAllocatedMsgRateOut sets AllocatedMsgRateOut field to given value.

### HasAllocatedMsgRateOut

`func (o *LoadReport) HasAllocatedMsgRateOut() bool`

HasAllocatedMsgRateOut returns a boolean if a field has been set.

### GetPreAllocatedCPU

`func (o *LoadReport) GetPreAllocatedCPU() float64`

GetPreAllocatedCPU returns the PreAllocatedCPU field if non-nil, zero value otherwise.

### GetPreAllocatedCPUOk

`func (o *LoadReport) GetPreAllocatedCPUOk() (*float64, bool)`

GetPreAllocatedCPUOk returns a tuple with the PreAllocatedCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAllocatedCPU

`func (o *LoadReport) SetPreAllocatedCPU(v float64)`

SetPreAllocatedCPU sets PreAllocatedCPU field to given value.

### HasPreAllocatedCPU

`func (o *LoadReport) HasPreAllocatedCPU() bool`

HasPreAllocatedCPU returns a boolean if a field has been set.

### GetPreAllocatedMemory

`func (o *LoadReport) GetPreAllocatedMemory() float64`

GetPreAllocatedMemory returns the PreAllocatedMemory field if non-nil, zero value otherwise.

### GetPreAllocatedMemoryOk

`func (o *LoadReport) GetPreAllocatedMemoryOk() (*float64, bool)`

GetPreAllocatedMemoryOk returns a tuple with the PreAllocatedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAllocatedMemory

`func (o *LoadReport) SetPreAllocatedMemory(v float64)`

SetPreAllocatedMemory sets PreAllocatedMemory field to given value.

### HasPreAllocatedMemory

`func (o *LoadReport) HasPreAllocatedMemory() bool`

HasPreAllocatedMemory returns a boolean if a field has been set.

### GetPreAllocatedBandwidthIn

`func (o *LoadReport) GetPreAllocatedBandwidthIn() float64`

GetPreAllocatedBandwidthIn returns the PreAllocatedBandwidthIn field if non-nil, zero value otherwise.

### GetPreAllocatedBandwidthInOk

`func (o *LoadReport) GetPreAllocatedBandwidthInOk() (*float64, bool)`

GetPreAllocatedBandwidthInOk returns a tuple with the PreAllocatedBandwidthIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAllocatedBandwidthIn

`func (o *LoadReport) SetPreAllocatedBandwidthIn(v float64)`

SetPreAllocatedBandwidthIn sets PreAllocatedBandwidthIn field to given value.

### HasPreAllocatedBandwidthIn

`func (o *LoadReport) HasPreAllocatedBandwidthIn() bool`

HasPreAllocatedBandwidthIn returns a boolean if a field has been set.

### GetPreAllocatedBandwidthOut

`func (o *LoadReport) GetPreAllocatedBandwidthOut() float64`

GetPreAllocatedBandwidthOut returns the PreAllocatedBandwidthOut field if non-nil, zero value otherwise.

### GetPreAllocatedBandwidthOutOk

`func (o *LoadReport) GetPreAllocatedBandwidthOutOk() (*float64, bool)`

GetPreAllocatedBandwidthOutOk returns a tuple with the PreAllocatedBandwidthOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAllocatedBandwidthOut

`func (o *LoadReport) SetPreAllocatedBandwidthOut(v float64)`

SetPreAllocatedBandwidthOut sets PreAllocatedBandwidthOut field to given value.

### HasPreAllocatedBandwidthOut

`func (o *LoadReport) HasPreAllocatedBandwidthOut() bool`

HasPreAllocatedBandwidthOut returns a boolean if a field has been set.

### GetPreAllocatedMsgRateIn

`func (o *LoadReport) GetPreAllocatedMsgRateIn() float64`

GetPreAllocatedMsgRateIn returns the PreAllocatedMsgRateIn field if non-nil, zero value otherwise.

### GetPreAllocatedMsgRateInOk

`func (o *LoadReport) GetPreAllocatedMsgRateInOk() (*float64, bool)`

GetPreAllocatedMsgRateInOk returns a tuple with the PreAllocatedMsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAllocatedMsgRateIn

`func (o *LoadReport) SetPreAllocatedMsgRateIn(v float64)`

SetPreAllocatedMsgRateIn sets PreAllocatedMsgRateIn field to given value.

### HasPreAllocatedMsgRateIn

`func (o *LoadReport) HasPreAllocatedMsgRateIn() bool`

HasPreAllocatedMsgRateIn returns a boolean if a field has been set.

### GetPreAllocatedMsgRateOut

`func (o *LoadReport) GetPreAllocatedMsgRateOut() float64`

GetPreAllocatedMsgRateOut returns the PreAllocatedMsgRateOut field if non-nil, zero value otherwise.

### GetPreAllocatedMsgRateOutOk

`func (o *LoadReport) GetPreAllocatedMsgRateOutOk() (*float64, bool)`

GetPreAllocatedMsgRateOutOk returns a tuple with the PreAllocatedMsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAllocatedMsgRateOut

`func (o *LoadReport) SetPreAllocatedMsgRateOut(v float64)`

SetPreAllocatedMsgRateOut sets PreAllocatedMsgRateOut field to given value.

### HasPreAllocatedMsgRateOut

`func (o *LoadReport) HasPreAllocatedMsgRateOut() bool`

HasPreAllocatedMsgRateOut returns a boolean if a field has been set.

### GetUnderLoaded

`func (o *LoadReport) GetUnderLoaded() bool`

GetUnderLoaded returns the UnderLoaded field if non-nil, zero value otherwise.

### GetUnderLoadedOk

`func (o *LoadReport) GetUnderLoadedOk() (*bool, bool)`

GetUnderLoadedOk returns a tuple with the UnderLoaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderLoaded

`func (o *LoadReport) SetUnderLoaded(v bool)`

SetUnderLoaded sets UnderLoaded field to given value.

### HasUnderLoaded

`func (o *LoadReport) HasUnderLoaded() bool`

HasUnderLoaded returns a boolean if a field has been set.

### GetOverLoaded

`func (o *LoadReport) GetOverLoaded() bool`

GetOverLoaded returns the OverLoaded field if non-nil, zero value otherwise.

### GetOverLoadedOk

`func (o *LoadReport) GetOverLoadedOk() (*bool, bool)`

GetOverLoadedOk returns a tuple with the OverLoaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverLoaded

`func (o *LoadReport) SetOverLoaded(v bool)`

SetOverLoaded sets OverLoaded field to given value.

### HasOverLoaded

`func (o *LoadReport) HasOverLoaded() bool`

HasOverLoaded returns a boolean if a field has been set.

### GetLoadReportType

`func (o *LoadReport) GetLoadReportType() string`

GetLoadReportType returns the LoadReportType field if non-nil, zero value otherwise.

### GetLoadReportTypeOk

`func (o *LoadReport) GetLoadReportTypeOk() (*string, bool)`

GetLoadReportTypeOk returns a tuple with the LoadReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadReportType

`func (o *LoadReport) SetLoadReportType(v string)`

SetLoadReportType sets LoadReportType field to given value.

### HasLoadReportType

`func (o *LoadReport) HasLoadReportType() bool`

HasLoadReportType returns a boolean if a field has been set.

### GetMsgThroughputIn

`func (o *LoadReport) GetMsgThroughputIn() float64`

GetMsgThroughputIn returns the MsgThroughputIn field if non-nil, zero value otherwise.

### GetMsgThroughputInOk

`func (o *LoadReport) GetMsgThroughputInOk() (*float64, bool)`

GetMsgThroughputInOk returns a tuple with the MsgThroughputIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputIn

`func (o *LoadReport) SetMsgThroughputIn(v float64)`

SetMsgThroughputIn sets MsgThroughputIn field to given value.

### HasMsgThroughputIn

`func (o *LoadReport) HasMsgThroughputIn() bool`

HasMsgThroughputIn returns a boolean if a field has been set.

### GetLastUpdate

`func (o *LoadReport) GetLastUpdate() int64`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *LoadReport) GetLastUpdateOk() (*int64, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *LoadReport) SetLastUpdate(v int64)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *LoadReport) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetDirectMemory

`func (o *LoadReport) GetDirectMemory() ResourceUsage`

GetDirectMemory returns the DirectMemory field if non-nil, zero value otherwise.

### GetDirectMemoryOk

`func (o *LoadReport) GetDirectMemoryOk() (*ResourceUsage, bool)`

GetDirectMemoryOk returns a tuple with the DirectMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMemory

`func (o *LoadReport) SetDirectMemory(v ResourceUsage)`

SetDirectMemory sets DirectMemory field to given value.

### HasDirectMemory

`func (o *LoadReport) HasDirectMemory() bool`

HasDirectMemory returns a boolean if a field has been set.

### GetMsgThroughputOut

`func (o *LoadReport) GetMsgThroughputOut() float64`

GetMsgThroughputOut returns the MsgThroughputOut field if non-nil, zero value otherwise.

### GetMsgThroughputOutOk

`func (o *LoadReport) GetMsgThroughputOutOk() (*float64, bool)`

GetMsgThroughputOutOk returns a tuple with the MsgThroughputOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputOut

`func (o *LoadReport) SetMsgThroughputOut(v float64)`

SetMsgThroughputOut sets MsgThroughputOut field to given value.

### HasMsgThroughputOut

`func (o *LoadReport) HasMsgThroughputOut() bool`

HasMsgThroughputOut returns a boolean if a field has been set.

### GetCpu

`func (o *LoadReport) GetCpu() ResourceUsage`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *LoadReport) GetCpuOk() (*ResourceUsage, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *LoadReport) SetCpu(v ResourceUsage)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *LoadReport) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *LoadReport) GetMemory() ResourceUsage`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *LoadReport) GetMemoryOk() (*ResourceUsage, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *LoadReport) SetMemory(v ResourceUsage)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *LoadReport) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetBandwidthIn

`func (o *LoadReport) GetBandwidthIn() ResourceUsage`

GetBandwidthIn returns the BandwidthIn field if non-nil, zero value otherwise.

### GetBandwidthInOk

`func (o *LoadReport) GetBandwidthInOk() (*ResourceUsage, bool)`

GetBandwidthInOk returns a tuple with the BandwidthIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthIn

`func (o *LoadReport) SetBandwidthIn(v ResourceUsage)`

SetBandwidthIn sets BandwidthIn field to given value.

### HasBandwidthIn

`func (o *LoadReport) HasBandwidthIn() bool`

HasBandwidthIn returns a boolean if a field has been set.

### GetBandwidthOut

`func (o *LoadReport) GetBandwidthOut() ResourceUsage`

GetBandwidthOut returns the BandwidthOut field if non-nil, zero value otherwise.

### GetBandwidthOutOk

`func (o *LoadReport) GetBandwidthOutOk() (*ResourceUsage, bool)`

GetBandwidthOutOk returns a tuple with the BandwidthOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthOut

`func (o *LoadReport) SetBandwidthOut(v ResourceUsage)`

SetBandwidthOut sets BandwidthOut field to given value.

### HasBandwidthOut

`func (o *LoadReport) HasBandwidthOut() bool`

HasBandwidthOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


