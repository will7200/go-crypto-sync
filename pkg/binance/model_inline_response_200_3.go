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
// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	Mins  *int32  `json:"mins,omitempty"`
	Price *string `json:"price,omitempty"`
}

// NewInlineResponse2003 instantiates a new InlineResponse2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003WithDefaults() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// GetMins returns the Mins field value if set, zero value otherwise.
func (o *InlineResponse2003) GetMins() int32 {
	if o == nil || o.Mins == nil {
		var ret int32
		return ret
	}
	return *o.Mins
}

// GetMinsOk returns a tuple with the Mins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetMinsOk() (*int32, bool) {
	if o == nil || o.Mins == nil {
		return nil, false
	}
	return o.Mins, true
}

// HasMins returns a boolean if a field has been set.
func (o *InlineResponse2003) HasMins() bool {
	if o != nil && o.Mins != nil {
		return true
	}

	return false
}

// SetMins gets a reference to the given int32 and assigns it to the Mins field.
func (o *InlineResponse2003) SetMins(v int32) {
	o.Mins = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InlineResponse2003) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InlineResponse2003) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *InlineResponse2003) SetPrice(v string) {
	o.Price = &v
}

func (o InlineResponse2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mins != nil {
		toSerialize["mins"] = o.Mins
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003 struct {
	value *InlineResponse2003
	isSet bool
}

func (v NullableInlineResponse2003) Get() *InlineResponse2003 {
	return v.value
}

func (v *NullableInlineResponse2003) Set(val *InlineResponse2003) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003(val *InlineResponse2003) *NullableInlineResponse2003 {
	return &NullableInlineResponse2003{value: val, isSet: true}
}

func (v NullableInlineResponse2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}