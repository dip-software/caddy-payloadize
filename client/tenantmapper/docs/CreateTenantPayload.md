# CreateTenantPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminClaim** | Pointer to **string** | The admin claim of the tenant | [optional] 
**Admins** | Pointer to **string** | Command separated list of admins | [optional] 
**EditorClaim** | Pointer to **string** | The editor claim of the tenant | [optional] 
**Guid** | **string** | The GUID of the IDP | 
**LogRetentionPeriod** | Pointer to **string** | The log retention period of the tenant | [optional] 
**Logging** | Pointer to **bool** | Logging support enabled for this tenant | [optional] 
**Metrics** | Pointer to **bool** | Metrics support enabled for this tenant | [optional] 
**MetricsRetentionPeriod** | Pointer to **string** | The metrics retention period of the tenant | [optional] 
**Name** | **string** | Name of the tenant | 
**TraceRetentionPeriod** | Pointer to **string** | The trace retention period of the tenant | [optional] 
**Tracing** | Pointer to **bool** | Tracing support enabled for this tenant | [optional] 
**ViewerClaim** | Pointer to **string** | The viewer claim of the tenant | [optional] 

## Methods

### NewCreateTenantPayload

`func NewCreateTenantPayload(guid string, name string, ) *CreateTenantPayload`

NewCreateTenantPayload instantiates a new CreateTenantPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantPayloadWithDefaults

`func NewCreateTenantPayloadWithDefaults() *CreateTenantPayload`

NewCreateTenantPayloadWithDefaults instantiates a new CreateTenantPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminClaim

`func (o *CreateTenantPayload) GetAdminClaim() string`

GetAdminClaim returns the AdminClaim field if non-nil, zero value otherwise.

### GetAdminClaimOk

`func (o *CreateTenantPayload) GetAdminClaimOk() (*string, bool)`

GetAdminClaimOk returns a tuple with the AdminClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminClaim

`func (o *CreateTenantPayload) SetAdminClaim(v string)`

SetAdminClaim sets AdminClaim field to given value.

### HasAdminClaim

`func (o *CreateTenantPayload) HasAdminClaim() bool`

HasAdminClaim returns a boolean if a field has been set.

### GetAdmins

`func (o *CreateTenantPayload) GetAdmins() string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *CreateTenantPayload) GetAdminsOk() (*string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *CreateTenantPayload) SetAdmins(v string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *CreateTenantPayload) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetEditorClaim

`func (o *CreateTenantPayload) GetEditorClaim() string`

GetEditorClaim returns the EditorClaim field if non-nil, zero value otherwise.

### GetEditorClaimOk

`func (o *CreateTenantPayload) GetEditorClaimOk() (*string, bool)`

GetEditorClaimOk returns a tuple with the EditorClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorClaim

`func (o *CreateTenantPayload) SetEditorClaim(v string)`

SetEditorClaim sets EditorClaim field to given value.

### HasEditorClaim

`func (o *CreateTenantPayload) HasEditorClaim() bool`

HasEditorClaim returns a boolean if a field has been set.

### GetGuid

`func (o *CreateTenantPayload) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *CreateTenantPayload) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *CreateTenantPayload) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLogRetentionPeriod

`func (o *CreateTenantPayload) GetLogRetentionPeriod() string`

GetLogRetentionPeriod returns the LogRetentionPeriod field if non-nil, zero value otherwise.

### GetLogRetentionPeriodOk

`func (o *CreateTenantPayload) GetLogRetentionPeriodOk() (*string, bool)`

GetLogRetentionPeriodOk returns a tuple with the LogRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetentionPeriod

`func (o *CreateTenantPayload) SetLogRetentionPeriod(v string)`

SetLogRetentionPeriod sets LogRetentionPeriod field to given value.

### HasLogRetentionPeriod

`func (o *CreateTenantPayload) HasLogRetentionPeriod() bool`

HasLogRetentionPeriod returns a boolean if a field has been set.

### GetLogging

`func (o *CreateTenantPayload) GetLogging() bool`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *CreateTenantPayload) GetLoggingOk() (*bool, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *CreateTenantPayload) SetLogging(v bool)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *CreateTenantPayload) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetMetrics

`func (o *CreateTenantPayload) GetMetrics() bool`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CreateTenantPayload) GetMetricsOk() (*bool, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CreateTenantPayload) SetMetrics(v bool)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CreateTenantPayload) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricsRetentionPeriod

`func (o *CreateTenantPayload) GetMetricsRetentionPeriod() string`

GetMetricsRetentionPeriod returns the MetricsRetentionPeriod field if non-nil, zero value otherwise.

### GetMetricsRetentionPeriodOk

`func (o *CreateTenantPayload) GetMetricsRetentionPeriodOk() (*string, bool)`

GetMetricsRetentionPeriodOk returns a tuple with the MetricsRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionPeriod

`func (o *CreateTenantPayload) SetMetricsRetentionPeriod(v string)`

SetMetricsRetentionPeriod sets MetricsRetentionPeriod field to given value.

### HasMetricsRetentionPeriod

`func (o *CreateTenantPayload) HasMetricsRetentionPeriod() bool`

HasMetricsRetentionPeriod returns a boolean if a field has been set.

### GetName

`func (o *CreateTenantPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTenantPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTenantPayload) SetName(v string)`

SetName sets Name field to given value.


### GetTraceRetentionPeriod

`func (o *CreateTenantPayload) GetTraceRetentionPeriod() string`

GetTraceRetentionPeriod returns the TraceRetentionPeriod field if non-nil, zero value otherwise.

### GetTraceRetentionPeriodOk

`func (o *CreateTenantPayload) GetTraceRetentionPeriodOk() (*string, bool)`

GetTraceRetentionPeriodOk returns a tuple with the TraceRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRetentionPeriod

`func (o *CreateTenantPayload) SetTraceRetentionPeriod(v string)`

SetTraceRetentionPeriod sets TraceRetentionPeriod field to given value.

### HasTraceRetentionPeriod

`func (o *CreateTenantPayload) HasTraceRetentionPeriod() bool`

HasTraceRetentionPeriod returns a boolean if a field has been set.

### GetTracing

`func (o *CreateTenantPayload) GetTracing() bool`

GetTracing returns the Tracing field if non-nil, zero value otherwise.

### GetTracingOk

`func (o *CreateTenantPayload) GetTracingOk() (*bool, bool)`

GetTracingOk returns a tuple with the Tracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracing

`func (o *CreateTenantPayload) SetTracing(v bool)`

SetTracing sets Tracing field to given value.

### HasTracing

`func (o *CreateTenantPayload) HasTracing() bool`

HasTracing returns a boolean if a field has been set.

### GetViewerClaim

`func (o *CreateTenantPayload) GetViewerClaim() string`

GetViewerClaim returns the ViewerClaim field if non-nil, zero value otherwise.

### GetViewerClaimOk

`func (o *CreateTenantPayload) GetViewerClaimOk() (*string, bool)`

GetViewerClaimOk returns a tuple with the ViewerClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerClaim

`func (o *CreateTenantPayload) SetViewerClaim(v string)`

SetViewerClaim sets ViewerClaim field to given value.

### HasViewerClaim

`func (o *CreateTenantPayload) HasViewerClaim() bool`

HasViewerClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


