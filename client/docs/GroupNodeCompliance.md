# GroupNodeCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the node | 
**Name** | **string** | Name of the node | 
**PolicyMode** | **string** | Policy mode of the node | 
**Compliance** | **float32** | Node compliance level | 
**ComplianceDetails** | [**NodeComplianceComponentComplianceDetails**](NodeComplianceComponentComplianceDetails.md) |  | 
**Rules** | [**[]GroupNodeComplianceRules**](GroupNodeComplianceRules.md) |  | 

## Methods

### NewGroupNodeCompliance

`func NewGroupNodeCompliance(id string, name string, policyMode string, compliance float32, complianceDetails NodeComplianceComponentComplianceDetails, rules []GroupNodeComplianceRules, ) *GroupNodeCompliance`

NewGroupNodeCompliance instantiates a new GroupNodeCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupNodeComplianceWithDefaults

`func NewGroupNodeComplianceWithDefaults() *GroupNodeCompliance`

NewGroupNodeComplianceWithDefaults instantiates a new GroupNodeCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupNodeCompliance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupNodeCompliance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupNodeCompliance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupNodeCompliance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupNodeCompliance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupNodeCompliance) SetName(v string)`

SetName sets Name field to given value.


### GetPolicyMode

`func (o *GroupNodeCompliance) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *GroupNodeCompliance) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *GroupNodeCompliance) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.


### GetCompliance

`func (o *GroupNodeCompliance) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GroupNodeCompliance) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GroupNodeCompliance) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GroupNodeCompliance) GetComplianceDetails() NodeComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GroupNodeCompliance) GetComplianceDetailsOk() (*NodeComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GroupNodeCompliance) SetComplianceDetails(v NodeComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetRules

`func (o *GroupNodeCompliance) GetRules() []GroupNodeComplianceRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GroupNodeCompliance) GetRulesOk() (*[]GroupNodeComplianceRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GroupNodeCompliance) SetRules(v []GroupNodeComplianceRules)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


