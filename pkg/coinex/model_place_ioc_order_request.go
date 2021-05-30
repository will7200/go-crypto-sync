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

// PlaceIocOrderRequest struct for PlaceIocOrderRequest
type PlaceIocOrderRequest struct {
	// access_id
	AccessId string `json:"access_id"`
	// See <API invocation description market>
	Market string `json:"market"`
	// sell: sell order
	Type string `json:"type"`
}

// NewPlaceIocOrderRequest instantiates a new PlaceIocOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceIocOrderRequest(accessId string, market string, type_ string) *PlaceIocOrderRequest {
	this := PlaceIocOrderRequest{}
	this.AccessId = accessId
	this.Market = market
	this.Type = type_
	return &this
}

// NewPlaceIocOrderRequestWithDefaults instantiates a new PlaceIocOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceIocOrderRequestWithDefaults() *PlaceIocOrderRequest {
	this := PlaceIocOrderRequest{}
	return &this
}

// GetAccessId returns the AccessId field value
func (o *PlaceIocOrderRequest) GetAccessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value
// and a boolean to check if the value has been set.
func (o *PlaceIocOrderRequest) GetAccessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessId, true
}

// SetAccessId sets field value
func (o *PlaceIocOrderRequest) SetAccessId(v string) {
	o.AccessId = v
}

// GetMarket returns the Market field value
func (o *PlaceIocOrderRequest) GetMarket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *PlaceIocOrderRequest) GetMarketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *PlaceIocOrderRequest) SetMarket(v string) {
	o.Market = v
}

// GetType returns the Type field value
func (o *PlaceIocOrderRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PlaceIocOrderRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PlaceIocOrderRequest) SetType(v string) {
	o.Type = v
}

func (o PlaceIocOrderRequest) MarshalJSON() ([]byte, error) {
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

type NullablePlaceIocOrderRequest struct {
	value *PlaceIocOrderRequest
	isSet bool
}

func (v NullablePlaceIocOrderRequest) Get() *PlaceIocOrderRequest {
	return v.value
}

func (v *NullablePlaceIocOrderRequest) Set(val *PlaceIocOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceIocOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceIocOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceIocOrderRequest(val *PlaceIocOrderRequest) *NullablePlaceIocOrderRequest {
	return &NullablePlaceIocOrderRequest{value: val, isSet: true}
}

func (v NullablePlaceIocOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceIocOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}