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

// checks if the Network type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Network{}

// Network struct for Network
type Network struct {
	Id NetworkIdentifier `json:"id"`
	// The human-readable name of the blockchain network
	DisplayName string `json:"display_name"`
	// The chain ID of the blockchain network
	ChainId int32 `json:"chain_id"`
	// The protocol family of the blockchain network
	ProtocolFamily string `json:"protocol_family"`
	// Whether the network is a testnet or not
	IsTestnet bool `json:"is_testnet"`
	NativeAsset Asset `json:"native_asset"`
	FeatureSet FeatureSet `json:"feature_set"`
	// The BIP44 path prefix for the network
	AddressPathPrefix *string `json:"address_path_prefix,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Network Network

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork(id NetworkIdentifier, displayName string, chainId int32, protocolFamily string, isTestnet bool, nativeAsset Asset, featureSet FeatureSet) *Network {
	this := Network{}
	this.Id = id
	this.DisplayName = displayName
	this.ChainId = chainId
	this.ProtocolFamily = protocolFamily
	this.IsTestnet = isTestnet
	this.NativeAsset = nativeAsset
	this.FeatureSet = featureSet
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetId returns the Id field value
func (o *Network) GetId() NetworkIdentifier {
	if o == nil {
		var ret NetworkIdentifier
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Network) GetIdOk() (*NetworkIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Network) SetId(v NetworkIdentifier) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value
func (o *Network) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Network) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Network) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetChainId returns the ChainId field value
func (o *Network) GetChainId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *Network) GetChainIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *Network) SetChainId(v int32) {
	o.ChainId = v
}

// GetProtocolFamily returns the ProtocolFamily field value
func (o *Network) GetProtocolFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtocolFamily
}

// GetProtocolFamilyOk returns a tuple with the ProtocolFamily field value
// and a boolean to check if the value has been set.
func (o *Network) GetProtocolFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtocolFamily, true
}

// SetProtocolFamily sets field value
func (o *Network) SetProtocolFamily(v string) {
	o.ProtocolFamily = v
}

// GetIsTestnet returns the IsTestnet field value
func (o *Network) GetIsTestnet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTestnet
}

// GetIsTestnetOk returns a tuple with the IsTestnet field value
// and a boolean to check if the value has been set.
func (o *Network) GetIsTestnetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTestnet, true
}

// SetIsTestnet sets field value
func (o *Network) SetIsTestnet(v bool) {
	o.IsTestnet = v
}

// GetNativeAsset returns the NativeAsset field value
func (o *Network) GetNativeAsset() Asset {
	if o == nil {
		var ret Asset
		return ret
	}

	return o.NativeAsset
}

// GetNativeAssetOk returns a tuple with the NativeAsset field value
// and a boolean to check if the value has been set.
func (o *Network) GetNativeAssetOk() (*Asset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NativeAsset, true
}

// SetNativeAsset sets field value
func (o *Network) SetNativeAsset(v Asset) {
	o.NativeAsset = v
}

// GetFeatureSet returns the FeatureSet field value
func (o *Network) GetFeatureSet() FeatureSet {
	if o == nil {
		var ret FeatureSet
		return ret
	}

	return o.FeatureSet
}

// GetFeatureSetOk returns a tuple with the FeatureSet field value
// and a boolean to check if the value has been set.
func (o *Network) GetFeatureSetOk() (*FeatureSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureSet, true
}

// SetFeatureSet sets field value
func (o *Network) SetFeatureSet(v FeatureSet) {
	o.FeatureSet = v
}

// GetAddressPathPrefix returns the AddressPathPrefix field value if set, zero value otherwise.
func (o *Network) GetAddressPathPrefix() string {
	if o == nil || IsNil(o.AddressPathPrefix) {
		var ret string
		return ret
	}
	return *o.AddressPathPrefix
}

// GetAddressPathPrefixOk returns a tuple with the AddressPathPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetAddressPathPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.AddressPathPrefix) {
		return nil, false
	}
	return o.AddressPathPrefix, true
}

// HasAddressPathPrefix returns a boolean if a field has been set.
func (o *Network) HasAddressPathPrefix() bool {
	if o != nil && !IsNil(o.AddressPathPrefix) {
		return true
	}

	return false
}

// SetAddressPathPrefix gets a reference to the given string and assigns it to the AddressPathPrefix field.
func (o *Network) SetAddressPathPrefix(v string) {
	o.AddressPathPrefix = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Network) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["display_name"] = o.DisplayName
	toSerialize["chain_id"] = o.ChainId
	toSerialize["protocol_family"] = o.ProtocolFamily
	toSerialize["is_testnet"] = o.IsTestnet
	toSerialize["native_asset"] = o.NativeAsset
	toSerialize["feature_set"] = o.FeatureSet
	if !IsNil(o.AddressPathPrefix) {
		toSerialize["address_path_prefix"] = o.AddressPathPrefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Network) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"display_name",
		"chain_id",
		"protocol_family",
		"is_testnet",
		"native_asset",
		"feature_set",
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

	varNetwork := _Network{}

	err = json.Unmarshal(data, &varNetwork)

	if err != nil {
		return err
	}

	*o = Network(varNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "protocol_family")
		delete(additionalProperties, "is_testnet")
		delete(additionalProperties, "native_asset")
		delete(additionalProperties, "feature_set")
		delete(additionalProperties, "address_path_prefix")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


