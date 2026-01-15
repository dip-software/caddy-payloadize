# CreateKeyPayload2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | **string** | The admin user | 
**Expiry** | **string** | Expiry time time.Duration format. A value of 0s means no expiry | 
**Format** | **string** | The format of the key | 

## Methods

### NewCreateKeyPayload2

`func NewCreateKeyPayload2(admin string, expiry string, format string, ) *CreateKeyPayload2`

NewCreateKeyPayload2 instantiates a new CreateKeyPayload2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyPayload2WithDefaults

`func NewCreateKeyPayload2WithDefaults() *CreateKeyPayload2`

NewCreateKeyPayload2WithDefaults instantiates a new CreateKeyPayload2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *CreateKeyPayload2) GetAdmin() string`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *CreateKeyPayload2) GetAdminOk() (*string, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *CreateKeyPayload2) SetAdmin(v string)`

SetAdmin sets Admin field to given value.


### GetExpiry

`func (o *CreateKeyPayload2) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *CreateKeyPayload2) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *CreateKeyPayload2) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.


### GetFormat

`func (o *CreateKeyPayload2) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateKeyPayload2) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateKeyPayload2) SetFormat(v string)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


