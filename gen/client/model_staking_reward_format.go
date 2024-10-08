/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// StakingRewardFormat The format in which the rewards are to be fetched i.e native or in equivalent USD
type StakingRewardFormat string

// List of StakingRewardFormat
const (
	STAKINGREWARDFORMAT_USD StakingRewardFormat = "usd"
	STAKINGREWARDFORMAT_NATIVE StakingRewardFormat = "native"
)

// All allowed values of StakingRewardFormat enum
var AllowedStakingRewardFormatEnumValues = []StakingRewardFormat{
	"usd",
	"native",
}

func (v *StakingRewardFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StakingRewardFormat(value)
	for _, existing := range AllowedStakingRewardFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StakingRewardFormat", value)
}

// NewStakingRewardFormatFromValue returns a pointer to a valid StakingRewardFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStakingRewardFormatFromValue(v string) (*StakingRewardFormat, error) {
	ev := StakingRewardFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StakingRewardFormat: valid values are %v", v, AllowedStakingRewardFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StakingRewardFormat) IsValid() bool {
	for _, existing := range AllowedStakingRewardFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StakingRewardFormat value
func (v StakingRewardFormat) Ptr() *StakingRewardFormat {
	return &v
}

type NullableStakingRewardFormat struct {
	value *StakingRewardFormat
	isSet bool
}

func (v NullableStakingRewardFormat) Get() *StakingRewardFormat {
	return v.value
}

func (v *NullableStakingRewardFormat) Set(val *StakingRewardFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableStakingRewardFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableStakingRewardFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakingRewardFormat(val *StakingRewardFormat) *NullableStakingRewardFormat {
	return &NullableStakingRewardFormat{value: val, isSet: true}
}

func (v NullableStakingRewardFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakingRewardFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

