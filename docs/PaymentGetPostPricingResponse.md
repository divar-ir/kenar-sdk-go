# PaymentGetPostPricingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reorder** | Pointer to [**GetPostPricingResponseReorder**](GetPostPricingResponseReorder.md) |  | [optional] 
**Submit** | Pointer to [**GetPostPricingResponseSubmit**](GetPostPricingResponseSubmit.md) |  | [optional] 

## Methods

### NewPaymentGetPostPricingResponse

`func NewPaymentGetPostPricingResponse() *PaymentGetPostPricingResponse`

NewPaymentGetPostPricingResponse instantiates a new PaymentGetPostPricingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentGetPostPricingResponseWithDefaults

`func NewPaymentGetPostPricingResponseWithDefaults() *PaymentGetPostPricingResponse`

NewPaymentGetPostPricingResponseWithDefaults instantiates a new PaymentGetPostPricingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReorder

`func (o *PaymentGetPostPricingResponse) GetReorder() GetPostPricingResponseReorder`

GetReorder returns the Reorder field if non-nil, zero value otherwise.

### GetReorderOk

`func (o *PaymentGetPostPricingResponse) GetReorderOk() (*GetPostPricingResponseReorder, bool)`

GetReorderOk returns a tuple with the Reorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorder

`func (o *PaymentGetPostPricingResponse) SetReorder(v GetPostPricingResponseReorder)`

SetReorder sets Reorder field to given value.

### HasReorder

`func (o *PaymentGetPostPricingResponse) HasReorder() bool`

HasReorder returns a boolean if a field has been set.

### GetSubmit

`func (o *PaymentGetPostPricingResponse) GetSubmit() GetPostPricingResponseSubmit`

GetSubmit returns the Submit field if non-nil, zero value otherwise.

### GetSubmitOk

`func (o *PaymentGetPostPricingResponse) GetSubmitOk() (*GetPostPricingResponseSubmit, bool)`

GetSubmitOk returns a tuple with the Submit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmit

`func (o *PaymentGetPostPricingResponse) SetSubmit(v GetPostPricingResponseSubmit)`

SetSubmit sets Submit field to given value.

### HasSubmit

`func (o *PaymentGetPostPricingResponse) HasSubmit() bool`

HasSubmit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


