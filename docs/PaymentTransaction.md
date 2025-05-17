# PaymentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Same uuid when creating transaction | [optional] 
**CostRials** | Pointer to **string** | The cost of the transaction in rials for your application | [optional] 
**ExtraDetails** | Pointer to **string** | Same Additional details which you sent in the request | [optional] 
**State** | Pointer to [**PaymentTransactionState**](PaymentTransactionState.md) |  | [optional] 
**Type** | Pointer to [**PaymentTransactionType**](PaymentTransactionType.md) |  | [optional] 

## Methods

### NewPaymentTransaction

`func NewPaymentTransaction() *PaymentTransaction`

NewPaymentTransaction instantiates a new PaymentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTransactionWithDefaults

`func NewPaymentTransactionWithDefaults() *PaymentTransaction`

NewPaymentTransactionWithDefaults instantiates a new PaymentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCostRials

`func (o *PaymentTransaction) GetCostRials() string`

GetCostRials returns the CostRials field if non-nil, zero value otherwise.

### GetCostRialsOk

`func (o *PaymentTransaction) GetCostRialsOk() (*string, bool)`

GetCostRialsOk returns a tuple with the CostRials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostRials

`func (o *PaymentTransaction) SetCostRials(v string)`

SetCostRials sets CostRials field to given value.

### HasCostRials

`func (o *PaymentTransaction) HasCostRials() bool`

HasCostRials returns a boolean if a field has been set.

### GetExtraDetails

`func (o *PaymentTransaction) GetExtraDetails() string`

GetExtraDetails returns the ExtraDetails field if non-nil, zero value otherwise.

### GetExtraDetailsOk

`func (o *PaymentTransaction) GetExtraDetailsOk() (*string, bool)`

GetExtraDetailsOk returns a tuple with the ExtraDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDetails

`func (o *PaymentTransaction) SetExtraDetails(v string)`

SetExtraDetails sets ExtraDetails field to given value.

### HasExtraDetails

`func (o *PaymentTransaction) HasExtraDetails() bool`

HasExtraDetails returns a boolean if a field has been set.

### GetState

`func (o *PaymentTransaction) GetState() PaymentTransactionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PaymentTransaction) GetStateOk() (*PaymentTransactionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PaymentTransaction) SetState(v PaymentTransactionState)`

SetState sets State field to given value.

### HasState

`func (o *PaymentTransaction) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *PaymentTransaction) GetType() PaymentTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentTransaction) GetTypeOk() (*PaymentTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentTransaction) SetType(v PaymentTransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentTransaction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


