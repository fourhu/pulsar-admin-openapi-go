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

// Locale struct for Locale
type Locale struct {
	Script *string `json:"script,omitempty"`
	Country *string `json:"country,omitempty"`
	Variant *string `json:"variant,omitempty"`
	ExtensionKeys *[]string `json:"extensionKeys,omitempty"`
	UnicodeLocaleAttributes *[]string `json:"unicodeLocaleAttributes,omitempty"`
	UnicodeLocaleKeys *[]string `json:"unicodeLocaleKeys,omitempty"`
	Iso3Language *string `json:"iso3Language,omitempty"`
	Iso3Country *string `json:"iso3Country,omitempty"`
	DisplayLanguage *string `json:"displayLanguage,omitempty"`
	DisplayScript *string `json:"displayScript,omitempty"`
	DisplayCountry *string `json:"displayCountry,omitempty"`
	DisplayVariant *string `json:"displayVariant,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Language *string `json:"language,omitempty"`
}

// NewLocale instantiates a new Locale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocale() *Locale {
	this := Locale{}
	return &this
}

// NewLocaleWithDefaults instantiates a new Locale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocaleWithDefaults() *Locale {
	this := Locale{}
	return &this
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *Locale) GetScript() string {
	if o == nil || o.Script == nil {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetScriptOk() (*string, bool) {
	if o == nil || o.Script == nil {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *Locale) HasScript() bool {
	if o != nil && o.Script != nil {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *Locale) SetScript(v string) {
	o.Script = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Locale) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Locale) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Locale) SetCountry(v string) {
	o.Country = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *Locale) GetVariant() string {
	if o == nil || o.Variant == nil {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetVariantOk() (*string, bool) {
	if o == nil || o.Variant == nil {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *Locale) HasVariant() bool {
	if o != nil && o.Variant != nil {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *Locale) SetVariant(v string) {
	o.Variant = &v
}

// GetExtensionKeys returns the ExtensionKeys field value if set, zero value otherwise.
func (o *Locale) GetExtensionKeys() []string {
	if o == nil || o.ExtensionKeys == nil {
		var ret []string
		return ret
	}
	return *o.ExtensionKeys
}

// GetExtensionKeysOk returns a tuple with the ExtensionKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetExtensionKeysOk() (*[]string, bool) {
	if o == nil || o.ExtensionKeys == nil {
		return nil, false
	}
	return o.ExtensionKeys, true
}

// HasExtensionKeys returns a boolean if a field has been set.
func (o *Locale) HasExtensionKeys() bool {
	if o != nil && o.ExtensionKeys != nil {
		return true
	}

	return false
}

// SetExtensionKeys gets a reference to the given []string and assigns it to the ExtensionKeys field.
func (o *Locale) SetExtensionKeys(v []string) {
	o.ExtensionKeys = &v
}

// GetUnicodeLocaleAttributes returns the UnicodeLocaleAttributes field value if set, zero value otherwise.
func (o *Locale) GetUnicodeLocaleAttributes() []string {
	if o == nil || o.UnicodeLocaleAttributes == nil {
		var ret []string
		return ret
	}
	return *o.UnicodeLocaleAttributes
}

// GetUnicodeLocaleAttributesOk returns a tuple with the UnicodeLocaleAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetUnicodeLocaleAttributesOk() (*[]string, bool) {
	if o == nil || o.UnicodeLocaleAttributes == nil {
		return nil, false
	}
	return o.UnicodeLocaleAttributes, true
}

// HasUnicodeLocaleAttributes returns a boolean if a field has been set.
func (o *Locale) HasUnicodeLocaleAttributes() bool {
	if o != nil && o.UnicodeLocaleAttributes != nil {
		return true
	}

	return false
}

// SetUnicodeLocaleAttributes gets a reference to the given []string and assigns it to the UnicodeLocaleAttributes field.
func (o *Locale) SetUnicodeLocaleAttributes(v []string) {
	o.UnicodeLocaleAttributes = &v
}

// GetUnicodeLocaleKeys returns the UnicodeLocaleKeys field value if set, zero value otherwise.
func (o *Locale) GetUnicodeLocaleKeys() []string {
	if o == nil || o.UnicodeLocaleKeys == nil {
		var ret []string
		return ret
	}
	return *o.UnicodeLocaleKeys
}

// GetUnicodeLocaleKeysOk returns a tuple with the UnicodeLocaleKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetUnicodeLocaleKeysOk() (*[]string, bool) {
	if o == nil || o.UnicodeLocaleKeys == nil {
		return nil, false
	}
	return o.UnicodeLocaleKeys, true
}

// HasUnicodeLocaleKeys returns a boolean if a field has been set.
func (o *Locale) HasUnicodeLocaleKeys() bool {
	if o != nil && o.UnicodeLocaleKeys != nil {
		return true
	}

	return false
}

// SetUnicodeLocaleKeys gets a reference to the given []string and assigns it to the UnicodeLocaleKeys field.
func (o *Locale) SetUnicodeLocaleKeys(v []string) {
	o.UnicodeLocaleKeys = &v
}

// GetIso3Language returns the Iso3Language field value if set, zero value otherwise.
func (o *Locale) GetIso3Language() string {
	if o == nil || o.Iso3Language == nil {
		var ret string
		return ret
	}
	return *o.Iso3Language
}

// GetIso3LanguageOk returns a tuple with the Iso3Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetIso3LanguageOk() (*string, bool) {
	if o == nil || o.Iso3Language == nil {
		return nil, false
	}
	return o.Iso3Language, true
}

// HasIso3Language returns a boolean if a field has been set.
func (o *Locale) HasIso3Language() bool {
	if o != nil && o.Iso3Language != nil {
		return true
	}

	return false
}

// SetIso3Language gets a reference to the given string and assigns it to the Iso3Language field.
func (o *Locale) SetIso3Language(v string) {
	o.Iso3Language = &v
}

// GetIso3Country returns the Iso3Country field value if set, zero value otherwise.
func (o *Locale) GetIso3Country() string {
	if o == nil || o.Iso3Country == nil {
		var ret string
		return ret
	}
	return *o.Iso3Country
}

// GetIso3CountryOk returns a tuple with the Iso3Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetIso3CountryOk() (*string, bool) {
	if o == nil || o.Iso3Country == nil {
		return nil, false
	}
	return o.Iso3Country, true
}

// HasIso3Country returns a boolean if a field has been set.
func (o *Locale) HasIso3Country() bool {
	if o != nil && o.Iso3Country != nil {
		return true
	}

	return false
}

// SetIso3Country gets a reference to the given string and assigns it to the Iso3Country field.
func (o *Locale) SetIso3Country(v string) {
	o.Iso3Country = &v
}

// GetDisplayLanguage returns the DisplayLanguage field value if set, zero value otherwise.
func (o *Locale) GetDisplayLanguage() string {
	if o == nil || o.DisplayLanguage == nil {
		var ret string
		return ret
	}
	return *o.DisplayLanguage
}

// GetDisplayLanguageOk returns a tuple with the DisplayLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetDisplayLanguageOk() (*string, bool) {
	if o == nil || o.DisplayLanguage == nil {
		return nil, false
	}
	return o.DisplayLanguage, true
}

// HasDisplayLanguage returns a boolean if a field has been set.
func (o *Locale) HasDisplayLanguage() bool {
	if o != nil && o.DisplayLanguage != nil {
		return true
	}

	return false
}

// SetDisplayLanguage gets a reference to the given string and assigns it to the DisplayLanguage field.
func (o *Locale) SetDisplayLanguage(v string) {
	o.DisplayLanguage = &v
}

// GetDisplayScript returns the DisplayScript field value if set, zero value otherwise.
func (o *Locale) GetDisplayScript() string {
	if o == nil || o.DisplayScript == nil {
		var ret string
		return ret
	}
	return *o.DisplayScript
}

// GetDisplayScriptOk returns a tuple with the DisplayScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetDisplayScriptOk() (*string, bool) {
	if o == nil || o.DisplayScript == nil {
		return nil, false
	}
	return o.DisplayScript, true
}

// HasDisplayScript returns a boolean if a field has been set.
func (o *Locale) HasDisplayScript() bool {
	if o != nil && o.DisplayScript != nil {
		return true
	}

	return false
}

// SetDisplayScript gets a reference to the given string and assigns it to the DisplayScript field.
func (o *Locale) SetDisplayScript(v string) {
	o.DisplayScript = &v
}

// GetDisplayCountry returns the DisplayCountry field value if set, zero value otherwise.
func (o *Locale) GetDisplayCountry() string {
	if o == nil || o.DisplayCountry == nil {
		var ret string
		return ret
	}
	return *o.DisplayCountry
}

// GetDisplayCountryOk returns a tuple with the DisplayCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetDisplayCountryOk() (*string, bool) {
	if o == nil || o.DisplayCountry == nil {
		return nil, false
	}
	return o.DisplayCountry, true
}

// HasDisplayCountry returns a boolean if a field has been set.
func (o *Locale) HasDisplayCountry() bool {
	if o != nil && o.DisplayCountry != nil {
		return true
	}

	return false
}

// SetDisplayCountry gets a reference to the given string and assigns it to the DisplayCountry field.
func (o *Locale) SetDisplayCountry(v string) {
	o.DisplayCountry = &v
}

// GetDisplayVariant returns the DisplayVariant field value if set, zero value otherwise.
func (o *Locale) GetDisplayVariant() string {
	if o == nil || o.DisplayVariant == nil {
		var ret string
		return ret
	}
	return *o.DisplayVariant
}

// GetDisplayVariantOk returns a tuple with the DisplayVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetDisplayVariantOk() (*string, bool) {
	if o == nil || o.DisplayVariant == nil {
		return nil, false
	}
	return o.DisplayVariant, true
}

// HasDisplayVariant returns a boolean if a field has been set.
func (o *Locale) HasDisplayVariant() bool {
	if o != nil && o.DisplayVariant != nil {
		return true
	}

	return false
}

// SetDisplayVariant gets a reference to the given string and assigns it to the DisplayVariant field.
func (o *Locale) SetDisplayVariant(v string) {
	o.DisplayVariant = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Locale) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Locale) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Locale) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Locale) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Locale) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Locale) SetLanguage(v string) {
	o.Language = &v
}

func (o Locale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Script != nil {
		toSerialize["script"] = o.Script
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Variant != nil {
		toSerialize["variant"] = o.Variant
	}
	if o.ExtensionKeys != nil {
		toSerialize["extensionKeys"] = o.ExtensionKeys
	}
	if o.UnicodeLocaleAttributes != nil {
		toSerialize["unicodeLocaleAttributes"] = o.UnicodeLocaleAttributes
	}
	if o.UnicodeLocaleKeys != nil {
		toSerialize["unicodeLocaleKeys"] = o.UnicodeLocaleKeys
	}
	if o.Iso3Language != nil {
		toSerialize["iso3Language"] = o.Iso3Language
	}
	if o.Iso3Country != nil {
		toSerialize["iso3Country"] = o.Iso3Country
	}
	if o.DisplayLanguage != nil {
		toSerialize["displayLanguage"] = o.DisplayLanguage
	}
	if o.DisplayScript != nil {
		toSerialize["displayScript"] = o.DisplayScript
	}
	if o.DisplayCountry != nil {
		toSerialize["displayCountry"] = o.DisplayCountry
	}
	if o.DisplayVariant != nil {
		toSerialize["displayVariant"] = o.DisplayVariant
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	return json.Marshal(toSerialize)
}

type NullableLocale struct {
	value *Locale
	isSet bool
}

func (v NullableLocale) Get() *Locale {
	return v.value
}

func (v *NullableLocale) Set(val *Locale) {
	v.value = val
	v.isSet = true
}

func (v NullableLocale) IsSet() bool {
	return v.isSet
}

func (v *NullableLocale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocale(val *Locale) *NullableLocale {
	return &NullableLocale{value: val, isSet: true}
}

func (v NullableLocale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


