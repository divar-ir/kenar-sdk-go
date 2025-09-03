# PostHomePresellFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasOwnImage** | **bool** | تصاویر مربوط به خود ملک بوده و تزئینی نیستند. | 
**BasePricePerSquareMeter** | Pointer to **string** | Base price per square meter in Toman | [optional] 
**ConstructionPhase** | Pointer to [**HomePresellFieldsConstructionPhase**](HomePresellFieldsConstructionPhase.md) |  | [optional] 
**DeliveryMonth** | Pointer to [**HomePresellFieldsDeliveryMonth**](HomePresellFieldsDeliveryMonth.md) |  | [optional] 
**DeliveryPaymentPercentage** | Pointer to **int32** | Payment percentage required upon delivery | [optional] 
**DeliveryYear** | Pointer to [**HomePresellFieldsDeliveryYear**](HomePresellFieldsDeliveryYear.md) |  | [optional] 
**DeveloperCompanyName** | Pointer to **string** | Name of the developer company | [optional] 
**DownPaymentPercentage** | Pointer to **int32** | Down payment percentage required | [optional] 
**MinUnitSize** | Pointer to **int32** | Minimum unit size in square meters | [optional] 
**ProjectName** | Pointer to **string** | Name of the home presell project | [optional] 
**ProjectPhysicalProgressPercentage** | Pointer to **int32** | Physical progress percentage of the project | [optional] 
**UnitTypesOffered** | Pointer to [**[]HomePresellFieldsUnitType**](HomePresellFieldsUnitType.md) | List of unit types offered in the project | [optional] 

## Methods

### NewPostHomePresellFields

`func NewPostHomePresellFields(hasOwnImage bool, ) *PostHomePresellFields`

NewPostHomePresellFields instantiates a new PostHomePresellFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostHomePresellFieldsWithDefaults

`func NewPostHomePresellFieldsWithDefaults() *PostHomePresellFields`

NewPostHomePresellFieldsWithDefaults instantiates a new PostHomePresellFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasOwnImage

`func (o *PostHomePresellFields) GetHasOwnImage() bool`

GetHasOwnImage returns the HasOwnImage field if non-nil, zero value otherwise.

### GetHasOwnImageOk

`func (o *PostHomePresellFields) GetHasOwnImageOk() (*bool, bool)`

GetHasOwnImageOk returns a tuple with the HasOwnImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOwnImage

`func (o *PostHomePresellFields) SetHasOwnImage(v bool)`

SetHasOwnImage sets HasOwnImage field to given value.


### GetBasePricePerSquareMeter

`func (o *PostHomePresellFields) GetBasePricePerSquareMeter() string`

GetBasePricePerSquareMeter returns the BasePricePerSquareMeter field if non-nil, zero value otherwise.

### GetBasePricePerSquareMeterOk

`func (o *PostHomePresellFields) GetBasePricePerSquareMeterOk() (*string, bool)`

GetBasePricePerSquareMeterOk returns a tuple with the BasePricePerSquareMeter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePricePerSquareMeter

`func (o *PostHomePresellFields) SetBasePricePerSquareMeter(v string)`

SetBasePricePerSquareMeter sets BasePricePerSquareMeter field to given value.

### HasBasePricePerSquareMeter

`func (o *PostHomePresellFields) HasBasePricePerSquareMeter() bool`

HasBasePricePerSquareMeter returns a boolean if a field has been set.

### GetConstructionPhase

`func (o *PostHomePresellFields) GetConstructionPhase() HomePresellFieldsConstructionPhase`

GetConstructionPhase returns the ConstructionPhase field if non-nil, zero value otherwise.

### GetConstructionPhaseOk

`func (o *PostHomePresellFields) GetConstructionPhaseOk() (*HomePresellFieldsConstructionPhase, bool)`

GetConstructionPhaseOk returns a tuple with the ConstructionPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstructionPhase

`func (o *PostHomePresellFields) SetConstructionPhase(v HomePresellFieldsConstructionPhase)`

SetConstructionPhase sets ConstructionPhase field to given value.

### HasConstructionPhase

`func (o *PostHomePresellFields) HasConstructionPhase() bool`

HasConstructionPhase returns a boolean if a field has been set.

### GetDeliveryMonth

`func (o *PostHomePresellFields) GetDeliveryMonth() HomePresellFieldsDeliveryMonth`

GetDeliveryMonth returns the DeliveryMonth field if non-nil, zero value otherwise.

### GetDeliveryMonthOk

`func (o *PostHomePresellFields) GetDeliveryMonthOk() (*HomePresellFieldsDeliveryMonth, bool)`

GetDeliveryMonthOk returns a tuple with the DeliveryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMonth

`func (o *PostHomePresellFields) SetDeliveryMonth(v HomePresellFieldsDeliveryMonth)`

SetDeliveryMonth sets DeliveryMonth field to given value.

### HasDeliveryMonth

`func (o *PostHomePresellFields) HasDeliveryMonth() bool`

HasDeliveryMonth returns a boolean if a field has been set.

### GetDeliveryPaymentPercentage

`func (o *PostHomePresellFields) GetDeliveryPaymentPercentage() int32`

GetDeliveryPaymentPercentage returns the DeliveryPaymentPercentage field if non-nil, zero value otherwise.

### GetDeliveryPaymentPercentageOk

`func (o *PostHomePresellFields) GetDeliveryPaymentPercentageOk() (*int32, bool)`

GetDeliveryPaymentPercentageOk returns a tuple with the DeliveryPaymentPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPaymentPercentage

`func (o *PostHomePresellFields) SetDeliveryPaymentPercentage(v int32)`

SetDeliveryPaymentPercentage sets DeliveryPaymentPercentage field to given value.

### HasDeliveryPaymentPercentage

`func (o *PostHomePresellFields) HasDeliveryPaymentPercentage() bool`

HasDeliveryPaymentPercentage returns a boolean if a field has been set.

### GetDeliveryYear

`func (o *PostHomePresellFields) GetDeliveryYear() HomePresellFieldsDeliveryYear`

GetDeliveryYear returns the DeliveryYear field if non-nil, zero value otherwise.

### GetDeliveryYearOk

`func (o *PostHomePresellFields) GetDeliveryYearOk() (*HomePresellFieldsDeliveryYear, bool)`

GetDeliveryYearOk returns a tuple with the DeliveryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryYear

`func (o *PostHomePresellFields) SetDeliveryYear(v HomePresellFieldsDeliveryYear)`

SetDeliveryYear sets DeliveryYear field to given value.

### HasDeliveryYear

`func (o *PostHomePresellFields) HasDeliveryYear() bool`

HasDeliveryYear returns a boolean if a field has been set.

### GetDeveloperCompanyName

`func (o *PostHomePresellFields) GetDeveloperCompanyName() string`

GetDeveloperCompanyName returns the DeveloperCompanyName field if non-nil, zero value otherwise.

### GetDeveloperCompanyNameOk

`func (o *PostHomePresellFields) GetDeveloperCompanyNameOk() (*string, bool)`

GetDeveloperCompanyNameOk returns a tuple with the DeveloperCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperCompanyName

`func (o *PostHomePresellFields) SetDeveloperCompanyName(v string)`

SetDeveloperCompanyName sets DeveloperCompanyName field to given value.

### HasDeveloperCompanyName

`func (o *PostHomePresellFields) HasDeveloperCompanyName() bool`

HasDeveloperCompanyName returns a boolean if a field has been set.

### GetDownPaymentPercentage

`func (o *PostHomePresellFields) GetDownPaymentPercentage() int32`

GetDownPaymentPercentage returns the DownPaymentPercentage field if non-nil, zero value otherwise.

### GetDownPaymentPercentageOk

`func (o *PostHomePresellFields) GetDownPaymentPercentageOk() (*int32, bool)`

GetDownPaymentPercentageOk returns a tuple with the DownPaymentPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPaymentPercentage

`func (o *PostHomePresellFields) SetDownPaymentPercentage(v int32)`

SetDownPaymentPercentage sets DownPaymentPercentage field to given value.

### HasDownPaymentPercentage

`func (o *PostHomePresellFields) HasDownPaymentPercentage() bool`

HasDownPaymentPercentage returns a boolean if a field has been set.

### GetMinUnitSize

`func (o *PostHomePresellFields) GetMinUnitSize() int32`

GetMinUnitSize returns the MinUnitSize field if non-nil, zero value otherwise.

### GetMinUnitSizeOk

`func (o *PostHomePresellFields) GetMinUnitSizeOk() (*int32, bool)`

GetMinUnitSizeOk returns a tuple with the MinUnitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUnitSize

`func (o *PostHomePresellFields) SetMinUnitSize(v int32)`

SetMinUnitSize sets MinUnitSize field to given value.

### HasMinUnitSize

`func (o *PostHomePresellFields) HasMinUnitSize() bool`

HasMinUnitSize returns a boolean if a field has been set.

### GetProjectName

`func (o *PostHomePresellFields) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *PostHomePresellFields) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *PostHomePresellFields) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *PostHomePresellFields) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetProjectPhysicalProgressPercentage

`func (o *PostHomePresellFields) GetProjectPhysicalProgressPercentage() int32`

GetProjectPhysicalProgressPercentage returns the ProjectPhysicalProgressPercentage field if non-nil, zero value otherwise.

### GetProjectPhysicalProgressPercentageOk

`func (o *PostHomePresellFields) GetProjectPhysicalProgressPercentageOk() (*int32, bool)`

GetProjectPhysicalProgressPercentageOk returns a tuple with the ProjectPhysicalProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPhysicalProgressPercentage

`func (o *PostHomePresellFields) SetProjectPhysicalProgressPercentage(v int32)`

SetProjectPhysicalProgressPercentage sets ProjectPhysicalProgressPercentage field to given value.

### HasProjectPhysicalProgressPercentage

`func (o *PostHomePresellFields) HasProjectPhysicalProgressPercentage() bool`

HasProjectPhysicalProgressPercentage returns a boolean if a field has been set.

### GetUnitTypesOffered

`func (o *PostHomePresellFields) GetUnitTypesOffered() []HomePresellFieldsUnitType`

GetUnitTypesOffered returns the UnitTypesOffered field if non-nil, zero value otherwise.

### GetUnitTypesOfferedOk

`func (o *PostHomePresellFields) GetUnitTypesOfferedOk() (*[]HomePresellFieldsUnitType, bool)`

GetUnitTypesOfferedOk returns a tuple with the UnitTypesOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitTypesOffered

`func (o *PostHomePresellFields) SetUnitTypesOffered(v []HomePresellFieldsUnitType)`

SetUnitTypesOffered sets UnitTypesOffered field to given value.

### HasUnitTypesOffered

`func (o *PostHomePresellFields) HasUnitTypesOffered() bool`

HasUnitTypesOffered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


