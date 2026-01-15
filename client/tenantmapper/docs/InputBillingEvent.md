# InputBillingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Count | 
**Organization** | **string** | Organization, also known as tenant | 
**Service** | **string** | Service | 
**Source** | Pointer to **string** | Source of the event | [optional] 
**Timestamp** | **string** | Timestamp | 
**Unit** | Pointer to **string** | Unit of value (B, MB, GB). Default is &#x60;B&#x60; | [optional] 
**Value** | **int64** | Value | 

## Methods

### NewInputBillingEvent

`func NewInputBillingEvent(count int64, organization string, service string, timestamp string, value int64, ) *InputBillingEvent`

NewInputBillingEvent instantiates a new InputBillingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputBillingEventWithDefaults

`func NewInputBillingEventWithDefaults() *InputBillingEvent`

NewInputBillingEventWithDefaults instantiates a new InputBillingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *InputBillingEvent) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InputBillingEvent) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InputBillingEvent) SetCount(v int64)`

SetCount sets Count field to given value.


### GetOrganization

`func (o *InputBillingEvent) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InputBillingEvent) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InputBillingEvent) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetService

`func (o *InputBillingEvent) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *InputBillingEvent) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *InputBillingEvent) SetService(v string)`

SetService sets Service field to given value.


### GetSource

`func (o *InputBillingEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputBillingEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputBillingEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InputBillingEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTimestamp

`func (o *InputBillingEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InputBillingEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InputBillingEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetUnit

`func (o *InputBillingEvent) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InputBillingEvent) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InputBillingEvent) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *InputBillingEvent) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *InputBillingEvent) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputBillingEvent) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputBillingEvent) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


