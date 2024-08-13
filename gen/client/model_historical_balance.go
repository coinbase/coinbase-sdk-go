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
	"bytes"
	"fmt"
)

// checks if the HistoricalBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricalBalance{}

// HistoricalBalance The balance of an asset onchain at a particular block
type HistoricalBalance struct {
	// The amount in the atomic units of the asset
	Amount string `json:"amount"`
	// The hash of the block at which the balance was recorded
	BlockHash string `json:"block_hash"`
	// The block height at which the balance was recorded
	BlockHeight string `json:"block_height"`
	Asset Asset `json:"asset"`
}

type _HistoricalBalance HistoricalBalance

// NewHistoricalBalance instantiates a new HistoricalBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalBalance(amount string, blockHash string, blockHeight string, asset Asset) *HistoricalBalance {
	this := HistoricalBalance{}
	this.Amount = amount
	this.BlockHash = blockHash
	this.BlockHeight = blockHeight
	this.Asset = asset
	return &this
}

// NewHistoricalBalanceWithDefaults instantiates a new HistoricalBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalBalanceWithDefaults() *HistoricalBalance {
	this := HistoricalBalance{}
	return &this
}

// GetAmount returns the Amount field value
func (o *HistoricalBalance) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *HistoricalBalance) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *HistoricalBalance) SetAmount(v string) {
	o.Amount = v
}

// GetBlockHash returns the BlockHash field value
func (o *HistoricalBalance) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *HistoricalBalance) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *HistoricalBalance) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *HistoricalBalance) GetBlockHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *HistoricalBalance) GetBlockHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *HistoricalBalance) SetBlockHeight(v string) {
	o.BlockHeight = v
}

// GetAsset returns the Asset field value
func (o *HistoricalBalance) GetAsset() Asset {
	if o == nil {
		var ret Asset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *HistoricalBalance) GetAssetOk() (*Asset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *HistoricalBalance) SetAsset(v Asset) {
	o.Asset = v
}

func (o HistoricalBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["block_hash"] = o.BlockHash
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *HistoricalBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"block_hash",
		"block_height",
		"asset",
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

	varHistoricalBalance := _HistoricalBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHistoricalBalance)

	if err != nil {
		return err
	}

	*o = HistoricalBalance(varHistoricalBalance)

	return err
}

type NullableHistoricalBalance struct {
	value *HistoricalBalance
	isSet bool
}

func (v NullableHistoricalBalance) Get() *HistoricalBalance {
	return v.value
}

func (v *NullableHistoricalBalance) Set(val *HistoricalBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalBalance(val *HistoricalBalance) *NullableHistoricalBalance {
	return &NullableHistoricalBalance{value: val, isSet: true}
}

func (v NullableHistoricalBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


