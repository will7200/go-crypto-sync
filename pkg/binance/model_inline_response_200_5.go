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
// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	Rows  *[]InlineResponse2005Rows `json:"rows,omitempty"`
	Total *int32                    `json:"total,omitempty"`
}

// NewInlineResponse2005 instantiates a new InlineResponse2005 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2005() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// NewInlineResponse2005WithDefaults instantiates a new InlineResponse2005 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2005WithDefaults() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *InlineResponse2005) GetRows() []InlineResponse2005Rows {
	if o == nil || o.Rows == nil {
		var ret []InlineResponse2005Rows
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005) GetRowsOk() (*[]InlineResponse2005Rows, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *InlineResponse2005) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given []InlineResponse2005Rows and assigns it to the Rows field.
func (o *InlineResponse2005) SetRows(v []InlineResponse2005Rows) {
	o.Rows = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse2005) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse2005) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse2005) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse2005) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2005 struct {
	value *InlineResponse2005
	isSet bool
}

func (v NullableInlineResponse2005) Get() *InlineResponse2005 {
	return v.value
}

func (v *NullableInlineResponse2005) Set(val *InlineResponse2005) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2005) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2005) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2005(val *InlineResponse2005) *NullableInlineResponse2005 {
	return &NullableInlineResponse2005{value: val, isSet: true}
}

func (v NullableInlineResponse2005) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2005) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
