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

// AcquireAllMarketDataResponse AcquireAllMarketData Response Value
type AcquireAllMarketDataResponse struct {
	// server time when returning
	Date *string `json:"date,omitempty"`
	// buy 1
	Buy *string `json:"buy,omitempty"`
	// buy 1 amount
	BuyAmount *string `json:"buy_amount,omitempty"`
	// 24H highest price
	High *string `json:"high,omitempty"`
	// latest price
	Last *string `json:"last,omitempty"`
	// 24H lowest price
	Low *string `json:"low,omitempty"`
	// sell 1
	Sell *string `json:"sell,omitempty"`
	// sell 1 amount
	SellAmount *string `json:"sell_amount,omitempty"`
	// 24H volumn
	Vol *string `json:"vol,omitempty"`
}

// NewAcquireAllMarketDataResponse instantiates a new AcquireAllMarketDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcquireAllMarketDataResponse() *AcquireAllMarketDataResponse {
	this := AcquireAllMarketDataResponse{}
	return &this
}

// NewAcquireAllMarketDataResponseWithDefaults instantiates a new AcquireAllMarketDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcquireAllMarketDataResponseWithDefaults() *AcquireAllMarketDataResponse {
	this := AcquireAllMarketDataResponse{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AcquireAllMarketDataResponse) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireAllMarketDataResponse) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AcquireAllMarketDataResponse) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AcquireAllMarketDataResponse) SetDate(v string) {
	o.Date = &v
}

// GetBuy returns the Buy field value if set, zero value otherwise.
func (o *AcquireAllMarketDataResponse) GetBuy() string {
	if o == nil || o.Buy == nil {
		var ret string
		return ret
	}
	return *o.Buy
}

// GetBuyOk returns a tuple with the Buy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireAllMarketDataResponse) GetBuyOk() (*string, bool) {
	if o == nil || o.Buy == nil {
		return nil, false
	}
	return o.Buy, true
}

// HasBuy returns a boolean if a field has been set.
func (o *AcquireAllMarketDataResponse) HasBuy() bool {
	if o != nil && o.Buy != nil {
		return true
	}

	return false
}

// SetBuy gets a reference to the given string and assigns it to the Buy field.
func (o *AcquireAllMarketDataResponse) SetBuy(v string) {
	o.Buy = &v
}

// GetBuyAmount returns the BuyAmount field value if set, zero value otherwise.
func (o *AcquireAllMarketDataResponse) GetBuyAmount() string {
	if o == nil || o.BuyAmount == nil {
		var ret string
		return ret
	}
	return *o.BuyAmount
}

// GetBuyAmountOk returns a tuple with the BuyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireAllMarketDataResponse) GetBuyAmountOk() (*string, bool) {
	if o == nil || o.BuyAmount == nil {
		return nil, false
	}
	return o.BuyAmount, true
}

// HasBuyAmount returns a boolean if a field has been set.
func (o *AcquireAllMarketDataResponse) HasBuyAmount() bool {
	if o != nil && o.BuyAmount != nil {
		return true
	}

	return false
}

// SetBuyAmount gets a reference to the given string and assigns it to the BuyAmount field.
func (o *AcquireAllMarketDataResponse) SetBuyAmount(v string) {
	o.BuyAmount = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *AcquireAllMarketDataResponse) GetHigh() string {
	if o == nil || o.High == nil {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireAllMarketDataResponse) GetHighOk() (*string, bool) {
	if o == nil || o.High == nil {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *AcquireAllMarketDataResponse) HasHigh() bool {
	if o != nil && o.High != nil {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *AcquireAllMarketDataResponse) SetHigh(v string) {
	o.High = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *AcquireAllMarketDataResponse) GetLast() string {
	if o == nil || o.Last == nil {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireAllMarketDataResponse) GetLastOk() (*string, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *AcquireAllMarketDataResponse) HasLast() bool {
	if o != nil && o.Last != nil {
		return true
	}

	return false
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *AcquireAllMarketDataResponse) SetLast(v string) {
	o.Last = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *AcquireAllMarketDataResponse) GetLow() string {
	if o == nil || o.Low == nil {
		var ret string
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireAllMarketDataResponse) GetLowOk() (*string, bool) {
	if o == nil || o.Low == nil {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *AcquireAllMarketDataResponse) HasLow() bool {
	if o != nil && o.Low != nil {
		return true
	}

	return false
}

// SetLow gets a reference to the given string and assigns it to the Low field.
func (o *AcquireAllMarketDataResponse) SetLow(v string) {
	o.Low = &v
}

// GetSell returns the Sell field value if set, zero value otherwise.
func (o *AcquireAllMarketDataResponse) GetSell() string {
	if o == nil || o.Sell == nil {
		var ret string
		return ret
	}
	return *o.Sell
}

// GetSellOk returns a tuple with the Sell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireAllMarketDataResponse) GetSellOk() (*string, bool) {
	if o == nil || o.Sell == nil {
		return nil, false
	}
	return o.Sell, true
}

// HasSell returns a boolean if a field has been set.
func (o *AcquireAllMarketDataResponse) HasSell() bool {
	if o != nil && o.Sell != nil {
		return true
	}

	return false
}

// SetSell gets a reference to the given string and assigns it to the Sell field.
func (o *AcquireAllMarketDataResponse) SetSell(v string) {
	o.Sell = &v
}

// GetSellAmount returns the SellAmount field value if set, zero value otherwise.
func (o *AcquireAllMarketDataResponse) GetSellAmount() string {
	if o == nil || o.SellAmount == nil {
		var ret string
		return ret
	}
	return *o.SellAmount
}

// GetSellAmountOk returns a tuple with the SellAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireAllMarketDataResponse) GetSellAmountOk() (*string, bool) {
	if o == nil || o.SellAmount == nil {
		return nil, false
	}
	return o.SellAmount, true
}

// HasSellAmount returns a boolean if a field has been set.
func (o *AcquireAllMarketDataResponse) HasSellAmount() bool {
	if o != nil && o.SellAmount != nil {
		return true
	}

	return false
}

// SetSellAmount gets a reference to the given string and assigns it to the SellAmount field.
func (o *AcquireAllMarketDataResponse) SetSellAmount(v string) {
	o.SellAmount = &v
}

// GetVol returns the Vol field value if set, zero value otherwise.
func (o *AcquireAllMarketDataResponse) GetVol() string {
	if o == nil || o.Vol == nil {
		var ret string
		return ret
	}
	return *o.Vol
}

// GetVolOk returns a tuple with the Vol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireAllMarketDataResponse) GetVolOk() (*string, bool) {
	if o == nil || o.Vol == nil {
		return nil, false
	}
	return o.Vol, true
}

// HasVol returns a boolean if a field has been set.
func (o *AcquireAllMarketDataResponse) HasVol() bool {
	if o != nil && o.Vol != nil {
		return true
	}

	return false
}

// SetVol gets a reference to the given string and assigns it to the Vol field.
func (o *AcquireAllMarketDataResponse) SetVol(v string) {
	o.Vol = &v
}

func (o AcquireAllMarketDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Buy != nil {
		toSerialize["buy"] = o.Buy
	}
	if o.BuyAmount != nil {
		toSerialize["buy_amount"] = o.BuyAmount
	}
	if o.High != nil {
		toSerialize["high"] = o.High
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.Low != nil {
		toSerialize["low"] = o.Low
	}
	if o.Sell != nil {
		toSerialize["sell"] = o.Sell
	}
	if o.SellAmount != nil {
		toSerialize["sell_amount"] = o.SellAmount
	}
	if o.Vol != nil {
		toSerialize["vol"] = o.Vol
	}
	return json.Marshal(toSerialize)
}

type NullableAcquireAllMarketDataResponse struct {
	value *AcquireAllMarketDataResponse
	isSet bool
}

func (v NullableAcquireAllMarketDataResponse) Get() *AcquireAllMarketDataResponse {
	return v.value
}

func (v *NullableAcquireAllMarketDataResponse) Set(val *AcquireAllMarketDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcquireAllMarketDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcquireAllMarketDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcquireAllMarketDataResponse(val *AcquireAllMarketDataResponse) *NullableAcquireAllMarketDataResponse {
	return &NullableAcquireAllMarketDataResponse{value: val, isSet: true}
}

func (v NullableAcquireAllMarketDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcquireAllMarketDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
