# InlineResponse2006

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Nomics ID of the exchange | [optional] 
**Name** | Pointer to **string** | Name of the exchange | [optional] 
**LogoUrl** | Pointer to **string** | URL to logo of the exchange | [optional] 
**Rank** | Pointer to **string** | Rank of the exchange relative to other exchanges after the &#x60;status&#x60; and &#x60;type&#x60; filters are applied | [optional] 
**TransparencyGrade** | Pointer to **string** | Transparency grade of the exchange | [optional] 
**CoverageType** | Pointer to **string** | Coverage type of the exchange (trades, candles, tickers) | [optional] 
**OrderBooks** | Pointer to **bool** | Indicates if order books are available | [optional] 
**FirstTrade** | Pointer to **string** | Date of the first trade available | [optional] 
**FirstCandle** | Pointer to **string** | Date of the first candle available | [optional] 
**FirstOrderBook** | Pointer to **string** | Date of the first order book available | [optional] 
**LastUpdated** | Pointer to **string** | Timestamp of the last update of data | [optional] 
**FiatCurrencies** | Pointer to **[]string** | List of fiat currencies available | [optional] 
**NumPairs** | Pointer to **string** | Number of currency pairs (markets) available | [optional] 
**NumPairsUnmapped** | Pointer to **string** | Number of currency pairs (markets) that we know about, but are not yet available | [optional] 
**Centralized** | Pointer to **bool** | Whether or not the exchange is centralized | [optional] 
**Decentralized** | Pointer to **bool** | Whether or not the exchange is decentralized | [optional] 
**Interval** | Pointer to [**ExchangesTickerInterval**](_exchanges_ticker_interval.md) |  | [optional] 

## Methods

### NewInlineResponse2006

`func NewInlineResponse2006() *InlineResponse2006`

NewInlineResponse2006 instantiates a new InlineResponse2006 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2006WithDefaults

`func NewInlineResponse2006WithDefaults() *InlineResponse2006`

NewInlineResponse2006WithDefaults instantiates a new InlineResponse2006 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2006) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2006) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2006) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2006) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2006) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2006) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2006) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2006) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *InlineResponse2006) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *InlineResponse2006) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *InlineResponse2006) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *InlineResponse2006) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetRank

`func (o *InlineResponse2006) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *InlineResponse2006) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *InlineResponse2006) SetRank(v string)`

SetRank sets Rank field to given value.

### HasRank

`func (o *InlineResponse2006) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetTransparencyGrade

`func (o *InlineResponse2006) GetTransparencyGrade() string`

GetTransparencyGrade returns the TransparencyGrade field if non-nil, zero value otherwise.

### GetTransparencyGradeOk

`func (o *InlineResponse2006) GetTransparencyGradeOk() (*string, bool)`

GetTransparencyGradeOk returns a tuple with the TransparencyGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparencyGrade

`func (o *InlineResponse2006) SetTransparencyGrade(v string)`

SetTransparencyGrade sets TransparencyGrade field to given value.

### HasTransparencyGrade

`func (o *InlineResponse2006) HasTransparencyGrade() bool`

HasTransparencyGrade returns a boolean if a field has been set.

### GetCoverageType

`func (o *InlineResponse2006) GetCoverageType() string`

GetCoverageType returns the CoverageType field if non-nil, zero value otherwise.

### GetCoverageTypeOk

`func (o *InlineResponse2006) GetCoverageTypeOk() (*string, bool)`

GetCoverageTypeOk returns a tuple with the CoverageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageType

`func (o *InlineResponse2006) SetCoverageType(v string)`

SetCoverageType sets CoverageType field to given value.

### HasCoverageType

`func (o *InlineResponse2006) HasCoverageType() bool`

HasCoverageType returns a boolean if a field has been set.

### GetOrderBooks

`func (o *InlineResponse2006) GetOrderBooks() bool`

GetOrderBooks returns the OrderBooks field if non-nil, zero value otherwise.

### GetOrderBooksOk

`func (o *InlineResponse2006) GetOrderBooksOk() (*bool, bool)`

GetOrderBooksOk returns a tuple with the OrderBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBooks

`func (o *InlineResponse2006) SetOrderBooks(v bool)`

SetOrderBooks sets OrderBooks field to given value.

### HasOrderBooks

`func (o *InlineResponse2006) HasOrderBooks() bool`

HasOrderBooks returns a boolean if a field has been set.

### GetFirstTrade

`func (o *InlineResponse2006) GetFirstTrade() string`

GetFirstTrade returns the FirstTrade field if non-nil, zero value otherwise.

### GetFirstTradeOk

`func (o *InlineResponse2006) GetFirstTradeOk() (*string, bool)`

GetFirstTradeOk returns a tuple with the FirstTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTrade

`func (o *InlineResponse2006) SetFirstTrade(v string)`

SetFirstTrade sets FirstTrade field to given value.

### HasFirstTrade

`func (o *InlineResponse2006) HasFirstTrade() bool`

HasFirstTrade returns a boolean if a field has been set.

### GetFirstCandle

`func (o *InlineResponse2006) GetFirstCandle() string`

GetFirstCandle returns the FirstCandle field if non-nil, zero value otherwise.

### GetFirstCandleOk

`func (o *InlineResponse2006) GetFirstCandleOk() (*string, bool)`

GetFirstCandleOk returns a tuple with the FirstCandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCandle

`func (o *InlineResponse2006) SetFirstCandle(v string)`

SetFirstCandle sets FirstCandle field to given value.

### HasFirstCandle

`func (o *InlineResponse2006) HasFirstCandle() bool`

HasFirstCandle returns a boolean if a field has been set.

### GetFirstOrderBook

`func (o *InlineResponse2006) GetFirstOrderBook() string`

GetFirstOrderBook returns the FirstOrderBook field if non-nil, zero value otherwise.

### GetFirstOrderBookOk

`func (o *InlineResponse2006) GetFirstOrderBookOk() (*string, bool)`

GetFirstOrderBookOk returns a tuple with the FirstOrderBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOrderBook

`func (o *InlineResponse2006) SetFirstOrderBook(v string)`

SetFirstOrderBook sets FirstOrderBook field to given value.

### HasFirstOrderBook

`func (o *InlineResponse2006) HasFirstOrderBook() bool`

HasFirstOrderBook returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse2006) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse2006) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse2006) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse2006) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetFiatCurrencies

`func (o *InlineResponse2006) GetFiatCurrencies() []string`

GetFiatCurrencies returns the FiatCurrencies field if non-nil, zero value otherwise.

### GetFiatCurrenciesOk

`func (o *InlineResponse2006) GetFiatCurrenciesOk() (*[]string, bool)`

GetFiatCurrenciesOk returns a tuple with the FiatCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatCurrencies

`func (o *InlineResponse2006) SetFiatCurrencies(v []string)`

SetFiatCurrencies sets FiatCurrencies field to given value.

### HasFiatCurrencies

`func (o *InlineResponse2006) HasFiatCurrencies() bool`

HasFiatCurrencies returns a boolean if a field has been set.

### GetNumPairs

`func (o *InlineResponse2006) GetNumPairs() string`

GetNumPairs returns the NumPairs field if non-nil, zero value otherwise.

### GetNumPairsOk

`func (o *InlineResponse2006) GetNumPairsOk() (*string, bool)`

GetNumPairsOk returns a tuple with the NumPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPairs

`func (o *InlineResponse2006) SetNumPairs(v string)`

SetNumPairs sets NumPairs field to given value.

### HasNumPairs

`func (o *InlineResponse2006) HasNumPairs() bool`

HasNumPairs returns a boolean if a field has been set.

### GetNumPairsUnmapped

`func (o *InlineResponse2006) GetNumPairsUnmapped() string`

GetNumPairsUnmapped returns the NumPairsUnmapped field if non-nil, zero value otherwise.

### GetNumPairsUnmappedOk

`func (o *InlineResponse2006) GetNumPairsUnmappedOk() (*string, bool)`

GetNumPairsUnmappedOk returns a tuple with the NumPairsUnmapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPairsUnmapped

`func (o *InlineResponse2006) SetNumPairsUnmapped(v string)`

SetNumPairsUnmapped sets NumPairsUnmapped field to given value.

### HasNumPairsUnmapped

`func (o *InlineResponse2006) HasNumPairsUnmapped() bool`

HasNumPairsUnmapped returns a boolean if a field has been set.

### GetCentralized

`func (o *InlineResponse2006) GetCentralized() bool`

GetCentralized returns the Centralized field if non-nil, zero value otherwise.

### GetCentralizedOk

`func (o *InlineResponse2006) GetCentralizedOk() (*bool, bool)`

GetCentralizedOk returns a tuple with the Centralized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentralized

`func (o *InlineResponse2006) SetCentralized(v bool)`

SetCentralized sets Centralized field to given value.

### HasCentralized

`func (o *InlineResponse2006) HasCentralized() bool`

HasCentralized returns a boolean if a field has been set.

### GetDecentralized

`func (o *InlineResponse2006) GetDecentralized() bool`

GetDecentralized returns the Decentralized field if non-nil, zero value otherwise.

### GetDecentralizedOk

`func (o *InlineResponse2006) GetDecentralizedOk() (*bool, bool)`

GetDecentralizedOk returns a tuple with the Decentralized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecentralized

`func (o *InlineResponse2006) SetDecentralized(v bool)`

SetDecentralized sets Decentralized field to given value.

### HasDecentralized

`func (o *InlineResponse2006) HasDecentralized() bool`

HasDecentralized returns a boolean if a field has been set.

### GetInterval

`func (o *InlineResponse2006) GetInterval() ExchangesTickerInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InlineResponse2006) GetIntervalOk() (*ExchangesTickerInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InlineResponse2006) SetInterval(v ExchangesTickerInterval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InlineResponse2006) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


