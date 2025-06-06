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

// checks if the AddonsGroupInfoRow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonsGroupInfoRow{}

// AddonsGroupInfoRow struct for AddonsGroupInfoRow
type AddonsGroupInfoRow struct {
	HasDivider *bool `json:"has_divider,omitempty"`
	Items []AddonsGroupInfoRowGroupInfoItem `json:"items,omitempty"`
}

// NewAddonsGroupInfoRow instantiates a new AddonsGroupInfoRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonsGroupInfoRow() *AddonsGroupInfoRow {
	this := AddonsGroupInfoRow{}
	return &this
}

// NewAddonsGroupInfoRowWithDefaults instantiates a new AddonsGroupInfoRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonsGroupInfoRowWithDefaults() *AddonsGroupInfoRow {
	this := AddonsGroupInfoRow{}
	return &this
}

// GetHasDivider returns the HasDivider field value if set, zero value otherwise.
func (o *AddonsGroupInfoRow) GetHasDivider() bool {
	if o == nil || IsNil(o.HasDivider) {
		var ret bool
		return ret
	}
	return *o.HasDivider
}

// GetHasDividerOk returns a tuple with the HasDivider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsGroupInfoRow) GetHasDividerOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDivider) {
		return nil, false
	}
	return o.HasDivider, true
}

// HasHasDivider returns a boolean if a field has been set.
func (o *AddonsGroupInfoRow) HasHasDivider() bool {
	if o != nil && !IsNil(o.HasDivider) {
		return true
	}

	return false
}

// SetHasDivider gets a reference to the given bool and assigns it to the HasDivider field.
func (o *AddonsGroupInfoRow) SetHasDivider(v bool) {
	o.HasDivider = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AddonsGroupInfoRow) GetItems() []AddonsGroupInfoRowGroupInfoItem {
	if o == nil || IsNil(o.Items) {
		var ret []AddonsGroupInfoRowGroupInfoItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsGroupInfoRow) GetItemsOk() ([]AddonsGroupInfoRowGroupInfoItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AddonsGroupInfoRow) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AddonsGroupInfoRowGroupInfoItem and assigns it to the Items field.
func (o *AddonsGroupInfoRow) SetItems(v []AddonsGroupInfoRowGroupInfoItem) {
	o.Items = v
}

func (o AddonsGroupInfoRow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonsGroupInfoRow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HasDivider) {
		toSerialize["has_divider"] = o.HasDivider
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableAddonsGroupInfoRow struct {
	value *AddonsGroupInfoRow
	isSet bool
}

func (v NullableAddonsGroupInfoRow) Get() *AddonsGroupInfoRow {
	return v.value
}

func (v *NullableAddonsGroupInfoRow) Set(val *AddonsGroupInfoRow) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonsGroupInfoRow) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonsGroupInfoRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonsGroupInfoRow(val *AddonsGroupInfoRow) *NullableAddonsGroupInfoRow {
	return &NullableAddonsGroupInfoRow{value: val, isSet: true}
}

func (v NullableAddonsGroupInfoRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonsGroupInfoRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


