# ApiAccountUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human readable name of the API account | [optional] 
**Description** | Pointer to **string** | One line API account description | [optional] 
**Status** | Pointer to **string** | Status of that API account | [optional] 
**ExpirationPolicy** | Pointer to **string** | Optionally specify the expiration policy of the API account. Default: never | [optional] 
**ExpirationDate** | Pointer to **string** | Optional expiration date of the API account | [optional] 
**Tenants** | Pointer to **string** | a comma separated list of tenants allowed for that account. See tenant documentation for syntax. | [optional] 
**AuthorizationType** | Pointer to **string** | Authorization type for that account:  - ro: read-only (GET) access to all APIs - rw: read (GET) and write (other verbs) access to all APIs - acl: access control list based permission for API (need api-authorization plugin) | [optional] 
**Acl** | Pointer to [**[]ApiAccountNewAclInner**](ApiAccountNewAclInner.md) | in case of ACL authorization type, the actual ACL | [optional] 

## Methods

### NewApiAccountUpdate

`func NewApiAccountUpdate() *ApiAccountUpdate`

NewApiAccountUpdate instantiates a new ApiAccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccountUpdateWithDefaults

`func NewApiAccountUpdateWithDefaults() *ApiAccountUpdate`

NewApiAccountUpdateWithDefaults instantiates a new ApiAccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiAccountUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAccountUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAccountUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAccountUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiAccountUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiAccountUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiAccountUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiAccountUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAccountUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAccountUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAccountUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAccountUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *ApiAccountUpdate) GetExpirationPolicy() string`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *ApiAccountUpdate) GetExpirationPolicyOk() (*string, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *ApiAccountUpdate) SetExpirationPolicy(v string)`

SetExpirationPolicy sets ExpirationPolicy field to given value.

### HasExpirationPolicy

`func (o *ApiAccountUpdate) HasExpirationPolicy() bool`

HasExpirationPolicy returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ApiAccountUpdate) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiAccountUpdate) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiAccountUpdate) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiAccountUpdate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTenants

`func (o *ApiAccountUpdate) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApiAccountUpdate) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApiAccountUpdate) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApiAccountUpdate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetAuthorizationType

`func (o *ApiAccountUpdate) GetAuthorizationType() string`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *ApiAccountUpdate) GetAuthorizationTypeOk() (*string, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *ApiAccountUpdate) SetAuthorizationType(v string)`

SetAuthorizationType sets AuthorizationType field to given value.

### HasAuthorizationType

`func (o *ApiAccountUpdate) HasAuthorizationType() bool`

HasAuthorizationType returns a boolean if a field has been set.

### GetAcl

`func (o *ApiAccountUpdate) GetAcl() []ApiAccountNewAclInner`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *ApiAccountUpdate) GetAclOk() (*[]ApiAccountNewAclInner, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *ApiAccountUpdate) SetAcl(v []ApiAccountNewAclInner)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *ApiAccountUpdate) HasAcl() bool`

HasAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


