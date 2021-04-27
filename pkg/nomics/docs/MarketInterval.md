# MarketInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | Pointer to **string** | Nomics Exchange ID | [optional] 
**Quote** | Pointer to **string** | Nomics Currency ID of the asset used to quote a price | [optional] 
**Volume** | Pointer to **string** | Total volume in USD | [optional] 
**Open** | Pointer to **string** | Open price in USD | [optional] 
**OpenTimestamp** | Pointer to **string** | Timestamp of the open price in RFC 3339 | [optional] 
**Close** | Pointer to **string** | Close price in USD | [optional] 
**CloseTimestamp** | Pointer to **string** | Timestamp of the close price in RFC 3339 | [optional] 

## Methods

### NewMarketInterval

`func NewMarketInterval() *MarketInterval`

NewMarketInterval instantiates a new MarketInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketIntervalWithDefaults

`func NewMarketIntervalWithDefaults() *MarketInterval`

NewMarketIntervalWithDefaults instantiates a new MarketInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *MarketInterval) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *MarketInterval) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *MarketInterval) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *MarketInterval) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetQuote

`func (o *MarketInterval) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *MarketInterval) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *MarketInterval) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *MarketInterval) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetVolume

`func (o *MarketInterval) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *MarketInterval) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *MarketInterval) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *MarketInterval) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetOpen

`func (o *MarketInterval) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *MarketInterval) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *MarketInterval) SetOpen(v string)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *MarketInterval) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetOpenTimestamp

`func (o *MarketInterval) GetOpenTimestamp() string`

GetOpenTimestamp returns the OpenTimestamp field if non-nil, zero value otherwise.

### GetOpenTimestampOk

`func (o *MarketInterval) GetOpenTimestampOk() (*string, bool)`

GetOpenTimestampOk returns a tuple with the OpenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTimestamp

`func (o *MarketInterval) SetOpenTimestamp(v string)`

SetOpenTimestamp sets OpenTimestamp field to given value.

### HasOpenTimestamp

`func (o *MarketInterval) HasOpenTimestamp() bool`

HasOpenTimestamp returns a boolean if a field has been set.

### GetClose

`func (o *MarketInterval) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *MarketInterval) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *MarketInterval) SetClose(v string)`

SetClose sets Close field to given value.

### HasClose

`func (o *MarketInterval) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetCloseTimestamp

`func (o *MarketInterval) GetCloseTimestamp() string`

GetCloseTimestamp returns the CloseTimestamp field if non-nil, zero value otherwise.

### GetCloseTimestampOk

`func (o *MarketInterval) GetCloseTimestampOk() (*string, bool)`

GetCloseTimestampOk returns a tuple with the CloseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTimestamp

`func (o *MarketInterval) SetCloseTimestamp(v string)`

SetCloseTimestamp sets CloseTimestamp field to given value.

### HasCloseTimestamp

`func (o *MarketInterval) HasCloseTimestamp() bool`

HasCloseTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


