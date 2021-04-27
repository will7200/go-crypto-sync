# ExchangeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency ID | [optional] 
**Rate** | Pointer to **string** | Exchange rate to USD (Currency * Rate &#x3D; USD) | [optional] 
**Timestamp** | Pointer to **string** | Timestamp this rate was updated at in RFC3339 | [optional] 

## Methods

### NewExchangeRate

`func NewExchangeRate() *ExchangeRate`

NewExchangeRate instantiates a new ExchangeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRateWithDefaults

`func NewExchangeRateWithDefaults() *ExchangeRate`

NewExchangeRateWithDefaults instantiates a new ExchangeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ExchangeRate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExchangeRate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExchangeRate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ExchangeRate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRate

`func (o *ExchangeRate) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ExchangeRate) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ExchangeRate) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ExchangeRate) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTimestamp

`func (o *ExchangeRate) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExchangeRate) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExchangeRate) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExchangeRate) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


