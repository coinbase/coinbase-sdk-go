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

// checks if the WebhookEventFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEventFilter{}

// WebhookEventFilter The event_filter parameter specifies the criteria to filter events from the blockchain. It allows filtering events by contract address, sender address and receiver address. For a single event filter, not all of the properties need to be presented.
type WebhookEventFilter struct {
	// The onchain contract address of the token for which the events should be tracked.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The onchain address of the sender. Set this filter to track all transfer events originating from your address.
	FromAddress *string `json:"from_address,omitempty"`
	// The onchain address of the receiver. Set this filter to track all transfer events sent to your address.
	ToAddress *string `json:"to_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebhookEventFilter WebhookEventFilter

// NewWebhookEventFilter instantiates a new WebhookEventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEventFilter() *WebhookEventFilter {
	this := WebhookEventFilter{}
	return &this
}

// NewWebhookEventFilterWithDefaults instantiates a new WebhookEventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventFilterWithDefaults() *WebhookEventFilter {
	this := WebhookEventFilter{}
	return &this
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *WebhookEventFilter) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookEventFilter) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *WebhookEventFilter) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *WebhookEventFilter) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *WebhookEventFilter) GetFromAddress() string {
	if o == nil || IsNil(o.FromAddress) {
		var ret string
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookEventFilter) GetFromAddressOk() (*string, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *WebhookEventFilter) HasFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given string and assigns it to the FromAddress field.
func (o *WebhookEventFilter) SetFromAddress(v string) {
	o.FromAddress = &v
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise.
func (o *WebhookEventFilter) GetToAddress() string {
	if o == nil || IsNil(o.ToAddress) {
		var ret string
		return ret
	}
	return *o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookEventFilter) GetToAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ToAddress) {
		return nil, false
	}
	return o.ToAddress, true
}

// HasToAddress returns a boolean if a field has been set.
func (o *WebhookEventFilter) HasToAddress() bool {
	if o != nil && !IsNil(o.ToAddress) {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given string and assigns it to the ToAddress field.
func (o *WebhookEventFilter) SetToAddress(v string) {
	o.ToAddress = &v
}

func (o WebhookEventFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEventFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if !IsNil(o.FromAddress) {
		toSerialize["from_address"] = o.FromAddress
	}
	if !IsNil(o.ToAddress) {
		toSerialize["to_address"] = o.ToAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookEventFilter) UnmarshalJSON(data []byte) (err error) {
	varWebhookEventFilter := _WebhookEventFilter{}

	err = json.Unmarshal(data, &varWebhookEventFilter)

	if err != nil {
		return err
	}

	*o = WebhookEventFilter(varWebhookEventFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "from_address")
		delete(additionalProperties, "to_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookEventFilter struct {
	value *WebhookEventFilter
	isSet bool
}

func (v NullableWebhookEventFilter) Get() *WebhookEventFilter {
	return v.value
}

func (v *NullableWebhookEventFilter) Set(val *WebhookEventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventFilter(val *WebhookEventFilter) *NullableWebhookEventFilter {
	return &NullableWebhookEventFilter{value: val, isSet: true}
}

func (v NullableWebhookEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


