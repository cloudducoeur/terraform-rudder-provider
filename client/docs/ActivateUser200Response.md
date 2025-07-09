# ActivateUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The id of the action | 
**Id** | **string** | The id of the updated user | 
**Result** | **string** | Result of the request | 
**Data** | [**ActivateUser200ResponseData**](ActivateUser200ResponseData.md) |  | 

## Methods

### NewActivateUser200Response

`func NewActivateUser200Response(action string, id string, result string, data ActivateUser200ResponseData, ) *ActivateUser200Response`

NewActivateUser200Response instantiates a new ActivateUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateUser200ResponseWithDefaults

`func NewActivateUser200ResponseWithDefaults() *ActivateUser200Response`

NewActivateUser200ResponseWithDefaults instantiates a new ActivateUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ActivateUser200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActivateUser200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActivateUser200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *ActivateUser200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivateUser200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivateUser200Response) SetId(v string)`

SetId sets Id field to given value.


### GetResult

`func (o *ActivateUser200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ActivateUser200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ActivateUser200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetData

`func (o *ActivateUser200Response) GetData() ActivateUser200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ActivateUser200Response) GetDataOk() (*ActivateUser200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ActivateUser200Response) SetData(v ActivateUser200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


