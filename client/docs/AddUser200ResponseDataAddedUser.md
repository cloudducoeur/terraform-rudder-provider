# AddUser200ResponseDataAddedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | New username | 
**Password** | **string** | New given password, hashed or not | 
**Permissions** | **[]string** | Permissions of new user | 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**OtherInfo** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAddUser200ResponseDataAddedUser

`func NewAddUser200ResponseDataAddedUser(username string, password string, permissions []string, ) *AddUser200ResponseDataAddedUser`

NewAddUser200ResponseDataAddedUser instantiates a new AddUser200ResponseDataAddedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUser200ResponseDataAddedUserWithDefaults

`func NewAddUser200ResponseDataAddedUserWithDefaults() *AddUser200ResponseDataAddedUser`

NewAddUser200ResponseDataAddedUserWithDefaults instantiates a new AddUser200ResponseDataAddedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AddUser200ResponseDataAddedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddUser200ResponseDataAddedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddUser200ResponseDataAddedUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AddUser200ResponseDataAddedUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddUser200ResponseDataAddedUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddUser200ResponseDataAddedUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPermissions

`func (o *AddUser200ResponseDataAddedUser) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AddUser200ResponseDataAddedUser) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AddUser200ResponseDataAddedUser) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetName

`func (o *AddUser200ResponseDataAddedUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddUser200ResponseDataAddedUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddUser200ResponseDataAddedUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddUser200ResponseDataAddedUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *AddUser200ResponseDataAddedUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddUser200ResponseDataAddedUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddUser200ResponseDataAddedUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddUser200ResponseDataAddedUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOtherInfo

`func (o *AddUser200ResponseDataAddedUser) GetOtherInfo() map[string]string`

GetOtherInfo returns the OtherInfo field if non-nil, zero value otherwise.

### GetOtherInfoOk

`func (o *AddUser200ResponseDataAddedUser) GetOtherInfoOk() (*map[string]string, bool)`

GetOtherInfoOk returns a tuple with the OtherInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherInfo

`func (o *AddUser200ResponseDataAddedUser) SetOtherInfo(v map[string]string)`

SetOtherInfo sets OtherInfo field to given value.

### HasOtherInfo

`func (o *AddUser200ResponseDataAddedUser) HasOtherInfo() bool`

HasOtherInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


