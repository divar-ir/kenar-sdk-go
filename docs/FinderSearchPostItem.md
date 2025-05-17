# FinderSearchPostItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**SearchPostItemPrice**](SearchPostItemPrice.md) |  | [optional] 
**RealEstateFields** | Pointer to [**SearchPostItemRealEstateFields**](SearchPostItemRealEstateFields.md) |  | [optional] 
**VehiclesFields** | Pointer to [**SearchPostItemVehiclesFields**](SearchPostItemVehiclesFields.md) |  | [optional] 
**ElectronicDevicesFields** | Pointer to **map[string]interface{}** |  | [optional] 
**HomeKitchenFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ServicesFields** | Pointer to **map[string]interface{}** |  | [optional] 
**PersonalGoodsFields** | Pointer to **map[string]interface{}** |  | [optional] 
**LeisureHobbiesFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CommunityFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ToolsMaterialsEquipmentFields** | Pointer to **map[string]interface{}** |  | [optional] 
**JobsFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFinderSearchPostItem

`func NewFinderSearchPostItem() *FinderSearchPostItem`

NewFinderSearchPostItem instantiates a new FinderSearchPostItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinderSearchPostItemWithDefaults

`func NewFinderSearchPostItemWithDefaults() *FinderSearchPostItem`

NewFinderSearchPostItemWithDefaults instantiates a new FinderSearchPostItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *FinderSearchPostItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FinderSearchPostItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FinderSearchPostItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FinderSearchPostItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCategory

`func (o *FinderSearchPostItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FinderSearchPostItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FinderSearchPostItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FinderSearchPostItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *FinderSearchPostItem) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *FinderSearchPostItem) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *FinderSearchPostItem) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *FinderSearchPostItem) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetCity

`func (o *FinderSearchPostItem) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *FinderSearchPostItem) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *FinderSearchPostItem) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *FinderSearchPostItem) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetTitle

`func (o *FinderSearchPostItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FinderSearchPostItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FinderSearchPostItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FinderSearchPostItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPrice

`func (o *FinderSearchPostItem) GetPrice() SearchPostItemPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FinderSearchPostItem) GetPriceOk() (*SearchPostItemPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FinderSearchPostItem) SetPrice(v SearchPostItemPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *FinderSearchPostItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRealEstateFields

`func (o *FinderSearchPostItem) GetRealEstateFields() SearchPostItemRealEstateFields`

GetRealEstateFields returns the RealEstateFields field if non-nil, zero value otherwise.

### GetRealEstateFieldsOk

`func (o *FinderSearchPostItem) GetRealEstateFieldsOk() (*SearchPostItemRealEstateFields, bool)`

GetRealEstateFieldsOk returns a tuple with the RealEstateFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealEstateFields

`func (o *FinderSearchPostItem) SetRealEstateFields(v SearchPostItemRealEstateFields)`

SetRealEstateFields sets RealEstateFields field to given value.

### HasRealEstateFields

`func (o *FinderSearchPostItem) HasRealEstateFields() bool`

HasRealEstateFields returns a boolean if a field has been set.

### GetVehiclesFields

`func (o *FinderSearchPostItem) GetVehiclesFields() SearchPostItemVehiclesFields`

GetVehiclesFields returns the VehiclesFields field if non-nil, zero value otherwise.

### GetVehiclesFieldsOk

`func (o *FinderSearchPostItem) GetVehiclesFieldsOk() (*SearchPostItemVehiclesFields, bool)`

GetVehiclesFieldsOk returns a tuple with the VehiclesFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehiclesFields

`func (o *FinderSearchPostItem) SetVehiclesFields(v SearchPostItemVehiclesFields)`

SetVehiclesFields sets VehiclesFields field to given value.

### HasVehiclesFields

`func (o *FinderSearchPostItem) HasVehiclesFields() bool`

HasVehiclesFields returns a boolean if a field has been set.

### GetElectronicDevicesFields

`func (o *FinderSearchPostItem) GetElectronicDevicesFields() map[string]interface{}`

GetElectronicDevicesFields returns the ElectronicDevicesFields field if non-nil, zero value otherwise.

### GetElectronicDevicesFieldsOk

`func (o *FinderSearchPostItem) GetElectronicDevicesFieldsOk() (*map[string]interface{}, bool)`

GetElectronicDevicesFieldsOk returns a tuple with the ElectronicDevicesFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronicDevicesFields

`func (o *FinderSearchPostItem) SetElectronicDevicesFields(v map[string]interface{})`

SetElectronicDevicesFields sets ElectronicDevicesFields field to given value.

### HasElectronicDevicesFields

`func (o *FinderSearchPostItem) HasElectronicDevicesFields() bool`

HasElectronicDevicesFields returns a boolean if a field has been set.

### GetHomeKitchenFields

`func (o *FinderSearchPostItem) GetHomeKitchenFields() map[string]interface{}`

GetHomeKitchenFields returns the HomeKitchenFields field if non-nil, zero value otherwise.

### GetHomeKitchenFieldsOk

`func (o *FinderSearchPostItem) GetHomeKitchenFieldsOk() (*map[string]interface{}, bool)`

GetHomeKitchenFieldsOk returns a tuple with the HomeKitchenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeKitchenFields

`func (o *FinderSearchPostItem) SetHomeKitchenFields(v map[string]interface{})`

SetHomeKitchenFields sets HomeKitchenFields field to given value.

### HasHomeKitchenFields

`func (o *FinderSearchPostItem) HasHomeKitchenFields() bool`

HasHomeKitchenFields returns a boolean if a field has been set.

### GetServicesFields

`func (o *FinderSearchPostItem) GetServicesFields() map[string]interface{}`

GetServicesFields returns the ServicesFields field if non-nil, zero value otherwise.

### GetServicesFieldsOk

`func (o *FinderSearchPostItem) GetServicesFieldsOk() (*map[string]interface{}, bool)`

GetServicesFieldsOk returns a tuple with the ServicesFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesFields

`func (o *FinderSearchPostItem) SetServicesFields(v map[string]interface{})`

SetServicesFields sets ServicesFields field to given value.

### HasServicesFields

`func (o *FinderSearchPostItem) HasServicesFields() bool`

HasServicesFields returns a boolean if a field has been set.

### GetPersonalGoodsFields

`func (o *FinderSearchPostItem) GetPersonalGoodsFields() map[string]interface{}`

GetPersonalGoodsFields returns the PersonalGoodsFields field if non-nil, zero value otherwise.

### GetPersonalGoodsFieldsOk

`func (o *FinderSearchPostItem) GetPersonalGoodsFieldsOk() (*map[string]interface{}, bool)`

GetPersonalGoodsFieldsOk returns a tuple with the PersonalGoodsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalGoodsFields

`func (o *FinderSearchPostItem) SetPersonalGoodsFields(v map[string]interface{})`

SetPersonalGoodsFields sets PersonalGoodsFields field to given value.

### HasPersonalGoodsFields

`func (o *FinderSearchPostItem) HasPersonalGoodsFields() bool`

HasPersonalGoodsFields returns a boolean if a field has been set.

### GetLeisureHobbiesFields

`func (o *FinderSearchPostItem) GetLeisureHobbiesFields() map[string]interface{}`

GetLeisureHobbiesFields returns the LeisureHobbiesFields field if non-nil, zero value otherwise.

### GetLeisureHobbiesFieldsOk

`func (o *FinderSearchPostItem) GetLeisureHobbiesFieldsOk() (*map[string]interface{}, bool)`

GetLeisureHobbiesFieldsOk returns a tuple with the LeisureHobbiesFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeisureHobbiesFields

`func (o *FinderSearchPostItem) SetLeisureHobbiesFields(v map[string]interface{})`

SetLeisureHobbiesFields sets LeisureHobbiesFields field to given value.

### HasLeisureHobbiesFields

`func (o *FinderSearchPostItem) HasLeisureHobbiesFields() bool`

HasLeisureHobbiesFields returns a boolean if a field has been set.

### GetCommunityFields

`func (o *FinderSearchPostItem) GetCommunityFields() map[string]interface{}`

GetCommunityFields returns the CommunityFields field if non-nil, zero value otherwise.

### GetCommunityFieldsOk

`func (o *FinderSearchPostItem) GetCommunityFieldsOk() (*map[string]interface{}, bool)`

GetCommunityFieldsOk returns a tuple with the CommunityFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityFields

`func (o *FinderSearchPostItem) SetCommunityFields(v map[string]interface{})`

SetCommunityFields sets CommunityFields field to given value.

### HasCommunityFields

`func (o *FinderSearchPostItem) HasCommunityFields() bool`

HasCommunityFields returns a boolean if a field has been set.

### GetToolsMaterialsEquipmentFields

`func (o *FinderSearchPostItem) GetToolsMaterialsEquipmentFields() map[string]interface{}`

GetToolsMaterialsEquipmentFields returns the ToolsMaterialsEquipmentFields field if non-nil, zero value otherwise.

### GetToolsMaterialsEquipmentFieldsOk

`func (o *FinderSearchPostItem) GetToolsMaterialsEquipmentFieldsOk() (*map[string]interface{}, bool)`

GetToolsMaterialsEquipmentFieldsOk returns a tuple with the ToolsMaterialsEquipmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsMaterialsEquipmentFields

`func (o *FinderSearchPostItem) SetToolsMaterialsEquipmentFields(v map[string]interface{})`

SetToolsMaterialsEquipmentFields sets ToolsMaterialsEquipmentFields field to given value.

### HasToolsMaterialsEquipmentFields

`func (o *FinderSearchPostItem) HasToolsMaterialsEquipmentFields() bool`

HasToolsMaterialsEquipmentFields returns a boolean if a field has been set.

### GetJobsFields

`func (o *FinderSearchPostItem) GetJobsFields() map[string]interface{}`

GetJobsFields returns the JobsFields field if non-nil, zero value otherwise.

### GetJobsFieldsOk

`func (o *FinderSearchPostItem) GetJobsFieldsOk() (*map[string]interface{}, bool)`

GetJobsFieldsOk returns a tuple with the JobsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsFields

`func (o *FinderSearchPostItem) SetJobsFields(v map[string]interface{})`

SetJobsFields sets JobsFields field to given value.

### HasJobsFields

`func (o *FinderSearchPostItem) HasJobsFields() bool`

HasJobsFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


