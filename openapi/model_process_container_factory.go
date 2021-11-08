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

// ProcessContainerFactory struct for ProcessContainerFactory
type ProcessContainerFactory struct {
	JavaInstanceJarLocation *string `json:"javaInstanceJarLocation,omitempty"`
	PythonInstanceLocation *string `json:"pythonInstanceLocation,omitempty"`
	LogDirectory *string `json:"logDirectory,omitempty"`
	ExtraFunctionDependenciesDir *string `json:"extraFunctionDependenciesDir,omitempty"`
}

// NewProcessContainerFactory instantiates a new ProcessContainerFactory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessContainerFactory() *ProcessContainerFactory {
	this := ProcessContainerFactory{}
	return &this
}

// NewProcessContainerFactoryWithDefaults instantiates a new ProcessContainerFactory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessContainerFactoryWithDefaults() *ProcessContainerFactory {
	this := ProcessContainerFactory{}
	return &this
}

// GetJavaInstanceJarLocation returns the JavaInstanceJarLocation field value if set, zero value otherwise.
func (o *ProcessContainerFactory) GetJavaInstanceJarLocation() string {
	if o == nil || o.JavaInstanceJarLocation == nil {
		var ret string
		return ret
	}
	return *o.JavaInstanceJarLocation
}

// GetJavaInstanceJarLocationOk returns a tuple with the JavaInstanceJarLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessContainerFactory) GetJavaInstanceJarLocationOk() (*string, bool) {
	if o == nil || o.JavaInstanceJarLocation == nil {
		return nil, false
	}
	return o.JavaInstanceJarLocation, true
}

// HasJavaInstanceJarLocation returns a boolean if a field has been set.
func (o *ProcessContainerFactory) HasJavaInstanceJarLocation() bool {
	if o != nil && o.JavaInstanceJarLocation != nil {
		return true
	}

	return false
}

// SetJavaInstanceJarLocation gets a reference to the given string and assigns it to the JavaInstanceJarLocation field.
func (o *ProcessContainerFactory) SetJavaInstanceJarLocation(v string) {
	o.JavaInstanceJarLocation = &v
}

// GetPythonInstanceLocation returns the PythonInstanceLocation field value if set, zero value otherwise.
func (o *ProcessContainerFactory) GetPythonInstanceLocation() string {
	if o == nil || o.PythonInstanceLocation == nil {
		var ret string
		return ret
	}
	return *o.PythonInstanceLocation
}

// GetPythonInstanceLocationOk returns a tuple with the PythonInstanceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessContainerFactory) GetPythonInstanceLocationOk() (*string, bool) {
	if o == nil || o.PythonInstanceLocation == nil {
		return nil, false
	}
	return o.PythonInstanceLocation, true
}

// HasPythonInstanceLocation returns a boolean if a field has been set.
func (o *ProcessContainerFactory) HasPythonInstanceLocation() bool {
	if o != nil && o.PythonInstanceLocation != nil {
		return true
	}

	return false
}

// SetPythonInstanceLocation gets a reference to the given string and assigns it to the PythonInstanceLocation field.
func (o *ProcessContainerFactory) SetPythonInstanceLocation(v string) {
	o.PythonInstanceLocation = &v
}

// GetLogDirectory returns the LogDirectory field value if set, zero value otherwise.
func (o *ProcessContainerFactory) GetLogDirectory() string {
	if o == nil || o.LogDirectory == nil {
		var ret string
		return ret
	}
	return *o.LogDirectory
}

// GetLogDirectoryOk returns a tuple with the LogDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessContainerFactory) GetLogDirectoryOk() (*string, bool) {
	if o == nil || o.LogDirectory == nil {
		return nil, false
	}
	return o.LogDirectory, true
}

// HasLogDirectory returns a boolean if a field has been set.
func (o *ProcessContainerFactory) HasLogDirectory() bool {
	if o != nil && o.LogDirectory != nil {
		return true
	}

	return false
}

// SetLogDirectory gets a reference to the given string and assigns it to the LogDirectory field.
func (o *ProcessContainerFactory) SetLogDirectory(v string) {
	o.LogDirectory = &v
}

// GetExtraFunctionDependenciesDir returns the ExtraFunctionDependenciesDir field value if set, zero value otherwise.
func (o *ProcessContainerFactory) GetExtraFunctionDependenciesDir() string {
	if o == nil || o.ExtraFunctionDependenciesDir == nil {
		var ret string
		return ret
	}
	return *o.ExtraFunctionDependenciesDir
}

// GetExtraFunctionDependenciesDirOk returns a tuple with the ExtraFunctionDependenciesDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessContainerFactory) GetExtraFunctionDependenciesDirOk() (*string, bool) {
	if o == nil || o.ExtraFunctionDependenciesDir == nil {
		return nil, false
	}
	return o.ExtraFunctionDependenciesDir, true
}

// HasExtraFunctionDependenciesDir returns a boolean if a field has been set.
func (o *ProcessContainerFactory) HasExtraFunctionDependenciesDir() bool {
	if o != nil && o.ExtraFunctionDependenciesDir != nil {
		return true
	}

	return false
}

// SetExtraFunctionDependenciesDir gets a reference to the given string and assigns it to the ExtraFunctionDependenciesDir field.
func (o *ProcessContainerFactory) SetExtraFunctionDependenciesDir(v string) {
	o.ExtraFunctionDependenciesDir = &v
}

func (o ProcessContainerFactory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JavaInstanceJarLocation != nil {
		toSerialize["javaInstanceJarLocation"] = o.JavaInstanceJarLocation
	}
	if o.PythonInstanceLocation != nil {
		toSerialize["pythonInstanceLocation"] = o.PythonInstanceLocation
	}
	if o.LogDirectory != nil {
		toSerialize["logDirectory"] = o.LogDirectory
	}
	if o.ExtraFunctionDependenciesDir != nil {
		toSerialize["extraFunctionDependenciesDir"] = o.ExtraFunctionDependenciesDir
	}
	return json.Marshal(toSerialize)
}

type NullableProcessContainerFactory struct {
	value *ProcessContainerFactory
	isSet bool
}

func (v NullableProcessContainerFactory) Get() *ProcessContainerFactory {
	return v.value
}

func (v *NullableProcessContainerFactory) Set(val *ProcessContainerFactory) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessContainerFactory) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessContainerFactory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessContainerFactory(val *ProcessContainerFactory) *NullableProcessContainerFactory {
	return &NullableProcessContainerFactory{value: val, isSet: true}
}

func (v NullableProcessContainerFactory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessContainerFactory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


