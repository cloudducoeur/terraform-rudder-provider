# ChangePendingNodesStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **[]string** | List of nodes to change status | [optional] 
**Status** | Pointer to **string** | New status of the pending node | [optional] 

## Methods

### NewChangePendingNodesStatusRequest

`func NewChangePendingNodesStatusRequest() *ChangePendingNodesStatusRequest`

NewChangePendingNodesStatusRequest instantiates a new ChangePendingNodesStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePendingNodesStatusRequestWithDefaults

`func NewChangePendingNodesStatusRequestWithDefaults() *ChangePendingNodesStatusRequest`

NewChangePendingNodesStatusRequestWithDefaults instantiates a new ChangePendingNodesStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *ChangePendingNodesStatusRequest) GetNodeId() []string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ChangePendingNodesStatusRequest) GetNodeIdOk() (*[]string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ChangePendingNodesStatusRequest) SetNodeId(v []string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ChangePendingNodesStatusRequest) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetStatus

`func (o *ChangePendingNodesStatusRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangePendingNodesStatusRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangePendingNodesStatusRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChangePendingNodesStatusRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


