# DeserializationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProblemHandlers** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultPropertyInclusion** | Pointer to [**Value**](Value.md) |  | [optional] 
**NodeFactory** | Pointer to **map[string]interface{}** |  | [optional] 
**AnnotationIntrospector** | Pointer to **map[string]interface{}** |  | [optional] 
**DeserializationFeatures** | Pointer to **int32** |  | [optional] 
**FullRootName** | Pointer to [**PropertyName**](PropertyName.md) |  | [optional] 
**SubtypeResolver** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultVisibilityChecker** | Pointer to **map[string]interface{}** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**RootName** | Pointer to **string** |  | [optional] 
**PropertyNamingStrategy** | Pointer to **map[string]interface{}** |  | [optional] 
**TypeFactory** | Pointer to [**TypeFactory**](TypeFactory.md) |  | [optional] 
**DateFormat** | Pointer to [**DateFormat**](DateFormat.md) |  | [optional] 
**Base64Variant** | Pointer to [**Base64Variant**](Base64Variant.md) |  | [optional] 
**HandlerInstantiator** | Pointer to **map[string]interface{}** |  | [optional] 
**AnnotationProcessingEnabled** | Pointer to **bool** |  | [optional] 
**ClassIntrospector** | Pointer to **map[string]interface{}** |  | [optional] 
**TimeZone** | Pointer to [**TimeZone**](TimeZone.md) |  | [optional] 
**Locale** | Pointer to [**Locale**](Locale.md) |  | [optional] 

## Methods

### NewDeserializationConfig

`func NewDeserializationConfig() *DeserializationConfig`

NewDeserializationConfig instantiates a new DeserializationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeserializationConfigWithDefaults

`func NewDeserializationConfigWithDefaults() *DeserializationConfig`

NewDeserializationConfigWithDefaults instantiates a new DeserializationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProblemHandlers

`func (o *DeserializationConfig) GetProblemHandlers() map[string]interface{}`

GetProblemHandlers returns the ProblemHandlers field if non-nil, zero value otherwise.

### GetProblemHandlersOk

`func (o *DeserializationConfig) GetProblemHandlersOk() (*map[string]interface{}, bool)`

GetProblemHandlersOk returns a tuple with the ProblemHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemHandlers

`func (o *DeserializationConfig) SetProblemHandlers(v map[string]interface{})`

SetProblemHandlers sets ProblemHandlers field to given value.

### HasProblemHandlers

`func (o *DeserializationConfig) HasProblemHandlers() bool`

HasProblemHandlers returns a boolean if a field has been set.

### GetDefaultPropertyInclusion

`func (o *DeserializationConfig) GetDefaultPropertyInclusion() Value`

GetDefaultPropertyInclusion returns the DefaultPropertyInclusion field if non-nil, zero value otherwise.

### GetDefaultPropertyInclusionOk

`func (o *DeserializationConfig) GetDefaultPropertyInclusionOk() (*Value, bool)`

GetDefaultPropertyInclusionOk returns a tuple with the DefaultPropertyInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPropertyInclusion

`func (o *DeserializationConfig) SetDefaultPropertyInclusion(v Value)`

SetDefaultPropertyInclusion sets DefaultPropertyInclusion field to given value.

### HasDefaultPropertyInclusion

`func (o *DeserializationConfig) HasDefaultPropertyInclusion() bool`

HasDefaultPropertyInclusion returns a boolean if a field has been set.

### GetNodeFactory

`func (o *DeserializationConfig) GetNodeFactory() map[string]interface{}`

GetNodeFactory returns the NodeFactory field if non-nil, zero value otherwise.

### GetNodeFactoryOk

`func (o *DeserializationConfig) GetNodeFactoryOk() (*map[string]interface{}, bool)`

GetNodeFactoryOk returns a tuple with the NodeFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFactory

`func (o *DeserializationConfig) SetNodeFactory(v map[string]interface{})`

SetNodeFactory sets NodeFactory field to given value.

### HasNodeFactory

`func (o *DeserializationConfig) HasNodeFactory() bool`

HasNodeFactory returns a boolean if a field has been set.

### GetAnnotationIntrospector

`func (o *DeserializationConfig) GetAnnotationIntrospector() map[string]interface{}`

GetAnnotationIntrospector returns the AnnotationIntrospector field if non-nil, zero value otherwise.

### GetAnnotationIntrospectorOk

`func (o *DeserializationConfig) GetAnnotationIntrospectorOk() (*map[string]interface{}, bool)`

GetAnnotationIntrospectorOk returns a tuple with the AnnotationIntrospector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationIntrospector

`func (o *DeserializationConfig) SetAnnotationIntrospector(v map[string]interface{})`

SetAnnotationIntrospector sets AnnotationIntrospector field to given value.

### HasAnnotationIntrospector

`func (o *DeserializationConfig) HasAnnotationIntrospector() bool`

HasAnnotationIntrospector returns a boolean if a field has been set.

### GetDeserializationFeatures

`func (o *DeserializationConfig) GetDeserializationFeatures() int32`

GetDeserializationFeatures returns the DeserializationFeatures field if non-nil, zero value otherwise.

### GetDeserializationFeaturesOk

`func (o *DeserializationConfig) GetDeserializationFeaturesOk() (*int32, bool)`

GetDeserializationFeaturesOk returns a tuple with the DeserializationFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeserializationFeatures

`func (o *DeserializationConfig) SetDeserializationFeatures(v int32)`

SetDeserializationFeatures sets DeserializationFeatures field to given value.

### HasDeserializationFeatures

`func (o *DeserializationConfig) HasDeserializationFeatures() bool`

HasDeserializationFeatures returns a boolean if a field has been set.

### GetFullRootName

`func (o *DeserializationConfig) GetFullRootName() PropertyName`

GetFullRootName returns the FullRootName field if non-nil, zero value otherwise.

### GetFullRootNameOk

`func (o *DeserializationConfig) GetFullRootNameOk() (*PropertyName, bool)`

GetFullRootNameOk returns a tuple with the FullRootName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRootName

`func (o *DeserializationConfig) SetFullRootName(v PropertyName)`

SetFullRootName sets FullRootName field to given value.

### HasFullRootName

`func (o *DeserializationConfig) HasFullRootName() bool`

HasFullRootName returns a boolean if a field has been set.

### GetSubtypeResolver

`func (o *DeserializationConfig) GetSubtypeResolver() map[string]interface{}`

GetSubtypeResolver returns the SubtypeResolver field if non-nil, zero value otherwise.

### GetSubtypeResolverOk

`func (o *DeserializationConfig) GetSubtypeResolverOk() (*map[string]interface{}, bool)`

GetSubtypeResolverOk returns a tuple with the SubtypeResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtypeResolver

`func (o *DeserializationConfig) SetSubtypeResolver(v map[string]interface{})`

SetSubtypeResolver sets SubtypeResolver field to given value.

### HasSubtypeResolver

`func (o *DeserializationConfig) HasSubtypeResolver() bool`

HasSubtypeResolver returns a boolean if a field has been set.

### GetDefaultVisibilityChecker

`func (o *DeserializationConfig) GetDefaultVisibilityChecker() map[string]interface{}`

GetDefaultVisibilityChecker returns the DefaultVisibilityChecker field if non-nil, zero value otherwise.

### GetDefaultVisibilityCheckerOk

`func (o *DeserializationConfig) GetDefaultVisibilityCheckerOk() (*map[string]interface{}, bool)`

GetDefaultVisibilityCheckerOk returns a tuple with the DefaultVisibilityChecker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVisibilityChecker

`func (o *DeserializationConfig) SetDefaultVisibilityChecker(v map[string]interface{})`

SetDefaultVisibilityChecker sets DefaultVisibilityChecker field to given value.

### HasDefaultVisibilityChecker

`func (o *DeserializationConfig) HasDefaultVisibilityChecker() bool`

HasDefaultVisibilityChecker returns a boolean if a field has been set.

### GetAttributes

`func (o *DeserializationConfig) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeserializationConfig) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeserializationConfig) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DeserializationConfig) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRootName

`func (o *DeserializationConfig) GetRootName() string`

GetRootName returns the RootName field if non-nil, zero value otherwise.

### GetRootNameOk

`func (o *DeserializationConfig) GetRootNameOk() (*string, bool)`

GetRootNameOk returns a tuple with the RootName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootName

`func (o *DeserializationConfig) SetRootName(v string)`

SetRootName sets RootName field to given value.

### HasRootName

`func (o *DeserializationConfig) HasRootName() bool`

HasRootName returns a boolean if a field has been set.

### GetPropertyNamingStrategy

`func (o *DeserializationConfig) GetPropertyNamingStrategy() map[string]interface{}`

GetPropertyNamingStrategy returns the PropertyNamingStrategy field if non-nil, zero value otherwise.

### GetPropertyNamingStrategyOk

`func (o *DeserializationConfig) GetPropertyNamingStrategyOk() (*map[string]interface{}, bool)`

GetPropertyNamingStrategyOk returns a tuple with the PropertyNamingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNamingStrategy

`func (o *DeserializationConfig) SetPropertyNamingStrategy(v map[string]interface{})`

SetPropertyNamingStrategy sets PropertyNamingStrategy field to given value.

### HasPropertyNamingStrategy

`func (o *DeserializationConfig) HasPropertyNamingStrategy() bool`

HasPropertyNamingStrategy returns a boolean if a field has been set.

### GetTypeFactory

`func (o *DeserializationConfig) GetTypeFactory() TypeFactory`

GetTypeFactory returns the TypeFactory field if non-nil, zero value otherwise.

### GetTypeFactoryOk

`func (o *DeserializationConfig) GetTypeFactoryOk() (*TypeFactory, bool)`

GetTypeFactoryOk returns a tuple with the TypeFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFactory

`func (o *DeserializationConfig) SetTypeFactory(v TypeFactory)`

SetTypeFactory sets TypeFactory field to given value.

### HasTypeFactory

`func (o *DeserializationConfig) HasTypeFactory() bool`

HasTypeFactory returns a boolean if a field has been set.

### GetDateFormat

`func (o *DeserializationConfig) GetDateFormat() DateFormat`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *DeserializationConfig) GetDateFormatOk() (*DateFormat, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *DeserializationConfig) SetDateFormat(v DateFormat)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *DeserializationConfig) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetBase64Variant

`func (o *DeserializationConfig) GetBase64Variant() Base64Variant`

GetBase64Variant returns the Base64Variant field if non-nil, zero value otherwise.

### GetBase64VariantOk

`func (o *DeserializationConfig) GetBase64VariantOk() (*Base64Variant, bool)`

GetBase64VariantOk returns a tuple with the Base64Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Variant

`func (o *DeserializationConfig) SetBase64Variant(v Base64Variant)`

SetBase64Variant sets Base64Variant field to given value.

### HasBase64Variant

`func (o *DeserializationConfig) HasBase64Variant() bool`

HasBase64Variant returns a boolean if a field has been set.

### GetHandlerInstantiator

`func (o *DeserializationConfig) GetHandlerInstantiator() map[string]interface{}`

GetHandlerInstantiator returns the HandlerInstantiator field if non-nil, zero value otherwise.

### GetHandlerInstantiatorOk

`func (o *DeserializationConfig) GetHandlerInstantiatorOk() (*map[string]interface{}, bool)`

GetHandlerInstantiatorOk returns a tuple with the HandlerInstantiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerInstantiator

`func (o *DeserializationConfig) SetHandlerInstantiator(v map[string]interface{})`

SetHandlerInstantiator sets HandlerInstantiator field to given value.

### HasHandlerInstantiator

`func (o *DeserializationConfig) HasHandlerInstantiator() bool`

HasHandlerInstantiator returns a boolean if a field has been set.

### GetAnnotationProcessingEnabled

`func (o *DeserializationConfig) GetAnnotationProcessingEnabled() bool`

GetAnnotationProcessingEnabled returns the AnnotationProcessingEnabled field if non-nil, zero value otherwise.

### GetAnnotationProcessingEnabledOk

`func (o *DeserializationConfig) GetAnnotationProcessingEnabledOk() (*bool, bool)`

GetAnnotationProcessingEnabledOk returns a tuple with the AnnotationProcessingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationProcessingEnabled

`func (o *DeserializationConfig) SetAnnotationProcessingEnabled(v bool)`

SetAnnotationProcessingEnabled sets AnnotationProcessingEnabled field to given value.

### HasAnnotationProcessingEnabled

`func (o *DeserializationConfig) HasAnnotationProcessingEnabled() bool`

HasAnnotationProcessingEnabled returns a boolean if a field has been set.

### GetClassIntrospector

`func (o *DeserializationConfig) GetClassIntrospector() map[string]interface{}`

GetClassIntrospector returns the ClassIntrospector field if non-nil, zero value otherwise.

### GetClassIntrospectorOk

`func (o *DeserializationConfig) GetClassIntrospectorOk() (*map[string]interface{}, bool)`

GetClassIntrospectorOk returns a tuple with the ClassIntrospector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassIntrospector

`func (o *DeserializationConfig) SetClassIntrospector(v map[string]interface{})`

SetClassIntrospector sets ClassIntrospector field to given value.

### HasClassIntrospector

`func (o *DeserializationConfig) HasClassIntrospector() bool`

HasClassIntrospector returns a boolean if a field has been set.

### GetTimeZone

`func (o *DeserializationConfig) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *DeserializationConfig) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *DeserializationConfig) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *DeserializationConfig) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetLocale

`func (o *DeserializationConfig) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *DeserializationConfig) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *DeserializationConfig) SetLocale(v Locale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *DeserializationConfig) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


