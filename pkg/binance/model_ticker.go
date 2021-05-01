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
// Ticker struct for Ticker
type Ticker struct {
	Symbol             *string `json:"symbol,omitempty"`
	PriceChange        *string `json:"priceChange,omitempty"`
	PriceChangePercent *string `json:"priceChangePercent,omitempty"`
	PrevClosePrice     *string `json:"prevClosePrice,omitempty"`
	LastPrice          *string `json:"lastPrice,omitempty"`
	BidPrice           *string `json:"bidPrice,omitempty"`
	BidQty             *string `json:"bidQty,omitempty"`
	AskPrice           *string `json:"askPrice,omitempty"`
	AskQty             *string `json:"askQty,omitempty"`
	OpenPrice          *string `json:"openPrice,omitempty"`
	HighPrice          *string `json:"highPrice,omitempty"`
	LowPrice           *string `json:"lowPrice,omitempty"`
	Volume             *string `json:"volume,omitempty"`
	QuoteVolume        *string `json:"quoteVolume,omitempty"`
	OpenTime           *int32  `json:"openTime,omitempty"`
	CloseTime          *int32  `json:"closeTime,omitempty"`
	FirstId            *int32  `json:"firstId,omitempty"`
	LastId             *int32  `json:"lastId,omitempty"`
	Count              *int32  `json:"count,omitempty"`
}

// NewTicker instantiates a new Ticker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker() *Ticker {
	this := Ticker{}
	return &this
}

// NewTickerWithDefaults instantiates a new Ticker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerWithDefaults() *Ticker {
	this := Ticker{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Ticker) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Ticker) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Ticker) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPriceChange returns the PriceChange field value if set, zero value otherwise.
func (o *Ticker) GetPriceChange() string {
	if o == nil || o.PriceChange == nil {
		var ret string
		return ret
	}
	return *o.PriceChange
}

// GetPriceChangeOk returns a tuple with the PriceChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetPriceChangeOk() (*string, bool) {
	if o == nil || o.PriceChange == nil {
		return nil, false
	}
	return o.PriceChange, true
}

// HasPriceChange returns a boolean if a field has been set.
func (o *Ticker) HasPriceChange() bool {
	if o != nil && o.PriceChange != nil {
		return true
	}

	return false
}

// SetPriceChange gets a reference to the given string and assigns it to the PriceChange field.
func (o *Ticker) SetPriceChange(v string) {
	o.PriceChange = &v
}

// GetPriceChangePercent returns the PriceChangePercent field value if set, zero value otherwise.
func (o *Ticker) GetPriceChangePercent() string {
	if o == nil || o.PriceChangePercent == nil {
		var ret string
		return ret
	}
	return *o.PriceChangePercent
}

// GetPriceChangePercentOk returns a tuple with the PriceChangePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetPriceChangePercentOk() (*string, bool) {
	if o == nil || o.PriceChangePercent == nil {
		return nil, false
	}
	return o.PriceChangePercent, true
}

// HasPriceChangePercent returns a boolean if a field has been set.
func (o *Ticker) HasPriceChangePercent() bool {
	if o != nil && o.PriceChangePercent != nil {
		return true
	}

	return false
}

// SetPriceChangePercent gets a reference to the given string and assigns it to the PriceChangePercent field.
func (o *Ticker) SetPriceChangePercent(v string) {
	o.PriceChangePercent = &v
}

// GetPrevClosePrice returns the PrevClosePrice field value if set, zero value otherwise.
func (o *Ticker) GetPrevClosePrice() string {
	if o == nil || o.PrevClosePrice == nil {
		var ret string
		return ret
	}
	return *o.PrevClosePrice
}

// GetPrevClosePriceOk returns a tuple with the PrevClosePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetPrevClosePriceOk() (*string, bool) {
	if o == nil || o.PrevClosePrice == nil {
		return nil, false
	}
	return o.PrevClosePrice, true
}

// HasPrevClosePrice returns a boolean if a field has been set.
func (o *Ticker) HasPrevClosePrice() bool {
	if o != nil && o.PrevClosePrice != nil {
		return true
	}

	return false
}

// SetPrevClosePrice gets a reference to the given string and assigns it to the PrevClosePrice field.
func (o *Ticker) SetPrevClosePrice(v string) {
	o.PrevClosePrice = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *Ticker) GetLastPrice() string {
	if o == nil || o.LastPrice == nil {
		var ret string
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetLastPriceOk() (*string, bool) {
	if o == nil || o.LastPrice == nil {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *Ticker) HasLastPrice() bool {
	if o != nil && o.LastPrice != nil {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given string and assigns it to the LastPrice field.
func (o *Ticker) SetLastPrice(v string) {
	o.LastPrice = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *Ticker) GetBidPrice() string {
	if o == nil || o.BidPrice == nil {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetBidPriceOk() (*string, bool) {
	if o == nil || o.BidPrice == nil {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *Ticker) HasBidPrice() bool {
	if o != nil && o.BidPrice != nil {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *Ticker) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetBidQty returns the BidQty field value if set, zero value otherwise.
func (o *Ticker) GetBidQty() string {
	if o == nil || o.BidQty == nil {
		var ret string
		return ret
	}
	return *o.BidQty
}

// GetBidQtyOk returns a tuple with the BidQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetBidQtyOk() (*string, bool) {
	if o == nil || o.BidQty == nil {
		return nil, false
	}
	return o.BidQty, true
}

// HasBidQty returns a boolean if a field has been set.
func (o *Ticker) HasBidQty() bool {
	if o != nil && o.BidQty != nil {
		return true
	}

	return false
}

// SetBidQty gets a reference to the given string and assigns it to the BidQty field.
func (o *Ticker) SetBidQty(v string) {
	o.BidQty = &v
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *Ticker) GetAskPrice() string {
	if o == nil || o.AskPrice == nil {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetAskPriceOk() (*string, bool) {
	if o == nil || o.AskPrice == nil {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *Ticker) HasAskPrice() bool {
	if o != nil && o.AskPrice != nil {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *Ticker) SetAskPrice(v string) {
	o.AskPrice = &v
}

// GetAskQty returns the AskQty field value if set, zero value otherwise.
func (o *Ticker) GetAskQty() string {
	if o == nil || o.AskQty == nil {
		var ret string
		return ret
	}
	return *o.AskQty
}

// GetAskQtyOk returns a tuple with the AskQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetAskQtyOk() (*string, bool) {
	if o == nil || o.AskQty == nil {
		return nil, false
	}
	return o.AskQty, true
}

// HasAskQty returns a boolean if a field has been set.
func (o *Ticker) HasAskQty() bool {
	if o != nil && o.AskQty != nil {
		return true
	}

	return false
}

// SetAskQty gets a reference to the given string and assigns it to the AskQty field.
func (o *Ticker) SetAskQty(v string) {
	o.AskQty = &v
}

// GetOpenPrice returns the OpenPrice field value if set, zero value otherwise.
func (o *Ticker) GetOpenPrice() string {
	if o == nil || o.OpenPrice == nil {
		var ret string
		return ret
	}
	return *o.OpenPrice
}

// GetOpenPriceOk returns a tuple with the OpenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetOpenPriceOk() (*string, bool) {
	if o == nil || o.OpenPrice == nil {
		return nil, false
	}
	return o.OpenPrice, true
}

// HasOpenPrice returns a boolean if a field has been set.
func (o *Ticker) HasOpenPrice() bool {
	if o != nil && o.OpenPrice != nil {
		return true
	}

	return false
}

// SetOpenPrice gets a reference to the given string and assigns it to the OpenPrice field.
func (o *Ticker) SetOpenPrice(v string) {
	o.OpenPrice = &v
}

// GetHighPrice returns the HighPrice field value if set, zero value otherwise.
func (o *Ticker) GetHighPrice() string {
	if o == nil || o.HighPrice == nil {
		var ret string
		return ret
	}
	return *o.HighPrice
}

// GetHighPriceOk returns a tuple with the HighPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetHighPriceOk() (*string, bool) {
	if o == nil || o.HighPrice == nil {
		return nil, false
	}
	return o.HighPrice, true
}

// HasHighPrice returns a boolean if a field has been set.
func (o *Ticker) HasHighPrice() bool {
	if o != nil && o.HighPrice != nil {
		return true
	}

	return false
}

// SetHighPrice gets a reference to the given string and assigns it to the HighPrice field.
func (o *Ticker) SetHighPrice(v string) {
	o.HighPrice = &v
}

// GetLowPrice returns the LowPrice field value if set, zero value otherwise.
func (o *Ticker) GetLowPrice() string {
	if o == nil || o.LowPrice == nil {
		var ret string
		return ret
	}
	return *o.LowPrice
}

// GetLowPriceOk returns a tuple with the LowPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetLowPriceOk() (*string, bool) {
	if o == nil || o.LowPrice == nil {
		return nil, false
	}
	return o.LowPrice, true
}

// HasLowPrice returns a boolean if a field has been set.
func (o *Ticker) HasLowPrice() bool {
	if o != nil && o.LowPrice != nil {
		return true
	}

	return false
}

// SetLowPrice gets a reference to the given string and assigns it to the LowPrice field.
func (o *Ticker) SetLowPrice(v string) {
	o.LowPrice = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Ticker) GetVolume() string {
	if o == nil || o.Volume == nil {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetVolumeOk() (*string, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *Ticker) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *Ticker) SetVolume(v string) {
	o.Volume = &v
}

// GetQuoteVolume returns the QuoteVolume field value if set, zero value otherwise.
func (o *Ticker) GetQuoteVolume() string {
	if o == nil || o.QuoteVolume == nil {
		var ret string
		return ret
	}
	return *o.QuoteVolume
}

// GetQuoteVolumeOk returns a tuple with the QuoteVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetQuoteVolumeOk() (*string, bool) {
	if o == nil || o.QuoteVolume == nil {
		return nil, false
	}
	return o.QuoteVolume, true
}

// HasQuoteVolume returns a boolean if a field has been set.
func (o *Ticker) HasQuoteVolume() bool {
	if o != nil && o.QuoteVolume != nil {
		return true
	}

	return false
}

// SetQuoteVolume gets a reference to the given string and assigns it to the QuoteVolume field.
func (o *Ticker) SetQuoteVolume(v string) {
	o.QuoteVolume = &v
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *Ticker) GetOpenTime() int32 {
	if o == nil || o.OpenTime == nil {
		var ret int32
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetOpenTimeOk() (*int32, bool) {
	if o == nil || o.OpenTime == nil {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *Ticker) HasOpenTime() bool {
	if o != nil && o.OpenTime != nil {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given int32 and assigns it to the OpenTime field.
func (o *Ticker) SetOpenTime(v int32) {
	o.OpenTime = &v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *Ticker) GetCloseTime() int32 {
	if o == nil || o.CloseTime == nil {
		var ret int32
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetCloseTimeOk() (*int32, bool) {
	if o == nil || o.CloseTime == nil {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *Ticker) HasCloseTime() bool {
	if o != nil && o.CloseTime != nil {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given int32 and assigns it to the CloseTime field.
func (o *Ticker) SetCloseTime(v int32) {
	o.CloseTime = &v
}

// GetFirstId returns the FirstId field value if set, zero value otherwise.
func (o *Ticker) GetFirstId() int32 {
	if o == nil || o.FirstId == nil {
		var ret int32
		return ret
	}
	return *o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetFirstIdOk() (*int32, bool) {
	if o == nil || o.FirstId == nil {
		return nil, false
	}
	return o.FirstId, true
}

// HasFirstId returns a boolean if a field has been set.
func (o *Ticker) HasFirstId() bool {
	if o != nil && o.FirstId != nil {
		return true
	}

	return false
}

// SetFirstId gets a reference to the given int32 and assigns it to the FirstId field.
func (o *Ticker) SetFirstId(v int32) {
	o.FirstId = &v
}

// GetLastId returns the LastId field value if set, zero value otherwise.
func (o *Ticker) GetLastId() int32 {
	if o == nil || o.LastId == nil {
		var ret int32
		return ret
	}
	return *o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetLastIdOk() (*int32, bool) {
	if o == nil || o.LastId == nil {
		return nil, false
	}
	return o.LastId, true
}

// HasLastId returns a boolean if a field has been set.
func (o *Ticker) HasLastId() bool {
	if o != nil && o.LastId != nil {
		return true
	}

	return false
}

// SetLastId gets a reference to the given int32 and assigns it to the LastId field.
func (o *Ticker) SetLastId(v int32) {
	o.LastId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Ticker) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Ticker) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Ticker) SetCount(v int32) {
	o.Count = &v
}

func (o Ticker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.PriceChange != nil {
		toSerialize["priceChange"] = o.PriceChange
	}
	if o.PriceChangePercent != nil {
		toSerialize["priceChangePercent"] = o.PriceChangePercent
	}
	if o.PrevClosePrice != nil {
		toSerialize["prevClosePrice"] = o.PrevClosePrice
	}
	if o.LastPrice != nil {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if o.BidPrice != nil {
		toSerialize["bidPrice"] = o.BidPrice
	}
	if o.BidQty != nil {
		toSerialize["bidQty"] = o.BidQty
	}
	if o.AskPrice != nil {
		toSerialize["askPrice"] = o.AskPrice
	}
	if o.AskQty != nil {
		toSerialize["askQty"] = o.AskQty
	}
	if o.OpenPrice != nil {
		toSerialize["openPrice"] = o.OpenPrice
	}
	if o.HighPrice != nil {
		toSerialize["highPrice"] = o.HighPrice
	}
	if o.LowPrice != nil {
		toSerialize["lowPrice"] = o.LowPrice
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if o.QuoteVolume != nil {
		toSerialize["quoteVolume"] = o.QuoteVolume
	}
	if o.OpenTime != nil {
		toSerialize["openTime"] = o.OpenTime
	}
	if o.CloseTime != nil {
		toSerialize["closeTime"] = o.CloseTime
	}
	if o.FirstId != nil {
		toSerialize["firstId"] = o.FirstId
	}
	if o.LastId != nil {
		toSerialize["lastId"] = o.LastId
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableTicker struct {
	value *Ticker
	isSet bool
}

func (v NullableTicker) Get() *Ticker {
	return v.value
}

func (v *NullableTicker) Set(val *Ticker) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicker(val *Ticker) *NullableTicker {
	return &NullableTicker{value: val, isSet: true}
}

func (v NullableTicker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
