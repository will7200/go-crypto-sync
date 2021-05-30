# AcquireMiningDifficultyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Difficulty** | Pointer to **string** | Mining difficulty (CET/Hour) | [optional] 
**Prediction** | Pointer to **string** | EST. hourly mining yield to distribute | [optional] 
**UpdateTime** | Pointer to **int64** | The time for an update on parameter \&quot;prediction\&quot; | [optional] 

## Methods

### NewAcquireMiningDifficultyResponse

`func NewAcquireMiningDifficultyResponse() *AcquireMiningDifficultyResponse`

NewAcquireMiningDifficultyResponse instantiates a new AcquireMiningDifficultyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireMiningDifficultyResponseWithDefaults

`func NewAcquireMiningDifficultyResponseWithDefaults() *AcquireMiningDifficultyResponse`

NewAcquireMiningDifficultyResponseWithDefaults instantiates a new AcquireMiningDifficultyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifficulty

`func (o *AcquireMiningDifficultyResponse) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *AcquireMiningDifficultyResponse) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *AcquireMiningDifficultyResponse) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *AcquireMiningDifficultyResponse) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetPrediction

`func (o *AcquireMiningDifficultyResponse) GetPrediction() string`

GetPrediction returns the Prediction field if non-nil, zero value otherwise.

### GetPredictionOk

`func (o *AcquireMiningDifficultyResponse) GetPredictionOk() (*string, bool)`

GetPredictionOk returns a tuple with the Prediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrediction

`func (o *AcquireMiningDifficultyResponse) SetPrediction(v string)`

SetPrediction sets Prediction field to given value.

### HasPrediction

`func (o *AcquireMiningDifficultyResponse) HasPrediction() bool`

HasPrediction returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AcquireMiningDifficultyResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AcquireMiningDifficultyResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AcquireMiningDifficultyResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AcquireMiningDifficultyResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


