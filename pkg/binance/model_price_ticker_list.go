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
// PriceTickerList struct for PriceTickerList
type PriceTickerList struct {
	Items []PriceTicker
}

// NewPriceTickerList instantiates a new PriceTickerList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceTickerList() *PriceTickerList {
	this := PriceTickerList{}
	return &this
}

// NewPriceTickerListWithDefaults instantiates a new PriceTickerList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceTickerListWithDefaults() *PriceTickerList {
	this := PriceTickerList{}
	return &this
}

func (o PriceTickerList) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return json.Marshal(toSerialize)
}

func (o *PriceTickerList) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &o.Items)
}

type NullablePriceTickerList struct {
	value *PriceTickerList
	isSet bool
}

func (v NullablePriceTickerList) Get() *PriceTickerList {
	return v.value
}

func (v *NullablePriceTickerList) Set(val *PriceTickerList) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTickerList) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTickerList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTickerList(val *PriceTickerList) *NullablePriceTickerList {
	return &NullablePriceTickerList{value: val, isSet: true}
}

func (v NullablePriceTickerList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTickerList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}