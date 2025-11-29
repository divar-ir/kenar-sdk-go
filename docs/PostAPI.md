# \PostAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCanUserSubmitPost**](PostAPI.md#PostCanUserSubmitPost) | **Get** /experimental/open-platform/user-posts/can-submit | بررسی اینکه آیا کاربر می‌تواند آگهی ارسال کند
[**PostCreateBusinessCustomizedButton**](PostAPI.md#PostCreateBusinessCustomizedButton) | **Post** /experimental/open-platform/business/{business_token}/customized-button | ایجاد دکمه اختصاصی برای آگهی‌های کسب‌و‌کار
[**PostDeleteBusinessCustomizedButton**](PostAPI.md#PostDeleteBusinessCustomizedButton) | **Delete** /experimental/open-platform/business/{business_token}/customized-button | حذف دکمه اختصاصی از آگهی‌های کسب‌و‌کار
[**PostDeletePostCustomizedButton**](PostAPI.md#PostDeletePostCustomizedButton) | **Delete** /experimental/open-platform/posts/{post_token}/customized-button | حذف دکمه اختصاصی از آگهی
[**PostDeleteUserPost**](PostAPI.md#PostDeleteUserPost) | **Delete** /v1/open-platform/post/{post_token} | حذف آگهی
[**PostEditPost**](PostAPI.md#PostEditPost) | **Put** /v1/open-platform/post/{post_token} | ویرایش آگهی
[**PostEditPostV2**](PostAPI.md#PostEditPostV2) | **Put** /v2/open-platform/post/{post_token} | ویرایش آگهی با پشتیبانی از فیلد ماسک
[**PostGetImageUploadURL**](PostAPI.md#PostGetImageUploadURL) | **Get** /v1/open-platform/post/image-upload-url | دریافت آدرس اپلود تصاویر آگهی (منسوخ شده)
[**PostGetPostStats**](PostAPI.md#PostGetPostStats) | **Get** /experimental/open-platform/posts/{post_token}/stats | دریافت آمارهای آگهی
[**PostGetUploadURLsV2**](PostAPI.md#PostGetUploadURLsV2) | **Get** /v2/open-platform/post/upload-urls | دریافت آدرس آپلود برای تصاویر و ویدیو‌ی آگهی‌ها
[**PostGetUserPost**](PostAPI.md#PostGetUserPost) | **Get** /v1/open-platform/user-post/{token} | دریافت آگهی با توکن
[**PostSetPostCustomizedButton**](PostAPI.md#PostSetPostCustomizedButton) | **Post** /experimental/open-platform/posts/{post_token}/customized-button | تنظیم دکمه اختصاصی بر روی آگهی ثبت شده
[**PostSubmitPost**](PostAPI.md#PostSubmitPost) | **Post** /experimental/open-platform/posts/new | ثبت آگهی
[**PostSubmitPostV2**](PostAPI.md#PostSubmitPostV2) | **Post** /experimental/open-platform/posts/new-v2 | ثبت آگهی با استفاده از اعتبارسنجی قالب JSON
[**PostSubmitUserPost**](PostAPI.md#PostSubmitUserPost) | **Post** /experimental/open-platform/user-posts/new | ثبت آگهی به عنوان کاربر



## PostCanUserSubmitPost

> PostCanUserSubmitPostResponse PostCanUserSubmitPost(ctx).Execute()

بررسی اینکه آیا کاربر می‌تواند آگهی ارسال کند



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
	resp, r, err := apiClient.PostAPI.PostCanUserSubmitPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostCanUserSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCanUserSubmitPost`: PostCanUserSubmitPostResponse
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostCanUserSubmitPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCanUserSubmitPostRequest struct via the builder pattern


### Return type

[**PostCanUserSubmitPostResponse**](PostCanUserSubmitPostResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateBusinessCustomizedButton

> map[string]interface{} PostCreateBusinessCustomizedButton(ctx, businessToken).PostCreateBusinessCustomizedButtonBody(postCreateBusinessCustomizedButtonBody).Execute()

ایجاد دکمه اختصاصی برای آگهی‌های کسب‌و‌کار



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
	businessToken := "businessToken_example" // string | 
	postCreateBusinessCustomizedButtonBody := *openapiclient.NewPostCreateBusinessCustomizedButtonBody(*openapiclient.NewPostCustomizedButton(*openapiclient.NewAddonsAction(), openapiclient.postCustomizedButtonType("ACCOMMODATION_BOOKING"))) // PostCreateBusinessCustomizedButtonBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostCreateBusinessCustomizedButton(context.Background(), businessToken).PostCreateBusinessCustomizedButtonBody(postCreateBusinessCustomizedButtonBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostCreateBusinessCustomizedButton``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCreateBusinessCustomizedButton`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostCreateBusinessCustomizedButton`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateBusinessCustomizedButtonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCreateBusinessCustomizedButtonBody** | [**PostCreateBusinessCustomizedButtonBody**](PostCreateBusinessCustomizedButtonBody.md) |  | 

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


## PostDeleteBusinessCustomizedButton

> map[string]interface{} PostDeleteBusinessCustomizedButton(ctx, businessToken).Execute()

حذف دکمه اختصاصی از آگهی‌های کسب‌و‌کار



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
	businessToken := "businessToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostDeleteBusinessCustomizedButton(context.Background(), businessToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostDeleteBusinessCustomizedButton``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeleteBusinessCustomizedButton`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostDeleteBusinessCustomizedButton`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeleteBusinessCustomizedButtonRequest struct via the builder pattern


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


## PostDeletePostCustomizedButton

> map[string]interface{} PostDeletePostCustomizedButton(ctx, postToken).Execute()

حذف دکمه اختصاصی از آگهی



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostDeletePostCustomizedButton(context.Background(), postToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostDeletePostCustomizedButton``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeletePostCustomizedButton`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostDeletePostCustomizedButton`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeletePostCustomizedButtonRequest struct via the builder pattern


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


## PostDeleteUserPost

> map[string]interface{} PostDeleteUserPost(ctx, postToken).Execute()

حذف آگهی



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
	postToken := "postToken_example" // string | توکن آگهی برای حذف

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostDeleteUserPost(context.Background(), postToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostDeleteUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeleteUserPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostDeleteUserPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** | توکن آگهی برای حذف | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeleteUserPostRequest struct via the builder pattern


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


## PostEditPostV2

> map[string]interface{} PostEditPostV2(ctx, postToken).PostEditPostV2Body(postEditPostV2Body).Execute()

ویرایش آگهی با پشتیبانی از فیلد ماسک



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
	postEditPostV2Body := *openapiclient.NewPostEditPostV2Body([]string{"UpdateMask_example"}) // PostEditPostV2Body | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostEditPostV2(context.Background(), postToken).PostEditPostV2Body(postEditPostV2Body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostEditPostV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEditPostV2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostEditPostV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** | توکن آگهی | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEditPostV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postEditPostV2Body** | [**PostEditPostV2Body**](PostEditPostV2Body.md) |  | 

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

دریافت آدرس اپلود تصاویر آگهی (منسوخ شده)



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


## PostGetUploadURLsV2

> PostGetUploadURLsV2Response PostGetUploadURLsV2(ctx).Execute()

دریافت آدرس آپلود برای تصاویر و ویدیو‌ی آگهی‌ها



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
	resp, r, err := apiClient.PostAPI.PostGetUploadURLsV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostGetUploadURLsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGetUploadURLsV2`: PostGetUploadURLsV2Response
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostGetUploadURLsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostGetUploadURLsV2Request struct via the builder pattern


### Return type

[**PostGetUploadURLsV2Response**](PostGetUploadURLsV2Response.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGetUserPost

> PostGetUserPostResponse PostGetUserPost(ctx, token).Execute()

دریافت آگهی با توکن



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
	resp, r, err := apiClient.PostAPI.PostGetUserPost(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostGetUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGetUserPost`: PostGetUserPostResponse
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostGetUserPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGetUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostGetUserPostResponse**](PostGetUserPostResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetPostCustomizedButton

> map[string]interface{} PostSetPostCustomizedButton(ctx, postToken).PostSetPostCustomizedButtonBody(postSetPostCustomizedButtonBody).Execute()

تنظیم دکمه اختصاصی بر روی آگهی ثبت شده



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
	postSetPostCustomizedButtonBody := *openapiclient.NewPostSetPostCustomizedButtonBody(*openapiclient.NewPostCustomizedButton(*openapiclient.NewAddonsAction(), openapiclient.postCustomizedButtonType("ACCOMMODATION_BOOKING"))) // PostSetPostCustomizedButtonBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.PostSetPostCustomizedButton(context.Background(), postToken).PostSetPostCustomizedButtonBody(postSetPostCustomizedButtonBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.PostSetPostCustomizedButton``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetPostCustomizedButton`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.PostSetPostCustomizedButton`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSetPostCustomizedButtonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postSetPostCustomizedButtonBody** | [**PostSetPostCustomizedButtonBody**](PostSetPostCustomizedButtonBody.md) |  | 

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
	postSubmitPostV2Request := *openapiclient.NewPostSubmitPostV2Request(map[string]interface{}(123), *openapiclient.NewPostPostGeneralData("apartment-sell", true, "tehran", "I'm available only in chat.", true, []string{"Images_example"}, openapiclient.postLocationType("LOCATION_TYPE_EMPTY"), "Temporary Residence for Rent in Tehran")) // PostSubmitPostV2Request | 

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
	postSubmitUserPostRequest := *openapiclient.NewPostSubmitUserPostRequest(map[string]interface{}(123), *openapiclient.NewPostPostGeneralData("apartment-sell", true, "tehran", "I'm available only in chat.", true, []string{"Images_example"}, openapiclient.postLocationType("LOCATION_TYPE_EMPTY"), "Temporary Residence for Rent in Tehran")) // PostSubmitUserPostRequest | 

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

