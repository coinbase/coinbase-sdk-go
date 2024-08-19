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
	"bytes"
	"fmt"
)

// checks if the BroadcastTradeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastTradeRequest{}

// BroadcastTradeRequest struct for BroadcastTradeRequest
type BroadcastTradeRequest struct {
	// The hex-encoded signed payload of the trade
	SignedPayload string `json:"signed_payload"`
	// The hex-encoded signed payload of the approval transaction
	ApproveTransactionSignedPayload *string `json:"approve_transaction_signed_payload,omitempty"`
}

type _BroadcastTradeRequest BroadcastTradeRequest

// NewBroadcastTradeRequest instantiates a new BroadcastTradeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastTradeRequest(signedPayload string) *BroadcastTradeRequest {
	this := BroadcastTradeRequest{}
	this.SignedPayload = signedPayload
	return &this
}

// NewBroadcastTradeRequestWithDefaults instantiates a new BroadcastTradeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastTradeRequestWithDefaults() *BroadcastTradeRequest {
	this := BroadcastTradeRequest{}
	return &this
}

// GetSignedPayload returns the SignedPayload field value
func (o *BroadcastTradeRequest) GetSignedPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedPayload
}

// GetSignedPayloadOk returns a tuple with the SignedPayload field value
// and a boolean to check if the value has been set.
func (o *BroadcastTradeRequest) GetSignedPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedPayload, true
}

// SetSignedPayload sets field value
func (o *BroadcastTradeRequest) SetSignedPayload(v string) {
	o.SignedPayload = v
}

// GetApproveTransactionSignedPayload returns the ApproveTransactionSignedPayload field value if set, zero value otherwise.
func (o *BroadcastTradeRequest) GetApproveTransactionSignedPayload() string {
	if o == nil || IsNil(o.ApproveTransactionSignedPayload) {
		var ret string
		return ret
	}
	return *o.ApproveTransactionSignedPayload
}

// GetApproveTransactionSignedPayloadOk returns a tuple with the ApproveTransactionSignedPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTradeRequest) GetApproveTransactionSignedPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.ApproveTransactionSignedPayload) {
		return nil, false
	}
	return o.ApproveTransactionSignedPayload, true
}

// HasApproveTransactionSignedPayload returns a boolean if a field has been set.
func (o *BroadcastTradeRequest) HasApproveTransactionSignedPayload() bool {
	if o != nil && !IsNil(o.ApproveTransactionSignedPayload) {
		return true
	}

	return false
}

// SetApproveTransactionSignedPayload gets a reference to the given string and assigns it to the ApproveTransactionSignedPayload field.
func (o *BroadcastTradeRequest) SetApproveTransactionSignedPayload(v string) {
	o.ApproveTransactionSignedPayload = &v
}

func (o BroadcastTradeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastTradeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signed_payload"] = o.SignedPayload
	if !IsNil(o.ApproveTransactionSignedPayload) {
		toSerialize["approve_transaction_signed_payload"] = o.ApproveTransactionSignedPayload
	}
	return toSerialize, nil
}

func (o *BroadcastTradeRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBroadcastTradeRequest := _BroadcastTradeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastTradeRequest)

	if err != nil {
		return err
	}

	*o = BroadcastTradeRequest(varBroadcastTradeRequest)

	return err
}

type NullableBroadcastTradeRequest struct {
	value *BroadcastTradeRequest
	isSet bool
}

func (v NullableBroadcastTradeRequest) Get() *BroadcastTradeRequest {
	return v.value
}

func (v *NullableBroadcastTradeRequest) Set(val *BroadcastTradeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTradeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTradeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTradeRequest(val *BroadcastTradeRequest) *NullableBroadcastTradeRequest {
	return &NullableBroadcastTradeRequest{value: val, isSet: true}
}

func (v NullableBroadcastTradeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTradeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


