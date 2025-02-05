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

// checks if the RegisterSmartContractRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterSmartContractRequest{}

// RegisterSmartContractRequest Smart Contract data to be registered
type RegisterSmartContractRequest struct {
	// ABI of the smart contract
	Abi string `json:"abi"`
	// Name of the smart contract
	ContractName *string `json:"contract_name,omitempty"`
}

type _RegisterSmartContractRequest RegisterSmartContractRequest

// NewRegisterSmartContractRequest instantiates a new RegisterSmartContractRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterSmartContractRequest(abi string) *RegisterSmartContractRequest {
	this := RegisterSmartContractRequest{}
	this.Abi = abi
	return &this
}

// NewRegisterSmartContractRequestWithDefaults instantiates a new RegisterSmartContractRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterSmartContractRequestWithDefaults() *RegisterSmartContractRequest {
	this := RegisterSmartContractRequest{}
	return &this
}

// GetAbi returns the Abi field value
func (o *RegisterSmartContractRequest) GetAbi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Abi
}

// GetAbiOk returns a tuple with the Abi field value
// and a boolean to check if the value has been set.
func (o *RegisterSmartContractRequest) GetAbiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Abi, true
}

// SetAbi sets field value
func (o *RegisterSmartContractRequest) SetAbi(v string) {
	o.Abi = v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *RegisterSmartContractRequest) GetContractName() string {
	if o == nil || IsNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSmartContractRequest) GetContractNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContractName) {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *RegisterSmartContractRequest) HasContractName() bool {
	if o != nil && !IsNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *RegisterSmartContractRequest) SetContractName(v string) {
	o.ContractName = &v
}

func (o RegisterSmartContractRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterSmartContractRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["abi"] = o.Abi
	if !IsNil(o.ContractName) {
		toSerialize["contract_name"] = o.ContractName
	}
	return toSerialize, nil
}

func (o *RegisterSmartContractRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"abi",
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

	varRegisterSmartContractRequest := _RegisterSmartContractRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegisterSmartContractRequest)

	if err != nil {
		return err
	}

	*o = RegisterSmartContractRequest(varRegisterSmartContractRequest)

	return err
}

type NullableRegisterSmartContractRequest struct {
	value *RegisterSmartContractRequest
	isSet bool
}

func (v NullableRegisterSmartContractRequest) Get() *RegisterSmartContractRequest {
	return v.value
}

func (v *NullableRegisterSmartContractRequest) Set(val *RegisterSmartContractRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterSmartContractRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterSmartContractRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterSmartContractRequest(val *RegisterSmartContractRequest) *NullableRegisterSmartContractRequest {
	return &NullableRegisterSmartContractRequest{value: val, isSet: true}
}

func (v NullableRegisterSmartContractRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterSmartContractRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


