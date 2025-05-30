/*
Kenar API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kenarapi

import (
	"encoding/json"
	"fmt"
)

// AddonsStatus the model 'AddonsStatus'
type AddonsStatus string

// List of addonsStatus
const (
	ADDONSSTATUS_ACTIVE AddonsStatus = "ACTIVE"
	ADDONSSTATUS_HIDDEN AddonsStatus = "HIDDEN"
	ADDONSSTATUS_DELETED AddonsStatus = "DELETED"
)

// All allowed values of AddonsStatus enum
var AllowedAddonsStatusEnumValues = []AddonsStatus{
	"ACTIVE",
	"HIDDEN",
	"DELETED",
}

func (v *AddonsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AddonsStatus(value)
	for _, existing := range AllowedAddonsStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AddonsStatus", value)
}

// NewAddonsStatusFromValue returns a pointer to a valid AddonsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAddonsStatusFromValue(v string) (*AddonsStatus, error) {
	ev := AddonsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AddonsStatus: valid values are %v", v, AllowedAddonsStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AddonsStatus) IsValid() bool {
	for _, existing := range AllowedAddonsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to addonsStatus value
func (v AddonsStatus) Ptr() *AddonsStatus {
	return &v
}

type NullableAddonsStatus struct {
	value *AddonsStatus
	isSet bool
}

func (v NullableAddonsStatus) Get() *AddonsStatus {
	return v.value
}

func (v *NullableAddonsStatus) Set(val *AddonsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonsStatus(val *AddonsStatus) *NullableAddonsStatus {
	return &NullableAddonsStatus{value: val, isSet: true}
}

func (v NullableAddonsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

