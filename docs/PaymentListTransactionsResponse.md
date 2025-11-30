# PaymentListTransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | توکن برای صفحه بعدی نتایج. | [optional] 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | فهرست تراکنش‌هایی که با درخواست مطابقت دارند | [optional] 

## Methods

### NewPaymentListTransactionsResponse

`func NewPaymentListTransactionsResponse() *PaymentListTransactionsResponse`

NewPaymentListTransactionsResponse instantiates a new PaymentListTransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentListTransactionsResponseWithDefaults

`func NewPaymentListTransactionsResponseWithDefaults() *PaymentListTransactionsResponse`

NewPaymentListTransactionsResponseWithDefaults instantiates a new PaymentListTransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *PaymentListTransactionsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *PaymentListTransactionsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *PaymentListTransactionsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *PaymentListTransactionsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTransactions

`func (o *PaymentListTransactionsResponse) GetTransactions() []PaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PaymentListTransactionsResponse) GetTransactionsOk() (*[]PaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PaymentListTransactionsResponse) SetTransactions(v []PaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *PaymentListTransactionsResponse) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


