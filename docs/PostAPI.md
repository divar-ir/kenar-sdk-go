# \PostAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostEditPost**](PostAPI.md#PostEditPost) | **Put** /v1/open-platform/post/{post_token} | ویرایش آگهی
[**PostGetImageUploadURL**](PostAPI.md#PostGetImageUploadURL) | **Get** /v1/open-platform/post/image-upload-url | دریافت URL آپلود تصویر
[**PostGetPostStats**](PostAPI.md#PostGetPostStats) | **Get** /experimental/open-platform/posts/{post_token}/stats | دریافت آمارهای آگهی
[**PostSubmitEmergencyResidencePost**](PostAPI.md#PostSubmitEmergencyResidencePost) | **Post** /experimental/open-platform/posts/emergency-residence | Submit an emergency residence post
[**PostSubmitPost**](PostAPI.md#PostSubmitPost) | **Post** /experimental/open-platform/posts/new | Submit a post



## PostEditPost

> map[string]interface{} PostEditPost(ctx, postToken).PostEditPostBody(postEditPostBody).Execute()

ویرایش آگهی



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
	postToken := "postToken_example" // string | 
	postEditPostBody := *openapiclient.NewPostEditPostBody() // PostEditPostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostEditPost(context.Background(), postToken).PostEditPostBody(postEditPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEditPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostEditPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postEditPostBody** | [**PostEditPostBody**](PostEditPostBody.md) |  | 

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


## PostGetImageUploadURL

> PostGetImageUploadURLResponse PostGetImageUploadURL(ctx).Execute()

دریافت URL آپلود تصویر



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
	resp, r, err := apiClient.PostAPI.PostGetImageUploadURL(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostGetImageUploadURL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGetImageUploadURL`: PostGetImageUploadURLResponse
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostGetImageUploadURL`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostGetImageUploadURLRequest struct via the builder pattern


### Return type

[**PostGetImageUploadURLResponse**](PostGetImageUploadURLResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGetPostStats

> PostGetPostStatsResponse PostGetPostStats(ctx, postToken).Execute()

دریافت آمارهای آگهی



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
	postToken := "postToken_example" // string | توکن آگهی

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostGetPostStats(context.Background(), postToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostGetPostStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGetPostStats`: PostGetPostStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostGetPostStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** | توکن آگهی | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGetPostStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostGetPostStatsResponse**](PostGetPostStatsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSubmitEmergencyResidencePost

> PostSubmitEmergencyResidencePostResponse PostSubmitEmergencyResidencePost(ctx).PostSubmitEmergencyResidencePostRequest(postSubmitEmergencyResidencePostRequest).Execute()

Submit an emergency residence post

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
	postSubmitEmergencyResidencePostRequest := *openapiclient.NewPostSubmitEmergencyResidencePostRequest() // PostSubmitEmergencyResidencePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostSubmitEmergencyResidencePost(context.Background()).PostSubmitEmergencyResidencePostRequest(postSubmitEmergencyResidencePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostSubmitEmergencyResidencePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSubmitEmergencyResidencePost`: PostSubmitEmergencyResidencePostResponse
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostSubmitEmergencyResidencePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSubmitEmergencyResidencePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSubmitEmergencyResidencePostRequest** | [**PostSubmitEmergencyResidencePostRequest**](PostSubmitEmergencyResidencePostRequest.md) |  | 

### Return type

[**PostSubmitEmergencyResidencePostResponse**](PostSubmitEmergencyResidencePostResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSubmitPost

> PostSubmitPostResponse PostSubmitPost(ctx).PostSubmitPostRequest(postSubmitPostRequest).Execute()

Submit a post



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
	postSubmitPostRequest := *openapiclient.NewPostSubmitPostRequest() // PostSubmitPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostSubmitPost(context.Background()).PostSubmitPostRequest(postSubmitPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSubmitPost`: PostSubmitPostResponse
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSubmitPostRequest** | [**PostSubmitPostRequest**](PostSubmitPostRequest.md) |  | 

### Return type

[**PostSubmitPostResponse**](PostSubmitPostResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

