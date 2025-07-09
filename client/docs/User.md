# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**OtherInfo** | **map[string]string** |  | 
**Status** | **string** |  | 
**Authz** | **[]string** |  | 
**Permissions** | **[]string** |  | 
**RolesCoverage** | **[]string** |  | 
**CustomRights** | **[]string** |  | 
**Providers** | **[]string** |  | 
**ProvidersInfo** | [**UserProvidersInfo**](UserProvidersInfo.md) |  | 
**Tenants** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUser

`func NewUser(login string, otherInfo map[string]string, status string, authz []string, permissions []string, rolesCoverage []string, customRights []string, providers []string, providersInfo UserProvidersInfo, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *User) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *User) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *User) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOtherInfo

`func (o *User) GetOtherInfo() map[string]string`

GetOtherInfo returns the OtherInfo field if non-nil, zero value otherwise.

### GetOtherInfoOk

`func (o *User) GetOtherInfoOk() (*map[string]string, bool)`

GetOtherInfoOk returns a tuple with the OtherInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherInfo

`func (o *User) SetOtherInfo(v map[string]string)`

SetOtherInfo sets OtherInfo field to given value.


### GetStatus

`func (o *User) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAuthz

`func (o *User) GetAuthz() []string`

GetAuthz returns the Authz field if non-nil, zero value otherwise.

### GetAuthzOk

`func (o *User) GetAuthzOk() (*[]string, bool)`

GetAuthzOk returns a tuple with the Authz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthz

`func (o *User) SetAuthz(v []string)`

SetAuthz sets Authz field to given value.


### GetPermissions

`func (o *User) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *User) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *User) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetRolesCoverage

`func (o *User) GetRolesCoverage() []string`

GetRolesCoverage returns the RolesCoverage field if non-nil, zero value otherwise.

### GetRolesCoverageOk

`func (o *User) GetRolesCoverageOk() (*[]string, bool)`

GetRolesCoverageOk returns a tuple with the RolesCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesCoverage

`func (o *User) SetRolesCoverage(v []string)`

SetRolesCoverage sets RolesCoverage field to given value.


### GetCustomRights

`func (o *User) GetCustomRights() []string`

GetCustomRights returns the CustomRights field if non-nil, zero value otherwise.

### GetCustomRightsOk

`func (o *User) GetCustomRightsOk() (*[]string, bool)`

GetCustomRightsOk returns a tuple with the CustomRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRights

`func (o *User) SetCustomRights(v []string)`

SetCustomRights sets CustomRights field to given value.


### GetProviders

`func (o *User) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *User) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *User) SetProviders(v []string)`

SetProviders sets Providers field to given value.


### GetProvidersInfo

`func (o *User) GetProvidersInfo() UserProvidersInfo`

GetProvidersInfo returns the ProvidersInfo field if non-nil, zero value otherwise.

### GetProvidersInfoOk

`func (o *User) GetProvidersInfoOk() (*UserProvidersInfo, bool)`

GetProvidersInfoOk returns a tuple with the ProvidersInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidersInfo

`func (o *User) SetProvidersInfo(v UserProvidersInfo)`

SetProvidersInfo sets ProvidersInfo field to given value.


### GetTenants

`func (o *User) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *User) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *User) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *User) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetLastLogin

`func (o *User) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


