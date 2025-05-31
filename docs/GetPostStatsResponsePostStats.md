# GetPostStatsResponsePostStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Daily** | Pointer to [**[]GetPostStatsResponseDailyStats**](GetPostStatsResponseDailyStats.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetPostStatsResponsePostStats

`func NewGetPostStatsResponsePostStats() *GetPostStatsResponsePostStats`

NewGetPostStatsResponsePostStats instantiates a new GetPostStatsResponsePostStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPostStatsResponsePostStatsWithDefaults

`func NewGetPostStatsResponsePostStatsWithDefaults() *GetPostStatsResponsePostStats`

NewGetPostStatsResponsePostStatsWithDefaults instantiates a new GetPostStatsResponsePostStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaily

`func (o *GetPostStatsResponsePostStats) GetDaily() []GetPostStatsResponseDailyStats`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *GetPostStatsResponsePostStats) GetDailyOk() (*[]GetPostStatsResponseDailyStats, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *GetPostStatsResponsePostStats) SetDaily(v []GetPostStatsResponseDailyStats)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *GetPostStatsResponsePostStats) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetTotal

`func (o *GetPostStatsResponsePostStats) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetPostStatsResponsePostStats) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetPostStatsResponsePostStats) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetPostStatsResponsePostStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


