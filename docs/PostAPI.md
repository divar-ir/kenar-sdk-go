# \PostAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostEditPost**](PostAPI.md#PostEditPost) | **Put** /v1/open-platform/post/{post_token} | ویرایش آگهی
[**PostGetImageUploadURL**](PostAPI.md#PostGetImageUploadURL) | **Get** /v1/open-platform/post/image-upload-url | دریافت URL آپلود تصویر



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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
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

