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

// checks if the DedicatedEthStakeTransactionDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DedicatedEthStakeTransactionDetail{}

// DedicatedEthStakeTransactionDetail A Dedicated ETH staking transaction containing both tx related details and non-tx related metadata.
type DedicatedEthStakeTransactionDetail struct {
	Transaction Transaction `json:"transaction"`
	EthereumValidatorDepositData EthereumValidatorDepositData `json:"ethereum_validator_deposit_data"`
}

type _DedicatedEthStakeTransactionDetail DedicatedEthStakeTransactionDetail

// NewDedicatedEthStakeTransactionDetail instantiates a new DedicatedEthStakeTransactionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedEthStakeTransactionDetail(transaction Transaction, ethereumValidatorDepositData EthereumValidatorDepositData) *DedicatedEthStakeTransactionDetail {
	this := DedicatedEthStakeTransactionDetail{}
	this.Transaction = transaction
	this.EthereumValidatorDepositData = ethereumValidatorDepositData
	return &this
}

// NewDedicatedEthStakeTransactionDetailWithDefaults instantiates a new DedicatedEthStakeTransactionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDedicatedEthStakeTransactionDetailWithDefaults() *DedicatedEthStakeTransactionDetail {
	this := DedicatedEthStakeTransactionDetail{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *DedicatedEthStakeTransactionDetail) GetTransaction() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *DedicatedEthStakeTransactionDetail) GetTransactionOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *DedicatedEthStakeTransactionDetail) SetTransaction(v Transaction) {
	o.Transaction = v
}

// GetEthereumValidatorDepositData returns the EthereumValidatorDepositData field value
func (o *DedicatedEthStakeTransactionDetail) GetEthereumValidatorDepositData() EthereumValidatorDepositData {
	if o == nil {
		var ret EthereumValidatorDepositData
		return ret
	}

	return o.EthereumValidatorDepositData
}

// GetEthereumValidatorDepositDataOk returns a tuple with the EthereumValidatorDepositData field value
// and a boolean to check if the value has been set.
func (o *DedicatedEthStakeTransactionDetail) GetEthereumValidatorDepositDataOk() (*EthereumValidatorDepositData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EthereumValidatorDepositData, true
}

// SetEthereumValidatorDepositData sets field value
func (o *DedicatedEthStakeTransactionDetail) SetEthereumValidatorDepositData(v EthereumValidatorDepositData) {
	o.EthereumValidatorDepositData = v
}

func (o DedicatedEthStakeTransactionDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DedicatedEthStakeTransactionDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction"] = o.Transaction
	toSerialize["ethereum_validator_deposit_data"] = o.EthereumValidatorDepositData
	return toSerialize, nil
}

func (o *DedicatedEthStakeTransactionDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction",
		"ethereum_validator_deposit_data",
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

	varDedicatedEthStakeTransactionDetail := _DedicatedEthStakeTransactionDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDedicatedEthStakeTransactionDetail)

	if err != nil {
		return err
	}

	*o = DedicatedEthStakeTransactionDetail(varDedicatedEthStakeTransactionDetail)

	return err
}

type NullableDedicatedEthStakeTransactionDetail struct {
	value *DedicatedEthStakeTransactionDetail
	isSet bool
}

func (v NullableDedicatedEthStakeTransactionDetail) Get() *DedicatedEthStakeTransactionDetail {
	return v.value
}

func (v *NullableDedicatedEthStakeTransactionDetail) Set(val *DedicatedEthStakeTransactionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableDedicatedEthStakeTransactionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableDedicatedEthStakeTransactionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDedicatedEthStakeTransactionDetail(val *DedicatedEthStakeTransactionDetail) *NullableDedicatedEthStakeTransactionDetail {
	return &NullableDedicatedEthStakeTransactionDetail{value: val, isSet: true}
}

func (v NullableDedicatedEthStakeTransactionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDedicatedEthStakeTransactionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

