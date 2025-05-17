# PaymentGetTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | Pointer to [**PaymentTransaction**](PaymentTransaction.md) |  | [optional] 

## Methods

### NewPaymentGetTransactionResponse

`func NewPaymentGetTransactionResponse() *PaymentGetTransactionResponse`

NewPaymentGetTransactionResponse instantiates a new PaymentGetTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentGetTransactionResponseWithDefaults

`func NewPaymentGetTransactionResponseWithDefaults() *PaymentGetTransactionResponse`

NewPaymentGetTransactionResponseWithDefaults instantiates a new PaymentGetTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *PaymentGetTransactionResponse) GetTransaction() PaymentTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *PaymentGetTransactionResponse) GetTransactionOk() (*PaymentTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *PaymentGetTransactionResponse) SetTransaction(v PaymentTransaction)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *PaymentGetTransactionResponse) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


