# \PaymentAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentCommitWalletTransaction**](PaymentAPI.md#PaymentCommitWalletTransaction) | **Post** /experimental/open-platform/wallet/payments/commit | تایید تراکنش کیف پول
[**PaymentCreateWalletPayment**](PaymentAPI.md#PaymentCreateWalletPayment) | **Post** /experimental/open-platform/wallet/payments/create | ایجاد پرداخت کیف پول
[**PaymentGetBalance**](PaymentAPI.md#PaymentGetBalance) | **Get** /experimental/open-platform/balance | دریافت موجودی اپلیکیشن
[**PaymentGetPostPricing**](PaymentAPI.md#PaymentGetPostPricing) | **Get** /v1/open-platform/post/{post_token}/pricing | دریافت هزینه سرویس
[**PaymentGetTransaction**](PaymentAPI.md#PaymentGetTransaction) | **Get** /experimental/open-platform/transactions/{id} | دریافت جزئیات تراکنش
[**PaymentListTransactions**](PaymentAPI.md#PaymentListTransactions) | **Get** /experimental/open-platform/transactions | لیست تراکنش‌ها
[**PaymentPublishUserPost**](PaymentAPI.md#PaymentPublishUserPost) | **Post** /experimental/open-platform/post/{post_token}/publish | پرداخت هزینه ثبت آگهی کاربر از طرف ارائه‌دهنده
[**PaymentRenewPost**](PaymentAPI.md#PaymentRenewPost) | **Post** /experimental/open-platform/post/{post_token}/renew | تمدید آگهی
[**PaymentReorderPost**](PaymentAPI.md#PaymentReorderPost) | **Post** /experimental/open-platform/post/{post_token}/reorder | نردبان آگهی
[**PaymentRetrieveWalletTransaction**](PaymentAPI.md#PaymentRetrieveWalletTransaction) | **Get** /experimental/open-platform/wallet/payments/{token} | بازیابی تراکنش کیف پول
[**PaymentSubmitUserPayment**](PaymentAPI.md#PaymentSubmitUserPayment) | **Post** /v1/open-platform/user-payments | ثبت پرداخت کاربر



## PaymentCommitWalletTransaction

> PaymentCommitWalletTransactionResponse PaymentCommitWalletTransaction(ctx).PaymentCommitWalletTransactionRequest(paymentCommitWalletTransactionRequest).Execute()

تایید تراکنش کیف پول



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
	paymentCommitWalletTransactionRequest := *openapiclient.NewPaymentCommitWalletTransactionRequest() // PaymentCommitWalletTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentCommitWalletTransaction(context.Background()).PaymentCommitWalletTransactionRequest(paymentCommitWalletTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentCommitWalletTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentCommitWalletTransaction`: PaymentCommitWalletTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentCommitWalletTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentCommitWalletTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentCommitWalletTransactionRequest** | [**PaymentCommitWalletTransactionRequest**](PaymentCommitWalletTransactionRequest.md) |  | 

### Return type

[**PaymentCommitWalletTransactionResponse**](PaymentCommitWalletTransactionResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentCreateWalletPayment

> PaymentCreateWalletPaymentResponse PaymentCreateWalletPayment(ctx).PaymentCreateWalletPaymentRequest(paymentCreateWalletPaymentRequest).Execute()

ایجاد پرداخت کیف پول



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
	paymentCreateWalletPaymentRequest := *openapiclient.NewPaymentCreateWalletPaymentRequest("AmountRials_example", "Description_example", "RedirectUrl_example") // PaymentCreateWalletPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentCreateWalletPayment(context.Background()).PaymentCreateWalletPaymentRequest(paymentCreateWalletPaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentCreateWalletPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentCreateWalletPayment`: PaymentCreateWalletPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentCreateWalletPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentCreateWalletPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentCreateWalletPaymentRequest** | [**PaymentCreateWalletPaymentRequest**](PaymentCreateWalletPaymentRequest.md) |  | 

### Return type

[**PaymentCreateWalletPaymentResponse**](PaymentCreateWalletPaymentResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentGetBalance

> PaymentGetBalanceResponse PaymentGetBalance(ctx).Execute()

دریافت موجودی اپلیکیشن



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
	resp, r, err := apiClient.PaymentAPI.PaymentGetBalance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentGetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentGetBalance`: PaymentGetBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentGetBalance`: %v\n", resp)
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
	openapiclient "github.com/divar-ir/kenar-sdk-go"
)

func main() {
	postToken := "postToken_example" // string | شناسه منحصر به فرد 8-9 کاراکتری برای آگهی

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentGetPostPricing(context.Background(), postToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentGetPostPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentGetPostPricing`: PaymentGetPostPricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentGetPostPricing`: %v\n", resp)
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

دریافت جزئیات تراکنش



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
	id := "id_example" // string | شناسه منحصر به فرد برای تراکنش، همان id در درخواست

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentGetTransaction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentGetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentGetTransaction`: PaymentGetTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentGetTransaction`: %v\n", resp)
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

لیست تراکنش‌ها



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
	pageSize := int32(56) // int32 | تعداد تراکنش‌ها برای برگرداندن در هر صفحه (optional)
	pageToken := "pageToken_example" // string | توکن برای صفحه بعدی نتایج (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentListTransactions(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentListTransactions`: PaymentListTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | تعداد تراکنش‌ها برای برگرداندن در هر صفحه | 
 **pageToken** | **string** | توکن برای صفحه بعدی نتایج | 

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


## PaymentPublishUserPost

> PaymentPublishUserPostResponse PaymentPublishUserPost(ctx, postToken).PaymentPublishUserPostBody(paymentPublishUserPostBody).Execute()

پرداخت هزینه ثبت آگهی کاربر از طرف ارائه‌دهنده



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
	postToken := "postToken_example" // string | توکن آگهی دریافت شده از RPC SubmitUserPost. شناسه منحصر به فرد 8-9 کاراکتری برای آگهی ثبت شده توسط کاربر.
	paymentPublishUserPostBody := *openapiclient.NewPaymentPublishUserPostBody() // PaymentPublishUserPostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentPublishUserPost(context.Background(), postToken).PaymentPublishUserPostBody(paymentPublishUserPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentPublishUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentPublishUserPost`: PaymentPublishUserPostResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentPublishUserPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** | توکن آگهی دریافت شده از RPC SubmitUserPost. شناسه منحصر به فرد 8-9 کاراکتری برای آگهی ثبت شده توسط کاربر. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentPublishUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentPublishUserPostBody** | [**PaymentPublishUserPostBody**](PaymentPublishUserPostBody.md) |  | 

### Return type

[**PaymentPublishUserPostResponse**](PaymentPublishUserPostResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentRenewPost

> PaymentRenewPostResponse PaymentRenewPost(ctx, postToken).PaymentRenewPostBody(paymentRenewPostBody).Execute()

تمدید آگهی



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
	paymentRenewPostBody := *openapiclient.NewPaymentRenewPostBody() // PaymentRenewPostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentRenewPost(context.Background(), postToken).PaymentRenewPostBody(paymentRenewPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentRenewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRenewPost`: PaymentRenewPostResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentRenewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRenewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentRenewPostBody** | [**PaymentRenewPostBody**](PaymentRenewPostBody.md) |  | 

### Return type

[**PaymentRenewPostResponse**](PaymentRenewPostResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentReorderPost

> PaymentReorderPostResponse PaymentReorderPost(ctx, postToken).PaymentReorderPostBody(paymentReorderPostBody).Execute()

نردبان آگهی



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
	paymentReorderPostBody := *openapiclient.NewPaymentReorderPostBody() // PaymentReorderPostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentReorderPost(context.Background(), postToken).PaymentReorderPostBody(paymentReorderPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentReorderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentReorderPost`: PaymentReorderPostResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentReorderPost`: %v\n", resp)
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


## PaymentRetrieveWalletTransaction

> PaymentRetrieveWalletTransactionResponse PaymentRetrieveWalletTransaction(ctx, token).Execute()

بازیابی تراکنش کیف پول



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
	token := "token_example" // string | توکن تراکنشی که می‌خواهید بازیابی کنید

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentRetrieveWalletTransaction(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentRetrieveWalletTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRetrieveWalletTransaction`: PaymentRetrieveWalletTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentRetrieveWalletTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | توکن تراکنشی که می‌خواهید بازیابی کنید | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRetrieveWalletTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentRetrieveWalletTransactionResponse**](PaymentRetrieveWalletTransactionResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentSubmitUserPayment

> map[string]interface{} PaymentSubmitUserPayment(ctx).PaymentSubmitUserPaymentRequest(paymentSubmitUserPaymentRequest).Execute()

ثبت پرداخت کاربر



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
	paymentSubmitUserPaymentRequest := *openapiclient.NewPaymentSubmitUserPaymentRequest("AmountRials_example", "ProfitRials_example", "ReferenceId_example", []string{"Services_example"}) // PaymentSubmitUserPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentSubmitUserPayment(context.Background()).PaymentSubmitUserPaymentRequest(paymentSubmitUserPaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentSubmitUserPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentSubmitUserPayment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentSubmitUserPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentSubmitUserPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentSubmitUserPaymentRequest** | [**PaymentSubmitUserPaymentRequest**](PaymentSubmitUserPaymentRequest.md) |  | 

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

