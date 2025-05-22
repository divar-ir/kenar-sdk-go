# FinderSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandModel** | Pointer to **[]string** |  | [optional] 
**Credit** | Pointer to [**FinderSearchQueryNumberRange**](FinderSearchQueryNumberRange.md) |  | [optional] 
**OnlyWithParking** | Pointer to **bool** |  | [optional] 
**ProductionYear** | Pointer to [**FinderSearchQueryNumberRange**](FinderSearchQueryNumberRange.md) |  | [optional] 
**Rent** | Pointer to [**FinderSearchQueryNumberRange**](FinderSearchQueryNumberRange.md) |  | [optional] 
**Rooms** | Pointer to **[]string** |  | [optional] 
**Size** | Pointer to [**FinderSearchQueryNumberRange**](FinderSearchQueryNumberRange.md) |  | [optional] 
**Usage** | Pointer to [**FinderSearchQueryNumberRange**](FinderSearchQueryNumberRange.md) |  | [optional] 

## Methods

### NewFinderSearchQuery

`func NewFinderSearchQuery() *FinderSearchQuery`

NewFinderSearchQuery instantiates a new FinderSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinderSearchQueryWithDefaults

`func NewFinderSearchQueryWithDefaults() *FinderSearchQuery`

NewFinderSearchQueryWithDefaults instantiates a new FinderSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandModel

`func (o *FinderSearchQuery) GetBrandModel() []string`

GetBrandModel returns the BrandModel field if non-nil, zero value otherwise.

### GetBrandModelOk

`func (o *FinderSearchQuery) GetBrandModelOk() (*[]string, bool)`

GetBrandModelOk returns a tuple with the BrandModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandModel

`func (o *FinderSearchQuery) SetBrandModel(v []string)`

SetBrandModel sets BrandModel field to given value.

### HasBrandModel

`func (o *FinderSearchQuery) HasBrandModel() bool`

HasBrandModel returns a boolean if a field has been set.

### GetCredit

`func (o *FinderSearchQuery) GetCredit() FinderSearchQueryNumberRange`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *FinderSearchQuery) GetCreditOk() (*FinderSearchQueryNumberRange, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *FinderSearchQuery) SetCredit(v FinderSearchQueryNumberRange)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *FinderSearchQuery) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetOnlyWithParking

`func (o *FinderSearchQuery) GetOnlyWithParking() bool`

GetOnlyWithParking returns the OnlyWithParking field if non-nil, zero value otherwise.

### GetOnlyWithParkingOk

`func (o *FinderSearchQuery) GetOnlyWithParkingOk() (*bool, bool)`

GetOnlyWithParkingOk returns a tuple with the OnlyWithParking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithParking

`func (o *FinderSearchQuery) SetOnlyWithParking(v bool)`

SetOnlyWithParking sets OnlyWithParking field to given value.

### HasOnlyWithParking

`func (o *FinderSearchQuery) HasOnlyWithParking() bool`

HasOnlyWithParking returns a boolean if a field has been set.

### GetProductionYear

`func (o *FinderSearchQuery) GetProductionYear() FinderSearchQueryNumberRange`

GetProductionYear returns the ProductionYear field if non-nil, zero value otherwise.

### GetProductionYearOk

`func (o *FinderSearchQuery) GetProductionYearOk() (*FinderSearchQueryNumberRange, bool)`

GetProductionYearOk returns a tuple with the ProductionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionYear

`func (o *FinderSearchQuery) SetProductionYear(v FinderSearchQueryNumberRange)`

SetProductionYear sets ProductionYear field to given value.

### HasProductionYear

`func (o *FinderSearchQuery) HasProductionYear() bool`

HasProductionYear returns a boolean if a field has been set.

### GetRent

`func (o *FinderSearchQuery) GetRent() FinderSearchQueryNumberRange`

GetRent returns the Rent field if non-nil, zero value otherwise.

### GetRentOk

`func (o *FinderSearchQuery) GetRentOk() (*FinderSearchQueryNumberRange, bool)`

GetRentOk returns a tuple with the Rent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRent

`func (o *FinderSearchQuery) SetRent(v FinderSearchQueryNumberRange)`

SetRent sets Rent field to given value.

### HasRent

`func (o *FinderSearchQuery) HasRent() bool`

HasRent returns a boolean if a field has been set.

### GetRooms

`func (o *FinderSearchQuery) GetRooms() []string`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *FinderSearchQuery) GetRoomsOk() (*[]string, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *FinderSearchQuery) SetRooms(v []string)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *FinderSearchQuery) HasRooms() bool`

HasRooms returns a boolean if a field has been set.

### GetSize

`func (o *FinderSearchQuery) GetSize() FinderSearchQueryNumberRange`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FinderSearchQuery) GetSizeOk() (*FinderSearchQueryNumberRange, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FinderSearchQuery) SetSize(v FinderSearchQueryNumberRange)`

SetSize sets Size field to given value.

### HasSize

`func (o *FinderSearchQuery) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUsage

`func (o *FinderSearchQuery) GetUsage() FinderSearchQueryNumberRange`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *FinderSearchQuery) GetUsageOk() (*FinderSearchQueryNumberRange, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *FinderSearchQuery) SetUsage(v FinderSearchQueryNumberRange)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *FinderSearchQuery) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


