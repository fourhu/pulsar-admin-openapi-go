# Locale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Variant** | Pointer to **string** |  | [optional] 
**ExtensionKeys** | Pointer to **[]string** |  | [optional] 
**UnicodeLocaleAttributes** | Pointer to **[]string** |  | [optional] 
**UnicodeLocaleKeys** | Pointer to **[]string** |  | [optional] 
**Iso3Language** | Pointer to **string** |  | [optional] 
**Iso3Country** | Pointer to **string** |  | [optional] 
**DisplayLanguage** | Pointer to **string** |  | [optional] 
**DisplayScript** | Pointer to **string** |  | [optional] 
**DisplayCountry** | Pointer to **string** |  | [optional] 
**DisplayVariant** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 

## Methods

### NewLocale

`func NewLocale() *Locale`

NewLocale instantiates a new Locale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocaleWithDefaults

`func NewLocaleWithDefaults() *Locale`

NewLocaleWithDefaults instantiates a new Locale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *Locale) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *Locale) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *Locale) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *Locale) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetCountry

`func (o *Locale) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Locale) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Locale) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Locale) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetVariant

`func (o *Locale) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *Locale) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *Locale) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *Locale) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetExtensionKeys

`func (o *Locale) GetExtensionKeys() []string`

GetExtensionKeys returns the ExtensionKeys field if non-nil, zero value otherwise.

### GetExtensionKeysOk

`func (o *Locale) GetExtensionKeysOk() (*[]string, bool)`

GetExtensionKeysOk returns a tuple with the ExtensionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionKeys

`func (o *Locale) SetExtensionKeys(v []string)`

SetExtensionKeys sets ExtensionKeys field to given value.

### HasExtensionKeys

`func (o *Locale) HasExtensionKeys() bool`

HasExtensionKeys returns a boolean if a field has been set.

### GetUnicodeLocaleAttributes

`func (o *Locale) GetUnicodeLocaleAttributes() []string`

GetUnicodeLocaleAttributes returns the UnicodeLocaleAttributes field if non-nil, zero value otherwise.

### GetUnicodeLocaleAttributesOk

`func (o *Locale) GetUnicodeLocaleAttributesOk() (*[]string, bool)`

GetUnicodeLocaleAttributesOk returns a tuple with the UnicodeLocaleAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeLocaleAttributes

`func (o *Locale) SetUnicodeLocaleAttributes(v []string)`

SetUnicodeLocaleAttributes sets UnicodeLocaleAttributes field to given value.

### HasUnicodeLocaleAttributes

`func (o *Locale) HasUnicodeLocaleAttributes() bool`

HasUnicodeLocaleAttributes returns a boolean if a field has been set.

### GetUnicodeLocaleKeys

`func (o *Locale) GetUnicodeLocaleKeys() []string`

GetUnicodeLocaleKeys returns the UnicodeLocaleKeys field if non-nil, zero value otherwise.

### GetUnicodeLocaleKeysOk

`func (o *Locale) GetUnicodeLocaleKeysOk() (*[]string, bool)`

GetUnicodeLocaleKeysOk returns a tuple with the UnicodeLocaleKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeLocaleKeys

`func (o *Locale) SetUnicodeLocaleKeys(v []string)`

SetUnicodeLocaleKeys sets UnicodeLocaleKeys field to given value.

### HasUnicodeLocaleKeys

`func (o *Locale) HasUnicodeLocaleKeys() bool`

HasUnicodeLocaleKeys returns a boolean if a field has been set.

### GetIso3Language

`func (o *Locale) GetIso3Language() string`

GetIso3Language returns the Iso3Language field if non-nil, zero value otherwise.

### GetIso3LanguageOk

`func (o *Locale) GetIso3LanguageOk() (*string, bool)`

GetIso3LanguageOk returns a tuple with the Iso3Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso3Language

`func (o *Locale) SetIso3Language(v string)`

SetIso3Language sets Iso3Language field to given value.

### HasIso3Language

`func (o *Locale) HasIso3Language() bool`

HasIso3Language returns a boolean if a field has been set.

### GetIso3Country

`func (o *Locale) GetIso3Country() string`

GetIso3Country returns the Iso3Country field if non-nil, zero value otherwise.

### GetIso3CountryOk

`func (o *Locale) GetIso3CountryOk() (*string, bool)`

GetIso3CountryOk returns a tuple with the Iso3Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso3Country

`func (o *Locale) SetIso3Country(v string)`

SetIso3Country sets Iso3Country field to given value.

### HasIso3Country

`func (o *Locale) HasIso3Country() bool`

HasIso3Country returns a boolean if a field has been set.

### GetDisplayLanguage

`func (o *Locale) GetDisplayLanguage() string`

GetDisplayLanguage returns the DisplayLanguage field if non-nil, zero value otherwise.

### GetDisplayLanguageOk

`func (o *Locale) GetDisplayLanguageOk() (*string, bool)`

GetDisplayLanguageOk returns a tuple with the DisplayLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLanguage

`func (o *Locale) SetDisplayLanguage(v string)`

SetDisplayLanguage sets DisplayLanguage field to given value.

### HasDisplayLanguage

`func (o *Locale) HasDisplayLanguage() bool`

HasDisplayLanguage returns a boolean if a field has been set.

### GetDisplayScript

`func (o *Locale) GetDisplayScript() string`

GetDisplayScript returns the DisplayScript field if non-nil, zero value otherwise.

### GetDisplayScriptOk

`func (o *Locale) GetDisplayScriptOk() (*string, bool)`

GetDisplayScriptOk returns a tuple with the DisplayScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayScript

`func (o *Locale) SetDisplayScript(v string)`

SetDisplayScript sets DisplayScript field to given value.

### HasDisplayScript

`func (o *Locale) HasDisplayScript() bool`

HasDisplayScript returns a boolean if a field has been set.

### GetDisplayCountry

`func (o *Locale) GetDisplayCountry() string`

GetDisplayCountry returns the DisplayCountry field if non-nil, zero value otherwise.

### GetDisplayCountryOk

`func (o *Locale) GetDisplayCountryOk() (*string, bool)`

GetDisplayCountryOk returns a tuple with the DisplayCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCountry

`func (o *Locale) SetDisplayCountry(v string)`

SetDisplayCountry sets DisplayCountry field to given value.

### HasDisplayCountry

`func (o *Locale) HasDisplayCountry() bool`

HasDisplayCountry returns a boolean if a field has been set.

### GetDisplayVariant

`func (o *Locale) GetDisplayVariant() string`

GetDisplayVariant returns the DisplayVariant field if non-nil, zero value otherwise.

### GetDisplayVariantOk

`func (o *Locale) GetDisplayVariantOk() (*string, bool)`

GetDisplayVariantOk returns a tuple with the DisplayVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayVariant

`func (o *Locale) SetDisplayVariant(v string)`

SetDisplayVariant sets DisplayVariant field to given value.

### HasDisplayVariant

`func (o *Locale) HasDisplayVariant() bool`

HasDisplayVariant returns a boolean if a field has been set.

### GetDisplayName

`func (o *Locale) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Locale) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Locale) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Locale) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLanguage

`func (o *Locale) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Locale) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Locale) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Locale) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


