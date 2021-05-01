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
// InlineResponse20014SnapshotVos struct for InlineResponse20014SnapshotVos
type InlineResponse20014SnapshotVos struct {
	Data       *InlineResponse20014Data `json:"data,omitempty"`
	Type       *string                  `json:"type,omitempty"`
	UpdateTime *int32                   `json:"updateTime,omitempty"`
}

// NewInlineResponse20014SnapshotVos instantiates a new InlineResponse20014SnapshotVos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014SnapshotVos() *InlineResponse20014SnapshotVos {
	this := InlineResponse20014SnapshotVos{}
	return &this
}

// NewInlineResponse20014SnapshotVosWithDefaults instantiates a new InlineResponse20014SnapshotVos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014SnapshotVosWithDefaults() *InlineResponse20014SnapshotVos {
	this := InlineResponse20014SnapshotVos{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse20014SnapshotVos) GetData() InlineResponse20014Data {
	if o == nil || o.Data == nil {
		var ret InlineResponse20014Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014SnapshotVos) GetDataOk() (*InlineResponse20014Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse20014SnapshotVos) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given InlineResponse20014Data and assigns it to the Data field.
func (o *InlineResponse20014SnapshotVos) SetData(v InlineResponse20014Data) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20014SnapshotVos) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014SnapshotVos) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20014SnapshotVos) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20014SnapshotVos) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *InlineResponse20014SnapshotVos) GetUpdateTime() int32 {
	if o == nil || o.UpdateTime == nil {
		var ret int32
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014SnapshotVos) GetUpdateTimeOk() (*int32, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *InlineResponse20014SnapshotVos) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int32 and assigns it to the UpdateTime field.
func (o *InlineResponse20014SnapshotVos) SetUpdateTime(v int32) {
	o.UpdateTime = &v
}

func (o InlineResponse20014SnapshotVos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20014SnapshotVos struct {
	value *InlineResponse20014SnapshotVos
	isSet bool
}

func (v NullableInlineResponse20014SnapshotVos) Get() *InlineResponse20014SnapshotVos {
	return v.value
}

func (v *NullableInlineResponse20014SnapshotVos) Set(val *InlineResponse20014SnapshotVos) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20014SnapshotVos) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20014SnapshotVos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20014SnapshotVos(val *InlineResponse20014SnapshotVos) *NullableInlineResponse20014SnapshotVos {
	return &NullableInlineResponse20014SnapshotVos{value: val, isSet: true}
}

func (v NullableInlineResponse20014SnapshotVos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20014SnapshotVos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
