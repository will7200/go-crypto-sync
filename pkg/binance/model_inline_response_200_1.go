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
// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	Timezone        *string                         `json:"timezone,omitempty"`
	ServerTime      *int32                          `json:"serverTime,omitempty"`
	RateLimits      *[]InlineResponse2001RateLimits `json:"rateLimits,omitempty"`
	ExchangeFilters *[]map[string]interface{}       `json:"exchangeFilters,omitempty"`
	Symbols         *[]InlineResponse2001Symbols    `json:"symbols,omitempty"`
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

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *InlineResponse2001) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *InlineResponse2001) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *InlineResponse2001) SetTimezone(v string) {
	o.Timezone = &v
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *InlineResponse2001) GetServerTime() int32 {
	if o == nil || o.ServerTime == nil {
		var ret int32
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetServerTimeOk() (*int32, bool) {
	if o == nil || o.ServerTime == nil {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *InlineResponse2001) HasServerTime() bool {
	if o != nil && o.ServerTime != nil {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int32 and assigns it to the ServerTime field.
func (o *InlineResponse2001) SetServerTime(v int32) {
	o.ServerTime = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *InlineResponse2001) GetRateLimits() []InlineResponse2001RateLimits {
	if o == nil || o.RateLimits == nil {
		var ret []InlineResponse2001RateLimits
		return ret
	}
	return *o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetRateLimitsOk() (*[]InlineResponse2001RateLimits, bool) {
	if o == nil || o.RateLimits == nil {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *InlineResponse2001) HasRateLimits() bool {
	if o != nil && o.RateLimits != nil {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []InlineResponse2001RateLimits and assigns it to the RateLimits field.
func (o *InlineResponse2001) SetRateLimits(v []InlineResponse2001RateLimits) {
	o.RateLimits = &v
}

// GetExchangeFilters returns the ExchangeFilters field value if set, zero value otherwise.
func (o *InlineResponse2001) GetExchangeFilters() []map[string]interface{} {
	if o == nil || o.ExchangeFilters == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.ExchangeFilters
}

// GetExchangeFiltersOk returns a tuple with the ExchangeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetExchangeFiltersOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ExchangeFilters == nil {
		return nil, false
	}
	return o.ExchangeFilters, true
}

// HasExchangeFilters returns a boolean if a field has been set.
func (o *InlineResponse2001) HasExchangeFilters() bool {
	if o != nil && o.ExchangeFilters != nil {
		return true
	}

	return false
}

// SetExchangeFilters gets a reference to the given []map[string]interface{} and assigns it to the ExchangeFilters field.
func (o *InlineResponse2001) SetExchangeFilters(v []map[string]interface{}) {
	o.ExchangeFilters = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *InlineResponse2001) GetSymbols() []InlineResponse2001Symbols {
	if o == nil || o.Symbols == nil {
		var ret []InlineResponse2001Symbols
		return ret
	}
	return *o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetSymbolsOk() (*[]InlineResponse2001Symbols, bool) {
	if o == nil || o.Symbols == nil {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *InlineResponse2001) HasSymbols() bool {
	if o != nil && o.Symbols != nil {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []InlineResponse2001Symbols and assigns it to the Symbols field.
func (o *InlineResponse2001) SetSymbols(v []InlineResponse2001Symbols) {
	o.Symbols = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.ServerTime != nil {
		toSerialize["serverTime"] = o.ServerTime
	}
	if o.RateLimits != nil {
		toSerialize["rateLimits"] = o.RateLimits
	}
	if o.ExchangeFilters != nil {
		toSerialize["exchangeFilters"] = o.ExchangeFilters
	}
	if o.Symbols != nil {
		toSerialize["symbols"] = o.Symbols
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
