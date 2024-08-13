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
	"time"
)

// checks if the ContractEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractEvent{}

// ContractEvent Represents a single decoded event emitted by a smart contract
type ContractEvent struct {
	// The name of the blockchain network
	NetworkName *string `json:"network_name,omitempty"`
	// The name of the blockchain project or protocol
	ProtocolName *string `json:"protocol_name,omitempty"`
	// The name of the specific contract within the project
	ContractName *string `json:"contract_name,omitempty"`
	// The name of the event emitted by the contract
	EventName *string `json:"event_name,omitempty"`
	// The signature of the event, including parameter types
	Sig *string `json:"sig,omitempty"`
	// The first four bytes of the Keccak hash of the event signature
	FourBytes *string `json:"fourBytes,omitempty"`
	// The EVM address of the smart contract
	ContractAddress *string `json:"contract_address,omitempty"`
	// The timestamp of the block in which the event was emitted
	BlockTime *time.Time `json:"block_time,omitempty"`
	// The block number in which the event was emitted
	BlockHeight *int32 `json:"block_height,omitempty"`
	// The transaction hash in which the event was emitted
	TxHash *string `json:"tx_hash,omitempty"`
	// The index of the transaction within the block
	TxIndex *int32 `json:"tx_index,omitempty"`
	// The index of the event within the transaction
	EventIndex *int32 `json:"event_index,omitempty"`
	// The event data in a stringified format
	Data *string `json:"data,omitempty"`
}

// NewContractEvent instantiates a new ContractEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractEvent() *ContractEvent {
	this := ContractEvent{}
	return &this
}

// NewContractEventWithDefaults instantiates a new ContractEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractEventWithDefaults() *ContractEvent {
	this := ContractEvent{}
	return &this
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *ContractEvent) GetNetworkName() string {
	if o == nil || IsNil(o.NetworkName) {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetNetworkNameOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkName) {
		return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *ContractEvent) HasNetworkName() bool {
	if o != nil && !IsNil(o.NetworkName) {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *ContractEvent) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetProtocolName returns the ProtocolName field value if set, zero value otherwise.
func (o *ContractEvent) GetProtocolName() string {
	if o == nil || IsNil(o.ProtocolName) {
		var ret string
		return ret
	}
	return *o.ProtocolName
}

// GetProtocolNameOk returns a tuple with the ProtocolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetProtocolNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolName) {
		return nil, false
	}
	return o.ProtocolName, true
}

// HasProtocolName returns a boolean if a field has been set.
func (o *ContractEvent) HasProtocolName() bool {
	if o != nil && !IsNil(o.ProtocolName) {
		return true
	}

	return false
}

// SetProtocolName gets a reference to the given string and assigns it to the ProtocolName field.
func (o *ContractEvent) SetProtocolName(v string) {
	o.ProtocolName = &v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *ContractEvent) GetContractName() string {
	if o == nil || IsNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetContractNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContractName) {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *ContractEvent) HasContractName() bool {
	if o != nil && !IsNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *ContractEvent) SetContractName(v string) {
	o.ContractName = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *ContractEvent) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *ContractEvent) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *ContractEvent) SetEventName(v string) {
	o.EventName = &v
}

// GetSig returns the Sig field value if set, zero value otherwise.
func (o *ContractEvent) GetSig() string {
	if o == nil || IsNil(o.Sig) {
		var ret string
		return ret
	}
	return *o.Sig
}

// GetSigOk returns a tuple with the Sig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetSigOk() (*string, bool) {
	if o == nil || IsNil(o.Sig) {
		return nil, false
	}
	return o.Sig, true
}

// HasSig returns a boolean if a field has been set.
func (o *ContractEvent) HasSig() bool {
	if o != nil && !IsNil(o.Sig) {
		return true
	}

	return false
}

// SetSig gets a reference to the given string and assigns it to the Sig field.
func (o *ContractEvent) SetSig(v string) {
	o.Sig = &v
}

// GetFourBytes returns the FourBytes field value if set, zero value otherwise.
func (o *ContractEvent) GetFourBytes() string {
	if o == nil || IsNil(o.FourBytes) {
		var ret string
		return ret
	}
	return *o.FourBytes
}

// GetFourBytesOk returns a tuple with the FourBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetFourBytesOk() (*string, bool) {
	if o == nil || IsNil(o.FourBytes) {
		return nil, false
	}
	return o.FourBytes, true
}

// HasFourBytes returns a boolean if a field has been set.
func (o *ContractEvent) HasFourBytes() bool {
	if o != nil && !IsNil(o.FourBytes) {
		return true
	}

	return false
}

// SetFourBytes gets a reference to the given string and assigns it to the FourBytes field.
func (o *ContractEvent) SetFourBytes(v string) {
	o.FourBytes = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *ContractEvent) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *ContractEvent) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *ContractEvent) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetBlockTime returns the BlockTime field value if set, zero value otherwise.
func (o *ContractEvent) GetBlockTime() time.Time {
	if o == nil || IsNil(o.BlockTime) {
		var ret time.Time
		return ret
	}
	return *o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetBlockTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BlockTime) {
		return nil, false
	}
	return o.BlockTime, true
}

// HasBlockTime returns a boolean if a field has been set.
func (o *ContractEvent) HasBlockTime() bool {
	if o != nil && !IsNil(o.BlockTime) {
		return true
	}

	return false
}

// SetBlockTime gets a reference to the given time.Time and assigns it to the BlockTime field.
func (o *ContractEvent) SetBlockTime(v time.Time) {
	o.BlockTime = &v
}

// GetBlockHeight returns the BlockHeight field value if set, zero value otherwise.
func (o *ContractEvent) GetBlockHeight() int32 {
	if o == nil || IsNil(o.BlockHeight) {
		var ret int32
		return ret
	}
	return *o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetBlockHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockHeight) {
		return nil, false
	}
	return o.BlockHeight, true
}

// HasBlockHeight returns a boolean if a field has been set.
func (o *ContractEvent) HasBlockHeight() bool {
	if o != nil && !IsNil(o.BlockHeight) {
		return true
	}

	return false
}

// SetBlockHeight gets a reference to the given int32 and assigns it to the BlockHeight field.
func (o *ContractEvent) SetBlockHeight(v int32) {
	o.BlockHeight = &v
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *ContractEvent) GetTxHash() string {
	if o == nil || IsNil(o.TxHash) {
		var ret string
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetTxHashOk() (*string, bool) {
	if o == nil || IsNil(o.TxHash) {
		return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *ContractEvent) HasTxHash() bool {
	if o != nil && !IsNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given string and assigns it to the TxHash field.
func (o *ContractEvent) SetTxHash(v string) {
	o.TxHash = &v
}

// GetTxIndex returns the TxIndex field value if set, zero value otherwise.
func (o *ContractEvent) GetTxIndex() int32 {
	if o == nil || IsNil(o.TxIndex) {
		var ret int32
		return ret
	}
	return *o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetTxIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.TxIndex) {
		return nil, false
	}
	return o.TxIndex, true
}

// HasTxIndex returns a boolean if a field has been set.
func (o *ContractEvent) HasTxIndex() bool {
	if o != nil && !IsNil(o.TxIndex) {
		return true
	}

	return false
}

// SetTxIndex gets a reference to the given int32 and assigns it to the TxIndex field.
func (o *ContractEvent) SetTxIndex(v int32) {
	o.TxIndex = &v
}

// GetEventIndex returns the EventIndex field value if set, zero value otherwise.
func (o *ContractEvent) GetEventIndex() int32 {
	if o == nil || IsNil(o.EventIndex) {
		var ret int32
		return ret
	}
	return *o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetEventIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.EventIndex) {
		return nil, false
	}
	return o.EventIndex, true
}

// HasEventIndex returns a boolean if a field has been set.
func (o *ContractEvent) HasEventIndex() bool {
	if o != nil && !IsNil(o.EventIndex) {
		return true
	}

	return false
}

// SetEventIndex gets a reference to the given int32 and assigns it to the EventIndex field.
func (o *ContractEvent) SetEventIndex(v int32) {
	o.EventIndex = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContractEvent) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContractEvent) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *ContractEvent) SetData(v string) {
	o.Data = &v
}

func (o ContractEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkName) {
		toSerialize["network_name"] = o.NetworkName
	}
	if !IsNil(o.ProtocolName) {
		toSerialize["protocol_name"] = o.ProtocolName
	}
	if !IsNil(o.ContractName) {
		toSerialize["contract_name"] = o.ContractName
	}
	if !IsNil(o.EventName) {
		toSerialize["event_name"] = o.EventName
	}
	if !IsNil(o.Sig) {
		toSerialize["sig"] = o.Sig
	}
	if !IsNil(o.FourBytes) {
		toSerialize["fourBytes"] = o.FourBytes
	}
	if !IsNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if !IsNil(o.BlockTime) {
		toSerialize["block_time"] = o.BlockTime
	}
	if !IsNil(o.BlockHeight) {
		toSerialize["block_height"] = o.BlockHeight
	}
	if !IsNil(o.TxHash) {
		toSerialize["tx_hash"] = o.TxHash
	}
	if !IsNil(o.TxIndex) {
		toSerialize["tx_index"] = o.TxIndex
	}
	if !IsNil(o.EventIndex) {
		toSerialize["event_index"] = o.EventIndex
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableContractEvent struct {
	value *ContractEvent
	isSet bool
}

func (v NullableContractEvent) Get() *ContractEvent {
	return v.value
}

func (v *NullableContractEvent) Set(val *ContractEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableContractEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableContractEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractEvent(val *ContractEvent) *NullableContractEvent {
	return &NullableContractEvent{value: val, isSet: true}
}

func (v NullableContractEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


