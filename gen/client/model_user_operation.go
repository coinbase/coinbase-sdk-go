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

// checks if the UserOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserOperation{}

// UserOperation struct for UserOperation
type UserOperation struct {
	// The ID of the user operation.
	Id string `json:"id"`
	// The ID of the network the user operation is being created on.
	NetworkId string `json:"network_id"`
	// The list of calls to make from the smart wallet.
	Calls []Call `json:"calls"`
	// The hex-encoded hash that must be signed by the user.
	UnsignedPayload string `json:"unsigned_payload"`
	// The hex-encoded signature of the user operation.
	Signature *string `json:"signature,omitempty"`
	// The status of the user operation.
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserOperation UserOperation

// NewUserOperation instantiates a new UserOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOperation(id string, networkId string, calls []Call, unsignedPayload string) *UserOperation {
	this := UserOperation{}
	this.Id = id
	this.NetworkId = networkId
	this.Calls = calls
	this.UnsignedPayload = unsignedPayload
	return &this
}

// NewUserOperationWithDefaults instantiates a new UserOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserOperationWithDefaults() *UserOperation {
	this := UserOperation{}
	return &this
}

// GetId returns the Id field value
func (o *UserOperation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserOperation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserOperation) SetId(v string) {
	o.Id = v
}

// GetNetworkId returns the NetworkId field value
func (o *UserOperation) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *UserOperation) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *UserOperation) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetCalls returns the Calls field value
func (o *UserOperation) GetCalls() []Call {
	if o == nil {
		var ret []Call
		return ret
	}

	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value
// and a boolean to check if the value has been set.
func (o *UserOperation) GetCallsOk() ([]Call, bool) {
	if o == nil {
		return nil, false
	}
	return o.Calls, true
}

// SetCalls sets field value
func (o *UserOperation) SetCalls(v []Call) {
	o.Calls = v
}

// GetUnsignedPayload returns the UnsignedPayload field value
func (o *UserOperation) GetUnsignedPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnsignedPayload
}

// GetUnsignedPayloadOk returns a tuple with the UnsignedPayload field value
// and a boolean to check if the value has been set.
func (o *UserOperation) GetUnsignedPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnsignedPayload, true
}

// SetUnsignedPayload sets field value
func (o *UserOperation) SetUnsignedPayload(v string) {
	o.UnsignedPayload = v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *UserOperation) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOperation) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *UserOperation) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *UserOperation) SetSignature(v string) {
	o.Signature = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserOperation) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOperation) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserOperation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserOperation) SetStatus(v string) {
	o.Status = &v
}

func (o UserOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["network_id"] = o.NetworkId
	toSerialize["calls"] = o.Calls
	toSerialize["unsigned_payload"] = o.UnsignedPayload
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"network_id",
		"calls",
		"unsigned_payload",
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

	varUserOperation := _UserOperation{}

	err = json.Unmarshal(data, &varUserOperation)

	if err != nil {
		return err
	}

	*o = UserOperation(varUserOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "network_id")
		delete(additionalProperties, "calls")
		delete(additionalProperties, "unsigned_payload")
		delete(additionalProperties, "signature")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserOperation struct {
	value *UserOperation
	isSet bool
}

func (v NullableUserOperation) Get() *UserOperation {
	return v.value
}

func (v *NullableUserOperation) Set(val *UserOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOperation(val *UserOperation) *NullableUserOperation {
	return &NullableUserOperation{value: val, isSet: true}
}

func (v NullableUserOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


