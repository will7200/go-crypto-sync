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

// InlineResponse20013 struct for InlineResponse20013
type InlineResponse20013 struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	// Unknown Response
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewInlineResponse20013 instantiates a new InlineResponse20013 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20013() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20013WithDefaults() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse20013) GetCode() int64 {
	if o == nil || o.Code == nil {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetCodeOk() (*int64, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20013) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *InlineResponse20013) SetCode(v int64) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse20013) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse20013) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse20013) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse20013) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse20013) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *InlineResponse20013) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o InlineResponse20013) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20013 struct {
	value *InlineResponse20013
	isSet bool
}

func (v NullableInlineResponse20013) Get() *InlineResponse20013 {
	return v.value
}

func (v *NullableInlineResponse20013) Set(val *InlineResponse20013) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20013) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20013) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20013(val *InlineResponse20013) *NullableInlineResponse20013 {
	return &NullableInlineResponse20013{value: val, isSet: true}
}

func (v NullableInlineResponse20013) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20013) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
