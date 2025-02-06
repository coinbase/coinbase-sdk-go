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

// checks if the FeatureSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureSet{}

// FeatureSet struct for FeatureSet
type FeatureSet struct {
	// Whether the network supports a faucet
	Faucet bool `json:"faucet"`
	// Whether the network supports Server-Signers
	ServerSigner bool `json:"server_signer"`
	// Whether the network supports transfers
	Transfer bool `json:"transfer"`
	// Whether the network supports trading
	Trade bool `json:"trade"`
	// Whether the network supports staking
	Stake bool `json:"stake"`
	// Whether the network supports gasless sends
	GaslessSend bool `json:"gasless_send"`
	AdditionalProperties map[string]interface{}
}

type _FeatureSet FeatureSet

// NewFeatureSet instantiates a new FeatureSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureSet(faucet bool, serverSigner bool, transfer bool, trade bool, stake bool, gaslessSend bool) *FeatureSet {
	this := FeatureSet{}
	this.Faucet = faucet
	this.ServerSigner = serverSigner
	this.Transfer = transfer
	this.Trade = trade
	this.Stake = stake
	this.GaslessSend = gaslessSend
	return &this
}

// NewFeatureSetWithDefaults instantiates a new FeatureSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureSetWithDefaults() *FeatureSet {
	this := FeatureSet{}
	return &this
}

// GetFaucet returns the Faucet field value
func (o *FeatureSet) GetFaucet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Faucet
}

// GetFaucetOk returns a tuple with the Faucet field value
// and a boolean to check if the value has been set.
func (o *FeatureSet) GetFaucetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Faucet, true
}

// SetFaucet sets field value
func (o *FeatureSet) SetFaucet(v bool) {
	o.Faucet = v
}

// GetServerSigner returns the ServerSigner field value
func (o *FeatureSet) GetServerSigner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ServerSigner
}

// GetServerSignerOk returns a tuple with the ServerSigner field value
// and a boolean to check if the value has been set.
func (o *FeatureSet) GetServerSignerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerSigner, true
}

// SetServerSigner sets field value
func (o *FeatureSet) SetServerSigner(v bool) {
	o.ServerSigner = v
}

// GetTransfer returns the Transfer field value
func (o *FeatureSet) GetTransfer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value
// and a boolean to check if the value has been set.
func (o *FeatureSet) GetTransferOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transfer, true
}

// SetTransfer sets field value
func (o *FeatureSet) SetTransfer(v bool) {
	o.Transfer = v
}

// GetTrade returns the Trade field value
func (o *FeatureSet) GetTrade() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trade
}

// GetTradeOk returns a tuple with the Trade field value
// and a boolean to check if the value has been set.
func (o *FeatureSet) GetTradeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trade, true
}

// SetTrade sets field value
func (o *FeatureSet) SetTrade(v bool) {
	o.Trade = v
}

// GetStake returns the Stake field value
func (o *FeatureSet) GetStake() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Stake
}

// GetStakeOk returns a tuple with the Stake field value
// and a boolean to check if the value has been set.
func (o *FeatureSet) GetStakeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stake, true
}

// SetStake sets field value
func (o *FeatureSet) SetStake(v bool) {
	o.Stake = v
}

// GetGaslessSend returns the GaslessSend field value
func (o *FeatureSet) GetGaslessSend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GaslessSend
}

// GetGaslessSendOk returns a tuple with the GaslessSend field value
// and a boolean to check if the value has been set.
func (o *FeatureSet) GetGaslessSendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GaslessSend, true
}

// SetGaslessSend sets field value
func (o *FeatureSet) SetGaslessSend(v bool) {
	o.GaslessSend = v
}

func (o FeatureSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["faucet"] = o.Faucet
	toSerialize["server_signer"] = o.ServerSigner
	toSerialize["transfer"] = o.Transfer
	toSerialize["trade"] = o.Trade
	toSerialize["stake"] = o.Stake
	toSerialize["gasless_send"] = o.GaslessSend

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FeatureSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"faucet",
		"server_signer",
		"transfer",
		"trade",
		"stake",
		"gasless_send",
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

	varFeatureSet := _FeatureSet{}

	err = json.Unmarshal(data, &varFeatureSet)

	if err != nil {
		return err
	}

	*o = FeatureSet(varFeatureSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "faucet")
		delete(additionalProperties, "server_signer")
		delete(additionalProperties, "transfer")
		delete(additionalProperties, "trade")
		delete(additionalProperties, "stake")
		delete(additionalProperties, "gasless_send")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeatureSet struct {
	value *FeatureSet
	isSet bool
}

func (v NullableFeatureSet) Get() *FeatureSet {
	return v.value
}

func (v *NullableFeatureSet) Set(val *FeatureSet) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureSet) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureSet(val *FeatureSet) *NullableFeatureSet {
	return &NullableFeatureSet{value: val, isSet: true}
}

func (v NullableFeatureSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


