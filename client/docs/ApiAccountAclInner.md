# ApiAccountAclInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | API path managed by that access control | [optional] 
**Verb** | Pointer to **string** | HTTP verb allowed | [optional] 

## Methods

### NewApiAccountAclInner

`func NewApiAccountAclInner() *ApiAccountAclInner`

NewApiAccountAclInner instantiates a new ApiAccountAclInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccountAclInnerWithDefaults

`func NewApiAccountAclInnerWithDefaults() *ApiAccountAclInner`

NewApiAccountAclInnerWithDefaults instantiates a new ApiAccountAclInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ApiAccountAclInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiAccountAclInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiAccountAclInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApiAccountAclInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVerb

`func (o *ApiAccountAclInner) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *ApiAccountAclInner) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *ApiAccountAclInner) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *ApiAccountAclInner) HasVerb() bool`

HasVerb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


