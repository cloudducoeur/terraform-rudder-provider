# UserNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | A new username | 
**Password** | **string** | This password will be hashed for you if the &#x60;isPreHashed&#x60; is set on false | 
**Permissions** | **[]string** | Roles or authorizations | 
**IsPreHashed** | **bool** | If you want to provide hashed password set this property to &#x60;true&#x60; otherwise we will hash the plain password and store the hash | 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**OtherInfo** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUserNew

`func NewUserNew(username string, password string, permissions []string, isPreHashed bool, ) *UserNew`

NewUserNew instantiates a new UserNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNewWithDefaults

`func NewUserNewWithDefaults() *UserNew`

NewUserNewWithDefaults instantiates a new UserNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserNew) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserNew) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserNew) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserNew) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserNew) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserNew) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPermissions

`func (o *UserNew) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserNew) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserNew) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetIsPreHashed

`func (o *UserNew) GetIsPreHashed() bool`

GetIsPreHashed returns the IsPreHashed field if non-nil, zero value otherwise.

### GetIsPreHashedOk

`func (o *UserNew) GetIsPreHashedOk() (*bool, bool)`

GetIsPreHashedOk returns a tuple with the IsPreHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreHashed

`func (o *UserNew) SetIsPreHashed(v bool)`

SetIsPreHashed sets IsPreHashed field to given value.


### GetName

`func (o *UserNew) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserNew) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserNew) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserNew) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UserNew) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserNew) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserNew) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserNew) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOtherInfo

`func (o *UserNew) GetOtherInfo() map[string]string`

GetOtherInfo returns the OtherInfo field if non-nil, zero value otherwise.

### GetOtherInfoOk

`func (o *UserNew) GetOtherInfoOk() (*map[string]string, bool)`

GetOtherInfoOk returns a tuple with the OtherInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherInfo

`func (o *UserNew) SetOtherInfo(v map[string]string)`

SetOtherInfo sets OtherInfo field to given value.

### HasOtherInfo

`func (o *UserNew) HasOtherInfo() bool`

HasOtherInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


