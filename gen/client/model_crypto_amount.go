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

// checks if the CryptoAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptoAmount{}

// CryptoAmount An amount in cryptocurrency
type CryptoAmount struct {
	// The amount of the crypto in atomic units
	Amount string `json:"amount"`
	Asset Asset `json:"asset"`
}

type _CryptoAmount CryptoAmount

// NewCryptoAmount instantiates a new CryptoAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoAmount(amount string, asset Asset) *CryptoAmount {
	this := CryptoAmount{}
	this.Amount = amount
	this.Asset = asset
	return &this
}

// NewCryptoAmountWithDefaults instantiates a new CryptoAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoAmountWithDefaults() *CryptoAmount {
	this := CryptoAmount{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CryptoAmount) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CryptoAmount) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CryptoAmount) SetAmount(v string) {
	o.Amount = v
}

// GetAsset returns the Asset field value
func (o *CryptoAmount) GetAsset() Asset {
	if o == nil {
		var ret Asset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *CryptoAmount) GetAssetOk() (*Asset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *CryptoAmount) SetAsset(v Asset) {
	o.Asset = v
}

func (o CryptoAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptoAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *CryptoAmount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
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

	varCryptoAmount := _CryptoAmount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCryptoAmount)

	if err != nil {
		return err
	}

	*o = CryptoAmount(varCryptoAmount)

	return err
}

type NullableCryptoAmount struct {
	value *CryptoAmount
	isSet bool
}

func (v NullableCryptoAmount) Get() *CryptoAmount {
	return v.value
}

func (v *NullableCryptoAmount) Set(val *CryptoAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoAmount(val *CryptoAmount) *NullableCryptoAmount {
	return &NullableCryptoAmount{value: val, isSet: true}
}

func (v NullableCryptoAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

