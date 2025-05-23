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

// WebhookStatus The status of the webhook.
type WebhookStatus string

// List of WebhookStatus
const (
	WEBHOOKSTATUS_ACTIVE WebhookStatus = "active"
	WEBHOOKSTATUS_INACTIVE WebhookStatus = "inactive"
)

// All allowed values of WebhookStatus enum
var AllowedWebhookStatusEnumValues = []WebhookStatus{
	"active",
	"inactive",
}

func (v *WebhookStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookStatus(value)
	for _, existing := range AllowedWebhookStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookStatus", value)
}

// NewWebhookStatusFromValue returns a pointer to a valid WebhookStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookStatusFromValue(v string) (*WebhookStatus, error) {
	ev := WebhookStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookStatus: valid values are %v", v, AllowedWebhookStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookStatus) IsValid() bool {
	for _, existing := range AllowedWebhookStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookStatus value
func (v WebhookStatus) Ptr() *WebhookStatus {
	return &v
}

type NullableWebhookStatus struct {
	value *WebhookStatus
	isSet bool
}

func (v NullableWebhookStatus) Get() *WebhookStatus {
	return v.value
}

func (v *NullableWebhookStatus) Set(val *WebhookStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookStatus(val *WebhookStatus) *NullableWebhookStatus {
	return &NullableWebhookStatus{value: val, isSet: true}
}

func (v NullableWebhookStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

