# PaymentCreateWalletPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountRials** | **string** | مبلغ تراکنش به ریال | 
**Description** | **string** | توضیحات تراکنش | 
**RedirectUrl** | **string** | آدرسی که کاربر باید بعد از پرداخت مبلغ به آن هدایت شود | 

## Methods

### NewPaymentCreateWalletPaymentRequest

`func NewPaymentCreateWalletPaymentRequest(amountRials string, description string, redirectUrl string, ) *PaymentCreateWalletPaymentRequest`

NewPaymentCreateWalletPaymentRequest instantiates a new PaymentCreateWalletPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCreateWalletPaymentRequestWithDefaults

`func NewPaymentCreateWalletPaymentRequestWithDefaults() *PaymentCreateWalletPaymentRequest`

NewPaymentCreateWalletPaymentRequestWithDefaults instantiates a new PaymentCreateWalletPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountRials

`func (o *PaymentCreateWalletPaymentRequest) GetAmountRials() string`

GetAmountRials returns the AmountRials field if non-nil, zero value otherwise.

### GetAmountRialsOk

`func (o *PaymentCreateWalletPaymentRequest) GetAmountRialsOk() (*string, bool)`

GetAmountRialsOk returns a tuple with the AmountRials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRials

`func (o *PaymentCreateWalletPaymentRequest) SetAmountRials(v string)`

SetAmountRials sets AmountRials field to given value.


### GetDescription

`func (o *PaymentCreateWalletPaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentCreateWalletPaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentCreateWalletPaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRedirectUrl

`func (o *PaymentCreateWalletPaymentRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PaymentCreateWalletPaymentRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PaymentCreateWalletPaymentRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


