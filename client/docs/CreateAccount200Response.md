# CreateAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Id** | **string** | ID of the created account | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllAccounts200ResponseData**](GetAllAccounts200ResponseData.md) |  | 

## Methods

### NewCreateAccount200Response

`func NewCreateAccount200Response(result string, id string, action string, data GetAllAccounts200ResponseData, ) *CreateAccount200Response`

NewCreateAccount200Response instantiates a new CreateAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccount200ResponseWithDefaults

`func NewCreateAccount200ResponseWithDefaults() *CreateAccount200Response`

NewCreateAccount200ResponseWithDefaults instantiates a new CreateAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CreateAccount200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateAccount200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateAccount200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetId

`func (o *CreateAccount200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAccount200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAccount200Response) SetId(v string)`

SetId sets Id field to given value.


### GetAction

`func (o *CreateAccount200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateAccount200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateAccount200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *CreateAccount200Response) GetData() GetAllAccounts200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateAccount200Response) GetDataOk() (*GetAllAccounts200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateAccount200Response) SetData(v GetAllAccounts200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


