# Go API client for kenarapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 0.1.0
- Generator version: 7.13.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import kenarapi "github.com/GIT_USER_ID/GIT_REPO_ID/kenarapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `kenarapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), kenarapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `kenarapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), kenarapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `kenarapi.ContextOperationServerIndices` and `kenarapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), kenarapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), kenarapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://open-api.divar.ir*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddonsAPI* | [**AddonsCreatePostAddonV2**](docs/AddonsAPI.md#addonscreatepostaddonv2) | **Post** /v2/open-platform/addons/post/{token} | Attach a new Addon to a post
*AddonsAPI* | [**AddonsCreateUserAddonV2**](docs/AddonsAPI.md#addonscreateuseraddonv2) | **Post** /v2/open-platform/addons/user/{phone} | Attach a new Addon to a user
*AddonsAPI* | [**AddonsCreateUserAddonV22**](docs/AddonsAPI.md#addonscreateuseraddonv22) | **Post** /v2/open-platform/addons/users/{divar_user_id} | Attach a new Addon to a user
*AddonsAPI* | [**AddonsDeletePostAddon**](docs/AddonsAPI.md#addonsdeletepostaddon) | **Delete** /v1/open-platform/add-ons/post/{token} | Delete an Addon from a post
*AddonsAPI* | [**AddonsDeletePostAddon2**](docs/AddonsAPI.md#addonsdeletepostaddon2) | **Delete** /v1/open-platform/addons/post/{token} | Delete an Addon from a post
*AddonsAPI* | [**AddonsDeleteUserAddon**](docs/AddonsAPI.md#addonsdeleteuseraddon) | **Delete** /v1/open-platform/addons/user/{id} | Delete an UserAddon
*AddonsAPI* | [**AddonsGetUserAddons**](docs/AddonsAPI.md#addonsgetuseraddons) | **Get** /v1/open-platform/addons/user/{phone} | Retrieve all UserAddons
*AddonsAPI* | [**AddonsGetUserAddons2**](docs/AddonsAPI.md#addonsgetuseraddons2) | **Get** /v2/open-platform/addons/users/{divar_user_id} | Retrieve all UserAddons
*AssetsAPI* | [**AssetsGetBodyStatuses**](docs/AssetsAPI.md#assetsgetbodystatuses) | **Get** /v1/open-platform/assets/body-status | Get available body-status options in Divar vehicle categories
*AssetsAPI* | [**AssetsGetBrandModels**](docs/AssetsAPI.md#assetsgetbrandmodels) | **Get** /v1/open-platform/assets/brand-model/{category} | Get brand models in a Divar category
*AssetsAPI* | [**AssetsGetCategories**](docs/AssetsAPI.md#assetsgetcategories) | **Get** /v1/open-platform/assets/category | Get Divar categories
*AssetsAPI* | [**AssetsGetCities**](docs/AssetsAPI.md#assetsgetcities) | **Get** /v1/open-platform/assets/city | Get Divar cities
*AssetsAPI* | [**AssetsGetColors**](docs/AssetsAPI.md#assetsgetcolors) | **Get** /v1/open-platform/assets/color/{category} | Get colors in a Divar category
*AssetsAPI* | [**AssetsGetDistricts**](docs/AssetsAPI.md#assetsgetdistricts) | **Get** /v1/open-platform/assets/district | Get Divar districts
*AssetsAPI* | [**AssetsGetDistricts2**](docs/AssetsAPI.md#assetsgetdistricts2) | **Get** /v1/open-platform/assets/district/{city_slug} | Get Divar districts
*AssetsAPI* | [**AssetsGetInternalStorages**](docs/AssetsAPI.md#assetsgetinternalstorages) | **Get** /v1/open-platform/assets/internal-storage | Get available internal-storage options in Divar mobile/tablet/laptop categories
*AssetsAPI* | [**AssetsGetOAuthScopes**](docs/AssetsAPI.md#assetsgetoauthscopes) | **Get** /v1/open-platform/assets/oauth-scope | Get KenarDivar OAuth scopes
*AssetsAPI* | [**AssetsGetPermissions**](docs/AssetsAPI.md#assetsgetpermissions) | **Get** /v1/open-platform/assets/permission | Get KenarDivar permissions
*AssetsAPI* | [**AssetsGetRamMemories**](docs/AssetsAPI.md#assetsgetrammemories) | **Get** /v1/open-platform/assets/ram-memory | Get available ram-memory options in Divar mobile/tablet/laptop categories
*AssetsAPI* | [**AssetsGetServiceTypes**](docs/AssetsAPI.md#assetsgetservicetypes) | **Get** /v1/open-platform/assets/service-type | Get available service types in KenarDivar
*ChatAPIAPI* | [**ChatAPIChatBotSendMessage**](docs/ChatAPIAPI.md#chatapichatbotsendmessage) | **Post** /experimental/open-platform/chatbot-conversations/{conversation_id}/messages | Send a message to a ChatBot conversation
*ChatAPIAPI* | [**ChatAPIChatBotSendMessage2**](docs/ChatAPIAPI.md#chatapichatbotsendmessage2) | **Post** /experimental/open-platform/chat/bot/users/{user_id}/messages | Send a message to a ChatBot conversation
*ChatAPIAPI* | [**ChatAPIChatBotSendMessage3**](docs/ChatAPIAPI.md#chatapichatbotsendmessage3) | **Post** /experimental/open-platform/chat/bot/conversations/{conversation_id}/messages | Send a message to a ChatBot conversation
*ChatAPIAPI* | [**ChatAPIConversationSendMessage**](docs/ChatAPIAPI.md#chatapiconversationsendmessage) | **Post** /v2/open-platform/conversations/{conversation_id}/messages | Send a message to a conversation
*ChatAPIAPI* | [**ChatAPIGenerateUploadToken**](docs/ChatAPIAPI.md#chatapigenerateuploadtoken) | **Post** /experimental/open-platform/chat/upload | Generate an upload token
*ChatAPIAPI* | [**ChatAPIGetConversation**](docs/ChatAPIAPI.md#chatapigetconversation) | **Get** /v1/open-platform/chat/conversations/{conversation_id} | Get Conversation by it&#39;s ID
*EventsAPI* | [**EventsRegisterEventSubscription**](docs/EventsAPI.md#eventsregistereventsubscription) | **Post** /v1/open-platform/events/subscriptions | Subscribe to an event
*FinderAPI* | [**FinderGetPost**](docs/FinderAPI.md#findergetpost) | **Get** /v1/open-platform/finder/post/{token} | Get a Divar post
*FinderAPI* | [**FinderGetUser**](docs/FinderAPI.md#findergetuser) | **Post** /v1/open-platform/users | Get user information
*FinderAPI* | [**FinderGetUser2**](docs/FinderAPI.md#findergetuser2) | **Get** /v1/open-platform/users | Get user information
*FinderAPI* | [**FinderGetUserPosts**](docs/FinderAPI.md#findergetuserposts) | **Get** /v1/open-platform/finder/user-posts | Get user posts
*FinderAPI* | [**FinderSearchPostV2**](docs/FinderAPI.md#findersearchpostv2) | **Post** /v2/open-platform/finder/post | Search Divar posts with some filters
*LimitedAPI* | [**PaymentGetBalance**](docs/LimitedAPI.md#paymentgetbalance) | **Get** /experimental/open-platform/balance | 
*LimitedAPI* | [**PaymentGetPostPricing**](docs/LimitedAPI.md#paymentgetpostpricing) | **Get** /v1/open-platform/post/{post_token}/pricing | Retrieve the cost of the service
*LimitedAPI* | [**PaymentGetTransaction**](docs/LimitedAPI.md#paymentgettransaction) | **Get** /experimental/open-platform/transactions/{id} | 
*LimitedAPI* | [**PaymentReorderPost**](docs/LimitedAPI.md#paymentreorderpost) | **Post** /experimental/open-platform/post/{post_token}/reorder | 
*PaymentTicketAPI* | [**PaymentTicketValidate**](docs/PaymentTicketAPI.md#paymentticketvalidate) | **Post** /v1/open-platform/payment-ticket/validate | Validate a payment ticket
*PostAPI* | [**PostEditPost**](docs/PostAPI.md#posteditpost) | **Put** /v1/open-platform/post/{post_token} | Edit a post
*PostAPI* | [**PostGetImageUploadURL**](docs/PostAPI.md#postgetimageuploadurl) | **Get** /v1/open-platform/post/image-upload-url | Get image upload URL
*SemanticAPI* | [**SemanticCreatePostSemantic**](docs/SemanticAPI.md#semanticcreatepostsemantic) | **Post** /experimental/open-platform/semantic/post/{token} | Create Post Semantic
*SemanticAPI* | [**SemanticCreateUserSemantic**](docs/SemanticAPI.md#semanticcreateusersemantic) | **Post** /v1/open-platform/semantic/user/{phone} | Create User Semantic
*SemanticAPI* | [**SemanticCreateUserSemantic2**](docs/SemanticAPI.md#semanticcreateusersemantic2) | **Post** /v1/open-platform/semantic/users/{divar_user_id} | Create User Semantic
*SemanticAPI* | [**SemanticDeleteUserSemantic**](docs/SemanticAPI.md#semanticdeleteusersemantic) | **Delete** /v1/open-platform/semantic/user/{phone} | Delete User Semantic
*SemanticAPI* | [**SemanticDeleteUserSemantic2**](docs/SemanticAPI.md#semanticdeleteusersemantic2) | **Delete** /v1/open-platform/semantic/users/{divar_user_id} | Delete User Semantic


## Documentation For Models

 - [AddonSecondaryLinkagePosition](docs/AddonSecondaryLinkagePosition.md)
 - [AddonSemanticCarVerificationStage](docs/AddonSemanticCarVerificationStage.md)
 - [AddonSemanticIdentityVerificationResult](docs/AddonSemanticIdentityVerificationResult.md)
 - [AddonSemanticIdentityVerificationStage](docs/AddonSemanticIdentityVerificationStage.md)
 - [AddonSemanticInspectionResult](docs/AddonSemanticInspectionResult.md)
 - [AddonSemanticNewFaceVerificationResult](docs/AddonSemanticNewFaceVerificationResult.md)
 - [AddonSemanticOwnershipResult](docs/AddonSemanticOwnershipResult.md)
 - [AddonSemanticPostVerificationResult](docs/AddonSemanticPostVerificationResult.md)
 - [AddonsAction](docs/AddonsAction.md)
 - [AddonsAddonLinkage](docs/AddonsAddonLinkage.md)
 - [AddonsAddonMetaData](docs/AddonsAddonMetaData.md)
 - [AddonsAddonSecondaryLinkage](docs/AddonsAddonSecondaryLinkage.md)
 - [AddonsAddonSecondaryLinks](docs/AddonsAddonSecondaryLinks.md)
 - [AddonsAddonSelector](docs/AddonsAddonSelector.md)
 - [AddonsAddonSemantic](docs/AddonsAddonSemantic.md)
 - [AddonsAddonSemanticPaymentMethod](docs/AddonsAddonSemanticPaymentMethod.md)
 - [AddonsAddonSemanticStatus](docs/AddonsAddonSemanticStatus.md)
 - [AddonsBusinessAddon](docs/AddonsBusinessAddon.md)
 - [AddonsButtonBar](docs/AddonsButtonBar.md)
 - [AddonsCreatePostAddonRequest](docs/AddonsCreatePostAddonRequest.md)
 - [AddonsCreatePostAddonV2Body](docs/AddonsCreatePostAddonV2Body.md)
 - [AddonsCreateUserAddonResponseV2](docs/AddonsCreateUserAddonResponseV2.md)
 - [AddonsCreateUserAddonV2Body](docs/AddonsCreateUserAddonV2Body.md)
 - [AddonsDescriptionRow](docs/AddonsDescriptionRow.md)
 - [AddonsEvaluationRow](docs/AddonsEvaluationRow.md)
 - [AddonsEvaluationRowSection](docs/AddonsEvaluationRowSection.md)
 - [AddonsEventRow](docs/AddonsEventRow.md)
 - [AddonsGetBusinessAddonsResponse](docs/AddonsGetBusinessAddonsResponse.md)
 - [AddonsGetDynamicAction](docs/AddonsGetDynamicAction.md)
 - [AddonsGetPostAddonsResponse](docs/AddonsGetPostAddonsResponse.md)
 - [AddonsGetUserAddonsResponse](docs/AddonsGetUserAddonsResponse.md)
 - [AddonsGetUserPostAddonsResponse](docs/AddonsGetUserPostAddonsResponse.md)
 - [AddonsGroupInfoRow](docs/AddonsGroupInfoRow.md)
 - [AddonsGroupInfoRowGroupInfoItem](docs/AddonsGroupInfoRowGroupInfoItem.md)
 - [AddonsImageCarouselRow](docs/AddonsImageCarouselRow.md)
 - [AddonsOpenServerLink](docs/AddonsOpenServerLink.md)
 - [AddonsPostAddon](docs/AddonsPostAddon.md)
 - [AddonsScoreRow](docs/AddonsScoreRow.md)
 - [AddonsSelectorRow](docs/AddonsSelectorRow.md)
 - [AddonsStatus](docs/AddonsStatus.md)
 - [AddonsSubtitleRow](docs/AddonsSubtitleRow.md)
 - [AddonsTitleRow](docs/AddonsTitleRow.md)
 - [AddonsUserAddon](docs/AddonsUserAddon.md)
 - [AddonsUserAddonFilters](docs/AddonsUserAddonFilters.md)
 - [AddonsWidget](docs/AddonsWidget.md)
 - [AddonsWidgetColor](docs/AddonsWidgetColor.md)
 - [AppsApp](docs/AppsApp.md)
 - [AppsAppStatus](docs/AppsAppStatus.md)
 - [AppsAppStatusStatus](docs/AppsAppStatusStatus.md)
 - [AppsServiceTag](docs/AppsServiceTag.md)
 - [AppsServiceType](docs/AppsServiceType.md)
 - [AssetsEnumOption](docs/AssetsEnumOption.md)
 - [AssetsGetBodyStatusesResponse](docs/AssetsGetBodyStatusesResponse.md)
 - [AssetsGetBrandModelsResponse](docs/AssetsGetBrandModelsResponse.md)
 - [AssetsGetCategoriesResponse](docs/AssetsGetCategoriesResponse.md)
 - [AssetsGetCitiesResponse](docs/AssetsGetCitiesResponse.md)
 - [AssetsGetColorsResponse](docs/AssetsGetColorsResponse.md)
 - [AssetsGetDistrictsResponse](docs/AssetsGetDistrictsResponse.md)
 - [AssetsGetInternalStoragesResponse](docs/AssetsGetInternalStoragesResponse.md)
 - [AssetsGetOAuthScopesResponse](docs/AssetsGetOAuthScopesResponse.md)
 - [AssetsGetOAuthScopesResponseAppOauthScope](docs/AssetsGetOAuthScopesResponseAppOauthScope.md)
 - [AssetsGetOAuthScopesResponseLifeCycleState](docs/AssetsGetOAuthScopesResponseLifeCycleState.md)
 - [AssetsGetPermissionsResponse](docs/AssetsGetPermissionsResponse.md)
 - [AssetsGetPermissionsResponseLifeCycleState](docs/AssetsGetPermissionsResponseLifeCycleState.md)
 - [AssetsGetPermissionsResponsePermission](docs/AssetsGetPermissionsResponsePermission.md)
 - [AssetsGetRamMemoriesResponse](docs/AssetsGetRamMemoriesResponse.md)
 - [AssetsGetServiceTypesResponse](docs/AssetsGetServiceTypesResponse.md)
 - [AuthorizationAPICallerInfo](docs/AuthorizationAPICallerInfo.md)
 - [AuthorizationOAuthScope](docs/AuthorizationOAuthScope.md)
 - [ChatAPIChatBotSendMessageBody](docs/ChatAPIChatBotSendMessageBody.md)
 - [ChatAPIConversationSendMessageBody](docs/ChatAPIConversationSendMessageBody.md)
 - [ChatapiChatBotSendMessageResponse](docs/ChatapiChatBotSendMessageResponse.md)
 - [ChatapiChatButton](docs/ChatapiChatButton.md)
 - [ChatapiChatButtonGrid](docs/ChatapiChatButtonGrid.md)
 - [ChatapiChatButtonRow](docs/ChatapiChatButtonRow.md)
 - [ChatapiConversation](docs/ChatapiConversation.md)
 - [ChatapiConversationSendMessageResponse](docs/ChatapiConversationSendMessageResponse.md)
 - [ChatapiConversationType](docs/ChatapiConversationType.md)
 - [ChatapiGenerateUploadTokenResponse](docs/ChatapiGenerateUploadTokenResponse.md)
 - [ChatapiGetConversationResponse](docs/ChatapiGetConversationResponse.md)
 - [ChatapiMessage](docs/ChatapiMessage.md)
 - [ChatapiMessageSenderSide](docs/ChatapiMessageSenderSide.md)
 - [ChatapiMessageSenderType](docs/ChatapiMessageSenderType.md)
 - [ChatapiMessageType](docs/ChatapiMessageType.md)
 - [DivarIconsIconName](docs/DivarIconsIconName.md)
 - [EventsRegisterEventSubscriptionRequest](docs/EventsRegisterEventSubscriptionRequest.md)
 - [EventsRegisterEventSubscriptionRequestEventType](docs/EventsRegisterEventSubscriptionRequestEventType.md)
 - [FinderGetAllDevelopmentPostsResponse](docs/FinderGetAllDevelopmentPostsResponse.md)
 - [FinderGetPostResponse](docs/FinderGetPostResponse.md)
 - [FinderGetUserPostsResponse](docs/FinderGetUserPostsResponse.md)
 - [FinderGetUserPostsResponsePost](docs/FinderGetUserPostsResponsePost.md)
 - [FinderPostExtState](docs/FinderPostExtState.md)
 - [FinderSearchPostItem](docs/FinderSearchPostItem.md)
 - [FinderSearchPostV2Response](docs/FinderSearchPostV2Response.md)
 - [FinderSearchPostsV2Request](docs/FinderSearchPostsV2Request.md)
 - [FinderSearchQuery](docs/FinderSearchQuery.md)
 - [FinderSearchQueryNumberRange](docs/FinderSearchQueryNumberRange.md)
 - [FinderUser](docs/FinderUser.md)
 - [GetPostPricingResponseReorder](docs/GetPostPricingResponseReorder.md)
 - [GetPostResponseBusinessData](docs/GetPostResponseBusinessData.md)
 - [GetServiceTypesResponseServiceTypeData](docs/GetServiceTypesResponseServiceTypeData.md)
 - [GooglerpcStatus](docs/GooglerpcStatus.md)
 - [ImageCarouselRowImageItem](docs/ImageCarouselRowImageItem.md)
 - [ManagementDevelopmentPost](docs/ManagementDevelopmentPost.md)
 - [ManagementPreset](docs/ManagementPreset.md)
 - [MessageFileData](docs/MessageFileData.md)
 - [MessageImageData](docs/MessageImageData.md)
 - [MessageLocationData](docs/MessageLocationData.md)
 - [MessageSender](docs/MessageSender.md)
 - [MessageVideoData](docs/MessageVideoData.md)
 - [MessageVoiceData](docs/MessageVoiceData.md)
 - [OAuthScopeScope](docs/OAuthScopeScope.md)
 - [PaymentGetBalanceResponse](docs/PaymentGetBalanceResponse.md)
 - [PaymentGetPostPricingResponse](docs/PaymentGetPostPricingResponse.md)
 - [PaymentGetTransactionResponse](docs/PaymentGetTransactionResponse.md)
 - [PaymentReorderPostBody](docs/PaymentReorderPostBody.md)
 - [PaymentReorderPostResponse](docs/PaymentReorderPostResponse.md)
 - [PaymentTicketGenerateResponse](docs/PaymentTicketGenerateResponse.md)
 - [PaymentTicketValidateRequest](docs/PaymentTicketValidateRequest.md)
 - [PaymentTicketValidateResponse](docs/PaymentTicketValidateResponse.md)
 - [PaymentTransaction](docs/PaymentTransaction.md)
 - [PaymentTransactionState](docs/PaymentTransactionState.md)
 - [PaymentTransactionType](docs/PaymentTransactionType.md)
 - [PostEditPostBody](docs/PostEditPostBody.md)
 - [PostGetImageUploadURLResponse](docs/PostGetImageUploadURLResponse.md)
 - [PremiumPanelBusinessDataSubBusinessType](docs/PremiumPanelBusinessDataSubBusinessType.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [SearchPostItemPrice](docs/SearchPostItemPrice.md)
 - [SearchPostItemRealEstateFields](docs/SearchPostItemRealEstateFields.md)
 - [SearchPostItemVehiclesFields](docs/SearchPostItemVehiclesFields.md)
 - [SemanticCreatePostSemanticBody](docs/SemanticCreatePostSemanticBody.md)
 - [SemanticCreateUserSemanticBody](docs/SemanticCreateUserSemanticBody.md)
 - [SemanticCreateUserSemanticResponse](docs/SemanticCreateUserSemanticResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### APIKey

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: APIKey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		kenarapi.ContextAPIKeys,
		map[string]kenarapi.APIKey{
			"APIKey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### OAuth


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: oauth.divar.ir/oauth2/auth
- **Scopes**: 
 - **BUSINESS_ADDON_CREATE.{resource_id}**: BUSINESS_ADDON_CREATE.{resource_id}
 - **CHAT_BOT_USER_MESSAGE_SEND**: CHAT_BOT_USER_MESSAGE_SEND
 - **CHAT_CONVERSATION_READ.{resource_id}**: CHAT_CONVERSATION_READ.{resource_id}
 - **CHAT_MESSAGE_SEND.{resource_id}**: CHAT_MESSAGE_SEND.{resource_id}
 - **CHAT_POST_CONVERSATIONS_MESSAGE_SEND.{resource_id}**: CHAT_POST_CONVERSATIONS_MESSAGE_SEND.{resource_id}
 - **CHAT_POST_CONVERSATIONS_READ.{resource_id}**: CHAT_POST_CONVERSATIONS_READ.{resource_id}
 - **CHAT_SUPPLIER_ALL_CONVERSATIONS_MESSAGE_SEND**: CHAT_SUPPLIER_ALL_CONVERSATIONS_MESSAGE_SEND
 - **CHAT_SUPPLIER_ALL_CONVERSATIONS_READ**: CHAT_SUPPLIER_ALL_CONVERSATIONS_READ
 - **CONVERSATION_SEND_MESSAGE.{resource_id}**: CONVERSATION_SEND_MESSAGE.{resource_id}
 - **MANAGEMENT_APPS_READ.{resource_id}**: MANAGEMENT_APPS_READ.{resource_id}
 - **MANAGEMENT_APPS_WRITE.{resource_id}**: MANAGEMENT_APPS_WRITE.{resource_id}
 - **NOTIFICATION_ACCESS_REVOCATION**: NOTIFICATION_ACCESS_REVOCATION
 - **PAYMENT_ALL_POSTS_PRICING_READ**: PAYMENT_ALL_POSTS_PRICING_READ
 - **PAYMENT_ALL_POSTS_REORDER**: PAYMENT_ALL_POSTS_REORDER
 - **POST_ADDON_CREATE.{resource_id}**: POST_ADDON_CREATE.{resource_id}
 - **POST_EDIT.{resource_id}**: POST_EDIT.{resource_id}
 - **POST_ONGOING_IMAGES_GET.{resource_id}**: POST_ONGOING_IMAGES_GET.{resource_id}
 - **POST_SEMANTIC_CREATE.{resource_id}**: POST_SEMANTIC_CREATE.{resource_id}
 - **USER_ADDON_CREATE**: USER_ADDON_CREATE
 - **USER_ID**: USER_ID
 - **USER_PHONE**: USER_PHONE
 - **USER_POSTS_ADDON_CREATE**: USER_POSTS_ADDON_CREATE
 - **USER_POSTS_GET**: USER_POSTS_GET
 - **USER_VERIFICATION_CREATE**: USER_VERIFICATION_CREATE
 - **offline_access**: offline_access
 - **openid**: openid

Example

```go
auth := context.WithValue(context.Background(), kenarapi.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, kenarapi.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



