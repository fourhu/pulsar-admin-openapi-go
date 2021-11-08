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

// JsonGenerator struct for JsonGenerator
type JsonGenerator struct {
	Closed *bool `json:"closed,omitempty"`
	Codec *ObjectCodec `json:"codec,omitempty"`
	CharacterEscapes *CharacterEscapes `json:"characterEscapes,omitempty"`
	CurrentValue *map[string]interface{} `json:"currentValue,omitempty"`
	FeatureMask *int32 `json:"featureMask,omitempty"`
	FormatFeatures *int32 `json:"formatFeatures,omitempty"`
	PrettyPrinter *map[string]interface{} `json:"prettyPrinter,omitempty"`
	OutputContext *JsonStreamContext `json:"outputContext,omitempty"`
	OutputTarget *map[string]interface{} `json:"outputTarget,omitempty"`
	HighestEscapedChar *int32 `json:"highestEscapedChar,omitempty"`
	OutputBuffered *int32 `json:"outputBuffered,omitempty"`
	Schema *FormatSchema `json:"schema,omitempty"`
}

// NewJsonGenerator instantiates a new JsonGenerator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonGenerator() *JsonGenerator {
	this := JsonGenerator{}
	return &this
}

// NewJsonGeneratorWithDefaults instantiates a new JsonGenerator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonGeneratorWithDefaults() *JsonGenerator {
	this := JsonGenerator{}
	return &this
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *JsonGenerator) GetClosed() bool {
	if o == nil || o.Closed == nil {
		var ret bool
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetClosedOk() (*bool, bool) {
	if o == nil || o.Closed == nil {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *JsonGenerator) HasClosed() bool {
	if o != nil && o.Closed != nil {
		return true
	}

	return false
}

// SetClosed gets a reference to the given bool and assigns it to the Closed field.
func (o *JsonGenerator) SetClosed(v bool) {
	o.Closed = &v
}

// GetCodec returns the Codec field value if set, zero value otherwise.
func (o *JsonGenerator) GetCodec() ObjectCodec {
	if o == nil || o.Codec == nil {
		var ret ObjectCodec
		return ret
	}
	return *o.Codec
}

// GetCodecOk returns a tuple with the Codec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetCodecOk() (*ObjectCodec, bool) {
	if o == nil || o.Codec == nil {
		return nil, false
	}
	return o.Codec, true
}

// HasCodec returns a boolean if a field has been set.
func (o *JsonGenerator) HasCodec() bool {
	if o != nil && o.Codec != nil {
		return true
	}

	return false
}

// SetCodec gets a reference to the given ObjectCodec and assigns it to the Codec field.
func (o *JsonGenerator) SetCodec(v ObjectCodec) {
	o.Codec = &v
}

// GetCharacterEscapes returns the CharacterEscapes field value if set, zero value otherwise.
func (o *JsonGenerator) GetCharacterEscapes() CharacterEscapes {
	if o == nil || o.CharacterEscapes == nil {
		var ret CharacterEscapes
		return ret
	}
	return *o.CharacterEscapes
}

// GetCharacterEscapesOk returns a tuple with the CharacterEscapes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetCharacterEscapesOk() (*CharacterEscapes, bool) {
	if o == nil || o.CharacterEscapes == nil {
		return nil, false
	}
	return o.CharacterEscapes, true
}

// HasCharacterEscapes returns a boolean if a field has been set.
func (o *JsonGenerator) HasCharacterEscapes() bool {
	if o != nil && o.CharacterEscapes != nil {
		return true
	}

	return false
}

// SetCharacterEscapes gets a reference to the given CharacterEscapes and assigns it to the CharacterEscapes field.
func (o *JsonGenerator) SetCharacterEscapes(v CharacterEscapes) {
	o.CharacterEscapes = &v
}

// GetCurrentValue returns the CurrentValue field value if set, zero value otherwise.
func (o *JsonGenerator) GetCurrentValue() map[string]interface{} {
	if o == nil || o.CurrentValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetCurrentValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.CurrentValue == nil {
		return nil, false
	}
	return o.CurrentValue, true
}

// HasCurrentValue returns a boolean if a field has been set.
func (o *JsonGenerator) HasCurrentValue() bool {
	if o != nil && o.CurrentValue != nil {
		return true
	}

	return false
}

// SetCurrentValue gets a reference to the given map[string]interface{} and assigns it to the CurrentValue field.
func (o *JsonGenerator) SetCurrentValue(v map[string]interface{}) {
	o.CurrentValue = &v
}

// GetFeatureMask returns the FeatureMask field value if set, zero value otherwise.
func (o *JsonGenerator) GetFeatureMask() int32 {
	if o == nil || o.FeatureMask == nil {
		var ret int32
		return ret
	}
	return *o.FeatureMask
}

// GetFeatureMaskOk returns a tuple with the FeatureMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetFeatureMaskOk() (*int32, bool) {
	if o == nil || o.FeatureMask == nil {
		return nil, false
	}
	return o.FeatureMask, true
}

// HasFeatureMask returns a boolean if a field has been set.
func (o *JsonGenerator) HasFeatureMask() bool {
	if o != nil && o.FeatureMask != nil {
		return true
	}

	return false
}

// SetFeatureMask gets a reference to the given int32 and assigns it to the FeatureMask field.
func (o *JsonGenerator) SetFeatureMask(v int32) {
	o.FeatureMask = &v
}

// GetFormatFeatures returns the FormatFeatures field value if set, zero value otherwise.
func (o *JsonGenerator) GetFormatFeatures() int32 {
	if o == nil || o.FormatFeatures == nil {
		var ret int32
		return ret
	}
	return *o.FormatFeatures
}

// GetFormatFeaturesOk returns a tuple with the FormatFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetFormatFeaturesOk() (*int32, bool) {
	if o == nil || o.FormatFeatures == nil {
		return nil, false
	}
	return o.FormatFeatures, true
}

// HasFormatFeatures returns a boolean if a field has been set.
func (o *JsonGenerator) HasFormatFeatures() bool {
	if o != nil && o.FormatFeatures != nil {
		return true
	}

	return false
}

// SetFormatFeatures gets a reference to the given int32 and assigns it to the FormatFeatures field.
func (o *JsonGenerator) SetFormatFeatures(v int32) {
	o.FormatFeatures = &v
}

// GetPrettyPrinter returns the PrettyPrinter field value if set, zero value otherwise.
func (o *JsonGenerator) GetPrettyPrinter() map[string]interface{} {
	if o == nil || o.PrettyPrinter == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.PrettyPrinter
}

// GetPrettyPrinterOk returns a tuple with the PrettyPrinter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetPrettyPrinterOk() (*map[string]interface{}, bool) {
	if o == nil || o.PrettyPrinter == nil {
		return nil, false
	}
	return o.PrettyPrinter, true
}

// HasPrettyPrinter returns a boolean if a field has been set.
func (o *JsonGenerator) HasPrettyPrinter() bool {
	if o != nil && o.PrettyPrinter != nil {
		return true
	}

	return false
}

// SetPrettyPrinter gets a reference to the given map[string]interface{} and assigns it to the PrettyPrinter field.
func (o *JsonGenerator) SetPrettyPrinter(v map[string]interface{}) {
	o.PrettyPrinter = &v
}

// GetOutputContext returns the OutputContext field value if set, zero value otherwise.
func (o *JsonGenerator) GetOutputContext() JsonStreamContext {
	if o == nil || o.OutputContext == nil {
		var ret JsonStreamContext
		return ret
	}
	return *o.OutputContext
}

// GetOutputContextOk returns a tuple with the OutputContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetOutputContextOk() (*JsonStreamContext, bool) {
	if o == nil || o.OutputContext == nil {
		return nil, false
	}
	return o.OutputContext, true
}

// HasOutputContext returns a boolean if a field has been set.
func (o *JsonGenerator) HasOutputContext() bool {
	if o != nil && o.OutputContext != nil {
		return true
	}

	return false
}

// SetOutputContext gets a reference to the given JsonStreamContext and assigns it to the OutputContext field.
func (o *JsonGenerator) SetOutputContext(v JsonStreamContext) {
	o.OutputContext = &v
}

// GetOutputTarget returns the OutputTarget field value if set, zero value otherwise.
func (o *JsonGenerator) GetOutputTarget() map[string]interface{} {
	if o == nil || o.OutputTarget == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.OutputTarget
}

// GetOutputTargetOk returns a tuple with the OutputTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetOutputTargetOk() (*map[string]interface{}, bool) {
	if o == nil || o.OutputTarget == nil {
		return nil, false
	}
	return o.OutputTarget, true
}

// HasOutputTarget returns a boolean if a field has been set.
func (o *JsonGenerator) HasOutputTarget() bool {
	if o != nil && o.OutputTarget != nil {
		return true
	}

	return false
}

// SetOutputTarget gets a reference to the given map[string]interface{} and assigns it to the OutputTarget field.
func (o *JsonGenerator) SetOutputTarget(v map[string]interface{}) {
	o.OutputTarget = &v
}

// GetHighestEscapedChar returns the HighestEscapedChar field value if set, zero value otherwise.
func (o *JsonGenerator) GetHighestEscapedChar() int32 {
	if o == nil || o.HighestEscapedChar == nil {
		var ret int32
		return ret
	}
	return *o.HighestEscapedChar
}

// GetHighestEscapedCharOk returns a tuple with the HighestEscapedChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetHighestEscapedCharOk() (*int32, bool) {
	if o == nil || o.HighestEscapedChar == nil {
		return nil, false
	}
	return o.HighestEscapedChar, true
}

// HasHighestEscapedChar returns a boolean if a field has been set.
func (o *JsonGenerator) HasHighestEscapedChar() bool {
	if o != nil && o.HighestEscapedChar != nil {
		return true
	}

	return false
}

// SetHighestEscapedChar gets a reference to the given int32 and assigns it to the HighestEscapedChar field.
func (o *JsonGenerator) SetHighestEscapedChar(v int32) {
	o.HighestEscapedChar = &v
}

// GetOutputBuffered returns the OutputBuffered field value if set, zero value otherwise.
func (o *JsonGenerator) GetOutputBuffered() int32 {
	if o == nil || o.OutputBuffered == nil {
		var ret int32
		return ret
	}
	return *o.OutputBuffered
}

// GetOutputBufferedOk returns a tuple with the OutputBuffered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetOutputBufferedOk() (*int32, bool) {
	if o == nil || o.OutputBuffered == nil {
		return nil, false
	}
	return o.OutputBuffered, true
}

// HasOutputBuffered returns a boolean if a field has been set.
func (o *JsonGenerator) HasOutputBuffered() bool {
	if o != nil && o.OutputBuffered != nil {
		return true
	}

	return false
}

// SetOutputBuffered gets a reference to the given int32 and assigns it to the OutputBuffered field.
func (o *JsonGenerator) SetOutputBuffered(v int32) {
	o.OutputBuffered = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *JsonGenerator) GetSchema() FormatSchema {
	if o == nil || o.Schema == nil {
		var ret FormatSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonGenerator) GetSchemaOk() (*FormatSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *JsonGenerator) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given FormatSchema and assigns it to the Schema field.
func (o *JsonGenerator) SetSchema(v FormatSchema) {
	o.Schema = &v
}

func (o JsonGenerator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Closed != nil {
		toSerialize["closed"] = o.Closed
	}
	if o.Codec != nil {
		toSerialize["codec"] = o.Codec
	}
	if o.CharacterEscapes != nil {
		toSerialize["characterEscapes"] = o.CharacterEscapes
	}
	if o.CurrentValue != nil {
		toSerialize["currentValue"] = o.CurrentValue
	}
	if o.FeatureMask != nil {
		toSerialize["featureMask"] = o.FeatureMask
	}
	if o.FormatFeatures != nil {
		toSerialize["formatFeatures"] = o.FormatFeatures
	}
	if o.PrettyPrinter != nil {
		toSerialize["prettyPrinter"] = o.PrettyPrinter
	}
	if o.OutputContext != nil {
		toSerialize["outputContext"] = o.OutputContext
	}
	if o.OutputTarget != nil {
		toSerialize["outputTarget"] = o.OutputTarget
	}
	if o.HighestEscapedChar != nil {
		toSerialize["highestEscapedChar"] = o.HighestEscapedChar
	}
	if o.OutputBuffered != nil {
		toSerialize["outputBuffered"] = o.OutputBuffered
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	return json.Marshal(toSerialize)
}

type NullableJsonGenerator struct {
	value *JsonGenerator
	isSet bool
}

func (v NullableJsonGenerator) Get() *JsonGenerator {
	return v.value
}

func (v *NullableJsonGenerator) Set(val *JsonGenerator) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonGenerator) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonGenerator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonGenerator(val *JsonGenerator) *NullableJsonGenerator {
	return &NullableJsonGenerator{value: val, isSet: true}
}

func (v NullableJsonGenerator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonGenerator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


