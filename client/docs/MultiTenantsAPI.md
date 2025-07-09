# \MultiTenantsAPI

All URIs are relative to *https://rudder.example.com/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTenants**](MultiTenantsAPI.md#AddTenants) | **Post** /tenants | Add one or more tenants
[**DeleteNodeSecurityTag**](MultiTenantsAPI.md#DeleteNodeSecurityTag) | **Delete** /securitytags/nodes | Remove tenant value on nodes
[**DeleteTenants**](MultiTenantsAPI.md#DeleteTenants) | **Delete** /tenants | Delete a list of tenants
[**GetTenant**](MultiTenantsAPI.md#GetTenant) | **Get** /tenants/{id} | Get tenant details
[**GetTenants**](MultiTenantsAPI.md#GetTenants) | **Get** /tenants | Get the list of all tenants.
[**SetNodeSecurityTags**](MultiTenantsAPI.md#SetNodeSecurityTags) | **Post** /securitytags/nodes | Define a tenant on nodes



## AddTenants

> AddTenants200Response AddTenants(ctx).Tenant(tenant).Execute()

Add one or more tenants



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
	tenant := []openapiclient.Tenant{*openapiclient.NewTenant()} // []Tenant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiTenantsAPI.AddTenants(context.Background()).Tenant(tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiTenantsAPI.AddTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTenants`: AddTenants200Response
	fmt.Fprintf(os.Stdout, "Response from `MultiTenantsAPI.AddTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant** | [**[]Tenant**](Tenant.md) |  | 

### Return type

[**AddTenants200Response**](AddTenants200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNodeSecurityTag

> DeleteNodeSecurityTag200Response DeleteNodeSecurityTag(ctx).RequestBody(requestBody).Execute()

Remove tenant value on nodes



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
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiTenantsAPI.DeleteNodeSecurityTag(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiTenantsAPI.DeleteNodeSecurityTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNodeSecurityTag`: DeleteNodeSecurityTag200Response
	fmt.Fprintf(os.Stdout, "Response from `MultiTenantsAPI.DeleteNodeSecurityTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeSecurityTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**DeleteNodeSecurityTag200Response**](DeleteNodeSecurityTag200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenants

> DeleteTenants200Response DeleteTenants(ctx).Execute()

Delete a list of tenants



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
	resp, r, err := apiClient.MultiTenantsAPI.DeleteTenants(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiTenantsAPI.DeleteTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenants`: DeleteTenants200Response
	fmt.Fprintf(os.Stdout, "Response from `MultiTenantsAPI.DeleteTenants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantsRequest struct via the builder pattern


### Return type

[**DeleteTenants200Response**](DeleteTenants200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenant

> GetTenants200Response GetTenant(ctx, id).Execute()

Get tenant details



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
	id := "zone-1" // string | Id of the tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiTenantsAPI.GetTenant(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiTenantsAPI.GetTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenant`: GetTenants200Response
	fmt.Fprintf(os.Stdout, "Response from `MultiTenantsAPI.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTenants200Response**](GetTenants200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenants

> GetTenants200Response GetTenants(ctx).Execute()

Get the list of all tenants.



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
	resp, r, err := apiClient.MultiTenantsAPI.GetTenants(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiTenantsAPI.GetTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenants`: GetTenants200Response
	fmt.Fprintf(os.Stdout, "Response from `MultiTenantsAPI.GetTenants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsRequest struct via the builder pattern


### Return type

[**GetTenants200Response**](GetTenants200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNodeSecurityTags

> SetNodeSecurityTags200Response SetNodeSecurityTags(ctx).SetNodeSecurityTagsRequestInner(setNodeSecurityTagsRequestInner).Execute()

Define a tenant on nodes



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
	setNodeSecurityTagsRequestInner := []openapiclient.SetNodeSecurityTagsRequestInner{*openapiclient.NewSetNodeSecurityTagsRequestInner("TenantId_example", []string{"NodeIds_example"})} // []SetNodeSecurityTagsRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiTenantsAPI.SetNodeSecurityTags(context.Background()).SetNodeSecurityTagsRequestInner(setNodeSecurityTagsRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiTenantsAPI.SetNodeSecurityTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodeSecurityTags`: SetNodeSecurityTags200Response
	fmt.Fprintf(os.Stdout, "Response from `MultiTenantsAPI.SetNodeSecurityTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNodeSecurityTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setNodeSecurityTagsRequestInner** | [**[]SetNodeSecurityTagsRequestInner**](SetNodeSecurityTagsRequestInner.md) |  | 

### Return type

[**SetNodeSecurityTags200Response**](SetNodeSecurityTags200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

