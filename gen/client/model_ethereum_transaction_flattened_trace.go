/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the EthereumTransactionFlattenedTrace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EthereumTransactionFlattenedTrace{}

// EthereumTransactionFlattenedTrace struct for EthereumTransactionFlattenedTrace
type EthereumTransactionFlattenedTrace struct {
	Error *string `json:"error,omitempty"`
	Type *string `json:"type,omitempty"`
	From *string `json:"from,omitempty"`
	To *string `json:"to,omitempty"`
	Value *string `json:"value,omitempty"`
	Gas *int64 `json:"gas,omitempty"`
	GasUsed *int64 `json:"gas_used,omitempty"`
	Input *string `json:"input,omitempty"`
	Output *string `json:"output,omitempty"`
	SubTraces *int64 `json:"sub_traces,omitempty"`
	TraceAddress []int64 `json:"trace_address,omitempty"`
	TraceType *string `json:"trace_type,omitempty"`
	CallType *string `json:"call_type,omitempty"`
	TraceId *string `json:"trace_id,omitempty"`
	Status *int64 `json:"status,omitempty"`
	BlockHash *string `json:"block_hash,omitempty"`
	BlockNumber *int64 `json:"block_number,omitempty"`
	TransactionHash *string `json:"transaction_hash,omitempty"`
	TransactionIndex *int64 `json:"transaction_index,omitempty"`
}

// NewEthereumTransactionFlattenedTrace instantiates a new EthereumTransactionFlattenedTrace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthereumTransactionFlattenedTrace() *EthereumTransactionFlattenedTrace {
	this := EthereumTransactionFlattenedTrace{}
	return &this
}

// NewEthereumTransactionFlattenedTraceWithDefaults instantiates a new EthereumTransactionFlattenedTrace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthereumTransactionFlattenedTraceWithDefaults() *EthereumTransactionFlattenedTrace {
	this := EthereumTransactionFlattenedTrace{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *EthereumTransactionFlattenedTrace) SetError(v string) {
	o.Error = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EthereumTransactionFlattenedTrace) SetType(v string) {
	o.Type = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *EthereumTransactionFlattenedTrace) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *EthereumTransactionFlattenedTrace) SetTo(v string) {
	o.To = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EthereumTransactionFlattenedTrace) SetValue(v string) {
	o.Value = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetGas() int64 {
	if o == nil || IsNil(o.Gas) {
		var ret int64
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetGasOk() (*int64, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given int64 and assigns it to the Gas field.
func (o *EthereumTransactionFlattenedTrace) SetGas(v int64) {
	o.Gas = &v
}

// GetGasUsed returns the GasUsed field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetGasUsed() int64 {
	if o == nil || IsNil(o.GasUsed) {
		var ret int64
		return ret
	}
	return *o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetGasUsedOk() (*int64, bool) {
	if o == nil || IsNil(o.GasUsed) {
		return nil, false
	}
	return o.GasUsed, true
}

// HasGasUsed returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasGasUsed() bool {
	if o != nil && !IsNil(o.GasUsed) {
		return true
	}

	return false
}

// SetGasUsed gets a reference to the given int64 and assigns it to the GasUsed field.
func (o *EthereumTransactionFlattenedTrace) SetGasUsed(v int64) {
	o.GasUsed = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetInput() string {
	if o == nil || IsNil(o.Input) {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetInputOk() (*string, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *EthereumTransactionFlattenedTrace) SetInput(v string) {
	o.Input = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetOutput() string {
	if o == nil || IsNil(o.Output) {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetOutputOk() (*string, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *EthereumTransactionFlattenedTrace) SetOutput(v string) {
	o.Output = &v
}

// GetSubTraces returns the SubTraces field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetSubTraces() int64 {
	if o == nil || IsNil(o.SubTraces) {
		var ret int64
		return ret
	}
	return *o.SubTraces
}

// GetSubTracesOk returns a tuple with the SubTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetSubTracesOk() (*int64, bool) {
	if o == nil || IsNil(o.SubTraces) {
		return nil, false
	}
	return o.SubTraces, true
}

// HasSubTraces returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasSubTraces() bool {
	if o != nil && !IsNil(o.SubTraces) {
		return true
	}

	return false
}

// SetSubTraces gets a reference to the given int64 and assigns it to the SubTraces field.
func (o *EthereumTransactionFlattenedTrace) SetSubTraces(v int64) {
	o.SubTraces = &v
}

// GetTraceAddress returns the TraceAddress field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetTraceAddress() []int64 {
	if o == nil || IsNil(o.TraceAddress) {
		var ret []int64
		return ret
	}
	return o.TraceAddress
}

// GetTraceAddressOk returns a tuple with the TraceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetTraceAddressOk() ([]int64, bool) {
	if o == nil || IsNil(o.TraceAddress) {
		return nil, false
	}
	return o.TraceAddress, true
}

// HasTraceAddress returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasTraceAddress() bool {
	if o != nil && !IsNil(o.TraceAddress) {
		return true
	}

	return false
}

// SetTraceAddress gets a reference to the given []int64 and assigns it to the TraceAddress field.
func (o *EthereumTransactionFlattenedTrace) SetTraceAddress(v []int64) {
	o.TraceAddress = v
}

// GetTraceType returns the TraceType field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetTraceType() string {
	if o == nil || IsNil(o.TraceType) {
		var ret string
		return ret
	}
	return *o.TraceType
}

// GetTraceTypeOk returns a tuple with the TraceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetTraceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TraceType) {
		return nil, false
	}
	return o.TraceType, true
}

// HasTraceType returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasTraceType() bool {
	if o != nil && !IsNil(o.TraceType) {
		return true
	}

	return false
}

// SetTraceType gets a reference to the given string and assigns it to the TraceType field.
func (o *EthereumTransactionFlattenedTrace) SetTraceType(v string) {
	o.TraceType = &v
}

// GetCallType returns the CallType field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetCallType() string {
	if o == nil || IsNil(o.CallType) {
		var ret string
		return ret
	}
	return *o.CallType
}

// GetCallTypeOk returns a tuple with the CallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetCallTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CallType) {
		return nil, false
	}
	return o.CallType, true
}

// HasCallType returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasCallType() bool {
	if o != nil && !IsNil(o.CallType) {
		return true
	}

	return false
}

// SetCallType gets a reference to the given string and assigns it to the CallType field.
func (o *EthereumTransactionFlattenedTrace) SetCallType(v string) {
	o.CallType = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *EthereumTransactionFlattenedTrace) SetTraceId(v string) {
	o.TraceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *EthereumTransactionFlattenedTrace) SetStatus(v int64) {
	o.Status = &v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash) {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetBlockHashOk() (*string, bool) {
	if o == nil || IsNil(o.BlockHash) {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasBlockHash() bool {
	if o != nil && !IsNil(o.BlockHash) {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *EthereumTransactionFlattenedTrace) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetBlockNumber() int64 {
	if o == nil || IsNil(o.BlockNumber) {
		var ret int64
		return ret
	}
	return *o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetBlockNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockNumber) {
		return nil, false
	}
	return o.BlockNumber, true
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasBlockNumber() bool {
	if o != nil && !IsNil(o.BlockNumber) {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given int64 and assigns it to the BlockNumber field.
func (o *EthereumTransactionFlattenedTrace) SetBlockNumber(v int64) {
	o.BlockNumber = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *EthereumTransactionFlattenedTrace) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

// GetTransactionIndex returns the TransactionIndex field value if set, zero value otherwise.
func (o *EthereumTransactionFlattenedTrace) GetTransactionIndex() int64 {
	if o == nil || IsNil(o.TransactionIndex) {
		var ret int64
		return ret
	}
	return *o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionFlattenedTrace) GetTransactionIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionIndex) {
		return nil, false
	}
	return o.TransactionIndex, true
}

// HasTransactionIndex returns a boolean if a field has been set.
func (o *EthereumTransactionFlattenedTrace) HasTransactionIndex() bool {
	if o != nil && !IsNil(o.TransactionIndex) {
		return true
	}

	return false
}

// SetTransactionIndex gets a reference to the given int64 and assigns it to the TransactionIndex field.
func (o *EthereumTransactionFlattenedTrace) SetTransactionIndex(v int64) {
	o.TransactionIndex = &v
}

func (o EthereumTransactionFlattenedTrace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthereumTransactionFlattenedTrace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !IsNil(o.GasUsed) {
		toSerialize["gas_used"] = o.GasUsed
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.SubTraces) {
		toSerialize["sub_traces"] = o.SubTraces
	}
	if !IsNil(o.TraceAddress) {
		toSerialize["trace_address"] = o.TraceAddress
	}
	if !IsNil(o.TraceType) {
		toSerialize["trace_type"] = o.TraceType
	}
	if !IsNil(o.CallType) {
		toSerialize["call_type"] = o.CallType
	}
	if !IsNil(o.TraceId) {
		toSerialize["trace_id"] = o.TraceId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.BlockHash) {
		toSerialize["block_hash"] = o.BlockHash
	}
	if !IsNil(o.BlockNumber) {
		toSerialize["block_number"] = o.BlockNumber
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transaction_hash"] = o.TransactionHash
	}
	if !IsNil(o.TransactionIndex) {
		toSerialize["transaction_index"] = o.TransactionIndex
	}
	return toSerialize, nil
}

type NullableEthereumTransactionFlattenedTrace struct {
	value *EthereumTransactionFlattenedTrace
	isSet bool
}

func (v NullableEthereumTransactionFlattenedTrace) Get() *EthereumTransactionFlattenedTrace {
	return v.value
}

func (v *NullableEthereumTransactionFlattenedTrace) Set(val *EthereumTransactionFlattenedTrace) {
	v.value = val
	v.isSet = true
}

func (v NullableEthereumTransactionFlattenedTrace) IsSet() bool {
	return v.isSet
}

func (v *NullableEthereumTransactionFlattenedTrace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthereumTransactionFlattenedTrace(val *EthereumTransactionFlattenedTrace) *NullableEthereumTransactionFlattenedTrace {
	return &NullableEthereumTransactionFlattenedTrace{value: val, isSet: true}
}

func (v NullableEthereumTransactionFlattenedTrace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthereumTransactionFlattenedTrace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


