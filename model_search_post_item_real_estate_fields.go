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

// checks if the SearchPostItemRealEstateFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchPostItemRealEstateFields{}

// SearchPostItemRealEstateFields struct for SearchPostItemRealEstateFields
type SearchPostItemRealEstateFields struct {
	Credit *SearchPostItemPrice `json:"credit,omitempty"`
	DailyRent *string `json:"daily_rent,omitempty"`
	Floor *int32 `json:"floor,omitempty"`
	HasElevator *bool `json:"has_elevator,omitempty"`
	HasParking *bool `json:"has_parking,omitempty"`
	Rent *SearchPostItemPrice `json:"rent,omitempty"`
	Rooms *string `json:"rooms,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Year *int64 `json:"year,omitempty"`
}

// NewSearchPostItemRealEstateFields instantiates a new SearchPostItemRealEstateFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPostItemRealEstateFields() *SearchPostItemRealEstateFields {
	this := SearchPostItemRealEstateFields{}
	return &this
}

// NewSearchPostItemRealEstateFieldsWithDefaults instantiates a new SearchPostItemRealEstateFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPostItemRealEstateFieldsWithDefaults() *SearchPostItemRealEstateFields {
	this := SearchPostItemRealEstateFields{}
	return &this
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *SearchPostItemRealEstateFields) GetCredit() SearchPostItemPrice {
	if o == nil || IsNil(o.Credit) {
		var ret SearchPostItemPrice
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostItemRealEstateFields) GetCreditOk() (*SearchPostItemPrice, bool) {
	if o == nil || IsNil(o.Credit) {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *SearchPostItemRealEstateFields) HasCredit() bool {
	if o != nil && !IsNil(o.Credit) {
		return true
	}

	return false
}

// SetCredit gets a reference to the given SearchPostItemPrice and assigns it to the Credit field.
func (o *SearchPostItemRealEstateFields) SetCredit(v SearchPostItemPrice) {
	o.Credit = &v
}

// GetDailyRent returns the DailyRent field value if set, zero value otherwise.
func (o *SearchPostItemRealEstateFields) GetDailyRent() string {
	if o == nil || IsNil(o.DailyRent) {
		var ret string
		return ret
	}
	return *o.DailyRent
}

// GetDailyRentOk returns a tuple with the DailyRent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostItemRealEstateFields) GetDailyRentOk() (*string, bool) {
	if o == nil || IsNil(o.DailyRent) {
		return nil, false
	}
	return o.DailyRent, true
}

// HasDailyRent returns a boolean if a field has been set.
func (o *SearchPostItemRealEstateFields) HasDailyRent() bool {
	if o != nil && !IsNil(o.DailyRent) {
		return true
	}

	return false
}

// SetDailyRent gets a reference to the given string and assigns it to the DailyRent field.
func (o *SearchPostItemRealEstateFields) SetDailyRent(v string) {
	o.DailyRent = &v
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *SearchPostItemRealEstateFields) GetFloor() int32 {
	if o == nil || IsNil(o.Floor) {
		var ret int32
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostItemRealEstateFields) GetFloorOk() (*int32, bool) {
	if o == nil || IsNil(o.Floor) {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *SearchPostItemRealEstateFields) HasFloor() bool {
	if o != nil && !IsNil(o.Floor) {
		return true
	}

	return false
}

// SetFloor gets a reference to the given int32 and assigns it to the Floor field.
func (o *SearchPostItemRealEstateFields) SetFloor(v int32) {
	o.Floor = &v
}

// GetHasElevator returns the HasElevator field value if set, zero value otherwise.
func (o *SearchPostItemRealEstateFields) GetHasElevator() bool {
	if o == nil || IsNil(o.HasElevator) {
		var ret bool
		return ret
	}
	return *o.HasElevator
}

// GetHasElevatorOk returns a tuple with the HasElevator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostItemRealEstateFields) GetHasElevatorOk() (*bool, bool) {
	if o == nil || IsNil(o.HasElevator) {
		return nil, false
	}
	return o.HasElevator, true
}

// HasHasElevator returns a boolean if a field has been set.
func (o *SearchPostItemRealEstateFields) HasHasElevator() bool {
	if o != nil && !IsNil(o.HasElevator) {
		return true
	}

	return false
}

// SetHasElevator gets a reference to the given bool and assigns it to the HasElevator field.
func (o *SearchPostItemRealEstateFields) SetHasElevator(v bool) {
	o.HasElevator = &v
}

// GetHasParking returns the HasParking field value if set, zero value otherwise.
func (o *SearchPostItemRealEstateFields) GetHasParking() bool {
	if o == nil || IsNil(o.HasParking) {
		var ret bool
		return ret
	}
	return *o.HasParking
}

// GetHasParkingOk returns a tuple with the HasParking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostItemRealEstateFields) GetHasParkingOk() (*bool, bool) {
	if o == nil || IsNil(o.HasParking) {
		return nil, false
	}
	return o.HasParking, true
}

// HasHasParking returns a boolean if a field has been set.
func (o *SearchPostItemRealEstateFields) HasHasParking() bool {
	if o != nil && !IsNil(o.HasParking) {
		return true
	}

	return false
}

// SetHasParking gets a reference to the given bool and assigns it to the HasParking field.
func (o *SearchPostItemRealEstateFields) SetHasParking(v bool) {
	o.HasParking = &v
}

// GetRent returns the Rent field value if set, zero value otherwise.
func (o *SearchPostItemRealEstateFields) GetRent() SearchPostItemPrice {
	if o == nil || IsNil(o.Rent) {
		var ret SearchPostItemPrice
		return ret
	}
	return *o.Rent
}

// GetRentOk returns a tuple with the Rent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostItemRealEstateFields) GetRentOk() (*SearchPostItemPrice, bool) {
	if o == nil || IsNil(o.Rent) {
		return nil, false
	}
	return o.Rent, true
}

// HasRent returns a boolean if a field has been set.
func (o *SearchPostItemRealEstateFields) HasRent() bool {
	if o != nil && !IsNil(o.Rent) {
		return true
	}

	return false
}

// SetRent gets a reference to the given SearchPostItemPrice and assigns it to the Rent field.
func (o *SearchPostItemRealEstateFields) SetRent(v SearchPostItemPrice) {
	o.Rent = &v
}

// GetRooms returns the Rooms field value if set, zero value otherwise.
func (o *SearchPostItemRealEstateFields) GetRooms() string {
	if o == nil || IsNil(o.Rooms) {
		var ret string
		return ret
	}
	return *o.Rooms
}

// GetRoomsOk returns a tuple with the Rooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostItemRealEstateFields) GetRoomsOk() (*string, bool) {
	if o == nil || IsNil(o.Rooms) {
		return nil, false
	}
	return o.Rooms, true
}

// HasRooms returns a boolean if a field has been set.
func (o *SearchPostItemRealEstateFields) HasRooms() bool {
	if o != nil && !IsNil(o.Rooms) {
		return true
	}

	return false
}

// SetRooms gets a reference to the given string and assigns it to the Rooms field.
func (o *SearchPostItemRealEstateFields) SetRooms(v string) {
	o.Rooms = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SearchPostItemRealEstateFields) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostItemRealEstateFields) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SearchPostItemRealEstateFields) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *SearchPostItemRealEstateFields) SetSize(v int32) {
	o.Size = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *SearchPostItemRealEstateFields) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPostItemRealEstateFields) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *SearchPostItemRealEstateFields) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *SearchPostItemRealEstateFields) SetYear(v int64) {
	o.Year = &v
}

func (o SearchPostItemRealEstateFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchPostItemRealEstateFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credit) {
		toSerialize["credit"] = o.Credit
	}
	if !IsNil(o.DailyRent) {
		toSerialize["daily_rent"] = o.DailyRent
	}
	if !IsNil(o.Floor) {
		toSerialize["floor"] = o.Floor
	}
	if !IsNil(o.HasElevator) {
		toSerialize["has_elevator"] = o.HasElevator
	}
	if !IsNil(o.HasParking) {
		toSerialize["has_parking"] = o.HasParking
	}
	if !IsNil(o.Rent) {
		toSerialize["rent"] = o.Rent
	}
	if !IsNil(o.Rooms) {
		toSerialize["rooms"] = o.Rooms
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	return toSerialize, nil
}

type NullableSearchPostItemRealEstateFields struct {
	value *SearchPostItemRealEstateFields
	isSet bool
}

func (v NullableSearchPostItemRealEstateFields) Get() *SearchPostItemRealEstateFields {
	return v.value
}

func (v *NullableSearchPostItemRealEstateFields) Set(val *SearchPostItemRealEstateFields) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPostItemRealEstateFields) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPostItemRealEstateFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPostItemRealEstateFields(val *SearchPostItemRealEstateFields) *NullableSearchPostItemRealEstateFields {
	return &NullableSearchPostItemRealEstateFields{value: val, isSet: true}
}

func (v NullableSearchPostItemRealEstateFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPostItemRealEstateFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


