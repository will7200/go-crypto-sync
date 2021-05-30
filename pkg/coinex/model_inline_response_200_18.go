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

// InlineResponse20018 struct for InlineResponse20018
type InlineResponse20018 struct {
	Code    *int64                      `json:"code,omitempty"`
	Message *string                     `json:"message,omitempty"`
	Data    *AcquireOrderStatusResponse `json:"data,omitempty"`
}

// NewInlineResponse20018 instantiates a new InlineResponse20018 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20018() *InlineResponse20018 {
	this := InlineResponse20018{}
	return &this
}

// NewInlineResponse20018WithDefaults instantiates a new InlineResponse20018 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20018WithDefaults() *InlineResponse20018 {
	this := InlineResponse20018{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse20018) GetCode() int64 {
	if o == nil || o.Code == nil {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018) GetCodeOk() (*int64, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20018) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *InlineResponse20018) SetCode(v int64) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse20018) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse20018) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse20018) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse20018) GetData() AcquireOrderStatusResponse {
	if o == nil || o.Data == nil {
		var ret AcquireOrderStatusResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018) GetDataOk() (*AcquireOrderStatusResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse20018) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AcquireOrderStatusResponse and assigns it to the Data field.
func (o *InlineResponse20018) SetData(v AcquireOrderStatusResponse) {
	o.Data = &v
}

func (o InlineResponse20018) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20018 struct {
	value *InlineResponse20018
	isSet bool
}

func (v NullableInlineResponse20018) Get() *InlineResponse20018 {
	return v.value
}

func (v *NullableInlineResponse20018) Set(val *InlineResponse20018) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20018) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20018) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20018(val *InlineResponse20018) *NullableInlineResponse20018 {
	return &NullableInlineResponse20018{value: val, isSet: true}
}

func (v NullableInlineResponse20018) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20018) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
