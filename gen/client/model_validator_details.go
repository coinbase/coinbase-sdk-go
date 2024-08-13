/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
Contact: yuga.cohler@coinbase.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// ValidatorDetails - struct for ValidatorDetails
type ValidatorDetails struct {
	EthereumValidatorMetadata *EthereumValidatorMetadata
}

// EthereumValidatorMetadataAsValidatorDetails is a convenience function that returns EthereumValidatorMetadata wrapped in ValidatorDetails
func EthereumValidatorMetadataAsValidatorDetails(v *EthereumValidatorMetadata) ValidatorDetails {
	return ValidatorDetails{
		EthereumValidatorMetadata: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ValidatorDetails) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EthereumValidatorMetadata
	err = newStrictDecoder(data).Decode(&dst.EthereumValidatorMetadata)
	if err == nil {
		jsonEthereumValidatorMetadata, _ := json.Marshal(dst.EthereumValidatorMetadata)
		if string(jsonEthereumValidatorMetadata) == "{}" { // empty struct
			dst.EthereumValidatorMetadata = nil
		} else {
			match++
		}
	} else {
		dst.EthereumValidatorMetadata = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EthereumValidatorMetadata = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ValidatorDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ValidatorDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ValidatorDetails) MarshalJSON() ([]byte, error) {
	if src.EthereumValidatorMetadata != nil {
		return json.Marshal(&src.EthereumValidatorMetadata)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ValidatorDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EthereumValidatorMetadata != nil {
		return obj.EthereumValidatorMetadata
	}

	// all schemas are nil
	return nil
}

type NullableValidatorDetails struct {
	value *ValidatorDetails
	isSet bool
}

func (v NullableValidatorDetails) Get() *ValidatorDetails {
	return v.value
}

func (v *NullableValidatorDetails) Set(val *ValidatorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatorDetails(val *ValidatorDetails) *NullableValidatorDetails {
	return &NullableValidatorDetails{value: val, isSet: true}
}

func (v NullableValidatorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


