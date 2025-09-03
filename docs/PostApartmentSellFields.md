# PostApartmentSellFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Floor** | **int32** | Floor of the property. Use -1 for under ground floor and 0 for ground floor. Use 1 for first floor and so on. | 
**HasElevator** | **bool** | Whether the property has an elevator | 
**HasOwnImage** | **bool** | تصاویر مربوط به خود ملک بوده و تزئینی نیستند. | 
**HasParking** | **bool** | Whether the property has a parking | 
**HasWarehouse** | **bool** | Whether the property has a warehouse | 
**Price** | **string** | Price of the property in Toman | 
**RoomsCount** | [**PostRoomsCount**](PostRoomsCount.md) |  | 
**Size** | **int32** | Size of the property in square meters | 
**YearBuilt** | **int32** | Year the property was built (Persian/Shamsi calendar) | 

## Methods

### NewPostApartmentSellFields

`func NewPostApartmentSellFields(floor int32, hasElevator bool, hasOwnImage bool, hasParking bool, hasWarehouse bool, price string, roomsCount PostRoomsCount, size int32, yearBuilt int32, ) *PostApartmentSellFields`

NewPostApartmentSellFields instantiates a new PostApartmentSellFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApartmentSellFieldsWithDefaults

`func NewPostApartmentSellFieldsWithDefaults() *PostApartmentSellFields`

NewPostApartmentSellFieldsWithDefaults instantiates a new PostApartmentSellFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloor

`func (o *PostApartmentSellFields) GetFloor() int32`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *PostApartmentSellFields) GetFloorOk() (*int32, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *PostApartmentSellFields) SetFloor(v int32)`

SetFloor sets Floor field to given value.


### GetHasElevator

`func (o *PostApartmentSellFields) GetHasElevator() bool`

GetHasElevator returns the HasElevator field if non-nil, zero value otherwise.

### GetHasElevatorOk

`func (o *PostApartmentSellFields) GetHasElevatorOk() (*bool, bool)`

GetHasElevatorOk returns a tuple with the HasElevator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasElevator

`func (o *PostApartmentSellFields) SetHasElevator(v bool)`

SetHasElevator sets HasElevator field to given value.


### GetHasOwnImage

`func (o *PostApartmentSellFields) GetHasOwnImage() bool`

GetHasOwnImage returns the HasOwnImage field if non-nil, zero value otherwise.

### GetHasOwnImageOk

`func (o *PostApartmentSellFields) GetHasOwnImageOk() (*bool, bool)`

GetHasOwnImageOk returns a tuple with the HasOwnImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOwnImage

`func (o *PostApartmentSellFields) SetHasOwnImage(v bool)`

SetHasOwnImage sets HasOwnImage field to given value.


### GetHasParking

`func (o *PostApartmentSellFields) GetHasParking() bool`

GetHasParking returns the HasParking field if non-nil, zero value otherwise.

### GetHasParkingOk

`func (o *PostApartmentSellFields) GetHasParkingOk() (*bool, bool)`

GetHasParkingOk returns a tuple with the HasParking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParking

`func (o *PostApartmentSellFields) SetHasParking(v bool)`

SetHasParking sets HasParking field to given value.


### GetHasWarehouse

`func (o *PostApartmentSellFields) GetHasWarehouse() bool`

GetHasWarehouse returns the HasWarehouse field if non-nil, zero value otherwise.

### GetHasWarehouseOk

`func (o *PostApartmentSellFields) GetHasWarehouseOk() (*bool, bool)`

GetHasWarehouseOk returns a tuple with the HasWarehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWarehouse

`func (o *PostApartmentSellFields) SetHasWarehouse(v bool)`

SetHasWarehouse sets HasWarehouse field to given value.


### GetPrice

`func (o *PostApartmentSellFields) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PostApartmentSellFields) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PostApartmentSellFields) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetRoomsCount

`func (o *PostApartmentSellFields) GetRoomsCount() PostRoomsCount`

GetRoomsCount returns the RoomsCount field if non-nil, zero value otherwise.

### GetRoomsCountOk

`func (o *PostApartmentSellFields) GetRoomsCountOk() (*PostRoomsCount, bool)`

GetRoomsCountOk returns a tuple with the RoomsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsCount

`func (o *PostApartmentSellFields) SetRoomsCount(v PostRoomsCount)`

SetRoomsCount sets RoomsCount field to given value.


### GetSize

`func (o *PostApartmentSellFields) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostApartmentSellFields) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostApartmentSellFields) SetSize(v int32)`

SetSize sets Size field to given value.


### GetYearBuilt

`func (o *PostApartmentSellFields) GetYearBuilt() int32`

GetYearBuilt returns the YearBuilt field if non-nil, zero value otherwise.

### GetYearBuiltOk

`func (o *PostApartmentSellFields) GetYearBuiltOk() (*int32, bool)`

GetYearBuiltOk returns a tuple with the YearBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearBuilt

`func (o *PostApartmentSellFields) SetYearBuilt(v int32)`

SetYearBuilt sets YearBuilt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


