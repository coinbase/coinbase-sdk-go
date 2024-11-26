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

// checks if the EthereumTokenTransfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EthereumTokenTransfer{}

// EthereumTokenTransfer struct for EthereumTokenTransfer
type EthereumTokenTransfer struct {
	ContractAddress string `json:"contract_address"`
	FromAddress string `json:"from_address"`
	ToAddress string `json:"to_address"`
	// The value of the transaction in atomic units of the token being transfer for ERC20 or ERC1155 contracts.
	Value *string `json:"value,omitempty"`
	// The ID of ERC721 or ERC1155 token being transferred.
	TokenId *string `json:"token_id,omitempty"`
	LogIndex int64 `json:"log_index"`
	TokenTransferType TokenTransferType `json:"token_transfer_type"`
}

type _EthereumTokenTransfer EthereumTokenTransfer

// NewEthereumTokenTransfer instantiates a new EthereumTokenTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthereumTokenTransfer(contractAddress string, fromAddress string, toAddress string, logIndex int64, tokenTransferType TokenTransferType) *EthereumTokenTransfer {
	this := EthereumTokenTransfer{}
	this.ContractAddress = contractAddress
	this.FromAddress = fromAddress
	this.ToAddress = toAddress
	this.LogIndex = logIndex
	this.TokenTransferType = tokenTransferType
	return &this
}

// NewEthereumTokenTransferWithDefaults instantiates a new EthereumTokenTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthereumTokenTransferWithDefaults() *EthereumTokenTransfer {
	this := EthereumTokenTransfer{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *EthereumTokenTransfer) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *EthereumTokenTransfer) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *EthereumTokenTransfer) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetFromAddress returns the FromAddress field value
func (o *EthereumTokenTransfer) GetFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *EthereumTokenTransfer) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *EthereumTokenTransfer) SetFromAddress(v string) {
	o.FromAddress = v
}

// GetToAddress returns the ToAddress field value
func (o *EthereumTokenTransfer) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *EthereumTokenTransfer) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *EthereumTokenTransfer) SetToAddress(v string) {
	o.ToAddress = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EthereumTokenTransfer) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTokenTransfer) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EthereumTokenTransfer) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EthereumTokenTransfer) SetValue(v string) {
	o.Value = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *EthereumTokenTransfer) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTokenTransfer) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *EthereumTokenTransfer) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *EthereumTokenTransfer) SetTokenId(v string) {
	o.TokenId = &v
}

// GetLogIndex returns the LogIndex field value
func (o *EthereumTokenTransfer) GetLogIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LogIndex
}

// GetLogIndexOk returns a tuple with the LogIndex field value
// and a boolean to check if the value has been set.
func (o *EthereumTokenTransfer) GetLogIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogIndex, true
}

// SetLogIndex sets field value
func (o *EthereumTokenTransfer) SetLogIndex(v int64) {
	o.LogIndex = v
}

// GetTokenTransferType returns the TokenTransferType field value
func (o *EthereumTokenTransfer) GetTokenTransferType() TokenTransferType {
	if o == nil {
		var ret TokenTransferType
		return ret
	}

	return o.TokenTransferType
}

// GetTokenTransferTypeOk returns a tuple with the TokenTransferType field value
// and a boolean to check if the value has been set.
func (o *EthereumTokenTransfer) GetTokenTransferTypeOk() (*TokenTransferType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenTransferType, true
}

// SetTokenTransferType sets field value
func (o *EthereumTokenTransfer) SetTokenTransferType(v TokenTransferType) {
	o.TokenTransferType = v
}

func (o EthereumTokenTransfer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthereumTokenTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["from_address"] = o.FromAddress
	toSerialize["to_address"] = o.ToAddress
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	toSerialize["log_index"] = o.LogIndex
	toSerialize["token_transfer_type"] = o.TokenTransferType
	return toSerialize, nil
}

func (o *EthereumTokenTransfer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contract_address",
		"from_address",
		"to_address",
		"log_index",
		"token_transfer_type",
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

	varEthereumTokenTransfer := _EthereumTokenTransfer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEthereumTokenTransfer)

	if err != nil {
		return err
	}

	*o = EthereumTokenTransfer(varEthereumTokenTransfer)

	return err
}

type NullableEthereumTokenTransfer struct {
	value *EthereumTokenTransfer
	isSet bool
}

func (v NullableEthereumTokenTransfer) Get() *EthereumTokenTransfer {
	return v.value
}

func (v *NullableEthereumTokenTransfer) Set(val *EthereumTokenTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableEthereumTokenTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableEthereumTokenTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthereumTokenTransfer(val *EthereumTokenTransfer) *NullableEthereumTokenTransfer {
	return &NullableEthereumTokenTransfer{value: val, isSet: true}
}

func (v NullableEthereumTokenTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthereumTokenTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


