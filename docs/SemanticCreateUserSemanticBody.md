# SemanticCreateUserSemanticBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** | شماره موبایل کاربر | 
**Semantic** | **map[string]string** | مپ key-value اطلاعات معنایی برای ذخیره | 
**Cost** | Pointer to **int32** | هزینه مرتبط با عملیات | [optional] 
**TicketUuid** | Pointer to **string** | UUID تیکت پرداخت (اختیاری) | [optional] 

## Methods

### NewSemanticCreateUserSemanticBody

`func NewSemanticCreateUserSemanticBody(phone string, semantic map[string]string, ) *SemanticCreateUserSemanticBody`

NewSemanticCreateUserSemanticBody instantiates a new SemanticCreateUserSemanticBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSemanticCreateUserSemanticBodyWithDefaults

`func NewSemanticCreateUserSemanticBodyWithDefaults() *SemanticCreateUserSemanticBody`

NewSemanticCreateUserSemanticBodyWithDefaults instantiates a new SemanticCreateUserSemanticBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SemanticCreateUserSemanticBody) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SemanticCreateUserSemanticBody) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SemanticCreateUserSemanticBody) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetSemantic

`func (o *SemanticCreateUserSemanticBody) GetSemantic() map[string]string`

GetSemantic returns the Semantic field if non-nil, zero value otherwise.

### GetSemanticOk

`func (o *SemanticCreateUserSemanticBody) GetSemanticOk() (*map[string]string, bool)`

GetSemanticOk returns a tuple with the Semantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantic

`func (o *SemanticCreateUserSemanticBody) SetSemantic(v map[string]string)`

SetSemantic sets Semantic field to given value.


### GetCost

`func (o *SemanticCreateUserSemanticBody) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SemanticCreateUserSemanticBody) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SemanticCreateUserSemanticBody) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SemanticCreateUserSemanticBody) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetTicketUuid

`func (o *SemanticCreateUserSemanticBody) GetTicketUuid() string`

GetTicketUuid returns the TicketUuid field if non-nil, zero value otherwise.

### GetTicketUuidOk

`func (o *SemanticCreateUserSemanticBody) GetTicketUuidOk() (*string, bool)`

GetTicketUuidOk returns a tuple with the TicketUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketUuid

`func (o *SemanticCreateUserSemanticBody) SetTicketUuid(v string)`

SetTicketUuid sets TicketUuid field to given value.

### HasTicketUuid

`func (o *SemanticCreateUserSemanticBody) HasTicketUuid() bool`

HasTicketUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


