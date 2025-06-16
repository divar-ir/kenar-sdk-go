# \AddonsAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddonsCreateBusinessAddon**](AddonsAPI.md#AddonsCreateBusinessAddon) | **Post** /v1/open-platform/addons/business/{business_token} | افزودن افزونه جدید به آگهی‌های کسب‌و‌کار
[**AddonsCreatePostAddonV2**](AddonsAPI.md#AddonsCreatePostAddonV2) | **Post** /v2/open-platform/addons/post/{token} | افزودن افزونه جدید به آگهی
[**AddonsCreateUserAddonV2**](AddonsAPI.md#AddonsCreateUserAddonV2) | **Post** /v2/open-platform/addons/user/{phone} | افزودن افزونه جدید به کاربر
[**AddonsCreateUserAddonV22**](AddonsAPI.md#AddonsCreateUserAddonV22) | **Post** /v2/open-platform/addons/users/{divar_user_id} | افزودن افزونه جدید به کاربر
[**AddonsDeletePostAddon**](AddonsAPI.md#AddonsDeletePostAddon) | **Delete** /v1/open-platform/add-ons/post/{token} | حذف افزونه از آگهی
[**AddonsDeletePostAddon2**](AddonsAPI.md#AddonsDeletePostAddon2) | **Delete** /v1/open-platform/addons/post/{token} | حذف افزونه از آگهی
[**AddonsDeleteUserAddon**](AddonsAPI.md#AddonsDeleteUserAddon) | **Delete** /v1/open-platform/addons/user/{id} | حذف افزونه کاربر
[**AddonsGetUserAddons**](AddonsAPI.md#AddonsGetUserAddons) | **Get** /v1/open-platform/addons/user/{phone} | دریافت تمام افزونه‌های کاربر
[**AddonsGetUserAddons2**](AddonsAPI.md#AddonsGetUserAddons2) | **Get** /v2/open-platform/addons/users/{divar_user_id} | دریافت تمام افزونه‌های کاربر



## AddonsCreateBusinessAddon

> AddonsCreateBusinessAddonResponse AddonsCreateBusinessAddon(ctx, businessToken).AddonsCreateBusinessAddonBody(addonsCreateBusinessAddonBody).Execute()

افزودن افزونه جدید به آگهی‌های کسب‌و‌کار



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
)

func main() {
	businessToken := "businessToken_example" // string | 
	addonsCreateBusinessAddonBody := *openapiclient.NewAddonsCreateBusinessAddonBody() // AddonsCreateBusinessAddonBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsCreateBusinessAddon(context.Background(), businessToken).AddonsCreateBusinessAddonBody(addonsCreateBusinessAddonBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsCreateBusinessAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsCreateBusinessAddon`: AddonsCreateBusinessAddonResponse
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsCreateBusinessAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsCreateBusinessAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addonsCreateBusinessAddonBody** | [**AddonsCreateBusinessAddonBody**](AddonsCreateBusinessAddonBody.md) |  | 

### Return type

[**AddonsCreateBusinessAddonResponse**](AddonsCreateBusinessAddonResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsCreatePostAddonV2

> map[string]interface{} AddonsCreatePostAddonV2(ctx, token).AddonsCreatePostAddonV2Body(addonsCreatePostAddonV2Body).Execute()

افزودن افزونه جدید به آگهی



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
)

func main() {
	token := "token_example" // string | 
	addonsCreatePostAddonV2Body := *openapiclient.NewAddonsCreatePostAddonV2Body() // AddonsCreatePostAddonV2Body | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsCreatePostAddonV2(context.Background(), token).AddonsCreatePostAddonV2Body(addonsCreatePostAddonV2Body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsCreatePostAddonV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsCreatePostAddonV2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsCreatePostAddonV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsCreatePostAddonV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addonsCreatePostAddonV2Body** | [**AddonsCreatePostAddonV2Body**](AddonsCreatePostAddonV2Body.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsCreateUserAddonV2

> AddonsCreateUserAddonResponseV2 AddonsCreateUserAddonV2(ctx, phone).AddonsCreateUserAddonV2Body(addonsCreateUserAddonV2Body).Execute()

افزودن افزونه جدید به کاربر



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
)

func main() {
	phone := "phone_example" // string | 
	addonsCreateUserAddonV2Body := *openapiclient.NewAddonsCreateUserAddonV2Body() // AddonsCreateUserAddonV2Body | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsCreateUserAddonV2(context.Background(), phone).AddonsCreateUserAddonV2Body(addonsCreateUserAddonV2Body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsCreateUserAddonV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsCreateUserAddonV2`: AddonsCreateUserAddonResponseV2
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsCreateUserAddonV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsCreateUserAddonV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addonsCreateUserAddonV2Body** | [**AddonsCreateUserAddonV2Body**](AddonsCreateUserAddonV2Body.md) |  | 

### Return type

[**AddonsCreateUserAddonResponseV2**](AddonsCreateUserAddonResponseV2.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsCreateUserAddonV22

> AddonsCreateUserAddonResponseV2 AddonsCreateUserAddonV22(ctx, divarUserId).AddonsCreateUserAddonV2Body(addonsCreateUserAddonV2Body).Execute()

افزودن افزونه جدید به کاربر



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
)

func main() {
	divarUserId := "divarUserId_example" // string | 
	addonsCreateUserAddonV2Body := *openapiclient.NewAddonsCreateUserAddonV2Body() // AddonsCreateUserAddonV2Body | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsCreateUserAddonV22(context.Background(), divarUserId).AddonsCreateUserAddonV2Body(addonsCreateUserAddonV2Body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsCreateUserAddonV22``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsCreateUserAddonV22`: AddonsCreateUserAddonResponseV2
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsCreateUserAddonV22`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**divarUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsCreateUserAddonV22Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addonsCreateUserAddonV2Body** | [**AddonsCreateUserAddonV2Body**](AddonsCreateUserAddonV2Body.md) |  | 

### Return type

[**AddonsCreateUserAddonResponseV2**](AddonsCreateUserAddonResponseV2.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsDeletePostAddon

> map[string]interface{} AddonsDeletePostAddon(ctx, token).Execute()

حذف افزونه از آگهی



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
)

func main() {
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsDeletePostAddon(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsDeletePostAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsDeletePostAddon`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsDeletePostAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsDeletePostAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsDeletePostAddon2

> map[string]interface{} AddonsDeletePostAddon2(ctx, token).Execute()

حذف افزونه از آگهی



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
)

func main() {
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsDeletePostAddon2(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsDeletePostAddon2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsDeletePostAddon2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsDeletePostAddon2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsDeletePostAddon2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsDeleteUserAddon

> map[string]interface{} AddonsDeleteUserAddon(ctx, id).Execute()

حذف افزونه کاربر



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsDeleteUserAddon(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsDeleteUserAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsDeleteUserAddon`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsDeleteUserAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsDeleteUserAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsGetUserAddons

> AddonsGetUserAddonsResponse AddonsGetUserAddons(ctx, phone).DivarUserId(divarUserId).Execute()

دریافت تمام افزونه‌های کاربر



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
)

func main() {
	phone := "phone_example" // string | 
	divarUserId := "divarUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsGetUserAddons(context.Background(), phone).DivarUserId(divarUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsGetUserAddons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsGetUserAddons`: AddonsGetUserAddonsResponse
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsGetUserAddons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsGetUserAddonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **divarUserId** | **string** |  | 

### Return type

[**AddonsGetUserAddonsResponse**](AddonsGetUserAddonsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsGetUserAddons2

> AddonsGetUserAddonsResponse AddonsGetUserAddons2(ctx, divarUserId).Phone(phone).Execute()

دریافت تمام افزونه‌های کاربر



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
)

func main() {
	divarUserId := "divarUserId_example" // string | 
	phone := "phone_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsGetUserAddons2(context.Background(), divarUserId).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsGetUserAddons2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsGetUserAddons2`: AddonsGetUserAddonsResponse
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsGetUserAddons2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**divarUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsGetUserAddons2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phone** | **string** |  | 

### Return type

[**AddonsGetUserAddonsResponse**](AddonsGetUserAddonsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

