# GetPluginsInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetPluginsInfo200ResponseData**](GetPluginsInfo200ResponseData.md) |  | 

## Methods

### NewGetPluginsInfo200Response

`func NewGetPluginsInfo200Response(result string, action string, data GetPluginsInfo200ResponseData, ) *GetPluginsInfo200Response`

NewGetPluginsInfo200Response instantiates a new GetPluginsInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPluginsInfo200ResponseWithDefaults

`func NewGetPluginsInfo200ResponseWithDefaults() *GetPluginsInfo200Response`

NewGetPluginsInfo200ResponseWithDefaults instantiates a new GetPluginsInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetPluginsInfo200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetPluginsInfo200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetPluginsInfo200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *GetPluginsInfo200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetPluginsInfo200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetPluginsInfo200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *GetPluginsInfo200Response) GetData() GetPluginsInfo200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPluginsInfo200Response) GetDataOk() (*GetPluginsInfo200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPluginsInfo200Response) SetData(v GetPluginsInfo200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


