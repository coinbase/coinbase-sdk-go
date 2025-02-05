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
	// The unique identifier of the smart contract.
	SmartContractId string `json:"smart_contract_id"`
	// The name of the blockchain network
	NetworkId string `json:"network_id"`
	// The ID of the wallet that deployed the smart contract. If this smart contract was deployed externally, this will be omitted.
	WalletId *string `json:"wallet_id,omitempty"`
	// The EVM address of the smart contract
	ContractAddress string `json:"contract_address"`
	// The name of the smart contract
	ContractName string `json:"contract_name"`
	// The EVM address of the account that deployed the smart contract. If this smart contract was deployed externally, this will be omitted.
	DeployerAddress *string `json:"deployer_address,omitempty"`
	Type SmartContractType `json:"type"`
	Options *SmartContractOptions `json:"options,omitempty"`
	// The JSON-encoded ABI of the contract
	Abi string `json:"abi"`
	Transaction *Transaction `json:"transaction,omitempty"`
	// Whether the smart contract was deployed externally. If true, the deployer_address and transaction will be omitted.
	IsExternal bool `json:"is_external"`
	// The ID of the compiled smart contract that was used to deploy this contract
	CompiledSmartContractId *string `json:"compiled_smart_contract_id,omitempty"`
}

type _SmartContract SmartContract

// NewSmartContract instantiates a new SmartContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContract(smartContractId string, networkId string, contractAddress string, contractName string, type_ SmartContractType, abi string, isExternal bool) *SmartContract {
	this := SmartContract{}
	this.SmartContractId = smartContractId
	this.NetworkId = networkId
	this.ContractAddress = contractAddress
	this.ContractName = contractName
	this.Type = type_
	this.Abi = abi
	this.IsExternal = isExternal
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

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *SmartContract) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContract) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *SmartContract) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *SmartContract) SetWalletId(v string) {
	o.WalletId = &v
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

// GetContractName returns the ContractName field value
func (o *SmartContract) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *SmartContract) SetContractName(v string) {
	o.ContractName = v
}

// GetDeployerAddress returns the DeployerAddress field value if set, zero value otherwise.
func (o *SmartContract) GetDeployerAddress() string {
	if o == nil || IsNil(o.DeployerAddress) {
		var ret string
		return ret
	}
	return *o.DeployerAddress
}

// GetDeployerAddressOk returns a tuple with the DeployerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContract) GetDeployerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DeployerAddress) {
		return nil, false
	}
	return o.DeployerAddress, true
}

// HasDeployerAddress returns a boolean if a field has been set.
func (o *SmartContract) HasDeployerAddress() bool {
	if o != nil && !IsNil(o.DeployerAddress) {
		return true
	}

	return false
}

// SetDeployerAddress gets a reference to the given string and assigns it to the DeployerAddress field.
func (o *SmartContract) SetDeployerAddress(v string) {
	o.DeployerAddress = &v
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

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SmartContract) GetOptions() SmartContractOptions {
	if o == nil || IsNil(o.Options) {
		var ret SmartContractOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContract) GetOptionsOk() (*SmartContractOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SmartContract) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SmartContractOptions and assigns it to the Options field.
func (o *SmartContract) SetOptions(v SmartContractOptions) {
	o.Options = &v
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

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *SmartContract) GetTransaction() Transaction {
	if o == nil || IsNil(o.Transaction) {
		var ret Transaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContract) GetTransactionOk() (*Transaction, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *SmartContract) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given Transaction and assigns it to the Transaction field.
func (o *SmartContract) SetTransaction(v Transaction) {
	o.Transaction = &v
}

// GetIsExternal returns the IsExternal field value
func (o *SmartContract) GetIsExternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExternal
}

// GetIsExternalOk returns a tuple with the IsExternal field value
// and a boolean to check if the value has been set.
func (o *SmartContract) GetIsExternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExternal, true
}

// SetIsExternal sets field value
func (o *SmartContract) SetIsExternal(v bool) {
	o.IsExternal = v
}

// GetCompiledSmartContractId returns the CompiledSmartContractId field value if set, zero value otherwise.
func (o *SmartContract) GetCompiledSmartContractId() string {
	if o == nil || IsNil(o.CompiledSmartContractId) {
		var ret string
		return ret
	}
	return *o.CompiledSmartContractId
}

// GetCompiledSmartContractIdOk returns a tuple with the CompiledSmartContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContract) GetCompiledSmartContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompiledSmartContractId) {
		return nil, false
	}
	return o.CompiledSmartContractId, true
}

// HasCompiledSmartContractId returns a boolean if a field has been set.
func (o *SmartContract) HasCompiledSmartContractId() bool {
	if o != nil && !IsNil(o.CompiledSmartContractId) {
		return true
	}

	return false
}

// SetCompiledSmartContractId gets a reference to the given string and assigns it to the CompiledSmartContractId field.
func (o *SmartContract) SetCompiledSmartContractId(v string) {
	o.CompiledSmartContractId = &v
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
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["contract_name"] = o.ContractName
	if !IsNil(o.DeployerAddress) {
		toSerialize["deployer_address"] = o.DeployerAddress
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["abi"] = o.Abi
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	toSerialize["is_external"] = o.IsExternal
	if !IsNil(o.CompiledSmartContractId) {
		toSerialize["compiled_smart_contract_id"] = o.CompiledSmartContractId
	}
	return toSerialize, nil
}

func (o *SmartContract) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"smart_contract_id",
		"network_id",
		"contract_address",
		"contract_name",
		"type",
		"abi",
		"is_external",
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


