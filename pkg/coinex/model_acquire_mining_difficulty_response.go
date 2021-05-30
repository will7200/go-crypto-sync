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

// AcquireMiningDifficultyResponse AcquireMiningDifficulty Response Value
type AcquireMiningDifficultyResponse struct {
	// Mining difficulty (CET/Hour)
	Difficulty *string `json:"difficulty,omitempty"`
	// EST. hourly mining yield to distribute
	Prediction *string `json:"prediction,omitempty"`
	// The time for an update on parameter \"prediction\"
	UpdateTime *int64 `json:"update_time,omitempty"`
}

// NewAcquireMiningDifficultyResponse instantiates a new AcquireMiningDifficultyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcquireMiningDifficultyResponse() *AcquireMiningDifficultyResponse {
	this := AcquireMiningDifficultyResponse{}
	return &this
}

// NewAcquireMiningDifficultyResponseWithDefaults instantiates a new AcquireMiningDifficultyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcquireMiningDifficultyResponseWithDefaults() *AcquireMiningDifficultyResponse {
	this := AcquireMiningDifficultyResponse{}
	return &this
}

// GetDifficulty returns the Difficulty field value if set, zero value otherwise.
func (o *AcquireMiningDifficultyResponse) GetDifficulty() string {
	if o == nil || o.Difficulty == nil {
		var ret string
		return ret
	}
	return *o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireMiningDifficultyResponse) GetDifficultyOk() (*string, bool) {
	if o == nil || o.Difficulty == nil {
		return nil, false
	}
	return o.Difficulty, true
}

// HasDifficulty returns a boolean if a field has been set.
func (o *AcquireMiningDifficultyResponse) HasDifficulty() bool {
	if o != nil && o.Difficulty != nil {
		return true
	}

	return false
}

// SetDifficulty gets a reference to the given string and assigns it to the Difficulty field.
func (o *AcquireMiningDifficultyResponse) SetDifficulty(v string) {
	o.Difficulty = &v
}

// GetPrediction returns the Prediction field value if set, zero value otherwise.
func (o *AcquireMiningDifficultyResponse) GetPrediction() string {
	if o == nil || o.Prediction == nil {
		var ret string
		return ret
	}
	return *o.Prediction
}

// GetPredictionOk returns a tuple with the Prediction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireMiningDifficultyResponse) GetPredictionOk() (*string, bool) {
	if o == nil || o.Prediction == nil {
		return nil, false
	}
	return o.Prediction, true
}

// HasPrediction returns a boolean if a field has been set.
func (o *AcquireMiningDifficultyResponse) HasPrediction() bool {
	if o != nil && o.Prediction != nil {
		return true
	}

	return false
}

// SetPrediction gets a reference to the given string and assigns it to the Prediction field.
func (o *AcquireMiningDifficultyResponse) SetPrediction(v string) {
	o.Prediction = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AcquireMiningDifficultyResponse) GetUpdateTime() int64 {
	if o == nil || o.UpdateTime == nil {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquireMiningDifficultyResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AcquireMiningDifficultyResponse) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AcquireMiningDifficultyResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o AcquireMiningDifficultyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Difficulty != nil {
		toSerialize["difficulty"] = o.Difficulty
	}
	if o.Prediction != nil {
		toSerialize["prediction"] = o.Prediction
	}
	if o.UpdateTime != nil {
		toSerialize["update_time"] = o.UpdateTime
	}
	return json.Marshal(toSerialize)
}

type NullableAcquireMiningDifficultyResponse struct {
	value *AcquireMiningDifficultyResponse
	isSet bool
}

func (v NullableAcquireMiningDifficultyResponse) Get() *AcquireMiningDifficultyResponse {
	return v.value
}

func (v *NullableAcquireMiningDifficultyResponse) Set(val *AcquireMiningDifficultyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcquireMiningDifficultyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcquireMiningDifficultyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcquireMiningDifficultyResponse(val *AcquireMiningDifficultyResponse) *NullableAcquireMiningDifficultyResponse {
	return &NullableAcquireMiningDifficultyResponse{value: val, isSet: true}
}

func (v NullableAcquireMiningDifficultyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcquireMiningDifficultyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
