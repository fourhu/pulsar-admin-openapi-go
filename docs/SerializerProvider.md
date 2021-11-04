# SerializerProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterProvider** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultNullKeySerializer** | Pointer to [**JsonSerializerObject**](JsonSerializerObject.md) |  | [optional] 
**DefaultNullValueSerializer** | Pointer to [**JsonSerializerObject**](JsonSerializerObject.md) |  | [optional] 
**TypeFactory** | Pointer to [**TypeFactory**](TypeFactory.md) |  | [optional] 
**AnnotationIntrospector** | Pointer to **map[string]interface{}** |  | [optional] 
**TimeZone** | Pointer to [**TimeZone**](TimeZone.md) |  | [optional] 
**Locale** | Pointer to [**Locale**](Locale.md) |  | [optional] 
**Config** | Pointer to [**SerializationConfig**](SerializationConfig.md) |  | [optional] 
**Generator** | Pointer to [**JsonGenerator**](JsonGenerator.md) |  | [optional] 

## Methods

### NewSerializerProvider

`func NewSerializerProvider() *SerializerProvider`

NewSerializerProvider instantiates a new SerializerProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSerializerProviderWithDefaults

`func NewSerializerProviderWithDefaults() *SerializerProvider`

NewSerializerProviderWithDefaults instantiates a new SerializerProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterProvider

`func (o *SerializerProvider) GetFilterProvider() map[string]interface{}`

GetFilterProvider returns the FilterProvider field if non-nil, zero value otherwise.

### GetFilterProviderOk

`func (o *SerializerProvider) GetFilterProviderOk() (*map[string]interface{}, bool)`

GetFilterProviderOk returns a tuple with the FilterProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterProvider

`func (o *SerializerProvider) SetFilterProvider(v map[string]interface{})`

SetFilterProvider sets FilterProvider field to given value.

### HasFilterProvider

`func (o *SerializerProvider) HasFilterProvider() bool`

HasFilterProvider returns a boolean if a field has been set.

### GetDefaultNullKeySerializer

`func (o *SerializerProvider) GetDefaultNullKeySerializer() JsonSerializerObject`

GetDefaultNullKeySerializer returns the DefaultNullKeySerializer field if non-nil, zero value otherwise.

### GetDefaultNullKeySerializerOk

`func (o *SerializerProvider) GetDefaultNullKeySerializerOk() (*JsonSerializerObject, bool)`

GetDefaultNullKeySerializerOk returns a tuple with the DefaultNullKeySerializer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNullKeySerializer

`func (o *SerializerProvider) SetDefaultNullKeySerializer(v JsonSerializerObject)`

SetDefaultNullKeySerializer sets DefaultNullKeySerializer field to given value.

### HasDefaultNullKeySerializer

`func (o *SerializerProvider) HasDefaultNullKeySerializer() bool`

HasDefaultNullKeySerializer returns a boolean if a field has been set.

### GetDefaultNullValueSerializer

`func (o *SerializerProvider) GetDefaultNullValueSerializer() JsonSerializerObject`

GetDefaultNullValueSerializer returns the DefaultNullValueSerializer field if non-nil, zero value otherwise.

### GetDefaultNullValueSerializerOk

`func (o *SerializerProvider) GetDefaultNullValueSerializerOk() (*JsonSerializerObject, bool)`

GetDefaultNullValueSerializerOk returns a tuple with the DefaultNullValueSerializer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNullValueSerializer

`func (o *SerializerProvider) SetDefaultNullValueSerializer(v JsonSerializerObject)`

SetDefaultNullValueSerializer sets DefaultNullValueSerializer field to given value.

### HasDefaultNullValueSerializer

`func (o *SerializerProvider) HasDefaultNullValueSerializer() bool`

HasDefaultNullValueSerializer returns a boolean if a field has been set.

### GetTypeFactory

`func (o *SerializerProvider) GetTypeFactory() TypeFactory`

GetTypeFactory returns the TypeFactory field if non-nil, zero value otherwise.

### GetTypeFactoryOk

`func (o *SerializerProvider) GetTypeFactoryOk() (*TypeFactory, bool)`

GetTypeFactoryOk returns a tuple with the TypeFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFactory

`func (o *SerializerProvider) SetTypeFactory(v TypeFactory)`

SetTypeFactory sets TypeFactory field to given value.

### HasTypeFactory

`func (o *SerializerProvider) HasTypeFactory() bool`

HasTypeFactory returns a boolean if a field has been set.

### GetAnnotationIntrospector

`func (o *SerializerProvider) GetAnnotationIntrospector() map[string]interface{}`

GetAnnotationIntrospector returns the AnnotationIntrospector field if non-nil, zero value otherwise.

### GetAnnotationIntrospectorOk

`func (o *SerializerProvider) GetAnnotationIntrospectorOk() (*map[string]interface{}, bool)`

GetAnnotationIntrospectorOk returns a tuple with the AnnotationIntrospector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationIntrospector

`func (o *SerializerProvider) SetAnnotationIntrospector(v map[string]interface{})`

SetAnnotationIntrospector sets AnnotationIntrospector field to given value.

### HasAnnotationIntrospector

`func (o *SerializerProvider) HasAnnotationIntrospector() bool`

HasAnnotationIntrospector returns a boolean if a field has been set.

### GetTimeZone

`func (o *SerializerProvider) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SerializerProvider) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SerializerProvider) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SerializerProvider) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetLocale

`func (o *SerializerProvider) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SerializerProvider) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SerializerProvider) SetLocale(v Locale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SerializerProvider) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetConfig

`func (o *SerializerProvider) GetConfig() SerializationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SerializerProvider) GetConfigOk() (*SerializationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SerializerProvider) SetConfig(v SerializationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SerializerProvider) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetGenerator

`func (o *SerializerProvider) GetGenerator() JsonGenerator`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *SerializerProvider) GetGeneratorOk() (*JsonGenerator, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *SerializerProvider) SetGenerator(v JsonGenerator)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *SerializerProvider) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


