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
// Trade struct for Trade
type Trade struct {
	// trade id
	Id *int32 `json:"id,omitempty"`
	// price
	Price *string `json:"price,omitempty"`
	// amount of base asset
	Qty *string `json:"qty,omitempty"`
	// amount of quote asset
	QuoteQty *string `json:"quoteQty,omitempty"`
	// trade timestamp
	Time         *int32 `json:"time,omitempty"`
	IsBuyerMaker *bool  `json:"isBuyerMaker,omitempty"`
	IsBestMatch  *bool  `json:"isBestMatch,omitempty"`
}

// NewTrade instantiates a new Trade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrade() *Trade {
	this := Trade{}
	return &this
}

// NewTradeWithDefaults instantiates a new Trade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeWithDefaults() *Trade {
	this := Trade{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Trade) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Trade) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Trade) SetId(v int32) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Trade) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Trade) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *Trade) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *Trade) GetQty() string {
	if o == nil || o.Qty == nil {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetQtyOk() (*string, bool) {
	if o == nil || o.Qty == nil {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *Trade) HasQty() bool {
	if o != nil && o.Qty != nil {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *Trade) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *Trade) GetQuoteQty() string {
	if o == nil || o.QuoteQty == nil {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetQuoteQtyOk() (*string, bool) {
	if o == nil || o.QuoteQty == nil {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *Trade) HasQuoteQty() bool {
	if o != nil && o.QuoteQty != nil {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *Trade) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Trade) GetTime() int32 {
	if o == nil || o.Time == nil {
		var ret int32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetTimeOk() (*int32, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Trade) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int32 and assigns it to the Time field.
func (o *Trade) SetTime(v int32) {
	o.Time = &v
}

// GetIsBuyerMaker returns the IsBuyerMaker field value if set, zero value otherwise.
func (o *Trade) GetIsBuyerMaker() bool {
	if o == nil || o.IsBuyerMaker == nil {
		var ret bool
		return ret
	}
	return *o.IsBuyerMaker
}

// GetIsBuyerMakerOk returns a tuple with the IsBuyerMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetIsBuyerMakerOk() (*bool, bool) {
	if o == nil || o.IsBuyerMaker == nil {
		return nil, false
	}
	return o.IsBuyerMaker, true
}

// HasIsBuyerMaker returns a boolean if a field has been set.
func (o *Trade) HasIsBuyerMaker() bool {
	if o != nil && o.IsBuyerMaker != nil {
		return true
	}

	return false
}

// SetIsBuyerMaker gets a reference to the given bool and assigns it to the IsBuyerMaker field.
func (o *Trade) SetIsBuyerMaker(v bool) {
	o.IsBuyerMaker = &v
}

// GetIsBestMatch returns the IsBestMatch field value if set, zero value otherwise.
func (o *Trade) GetIsBestMatch() bool {
	if o == nil || o.IsBestMatch == nil {
		var ret bool
		return ret
	}
	return *o.IsBestMatch
}

// GetIsBestMatchOk returns a tuple with the IsBestMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetIsBestMatchOk() (*bool, bool) {
	if o == nil || o.IsBestMatch == nil {
		return nil, false
	}
	return o.IsBestMatch, true
}

// HasIsBestMatch returns a boolean if a field has been set.
func (o *Trade) HasIsBestMatch() bool {
	if o != nil && o.IsBestMatch != nil {
		return true
	}

	return false
}

// SetIsBestMatch gets a reference to the given bool and assigns it to the IsBestMatch field.
func (o *Trade) SetIsBestMatch(v bool) {
	o.IsBestMatch = &v
}

func (o Trade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Qty != nil {
		toSerialize["qty"] = o.Qty
	}
	if o.QuoteQty != nil {
		toSerialize["quoteQty"] = o.QuoteQty
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.IsBuyerMaker != nil {
		toSerialize["isBuyerMaker"] = o.IsBuyerMaker
	}
	if o.IsBestMatch != nil {
		toSerialize["isBestMatch"] = o.IsBestMatch
	}
	return json.Marshal(toSerialize)
}

type NullableTrade struct {
	value *Trade
	isSet bool
}

func (v NullableTrade) Get() *Trade {
	return v.value
}

func (v *NullableTrade) Set(val *Trade) {
	v.value = val
	v.isSet = true
}

func (v NullableTrade) IsSet() bool {
	return v.isSet
}

func (v *NullableTrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrade(val *Trade) *NullableTrade {
	return &NullableTrade{value: val, isSet: true}
}

func (v NullableTrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}