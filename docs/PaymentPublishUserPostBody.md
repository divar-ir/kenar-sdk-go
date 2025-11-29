# PaymentPublishUserPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraDetails** | Pointer to **string** | جزئیات اضافی که می‌خواهید به سمت کنار ارسال کنید. این فیلد اختیاری است و می‌تواند برای حل ناسازگاری‌ها در تراکنش استفاده شود. | [optional] 
**Id** | Pointer to **string** | یک uuid نسخه 4 که باید برای هر پرداخت منحصر به فرد باشد. این uuid باید در سمت شما تولید شده و در درخواست ارسال شود. اگر id ای ارسال شود که تراکنش موفق یا نیمه موفقی در سمت کنار داشته باشد، خطا دریافت خواهید کرد. | [optional] 

## Methods

### NewPaymentPublishUserPostBody

`func NewPaymentPublishUserPostBody() *PaymentPublishUserPostBody`

NewPaymentPublishUserPostBody instantiates a new PaymentPublishUserPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPublishUserPostBodyWithDefaults

`func NewPaymentPublishUserPostBodyWithDefaults() *PaymentPublishUserPostBody`

NewPaymentPublishUserPostBodyWithDefaults instantiates a new PaymentPublishUserPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraDetails

`func (o *PaymentPublishUserPostBody) GetExtraDetails() string`

GetExtraDetails returns the ExtraDetails field if non-nil, zero value otherwise.

### GetExtraDetailsOk

`func (o *PaymentPublishUserPostBody) GetExtraDetailsOk() (*string, bool)`

GetExtraDetailsOk returns a tuple with the ExtraDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDetails

`func (o *PaymentPublishUserPostBody) SetExtraDetails(v string)`

SetExtraDetails sets ExtraDetails field to given value.

### HasExtraDetails

`func (o *PaymentPublishUserPostBody) HasExtraDetails() bool`

HasExtraDetails returns a boolean if a field has been set.

### GetId

`func (o *PaymentPublishUserPostBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentPublishUserPostBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentPublishUserPostBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentPublishUserPostBody) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


