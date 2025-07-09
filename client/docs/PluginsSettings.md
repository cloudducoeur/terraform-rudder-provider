# PluginsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | username to use for Rudder account | 
**Password** | **string** | password to access Rudder account | 
**Url** | Pointer to **string** | URL for getting plugins | [optional] 
**ProxyUser** | Pointer to **string** | if an authenticated proxy is necessary, username of proxy | [optional] 
**ProxyPassword** | Pointer to **string** | if an authenticated proxy is necessary, password of the proxy | [optional] 
**ProxyUrl** | Pointer to **string** | proxy URL to use | [optional] 

## Methods

### NewPluginsSettings

`func NewPluginsSettings(username string, password string, ) *PluginsSettings`

NewPluginsSettings instantiates a new PluginsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsSettingsWithDefaults

`func NewPluginsSettingsWithDefaults() *PluginsSettings`

NewPluginsSettingsWithDefaults instantiates a new PluginsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *PluginsSettings) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PluginsSettings) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PluginsSettings) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *PluginsSettings) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PluginsSettings) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PluginsSettings) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUrl

`func (o *PluginsSettings) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PluginsSettings) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PluginsSettings) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PluginsSettings) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetProxyUser

`func (o *PluginsSettings) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *PluginsSettings) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *PluginsSettings) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *PluginsSettings) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### GetProxyPassword

`func (o *PluginsSettings) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *PluginsSettings) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *PluginsSettings) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *PluginsSettings) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetProxyUrl

`func (o *PluginsSettings) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *PluginsSettings) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *PluginsSettings) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *PluginsSettings) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


