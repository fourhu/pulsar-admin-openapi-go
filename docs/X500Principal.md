# X500Principal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Encoded** | Pointer to **[]string** |  | [optional] 

## Methods

### NewX500Principal

`func NewX500Principal() *X500Principal`

NewX500Principal instantiates a new X500Principal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX500PrincipalWithDefaults

`func NewX500PrincipalWithDefaults() *X500Principal`

NewX500PrincipalWithDefaults instantiates a new X500Principal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *X500Principal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *X500Principal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *X500Principal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *X500Principal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEncoded

`func (o *X500Principal) GetEncoded() []string`

GetEncoded returns the Encoded field if non-nil, zero value otherwise.

### GetEncodedOk

`func (o *X500Principal) GetEncodedOk() (*[]string, bool)`

GetEncodedOk returns a tuple with the Encoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoded

`func (o *X500Principal) SetEncoded(v []string)`

SetEncoded sets Encoded field to given value.

### HasEncoded

`func (o *X500Principal) HasEncoded() bool`

HasEncoded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


