# FinderGetUserBusinessesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Businesses** | Pointer to [**[]FinderGetUserBusinessesResponseBusiness**](FinderGetUserBusinessesResponseBusiness.md) | فهرست کسب‌وکارهای مرتبط با کاربر احراز هویت شده | [optional] 

## Methods

### NewFinderGetUserBusinessesResponse

`func NewFinderGetUserBusinessesResponse() *FinderGetUserBusinessesResponse`

NewFinderGetUserBusinessesResponse instantiates a new FinderGetUserBusinessesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinderGetUserBusinessesResponseWithDefaults

`func NewFinderGetUserBusinessesResponseWithDefaults() *FinderGetUserBusinessesResponse`

NewFinderGetUserBusinessesResponseWithDefaults instantiates a new FinderGetUserBusinessesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinesses

`func (o *FinderGetUserBusinessesResponse) GetBusinesses() []FinderGetUserBusinessesResponseBusiness`

GetBusinesses returns the Businesses field if non-nil, zero value otherwise.

### GetBusinessesOk

`func (o *FinderGetUserBusinessesResponse) GetBusinessesOk() (*[]FinderGetUserBusinessesResponseBusiness, bool)`

GetBusinessesOk returns a tuple with the Businesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinesses

`func (o *FinderGetUserBusinessesResponse) SetBusinesses(v []FinderGetUserBusinessesResponseBusiness)`

SetBusinesses sets Businesses field to given value.

### HasBusinesses

`func (o *FinderGetUserBusinessesResponse) HasBusinesses() bool`

HasBusinesses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


