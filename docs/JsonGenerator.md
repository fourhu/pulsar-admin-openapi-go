# JsonGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Closed** | Pointer to **bool** |  | [optional] 
**Codec** | Pointer to [**ObjectCodec**](ObjectCodec.md) |  | [optional] 
**CharacterEscapes** | Pointer to [**CharacterEscapes**](CharacterEscapes.md) |  | [optional] 
**CurrentValue** | Pointer to **map[string]interface{}** |  | [optional] 
**FeatureMask** | Pointer to **int32** |  | [optional] 
**FormatFeatures** | Pointer to **int32** |  | [optional] 
**PrettyPrinter** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputContext** | Pointer to [**JsonStreamContext**](JsonStreamContext.md) |  | [optional] 
**OutputTarget** | Pointer to **map[string]interface{}** |  | [optional] 
**HighestEscapedChar** | Pointer to **int32** |  | [optional] 
**OutputBuffered** | Pointer to **int32** |  | [optional] 
**Schema** | Pointer to [**FormatSchema**](FormatSchema.md) |  | [optional] 

## Methods

### NewJsonGenerator

`func NewJsonGenerator() *JsonGenerator`

NewJsonGenerator instantiates a new JsonGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonGeneratorWithDefaults

`func NewJsonGeneratorWithDefaults() *JsonGenerator`

NewJsonGeneratorWithDefaults instantiates a new JsonGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosed

`func (o *JsonGenerator) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *JsonGenerator) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *JsonGenerator) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *JsonGenerator) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetCodec

`func (o *JsonGenerator) GetCodec() ObjectCodec`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *JsonGenerator) GetCodecOk() (*ObjectCodec, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *JsonGenerator) SetCodec(v ObjectCodec)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *JsonGenerator) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetCharacterEscapes

`func (o *JsonGenerator) GetCharacterEscapes() CharacterEscapes`

GetCharacterEscapes returns the CharacterEscapes field if non-nil, zero value otherwise.

### GetCharacterEscapesOk

`func (o *JsonGenerator) GetCharacterEscapesOk() (*CharacterEscapes, bool)`

GetCharacterEscapesOk returns a tuple with the CharacterEscapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterEscapes

`func (o *JsonGenerator) SetCharacterEscapes(v CharacterEscapes)`

SetCharacterEscapes sets CharacterEscapes field to given value.

### HasCharacterEscapes

`func (o *JsonGenerator) HasCharacterEscapes() bool`

HasCharacterEscapes returns a boolean if a field has been set.

### GetCurrentValue

`func (o *JsonGenerator) GetCurrentValue() map[string]interface{}`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *JsonGenerator) GetCurrentValueOk() (*map[string]interface{}, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *JsonGenerator) SetCurrentValue(v map[string]interface{})`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *JsonGenerator) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetFeatureMask

`func (o *JsonGenerator) GetFeatureMask() int32`

GetFeatureMask returns the FeatureMask field if non-nil, zero value otherwise.

### GetFeatureMaskOk

`func (o *JsonGenerator) GetFeatureMaskOk() (*int32, bool)`

GetFeatureMaskOk returns a tuple with the FeatureMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureMask

`func (o *JsonGenerator) SetFeatureMask(v int32)`

SetFeatureMask sets FeatureMask field to given value.

### HasFeatureMask

`func (o *JsonGenerator) HasFeatureMask() bool`

HasFeatureMask returns a boolean if a field has been set.

### GetFormatFeatures

`func (o *JsonGenerator) GetFormatFeatures() int32`

GetFormatFeatures returns the FormatFeatures field if non-nil, zero value otherwise.

### GetFormatFeaturesOk

`func (o *JsonGenerator) GetFormatFeaturesOk() (*int32, bool)`

GetFormatFeaturesOk returns a tuple with the FormatFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatFeatures

`func (o *JsonGenerator) SetFormatFeatures(v int32)`

SetFormatFeatures sets FormatFeatures field to given value.

### HasFormatFeatures

`func (o *JsonGenerator) HasFormatFeatures() bool`

HasFormatFeatures returns a boolean if a field has been set.

### GetPrettyPrinter

`func (o *JsonGenerator) GetPrettyPrinter() map[string]interface{}`

GetPrettyPrinter returns the PrettyPrinter field if non-nil, zero value otherwise.

### GetPrettyPrinterOk

`func (o *JsonGenerator) GetPrettyPrinterOk() (*map[string]interface{}, bool)`

GetPrettyPrinterOk returns a tuple with the PrettyPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrettyPrinter

`func (o *JsonGenerator) SetPrettyPrinter(v map[string]interface{})`

SetPrettyPrinter sets PrettyPrinter field to given value.

### HasPrettyPrinter

`func (o *JsonGenerator) HasPrettyPrinter() bool`

HasPrettyPrinter returns a boolean if a field has been set.

### GetOutputContext

`func (o *JsonGenerator) GetOutputContext() JsonStreamContext`

GetOutputContext returns the OutputContext field if non-nil, zero value otherwise.

### GetOutputContextOk

`func (o *JsonGenerator) GetOutputContextOk() (*JsonStreamContext, bool)`

GetOutputContextOk returns a tuple with the OutputContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputContext

`func (o *JsonGenerator) SetOutputContext(v JsonStreamContext)`

SetOutputContext sets OutputContext field to given value.

### HasOutputContext

`func (o *JsonGenerator) HasOutputContext() bool`

HasOutputContext returns a boolean if a field has been set.

### GetOutputTarget

`func (o *JsonGenerator) GetOutputTarget() map[string]interface{}`

GetOutputTarget returns the OutputTarget field if non-nil, zero value otherwise.

### GetOutputTargetOk

`func (o *JsonGenerator) GetOutputTargetOk() (*map[string]interface{}, bool)`

GetOutputTargetOk returns a tuple with the OutputTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTarget

`func (o *JsonGenerator) SetOutputTarget(v map[string]interface{})`

SetOutputTarget sets OutputTarget field to given value.

### HasOutputTarget

`func (o *JsonGenerator) HasOutputTarget() bool`

HasOutputTarget returns a boolean if a field has been set.

### GetHighestEscapedChar

`func (o *JsonGenerator) GetHighestEscapedChar() int32`

GetHighestEscapedChar returns the HighestEscapedChar field if non-nil, zero value otherwise.

### GetHighestEscapedCharOk

`func (o *JsonGenerator) GetHighestEscapedCharOk() (*int32, bool)`

GetHighestEscapedCharOk returns a tuple with the HighestEscapedChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestEscapedChar

`func (o *JsonGenerator) SetHighestEscapedChar(v int32)`

SetHighestEscapedChar sets HighestEscapedChar field to given value.

### HasHighestEscapedChar

`func (o *JsonGenerator) HasHighestEscapedChar() bool`

HasHighestEscapedChar returns a boolean if a field has been set.

### GetOutputBuffered

`func (o *JsonGenerator) GetOutputBuffered() int32`

GetOutputBuffered returns the OutputBuffered field if non-nil, zero value otherwise.

### GetOutputBufferedOk

`func (o *JsonGenerator) GetOutputBufferedOk() (*int32, bool)`

GetOutputBufferedOk returns a tuple with the OutputBuffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputBuffered

`func (o *JsonGenerator) SetOutputBuffered(v int32)`

SetOutputBuffered sets OutputBuffered field to given value.

### HasOutputBuffered

`func (o *JsonGenerator) HasOutputBuffered() bool`

HasOutputBuffered returns a boolean if a field has been set.

### GetSchema

`func (o *JsonGenerator) GetSchema() FormatSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *JsonGenerator) GetSchemaOk() (*FormatSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *JsonGenerator) SetSchema(v FormatSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *JsonGenerator) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


