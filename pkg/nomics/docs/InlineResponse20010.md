# InlineResponse20010

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | Pointer to **string** | Nomics Exchange ID | [optional] 
**Market** | Pointer to **string** | Nomics Market ID | [optional] 
**Type** | Pointer to **string** | Market Type | [optional] 
**Subtype** | Pointer to **string** | Market Subtype | [optional] 
**Aggregated** | Pointer to **bool** | Indicates if the market is used to price the base and quote currencies. See [Nomics Pricing Methodology](https://blog.nomics.com/newsletter/nomics-pricing-methodology-explained/) for more information. | [optional] 
**PriceExclude** | Pointer to **bool** | Indicates if the market has been excluded from pricing the base and quote currencies. | [optional] 
**VolumeExclude** | Pointer to **bool** | Indicates if the market has been excluded from counting towards volume for base and quote currencies. | [optional] 
**Base** | Pointer to **string** | Nomics Currency ID of the base of the market | [optional] 
**Quote** | Pointer to **string** | Nomics Currency ID of the quote of the market | [optional] 
**BaseSymbol** | Pointer to **string** | Nomics display symbol of the base of the market | [optional] 
**QuoteSymbol** | Pointer to **string** | Nomics display symbol of the quote of the market | [optional] 
**Price** | Pointer to **string** | Latest price of the market in USD | [optional] 
**PriceQuote** | Pointer to **string** | Latest price of the market in the quote currency | [optional] 
**VolumeUsd** | Pointer to **string** | Market volume in USD based on the first &#x60;interval&#x60; (or &#x60;1d&#x60; if none specified) | [optional] 
**LastUpdated** | Pointer to **string** | Timestamp of when the market was last updated | [optional] 
**Interval** | Pointer to [**ExchangeMarketsTickerInterval**](_exchange_markets_ticker_interval.md) |  | [optional] 

## Methods

### NewInlineResponse20010

`func NewInlineResponse20010() *InlineResponse20010`

NewInlineResponse20010 instantiates a new InlineResponse20010 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20010WithDefaults

`func NewInlineResponse20010WithDefaults() *InlineResponse20010`

NewInlineResponse20010WithDefaults instantiates a new InlineResponse20010 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *InlineResponse20010) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InlineResponse20010) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InlineResponse20010) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InlineResponse20010) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMarket

`func (o *InlineResponse20010) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *InlineResponse20010) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *InlineResponse20010) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *InlineResponse20010) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20010) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20010) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20010) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20010) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *InlineResponse20010) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InlineResponse20010) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InlineResponse20010) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *InlineResponse20010) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetAggregated

`func (o *InlineResponse20010) GetAggregated() bool`

GetAggregated returns the Aggregated field if non-nil, zero value otherwise.

### GetAggregatedOk

`func (o *InlineResponse20010) GetAggregatedOk() (*bool, bool)`

GetAggregatedOk returns a tuple with the Aggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregated

`func (o *InlineResponse20010) SetAggregated(v bool)`

SetAggregated sets Aggregated field to given value.

### HasAggregated

`func (o *InlineResponse20010) HasAggregated() bool`

HasAggregated returns a boolean if a field has been set.

### GetPriceExclude

`func (o *InlineResponse20010) GetPriceExclude() bool`

GetPriceExclude returns the PriceExclude field if non-nil, zero value otherwise.

### GetPriceExcludeOk

`func (o *InlineResponse20010) GetPriceExcludeOk() (*bool, bool)`

GetPriceExcludeOk returns a tuple with the PriceExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceExclude

`func (o *InlineResponse20010) SetPriceExclude(v bool)`

SetPriceExclude sets PriceExclude field to given value.

### HasPriceExclude

`func (o *InlineResponse20010) HasPriceExclude() bool`

HasPriceExclude returns a boolean if a field has been set.

### GetVolumeExclude

`func (o *InlineResponse20010) GetVolumeExclude() bool`

GetVolumeExclude returns the VolumeExclude field if non-nil, zero value otherwise.

### GetVolumeExcludeOk

`func (o *InlineResponse20010) GetVolumeExcludeOk() (*bool, bool)`

GetVolumeExcludeOk returns a tuple with the VolumeExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeExclude

`func (o *InlineResponse20010) SetVolumeExclude(v bool)`

SetVolumeExclude sets VolumeExclude field to given value.

### HasVolumeExclude

`func (o *InlineResponse20010) HasVolumeExclude() bool`

HasVolumeExclude returns a boolean if a field has been set.

### GetBase

`func (o *InlineResponse20010) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *InlineResponse20010) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *InlineResponse20010) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *InlineResponse20010) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetQuote

`func (o *InlineResponse20010) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *InlineResponse20010) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *InlineResponse20010) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *InlineResponse20010) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetBaseSymbol

`func (o *InlineResponse20010) GetBaseSymbol() string`

GetBaseSymbol returns the BaseSymbol field if non-nil, zero value otherwise.

### GetBaseSymbolOk

`func (o *InlineResponse20010) GetBaseSymbolOk() (*string, bool)`

GetBaseSymbolOk returns a tuple with the BaseSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSymbol

`func (o *InlineResponse20010) SetBaseSymbol(v string)`

SetBaseSymbol sets BaseSymbol field to given value.

### HasBaseSymbol

`func (o *InlineResponse20010) HasBaseSymbol() bool`

HasBaseSymbol returns a boolean if a field has been set.

### GetQuoteSymbol

`func (o *InlineResponse20010) GetQuoteSymbol() string`

GetQuoteSymbol returns the QuoteSymbol field if non-nil, zero value otherwise.

### GetQuoteSymbolOk

`func (o *InlineResponse20010) GetQuoteSymbolOk() (*string, bool)`

GetQuoteSymbolOk returns a tuple with the QuoteSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSymbol

`func (o *InlineResponse20010) SetQuoteSymbol(v string)`

SetQuoteSymbol sets QuoteSymbol field to given value.

### HasQuoteSymbol

`func (o *InlineResponse20010) HasQuoteSymbol() bool`

HasQuoteSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *InlineResponse20010) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InlineResponse20010) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InlineResponse20010) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InlineResponse20010) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceQuote

`func (o *InlineResponse20010) GetPriceQuote() string`

GetPriceQuote returns the PriceQuote field if non-nil, zero value otherwise.

### GetPriceQuoteOk

`func (o *InlineResponse20010) GetPriceQuoteOk() (*string, bool)`

GetPriceQuoteOk returns a tuple with the PriceQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceQuote

`func (o *InlineResponse20010) SetPriceQuote(v string)`

SetPriceQuote sets PriceQuote field to given value.

### HasPriceQuote

`func (o *InlineResponse20010) HasPriceQuote() bool`

HasPriceQuote returns a boolean if a field has been set.

### GetVolumeUsd

`func (o *InlineResponse20010) GetVolumeUsd() string`

GetVolumeUsd returns the VolumeUsd field if non-nil, zero value otherwise.

### GetVolumeUsdOk

`func (o *InlineResponse20010) GetVolumeUsdOk() (*string, bool)`

GetVolumeUsdOk returns a tuple with the VolumeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUsd

`func (o *InlineResponse20010) SetVolumeUsd(v string)`

SetVolumeUsd sets VolumeUsd field to given value.

### HasVolumeUsd

`func (o *InlineResponse20010) HasVolumeUsd() bool`

HasVolumeUsd returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20010) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20010) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20010) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20010) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetInterval

`func (o *InlineResponse20010) GetInterval() ExchangeMarketsTickerInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InlineResponse20010) GetIntervalOk() (*ExchangeMarketsTickerInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InlineResponse20010) SetInterval(v ExchangeMarketsTickerInterval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InlineResponse20010) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


