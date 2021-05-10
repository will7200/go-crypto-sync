/*
 * CoinBase API
 *
 * Coinbase provides a simple and powerful REST API to integrate bitcoin, bitcoin cash, litecoin and ethereum payments into your business or application.  This API reference provides information on available endpoints and how to interact with it. To read more about the API, visit our API documentation.
 *
 * API version: 2019-11-15
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coinbase

import (
	"encoding/json"
)

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	Data *ExchangeRates `json:"data,omitempty"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse2001) GetData() ExchangeRates {
	if o == nil || o.Data == nil {
		var ret ExchangeRates
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetDataOk() (*ExchangeRates, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse2001) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ExchangeRates and assigns it to the Data field.
func (o *InlineResponse2001) SetData(v ExchangeRates) {
	o.Data = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
