# SparklineRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency ID | [optional] 
**Timestamps** | Pointer to **[]string** | Time values matching the close values | [optional] 
**Closes** | Pointer to **[]string** | Closing price in USD corresponding to timestamp value | [optional] 

## Methods

### NewSparklineRow

`func NewSparklineRow() *SparklineRow`

NewSparklineRow instantiates a new SparklineRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSparklineRowWithDefaults

`func NewSparklineRowWithDefaults() *SparklineRow`

NewSparklineRowWithDefaults instantiates a new SparklineRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *SparklineRow) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SparklineRow) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SparklineRow) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SparklineRow) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTimestamps

`func (o *SparklineRow) GetTimestamps() []string`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *SparklineRow) GetTimestampsOk() (*[]string, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *SparklineRow) SetTimestamps(v []string)`

SetTimestamps sets Timestamps field to given value.

### HasTimestamps

`func (o *SparklineRow) HasTimestamps() bool`

HasTimestamps returns a boolean if a field has been set.

### GetCloses

`func (o *SparklineRow) GetCloses() []string`

GetCloses returns the Closes field if non-nil, zero value otherwise.

### GetClosesOk

`func (o *SparklineRow) GetClosesOk() (*[]string, bool)`

GetClosesOk returns a tuple with the Closes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloses

`func (o *SparklineRow) SetCloses(v []string)`

SetCloses sets Closes field to given value.

### HasCloses

`func (o *SparklineRow) HasCloses() bool`

HasCloses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


