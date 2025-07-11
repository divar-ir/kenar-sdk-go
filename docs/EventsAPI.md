# \EventsAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsRegisterEventSubscription**](EventsAPI.md#EventsRegisterEventSubscription) | **Post** /v1/open-platform/events/subscriptions | اشتراک در رویداد



## EventsRegisterEventSubscription

> map[string]interface{} EventsRegisterEventSubscription(ctx).EventsRegisterEventSubscriptionRequest(eventsRegisterEventSubscriptionRequest).Execute()

اشتراک در رویداد



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
	eventsRegisterEventSubscriptionRequest := *openapiclient.NewEventsRegisterEventSubscriptionRequest() // EventsRegisterEventSubscriptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsRegisterEventSubscription(context.Background()).EventsRegisterEventSubscriptionRequest(eventsRegisterEventSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRegisterEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsRegisterEventSubscription`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsRegisterEventSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsRegisterEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventsRegisterEventSubscriptionRequest** | [**EventsRegisterEventSubscriptionRequest**](EventsRegisterEventSubscriptionRequest.md) |  | 

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

