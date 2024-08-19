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

// checks if the AddressBalanceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressBalanceList{}

// AddressBalanceList 
type AddressBalanceList struct {
	Data []Balance `json:"data"`
	// True if this list has another page of items after this one that can be fetched.
	HasMore bool `json:"has_more"`
	// The page token to be used to fetch the next page.
	NextPage string `json:"next_page"`
	// The total number of balances for the wallet.
	TotalCount int32 `json:"total_count"`
}

type _AddressBalanceList AddressBalanceList

// NewAddressBalanceList instantiates a new AddressBalanceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressBalanceList(data []Balance, hasMore bool, nextPage string, totalCount int32) *AddressBalanceList {
	this := AddressBalanceList{}
	this.Data = data
	this.HasMore = hasMore
	this.NextPage = nextPage
	this.TotalCount = totalCount
	return &this
}

// NewAddressBalanceListWithDefaults instantiates a new AddressBalanceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressBalanceListWithDefaults() *AddressBalanceList {
	this := AddressBalanceList{}
	return &this
}

// GetData returns the Data field value
func (o *AddressBalanceList) GetData() []Balance {
	if o == nil {
		var ret []Balance
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressBalanceList) GetDataOk() ([]Balance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AddressBalanceList) SetData(v []Balance) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *AddressBalanceList) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *AddressBalanceList) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *AddressBalanceList) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextPage returns the NextPage field value
func (o *AddressBalanceList) GetNextPage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *AddressBalanceList) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *AddressBalanceList) SetNextPage(v string) {
	o.NextPage = v
}

// GetTotalCount returns the TotalCount field value
func (o *AddressBalanceList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *AddressBalanceList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *AddressBalanceList) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o AddressBalanceList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressBalanceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["has_more"] = o.HasMore
	toSerialize["next_page"] = o.NextPage
	toSerialize["total_count"] = o.TotalCount
	return toSerialize, nil
}

func (o *AddressBalanceList) UnmarshalJSON(data []byte) (err error) {
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

	varAddressBalanceList := _AddressBalanceList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressBalanceList)

	if err != nil {
		return err
	}

	*o = AddressBalanceList(varAddressBalanceList)

	return err
}

type NullableAddressBalanceList struct {
	value *AddressBalanceList
	isSet bool
}

func (v NullableAddressBalanceList) Get() *AddressBalanceList {
	return v.value
}

func (v *NullableAddressBalanceList) Set(val *AddressBalanceList) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressBalanceList) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressBalanceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressBalanceList(val *AddressBalanceList) *NullableAddressBalanceList {
	return &NullableAddressBalanceList{value: val, isSet: true}
}

func (v NullableAddressBalanceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressBalanceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


