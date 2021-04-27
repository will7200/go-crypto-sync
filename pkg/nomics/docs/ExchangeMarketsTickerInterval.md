# ExchangeMarketsTickerInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | Pointer to **string** | Volume in USD | [optional] 
**VolumeBase** | Pointer to **string** | Volume in the base currency | [optional] 
**VolumeChange** | Pointer to **string** | Change of the volume in USD | [optional] 
**VolumeBaseChange** | Pointer to **string** | Change of the volume in the base currency | [optional] 
**Trades** | Pointer to **string** | Number of trades | [optional] 
**TradesChange** | Pointer to **string** | Change in the number of trades | [optional] 
**PriceChange** | Pointer to **string** | Change in the price in USD | [optional] 
**PriceQuoteChange** | Pointer to **string** | Change in the price in the quote currency | [optional] 

## Methods

### NewExchangeMarketsTickerInterval

`func NewExchangeMarketsTickerInterval() *ExchangeMarketsTickerInterval`

NewExchangeMarketsTickerInterval instantiates a new ExchangeMarketsTickerInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeMarketsTickerIntervalWithDefaults

`func NewExchangeMarketsTickerIntervalWithDefaults() *ExchangeMarketsTickerInterval`

NewExchangeMarketsTickerIntervalWithDefaults instantiates a new ExchangeMarketsTickerInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *ExchangeMarketsTickerInterval) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ExchangeMarketsTickerInterval) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ExchangeMarketsTickerInterval) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ExchangeMarketsTickerInterval) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetVolumeBase

`func (o *ExchangeMarketsTickerInterval) GetVolumeBase() string`

GetVolumeBase returns the VolumeBase field if non-nil, zero value otherwise.

### GetVolumeBaseOk

`func (o *ExchangeMarketsTickerInterval) GetVolumeBaseOk() (*string, bool)`

GetVolumeBaseOk returns a tuple with the VolumeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBase

`func (o *ExchangeMarketsTickerInterval) SetVolumeBase(v string)`

SetVolumeBase sets VolumeBase field to given value.

### HasVolumeBase

`func (o *ExchangeMarketsTickerInterval) HasVolumeBase() bool`

HasVolumeBase returns a boolean if a field has been set.

### GetVolumeChange

`func (o *ExchangeMarketsTickerInterval) GetVolumeChange() string`

GetVolumeChange returns the VolumeChange field if non-nil, zero value otherwise.

### GetVolumeChangeOk

`func (o *ExchangeMarketsTickerInterval) GetVolumeChangeOk() (*string, bool)`

GetVolumeChangeOk returns a tuple with the VolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeChange

`func (o *ExchangeMarketsTickerInterval) SetVolumeChange(v string)`

SetVolumeChange sets VolumeChange field to given value.

### HasVolumeChange

`func (o *ExchangeMarketsTickerInterval) HasVolumeChange() bool`

HasVolumeChange returns a boolean if a field has been set.

### GetVolumeBaseChange

`func (o *ExchangeMarketsTickerInterval) GetVolumeBaseChange() string`

GetVolumeBaseChange returns the VolumeBaseChange field if non-nil, zero value otherwise.

### GetVolumeBaseChangeOk

`func (o *ExchangeMarketsTickerInterval) GetVolumeBaseChangeOk() (*string, bool)`

GetVolumeBaseChangeOk returns a tuple with the VolumeBaseChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBaseChange

`func (o *ExchangeMarketsTickerInterval) SetVolumeBaseChange(v string)`

SetVolumeBaseChange sets VolumeBaseChange field to given value.

### HasVolumeBaseChange

`func (o *ExchangeMarketsTickerInterval) HasVolumeBaseChange() bool`

HasVolumeBaseChange returns a boolean if a field has been set.

### GetTrades

`func (o *ExchangeMarketsTickerInterval) GetTrades() string`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *ExchangeMarketsTickerInterval) GetTradesOk() (*string, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *ExchangeMarketsTickerInterval) SetTrades(v string)`

SetTrades sets Trades field to given value.

### HasTrades

`func (o *ExchangeMarketsTickerInterval) HasTrades() bool`

HasTrades returns a boolean if a field has been set.

### GetTradesChange

`func (o *ExchangeMarketsTickerInterval) GetTradesChange() string`

GetTradesChange returns the TradesChange field if non-nil, zero value otherwise.

### GetTradesChangeOk

`func (o *ExchangeMarketsTickerInterval) GetTradesChangeOk() (*string, bool)`

GetTradesChangeOk returns a tuple with the TradesChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradesChange

`func (o *ExchangeMarketsTickerInterval) SetTradesChange(v string)`

SetTradesChange sets TradesChange field to given value.

### HasTradesChange

`func (o *ExchangeMarketsTickerInterval) HasTradesChange() bool`

HasTradesChange returns a boolean if a field has been set.

### GetPriceChange

`func (o *ExchangeMarketsTickerInterval) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *ExchangeMarketsTickerInterval) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *ExchangeMarketsTickerInterval) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *ExchangeMarketsTickerInterval) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetPriceQuoteChange

`func (o *ExchangeMarketsTickerInterval) GetPriceQuoteChange() string`

GetPriceQuoteChange returns the PriceQuoteChange field if non-nil, zero value otherwise.

### GetPriceQuoteChangeOk

`func (o *ExchangeMarketsTickerInterval) GetPriceQuoteChangeOk() (*string, bool)`

GetPriceQuoteChangeOk returns a tuple with the PriceQuoteChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceQuoteChange

`func (o *ExchangeMarketsTickerInterval) SetPriceQuoteChange(v string)`

SetPriceQuoteChange sets PriceQuoteChange field to given value.

### HasPriceQuoteChange

`func (o *ExchangeMarketsTickerInterval) HasPriceQuoteChange() bool`

HasPriceQuoteChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


