# OpenPlatformpostServicesFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**PostServicesFieldsCategory**](PostServicesFieldsCategory.md) |  | [optional] 
**ExpertiseIds** | Pointer to **[]string** | List of expertise ids | [optional] 
**WorkHoursEnd** | Pointer to **int32** | End hour of work in 24-hour format (e.g. 18 for 18:00). Only applicable if &#x60;works_24_7&#x60; is false. | [optional] 
**WorkHoursStart** | Pointer to **int32** | Start hour of work in 24-hour format (e.g. 9 for 9:00). Only applicable if &#x60;works_24_7&#x60; is false. | [optional] 
**WorkOnHolidays** | Pointer to **bool** | Whether the service provider works on holidays | [optional] 
**Works247** | Pointer to **bool** | Whether the service provider is available 24/7. If true, &#x60;work_hours_start&#x60; and &#x60;work_hours_end&#x60; are ignored. | [optional] 

## Methods

### NewOpenPlatformpostServicesFields

`func NewOpenPlatformpostServicesFields() *OpenPlatformpostServicesFields`

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

### HasCategory

`func (o *OpenPlatformpostServicesFields) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

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

### HasExpertiseIds

`func (o *OpenPlatformpostServicesFields) HasExpertiseIds() bool`

HasExpertiseIds returns a boolean if a field has been set.

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

### HasWorkHoursEnd

`func (o *OpenPlatformpostServicesFields) HasWorkHoursEnd() bool`

HasWorkHoursEnd returns a boolean if a field has been set.

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

### HasWorkHoursStart

`func (o *OpenPlatformpostServicesFields) HasWorkHoursStart() bool`

HasWorkHoursStart returns a boolean if a field has been set.

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

### HasWorkOnHolidays

`func (o *OpenPlatformpostServicesFields) HasWorkOnHolidays() bool`

HasWorkOnHolidays returns a boolean if a field has been set.

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

### HasWorks247

`func (o *OpenPlatformpostServicesFields) HasWorks247() bool`

HasWorks247 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


