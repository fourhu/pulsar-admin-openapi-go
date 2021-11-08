# OffloadPoliciesImpl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OffloadersDirectory** | Pointer to **string** |  | [optional] 
**ManagedLedgerOffloadDriver** | Pointer to **string** |  | [optional] 
**ManagedLedgerOffloadMaxThreads** | Pointer to **int32** |  | [optional] 
**ManagedLedgerOffloadPrefetchRounds** | Pointer to **int32** |  | [optional] 
**ManagedLedgerOffloadThresholdInBytes** | Pointer to **int64** |  | [optional] 
**ManagedLedgerOffloadDeletionLagInMillis** | Pointer to **int64** |  | [optional] 
**ManagedLedgerOffloadedReadPriority** | Pointer to **string** |  | [optional] 
**S3ManagedLedgerOffloadRegion** | Pointer to **string** |  | [optional] 
**S3ManagedLedgerOffloadBucket** | Pointer to **string** |  | [optional] 
**S3ManagedLedgerOffloadServiceEndpoint** | Pointer to **string** |  | [optional] 
**S3ManagedLedgerOffloadMaxBlockSizeInBytes** | Pointer to **int32** |  | [optional] 
**S3ManagedLedgerOffloadReadBufferSizeInBytes** | Pointer to **int32** |  | [optional] 
**S3ManagedLedgerOffloadCredentialId** | Pointer to **string** |  | [optional] 
**S3ManagedLedgerOffloadCredentialSecret** | Pointer to **string** |  | [optional] 
**S3ManagedLedgerOffloadRole** | Pointer to **string** |  | [optional] 
**S3ManagedLedgerOffloadRoleSessionName** | Pointer to **string** |  | [optional] 
**GcsManagedLedgerOffloadRegion** | Pointer to **string** |  | [optional] 
**GcsManagedLedgerOffloadBucket** | Pointer to **string** |  | [optional] 
**GcsManagedLedgerOffloadMaxBlockSizeInBytes** | Pointer to **int32** |  | [optional] 
**GcsManagedLedgerOffloadReadBufferSizeInBytes** | Pointer to **int32** |  | [optional] 
**GcsManagedLedgerOffloadServiceAccountKeyFile** | Pointer to **string** |  | [optional] 
**FileSystemProfilePath** | Pointer to **string** |  | [optional] 
**FileSystemURI** | Pointer to **string** |  | [optional] 
**ManagedLedgerOffloadBucket** | Pointer to **string** |  | [optional] 
**ManagedLedgerOffloadRegion** | Pointer to **string** |  | [optional] 
**ManagedLedgerOffloadServiceEndpoint** | Pointer to **string** |  | [optional] 
**ManagedLedgerOffloadMaxBlockSizeInBytes** | Pointer to **int32** |  | [optional] 
**ManagedLedgerOffloadReadBufferSizeInBytes** | Pointer to **int32** |  | [optional] 
**S3Driver** | Pointer to **bool** |  | [optional] 
**GcsDriver** | Pointer to **bool** |  | [optional] 
**FileSystemDriver** | Pointer to **bool** |  | [optional] 

## Methods

### NewOffloadPoliciesImpl

`func NewOffloadPoliciesImpl() *OffloadPoliciesImpl`

NewOffloadPoliciesImpl instantiates a new OffloadPoliciesImpl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffloadPoliciesImplWithDefaults

`func NewOffloadPoliciesImplWithDefaults() *OffloadPoliciesImpl`

NewOffloadPoliciesImplWithDefaults instantiates a new OffloadPoliciesImpl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffloadersDirectory

`func (o *OffloadPoliciesImpl) GetOffloadersDirectory() string`

GetOffloadersDirectory returns the OffloadersDirectory field if non-nil, zero value otherwise.

### GetOffloadersDirectoryOk

`func (o *OffloadPoliciesImpl) GetOffloadersDirectoryOk() (*string, bool)`

GetOffloadersDirectoryOk returns a tuple with the OffloadersDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffloadersDirectory

`func (o *OffloadPoliciesImpl) SetOffloadersDirectory(v string)`

SetOffloadersDirectory sets OffloadersDirectory field to given value.

### HasOffloadersDirectory

`func (o *OffloadPoliciesImpl) HasOffloadersDirectory() bool`

HasOffloadersDirectory returns a boolean if a field has been set.

### GetManagedLedgerOffloadDriver

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadDriver() string`

GetManagedLedgerOffloadDriver returns the ManagedLedgerOffloadDriver field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadDriverOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadDriverOk() (*string, bool)`

GetManagedLedgerOffloadDriverOk returns a tuple with the ManagedLedgerOffloadDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadDriver

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadDriver(v string)`

SetManagedLedgerOffloadDriver sets ManagedLedgerOffloadDriver field to given value.

### HasManagedLedgerOffloadDriver

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadDriver() bool`

HasManagedLedgerOffloadDriver returns a boolean if a field has been set.

### GetManagedLedgerOffloadMaxThreads

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadMaxThreads() int32`

GetManagedLedgerOffloadMaxThreads returns the ManagedLedgerOffloadMaxThreads field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadMaxThreadsOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadMaxThreadsOk() (*int32, bool)`

GetManagedLedgerOffloadMaxThreadsOk returns a tuple with the ManagedLedgerOffloadMaxThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadMaxThreads

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadMaxThreads(v int32)`

SetManagedLedgerOffloadMaxThreads sets ManagedLedgerOffloadMaxThreads field to given value.

### HasManagedLedgerOffloadMaxThreads

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadMaxThreads() bool`

HasManagedLedgerOffloadMaxThreads returns a boolean if a field has been set.

### GetManagedLedgerOffloadPrefetchRounds

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadPrefetchRounds() int32`

GetManagedLedgerOffloadPrefetchRounds returns the ManagedLedgerOffloadPrefetchRounds field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadPrefetchRoundsOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadPrefetchRoundsOk() (*int32, bool)`

GetManagedLedgerOffloadPrefetchRoundsOk returns a tuple with the ManagedLedgerOffloadPrefetchRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadPrefetchRounds

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadPrefetchRounds(v int32)`

SetManagedLedgerOffloadPrefetchRounds sets ManagedLedgerOffloadPrefetchRounds field to given value.

### HasManagedLedgerOffloadPrefetchRounds

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadPrefetchRounds() bool`

HasManagedLedgerOffloadPrefetchRounds returns a boolean if a field has been set.

### GetManagedLedgerOffloadThresholdInBytes

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadThresholdInBytes() int64`

GetManagedLedgerOffloadThresholdInBytes returns the ManagedLedgerOffloadThresholdInBytes field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadThresholdInBytesOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadThresholdInBytesOk() (*int64, bool)`

GetManagedLedgerOffloadThresholdInBytesOk returns a tuple with the ManagedLedgerOffloadThresholdInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadThresholdInBytes

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadThresholdInBytes(v int64)`

SetManagedLedgerOffloadThresholdInBytes sets ManagedLedgerOffloadThresholdInBytes field to given value.

### HasManagedLedgerOffloadThresholdInBytes

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadThresholdInBytes() bool`

HasManagedLedgerOffloadThresholdInBytes returns a boolean if a field has been set.

### GetManagedLedgerOffloadDeletionLagInMillis

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadDeletionLagInMillis() int64`

GetManagedLedgerOffloadDeletionLagInMillis returns the ManagedLedgerOffloadDeletionLagInMillis field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadDeletionLagInMillisOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadDeletionLagInMillisOk() (*int64, bool)`

GetManagedLedgerOffloadDeletionLagInMillisOk returns a tuple with the ManagedLedgerOffloadDeletionLagInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadDeletionLagInMillis

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadDeletionLagInMillis(v int64)`

SetManagedLedgerOffloadDeletionLagInMillis sets ManagedLedgerOffloadDeletionLagInMillis field to given value.

### HasManagedLedgerOffloadDeletionLagInMillis

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadDeletionLagInMillis() bool`

HasManagedLedgerOffloadDeletionLagInMillis returns a boolean if a field has been set.

### GetManagedLedgerOffloadedReadPriority

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadedReadPriority() string`

GetManagedLedgerOffloadedReadPriority returns the ManagedLedgerOffloadedReadPriority field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadedReadPriorityOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadedReadPriorityOk() (*string, bool)`

GetManagedLedgerOffloadedReadPriorityOk returns a tuple with the ManagedLedgerOffloadedReadPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadedReadPriority

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadedReadPriority(v string)`

SetManagedLedgerOffloadedReadPriority sets ManagedLedgerOffloadedReadPriority field to given value.

### HasManagedLedgerOffloadedReadPriority

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadedReadPriority() bool`

HasManagedLedgerOffloadedReadPriority returns a boolean if a field has been set.

### GetS3ManagedLedgerOffloadRegion

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadRegion() string`

GetS3ManagedLedgerOffloadRegion returns the S3ManagedLedgerOffloadRegion field if non-nil, zero value otherwise.

### GetS3ManagedLedgerOffloadRegionOk

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadRegionOk() (*string, bool)`

GetS3ManagedLedgerOffloadRegionOk returns a tuple with the S3ManagedLedgerOffloadRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ManagedLedgerOffloadRegion

`func (o *OffloadPoliciesImpl) SetS3ManagedLedgerOffloadRegion(v string)`

SetS3ManagedLedgerOffloadRegion sets S3ManagedLedgerOffloadRegion field to given value.

### HasS3ManagedLedgerOffloadRegion

`func (o *OffloadPoliciesImpl) HasS3ManagedLedgerOffloadRegion() bool`

HasS3ManagedLedgerOffloadRegion returns a boolean if a field has been set.

### GetS3ManagedLedgerOffloadBucket

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadBucket() string`

GetS3ManagedLedgerOffloadBucket returns the S3ManagedLedgerOffloadBucket field if non-nil, zero value otherwise.

### GetS3ManagedLedgerOffloadBucketOk

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadBucketOk() (*string, bool)`

GetS3ManagedLedgerOffloadBucketOk returns a tuple with the S3ManagedLedgerOffloadBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ManagedLedgerOffloadBucket

`func (o *OffloadPoliciesImpl) SetS3ManagedLedgerOffloadBucket(v string)`

SetS3ManagedLedgerOffloadBucket sets S3ManagedLedgerOffloadBucket field to given value.

### HasS3ManagedLedgerOffloadBucket

`func (o *OffloadPoliciesImpl) HasS3ManagedLedgerOffloadBucket() bool`

HasS3ManagedLedgerOffloadBucket returns a boolean if a field has been set.

### GetS3ManagedLedgerOffloadServiceEndpoint

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadServiceEndpoint() string`

GetS3ManagedLedgerOffloadServiceEndpoint returns the S3ManagedLedgerOffloadServiceEndpoint field if non-nil, zero value otherwise.

### GetS3ManagedLedgerOffloadServiceEndpointOk

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadServiceEndpointOk() (*string, bool)`

GetS3ManagedLedgerOffloadServiceEndpointOk returns a tuple with the S3ManagedLedgerOffloadServiceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ManagedLedgerOffloadServiceEndpoint

`func (o *OffloadPoliciesImpl) SetS3ManagedLedgerOffloadServiceEndpoint(v string)`

SetS3ManagedLedgerOffloadServiceEndpoint sets S3ManagedLedgerOffloadServiceEndpoint field to given value.

### HasS3ManagedLedgerOffloadServiceEndpoint

`func (o *OffloadPoliciesImpl) HasS3ManagedLedgerOffloadServiceEndpoint() bool`

HasS3ManagedLedgerOffloadServiceEndpoint returns a boolean if a field has been set.

### GetS3ManagedLedgerOffloadMaxBlockSizeInBytes

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadMaxBlockSizeInBytes() int32`

GetS3ManagedLedgerOffloadMaxBlockSizeInBytes returns the S3ManagedLedgerOffloadMaxBlockSizeInBytes field if non-nil, zero value otherwise.

### GetS3ManagedLedgerOffloadMaxBlockSizeInBytesOk

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadMaxBlockSizeInBytesOk() (*int32, bool)`

GetS3ManagedLedgerOffloadMaxBlockSizeInBytesOk returns a tuple with the S3ManagedLedgerOffloadMaxBlockSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ManagedLedgerOffloadMaxBlockSizeInBytes

`func (o *OffloadPoliciesImpl) SetS3ManagedLedgerOffloadMaxBlockSizeInBytes(v int32)`

SetS3ManagedLedgerOffloadMaxBlockSizeInBytes sets S3ManagedLedgerOffloadMaxBlockSizeInBytes field to given value.

### HasS3ManagedLedgerOffloadMaxBlockSizeInBytes

`func (o *OffloadPoliciesImpl) HasS3ManagedLedgerOffloadMaxBlockSizeInBytes() bool`

HasS3ManagedLedgerOffloadMaxBlockSizeInBytes returns a boolean if a field has been set.

### GetS3ManagedLedgerOffloadReadBufferSizeInBytes

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadReadBufferSizeInBytes() int32`

GetS3ManagedLedgerOffloadReadBufferSizeInBytes returns the S3ManagedLedgerOffloadReadBufferSizeInBytes field if non-nil, zero value otherwise.

### GetS3ManagedLedgerOffloadReadBufferSizeInBytesOk

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadReadBufferSizeInBytesOk() (*int32, bool)`

GetS3ManagedLedgerOffloadReadBufferSizeInBytesOk returns a tuple with the S3ManagedLedgerOffloadReadBufferSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ManagedLedgerOffloadReadBufferSizeInBytes

`func (o *OffloadPoliciesImpl) SetS3ManagedLedgerOffloadReadBufferSizeInBytes(v int32)`

SetS3ManagedLedgerOffloadReadBufferSizeInBytes sets S3ManagedLedgerOffloadReadBufferSizeInBytes field to given value.

### HasS3ManagedLedgerOffloadReadBufferSizeInBytes

`func (o *OffloadPoliciesImpl) HasS3ManagedLedgerOffloadReadBufferSizeInBytes() bool`

HasS3ManagedLedgerOffloadReadBufferSizeInBytes returns a boolean if a field has been set.

### GetS3ManagedLedgerOffloadCredentialId

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadCredentialId() string`

GetS3ManagedLedgerOffloadCredentialId returns the S3ManagedLedgerOffloadCredentialId field if non-nil, zero value otherwise.

### GetS3ManagedLedgerOffloadCredentialIdOk

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadCredentialIdOk() (*string, bool)`

GetS3ManagedLedgerOffloadCredentialIdOk returns a tuple with the S3ManagedLedgerOffloadCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ManagedLedgerOffloadCredentialId

`func (o *OffloadPoliciesImpl) SetS3ManagedLedgerOffloadCredentialId(v string)`

SetS3ManagedLedgerOffloadCredentialId sets S3ManagedLedgerOffloadCredentialId field to given value.

### HasS3ManagedLedgerOffloadCredentialId

`func (o *OffloadPoliciesImpl) HasS3ManagedLedgerOffloadCredentialId() bool`

HasS3ManagedLedgerOffloadCredentialId returns a boolean if a field has been set.

### GetS3ManagedLedgerOffloadCredentialSecret

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadCredentialSecret() string`

GetS3ManagedLedgerOffloadCredentialSecret returns the S3ManagedLedgerOffloadCredentialSecret field if non-nil, zero value otherwise.

### GetS3ManagedLedgerOffloadCredentialSecretOk

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadCredentialSecretOk() (*string, bool)`

GetS3ManagedLedgerOffloadCredentialSecretOk returns a tuple with the S3ManagedLedgerOffloadCredentialSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ManagedLedgerOffloadCredentialSecret

`func (o *OffloadPoliciesImpl) SetS3ManagedLedgerOffloadCredentialSecret(v string)`

SetS3ManagedLedgerOffloadCredentialSecret sets S3ManagedLedgerOffloadCredentialSecret field to given value.

### HasS3ManagedLedgerOffloadCredentialSecret

`func (o *OffloadPoliciesImpl) HasS3ManagedLedgerOffloadCredentialSecret() bool`

HasS3ManagedLedgerOffloadCredentialSecret returns a boolean if a field has been set.

### GetS3ManagedLedgerOffloadRole

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadRole() string`

GetS3ManagedLedgerOffloadRole returns the S3ManagedLedgerOffloadRole field if non-nil, zero value otherwise.

### GetS3ManagedLedgerOffloadRoleOk

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadRoleOk() (*string, bool)`

GetS3ManagedLedgerOffloadRoleOk returns a tuple with the S3ManagedLedgerOffloadRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ManagedLedgerOffloadRole

`func (o *OffloadPoliciesImpl) SetS3ManagedLedgerOffloadRole(v string)`

SetS3ManagedLedgerOffloadRole sets S3ManagedLedgerOffloadRole field to given value.

### HasS3ManagedLedgerOffloadRole

`func (o *OffloadPoliciesImpl) HasS3ManagedLedgerOffloadRole() bool`

HasS3ManagedLedgerOffloadRole returns a boolean if a field has been set.

### GetS3ManagedLedgerOffloadRoleSessionName

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadRoleSessionName() string`

GetS3ManagedLedgerOffloadRoleSessionName returns the S3ManagedLedgerOffloadRoleSessionName field if non-nil, zero value otherwise.

### GetS3ManagedLedgerOffloadRoleSessionNameOk

`func (o *OffloadPoliciesImpl) GetS3ManagedLedgerOffloadRoleSessionNameOk() (*string, bool)`

GetS3ManagedLedgerOffloadRoleSessionNameOk returns a tuple with the S3ManagedLedgerOffloadRoleSessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ManagedLedgerOffloadRoleSessionName

`func (o *OffloadPoliciesImpl) SetS3ManagedLedgerOffloadRoleSessionName(v string)`

SetS3ManagedLedgerOffloadRoleSessionName sets S3ManagedLedgerOffloadRoleSessionName field to given value.

### HasS3ManagedLedgerOffloadRoleSessionName

`func (o *OffloadPoliciesImpl) HasS3ManagedLedgerOffloadRoleSessionName() bool`

HasS3ManagedLedgerOffloadRoleSessionName returns a boolean if a field has been set.

### GetGcsManagedLedgerOffloadRegion

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadRegion() string`

GetGcsManagedLedgerOffloadRegion returns the GcsManagedLedgerOffloadRegion field if non-nil, zero value otherwise.

### GetGcsManagedLedgerOffloadRegionOk

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadRegionOk() (*string, bool)`

GetGcsManagedLedgerOffloadRegionOk returns a tuple with the GcsManagedLedgerOffloadRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsManagedLedgerOffloadRegion

`func (o *OffloadPoliciesImpl) SetGcsManagedLedgerOffloadRegion(v string)`

SetGcsManagedLedgerOffloadRegion sets GcsManagedLedgerOffloadRegion field to given value.

### HasGcsManagedLedgerOffloadRegion

`func (o *OffloadPoliciesImpl) HasGcsManagedLedgerOffloadRegion() bool`

HasGcsManagedLedgerOffloadRegion returns a boolean if a field has been set.

### GetGcsManagedLedgerOffloadBucket

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadBucket() string`

GetGcsManagedLedgerOffloadBucket returns the GcsManagedLedgerOffloadBucket field if non-nil, zero value otherwise.

### GetGcsManagedLedgerOffloadBucketOk

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadBucketOk() (*string, bool)`

GetGcsManagedLedgerOffloadBucketOk returns a tuple with the GcsManagedLedgerOffloadBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsManagedLedgerOffloadBucket

`func (o *OffloadPoliciesImpl) SetGcsManagedLedgerOffloadBucket(v string)`

SetGcsManagedLedgerOffloadBucket sets GcsManagedLedgerOffloadBucket field to given value.

### HasGcsManagedLedgerOffloadBucket

`func (o *OffloadPoliciesImpl) HasGcsManagedLedgerOffloadBucket() bool`

HasGcsManagedLedgerOffloadBucket returns a boolean if a field has been set.

### GetGcsManagedLedgerOffloadMaxBlockSizeInBytes

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadMaxBlockSizeInBytes() int32`

GetGcsManagedLedgerOffloadMaxBlockSizeInBytes returns the GcsManagedLedgerOffloadMaxBlockSizeInBytes field if non-nil, zero value otherwise.

### GetGcsManagedLedgerOffloadMaxBlockSizeInBytesOk

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadMaxBlockSizeInBytesOk() (*int32, bool)`

GetGcsManagedLedgerOffloadMaxBlockSizeInBytesOk returns a tuple with the GcsManagedLedgerOffloadMaxBlockSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsManagedLedgerOffloadMaxBlockSizeInBytes

`func (o *OffloadPoliciesImpl) SetGcsManagedLedgerOffloadMaxBlockSizeInBytes(v int32)`

SetGcsManagedLedgerOffloadMaxBlockSizeInBytes sets GcsManagedLedgerOffloadMaxBlockSizeInBytes field to given value.

### HasGcsManagedLedgerOffloadMaxBlockSizeInBytes

`func (o *OffloadPoliciesImpl) HasGcsManagedLedgerOffloadMaxBlockSizeInBytes() bool`

HasGcsManagedLedgerOffloadMaxBlockSizeInBytes returns a boolean if a field has been set.

### GetGcsManagedLedgerOffloadReadBufferSizeInBytes

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadReadBufferSizeInBytes() int32`

GetGcsManagedLedgerOffloadReadBufferSizeInBytes returns the GcsManagedLedgerOffloadReadBufferSizeInBytes field if non-nil, zero value otherwise.

### GetGcsManagedLedgerOffloadReadBufferSizeInBytesOk

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadReadBufferSizeInBytesOk() (*int32, bool)`

GetGcsManagedLedgerOffloadReadBufferSizeInBytesOk returns a tuple with the GcsManagedLedgerOffloadReadBufferSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsManagedLedgerOffloadReadBufferSizeInBytes

`func (o *OffloadPoliciesImpl) SetGcsManagedLedgerOffloadReadBufferSizeInBytes(v int32)`

SetGcsManagedLedgerOffloadReadBufferSizeInBytes sets GcsManagedLedgerOffloadReadBufferSizeInBytes field to given value.

### HasGcsManagedLedgerOffloadReadBufferSizeInBytes

`func (o *OffloadPoliciesImpl) HasGcsManagedLedgerOffloadReadBufferSizeInBytes() bool`

HasGcsManagedLedgerOffloadReadBufferSizeInBytes returns a boolean if a field has been set.

### GetGcsManagedLedgerOffloadServiceAccountKeyFile

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadServiceAccountKeyFile() string`

GetGcsManagedLedgerOffloadServiceAccountKeyFile returns the GcsManagedLedgerOffloadServiceAccountKeyFile field if non-nil, zero value otherwise.

### GetGcsManagedLedgerOffloadServiceAccountKeyFileOk

`func (o *OffloadPoliciesImpl) GetGcsManagedLedgerOffloadServiceAccountKeyFileOk() (*string, bool)`

GetGcsManagedLedgerOffloadServiceAccountKeyFileOk returns a tuple with the GcsManagedLedgerOffloadServiceAccountKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsManagedLedgerOffloadServiceAccountKeyFile

`func (o *OffloadPoliciesImpl) SetGcsManagedLedgerOffloadServiceAccountKeyFile(v string)`

SetGcsManagedLedgerOffloadServiceAccountKeyFile sets GcsManagedLedgerOffloadServiceAccountKeyFile field to given value.

### HasGcsManagedLedgerOffloadServiceAccountKeyFile

`func (o *OffloadPoliciesImpl) HasGcsManagedLedgerOffloadServiceAccountKeyFile() bool`

HasGcsManagedLedgerOffloadServiceAccountKeyFile returns a boolean if a field has been set.

### GetFileSystemProfilePath

`func (o *OffloadPoliciesImpl) GetFileSystemProfilePath() string`

GetFileSystemProfilePath returns the FileSystemProfilePath field if non-nil, zero value otherwise.

### GetFileSystemProfilePathOk

`func (o *OffloadPoliciesImpl) GetFileSystemProfilePathOk() (*string, bool)`

GetFileSystemProfilePathOk returns a tuple with the FileSystemProfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemProfilePath

`func (o *OffloadPoliciesImpl) SetFileSystemProfilePath(v string)`

SetFileSystemProfilePath sets FileSystemProfilePath field to given value.

### HasFileSystemProfilePath

`func (o *OffloadPoliciesImpl) HasFileSystemProfilePath() bool`

HasFileSystemProfilePath returns a boolean if a field has been set.

### GetFileSystemURI

`func (o *OffloadPoliciesImpl) GetFileSystemURI() string`

GetFileSystemURI returns the FileSystemURI field if non-nil, zero value otherwise.

### GetFileSystemURIOk

`func (o *OffloadPoliciesImpl) GetFileSystemURIOk() (*string, bool)`

GetFileSystemURIOk returns a tuple with the FileSystemURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemURI

`func (o *OffloadPoliciesImpl) SetFileSystemURI(v string)`

SetFileSystemURI sets FileSystemURI field to given value.

### HasFileSystemURI

`func (o *OffloadPoliciesImpl) HasFileSystemURI() bool`

HasFileSystemURI returns a boolean if a field has been set.

### GetManagedLedgerOffloadBucket

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadBucket() string`

GetManagedLedgerOffloadBucket returns the ManagedLedgerOffloadBucket field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadBucketOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadBucketOk() (*string, bool)`

GetManagedLedgerOffloadBucketOk returns a tuple with the ManagedLedgerOffloadBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadBucket

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadBucket(v string)`

SetManagedLedgerOffloadBucket sets ManagedLedgerOffloadBucket field to given value.

### HasManagedLedgerOffloadBucket

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadBucket() bool`

HasManagedLedgerOffloadBucket returns a boolean if a field has been set.

### GetManagedLedgerOffloadRegion

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadRegion() string`

GetManagedLedgerOffloadRegion returns the ManagedLedgerOffloadRegion field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadRegionOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadRegionOk() (*string, bool)`

GetManagedLedgerOffloadRegionOk returns a tuple with the ManagedLedgerOffloadRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadRegion

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadRegion(v string)`

SetManagedLedgerOffloadRegion sets ManagedLedgerOffloadRegion field to given value.

### HasManagedLedgerOffloadRegion

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadRegion() bool`

HasManagedLedgerOffloadRegion returns a boolean if a field has been set.

### GetManagedLedgerOffloadServiceEndpoint

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadServiceEndpoint() string`

GetManagedLedgerOffloadServiceEndpoint returns the ManagedLedgerOffloadServiceEndpoint field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadServiceEndpointOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadServiceEndpointOk() (*string, bool)`

GetManagedLedgerOffloadServiceEndpointOk returns a tuple with the ManagedLedgerOffloadServiceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadServiceEndpoint

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadServiceEndpoint(v string)`

SetManagedLedgerOffloadServiceEndpoint sets ManagedLedgerOffloadServiceEndpoint field to given value.

### HasManagedLedgerOffloadServiceEndpoint

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadServiceEndpoint() bool`

HasManagedLedgerOffloadServiceEndpoint returns a boolean if a field has been set.

### GetManagedLedgerOffloadMaxBlockSizeInBytes

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadMaxBlockSizeInBytes() int32`

GetManagedLedgerOffloadMaxBlockSizeInBytes returns the ManagedLedgerOffloadMaxBlockSizeInBytes field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadMaxBlockSizeInBytesOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadMaxBlockSizeInBytesOk() (*int32, bool)`

GetManagedLedgerOffloadMaxBlockSizeInBytesOk returns a tuple with the ManagedLedgerOffloadMaxBlockSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadMaxBlockSizeInBytes

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadMaxBlockSizeInBytes(v int32)`

SetManagedLedgerOffloadMaxBlockSizeInBytes sets ManagedLedgerOffloadMaxBlockSizeInBytes field to given value.

### HasManagedLedgerOffloadMaxBlockSizeInBytes

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadMaxBlockSizeInBytes() bool`

HasManagedLedgerOffloadMaxBlockSizeInBytes returns a boolean if a field has been set.

### GetManagedLedgerOffloadReadBufferSizeInBytes

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadReadBufferSizeInBytes() int32`

GetManagedLedgerOffloadReadBufferSizeInBytes returns the ManagedLedgerOffloadReadBufferSizeInBytes field if non-nil, zero value otherwise.

### GetManagedLedgerOffloadReadBufferSizeInBytesOk

`func (o *OffloadPoliciesImpl) GetManagedLedgerOffloadReadBufferSizeInBytesOk() (*int32, bool)`

GetManagedLedgerOffloadReadBufferSizeInBytesOk returns a tuple with the ManagedLedgerOffloadReadBufferSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLedgerOffloadReadBufferSizeInBytes

`func (o *OffloadPoliciesImpl) SetManagedLedgerOffloadReadBufferSizeInBytes(v int32)`

SetManagedLedgerOffloadReadBufferSizeInBytes sets ManagedLedgerOffloadReadBufferSizeInBytes field to given value.

### HasManagedLedgerOffloadReadBufferSizeInBytes

`func (o *OffloadPoliciesImpl) HasManagedLedgerOffloadReadBufferSizeInBytes() bool`

HasManagedLedgerOffloadReadBufferSizeInBytes returns a boolean if a field has been set.

### GetS3Driver

`func (o *OffloadPoliciesImpl) GetS3Driver() bool`

GetS3Driver returns the S3Driver field if non-nil, zero value otherwise.

### GetS3DriverOk

`func (o *OffloadPoliciesImpl) GetS3DriverOk() (*bool, bool)`

GetS3DriverOk returns a tuple with the S3Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Driver

`func (o *OffloadPoliciesImpl) SetS3Driver(v bool)`

SetS3Driver sets S3Driver field to given value.

### HasS3Driver

`func (o *OffloadPoliciesImpl) HasS3Driver() bool`

HasS3Driver returns a boolean if a field has been set.

### GetGcsDriver

`func (o *OffloadPoliciesImpl) GetGcsDriver() bool`

GetGcsDriver returns the GcsDriver field if non-nil, zero value otherwise.

### GetGcsDriverOk

`func (o *OffloadPoliciesImpl) GetGcsDriverOk() (*bool, bool)`

GetGcsDriverOk returns a tuple with the GcsDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsDriver

`func (o *OffloadPoliciesImpl) SetGcsDriver(v bool)`

SetGcsDriver sets GcsDriver field to given value.

### HasGcsDriver

`func (o *OffloadPoliciesImpl) HasGcsDriver() bool`

HasGcsDriver returns a boolean if a field has been set.

### GetFileSystemDriver

`func (o *OffloadPoliciesImpl) GetFileSystemDriver() bool`

GetFileSystemDriver returns the FileSystemDriver field if non-nil, zero value otherwise.

### GetFileSystemDriverOk

`func (o *OffloadPoliciesImpl) GetFileSystemDriverOk() (*bool, bool)`

GetFileSystemDriverOk returns a tuple with the FileSystemDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemDriver

`func (o *OffloadPoliciesImpl) SetFileSystemDriver(v bool)`

SetFileSystemDriver sets FileSystemDriver field to given value.

### HasFileSystemDriver

`func (o *OffloadPoliciesImpl) HasFileSystemDriver() bool`

HasFileSystemDriver returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


