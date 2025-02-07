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

// checks if the SeedCreationEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeedCreationEvent{}

// SeedCreationEvent An event representing a seed creation.
type SeedCreationEvent struct {
	// The ID of the wallet that the server-signer should create the seed for
	WalletId string `json:"wallet_id"`
	// The ID of the user that the wallet belongs to
	WalletUserId string `json:"wallet_user_id"`
	AdditionalProperties map[string]interface{}
}

type _SeedCreationEvent SeedCreationEvent

// NewSeedCreationEvent instantiates a new SeedCreationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeedCreationEvent(walletId string, walletUserId string) *SeedCreationEvent {
	this := SeedCreationEvent{}
	this.WalletId = walletId
	this.WalletUserId = walletUserId
	return &this
}

// NewSeedCreationEventWithDefaults instantiates a new SeedCreationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeedCreationEventWithDefaults() *SeedCreationEvent {
	this := SeedCreationEvent{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *SeedCreationEvent) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *SeedCreationEvent) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *SeedCreationEvent) SetWalletId(v string) {
	o.WalletId = v
}

// GetWalletUserId returns the WalletUserId field value
func (o *SeedCreationEvent) GetWalletUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletUserId
}

// GetWalletUserIdOk returns a tuple with the WalletUserId field value
// and a boolean to check if the value has been set.
func (o *SeedCreationEvent) GetWalletUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletUserId, true
}

// SetWalletUserId sets field value
func (o *SeedCreationEvent) SetWalletUserId(v string) {
	o.WalletUserId = v
}

func (o SeedCreationEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeedCreationEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["wallet_user_id"] = o.WalletUserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SeedCreationEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet_id",
		"wallet_user_id",
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

	varSeedCreationEvent := _SeedCreationEvent{}

	err = json.Unmarshal(data, &varSeedCreationEvent)

	if err != nil {
		return err
	}

	*o = SeedCreationEvent(varSeedCreationEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "wallet_user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSeedCreationEvent struct {
	value *SeedCreationEvent
	isSet bool
}

func (v NullableSeedCreationEvent) Get() *SeedCreationEvent {
	return v.value
}

func (v *NullableSeedCreationEvent) Set(val *SeedCreationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSeedCreationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSeedCreationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeedCreationEvent(val *SeedCreationEvent) *NullableSeedCreationEvent {
	return &NullableSeedCreationEvent{value: val, isSet: true}
}

func (v NullableSeedCreationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeedCreationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


