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

// PlaceMultipleLimitOrdersRequest struct for PlaceMultipleLimitOrdersRequest
type PlaceMultipleLimitOrdersRequest struct {
	// access_id
	AccessId    string  `json:"access_id"`
	BatchOrders *string `json:"batch_orders,omitempty"`
}

// NewPlaceMultipleLimitOrdersRequest instantiates a new PlaceMultipleLimitOrdersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceMultipleLimitOrdersRequest(accessId string) *PlaceMultipleLimitOrdersRequest {
	this := PlaceMultipleLimitOrdersRequest{}
	this.AccessId = accessId
	return &this
}

// NewPlaceMultipleLimitOrdersRequestWithDefaults instantiates a new PlaceMultipleLimitOrdersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceMultipleLimitOrdersRequestWithDefaults() *PlaceMultipleLimitOrdersRequest {
	this := PlaceMultipleLimitOrdersRequest{}
	return &this
}

// GetAccessId returns the AccessId field value
func (o *PlaceMultipleLimitOrdersRequest) GetAccessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value
// and a boolean to check if the value has been set.
func (o *PlaceMultipleLimitOrdersRequest) GetAccessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessId, true
}

// SetAccessId sets field value
func (o *PlaceMultipleLimitOrdersRequest) SetAccessId(v string) {
	o.AccessId = v
}

// GetBatchOrders returns the BatchOrders field value if set, zero value otherwise.
func (o *PlaceMultipleLimitOrdersRequest) GetBatchOrders() string {
	if o == nil || o.BatchOrders == nil {
		var ret string
		return ret
	}
	return *o.BatchOrders
}

// GetBatchOrdersOk returns a tuple with the BatchOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleLimitOrdersRequest) GetBatchOrdersOk() (*string, bool) {
	if o == nil || o.BatchOrders == nil {
		return nil, false
	}
	return o.BatchOrders, true
}

// HasBatchOrders returns a boolean if a field has been set.
func (o *PlaceMultipleLimitOrdersRequest) HasBatchOrders() bool {
	if o != nil && o.BatchOrders != nil {
		return true
	}

	return false
}

// SetBatchOrders gets a reference to the given string and assigns it to the BatchOrders field.
func (o *PlaceMultipleLimitOrdersRequest) SetBatchOrders(v string) {
	o.BatchOrders = &v
}

func (o PlaceMultipleLimitOrdersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_id"] = o.AccessId
	}
	if o.BatchOrders != nil {
		toSerialize["batch_orders"] = o.BatchOrders
	}
	return json.Marshal(toSerialize)
}

type NullablePlaceMultipleLimitOrdersRequest struct {
	value *PlaceMultipleLimitOrdersRequest
	isSet bool
}

func (v NullablePlaceMultipleLimitOrdersRequest) Get() *PlaceMultipleLimitOrdersRequest {
	return v.value
}

func (v *NullablePlaceMultipleLimitOrdersRequest) Set(val *PlaceMultipleLimitOrdersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleLimitOrdersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleLimitOrdersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleLimitOrdersRequest(val *PlaceMultipleLimitOrdersRequest) *NullablePlaceMultipleLimitOrdersRequest {
	return &NullablePlaceMultipleLimitOrdersRequest{value: val, isSet: true}
}

func (v NullablePlaceMultipleLimitOrdersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleLimitOrdersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
