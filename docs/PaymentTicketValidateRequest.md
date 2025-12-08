# PaymentTicketValidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceCost** | **int32** | هزینه سرویس به ریال | 
**TicketUuid** | **string** | شناسه منحصر به فرد تیکت پرداخت | 
**PhoneNumber** | Pointer to **string** | شماره تلفن کاربر (به جای آن از user_id استفاده کنید) | [optional] 
**UserId** | Pointer to **string** | شناسه منحصر به فرد کاربر | [optional] 

## Methods

### NewPaymentTicketValidateRequest

`func NewPaymentTicketValidateRequest(serviceCost int32, ticketUuid string, ) *PaymentTicketValidateRequest`

NewPaymentTicketValidateRequest instantiates a new PaymentTicketValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTicketValidateRequestWithDefaults

`func NewPaymentTicketValidateRequestWithDefaults() *PaymentTicketValidateRequest`

NewPaymentTicketValidateRequestWithDefaults instantiates a new PaymentTicketValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceCost

`func (o *PaymentTicketValidateRequest) GetServiceCost() int32`

GetServiceCost returns the ServiceCost field if non-nil, zero value otherwise.

### GetServiceCostOk

`func (o *PaymentTicketValidateRequest) GetServiceCostOk() (*int32, bool)`

GetServiceCostOk returns a tuple with the ServiceCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCost

`func (o *PaymentTicketValidateRequest) SetServiceCost(v int32)`

SetServiceCost sets ServiceCost field to given value.


### GetTicketUuid

`func (o *PaymentTicketValidateRequest) GetTicketUuid() string`

GetTicketUuid returns the TicketUuid field if non-nil, zero value otherwise.

### GetTicketUuidOk

`func (o *PaymentTicketValidateRequest) GetTicketUuidOk() (*string, bool)`

GetTicketUuidOk returns a tuple with the TicketUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketUuid

`func (o *PaymentTicketValidateRequest) SetTicketUuid(v string)`

SetTicketUuid sets TicketUuid field to given value.


### GetPhoneNumber

`func (o *PaymentTicketValidateRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PaymentTicketValidateRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PaymentTicketValidateRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PaymentTicketValidateRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetUserId

`func (o *PaymentTicketValidateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PaymentTicketValidateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PaymentTicketValidateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PaymentTicketValidateRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


