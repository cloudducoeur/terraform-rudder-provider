# GetNodeGroupComplianceId200ResponseDataNodeGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the group | 
**Mode** | **string** |  | 
**Compliance** | **float32** | Group compliance level | 
**ComplianceDetails** | [**GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails**](GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails.md) |  | 
**Rules** | [**[]GroupRuleCompliance**](GroupRuleCompliance.md) |  | 
**Nodes** | [**[]GroupNodeCompliance**](GroupNodeCompliance.md) |  | 

## Methods

### NewGetNodeGroupComplianceId200ResponseDataNodeGroupsInner

`func NewGetNodeGroupComplianceId200ResponseDataNodeGroupsInner(id string, mode string, compliance float32, complianceDetails GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails, rules []GroupRuleCompliance, nodes []GroupNodeCompliance, ) *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner`

NewGetNodeGroupComplianceId200ResponseDataNodeGroupsInner instantiates a new GetNodeGroupComplianceId200ResponseDataNodeGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodeGroupComplianceId200ResponseDataNodeGroupsInnerWithDefaults

`func NewGetNodeGroupComplianceId200ResponseDataNodeGroupsInnerWithDefaults() *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner`

NewGetNodeGroupComplianceId200ResponseDataNodeGroupsInnerWithDefaults instantiates a new GetNodeGroupComplianceId200ResponseDataNodeGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) SetId(v string)`

SetId sets Id field to given value.


### GetMode

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) SetMode(v string)`

SetMode sets Mode field to given value.


### GetCompliance

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetComplianceDetails() GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetComplianceDetailsOk() (*GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) SetComplianceDetails(v GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetRules

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetRules() []GroupRuleCompliance`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetRulesOk() (*[]GroupRuleCompliance, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) SetRules(v []GroupRuleCompliance)`

SetRules sets Rules field to given value.


### GetNodes

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetNodes() []GroupNodeCompliance`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) GetNodesOk() (*[]GroupNodeCompliance, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetNodeGroupComplianceId200ResponseDataNodeGroupsInner) SetNodes(v []GroupNodeCompliance)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


