# RoleCoverageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | **[]string** | List of roles | 
**Authz** | **[]string** | List of authorizations | 

## Methods

### NewRoleCoverageRequest

`func NewRoleCoverageRequest(permissions []string, authz []string, ) *RoleCoverageRequest`

NewRoleCoverageRequest instantiates a new RoleCoverageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleCoverageRequestWithDefaults

`func NewRoleCoverageRequestWithDefaults() *RoleCoverageRequest`

NewRoleCoverageRequestWithDefaults instantiates a new RoleCoverageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *RoleCoverageRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleCoverageRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleCoverageRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetAuthz

`func (o *RoleCoverageRequest) GetAuthz() []string`

GetAuthz returns the Authz field if non-nil, zero value otherwise.

### GetAuthzOk

`func (o *RoleCoverageRequest) GetAuthzOk() (*[]string, bool)`

GetAuthzOk returns a tuple with the Authz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthz

`func (o *RoleCoverageRequest) SetAuthz(v []string)`

SetAuthz sets Authz field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


