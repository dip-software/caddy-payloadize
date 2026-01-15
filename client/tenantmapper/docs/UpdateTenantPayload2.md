# UpdateTenantPayload2

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
**TraceRetentionPeriod** | Pointer to **string** | The trace retention period of the tenant | [optional] 
**Tracing** | Pointer to **bool** | Tracing support enabled for this tenant | [optional] 
**ViewerClaim** | Pointer to **string** | The viewer claim of the tenant | [optional] 

## Methods

### NewUpdateTenantPayload2

`func NewUpdateTenantPayload2() *UpdateTenantPayload2`

NewUpdateTenantPayload2 instantiates a new UpdateTenantPayload2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantPayload2WithDefaults

`func NewUpdateTenantPayload2WithDefaults() *UpdateTenantPayload2`

NewUpdateTenantPayload2WithDefaults instantiates a new UpdateTenantPayload2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminClaim

`func (o *UpdateTenantPayload2) GetAdminClaim() string`

GetAdminClaim returns the AdminClaim field if non-nil, zero value otherwise.

### GetAdminClaimOk

`func (o *UpdateTenantPayload2) GetAdminClaimOk() (*string, bool)`

GetAdminClaimOk returns a tuple with the AdminClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminClaim

`func (o *UpdateTenantPayload2) SetAdminClaim(v string)`

SetAdminClaim sets AdminClaim field to given value.

### HasAdminClaim

`func (o *UpdateTenantPayload2) HasAdminClaim() bool`

HasAdminClaim returns a boolean if a field has been set.

### GetAdmins

`func (o *UpdateTenantPayload2) GetAdmins() string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *UpdateTenantPayload2) GetAdminsOk() (*string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *UpdateTenantPayload2) SetAdmins(v string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *UpdateTenantPayload2) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetEditorClaim

`func (o *UpdateTenantPayload2) GetEditorClaim() string`

GetEditorClaim returns the EditorClaim field if non-nil, zero value otherwise.

### GetEditorClaimOk

`func (o *UpdateTenantPayload2) GetEditorClaimOk() (*string, bool)`

GetEditorClaimOk returns a tuple with the EditorClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorClaim

`func (o *UpdateTenantPayload2) SetEditorClaim(v string)`

SetEditorClaim sets EditorClaim field to given value.

### HasEditorClaim

`func (o *UpdateTenantPayload2) HasEditorClaim() bool`

HasEditorClaim returns a boolean if a field has been set.

### GetLogRetentionPeriod

`func (o *UpdateTenantPayload2) GetLogRetentionPeriod() string`

GetLogRetentionPeriod returns the LogRetentionPeriod field if non-nil, zero value otherwise.

### GetLogRetentionPeriodOk

`func (o *UpdateTenantPayload2) GetLogRetentionPeriodOk() (*string, bool)`

GetLogRetentionPeriodOk returns a tuple with the LogRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetentionPeriod

`func (o *UpdateTenantPayload2) SetLogRetentionPeriod(v string)`

SetLogRetentionPeriod sets LogRetentionPeriod field to given value.

### HasLogRetentionPeriod

`func (o *UpdateTenantPayload2) HasLogRetentionPeriod() bool`

HasLogRetentionPeriod returns a boolean if a field has been set.

### GetLogging

`func (o *UpdateTenantPayload2) GetLogging() bool`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *UpdateTenantPayload2) GetLoggingOk() (*bool, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *UpdateTenantPayload2) SetLogging(v bool)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *UpdateTenantPayload2) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetMetrics

`func (o *UpdateTenantPayload2) GetMetrics() bool`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *UpdateTenantPayload2) GetMetricsOk() (*bool, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *UpdateTenantPayload2) SetMetrics(v bool)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *UpdateTenantPayload2) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricsRetentionPeriod

`func (o *UpdateTenantPayload2) GetMetricsRetentionPeriod() string`

GetMetricsRetentionPeriod returns the MetricsRetentionPeriod field if non-nil, zero value otherwise.

### GetMetricsRetentionPeriodOk

`func (o *UpdateTenantPayload2) GetMetricsRetentionPeriodOk() (*string, bool)`

GetMetricsRetentionPeriodOk returns a tuple with the MetricsRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionPeriod

`func (o *UpdateTenantPayload2) SetMetricsRetentionPeriod(v string)`

SetMetricsRetentionPeriod sets MetricsRetentionPeriod field to given value.

### HasMetricsRetentionPeriod

`func (o *UpdateTenantPayload2) HasMetricsRetentionPeriod() bool`

HasMetricsRetentionPeriod returns a boolean if a field has been set.

### GetTraceRetentionPeriod

`func (o *UpdateTenantPayload2) GetTraceRetentionPeriod() string`

GetTraceRetentionPeriod returns the TraceRetentionPeriod field if non-nil, zero value otherwise.

### GetTraceRetentionPeriodOk

`func (o *UpdateTenantPayload2) GetTraceRetentionPeriodOk() (*string, bool)`

GetTraceRetentionPeriodOk returns a tuple with the TraceRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRetentionPeriod

`func (o *UpdateTenantPayload2) SetTraceRetentionPeriod(v string)`

SetTraceRetentionPeriod sets TraceRetentionPeriod field to given value.

### HasTraceRetentionPeriod

`func (o *UpdateTenantPayload2) HasTraceRetentionPeriod() bool`

HasTraceRetentionPeriod returns a boolean if a field has been set.

### GetTracing

`func (o *UpdateTenantPayload2) GetTracing() bool`

GetTracing returns the Tracing field if non-nil, zero value otherwise.

### GetTracingOk

`func (o *UpdateTenantPayload2) GetTracingOk() (*bool, bool)`

GetTracingOk returns a tuple with the Tracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracing

`func (o *UpdateTenantPayload2) SetTracing(v bool)`

SetTracing sets Tracing field to given value.

### HasTracing

`func (o *UpdateTenantPayload2) HasTracing() bool`

HasTracing returns a boolean if a field has been set.

### GetViewerClaim

`func (o *UpdateTenantPayload2) GetViewerClaim() string`

GetViewerClaim returns the ViewerClaim field if non-nil, zero value otherwise.

### GetViewerClaimOk

`func (o *UpdateTenantPayload2) GetViewerClaimOk() (*string, bool)`

GetViewerClaimOk returns a tuple with the ViewerClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerClaim

`func (o *UpdateTenantPayload2) SetViewerClaim(v string)`

SetViewerClaim sets ViewerClaim field to given value.

### HasViewerClaim

`func (o *UpdateTenantPayload2) HasViewerClaim() bool`

HasViewerClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


