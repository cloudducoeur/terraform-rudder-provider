# PluginsInfoGlobalLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licensees** | Pointer to **[]string** | list of licensees for these plugins | [optional] 
**StartDate** | Pointer to **string** | the latest date of start of validity for plugins | [optional] 
**EndDate** | Pointer to **string** | the earliest date of end of validity for plugins | [optional] 
**MaxNodes** | Pointer to **int32** | the lowest limit on maximum number of nodes for plugins | [optional] 

## Methods

### NewPluginsInfoGlobalLimits

`func NewPluginsInfoGlobalLimits() *PluginsInfoGlobalLimits`

NewPluginsInfoGlobalLimits instantiates a new PluginsInfoGlobalLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsInfoGlobalLimitsWithDefaults

`func NewPluginsInfoGlobalLimitsWithDefaults() *PluginsInfoGlobalLimits`

NewPluginsInfoGlobalLimitsWithDefaults instantiates a new PluginsInfoGlobalLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicensees

`func (o *PluginsInfoGlobalLimits) GetLicensees() []string`

GetLicensees returns the Licensees field if non-nil, zero value otherwise.

### GetLicenseesOk

`func (o *PluginsInfoGlobalLimits) GetLicenseesOk() (*[]string, bool)`

GetLicenseesOk returns a tuple with the Licensees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensees

`func (o *PluginsInfoGlobalLimits) SetLicensees(v []string)`

SetLicensees sets Licensees field to given value.

### HasLicensees

`func (o *PluginsInfoGlobalLimits) HasLicensees() bool`

HasLicensees returns a boolean if a field has been set.

### GetStartDate

`func (o *PluginsInfoGlobalLimits) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PluginsInfoGlobalLimits) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PluginsInfoGlobalLimits) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PluginsInfoGlobalLimits) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PluginsInfoGlobalLimits) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PluginsInfoGlobalLimits) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PluginsInfoGlobalLimits) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PluginsInfoGlobalLimits) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMaxNodes

`func (o *PluginsInfoGlobalLimits) GetMaxNodes() int32`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *PluginsInfoGlobalLimits) GetMaxNodesOk() (*int32, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *PluginsInfoGlobalLimits) SetMaxNodes(v int32)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *PluginsInfoGlobalLimits) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


