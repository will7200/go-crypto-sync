# InlineResponse20021

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Currency ID | [optional] 
**Predictions** | Pointer to [**[]CurrenciesPredictionsTickerPredictions**](_currencies_predictions_ticker_predictions.md) |  | [optional] 

## Methods

### NewInlineResponse20021

`func NewInlineResponse20021() *InlineResponse20021`

NewInlineResponse20021 instantiates a new InlineResponse20021 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20021WithDefaults

`func NewInlineResponse20021WithDefaults() *InlineResponse20021`

NewInlineResponse20021WithDefaults instantiates a new InlineResponse20021 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20021) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20021) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20021) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20021) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPredictions

`func (o *InlineResponse20021) GetPredictions() []CurrenciesPredictionsTickerPredictions`

GetPredictions returns the Predictions field if non-nil, zero value otherwise.

### GetPredictionsOk

`func (o *InlineResponse20021) GetPredictionsOk() (*[]CurrenciesPredictionsTickerPredictions, bool)`

GetPredictionsOk returns a tuple with the Predictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictions

`func (o *InlineResponse20021) SetPredictions(v []CurrenciesPredictionsTickerPredictions)`

SetPredictions sets Predictions field to given value.

### HasPredictions

`func (o *InlineResponse20021) HasPredictions() bool`

HasPredictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


