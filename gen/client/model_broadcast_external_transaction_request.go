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

// checks if the BroadcastExternalTransactionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastExternalTransactionRequest{}

// BroadcastExternalTransactionRequest struct for BroadcastExternalTransactionRequest
type BroadcastExternalTransactionRequest struct {
	// The hex-encoded signed payload of the external address transaction.
	SignedPayload string `json:"signed_payload"`
	AdditionalProperties map[string]interface{}
}

type _BroadcastExternalTransactionRequest BroadcastExternalTransactionRequest

// NewBroadcastExternalTransactionRequest instantiates a new BroadcastExternalTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastExternalTransactionRequest(signedPayload string) *BroadcastExternalTransactionRequest {
	this := BroadcastExternalTransactionRequest{}
	this.SignedPayload = signedPayload
	return &this
}

// NewBroadcastExternalTransactionRequestWithDefaults instantiates a new BroadcastExternalTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastExternalTransactionRequestWithDefaults() *BroadcastExternalTransactionRequest {
	this := BroadcastExternalTransactionRequest{}
	return &this
}

// GetSignedPayload returns the SignedPayload field value
func (o *BroadcastExternalTransactionRequest) GetSignedPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedPayload
}

// GetSignedPayloadOk returns a tuple with the SignedPayload field value
// and a boolean to check if the value has been set.
func (o *BroadcastExternalTransactionRequest) GetSignedPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedPayload, true
}

// SetSignedPayload sets field value
func (o *BroadcastExternalTransactionRequest) SetSignedPayload(v string) {
	o.SignedPayload = v
}

func (o BroadcastExternalTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastExternalTransactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signed_payload"] = o.SignedPayload

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BroadcastExternalTransactionRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBroadcastExternalTransactionRequest := _BroadcastExternalTransactionRequest{}

	err = json.Unmarshal(data, &varBroadcastExternalTransactionRequest)

	if err != nil {
		return err
	}

	*o = BroadcastExternalTransactionRequest(varBroadcastExternalTransactionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signed_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBroadcastExternalTransactionRequest struct {
	value *BroadcastExternalTransactionRequest
	isSet bool
}

func (v NullableBroadcastExternalTransactionRequest) Get() *BroadcastExternalTransactionRequest {
	return v.value
}

func (v *NullableBroadcastExternalTransactionRequest) Set(val *BroadcastExternalTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastExternalTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastExternalTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastExternalTransactionRequest(val *BroadcastExternalTransactionRequest) *NullableBroadcastExternalTransactionRequest {
	return &NullableBroadcastExternalTransactionRequest{value: val, isSet: true}
}

func (v NullableBroadcastExternalTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastExternalTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


