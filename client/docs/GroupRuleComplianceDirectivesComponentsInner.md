# GroupRuleComplianceDirectivesComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the component | 
**Compliance** | **float32** | Component compliance level | 
**ComplianceDetails** | [**RuleComplianceComponentComplianceDetails**](RuleComplianceComponentComplianceDetails.md) |  | 
**Nodes** | [**[]NodeComplianceComponent**](NodeComplianceComponent.md) |  | 

## Methods

### NewGroupRuleComplianceDirectivesComponentsInner

`func NewGroupRuleComplianceDirectivesComponentsInner(name string, compliance float32, complianceDetails RuleComplianceComponentComplianceDetails, nodes []NodeComplianceComponent, ) *GroupRuleComplianceDirectivesComponentsInner`

NewGroupRuleComplianceDirectivesComponentsInner instantiates a new GroupRuleComplianceDirectivesComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRuleComplianceDirectivesComponentsInnerWithDefaults

`func NewGroupRuleComplianceDirectivesComponentsInnerWithDefaults() *GroupRuleComplianceDirectivesComponentsInner`

NewGroupRuleComplianceDirectivesComponentsInnerWithDefaults instantiates a new GroupRuleComplianceDirectivesComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupRuleComplianceDirectivesComponentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRuleComplianceDirectivesComponentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRuleComplianceDirectivesComponentsInner) SetName(v string)`

SetName sets Name field to given value.


### GetCompliance

`func (o *GroupRuleComplianceDirectivesComponentsInner) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GroupRuleComplianceDirectivesComponentsInner) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GroupRuleComplianceDirectivesComponentsInner) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GroupRuleComplianceDirectivesComponentsInner) GetComplianceDetails() RuleComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GroupRuleComplianceDirectivesComponentsInner) GetComplianceDetailsOk() (*RuleComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GroupRuleComplianceDirectivesComponentsInner) SetComplianceDetails(v RuleComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetNodes

`func (o *GroupRuleComplianceDirectivesComponentsInner) GetNodes() []NodeComplianceComponent`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GroupRuleComplianceDirectivesComponentsInner) GetNodesOk() (*[]NodeComplianceComponent, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GroupRuleComplianceDirectivesComponentsInner) SetNodes(v []NodeComplianceComponent)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


