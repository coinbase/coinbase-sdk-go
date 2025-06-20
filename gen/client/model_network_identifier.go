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

// NetworkIdentifier The ID of the blockchain network. This is unique across all networks, and takes the form of `<blockchain>-<network>`.
type NetworkIdentifier string

// List of NetworkIdentifier
const (
	NETWORKIDENTIFIER_BASE_SEPOLIA NetworkIdentifier = "base-sepolia"
	NETWORKIDENTIFIER_BASE_MAINNET NetworkIdentifier = "base-mainnet"
	NETWORKIDENTIFIER_ETHEREUM_HOODI NetworkIdentifier = "ethereum-hoodi"
	NETWORKIDENTIFIER_ETHEREUM_HOODI2 NetworkIdentifier = "ethereum-hoodi"
	NETWORKIDENTIFIER_ETHEREUM_SEPOLIA NetworkIdentifier = "ethereum-sepolia"
	NETWORKIDENTIFIER_ETHEREUM_MAINNET NetworkIdentifier = "ethereum-mainnet"
	NETWORKIDENTIFIER_POLYGON_MAINNET NetworkIdentifier = "polygon-mainnet"
	NETWORKIDENTIFIER_SOLANA_DEVNET NetworkIdentifier = "solana-devnet"
	NETWORKIDENTIFIER_SOLANA_MAINNET NetworkIdentifier = "solana-mainnet"
	NETWORKIDENTIFIER_ARBITRUM_MAINNET NetworkIdentifier = "arbitrum-mainnet"
	NETWORKIDENTIFIER_ARBITRUM_SEPOLIA NetworkIdentifier = "arbitrum-sepolia"
	NETWORKIDENTIFIER_BITCOIN_MAINNET NetworkIdentifier = "bitcoin-mainnet"
	NETWORKIDENTIFIER_NEAR_TESTNET NetworkIdentifier = "near-testnet"
	NETWORKIDENTIFIER_NEAR_MAINNET NetworkIdentifier = "near-mainnet"
)

// All allowed values of NetworkIdentifier enum
var AllowedNetworkIdentifierEnumValues = []NetworkIdentifier{
	"base-sepolia",
	"base-mainnet",
	"ethereum-hoodi",
	"ethereum-hoodi",
	"ethereum-sepolia",
	"ethereum-mainnet",
	"polygon-mainnet",
	"solana-devnet",
	"solana-mainnet",
	"arbitrum-mainnet",
	"arbitrum-sepolia",
	"bitcoin-mainnet",
	"near-testnet",
	"near-mainnet",
}

func (v *NetworkIdentifier) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkIdentifier(value)
	for _, existing := range AllowedNetworkIdentifierEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkIdentifier", value)
}

// NewNetworkIdentifierFromValue returns a pointer to a valid NetworkIdentifier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkIdentifierFromValue(v string) (*NetworkIdentifier, error) {
	ev := NetworkIdentifier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkIdentifier: valid values are %v", v, AllowedNetworkIdentifierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkIdentifier) IsValid() bool {
	for _, existing := range AllowedNetworkIdentifierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkIdentifier value
func (v NetworkIdentifier) Ptr() *NetworkIdentifier {
	return &v
}

type NullableNetworkIdentifier struct {
	value *NetworkIdentifier
	isSet bool
}

func (v NullableNetworkIdentifier) Get() *NetworkIdentifier {
	return v.value
}

func (v *NullableNetworkIdentifier) Set(val *NetworkIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkIdentifier(val *NetworkIdentifier) *NullableNetworkIdentifier {
	return &NullableNetworkIdentifier{value: val, isSet: true}
}

func (v NullableNetworkIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

