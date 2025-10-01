# PaymentSubmitUserPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountRials** | **string** | کل مبلغ پرداختی توسط کاربر، به ریال | 
**ProfitRials** | **string** | بخشی از مبلغ پرداختی که به شما تعلق می‌گیرد (سود یا کمیسیون)، به ریال | 
**ReferenceId** | **string** | شناسه مرجع فاکتور یا تراکنش | 
**Services** | **[]string** | لیست شناسه خدماتی که کاربر برای آنها پرداخت انجام داده است (مثلاً «banner»، «title_refinement» و ...). توصیه می‌شود از نام‌های انگلیسی کوتاه و توصیفی به‌عنوان شناسه خدمت استفاده شود. | 

## Methods

### NewPaymentSubmitUserPaymentRequest

`func NewPaymentSubmitUserPaymentRequest(amountRials string, profitRials string, referenceId string, services []string, ) *PaymentSubmitUserPaymentRequest`

NewPaymentSubmitUserPaymentRequest instantiates a new PaymentSubmitUserPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSubmitUserPaymentRequestWithDefaults

`func NewPaymentSubmitUserPaymentRequestWithDefaults() *PaymentSubmitUserPaymentRequest`

NewPaymentSubmitUserPaymentRequestWithDefaults instantiates a new PaymentSubmitUserPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountRials

`func (o *PaymentSubmitUserPaymentRequest) GetAmountRials() string`

GetAmountRials returns the AmountRials field if non-nil, zero value otherwise.

### GetAmountRialsOk

`func (o *PaymentSubmitUserPaymentRequest) GetAmountRialsOk() (*string, bool)`

GetAmountRialsOk returns a tuple with the AmountRials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRials

`func (o *PaymentSubmitUserPaymentRequest) SetAmountRials(v string)`

SetAmountRials sets AmountRials field to given value.


### GetProfitRials

`func (o *PaymentSubmitUserPaymentRequest) GetProfitRials() string`

GetProfitRials returns the ProfitRials field if non-nil, zero value otherwise.

### GetProfitRialsOk

`func (o *PaymentSubmitUserPaymentRequest) GetProfitRialsOk() (*string, bool)`

GetProfitRialsOk returns a tuple with the ProfitRials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitRials

`func (o *PaymentSubmitUserPaymentRequest) SetProfitRials(v string)`

SetProfitRials sets ProfitRials field to given value.


### GetReferenceId

`func (o *PaymentSubmitUserPaymentRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentSubmitUserPaymentRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentSubmitUserPaymentRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetServices

`func (o *PaymentSubmitUserPaymentRequest) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *PaymentSubmitUserPaymentRequest) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *PaymentSubmitUserPaymentRequest) SetServices(v []string)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


