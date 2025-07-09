# GetUserInfo200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleListOverride** | [**RoleListOverride**](RoleListOverride.md) |  | 
**AuthenticationBackends** | **[]string** |  | 
**ProviderProperties** | [**map[string]GetUserInfo200ResponseDataProviderPropertiesValue**](GetUserInfo200ResponseDataProviderPropertiesValue.md) |  | 
**Users** | [**[]User**](User.md) |  | 
**Digest** | **string** | Default password hash algorithm | 
**TenantsEnabled** | **bool** | Whether the tenants plugin is enabled or not, indicating if users only belong to their assigned tenants | 

## Methods

### NewGetUserInfo200ResponseData

`func NewGetUserInfo200ResponseData(roleListOverride RoleListOverride, authenticationBackends []string, providerProperties map[string]GetUserInfo200ResponseDataProviderPropertiesValue, users []User, digest string, tenantsEnabled bool, ) *GetUserInfo200ResponseData`

NewGetUserInfo200ResponseData instantiates a new GetUserInfo200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserInfo200ResponseDataWithDefaults

`func NewGetUserInfo200ResponseDataWithDefaults() *GetUserInfo200ResponseData`

NewGetUserInfo200ResponseDataWithDefaults instantiates a new GetUserInfo200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleListOverride

`func (o *GetUserInfo200ResponseData) GetRoleListOverride() RoleListOverride`

GetRoleListOverride returns the RoleListOverride field if non-nil, zero value otherwise.

### GetRoleListOverrideOk

`func (o *GetUserInfo200ResponseData) GetRoleListOverrideOk() (*RoleListOverride, bool)`

GetRoleListOverrideOk returns a tuple with the RoleListOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleListOverride

`func (o *GetUserInfo200ResponseData) SetRoleListOverride(v RoleListOverride)`

SetRoleListOverride sets RoleListOverride field to given value.


### GetAuthenticationBackends

`func (o *GetUserInfo200ResponseData) GetAuthenticationBackends() []string`

GetAuthenticationBackends returns the AuthenticationBackends field if non-nil, zero value otherwise.

### GetAuthenticationBackendsOk

`func (o *GetUserInfo200ResponseData) GetAuthenticationBackendsOk() (*[]string, bool)`

GetAuthenticationBackendsOk returns a tuple with the AuthenticationBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBackends

`func (o *GetUserInfo200ResponseData) SetAuthenticationBackends(v []string)`

SetAuthenticationBackends sets AuthenticationBackends field to given value.


### GetProviderProperties

`func (o *GetUserInfo200ResponseData) GetProviderProperties() map[string]GetUserInfo200ResponseDataProviderPropertiesValue`

GetProviderProperties returns the ProviderProperties field if non-nil, zero value otherwise.

### GetProviderPropertiesOk

`func (o *GetUserInfo200ResponseData) GetProviderPropertiesOk() (*map[string]GetUserInfo200ResponseDataProviderPropertiesValue, bool)`

GetProviderPropertiesOk returns a tuple with the ProviderProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderProperties

`func (o *GetUserInfo200ResponseData) SetProviderProperties(v map[string]GetUserInfo200ResponseDataProviderPropertiesValue)`

SetProviderProperties sets ProviderProperties field to given value.


### GetUsers

`func (o *GetUserInfo200ResponseData) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetUserInfo200ResponseData) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetUserInfo200ResponseData) SetUsers(v []User)`

SetUsers sets Users field to given value.


### GetDigest

`func (o *GetUserInfo200ResponseData) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *GetUserInfo200ResponseData) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *GetUserInfo200ResponseData) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetTenantsEnabled

`func (o *GetUserInfo200ResponseData) GetTenantsEnabled() bool`

GetTenantsEnabled returns the TenantsEnabled field if non-nil, zero value otherwise.

### GetTenantsEnabledOk

`func (o *GetUserInfo200ResponseData) GetTenantsEnabledOk() (*bool, bool)`

GetTenantsEnabledOk returns a tuple with the TenantsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantsEnabled

`func (o *GetUserInfo200ResponseData) SetTenantsEnabled(v bool)`

SetTenantsEnabled sets TenantsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


