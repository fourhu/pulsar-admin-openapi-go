# ObjectCodec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonFactory** | Pointer to [**JsonFactory**](JsonFactory.md) |  | [optional] 
**Factory** | Pointer to [**JsonFactory**](JsonFactory.md) |  | [optional] 

## Methods

### NewObjectCodec

`func NewObjectCodec() *ObjectCodec`

NewObjectCodec instantiates a new ObjectCodec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectCodecWithDefaults

`func NewObjectCodecWithDefaults() *ObjectCodec`

NewObjectCodecWithDefaults instantiates a new ObjectCodec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonFactory

`func (o *ObjectCodec) GetJsonFactory() JsonFactory`

GetJsonFactory returns the JsonFactory field if non-nil, zero value otherwise.

### GetJsonFactoryOk

`func (o *ObjectCodec) GetJsonFactoryOk() (*JsonFactory, bool)`

GetJsonFactoryOk returns a tuple with the JsonFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonFactory

`func (o *ObjectCodec) SetJsonFactory(v JsonFactory)`

SetJsonFactory sets JsonFactory field to given value.

### HasJsonFactory

`func (o *ObjectCodec) HasJsonFactory() bool`

HasJsonFactory returns a boolean if a field has been set.

### GetFactory

`func (o *ObjectCodec) GetFactory() JsonFactory`

GetFactory returns the Factory field if non-nil, zero value otherwise.

### GetFactoryOk

`func (o *ObjectCodec) GetFactoryOk() (*JsonFactory, bool)`

GetFactoryOk returns a tuple with the Factory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactory

`func (o *ObjectCodec) SetFactory(v JsonFactory)`

SetFactory sets Factory field to given value.

### HasFactory

`func (o *ObjectCodec) HasFactory() bool`

HasFactory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


