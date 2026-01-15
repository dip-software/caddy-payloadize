# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminClaim** | Pointer to **string** | The admin claim of the tenant | [optional] 
**Admins** | Pointer to **string** | Comma separated list of admins | [optional] 
**EditorClaim** | Pointer to **string** | The editor claim of the tenant | [optional] 
**Guid** | Pointer to **string** | The GUID of the IDP | [optional] 
**Id** | Pointer to **string** | Unique account identifier | [optional] 
**LogRetentionPeriod** | Pointer to **string** | The log retention period of the tenant, format is Go&#39;s time.Duration | [optional] 
**Logging** | Pointer to **bool** | Logging enabled | [optional] 
**Metrics** | Pointer to **bool** | Metrics enabled | [optional] 
**MetricsRetentionPeriod** | Pointer to **string** | The metrics retention period of the tenant, format is Go&#39;s time.Duration | [optional] 
**Name** | Pointer to **string** | Name of the tenant | [optional] 
**OrgId** | Pointer to **int64** | The Org ID of the application | [optional] 
**Profiles** | Pointer to **bool** | Profiles enabled | [optional] 
**TraceRetentionPeriod** | Pointer to **string** | The trace retention period of the tenant, format is Go&#39;s time.Duration | [optional] 
**Tracing** | Pointer to **bool** | Tracing enabled | [optional] 
**ViewerClaim** | Pointer to **string** | The viewer claim of the tenant | [optional] 

## Methods

### NewTenant

`func NewTenant() *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminClaim

`func (o *Tenant) GetAdminClaim() string`

GetAdminClaim returns the AdminClaim field if non-nil, zero value otherwise.

### GetAdminClaimOk

`func (o *Tenant) GetAdminClaimOk() (*string, bool)`

GetAdminClaimOk returns a tuple with the AdminClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminClaim

`func (o *Tenant) SetAdminClaim(v string)`

SetAdminClaim sets AdminClaim field to given value.

### HasAdminClaim

`func (o *Tenant) HasAdminClaim() bool`

HasAdminClaim returns a boolean if a field has been set.

### GetAdmins

`func (o *Tenant) GetAdmins() string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *Tenant) GetAdminsOk() (*string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *Tenant) SetAdmins(v string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *Tenant) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetEditorClaim

`func (o *Tenant) GetEditorClaim() string`

GetEditorClaim returns the EditorClaim field if non-nil, zero value otherwise.

### GetEditorClaimOk

`func (o *Tenant) GetEditorClaimOk() (*string, bool)`

GetEditorClaimOk returns a tuple with the EditorClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorClaim

`func (o *Tenant) SetEditorClaim(v string)`

SetEditorClaim sets EditorClaim field to given value.

### HasEditorClaim

`func (o *Tenant) HasEditorClaim() bool`

HasEditorClaim returns a boolean if a field has been set.

### GetGuid

`func (o *Tenant) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Tenant) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Tenant) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Tenant) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetId

`func (o *Tenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tenant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Tenant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogRetentionPeriod

`func (o *Tenant) GetLogRetentionPeriod() string`

GetLogRetentionPeriod returns the LogRetentionPeriod field if non-nil, zero value otherwise.

### GetLogRetentionPeriodOk

`func (o *Tenant) GetLogRetentionPeriodOk() (*string, bool)`

GetLogRetentionPeriodOk returns a tuple with the LogRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetentionPeriod

`func (o *Tenant) SetLogRetentionPeriod(v string)`

SetLogRetentionPeriod sets LogRetentionPeriod field to given value.

### HasLogRetentionPeriod

`func (o *Tenant) HasLogRetentionPeriod() bool`

HasLogRetentionPeriod returns a boolean if a field has been set.

### GetLogging

`func (o *Tenant) GetLogging() bool`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *Tenant) GetLoggingOk() (*bool, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *Tenant) SetLogging(v bool)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *Tenant) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetMetrics

`func (o *Tenant) GetMetrics() bool`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Tenant) GetMetricsOk() (*bool, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Tenant) SetMetrics(v bool)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Tenant) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricsRetentionPeriod

`func (o *Tenant) GetMetricsRetentionPeriod() string`

GetMetricsRetentionPeriod returns the MetricsRetentionPeriod field if non-nil, zero value otherwise.

### GetMetricsRetentionPeriodOk

`func (o *Tenant) GetMetricsRetentionPeriodOk() (*string, bool)`

GetMetricsRetentionPeriodOk returns a tuple with the MetricsRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionPeriod

`func (o *Tenant) SetMetricsRetentionPeriod(v string)`

SetMetricsRetentionPeriod sets MetricsRetentionPeriod field to given value.

### HasMetricsRetentionPeriod

`func (o *Tenant) HasMetricsRetentionPeriod() bool`

HasMetricsRetentionPeriod returns a boolean if a field has been set.

### GetName

`func (o *Tenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tenant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Tenant) GetOrgId() int64`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Tenant) GetOrgIdOk() (*int64, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Tenant) SetOrgId(v int64)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Tenant) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProfiles

`func (o *Tenant) GetProfiles() bool`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *Tenant) GetProfilesOk() (*bool, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *Tenant) SetProfiles(v bool)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *Tenant) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetTraceRetentionPeriod

`func (o *Tenant) GetTraceRetentionPeriod() string`

GetTraceRetentionPeriod returns the TraceRetentionPeriod field if non-nil, zero value otherwise.

### GetTraceRetentionPeriodOk

`func (o *Tenant) GetTraceRetentionPeriodOk() (*string, bool)`

GetTraceRetentionPeriodOk returns a tuple with the TraceRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRetentionPeriod

`func (o *Tenant) SetTraceRetentionPeriod(v string)`

SetTraceRetentionPeriod sets TraceRetentionPeriod field to given value.

### HasTraceRetentionPeriod

`func (o *Tenant) HasTraceRetentionPeriod() bool`

HasTraceRetentionPeriod returns a boolean if a field has been set.

### GetTracing

`func (o *Tenant) GetTracing() bool`

GetTracing returns the Tracing field if non-nil, zero value otherwise.

### GetTracingOk

`func (o *Tenant) GetTracingOk() (*bool, bool)`

GetTracingOk returns a tuple with the Tracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracing

`func (o *Tenant) SetTracing(v bool)`

SetTracing sets Tracing field to given value.

### HasTracing

`func (o *Tenant) HasTracing() bool`

HasTracing returns a boolean if a field has been set.

### GetViewerClaim

`func (o *Tenant) GetViewerClaim() string`

GetViewerClaim returns the ViewerClaim field if non-nil, zero value otherwise.

### GetViewerClaimOk

`func (o *Tenant) GetViewerClaimOk() (*string, bool)`

GetViewerClaimOk returns a tuple with the ViewerClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerClaim

`func (o *Tenant) SetViewerClaim(v string)`

SetViewerClaim sets ViewerClaim field to given value.

### HasViewerClaim

`func (o *Tenant) HasViewerClaim() bool`

HasViewerClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


