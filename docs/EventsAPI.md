# \EventsAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsRegisterEventSubscription**](EventsAPI.md#EventsRegisterEventSubscription) | **Post** /v1/open-platform/events/subscriptions | اشتراک در رویداد
[**EventsSendEvent**](EventsAPI.md#EventsSendEvent) | **Post** /experimental/open-platform/events/send | ارسال رویداد به کاربر با استفاده از API



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


## EventsSendEvent

> map[string]interface{} EventsSendEvent(ctx).Message(message).TargetType(targetType).TargetResourceId(targetResourceId).Execute()

ارسال رویداد به کاربر با استفاده از API



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
	message := "message_example" // string | پیام رویداد برای نمایش به کاربر (optional)
	targetType := "targetType_example" // string | هدف رویداد؛ USER یا POST (optional)
	targetResourceId := "targetResourceId_example" // string | شناسه هدف. وقتی نوع هدف USER است، باید شناسه کاربر دیوار آن کاربر باشد و وقتی نوع هدف POST است، باید توکن آگهی باشد. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsSendEvent(context.Background()).Message(message).TargetType(targetType).TargetResourceId(targetResourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsSendEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsSendEvent`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsSendEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsSendEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **string** | پیام رویداد برای نمایش به کاربر | 
 **targetType** | **string** | هدف رویداد؛ USER یا POST | 
 **targetResourceId** | **string** | شناسه هدف. وقتی نوع هدف USER است، باید شناسه کاربر دیوار آن کاربر باشد و وقتی نوع هدف POST است، باید توکن آگهی باشد. | 

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

