# GroupNodeComplianceRulesDirectivesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the directive | 
**Name** | **string** | Name of the directive | 
**PolicyMode** | Pointer to **string** | Policy mode of the directive | [optional] 
**Compliance** | **float32** | Directive compliance level | 
**ComplianceDetails** | [**RuleComplianceComponentComplianceDetails**](RuleComplianceComponentComplianceDetails.md) |  | 
**Components** | [**[]GroupNodeComplianceRulesDirectivesInnerComponentsInner**](GroupNodeComplianceRulesDirectivesInnerComponentsInner.md) |  | 

## Methods

### NewGroupNodeComplianceRulesDirectivesInner

`func NewGroupNodeComplianceRulesDirectivesInner(id string, name string, compliance float32, complianceDetails RuleComplianceComponentComplianceDetails, components []GroupNodeComplianceRulesDirectivesInnerComponentsInner, ) *GroupNodeComplianceRulesDirectivesInner`

NewGroupNodeComplianceRulesDirectivesInner instantiates a new GroupNodeComplianceRulesDirectivesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupNodeComplianceRulesDirectivesInnerWithDefaults

`func NewGroupNodeComplianceRulesDirectivesInnerWithDefaults() *GroupNodeComplianceRulesDirectivesInner`

NewGroupNodeComplianceRulesDirectivesInnerWithDefaults instantiates a new GroupNodeComplianceRulesDirectivesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupNodeComplianceRulesDirectivesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupNodeComplianceRulesDirectivesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupNodeComplianceRulesDirectivesInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupNodeComplianceRulesDirectivesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupNodeComplianceRulesDirectivesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupNodeComplianceRulesDirectivesInner) SetName(v string)`

SetName sets Name field to given value.


### GetPolicyMode

`func (o *GroupNodeComplianceRulesDirectivesInner) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *GroupNodeComplianceRulesDirectivesInner) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *GroupNodeComplianceRulesDirectivesInner) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.

### HasPolicyMode

`func (o *GroupNodeComplianceRulesDirectivesInner) HasPolicyMode() bool`

HasPolicyMode returns a boolean if a field has been set.

### GetCompliance

`func (o *GroupNodeComplianceRulesDirectivesInner) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GroupNodeComplianceRulesDirectivesInner) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GroupNodeComplianceRulesDirectivesInner) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GroupNodeComplianceRulesDirectivesInner) GetComplianceDetails() RuleComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GroupNodeComplianceRulesDirectivesInner) GetComplianceDetailsOk() (*RuleComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GroupNodeComplianceRulesDirectivesInner) SetComplianceDetails(v RuleComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetComponents

`func (o *GroupNodeComplianceRulesDirectivesInner) GetComponents() []GroupNodeComplianceRulesDirectivesInnerComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GroupNodeComplianceRulesDirectivesInner) GetComponentsOk() (*[]GroupNodeComplianceRulesDirectivesInnerComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GroupNodeComplianceRulesDirectivesInner) SetComponents(v []GroupNodeComplianceRulesDirectivesInnerComponentsInner)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


