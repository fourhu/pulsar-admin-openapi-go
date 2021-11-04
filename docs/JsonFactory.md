# JsonFactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codec** | Pointer to [**ObjectCodec**](ObjectCodec.md) |  | [optional] 
**InputDecorator** | Pointer to **map[string]interface{}** |  | [optional] 
**CharacterEscapes** | Pointer to [**CharacterEscapes**](CharacterEscapes.md) |  | [optional] 
**OutputDecorator** | Pointer to **map[string]interface{}** |  | [optional] 
**RootValueSeparator** | Pointer to **string** |  | [optional] 
**FormatName** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonFactory

`func NewJsonFactory() *JsonFactory`

NewJsonFactory instantiates a new JsonFactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonFactoryWithDefaults

`func NewJsonFactoryWithDefaults() *JsonFactory`

NewJsonFactoryWithDefaults instantiates a new JsonFactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodec

`func (o *JsonFactory) GetCodec() ObjectCodec`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *JsonFactory) GetCodecOk() (*ObjectCodec, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *JsonFactory) SetCodec(v ObjectCodec)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *JsonFactory) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetInputDecorator

`func (o *JsonFactory) GetInputDecorator() map[string]interface{}`

GetInputDecorator returns the InputDecorator field if non-nil, zero value otherwise.

### GetInputDecoratorOk

`func (o *JsonFactory) GetInputDecoratorOk() (*map[string]interface{}, bool)`

GetInputDecoratorOk returns a tuple with the InputDecorator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDecorator

`func (o *JsonFactory) SetInputDecorator(v map[string]interface{})`

SetInputDecorator sets InputDecorator field to given value.

### HasInputDecorator

`func (o *JsonFactory) HasInputDecorator() bool`

HasInputDecorator returns a boolean if a field has been set.

### GetCharacterEscapes

`func (o *JsonFactory) GetCharacterEscapes() CharacterEscapes`

GetCharacterEscapes returns the CharacterEscapes field if non-nil, zero value otherwise.

### GetCharacterEscapesOk

`func (o *JsonFactory) GetCharacterEscapesOk() (*CharacterEscapes, bool)`

GetCharacterEscapesOk returns a tuple with the CharacterEscapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterEscapes

`func (o *JsonFactory) SetCharacterEscapes(v CharacterEscapes)`

SetCharacterEscapes sets CharacterEscapes field to given value.

### HasCharacterEscapes

`func (o *JsonFactory) HasCharacterEscapes() bool`

HasCharacterEscapes returns a boolean if a field has been set.

### GetOutputDecorator

`func (o *JsonFactory) GetOutputDecorator() map[string]interface{}`

GetOutputDecorator returns the OutputDecorator field if non-nil, zero value otherwise.

### GetOutputDecoratorOk

`func (o *JsonFactory) GetOutputDecoratorOk() (*map[string]interface{}, bool)`

GetOutputDecoratorOk returns a tuple with the OutputDecorator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDecorator

`func (o *JsonFactory) SetOutputDecorator(v map[string]interface{})`

SetOutputDecorator sets OutputDecorator field to given value.

### HasOutputDecorator

`func (o *JsonFactory) HasOutputDecorator() bool`

HasOutputDecorator returns a boolean if a field has been set.

### GetRootValueSeparator

`func (o *JsonFactory) GetRootValueSeparator() string`

GetRootValueSeparator returns the RootValueSeparator field if non-nil, zero value otherwise.

### GetRootValueSeparatorOk

`func (o *JsonFactory) GetRootValueSeparatorOk() (*string, bool)`

GetRootValueSeparatorOk returns a tuple with the RootValueSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootValueSeparator

`func (o *JsonFactory) SetRootValueSeparator(v string)`

SetRootValueSeparator sets RootValueSeparator field to given value.

### HasRootValueSeparator

`func (o *JsonFactory) HasRootValueSeparator() bool`

HasRootValueSeparator returns a boolean if a field has been set.

### GetFormatName

`func (o *JsonFactory) GetFormatName() string`

GetFormatName returns the FormatName field if non-nil, zero value otherwise.

### GetFormatNameOk

`func (o *JsonFactory) GetFormatNameOk() (*string, bool)`

GetFormatNameOk returns a tuple with the FormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatName

`func (o *JsonFactory) SetFormatName(v string)`

SetFormatName sets FormatName field to given value.

### HasFormatName

`func (o *JsonFactory) HasFormatName() bool`

HasFormatName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


