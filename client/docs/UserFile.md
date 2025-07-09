# UserFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | A new username or an empty string to ignore the update of username. | 
**Password** | **string** | This password will be hashed for you if the &#x60;isPreHashed&#x60; is set on false. Empty string means that no update is requested. | 
**Permissions** | **[]string** | Roles or authorizations | 
**IsPreHashed** | **bool** | If you want to provide hashed password set this property to &#x60;true&#x60; otherwise we will hash the plain password and store the hash | 

## Methods

### NewUserFile

`func NewUserFile(username string, password string, permissions []string, isPreHashed bool, ) *UserFile`

NewUserFile instantiates a new UserFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFileWithDefaults

`func NewUserFileWithDefaults() *UserFile`

NewUserFileWithDefaults instantiates a new UserFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserFile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserFile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserFile) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserFile) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserFile) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserFile) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPermissions

`func (o *UserFile) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserFile) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserFile) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetIsPreHashed

`func (o *UserFile) GetIsPreHashed() bool`

GetIsPreHashed returns the IsPreHashed field if non-nil, zero value otherwise.

### GetIsPreHashedOk

`func (o *UserFile) GetIsPreHashedOk() (*bool, bool)`

GetIsPreHashedOk returns a tuple with the IsPreHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreHashed

`func (o *UserFile) SetIsPreHashed(v bool)`

SetIsPreHashed sets IsPreHashed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


