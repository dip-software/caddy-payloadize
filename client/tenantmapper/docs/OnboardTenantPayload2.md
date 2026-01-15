# OnboardTenantPayload2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admins** | Pointer to **string** | Comma separated list of authorized admins | [optional] 
**AwsRegion** | **string** | The AWS region to onboard onto | 
**GhEditorTeam** | Pointer to **string** | The GitHub editor team | [optional] 
**GhViewerTeam** | Pointer to **string** | The GitHub viewer team | [optional] 
**IamOrgId** | Pointer to **string** | The IAM Org ID | [optional] 
**LogRetention** | Pointer to **string** | The log retention period of the tenant | [optional] 
**MetricsRetention** | Pointer to **string** | The metrics retention period of the tenant | [optional] 
**Name** | **string** | Tenant name | 
**TraceRetention** | Pointer to **string** | The tracing retention period of the tenant | [optional] 

## Methods

### NewOnboardTenantPayload2

`func NewOnboardTenantPayload2(awsRegion string, name string, ) *OnboardTenantPayload2`

NewOnboardTenantPayload2 instantiates a new OnboardTenantPayload2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardTenantPayload2WithDefaults

`func NewOnboardTenantPayload2WithDefaults() *OnboardTenantPayload2`

NewOnboardTenantPayload2WithDefaults instantiates a new OnboardTenantPayload2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmins

`func (o *OnboardTenantPayload2) GetAdmins() string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *OnboardTenantPayload2) GetAdminsOk() (*string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *OnboardTenantPayload2) SetAdmins(v string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *OnboardTenantPayload2) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetAwsRegion

`func (o *OnboardTenantPayload2) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *OnboardTenantPayload2) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *OnboardTenantPayload2) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### GetGhEditorTeam

`func (o *OnboardTenantPayload2) GetGhEditorTeam() string`

GetGhEditorTeam returns the GhEditorTeam field if non-nil, zero value otherwise.

### GetGhEditorTeamOk

`func (o *OnboardTenantPayload2) GetGhEditorTeamOk() (*string, bool)`

GetGhEditorTeamOk returns a tuple with the GhEditorTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhEditorTeam

`func (o *OnboardTenantPayload2) SetGhEditorTeam(v string)`

SetGhEditorTeam sets GhEditorTeam field to given value.

### HasGhEditorTeam

`func (o *OnboardTenantPayload2) HasGhEditorTeam() bool`

HasGhEditorTeam returns a boolean if a field has been set.

### GetGhViewerTeam

`func (o *OnboardTenantPayload2) GetGhViewerTeam() string`

GetGhViewerTeam returns the GhViewerTeam field if non-nil, zero value otherwise.

### GetGhViewerTeamOk

`func (o *OnboardTenantPayload2) GetGhViewerTeamOk() (*string, bool)`

GetGhViewerTeamOk returns a tuple with the GhViewerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhViewerTeam

`func (o *OnboardTenantPayload2) SetGhViewerTeam(v string)`

SetGhViewerTeam sets GhViewerTeam field to given value.

### HasGhViewerTeam

`func (o *OnboardTenantPayload2) HasGhViewerTeam() bool`

HasGhViewerTeam returns a boolean if a field has been set.

### GetIamOrgId

`func (o *OnboardTenantPayload2) GetIamOrgId() string`

GetIamOrgId returns the IamOrgId field if non-nil, zero value otherwise.

### GetIamOrgIdOk

`func (o *OnboardTenantPayload2) GetIamOrgIdOk() (*string, bool)`

GetIamOrgIdOk returns a tuple with the IamOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamOrgId

`func (o *OnboardTenantPayload2) SetIamOrgId(v string)`

SetIamOrgId sets IamOrgId field to given value.

### HasIamOrgId

`func (o *OnboardTenantPayload2) HasIamOrgId() bool`

HasIamOrgId returns a boolean if a field has been set.

### GetLogRetention

`func (o *OnboardTenantPayload2) GetLogRetention() string`

GetLogRetention returns the LogRetention field if non-nil, zero value otherwise.

### GetLogRetentionOk

`func (o *OnboardTenantPayload2) GetLogRetentionOk() (*string, bool)`

GetLogRetentionOk returns a tuple with the LogRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetention

`func (o *OnboardTenantPayload2) SetLogRetention(v string)`

SetLogRetention sets LogRetention field to given value.

### HasLogRetention

`func (o *OnboardTenantPayload2) HasLogRetention() bool`

HasLogRetention returns a boolean if a field has been set.

### GetMetricsRetention

`func (o *OnboardTenantPayload2) GetMetricsRetention() string`

GetMetricsRetention returns the MetricsRetention field if non-nil, zero value otherwise.

### GetMetricsRetentionOk

`func (o *OnboardTenantPayload2) GetMetricsRetentionOk() (*string, bool)`

GetMetricsRetentionOk returns a tuple with the MetricsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetention

`func (o *OnboardTenantPayload2) SetMetricsRetention(v string)`

SetMetricsRetention sets MetricsRetention field to given value.

### HasMetricsRetention

`func (o *OnboardTenantPayload2) HasMetricsRetention() bool`

HasMetricsRetention returns a boolean if a field has been set.

### GetName

`func (o *OnboardTenantPayload2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnboardTenantPayload2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnboardTenantPayload2) SetName(v string)`

SetName sets Name field to given value.


### GetTraceRetention

`func (o *OnboardTenantPayload2) GetTraceRetention() string`

GetTraceRetention returns the TraceRetention field if non-nil, zero value otherwise.

### GetTraceRetentionOk

`func (o *OnboardTenantPayload2) GetTraceRetentionOk() (*string, bool)`

GetTraceRetentionOk returns a tuple with the TraceRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRetention

`func (o *OnboardTenantPayload2) SetTraceRetention(v string)`

SetTraceRetention sets TraceRetention field to given value.

### HasTraceRetention

`func (o *OnboardTenantPayload2) HasTraceRetention() bool`

HasTraceRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


