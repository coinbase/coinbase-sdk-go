/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SolidityValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SolidityValue{}

// SolidityValue struct for SolidityValue
type SolidityValue struct {
	Type string `json:"type"`
	// The field name for tuple types. Not used for other types.
	Name *string `json:"name,omitempty"`
	// The value as a string for simple types. Not used for complex types (array, tuple).
	Value *string `json:"value,omitempty"`
	// For array and tuple types, the components of the value
	Values []SolidityValue `json:"values,omitempty"`
}

type _SolidityValue SolidityValue

// NewSolidityValue instantiates a new SolidityValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolidityValue(type_ string) *SolidityValue {
	this := SolidityValue{}
	this.Type = type_
	return &this
}

// NewSolidityValueWithDefaults instantiates a new SolidityValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolidityValueWithDefaults() *SolidityValue {
	this := SolidityValue{}
	return &this
}

// GetType returns the Type field value
func (o *SolidityValue) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SolidityValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SolidityValue) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SolidityValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SolidityValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SolidityValue) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SolidityValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SolidityValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SolidityValue) SetValue(v string) {
	o.Value = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SolidityValue) GetValues() []SolidityValue {
	if o == nil || IsNil(o.Values) {
		var ret []SolidityValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityValue) GetValuesOk() ([]SolidityValue, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SolidityValue) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []SolidityValue and assigns it to the Values field.
func (o *SolidityValue) SetValues(v []SolidityValue) {
	o.Values = v
}

func (o SolidityValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SolidityValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

func (o *SolidityValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSolidityValue := _SolidityValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSolidityValue)

	if err != nil {
		return err
	}

	*o = SolidityValue(varSolidityValue)

	return err
}

type NullableSolidityValue struct {
	value *SolidityValue
	isSet bool
}

func (v NullableSolidityValue) Get() *SolidityValue {
	return v.value
}

func (v *NullableSolidityValue) Set(val *SolidityValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSolidityValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSolidityValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolidityValue(val *SolidityValue) *NullableSolidityValue {
	return &NullableSolidityValue{value: val, isSet: true}
}

func (v NullableSolidityValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolidityValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


