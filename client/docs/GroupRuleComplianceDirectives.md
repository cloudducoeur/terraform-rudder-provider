# GroupRuleComplianceDirectives

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the directive | 
**Name** | **string** | Name of the directive | 
**PolicyMode** | **string** | Policy mode of the rule | 
**Compliance** | **float32** | Directive compliance level | 
**ComplianceDetails** | [**GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails**](GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails.md) |  | 
**Components** | [**[]GroupRuleComplianceDirectivesComponentsInner**](GroupRuleComplianceDirectivesComponentsInner.md) |  | 

## Methods

### NewGroupRuleComplianceDirectives

`func NewGroupRuleComplianceDirectives(id string, name string, policyMode string, compliance float32, complianceDetails GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails, components []GroupRuleComplianceDirectivesComponentsInner, ) *GroupRuleComplianceDirectives`

NewGroupRuleComplianceDirectives instantiates a new GroupRuleComplianceDirectives object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRuleComplianceDirectivesWithDefaults

`func NewGroupRuleComplianceDirectivesWithDefaults() *GroupRuleComplianceDirectives`

NewGroupRuleComplianceDirectivesWithDefaults instantiates a new GroupRuleComplianceDirectives object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupRuleComplianceDirectives) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupRuleComplianceDirectives) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupRuleComplianceDirectives) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupRuleComplianceDirectives) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRuleComplianceDirectives) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRuleComplianceDirectives) SetName(v string)`

SetName sets Name field to given value.


### GetPolicyMode

`func (o *GroupRuleComplianceDirectives) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *GroupRuleComplianceDirectives) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *GroupRuleComplianceDirectives) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.


### GetCompliance

`func (o *GroupRuleComplianceDirectives) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GroupRuleComplianceDirectives) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GroupRuleComplianceDirectives) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GroupRuleComplianceDirectives) GetComplianceDetails() GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GroupRuleComplianceDirectives) GetComplianceDetailsOk() (*GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GroupRuleComplianceDirectives) SetComplianceDetails(v GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetComponents

`func (o *GroupRuleComplianceDirectives) GetComponents() []GroupRuleComplianceDirectivesComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GroupRuleComplianceDirectives) GetComponentsOk() (*[]GroupRuleComplianceDirectivesComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GroupRuleComplianceDirectives) SetComponents(v []GroupRuleComplianceDirectivesComponentsInner)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


