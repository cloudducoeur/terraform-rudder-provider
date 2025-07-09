# ApiAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | API account id | [optional] 
**Name** | Pointer to **string** | Human readable name of the API account | [optional] 
**Description** | Pointer to **string** | One line API account description | [optional] 
**Status** | Pointer to **string** | Status of that API account | [optional] 
**CreationDate** | Pointer to **string** | Creation date of the API account | [optional] 
**ExpirationPolicy** | Pointer to **string** | Expiration policy of the API account (not the token) | [optional] 
**ExpirationDate** | Pointer to **string** | Optional expiration date of the API account | [optional] 
**TokenState** | Pointer to **string** | State of the token for that account. An account with token state &#39;missing&#39; means that there is no token for that account. | [optional] 
**Token** | Pointer to **string** | The token secret. It will be returned only once when a new token is generated. | [optional] 
**TokenGenerationDate** | Pointer to **string** | Optional generation date for the token when tokenState is &#39;generated&#39; | [optional] 
**Tenants** | Pointer to **string** | a comma separated list of tenants allowed for that account. See tenant documentation for syntax. | [optional] 
**AuthorizationType** | Pointer to **string** | Authorization type for that account:  - ro: read-only (GET) access to all APIs - rw: read (GET) and write (other verbs) access to all APIs - acl: access control list based permission for API (need api-authorization plugin) | [optional] 
**Acl** | Pointer to [**[]ApiAccountAclInner**](ApiAccountAclInner.md) | in case of ACL authorization type, the actual ACL | [optional] 

## Methods

### NewApiAccount

`func NewApiAccount() *ApiAccount`

NewApiAccount instantiates a new ApiAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccountWithDefaults

`func NewApiAccountWithDefaults() *ApiAccount`

NewApiAccountWithDefaults instantiates a new ApiAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApiAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreationDate

`func (o *ApiAccount) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ApiAccount) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ApiAccount) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ApiAccount) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *ApiAccount) GetExpirationPolicy() string`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *ApiAccount) GetExpirationPolicyOk() (*string, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *ApiAccount) SetExpirationPolicy(v string)`

SetExpirationPolicy sets ExpirationPolicy field to given value.

### HasExpirationPolicy

`func (o *ApiAccount) HasExpirationPolicy() bool`

HasExpirationPolicy returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ApiAccount) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiAccount) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiAccount) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiAccount) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTokenState

`func (o *ApiAccount) GetTokenState() string`

GetTokenState returns the TokenState field if non-nil, zero value otherwise.

### GetTokenStateOk

`func (o *ApiAccount) GetTokenStateOk() (*string, bool)`

GetTokenStateOk returns a tuple with the TokenState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenState

`func (o *ApiAccount) SetTokenState(v string)`

SetTokenState sets TokenState field to given value.

### HasTokenState

`func (o *ApiAccount) HasTokenState() bool`

HasTokenState returns a boolean if a field has been set.

### GetToken

`func (o *ApiAccount) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiAccount) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiAccount) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiAccount) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenGenerationDate

`func (o *ApiAccount) GetTokenGenerationDate() string`

GetTokenGenerationDate returns the TokenGenerationDate field if non-nil, zero value otherwise.

### GetTokenGenerationDateOk

`func (o *ApiAccount) GetTokenGenerationDateOk() (*string, bool)`

GetTokenGenerationDateOk returns a tuple with the TokenGenerationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenGenerationDate

`func (o *ApiAccount) SetTokenGenerationDate(v string)`

SetTokenGenerationDate sets TokenGenerationDate field to given value.

### HasTokenGenerationDate

`func (o *ApiAccount) HasTokenGenerationDate() bool`

HasTokenGenerationDate returns a boolean if a field has been set.

### GetTenants

`func (o *ApiAccount) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApiAccount) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApiAccount) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApiAccount) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetAuthorizationType

`func (o *ApiAccount) GetAuthorizationType() string`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *ApiAccount) GetAuthorizationTypeOk() (*string, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *ApiAccount) SetAuthorizationType(v string)`

SetAuthorizationType sets AuthorizationType field to given value.

### HasAuthorizationType

`func (o *ApiAccount) HasAuthorizationType() bool`

HasAuthorizationType returns a boolean if a field has been set.

### GetAcl

`func (o *ApiAccount) GetAcl() []ApiAccountAclInner`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *ApiAccount) GetAclOk() (*[]ApiAccountAclInner, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *ApiAccount) SetAcl(v []ApiAccountAclInner)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *ApiAccount) HasAcl() bool`

HasAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


