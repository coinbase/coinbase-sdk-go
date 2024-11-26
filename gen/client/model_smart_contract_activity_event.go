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
)

// checks if the SmartContractActivityEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartContractActivityEvent{}

// SmartContractActivityEvent Represents an event triggered by a smart contract activity on the blockchain. Contains information about the function, transaction, block, and involved addresses.
type SmartContractActivityEvent struct {
	// Unique identifier for the webhook that triggered this event.
	WebhookId *string `json:"webhookId,omitempty"`
	// Type of event, in this case, an ERC-721 token transfer.
	EventType *string `json:"eventType,omitempty"`
	// Blockchain network where the event occurred.
	Network *string `json:"network,omitempty"`
	// Name of the project this smart contract belongs to.
	ProjectName *string `json:"projectName,omitempty"`
	// Name of the contract.
	ContractName *string `json:"contractName,omitempty"`
	// Name of the function.
	Func *string `json:"func,omitempty"`
	// Signature of the function.
	Sig *string `json:"sig,omitempty"`
	// First 4 bytes of the Transaction, a unique ID.
	FourBytes *string `json:"fourBytes,omitempty"`
	// Address of the smart contract.
	ContractAddress *string `json:"contractAddress,omitempty"`
	// Hash of the block containing the transaction.
	BlockHash *string `json:"blockHash,omitempty"`
	// Number of the block containing the transaction.
	BlockNumber *int32 `json:"blockNumber,omitempty"`
	// Timestamp when the block was mined.
	BlockTime *time.Time `json:"blockTime,omitempty"`
	// Hash of the transaction that triggered the event.
	TransactionHash *string `json:"transactionHash,omitempty"`
	// Position of the transaction within the block.
	TransactionIndex *int32 `json:"transactionIndex,omitempty"`
	// Position of the event log within the transaction.
	LogIndex *int32 `json:"logIndex,omitempty"`
	// Address of the initiator in the transfer.
	From *string `json:"from,omitempty"`
	// Address of the recipient in the transfer.
	To *string `json:"to,omitempty"`
	// Amount of tokens transferred, typically in the smallest unit (e.g., wei for Ethereum).
	Value *int32 `json:"value,omitempty"`
}

// NewSmartContractActivityEvent instantiates a new SmartContractActivityEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractActivityEvent() *SmartContractActivityEvent {
	this := SmartContractActivityEvent{}
	return &this
}

// NewSmartContractActivityEventWithDefaults instantiates a new SmartContractActivityEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractActivityEventWithDefaults() *SmartContractActivityEvent {
	this := SmartContractActivityEvent{}
	return &this
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetWebhookId() string {
	if o == nil || IsNil(o.WebhookId) {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetWebhookIdOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookId) {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasWebhookId() bool {
	if o != nil && !IsNil(o.WebhookId) {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *SmartContractActivityEvent) SetWebhookId(v string) {
	o.WebhookId = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *SmartContractActivityEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *SmartContractActivityEvent) SetNetwork(v string) {
	o.Network = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *SmartContractActivityEvent) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetContractName() string {
	if o == nil || IsNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetContractNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContractName) {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasContractName() bool {
	if o != nil && !IsNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *SmartContractActivityEvent) SetContractName(v string) {
	o.ContractName = &v
}

// GetFunc returns the Func field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetFunc() string {
	if o == nil || IsNil(o.Func) {
		var ret string
		return ret
	}
	return *o.Func
}

// GetFuncOk returns a tuple with the Func field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetFuncOk() (*string, bool) {
	if o == nil || IsNil(o.Func) {
		return nil, false
	}
	return o.Func, true
}

// HasFunc returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasFunc() bool {
	if o != nil && !IsNil(o.Func) {
		return true
	}

	return false
}

// SetFunc gets a reference to the given string and assigns it to the Func field.
func (o *SmartContractActivityEvent) SetFunc(v string) {
	o.Func = &v
}

// GetSig returns the Sig field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetSig() string {
	if o == nil || IsNil(o.Sig) {
		var ret string
		return ret
	}
	return *o.Sig
}

// GetSigOk returns a tuple with the Sig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetSigOk() (*string, bool) {
	if o == nil || IsNil(o.Sig) {
		return nil, false
	}
	return o.Sig, true
}

// HasSig returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasSig() bool {
	if o != nil && !IsNil(o.Sig) {
		return true
	}

	return false
}

// SetSig gets a reference to the given string and assigns it to the Sig field.
func (o *SmartContractActivityEvent) SetSig(v string) {
	o.Sig = &v
}

// GetFourBytes returns the FourBytes field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetFourBytes() string {
	if o == nil || IsNil(o.FourBytes) {
		var ret string
		return ret
	}
	return *o.FourBytes
}

// GetFourBytesOk returns a tuple with the FourBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetFourBytesOk() (*string, bool) {
	if o == nil || IsNil(o.FourBytes) {
		return nil, false
	}
	return o.FourBytes, true
}

// HasFourBytes returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasFourBytes() bool {
	if o != nil && !IsNil(o.FourBytes) {
		return true
	}

	return false
}

// SetFourBytes gets a reference to the given string and assigns it to the FourBytes field.
func (o *SmartContractActivityEvent) SetFourBytes(v string) {
	o.FourBytes = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *SmartContractActivityEvent) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash) {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetBlockHashOk() (*string, bool) {
	if o == nil || IsNil(o.BlockHash) {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasBlockHash() bool {
	if o != nil && !IsNil(o.BlockHash) {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *SmartContractActivityEvent) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetBlockNumber() int32 {
	if o == nil || IsNil(o.BlockNumber) {
		var ret int32
		return ret
	}
	return *o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetBlockNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockNumber) {
		return nil, false
	}
	return o.BlockNumber, true
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasBlockNumber() bool {
	if o != nil && !IsNil(o.BlockNumber) {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given int32 and assigns it to the BlockNumber field.
func (o *SmartContractActivityEvent) SetBlockNumber(v int32) {
	o.BlockNumber = &v
}

// GetBlockTime returns the BlockTime field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetBlockTime() time.Time {
	if o == nil || IsNil(o.BlockTime) {
		var ret time.Time
		return ret
	}
	return *o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetBlockTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BlockTime) {
		return nil, false
	}
	return o.BlockTime, true
}

// HasBlockTime returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasBlockTime() bool {
	if o != nil && !IsNil(o.BlockTime) {
		return true
	}

	return false
}

// SetBlockTime gets a reference to the given time.Time and assigns it to the BlockTime field.
func (o *SmartContractActivityEvent) SetBlockTime(v time.Time) {
	o.BlockTime = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *SmartContractActivityEvent) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

// GetTransactionIndex returns the TransactionIndex field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetTransactionIndex() int32 {
	if o == nil || IsNil(o.TransactionIndex) {
		var ret int32
		return ret
	}
	return *o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetTransactionIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.TransactionIndex) {
		return nil, false
	}
	return o.TransactionIndex, true
}

// HasTransactionIndex returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasTransactionIndex() bool {
	if o != nil && !IsNil(o.TransactionIndex) {
		return true
	}

	return false
}

// SetTransactionIndex gets a reference to the given int32 and assigns it to the TransactionIndex field.
func (o *SmartContractActivityEvent) SetTransactionIndex(v int32) {
	o.TransactionIndex = &v
}

// GetLogIndex returns the LogIndex field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetLogIndex() int32 {
	if o == nil || IsNil(o.LogIndex) {
		var ret int32
		return ret
	}
	return *o.LogIndex
}

// GetLogIndexOk returns a tuple with the LogIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetLogIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.LogIndex) {
		return nil, false
	}
	return o.LogIndex, true
}

// HasLogIndex returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasLogIndex() bool {
	if o != nil && !IsNil(o.LogIndex) {
		return true
	}

	return false
}

// SetLogIndex gets a reference to the given int32 and assigns it to the LogIndex field.
func (o *SmartContractActivityEvent) SetLogIndex(v int32) {
	o.LogIndex = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *SmartContractActivityEvent) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *SmartContractActivityEvent) SetTo(v string) {
	o.To = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SmartContractActivityEvent) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractActivityEvent) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SmartContractActivityEvent) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *SmartContractActivityEvent) SetValue(v int32) {
	o.Value = &v
}

func (o SmartContractActivityEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractActivityEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WebhookId) {
		toSerialize["webhookId"] = o.WebhookId
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ProjectName) {
		toSerialize["projectName"] = o.ProjectName
	}
	if !IsNil(o.ContractName) {
		toSerialize["contractName"] = o.ContractName
	}
	if !IsNil(o.Func) {
		toSerialize["func"] = o.Func
	}
	if !IsNil(o.Sig) {
		toSerialize["sig"] = o.Sig
	}
	if !IsNil(o.FourBytes) {
		toSerialize["fourBytes"] = o.FourBytes
	}
	if !IsNil(o.ContractAddress) {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if !IsNil(o.BlockHash) {
		toSerialize["blockHash"] = o.BlockHash
	}
	if !IsNil(o.BlockNumber) {
		toSerialize["blockNumber"] = o.BlockNumber
	}
	if !IsNil(o.BlockTime) {
		toSerialize["blockTime"] = o.BlockTime
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transactionHash"] = o.TransactionHash
	}
	if !IsNil(o.TransactionIndex) {
		toSerialize["transactionIndex"] = o.TransactionIndex
	}
	if !IsNil(o.LogIndex) {
		toSerialize["logIndex"] = o.LogIndex
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSmartContractActivityEvent struct {
	value *SmartContractActivityEvent
	isSet bool
}

func (v NullableSmartContractActivityEvent) Get() *SmartContractActivityEvent {
	return v.value
}

func (v *NullableSmartContractActivityEvent) Set(val *SmartContractActivityEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractActivityEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractActivityEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractActivityEvent(val *SmartContractActivityEvent) *NullableSmartContractActivityEvent {
	return &NullableSmartContractActivityEvent{value: val, isSet: true}
}

func (v NullableSmartContractActivityEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractActivityEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


