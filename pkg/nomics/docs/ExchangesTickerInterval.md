# ExchangesTickerInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | Pointer to **string** | Total volume for the given interval in USD | [optional] 
**VolumeChange** | Pointer to **string** | Amount of volume change for the given interval in USD | [optional] 
**VolumeChangePct** | Pointer to **string** | Percent change of volume for the given interval | [optional] 
**SpotVolume** | Pointer to **string** | Total spot market volume for the given interval in USD | [optional] 
**SpotVolumeChange** | Pointer to **string** | Amount of spot market volume change for the given interval in USD | [optional] 
**SpotVolumeChangePct** | Pointer to **string** | Percent change of spot market volume for the given interval | [optional] 
**DerivativeVolume** | Pointer to **string** | Total derivative market volume for the given interval in USD | [optional] 
**DerivativeVolumeChange** | Pointer to **string** | Amount of derivative market volume change for the given interval in USD | [optional] 
**DerivativeVolumeChangePct** | Pointer to **string** | Percent change of derivative market volume for the given interval | [optional] 
**Trades** | Pointer to **string** | Total trades for the given interval in USD | [optional] 
**TradesChange** | Pointer to **string** | Amount of trades change for the given interval in USD | [optional] 
**TradesChangePct** | Pointer to **string** | Percent change of trades for the given interval | [optional] 

## Methods

### NewExchangesTickerInterval

`func NewExchangesTickerInterval() *ExchangesTickerInterval`

NewExchangesTickerInterval instantiates a new ExchangesTickerInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangesTickerIntervalWithDefaults

`func NewExchangesTickerIntervalWithDefaults() *ExchangesTickerInterval`

NewExchangesTickerIntervalWithDefaults instantiates a new ExchangesTickerInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *ExchangesTickerInterval) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ExchangesTickerInterval) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ExchangesTickerInterval) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ExchangesTickerInterval) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetVolumeChange

`func (o *ExchangesTickerInterval) GetVolumeChange() string`

GetVolumeChange returns the VolumeChange field if non-nil, zero value otherwise.

### GetVolumeChangeOk

`func (o *ExchangesTickerInterval) GetVolumeChangeOk() (*string, bool)`

GetVolumeChangeOk returns a tuple with the VolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeChange

`func (o *ExchangesTickerInterval) SetVolumeChange(v string)`

SetVolumeChange sets VolumeChange field to given value.

### HasVolumeChange

`func (o *ExchangesTickerInterval) HasVolumeChange() bool`

HasVolumeChange returns a boolean if a field has been set.

### GetVolumeChangePct

`func (o *ExchangesTickerInterval) GetVolumeChangePct() string`

GetVolumeChangePct returns the VolumeChangePct field if non-nil, zero value otherwise.

### GetVolumeChangePctOk

`func (o *ExchangesTickerInterval) GetVolumeChangePctOk() (*string, bool)`

GetVolumeChangePctOk returns a tuple with the VolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeChangePct

`func (o *ExchangesTickerInterval) SetVolumeChangePct(v string)`

SetVolumeChangePct sets VolumeChangePct field to given value.

### HasVolumeChangePct

`func (o *ExchangesTickerInterval) HasVolumeChangePct() bool`

HasVolumeChangePct returns a boolean if a field has been set.

### GetSpotVolume

`func (o *ExchangesTickerInterval) GetSpotVolume() string`

GetSpotVolume returns the SpotVolume field if non-nil, zero value otherwise.

### GetSpotVolumeOk

`func (o *ExchangesTickerInterval) GetSpotVolumeOk() (*string, bool)`

GetSpotVolumeOk returns a tuple with the SpotVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotVolume

`func (o *ExchangesTickerInterval) SetSpotVolume(v string)`

SetSpotVolume sets SpotVolume field to given value.

### HasSpotVolume

`func (o *ExchangesTickerInterval) HasSpotVolume() bool`

HasSpotVolume returns a boolean if a field has been set.

### GetSpotVolumeChange

`func (o *ExchangesTickerInterval) GetSpotVolumeChange() string`

GetSpotVolumeChange returns the SpotVolumeChange field if non-nil, zero value otherwise.

### GetSpotVolumeChangeOk

`func (o *ExchangesTickerInterval) GetSpotVolumeChangeOk() (*string, bool)`

GetSpotVolumeChangeOk returns a tuple with the SpotVolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotVolumeChange

`func (o *ExchangesTickerInterval) SetSpotVolumeChange(v string)`

SetSpotVolumeChange sets SpotVolumeChange field to given value.

### HasSpotVolumeChange

`func (o *ExchangesTickerInterval) HasSpotVolumeChange() bool`

HasSpotVolumeChange returns a boolean if a field has been set.

### GetSpotVolumeChangePct

`func (o *ExchangesTickerInterval) GetSpotVolumeChangePct() string`

GetSpotVolumeChangePct returns the SpotVolumeChangePct field if non-nil, zero value otherwise.

### GetSpotVolumeChangePctOk

`func (o *ExchangesTickerInterval) GetSpotVolumeChangePctOk() (*string, bool)`

GetSpotVolumeChangePctOk returns a tuple with the SpotVolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotVolumeChangePct

`func (o *ExchangesTickerInterval) SetSpotVolumeChangePct(v string)`

SetSpotVolumeChangePct sets SpotVolumeChangePct field to given value.

### HasSpotVolumeChangePct

`func (o *ExchangesTickerInterval) HasSpotVolumeChangePct() bool`

HasSpotVolumeChangePct returns a boolean if a field has been set.

### GetDerivativeVolume

`func (o *ExchangesTickerInterval) GetDerivativeVolume() string`

GetDerivativeVolume returns the DerivativeVolume field if non-nil, zero value otherwise.

### GetDerivativeVolumeOk

`func (o *ExchangesTickerInterval) GetDerivativeVolumeOk() (*string, bool)`

GetDerivativeVolumeOk returns a tuple with the DerivativeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeVolume

`func (o *ExchangesTickerInterval) SetDerivativeVolume(v string)`

SetDerivativeVolume sets DerivativeVolume field to given value.

### HasDerivativeVolume

`func (o *ExchangesTickerInterval) HasDerivativeVolume() bool`

HasDerivativeVolume returns a boolean if a field has been set.

### GetDerivativeVolumeChange

`func (o *ExchangesTickerInterval) GetDerivativeVolumeChange() string`

GetDerivativeVolumeChange returns the DerivativeVolumeChange field if non-nil, zero value otherwise.

### GetDerivativeVolumeChangeOk

`func (o *ExchangesTickerInterval) GetDerivativeVolumeChangeOk() (*string, bool)`

GetDerivativeVolumeChangeOk returns a tuple with the DerivativeVolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeVolumeChange

`func (o *ExchangesTickerInterval) SetDerivativeVolumeChange(v string)`

SetDerivativeVolumeChange sets DerivativeVolumeChange field to given value.

### HasDerivativeVolumeChange

`func (o *ExchangesTickerInterval) HasDerivativeVolumeChange() bool`

HasDerivativeVolumeChange returns a boolean if a field has been set.

### GetDerivativeVolumeChangePct

`func (o *ExchangesTickerInterval) GetDerivativeVolumeChangePct() string`

GetDerivativeVolumeChangePct returns the DerivativeVolumeChangePct field if non-nil, zero value otherwise.

### GetDerivativeVolumeChangePctOk

`func (o *ExchangesTickerInterval) GetDerivativeVolumeChangePctOk() (*string, bool)`

GetDerivativeVolumeChangePctOk returns a tuple with the DerivativeVolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeVolumeChangePct

`func (o *ExchangesTickerInterval) SetDerivativeVolumeChangePct(v string)`

SetDerivativeVolumeChangePct sets DerivativeVolumeChangePct field to given value.

### HasDerivativeVolumeChangePct

`func (o *ExchangesTickerInterval) HasDerivativeVolumeChangePct() bool`

HasDerivativeVolumeChangePct returns a boolean if a field has been set.

### GetTrades

`func (o *ExchangesTickerInterval) GetTrades() string`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *ExchangesTickerInterval) GetTradesOk() (*string, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *ExchangesTickerInterval) SetTrades(v string)`

SetTrades sets Trades field to given value.

### HasTrades

`func (o *ExchangesTickerInterval) HasTrades() bool`

HasTrades returns a boolean if a field has been set.

### GetTradesChange

`func (o *ExchangesTickerInterval) GetTradesChange() string`

GetTradesChange returns the TradesChange field if non-nil, zero value otherwise.

### GetTradesChangeOk

`func (o *ExchangesTickerInterval) GetTradesChangeOk() (*string, bool)`

GetTradesChangeOk returns a tuple with the TradesChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradesChange

`func (o *ExchangesTickerInterval) SetTradesChange(v string)`

SetTradesChange sets TradesChange field to given value.

### HasTradesChange

`func (o *ExchangesTickerInterval) HasTradesChange() bool`

HasTradesChange returns a boolean if a field has been set.

### GetTradesChangePct

`func (o *ExchangesTickerInterval) GetTradesChangePct() string`

GetTradesChangePct returns the TradesChangePct field if non-nil, zero value otherwise.

### GetTradesChangePctOk

`func (o *ExchangesTickerInterval) GetTradesChangePctOk() (*string, bool)`

GetTradesChangePctOk returns a tuple with the TradesChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradesChangePct

`func (o *ExchangesTickerInterval) SetTradesChangePct(v string)`

SetTradesChangePct sets TradesChangePct field to given value.

### HasTradesChangePct

`func (o *ExchangesTickerInterval) HasTradesChangePct() bool`

HasTradesChangePct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


