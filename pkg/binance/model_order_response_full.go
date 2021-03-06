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
// OrderResponseFull struct for OrderResponseFull
type OrderResponseFull struct {
	Symbol  *string `json:"symbol,omitempty"`
	OrderId *int32  `json:"orderId,omitempty"`
	// Unless OCO, value will be -1
	OrderListId         *int32    `json:"orderListId,omitempty"`
	ClientOrderId       *string   `json:"clientOrderId,omitempty"`
	TransactTime        *int32    `json:"transactTime,omitempty"`
	Price               *float32  `json:"price,omitempty"`
	OrigQty             *float32  `json:"origQty,omitempty"`
	ExecutedQty         *float32  `json:"executedQty,omitempty"`
	CummulativeQuoteQty *float32  `json:"cummulativeQuoteQty,omitempty"`
	Status              *string   `json:"status,omitempty"`
	TimeInForce         *string   `json:"timeInForce,omitempty"`
	Type                *string   `json:"type,omitempty"`
	Side                *string   `json:"side,omitempty"`
	Fills               *[]string `json:"fills,omitempty"`
}

// NewOrderResponseFull instantiates a new OrderResponseFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseFull() *OrderResponseFull {
	this := OrderResponseFull{}
	return &this
}

// NewOrderResponseFullWithDefaults instantiates a new OrderResponseFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseFullWithDefaults() *OrderResponseFull {
	this := OrderResponseFull{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderResponseFull) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderResponseFull) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderResponseFull) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderResponseFull) GetOrderId() int32 {
	if o == nil || o.OrderId == nil {
		var ret int32
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetOrderIdOk() (*int32, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderResponseFull) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int32 and assigns it to the OrderId field.
func (o *OrderResponseFull) SetOrderId(v int32) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderResponseFull) GetOrderListId() int32 {
	if o == nil || o.OrderListId == nil {
		var ret int32
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetOrderListIdOk() (*int32, bool) {
	if o == nil || o.OrderListId == nil {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderResponseFull) HasOrderListId() bool {
	if o != nil && o.OrderListId != nil {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int32 and assigns it to the OrderListId field.
func (o *OrderResponseFull) SetOrderListId(v int32) {
	o.OrderListId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OrderResponseFull) GetClientOrderId() string {
	if o == nil || o.ClientOrderId == nil {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetClientOrderIdOk() (*string, bool) {
	if o == nil || o.ClientOrderId == nil {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OrderResponseFull) HasClientOrderId() bool {
	if o != nil && o.ClientOrderId != nil {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OrderResponseFull) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *OrderResponseFull) GetTransactTime() int32 {
	if o == nil || o.TransactTime == nil {
		var ret int32
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetTransactTimeOk() (*int32, bool) {
	if o == nil || o.TransactTime == nil {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *OrderResponseFull) HasTransactTime() bool {
	if o != nil && o.TransactTime != nil {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int32 and assigns it to the TransactTime field.
func (o *OrderResponseFull) SetTransactTime(v int32) {
	o.TransactTime = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OrderResponseFull) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OrderResponseFull) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *OrderResponseFull) SetPrice(v float32) {
	o.Price = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *OrderResponseFull) GetOrigQty() float32 {
	if o == nil || o.OrigQty == nil {
		var ret float32
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetOrigQtyOk() (*float32, bool) {
	if o == nil || o.OrigQty == nil {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *OrderResponseFull) HasOrigQty() bool {
	if o != nil && o.OrigQty != nil {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given float32 and assigns it to the OrigQty field.
func (o *OrderResponseFull) SetOrigQty(v float32) {
	o.OrigQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *OrderResponseFull) GetExecutedQty() float32 {
	if o == nil || o.ExecutedQty == nil {
		var ret float32
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetExecutedQtyOk() (*float32, bool) {
	if o == nil || o.ExecutedQty == nil {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *OrderResponseFull) HasExecutedQty() bool {
	if o != nil && o.ExecutedQty != nil {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given float32 and assigns it to the ExecutedQty field.
func (o *OrderResponseFull) SetExecutedQty(v float32) {
	o.ExecutedQty = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *OrderResponseFull) GetCummulativeQuoteQty() float32 {
	if o == nil || o.CummulativeQuoteQty == nil {
		var ret float32
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetCummulativeQuoteQtyOk() (*float32, bool) {
	if o == nil || o.CummulativeQuoteQty == nil {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *OrderResponseFull) HasCummulativeQuoteQty() bool {
	if o != nil && o.CummulativeQuoteQty != nil {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given float32 and assigns it to the CummulativeQuoteQty field.
func (o *OrderResponseFull) SetCummulativeQuoteQty(v float32) {
	o.CummulativeQuoteQty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderResponseFull) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderResponseFull) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderResponseFull) SetStatus(v string) {
	o.Status = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *OrderResponseFull) GetTimeInForce() string {
	if o == nil || o.TimeInForce == nil {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetTimeInForceOk() (*string, bool) {
	if o == nil || o.TimeInForce == nil {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *OrderResponseFull) HasTimeInForce() bool {
	if o != nil && o.TimeInForce != nil {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *OrderResponseFull) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderResponseFull) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderResponseFull) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderResponseFull) SetType(v string) {
	o.Type = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OrderResponseFull) GetSide() string {
	if o == nil || o.Side == nil {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetSideOk() (*string, bool) {
	if o == nil || o.Side == nil {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OrderResponseFull) HasSide() bool {
	if o != nil && o.Side != nil {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *OrderResponseFull) SetSide(v string) {
	o.Side = &v
}

// GetFills returns the Fills field value if set, zero value otherwise.
func (o *OrderResponseFull) GetFills() []string {
	if o == nil || o.Fills == nil {
		var ret []string
		return ret
	}
	return *o.Fills
}

// GetFillsOk returns a tuple with the Fills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFull) GetFillsOk() (*[]string, bool) {
	if o == nil || o.Fills == nil {
		return nil, false
	}
	return o.Fills, true
}

// HasFills returns a boolean if a field has been set.
func (o *OrderResponseFull) HasFills() bool {
	if o != nil && o.Fills != nil {
		return true
	}

	return false
}

// SetFills gets a reference to the given []string and assigns it to the Fills field.
func (o *OrderResponseFull) SetFills(v []string) {
	o.Fills = &v
}

func (o OrderResponseFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.OrderId != nil {
		toSerialize["orderId"] = o.OrderId
	}
	if o.OrderListId != nil {
		toSerialize["orderListId"] = o.OrderListId
	}
	if o.ClientOrderId != nil {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if o.TransactTime != nil {
		toSerialize["transactTime"] = o.TransactTime
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.OrigQty != nil {
		toSerialize["origQty"] = o.OrigQty
	}
	if o.ExecutedQty != nil {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if o.CummulativeQuoteQty != nil {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TimeInForce != nil {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Side != nil {
		toSerialize["side"] = o.Side
	}
	if o.Fills != nil {
		toSerialize["fills"] = o.Fills
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResponseFull struct {
	value *OrderResponseFull
	isSet bool
}

func (v NullableOrderResponseFull) Get() *OrderResponseFull {
	return v.value
}

func (v *NullableOrderResponseFull) Set(val *OrderResponseFull) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseFull) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseFull(val *OrderResponseFull) *NullableOrderResponseFull {
	return &NullableOrderResponseFull{value: val, isSet: true}
}

func (v NullableOrderResponseFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
