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

// InlineResponse20023 struct for InlineResponse20023
type InlineResponse20023 struct {
	Code    *int64                      `json:"code,omitempty"`
	Message *string                     `json:"message,omitempty"`
	Data    *InquireDepositListResponse `json:"data,omitempty"`
}

// NewInlineResponse20023 instantiates a new InlineResponse20023 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023() *InlineResponse20023 {
	this := InlineResponse20023{}
	return &this
}

// NewInlineResponse20023WithDefaults instantiates a new InlineResponse20023 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023WithDefaults() *InlineResponse20023 {
	this := InlineResponse20023{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse20023) GetCode() int64 {
	if o == nil || o.Code == nil {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetCodeOk() (*int64, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20023) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *InlineResponse20023) SetCode(v int64) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse20023) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse20023) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse20023) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse20023) GetData() InquireDepositListResponse {
	if o == nil || o.Data == nil {
		var ret InquireDepositListResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetDataOk() (*InquireDepositListResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse20023) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given InquireDepositListResponse and assigns it to the Data field.
func (o *InlineResponse20023) SetData(v InquireDepositListResponse) {
	o.Data = &v
}

func (o InlineResponse20023) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023 struct {
	value *InlineResponse20023
	isSet bool
}

func (v NullableInlineResponse20023) Get() *InlineResponse20023 {
	return v.value
}

func (v *NullableInlineResponse20023) Set(val *InlineResponse20023) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023(val *InlineResponse20023) *NullableInlineResponse20023 {
	return &NullableInlineResponse20023{value: val, isSet: true}
}

func (v NullableInlineResponse20023) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
