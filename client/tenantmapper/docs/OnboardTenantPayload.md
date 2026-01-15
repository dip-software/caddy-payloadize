# OnboardTenantPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admins** | Pointer to **string** | Comma separated list of authorized admins | [optional] 
**AwsRegion** | **string** | The AWS region to onboard onto | 
**GhEditorTeam** | Pointer to **string** | The GitHub editor team | [optional] 
**GhViewerTeam** | Pointer to **string** | The GitHub viewer team | [optional] 
**IamOrgId** | Pointer to **string** | The IAM Org ID | [optional] 
**IdToken** | **string** | ID token | 
**LogRetention** | Pointer to **string** | The log retention period of the tenant | [optional] 
**MetricsRetention** | Pointer to **string** | The metrics retention period of the tenant | [optional] 
**Name** | **string** | Tenant name | 
**TraceRetention** | Pointer to **string** | The tracing retention period of the tenant | [optional] 

## Methods

### NewOnboardTenantPayload

`func NewOnboardTenantPayload(awsRegion string, idToken string, name string, ) *OnboardTenantPayload`

NewOnboardTenantPayload instantiates a new OnboardTenantPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardTenantPayloadWithDefaults

`func NewOnboardTenantPayloadWithDefaults() *OnboardTenantPayload`

NewOnboardTenantPayloadWithDefaults instantiates a new OnboardTenantPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmins

`func (o *OnboardTenantPayload) GetAdmins() string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *OnboardTenantPayload) GetAdminsOk() (*string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *OnboardTenantPayload) SetAdmins(v string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *OnboardTenantPayload) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetAwsRegion

`func (o *OnboardTenantPayload) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *OnboardTenantPayload) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *OnboardTenantPayload) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### GetGhEditorTeam

`func (o *OnboardTenantPayload) GetGhEditorTeam() string`

GetGhEditorTeam returns the GhEditorTeam field if non-nil, zero value otherwise.

### GetGhEditorTeamOk

`func (o *OnboardTenantPayload) GetGhEditorTeamOk() (*string, bool)`

GetGhEditorTeamOk returns a tuple with the GhEditorTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhEditorTeam

`func (o *OnboardTenantPayload) SetGhEditorTeam(v string)`

SetGhEditorTeam sets GhEditorTeam field to given value.

### HasGhEditorTeam

`func (o *OnboardTenantPayload) HasGhEditorTeam() bool`

HasGhEditorTeam returns a boolean if a field has been set.

### GetGhViewerTeam

`func (o *OnboardTenantPayload) GetGhViewerTeam() string`

GetGhViewerTeam returns the GhViewerTeam field if non-nil, zero value otherwise.

### GetGhViewerTeamOk

`func (o *OnboardTenantPayload) GetGhViewerTeamOk() (*string, bool)`

GetGhViewerTeamOk returns a tuple with the GhViewerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhViewerTeam

`func (o *OnboardTenantPayload) SetGhViewerTeam(v string)`

SetGhViewerTeam sets GhViewerTeam field to given value.

### HasGhViewerTeam

`func (o *OnboardTenantPayload) HasGhViewerTeam() bool`

HasGhViewerTeam returns a boolean if a field has been set.

### GetIamOrgId

`func (o *OnboardTenantPayload) GetIamOrgId() string`

GetIamOrgId returns the IamOrgId field if non-nil, zero value otherwise.

### GetIamOrgIdOk

`func (o *OnboardTenantPayload) GetIamOrgIdOk() (*string, bool)`

GetIamOrgIdOk returns a tuple with the IamOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamOrgId

`func (o *OnboardTenantPayload) SetIamOrgId(v string)`

SetIamOrgId sets IamOrgId field to given value.

### HasIamOrgId

`func (o *OnboardTenantPayload) HasIamOrgId() bool`

HasIamOrgId returns a boolean if a field has been set.

### GetIdToken

`func (o *OnboardTenantPayload) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *OnboardTenantPayload) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *OnboardTenantPayload) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.


### GetLogRetention

`func (o *OnboardTenantPayload) GetLogRetention() string`

GetLogRetention returns the LogRetention field if non-nil, zero value otherwise.

### GetLogRetentionOk

`func (o *OnboardTenantPayload) GetLogRetentionOk() (*string, bool)`

GetLogRetentionOk returns a tuple with the LogRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetention

`func (o *OnboardTenantPayload) SetLogRetention(v string)`

SetLogRetention sets LogRetention field to given value.

### HasLogRetention

`func (o *OnboardTenantPayload) HasLogRetention() bool`

HasLogRetention returns a boolean if a field has been set.

### GetMetricsRetention

`func (o *OnboardTenantPayload) GetMetricsRetention() string`

GetMetricsRetention returns the MetricsRetention field if non-nil, zero value otherwise.

### GetMetricsRetentionOk

`func (o *OnboardTenantPayload) GetMetricsRetentionOk() (*string, bool)`

GetMetricsRetentionOk returns a tuple with the MetricsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetention

`func (o *OnboardTenantPayload) SetMetricsRetention(v string)`

SetMetricsRetention sets MetricsRetention field to given value.

### HasMetricsRetention

`func (o *OnboardTenantPayload) HasMetricsRetention() bool`

HasMetricsRetention returns a boolean if a field has been set.

### GetName

`func (o *OnboardTenantPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnboardTenantPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnboardTenantPayload) SetName(v string)`

SetName sets Name field to given value.


### GetTraceRetention

`func (o *OnboardTenantPayload) GetTraceRetention() string`

GetTraceRetention returns the TraceRetention field if non-nil, zero value otherwise.

### GetTraceRetentionOk

`func (o *OnboardTenantPayload) GetTraceRetentionOk() (*string, bool)`

GetTraceRetentionOk returns a tuple with the TraceRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRetention

`func (o *OnboardTenantPayload) SetTraceRetention(v string)`

SetTraceRetention sets TraceRetention field to given value.

### HasTraceRetention

`func (o *OnboardTenantPayload) HasTraceRetention() bool`

HasTraceRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


