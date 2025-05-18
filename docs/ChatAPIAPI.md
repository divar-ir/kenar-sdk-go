# \ChatAPIAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatAPIChatBotSendMessage**](ChatAPIAPI.md#ChatAPIChatBotSendMessage) | **Post** /experimental/open-platform/chatbot-conversations/{conversation_id}/messages | Send a message to a ChatBot conversation
[**ChatAPIChatBotSendMessage2**](ChatAPIAPI.md#ChatAPIChatBotSendMessage2) | **Post** /experimental/open-platform/chat/bot/users/{user_id}/messages | Send a message to a ChatBot conversation
[**ChatAPIChatBotSendMessage3**](ChatAPIAPI.md#ChatAPIChatBotSendMessage3) | **Post** /experimental/open-platform/chat/bot/conversations/{conversation_id}/messages | Send a message to a ChatBot conversation
[**ChatAPIConversationSendMessage**](ChatAPIAPI.md#ChatAPIConversationSendMessage) | **Post** /v2/open-platform/conversations/{conversation_id}/messages | Send a message to a conversation
[**ChatAPIGenerateUploadToken**](ChatAPIAPI.md#ChatAPIGenerateUploadToken) | **Post** /experimental/open-platform/chat/upload | Generate an upload token
[**ChatAPIGetConversation**](ChatAPIAPI.md#ChatAPIGetConversation) | **Get** /v1/open-platform/chat/conversations/{conversation_id} | Get Conversation by it&#39;s ID



## ChatAPIChatBotSendMessage

> ChatapiChatBotSendMessageResponse ChatAPIChatBotSendMessage(ctx, conversationId).ChatAPIChatBotSendMessageBody(chatAPIChatBotSendMessageBody).Execute()

Send a message to a ChatBot conversation



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
	conversationId := "conversationId_example" // string | Unique identifier for the conversation
	chatAPIChatBotSendMessageBody := *openapiclient.NewChatAPIChatBotSendMessageBody("Thank you for your inquiry. How can I help you?") // ChatAPIChatBotSendMessageBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.ChatAPIChatBotSendMessage(context.Background(), conversationId).ChatAPIChatBotSendMessageBody(chatAPIChatBotSendMessageBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.ChatAPIChatBotSendMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatAPIChatBotSendMessage`: ChatapiChatBotSendMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.ChatAPIChatBotSendMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | Unique identifier for the conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatAPIChatBotSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatAPIChatBotSendMessageBody** | [**ChatAPIChatBotSendMessageBody**](ChatAPIChatBotSendMessageBody.md) |  | 

### Return type

[**ChatapiChatBotSendMessageResponse**](ChatapiChatBotSendMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatAPIChatBotSendMessage2

> ChatapiChatBotSendMessageResponse ChatAPIChatBotSendMessage2(ctx, userId).ChatAPIChatBotSendMessageBody(chatAPIChatBotSendMessageBody).Execute()

Send a message to a ChatBot conversation



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
	userId := "userId_example" // string | Unique identifier for the user to start or continue a conversation with
	chatAPIChatBotSendMessageBody := *openapiclient.NewChatAPIChatBotSendMessageBody("Thank you for your inquiry. How can I help you?") // ChatAPIChatBotSendMessageBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.ChatAPIChatBotSendMessage2(context.Background(), userId).ChatAPIChatBotSendMessageBody(chatAPIChatBotSendMessageBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.ChatAPIChatBotSendMessage2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatAPIChatBotSendMessage2`: ChatapiChatBotSendMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.ChatAPIChatBotSendMessage2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Unique identifier for the user to start or continue a conversation with | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatAPIChatBotSendMessage2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatAPIChatBotSendMessageBody** | [**ChatAPIChatBotSendMessageBody**](ChatAPIChatBotSendMessageBody.md) |  | 

### Return type

[**ChatapiChatBotSendMessageResponse**](ChatapiChatBotSendMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatAPIChatBotSendMessage3

> ChatapiChatBotSendMessageResponse ChatAPIChatBotSendMessage3(ctx, conversationId).ChatAPIChatBotSendMessageBody(chatAPIChatBotSendMessageBody).Execute()

Send a message to a ChatBot conversation



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
	conversationId := "conversationId_example" // string | Unique identifier for the conversation
	chatAPIChatBotSendMessageBody := *openapiclient.NewChatAPIChatBotSendMessageBody("Thank you for your inquiry. How can I help you?") // ChatAPIChatBotSendMessageBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.ChatAPIChatBotSendMessage3(context.Background(), conversationId).ChatAPIChatBotSendMessageBody(chatAPIChatBotSendMessageBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.ChatAPIChatBotSendMessage3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatAPIChatBotSendMessage3`: ChatapiChatBotSendMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.ChatAPIChatBotSendMessage3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | Unique identifier for the conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatAPIChatBotSendMessage3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatAPIChatBotSendMessageBody** | [**ChatAPIChatBotSendMessageBody**](ChatAPIChatBotSendMessageBody.md) |  | 

### Return type

[**ChatapiChatBotSendMessageResponse**](ChatapiChatBotSendMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatAPIConversationSendMessage

> ChatapiConversationSendMessageResponse ChatAPIConversationSendMessage(ctx, conversationId).ChatAPIConversationSendMessageBody(chatAPIConversationSendMessageBody).Execute()

Send a message to a conversation



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
	conversationId := "conversationId_example" // string | Unique identifier for the conversation
	chatAPIConversationSendMessageBody := *openapiclient.NewChatAPIConversationSendMessageBody("Hello, I'm interested in your product") // ChatAPIConversationSendMessageBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.ChatAPIConversationSendMessage(context.Background(), conversationId).ChatAPIConversationSendMessageBody(chatAPIConversationSendMessageBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.ChatAPIConversationSendMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatAPIConversationSendMessage`: ChatapiConversationSendMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.ChatAPIConversationSendMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | Unique identifier for the conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatAPIConversationSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatAPIConversationSendMessageBody** | [**ChatAPIConversationSendMessageBody**](ChatAPIConversationSendMessageBody.md) |  | 

### Return type

[**ChatapiConversationSendMessageResponse**](ChatapiConversationSendMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatAPIGenerateUploadToken

> ChatapiGenerateUploadTokenResponse ChatAPIGenerateUploadToken(ctx).Body(body).Execute()

Generate an upload token



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.ChatAPIGenerateUploadToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.ChatAPIGenerateUploadToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatAPIGenerateUploadToken`: ChatapiGenerateUploadTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.ChatAPIGenerateUploadToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatAPIGenerateUploadTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**ChatapiGenerateUploadTokenResponse**](ChatapiGenerateUploadTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatAPIGetConversation

> ChatapiGetConversationResponse ChatAPIGetConversation(ctx, conversationId).Execute()

Get Conversation by it's ID



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
	conversationId := "conversationId_example" // string | Unique identifier for the conversation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.ChatAPIGetConversation(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.ChatAPIGetConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatAPIGetConversation`: ChatapiGetConversationResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.ChatAPIGetConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | Unique identifier for the conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatAPIGetConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatapiGetConversationResponse**](ChatapiGetConversationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

