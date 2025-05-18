# \EventsAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsRegisterEventSubscription**](EventsAPI.md#EventsRegisterEventSubscription) | **Post** /v1/open-platform/events/subscriptions | Subscribe to an event



## EventsRegisterEventSubscription

> map[string]interface{} EventsRegisterEventSubscription(ctx).EventsRegisterEventSubscriptionRequest(eventsRegisterEventSubscriptionRequest).Execute()

Subscribe to an event



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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

