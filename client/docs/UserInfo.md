# UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**OtherInfo** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUserInfo

`func NewUserInfo() *UserInfo`

NewUserInfo instantiates a new UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoWithDefaults

`func NewUserInfoWithDefaults() *UserInfo`

NewUserInfoWithDefaults instantiates a new UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOtherInfo

`func (o *UserInfo) GetOtherInfo() map[string]string`

GetOtherInfo returns the OtherInfo field if non-nil, zero value otherwise.

### GetOtherInfoOk

`func (o *UserInfo) GetOtherInfoOk() (*map[string]string, bool)`

GetOtherInfoOk returns a tuple with the OtherInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherInfo

`func (o *UserInfo) SetOtherInfo(v map[string]string)`

SetOtherInfo sets OtherInfo field to given value.

### HasOtherInfo

`func (o *UserInfo) HasOtherInfo() bool`

HasOtherInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


