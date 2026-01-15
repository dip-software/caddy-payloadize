# InputBillingEvent2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Count | 
**Service** | **string** | Service | 
**Source** | Pointer to **string** | Source of the event | [optional] 
**Timestamp** | **string** | Timestamp | 
**Unit** | Pointer to **string** | Unit of value (B, MB, GB). Default is &#x60;B&#x60; | [optional] 
**Value** | **int64** | Value | 

## Methods

### NewInputBillingEvent2

`func NewInputBillingEvent2(count int64, service string, timestamp string, value int64, ) *InputBillingEvent2`

NewInputBillingEvent2 instantiates a new InputBillingEvent2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputBillingEvent2WithDefaults

`func NewInputBillingEvent2WithDefaults() *InputBillingEvent2`

NewInputBillingEvent2WithDefaults instantiates a new InputBillingEvent2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *InputBillingEvent2) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InputBillingEvent2) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InputBillingEvent2) SetCount(v int64)`

SetCount sets Count field to given value.


### GetService

`func (o *InputBillingEvent2) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *InputBillingEvent2) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *InputBillingEvent2) SetService(v string)`

SetService sets Service field to given value.


### GetSource

`func (o *InputBillingEvent2) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputBillingEvent2) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputBillingEvent2) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InputBillingEvent2) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTimestamp

`func (o *InputBillingEvent2) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InputBillingEvent2) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InputBillingEvent2) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetUnit

`func (o *InputBillingEvent2) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InputBillingEvent2) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InputBillingEvent2) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *InputBillingEvent2) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *InputBillingEvent2) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputBillingEvent2) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputBillingEvent2) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


