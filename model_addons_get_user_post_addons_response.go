/*
Kenar API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kenarapi

import (
	"encoding/json"
)

// checks if the AddonsGetUserPostAddonsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonsGetUserPostAddonsResponse{}

// AddonsGetUserPostAddonsResponse struct for AddonsGetUserPostAddonsResponse
type AddonsGetUserPostAddonsResponse struct {
	Addons []AddonsPostAddon `json:"addons,omitempty"`
}

// NewAddonsGetUserPostAddonsResponse instantiates a new AddonsGetUserPostAddonsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonsGetUserPostAddonsResponse() *AddonsGetUserPostAddonsResponse {
	this := AddonsGetUserPostAddonsResponse{}
	return &this
}

// NewAddonsGetUserPostAddonsResponseWithDefaults instantiates a new AddonsGetUserPostAddonsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonsGetUserPostAddonsResponseWithDefaults() *AddonsGetUserPostAddonsResponse {
	this := AddonsGetUserPostAddonsResponse{}
	return &this
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *AddonsGetUserPostAddonsResponse) GetAddons() []AddonsPostAddon {
	if o == nil || IsNil(o.Addons) {
		var ret []AddonsPostAddon
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsGetUserPostAddonsResponse) GetAddonsOk() ([]AddonsPostAddon, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *AddonsGetUserPostAddonsResponse) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []AddonsPostAddon and assigns it to the Addons field.
func (o *AddonsGetUserPostAddonsResponse) SetAddons(v []AddonsPostAddon) {
	o.Addons = v
}

func (o AddonsGetUserPostAddonsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonsGetUserPostAddonsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	return toSerialize, nil
}

type NullableAddonsGetUserPostAddonsResponse struct {
	value *AddonsGetUserPostAddonsResponse
	isSet bool
}

func (v NullableAddonsGetUserPostAddonsResponse) Get() *AddonsGetUserPostAddonsResponse {
	return v.value
}

func (v *NullableAddonsGetUserPostAddonsResponse) Set(val *AddonsGetUserPostAddonsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonsGetUserPostAddonsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonsGetUserPostAddonsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonsGetUserPostAddonsResponse(val *AddonsGetUserPostAddonsResponse) *NullableAddonsGetUserPostAddonsResponse {
	return &NullableAddonsGetUserPostAddonsResponse{value: val, isSet: true}
}

func (v NullableAddonsGetUserPostAddonsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonsGetUserPostAddonsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


