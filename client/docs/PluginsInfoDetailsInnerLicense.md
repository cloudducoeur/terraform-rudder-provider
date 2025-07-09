# PluginsInfoDetailsInnerLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licensee** | Pointer to **string** | name of the licensee for that plugin | [optional] 
**SoftwareId** | Pointer to **string** | the fully qualified name of the plugin for which that license was issued | [optional] 
**MinVersion** | Pointer to **string** | lowest version of the software for which that license is valid | [optional] 
**MaxVersion** | Pointer to **string** | highest version of the software for which that license is valid | [optional] 
**StartDate** | Pointer to **string** | start of validity date | [optional] 
**EndDate** | Pointer to **string** | end of validity date | [optional] 
**MaxNodes** | Pointer to **int32** | maximum number of node in Rudder for that license | [optional] 
**AdditionalInfo** | Pointer to **map[string]interface{}** | additional information provided by the license | [optional] 

## Methods

### NewPluginsInfoDetailsInnerLicense

`func NewPluginsInfoDetailsInnerLicense() *PluginsInfoDetailsInnerLicense`

NewPluginsInfoDetailsInnerLicense instantiates a new PluginsInfoDetailsInnerLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsInfoDetailsInnerLicenseWithDefaults

`func NewPluginsInfoDetailsInnerLicenseWithDefaults() *PluginsInfoDetailsInnerLicense`

NewPluginsInfoDetailsInnerLicenseWithDefaults instantiates a new PluginsInfoDetailsInnerLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicensee

`func (o *PluginsInfoDetailsInnerLicense) GetLicensee() string`

GetLicensee returns the Licensee field if non-nil, zero value otherwise.

### GetLicenseeOk

`func (o *PluginsInfoDetailsInnerLicense) GetLicenseeOk() (*string, bool)`

GetLicenseeOk returns a tuple with the Licensee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensee

`func (o *PluginsInfoDetailsInnerLicense) SetLicensee(v string)`

SetLicensee sets Licensee field to given value.

### HasLicensee

`func (o *PluginsInfoDetailsInnerLicense) HasLicensee() bool`

HasLicensee returns a boolean if a field has been set.

### GetSoftwareId

`func (o *PluginsInfoDetailsInnerLicense) GetSoftwareId() string`

GetSoftwareId returns the SoftwareId field if non-nil, zero value otherwise.

### GetSoftwareIdOk

`func (o *PluginsInfoDetailsInnerLicense) GetSoftwareIdOk() (*string, bool)`

GetSoftwareIdOk returns a tuple with the SoftwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareId

`func (o *PluginsInfoDetailsInnerLicense) SetSoftwareId(v string)`

SetSoftwareId sets SoftwareId field to given value.

### HasSoftwareId

`func (o *PluginsInfoDetailsInnerLicense) HasSoftwareId() bool`

HasSoftwareId returns a boolean if a field has been set.

### GetMinVersion

`func (o *PluginsInfoDetailsInnerLicense) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *PluginsInfoDetailsInnerLicense) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *PluginsInfoDetailsInnerLicense) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *PluginsInfoDetailsInnerLicense) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetMaxVersion

`func (o *PluginsInfoDetailsInnerLicense) GetMaxVersion() string`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *PluginsInfoDetailsInnerLicense) GetMaxVersionOk() (*string, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *PluginsInfoDetailsInnerLicense) SetMaxVersion(v string)`

SetMaxVersion sets MaxVersion field to given value.

### HasMaxVersion

`func (o *PluginsInfoDetailsInnerLicense) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.

### GetStartDate

`func (o *PluginsInfoDetailsInnerLicense) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PluginsInfoDetailsInnerLicense) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PluginsInfoDetailsInnerLicense) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PluginsInfoDetailsInnerLicense) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PluginsInfoDetailsInnerLicense) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PluginsInfoDetailsInnerLicense) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PluginsInfoDetailsInnerLicense) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PluginsInfoDetailsInnerLicense) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMaxNodes

`func (o *PluginsInfoDetailsInnerLicense) GetMaxNodes() int32`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *PluginsInfoDetailsInnerLicense) GetMaxNodesOk() (*int32, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *PluginsInfoDetailsInnerLicense) SetMaxNodes(v int32)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *PluginsInfoDetailsInnerLicense) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *PluginsInfoDetailsInnerLicense) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *PluginsInfoDetailsInnerLicense) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *PluginsInfoDetailsInnerLicense) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *PluginsInfoDetailsInnerLicense) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


