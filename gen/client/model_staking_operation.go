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

// checks if the StakingOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StakingOperation{}

// StakingOperation A list of onchain transactions to help realize a staking action.
type StakingOperation struct {
	// The unique ID of the staking operation.
	Id string `json:"id"`
	// The ID of the wallet that owns the address.
	WalletId *string `json:"wallet_id,omitempty"`
	// The ID of the blockchain network.
	NetworkId string `json:"network_id"`
	// The onchain address orchestrating the staking operation.
	AddressId string `json:"address_id"`
	// The status of the staking operation.
	Status string `json:"status"`
	// The transaction(s) that will execute the staking operation onchain.
	Transactions []Transaction `json:"transactions"`
	Metadata *StakingOperationMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StakingOperation StakingOperation

// NewStakingOperation instantiates a new StakingOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStakingOperation(id string, networkId string, addressId string, status string, transactions []Transaction) *StakingOperation {
	this := StakingOperation{}
	this.Id = id
	this.NetworkId = networkId
	this.AddressId = addressId
	this.Status = status
	this.Transactions = transactions
	return &this
}

// NewStakingOperationWithDefaults instantiates a new StakingOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStakingOperationWithDefaults() *StakingOperation {
	this := StakingOperation{}
	return &this
}

// GetId returns the Id field value
func (o *StakingOperation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StakingOperation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StakingOperation) SetId(v string) {
	o.Id = v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *StakingOperation) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingOperation) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *StakingOperation) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *StakingOperation) SetWalletId(v string) {
	o.WalletId = &v
}

// GetNetworkId returns the NetworkId field value
func (o *StakingOperation) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *StakingOperation) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *StakingOperation) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetAddressId returns the AddressId field value
func (o *StakingOperation) GetAddressId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value
// and a boolean to check if the value has been set.
func (o *StakingOperation) GetAddressIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressId, true
}

// SetAddressId sets field value
func (o *StakingOperation) SetAddressId(v string) {
	o.AddressId = v
}

// GetStatus returns the Status field value
func (o *StakingOperation) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StakingOperation) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StakingOperation) SetStatus(v string) {
	o.Status = v
}

// GetTransactions returns the Transactions field value
func (o *StakingOperation) GetTransactions() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *StakingOperation) GetTransactionsOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *StakingOperation) SetTransactions(v []Transaction) {
	o.Transactions = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *StakingOperation) GetMetadata() StakingOperationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret StakingOperationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingOperation) GetMetadataOk() (*StakingOperationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *StakingOperation) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given StakingOperationMetadata and assigns it to the Metadata field.
func (o *StakingOperation) SetMetadata(v StakingOperationMetadata) {
	o.Metadata = &v
}

func (o StakingOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StakingOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	toSerialize["network_id"] = o.NetworkId
	toSerialize["address_id"] = o.AddressId
	toSerialize["status"] = o.Status
	toSerialize["transactions"] = o.Transactions
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StakingOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"network_id",
		"address_id",
		"status",
		"transactions",
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

	varStakingOperation := _StakingOperation{}

	err = json.Unmarshal(data, &varStakingOperation)

	if err != nil {
		return err
	}

	*o = StakingOperation(varStakingOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "network_id")
		delete(additionalProperties, "address_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "transactions")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStakingOperation struct {
	value *StakingOperation
	isSet bool
}

func (v NullableStakingOperation) Get() *StakingOperation {
	return v.value
}

func (v *NullableStakingOperation) Set(val *StakingOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableStakingOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableStakingOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakingOperation(val *StakingOperation) *NullableStakingOperation {
	return &NullableStakingOperation{value: val, isSet: true}
}

func (v NullableStakingOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakingOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


