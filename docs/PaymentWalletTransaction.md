# PaymentWalletTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountRials** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PaymentWalletTransactionStatus**](PaymentWalletTransactionStatus.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentWalletTransaction

`func NewPaymentWalletTransaction() *PaymentWalletTransaction`

NewPaymentWalletTransaction instantiates a new PaymentWalletTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWalletTransactionWithDefaults

`func NewPaymentWalletTransactionWithDefaults() *PaymentWalletTransaction`

NewPaymentWalletTransactionWithDefaults instantiates a new PaymentWalletTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountRials

`func (o *PaymentWalletTransaction) GetAmountRials() string`

GetAmountRials returns the AmountRials field if non-nil, zero value otherwise.

### GetAmountRialsOk

`func (o *PaymentWalletTransaction) GetAmountRialsOk() (*string, bool)`

GetAmountRialsOk returns a tuple with the AmountRials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRials

`func (o *PaymentWalletTransaction) SetAmountRials(v string)`

SetAmountRials sets AmountRials field to given value.

### HasAmountRials

`func (o *PaymentWalletTransaction) HasAmountRials() bool`

HasAmountRials returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentWalletTransaction) GetStatus() PaymentWalletTransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentWalletTransaction) GetStatusOk() (*PaymentWalletTransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentWalletTransaction) SetStatus(v PaymentWalletTransactionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentWalletTransaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *PaymentWalletTransaction) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentWalletTransaction) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentWalletTransaction) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentWalletTransaction) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


