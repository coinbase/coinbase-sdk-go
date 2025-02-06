/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the FundQuote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FundQuote{}

// FundQuote A quote for a fund operation
type FundQuote struct {
	// The ID of the fund quote
	FundQuoteId string `json:"fund_quote_id"`
	// The ID of the blockchain network
	NetworkId string `json:"network_id"`
	// The ID of the wallet that will receive the crypto
	WalletId string `json:"wallet_id"`
	// The ID of the address that will receive the crypto
	AddressId string `json:"address_id"`
	CryptoAmount CryptoAmount `json:"crypto_amount"`
	FiatAmount FiatAmount `json:"fiat_amount"`
	// The time at which the quote expires
	ExpiresAt time.Time `json:"expires_at"`
	Fees FundOperationFees `json:"fees"`
	AdditionalProperties map[string]interface{}
}

type _FundQuote FundQuote

// NewFundQuote instantiates a new FundQuote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundQuote(fundQuoteId string, networkId string, walletId string, addressId string, cryptoAmount CryptoAmount, fiatAmount FiatAmount, expiresAt time.Time, fees FundOperationFees) *FundQuote {
	this := FundQuote{}
	this.FundQuoteId = fundQuoteId
	this.NetworkId = networkId
	this.WalletId = walletId
	this.AddressId = addressId
	this.CryptoAmount = cryptoAmount
	this.FiatAmount = fiatAmount
	this.ExpiresAt = expiresAt
	this.Fees = fees
	return &this
}

// NewFundQuoteWithDefaults instantiates a new FundQuote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundQuoteWithDefaults() *FundQuote {
	this := FundQuote{}
	return &this
}

// GetFundQuoteId returns the FundQuoteId field value
func (o *FundQuote) GetFundQuoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundQuoteId
}

// GetFundQuoteIdOk returns a tuple with the FundQuoteId field value
// and a boolean to check if the value has been set.
func (o *FundQuote) GetFundQuoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundQuoteId, true
}

// SetFundQuoteId sets field value
func (o *FundQuote) SetFundQuoteId(v string) {
	o.FundQuoteId = v
}

// GetNetworkId returns the NetworkId field value
func (o *FundQuote) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *FundQuote) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *FundQuote) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetWalletId returns the WalletId field value
func (o *FundQuote) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *FundQuote) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *FundQuote) SetWalletId(v string) {
	o.WalletId = v
}

// GetAddressId returns the AddressId field value
func (o *FundQuote) GetAddressId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value
// and a boolean to check if the value has been set.
func (o *FundQuote) GetAddressIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressId, true
}

// SetAddressId sets field value
func (o *FundQuote) SetAddressId(v string) {
	o.AddressId = v
}

// GetCryptoAmount returns the CryptoAmount field value
func (o *FundQuote) GetCryptoAmount() CryptoAmount {
	if o == nil {
		var ret CryptoAmount
		return ret
	}

	return o.CryptoAmount
}

// GetCryptoAmountOk returns a tuple with the CryptoAmount field value
// and a boolean to check if the value has been set.
func (o *FundQuote) GetCryptoAmountOk() (*CryptoAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoAmount, true
}

// SetCryptoAmount sets field value
func (o *FundQuote) SetCryptoAmount(v CryptoAmount) {
	o.CryptoAmount = v
}

// GetFiatAmount returns the FiatAmount field value
func (o *FundQuote) GetFiatAmount() FiatAmount {
	if o == nil {
		var ret FiatAmount
		return ret
	}

	return o.FiatAmount
}

// GetFiatAmountOk returns a tuple with the FiatAmount field value
// and a boolean to check if the value has been set.
func (o *FundQuote) GetFiatAmountOk() (*FiatAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiatAmount, true
}

// SetFiatAmount sets field value
func (o *FundQuote) SetFiatAmount(v FiatAmount) {
	o.FiatAmount = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *FundQuote) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *FundQuote) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *FundQuote) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetFees returns the Fees field value
func (o *FundQuote) GetFees() FundOperationFees {
	if o == nil {
		var ret FundOperationFees
		return ret
	}

	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value
// and a boolean to check if the value has been set.
func (o *FundQuote) GetFeesOk() (*FundOperationFees, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fees, true
}

// SetFees sets field value
func (o *FundQuote) SetFees(v FundOperationFees) {
	o.Fees = v
}

func (o FundQuote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundQuote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fund_quote_id"] = o.FundQuoteId
	toSerialize["network_id"] = o.NetworkId
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["address_id"] = o.AddressId
	toSerialize["crypto_amount"] = o.CryptoAmount
	toSerialize["fiat_amount"] = o.FiatAmount
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["fees"] = o.Fees

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FundQuote) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fund_quote_id",
		"network_id",
		"wallet_id",
		"address_id",
		"crypto_amount",
		"fiat_amount",
		"expires_at",
		"fees",
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

	varFundQuote := _FundQuote{}

	err = json.Unmarshal(data, &varFundQuote)

	if err != nil {
		return err
	}

	*o = FundQuote(varFundQuote)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fund_quote_id")
		delete(additionalProperties, "network_id")
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "address_id")
		delete(additionalProperties, "crypto_amount")
		delete(additionalProperties, "fiat_amount")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "fees")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFundQuote struct {
	value *FundQuote
	isSet bool
}

func (v NullableFundQuote) Get() *FundQuote {
	return v.value
}

func (v *NullableFundQuote) Set(val *FundQuote) {
	v.value = val
	v.isSet = true
}

func (v NullableFundQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableFundQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundQuote(val *FundQuote) *NullableFundQuote {
	return &NullableFundQuote{value: val, isSet: true}
}

func (v NullableFundQuote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


