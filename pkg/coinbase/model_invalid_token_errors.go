/*
 * CoinBase API
 *
 * Coinbase provides a simple and powerful REST API to integrate bitcoin, bitcoin cash, litecoin and ethereum payments into your business or application.  This API reference provides information on available endpoints and how to interact with it. To read more about the API, visit our API documentation.
 *
 * API version: 2019-11-15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package coinbase

import (
	"encoding/json"
)

// InvalidTokenErrors struct for InvalidTokenErrors
type InvalidTokenErrors struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

// NewInvalidTokenErrors instantiates a new InvalidTokenErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidTokenErrors(id string, message string) *InvalidTokenErrors {
	this := InvalidTokenErrors{}
	this.Id = id
	this.Message = message
	return &this
}

// NewInvalidTokenErrorsWithDefaults instantiates a new InvalidTokenErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidTokenErrorsWithDefaults() *InvalidTokenErrors {
	this := InvalidTokenErrors{}
	return &this
}

// GetId returns the Id field value
func (o *InvalidTokenErrors) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvalidTokenErrors) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvalidTokenErrors) SetId(v string) {
	o.Id = v
}

// GetMessage returns the Message field value
func (o *InvalidTokenErrors) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvalidTokenErrors) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvalidTokenErrors) SetMessage(v string) {
	o.Message = v
}

func (o InvalidTokenErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableInvalidTokenErrors struct {
	value *InvalidTokenErrors
	isSet bool
}

func (v NullableInvalidTokenErrors) Get() *InvalidTokenErrors {
	return v.value
}

func (v *NullableInvalidTokenErrors) Set(val *InvalidTokenErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidTokenErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidTokenErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidTokenErrors(val *InvalidTokenErrors) *NullableInvalidTokenErrors {
	return &NullableInvalidTokenErrors{value: val, isSet: true}
}

func (v NullableInvalidTokenErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidTokenErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
