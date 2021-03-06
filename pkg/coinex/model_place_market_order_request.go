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

// PlaceMarketOrderRequest struct for PlaceMarketOrderRequest
type PlaceMarketOrderRequest struct {
	// access_id
	AccessId string `json:"access_id"`
	// See <API invocation description market>
	Market string `json:"market"`
	// sell: Sell order
	Type string `json:"type"`
}

// NewPlaceMarketOrderRequest instantiates a new PlaceMarketOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceMarketOrderRequest(accessId string, market string, type_ string) *PlaceMarketOrderRequest {
	this := PlaceMarketOrderRequest{}
	this.AccessId = accessId
	this.Market = market
	this.Type = type_
	return &this
}

// NewPlaceMarketOrderRequestWithDefaults instantiates a new PlaceMarketOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceMarketOrderRequestWithDefaults() *PlaceMarketOrderRequest {
	this := PlaceMarketOrderRequest{}
	return &this
}

// GetAccessId returns the AccessId field value
func (o *PlaceMarketOrderRequest) GetAccessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value
// and a boolean to check if the value has been set.
func (o *PlaceMarketOrderRequest) GetAccessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessId, true
}

// SetAccessId sets field value
func (o *PlaceMarketOrderRequest) SetAccessId(v string) {
	o.AccessId = v
}

// GetMarket returns the Market field value
func (o *PlaceMarketOrderRequest) GetMarket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *PlaceMarketOrderRequest) GetMarketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *PlaceMarketOrderRequest) SetMarket(v string) {
	o.Market = v
}

// GetType returns the Type field value
func (o *PlaceMarketOrderRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PlaceMarketOrderRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PlaceMarketOrderRequest) SetType(v string) {
	o.Type = v
}

func (o PlaceMarketOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_id"] = o.AccessId
	}
	if true {
		toSerialize["market"] = o.Market
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePlaceMarketOrderRequest struct {
	value *PlaceMarketOrderRequest
	isSet bool
}

func (v NullablePlaceMarketOrderRequest) Get() *PlaceMarketOrderRequest {
	return v.value
}

func (v *NullablePlaceMarketOrderRequest) Set(val *PlaceMarketOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMarketOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMarketOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMarketOrderRequest(val *PlaceMarketOrderRequest) *NullablePlaceMarketOrderRequest {
	return &NullablePlaceMarketOrderRequest{value: val, isSet: true}
}

func (v NullablePlaceMarketOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMarketOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
