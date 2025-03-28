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

// checks if the CreateTradeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeRequest{}

// CreateTradeRequest struct for CreateTradeRequest
type CreateTradeRequest struct {
	// The amount to trade
	Amount string `json:"amount"`
	// The ID of the asset to trade
	FromAssetId string `json:"from_asset_id"`
	// The ID of the asset to receive from the trade
	ToAssetId string `json:"to_asset_id"`
	AdditionalProperties map[string]interface{}
}

type _CreateTradeRequest CreateTradeRequest

// NewCreateTradeRequest instantiates a new CreateTradeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeRequest(amount string, fromAssetId string, toAssetId string) *CreateTradeRequest {
	this := CreateTradeRequest{}
	this.Amount = amount
	this.FromAssetId = fromAssetId
	this.ToAssetId = toAssetId
	return &this
}

// NewCreateTradeRequestWithDefaults instantiates a new CreateTradeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeRequestWithDefaults() *CreateTradeRequest {
	this := CreateTradeRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CreateTradeRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateTradeRequest) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateTradeRequest) SetAmount(v string) {
	o.Amount = v
}

// GetFromAssetId returns the FromAssetId field value
func (o *CreateTradeRequest) GetFromAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAssetId
}

// GetFromAssetIdOk returns a tuple with the FromAssetId field value
// and a boolean to check if the value has been set.
func (o *CreateTradeRequest) GetFromAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAssetId, true
}

// SetFromAssetId sets field value
func (o *CreateTradeRequest) SetFromAssetId(v string) {
	o.FromAssetId = v
}

// GetToAssetId returns the ToAssetId field value
func (o *CreateTradeRequest) GetToAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAssetId
}

// GetToAssetIdOk returns a tuple with the ToAssetId field value
// and a boolean to check if the value has been set.
func (o *CreateTradeRequest) GetToAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAssetId, true
}

// SetToAssetId sets field value
func (o *CreateTradeRequest) SetToAssetId(v string) {
	o.ToAssetId = v
}

func (o CreateTradeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["from_asset_id"] = o.FromAssetId
	toSerialize["to_asset_id"] = o.ToAssetId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateTradeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"from_asset_id",
		"to_asset_id",
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

	varCreateTradeRequest := _CreateTradeRequest{}

	err = json.Unmarshal(data, &varCreateTradeRequest)

	if err != nil {
		return err
	}

	*o = CreateTradeRequest(varCreateTradeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "from_asset_id")
		delete(additionalProperties, "to_asset_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateTradeRequest struct {
	value *CreateTradeRequest
	isSet bool
}

func (v NullableCreateTradeRequest) Get() *CreateTradeRequest {
	return v.value
}

func (v *NullableCreateTradeRequest) Set(val *CreateTradeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeRequest(val *CreateTradeRequest) *NullableCreateTradeRequest {
	return &NullableCreateTradeRequest{value: val, isSet: true}
}

func (v NullableCreateTradeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


