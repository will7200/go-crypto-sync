/*
 * CoinEx API
 *
 * Coinex doesn't have a openapi docs sadville  Open and simple, CoinEx API makes sure that you can build your own trading tools to achieve a more effective trading strategy. CoinEx API is now available for these features:
 *
 * API version: 2021-05-29
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coinex

import (
	"encoding/json"
)

// AcquireSingleMarketInformationResponse AcquireSingleMarketInformation Response Value
type AcquireSingleMarketInformationResponse struct {
	// market name
	Name *string `json:"name,omitempty"`
	// taker fee
	TakerFeeRate *string `json:"taker_fee_rate,omitempty"`
	// maker fee
	MakerFeeRate *string `json:"maker_fee_rate,omitempty"`
	// Min. amount of transaction
	MinAmount *string `json:"min_amount,omitempty"`
	// tradiing coin type
	TradingName *string `json:"trading_name,omitempty"`
	// decimal of tradiing coin
	TradingDecimal *int64 `json:"trading_decimal,omitempty"`
	// pricing coin type
	PricingName *string `json:"pricing_name,omitempty"`
	// decimal of pricing coin
	PricingDecimal *int64 `json:"pricing_decimal,omitempty"`
}

// NewAcquireSingleMarketInformationResponse instantiates a new AcquireSingleMarketInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcquireSingleMarketInformationResponse() *AcquireSingleMarketInformationResponse {
	this := AcquireSingleMarketInformationResponse{}
	return &this
}

// NewAcquireSingleMarketInformationResponseWithDefaults instantiates a new AcquireSingleMarketInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcquireSingleMarketInformationResponseWithDefaults() *AcquireSingleMarketInformationResponse {
	this := AcquireSingleMarketInformationResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AcquireSingleMarketInformationResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireSingleMarketInformationResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AcquireSingleMarketInformationResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AcquireSingleMarketInformationResponse) SetName(v string) {
	o.Name = &v
}

// GetTakerFeeRate returns the TakerFeeRate field value if set, zero value otherwise.
func (o *AcquireSingleMarketInformationResponse) GetTakerFeeRate() string {
	if o == nil || o.TakerFeeRate == nil {
		var ret string
		return ret
	}
	return *o.TakerFeeRate
}

// GetTakerFeeRateOk returns a tuple with the TakerFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireSingleMarketInformationResponse) GetTakerFeeRateOk() (*string, bool) {
	if o == nil || o.TakerFeeRate == nil {
		return nil, false
	}
	return o.TakerFeeRate, true
}

// HasTakerFeeRate returns a boolean if a field has been set.
func (o *AcquireSingleMarketInformationResponse) HasTakerFeeRate() bool {
	if o != nil && o.TakerFeeRate != nil {
		return true
	}

	return false
}

// SetTakerFeeRate gets a reference to the given string and assigns it to the TakerFeeRate field.
func (o *AcquireSingleMarketInformationResponse) SetTakerFeeRate(v string) {
	o.TakerFeeRate = &v
}

// GetMakerFeeRate returns the MakerFeeRate field value if set, zero value otherwise.
func (o *AcquireSingleMarketInformationResponse) GetMakerFeeRate() string {
	if o == nil || o.MakerFeeRate == nil {
		var ret string
		return ret
	}
	return *o.MakerFeeRate
}

// GetMakerFeeRateOk returns a tuple with the MakerFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireSingleMarketInformationResponse) GetMakerFeeRateOk() (*string, bool) {
	if o == nil || o.MakerFeeRate == nil {
		return nil, false
	}
	return o.MakerFeeRate, true
}

// HasMakerFeeRate returns a boolean if a field has been set.
func (o *AcquireSingleMarketInformationResponse) HasMakerFeeRate() bool {
	if o != nil && o.MakerFeeRate != nil {
		return true
	}

	return false
}

// SetMakerFeeRate gets a reference to the given string and assigns it to the MakerFeeRate field.
func (o *AcquireSingleMarketInformationResponse) SetMakerFeeRate(v string) {
	o.MakerFeeRate = &v
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise.
func (o *AcquireSingleMarketInformationResponse) GetMinAmount() string {
	if o == nil || o.MinAmount == nil {
		var ret string
		return ret
	}
	return *o.MinAmount
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireSingleMarketInformationResponse) GetMinAmountOk() (*string, bool) {
	if o == nil || o.MinAmount == nil {
		return nil, false
	}
	return o.MinAmount, true
}

// HasMinAmount returns a boolean if a field has been set.
func (o *AcquireSingleMarketInformationResponse) HasMinAmount() bool {
	if o != nil && o.MinAmount != nil {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given string and assigns it to the MinAmount field.
func (o *AcquireSingleMarketInformationResponse) SetMinAmount(v string) {
	o.MinAmount = &v
}

// GetTradingName returns the TradingName field value if set, zero value otherwise.
func (o *AcquireSingleMarketInformationResponse) GetTradingName() string {
	if o == nil || o.TradingName == nil {
		var ret string
		return ret
	}
	return *o.TradingName
}

// GetTradingNameOk returns a tuple with the TradingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireSingleMarketInformationResponse) GetTradingNameOk() (*string, bool) {
	if o == nil || o.TradingName == nil {
		return nil, false
	}
	return o.TradingName, true
}

// HasTradingName returns a boolean if a field has been set.
func (o *AcquireSingleMarketInformationResponse) HasTradingName() bool {
	if o != nil && o.TradingName != nil {
		return true
	}

	return false
}

// SetTradingName gets a reference to the given string and assigns it to the TradingName field.
func (o *AcquireSingleMarketInformationResponse) SetTradingName(v string) {
	o.TradingName = &v
}

// GetTradingDecimal returns the TradingDecimal field value if set, zero value otherwise.
func (o *AcquireSingleMarketInformationResponse) GetTradingDecimal() int64 {
	if o == nil || o.TradingDecimal == nil {
		var ret int64
		return ret
	}
	return *o.TradingDecimal
}

// GetTradingDecimalOk returns a tuple with the TradingDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireSingleMarketInformationResponse) GetTradingDecimalOk() (*int64, bool) {
	if o == nil || o.TradingDecimal == nil {
		return nil, false
	}
	return o.TradingDecimal, true
}

// HasTradingDecimal returns a boolean if a field has been set.
func (o *AcquireSingleMarketInformationResponse) HasTradingDecimal() bool {
	if o != nil && o.TradingDecimal != nil {
		return true
	}

	return false
}

// SetTradingDecimal gets a reference to the given int64 and assigns it to the TradingDecimal field.
func (o *AcquireSingleMarketInformationResponse) SetTradingDecimal(v int64) {
	o.TradingDecimal = &v
}

// GetPricingName returns the PricingName field value if set, zero value otherwise.
func (o *AcquireSingleMarketInformationResponse) GetPricingName() string {
	if o == nil || o.PricingName == nil {
		var ret string
		return ret
	}
	return *o.PricingName
}

// GetPricingNameOk returns a tuple with the PricingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireSingleMarketInformationResponse) GetPricingNameOk() (*string, bool) {
	if o == nil || o.PricingName == nil {
		return nil, false
	}
	return o.PricingName, true
}

// HasPricingName returns a boolean if a field has been set.
func (o *AcquireSingleMarketInformationResponse) HasPricingName() bool {
	if o != nil && o.PricingName != nil {
		return true
	}

	return false
}

// SetPricingName gets a reference to the given string and assigns it to the PricingName field.
func (o *AcquireSingleMarketInformationResponse) SetPricingName(v string) {
	o.PricingName = &v
}

// GetPricingDecimal returns the PricingDecimal field value if set, zero value otherwise.
func (o *AcquireSingleMarketInformationResponse) GetPricingDecimal() int64 {
	if o == nil || o.PricingDecimal == nil {
		var ret int64
		return ret
	}
	return *o.PricingDecimal
}

// GetPricingDecimalOk returns a tuple with the PricingDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireSingleMarketInformationResponse) GetPricingDecimalOk() (*int64, bool) {
	if o == nil || o.PricingDecimal == nil {
		return nil, false
	}
	return o.PricingDecimal, true
}

// HasPricingDecimal returns a boolean if a field has been set.
func (o *AcquireSingleMarketInformationResponse) HasPricingDecimal() bool {
	if o != nil && o.PricingDecimal != nil {
		return true
	}

	return false
}

// SetPricingDecimal gets a reference to the given int64 and assigns it to the PricingDecimal field.
func (o *AcquireSingleMarketInformationResponse) SetPricingDecimal(v int64) {
	o.PricingDecimal = &v
}

func (o AcquireSingleMarketInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TakerFeeRate != nil {
		toSerialize["taker_fee_rate"] = o.TakerFeeRate
	}
	if o.MakerFeeRate != nil {
		toSerialize["maker_fee_rate"] = o.MakerFeeRate
	}
	if o.MinAmount != nil {
		toSerialize["min_amount"] = o.MinAmount
	}
	if o.TradingName != nil {
		toSerialize["trading_name"] = o.TradingName
	}
	if o.TradingDecimal != nil {
		toSerialize["trading_decimal"] = o.TradingDecimal
	}
	if o.PricingName != nil {
		toSerialize["pricing_name"] = o.PricingName
	}
	if o.PricingDecimal != nil {
		toSerialize["pricing_decimal"] = o.PricingDecimal
	}
	return json.Marshal(toSerialize)
}

type NullableAcquireSingleMarketInformationResponse struct {
	value *AcquireSingleMarketInformationResponse
	isSet bool
}

func (v NullableAcquireSingleMarketInformationResponse) Get() *AcquireSingleMarketInformationResponse {
	return v.value
}

func (v *NullableAcquireSingleMarketInformationResponse) Set(val *AcquireSingleMarketInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcquireSingleMarketInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcquireSingleMarketInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcquireSingleMarketInformationResponse(val *AcquireSingleMarketInformationResponse) *NullableAcquireSingleMarketInformationResponse {
	return &NullableAcquireSingleMarketInformationResponse{value: val, isSet: true}
}

func (v NullableAcquireSingleMarketInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcquireSingleMarketInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
