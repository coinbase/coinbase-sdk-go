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

// checks if the SmartContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartContract{}

// SmartContract Represents a smart contract on the blockchain
type SmartContract struct {
	// The unique identifier of the smart contract
	SmartContractId string `json:"smart_contract_id"`
	// The name of the blockchain network
	NetworkId string `json:"network_id"`
	// The ID of the wallet that deployed the smart contract
	WalletId string `json:"wallet_id"`
	// The EVM address of the smart contract
	ContractAddress string `json:"contract_address"`
	// The EVM address of the account that deployed the smart contract
	DeployerAddress string `json:"deployer_address"`
	Type SmartContractType `json:"type"`
	Options SmartContractOptions `json:"options"`
	// The JSON-encoded ABI of the contract
	Abi string `json:"abi"`
	Transaction Transaction `json:"transaction"`
}

type _SmartContract SmartContract

// NewSmartContract instantiates a new SmartContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContract(smartContractId string, networkId string, walletId string, contractAddress string, deployerAddress string, type_ SmartContractType, options SmartContractOptions, abi string, transaction Transaction) *SmartContract {
	this := SmartContract{}
	this.SmartContractId = smartContractId
	this.NetworkId = networkId
	this.WalletId = walletId
	this.ContractAddress = contractAddress
	this.DeployerAddress = deployerAddress
	this.Type = type_
	this.Options = options
	this.Abi = abi
	this.Transaction = transaction
	return &this
}

// NewSmartContractWithDefaults instantiates a new SmartContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractWithDefaults() *SmartContract {
	this := SmartContract{}
	return &this
}

// GetSmartContractId returns the SmartContractId field value
func (o *SmartContract) GetSmartContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmartContractId
}

// GetSmartContractIdOk returns a tuple with the SmartContractId field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetSmartContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmartContractId, true
}

// SetSmartContractId sets field value
func (o *SmartContract) SetSmartContractId(v string) {
	o.SmartContractId = v
}

// GetNetworkId returns the NetworkId field value
func (o *SmartContract) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *SmartContract) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetWalletId returns the WalletId field value
func (o *SmartContract) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *SmartContract) SetWalletId(v string) {
	o.WalletId = v
}

// GetContractAddress returns the ContractAddress field value
func (o *SmartContract) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *SmartContract) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetDeployerAddress returns the DeployerAddress field value
func (o *SmartContract) GetDeployerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployerAddress
}

// GetDeployerAddressOk returns a tuple with the DeployerAddress field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetDeployerAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployerAddress, true
}

// SetDeployerAddress sets field value
func (o *SmartContract) SetDeployerAddress(v string) {
	o.DeployerAddress = v
}

// GetType returns the Type field value
func (o *SmartContract) GetType() SmartContractType {
	if o == nil {
		var ret SmartContractType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetTypeOk() (*SmartContractType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SmartContract) SetType(v SmartContractType) {
	o.Type = v
}

// GetOptions returns the Options field value
func (o *SmartContract) GetOptions() SmartContractOptions {
	if o == nil {
		var ret SmartContractOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetOptionsOk() (*SmartContractOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *SmartContract) SetOptions(v SmartContractOptions) {
	o.Options = v
}

// GetAbi returns the Abi field value
func (o *SmartContract) GetAbi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Abi
}

// GetAbiOk returns a tuple with the Abi field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetAbiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Abi, true
}

// SetAbi sets field value
func (o *SmartContract) SetAbi(v string) {
	o.Abi = v
}

// GetTransaction returns the Transaction field value
func (o *SmartContract) GetTransaction() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetTransactionOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *SmartContract) SetTransaction(v Transaction) {
	o.Transaction = v
}

func (o SmartContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["smart_contract_id"] = o.SmartContractId
	toSerialize["network_id"] = o.NetworkId
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["deployer_address"] = o.DeployerAddress
	toSerialize["type"] = o.Type
	toSerialize["options"] = o.Options
	toSerialize["abi"] = o.Abi
	toSerialize["transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *SmartContract) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"smart_contract_id",
		"network_id",
		"wallet_id",
		"contract_address",
		"deployer_address",
		"type",
		"options",
		"abi",
		"transaction",
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

	varSmartContract := _SmartContract{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContract)

	if err != nil {
		return err
	}

	*o = SmartContract(varSmartContract)

	return err
}

type NullableSmartContract struct {
	value *SmartContract
	isSet bool
}

func (v NullableSmartContract) Get() *SmartContract {
	return v.value
}

func (v *NullableSmartContract) Set(val *SmartContract) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContract) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContract(val *SmartContract) *NullableSmartContract {
	return &NullableSmartContract{value: val, isSet: true}
}

func (v NullableSmartContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

