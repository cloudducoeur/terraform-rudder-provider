# UpdateUser200ResponseDataUpdatedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Given username | 
**Password** | **string** | Given password, empty means no update was made | 
**Permissions** | **[]string** | Given user permissions | 

## Methods

### NewUpdateUser200ResponseDataUpdatedUser

`func NewUpdateUser200ResponseDataUpdatedUser(username string, password string, permissions []string, ) *UpdateUser200ResponseDataUpdatedUser`

NewUpdateUser200ResponseDataUpdatedUser instantiates a new UpdateUser200ResponseDataUpdatedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUser200ResponseDataUpdatedUserWithDefaults

`func NewUpdateUser200ResponseDataUpdatedUserWithDefaults() *UpdateUser200ResponseDataUpdatedUser`

NewUpdateUser200ResponseDataUpdatedUserWithDefaults instantiates a new UpdateUser200ResponseDataUpdatedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UpdateUser200ResponseDataUpdatedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUser200ResponseDataUpdatedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUser200ResponseDataUpdatedUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UpdateUser200ResponseDataUpdatedUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUser200ResponseDataUpdatedUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUser200ResponseDataUpdatedUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPermissions

`func (o *UpdateUser200ResponseDataUpdatedUser) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateUser200ResponseDataUpdatedUser) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateUser200ResponseDataUpdatedUser) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


