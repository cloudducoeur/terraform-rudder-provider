# RegenerateToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Id** | **string** | Id of the API account | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllAccounts200ResponseData**](GetAllAccounts200ResponseData.md) |  | 

## Methods

### NewRegenerateToken200Response

`func NewRegenerateToken200Response(result string, id string, action string, data GetAllAccounts200ResponseData, ) *RegenerateToken200Response`

NewRegenerateToken200Response instantiates a new RegenerateToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegenerateToken200ResponseWithDefaults

`func NewRegenerateToken200ResponseWithDefaults() *RegenerateToken200Response`

NewRegenerateToken200ResponseWithDefaults instantiates a new RegenerateToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *RegenerateToken200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RegenerateToken200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RegenerateToken200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetId

`func (o *RegenerateToken200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegenerateToken200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegenerateToken200Response) SetId(v string)`

SetId sets Id field to given value.


### GetAction

`func (o *RegenerateToken200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RegenerateToken200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RegenerateToken200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *RegenerateToken200Response) GetData() GetAllAccounts200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RegenerateToken200Response) GetDataOk() (*GetAllAccounts200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RegenerateToken200Response) SetData(v GetAllAccounts200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


