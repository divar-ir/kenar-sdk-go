# GetPostResponseBusinessData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessName** | Pointer to **string** |  | [optional] 
**BusinessType** | Pointer to [**PremiumPanelBusinessDataSubBusinessType**](PremiumPanelBusinessDataSubBusinessType.md) |  | [optional] 

## Methods

### NewGetPostResponseBusinessData

`func NewGetPostResponseBusinessData() *GetPostResponseBusinessData`

NewGetPostResponseBusinessData instantiates a new GetPostResponseBusinessData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPostResponseBusinessDataWithDefaults

`func NewGetPostResponseBusinessDataWithDefaults() *GetPostResponseBusinessData`

NewGetPostResponseBusinessDataWithDefaults instantiates a new GetPostResponseBusinessData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessName

`func (o *GetPostResponseBusinessData) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *GetPostResponseBusinessData) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *GetPostResponseBusinessData) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *GetPostResponseBusinessData) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetBusinessType

`func (o *GetPostResponseBusinessData) GetBusinessType() PremiumPanelBusinessDataSubBusinessType`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *GetPostResponseBusinessData) GetBusinessTypeOk() (*PremiumPanelBusinessDataSubBusinessType, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *GetPostResponseBusinessData) SetBusinessType(v PremiumPanelBusinessDataSubBusinessType)`

SetBusinessType sets BusinessType field to given value.

### HasBusinessType

`func (o *GetPostResponseBusinessData) HasBusinessType() bool`

HasBusinessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


