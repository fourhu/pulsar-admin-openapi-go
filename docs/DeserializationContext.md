# DeserializationContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeFactory** | Pointer to [**TypeFactory**](TypeFactory.md) |  | [optional] 
**NodeFactory** | Pointer to **map[string]interface{}** |  | [optional] 
**AnnotationIntrospector** | Pointer to **map[string]interface{}** |  | [optional] 
**ContextualType** | Pointer to [**JavaType**](JavaType.md) |  | [optional] 
**DeserializationFeatures** | Pointer to **int32** |  | [optional] 
**Base64Variant** | Pointer to [**Base64Variant**](Base64Variant.md) |  | [optional] 
**ArrayBuilders** | Pointer to [**ArrayBuilders**](ArrayBuilders.md) |  | [optional] 
**TimeZone** | Pointer to [**TimeZone**](TimeZone.md) |  | [optional] 
**Locale** | Pointer to [**Locale**](Locale.md) |  | [optional] 
**Factory** | Pointer to **map[string]interface{}** |  | [optional] 
**Config** | Pointer to [**DeserializationConfig**](DeserializationConfig.md) |  | [optional] 
**Parser** | Pointer to [**JsonParser**](JsonParser.md) |  | [optional] 

## Methods

### NewDeserializationContext

`func NewDeserializationContext() *DeserializationContext`

NewDeserializationContext instantiates a new DeserializationContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeserializationContextWithDefaults

`func NewDeserializationContextWithDefaults() *DeserializationContext`

NewDeserializationContextWithDefaults instantiates a new DeserializationContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeFactory

`func (o *DeserializationContext) GetTypeFactory() TypeFactory`

GetTypeFactory returns the TypeFactory field if non-nil, zero value otherwise.

### GetTypeFactoryOk

`func (o *DeserializationContext) GetTypeFactoryOk() (*TypeFactory, bool)`

GetTypeFactoryOk returns a tuple with the TypeFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFactory

`func (o *DeserializationContext) SetTypeFactory(v TypeFactory)`

SetTypeFactory sets TypeFactory field to given value.

### HasTypeFactory

`func (o *DeserializationContext) HasTypeFactory() bool`

HasTypeFactory returns a boolean if a field has been set.

### GetNodeFactory

`func (o *DeserializationContext) GetNodeFactory() map[string]interface{}`

GetNodeFactory returns the NodeFactory field if non-nil, zero value otherwise.

### GetNodeFactoryOk

`func (o *DeserializationContext) GetNodeFactoryOk() (*map[string]interface{}, bool)`

GetNodeFactoryOk returns a tuple with the NodeFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFactory

`func (o *DeserializationContext) SetNodeFactory(v map[string]interface{})`

SetNodeFactory sets NodeFactory field to given value.

### HasNodeFactory

`func (o *DeserializationContext) HasNodeFactory() bool`

HasNodeFactory returns a boolean if a field has been set.

### GetAnnotationIntrospector

`func (o *DeserializationContext) GetAnnotationIntrospector() map[string]interface{}`

GetAnnotationIntrospector returns the AnnotationIntrospector field if non-nil, zero value otherwise.

### GetAnnotationIntrospectorOk

`func (o *DeserializationContext) GetAnnotationIntrospectorOk() (*map[string]interface{}, bool)`

GetAnnotationIntrospectorOk returns a tuple with the AnnotationIntrospector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationIntrospector

`func (o *DeserializationContext) SetAnnotationIntrospector(v map[string]interface{})`

SetAnnotationIntrospector sets AnnotationIntrospector field to given value.

### HasAnnotationIntrospector

`func (o *DeserializationContext) HasAnnotationIntrospector() bool`

HasAnnotationIntrospector returns a boolean if a field has been set.

### GetContextualType

`func (o *DeserializationContext) GetContextualType() JavaType`

GetContextualType returns the ContextualType field if non-nil, zero value otherwise.

### GetContextualTypeOk

`func (o *DeserializationContext) GetContextualTypeOk() (*JavaType, bool)`

GetContextualTypeOk returns a tuple with the ContextualType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualType

`func (o *DeserializationContext) SetContextualType(v JavaType)`

SetContextualType sets ContextualType field to given value.

### HasContextualType

`func (o *DeserializationContext) HasContextualType() bool`

HasContextualType returns a boolean if a field has been set.

### GetDeserializationFeatures

`func (o *DeserializationContext) GetDeserializationFeatures() int32`

GetDeserializationFeatures returns the DeserializationFeatures field if non-nil, zero value otherwise.

### GetDeserializationFeaturesOk

`func (o *DeserializationContext) GetDeserializationFeaturesOk() (*int32, bool)`

GetDeserializationFeaturesOk returns a tuple with the DeserializationFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeserializationFeatures

`func (o *DeserializationContext) SetDeserializationFeatures(v int32)`

SetDeserializationFeatures sets DeserializationFeatures field to given value.

### HasDeserializationFeatures

`func (o *DeserializationContext) HasDeserializationFeatures() bool`

HasDeserializationFeatures returns a boolean if a field has been set.

### GetBase64Variant

`func (o *DeserializationContext) GetBase64Variant() Base64Variant`

GetBase64Variant returns the Base64Variant field if non-nil, zero value otherwise.

### GetBase64VariantOk

`func (o *DeserializationContext) GetBase64VariantOk() (*Base64Variant, bool)`

GetBase64VariantOk returns a tuple with the Base64Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Variant

`func (o *DeserializationContext) SetBase64Variant(v Base64Variant)`

SetBase64Variant sets Base64Variant field to given value.

### HasBase64Variant

`func (o *DeserializationContext) HasBase64Variant() bool`

HasBase64Variant returns a boolean if a field has been set.

### GetArrayBuilders

`func (o *DeserializationContext) GetArrayBuilders() ArrayBuilders`

GetArrayBuilders returns the ArrayBuilders field if non-nil, zero value otherwise.

### GetArrayBuildersOk

`func (o *DeserializationContext) GetArrayBuildersOk() (*ArrayBuilders, bool)`

GetArrayBuildersOk returns a tuple with the ArrayBuilders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayBuilders

`func (o *DeserializationContext) SetArrayBuilders(v ArrayBuilders)`

SetArrayBuilders sets ArrayBuilders field to given value.

### HasArrayBuilders

`func (o *DeserializationContext) HasArrayBuilders() bool`

HasArrayBuilders returns a boolean if a field has been set.

### GetTimeZone

`func (o *DeserializationContext) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *DeserializationContext) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *DeserializationContext) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *DeserializationContext) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetLocale

`func (o *DeserializationContext) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *DeserializationContext) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *DeserializationContext) SetLocale(v Locale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *DeserializationContext) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetFactory

`func (o *DeserializationContext) GetFactory() map[string]interface{}`

GetFactory returns the Factory field if non-nil, zero value otherwise.

### GetFactoryOk

`func (o *DeserializationContext) GetFactoryOk() (*map[string]interface{}, bool)`

GetFactoryOk returns a tuple with the Factory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactory

`func (o *DeserializationContext) SetFactory(v map[string]interface{})`

SetFactory sets Factory field to given value.

### HasFactory

`func (o *DeserializationContext) HasFactory() bool`

HasFactory returns a boolean if a field has been set.

### GetConfig

`func (o *DeserializationContext) GetConfig() DeserializationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DeserializationContext) GetConfigOk() (*DeserializationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DeserializationContext) SetConfig(v DeserializationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DeserializationContext) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetParser

`func (o *DeserializationContext) GetParser() JsonParser`

GetParser returns the Parser field if non-nil, zero value otherwise.

### GetParserOk

`func (o *DeserializationContext) GetParserOk() (*JsonParser, bool)`

GetParserOk returns a tuple with the Parser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParser

`func (o *DeserializationContext) SetParser(v JsonParser)`

SetParser sets Parser field to given value.

### HasParser

`func (o *DeserializationContext) HasParser() bool`

HasParser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


