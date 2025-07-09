# UserProvidersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**Authz** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**CustomRights** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserProvidersInfo

`func NewUserProvidersInfo() *UserProvidersInfo`

NewUserProvidersInfo instantiates a new UserProvidersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProvidersInfoWithDefaults

`func NewUserProvidersInfoWithDefaults() *UserProvidersInfo`

NewUserProvidersInfoWithDefaults instantiates a new UserProvidersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *UserProvidersInfo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserProvidersInfo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserProvidersInfo) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserProvidersInfo) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetAuthz

`func (o *UserProvidersInfo) GetAuthz() []string`

GetAuthz returns the Authz field if non-nil, zero value otherwise.

### GetAuthzOk

`func (o *UserProvidersInfo) GetAuthzOk() (*[]string, bool)`

GetAuthzOk returns a tuple with the Authz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthz

`func (o *UserProvidersInfo) SetAuthz(v []string)`

SetAuthz sets Authz field to given value.

### HasAuthz

`func (o *UserProvidersInfo) HasAuthz() bool`

HasAuthz returns a boolean if a field has been set.

### GetRoles

`func (o *UserProvidersInfo) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserProvidersInfo) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserProvidersInfo) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserProvidersInfo) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetCustomRights

`func (o *UserProvidersInfo) GetCustomRights() []string`

GetCustomRights returns the CustomRights field if non-nil, zero value otherwise.

### GetCustomRightsOk

`func (o *UserProvidersInfo) GetCustomRightsOk() (*[]string, bool)`

GetCustomRightsOk returns a tuple with the CustomRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRights

`func (o *UserProvidersInfo) SetCustomRights(v []string)`

SetCustomRights sets CustomRights field to given value.

### HasCustomRights

`func (o *UserProvidersInfo) HasCustomRights() bool`

HasCustomRights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


