# CampaignScheduleDaily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to [**CampaignScheduleDailyStart**](CampaignScheduleDailyStart.md) |  | [optional] 
**End** | Pointer to [**CampaignScheduleDailyEnd**](CampaignScheduleDailyEnd.md) |  | [optional] 
**Tz** | Pointer to **string** | IANA timezone ID associated with the start and end times of the campaign schedule | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCampaignScheduleDaily

`func NewCampaignScheduleDaily() *CampaignScheduleDaily`

NewCampaignScheduleDaily instantiates a new CampaignScheduleDaily object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignScheduleDailyWithDefaults

`func NewCampaignScheduleDailyWithDefaults() *CampaignScheduleDaily`

NewCampaignScheduleDailyWithDefaults instantiates a new CampaignScheduleDaily object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *CampaignScheduleDaily) GetStart() CampaignScheduleDailyStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CampaignScheduleDaily) GetStartOk() (*CampaignScheduleDailyStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CampaignScheduleDaily) SetStart(v CampaignScheduleDailyStart)`

SetStart sets Start field to given value.

### HasStart

`func (o *CampaignScheduleDaily) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *CampaignScheduleDaily) GetEnd() CampaignScheduleDailyEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CampaignScheduleDaily) GetEndOk() (*CampaignScheduleDailyEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CampaignScheduleDaily) SetEnd(v CampaignScheduleDailyEnd)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CampaignScheduleDaily) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetTz

`func (o *CampaignScheduleDaily) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *CampaignScheduleDaily) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *CampaignScheduleDaily) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *CampaignScheduleDaily) HasTz() bool`

HasTz returns a boolean if a field has been set.

### GetType

`func (o *CampaignScheduleDaily) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignScheduleDaily) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignScheduleDaily) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CampaignScheduleDaily) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


