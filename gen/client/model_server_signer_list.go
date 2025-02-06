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

// checks if the ServerSignerList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerSignerList{}

// ServerSignerList 
type ServerSignerList struct {
	Data []ServerSigner `json:"data"`
	// True if this list has another page of items after this one that can be fetched.
	HasMore bool `json:"has_more"`
	// The page token to be used to fetch the next page.
	NextPage string `json:"next_page"`
	// The total number of server-signers for the project.
	TotalCount int32 `json:"total_count"`
	AdditionalProperties map[string]interface{}
}

type _ServerSignerList ServerSignerList

// NewServerSignerList instantiates a new ServerSignerList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerSignerList(data []ServerSigner, hasMore bool, nextPage string, totalCount int32) *ServerSignerList {
	this := ServerSignerList{}
	this.Data = data
	this.HasMore = hasMore
	this.NextPage = nextPage
	this.TotalCount = totalCount
	return &this
}

// NewServerSignerListWithDefaults instantiates a new ServerSignerList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerSignerListWithDefaults() *ServerSignerList {
	this := ServerSignerList{}
	return &this
}

// GetData returns the Data field value
func (o *ServerSignerList) GetData() []ServerSigner {
	if o == nil {
		var ret []ServerSigner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ServerSignerList) GetDataOk() ([]ServerSigner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ServerSignerList) SetData(v []ServerSigner) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *ServerSignerList) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ServerSignerList) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ServerSignerList) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextPage returns the NextPage field value
func (o *ServerSignerList) GetNextPage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *ServerSignerList) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *ServerSignerList) SetNextPage(v string) {
	o.NextPage = v
}

// GetTotalCount returns the TotalCount field value
func (o *ServerSignerList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ServerSignerList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ServerSignerList) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o ServerSignerList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerSignerList) ToMap() (map[string]interface{}, error) {
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

func (o *ServerSignerList) UnmarshalJSON(data []byte) (err error) {
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

	varServerSignerList := _ServerSignerList{}

	err = json.Unmarshal(data, &varServerSignerList)

	if err != nil {
		return err
	}

	*o = ServerSignerList(varServerSignerList)

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

type NullableServerSignerList struct {
	value *ServerSignerList
	isSet bool
}

func (v NullableServerSignerList) Get() *ServerSignerList {
	return v.value
}

func (v *NullableServerSignerList) Set(val *ServerSignerList) {
	v.value = val
	v.isSet = true
}

func (v NullableServerSignerList) IsSet() bool {
	return v.isSet
}

func (v *NullableServerSignerList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerSignerList(val *ServerSignerList) *NullableServerSignerList {
	return &NullableServerSignerList{value: val, isSet: true}
}

func (v NullableServerSignerList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerSignerList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


