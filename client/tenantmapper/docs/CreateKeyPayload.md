# CreateKeyPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | **string** | The admin user | 
**Expiry** | **string** | Expiry time time.Duration format. A value of 0s means no expiry | 
**Format** | **string** | The format of the key | 
**IdToken** | **string** | ID token | 
**Tenant** | **string** | Name of the tenant | 

## Methods

### NewCreateKeyPayload

`func NewCreateKeyPayload(admin string, expiry string, format string, idToken string, tenant string, ) *CreateKeyPayload`

NewCreateKeyPayload instantiates a new CreateKeyPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyPayloadWithDefaults

`func NewCreateKeyPayloadWithDefaults() *CreateKeyPayload`

NewCreateKeyPayloadWithDefaults instantiates a new CreateKeyPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *CreateKeyPayload) GetAdmin() string`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *CreateKeyPayload) GetAdminOk() (*string, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *CreateKeyPayload) SetAdmin(v string)`

SetAdmin sets Admin field to given value.


### GetExpiry

`func (o *CreateKeyPayload) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *CreateKeyPayload) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *CreateKeyPayload) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.


### GetFormat

`func (o *CreateKeyPayload) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateKeyPayload) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateKeyPayload) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetIdToken

`func (o *CreateKeyPayload) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *CreateKeyPayload) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *CreateKeyPayload) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.


### GetTenant

`func (o *CreateKeyPayload) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreateKeyPayload) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreateKeyPayload) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


