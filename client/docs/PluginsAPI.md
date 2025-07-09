# \PluginsAPI

All URIs are relative to *https://rudder.example.com/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPluginsInfo**](PluginsAPI.md#GetPluginsInfo) | **Get** /plugins/info | Information about installed plugins
[**PluginSettings**](PluginsAPI.md#PluginSettings) | **Get** /plugins/settings | Get plugins repository settings
[**UpdateSettings**](PluginsAPI.md#UpdateSettings) | **Post** /plugins/settings | Update plugins settings



## GetPluginsInfo

> GetPluginsInfo200Response GetPluginsInfo(ctx).Execute()

Information about installed plugins



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.GetPluginsInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPluginsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPluginsInfo`: GetPluginsInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPluginsInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginsInfoRequest struct via the builder pattern


### Return type

[**GetPluginsInfo200Response**](GetPluginsInfo200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginSettings

> PluginSettings200Response PluginSettings(ctx).Execute()

Get plugins repository settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginSettings`: PluginSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPluginSettingsRequest struct via the builder pattern


### Return type

[**PluginSettings200Response**](PluginSettings200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettings

> UpdateSettings200Response UpdateSettings(ctx).Execute()

Update plugins settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.UpdateSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.UpdateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSettings`: UpdateSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.UpdateSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsRequest struct via the builder pattern


### Return type

[**UpdateSettings200Response**](UpdateSettings200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

