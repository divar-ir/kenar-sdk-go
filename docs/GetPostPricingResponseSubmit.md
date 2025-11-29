# GetPostPricingResponseSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | نشان می‌دهد که آیا آگهی قابل ثبت است. اگر false باشد، API ثبت خطا برمی‌گرداند | [optional] 
**CostRials** | Pointer to **string** | هزینه ثبت آگهی به ریال | [optional] 

## Methods

### NewGetPostPricingResponseSubmit

`func NewGetPostPricingResponseSubmit() *GetPostPricingResponseSubmit`

NewGetPostPricingResponseSubmit instantiates a new GetPostPricingResponseSubmit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPostPricingResponseSubmitWithDefaults

`func NewGetPostPricingResponseSubmitWithDefaults() *GetPostPricingResponseSubmit`

NewGetPostPricingResponseSubmitWithDefaults instantiates a new GetPostPricingResponseSubmit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *GetPostPricingResponseSubmit) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetPostPricingResponseSubmit) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetPostPricingResponseSubmit) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *GetPostPricingResponseSubmit) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCostRials

`func (o *GetPostPricingResponseSubmit) GetCostRials() string`

GetCostRials returns the CostRials field if non-nil, zero value otherwise.

### GetCostRialsOk

`func (o *GetPostPricingResponseSubmit) GetCostRialsOk() (*string, bool)`

GetCostRialsOk returns a tuple with the CostRials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostRials

`func (o *GetPostPricingResponseSubmit) SetCostRials(v string)`

SetCostRials sets CostRials field to given value.

### HasCostRials

`func (o *GetPostPricingResponseSubmit) HasCostRials() bool`

HasCostRials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


