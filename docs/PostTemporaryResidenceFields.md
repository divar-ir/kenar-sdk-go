# PostTemporaryResidenceFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **int32** | متراژ اقامتگاه به متر مربع | [optional] 
**ExtraPersonCapacity** | Pointer to **int32** | تعداد افراد اضافه مجاز در اقامتگاه | [optional] 
**HasOwnImage** | Pointer to **bool** | تصاویر مربوط به خود ملک بوده و تزئینی نیستند. | [optional] 
**PriceCostPerExtraPerson** | Pointer to **string** | هزینه هر نفر اضافه به ازای هر شب به تومان | [optional] 
**PriceRegularDays** | Pointer to **string** | قیمت اقامتگاه در روزهای عادی (شنبه تا سه‌شنبه) به تومان | [optional] 
**PriceSpecialDays** | Pointer to **string** | قیمت اقامتگاه در روزهای خاص (تعطیلات و مناسبت‌ها) به تومان | [optional] 
**PriceWeekends** | Pointer to **string** | قیمت اقامتگاه در آخر هفته (چهارشنبه تا جمعه) به تومان | [optional] 
**RegularPersonCapacity** | Pointer to **int32** | ظرفیت استاندارد افراد در اقامتگاه | [optional] 
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


