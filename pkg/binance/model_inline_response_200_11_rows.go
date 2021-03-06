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
// InlineResponse20011Rows struct for InlineResponse20011Rows
type InlineResponse20011Rows struct {
	Asset               *string `json:"asset,omitempty"`
	Interest            *string `json:"interest,omitempty"`
	InterestAccuredTime *int32  `json:"interestAccuredTime,omitempty"`
	InterestRate        *string `json:"interestRate,omitempty"`
	Principal           *string `json:"principal,omitempty"`
	Type                *string `json:"type,omitempty"`
}

// NewInlineResponse20011Rows instantiates a new InlineResponse20011Rows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011Rows() *InlineResponse20011Rows {
	this := InlineResponse20011Rows{}
	return &this
}

// NewInlineResponse20011RowsWithDefaults instantiates a new InlineResponse20011Rows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011RowsWithDefaults() *InlineResponse20011Rows {
	this := InlineResponse20011Rows{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *InlineResponse20011Rows) GetAsset() string {
	if o == nil || o.Asset == nil {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Rows) GetAssetOk() (*string, bool) {
	if o == nil || o.Asset == nil {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *InlineResponse20011Rows) HasAsset() bool {
	if o != nil && o.Asset != nil {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *InlineResponse20011Rows) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *InlineResponse20011Rows) GetInterest() string {
	if o == nil || o.Interest == nil {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Rows) GetInterestOk() (*string, bool) {
	if o == nil || o.Interest == nil {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *InlineResponse20011Rows) HasInterest() bool {
	if o != nil && o.Interest != nil {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *InlineResponse20011Rows) SetInterest(v string) {
	o.Interest = &v
}

// GetInterestAccuredTime returns the InterestAccuredTime field value if set, zero value otherwise.
func (o *InlineResponse20011Rows) GetInterestAccuredTime() int32 {
	if o == nil || o.InterestAccuredTime == nil {
		var ret int32
		return ret
	}
	return *o.InterestAccuredTime
}

// GetInterestAccuredTimeOk returns a tuple with the InterestAccuredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Rows) GetInterestAccuredTimeOk() (*int32, bool) {
	if o == nil || o.InterestAccuredTime == nil {
		return nil, false
	}
	return o.InterestAccuredTime, true
}

// HasInterestAccuredTime returns a boolean if a field has been set.
func (o *InlineResponse20011Rows) HasInterestAccuredTime() bool {
	if o != nil && o.InterestAccuredTime != nil {
		return true
	}

	return false
}

// SetInterestAccuredTime gets a reference to the given int32 and assigns it to the InterestAccuredTime field.
func (o *InlineResponse20011Rows) SetInterestAccuredTime(v int32) {
	o.InterestAccuredTime = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *InlineResponse20011Rows) GetInterestRate() string {
	if o == nil || o.InterestRate == nil {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Rows) GetInterestRateOk() (*string, bool) {
	if o == nil || o.InterestRate == nil {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *InlineResponse20011Rows) HasInterestRate() bool {
	if o != nil && o.InterestRate != nil {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *InlineResponse20011Rows) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *InlineResponse20011Rows) GetPrincipal() string {
	if o == nil || o.Principal == nil {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Rows) GetPrincipalOk() (*string, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *InlineResponse20011Rows) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *InlineResponse20011Rows) SetPrincipal(v string) {
	o.Principal = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20011Rows) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Rows) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20011Rows) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20011Rows) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse20011Rows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Asset != nil {
		toSerialize["asset"] = o.Asset
	}
	if o.Interest != nil {
		toSerialize["interest"] = o.Interest
	}
	if o.InterestAccuredTime != nil {
		toSerialize["interestAccuredTime"] = o.InterestAccuredTime
	}
	if o.InterestRate != nil {
		toSerialize["interestRate"] = o.InterestRate
	}
	if o.Principal != nil {
		toSerialize["principal"] = o.Principal
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011Rows struct {
	value *InlineResponse20011Rows
	isSet bool
}

func (v NullableInlineResponse20011Rows) Get() *InlineResponse20011Rows {
	return v.value
}

func (v *NullableInlineResponse20011Rows) Set(val *InlineResponse20011Rows) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011Rows) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011Rows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011Rows(val *InlineResponse20011Rows) *NullableInlineResponse20011Rows {
	return &NullableInlineResponse20011Rows{value: val, isSet: true}
}

func (v NullableInlineResponse20011Rows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011Rows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
