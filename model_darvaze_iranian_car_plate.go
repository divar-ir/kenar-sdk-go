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

// checks if the DarvazeIranianCarPlate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DarvazeIranianCarPlate{}

// DarvazeIranianCarPlate struct for DarvazeIranianCarPlate
type DarvazeIranianCarPlate struct {
	AreaCode *int64 `json:"area_code,omitempty"`
	ProvinceCode *int64 `json:"province_code,omitempty"`
	SerialLetter *string `json:"serial_letter,omitempty"`
	SerialNumber *int64 `json:"serial_number,omitempty"`
}

// NewDarvazeIranianCarPlate instantiates a new DarvazeIranianCarPlate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDarvazeIranianCarPlate() *DarvazeIranianCarPlate {
	this := DarvazeIranianCarPlate{}
	return &this
}

// NewDarvazeIranianCarPlateWithDefaults instantiates a new DarvazeIranianCarPlate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDarvazeIranianCarPlateWithDefaults() *DarvazeIranianCarPlate {
	this := DarvazeIranianCarPlate{}
	return &this
}

// GetAreaCode returns the AreaCode field value if set, zero value otherwise.
func (o *DarvazeIranianCarPlate) GetAreaCode() int64 {
	if o == nil || IsNil(o.AreaCode) {
		var ret int64
		return ret
	}
	return *o.AreaCode
}

// GetAreaCodeOk returns a tuple with the AreaCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DarvazeIranianCarPlate) GetAreaCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.AreaCode) {
		return nil, false
	}
	return o.AreaCode, true
}

// HasAreaCode returns a boolean if a field has been set.
func (o *DarvazeIranianCarPlate) HasAreaCode() bool {
	if o != nil && !IsNil(o.AreaCode) {
		return true
	}

	return false
}

// SetAreaCode gets a reference to the given int64 and assigns it to the AreaCode field.
func (o *DarvazeIranianCarPlate) SetAreaCode(v int64) {
	o.AreaCode = &v
}

// GetProvinceCode returns the ProvinceCode field value if set, zero value otherwise.
func (o *DarvazeIranianCarPlate) GetProvinceCode() int64 {
	if o == nil || IsNil(o.ProvinceCode) {
		var ret int64
		return ret
	}
	return *o.ProvinceCode
}

// GetProvinceCodeOk returns a tuple with the ProvinceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DarvazeIranianCarPlate) GetProvinceCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.ProvinceCode) {
		return nil, false
	}
	return o.ProvinceCode, true
}

// HasProvinceCode returns a boolean if a field has been set.
func (o *DarvazeIranianCarPlate) HasProvinceCode() bool {
	if o != nil && !IsNil(o.ProvinceCode) {
		return true
	}

	return false
}

// SetProvinceCode gets a reference to the given int64 and assigns it to the ProvinceCode field.
func (o *DarvazeIranianCarPlate) SetProvinceCode(v int64) {
	o.ProvinceCode = &v
}

// GetSerialLetter returns the SerialLetter field value if set, zero value otherwise.
func (o *DarvazeIranianCarPlate) GetSerialLetter() string {
	if o == nil || IsNil(o.SerialLetter) {
		var ret string
		return ret
	}
	return *o.SerialLetter
}

// GetSerialLetterOk returns a tuple with the SerialLetter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DarvazeIranianCarPlate) GetSerialLetterOk() (*string, bool) {
	if o == nil || IsNil(o.SerialLetter) {
		return nil, false
	}
	return o.SerialLetter, true
}

// HasSerialLetter returns a boolean if a field has been set.
func (o *DarvazeIranianCarPlate) HasSerialLetter() bool {
	if o != nil && !IsNil(o.SerialLetter) {
		return true
	}

	return false
}

// SetSerialLetter gets a reference to the given string and assigns it to the SerialLetter field.
func (o *DarvazeIranianCarPlate) SetSerialLetter(v string) {
	o.SerialLetter = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *DarvazeIranianCarPlate) GetSerialNumber() int64 {
	if o == nil || IsNil(o.SerialNumber) {
		var ret int64
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DarvazeIranianCarPlate) GetSerialNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *DarvazeIranianCarPlate) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given int64 and assigns it to the SerialNumber field.
func (o *DarvazeIranianCarPlate) SetSerialNumber(v int64) {
	o.SerialNumber = &v
}

func (o DarvazeIranianCarPlate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DarvazeIranianCarPlate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AreaCode) {
		toSerialize["area_code"] = o.AreaCode
	}
	if !IsNil(o.ProvinceCode) {
		toSerialize["province_code"] = o.ProvinceCode
	}
	if !IsNil(o.SerialLetter) {
		toSerialize["serial_letter"] = o.SerialLetter
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serial_number"] = o.SerialNumber
	}
	return toSerialize, nil
}

type NullableDarvazeIranianCarPlate struct {
	value *DarvazeIranianCarPlate
	isSet bool
}

func (v NullableDarvazeIranianCarPlate) Get() *DarvazeIranianCarPlate {
	return v.value
}

func (v *NullableDarvazeIranianCarPlate) Set(val *DarvazeIranianCarPlate) {
	v.value = val
	v.isSet = true
}

func (v NullableDarvazeIranianCarPlate) IsSet() bool {
	return v.isSet
}

func (v *NullableDarvazeIranianCarPlate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDarvazeIranianCarPlate(val *DarvazeIranianCarPlate) *NullableDarvazeIranianCarPlate {
	return &NullableDarvazeIranianCarPlate{value: val, isSet: true}
}

func (v NullableDarvazeIranianCarPlate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDarvazeIranianCarPlate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


