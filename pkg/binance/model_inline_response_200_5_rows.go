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
// InlineResponse2005Rows struct for InlineResponse2005Rows
type InlineResponse2005Rows struct {
	Amount    *string `json:"amount,omitempty"`
	Asset     *string `json:"asset,omitempty"`
	Interest  *string `json:"interest,omitempty"`
	Principal *string `json:"principal,omitempty"`
	Status    *string `json:"status,omitempty"`
	Timestamp *int32  `json:"timestamp,omitempty"`
	TxId      *int32  `json:"txId,omitempty"`
}

// NewInlineResponse2005Rows instantiates a new InlineResponse2005Rows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2005Rows() *InlineResponse2005Rows {
	this := InlineResponse2005Rows{}
	return &this
}

// NewInlineResponse2005RowsWithDefaults instantiates a new InlineResponse2005Rows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2005RowsWithDefaults() *InlineResponse2005Rows {
	this := InlineResponse2005Rows{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *InlineResponse2005Rows) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005Rows) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *InlineResponse2005Rows) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *InlineResponse2005Rows) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *InlineResponse2005Rows) GetAsset() string {
	if o == nil || o.Asset == nil {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005Rows) GetAssetOk() (*string, bool) {
	if o == nil || o.Asset == nil {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *InlineResponse2005Rows) HasAsset() bool {
	if o != nil && o.Asset != nil {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *InlineResponse2005Rows) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *InlineResponse2005Rows) GetInterest() string {
	if o == nil || o.Interest == nil {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005Rows) GetInterestOk() (*string, bool) {
	if o == nil || o.Interest == nil {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *InlineResponse2005Rows) HasInterest() bool {
	if o != nil && o.Interest != nil {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *InlineResponse2005Rows) SetInterest(v string) {
	o.Interest = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *InlineResponse2005Rows) GetPrincipal() string {
	if o == nil || o.Principal == nil {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005Rows) GetPrincipalOk() (*string, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *InlineResponse2005Rows) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *InlineResponse2005Rows) SetPrincipal(v string) {
	o.Principal = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2005Rows) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005Rows) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2005Rows) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2005Rows) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *InlineResponse2005Rows) GetTimestamp() int32 {
	if o == nil || o.Timestamp == nil {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005Rows) GetTimestampOk() (*int32, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *InlineResponse2005Rows) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *InlineResponse2005Rows) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *InlineResponse2005Rows) GetTxId() int32 {
	if o == nil || o.TxId == nil {
		var ret int32
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005Rows) GetTxIdOk() (*int32, bool) {
	if o == nil || o.TxId == nil {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *InlineResponse2005Rows) HasTxId() bool {
	if o != nil && o.TxId != nil {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int32 and assigns it to the TxId field.
func (o *InlineResponse2005Rows) SetTxId(v int32) {
	o.TxId = &v
}

func (o InlineResponse2005Rows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Asset != nil {
		toSerialize["asset"] = o.Asset
	}
	if o.Interest != nil {
		toSerialize["interest"] = o.Interest
	}
	if o.Principal != nil {
		toSerialize["principal"] = o.Principal
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TxId != nil {
		toSerialize["txId"] = o.TxId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2005Rows struct {
	value *InlineResponse2005Rows
	isSet bool
}

func (v NullableInlineResponse2005Rows) Get() *InlineResponse2005Rows {
	return v.value
}

func (v *NullableInlineResponse2005Rows) Set(val *InlineResponse2005Rows) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2005Rows) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2005Rows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2005Rows(val *InlineResponse2005Rows) *NullableInlineResponse2005Rows {
	return &NullableInlineResponse2005Rows{value: val, isSet: true}
}

func (v NullableInlineResponse2005Rows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2005Rows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}