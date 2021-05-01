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
// InlineResponse20014Data struct for InlineResponse20014Data
type InlineResponse20014Data struct {
	Balances        *[]InlineResponse20014DataBalances `json:"balances,omitempty"`
	TotalAssetOfBtc *string                            `json:"totalAssetOfBtc,omitempty"`
}

// NewInlineResponse20014Data instantiates a new InlineResponse20014Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014Data() *InlineResponse20014Data {
	this := InlineResponse20014Data{}
	return &this
}

// NewInlineResponse20014DataWithDefaults instantiates a new InlineResponse20014Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014DataWithDefaults() *InlineResponse20014Data {
	this := InlineResponse20014Data{}
	return &this
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *InlineResponse20014Data) GetBalances() []InlineResponse20014DataBalances {
	if o == nil || o.Balances == nil {
		var ret []InlineResponse20014DataBalances
		return ret
	}
	return *o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014Data) GetBalancesOk() (*[]InlineResponse20014DataBalances, bool) {
	if o == nil || o.Balances == nil {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *InlineResponse20014Data) HasBalances() bool {
	if o != nil && o.Balances != nil {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []InlineResponse20014DataBalances and assigns it to the Balances field.
func (o *InlineResponse20014Data) SetBalances(v []InlineResponse20014DataBalances) {
	o.Balances = &v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *InlineResponse20014Data) GetTotalAssetOfBtc() string {
	if o == nil || o.TotalAssetOfBtc == nil {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014Data) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || o.TotalAssetOfBtc == nil {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *InlineResponse20014Data) HasTotalAssetOfBtc() bool {
	if o != nil && o.TotalAssetOfBtc != nil {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *InlineResponse20014Data) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

func (o InlineResponse20014Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Balances != nil {
		toSerialize["balances"] = o.Balances
	}
	if o.TotalAssetOfBtc != nil {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20014Data struct {
	value *InlineResponse20014Data
	isSet bool
}

func (v NullableInlineResponse20014Data) Get() *InlineResponse20014Data {
	return v.value
}

func (v *NullableInlineResponse20014Data) Set(val *InlineResponse20014Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20014Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20014Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20014Data(val *InlineResponse20014Data) *NullableInlineResponse20014Data {
	return &NullableInlineResponse20014Data{value: val, isSet: true}
}

func (v NullableInlineResponse20014Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20014Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}