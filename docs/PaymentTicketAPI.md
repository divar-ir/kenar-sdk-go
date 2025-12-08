# \PaymentTicketAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentTicketValidate**](PaymentTicketAPI.md#PaymentTicketValidate) | **Post** /v1/open-platform/payment-ticket/validate | اعتبارسنجی تیکت پرداخت



## PaymentTicketValidate

> PaymentTicketValidateResponse PaymentTicketValidate(ctx).PaymentTicketValidateRequest(paymentTicketValidateRequest).Execute()

اعتبارسنجی تیکت پرداخت



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
	paymentTicketValidateRequest := *openapiclient.NewPaymentTicketValidateRequest(int32(123), "TicketUuid_example") // PaymentTicketValidateRequest | 

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

