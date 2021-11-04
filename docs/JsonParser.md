# JsonParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Closed** | Pointer to **bool** |  | [optional] 
**CurrentLocation** | Pointer to [**JsonLocation**](JsonLocation.md) |  | [optional] 
**Codec** | Pointer to [**ObjectCodec**](ObjectCodec.md) |  | [optional] 
**BinaryValue** | Pointer to **[]string** |  | [optional] 
**CurrentValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ParsingContext** | Pointer to [**JsonStreamContext**](JsonStreamContext.md) |  | [optional] 
**TokenLocation** | Pointer to [**JsonLocation**](JsonLocation.md) |  | [optional] 
**InputSource** | Pointer to **map[string]interface{}** |  | [optional] 
**FeatureMask** | Pointer to **int32** |  | [optional] 
**FormatFeatures** | Pointer to **int32** |  | [optional] 
**CurrentToken** | Pointer to **string** |  | [optional] 
**CurrentTokenId** | Pointer to **int32** |  | [optional] 
**ExpectedStartArrayToken** | Pointer to **bool** |  | [optional] 
**ExpectedStartObjectToken** | Pointer to **bool** |  | [optional] 
**LastClearedToken** | Pointer to **string** |  | [optional] 
**CurrentName** | Pointer to **string** |  | [optional] 
**TextOffset** | Pointer to **int32** |  | [optional] 
**NumberValue** | Pointer to **map[string]interface{}** |  | [optional] 
**NumberType** | Pointer to **string** |  | [optional] 
**BigIntegerValue** | Pointer to **int32** |  | [optional] 
**DecimalValue** | Pointer to **float32** |  | [optional] 
**EmbeddedObject** | Pointer to **map[string]interface{}** |  | [optional] 
**ValueAsInt** | Pointer to **int32** |  | [optional] 
**ValueAsLong** | Pointer to **int64** |  | [optional] 
**ValueAsDouble** | Pointer to **float64** |  | [optional] 
**ValueAsBoolean** | Pointer to **bool** |  | [optional] 
**ValueAsString** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **map[string]interface{}** |  | [optional] 
**TypeId** | Pointer to **map[string]interface{}** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**TextCharacters** | Pointer to **[]string** |  | [optional] 
**TextLength** | Pointer to **int32** |  | [optional] 
**BooleanValue** | Pointer to **bool** |  | [optional] 
**ByteValue** | Pointer to **string** |  | [optional] 
**DoubleValue** | Pointer to **float64** |  | [optional] 
**FloatValue** | Pointer to **float32** |  | [optional] 
**LongValue** | Pointer to **int64** |  | [optional] 
**ShortValue** | Pointer to **int32** |  | [optional] 
**IntValue** | Pointer to **int32** |  | [optional] 
**Schema** | Pointer to [**FormatSchema**](FormatSchema.md) |  | [optional] 

## Methods

### NewJsonParser

`func NewJsonParser() *JsonParser`

NewJsonParser instantiates a new JsonParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonParserWithDefaults

`func NewJsonParserWithDefaults() *JsonParser`

NewJsonParserWithDefaults instantiates a new JsonParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosed

`func (o *JsonParser) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *JsonParser) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *JsonParser) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *JsonParser) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetCurrentLocation

`func (o *JsonParser) GetCurrentLocation() JsonLocation`

GetCurrentLocation returns the CurrentLocation field if non-nil, zero value otherwise.

### GetCurrentLocationOk

`func (o *JsonParser) GetCurrentLocationOk() (*JsonLocation, bool)`

GetCurrentLocationOk returns a tuple with the CurrentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLocation

`func (o *JsonParser) SetCurrentLocation(v JsonLocation)`

SetCurrentLocation sets CurrentLocation field to given value.

### HasCurrentLocation

`func (o *JsonParser) HasCurrentLocation() bool`

HasCurrentLocation returns a boolean if a field has been set.

### GetCodec

`func (o *JsonParser) GetCodec() ObjectCodec`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *JsonParser) GetCodecOk() (*ObjectCodec, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *JsonParser) SetCodec(v ObjectCodec)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *JsonParser) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetBinaryValue

`func (o *JsonParser) GetBinaryValue() []string`

GetBinaryValue returns the BinaryValue field if non-nil, zero value otherwise.

### GetBinaryValueOk

`func (o *JsonParser) GetBinaryValueOk() (*[]string, bool)`

GetBinaryValueOk returns a tuple with the BinaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryValue

`func (o *JsonParser) SetBinaryValue(v []string)`

SetBinaryValue sets BinaryValue field to given value.

### HasBinaryValue

`func (o *JsonParser) HasBinaryValue() bool`

HasBinaryValue returns a boolean if a field has been set.

### GetCurrentValue

`func (o *JsonParser) GetCurrentValue() map[string]interface{}`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *JsonParser) GetCurrentValueOk() (*map[string]interface{}, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *JsonParser) SetCurrentValue(v map[string]interface{})`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *JsonParser) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetParsingContext

`func (o *JsonParser) GetParsingContext() JsonStreamContext`

GetParsingContext returns the ParsingContext field if non-nil, zero value otherwise.

### GetParsingContextOk

`func (o *JsonParser) GetParsingContextOk() (*JsonStreamContext, bool)`

GetParsingContextOk returns a tuple with the ParsingContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingContext

`func (o *JsonParser) SetParsingContext(v JsonStreamContext)`

SetParsingContext sets ParsingContext field to given value.

### HasParsingContext

`func (o *JsonParser) HasParsingContext() bool`

HasParsingContext returns a boolean if a field has been set.

### GetTokenLocation

`func (o *JsonParser) GetTokenLocation() JsonLocation`

GetTokenLocation returns the TokenLocation field if non-nil, zero value otherwise.

### GetTokenLocationOk

`func (o *JsonParser) GetTokenLocationOk() (*JsonLocation, bool)`

GetTokenLocationOk returns a tuple with the TokenLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLocation

`func (o *JsonParser) SetTokenLocation(v JsonLocation)`

SetTokenLocation sets TokenLocation field to given value.

### HasTokenLocation

`func (o *JsonParser) HasTokenLocation() bool`

HasTokenLocation returns a boolean if a field has been set.

### GetInputSource

`func (o *JsonParser) GetInputSource() map[string]interface{}`

GetInputSource returns the InputSource field if non-nil, zero value otherwise.

### GetInputSourceOk

`func (o *JsonParser) GetInputSourceOk() (*map[string]interface{}, bool)`

GetInputSourceOk returns a tuple with the InputSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSource

`func (o *JsonParser) SetInputSource(v map[string]interface{})`

SetInputSource sets InputSource field to given value.

### HasInputSource

`func (o *JsonParser) HasInputSource() bool`

HasInputSource returns a boolean if a field has been set.

### GetFeatureMask

`func (o *JsonParser) GetFeatureMask() int32`

GetFeatureMask returns the FeatureMask field if non-nil, zero value otherwise.

### GetFeatureMaskOk

`func (o *JsonParser) GetFeatureMaskOk() (*int32, bool)`

GetFeatureMaskOk returns a tuple with the FeatureMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureMask

`func (o *JsonParser) SetFeatureMask(v int32)`

SetFeatureMask sets FeatureMask field to given value.

### HasFeatureMask

`func (o *JsonParser) HasFeatureMask() bool`

HasFeatureMask returns a boolean if a field has been set.

### GetFormatFeatures

`func (o *JsonParser) GetFormatFeatures() int32`

GetFormatFeatures returns the FormatFeatures field if non-nil, zero value otherwise.

### GetFormatFeaturesOk

`func (o *JsonParser) GetFormatFeaturesOk() (*int32, bool)`

GetFormatFeaturesOk returns a tuple with the FormatFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatFeatures

`func (o *JsonParser) SetFormatFeatures(v int32)`

SetFormatFeatures sets FormatFeatures field to given value.

### HasFormatFeatures

`func (o *JsonParser) HasFormatFeatures() bool`

HasFormatFeatures returns a boolean if a field has been set.

### GetCurrentToken

`func (o *JsonParser) GetCurrentToken() string`

GetCurrentToken returns the CurrentToken field if non-nil, zero value otherwise.

### GetCurrentTokenOk

`func (o *JsonParser) GetCurrentTokenOk() (*string, bool)`

GetCurrentTokenOk returns a tuple with the CurrentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentToken

`func (o *JsonParser) SetCurrentToken(v string)`

SetCurrentToken sets CurrentToken field to given value.

### HasCurrentToken

`func (o *JsonParser) HasCurrentToken() bool`

HasCurrentToken returns a boolean if a field has been set.

### GetCurrentTokenId

`func (o *JsonParser) GetCurrentTokenId() int32`

GetCurrentTokenId returns the CurrentTokenId field if non-nil, zero value otherwise.

### GetCurrentTokenIdOk

`func (o *JsonParser) GetCurrentTokenIdOk() (*int32, bool)`

GetCurrentTokenIdOk returns a tuple with the CurrentTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTokenId

`func (o *JsonParser) SetCurrentTokenId(v int32)`

SetCurrentTokenId sets CurrentTokenId field to given value.

### HasCurrentTokenId

`func (o *JsonParser) HasCurrentTokenId() bool`

HasCurrentTokenId returns a boolean if a field has been set.

### GetExpectedStartArrayToken

`func (o *JsonParser) GetExpectedStartArrayToken() bool`

GetExpectedStartArrayToken returns the ExpectedStartArrayToken field if non-nil, zero value otherwise.

### GetExpectedStartArrayTokenOk

`func (o *JsonParser) GetExpectedStartArrayTokenOk() (*bool, bool)`

GetExpectedStartArrayTokenOk returns a tuple with the ExpectedStartArrayToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStartArrayToken

`func (o *JsonParser) SetExpectedStartArrayToken(v bool)`

SetExpectedStartArrayToken sets ExpectedStartArrayToken field to given value.

### HasExpectedStartArrayToken

`func (o *JsonParser) HasExpectedStartArrayToken() bool`

HasExpectedStartArrayToken returns a boolean if a field has been set.

### GetExpectedStartObjectToken

`func (o *JsonParser) GetExpectedStartObjectToken() bool`

GetExpectedStartObjectToken returns the ExpectedStartObjectToken field if non-nil, zero value otherwise.

### GetExpectedStartObjectTokenOk

`func (o *JsonParser) GetExpectedStartObjectTokenOk() (*bool, bool)`

GetExpectedStartObjectTokenOk returns a tuple with the ExpectedStartObjectToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStartObjectToken

`func (o *JsonParser) SetExpectedStartObjectToken(v bool)`

SetExpectedStartObjectToken sets ExpectedStartObjectToken field to given value.

### HasExpectedStartObjectToken

`func (o *JsonParser) HasExpectedStartObjectToken() bool`

HasExpectedStartObjectToken returns a boolean if a field has been set.

### GetLastClearedToken

`func (o *JsonParser) GetLastClearedToken() string`

GetLastClearedToken returns the LastClearedToken field if non-nil, zero value otherwise.

### GetLastClearedTokenOk

`func (o *JsonParser) GetLastClearedTokenOk() (*string, bool)`

GetLastClearedTokenOk returns a tuple with the LastClearedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastClearedToken

`func (o *JsonParser) SetLastClearedToken(v string)`

SetLastClearedToken sets LastClearedToken field to given value.

### HasLastClearedToken

`func (o *JsonParser) HasLastClearedToken() bool`

HasLastClearedToken returns a boolean if a field has been set.

### GetCurrentName

`func (o *JsonParser) GetCurrentName() string`

GetCurrentName returns the CurrentName field if non-nil, zero value otherwise.

### GetCurrentNameOk

`func (o *JsonParser) GetCurrentNameOk() (*string, bool)`

GetCurrentNameOk returns a tuple with the CurrentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentName

`func (o *JsonParser) SetCurrentName(v string)`

SetCurrentName sets CurrentName field to given value.

### HasCurrentName

`func (o *JsonParser) HasCurrentName() bool`

HasCurrentName returns a boolean if a field has been set.

### GetTextOffset

`func (o *JsonParser) GetTextOffset() int32`

GetTextOffset returns the TextOffset field if non-nil, zero value otherwise.

### GetTextOffsetOk

`func (o *JsonParser) GetTextOffsetOk() (*int32, bool)`

GetTextOffsetOk returns a tuple with the TextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextOffset

`func (o *JsonParser) SetTextOffset(v int32)`

SetTextOffset sets TextOffset field to given value.

### HasTextOffset

`func (o *JsonParser) HasTextOffset() bool`

HasTextOffset returns a boolean if a field has been set.

### GetNumberValue

`func (o *JsonParser) GetNumberValue() map[string]interface{}`

GetNumberValue returns the NumberValue field if non-nil, zero value otherwise.

### GetNumberValueOk

`func (o *JsonParser) GetNumberValueOk() (*map[string]interface{}, bool)`

GetNumberValueOk returns a tuple with the NumberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValue

`func (o *JsonParser) SetNumberValue(v map[string]interface{})`

SetNumberValue sets NumberValue field to given value.

### HasNumberValue

`func (o *JsonParser) HasNumberValue() bool`

HasNumberValue returns a boolean if a field has been set.

### GetNumberType

`func (o *JsonParser) GetNumberType() string`

GetNumberType returns the NumberType field if non-nil, zero value otherwise.

### GetNumberTypeOk

`func (o *JsonParser) GetNumberTypeOk() (*string, bool)`

GetNumberTypeOk returns a tuple with the NumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberType

`func (o *JsonParser) SetNumberType(v string)`

SetNumberType sets NumberType field to given value.

### HasNumberType

`func (o *JsonParser) HasNumberType() bool`

HasNumberType returns a boolean if a field has been set.

### GetBigIntegerValue

`func (o *JsonParser) GetBigIntegerValue() int32`

GetBigIntegerValue returns the BigIntegerValue field if non-nil, zero value otherwise.

### GetBigIntegerValueOk

`func (o *JsonParser) GetBigIntegerValueOk() (*int32, bool)`

GetBigIntegerValueOk returns a tuple with the BigIntegerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigIntegerValue

`func (o *JsonParser) SetBigIntegerValue(v int32)`

SetBigIntegerValue sets BigIntegerValue field to given value.

### HasBigIntegerValue

`func (o *JsonParser) HasBigIntegerValue() bool`

HasBigIntegerValue returns a boolean if a field has been set.

### GetDecimalValue

`func (o *JsonParser) GetDecimalValue() float32`

GetDecimalValue returns the DecimalValue field if non-nil, zero value otherwise.

### GetDecimalValueOk

`func (o *JsonParser) GetDecimalValueOk() (*float32, bool)`

GetDecimalValueOk returns a tuple with the DecimalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalValue

`func (o *JsonParser) SetDecimalValue(v float32)`

SetDecimalValue sets DecimalValue field to given value.

### HasDecimalValue

`func (o *JsonParser) HasDecimalValue() bool`

HasDecimalValue returns a boolean if a field has been set.

### GetEmbeddedObject

`func (o *JsonParser) GetEmbeddedObject() map[string]interface{}`

GetEmbeddedObject returns the EmbeddedObject field if non-nil, zero value otherwise.

### GetEmbeddedObjectOk

`func (o *JsonParser) GetEmbeddedObjectOk() (*map[string]interface{}, bool)`

GetEmbeddedObjectOk returns a tuple with the EmbeddedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddedObject

`func (o *JsonParser) SetEmbeddedObject(v map[string]interface{})`

SetEmbeddedObject sets EmbeddedObject field to given value.

### HasEmbeddedObject

`func (o *JsonParser) HasEmbeddedObject() bool`

HasEmbeddedObject returns a boolean if a field has been set.

### GetValueAsInt

`func (o *JsonParser) GetValueAsInt() int32`

GetValueAsInt returns the ValueAsInt field if non-nil, zero value otherwise.

### GetValueAsIntOk

`func (o *JsonParser) GetValueAsIntOk() (*int32, bool)`

GetValueAsIntOk returns a tuple with the ValueAsInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsInt

`func (o *JsonParser) SetValueAsInt(v int32)`

SetValueAsInt sets ValueAsInt field to given value.

### HasValueAsInt

`func (o *JsonParser) HasValueAsInt() bool`

HasValueAsInt returns a boolean if a field has been set.

### GetValueAsLong

`func (o *JsonParser) GetValueAsLong() int64`

GetValueAsLong returns the ValueAsLong field if non-nil, zero value otherwise.

### GetValueAsLongOk

`func (o *JsonParser) GetValueAsLongOk() (*int64, bool)`

GetValueAsLongOk returns a tuple with the ValueAsLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsLong

`func (o *JsonParser) SetValueAsLong(v int64)`

SetValueAsLong sets ValueAsLong field to given value.

### HasValueAsLong

`func (o *JsonParser) HasValueAsLong() bool`

HasValueAsLong returns a boolean if a field has been set.

### GetValueAsDouble

`func (o *JsonParser) GetValueAsDouble() float64`

GetValueAsDouble returns the ValueAsDouble field if non-nil, zero value otherwise.

### GetValueAsDoubleOk

`func (o *JsonParser) GetValueAsDoubleOk() (*float64, bool)`

GetValueAsDoubleOk returns a tuple with the ValueAsDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsDouble

`func (o *JsonParser) SetValueAsDouble(v float64)`

SetValueAsDouble sets ValueAsDouble field to given value.

### HasValueAsDouble

`func (o *JsonParser) HasValueAsDouble() bool`

HasValueAsDouble returns a boolean if a field has been set.

### GetValueAsBoolean

`func (o *JsonParser) GetValueAsBoolean() bool`

GetValueAsBoolean returns the ValueAsBoolean field if non-nil, zero value otherwise.

### GetValueAsBooleanOk

`func (o *JsonParser) GetValueAsBooleanOk() (*bool, bool)`

GetValueAsBooleanOk returns a tuple with the ValueAsBoolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsBoolean

`func (o *JsonParser) SetValueAsBoolean(v bool)`

SetValueAsBoolean sets ValueAsBoolean field to given value.

### HasValueAsBoolean

`func (o *JsonParser) HasValueAsBoolean() bool`

HasValueAsBoolean returns a boolean if a field has been set.

### GetValueAsString

`func (o *JsonParser) GetValueAsString() string`

GetValueAsString returns the ValueAsString field if non-nil, zero value otherwise.

### GetValueAsStringOk

`func (o *JsonParser) GetValueAsStringOk() (*string, bool)`

GetValueAsStringOk returns a tuple with the ValueAsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsString

`func (o *JsonParser) SetValueAsString(v string)`

SetValueAsString sets ValueAsString field to given value.

### HasValueAsString

`func (o *JsonParser) HasValueAsString() bool`

HasValueAsString returns a boolean if a field has been set.

### GetObjectId

`func (o *JsonParser) GetObjectId() map[string]interface{}`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *JsonParser) GetObjectIdOk() (*map[string]interface{}, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *JsonParser) SetObjectId(v map[string]interface{})`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *JsonParser) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetTypeId

`func (o *JsonParser) GetTypeId() map[string]interface{}`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *JsonParser) GetTypeIdOk() (*map[string]interface{}, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *JsonParser) SetTypeId(v map[string]interface{})`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *JsonParser) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetText

`func (o *JsonParser) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *JsonParser) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *JsonParser) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *JsonParser) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTextCharacters

`func (o *JsonParser) GetTextCharacters() []string`

GetTextCharacters returns the TextCharacters field if non-nil, zero value otherwise.

### GetTextCharactersOk

`func (o *JsonParser) GetTextCharactersOk() (*[]string, bool)`

GetTextCharactersOk returns a tuple with the TextCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCharacters

`func (o *JsonParser) SetTextCharacters(v []string)`

SetTextCharacters sets TextCharacters field to given value.

### HasTextCharacters

`func (o *JsonParser) HasTextCharacters() bool`

HasTextCharacters returns a boolean if a field has been set.

### GetTextLength

`func (o *JsonParser) GetTextLength() int32`

GetTextLength returns the TextLength field if non-nil, zero value otherwise.

### GetTextLengthOk

`func (o *JsonParser) GetTextLengthOk() (*int32, bool)`

GetTextLengthOk returns a tuple with the TextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextLength

`func (o *JsonParser) SetTextLength(v int32)`

SetTextLength sets TextLength field to given value.

### HasTextLength

`func (o *JsonParser) HasTextLength() bool`

HasTextLength returns a boolean if a field has been set.

### GetBooleanValue

`func (o *JsonParser) GetBooleanValue() bool`

GetBooleanValue returns the BooleanValue field if non-nil, zero value otherwise.

### GetBooleanValueOk

`func (o *JsonParser) GetBooleanValueOk() (*bool, bool)`

GetBooleanValueOk returns a tuple with the BooleanValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanValue

`func (o *JsonParser) SetBooleanValue(v bool)`

SetBooleanValue sets BooleanValue field to given value.

### HasBooleanValue

`func (o *JsonParser) HasBooleanValue() bool`

HasBooleanValue returns a boolean if a field has been set.

### GetByteValue

`func (o *JsonParser) GetByteValue() string`

GetByteValue returns the ByteValue field if non-nil, zero value otherwise.

### GetByteValueOk

`func (o *JsonParser) GetByteValueOk() (*string, bool)`

GetByteValueOk returns a tuple with the ByteValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteValue

`func (o *JsonParser) SetByteValue(v string)`

SetByteValue sets ByteValue field to given value.

### HasByteValue

`func (o *JsonParser) HasByteValue() bool`

HasByteValue returns a boolean if a field has been set.

### GetDoubleValue

`func (o *JsonParser) GetDoubleValue() float64`

GetDoubleValue returns the DoubleValue field if non-nil, zero value otherwise.

### GetDoubleValueOk

`func (o *JsonParser) GetDoubleValueOk() (*float64, bool)`

GetDoubleValueOk returns a tuple with the DoubleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleValue

`func (o *JsonParser) SetDoubleValue(v float64)`

SetDoubleValue sets DoubleValue field to given value.

### HasDoubleValue

`func (o *JsonParser) HasDoubleValue() bool`

HasDoubleValue returns a boolean if a field has been set.

### GetFloatValue

`func (o *JsonParser) GetFloatValue() float32`

GetFloatValue returns the FloatValue field if non-nil, zero value otherwise.

### GetFloatValueOk

`func (o *JsonParser) GetFloatValueOk() (*float32, bool)`

GetFloatValueOk returns a tuple with the FloatValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatValue

`func (o *JsonParser) SetFloatValue(v float32)`

SetFloatValue sets FloatValue field to given value.

### HasFloatValue

`func (o *JsonParser) HasFloatValue() bool`

HasFloatValue returns a boolean if a field has been set.

### GetLongValue

`func (o *JsonParser) GetLongValue() int64`

GetLongValue returns the LongValue field if non-nil, zero value otherwise.

### GetLongValueOk

`func (o *JsonParser) GetLongValueOk() (*int64, bool)`

GetLongValueOk returns a tuple with the LongValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongValue

`func (o *JsonParser) SetLongValue(v int64)`

SetLongValue sets LongValue field to given value.

### HasLongValue

`func (o *JsonParser) HasLongValue() bool`

HasLongValue returns a boolean if a field has been set.

### GetShortValue

`func (o *JsonParser) GetShortValue() int32`

GetShortValue returns the ShortValue field if non-nil, zero value otherwise.

### GetShortValueOk

`func (o *JsonParser) GetShortValueOk() (*int32, bool)`

GetShortValueOk returns a tuple with the ShortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortValue

`func (o *JsonParser) SetShortValue(v int32)`

SetShortValue sets ShortValue field to given value.

### HasShortValue

`func (o *JsonParser) HasShortValue() bool`

HasShortValue returns a boolean if a field has been set.

### GetIntValue

`func (o *JsonParser) GetIntValue() int32`

GetIntValue returns the IntValue field if non-nil, zero value otherwise.

### GetIntValueOk

`func (o *JsonParser) GetIntValueOk() (*int32, bool)`

GetIntValueOk returns a tuple with the IntValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntValue

`func (o *JsonParser) SetIntValue(v int32)`

SetIntValue sets IntValue field to given value.

### HasIntValue

`func (o *JsonParser) HasIntValue() bool`

HasIntValue returns a boolean if a field has been set.

### GetSchema

`func (o *JsonParser) GetSchema() FormatSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *JsonParser) GetSchemaOk() (*FormatSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *JsonParser) SetSchema(v FormatSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *JsonParser) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


