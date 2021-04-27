# InlineResponse20022

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Currency ID | [optional] 
**Interval** | Pointer to **string** | Prediction interval | [optional] 
**Predictions** | Pointer to [**[]InlineResponse20022Predictions**](inline_response_200_22_predictions.md) |  | [optional] 

## Methods

### NewInlineResponse20022

`func NewInlineResponse20022() *InlineResponse20022`

NewInlineResponse20022 instantiates a new InlineResponse20022 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20022WithDefaults

`func NewInlineResponse20022WithDefaults() *InlineResponse20022`

NewInlineResponse20022WithDefaults instantiates a new InlineResponse20022 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20022) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20022) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20022) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20022) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *InlineResponse20022) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InlineResponse20022) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InlineResponse20022) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InlineResponse20022) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPredictions

`func (o *InlineResponse20022) GetPredictions() []InlineResponse20022Predictions`

GetPredictions returns the Predictions field if non-nil, zero value otherwise.

### GetPredictionsOk

`func (o *InlineResponse20022) GetPredictionsOk() (*[]InlineResponse20022Predictions, bool)`

GetPredictionsOk returns a tuple with the Predictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictions

`func (o *InlineResponse20022) SetPredictions(v []InlineResponse20022Predictions)`

SetPredictions sets Predictions field to given value.

### HasPredictions

`func (o *InlineResponse20022) HasPredictions() bool`

HasPredictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


