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
// InlineResponse20011 struct for InlineResponse20011
type InlineResponse20011 struct {
	Rows  *[]InlineResponse20011Rows `json:"rows,omitempty"`
	Total *int32                     `json:"total,omitempty"`
}

// NewInlineResponse20011 instantiates a new InlineResponse20011 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011() *InlineResponse20011 {
	this := InlineResponse20011{}
	return &this
}

// NewInlineResponse20011WithDefaults instantiates a new InlineResponse20011 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011WithDefaults() *InlineResponse20011 {
	this := InlineResponse20011{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *InlineResponse20011) GetRows() []InlineResponse20011Rows {
	if o == nil || o.Rows == nil {
		var ret []InlineResponse20011Rows
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetRowsOk() (*[]InlineResponse20011Rows, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *InlineResponse20011) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given []InlineResponse20011Rows and assigns it to the Rows field.
func (o *InlineResponse20011) SetRows(v []InlineResponse20011Rows) {
	o.Rows = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse20011) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse20011) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse20011) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse20011) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011 struct {
	value *InlineResponse20011
	isSet bool
}

func (v NullableInlineResponse20011) Get() *InlineResponse20011 {
	return v.value
}

func (v *NullableInlineResponse20011) Set(val *InlineResponse20011) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011(val *InlineResponse20011) *NullableInlineResponse20011 {
	return &NullableInlineResponse20011{value: val, isSet: true}
}

func (v NullableInlineResponse20011) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
