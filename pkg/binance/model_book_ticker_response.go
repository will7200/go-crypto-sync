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
	"fmt"
)

//
// BookTickerResponse - struct for BookTickerResponse
type BookTickerResponse struct {
	BookTicker     *BookTicker
	BookTickerList *BookTickerList
}

// BookTickerAsBookTickerResponse is a convenience function that returns BookTicker wrapped in BookTickerResponse
func BookTickerAsBookTickerResponse(v *BookTicker) BookTickerResponse {
	return BookTickerResponse{BookTicker: v}
}

// BookTickerListAsBookTickerResponse is a convenience function that returns BookTickerList wrapped in BookTickerResponse
func BookTickerListAsBookTickerResponse(v *BookTickerList) BookTickerResponse {
	return BookTickerResponse{BookTickerList: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BookTickerResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BookTicker
	err = json.Unmarshal(data, &dst.BookTicker)
	if err == nil {
		jsonBookTicker, _ := json.Marshal(dst.BookTicker)
		if string(jsonBookTicker) == "{}" { // empty struct
			dst.BookTicker = nil
		} else {
			match++
		}
	} else {
		dst.BookTicker = nil
	}

	// try to unmarshal data into BookTickerList
	err = json.Unmarshal(data, &dst.BookTickerList)
	if err == nil {
		jsonBookTickerList, _ := json.Marshal(dst.BookTickerList)
		if string(jsonBookTickerList) == "{}" { // empty struct
			dst.BookTickerList = nil
		} else {
			match++
		}
	} else {
		dst.BookTickerList = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BookTicker = nil
		dst.BookTickerList = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(BookTickerResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(BookTickerResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BookTickerResponse) MarshalJSON() ([]byte, error) {
	if src.BookTicker != nil {
		return json.Marshal(&src.BookTicker)
	}

	if src.BookTickerList != nil {
		return json.Marshal(&src.BookTickerList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BookTickerResponse) GetActualInstance() interface{} {
	if obj.BookTicker != nil {
		return obj.BookTicker
	}

	if obj.BookTickerList != nil {
		return obj.BookTickerList
	}

	// all schemas are nil
	return nil
}

type NullableBookTickerResponse struct {
	value *BookTickerResponse
	isSet bool
}

func (v NullableBookTickerResponse) Get() *BookTickerResponse {
	return v.value
}

func (v *NullableBookTickerResponse) Set(val *BookTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBookTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBookTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookTickerResponse(val *BookTickerResponse) *NullableBookTickerResponse {
	return &NullableBookTickerResponse{value: val, isSet: true}
}

func (v NullableBookTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}