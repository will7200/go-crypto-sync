# MarketCapSparklineRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency ID | [optional] 
**Timestamps** | Pointer to **[]string** | Time values matching the close values | [optional] 
**Closes** | Pointer to **[]string** | Closing market cap in USD corresponding to timestamp value | [optional] 

## Methods

### NewMarketCapSparklineRow

`func NewMarketCapSparklineRow() *MarketCapSparklineRow`

NewMarketCapSparklineRow instantiates a new MarketCapSparklineRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketCapSparklineRowWithDefaults

`func NewMarketCapSparklineRowWithDefaults() *MarketCapSparklineRow`

NewMarketCapSparklineRowWithDefaults instantiates a new MarketCapSparklineRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *MarketCapSparklineRow) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketCapSparklineRow) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketCapSparklineRow) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketCapSparklineRow) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTimestamps

`func (o *MarketCapSparklineRow) GetTimestamps() []string`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *MarketCapSparklineRow) GetTimestampsOk() (*[]string, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *MarketCapSparklineRow) SetTimestamps(v []string)`

SetTimestamps sets Timestamps field to given value.

### HasTimestamps

`func (o *MarketCapSparklineRow) HasTimestamps() bool`

HasTimestamps returns a boolean if a field has been set.

### GetCloses

`func (o *MarketCapSparklineRow) GetCloses() []string`

GetCloses returns the Closes field if non-nil, zero value otherwise.

### GetClosesOk

`func (o *MarketCapSparklineRow) GetClosesOk() (*[]string, bool)`

GetClosesOk returns a tuple with the Closes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloses

`func (o *MarketCapSparklineRow) SetCloses(v []string)`

SetCloses sets Closes field to given value.

### HasCloses

`func (o *MarketCapSparklineRow) HasCloses() bool`

HasCloses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


