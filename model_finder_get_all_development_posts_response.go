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

// checks if the FinderGetAllDevelopmentPostsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinderGetAllDevelopmentPostsResponse{}

// FinderGetAllDevelopmentPostsResponse struct for FinderGetAllDevelopmentPostsResponse
type FinderGetAllDevelopmentPostsResponse struct {
	DevelopmentPosts []ManagementDevelopmentPost `json:"development_posts,omitempty"`
}

// NewFinderGetAllDevelopmentPostsResponse instantiates a new FinderGetAllDevelopmentPostsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinderGetAllDevelopmentPostsResponse() *FinderGetAllDevelopmentPostsResponse {
	this := FinderGetAllDevelopmentPostsResponse{}
	return &this
}

// NewFinderGetAllDevelopmentPostsResponseWithDefaults instantiates a new FinderGetAllDevelopmentPostsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinderGetAllDevelopmentPostsResponseWithDefaults() *FinderGetAllDevelopmentPostsResponse {
	this := FinderGetAllDevelopmentPostsResponse{}
	return &this
}

// GetDevelopmentPosts returns the DevelopmentPosts field value if set, zero value otherwise.
func (o *FinderGetAllDevelopmentPostsResponse) GetDevelopmentPosts() []ManagementDevelopmentPost {
	if o == nil || IsNil(o.DevelopmentPosts) {
		var ret []ManagementDevelopmentPost
		return ret
	}
	return o.DevelopmentPosts
}

// GetDevelopmentPostsOk returns a tuple with the DevelopmentPosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinderGetAllDevelopmentPostsResponse) GetDevelopmentPostsOk() ([]ManagementDevelopmentPost, bool) {
	if o == nil || IsNil(o.DevelopmentPosts) {
		return nil, false
	}
	return o.DevelopmentPosts, true
}

// HasDevelopmentPosts returns a boolean if a field has been set.
func (o *FinderGetAllDevelopmentPostsResponse) HasDevelopmentPosts() bool {
	if o != nil && !IsNil(o.DevelopmentPosts) {
		return true
	}

	return false
}

// SetDevelopmentPosts gets a reference to the given []ManagementDevelopmentPost and assigns it to the DevelopmentPosts field.
func (o *FinderGetAllDevelopmentPostsResponse) SetDevelopmentPosts(v []ManagementDevelopmentPost) {
	o.DevelopmentPosts = v
}

func (o FinderGetAllDevelopmentPostsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinderGetAllDevelopmentPostsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DevelopmentPosts) {
		toSerialize["development_posts"] = o.DevelopmentPosts
	}
	return toSerialize, nil
}

type NullableFinderGetAllDevelopmentPostsResponse struct {
	value *FinderGetAllDevelopmentPostsResponse
	isSet bool
}

func (v NullableFinderGetAllDevelopmentPostsResponse) Get() *FinderGetAllDevelopmentPostsResponse {
	return v.value
}

func (v *NullableFinderGetAllDevelopmentPostsResponse) Set(val *FinderGetAllDevelopmentPostsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFinderGetAllDevelopmentPostsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFinderGetAllDevelopmentPostsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinderGetAllDevelopmentPostsResponse(val *FinderGetAllDevelopmentPostsResponse) *NullableFinderGetAllDevelopmentPostsResponse {
	return &NullableFinderGetAllDevelopmentPostsResponse{value: val, isSet: true}
}

func (v NullableFinderGetAllDevelopmentPostsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinderGetAllDevelopmentPostsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


