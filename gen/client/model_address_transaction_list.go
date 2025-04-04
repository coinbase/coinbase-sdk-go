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

// checks if the AddressTransactionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressTransactionList{}

// AddressTransactionList 
type AddressTransactionList struct {
	Data []Transaction `json:"data"`
	// True if this list has another page of items after this one that can be fetched.
	HasMore bool `json:"has_more"`
	// The page token to be used to fetch the next page.
	NextPage string `json:"next_page"`
	AdditionalProperties map[string]interface{}
}

type _AddressTransactionList AddressTransactionList

// NewAddressTransactionList instantiates a new AddressTransactionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransactionList(data []Transaction, hasMore bool, nextPage string) *AddressTransactionList {
	this := AddressTransactionList{}
	this.Data = data
	this.HasMore = hasMore
	this.NextPage = nextPage
	return &this
}

// NewAddressTransactionListWithDefaults instantiates a new AddressTransactionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransactionListWithDefaults() *AddressTransactionList {
	this := AddressTransactionList{}
	return &this
}

// GetData returns the Data field value
func (o *AddressTransactionList) GetData() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionList) GetDataOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AddressTransactionList) SetData(v []Transaction) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *AddressTransactionList) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionList) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *AddressTransactionList) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextPage returns the NextPage field value
func (o *AddressTransactionList) GetNextPage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *AddressTransactionList) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *AddressTransactionList) SetNextPage(v string) {
	o.NextPage = v
}

func (o AddressTransactionList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransactionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["has_more"] = o.HasMore
	toSerialize["next_page"] = o.NextPage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddressTransactionList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"has_more",
		"next_page",
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

	varAddressTransactionList := _AddressTransactionList{}

	err = json.Unmarshal(data, &varAddressTransactionList)

	if err != nil {
		return err
	}

	*o = AddressTransactionList(varAddressTransactionList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "has_more")
		delete(additionalProperties, "next_page")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddressTransactionList struct {
	value *AddressTransactionList
	isSet bool
}

func (v NullableAddressTransactionList) Get() *AddressTransactionList {
	return v.value
}

func (v *NullableAddressTransactionList) Set(val *AddressTransactionList) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransactionList) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransactionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransactionList(val *AddressTransactionList) *NullableAddressTransactionList {
	return &NullableAddressTransactionList{value: val, isSet: true}
}

func (v NullableAddressTransactionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransactionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


