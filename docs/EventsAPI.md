# \EventsAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsRegisterEventSubscription**](EventsAPI.md#EventsRegisterEventSubscription) | **Post** /v1/open-platform/events/subscriptions | اشتراک در رویداد
[**EventsSendEvent**](EventsAPI.md#EventsSendEvent) | **Post** /experimental/open-platform/events/send | Send an event to a user



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

> map[string]interface{} EventsSendEvent(ctx).PrimaryButtonActionOpenPostPagePostToken(primaryButtonActionOpenPostPagePostToken).PrimaryButtonActionOpenPostManagePagePostToken(primaryButtonActionOpenPostManagePagePostToken).Message(message).PrimaryButtonTitle(primaryButtonTitle).PrimaryButtonActionOpenDirectLink(primaryButtonActionOpenDirectLink).PrimaryButtonActionOpenServerLinkData(primaryButtonActionOpenServerLinkData).PrimaryButtonActionGetDynamicActionData(primaryButtonActionGetDynamicActionData).TargetType(targetType).TargetResourceId(targetResourceId).Execute()

Send an event to a user



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
	primaryButtonActionOpenPostPagePostToken := "primaryButtonActionOpenPostPagePostToken_example" // string | Token of the post to open
	primaryButtonActionOpenPostManagePagePostToken := "primaryButtonActionOpenPostManagePagePostToken_example" // string | Token of the post to redirect to its management page
	message := "message_example" // string | The event message to display to the user (optional)
	primaryButtonTitle := "primaryButtonTitle_example" // string | The text to display on the button (optional)
	primaryButtonActionOpenDirectLink := "primaryButtonActionOpenDirectLink_example" // string | An action to send user to your URL directly with just a resource id (if applicable) (optional)
	primaryButtonActionOpenServerLinkData := map[string]interface{}{ ... } // map[string]interface{} | A data that you can set and will be returned to you upon user click to recognize the action (optional)
	primaryButtonActionGetDynamicActionData := map[string]interface{}{ ... } // map[string]interface{} | A data that you can set and will be returned to you upon user click to recognize the action (optional)
	targetType := "targetType_example" // string | Target of the event; USER or POST (optional)
	targetResourceId := "targetResourceId_example" // string | id of the target. When target type is USER, it should be the Divar User ID of that user and when target type is POST, it should be the post token.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsSendEvent(context.Background()).PrimaryButtonActionOpenPostPagePostToken(primaryButtonActionOpenPostPagePostToken).PrimaryButtonActionOpenPostManagePagePostToken(primaryButtonActionOpenPostManagePagePostToken).Message(message).PrimaryButtonTitle(primaryButtonTitle).PrimaryButtonActionOpenDirectLink(primaryButtonActionOpenDirectLink).PrimaryButtonActionOpenServerLinkData(primaryButtonActionOpenServerLinkData).PrimaryButtonActionGetDynamicActionData(primaryButtonActionGetDynamicActionData).TargetType(targetType).TargetResourceId(targetResourceId).Execute()
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
 **primaryButtonActionOpenPostPagePostToken** | **string** | Token of the post to open | 
 **primaryButtonActionOpenPostManagePagePostToken** | **string** | Token of the post to redirect to its management page | 
 **message** | **string** | The event message to display to the user | 
 **primaryButtonTitle** | **string** | The text to display on the button | 
 **primaryButtonActionOpenDirectLink** | **string** | An action to send user to your URL directly with just a resource id (if applicable) | 
 **primaryButtonActionOpenServerLinkData** | [**map[string]interface{}**](map[string]interface{}.md) | A data that you can set and will be returned to you upon user click to recognize the action | 
 **primaryButtonActionGetDynamicActionData** | [**map[string]interface{}**](map[string]interface{}.md) | A data that you can set and will be returned to you upon user click to recognize the action | 
 **targetType** | **string** | Target of the event; USER or POST | 
 **targetResourceId** | **string** | id of the target. When target type is USER, it should be the Divar User ID of that user and when target type is POST, it should be the post token.  | 

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

