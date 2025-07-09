# SetNodeSecurityTagsRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | the ID of the tenant to set on nodes | 
**NodeIds** | **[]string** |  | 

## Methods

### NewSetNodeSecurityTagsRequestInner

`func NewSetNodeSecurityTagsRequestInner(tenantId string, nodeIds []string, ) *SetNodeSecurityTagsRequestInner`

NewSetNodeSecurityTagsRequestInner instantiates a new SetNodeSecurityTagsRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNodeSecurityTagsRequestInnerWithDefaults

`func NewSetNodeSecurityTagsRequestInnerWithDefaults() *SetNodeSecurityTagsRequestInner`

NewSetNodeSecurityTagsRequestInnerWithDefaults instantiates a new SetNodeSecurityTagsRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *SetNodeSecurityTagsRequestInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SetNodeSecurityTagsRequestInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SetNodeSecurityTagsRequestInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetNodeIds

`func (o *SetNodeSecurityTagsRequestInner) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *SetNodeSecurityTagsRequestInner) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *SetNodeSecurityTagsRequestInner) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


