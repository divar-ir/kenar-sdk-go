# SearchPostItemRealEstateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credit** | Pointer to [**SearchPostItemPrice**](SearchPostItemPrice.md) |  | [optional] 
**Rent** | Pointer to [**SearchPostItemPrice**](SearchPostItemPrice.md) |  | [optional] 
**DailyRent** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Year** | Pointer to **int64** |  | [optional] 
**HasParking** | Pointer to **bool** |  | [optional] 
**HasElevator** | Pointer to **bool** |  | [optional] 
**Rooms** | Pointer to **string** |  | [optional] 
**Floor** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchPostItemRealEstateFields

`func NewSearchPostItemRealEstateFields() *SearchPostItemRealEstateFields`

NewSearchPostItemRealEstateFields instantiates a new SearchPostItemRealEstateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPostItemRealEstateFieldsWithDefaults

`func NewSearchPostItemRealEstateFieldsWithDefaults() *SearchPostItemRealEstateFields`

NewSearchPostItemRealEstateFieldsWithDefaults instantiates a new SearchPostItemRealEstateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredit

`func (o *SearchPostItemRealEstateFields) GetCredit() SearchPostItemPrice`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *SearchPostItemRealEstateFields) GetCreditOk() (*SearchPostItemPrice, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *SearchPostItemRealEstateFields) SetCredit(v SearchPostItemPrice)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *SearchPostItemRealEstateFields) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetRent

`func (o *SearchPostItemRealEstateFields) GetRent() SearchPostItemPrice`

GetRent returns the Rent field if non-nil, zero value otherwise.

### GetRentOk

`func (o *SearchPostItemRealEstateFields) GetRentOk() (*SearchPostItemPrice, bool)`

GetRentOk returns a tuple with the Rent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRent

`func (o *SearchPostItemRealEstateFields) SetRent(v SearchPostItemPrice)`

SetRent sets Rent field to given value.

### HasRent

`func (o *SearchPostItemRealEstateFields) HasRent() bool`

HasRent returns a boolean if a field has been set.

### GetDailyRent

`func (o *SearchPostItemRealEstateFields) GetDailyRent() string`

GetDailyRent returns the DailyRent field if non-nil, zero value otherwise.

### GetDailyRentOk

`func (o *SearchPostItemRealEstateFields) GetDailyRentOk() (*string, bool)`

GetDailyRentOk returns a tuple with the DailyRent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyRent

`func (o *SearchPostItemRealEstateFields) SetDailyRent(v string)`

SetDailyRent sets DailyRent field to given value.

### HasDailyRent

`func (o *SearchPostItemRealEstateFields) HasDailyRent() bool`

HasDailyRent returns a boolean if a field has been set.

### GetSize

`func (o *SearchPostItemRealEstateFields) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SearchPostItemRealEstateFields) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SearchPostItemRealEstateFields) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SearchPostItemRealEstateFields) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetYear

`func (o *SearchPostItemRealEstateFields) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SearchPostItemRealEstateFields) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SearchPostItemRealEstateFields) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *SearchPostItemRealEstateFields) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetHasParking

`func (o *SearchPostItemRealEstateFields) GetHasParking() bool`

GetHasParking returns the HasParking field if non-nil, zero value otherwise.

### GetHasParkingOk

`func (o *SearchPostItemRealEstateFields) GetHasParkingOk() (*bool, bool)`

GetHasParkingOk returns a tuple with the HasParking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParking

`func (o *SearchPostItemRealEstateFields) SetHasParking(v bool)`

SetHasParking sets HasParking field to given value.

### HasHasParking

`func (o *SearchPostItemRealEstateFields) HasHasParking() bool`

HasHasParking returns a boolean if a field has been set.

### GetHasElevator

`func (o *SearchPostItemRealEstateFields) GetHasElevator() bool`

GetHasElevator returns the HasElevator field if non-nil, zero value otherwise.

### GetHasElevatorOk

`func (o *SearchPostItemRealEstateFields) GetHasElevatorOk() (*bool, bool)`

GetHasElevatorOk returns a tuple with the HasElevator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasElevator

`func (o *SearchPostItemRealEstateFields) SetHasElevator(v bool)`

SetHasElevator sets HasElevator field to given value.

### HasHasElevator

`func (o *SearchPostItemRealEstateFields) HasHasElevator() bool`

HasHasElevator returns a boolean if a field has been set.

### GetRooms

`func (o *SearchPostItemRealEstateFields) GetRooms() string`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *SearchPostItemRealEstateFields) GetRoomsOk() (*string, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *SearchPostItemRealEstateFields) SetRooms(v string)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *SearchPostItemRealEstateFields) HasRooms() bool`

HasRooms returns a boolean if a field has been set.

### GetFloor

`func (o *SearchPostItemRealEstateFields) GetFloor() int32`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *SearchPostItemRealEstateFields) GetFloorOk() (*int32, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *SearchPostItemRealEstateFields) SetFloor(v int32)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *SearchPostItemRealEstateFields) HasFloor() bool`

HasFloor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


