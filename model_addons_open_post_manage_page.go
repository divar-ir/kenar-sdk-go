/*
Kenar API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kenarapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AddonsOpenPostManagePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonsOpenPostManagePage{}

// AddonsOpenPostManagePage An action to open a post management page in the app
type AddonsOpenPostManagePage struct {
	// Token of the post to redirect to its management page
	PostToken string `json:"post_token"`
}

type _AddonsOpenPostManagePage AddonsOpenPostManagePage

// NewAddonsOpenPostManagePage instantiates a new AddonsOpenPostManagePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonsOpenPostManagePage(postToken string) *AddonsOpenPostManagePage {
	this := AddonsOpenPostManagePage{}
	this.PostToken = postToken
	return &this
}

// NewAddonsOpenPostManagePageWithDefaults instantiates a new AddonsOpenPostManagePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonsOpenPostManagePageWithDefaults() *AddonsOpenPostManagePage {
	this := AddonsOpenPostManagePage{}
	return &this
}

// GetPostToken returns the PostToken field value
func (o *AddonsOpenPostManagePage) GetPostToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostToken
}

// GetPostTokenOk returns a tuple with the PostToken field value
// and a boolean to check if the value has been set.
func (o *AddonsOpenPostManagePage) GetPostTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostToken, true
}

// SetPostToken sets field value
func (o *AddonsOpenPostManagePage) SetPostToken(v string) {
	o.PostToken = v
}

func (o AddonsOpenPostManagePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonsOpenPostManagePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["post_token"] = o.PostToken
	return toSerialize, nil
}

func (o *AddonsOpenPostManagePage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"post_token",
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

	varAddonsOpenPostManagePage := _AddonsOpenPostManagePage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddonsOpenPostManagePage)

	if err != nil {
		return err
	}

	*o = AddonsOpenPostManagePage(varAddonsOpenPostManagePage)

	return err
}

type NullableAddonsOpenPostManagePage struct {
	value *AddonsOpenPostManagePage
	isSet bool
}

func (v NullableAddonsOpenPostManagePage) Get() *AddonsOpenPostManagePage {
	return v.value
}

func (v *NullableAddonsOpenPostManagePage) Set(val *AddonsOpenPostManagePage) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonsOpenPostManagePage) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonsOpenPostManagePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonsOpenPostManagePage(val *AddonsOpenPostManagePage) *NullableAddonsOpenPostManagePage {
	return &NullableAddonsOpenPostManagePage{value: val, isSet: true}
}

func (v NullableAddonsOpenPostManagePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonsOpenPostManagePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


