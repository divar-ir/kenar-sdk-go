# OpenPlatformpostServicesFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**PostServicesFieldsCategory**](PostServicesFieldsCategory.md) |  | 
**ExpertiseIds** | **[]string** | List of expertise ids | 
**WorkHoursEnd** | **int32** | End hour of work in 24-hour format (e.g. 18 for 18:00). Only applicable if &#x60;works_24_7&#x60; is false. | 
**WorkHoursStart** | **int32** | Start hour of work in 24-hour format (e.g. 9 for 9:00). Only applicable if &#x60;works_24_7&#x60; is false. | 
**WorkOnHolidays** | **bool** | Whether the service provider works on holidays | 
**Works247** | **bool** | Whether the service provider is available 24/7. If true, &#x60;work_hours_start&#x60; and &#x60;work_hours_end&#x60; are ignored. | 

## Methods

### NewOpenPlatformpostServicesFields

`func NewOpenPlatformpostServicesFields(category PostServicesFieldsCategory, expertiseIds []string, workHoursEnd int32, workHoursStart int32, workOnHolidays bool, works247 bool, ) *OpenPlatformpostServicesFields`

NewOpenPlatformpostServicesFields instantiates a new OpenPlatformpostServicesFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenPlatformpostServicesFieldsWithDefaults

`func NewOpenPlatformpostServicesFieldsWithDefaults() *OpenPlatformpostServicesFields`

NewOpenPlatformpostServicesFieldsWithDefaults instantiates a new OpenPlatformpostServicesFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *OpenPlatformpostServicesFields) GetCategory() PostServicesFieldsCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OpenPlatformpostServicesFields) GetCategoryOk() (*PostServicesFieldsCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OpenPlatformpostServicesFields) SetCategory(v PostServicesFieldsCategory)`

SetCategory sets Category field to given value.


### GetExpertiseIds

`func (o *OpenPlatformpostServicesFields) GetExpertiseIds() []string`

GetExpertiseIds returns the ExpertiseIds field if non-nil, zero value otherwise.

### GetExpertiseIdsOk

`func (o *OpenPlatformpostServicesFields) GetExpertiseIdsOk() (*[]string, bool)`

GetExpertiseIdsOk returns a tuple with the ExpertiseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpertiseIds

`func (o *OpenPlatformpostServicesFields) SetExpertiseIds(v []string)`

SetExpertiseIds sets ExpertiseIds field to given value.


### GetWorkHoursEnd

`func (o *OpenPlatformpostServicesFields) GetWorkHoursEnd() int32`

GetWorkHoursEnd returns the WorkHoursEnd field if non-nil, zero value otherwise.

### GetWorkHoursEndOk

`func (o *OpenPlatformpostServicesFields) GetWorkHoursEndOk() (*int32, bool)`

GetWorkHoursEndOk returns a tuple with the WorkHoursEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkHoursEnd

`func (o *OpenPlatformpostServicesFields) SetWorkHoursEnd(v int32)`

SetWorkHoursEnd sets WorkHoursEnd field to given value.


### GetWorkHoursStart

`func (o *OpenPlatformpostServicesFields) GetWorkHoursStart() int32`

GetWorkHoursStart returns the WorkHoursStart field if non-nil, zero value otherwise.

### GetWorkHoursStartOk

`func (o *OpenPlatformpostServicesFields) GetWorkHoursStartOk() (*int32, bool)`

GetWorkHoursStartOk returns a tuple with the WorkHoursStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkHoursStart

`func (o *OpenPlatformpostServicesFields) SetWorkHoursStart(v int32)`

SetWorkHoursStart sets WorkHoursStart field to given value.


### GetWorkOnHolidays

`func (o *OpenPlatformpostServicesFields) GetWorkOnHolidays() bool`

GetWorkOnHolidays returns the WorkOnHolidays field if non-nil, zero value otherwise.

### GetWorkOnHolidaysOk

`func (o *OpenPlatformpostServicesFields) GetWorkOnHolidaysOk() (*bool, bool)`

GetWorkOnHolidaysOk returns a tuple with the WorkOnHolidays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkOnHolidays

`func (o *OpenPlatformpostServicesFields) SetWorkOnHolidays(v bool)`

SetWorkOnHolidays sets WorkOnHolidays field to given value.


### GetWorks247

`func (o *OpenPlatformpostServicesFields) GetWorks247() bool`

GetWorks247 returns the Works247 field if non-nil, zero value otherwise.

### GetWorks247Ok

`func (o *OpenPlatformpostServicesFields) GetWorks247Ok() (*bool, bool)`

GetWorks247Ok returns a tuple with the Works247 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorks247

`func (o *OpenPlatformpostServicesFields) SetWorks247(v bool)`

SetWorks247 sets Works247 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


