# ExchangeRates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | 
**Rates** | Pointer to **map[string]string** |  | 

## Methods

### NewExchangeRates

`func NewExchangeRates(currency string, rates map[string]string, ) *ExchangeRates`

NewExchangeRates instantiates a new ExchangeRates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRatesWithDefaults

`func NewExchangeRatesWithDefaults() *ExchangeRates`

NewExchangeRatesWithDefaults instantiates a new ExchangeRates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ExchangeRates) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExchangeRates) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExchangeRates) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRates

`func (o *ExchangeRates) GetRates() map[string]string`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *ExchangeRates) GetRatesOk() (*map[string]string, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *ExchangeRates) SetRates(v map[string]string)`

SetRates sets Rates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


