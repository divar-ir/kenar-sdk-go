# PaymentRenewPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraDetails** | Pointer to **string** | اطلاعات تکمیلی که می‌خواهید به کنار ارسال کنید. این فیلد اختیاری است و برای رفع مغایرت‌های احتمالی در تراکنش کاربرد دارد. | [optional] 
**Id** | Pointer to **string** | یک UUID نسخه ۴ که باید برای هر پرداخت یکتا باشد. این UUID باید در سمت شما ساخته شده و در درخواست ارسال شود. اگر id ارسال‌شده قبلاً تراکنش موفق یا نیمه‌موفقی در کنار داشته باشد، خطا دریافت می‌کنید. | [optional] 

## Methods

### NewPaymentRenewPostBody

`func NewPaymentRenewPostBody() *PaymentRenewPostBody`

NewPaymentRenewPostBody instantiates a new PaymentRenewPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRenewPostBodyWithDefaults

`func NewPaymentRenewPostBodyWithDefaults() *PaymentRenewPostBody`

NewPaymentRenewPostBodyWithDefaults instantiates a new PaymentRenewPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraDetails

`func (o *PaymentRenewPostBody) GetExtraDetails() string`

GetExtraDetails returns the ExtraDetails field if non-nil, zero value otherwise.

### GetExtraDetailsOk

`func (o *PaymentRenewPostBody) GetExtraDetailsOk() (*string, bool)`

GetExtraDetailsOk returns a tuple with the ExtraDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDetails

`func (o *PaymentRenewPostBody) SetExtraDetails(v string)`

SetExtraDetails sets ExtraDetails field to given value.

### HasExtraDetails

`func (o *PaymentRenewPostBody) HasExtraDetails() bool`

HasExtraDetails returns a boolean if a field has been set.

### GetId

`func (o *PaymentRenewPostBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRenewPostBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRenewPostBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentRenewPostBody) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


