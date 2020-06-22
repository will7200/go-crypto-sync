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

// Pagination struct for Pagination
type Pagination struct {
	EndingBefore  *string  `json:"ending_before,omitempty"`
	StartingAfter *string  `json:"starting_after,omitempty"`
	Limit         *float32 `json:"limit,omitempty"`
	Order         *string  `json:"order,omitempty"`
	PreviousUri   *string  `json:"previous_uri,omitempty"`
	NextUri       *string  `json:"next_uri,omitempty"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination() *Pagination {
	this := Pagination{}
	var limit float32 = 25
	this.Limit = &limit
	var order string = "desc"
	this.Order = &order
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	var limit float32 = 25
	this.Limit = &limit
	var order string = "desc"
	this.Order = &order
	return &this
}

// GetEndingBefore returns the EndingBefore field value if set, zero value otherwise.
func (o *Pagination) GetEndingBefore() string {
	if o == nil || o.EndingBefore == nil {
		var ret string
		return ret
	}
	return *o.EndingBefore
}

// GetEndingBeforeOk returns a tuple with the EndingBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetEndingBeforeOk() (*string, bool) {
	if o == nil || o.EndingBefore == nil {
		return nil, false
	}
	return o.EndingBefore, true
}

// HasEndingBefore returns a boolean if a field has been set.
func (o *Pagination) HasEndingBefore() bool {
	if o != nil && o.EndingBefore != nil {
		return true
	}

	return false
}

// SetEndingBefore gets a reference to the given string and assigns it to the EndingBefore field.
func (o *Pagination) SetEndingBefore(v string) {
	o.EndingBefore = &v
}

// GetStartingAfter returns the StartingAfter field value if set, zero value otherwise.
func (o *Pagination) GetStartingAfter() string {
	if o == nil || o.StartingAfter == nil {
		var ret string
		return ret
	}
	return *o.StartingAfter
}

// GetStartingAfterOk returns a tuple with the StartingAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetStartingAfterOk() (*string, bool) {
	if o == nil || o.StartingAfter == nil {
		return nil, false
	}
	return o.StartingAfter, true
}

// HasStartingAfter returns a boolean if a field has been set.
func (o *Pagination) HasStartingAfter() bool {
	if o != nil && o.StartingAfter != nil {
		return true
	}

	return false
}

// SetStartingAfter gets a reference to the given string and assigns it to the StartingAfter field.
func (o *Pagination) SetStartingAfter(v string) {
	o.StartingAfter = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *Pagination) GetLimit() float32 {
	if o == nil || o.Limit == nil {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetLimitOk() (*float32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Pagination) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *Pagination) SetLimit(v float32) {
	o.Limit = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Pagination) GetOrder() string {
	if o == nil || o.Order == nil {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetOrderOk() (*string, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Pagination) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *Pagination) SetOrder(v string) {
	o.Order = &v
}

// GetPreviousUri returns the PreviousUri field value if set, zero value otherwise.
func (o *Pagination) GetPreviousUri() string {
	if o == nil || o.PreviousUri == nil {
		var ret string
		return ret
	}
	return *o.PreviousUri
}

// GetPreviousUriOk returns a tuple with the PreviousUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetPreviousUriOk() (*string, bool) {
	if o == nil || o.PreviousUri == nil {
		return nil, false
	}
	return o.PreviousUri, true
}

// HasPreviousUri returns a boolean if a field has been set.
func (o *Pagination) HasPreviousUri() bool {
	if o != nil && o.PreviousUri != nil {
		return true
	}

	return false
}

// SetPreviousUri gets a reference to the given string and assigns it to the PreviousUri field.
func (o *Pagination) SetPreviousUri(v string) {
	o.PreviousUri = &v
}

// GetNextUri returns the NextUri field value if set, zero value otherwise.
func (o *Pagination) GetNextUri() string {
	if o == nil || o.NextUri == nil {
		var ret string
		return ret
	}
	return *o.NextUri
}

// GetNextUriOk returns a tuple with the NextUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetNextUriOk() (*string, bool) {
	if o == nil || o.NextUri == nil {
		return nil, false
	}
	return o.NextUri, true
}

// HasNextUri returns a boolean if a field has been set.
func (o *Pagination) HasNextUri() bool {
	if o != nil && o.NextUri != nil {
		return true
	}

	return false
}

// SetNextUri gets a reference to the given string and assigns it to the NextUri field.
func (o *Pagination) SetNextUri(v string) {
	o.NextUri = &v
}

func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndingBefore != nil {
		toSerialize["ending_before"] = o.EndingBefore
	}
	if o.StartingAfter != nil {
		toSerialize["starting_after"] = o.StartingAfter
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.PreviousUri != nil {
		toSerialize["previous_uri"] = o.PreviousUri
	}
	if o.NextUri != nil {
		toSerialize["next_uri"] = o.NextUri
	}
	return json.Marshal(toSerialize)
}

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}