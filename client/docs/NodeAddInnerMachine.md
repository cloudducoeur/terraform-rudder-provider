# NodeAddInnerMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | list of groups to include in rule application | 
**Provider** | Pointer to **string** | The kind of virtual machine for the node | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the machine | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the machine | [optional] 

## Methods

### NewNodeAddInnerMachine

`func NewNodeAddInnerMachine(type_ string, ) *NodeAddInnerMachine`

NewNodeAddInnerMachine instantiates a new NodeAddInnerMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAddInnerMachineWithDefaults

`func NewNodeAddInnerMachineWithDefaults() *NodeAddInnerMachine`

NewNodeAddInnerMachineWithDefaults instantiates a new NodeAddInnerMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NodeAddInnerMachine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeAddInnerMachine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeAddInnerMachine) SetType(v string)`

SetType sets Type field to given value.


### GetProvider

`func (o *NodeAddInnerMachine) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NodeAddInnerMachine) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NodeAddInnerMachine) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NodeAddInnerMachine) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetManufacturer

`func (o *NodeAddInnerMachine) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *NodeAddInnerMachine) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *NodeAddInnerMachine) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *NodeAddInnerMachine) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NodeAddInnerMachine) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NodeAddInnerMachine) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NodeAddInnerMachine) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NodeAddInnerMachine) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


