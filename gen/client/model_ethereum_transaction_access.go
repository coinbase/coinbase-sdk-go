/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the EthereumTransactionAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EthereumTransactionAccess{}

// EthereumTransactionAccess struct for EthereumTransactionAccess
type EthereumTransactionAccess struct {
	Address *string `json:"address,omitempty"`
	StorageKeys []string `json:"storage_keys,omitempty"`
}

// NewEthereumTransactionAccess instantiates a new EthereumTransactionAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthereumTransactionAccess() *EthereumTransactionAccess {
	this := EthereumTransactionAccess{}
	return &this
}

// NewEthereumTransactionAccessWithDefaults instantiates a new EthereumTransactionAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthereumTransactionAccessWithDefaults() *EthereumTransactionAccess {
	this := EthereumTransactionAccess{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *EthereumTransactionAccess) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionAccess) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *EthereumTransactionAccess) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *EthereumTransactionAccess) SetAddress(v string) {
	o.Address = &v
}

// GetStorageKeys returns the StorageKeys field value if set, zero value otherwise.
func (o *EthereumTransactionAccess) GetStorageKeys() []string {
	if o == nil || IsNil(o.StorageKeys) {
		var ret []string
		return ret
	}
	return o.StorageKeys
}

// GetStorageKeysOk returns a tuple with the StorageKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionAccess) GetStorageKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.StorageKeys) {
		return nil, false
	}
	return o.StorageKeys, true
}

// HasStorageKeys returns a boolean if a field has been set.
func (o *EthereumTransactionAccess) HasStorageKeys() bool {
	if o != nil && !IsNil(o.StorageKeys) {
		return true
	}

	return false
}

// SetStorageKeys gets a reference to the given []string and assigns it to the StorageKeys field.
func (o *EthereumTransactionAccess) SetStorageKeys(v []string) {
	o.StorageKeys = v
}

func (o EthereumTransactionAccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthereumTransactionAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.StorageKeys) {
		toSerialize["storage_keys"] = o.StorageKeys
	}
	return toSerialize, nil
}

type NullableEthereumTransactionAccess struct {
	value *EthereumTransactionAccess
	isSet bool
}

func (v NullableEthereumTransactionAccess) Get() *EthereumTransactionAccess {
	return v.value
}

func (v *NullableEthereumTransactionAccess) Set(val *EthereumTransactionAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableEthereumTransactionAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableEthereumTransactionAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthereumTransactionAccess(val *EthereumTransactionAccess) *NullableEthereumTransactionAccess {
	return &NullableEthereumTransactionAccess{value: val, isSet: true}
}

func (v NullableEthereumTransactionAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthereumTransactionAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

