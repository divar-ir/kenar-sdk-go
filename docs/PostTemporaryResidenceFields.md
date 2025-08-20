# PostTemporaryResidenceFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **int32** | متراژ اقامتگاه به متر مربع | [optional] 
**CheckInTime** | Pointer to **string** | Check-in time | [optional] 
**CheckOutTime** | Pointer to **string** | Check-out time | [optional] 
**ComfortAmenities** | Pointer to [**[]TemporaryResidenceFieldsComfortAmenity**](TemporaryResidenceFieldsComfortAmenity.md) |  | [optional] 
**DamageDeposit** | Pointer to **string** | Damage deposit amount in Toman | [optional] 
**ExtraPersonCapacity** | Pointer to **int32** | تعداد افراد اضافه مجاز در اقامتگاه | [optional] 
**FullyFurnished** | Pointer to **bool** | Whether the residence is fully furnished | [optional] 
**HasOwnImage** | Pointer to **bool** | تصاویر مربوط به خود ملک بوده و تزئینی نیستند. | [optional] 
**HeatingCoolingSystem** | Pointer to [**[]TemporaryResidenceFieldsHeatingCoolingSystem**](TemporaryResidenceFieldsHeatingCoolingSystem.md) |  | [optional] 
**HouseRules** | Pointer to **string** | House rules and regulations | [optional] 
**MinimumStay** | Pointer to **int32** | Minimum number of days required for stay | [optional] 
**PetsAllowed** | Pointer to [**TemporaryResidenceFieldsPetsAllowed**](TemporaryResidenceFieldsPetsAllowed.md) |  | [optional] 
**PriceCostPerExtraPerson** | Pointer to **string** | هزینه هر نفر اضافه به ازای هر شب به تومان | [optional] 
**PriceRegularDays** | Pointer to **string** | قیمت اقامتگاه در روزهای عادی (شنبه تا سه‌شنبه) به تومان | [optional] 
**PriceSpecialDays** | Pointer to **string** | قیمت اقامتگاه در روزهای خاص (تعطیلات و مناسبت‌ها) به تومان | [optional] 
**PriceWeekends** | Pointer to **string** | قیمت اقامتگاه در آخر هفته (چهارشنبه تا جمعه) به تومان | [optional] 
**RegularPersonCapacity** | Pointer to **int32** | ظرفیت استاندارد افراد در اقامتگاه | [optional] 
**RentalPeriod** | Pointer to [**TemporaryResidenceFieldsRentalPeriod**](TemporaryResidenceFieldsRentalPeriod.md) |  | [optional] 
**RoomsCount** | Pointer to [**TemporaryResidenceFieldsRoomsCount**](TemporaryResidenceFieldsRoomsCount.md) |  | [optional] 

## Methods

### NewPostTemporaryResidenceFields

`func NewPostTemporaryResidenceFields() *PostTemporaryResidenceFields`

NewPostTemporaryResidenceFields instantiates a new PostTemporaryResidenceFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTemporaryResidenceFieldsWithDefaults

`func NewPostTemporaryResidenceFieldsWithDefaults() *PostTemporaryResidenceFields`

NewPostTemporaryResidenceFieldsWithDefaults instantiates a new PostTemporaryResidenceFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *PostTemporaryResidenceFields) GetArea() int32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *PostTemporaryResidenceFields) GetAreaOk() (*int32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *PostTemporaryResidenceFields) SetArea(v int32)`

SetArea sets Area field to given value.

### HasArea

`func (o *PostTemporaryResidenceFields) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCheckInTime

`func (o *PostTemporaryResidenceFields) GetCheckInTime() string`

GetCheckInTime returns the CheckInTime field if non-nil, zero value otherwise.

### GetCheckInTimeOk

`func (o *PostTemporaryResidenceFields) GetCheckInTimeOk() (*string, bool)`

GetCheckInTimeOk returns a tuple with the CheckInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInTime

`func (o *PostTemporaryResidenceFields) SetCheckInTime(v string)`

SetCheckInTime sets CheckInTime field to given value.

### HasCheckInTime

`func (o *PostTemporaryResidenceFields) HasCheckInTime() bool`

HasCheckInTime returns a boolean if a field has been set.

### GetCheckOutTime

`func (o *PostTemporaryResidenceFields) GetCheckOutTime() string`

GetCheckOutTime returns the CheckOutTime field if non-nil, zero value otherwise.

### GetCheckOutTimeOk

`func (o *PostTemporaryResidenceFields) GetCheckOutTimeOk() (*string, bool)`

GetCheckOutTimeOk returns a tuple with the CheckOutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOutTime

`func (o *PostTemporaryResidenceFields) SetCheckOutTime(v string)`

SetCheckOutTime sets CheckOutTime field to given value.

### HasCheckOutTime

`func (o *PostTemporaryResidenceFields) HasCheckOutTime() bool`

HasCheckOutTime returns a boolean if a field has been set.

### GetComfortAmenities

`func (o *PostTemporaryResidenceFields) GetComfortAmenities() []TemporaryResidenceFieldsComfortAmenity`

GetComfortAmenities returns the ComfortAmenities field if non-nil, zero value otherwise.

### GetComfortAmenitiesOk

`func (o *PostTemporaryResidenceFields) GetComfortAmenitiesOk() (*[]TemporaryResidenceFieldsComfortAmenity, bool)`

GetComfortAmenitiesOk returns a tuple with the ComfortAmenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComfortAmenities

`func (o *PostTemporaryResidenceFields) SetComfortAmenities(v []TemporaryResidenceFieldsComfortAmenity)`

SetComfortAmenities sets ComfortAmenities field to given value.

### HasComfortAmenities

`func (o *PostTemporaryResidenceFields) HasComfortAmenities() bool`

HasComfortAmenities returns a boolean if a field has been set.

### GetDamageDeposit

`func (o *PostTemporaryResidenceFields) GetDamageDeposit() string`

GetDamageDeposit returns the DamageDeposit field if non-nil, zero value otherwise.

### GetDamageDepositOk

`func (o *PostTemporaryResidenceFields) GetDamageDepositOk() (*string, bool)`

GetDamageDepositOk returns a tuple with the DamageDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageDeposit

`func (o *PostTemporaryResidenceFields) SetDamageDeposit(v string)`

SetDamageDeposit sets DamageDeposit field to given value.

### HasDamageDeposit

`func (o *PostTemporaryResidenceFields) HasDamageDeposit() bool`

HasDamageDeposit returns a boolean if a field has been set.

### GetExtraPersonCapacity

`func (o *PostTemporaryResidenceFields) GetExtraPersonCapacity() int32`

GetExtraPersonCapacity returns the ExtraPersonCapacity field if non-nil, zero value otherwise.

### GetExtraPersonCapacityOk

`func (o *PostTemporaryResidenceFields) GetExtraPersonCapacityOk() (*int32, bool)`

GetExtraPersonCapacityOk returns a tuple with the ExtraPersonCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPersonCapacity

`func (o *PostTemporaryResidenceFields) SetExtraPersonCapacity(v int32)`

SetExtraPersonCapacity sets ExtraPersonCapacity field to given value.

### HasExtraPersonCapacity

`func (o *PostTemporaryResidenceFields) HasExtraPersonCapacity() bool`

HasExtraPersonCapacity returns a boolean if a field has been set.

### GetFullyFurnished

`func (o *PostTemporaryResidenceFields) GetFullyFurnished() bool`

GetFullyFurnished returns the FullyFurnished field if non-nil, zero value otherwise.

### GetFullyFurnishedOk

`func (o *PostTemporaryResidenceFields) GetFullyFurnishedOk() (*bool, bool)`

GetFullyFurnishedOk returns a tuple with the FullyFurnished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyFurnished

`func (o *PostTemporaryResidenceFields) SetFullyFurnished(v bool)`

SetFullyFurnished sets FullyFurnished field to given value.

### HasFullyFurnished

`func (o *PostTemporaryResidenceFields) HasFullyFurnished() bool`

HasFullyFurnished returns a boolean if a field has been set.

### GetHasOwnImage

`func (o *PostTemporaryResidenceFields) GetHasOwnImage() bool`

GetHasOwnImage returns the HasOwnImage field if non-nil, zero value otherwise.

### GetHasOwnImageOk

`func (o *PostTemporaryResidenceFields) GetHasOwnImageOk() (*bool, bool)`

GetHasOwnImageOk returns a tuple with the HasOwnImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOwnImage

`func (o *PostTemporaryResidenceFields) SetHasOwnImage(v bool)`

SetHasOwnImage sets HasOwnImage field to given value.

### HasHasOwnImage

`func (o *PostTemporaryResidenceFields) HasHasOwnImage() bool`

HasHasOwnImage returns a boolean if a field has been set.

### GetHeatingCoolingSystem

`func (o *PostTemporaryResidenceFields) GetHeatingCoolingSystem() []TemporaryResidenceFieldsHeatingCoolingSystem`

GetHeatingCoolingSystem returns the HeatingCoolingSystem field if non-nil, zero value otherwise.

### GetHeatingCoolingSystemOk

`func (o *PostTemporaryResidenceFields) GetHeatingCoolingSystemOk() (*[]TemporaryResidenceFieldsHeatingCoolingSystem, bool)`

GetHeatingCoolingSystemOk returns a tuple with the HeatingCoolingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatingCoolingSystem

`func (o *PostTemporaryResidenceFields) SetHeatingCoolingSystem(v []TemporaryResidenceFieldsHeatingCoolingSystem)`

SetHeatingCoolingSystem sets HeatingCoolingSystem field to given value.

### HasHeatingCoolingSystem

`func (o *PostTemporaryResidenceFields) HasHeatingCoolingSystem() bool`

HasHeatingCoolingSystem returns a boolean if a field has been set.

### GetHouseRules

`func (o *PostTemporaryResidenceFields) GetHouseRules() string`

GetHouseRules returns the HouseRules field if non-nil, zero value otherwise.

### GetHouseRulesOk

`func (o *PostTemporaryResidenceFields) GetHouseRulesOk() (*string, bool)`

GetHouseRulesOk returns a tuple with the HouseRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseRules

`func (o *PostTemporaryResidenceFields) SetHouseRules(v string)`

SetHouseRules sets HouseRules field to given value.

### HasHouseRules

`func (o *PostTemporaryResidenceFields) HasHouseRules() bool`

HasHouseRules returns a boolean if a field has been set.

### GetMinimumStay

`func (o *PostTemporaryResidenceFields) GetMinimumStay() int32`

GetMinimumStay returns the MinimumStay field if non-nil, zero value otherwise.

### GetMinimumStayOk

`func (o *PostTemporaryResidenceFields) GetMinimumStayOk() (*int32, bool)`

GetMinimumStayOk returns a tuple with the MinimumStay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumStay

`func (o *PostTemporaryResidenceFields) SetMinimumStay(v int32)`

SetMinimumStay sets MinimumStay field to given value.

### HasMinimumStay

`func (o *PostTemporaryResidenceFields) HasMinimumStay() bool`

HasMinimumStay returns a boolean if a field has been set.

### GetPetsAllowed

`func (o *PostTemporaryResidenceFields) GetPetsAllowed() TemporaryResidenceFieldsPetsAllowed`

GetPetsAllowed returns the PetsAllowed field if non-nil, zero value otherwise.

### GetPetsAllowedOk

`func (o *PostTemporaryResidenceFields) GetPetsAllowedOk() (*TemporaryResidenceFieldsPetsAllowed, bool)`

GetPetsAllowedOk returns a tuple with the PetsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPetsAllowed

`func (o *PostTemporaryResidenceFields) SetPetsAllowed(v TemporaryResidenceFieldsPetsAllowed)`

SetPetsAllowed sets PetsAllowed field to given value.

### HasPetsAllowed

`func (o *PostTemporaryResidenceFields) HasPetsAllowed() bool`

HasPetsAllowed returns a boolean if a field has been set.

### GetPriceCostPerExtraPerson

`func (o *PostTemporaryResidenceFields) GetPriceCostPerExtraPerson() string`

GetPriceCostPerExtraPerson returns the PriceCostPerExtraPerson field if non-nil, zero value otherwise.

### GetPriceCostPerExtraPersonOk

`func (o *PostTemporaryResidenceFields) GetPriceCostPerExtraPersonOk() (*string, bool)`

GetPriceCostPerExtraPersonOk returns a tuple with the PriceCostPerExtraPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCostPerExtraPerson

`func (o *PostTemporaryResidenceFields) SetPriceCostPerExtraPerson(v string)`

SetPriceCostPerExtraPerson sets PriceCostPerExtraPerson field to given value.

### HasPriceCostPerExtraPerson

`func (o *PostTemporaryResidenceFields) HasPriceCostPerExtraPerson() bool`

HasPriceCostPerExtraPerson returns a boolean if a field has been set.

### GetPriceRegularDays

`func (o *PostTemporaryResidenceFields) GetPriceRegularDays() string`

GetPriceRegularDays returns the PriceRegularDays field if non-nil, zero value otherwise.

### GetPriceRegularDaysOk

`func (o *PostTemporaryResidenceFields) GetPriceRegularDaysOk() (*string, bool)`

GetPriceRegularDaysOk returns a tuple with the PriceRegularDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRegularDays

`func (o *PostTemporaryResidenceFields) SetPriceRegularDays(v string)`

SetPriceRegularDays sets PriceRegularDays field to given value.

### HasPriceRegularDays

`func (o *PostTemporaryResidenceFields) HasPriceRegularDays() bool`

HasPriceRegularDays returns a boolean if a field has been set.

### GetPriceSpecialDays

`func (o *PostTemporaryResidenceFields) GetPriceSpecialDays() string`

GetPriceSpecialDays returns the PriceSpecialDays field if non-nil, zero value otherwise.

### GetPriceSpecialDaysOk

`func (o *PostTemporaryResidenceFields) GetPriceSpecialDaysOk() (*string, bool)`

GetPriceSpecialDaysOk returns a tuple with the PriceSpecialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSpecialDays

`func (o *PostTemporaryResidenceFields) SetPriceSpecialDays(v string)`

SetPriceSpecialDays sets PriceSpecialDays field to given value.

### HasPriceSpecialDays

`func (o *PostTemporaryResidenceFields) HasPriceSpecialDays() bool`

HasPriceSpecialDays returns a boolean if a field has been set.

### GetPriceWeekends

`func (o *PostTemporaryResidenceFields) GetPriceWeekends() string`

GetPriceWeekends returns the PriceWeekends field if non-nil, zero value otherwise.

### GetPriceWeekendsOk

`func (o *PostTemporaryResidenceFields) GetPriceWeekendsOk() (*string, bool)`

GetPriceWeekendsOk returns a tuple with the PriceWeekends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceWeekends

`func (o *PostTemporaryResidenceFields) SetPriceWeekends(v string)`

SetPriceWeekends sets PriceWeekends field to given value.

### HasPriceWeekends

`func (o *PostTemporaryResidenceFields) HasPriceWeekends() bool`

HasPriceWeekends returns a boolean if a field has been set.

### GetRegularPersonCapacity

`func (o *PostTemporaryResidenceFields) GetRegularPersonCapacity() int32`

GetRegularPersonCapacity returns the RegularPersonCapacity field if non-nil, zero value otherwise.

### GetRegularPersonCapacityOk

`func (o *PostTemporaryResidenceFields) GetRegularPersonCapacityOk() (*int32, bool)`

GetRegularPersonCapacityOk returns a tuple with the RegularPersonCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularPersonCapacity

`func (o *PostTemporaryResidenceFields) SetRegularPersonCapacity(v int32)`

SetRegularPersonCapacity sets RegularPersonCapacity field to given value.

### HasRegularPersonCapacity

`func (o *PostTemporaryResidenceFields) HasRegularPersonCapacity() bool`

HasRegularPersonCapacity returns a boolean if a field has been set.

### GetRentalPeriod

`func (o *PostTemporaryResidenceFields) GetRentalPeriod() TemporaryResidenceFieldsRentalPeriod`

GetRentalPeriod returns the RentalPeriod field if non-nil, zero value otherwise.

### GetRentalPeriodOk

`func (o *PostTemporaryResidenceFields) GetRentalPeriodOk() (*TemporaryResidenceFieldsRentalPeriod, bool)`

GetRentalPeriodOk returns a tuple with the RentalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentalPeriod

`func (o *PostTemporaryResidenceFields) SetRentalPeriod(v TemporaryResidenceFieldsRentalPeriod)`

SetRentalPeriod sets RentalPeriod field to given value.

### HasRentalPeriod

`func (o *PostTemporaryResidenceFields) HasRentalPeriod() bool`

HasRentalPeriod returns a boolean if a field has been set.

### GetRoomsCount

`func (o *PostTemporaryResidenceFields) GetRoomsCount() TemporaryResidenceFieldsRoomsCount`

GetRoomsCount returns the RoomsCount field if non-nil, zero value otherwise.

### GetRoomsCountOk

`func (o *PostTemporaryResidenceFields) GetRoomsCountOk() (*TemporaryResidenceFieldsRoomsCount, bool)`

GetRoomsCountOk returns a tuple with the RoomsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsCount

`func (o *PostTemporaryResidenceFields) SetRoomsCount(v TemporaryResidenceFieldsRoomsCount)`

SetRoomsCount sets RoomsCount field to given value.

### HasRoomsCount

`func (o *PostTemporaryResidenceFields) HasRoomsCount() bool`

HasRoomsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


