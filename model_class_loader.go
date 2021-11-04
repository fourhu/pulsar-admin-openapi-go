/*
Pulsar Packages REST API

This provides the REST API for Pulsar Packages operations

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ClassLoader struct for ClassLoader
type ClassLoader struct {
	Parent *ClassLoader `json:"parent,omitempty"`
	Name *string `json:"name,omitempty"`
	UnnamedModule *Module `json:"unnamedModule,omitempty"`
	RegisteredAsParallelCapable *bool `json:"registeredAsParallelCapable,omitempty"`
	DefinedPackages *[]Package `json:"definedPackages,omitempty"`
}

// NewClassLoader instantiates a new ClassLoader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassLoader() *ClassLoader {
	this := ClassLoader{}
	return &this
}

// NewClassLoaderWithDefaults instantiates a new ClassLoader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassLoaderWithDefaults() *ClassLoader {
	this := ClassLoader{}
	return &this
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ClassLoader) GetParent() ClassLoader {
	if o == nil || o.Parent == nil {
		var ret ClassLoader
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLoader) GetParentOk() (*ClassLoader, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ClassLoader) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given ClassLoader and assigns it to the Parent field.
func (o *ClassLoader) SetParent(v ClassLoader) {
	o.Parent = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClassLoader) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLoader) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClassLoader) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClassLoader) SetName(v string) {
	o.Name = &v
}

// GetUnnamedModule returns the UnnamedModule field value if set, zero value otherwise.
func (o *ClassLoader) GetUnnamedModule() Module {
	if o == nil || o.UnnamedModule == nil {
		var ret Module
		return ret
	}
	return *o.UnnamedModule
}

// GetUnnamedModuleOk returns a tuple with the UnnamedModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLoader) GetUnnamedModuleOk() (*Module, bool) {
	if o == nil || o.UnnamedModule == nil {
		return nil, false
	}
	return o.UnnamedModule, true
}

// HasUnnamedModule returns a boolean if a field has been set.
func (o *ClassLoader) HasUnnamedModule() bool {
	if o != nil && o.UnnamedModule != nil {
		return true
	}

	return false
}

// SetUnnamedModule gets a reference to the given Module and assigns it to the UnnamedModule field.
func (o *ClassLoader) SetUnnamedModule(v Module) {
	o.UnnamedModule = &v
}

// GetRegisteredAsParallelCapable returns the RegisteredAsParallelCapable field value if set, zero value otherwise.
func (o *ClassLoader) GetRegisteredAsParallelCapable() bool {
	if o == nil || o.RegisteredAsParallelCapable == nil {
		var ret bool
		return ret
	}
	return *o.RegisteredAsParallelCapable
}

// GetRegisteredAsParallelCapableOk returns a tuple with the RegisteredAsParallelCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLoader) GetRegisteredAsParallelCapableOk() (*bool, bool) {
	if o == nil || o.RegisteredAsParallelCapable == nil {
		return nil, false
	}
	return o.RegisteredAsParallelCapable, true
}

// HasRegisteredAsParallelCapable returns a boolean if a field has been set.
func (o *ClassLoader) HasRegisteredAsParallelCapable() bool {
	if o != nil && o.RegisteredAsParallelCapable != nil {
		return true
	}

	return false
}

// SetRegisteredAsParallelCapable gets a reference to the given bool and assigns it to the RegisteredAsParallelCapable field.
func (o *ClassLoader) SetRegisteredAsParallelCapable(v bool) {
	o.RegisteredAsParallelCapable = &v
}

// GetDefinedPackages returns the DefinedPackages field value if set, zero value otherwise.
func (o *ClassLoader) GetDefinedPackages() []Package {
	if o == nil || o.DefinedPackages == nil {
		var ret []Package
		return ret
	}
	return *o.DefinedPackages
}

// GetDefinedPackagesOk returns a tuple with the DefinedPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassLoader) GetDefinedPackagesOk() (*[]Package, bool) {
	if o == nil || o.DefinedPackages == nil {
		return nil, false
	}
	return o.DefinedPackages, true
}

// HasDefinedPackages returns a boolean if a field has been set.
func (o *ClassLoader) HasDefinedPackages() bool {
	if o != nil && o.DefinedPackages != nil {
		return true
	}

	return false
}

// SetDefinedPackages gets a reference to the given []Package and assigns it to the DefinedPackages field.
func (o *ClassLoader) SetDefinedPackages(v []Package) {
	o.DefinedPackages = &v
}

func (o ClassLoader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UnnamedModule != nil {
		toSerialize["unnamedModule"] = o.UnnamedModule
	}
	if o.RegisteredAsParallelCapable != nil {
		toSerialize["registeredAsParallelCapable"] = o.RegisteredAsParallelCapable
	}
	if o.DefinedPackages != nil {
		toSerialize["definedPackages"] = o.DefinedPackages
	}
	return json.Marshal(toSerialize)
}

type NullableClassLoader struct {
	value *ClassLoader
	isSet bool
}

func (v NullableClassLoader) Get() *ClassLoader {
	return v.value
}

func (v *NullableClassLoader) Set(val *ClassLoader) {
	v.value = val
	v.isSet = true
}

func (v NullableClassLoader) IsSet() bool {
	return v.isSet
}

func (v *NullableClassLoader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassLoader(val *ClassLoader) *NullableClassLoader {
	return &NullableClassLoader{value: val, isSet: true}
}

func (v NullableClassLoader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassLoader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


