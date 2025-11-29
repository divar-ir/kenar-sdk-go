# \FinderAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FinderGetPost**](FinderAPI.md#FinderGetPost) | **Get** /v1/open-platform/finder/post/{token} | دریافت آگهی دیوار
[**FinderGetUser**](FinderAPI.md#FinderGetUser) | **Post** /v1/open-platform/users | دریافت اطلاعات کاربر
[**FinderGetUser2**](FinderAPI.md#FinderGetUser2) | **Get** /v1/open-platform/users | دریافت اطلاعات کاربر
[**FinderGetUserIDByPhone**](FinderAPI.md#FinderGetUserIDByPhone) | **Post** /v1/open-platform/get-user-id-by-phone | دریافت شناسه دیوار کاربر با شماره تلفن
[**FinderGetUserPosts**](FinderAPI.md#FinderGetUserPosts) | **Get** /v1/open-platform/finder/user-posts | دریافت آگهی‌های کاربر
[**FinderSearchPostV2**](FinderAPI.md#FinderSearchPostV2) | **Post** /v2/open-platform/finder/post | جستجو آگهی‌های دیوار با فیلترهایی



## FinderGetPost

> FinderGetPostResponse FinderGetPost(ctx, token).Execute()

دریافت آگهی دیوار



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinderAPI.FinderGetPost(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinderAPI.FinderGetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinderGetPost`: FinderGetPostResponse
	fmt.Fprintf(os.Stdout, "Response from `FinderAPI.FinderGetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinderGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FinderGetPostResponse**](FinderGetPostResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinderGetUser

> FinderUser FinderGetUser(ctx).Body(body).Execute()

دریافت اطلاعات کاربر



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinderAPI.FinderGetUser(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinderAPI.FinderGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinderGetUser`: FinderUser
	fmt.Fprintf(os.Stdout, "Response from `FinderAPI.FinderGetUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinderGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**FinderUser**](FinderUser.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinderGetUser2

> FinderUser FinderGetUser2(ctx).Execute()

دریافت اطلاعات کاربر



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinderAPI.FinderGetUser2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinderAPI.FinderGetUser2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinderGetUser2`: FinderUser
	fmt.Fprintf(os.Stdout, "Response from `FinderAPI.FinderGetUser2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFinderGetUser2Request struct via the builder pattern


### Return type

[**FinderUser**](FinderUser.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinderGetUserIDByPhone

> FinderGetUserIDByPhoneResponse FinderGetUserIDByPhone(ctx).FinderGetUserIDByPhoneRequest(finderGetUserIDByPhoneRequest).Execute()

دریافت شناسه دیوار کاربر با شماره تلفن



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
	finderGetUserIDByPhoneRequest := *openapiclient.NewFinderGetUserIDByPhoneRequest() // FinderGetUserIDByPhoneRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinderAPI.FinderGetUserIDByPhone(context.Background()).FinderGetUserIDByPhoneRequest(finderGetUserIDByPhoneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinderAPI.FinderGetUserIDByPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinderGetUserIDByPhone`: FinderGetUserIDByPhoneResponse
	fmt.Fprintf(os.Stdout, "Response from `FinderAPI.FinderGetUserIDByPhone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinderGetUserIDByPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **finderGetUserIDByPhoneRequest** | [**FinderGetUserIDByPhoneRequest**](FinderGetUserIDByPhoneRequest.md) |  | 

### Return type

[**FinderGetUserIDByPhoneResponse**](FinderGetUserIDByPhoneResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinderGetUserPosts

> FinderGetUserPostsResponse FinderGetUserPosts(ctx).Execute()

دریافت آگهی‌های کاربر



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinderAPI.FinderGetUserPosts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinderAPI.FinderGetUserPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinderGetUserPosts`: FinderGetUserPostsResponse
	fmt.Fprintf(os.Stdout, "Response from `FinderAPI.FinderGetUserPosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFinderGetUserPostsRequest struct via the builder pattern


### Return type

[**FinderGetUserPostsResponse**](FinderGetUserPostsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinderSearchPostV2

> FinderSearchPostV2Response FinderSearchPostV2(ctx).FinderSearchPostsV2Request(finderSearchPostsV2Request).Execute()

جستجو آگهی‌های دیوار با فیلترهایی



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
	finderSearchPostsV2Request := *openapiclient.NewFinderSearchPostsV2Request() // FinderSearchPostsV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinderAPI.FinderSearchPostV2(context.Background()).FinderSearchPostsV2Request(finderSearchPostsV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinderAPI.FinderSearchPostV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinderSearchPostV2`: FinderSearchPostV2Response
	fmt.Fprintf(os.Stdout, "Response from `FinderAPI.FinderSearchPostV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinderSearchPostV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **finderSearchPostsV2Request** | [**FinderSearchPostsV2Request**](FinderSearchPostsV2Request.md) |  | 

### Return type

[**FinderSearchPostV2Response**](FinderSearchPostV2Response.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

