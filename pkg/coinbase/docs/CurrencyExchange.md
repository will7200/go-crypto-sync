# CurrencyExchange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | 
**Currency** | Pointer to **string** |  | 
**Base** | Pointer to **string** |  | 

## Methods

### NewCurrencyExchange

`func NewCurrencyExchange(amount string, currency string, base string, ) *CurrencyExchange`

NewCurrencyExchange instantiates a new CurrencyExchange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyExchangeWithDefaults

`func NewCurrencyExchangeWithDefaults() *CurrencyExchange`

NewCurrencyExchangeWithDefaults instantiates a new CurrencyExchange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CurrencyExchange) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CurrencyExchange) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CurrencyExchange) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CurrencyExchange) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CurrencyExchange) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CurrencyExchange) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBase

`func (o *CurrencyExchange) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *CurrencyExchange) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *CurrencyExchange) SetBase(v string)`

SetBase sets Base field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


