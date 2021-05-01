# InlineResponse20020Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLocked** | Pointer to **bool** |  | [optional] 
**PlannedRecoverTime** | Pointer to **int32** |  | [optional] 
**TriggerCondition** | Pointer to [**InlineResponse20020StatusTriggerCondition**](InlineResponse20020StatusTriggerCondition.md) |  | [optional] 
**Indicators** | Pointer to [**InlineResponse20020StatusIndicators**](InlineResponse20020StatusIndicators.md) |  | [optional] 
**UpdateTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewInlineResponse20020Status

`func NewInlineResponse20020Status() *InlineResponse20020Status`

NewInlineResponse20020Status instantiates a new InlineResponse20020Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20020StatusWithDefaults

`func NewInlineResponse20020StatusWithDefaults() *InlineResponse20020Status`

NewInlineResponse20020StatusWithDefaults instantiates a new InlineResponse20020Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLocked

`func (o *InlineResponse20020Status) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *InlineResponse20020Status) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *InlineResponse20020Status) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *InlineResponse20020Status) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetPlannedRecoverTime

`func (o *InlineResponse20020Status) GetPlannedRecoverTime() int32`

GetPlannedRecoverTime returns the PlannedRecoverTime field if non-nil, zero value otherwise.

### GetPlannedRecoverTimeOk

`func (o *InlineResponse20020Status) GetPlannedRecoverTimeOk() (*int32, bool)`

GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRecoverTime

`func (o *InlineResponse20020Status) SetPlannedRecoverTime(v int32)`

SetPlannedRecoverTime sets PlannedRecoverTime field to given value.

### HasPlannedRecoverTime

`func (o *InlineResponse20020Status) HasPlannedRecoverTime() bool`

HasPlannedRecoverTime returns a boolean if a field has been set.

### GetTriggerCondition

`func (o *InlineResponse20020Status) GetTriggerCondition() InlineResponse20020StatusTriggerCondition`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *InlineResponse20020Status) GetTriggerConditionOk() (*InlineResponse20020StatusTriggerCondition, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *InlineResponse20020Status) SetTriggerCondition(v InlineResponse20020StatusTriggerCondition)`

SetTriggerCondition sets TriggerCondition field to given value.

### HasTriggerCondition

`func (o *InlineResponse20020Status) HasTriggerCondition() bool`

HasTriggerCondition returns a boolean if a field has been set.

### GetIndicators

`func (o *InlineResponse20020Status) GetIndicators() InlineResponse20020StatusIndicators`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *InlineResponse20020Status) GetIndicatorsOk() (*InlineResponse20020StatusIndicators, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *InlineResponse20020Status) SetIndicators(v InlineResponse20020StatusIndicators)`

SetIndicators sets Indicators field to given value.

### HasIndicators

`func (o *InlineResponse20020Status) HasIndicators() bool`

HasIndicators returns a boolean if a field has been set.

### GetUpdateTime

`func (o *InlineResponse20020Status) GetUpdateTime() int32`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *InlineResponse20020Status) GetUpdateTimeOk() (*int32, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *InlineResponse20020Status) SetUpdateTime(v int32)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *InlineResponse20020Status) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


