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

// checks if the WebhookWalletActivityFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookWalletActivityFilter{}

// WebhookWalletActivityFilter Filter for wallet activity events. This filter allows the client to specify one or more wallet addresses to monitor for activities such as transactions, transfers, or other types of events that are associated with the specified addresses. 
type WebhookWalletActivityFilter struct {
	// A list of wallet addresses to filter on.
	Addresses []string `json:"addresses,omitempty"`
	// The ID of the wallet that owns the webhook.
	WalletId *string `json:"wallet_id,omitempty"`
}

// NewWebhookWalletActivityFilter instantiates a new WebhookWalletActivityFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookWalletActivityFilter() *WebhookWalletActivityFilter {
	this := WebhookWalletActivityFilter{}
	return &this
}

// NewWebhookWalletActivityFilterWithDefaults instantiates a new WebhookWalletActivityFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWalletActivityFilterWithDefaults() *WebhookWalletActivityFilter {
	this := WebhookWalletActivityFilter{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *WebhookWalletActivityFilter) GetAddresses() []string {
	if o == nil || IsNil(o.Addresses) {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookWalletActivityFilter) GetAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *WebhookWalletActivityFilter) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *WebhookWalletActivityFilter) SetAddresses(v []string) {
	o.Addresses = v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *WebhookWalletActivityFilter) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookWalletActivityFilter) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *WebhookWalletActivityFilter) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *WebhookWalletActivityFilter) SetWalletId(v string) {
	o.WalletId = &v
}

func (o WebhookWalletActivityFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookWalletActivityFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	return toSerialize, nil
}

type NullableWebhookWalletActivityFilter struct {
	value *WebhookWalletActivityFilter
	isSet bool
}

func (v NullableWebhookWalletActivityFilter) Get() *WebhookWalletActivityFilter {
	return v.value
}

func (v *NullableWebhookWalletActivityFilter) Set(val *WebhookWalletActivityFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookWalletActivityFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookWalletActivityFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookWalletActivityFilter(val *WebhookWalletActivityFilter) *NullableWebhookWalletActivityFilter {
	return &NullableWebhookWalletActivityFilter{value: val, isSet: true}
}

func (v NullableWebhookWalletActivityFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookWalletActivityFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


