# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Nomics ID of the currency | [optional] 
**Id** | Pointer to **string** | Unique Nomics ID of the currency | [optional] 
**Status** | Pointer to **string** | Current status | [optional] 
**Price** | Pointer to **string** | Current price | [optional] 
**PriceDate** | Pointer to **string** | Date of the price | [optional] 
**PriceTimestamp** | Pointer to **string** | Timestamp of the price | [optional] 
**Symbol** | Pointer to **string** | Display symbol of the currency (may be duplicated) | [optional] 
**CirculatingSupply** | Pointer to **string** | Current circulating supply | [optional] 
**MaxSupply** | Pointer to **string** | Current maximum supply | [optional] 
**Name** | Pointer to **string** | Name of the currency | [optional] 
**LogoUrl** | Pointer to **string** | URL to logo of the currency | [optional] 
**MarketCap** | Pointer to **string** | Current market cap | [optional] 
**TransparentMarketCap** | Pointer to **string** | Current transparent market cap | [optional] 
**NumExchanges** | Pointer to **string** | The number of exchanges on which this asset is traded | [optional] 
**NumPairs** | Pointer to **string** | Number of currency pairs (markets) this asset is a part of, where both assets are known | [optional] 
**NumPairsUnmapped** | Pointer to **string** | Number of currency pairs (markets) this asset is a part of, but where the other asset is unknown | [optional] 
**FirstCandle** | Pointer to **string** | RFC3999 timestamp of the first &#x60;1d&#x60; candle available via the Nomics API | [optional] 
**FirstTrade** | Pointer to **string** | RFC3999 timestamp of the first trade available via the Nomics API | [optional] 
**FirstOrderBook** | Pointer to **string** | RFC3999 timestamp of the first order book snapshot available via the Nomics API | [optional] 
**FirstPricedAt** | Pointer to **string** | RFC3999 timestamp representing the currency was first priced by Nomics | [optional] 
**Rank** | Pointer to **string** | Rank of the currency by market cap | [optional] 
**RankDelta** | Pointer to **string** | Relative rank change based on the first specified &#x60;interval&#x60; other than &#x60;1h&#x60;.  This field is only present on &#x60;active&#x60; currencies. | [optional] 
**High** | Pointer to **string** | All time high price | [optional] 
**HighTimestamp** | Pointer to **string** | RFC3999 timestamp of the all time high | [optional] 
**Interval** | Pointer to [**CurrenciesTickerInterval**](_currencies_ticker_interval.md) |  | [optional] 

## Methods

### NewInlineResponse200

`func NewInlineResponse200() *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InlineResponse200) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InlineResponse200) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InlineResponse200) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InlineResponse200) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrice

`func (o *InlineResponse200) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InlineResponse200) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InlineResponse200) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InlineResponse200) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDate

`func (o *InlineResponse200) GetPriceDate() string`

GetPriceDate returns the PriceDate field if non-nil, zero value otherwise.

### GetPriceDateOk

`func (o *InlineResponse200) GetPriceDateOk() (*string, bool)`

GetPriceDateOk returns a tuple with the PriceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDate

`func (o *InlineResponse200) SetPriceDate(v string)`

SetPriceDate sets PriceDate field to given value.

### HasPriceDate

`func (o *InlineResponse200) HasPriceDate() bool`

HasPriceDate returns a boolean if a field has been set.

### GetPriceTimestamp

`func (o *InlineResponse200) GetPriceTimestamp() string`

GetPriceTimestamp returns the PriceTimestamp field if non-nil, zero value otherwise.

### GetPriceTimestampOk

`func (o *InlineResponse200) GetPriceTimestampOk() (*string, bool)`

GetPriceTimestampOk returns a tuple with the PriceTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTimestamp

`func (o *InlineResponse200) SetPriceTimestamp(v string)`

SetPriceTimestamp sets PriceTimestamp field to given value.

### HasPriceTimestamp

`func (o *InlineResponse200) HasPriceTimestamp() bool`

HasPriceTimestamp returns a boolean if a field has been set.

### GetSymbol

`func (o *InlineResponse200) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineResponse200) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineResponse200) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InlineResponse200) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCirculatingSupply

`func (o *InlineResponse200) GetCirculatingSupply() string`

GetCirculatingSupply returns the CirculatingSupply field if non-nil, zero value otherwise.

### GetCirculatingSupplyOk

`func (o *InlineResponse200) GetCirculatingSupplyOk() (*string, bool)`

GetCirculatingSupplyOk returns a tuple with the CirculatingSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCirculatingSupply

`func (o *InlineResponse200) SetCirculatingSupply(v string)`

SetCirculatingSupply sets CirculatingSupply field to given value.

### HasCirculatingSupply

`func (o *InlineResponse200) HasCirculatingSupply() bool`

HasCirculatingSupply returns a boolean if a field has been set.

### GetMaxSupply

`func (o *InlineResponse200) GetMaxSupply() string`

GetMaxSupply returns the MaxSupply field if non-nil, zero value otherwise.

### GetMaxSupplyOk

`func (o *InlineResponse200) GetMaxSupplyOk() (*string, bool)`

GetMaxSupplyOk returns a tuple with the MaxSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupply

`func (o *InlineResponse200) SetMaxSupply(v string)`

SetMaxSupply sets MaxSupply field to given value.

### HasMaxSupply

`func (o *InlineResponse200) HasMaxSupply() bool`

HasMaxSupply returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *InlineResponse200) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *InlineResponse200) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *InlineResponse200) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *InlineResponse200) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetMarketCap

`func (o *InlineResponse200) GetMarketCap() string`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *InlineResponse200) GetMarketCapOk() (*string, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *InlineResponse200) SetMarketCap(v string)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *InlineResponse200) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetTransparentMarketCap

`func (o *InlineResponse200) GetTransparentMarketCap() string`

GetTransparentMarketCap returns the TransparentMarketCap field if non-nil, zero value otherwise.

### GetTransparentMarketCapOk

`func (o *InlineResponse200) GetTransparentMarketCapOk() (*string, bool)`

GetTransparentMarketCapOk returns a tuple with the TransparentMarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentMarketCap

`func (o *InlineResponse200) SetTransparentMarketCap(v string)`

SetTransparentMarketCap sets TransparentMarketCap field to given value.

### HasTransparentMarketCap

`func (o *InlineResponse200) HasTransparentMarketCap() bool`

HasTransparentMarketCap returns a boolean if a field has been set.

### GetNumExchanges

`func (o *InlineResponse200) GetNumExchanges() string`

GetNumExchanges returns the NumExchanges field if non-nil, zero value otherwise.

### GetNumExchangesOk

`func (o *InlineResponse200) GetNumExchangesOk() (*string, bool)`

GetNumExchangesOk returns a tuple with the NumExchanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumExchanges

`func (o *InlineResponse200) SetNumExchanges(v string)`

SetNumExchanges sets NumExchanges field to given value.

### HasNumExchanges

`func (o *InlineResponse200) HasNumExchanges() bool`

HasNumExchanges returns a boolean if a field has been set.

### GetNumPairs

`func (o *InlineResponse200) GetNumPairs() string`

GetNumPairs returns the NumPairs field if non-nil, zero value otherwise.

### GetNumPairsOk

`func (o *InlineResponse200) GetNumPairsOk() (*string, bool)`

GetNumPairsOk returns a tuple with the NumPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPairs

`func (o *InlineResponse200) SetNumPairs(v string)`

SetNumPairs sets NumPairs field to given value.

### HasNumPairs

`func (o *InlineResponse200) HasNumPairs() bool`

HasNumPairs returns a boolean if a field has been set.

### GetNumPairsUnmapped

`func (o *InlineResponse200) GetNumPairsUnmapped() string`

GetNumPairsUnmapped returns the NumPairsUnmapped field if non-nil, zero value otherwise.

### GetNumPairsUnmappedOk

`func (o *InlineResponse200) GetNumPairsUnmappedOk() (*string, bool)`

GetNumPairsUnmappedOk returns a tuple with the NumPairsUnmapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPairsUnmapped

`func (o *InlineResponse200) SetNumPairsUnmapped(v string)`

SetNumPairsUnmapped sets NumPairsUnmapped field to given value.

### HasNumPairsUnmapped

`func (o *InlineResponse200) HasNumPairsUnmapped() bool`

HasNumPairsUnmapped returns a boolean if a field has been set.

### GetFirstCandle

`func (o *InlineResponse200) GetFirstCandle() string`

GetFirstCandle returns the FirstCandle field if non-nil, zero value otherwise.

### GetFirstCandleOk

`func (o *InlineResponse200) GetFirstCandleOk() (*string, bool)`

GetFirstCandleOk returns a tuple with the FirstCandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCandle

`func (o *InlineResponse200) SetFirstCandle(v string)`

SetFirstCandle sets FirstCandle field to given value.

### HasFirstCandle

`func (o *InlineResponse200) HasFirstCandle() bool`

HasFirstCandle returns a boolean if a field has been set.

### GetFirstTrade

`func (o *InlineResponse200) GetFirstTrade() string`

GetFirstTrade returns the FirstTrade field if non-nil, zero value otherwise.

### GetFirstTradeOk

`func (o *InlineResponse200) GetFirstTradeOk() (*string, bool)`

GetFirstTradeOk returns a tuple with the FirstTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTrade

`func (o *InlineResponse200) SetFirstTrade(v string)`

SetFirstTrade sets FirstTrade field to given value.

### HasFirstTrade

`func (o *InlineResponse200) HasFirstTrade() bool`

HasFirstTrade returns a boolean if a field has been set.

### GetFirstOrderBook

`func (o *InlineResponse200) GetFirstOrderBook() string`

GetFirstOrderBook returns the FirstOrderBook field if non-nil, zero value otherwise.

### GetFirstOrderBookOk

`func (o *InlineResponse200) GetFirstOrderBookOk() (*string, bool)`

GetFirstOrderBookOk returns a tuple with the FirstOrderBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOrderBook

`func (o *InlineResponse200) SetFirstOrderBook(v string)`

SetFirstOrderBook sets FirstOrderBook field to given value.

### HasFirstOrderBook

`func (o *InlineResponse200) HasFirstOrderBook() bool`

HasFirstOrderBook returns a boolean if a field has been set.

### GetFirstPricedAt

`func (o *InlineResponse200) GetFirstPricedAt() string`

GetFirstPricedAt returns the FirstPricedAt field if non-nil, zero value otherwise.

### GetFirstPricedAtOk

`func (o *InlineResponse200) GetFirstPricedAtOk() (*string, bool)`

GetFirstPricedAtOk returns a tuple with the FirstPricedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPricedAt

`func (o *InlineResponse200) SetFirstPricedAt(v string)`

SetFirstPricedAt sets FirstPricedAt field to given value.

### HasFirstPricedAt

`func (o *InlineResponse200) HasFirstPricedAt() bool`

HasFirstPricedAt returns a boolean if a field has been set.

### GetRank

`func (o *InlineResponse200) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *InlineResponse200) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *InlineResponse200) SetRank(v string)`

SetRank sets Rank field to given value.

### HasRank

`func (o *InlineResponse200) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRankDelta

`func (o *InlineResponse200) GetRankDelta() string`

GetRankDelta returns the RankDelta field if non-nil, zero value otherwise.

### GetRankDeltaOk

`func (o *InlineResponse200) GetRankDeltaOk() (*string, bool)`

GetRankDeltaOk returns a tuple with the RankDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankDelta

`func (o *InlineResponse200) SetRankDelta(v string)`

SetRankDelta sets RankDelta field to given value.

### HasRankDelta

`func (o *InlineResponse200) HasRankDelta() bool`

HasRankDelta returns a boolean if a field has been set.

### GetHigh

`func (o *InlineResponse200) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *InlineResponse200) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *InlineResponse200) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *InlineResponse200) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetHighTimestamp

`func (o *InlineResponse200) GetHighTimestamp() string`

GetHighTimestamp returns the HighTimestamp field if non-nil, zero value otherwise.

### GetHighTimestampOk

`func (o *InlineResponse200) GetHighTimestampOk() (*string, bool)`

GetHighTimestampOk returns a tuple with the HighTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighTimestamp

`func (o *InlineResponse200) SetHighTimestamp(v string)`

SetHighTimestamp sets HighTimestamp field to given value.

### HasHighTimestamp

`func (o *InlineResponse200) HasHighTimestamp() bool`

HasHighTimestamp returns a boolean if a field has been set.

### GetInterval

`func (o *InlineResponse200) GetInterval() CurrenciesTickerInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InlineResponse200) GetIntervalOk() (*CurrenciesTickerInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InlineResponse200) SetInterval(v CurrenciesTickerInterval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InlineResponse200) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


