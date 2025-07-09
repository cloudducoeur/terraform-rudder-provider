# UpdateAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllAccounts200ResponseData**](GetAllAccounts200ResponseData.md) |  | 

## Methods

### NewUpdateAccount200Response

`func NewUpdateAccount200Response(result string, action string, data GetAllAccounts200ResponseData, ) *UpdateAccount200Response`

NewUpdateAccount200Response instantiates a new UpdateAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccount200ResponseWithDefaults

`func NewUpdateAccount200ResponseWithDefaults() *UpdateAccount200Response`

NewUpdateAccount200ResponseWithDefaults instantiates a new UpdateAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdateAccount200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateAccount200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateAccount200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *UpdateAccount200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateAccount200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateAccount200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *UpdateAccount200Response) GetData() GetAllAccounts200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateAccount200Response) GetDataOk() (*GetAllAccounts200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateAccount200Response) SetData(v GetAllAccounts200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


