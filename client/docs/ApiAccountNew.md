# ApiAccountNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An optional ID for that account. It needs to match [a-zA-Z0-9_-]+ | [optional] 
**Name** | Pointer to **string** | Human readable name of the API account | [optional] 
**Description** | Pointer to **string** | One line API account description | [optional] 
**Status** | Pointer to **string** | Status of that API account | [optional] 
**GenerateToken** | Pointer to **bool** | Optionally specify if the token need to be generated (default true) | [optional] 
**ExpirationPolicy** | Pointer to **string** | Optionally specify the expiration policy of the API account. Default: never | [optional] 
**ExpirationDate** | Pointer to **string** | Optional expiration date of the API account | [optional] 
**Tenants** | Pointer to **string** | a comma separated list of tenants allowed for that account. See tenant documentation for syntax. | [optional] 
**AuthorizationType** | Pointer to **string** | Authorization type for that account:  - ro: read-only (GET) access to all APIs - rw: read (GET) and write (other verbs) access to all APIs - acl: access control list based permission for API (need api-authorization plugin) | [optional] 
**Acl** | Pointer to [**[]ApiAccountNewAclInner**](ApiAccountNewAclInner.md) | in case of ACL authorization type, the actual ACL | [optional] 

## Methods

### NewApiAccountNew

`func NewApiAccountNew() *ApiAccountNew`

NewApiAccountNew instantiates a new ApiAccountNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccountNewWithDefaults

`func NewApiAccountNewWithDefaults() *ApiAccountNew`

NewApiAccountNewWithDefaults instantiates a new ApiAccountNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAccountNew) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAccountNew) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAccountNew) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAccountNew) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApiAccountNew) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAccountNew) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAccountNew) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAccountNew) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiAccountNew) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiAccountNew) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiAccountNew) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiAccountNew) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAccountNew) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAccountNew) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAccountNew) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAccountNew) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGenerateToken

`func (o *ApiAccountNew) GetGenerateToken() bool`

GetGenerateToken returns the GenerateToken field if non-nil, zero value otherwise.

### GetGenerateTokenOk

`func (o *ApiAccountNew) GetGenerateTokenOk() (*bool, bool)`

GetGenerateTokenOk returns a tuple with the GenerateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateToken

`func (o *ApiAccountNew) SetGenerateToken(v bool)`

SetGenerateToken sets GenerateToken field to given value.

### HasGenerateToken

`func (o *ApiAccountNew) HasGenerateToken() bool`

HasGenerateToken returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *ApiAccountNew) GetExpirationPolicy() string`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *ApiAccountNew) GetExpirationPolicyOk() (*string, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *ApiAccountNew) SetExpirationPolicy(v string)`

SetExpirationPolicy sets ExpirationPolicy field to given value.

### HasExpirationPolicy

`func (o *ApiAccountNew) HasExpirationPolicy() bool`

HasExpirationPolicy returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ApiAccountNew) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiAccountNew) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiAccountNew) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiAccountNew) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTenants

`func (o *ApiAccountNew) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApiAccountNew) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApiAccountNew) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApiAccountNew) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetAuthorizationType

`func (o *ApiAccountNew) GetAuthorizationType() string`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *ApiAccountNew) GetAuthorizationTypeOk() (*string, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *ApiAccountNew) SetAuthorizationType(v string)`

SetAuthorizationType sets AuthorizationType field to given value.

### HasAuthorizationType

`func (o *ApiAccountNew) HasAuthorizationType() bool`

HasAuthorizationType returns a boolean if a field has been set.

### GetAcl

`func (o *ApiAccountNew) GetAcl() []ApiAccountNewAclInner`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *ApiAccountNew) GetAclOk() (*[]ApiAccountNewAclInner, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *ApiAccountNew) SetAcl(v []ApiAccountNewAclInner)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *ApiAccountNew) HasAcl() bool`

HasAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


