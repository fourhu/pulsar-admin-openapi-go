# PublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] 
**Algorithm** | Pointer to **string** |  | [optional] 
**Encoded** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPublicKey

`func NewPublicKey() *PublicKey`

NewPublicKey instantiates a new PublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicKeyWithDefaults

`func NewPublicKeyWithDefaults() *PublicKey`

NewPublicKeyWithDefaults instantiates a new PublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *PublicKey) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PublicKey) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PublicKey) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PublicKey) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetAlgorithm

`func (o *PublicKey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *PublicKey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *PublicKey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *PublicKey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetEncoded

`func (o *PublicKey) GetEncoded() []string`

GetEncoded returns the Encoded field if non-nil, zero value otherwise.

### GetEncodedOk

`func (o *PublicKey) GetEncodedOk() (*[]string, bool)`

GetEncodedOk returns a tuple with the Encoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoded

`func (o *PublicKey) SetEncoded(v []string)`

SetEncoded sets Encoded field to given value.

### HasEncoded

`func (o *PublicKey) HasEncoded() bool`

HasEncoded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


