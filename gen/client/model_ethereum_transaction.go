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
	"bytes"
	"fmt"
)

// checks if the EthereumTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EthereumTransaction{}

// EthereumTransaction struct for EthereumTransaction
type EthereumTransaction struct {
	// The onchain address of the sender.
	From string `json:"from"`
	// The amount of gas spent in the transaction.
	Gas *int32 `json:"gas,omitempty"`
	// The price per gas spent in the transaction in atomic units of the native asset.
	GasPrice *int32 `json:"gas_price,omitempty"`
	// The hash of the transaction as a hexadecimal string, prefixed with 0x.
	Hash *string `json:"hash,omitempty"`
	// The input data of the transaction.
	Input *string `json:"input,omitempty"`
	// The nonce of the transaction in the source address.
	Nonce *int32 `json:"nonce,omitempty"`
	// The onchain address of the receiver.
	To string `json:"to"`
	// The index of the transaction in the block.
	Index *int32 `json:"index,omitempty"`
	// The value of the transaction in atomic units of the native asset.
	Value *string `json:"value,omitempty"`
	// The EIP-2718 transaction type. See https://eips.ethereum.org/EIPS/eip-2718 for more details.
	Type *int32 `json:"type,omitempty"`
	// The max fee per gas as defined in EIP-1559. https://eips.ethereum.org/EIPS/eip-1559 for more details.
	MaxFeePerGas *int32 `json:"max_fee_per_gas,omitempty"`
	// The max priority fee per gas as defined in EIP-1559. https://eips.ethereum.org/EIPS/eip-1559 for more details.
	MaxPriorityFeePerGas *int32 `json:"max_priority_fee_per_gas,omitempty"`
	// The confirmed priority fee per gas as defined in EIP-1559. https://eips.ethereum.org/EIPS/eip-1559 for more details.
	PriorityFeePerGas *int32 `json:"priority_fee_per_gas,omitempty"`
	TransactionAccessList *EthereumTransactionAccessList `json:"transaction_access_list,omitempty"`
	FlattenedTraces []EthereumTransactionFlattenedTrace `json:"flattened_traces,omitempty"`
	// The timestamp of the block in which the event was emitted
	BlockTimestamp *time.Time `json:"block_timestamp,omitempty"`
	// This is for handling optimism rollup specific EIP-2718 transaction type field.
	Mint *string `json:"mint,omitempty"`
}

type _EthereumTransaction EthereumTransaction

// NewEthereumTransaction instantiates a new EthereumTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthereumTransaction(from string, to string) *EthereumTransaction {
	this := EthereumTransaction{}
	this.From = from
	this.To = to
	return &this
}

// NewEthereumTransactionWithDefaults instantiates a new EthereumTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthereumTransactionWithDefaults() *EthereumTransaction {
	this := EthereumTransaction{}
	return &this
}

// GetFrom returns the From field value
func (o *EthereumTransaction) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *EthereumTransaction) SetFrom(v string) {
	o.From = v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *EthereumTransaction) GetGas() int32 {
	if o == nil || IsNil(o.Gas) {
		var ret int32
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetGasOk() (*int32, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *EthereumTransaction) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given int32 and assigns it to the Gas field.
func (o *EthereumTransaction) SetGas(v int32) {
	o.Gas = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *EthereumTransaction) GetGasPrice() int32 {
	if o == nil || IsNil(o.GasPrice) {
		var ret int32
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetGasPriceOk() (*int32, bool) {
	if o == nil || IsNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *EthereumTransaction) HasGasPrice() bool {
	if o != nil && !IsNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given int32 and assigns it to the GasPrice field.
func (o *EthereumTransaction) SetGasPrice(v int32) {
	o.GasPrice = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *EthereumTransaction) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *EthereumTransaction) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *EthereumTransaction) SetHash(v string) {
	o.Hash = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *EthereumTransaction) GetInput() string {
	if o == nil || IsNil(o.Input) {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetInputOk() (*string, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *EthereumTransaction) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *EthereumTransaction) SetInput(v string) {
	o.Input = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *EthereumTransaction) GetNonce() int32 {
	if o == nil || IsNil(o.Nonce) {
		var ret int32
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetNonceOk() (*int32, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *EthereumTransaction) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given int32 and assigns it to the Nonce field.
func (o *EthereumTransaction) SetNonce(v int32) {
	o.Nonce = &v
}

// GetTo returns the To field value
func (o *EthereumTransaction) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *EthereumTransaction) SetTo(v string) {
	o.To = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *EthereumTransaction) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *EthereumTransaction) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *EthereumTransaction) SetIndex(v int32) {
	o.Index = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EthereumTransaction) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EthereumTransaction) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EthereumTransaction) SetValue(v string) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EthereumTransaction) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EthereumTransaction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *EthereumTransaction) SetType(v int32) {
	o.Type = &v
}

// GetMaxFeePerGas returns the MaxFeePerGas field value if set, zero value otherwise.
func (o *EthereumTransaction) GetMaxFeePerGas() int32 {
	if o == nil || IsNil(o.MaxFeePerGas) {
		var ret int32
		return ret
	}
	return *o.MaxFeePerGas
}

// GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetMaxFeePerGasOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxFeePerGas) {
		return nil, false
	}
	return o.MaxFeePerGas, true
}

// HasMaxFeePerGas returns a boolean if a field has been set.
func (o *EthereumTransaction) HasMaxFeePerGas() bool {
	if o != nil && !IsNil(o.MaxFeePerGas) {
		return true
	}

	return false
}

// SetMaxFeePerGas gets a reference to the given int32 and assigns it to the MaxFeePerGas field.
func (o *EthereumTransaction) SetMaxFeePerGas(v int32) {
	o.MaxFeePerGas = &v
}

// GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field value if set, zero value otherwise.
func (o *EthereumTransaction) GetMaxPriorityFeePerGas() int32 {
	if o == nil || IsNil(o.MaxPriorityFeePerGas) {
		var ret int32
		return ret
	}
	return *o.MaxPriorityFeePerGas
}

// GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetMaxPriorityFeePerGasOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPriorityFeePerGas) {
		return nil, false
	}
	return o.MaxPriorityFeePerGas, true
}

// HasMaxPriorityFeePerGas returns a boolean if a field has been set.
func (o *EthereumTransaction) HasMaxPriorityFeePerGas() bool {
	if o != nil && !IsNil(o.MaxPriorityFeePerGas) {
		return true
	}

	return false
}

// SetMaxPriorityFeePerGas gets a reference to the given int32 and assigns it to the MaxPriorityFeePerGas field.
func (o *EthereumTransaction) SetMaxPriorityFeePerGas(v int32) {
	o.MaxPriorityFeePerGas = &v
}

// GetPriorityFeePerGas returns the PriorityFeePerGas field value if set, zero value otherwise.
func (o *EthereumTransaction) GetPriorityFeePerGas() int32 {
	if o == nil || IsNil(o.PriorityFeePerGas) {
		var ret int32
		return ret
	}
	return *o.PriorityFeePerGas
}

// GetPriorityFeePerGasOk returns a tuple with the PriorityFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetPriorityFeePerGasOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityFeePerGas) {
		return nil, false
	}
	return o.PriorityFeePerGas, true
}

// HasPriorityFeePerGas returns a boolean if a field has been set.
func (o *EthereumTransaction) HasPriorityFeePerGas() bool {
	if o != nil && !IsNil(o.PriorityFeePerGas) {
		return true
	}

	return false
}

// SetPriorityFeePerGas gets a reference to the given int32 and assigns it to the PriorityFeePerGas field.
func (o *EthereumTransaction) SetPriorityFeePerGas(v int32) {
	o.PriorityFeePerGas = &v
}

// GetTransactionAccessList returns the TransactionAccessList field value if set, zero value otherwise.
func (o *EthereumTransaction) GetTransactionAccessList() EthereumTransactionAccessList {
	if o == nil || IsNil(o.TransactionAccessList) {
		var ret EthereumTransactionAccessList
		return ret
	}
	return *o.TransactionAccessList
}

// GetTransactionAccessListOk returns a tuple with the TransactionAccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetTransactionAccessListOk() (*EthereumTransactionAccessList, bool) {
	if o == nil || IsNil(o.TransactionAccessList) {
		return nil, false
	}
	return o.TransactionAccessList, true
}

// HasTransactionAccessList returns a boolean if a field has been set.
func (o *EthereumTransaction) HasTransactionAccessList() bool {
	if o != nil && !IsNil(o.TransactionAccessList) {
		return true
	}

	return false
}

// SetTransactionAccessList gets a reference to the given EthereumTransactionAccessList and assigns it to the TransactionAccessList field.
func (o *EthereumTransaction) SetTransactionAccessList(v EthereumTransactionAccessList) {
	o.TransactionAccessList = &v
}

// GetFlattenedTraces returns the FlattenedTraces field value if set, zero value otherwise.
func (o *EthereumTransaction) GetFlattenedTraces() []EthereumTransactionFlattenedTrace {
	if o == nil || IsNil(o.FlattenedTraces) {
		var ret []EthereumTransactionFlattenedTrace
		return ret
	}
	return o.FlattenedTraces
}

// GetFlattenedTracesOk returns a tuple with the FlattenedTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetFlattenedTracesOk() ([]EthereumTransactionFlattenedTrace, bool) {
	if o == nil || IsNil(o.FlattenedTraces) {
		return nil, false
	}
	return o.FlattenedTraces, true
}

// HasFlattenedTraces returns a boolean if a field has been set.
func (o *EthereumTransaction) HasFlattenedTraces() bool {
	if o != nil && !IsNil(o.FlattenedTraces) {
		return true
	}

	return false
}

// SetFlattenedTraces gets a reference to the given []EthereumTransactionFlattenedTrace and assigns it to the FlattenedTraces field.
func (o *EthereumTransaction) SetFlattenedTraces(v []EthereumTransactionFlattenedTrace) {
	o.FlattenedTraces = v
}

// GetBlockTimestamp returns the BlockTimestamp field value if set, zero value otherwise.
func (o *EthereumTransaction) GetBlockTimestamp() time.Time {
	if o == nil || IsNil(o.BlockTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.BlockTimestamp
}

// GetBlockTimestampOk returns a tuple with the BlockTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetBlockTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BlockTimestamp) {
		return nil, false
	}
	return o.BlockTimestamp, true
}

// HasBlockTimestamp returns a boolean if a field has been set.
func (o *EthereumTransaction) HasBlockTimestamp() bool {
	if o != nil && !IsNil(o.BlockTimestamp) {
		return true
	}

	return false
}

// SetBlockTimestamp gets a reference to the given time.Time and assigns it to the BlockTimestamp field.
func (o *EthereumTransaction) SetBlockTimestamp(v time.Time) {
	o.BlockTimestamp = &v
}

// GetMint returns the Mint field value if set, zero value otherwise.
func (o *EthereumTransaction) GetMint() string {
	if o == nil || IsNil(o.Mint) {
		var ret string
		return ret
	}
	return *o.Mint
}

// GetMintOk returns a tuple with the Mint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetMintOk() (*string, bool) {
	if o == nil || IsNil(o.Mint) {
		return nil, false
	}
	return o.Mint, true
}

// HasMint returns a boolean if a field has been set.
func (o *EthereumTransaction) HasMint() bool {
	if o != nil && !IsNil(o.Mint) {
		return true
	}

	return false
}

// SetMint gets a reference to the given string and assigns it to the Mint field.
func (o *EthereumTransaction) SetMint(v string) {
	o.Mint = &v
}

func (o EthereumTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthereumTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !IsNil(o.GasPrice) {
		toSerialize["gas_price"] = o.GasPrice
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	toSerialize["to"] = o.To
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.MaxFeePerGas) {
		toSerialize["max_fee_per_gas"] = o.MaxFeePerGas
	}
	if !IsNil(o.MaxPriorityFeePerGas) {
		toSerialize["max_priority_fee_per_gas"] = o.MaxPriorityFeePerGas
	}
	if !IsNil(o.PriorityFeePerGas) {
		toSerialize["priority_fee_per_gas"] = o.PriorityFeePerGas
	}
	if !IsNil(o.TransactionAccessList) {
		toSerialize["transaction_access_list"] = o.TransactionAccessList
	}
	if !IsNil(o.FlattenedTraces) {
		toSerialize["flattened_traces"] = o.FlattenedTraces
	}
	if !IsNil(o.BlockTimestamp) {
		toSerialize["block_timestamp"] = o.BlockTimestamp
	}
	if !IsNil(o.Mint) {
		toSerialize["mint"] = o.Mint
	}
	return toSerialize, nil
}

func (o *EthereumTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from",
		"to",
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

	varEthereumTransaction := _EthereumTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEthereumTransaction)

	if err != nil {
		return err
	}

	*o = EthereumTransaction(varEthereumTransaction)

	return err
}

type NullableEthereumTransaction struct {
	value *EthereumTransaction
	isSet bool
}

func (v NullableEthereumTransaction) Get() *EthereumTransaction {
	return v.value
}

func (v *NullableEthereumTransaction) Set(val *EthereumTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableEthereumTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableEthereumTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthereumTransaction(val *EthereumTransaction) *NullableEthereumTransaction {
	return &NullableEthereumTransaction{value: val, isSet: true}
}

func (v NullableEthereumTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthereumTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


