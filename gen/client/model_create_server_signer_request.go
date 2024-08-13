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

// checks if the CreateServerSignerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServerSignerRequest{}

// CreateServerSignerRequest struct for CreateServerSignerRequest
type CreateServerSignerRequest struct {
	// The ID of the server signer for the 1 of 1 server signer.
	ServerSignerId *string `json:"server_signer_id,omitempty"`
	// The enrollment data of the server signer. This will be the base64 encoded server-signer-id for the 1 of 1 server signer.
	EnrollmentData string `json:"enrollment_data"`
	// Whether the Server-Signer uses MPC.
	IsMpc bool `json:"is_mpc"`
}

type _CreateServerSignerRequest CreateServerSignerRequest

// NewCreateServerSignerRequest instantiates a new CreateServerSignerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServerSignerRequest(enrollmentData string, isMpc bool) *CreateServerSignerRequest {
	this := CreateServerSignerRequest{}
	this.EnrollmentData = enrollmentData
	this.IsMpc = isMpc
	return &this
}

// NewCreateServerSignerRequestWithDefaults instantiates a new CreateServerSignerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServerSignerRequestWithDefaults() *CreateServerSignerRequest {
	this := CreateServerSignerRequest{}
	return &this
}

// GetServerSignerId returns the ServerSignerId field value if set, zero value otherwise.
func (o *CreateServerSignerRequest) GetServerSignerId() string {
	if o == nil || IsNil(o.ServerSignerId) {
		var ret string
		return ret
	}
	return *o.ServerSignerId
}

// GetServerSignerIdOk returns a tuple with the ServerSignerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerSignerRequest) GetServerSignerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerSignerId) {
		return nil, false
	}
	return o.ServerSignerId, true
}

// HasServerSignerId returns a boolean if a field has been set.
func (o *CreateServerSignerRequest) HasServerSignerId() bool {
	if o != nil && !IsNil(o.ServerSignerId) {
		return true
	}

	return false
}

// SetServerSignerId gets a reference to the given string and assigns it to the ServerSignerId field.
func (o *CreateServerSignerRequest) SetServerSignerId(v string) {
	o.ServerSignerId = &v
}

// GetEnrollmentData returns the EnrollmentData field value
func (o *CreateServerSignerRequest) GetEnrollmentData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnrollmentData
}

// GetEnrollmentDataOk returns a tuple with the EnrollmentData field value
// and a boolean to check if the value has been set.
func (o *CreateServerSignerRequest) GetEnrollmentDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentData, true
}

// SetEnrollmentData sets field value
func (o *CreateServerSignerRequest) SetEnrollmentData(v string) {
	o.EnrollmentData = v
}

// GetIsMpc returns the IsMpc field value
func (o *CreateServerSignerRequest) GetIsMpc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMpc
}

// GetIsMpcOk returns a tuple with the IsMpc field value
// and a boolean to check if the value has been set.
func (o *CreateServerSignerRequest) GetIsMpcOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMpc, true
}

// SetIsMpc sets field value
func (o *CreateServerSignerRequest) SetIsMpc(v bool) {
	o.IsMpc = v
}

func (o CreateServerSignerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServerSignerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerSignerId) {
		toSerialize["server_signer_id"] = o.ServerSignerId
	}
	toSerialize["enrollment_data"] = o.EnrollmentData
	toSerialize["is_mpc"] = o.IsMpc
	return toSerialize, nil
}

func (o *CreateServerSignerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enrollment_data",
		"is_mpc",
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

	varCreateServerSignerRequest := _CreateServerSignerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateServerSignerRequest)

	if err != nil {
		return err
	}

	*o = CreateServerSignerRequest(varCreateServerSignerRequest)

	return err
}

type NullableCreateServerSignerRequest struct {
	value *CreateServerSignerRequest
	isSet bool
}

func (v NullableCreateServerSignerRequest) Get() *CreateServerSignerRequest {
	return v.value
}

func (v *NullableCreateServerSignerRequest) Set(val *CreateServerSignerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServerSignerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServerSignerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServerSignerRequest(val *CreateServerSignerRequest) *NullableCreateServerSignerRequest {
	return &NullableCreateServerSignerRequest{value: val, isSet: true}
}

func (v NullableCreateServerSignerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServerSignerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

