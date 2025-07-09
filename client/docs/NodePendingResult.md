# NodePendingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique id of the node | 
**Hostname** | **string** | Fully qualified name of the node | 
**Status** | **string** | Status of the node | 
**OsName** | Pointer to **string** | Operating system name (distribution on Linux, etc.) | [optional] 
**OsVersion** | Pointer to **string** | OS version | [optional] 
**MachineType** | Pointer to **string** | The kind of machine for the node (use vm for a generic VM) | [optional] 

## Methods

### NewNodePendingResult

`func NewNodePendingResult(id string, hostname string, status string, ) *NodePendingResult`

NewNodePendingResult instantiates a new NodePendingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePendingResultWithDefaults

`func NewNodePendingResultWithDefaults() *NodePendingResult`

NewNodePendingResultWithDefaults instantiates a new NodePendingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodePendingResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodePendingResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodePendingResult) SetId(v string)`

SetId sets Id field to given value.


### GetHostname

`func (o *NodePendingResult) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodePendingResult) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodePendingResult) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStatus

`func (o *NodePendingResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodePendingResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodePendingResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOsName

`func (o *NodePendingResult) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *NodePendingResult) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *NodePendingResult) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *NodePendingResult) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *NodePendingResult) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *NodePendingResult) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *NodePendingResult) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *NodePendingResult) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetMachineType

`func (o *NodePendingResult) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *NodePendingResult) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *NodePendingResult) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *NodePendingResult) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


