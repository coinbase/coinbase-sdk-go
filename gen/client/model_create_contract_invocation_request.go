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

// checks if the CreateContractInvocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContractInvocationRequest{}

// CreateContractInvocationRequest struct for CreateContractInvocationRequest
type CreateContractInvocationRequest struct {
	// The address of the contract to invoke.
	ContractAddress string `json:"contract_address"`
	// The method to invoke on the contract.
	Method string `json:"method"`
	// The JSON-encoded arguments to pass to the contract method. The keys should be the argument names and the values should be the argument values.
	Args string `json:"args"`
	// The JSON-encoded ABI of the contract.
	Abi *string `json:"abi,omitempty"`
	// The amount in atomic units of the native asset to send to the contract for a payable method
	Amount *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateContractInvocationRequest CreateContractInvocationRequest

// NewCreateContractInvocationRequest instantiates a new CreateContractInvocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContractInvocationRequest(contractAddress string, method string, args string) *CreateContractInvocationRequest {
	this := CreateContractInvocationRequest{}
	this.ContractAddress = contractAddress
	this.Method = method
	this.Args = args
	return &this
}

// NewCreateContractInvocationRequestWithDefaults instantiates a new CreateContractInvocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContractInvocationRequestWithDefaults() *CreateContractInvocationRequest {
	this := CreateContractInvocationRequest{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *CreateContractInvocationRequest) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *CreateContractInvocationRequest) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *CreateContractInvocationRequest) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetMethod returns the Method field value
func (o *CreateContractInvocationRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *CreateContractInvocationRequest) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *CreateContractInvocationRequest) SetMethod(v string) {
	o.Method = v
}

// GetArgs returns the Args field value
func (o *CreateContractInvocationRequest) GetArgs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
func (o *CreateContractInvocationRequest) GetArgsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *CreateContractInvocationRequest) SetArgs(v string) {
	o.Args = v
}

// GetAbi returns the Abi field value if set, zero value otherwise.
func (o *CreateContractInvocationRequest) GetAbi() string {
	if o == nil || IsNil(o.Abi) {
		var ret string
		return ret
	}
	return *o.Abi
}

// GetAbiOk returns a tuple with the Abi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContractInvocationRequest) GetAbiOk() (*string, bool) {
	if o == nil || IsNil(o.Abi) {
		return nil, false
	}
	return o.Abi, true
}

// HasAbi returns a boolean if a field has been set.
func (o *CreateContractInvocationRequest) HasAbi() bool {
	if o != nil && !IsNil(o.Abi) {
		return true
	}

	return false
}

// SetAbi gets a reference to the given string and assigns it to the Abi field.
func (o *CreateContractInvocationRequest) SetAbi(v string) {
	o.Abi = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CreateContractInvocationRequest) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContractInvocationRequest) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CreateContractInvocationRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *CreateContractInvocationRequest) SetAmount(v string) {
	o.Amount = &v
}

func (o CreateContractInvocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContractInvocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["method"] = o.Method
	toSerialize["args"] = o.Args
	if !IsNil(o.Abi) {
		toSerialize["abi"] = o.Abi
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateContractInvocationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contract_address",
		"method",
		"args",
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

	varCreateContractInvocationRequest := _CreateContractInvocationRequest{}

	err = json.Unmarshal(data, &varCreateContractInvocationRequest)

	if err != nil {
		return err
	}

	*o = CreateContractInvocationRequest(varCreateContractInvocationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "method")
		delete(additionalProperties, "args")
		delete(additionalProperties, "abi")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateContractInvocationRequest struct {
	value *CreateContractInvocationRequest
	isSet bool
}

func (v NullableCreateContractInvocationRequest) Get() *CreateContractInvocationRequest {
	return v.value
}

func (v *NullableCreateContractInvocationRequest) Set(val *CreateContractInvocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContractInvocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContractInvocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContractInvocationRequest(val *CreateContractInvocationRequest) *NullableCreateContractInvocationRequest {
	return &NullableCreateContractInvocationRequest{value: val, isSet: true}
}

func (v NullableCreateContractInvocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContractInvocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


