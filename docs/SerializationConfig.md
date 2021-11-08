# SerializationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPropertyInclusion** | Pointer to [**Value**](Value.md) |  | [optional] 
**FilterProvider** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultPrettyPrinter** | Pointer to **map[string]interface{}** |  | [optional] 
**AnnotationIntrospector** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializationInclusion** | Pointer to **string** |  | [optional] 
**SerializationFeatures** | Pointer to **int32** |  | [optional] 
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

### NewSerializationConfig

`func NewSerializationConfig() *SerializationConfig`

NewSerializationConfig instantiates a new SerializationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSerializationConfigWithDefaults

`func NewSerializationConfigWithDefaults() *SerializationConfig`

NewSerializationConfigWithDefaults instantiates a new SerializationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPropertyInclusion

`func (o *SerializationConfig) GetDefaultPropertyInclusion() Value`

GetDefaultPropertyInclusion returns the DefaultPropertyInclusion field if non-nil, zero value otherwise.

### GetDefaultPropertyInclusionOk

`func (o *SerializationConfig) GetDefaultPropertyInclusionOk() (*Value, bool)`

GetDefaultPropertyInclusionOk returns a tuple with the DefaultPropertyInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPropertyInclusion

`func (o *SerializationConfig) SetDefaultPropertyInclusion(v Value)`

SetDefaultPropertyInclusion sets DefaultPropertyInclusion field to given value.

### HasDefaultPropertyInclusion

`func (o *SerializationConfig) HasDefaultPropertyInclusion() bool`

HasDefaultPropertyInclusion returns a boolean if a field has been set.

### GetFilterProvider

`func (o *SerializationConfig) GetFilterProvider() map[string]interface{}`

GetFilterProvider returns the FilterProvider field if non-nil, zero value otherwise.

### GetFilterProviderOk

`func (o *SerializationConfig) GetFilterProviderOk() (*map[string]interface{}, bool)`

GetFilterProviderOk returns a tuple with the FilterProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterProvider

`func (o *SerializationConfig) SetFilterProvider(v map[string]interface{})`

SetFilterProvider sets FilterProvider field to given value.

### HasFilterProvider

`func (o *SerializationConfig) HasFilterProvider() bool`

HasFilterProvider returns a boolean if a field has been set.

### GetDefaultPrettyPrinter

`func (o *SerializationConfig) GetDefaultPrettyPrinter() map[string]interface{}`

GetDefaultPrettyPrinter returns the DefaultPrettyPrinter field if non-nil, zero value otherwise.

### GetDefaultPrettyPrinterOk

`func (o *SerializationConfig) GetDefaultPrettyPrinterOk() (*map[string]interface{}, bool)`

GetDefaultPrettyPrinterOk returns a tuple with the DefaultPrettyPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrettyPrinter

`func (o *SerializationConfig) SetDefaultPrettyPrinter(v map[string]interface{})`

SetDefaultPrettyPrinter sets DefaultPrettyPrinter field to given value.

### HasDefaultPrettyPrinter

`func (o *SerializationConfig) HasDefaultPrettyPrinter() bool`

HasDefaultPrettyPrinter returns a boolean if a field has been set.

### GetAnnotationIntrospector

`func (o *SerializationConfig) GetAnnotationIntrospector() map[string]interface{}`

GetAnnotationIntrospector returns the AnnotationIntrospector field if non-nil, zero value otherwise.

### GetAnnotationIntrospectorOk

`func (o *SerializationConfig) GetAnnotationIntrospectorOk() (*map[string]interface{}, bool)`

GetAnnotationIntrospectorOk returns a tuple with the AnnotationIntrospector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationIntrospector

`func (o *SerializationConfig) SetAnnotationIntrospector(v map[string]interface{})`

SetAnnotationIntrospector sets AnnotationIntrospector field to given value.

### HasAnnotationIntrospector

`func (o *SerializationConfig) HasAnnotationIntrospector() bool`

HasAnnotationIntrospector returns a boolean if a field has been set.

### GetSerializationInclusion

`func (o *SerializationConfig) GetSerializationInclusion() string`

GetSerializationInclusion returns the SerializationInclusion field if non-nil, zero value otherwise.

### GetSerializationInclusionOk

`func (o *SerializationConfig) GetSerializationInclusionOk() (*string, bool)`

GetSerializationInclusionOk returns a tuple with the SerializationInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationInclusion

`func (o *SerializationConfig) SetSerializationInclusion(v string)`

SetSerializationInclusion sets SerializationInclusion field to given value.

### HasSerializationInclusion

`func (o *SerializationConfig) HasSerializationInclusion() bool`

HasSerializationInclusion returns a boolean if a field has been set.

### GetSerializationFeatures

`func (o *SerializationConfig) GetSerializationFeatures() int32`

GetSerializationFeatures returns the SerializationFeatures field if non-nil, zero value otherwise.

### GetSerializationFeaturesOk

`func (o *SerializationConfig) GetSerializationFeaturesOk() (*int32, bool)`

GetSerializationFeaturesOk returns a tuple with the SerializationFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationFeatures

`func (o *SerializationConfig) SetSerializationFeatures(v int32)`

SetSerializationFeatures sets SerializationFeatures field to given value.

### HasSerializationFeatures

`func (o *SerializationConfig) HasSerializationFeatures() bool`

HasSerializationFeatures returns a boolean if a field has been set.

### GetFullRootName

`func (o *SerializationConfig) GetFullRootName() PropertyName`

GetFullRootName returns the FullRootName field if non-nil, zero value otherwise.

### GetFullRootNameOk

`func (o *SerializationConfig) GetFullRootNameOk() (*PropertyName, bool)`

GetFullRootNameOk returns a tuple with the FullRootName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRootName

`func (o *SerializationConfig) SetFullRootName(v PropertyName)`

SetFullRootName sets FullRootName field to given value.

### HasFullRootName

`func (o *SerializationConfig) HasFullRootName() bool`

HasFullRootName returns a boolean if a field has been set.

### GetSubtypeResolver

`func (o *SerializationConfig) GetSubtypeResolver() map[string]interface{}`

GetSubtypeResolver returns the SubtypeResolver field if non-nil, zero value otherwise.

### GetSubtypeResolverOk

`func (o *SerializationConfig) GetSubtypeResolverOk() (*map[string]interface{}, bool)`

GetSubtypeResolverOk returns a tuple with the SubtypeResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtypeResolver

`func (o *SerializationConfig) SetSubtypeResolver(v map[string]interface{})`

SetSubtypeResolver sets SubtypeResolver field to given value.

### HasSubtypeResolver

`func (o *SerializationConfig) HasSubtypeResolver() bool`

HasSubtypeResolver returns a boolean if a field has been set.

### GetDefaultVisibilityChecker

`func (o *SerializationConfig) GetDefaultVisibilityChecker() map[string]interface{}`

GetDefaultVisibilityChecker returns the DefaultVisibilityChecker field if non-nil, zero value otherwise.

### GetDefaultVisibilityCheckerOk

`func (o *SerializationConfig) GetDefaultVisibilityCheckerOk() (*map[string]interface{}, bool)`

GetDefaultVisibilityCheckerOk returns a tuple with the DefaultVisibilityChecker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVisibilityChecker

`func (o *SerializationConfig) SetDefaultVisibilityChecker(v map[string]interface{})`

SetDefaultVisibilityChecker sets DefaultVisibilityChecker field to given value.

### HasDefaultVisibilityChecker

`func (o *SerializationConfig) HasDefaultVisibilityChecker() bool`

HasDefaultVisibilityChecker returns a boolean if a field has been set.

### GetAttributes

`func (o *SerializationConfig) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SerializationConfig) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SerializationConfig) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SerializationConfig) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRootName

`func (o *SerializationConfig) GetRootName() string`

GetRootName returns the RootName field if non-nil, zero value otherwise.

### GetRootNameOk

`func (o *SerializationConfig) GetRootNameOk() (*string, bool)`

GetRootNameOk returns a tuple with the RootName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootName

`func (o *SerializationConfig) SetRootName(v string)`

SetRootName sets RootName field to given value.

### HasRootName

`func (o *SerializationConfig) HasRootName() bool`

HasRootName returns a boolean if a field has been set.

### GetPropertyNamingStrategy

`func (o *SerializationConfig) GetPropertyNamingStrategy() map[string]interface{}`

GetPropertyNamingStrategy returns the PropertyNamingStrategy field if non-nil, zero value otherwise.

### GetPropertyNamingStrategyOk

`func (o *SerializationConfig) GetPropertyNamingStrategyOk() (*map[string]interface{}, bool)`

GetPropertyNamingStrategyOk returns a tuple with the PropertyNamingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNamingStrategy

`func (o *SerializationConfig) SetPropertyNamingStrategy(v map[string]interface{})`

SetPropertyNamingStrategy sets PropertyNamingStrategy field to given value.

### HasPropertyNamingStrategy

`func (o *SerializationConfig) HasPropertyNamingStrategy() bool`

HasPropertyNamingStrategy returns a boolean if a field has been set.

### GetTypeFactory

`func (o *SerializationConfig) GetTypeFactory() TypeFactory`

GetTypeFactory returns the TypeFactory field if non-nil, zero value otherwise.

### GetTypeFactoryOk

`func (o *SerializationConfig) GetTypeFactoryOk() (*TypeFactory, bool)`

GetTypeFactoryOk returns a tuple with the TypeFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFactory

`func (o *SerializationConfig) SetTypeFactory(v TypeFactory)`

SetTypeFactory sets TypeFactory field to given value.

### HasTypeFactory

`func (o *SerializationConfig) HasTypeFactory() bool`

HasTypeFactory returns a boolean if a field has been set.

### GetDateFormat

`func (o *SerializationConfig) GetDateFormat() DateFormat`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *SerializationConfig) GetDateFormatOk() (*DateFormat, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *SerializationConfig) SetDateFormat(v DateFormat)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *SerializationConfig) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetBase64Variant

`func (o *SerializationConfig) GetBase64Variant() Base64Variant`

GetBase64Variant returns the Base64Variant field if non-nil, zero value otherwise.

### GetBase64VariantOk

`func (o *SerializationConfig) GetBase64VariantOk() (*Base64Variant, bool)`

GetBase64VariantOk returns a tuple with the Base64Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Variant

`func (o *SerializationConfig) SetBase64Variant(v Base64Variant)`

SetBase64Variant sets Base64Variant field to given value.

### HasBase64Variant

`func (o *SerializationConfig) HasBase64Variant() bool`

HasBase64Variant returns a boolean if a field has been set.

### GetHandlerInstantiator

`func (o *SerializationConfig) GetHandlerInstantiator() map[string]interface{}`

GetHandlerInstantiator returns the HandlerInstantiator field if non-nil, zero value otherwise.

### GetHandlerInstantiatorOk

`func (o *SerializationConfig) GetHandlerInstantiatorOk() (*map[string]interface{}, bool)`

GetHandlerInstantiatorOk returns a tuple with the HandlerInstantiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerInstantiator

`func (o *SerializationConfig) SetHandlerInstantiator(v map[string]interface{})`

SetHandlerInstantiator sets HandlerInstantiator field to given value.

### HasHandlerInstantiator

`func (o *SerializationConfig) HasHandlerInstantiator() bool`

HasHandlerInstantiator returns a boolean if a field has been set.

### GetAnnotationProcessingEnabled

`func (o *SerializationConfig) GetAnnotationProcessingEnabled() bool`

GetAnnotationProcessingEnabled returns the AnnotationProcessingEnabled field if non-nil, zero value otherwise.

### GetAnnotationProcessingEnabledOk

`func (o *SerializationConfig) GetAnnotationProcessingEnabledOk() (*bool, bool)`

GetAnnotationProcessingEnabledOk returns a tuple with the AnnotationProcessingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationProcessingEnabled

`func (o *SerializationConfig) SetAnnotationProcessingEnabled(v bool)`

SetAnnotationProcessingEnabled sets AnnotationProcessingEnabled field to given value.

### HasAnnotationProcessingEnabled

`func (o *SerializationConfig) HasAnnotationProcessingEnabled() bool`

HasAnnotationProcessingEnabled returns a boolean if a field has been set.

### GetClassIntrospector

`func (o *SerializationConfig) GetClassIntrospector() map[string]interface{}`

GetClassIntrospector returns the ClassIntrospector field if non-nil, zero value otherwise.

### GetClassIntrospectorOk

`func (o *SerializationConfig) GetClassIntrospectorOk() (*map[string]interface{}, bool)`

GetClassIntrospectorOk returns a tuple with the ClassIntrospector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassIntrospector

`func (o *SerializationConfig) SetClassIntrospector(v map[string]interface{})`

SetClassIntrospector sets ClassIntrospector field to given value.

### HasClassIntrospector

`func (o *SerializationConfig) HasClassIntrospector() bool`

HasClassIntrospector returns a boolean if a field has been set.

### GetTimeZone

`func (o *SerializationConfig) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SerializationConfig) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SerializationConfig) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SerializationConfig) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetLocale

`func (o *SerializationConfig) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SerializationConfig) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SerializationConfig) SetLocale(v Locale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SerializationConfig) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


