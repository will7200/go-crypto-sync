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
// InlineResponse20020StatusIndicatorsBTCUSDT struct for InlineResponse20020StatusIndicatorsBTCUSDT
type InlineResponse20020StatusIndicatorsBTCUSDT struct {
	I *string  `json:"i,omitempty"`
	C *int32   `json:"c,omitempty"`
	V *float32 `json:"v,omitempty"`
	T *float32 `json:"t,omitempty"`
}

// NewInlineResponse20020StatusIndicatorsBTCUSDT instantiates a new InlineResponse20020StatusIndicatorsBTCUSDT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20020StatusIndicatorsBTCUSDT() *InlineResponse20020StatusIndicatorsBTCUSDT {
	this := InlineResponse20020StatusIndicatorsBTCUSDT{}
	return &this
}

// NewInlineResponse20020StatusIndicatorsBTCUSDTWithDefaults instantiates a new InlineResponse20020StatusIndicatorsBTCUSDT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20020StatusIndicatorsBTCUSDTWithDefaults() *InlineResponse20020StatusIndicatorsBTCUSDT {
	this := InlineResponse20020StatusIndicatorsBTCUSDT{}
	return &this
}

// GetI returns the I field value if set, zero value otherwise.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) GetI() string {
	if o == nil || o.I == nil {
		var ret string
		return ret
	}
	return *o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) GetIOk() (*string, bool) {
	if o == nil || o.I == nil {
		return nil, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) HasI() bool {
	if o != nil && o.I != nil {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) SetI(v string) {
	o.I = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) GetC() int32 {
	if o == nil || o.C == nil {
		var ret int32
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) GetCOk() (*int32, bool) {
	if o == nil || o.C == nil {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) HasC() bool {
	if o != nil && o.C != nil {
		return true
	}

	return false
}

// SetC gets a reference to the given int32 and assigns it to the C field.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) SetC(v int32) {
	o.C = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) GetV() float32 {
	if o == nil || o.V == nil {
		var ret float32
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) GetVOk() (*float32, bool) {
	if o == nil || o.V == nil {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) HasV() bool {
	if o != nil && o.V != nil {
		return true
	}

	return false
}

// SetV gets a reference to the given float32 and assigns it to the V field.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) SetV(v float32) {
	o.V = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) GetT() float32 {
	if o == nil || o.T == nil {
		var ret float32
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) GetTOk() (*float32, bool) {
	if o == nil || o.T == nil {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) HasT() bool {
	if o != nil && o.T != nil {
		return true
	}

	return false
}

// SetT gets a reference to the given float32 and assigns it to the T field.
func (o *InlineResponse20020StatusIndicatorsBTCUSDT) SetT(v float32) {
	o.T = &v
}

func (o InlineResponse20020StatusIndicatorsBTCUSDT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.I != nil {
		toSerialize["i"] = o.I
	}
	if o.C != nil {
		toSerialize["c"] = o.C
	}
	if o.V != nil {
		toSerialize["v"] = o.V
	}
	if o.T != nil {
		toSerialize["t"] = o.T
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20020StatusIndicatorsBTCUSDT struct {
	value *InlineResponse20020StatusIndicatorsBTCUSDT
	isSet bool
}

func (v NullableInlineResponse20020StatusIndicatorsBTCUSDT) Get() *InlineResponse20020StatusIndicatorsBTCUSDT {
	return v.value
}

func (v *NullableInlineResponse20020StatusIndicatorsBTCUSDT) Set(val *InlineResponse20020StatusIndicatorsBTCUSDT) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20020StatusIndicatorsBTCUSDT) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20020StatusIndicatorsBTCUSDT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20020StatusIndicatorsBTCUSDT(val *InlineResponse20020StatusIndicatorsBTCUSDT) *NullableInlineResponse20020StatusIndicatorsBTCUSDT {
	return &NullableInlineResponse20020StatusIndicatorsBTCUSDT{value: val, isSet: true}
}

func (v NullableInlineResponse20020StatusIndicatorsBTCUSDT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20020StatusIndicatorsBTCUSDT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
