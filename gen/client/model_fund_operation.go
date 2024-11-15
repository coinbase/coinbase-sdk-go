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

// checks if the FundOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FundOperation{}

// FundOperation An operation to fund a wallet with crypto
type FundOperation struct {
	// The ID of the fund operation
	FundOperationId string `json:"fund_operation_id"`
	// The ID of the blockchain network
	NetworkId string `json:"network_id"`
	// The ID of the wallet that will receive the crypto
	WalletId string `json:"wallet_id"`
	// The ID of the address that will receive the crypto
	AddressId string `json:"address_id"`
	CryptoAmount CryptoAmount `json:"crypto_amount"`
	FiatAmount FiatAmount `json:"fiat_amount"`
	Fees FundOperationFees `json:"fees"`
	// The status of the fund operation
	Status string `json:"status"`
}

type _FundOperation FundOperation

// NewFundOperation instantiates a new FundOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundOperation(fundOperationId string, networkId string, walletId string, addressId string, cryptoAmount CryptoAmount, fiatAmount FiatAmount, fees FundOperationFees, status string) *FundOperation {
	this := FundOperation{}
	this.FundOperationId = fundOperationId
	this.NetworkId = networkId
	this.WalletId = walletId
	this.AddressId = addressId
	this.CryptoAmount = cryptoAmount
	this.FiatAmount = fiatAmount
	this.Fees = fees
	this.Status = status
	return &this
}

// NewFundOperationWithDefaults instantiates a new FundOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundOperationWithDefaults() *FundOperation {
	this := FundOperation{}
	return &this
}

// GetFundOperationId returns the FundOperationId field value
func (o *FundOperation) GetFundOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundOperationId
}

// GetFundOperationIdOk returns a tuple with the FundOperationId field value
// and a boolean to check if the value has been set.
func (o *FundOperation) GetFundOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundOperationId, true
}

// SetFundOperationId sets field value
func (o *FundOperation) SetFundOperationId(v string) {
	o.FundOperationId = v
}

// GetNetworkId returns the NetworkId field value
func (o *FundOperation) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *FundOperation) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *FundOperation) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetWalletId returns the WalletId field value
func (o *FundOperation) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *FundOperation) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *FundOperation) SetWalletId(v string) {
	o.WalletId = v
}

// GetAddressId returns the AddressId field value
func (o *FundOperation) GetAddressId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value
// and a boolean to check if the value has been set.
func (o *FundOperation) GetAddressIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressId, true
}

// SetAddressId sets field value
func (o *FundOperation) SetAddressId(v string) {
	o.AddressId = v
}

// GetCryptoAmount returns the CryptoAmount field value
func (o *FundOperation) GetCryptoAmount() CryptoAmount {
	if o == nil {
		var ret CryptoAmount
		return ret
	}

	return o.CryptoAmount
}

// GetCryptoAmountOk returns a tuple with the CryptoAmount field value
// and a boolean to check if the value has been set.
func (o *FundOperation) GetCryptoAmountOk() (*CryptoAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoAmount, true
}

// SetCryptoAmount sets field value
func (o *FundOperation) SetCryptoAmount(v CryptoAmount) {
	o.CryptoAmount = v
}

// GetFiatAmount returns the FiatAmount field value
func (o *FundOperation) GetFiatAmount() FiatAmount {
	if o == nil {
		var ret FiatAmount
		return ret
	}

	return o.FiatAmount
}

// GetFiatAmountOk returns a tuple with the FiatAmount field value
// and a boolean to check if the value has been set.
func (o *FundOperation) GetFiatAmountOk() (*FiatAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiatAmount, true
}

// SetFiatAmount sets field value
func (o *FundOperation) SetFiatAmount(v FiatAmount) {
	o.FiatAmount = v
}

// GetFees returns the Fees field value
func (o *FundOperation) GetFees() FundOperationFees {
	if o == nil {
		var ret FundOperationFees
		return ret
	}

	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value
// and a boolean to check if the value has been set.
func (o *FundOperation) GetFeesOk() (*FundOperationFees, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fees, true
}

// SetFees sets field value
func (o *FundOperation) SetFees(v FundOperationFees) {
	o.Fees = v
}

// GetStatus returns the Status field value
func (o *FundOperation) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FundOperation) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FundOperation) SetStatus(v string) {
	o.Status = v
}

func (o FundOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fund_operation_id"] = o.FundOperationId
	toSerialize["network_id"] = o.NetworkId
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["address_id"] = o.AddressId
	toSerialize["crypto_amount"] = o.CryptoAmount
	toSerialize["fiat_amount"] = o.FiatAmount
	toSerialize["fees"] = o.Fees
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *FundOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fund_operation_id",
		"network_id",
		"wallet_id",
		"address_id",
		"crypto_amount",
		"fiat_amount",
		"fees",
		"status",
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

	varFundOperation := _FundOperation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFundOperation)

	if err != nil {
		return err
	}

	*o = FundOperation(varFundOperation)

	return err
}

type NullableFundOperation struct {
	value *FundOperation
	isSet bool
}

func (v NullableFundOperation) Get() *FundOperation {
	return v.value
}

func (v *NullableFundOperation) Set(val *FundOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableFundOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableFundOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundOperation(val *FundOperation) *NullableFundOperation {
	return &NullableFundOperation{value: val, isSet: true}
}

func (v NullableFundOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


