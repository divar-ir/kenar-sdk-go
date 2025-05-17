# GetPostPricingResponseReorder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostRials** | Pointer to **string** | The cost of reordering in rials | [optional] 
**Available** | Pointer to **bool** | Indicates if the post can be reordered. If false, the reorder API will return an error | [optional] 

## Methods

### NewGetPostPricingResponseReorder

`func NewGetPostPricingResponseReorder() *GetPostPricingResponseReorder`

NewGetPostPricingResponseReorder instantiates a new GetPostPricingResponseReorder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPostPricingResponseReorderWithDefaults

`func NewGetPostPricingResponseReorderWithDefaults() *GetPostPricingResponseReorder`

NewGetPostPricingResponseReorderWithDefaults instantiates a new GetPostPricingResponseReorder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostRials

`func (o *GetPostPricingResponseReorder) GetCostRials() string`

GetCostRials returns the CostRials field if non-nil, zero value otherwise.

### GetCostRialsOk

`func (o *GetPostPricingResponseReorder) GetCostRialsOk() (*string, bool)`

GetCostRialsOk returns a tuple with the CostRials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostRials

`func (o *GetPostPricingResponseReorder) SetCostRials(v string)`

SetCostRials sets CostRials field to given value.

### HasCostRials

`func (o *GetPostPricingResponseReorder) HasCostRials() bool`

HasCostRials returns a boolean if a field has been set.

### GetAvailable

`func (o *GetPostPricingResponseReorder) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetPostPricingResponseReorder) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetPostPricingResponseReorder) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *GetPostPricingResponseReorder) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


