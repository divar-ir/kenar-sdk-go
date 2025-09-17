# \PostAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostEditPost**](PostAPI.md#PostEditPost) | **Put** /v1/open-platform/post/{post_token} | ویرایش آگهی
[**PostGetImageUploadURL**](PostAPI.md#PostGetImageUploadURL) | **Get** /v1/open-platform/post/image-upload-url | دریافت URL آپلود تصویر
[**PostGetPostStats**](PostAPI.md#PostGetPostStats) | **Get** /experimental/open-platform/posts/{post_token}/stats | دریافت آمارهای آگهی
[**PostSubmitPost**](PostAPI.md#PostSubmitPost) | **Post** /experimental/open-platform/posts/new | ثبت آگهی
[**PostSubmitPostV2**](PostAPI.md#PostSubmitPostV2) | **Post** /experimental/open-platform/posts/new-v2 | ثبت آگهی با استفاده از اعتبارسنجی قالب JSON
[**PostSubmitUserPost**](PostAPI.md#PostSubmitUserPost) | **Post** /experimental/open-platform/user-posts/new | ثبت آگهی به عنوان کاربر



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
	postEditPostBody := *openapiclient.NewPostEditPostBody("Description_example", "Title_example") // PostEditPostBody | 

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


## PostSubmitPost

> PostSubmitPostResponse PostSubmitPost(ctx).PostSubmitPostRequest(postSubmitPostRequest).Execute()

ثبت آگهی



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
	postSubmitPostRequest := *openapiclient.NewPostSubmitPostRequest(true, "tehran", "I'm available only in chat.", true, []string{"Images_example"}, openapiclient.postLocationType("LOCATION_TYPE_EMPTY"), "Temporary Residence for Rent in Tehran") // PostSubmitPostRequest | 

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


## PostSubmitPostV2

> PostSubmitPostResponse PostSubmitPostV2(ctx).PostSubmitPostV2Request(postSubmitPostV2Request).Execute()

ثبت آگهی با استفاده از اعتبارسنجی قالب JSON



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
	postSubmitPostV2Request := *openapiclient.NewPostSubmitPostV2Request(map[string]interface{}(123), *openapiclient.NewPostSubmitPostGeneralData("apartment-sell", true, "tehran", "I'm available only in chat.", true, []string{"Images_example"}, openapiclient.postLocationType("LOCATION_TYPE_EMPTY"), "Temporary Residence for Rent in Tehran")) // PostSubmitPostV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostSubmitPostV2(context.Background()).PostSubmitPostV2Request(postSubmitPostV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostSubmitPostV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSubmitPostV2`: PostSubmitPostResponse
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostSubmitPostV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSubmitPostV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSubmitPostV2Request** | [**PostSubmitPostV2Request**](PostSubmitPostV2Request.md) |  | 

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


## PostSubmitUserPost

> PostSubmitPostResponse PostSubmitUserPost(ctx).PostSubmitUserPostRequest(postSubmitUserPostRequest).Execute()

ثبت آگهی به عنوان کاربر



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
	postSubmitUserPostRequest := *openapiclient.NewPostSubmitUserPostRequest(map[string]interface{}(123), *openapiclient.NewPostSubmitPostGeneralData("apartment-sell", true, "tehran", "I'm available only in chat.", true, []string{"Images_example"}, openapiclient.postLocationType("LOCATION_TYPE_EMPTY"), "Temporary Residence for Rent in Tehran")) // PostSubmitUserPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostSubmitUserPost(context.Background()).PostSubmitUserPostRequest(postSubmitUserPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostSubmitUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSubmitUserPost`: PostSubmitPostResponse
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostSubmitUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSubmitUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSubmitUserPostRequest** | [**PostSubmitUserPostRequest**](PostSubmitUserPostRequest.md) |  | 

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

