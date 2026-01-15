# TenantAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Who created this key | [optional] 
**Description** | Pointer to **string** | Description of the key | [optional] 
**Id** | Pointer to **string** | Unique key identifier | [optional] 
**LastUsed** | Pointer to **string** | Last use of this key | [optional] 
**RevokedAt** | Pointer to **string** | When this key was revoked | [optional] 
**Scopes** | Pointer to **string** | Scopes of the key | [optional] 
**Signature** | Pointer to **string** | Signature of the key | [optional] 
**Tenant** | Pointer to **string** | Name of the tenant this key belongs to | [optional] 

## Methods

### NewTenantAPIKey

`func NewTenantAPIKey() *TenantAPIKey`

NewTenantAPIKey instantiates a new TenantAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAPIKeyWithDefaults

`func NewTenantAPIKeyWithDefaults() *TenantAPIKey`

NewTenantAPIKeyWithDefaults instantiates a new TenantAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *TenantAPIKey) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *TenantAPIKey) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *TenantAPIKey) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *TenantAPIKey) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetDescription

`func (o *TenantAPIKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantAPIKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantAPIKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantAPIKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TenantAPIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantAPIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantAPIKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantAPIKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsed

`func (o *TenantAPIKey) GetLastUsed() string`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *TenantAPIKey) GetLastUsedOk() (*string, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *TenantAPIKey) SetLastUsed(v string)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *TenantAPIKey) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetRevokedAt

`func (o *TenantAPIKey) GetRevokedAt() string`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *TenantAPIKey) GetRevokedAtOk() (*string, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *TenantAPIKey) SetRevokedAt(v string)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *TenantAPIKey) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetScopes

`func (o *TenantAPIKey) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TenantAPIKey) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TenantAPIKey) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TenantAPIKey) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSignature

`func (o *TenantAPIKey) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TenantAPIKey) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TenantAPIKey) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *TenantAPIKey) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTenant

`func (o *TenantAPIKey) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TenantAPIKey) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TenantAPIKey) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TenantAPIKey) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


