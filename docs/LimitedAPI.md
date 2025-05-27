# \LimitedAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentGetBalance**](LimitedAPI.md#PaymentGetBalance) | **Get** /experimental/open-platform/balance | 
[**PaymentGetPostPricing**](LimitedAPI.md#PaymentGetPostPricing) | **Get** /v1/open-platform/post/{post_token}/pricing | دریافت هزینه سرویس
[**PaymentGetTransaction**](LimitedAPI.md#PaymentGetTransaction) | **Get** /experimental/open-platform/transactions/{id} | 
[**PaymentListTransactions**](LimitedAPI.md#PaymentListTransactions) | **Get** /experimental/open-platform/transactions | 
[**PaymentReorderPost**](LimitedAPI.md#PaymentReorderPost) | **Post** /experimental/open-platform/post/{post_token}/reorder | 



## PaymentGetBalance

> PaymentGetBalanceResponse PaymentGetBalance(ctx).Execute()





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
	resp, r, err := apiClient.LimitedAPI.PaymentGetBalance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitedAPI.PaymentGetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentGetBalance`: PaymentGetBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitedAPI.PaymentGetBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentGetBalanceRequest struct via the builder pattern


### Return type

[**PaymentGetBalanceResponse**](PaymentGetBalanceResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentGetPostPricing

> PaymentGetPostPricingResponse PaymentGetPostPricing(ctx, postToken).Execute()

دریافت هزینه سرویس



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
	postToken := "postToken_example" // string | شناسه منحصر به فرد 8-9 کاراکتری برای آگهی

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitedAPI.PaymentGetPostPricing(context.Background(), postToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitedAPI.PaymentGetPostPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentGetPostPricing`: PaymentGetPostPricingResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitedAPI.PaymentGetPostPricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** | شناسه منحصر به فرد 8-9 کاراکتری برای آگهی | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentGetPostPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentGetPostPricingResponse**](PaymentGetPostPricingResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentGetTransaction

> PaymentGetTransactionResponse PaymentGetTransaction(ctx, id).Execute()





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
	id := "id_example" // string | شناسه منحصر به فرد برای تراکنش، همان id در درخواست

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitedAPI.PaymentGetTransaction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitedAPI.PaymentGetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentGetTransaction`: PaymentGetTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitedAPI.PaymentGetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | شناسه منحصر به فرد برای تراکنش، همان id در درخواست | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentGetTransactionResponse**](PaymentGetTransactionResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentListTransactions

> PaymentListTransactionsResponse PaymentListTransactions(ctx).PageSize(pageSize).PageToken(pageToken).Execute()





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
	pageSize := int32(56) // int32 | Number of transactions to return per page (optional)
	pageToken := "pageToken_example" // string | Token for the next page of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitedAPI.PaymentListTransactions(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitedAPI.PaymentListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentListTransactions`: PaymentListTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitedAPI.PaymentListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of transactions to return per page | 
 **pageToken** | **string** | Token for the next page of results | 

### Return type

[**PaymentListTransactionsResponse**](PaymentListTransactionsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentReorderPost

> PaymentReorderPostResponse PaymentReorderPost(ctx, postToken).PaymentReorderPostBody(paymentReorderPostBody).Execute()





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
	paymentReorderPostBody := *openapiclient.NewPaymentReorderPostBody() // PaymentReorderPostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitedAPI.PaymentReorderPost(context.Background(), postToken).PaymentReorderPostBody(paymentReorderPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitedAPI.PaymentReorderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentReorderPost`: PaymentReorderPostResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitedAPI.PaymentReorderPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentReorderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentReorderPostBody** | [**PaymentReorderPostBody**](PaymentReorderPostBody.md) |  | 

### Return type

[**PaymentReorderPostResponse**](PaymentReorderPostResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

