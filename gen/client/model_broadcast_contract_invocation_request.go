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

// checks if the BroadcastContractInvocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastContractInvocationRequest{}

// BroadcastContractInvocationRequest struct for BroadcastContractInvocationRequest
type BroadcastContractInvocationRequest struct {
	// The hex-encoded signed payload of the contract invocation
	SignedPayload string `json:"signed_payload"`
}

type _BroadcastContractInvocationRequest BroadcastContractInvocationRequest

// NewBroadcastContractInvocationRequest instantiates a new BroadcastContractInvocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastContractInvocationRequest(signedPayload string) *BroadcastContractInvocationRequest {
	this := BroadcastContractInvocationRequest{}
	this.SignedPayload = signedPayload
	return &this
}

// NewBroadcastContractInvocationRequestWithDefaults instantiates a new BroadcastContractInvocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastContractInvocationRequestWithDefaults() *BroadcastContractInvocationRequest {
	this := BroadcastContractInvocationRequest{}
	return &this
}

// GetSignedPayload returns the SignedPayload field value
func (o *BroadcastContractInvocationRequest) GetSignedPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedPayload
}

// GetSignedPayloadOk returns a tuple with the SignedPayload field value
// and a boolean to check if the value has been set.
func (o *BroadcastContractInvocationRequest) GetSignedPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedPayload, true
}

// SetSignedPayload sets field value
func (o *BroadcastContractInvocationRequest) SetSignedPayload(v string) {
	o.SignedPayload = v
}

func (o BroadcastContractInvocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastContractInvocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signed_payload"] = o.SignedPayload
	return toSerialize, nil
}

func (o *BroadcastContractInvocationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signed_payload",
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

	varBroadcastContractInvocationRequest := _BroadcastContractInvocationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastContractInvocationRequest)

	if err != nil {
		return err
	}

	*o = BroadcastContractInvocationRequest(varBroadcastContractInvocationRequest)

	return err
}

type NullableBroadcastContractInvocationRequest struct {
	value *BroadcastContractInvocationRequest
	isSet bool
}

func (v NullableBroadcastContractInvocationRequest) Get() *BroadcastContractInvocationRequest {
	return v.value
}

func (v *NullableBroadcastContractInvocationRequest) Set(val *BroadcastContractInvocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastContractInvocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastContractInvocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastContractInvocationRequest(val *BroadcastContractInvocationRequest) *NullableBroadcastContractInvocationRequest {
	return &NullableBroadcastContractInvocationRequest{value: val, isSet: true}
}

func (v NullableBroadcastContractInvocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastContractInvocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

