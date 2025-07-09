# GroupNodeComplianceRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the rule | 
**Name** | **string** | Name of the rule | 
**PolicyMode** | **string** | Policy mode of the rule | 
**Compliance** | **float32** | Rule compliance level | 
**ComplianceDetails** | [**RuleComplianceComponentComplianceDetails**](RuleComplianceComponentComplianceDetails.md) |  | 
**Directives** | [**[]GroupNodeComplianceRulesDirectivesInner**](GroupNodeComplianceRulesDirectivesInner.md) |  | 

## Methods

### NewGroupNodeComplianceRules

`func NewGroupNodeComplianceRules(id string, name string, policyMode string, compliance float32, complianceDetails RuleComplianceComponentComplianceDetails, directives []GroupNodeComplianceRulesDirectivesInner, ) *GroupNodeComplianceRules`

NewGroupNodeComplianceRules instantiates a new GroupNodeComplianceRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupNodeComplianceRulesWithDefaults

`func NewGroupNodeComplianceRulesWithDefaults() *GroupNodeComplianceRules`

NewGroupNodeComplianceRulesWithDefaults instantiates a new GroupNodeComplianceRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupNodeComplianceRules) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupNodeComplianceRules) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupNodeComplianceRules) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupNodeComplianceRules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupNodeComplianceRules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupNodeComplianceRules) SetName(v string)`

SetName sets Name field to given value.


### GetPolicyMode

`func (o *GroupNodeComplianceRules) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *GroupNodeComplianceRules) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *GroupNodeComplianceRules) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.


### GetCompliance

`func (o *GroupNodeComplianceRules) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GroupNodeComplianceRules) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GroupNodeComplianceRules) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GroupNodeComplianceRules) GetComplianceDetails() RuleComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GroupNodeComplianceRules) GetComplianceDetailsOk() (*RuleComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GroupNodeComplianceRules) SetComplianceDetails(v RuleComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetDirectives

`func (o *GroupNodeComplianceRules) GetDirectives() []GroupNodeComplianceRulesDirectivesInner`

GetDirectives returns the Directives field if non-nil, zero value otherwise.

### GetDirectivesOk

`func (o *GroupNodeComplianceRules) GetDirectivesOk() (*[]GroupNodeComplianceRulesDirectivesInner, bool)`

GetDirectivesOk returns a tuple with the Directives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectives

`func (o *GroupNodeComplianceRules) SetDirectives(v []GroupNodeComplianceRulesDirectivesInner)`

SetDirectives sets Directives field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


