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
// InlineResponse20025TradeFee struct for InlineResponse20025TradeFee
type InlineResponse20025TradeFee struct {
	Symbol *string `json:"symbol,omitempty"`
	Maker  *int32  `json:"maker,omitempty"`
	Taker  *int32  `json:"taker,omitempty"`
}

// NewInlineResponse20025TradeFee instantiates a new InlineResponse20025TradeFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20025TradeFee() *InlineResponse20025TradeFee {
	this := InlineResponse20025TradeFee{}
	return &this
}

// NewInlineResponse20025TradeFeeWithDefaults instantiates a new InlineResponse20025TradeFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20025TradeFeeWithDefaults() *InlineResponse20025TradeFee {
	this := InlineResponse20025TradeFee{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *InlineResponse20025TradeFee) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025TradeFee) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *InlineResponse20025TradeFee) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *InlineResponse20025TradeFee) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *InlineResponse20025TradeFee) GetMaker() int32 {
	if o == nil || o.Maker == nil {
		var ret int32
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025TradeFee) GetMakerOk() (*int32, bool) {
	if o == nil || o.Maker == nil {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *InlineResponse20025TradeFee) HasMaker() bool {
	if o != nil && o.Maker != nil {
		return true
	}

	return false
}

// SetMaker gets a reference to the given int32 and assigns it to the Maker field.
func (o *InlineResponse20025TradeFee) SetMaker(v int32) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *InlineResponse20025TradeFee) GetTaker() int32 {
	if o == nil || o.Taker == nil {
		var ret int32
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025TradeFee) GetTakerOk() (*int32, bool) {
	if o == nil || o.Taker == nil {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *InlineResponse20025TradeFee) HasTaker() bool {
	if o != nil && o.Taker != nil {
		return true
	}

	return false
}

// SetTaker gets a reference to the given int32 and assigns it to the Taker field.
func (o *InlineResponse20025TradeFee) SetTaker(v int32) {
	o.Taker = &v
}

func (o InlineResponse20025TradeFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Maker != nil {
		toSerialize["maker"] = o.Maker
	}
	if o.Taker != nil {
		toSerialize["taker"] = o.Taker
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20025TradeFee struct {
	value *InlineResponse20025TradeFee
	isSet bool
}

func (v NullableInlineResponse20025TradeFee) Get() *InlineResponse20025TradeFee {
	return v.value
}

func (v *NullableInlineResponse20025TradeFee) Set(val *InlineResponse20025TradeFee) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20025TradeFee) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20025TradeFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20025TradeFee(val *InlineResponse20025TradeFee) *NullableInlineResponse20025TradeFee {
	return &NullableInlineResponse20025TradeFee{value: val, isSet: true}
}

func (v NullableInlineResponse20025TradeFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20025TradeFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
