# PaymentReorderPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraDetails** | Pointer to **string** | Additional details that you want to send to the Kenar side. This field is optional and can be used to solve inconsistencies in the transaction. | [optional] 
**Id** | Pointer to **string** | A Version 4 uuid that must be unique for each payment. This uuid must be generated on your side and sent in the request. If an id is sent that has a successful or semi-successful transaction on the Kenar side, you will receive an error. | [optional] 

## Methods

### NewPaymentReorderPostBody

`func NewPaymentReorderPostBody() *PaymentReorderPostBody`

NewPaymentReorderPostBody instantiates a new PaymentReorderPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReorderPostBodyWithDefaults

`func NewPaymentReorderPostBodyWithDefaults() *PaymentReorderPostBody`

NewPaymentReorderPostBodyWithDefaults instantiates a new PaymentReorderPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraDetails

`func (o *PaymentReorderPostBody) GetExtraDetails() string`

GetExtraDetails returns the ExtraDetails field if non-nil, zero value otherwise.

### GetExtraDetailsOk

`func (o *PaymentReorderPostBody) GetExtraDetailsOk() (*string, bool)`

GetExtraDetailsOk returns a tuple with the ExtraDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDetails

`func (o *PaymentReorderPostBody) SetExtraDetails(v string)`

SetExtraDetails sets ExtraDetails field to given value.

### HasExtraDetails

`func (o *PaymentReorderPostBody) HasExtraDetails() bool`

HasExtraDetails returns a boolean if a field has been set.

### GetId

`func (o *PaymentReorderPostBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentReorderPostBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentReorderPostBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentReorderPostBody) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


