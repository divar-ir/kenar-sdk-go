# PaymentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppSlug** | Pointer to **string** | شناسه اپلیکیشنی که تراکنش را ایجاد کرده است | [optional] 
**CostRials** | Pointer to **string** | هزینه تراکنش به ریال برای اپلیکیشن شما | [optional] 
**CreatedAt** | Pointer to **time.Time** | زمان ایجاد تراکنش | [optional] 
**ExtraDetails** | Pointer to **string** | همان جزئیات اضافی که در درخواست ارسال کردید | [optional] 
**Id** | Pointer to **string** | همان uuid هنگام ایجاد تراکنش | [optional] 
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

### GetAppSlug

`func (o *PaymentTransaction) GetAppSlug() string`

GetAppSlug returns the AppSlug field if non-nil, zero value otherwise.

### GetAppSlugOk

`func (o *PaymentTransaction) GetAppSlugOk() (*string, bool)`

GetAppSlugOk returns a tuple with the AppSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSlug

`func (o *PaymentTransaction) SetAppSlug(v string)`

SetAppSlug sets AppSlug field to given value.

### HasAppSlug

`func (o *PaymentTransaction) HasAppSlug() bool`

HasAppSlug returns a boolean if a field has been set.

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

### GetCreatedAt

`func (o *PaymentTransaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentTransaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentTransaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentTransaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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


