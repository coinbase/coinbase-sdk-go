/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OnchainNameList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnchainNameList{}

// OnchainNameList A list of onchain events with pagination information
type OnchainNameList struct {
	// A list of onchain name objects
	Data []OnchainName `json:"data"`
	// True if this list has another page of items after this one that can be fetched.
	HasMore *bool `json:"has_more,omitempty"`
	// The page token to be used to fetch the next page.
	NextPage string `json:"next_page"`
	// The total number of payload signatures for the address.
	TotalCount *int32 `json:"total_count,omitempty"`
}

type _OnchainNameList OnchainNameList

// NewOnchainNameList instantiates a new OnchainNameList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnchainNameList(data []OnchainName, nextPage string) *OnchainNameList {
	this := OnchainNameList{}
	this.Data = data
	this.NextPage = nextPage
	return &this
}

// NewOnchainNameListWithDefaults instantiates a new OnchainNameList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnchainNameListWithDefaults() *OnchainNameList {
	this := OnchainNameList{}
	return &this
}

// GetData returns the Data field value
func (o *OnchainNameList) GetData() []OnchainName {
	if o == nil {
		var ret []OnchainName
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OnchainNameList) GetDataOk() ([]OnchainName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *OnchainNameList) SetData(v []OnchainName) {
	o.Data = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *OnchainNameList) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnchainNameList) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *OnchainNameList) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *OnchainNameList) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetNextPage returns the NextPage field value
func (o *OnchainNameList) GetNextPage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *OnchainNameList) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *OnchainNameList) SetNextPage(v string) {
	o.NextPage = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *OnchainNameList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnchainNameList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *OnchainNameList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *OnchainNameList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o OnchainNameList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnchainNameList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	toSerialize["next_page"] = o.NextPage
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

func (o *OnchainNameList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varOnchainNameList := _OnchainNameList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnchainNameList)

	if err != nil {
		return err
	}

	*o = OnchainNameList(varOnchainNameList)

	return err
}

type NullableOnchainNameList struct {
	value *OnchainNameList
	isSet bool
}

func (v NullableOnchainNameList) Get() *OnchainNameList {
	return v.value
}

func (v *NullableOnchainNameList) Set(val *OnchainNameList) {
	v.value = val
	v.isSet = true
}

func (v NullableOnchainNameList) IsSet() bool {
	return v.isSet
}

func (v *NullableOnchainNameList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnchainNameList(val *OnchainNameList) *NullableOnchainNameList {
	return &NullableOnchainNameList{value: val, isSet: true}
}

func (v NullableOnchainNameList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnchainNameList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

