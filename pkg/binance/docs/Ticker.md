# Ticker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**PriceChange** | Pointer to **string** |  | [optional] 
**PriceChangePercent** | Pointer to **string** |  | [optional] 
**PrevClosePrice** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**BidPrice** | Pointer to **string** |  | [optional] 
**BidQty** | Pointer to **string** |  | [optional] 
**AskPrice** | Pointer to **string** |  | [optional] 
**AskQty** | Pointer to **string** |  | [optional] 
**OpenPrice** | Pointer to **string** |  | [optional] 
**HighPrice** | Pointer to **string** |  | [optional] 
**LowPrice** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**QuoteVolume** | Pointer to **string** |  | [optional] 
**OpenTime** | Pointer to **int32** |  | [optional] 
**CloseTime** | Pointer to **int32** |  | [optional] 
**FirstId** | Pointer to **int32** |  | [optional] 
**LastId** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewTicker

`func NewTicker() *Ticker`

NewTicker instantiates a new Ticker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerWithDefaults

`func NewTickerWithDefaults() *Ticker`

NewTickerWithDefaults instantiates a new Ticker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Ticker) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Ticker) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Ticker) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Ticker) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPriceChange

`func (o *Ticker) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *Ticker) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *Ticker) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *Ticker) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetPriceChangePercent

`func (o *Ticker) GetPriceChangePercent() string`

GetPriceChangePercent returns the PriceChangePercent field if non-nil, zero value otherwise.

### GetPriceChangePercentOk

`func (o *Ticker) GetPriceChangePercentOk() (*string, bool)`

GetPriceChangePercentOk returns a tuple with the PriceChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangePercent

`func (o *Ticker) SetPriceChangePercent(v string)`

SetPriceChangePercent sets PriceChangePercent field to given value.

### HasPriceChangePercent

`func (o *Ticker) HasPriceChangePercent() bool`

HasPriceChangePercent returns a boolean if a field has been set.

### GetPrevClosePrice

`func (o *Ticker) GetPrevClosePrice() string`

GetPrevClosePrice returns the PrevClosePrice field if non-nil, zero value otherwise.

### GetPrevClosePriceOk

`func (o *Ticker) GetPrevClosePriceOk() (*string, bool)`

GetPrevClosePriceOk returns a tuple with the PrevClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevClosePrice

`func (o *Ticker) SetPrevClosePrice(v string)`

SetPrevClosePrice sets PrevClosePrice field to given value.

### HasPrevClosePrice

`func (o *Ticker) HasPrevClosePrice() bool`

HasPrevClosePrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *Ticker) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *Ticker) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *Ticker) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *Ticker) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetBidPrice

`func (o *Ticker) GetBidPrice() string`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *Ticker) GetBidPriceOk() (*string, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *Ticker) SetBidPrice(v string)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *Ticker) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidQty

`func (o *Ticker) GetBidQty() string`

GetBidQty returns the BidQty field if non-nil, zero value otherwise.

### GetBidQtyOk

`func (o *Ticker) GetBidQtyOk() (*string, bool)`

GetBidQtyOk returns a tuple with the BidQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidQty

`func (o *Ticker) SetBidQty(v string)`

SetBidQty sets BidQty field to given value.

### HasBidQty

`func (o *Ticker) HasBidQty() bool`

HasBidQty returns a boolean if a field has been set.

### GetAskPrice

`func (o *Ticker) GetAskPrice() string`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *Ticker) GetAskPriceOk() (*string, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *Ticker) SetAskPrice(v string)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *Ticker) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskQty

`func (o *Ticker) GetAskQty() string`

GetAskQty returns the AskQty field if non-nil, zero value otherwise.

### GetAskQtyOk

`func (o *Ticker) GetAskQtyOk() (*string, bool)`

GetAskQtyOk returns a tuple with the AskQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskQty

`func (o *Ticker) SetAskQty(v string)`

SetAskQty sets AskQty field to given value.

### HasAskQty

`func (o *Ticker) HasAskQty() bool`

HasAskQty returns a boolean if a field has been set.

### GetOpenPrice

`func (o *Ticker) GetOpenPrice() string`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *Ticker) GetOpenPriceOk() (*string, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *Ticker) SetOpenPrice(v string)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *Ticker) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetHighPrice

`func (o *Ticker) GetHighPrice() string`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *Ticker) GetHighPriceOk() (*string, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *Ticker) SetHighPrice(v string)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *Ticker) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLowPrice

`func (o *Ticker) GetLowPrice() string`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *Ticker) GetLowPriceOk() (*string, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *Ticker) SetLowPrice(v string)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *Ticker) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetVolume

`func (o *Ticker) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Ticker) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Ticker) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Ticker) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetQuoteVolume

`func (o *Ticker) GetQuoteVolume() string`

GetQuoteVolume returns the QuoteVolume field if non-nil, zero value otherwise.

### GetQuoteVolumeOk

`func (o *Ticker) GetQuoteVolumeOk() (*string, bool)`

GetQuoteVolumeOk returns a tuple with the QuoteVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVolume

`func (o *Ticker) SetQuoteVolume(v string)`

SetQuoteVolume sets QuoteVolume field to given value.

### HasQuoteVolume

`func (o *Ticker) HasQuoteVolume() bool`

HasQuoteVolume returns a boolean if a field has been set.

### GetOpenTime

`func (o *Ticker) GetOpenTime() int32`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *Ticker) GetOpenTimeOk() (*int32, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *Ticker) SetOpenTime(v int32)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *Ticker) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetCloseTime

`func (o *Ticker) GetCloseTime() int32`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *Ticker) GetCloseTimeOk() (*int32, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *Ticker) SetCloseTime(v int32)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *Ticker) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetFirstId

`func (o *Ticker) GetFirstId() int32`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *Ticker) GetFirstIdOk() (*int32, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *Ticker) SetFirstId(v int32)`

SetFirstId sets FirstId field to given value.

### HasFirstId

`func (o *Ticker) HasFirstId() bool`

HasFirstId returns a boolean if a field has been set.

### GetLastId

`func (o *Ticker) GetLastId() int32`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *Ticker) GetLastIdOk() (*int32, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *Ticker) SetLastId(v int32)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *Ticker) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### GetCount

`func (o *Ticker) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Ticker) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Ticker) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Ticker) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


