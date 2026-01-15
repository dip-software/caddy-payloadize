# UpdateTenantPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminClaim** | Pointer to **string** | The admin claim of the tenant | [optional] 
**Admins** | Pointer to **string** | Comma separated list of admins | [optional] 
**EditorClaim** | Pointer to **string** | The editor claim of the tenant | [optional] 
**LogRetentionPeriod** | Pointer to **string** | The log retention period of the tenant | [optional] 
**Logging** | Pointer to **bool** | Logging support enabled for this tenant | [optional] 
**Metrics** | Pointer to **bool** | Metrics support enabled for this tenant | [optional] 
**MetricsRetentionPeriod** | Pointer to **string** | The metrics retention period of the tenant | [optional] 
**TenantId** | **string** | Tenant ID | 
**TraceRetentionPeriod** | Pointer to **string** | The trace retention period of the tenant | [optional] 
**Tracing** | Pointer to **bool** | Tracing support enabled for this tenant | [optional] 
**ViewerClaim** | Pointer to **string** | The viewer claim of the tenant | [optional] 

## Methods

### NewUpdateTenantPayload

`func NewUpdateTenantPayload(tenantId string, ) *UpdateTenantPayload`

NewUpdateTenantPayload instantiates a new UpdateTenantPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantPayloadWithDefaults

`func NewUpdateTenantPayloadWithDefaults() *UpdateTenantPayload`

NewUpdateTenantPayloadWithDefaults instantiates a new UpdateTenantPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminClaim

`func (o *UpdateTenantPayload) GetAdminClaim() string`

GetAdminClaim returns the AdminClaim field if non-nil, zero value otherwise.

### GetAdminClaimOk

`func (o *UpdateTenantPayload) GetAdminClaimOk() (*string, bool)`

GetAdminClaimOk returns a tuple with the AdminClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminClaim

`func (o *UpdateTenantPayload) SetAdminClaim(v string)`

SetAdminClaim sets AdminClaim field to given value.

### HasAdminClaim

`func (o *UpdateTenantPayload) HasAdminClaim() bool`

HasAdminClaim returns a boolean if a field has been set.

### GetAdmins

`func (o *UpdateTenantPayload) GetAdmins() string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *UpdateTenantPayload) GetAdminsOk() (*string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *UpdateTenantPayload) SetAdmins(v string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *UpdateTenantPayload) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetEditorClaim

`func (o *UpdateTenantPayload) GetEditorClaim() string`

GetEditorClaim returns the EditorClaim field if non-nil, zero value otherwise.

### GetEditorClaimOk

`func (o *UpdateTenantPayload) GetEditorClaimOk() (*string, bool)`

GetEditorClaimOk returns a tuple with the EditorClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorClaim

`func (o *UpdateTenantPayload) SetEditorClaim(v string)`

SetEditorClaim sets EditorClaim field to given value.

### HasEditorClaim

`func (o *UpdateTenantPayload) HasEditorClaim() bool`

HasEditorClaim returns a boolean if a field has been set.

### GetLogRetentionPeriod

`func (o *UpdateTenantPayload) GetLogRetentionPeriod() string`

GetLogRetentionPeriod returns the LogRetentionPeriod field if non-nil, zero value otherwise.

### GetLogRetentionPeriodOk

`func (o *UpdateTenantPayload) GetLogRetentionPeriodOk() (*string, bool)`

GetLogRetentionPeriodOk returns a tuple with the LogRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetentionPeriod

`func (o *UpdateTenantPayload) SetLogRetentionPeriod(v string)`

SetLogRetentionPeriod sets LogRetentionPeriod field to given value.

### HasLogRetentionPeriod

`func (o *UpdateTenantPayload) HasLogRetentionPeriod() bool`

HasLogRetentionPeriod returns a boolean if a field has been set.

### GetLogging

`func (o *UpdateTenantPayload) GetLogging() bool`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *UpdateTenantPayload) GetLoggingOk() (*bool, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *UpdateTenantPayload) SetLogging(v bool)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *UpdateTenantPayload) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetMetrics

`func (o *UpdateTenantPayload) GetMetrics() bool`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *UpdateTenantPayload) GetMetricsOk() (*bool, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *UpdateTenantPayload) SetMetrics(v bool)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *UpdateTenantPayload) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricsRetentionPeriod

`func (o *UpdateTenantPayload) GetMetricsRetentionPeriod() string`

GetMetricsRetentionPeriod returns the MetricsRetentionPeriod field if non-nil, zero value otherwise.

### GetMetricsRetentionPeriodOk

`func (o *UpdateTenantPayload) GetMetricsRetentionPeriodOk() (*string, bool)`

GetMetricsRetentionPeriodOk returns a tuple with the MetricsRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionPeriod

`func (o *UpdateTenantPayload) SetMetricsRetentionPeriod(v string)`

SetMetricsRetentionPeriod sets MetricsRetentionPeriod field to given value.

### HasMetricsRetentionPeriod

`func (o *UpdateTenantPayload) HasMetricsRetentionPeriod() bool`

HasMetricsRetentionPeriod returns a boolean if a field has been set.

### GetTenantId

`func (o *UpdateTenantPayload) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateTenantPayload) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateTenantPayload) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTraceRetentionPeriod

`func (o *UpdateTenantPayload) GetTraceRetentionPeriod() string`

GetTraceRetentionPeriod returns the TraceRetentionPeriod field if non-nil, zero value otherwise.

### GetTraceRetentionPeriodOk

`func (o *UpdateTenantPayload) GetTraceRetentionPeriodOk() (*string, bool)`

GetTraceRetentionPeriodOk returns a tuple with the TraceRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRetentionPeriod

`func (o *UpdateTenantPayload) SetTraceRetentionPeriod(v string)`

SetTraceRetentionPeriod sets TraceRetentionPeriod field to given value.

### HasTraceRetentionPeriod

`func (o *UpdateTenantPayload) HasTraceRetentionPeriod() bool`

HasTraceRetentionPeriod returns a boolean if a field has been set.

### GetTracing

`func (o *UpdateTenantPayload) GetTracing() bool`

GetTracing returns the Tracing field if non-nil, zero value otherwise.

### GetTracingOk

`func (o *UpdateTenantPayload) GetTracingOk() (*bool, bool)`

GetTracingOk returns a tuple with the Tracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracing

`func (o *UpdateTenantPayload) SetTracing(v bool)`

SetTracing sets Tracing field to given value.

### HasTracing

`func (o *UpdateTenantPayload) HasTracing() bool`

HasTracing returns a boolean if a field has been set.

### GetViewerClaim

`func (o *UpdateTenantPayload) GetViewerClaim() string`

GetViewerClaim returns the ViewerClaim field if non-nil, zero value otherwise.

### GetViewerClaimOk

`func (o *UpdateTenantPayload) GetViewerClaimOk() (*string, bool)`

GetViewerClaimOk returns a tuple with the ViewerClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerClaim

`func (o *UpdateTenantPayload) SetViewerClaim(v string)`

SetViewerClaim sets ViewerClaim field to given value.

### HasViewerClaim

`func (o *UpdateTenantPayload) HasViewerClaim() bool`

HasViewerClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


