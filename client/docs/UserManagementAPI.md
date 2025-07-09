# \UserManagementAPI

All URIs are relative to *https://rudder.example.com/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUser**](UserManagementAPI.md#ActivateUser) | **Put** /usermanagement/status/activate/{username} | Activate user
[**AddUser**](UserManagementAPI.md#AddUser) | **Post** /usermanagement | Add user
[**DeleteUser**](UserManagementAPI.md#DeleteUser) | **Delete** /usermanagement/{username} | Delete an user
[**DisableUser**](UserManagementAPI.md#DisableUser) | **Put** /usermanagement/status/disable/{username} | Disable user
[**GetRole**](UserManagementAPI.md#GetRole) | **Get** /usermanagement/roles | List all roles
[**GetUserInfo**](UserManagementAPI.md#GetUserInfo) | **Get** /usermanagement/users | List all users
[**ReloadUserConf**](UserManagementAPI.md#ReloadUserConf) | **Post** /usermanagement/users/reload | Reload users
[**RoleCoverage**](UserManagementAPI.md#RoleCoverage) | **Post** /usermanagement/coverage/{username} | Compute the role coverage
[**UpdateUser**](UserManagementAPI.md#UpdateUser) | **Post** /usermanagement/update/{username} | Update user access to Rudder
[**UpdateUserInfo**](UserManagementAPI.md#UpdateUserInfo) | **Post** /usermanagement/update/info/{username} | Update user information



## ActivateUser

> ActivateUser200Response ActivateUser(ctx, username).Execute()

Activate user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {
	username := "JaneDoe" // string | Username of an user (unique and must not contain whitespace)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ActivateUser(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ActivateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateUser`: ActivateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ActivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of an user (unique and must not contain whitespace) | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivateUser200Response**](ActivateUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUser

> AddUser200Response AddUser(ctx).UserNew(userNew).Execute()

Add user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {
	userNew := *openapiclient.NewUserNew("johndoe", "passwdWillBeStoredHashed", []string{"admin"}, false) // UserNew | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.AddUser(context.Background()).UserNew(userNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.AddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUser`: AddUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.AddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userNew** | [**UserNew**](UserNew.md) |  | 

### Return type

[**AddUser200Response**](AddUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser200Response DeleteUser(ctx, username).Execute()

Delete an user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {
	username := "JaneDoe" // string | Username of an user (unique and must not contain whitespace)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.DeleteUser(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUser`: DeleteUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of an user (unique and must not contain whitespace) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteUser200Response**](DeleteUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableUser

> DisableUser200Response DisableUser(ctx, username).Execute()

Disable user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {
	username := "JaneDoe" // string | Username of an user (unique and must not contain whitespace)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.DisableUser(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DisableUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableUser`: DisableUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.DisableUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of an user (unique and must not contain whitespace) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisableUser200Response**](DisableUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> GetRole200Response GetRole(ctx).Execute()

List all roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.GetRole(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: GetRole200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


### Return type

[**GetRole200Response**](GetRole200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInfo

> GetUserInfo200Response GetUserInfo(ctx).Execute()

List all users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.GetUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.GetUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfo`: GetUserInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.GetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoRequest struct via the builder pattern


### Return type

[**GetUserInfo200Response**](GetUserInfo200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadUserConf

> ReloadUserConf200Response ReloadUserConf(ctx).Execute()

Reload users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.ReloadUserConf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ReloadUserConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadUserConf`: ReloadUserConf200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ReloadUserConf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReloadUserConfRequest struct via the builder pattern


### Return type

[**ReloadUserConf200Response**](ReloadUserConf200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleCoverage

> RoleCoverage200Response RoleCoverage(ctx, username).RoleCoverageRequest(roleCoverageRequest).Execute()

Compute the role coverage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {
	username := "JaneDoe" // string | Username of an user (unique and must not contain whitespace)
	roleCoverageRequest := *openapiclient.NewRoleCoverageRequest([]string{"Permissions_example"}, []string{"Authz_example"}) // RoleCoverageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.RoleCoverage(context.Background(), username).RoleCoverageRequest(roleCoverageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.RoleCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoleCoverage`: RoleCoverage200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.RoleCoverage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of an user (unique and must not contain whitespace) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRoleCoverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleCoverageRequest** | [**RoleCoverageRequest**](RoleCoverageRequest.md) |  | 

### Return type

[**RoleCoverage200Response**](RoleCoverage200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser200Response UpdateUser(ctx, username).UserFile(userFile).Execute()

Update user access to Rudder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {
	username := "JaneDoe" // string | Username of an user (unique and must not contain whitespace)
	userFile := *openapiclient.NewUserFile("johndoe", "passwdWillBeStoredHashed", []string{"admin"}, false) // UserFile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.UpdateUser(context.Background(), username).UserFile(userFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: UpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of an user (unique and must not contain whitespace) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userFile** | [**UserFile**](UserFile.md) |  | 

### Return type

[**UpdateUser200Response**](UpdateUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserInfo

> UpdateUserInfo200Response UpdateUserInfo(ctx, username).UserInfo(userInfo).Execute()

Update user information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudducoeur/terraform-rudder-provider/client"
)

func main() {
	username := "JaneDoe" // string | Username of an user (unique and must not contain whitespace)
	userInfo := *openapiclient.NewUserInfo() // UserInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.UpdateUserInfo(context.Background(), username).UserInfo(userInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UpdateUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserInfo`: UpdateUserInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.UpdateUserInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of an user (unique and must not contain whitespace) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userInfo** | [**UserInfo**](UserInfo.md) |  | 

### Return type

[**UpdateUserInfo200Response**](UpdateUserInfo200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

