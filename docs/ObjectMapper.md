# ObjectMapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonFactory** | Pointer to [**JsonFactory**](JsonFactory.md) |  | [optional] 
**SerializationConfig** | Pointer to [**SerializationConfig**](SerializationConfig.md) |  | [optional] 
**DeserializationConfig** | Pointer to [**DeserializationConfig**](DeserializationConfig.md) |  | [optional] 
**DeserializationContext** | Pointer to [**DeserializationContext**](DeserializationContext.md) |  | [optional] 
**SerializerFactory** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializerProvider** | Pointer to [**SerializerProvider**](SerializerProvider.md) |  | [optional] 
**SerializerProviderInstance** | Pointer to [**SerializerProvider**](SerializerProvider.md) |  | [optional] 
**VisibilityChecker** | Pointer to **map[string]interface{}** |  | [optional] 
**SubtypeResolver** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyNamingStrategy** | Pointer to **map[string]interface{}** |  | [optional] 
**TypeFactory** | Pointer to [**TypeFactory**](TypeFactory.md) |  | [optional] 
**NodeFactory** | Pointer to **map[string]interface{}** |  | [optional] 
**DateFormat** | Pointer to [**DateFormat**](DateFormat.md) |  | [optional] 
**InjectableValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Factory** | Pointer to [**JsonFactory**](JsonFactory.md) |  | [optional] 

## Methods

### NewObjectMapper

`func NewObjectMapper() *ObjectMapper`

NewObjectMapper instantiates a new ObjectMapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMapperWithDefaults

`func NewObjectMapperWithDefaults() *ObjectMapper`

NewObjectMapperWithDefaults instantiates a new ObjectMapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonFactory

`func (o *ObjectMapper) GetJsonFactory() JsonFactory`

GetJsonFactory returns the JsonFactory field if non-nil, zero value otherwise.

### GetJsonFactoryOk

`func (o *ObjectMapper) GetJsonFactoryOk() (*JsonFactory, bool)`

GetJsonFactoryOk returns a tuple with the JsonFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonFactory

`func (o *ObjectMapper) SetJsonFactory(v JsonFactory)`

SetJsonFactory sets JsonFactory field to given value.

### HasJsonFactory

`func (o *ObjectMapper) HasJsonFactory() bool`

HasJsonFactory returns a boolean if a field has been set.

### GetSerializationConfig

`func (o *ObjectMapper) GetSerializationConfig() SerializationConfig`

GetSerializationConfig returns the SerializationConfig field if non-nil, zero value otherwise.

### GetSerializationConfigOk

`func (o *ObjectMapper) GetSerializationConfigOk() (*SerializationConfig, bool)`

GetSerializationConfigOk returns a tuple with the SerializationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationConfig

`func (o *ObjectMapper) SetSerializationConfig(v SerializationConfig)`

SetSerializationConfig sets SerializationConfig field to given value.

### HasSerializationConfig

`func (o *ObjectMapper) HasSerializationConfig() bool`

HasSerializationConfig returns a boolean if a field has been set.

### GetDeserializationConfig

`func (o *ObjectMapper) GetDeserializationConfig() DeserializationConfig`

GetDeserializationConfig returns the DeserializationConfig field if non-nil, zero value otherwise.

### GetDeserializationConfigOk

`func (o *ObjectMapper) GetDeserializationConfigOk() (*DeserializationConfig, bool)`

GetDeserializationConfigOk returns a tuple with the DeserializationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeserializationConfig

`func (o *ObjectMapper) SetDeserializationConfig(v DeserializationConfig)`

SetDeserializationConfig sets DeserializationConfig field to given value.

### HasDeserializationConfig

`func (o *ObjectMapper) HasDeserializationConfig() bool`

HasDeserializationConfig returns a boolean if a field has been set.

### GetDeserializationContext

`func (o *ObjectMapper) GetDeserializationContext() DeserializationContext`

GetDeserializationContext returns the DeserializationContext field if non-nil, zero value otherwise.

### GetDeserializationContextOk

`func (o *ObjectMapper) GetDeserializationContextOk() (*DeserializationContext, bool)`

GetDeserializationContextOk returns a tuple with the DeserializationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeserializationContext

`func (o *ObjectMapper) SetDeserializationContext(v DeserializationContext)`

SetDeserializationContext sets DeserializationContext field to given value.

### HasDeserializationContext

`func (o *ObjectMapper) HasDeserializationContext() bool`

HasDeserializationContext returns a boolean if a field has been set.

### GetSerializerFactory

`func (o *ObjectMapper) GetSerializerFactory() map[string]interface{}`

GetSerializerFactory returns the SerializerFactory field if non-nil, zero value otherwise.

### GetSerializerFactoryOk

`func (o *ObjectMapper) GetSerializerFactoryOk() (*map[string]interface{}, bool)`

GetSerializerFactoryOk returns a tuple with the SerializerFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializerFactory

`func (o *ObjectMapper) SetSerializerFactory(v map[string]interface{})`

SetSerializerFactory sets SerializerFactory field to given value.

### HasSerializerFactory

`func (o *ObjectMapper) HasSerializerFactory() bool`

HasSerializerFactory returns a boolean if a field has been set.

### GetSerializerProvider

`func (o *ObjectMapper) GetSerializerProvider() SerializerProvider`

GetSerializerProvider returns the SerializerProvider field if non-nil, zero value otherwise.

### GetSerializerProviderOk

`func (o *ObjectMapper) GetSerializerProviderOk() (*SerializerProvider, bool)`

GetSerializerProviderOk returns a tuple with the SerializerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializerProvider

`func (o *ObjectMapper) SetSerializerProvider(v SerializerProvider)`

SetSerializerProvider sets SerializerProvider field to given value.

### HasSerializerProvider

`func (o *ObjectMapper) HasSerializerProvider() bool`

HasSerializerProvider returns a boolean if a field has been set.

### GetSerializerProviderInstance

`func (o *ObjectMapper) GetSerializerProviderInstance() SerializerProvider`

GetSerializerProviderInstance returns the SerializerProviderInstance field if non-nil, zero value otherwise.

### GetSerializerProviderInstanceOk

`func (o *ObjectMapper) GetSerializerProviderInstanceOk() (*SerializerProvider, bool)`

GetSerializerProviderInstanceOk returns a tuple with the SerializerProviderInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializerProviderInstance

`func (o *ObjectMapper) SetSerializerProviderInstance(v SerializerProvider)`

SetSerializerProviderInstance sets SerializerProviderInstance field to given value.

### HasSerializerProviderInstance

`func (o *ObjectMapper) HasSerializerProviderInstance() bool`

HasSerializerProviderInstance returns a boolean if a field has been set.

### GetVisibilityChecker

`func (o *ObjectMapper) GetVisibilityChecker() map[string]interface{}`

GetVisibilityChecker returns the VisibilityChecker field if non-nil, zero value otherwise.

### GetVisibilityCheckerOk

`func (o *ObjectMapper) GetVisibilityCheckerOk() (*map[string]interface{}, bool)`

GetVisibilityCheckerOk returns a tuple with the VisibilityChecker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityChecker

`func (o *ObjectMapper) SetVisibilityChecker(v map[string]interface{})`

SetVisibilityChecker sets VisibilityChecker field to given value.

### HasVisibilityChecker

`func (o *ObjectMapper) HasVisibilityChecker() bool`

HasVisibilityChecker returns a boolean if a field has been set.

### GetSubtypeResolver

`func (o *ObjectMapper) GetSubtypeResolver() map[string]interface{}`

GetSubtypeResolver returns the SubtypeResolver field if non-nil, zero value otherwise.

### GetSubtypeResolverOk

`func (o *ObjectMapper) GetSubtypeResolverOk() (*map[string]interface{}, bool)`

GetSubtypeResolverOk returns a tuple with the SubtypeResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtypeResolver

`func (o *ObjectMapper) SetSubtypeResolver(v map[string]interface{})`

SetSubtypeResolver sets SubtypeResolver field to given value.

### HasSubtypeResolver

`func (o *ObjectMapper) HasSubtypeResolver() bool`

HasSubtypeResolver returns a boolean if a field has been set.

### GetPropertyNamingStrategy

`func (o *ObjectMapper) GetPropertyNamingStrategy() map[string]interface{}`

GetPropertyNamingStrategy returns the PropertyNamingStrategy field if non-nil, zero value otherwise.

### GetPropertyNamingStrategyOk

`func (o *ObjectMapper) GetPropertyNamingStrategyOk() (*map[string]interface{}, bool)`

GetPropertyNamingStrategyOk returns a tuple with the PropertyNamingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNamingStrategy

`func (o *ObjectMapper) SetPropertyNamingStrategy(v map[string]interface{})`

SetPropertyNamingStrategy sets PropertyNamingStrategy field to given value.

### HasPropertyNamingStrategy

`func (o *ObjectMapper) HasPropertyNamingStrategy() bool`

HasPropertyNamingStrategy returns a boolean if a field has been set.

### GetTypeFactory

`func (o *ObjectMapper) GetTypeFactory() TypeFactory`

GetTypeFactory returns the TypeFactory field if non-nil, zero value otherwise.

### GetTypeFactoryOk

`func (o *ObjectMapper) GetTypeFactoryOk() (*TypeFactory, bool)`

GetTypeFactoryOk returns a tuple with the TypeFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFactory

`func (o *ObjectMapper) SetTypeFactory(v TypeFactory)`

SetTypeFactory sets TypeFactory field to given value.

### HasTypeFactory

`func (o *ObjectMapper) HasTypeFactory() bool`

HasTypeFactory returns a boolean if a field has been set.

### GetNodeFactory

`func (o *ObjectMapper) GetNodeFactory() map[string]interface{}`

GetNodeFactory returns the NodeFactory field if non-nil, zero value otherwise.

### GetNodeFactoryOk

`func (o *ObjectMapper) GetNodeFactoryOk() (*map[string]interface{}, bool)`

GetNodeFactoryOk returns a tuple with the NodeFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFactory

`func (o *ObjectMapper) SetNodeFactory(v map[string]interface{})`

SetNodeFactory sets NodeFactory field to given value.

### HasNodeFactory

`func (o *ObjectMapper) HasNodeFactory() bool`

HasNodeFactory returns a boolean if a field has been set.

### GetDateFormat

`func (o *ObjectMapper) GetDateFormat() DateFormat`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *ObjectMapper) GetDateFormatOk() (*DateFormat, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *ObjectMapper) SetDateFormat(v DateFormat)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *ObjectMapper) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetInjectableValues

`func (o *ObjectMapper) GetInjectableValues() map[string]interface{}`

GetInjectableValues returns the InjectableValues field if non-nil, zero value otherwise.

### GetInjectableValuesOk

`func (o *ObjectMapper) GetInjectableValuesOk() (*map[string]interface{}, bool)`

GetInjectableValuesOk returns a tuple with the InjectableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectableValues

`func (o *ObjectMapper) SetInjectableValues(v map[string]interface{})`

SetInjectableValues sets InjectableValues field to given value.

### HasInjectableValues

`func (o *ObjectMapper) HasInjectableValues() bool`

HasInjectableValues returns a boolean if a field has been set.

### GetFactory

`func (o *ObjectMapper) GetFactory() JsonFactory`

GetFactory returns the Factory field if non-nil, zero value otherwise.

### GetFactoryOk

`func (o *ObjectMapper) GetFactoryOk() (*JsonFactory, bool)`

GetFactoryOk returns a tuple with the Factory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactory

`func (o *ObjectMapper) SetFactory(v JsonFactory)`

SetFactory sets Factory field to given value.

### HasFactory

`func (o *ObjectMapper) HasFactory() bool`

HasFactory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


