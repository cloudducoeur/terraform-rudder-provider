# ApiAccountNewAclInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | API path managed by that access control | [optional] 
**Actions** | Pointer to **[]string** | List of HTTP verb allowed | [optional] 

## Methods

### NewApiAccountNewAclInner

`func NewApiAccountNewAclInner() *ApiAccountNewAclInner`

NewApiAccountNewAclInner instantiates a new ApiAccountNewAclInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccountNewAclInnerWithDefaults

`func NewApiAccountNewAclInnerWithDefaults() *ApiAccountNewAclInner`

NewApiAccountNewAclInnerWithDefaults instantiates a new ApiAccountNewAclInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ApiAccountNewAclInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiAccountNewAclInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiAccountNewAclInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApiAccountNewAclInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetActions

`func (o *ApiAccountNewAclInner) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiAccountNewAclInner) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiAccountNewAclInner) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApiAccountNewAclInner) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


