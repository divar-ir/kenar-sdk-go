# \SemanticAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SemanticCreatePostSemantic**](SemanticAPI.md#SemanticCreatePostSemantic) | **Post** /experimental/open-platform/semantic/post/{token} | ایجاد اطلاعات معنایی آگهی
[**SemanticCreateUserSemantic**](SemanticAPI.md#SemanticCreateUserSemantic) | **Post** /v1/open-platform/semantic/user/{phone} | ایجاد اطلاعات معنایی کاربر
[**SemanticCreateUserSemantic2**](SemanticAPI.md#SemanticCreateUserSemantic2) | **Post** /v1/open-platform/semantic/users/{divar_user_id} | ایجاد اطلاعات معنایی کاربر
[**SemanticDeleteUserSemantic**](SemanticAPI.md#SemanticDeleteUserSemantic) | **Delete** /v1/open-platform/semantic/user/{phone} | حذف اطلاعات معنایی کاربر
[**SemanticDeleteUserSemantic2**](SemanticAPI.md#SemanticDeleteUserSemantic2) | **Delete** /v1/open-platform/semantic/users/{divar_user_id} | حذف اطلاعات معنایی کاربر



## SemanticCreatePostSemantic

> map[string]interface{} SemanticCreatePostSemantic(ctx, token).SemanticCreatePostSemanticBody(semanticCreatePostSemanticBody).Execute()

ایجاد اطلاعات معنایی آگهی



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/divar-ir/kenar-sdk-go"
)

func main() {
	token := "token_example" // string | 
	semanticCreatePostSemanticBody := *openapiclient.NewSemanticCreatePostSemanticBody() // SemanticCreatePostSemanticBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SemanticAPI.SemanticCreatePostSemantic(context.Background(), token).SemanticCreatePostSemanticBody(semanticCreatePostSemanticBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SemanticAPI.SemanticCreatePostSemantic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SemanticCreatePostSemantic`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SemanticAPI.SemanticCreatePostSemantic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSemanticCreatePostSemanticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **semanticCreatePostSemanticBody** | [**SemanticCreatePostSemanticBody**](SemanticCreatePostSemanticBody.md) |  | 

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


## SemanticCreateUserSemantic

> SemanticCreateUserSemanticResponse SemanticCreateUserSemantic(ctx, phone).SemanticCreateUserSemanticBody(semanticCreateUserSemanticBody).Execute()

ایجاد اطلاعات معنایی کاربر



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/divar-ir/kenar-sdk-go"
)

func main() {
	phone := "phone_example" // string | 
	semanticCreateUserSemanticBody := *openapiclient.NewSemanticCreateUserSemanticBody() // SemanticCreateUserSemanticBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SemanticAPI.SemanticCreateUserSemantic(context.Background(), phone).SemanticCreateUserSemanticBody(semanticCreateUserSemanticBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SemanticAPI.SemanticCreateUserSemantic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SemanticCreateUserSemantic`: SemanticCreateUserSemanticResponse
	fmt.Fprintf(os.Stdout, "Response from `SemanticAPI.SemanticCreateUserSemantic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSemanticCreateUserSemanticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **semanticCreateUserSemanticBody** | [**SemanticCreateUserSemanticBody**](SemanticCreateUserSemanticBody.md) |  | 

### Return type

[**SemanticCreateUserSemanticResponse**](SemanticCreateUserSemanticResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SemanticCreateUserSemantic2

> SemanticCreateUserSemanticResponse SemanticCreateUserSemantic2(ctx, divarUserId).SemanticCreateUserSemanticBody(semanticCreateUserSemanticBody).Execute()

ایجاد اطلاعات معنایی کاربر



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/divar-ir/kenar-sdk-go"
)

func main() {
	divarUserId := "divarUserId_example" // string | 
	semanticCreateUserSemanticBody := *openapiclient.NewSemanticCreateUserSemanticBody() // SemanticCreateUserSemanticBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SemanticAPI.SemanticCreateUserSemantic2(context.Background(), divarUserId).SemanticCreateUserSemanticBody(semanticCreateUserSemanticBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SemanticAPI.SemanticCreateUserSemantic2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SemanticCreateUserSemantic2`: SemanticCreateUserSemanticResponse
	fmt.Fprintf(os.Stdout, "Response from `SemanticAPI.SemanticCreateUserSemantic2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**divarUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSemanticCreateUserSemantic2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **semanticCreateUserSemanticBody** | [**SemanticCreateUserSemanticBody**](SemanticCreateUserSemanticBody.md) |  | 

### Return type

[**SemanticCreateUserSemanticResponse**](SemanticCreateUserSemanticResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SemanticDeleteUserSemantic

> map[string]interface{} SemanticDeleteUserSemantic(ctx, phone).DivarUserId(divarUserId).Execute()

حذف اطلاعات معنایی کاربر



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/divar-ir/kenar-sdk-go"
)

func main() {
	phone := "phone_example" // string | 
	divarUserId := "divarUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SemanticAPI.SemanticDeleteUserSemantic(context.Background(), phone).DivarUserId(divarUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SemanticAPI.SemanticDeleteUserSemantic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SemanticDeleteUserSemantic`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SemanticAPI.SemanticDeleteUserSemantic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSemanticDeleteUserSemanticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **divarUserId** | **string** |  | 

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


## SemanticDeleteUserSemantic2

> map[string]interface{} SemanticDeleteUserSemantic2(ctx, divarUserId).Phone(phone).Execute()

حذف اطلاعات معنایی کاربر



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/divar-ir/kenar-sdk-go"
)

func main() {
	divarUserId := "divarUserId_example" // string | 
	phone := "phone_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SemanticAPI.SemanticDeleteUserSemantic2(context.Background(), divarUserId).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SemanticAPI.SemanticDeleteUserSemantic2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SemanticDeleteUserSemantic2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SemanticAPI.SemanticDeleteUserSemantic2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**divarUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSemanticDeleteUserSemantic2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phone** | **string** |  | 

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

