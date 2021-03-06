/*
 * Binance SPOT Public API
 *
 * The swagger file of Binance Public API  API documents:   - [https://github.com/binance/binance-spot-api-docs](https://github.com/binance/binance-spot-api-docs)   - [https://binance-docs.github.io/apidocs/spot/en](https://binance-docs.github.io/apidocs/spot/en)
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binance

import (
	"encoding/json"
)

//
// InlineResponse20020Status struct for InlineResponse20020Status
type InlineResponse20020Status struct {
	IsLocked           *bool                                      `json:"isLocked,omitempty"`
	PlannedRecoverTime *int32                                     `json:"plannedRecoverTime,omitempty"`
	TriggerCondition   *InlineResponse20020StatusTriggerCondition `json:"triggerCondition,omitempty"`
	Indicators         *InlineResponse20020StatusIndicators       `json:"indicators,omitempty"`
	UpdateTime         *int32                                     `json:"updateTime,omitempty"`
}

// NewInlineResponse20020Status instantiates a new InlineResponse20020Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20020Status() *InlineResponse20020Status {
	this := InlineResponse20020Status{}
	return &this
}

// NewInlineResponse20020StatusWithDefaults instantiates a new InlineResponse20020Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20020StatusWithDefaults() *InlineResponse20020Status {
	this := InlineResponse20020Status{}
	return &this
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *InlineResponse20020Status) GetIsLocked() bool {
	if o == nil || o.IsLocked == nil {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Status) GetIsLockedOk() (*bool, bool) {
	if o == nil || o.IsLocked == nil {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *InlineResponse20020Status) HasIsLocked() bool {
	if o != nil && o.IsLocked != nil {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *InlineResponse20020Status) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetPlannedRecoverTime returns the PlannedRecoverTime field value if set, zero value otherwise.
func (o *InlineResponse20020Status) GetPlannedRecoverTime() int32 {
	if o == nil || o.PlannedRecoverTime == nil {
		var ret int32
		return ret
	}
	return *o.PlannedRecoverTime
}

// GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Status) GetPlannedRecoverTimeOk() (*int32, bool) {
	if o == nil || o.PlannedRecoverTime == nil {
		return nil, false
	}
	return o.PlannedRecoverTime, true
}

// HasPlannedRecoverTime returns a boolean if a field has been set.
func (o *InlineResponse20020Status) HasPlannedRecoverTime() bool {
	if o != nil && o.PlannedRecoverTime != nil {
		return true
	}

	return false
}

// SetPlannedRecoverTime gets a reference to the given int32 and assigns it to the PlannedRecoverTime field.
func (o *InlineResponse20020Status) SetPlannedRecoverTime(v int32) {
	o.PlannedRecoverTime = &v
}

// GetTriggerCondition returns the TriggerCondition field value if set, zero value otherwise.
func (o *InlineResponse20020Status) GetTriggerCondition() InlineResponse20020StatusTriggerCondition {
	if o == nil || o.TriggerCondition == nil {
		var ret InlineResponse20020StatusTriggerCondition
		return ret
	}
	return *o.TriggerCondition
}

// GetTriggerConditionOk returns a tuple with the TriggerCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Status) GetTriggerConditionOk() (*InlineResponse20020StatusTriggerCondition, bool) {
	if o == nil || o.TriggerCondition == nil {
		return nil, false
	}
	return o.TriggerCondition, true
}

// HasTriggerCondition returns a boolean if a field has been set.
func (o *InlineResponse20020Status) HasTriggerCondition() bool {
	if o != nil && o.TriggerCondition != nil {
		return true
	}

	return false
}

// SetTriggerCondition gets a reference to the given InlineResponse20020StatusTriggerCondition and assigns it to the TriggerCondition field.
func (o *InlineResponse20020Status) SetTriggerCondition(v InlineResponse20020StatusTriggerCondition) {
	o.TriggerCondition = &v
}

// GetIndicators returns the Indicators field value if set, zero value otherwise.
func (o *InlineResponse20020Status) GetIndicators() InlineResponse20020StatusIndicators {
	if o == nil || o.Indicators == nil {
		var ret InlineResponse20020StatusIndicators
		return ret
	}
	return *o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Status) GetIndicatorsOk() (*InlineResponse20020StatusIndicators, bool) {
	if o == nil || o.Indicators == nil {
		return nil, false
	}
	return o.Indicators, true
}

// HasIndicators returns a boolean if a field has been set.
func (o *InlineResponse20020Status) HasIndicators() bool {
	if o != nil && o.Indicators != nil {
		return true
	}

	return false
}

// SetIndicators gets a reference to the given InlineResponse20020StatusIndicators and assigns it to the Indicators field.
func (o *InlineResponse20020Status) SetIndicators(v InlineResponse20020StatusIndicators) {
	o.Indicators = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *InlineResponse20020Status) GetUpdateTime() int32 {
	if o == nil || o.UpdateTime == nil {
		var ret int32
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Status) GetUpdateTimeOk() (*int32, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *InlineResponse20020Status) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int32 and assigns it to the UpdateTime field.
func (o *InlineResponse20020Status) SetUpdateTime(v int32) {
	o.UpdateTime = &v
}

func (o InlineResponse20020Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsLocked != nil {
		toSerialize["isLocked"] = o.IsLocked
	}
	if o.PlannedRecoverTime != nil {
		toSerialize["plannedRecoverTime"] = o.PlannedRecoverTime
	}
	if o.TriggerCondition != nil {
		toSerialize["triggerCondition"] = o.TriggerCondition
	}
	if o.Indicators != nil {
		toSerialize["indicators"] = o.Indicators
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20020Status struct {
	value *InlineResponse20020Status
	isSet bool
}

func (v NullableInlineResponse20020Status) Get() *InlineResponse20020Status {
	return v.value
}

func (v *NullableInlineResponse20020Status) Set(val *InlineResponse20020Status) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20020Status) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20020Status) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20020Status(val *InlineResponse20020Status) *NullableInlineResponse20020Status {
	return &NullableInlineResponse20020Status{value: val, isSet: true}
}

func (v NullableInlineResponse20020Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20020Status) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
