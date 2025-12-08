# SemanticCreatePostSemanticBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Semantic** | **map[string]string** | مپ key-value اطلاعات معنایی برای ذخیره | 
**Cost** | Pointer to **int32** | هزینه مرتبط با عملیات | [optional] 
**TicketUuid** | Pointer to **string** | UUID تیکت پرداخت (اختیاری) | [optional] 

## Methods

### NewSemanticCreatePostSemanticBody

`func NewSemanticCreatePostSemanticBody(semantic map[string]string, ) *SemanticCreatePostSemanticBody`

NewSemanticCreatePostSemanticBody instantiates a new SemanticCreatePostSemanticBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSemanticCreatePostSemanticBodyWithDefaults

`func NewSemanticCreatePostSemanticBodyWithDefaults() *SemanticCreatePostSemanticBody`

NewSemanticCreatePostSemanticBodyWithDefaults instantiates a new SemanticCreatePostSemanticBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSemantic

`func (o *SemanticCreatePostSemanticBody) GetSemantic() map[string]string`

GetSemantic returns the Semantic field if non-nil, zero value otherwise.

### GetSemanticOk

`func (o *SemanticCreatePostSemanticBody) GetSemanticOk() (*map[string]string, bool)`

GetSemanticOk returns a tuple with the Semantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantic

`func (o *SemanticCreatePostSemanticBody) SetSemantic(v map[string]string)`

SetSemantic sets Semantic field to given value.


### GetCost

`func (o *SemanticCreatePostSemanticBody) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SemanticCreatePostSemanticBody) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SemanticCreatePostSemanticBody) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SemanticCreatePostSemanticBody) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetTicketUuid

`func (o *SemanticCreatePostSemanticBody) GetTicketUuid() string`

GetTicketUuid returns the TicketUuid field if non-nil, zero value otherwise.

### GetTicketUuidOk

`func (o *SemanticCreatePostSemanticBody) GetTicketUuidOk() (*string, bool)`

GetTicketUuidOk returns a tuple with the TicketUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketUuid

`func (o *SemanticCreatePostSemanticBody) SetTicketUuid(v string)`

SetTicketUuid sets TicketUuid field to given value.

### HasTicketUuid

`func (o *SemanticCreatePostSemanticBody) HasTicketUuid() bool`

HasTicketUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


