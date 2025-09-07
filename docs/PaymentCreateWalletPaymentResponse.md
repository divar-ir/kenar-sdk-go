# PaymentCreateWalletPaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentUrl** | Pointer to **string** | آدرسی که باید کاربر را به آن هدایت کنید | [optional] 
**Token** | Pointer to **string** | توکن تراکنش. برای عملیات‌های بعدی استفاده می‌شود | [optional] 

## Methods

### NewPaymentCreateWalletPaymentResponse

`func NewPaymentCreateWalletPaymentResponse() *PaymentCreateWalletPaymentResponse`

NewPaymentCreateWalletPaymentResponse instantiates a new PaymentCreateWalletPaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCreateWalletPaymentResponseWithDefaults

`func NewPaymentCreateWalletPaymentResponseWithDefaults() *PaymentCreateWalletPaymentResponse`

NewPaymentCreateWalletPaymentResponseWithDefaults instantiates a new PaymentCreateWalletPaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentUrl

`func (o *PaymentCreateWalletPaymentResponse) GetPaymentUrl() string`

GetPaymentUrl returns the PaymentUrl field if non-nil, zero value otherwise.

### GetPaymentUrlOk

`func (o *PaymentCreateWalletPaymentResponse) GetPaymentUrlOk() (*string, bool)`

GetPaymentUrlOk returns a tuple with the PaymentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUrl

`func (o *PaymentCreateWalletPaymentResponse) SetPaymentUrl(v string)`

SetPaymentUrl sets PaymentUrl field to given value.

### HasPaymentUrl

`func (o *PaymentCreateWalletPaymentResponse) HasPaymentUrl() bool`

HasPaymentUrl returns a boolean if a field has been set.

### GetToken

`func (o *PaymentCreateWalletPaymentResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentCreateWalletPaymentResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentCreateWalletPaymentResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentCreateWalletPaymentResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


