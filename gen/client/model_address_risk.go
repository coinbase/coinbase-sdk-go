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

// checks if the AddressRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressRisk{}

// AddressRisk The risk score of a blockchain address.
type AddressRisk struct {
	// The lower the score is, the higher the risk is. The score lies between -100 to 0.
	RiskScore int32 `json:"risk_score"`
}

type _AddressRisk AddressRisk

// NewAddressRisk instantiates a new AddressRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressRisk(riskScore int32) *AddressRisk {
	this := AddressRisk{}
	this.RiskScore = riskScore
	return &this
}

// NewAddressRiskWithDefaults instantiates a new AddressRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressRiskWithDefaults() *AddressRisk {
	this := AddressRisk{}
	return &this
}

// GetRiskScore returns the RiskScore field value
func (o *AddressRisk) GetRiskScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value
// and a boolean to check if the value has been set.
func (o *AddressRisk) GetRiskScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskScore, true
}

// SetRiskScore sets field value
func (o *AddressRisk) SetRiskScore(v int32) {
	o.RiskScore = v
}

func (o AddressRisk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["risk_score"] = o.RiskScore
	return toSerialize, nil
}

func (o *AddressRisk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"risk_score",
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

	varAddressRisk := _AddressRisk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressRisk)

	if err != nil {
		return err
	}

	*o = AddressRisk(varAddressRisk)

	return err
}

type NullableAddressRisk struct {
	value *AddressRisk
	isSet bool
}

func (v NullableAddressRisk) Get() *AddressRisk {
	return v.value
}

func (v *NullableAddressRisk) Set(val *AddressRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressRisk(val *AddressRisk) *NullableAddressRisk {
	return &NullableAddressRisk{value: val, isSet: true}
}

func (v NullableAddressRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


