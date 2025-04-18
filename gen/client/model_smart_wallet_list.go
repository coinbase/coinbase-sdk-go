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

// checks if the SmartWalletList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartWalletList{}

// SmartWalletList Paginated list of smart wallets
type SmartWalletList struct {
	Data []SmartWallet `json:"data"`
	// True if this list has another page of items after this one that can be fetched.
	HasMore bool `json:"has_more"`
	// The page token to be used to fetch the next page.
	NextPage string `json:"next_page"`
	// The total number of wallets
	TotalCount int32 `json:"total_count"`
	AdditionalProperties map[string]interface{}
}

type _SmartWalletList SmartWalletList

// NewSmartWalletList instantiates a new SmartWalletList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartWalletList(data []SmartWallet, hasMore bool, nextPage string, totalCount int32) *SmartWalletList {
	this := SmartWalletList{}
	this.Data = data
	this.HasMore = hasMore
	this.NextPage = nextPage
	this.TotalCount = totalCount
	return &this
}

// NewSmartWalletListWithDefaults instantiates a new SmartWalletList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartWalletListWithDefaults() *SmartWalletList {
	this := SmartWalletList{}
	return &this
}

// GetData returns the Data field value
func (o *SmartWalletList) GetData() []SmartWallet {
	if o == nil {
		var ret []SmartWallet
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SmartWalletList) GetDataOk() ([]SmartWallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SmartWalletList) SetData(v []SmartWallet) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *SmartWalletList) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *SmartWalletList) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *SmartWalletList) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextPage returns the NextPage field value
func (o *SmartWalletList) GetNextPage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *SmartWalletList) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *SmartWalletList) SetNextPage(v string) {
	o.NextPage = v
}

// GetTotalCount returns the TotalCount field value
func (o *SmartWalletList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SmartWalletList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *SmartWalletList) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o SmartWalletList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartWalletList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["has_more"] = o.HasMore
	toSerialize["next_page"] = o.NextPage
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SmartWalletList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"has_more",
		"next_page",
		"total_count",
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

	varSmartWalletList := _SmartWalletList{}

	err = json.Unmarshal(data, &varSmartWalletList)

	if err != nil {
		return err
	}

	*o = SmartWalletList(varSmartWalletList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "has_more")
		delete(additionalProperties, "next_page")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSmartWalletList struct {
	value *SmartWalletList
	isSet bool
}

func (v NullableSmartWalletList) Get() *SmartWalletList {
	return v.value
}

func (v *NullableSmartWalletList) Set(val *SmartWalletList) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartWalletList) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartWalletList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartWalletList(val *SmartWalletList) *NullableSmartWalletList {
	return &NullableSmartWalletList{value: val, isSet: true}
}

func (v NullableSmartWalletList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartWalletList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


