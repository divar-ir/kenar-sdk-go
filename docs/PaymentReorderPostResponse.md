# PaymentReorderPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostRials** | Pointer to **string** | هزینه تراکنش به ریال برای اپلیکیشن شما | [optional] 
**Id** | Pointer to **string** | همان uuid هنگام ایجاد تراکنش | [optional] 

## Methods

### NewPaymentReorderPostResponse

`func NewPaymentReorderPostResponse() *PaymentReorderPostResponse`

NewPaymentReorderPostResponse instantiates a new PaymentReorderPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReorderPostResponseWithDefaults

`func NewPaymentReorderPostResponseWithDefaults() *PaymentReorderPostResponse`

NewPaymentReorderPostResponseWithDefaults instantiates a new PaymentReorderPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostRials

`func (o *PaymentReorderPostResponse) GetCostRials() string`

GetCostRials returns the CostRials field if non-nil, zero value otherwise.

### GetCostRialsOk

`func (o *PaymentReorderPostResponse) GetCostRialsOk() (*string, bool)`

GetCostRialsOk returns a tuple with the CostRials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostRials

`func (o *PaymentReorderPostResponse) SetCostRials(v string)`

SetCostRials sets CostRials field to given value.

### HasCostRials

`func (o *PaymentReorderPostResponse) HasCostRials() bool`

HasCostRials returns a boolean if a field has been set.

### GetId

`func (o *PaymentReorderPostResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentReorderPostResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentReorderPostResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentReorderPostResponse) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


