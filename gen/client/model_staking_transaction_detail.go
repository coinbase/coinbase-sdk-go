/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// StakingTransactionDetail - A network specific staking transaction containing both tx related details and non-tx related details.
type StakingTransactionDetail struct {
	DedicatedEthStakeTransactionDetail *DedicatedEthStakeTransactionDetail
	DedicatedEthUnstakeTransactionDetail *DedicatedEthUnstakeTransactionDetail
	PartialEthStakingTransactionDetail *PartialEthStakingTransactionDetail
	SolanaStakingTransactionDetail *SolanaStakingTransactionDetail
}

// DedicatedEthStakeTransactionDetailAsStakingTransactionDetail is a convenience function that returns DedicatedEthStakeTransactionDetail wrapped in StakingTransactionDetail
func DedicatedEthStakeTransactionDetailAsStakingTransactionDetail(v *DedicatedEthStakeTransactionDetail) StakingTransactionDetail {
	return StakingTransactionDetail{
		DedicatedEthStakeTransactionDetail: v,
	}
}

// DedicatedEthUnstakeTransactionDetailAsStakingTransactionDetail is a convenience function that returns DedicatedEthUnstakeTransactionDetail wrapped in StakingTransactionDetail
func DedicatedEthUnstakeTransactionDetailAsStakingTransactionDetail(v *DedicatedEthUnstakeTransactionDetail) StakingTransactionDetail {
	return StakingTransactionDetail{
		DedicatedEthUnstakeTransactionDetail: v,
	}
}

// PartialEthStakingTransactionDetailAsStakingTransactionDetail is a convenience function that returns PartialEthStakingTransactionDetail wrapped in StakingTransactionDetail
func PartialEthStakingTransactionDetailAsStakingTransactionDetail(v *PartialEthStakingTransactionDetail) StakingTransactionDetail {
	return StakingTransactionDetail{
		PartialEthStakingTransactionDetail: v,
	}
}

// SolanaStakingTransactionDetailAsStakingTransactionDetail is a convenience function that returns SolanaStakingTransactionDetail wrapped in StakingTransactionDetail
func SolanaStakingTransactionDetailAsStakingTransactionDetail(v *SolanaStakingTransactionDetail) StakingTransactionDetail {
	return StakingTransactionDetail{
		SolanaStakingTransactionDetail: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *StakingTransactionDetail) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DedicatedEthStakeTransactionDetail
	err = newStrictDecoder(data).Decode(&dst.DedicatedEthStakeTransactionDetail)
	if err == nil {
		jsonDedicatedEthStakeTransactionDetail, _ := json.Marshal(dst.DedicatedEthStakeTransactionDetail)
		if string(jsonDedicatedEthStakeTransactionDetail) == "{}" { // empty struct
			dst.DedicatedEthStakeTransactionDetail = nil
		} else {
			if err = validator.Validate(dst.DedicatedEthStakeTransactionDetail); err != nil {
				dst.DedicatedEthStakeTransactionDetail = nil
			} else {
				match++
			}
		}
	} else {
		dst.DedicatedEthStakeTransactionDetail = nil
	}

	// try to unmarshal data into DedicatedEthUnstakeTransactionDetail
	err = newStrictDecoder(data).Decode(&dst.DedicatedEthUnstakeTransactionDetail)
	if err == nil {
		jsonDedicatedEthUnstakeTransactionDetail, _ := json.Marshal(dst.DedicatedEthUnstakeTransactionDetail)
		if string(jsonDedicatedEthUnstakeTransactionDetail) == "{}" { // empty struct
			dst.DedicatedEthUnstakeTransactionDetail = nil
		} else {
			if err = validator.Validate(dst.DedicatedEthUnstakeTransactionDetail); err != nil {
				dst.DedicatedEthUnstakeTransactionDetail = nil
			} else {
				match++
			}
		}
	} else {
		dst.DedicatedEthUnstakeTransactionDetail = nil
	}

	// try to unmarshal data into PartialEthStakingTransactionDetail
	err = newStrictDecoder(data).Decode(&dst.PartialEthStakingTransactionDetail)
	if err == nil {
		jsonPartialEthStakingTransactionDetail, _ := json.Marshal(dst.PartialEthStakingTransactionDetail)
		if string(jsonPartialEthStakingTransactionDetail) == "{}" { // empty struct
			dst.PartialEthStakingTransactionDetail = nil
		} else {
			if err = validator.Validate(dst.PartialEthStakingTransactionDetail); err != nil {
				dst.PartialEthStakingTransactionDetail = nil
			} else {
				match++
			}
		}
	} else {
		dst.PartialEthStakingTransactionDetail = nil
	}

	// try to unmarshal data into SolanaStakingTransactionDetail
	err = newStrictDecoder(data).Decode(&dst.SolanaStakingTransactionDetail)
	if err == nil {
		jsonSolanaStakingTransactionDetail, _ := json.Marshal(dst.SolanaStakingTransactionDetail)
		if string(jsonSolanaStakingTransactionDetail) == "{}" { // empty struct
			dst.SolanaStakingTransactionDetail = nil
		} else {
			if err = validator.Validate(dst.SolanaStakingTransactionDetail); err != nil {
				dst.SolanaStakingTransactionDetail = nil
			} else {
				match++
			}
		}
	} else {
		dst.SolanaStakingTransactionDetail = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DedicatedEthStakeTransactionDetail = nil
		dst.DedicatedEthUnstakeTransactionDetail = nil
		dst.PartialEthStakingTransactionDetail = nil
		dst.SolanaStakingTransactionDetail = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StakingTransactionDetail)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StakingTransactionDetail)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StakingTransactionDetail) MarshalJSON() ([]byte, error) {
	if src.DedicatedEthStakeTransactionDetail != nil {
		return json.Marshal(&src.DedicatedEthStakeTransactionDetail)
	}

	if src.DedicatedEthUnstakeTransactionDetail != nil {
		return json.Marshal(&src.DedicatedEthUnstakeTransactionDetail)
	}

	if src.PartialEthStakingTransactionDetail != nil {
		return json.Marshal(&src.PartialEthStakingTransactionDetail)
	}

	if src.SolanaStakingTransactionDetail != nil {
		return json.Marshal(&src.SolanaStakingTransactionDetail)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StakingTransactionDetail) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DedicatedEthStakeTransactionDetail != nil {
		return obj.DedicatedEthStakeTransactionDetail
	}

	if obj.DedicatedEthUnstakeTransactionDetail != nil {
		return obj.DedicatedEthUnstakeTransactionDetail
	}

	if obj.PartialEthStakingTransactionDetail != nil {
		return obj.PartialEthStakingTransactionDetail
	}

	if obj.SolanaStakingTransactionDetail != nil {
		return obj.SolanaStakingTransactionDetail
	}

	// all schemas are nil
	return nil
}

type NullableStakingTransactionDetail struct {
	value *StakingTransactionDetail
	isSet bool
}

func (v NullableStakingTransactionDetail) Get() *StakingTransactionDetail {
	return v.value
}

func (v *NullableStakingTransactionDetail) Set(val *StakingTransactionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableStakingTransactionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableStakingTransactionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakingTransactionDetail(val *StakingTransactionDetail) *NullableStakingTransactionDetail {
	return &NullableStakingTransactionDetail{value: val, isSet: true}
}

func (v NullableStakingTransactionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakingTransactionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

