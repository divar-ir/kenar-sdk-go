# \AssetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsGetBodyStatuses**](AssetsAPI.md#AssetsGetBodyStatuses) | **Get** /v1/open-platform/assets/body-status | Get available body-status options in Divar vehicle categories
[**AssetsGetBrandModels**](AssetsAPI.md#AssetsGetBrandModels) | **Get** /v1/open-platform/assets/brand-model/{category} | Get brand models in a Divar category
[**AssetsGetCategories**](AssetsAPI.md#AssetsGetCategories) | **Get** /v1/open-platform/assets/category | Get Divar categories
[**AssetsGetCities**](AssetsAPI.md#AssetsGetCities) | **Get** /v1/open-platform/assets/city | Get Divar cities
[**AssetsGetColors**](AssetsAPI.md#AssetsGetColors) | **Get** /v1/open-platform/assets/color/{category} | Get colors in a Divar category
[**AssetsGetDistricts**](AssetsAPI.md#AssetsGetDistricts) | **Get** /v1/open-platform/assets/district | Get Divar districts
[**AssetsGetDistricts2**](AssetsAPI.md#AssetsGetDistricts2) | **Get** /v1/open-platform/assets/district/{city_slug} | Get Divar districts
[**AssetsGetInternalStorages**](AssetsAPI.md#AssetsGetInternalStorages) | **Get** /v1/open-platform/assets/internal-storage | Get available internal-storage options in Divar mobile/tablet/laptop categories
[**AssetsGetOAuthScopes**](AssetsAPI.md#AssetsGetOAuthScopes) | **Get** /v1/open-platform/assets/oauth-scope | Get KenarDivar OAuth scopes
[**AssetsGetPermissions**](AssetsAPI.md#AssetsGetPermissions) | **Get** /v1/open-platform/assets/permission | Get KenarDivar permissions
[**AssetsGetRamMemories**](AssetsAPI.md#AssetsGetRamMemories) | **Get** /v1/open-platform/assets/ram-memory | Get available ram-memory options in Divar mobile/tablet/laptop categories
[**AssetsGetServiceTypes**](AssetsAPI.md#AssetsGetServiceTypes) | **Get** /v1/open-platform/assets/service-type | Get available service types in KenarDivar



## AssetsGetBodyStatuses

> AssetsGetBodyStatusesResponse AssetsGetBodyStatuses(ctx).Execute()

Get available body-status options in Divar vehicle categories



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
	resp, r, err := apiClient.AssetsAPI.AssetsGetBodyStatuses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetBodyStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetBodyStatuses`: AssetsGetBodyStatusesResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetBodyStatuses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetBodyStatusesRequest struct via the builder pattern


### Return type

[**AssetsGetBodyStatusesResponse**](AssetsGetBodyStatusesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetBrandModels

> AssetsGetBrandModelsResponse AssetsGetBrandModels(ctx, category).Execute()

Get brand models in a Divar category



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
	category := "category_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AssetsGetBrandModels(context.Background(), category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetBrandModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetBrandModels`: AssetsGetBrandModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetBrandModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetBrandModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetsGetBrandModelsResponse**](AssetsGetBrandModelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetCategories

> AssetsGetCategoriesResponse AssetsGetCategories(ctx).Execute()

Get Divar categories



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
	resp, r, err := apiClient.AssetsAPI.AssetsGetCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetCategories`: AssetsGetCategoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetCategoriesRequest struct via the builder pattern


### Return type

[**AssetsGetCategoriesResponse**](AssetsGetCategoriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetCities

> AssetsGetCitiesResponse AssetsGetCities(ctx).Execute()

Get Divar cities



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
	resp, r, err := apiClient.AssetsAPI.AssetsGetCities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetCities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetCities`: AssetsGetCitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetCities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetCitiesRequest struct via the builder pattern


### Return type

[**AssetsGetCitiesResponse**](AssetsGetCitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetColors

> AssetsGetColorsResponse AssetsGetColors(ctx, category).Execute()

Get colors in a Divar category



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
	category := "category_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AssetsGetColors(context.Background(), category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetColors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetColors`: AssetsGetColorsResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetColors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetColorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetsGetColorsResponse**](AssetsGetColorsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetDistricts

> AssetsGetDistrictsResponse AssetsGetDistricts(ctx).CitySlug(citySlug).Execute()

Get Divar districts



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
	citySlug := "citySlug_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AssetsGetDistricts(context.Background()).CitySlug(citySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetDistricts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetDistricts`: AssetsGetDistrictsResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetDistricts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetDistrictsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citySlug** | **string** |  | 

### Return type

[**AssetsGetDistrictsResponse**](AssetsGetDistrictsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetDistricts2

> AssetsGetDistrictsResponse AssetsGetDistricts2(ctx, citySlug).Execute()

Get Divar districts



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
	citySlug := "citySlug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AssetsGetDistricts2(context.Background(), citySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetDistricts2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetDistricts2`: AssetsGetDistrictsResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetDistricts2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**citySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetDistricts2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetsGetDistrictsResponse**](AssetsGetDistrictsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetInternalStorages

> AssetsGetInternalStoragesResponse AssetsGetInternalStorages(ctx).Execute()

Get available internal-storage options in Divar mobile/tablet/laptop categories



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
	resp, r, err := apiClient.AssetsAPI.AssetsGetInternalStorages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetInternalStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetInternalStorages`: AssetsGetInternalStoragesResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetInternalStorages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetInternalStoragesRequest struct via the builder pattern


### Return type

[**AssetsGetInternalStoragesResponse**](AssetsGetInternalStoragesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetOAuthScopes

> AssetsGetOAuthScopesResponse AssetsGetOAuthScopes(ctx).Execute()

Get KenarDivar OAuth scopes



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
	resp, r, err := apiClient.AssetsAPI.AssetsGetOAuthScopes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetOAuthScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetOAuthScopes`: AssetsGetOAuthScopesResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetOAuthScopes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetOAuthScopesRequest struct via the builder pattern


### Return type

[**AssetsGetOAuthScopesResponse**](AssetsGetOAuthScopesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetPermissions

> AssetsGetPermissionsResponse AssetsGetPermissions(ctx).Execute()

Get KenarDivar permissions



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
	resp, r, err := apiClient.AssetsAPI.AssetsGetPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetPermissions`: AssetsGetPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetPermissionsRequest struct via the builder pattern


### Return type

[**AssetsGetPermissionsResponse**](AssetsGetPermissionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetRamMemories

> AssetsGetRamMemoriesResponse AssetsGetRamMemories(ctx).Execute()

Get available ram-memory options in Divar mobile/tablet/laptop categories



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
	resp, r, err := apiClient.AssetsAPI.AssetsGetRamMemories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetRamMemories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetRamMemories`: AssetsGetRamMemoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetRamMemories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetRamMemoriesRequest struct via the builder pattern


### Return type

[**AssetsGetRamMemoriesResponse**](AssetsGetRamMemoriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetServiceTypes

> AssetsGetServiceTypesResponse AssetsGetServiceTypes(ctx).Execute()

Get available service types in KenarDivar



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
	resp, r, err := apiClient.AssetsAPI.AssetsGetServiceTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AssetsGetServiceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsGetServiceTypes`: AssetsGetServiceTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AssetsGetServiceTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsGetServiceTypesRequest struct via the builder pattern


### Return type

[**AssetsGetServiceTypesResponse**](AssetsGetServiceTypesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

