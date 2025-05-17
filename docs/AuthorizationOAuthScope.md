# AuthorizationOAuthScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to [**OAuthScopeScope**](OAuthScopeScope.md) |  | [optional] 

## Methods

### NewAuthorizationOAuthScope

`func NewAuthorizationOAuthScope() *AuthorizationOAuthScope`

NewAuthorizationOAuthScope instantiates a new AuthorizationOAuthScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationOAuthScopeWithDefaults

`func NewAuthorizationOAuthScopeWithDefaults() *AuthorizationOAuthScope`

NewAuthorizationOAuthScopeWithDefaults instantiates a new AuthorizationOAuthScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *AuthorizationOAuthScope) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AuthorizationOAuthScope) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AuthorizationOAuthScope) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AuthorizationOAuthScope) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetScope

`func (o *AuthorizationOAuthScope) GetScope() OAuthScopeScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AuthorizationOAuthScope) GetScopeOk() (*OAuthScopeScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AuthorizationOAuthScope) SetScope(v OAuthScopeScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AuthorizationOAuthScope) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


