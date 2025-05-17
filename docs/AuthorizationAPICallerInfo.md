# AuthorizationAPICallerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **int32** |  | [optional] 
**App** | Pointer to [**AppsApp**](AppsApp.md) |  | [optional] 
**Scopes** | Pointer to [**[]AuthorizationOAuthScope**](AuthorizationOAuthScope.md) |  | [optional] 
**ApiKeyIdV2** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthorizationAPICallerInfo

`func NewAuthorizationAPICallerInfo() *AuthorizationAPICallerInfo`

NewAuthorizationAPICallerInfo instantiates a new AuthorizationAPICallerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationAPICallerInfoWithDefaults

`func NewAuthorizationAPICallerInfoWithDefaults() *AuthorizationAPICallerInfo`

NewAuthorizationAPICallerInfoWithDefaults instantiates a new AuthorizationAPICallerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *AuthorizationAPICallerInfo) GetApiKeyId() int32`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *AuthorizationAPICallerInfo) GetApiKeyIdOk() (*int32, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *AuthorizationAPICallerInfo) SetApiKeyId(v int32)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *AuthorizationAPICallerInfo) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetApp

`func (o *AuthorizationAPICallerInfo) GetApp() AppsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AuthorizationAPICallerInfo) GetAppOk() (*AppsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AuthorizationAPICallerInfo) SetApp(v AppsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *AuthorizationAPICallerInfo) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetScopes

`func (o *AuthorizationAPICallerInfo) GetScopes() []AuthorizationOAuthScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizationAPICallerInfo) GetScopesOk() (*[]AuthorizationOAuthScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizationAPICallerInfo) SetScopes(v []AuthorizationOAuthScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizationAPICallerInfo) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetApiKeyIdV2

`func (o *AuthorizationAPICallerInfo) GetApiKeyIdV2() string`

GetApiKeyIdV2 returns the ApiKeyIdV2 field if non-nil, zero value otherwise.

### GetApiKeyIdV2Ok

`func (o *AuthorizationAPICallerInfo) GetApiKeyIdV2Ok() (*string, bool)`

GetApiKeyIdV2Ok returns a tuple with the ApiKeyIdV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyIdV2

`func (o *AuthorizationAPICallerInfo) SetApiKeyIdV2(v string)`

SetApiKeyIdV2 sets ApiKeyIdV2 field to given value.

### HasApiKeyIdV2

`func (o *AuthorizationAPICallerInfo) HasApiKeyIdV2() bool`

HasApiKeyIdV2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


