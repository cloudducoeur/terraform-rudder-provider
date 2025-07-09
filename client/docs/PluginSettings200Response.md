# PluginSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**PluginSettings200ResponseData**](PluginSettings200ResponseData.md) |  | 

## Methods

### NewPluginSettings200Response

`func NewPluginSettings200Response(result string, action string, data PluginSettings200ResponseData, ) *PluginSettings200Response`

NewPluginSettings200Response instantiates a new PluginSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginSettings200ResponseWithDefaults

`func NewPluginSettings200ResponseWithDefaults() *PluginSettings200Response`

NewPluginSettings200ResponseWithDefaults instantiates a new PluginSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *PluginSettings200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PluginSettings200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PluginSettings200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *PluginSettings200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PluginSettings200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PluginSettings200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *PluginSettings200Response) GetData() PluginSettings200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PluginSettings200Response) GetDataOk() (*PluginSettings200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PluginSettings200Response) SetData(v PluginSettings200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


