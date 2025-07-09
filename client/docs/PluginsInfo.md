# PluginsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalLimits** | Pointer to [**PluginsInfoGlobalLimits**](PluginsInfoGlobalLimits.md) |  | [optional] 
**Details** | [**[]PluginsInfoDetailsInner**](PluginsInfoDetailsInner.md) | the list of details for each plugins | 

## Methods

### NewPluginsInfo

`func NewPluginsInfo(details []PluginsInfoDetailsInner, ) *PluginsInfo`

NewPluginsInfo instantiates a new PluginsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsInfoWithDefaults

`func NewPluginsInfoWithDefaults() *PluginsInfo`

NewPluginsInfoWithDefaults instantiates a new PluginsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalLimits

`func (o *PluginsInfo) GetGlobalLimits() PluginsInfoGlobalLimits`

GetGlobalLimits returns the GlobalLimits field if non-nil, zero value otherwise.

### GetGlobalLimitsOk

`func (o *PluginsInfo) GetGlobalLimitsOk() (*PluginsInfoGlobalLimits, bool)`

GetGlobalLimitsOk returns a tuple with the GlobalLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalLimits

`func (o *PluginsInfo) SetGlobalLimits(v PluginsInfoGlobalLimits)`

SetGlobalLimits sets GlobalLimits field to given value.

### HasGlobalLimits

`func (o *PluginsInfo) HasGlobalLimits() bool`

HasGlobalLimits returns a boolean if a field has been set.

### GetDetails

`func (o *PluginsInfo) GetDetails() []PluginsInfoDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PluginsInfo) GetDetailsOk() (*[]PluginsInfoDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PluginsInfo) SetDetails(v []PluginsInfoDetailsInner)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


