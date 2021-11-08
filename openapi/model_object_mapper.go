/*
Pulsar Admin REST API

This provides the REST API for admin operations

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ObjectMapper struct for ObjectMapper
type ObjectMapper struct {
	JsonFactory *JsonFactory `json:"jsonFactory,omitempty"`
	SerializationConfig *SerializationConfig `json:"serializationConfig,omitempty"`
	DeserializationConfig *DeserializationConfig `json:"deserializationConfig,omitempty"`
	DeserializationContext *DeserializationContext `json:"deserializationContext,omitempty"`
	SerializerFactory *map[string]interface{} `json:"serializerFactory,omitempty"`
	SerializerProvider *SerializerProvider `json:"serializerProvider,omitempty"`
	SerializerProviderInstance *SerializerProvider `json:"serializerProviderInstance,omitempty"`
	VisibilityChecker *map[string]interface{} `json:"visibilityChecker,omitempty"`
	SubtypeResolver *map[string]interface{} `json:"subtypeResolver,omitempty"`
	PropertyNamingStrategy *map[string]interface{} `json:"propertyNamingStrategy,omitempty"`
	TypeFactory *TypeFactory `json:"typeFactory,omitempty"`
	NodeFactory *map[string]interface{} `json:"nodeFactory,omitempty"`
	DateFormat *DateFormat `json:"dateFormat,omitempty"`
	InjectableValues *map[string]interface{} `json:"injectableValues,omitempty"`
	Factory *JsonFactory `json:"factory,omitempty"`
}

// NewObjectMapper instantiates a new ObjectMapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectMapper() *ObjectMapper {
	this := ObjectMapper{}
	return &this
}

// NewObjectMapperWithDefaults instantiates a new ObjectMapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectMapperWithDefaults() *ObjectMapper {
	this := ObjectMapper{}
	return &this
}

// GetJsonFactory returns the JsonFactory field value if set, zero value otherwise.
func (o *ObjectMapper) GetJsonFactory() JsonFactory {
	if o == nil || o.JsonFactory == nil {
		var ret JsonFactory
		return ret
	}
	return *o.JsonFactory
}

// GetJsonFactoryOk returns a tuple with the JsonFactory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetJsonFactoryOk() (*JsonFactory, bool) {
	if o == nil || o.JsonFactory == nil {
		return nil, false
	}
	return o.JsonFactory, true
}

// HasJsonFactory returns a boolean if a field has been set.
func (o *ObjectMapper) HasJsonFactory() bool {
	if o != nil && o.JsonFactory != nil {
		return true
	}

	return false
}

// SetJsonFactory gets a reference to the given JsonFactory and assigns it to the JsonFactory field.
func (o *ObjectMapper) SetJsonFactory(v JsonFactory) {
	o.JsonFactory = &v
}

// GetSerializationConfig returns the SerializationConfig field value if set, zero value otherwise.
func (o *ObjectMapper) GetSerializationConfig() SerializationConfig {
	if o == nil || o.SerializationConfig == nil {
		var ret SerializationConfig
		return ret
	}
	return *o.SerializationConfig
}

// GetSerializationConfigOk returns a tuple with the SerializationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetSerializationConfigOk() (*SerializationConfig, bool) {
	if o == nil || o.SerializationConfig == nil {
		return nil, false
	}
	return o.SerializationConfig, true
}

// HasSerializationConfig returns a boolean if a field has been set.
func (o *ObjectMapper) HasSerializationConfig() bool {
	if o != nil && o.SerializationConfig != nil {
		return true
	}

	return false
}

// SetSerializationConfig gets a reference to the given SerializationConfig and assigns it to the SerializationConfig field.
func (o *ObjectMapper) SetSerializationConfig(v SerializationConfig) {
	o.SerializationConfig = &v
}

// GetDeserializationConfig returns the DeserializationConfig field value if set, zero value otherwise.
func (o *ObjectMapper) GetDeserializationConfig() DeserializationConfig {
	if o == nil || o.DeserializationConfig == nil {
		var ret DeserializationConfig
		return ret
	}
	return *o.DeserializationConfig
}

// GetDeserializationConfigOk returns a tuple with the DeserializationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetDeserializationConfigOk() (*DeserializationConfig, bool) {
	if o == nil || o.DeserializationConfig == nil {
		return nil, false
	}
	return o.DeserializationConfig, true
}

// HasDeserializationConfig returns a boolean if a field has been set.
func (o *ObjectMapper) HasDeserializationConfig() bool {
	if o != nil && o.DeserializationConfig != nil {
		return true
	}

	return false
}

// SetDeserializationConfig gets a reference to the given DeserializationConfig and assigns it to the DeserializationConfig field.
func (o *ObjectMapper) SetDeserializationConfig(v DeserializationConfig) {
	o.DeserializationConfig = &v
}

// GetDeserializationContext returns the DeserializationContext field value if set, zero value otherwise.
func (o *ObjectMapper) GetDeserializationContext() DeserializationContext {
	if o == nil || o.DeserializationContext == nil {
		var ret DeserializationContext
		return ret
	}
	return *o.DeserializationContext
}

// GetDeserializationContextOk returns a tuple with the DeserializationContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetDeserializationContextOk() (*DeserializationContext, bool) {
	if o == nil || o.DeserializationContext == nil {
		return nil, false
	}
	return o.DeserializationContext, true
}

// HasDeserializationContext returns a boolean if a field has been set.
func (o *ObjectMapper) HasDeserializationContext() bool {
	if o != nil && o.DeserializationContext != nil {
		return true
	}

	return false
}

// SetDeserializationContext gets a reference to the given DeserializationContext and assigns it to the DeserializationContext field.
func (o *ObjectMapper) SetDeserializationContext(v DeserializationContext) {
	o.DeserializationContext = &v
}

// GetSerializerFactory returns the SerializerFactory field value if set, zero value otherwise.
func (o *ObjectMapper) GetSerializerFactory() map[string]interface{} {
	if o == nil || o.SerializerFactory == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SerializerFactory
}

// GetSerializerFactoryOk returns a tuple with the SerializerFactory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetSerializerFactoryOk() (*map[string]interface{}, bool) {
	if o == nil || o.SerializerFactory == nil {
		return nil, false
	}
	return o.SerializerFactory, true
}

// HasSerializerFactory returns a boolean if a field has been set.
func (o *ObjectMapper) HasSerializerFactory() bool {
	if o != nil && o.SerializerFactory != nil {
		return true
	}

	return false
}

// SetSerializerFactory gets a reference to the given map[string]interface{} and assigns it to the SerializerFactory field.
func (o *ObjectMapper) SetSerializerFactory(v map[string]interface{}) {
	o.SerializerFactory = &v
}

// GetSerializerProvider returns the SerializerProvider field value if set, zero value otherwise.
func (o *ObjectMapper) GetSerializerProvider() SerializerProvider {
	if o == nil || o.SerializerProvider == nil {
		var ret SerializerProvider
		return ret
	}
	return *o.SerializerProvider
}

// GetSerializerProviderOk returns a tuple with the SerializerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetSerializerProviderOk() (*SerializerProvider, bool) {
	if o == nil || o.SerializerProvider == nil {
		return nil, false
	}
	return o.SerializerProvider, true
}

// HasSerializerProvider returns a boolean if a field has been set.
func (o *ObjectMapper) HasSerializerProvider() bool {
	if o != nil && o.SerializerProvider != nil {
		return true
	}

	return false
}

// SetSerializerProvider gets a reference to the given SerializerProvider and assigns it to the SerializerProvider field.
func (o *ObjectMapper) SetSerializerProvider(v SerializerProvider) {
	o.SerializerProvider = &v
}

// GetSerializerProviderInstance returns the SerializerProviderInstance field value if set, zero value otherwise.
func (o *ObjectMapper) GetSerializerProviderInstance() SerializerProvider {
	if o == nil || o.SerializerProviderInstance == nil {
		var ret SerializerProvider
		return ret
	}
	return *o.SerializerProviderInstance
}

// GetSerializerProviderInstanceOk returns a tuple with the SerializerProviderInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetSerializerProviderInstanceOk() (*SerializerProvider, bool) {
	if o == nil || o.SerializerProviderInstance == nil {
		return nil, false
	}
	return o.SerializerProviderInstance, true
}

// HasSerializerProviderInstance returns a boolean if a field has been set.
func (o *ObjectMapper) HasSerializerProviderInstance() bool {
	if o != nil && o.SerializerProviderInstance != nil {
		return true
	}

	return false
}

// SetSerializerProviderInstance gets a reference to the given SerializerProvider and assigns it to the SerializerProviderInstance field.
func (o *ObjectMapper) SetSerializerProviderInstance(v SerializerProvider) {
	o.SerializerProviderInstance = &v
}

// GetVisibilityChecker returns the VisibilityChecker field value if set, zero value otherwise.
func (o *ObjectMapper) GetVisibilityChecker() map[string]interface{} {
	if o == nil || o.VisibilityChecker == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.VisibilityChecker
}

// GetVisibilityCheckerOk returns a tuple with the VisibilityChecker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetVisibilityCheckerOk() (*map[string]interface{}, bool) {
	if o == nil || o.VisibilityChecker == nil {
		return nil, false
	}
	return o.VisibilityChecker, true
}

// HasVisibilityChecker returns a boolean if a field has been set.
func (o *ObjectMapper) HasVisibilityChecker() bool {
	if o != nil && o.VisibilityChecker != nil {
		return true
	}

	return false
}

// SetVisibilityChecker gets a reference to the given map[string]interface{} and assigns it to the VisibilityChecker field.
func (o *ObjectMapper) SetVisibilityChecker(v map[string]interface{}) {
	o.VisibilityChecker = &v
}

// GetSubtypeResolver returns the SubtypeResolver field value if set, zero value otherwise.
func (o *ObjectMapper) GetSubtypeResolver() map[string]interface{} {
	if o == nil || o.SubtypeResolver == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SubtypeResolver
}

// GetSubtypeResolverOk returns a tuple with the SubtypeResolver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetSubtypeResolverOk() (*map[string]interface{}, bool) {
	if o == nil || o.SubtypeResolver == nil {
		return nil, false
	}
	return o.SubtypeResolver, true
}

// HasSubtypeResolver returns a boolean if a field has been set.
func (o *ObjectMapper) HasSubtypeResolver() bool {
	if o != nil && o.SubtypeResolver != nil {
		return true
	}

	return false
}

// SetSubtypeResolver gets a reference to the given map[string]interface{} and assigns it to the SubtypeResolver field.
func (o *ObjectMapper) SetSubtypeResolver(v map[string]interface{}) {
	o.SubtypeResolver = &v
}

// GetPropertyNamingStrategy returns the PropertyNamingStrategy field value if set, zero value otherwise.
func (o *ObjectMapper) GetPropertyNamingStrategy() map[string]interface{} {
	if o == nil || o.PropertyNamingStrategy == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.PropertyNamingStrategy
}

// GetPropertyNamingStrategyOk returns a tuple with the PropertyNamingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetPropertyNamingStrategyOk() (*map[string]interface{}, bool) {
	if o == nil || o.PropertyNamingStrategy == nil {
		return nil, false
	}
	return o.PropertyNamingStrategy, true
}

// HasPropertyNamingStrategy returns a boolean if a field has been set.
func (o *ObjectMapper) HasPropertyNamingStrategy() bool {
	if o != nil && o.PropertyNamingStrategy != nil {
		return true
	}

	return false
}

// SetPropertyNamingStrategy gets a reference to the given map[string]interface{} and assigns it to the PropertyNamingStrategy field.
func (o *ObjectMapper) SetPropertyNamingStrategy(v map[string]interface{}) {
	o.PropertyNamingStrategy = &v
}

// GetTypeFactory returns the TypeFactory field value if set, zero value otherwise.
func (o *ObjectMapper) GetTypeFactory() TypeFactory {
	if o == nil || o.TypeFactory == nil {
		var ret TypeFactory
		return ret
	}
	return *o.TypeFactory
}

// GetTypeFactoryOk returns a tuple with the TypeFactory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetTypeFactoryOk() (*TypeFactory, bool) {
	if o == nil || o.TypeFactory == nil {
		return nil, false
	}
	return o.TypeFactory, true
}

// HasTypeFactory returns a boolean if a field has been set.
func (o *ObjectMapper) HasTypeFactory() bool {
	if o != nil && o.TypeFactory != nil {
		return true
	}

	return false
}

// SetTypeFactory gets a reference to the given TypeFactory and assigns it to the TypeFactory field.
func (o *ObjectMapper) SetTypeFactory(v TypeFactory) {
	o.TypeFactory = &v
}

// GetNodeFactory returns the NodeFactory field value if set, zero value otherwise.
func (o *ObjectMapper) GetNodeFactory() map[string]interface{} {
	if o == nil || o.NodeFactory == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.NodeFactory
}

// GetNodeFactoryOk returns a tuple with the NodeFactory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetNodeFactoryOk() (*map[string]interface{}, bool) {
	if o == nil || o.NodeFactory == nil {
		return nil, false
	}
	return o.NodeFactory, true
}

// HasNodeFactory returns a boolean if a field has been set.
func (o *ObjectMapper) HasNodeFactory() bool {
	if o != nil && o.NodeFactory != nil {
		return true
	}

	return false
}

// SetNodeFactory gets a reference to the given map[string]interface{} and assigns it to the NodeFactory field.
func (o *ObjectMapper) SetNodeFactory(v map[string]interface{}) {
	o.NodeFactory = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *ObjectMapper) GetDateFormat() DateFormat {
	if o == nil || o.DateFormat == nil {
		var ret DateFormat
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetDateFormatOk() (*DateFormat, bool) {
	if o == nil || o.DateFormat == nil {
		return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *ObjectMapper) HasDateFormat() bool {
	if o != nil && o.DateFormat != nil {
		return true
	}

	return false
}

// SetDateFormat gets a reference to the given DateFormat and assigns it to the DateFormat field.
func (o *ObjectMapper) SetDateFormat(v DateFormat) {
	o.DateFormat = &v
}

// GetInjectableValues returns the InjectableValues field value if set, zero value otherwise.
func (o *ObjectMapper) GetInjectableValues() map[string]interface{} {
	if o == nil || o.InjectableValues == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.InjectableValues
}

// GetInjectableValuesOk returns a tuple with the InjectableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetInjectableValuesOk() (*map[string]interface{}, bool) {
	if o == nil || o.InjectableValues == nil {
		return nil, false
	}
	return o.InjectableValues, true
}

// HasInjectableValues returns a boolean if a field has been set.
func (o *ObjectMapper) HasInjectableValues() bool {
	if o != nil && o.InjectableValues != nil {
		return true
	}

	return false
}

// SetInjectableValues gets a reference to the given map[string]interface{} and assigns it to the InjectableValues field.
func (o *ObjectMapper) SetInjectableValues(v map[string]interface{}) {
	o.InjectableValues = &v
}

// GetFactory returns the Factory field value if set, zero value otherwise.
func (o *ObjectMapper) GetFactory() JsonFactory {
	if o == nil || o.Factory == nil {
		var ret JsonFactory
		return ret
	}
	return *o.Factory
}

// GetFactoryOk returns a tuple with the Factory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMapper) GetFactoryOk() (*JsonFactory, bool) {
	if o == nil || o.Factory == nil {
		return nil, false
	}
	return o.Factory, true
}

// HasFactory returns a boolean if a field has been set.
func (o *ObjectMapper) HasFactory() bool {
	if o != nil && o.Factory != nil {
		return true
	}

	return false
}

// SetFactory gets a reference to the given JsonFactory and assigns it to the Factory field.
func (o *ObjectMapper) SetFactory(v JsonFactory) {
	o.Factory = &v
}

func (o ObjectMapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JsonFactory != nil {
		toSerialize["jsonFactory"] = o.JsonFactory
	}
	if o.SerializationConfig != nil {
		toSerialize["serializationConfig"] = o.SerializationConfig
	}
	if o.DeserializationConfig != nil {
		toSerialize["deserializationConfig"] = o.DeserializationConfig
	}
	if o.DeserializationContext != nil {
		toSerialize["deserializationContext"] = o.DeserializationContext
	}
	if o.SerializerFactory != nil {
		toSerialize["serializerFactory"] = o.SerializerFactory
	}
	if o.SerializerProvider != nil {
		toSerialize["serializerProvider"] = o.SerializerProvider
	}
	if o.SerializerProviderInstance != nil {
		toSerialize["serializerProviderInstance"] = o.SerializerProviderInstance
	}
	if o.VisibilityChecker != nil {
		toSerialize["visibilityChecker"] = o.VisibilityChecker
	}
	if o.SubtypeResolver != nil {
		toSerialize["subtypeResolver"] = o.SubtypeResolver
	}
	if o.PropertyNamingStrategy != nil {
		toSerialize["propertyNamingStrategy"] = o.PropertyNamingStrategy
	}
	if o.TypeFactory != nil {
		toSerialize["typeFactory"] = o.TypeFactory
	}
	if o.NodeFactory != nil {
		toSerialize["nodeFactory"] = o.NodeFactory
	}
	if o.DateFormat != nil {
		toSerialize["dateFormat"] = o.DateFormat
	}
	if o.InjectableValues != nil {
		toSerialize["injectableValues"] = o.InjectableValues
	}
	if o.Factory != nil {
		toSerialize["factory"] = o.Factory
	}
	return json.Marshal(toSerialize)
}

type NullableObjectMapper struct {
	value *ObjectMapper
	isSet bool
}

func (v NullableObjectMapper) Get() *ObjectMapper {
	return v.value
}

func (v *NullableObjectMapper) Set(val *ObjectMapper) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectMapper) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectMapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectMapper(val *ObjectMapper) *NullableObjectMapper {
	return &NullableObjectMapper{value: val, isSet: true}
}

func (v NullableObjectMapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectMapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


