# \PaymentTicketAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentTicketValidate**](PaymentTicketAPI.md#PaymentTicketValidate) | **Post** /v1/open-platform/payment-ticket/validate | Validate a payment ticket



## PaymentTicketValidate

> PaymentTicketValidateResponse PaymentTicketValidate(ctx).PaymentTicketValidateRequest(paymentTicketValidateRequest).Execute()

Validate a payment ticket



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
	paymentTicketValidateRequest := *openapiclient.NewPaymentTicketValidateRequest() // PaymentTicketValidateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentTicketAPI.PaymentTicketValidate(context.Background()).PaymentTicketValidateRequest(paymentTicketValidateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentTicketAPI.PaymentTicketValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentTicketValidate`: PaymentTicketValidateResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentTicketAPI.PaymentTicketValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentTicketValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentTicketValidateRequest** | [**PaymentTicketValidateRequest**](PaymentTicketValidateRequest.md) |  | 

### Return type

[**PaymentTicketValidateResponse**](PaymentTicketValidateResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

