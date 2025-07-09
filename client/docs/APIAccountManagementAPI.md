# \APIAccountManagementAPI

All URIs are relative to *https://rudder.example.com/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](APIAccountManagementAPI.md#CreateAccount) | **Post** /apiaccounts | Create an API account
[**DeleteAccount**](APIAccountManagementAPI.md#DeleteAccount) | **Delete** /apiaccounts/${apiAccountId} | Delete an API account
[**GetAccount**](APIAccountManagementAPI.md#GetAccount) | **Get** /apiaccounts/${apiAccountId} | Get an API account details
[**GetAllAccounts**](APIAccountManagementAPI.md#GetAllAccounts) | **Get** /apiaccounts | List all API accounts
[**RegenerateToken**](APIAccountManagementAPI.md#RegenerateToken) | **Post** /apiaccounts/${apiAccountId}/regenerate | Regenerate the token of an API account
[**UpdateAccount**](APIAccountManagementAPI.md#UpdateAccount) | **Post** /apiaccounts/${apiAccountId} | Update an API account details



## CreateAccount

> CreateAccount200Response CreateAccount(ctx).ApiAccountNew(apiAccountNew).Execute()

Create an API account



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
	apiAccountNew := *openapiclient.NewApiAccountNew() // ApiAccountNew |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAccountManagementAPI.CreateAccount(context.Background()).ApiAccountNew(apiAccountNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAccountManagementAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: CreateAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAccountManagementAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAccountNew** | [**ApiAccountNew**](ApiAccountNew.md) |  | 

### Return type

[**CreateAccount200Response**](CreateAccount200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteAccount200Response DeleteAccount(ctx, apiAccountId).Execute()

Delete an API account



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
	apiAccountId := "1ebee8c7-c898-4ee5-9470-139bfd80c442" // string | Id of the API account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAccountManagementAPI.DeleteAccount(context.Background(), apiAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAccountManagementAPI.DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccount`: DeleteAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAccountManagementAPI.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAccountId** | **string** | Id of the API account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAccount200Response**](DeleteAccount200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> GetAccount200Response GetAccount(ctx, apiAccountId).Execute()

Get an API account details



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
	apiAccountId := "1ebee8c7-c898-4ee5-9470-139bfd80c442" // string | Id of the API account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAccountManagementAPI.GetAccount(context.Background(), apiAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAccountManagementAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: GetAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAccountManagementAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAccountId** | **string** | Id of the API account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccount200Response**](GetAccount200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAccounts

> GetAllAccounts200Response GetAllAccounts(ctx).Execute()

List all API accounts



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
	resp, r, err := apiClient.APIAccountManagementAPI.GetAllAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAccountManagementAPI.GetAllAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAccounts`: GetAllAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAccountManagementAPI.GetAllAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccountsRequest struct via the builder pattern


### Return type

[**GetAllAccounts200Response**](GetAllAccounts200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateToken

> RegenerateToken200Response RegenerateToken(ctx, apiAccountId).Execute()

Regenerate the token of an API account



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
	apiAccountId := "1ebee8c7-c898-4ee5-9470-139bfd80c442" // string | Id of the API account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAccountManagementAPI.RegenerateToken(context.Background(), apiAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAccountManagementAPI.RegenerateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateToken`: RegenerateToken200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAccountManagementAPI.RegenerateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAccountId** | **string** | Id of the API account | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegenerateToken200Response**](RegenerateToken200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> UpdateAccount200Response UpdateAccount(ctx, apiAccountId).ApiAccountUpdate(apiAccountUpdate).Execute()

Update an API account details



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
	apiAccountId := "1ebee8c7-c898-4ee5-9470-139bfd80c442" // string | Id of the API account
	apiAccountUpdate := *openapiclient.NewApiAccountUpdate() // ApiAccountUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAccountManagementAPI.UpdateAccount(context.Background(), apiAccountId).ApiAccountUpdate(apiAccountUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAccountManagementAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccount`: UpdateAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAccountManagementAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAccountId** | **string** | Id of the API account | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAccountUpdate** | [**ApiAccountUpdate**](ApiAccountUpdate.md) |  | 

### Return type

[**UpdateAccount200Response**](UpdateAccount200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

