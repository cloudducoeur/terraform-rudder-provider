# GroupRuleCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the rule | 
**Name** | **string** | Name of the rule | 
**Mode** | Pointer to **string** |  | [optional] 
**PolicyMode** | Pointer to **string** | Policy mode of the rule | [optional] 
**Compliance** | **float32** | Directive compliance level | 
**ComplianceDetails** | [**NodeComplianceComponentComplianceDetails**](NodeComplianceComponentComplianceDetails.md) |  | 
**Directives** | [**[]GroupRuleComplianceDirectives**](GroupRuleComplianceDirectives.md) |  | 

## Methods

### NewGroupRuleCompliance

`func NewGroupRuleCompliance(id string, name string, compliance float32, complianceDetails NodeComplianceComponentComplianceDetails, directives []GroupRuleComplianceDirectives, ) *GroupRuleCompliance`

NewGroupRuleCompliance instantiates a new GroupRuleCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRuleComplianceWithDefaults

`func NewGroupRuleComplianceWithDefaults() *GroupRuleCompliance`

NewGroupRuleComplianceWithDefaults instantiates a new GroupRuleCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupRuleCompliance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupRuleCompliance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupRuleCompliance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupRuleCompliance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRuleCompliance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRuleCompliance) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *GroupRuleCompliance) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GroupRuleCompliance) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GroupRuleCompliance) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GroupRuleCompliance) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPolicyMode

`func (o *GroupRuleCompliance) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *GroupRuleCompliance) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *GroupRuleCompliance) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.

### HasPolicyMode

`func (o *GroupRuleCompliance) HasPolicyMode() bool`

HasPolicyMode returns a boolean if a field has been set.

### GetCompliance

`func (o *GroupRuleCompliance) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GroupRuleCompliance) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GroupRuleCompliance) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GroupRuleCompliance) GetComplianceDetails() NodeComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GroupRuleCompliance) GetComplianceDetailsOk() (*NodeComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GroupRuleCompliance) SetComplianceDetails(v NodeComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetDirectives

`func (o *GroupRuleCompliance) GetDirectives() []GroupRuleComplianceDirectives`

GetDirectives returns the Directives field if non-nil, zero value otherwise.

### GetDirectivesOk

`func (o *GroupRuleCompliance) GetDirectivesOk() (*[]GroupRuleComplianceDirectives, bool)`

GetDirectivesOk returns a tuple with the Directives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectives

`func (o *GroupRuleCompliance) SetDirectives(v []GroupRuleComplianceDirectives)`

SetDirectives sets Directives field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


