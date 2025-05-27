# \AssetsAPI

All URIs are relative to *https://open-api.divar.ir*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsGetBodyStatuses**](AssetsAPI.md#AssetsGetBodyStatuses) | **Get** /v1/open-platform/assets/body-status | دریافت گزینه‌های وضعیت بدنه موجود در دسته‌بندی‌های خودرو دیوار
[**AssetsGetBrandModels**](AssetsAPI.md#AssetsGetBrandModels) | **Get** /v1/open-platform/assets/brand-model/{category} | دریافت مدل‌های برند در دسته‌بندی دیوار
[**AssetsGetCategories**](AssetsAPI.md#AssetsGetCategories) | **Get** /v1/open-platform/assets/category | دریافت دسته‌بندی‌های دیوار
[**AssetsGetCities**](AssetsAPI.md#AssetsGetCities) | **Get** /v1/open-platform/assets/city | دریافت شهرهای دیوار
[**AssetsGetColors**](AssetsAPI.md#AssetsGetColors) | **Get** /v1/open-platform/assets/color/{category} | دریافت رنگ‌ها در دسته‌بندی دیوار
[**AssetsGetDistricts**](AssetsAPI.md#AssetsGetDistricts) | **Get** /v1/open-platform/assets/district | دریافت مناطق دیوار
[**AssetsGetDistricts2**](AssetsAPI.md#AssetsGetDistricts2) | **Get** /v1/open-platform/assets/district/{city_slug} | دریافت مناطق دیوار
[**AssetsGetInternalStorages**](AssetsAPI.md#AssetsGetInternalStorages) | **Get** /v1/open-platform/assets/internal-storage | دریافت گزینه‌های حافظه داخلی موجود در دسته‌بندی‌های موبایل/تبلت/لپ‌تاپ دیوار
[**AssetsGetOAuthScopes**](AssetsAPI.md#AssetsGetOAuthScopes) | **Get** /v1/open-platform/assets/oauth-scope | دریافت دامنه‌های OAuth کنار دیوار
[**AssetsGetPermissions**](AssetsAPI.md#AssetsGetPermissions) | **Get** /v1/open-platform/assets/permission | دریافت مجوزهای کنار دیوار
[**AssetsGetRamMemories**](AssetsAPI.md#AssetsGetRamMemories) | **Get** /v1/open-platform/assets/ram-memory | دریافت گزینه‌های حافظه رم موجود در دسته‌بندی‌های موبایل/تبلت/لپ‌تاپ دیوار
[**AssetsGetServiceTypes**](AssetsAPI.md#AssetsGetServiceTypes) | **Get** /v1/open-platform/assets/service-type | دریافت انواع سرویس موجود در کنار دیوار



## AssetsGetBodyStatuses

> AssetsGetBodyStatusesResponse AssetsGetBodyStatuses(ctx).Execute()

دریافت گزینه‌های وضعیت بدنه موجود در دسته‌بندی‌های خودرو دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetBrandModels

> AssetsGetBrandModelsResponse AssetsGetBrandModels(ctx, category).Execute()

دریافت مدل‌های برند در دسته‌بندی دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetCategories

> AssetsGetCategoriesResponse AssetsGetCategories(ctx).Execute()

دریافت دسته‌بندی‌های دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetCities

> AssetsGetCitiesResponse AssetsGetCities(ctx).Execute()

دریافت شهرهای دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetColors

> AssetsGetColorsResponse AssetsGetColors(ctx, category).Execute()

دریافت رنگ‌ها در دسته‌بندی دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetDistricts

> AssetsGetDistrictsResponse AssetsGetDistricts(ctx).CitySlug(citySlug).Execute()

دریافت مناطق دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetDistricts2

> AssetsGetDistrictsResponse AssetsGetDistricts2(ctx, citySlug).Execute()

دریافت مناطق دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetInternalStorages

> AssetsGetInternalStoragesResponse AssetsGetInternalStorages(ctx).Execute()

دریافت گزینه‌های حافظه داخلی موجود در دسته‌بندی‌های موبایل/تبلت/لپ‌تاپ دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetOAuthScopes

> AssetsGetOAuthScopesResponse AssetsGetOAuthScopes(ctx).Execute()

دریافت دامنه‌های OAuth کنار دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetPermissions

> AssetsGetPermissionsResponse AssetsGetPermissions(ctx).Execute()

دریافت مجوزهای کنار دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetRamMemories

> AssetsGetRamMemoriesResponse AssetsGetRamMemories(ctx).Execute()

دریافت گزینه‌های حافظه رم موجود در دسته‌بندی‌های موبایل/تبلت/لپ‌تاپ دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsGetServiceTypes

> AssetsGetServiceTypesResponse AssetsGetServiceTypes(ctx).Execute()

دریافت انواع سرویس موجود در کنار دیوار



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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

