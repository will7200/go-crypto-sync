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
// InlineResponse20012UserAssets struct for InlineResponse20012UserAssets
type InlineResponse20012UserAssets struct {
	Asset    *string `json:"asset,omitempty"`
	Borrowed *string `json:"borrowed,omitempty"`
	Free     *string `json:"free,omitempty"`
	Interest *string `json:"interest,omitempty"`
	Locked   *string `json:"locked,omitempty"`
	NetAsset *string `json:"netAsset,omitempty"`
}

// NewInlineResponse20012UserAssets instantiates a new InlineResponse20012UserAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20012UserAssets() *InlineResponse20012UserAssets {
	this := InlineResponse20012UserAssets{}
	return &this
}

// NewInlineResponse20012UserAssetsWithDefaults instantiates a new InlineResponse20012UserAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20012UserAssetsWithDefaults() *InlineResponse20012UserAssets {
	this := InlineResponse20012UserAssets{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *InlineResponse20012UserAssets) GetAsset() string {
	if o == nil || o.Asset == nil {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012UserAssets) GetAssetOk() (*string, bool) {
	if o == nil || o.Asset == nil {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *InlineResponse20012UserAssets) HasAsset() bool {
	if o != nil && o.Asset != nil {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *InlineResponse20012UserAssets) SetAsset(v string) {
	o.Asset = &v
}

// GetBorrowed returns the Borrowed field value if set, zero value otherwise.
func (o *InlineResponse20012UserAssets) GetBorrowed() string {
	if o == nil || o.Borrowed == nil {
		var ret string
		return ret
	}
	return *o.Borrowed
}

// GetBorrowedOk returns a tuple with the Borrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012UserAssets) GetBorrowedOk() (*string, bool) {
	if o == nil || o.Borrowed == nil {
		return nil, false
	}
	return o.Borrowed, true
}

// HasBorrowed returns a boolean if a field has been set.
func (o *InlineResponse20012UserAssets) HasBorrowed() bool {
	if o != nil && o.Borrowed != nil {
		return true
	}

	return false
}

// SetBorrowed gets a reference to the given string and assigns it to the Borrowed field.
func (o *InlineResponse20012UserAssets) SetBorrowed(v string) {
	o.Borrowed = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *InlineResponse20012UserAssets) GetFree() string {
	if o == nil || o.Free == nil {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012UserAssets) GetFreeOk() (*string, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *InlineResponse20012UserAssets) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *InlineResponse20012UserAssets) SetFree(v string) {
	o.Free = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *InlineResponse20012UserAssets) GetInterest() string {
	if o == nil || o.Interest == nil {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012UserAssets) GetInterestOk() (*string, bool) {
	if o == nil || o.Interest == nil {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *InlineResponse20012UserAssets) HasInterest() bool {
	if o != nil && o.Interest != nil {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *InlineResponse20012UserAssets) SetInterest(v string) {
	o.Interest = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *InlineResponse20012UserAssets) GetLocked() string {
	if o == nil || o.Locked == nil {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012UserAssets) GetLockedOk() (*string, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *InlineResponse20012UserAssets) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *InlineResponse20012UserAssets) SetLocked(v string) {
	o.Locked = &v
}

// GetNetAsset returns the NetAsset field value if set, zero value otherwise.
func (o *InlineResponse20012UserAssets) GetNetAsset() string {
	if o == nil || o.NetAsset == nil {
		var ret string
		return ret
	}
	return *o.NetAsset
}

// GetNetAssetOk returns a tuple with the NetAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012UserAssets) GetNetAssetOk() (*string, bool) {
	if o == nil || o.NetAsset == nil {
		return nil, false
	}
	return o.NetAsset, true
}

// HasNetAsset returns a boolean if a field has been set.
func (o *InlineResponse20012UserAssets) HasNetAsset() bool {
	if o != nil && o.NetAsset != nil {
		return true
	}

	return false
}

// SetNetAsset gets a reference to the given string and assigns it to the NetAsset field.
func (o *InlineResponse20012UserAssets) SetNetAsset(v string) {
	o.NetAsset = &v
}

func (o InlineResponse20012UserAssets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Asset != nil {
		toSerialize["asset"] = o.Asset
	}
	if o.Borrowed != nil {
		toSerialize["borrowed"] = o.Borrowed
	}
	if o.Free != nil {
		toSerialize["free"] = o.Free
	}
	if o.Interest != nil {
		toSerialize["interest"] = o.Interest
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.NetAsset != nil {
		toSerialize["netAsset"] = o.NetAsset
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20012UserAssets struct {
	value *InlineResponse20012UserAssets
	isSet bool
}

func (v NullableInlineResponse20012UserAssets) Get() *InlineResponse20012UserAssets {
	return v.value
}

func (v *NullableInlineResponse20012UserAssets) Set(val *InlineResponse20012UserAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20012UserAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20012UserAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20012UserAssets(val *InlineResponse20012UserAssets) *NullableInlineResponse20012UserAssets {
	return &NullableInlineResponse20012UserAssets{value: val, isSet: true}
}

func (v NullableInlineResponse20012UserAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20012UserAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
