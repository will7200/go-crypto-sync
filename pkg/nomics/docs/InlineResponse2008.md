# InlineResponse2008

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Nomics Exchange ID | [optional] 
**CapabilityMarkets** | Pointer to **bool** | Flag that determines if the exchange provides markets data | [optional] 
**CapabilityTrades** | Pointer to **bool** | Flag that determines if the exchange provides trades data | [optional] 
**CapabilityTradesByTimestamp** | Pointer to **bool** | Flag that determines if the exchange provides trades by timestamp data | [optional] 
**CapabilityTradesSnapshot** | Pointer to **bool** | Flag that determines if the exchange provides recent-only trades data | [optional] 
**CapabilityOrdersSnapshot** | Pointer to **bool** | Flag that determines if the exchange provides orderbook snapshot data | [optional] 
**CapabilityCandles** | Pointer to **bool** | Flag that determines if the exchange provides OHLCV candle data | [optional] 
**CapabilityTicker** | Pointer to **bool** | Flag that determines if the exchange provides 24h ticker data | [optional] 
**Integrated** | Pointer to **bool** | Flag that determines if the exchange has completed a [Nomics Deep Data Integration](https://nomicsintegration.com) | [optional] 
**Name** | Pointer to **string** | Exchange Name | [optional] 
**Description** | Pointer to **string** | Exchange description | [optional] 
**Location** | Pointer to **string** | Primary exchange location country | [optional] 
**LogoUrl** | Pointer to **string** | Exchange logo URL | [optional] 
**WebsiteUrl** | Pointer to **string** | Exchange website URL | [optional] 
**FeesUrl** | Pointer to **string** | Exchange fees URL | [optional] 
**TwitterUrl** | Pointer to **string** | Exchange Twitter URL | [optional] 
**FacebookUrl** | Pointer to **string** | Exchange Facebook URL | [optional] 
**RedditUrl** | Pointer to **string** | Exchange Reddit URL | [optional] 
**ChatUrl** | Pointer to **string** | Exchange chat URL | [optional] 
**BlogUrl** | Pointer to **string** | Exchange blog URL | [optional] 
**Year** | Pointer to **string** | Year exchange became active | [optional] 
**TransparencyGrade** | Pointer to **string** | The [Nomics Transparency Rating](https://blog.nomics.com/essays/transparency-ratings/) for the exchange | [optional] 
**OrderBooksInterval** | Pointer to **float32** | The maximum frequency (in milliseconds) at which order book snapshots are taken for this exchange | [optional] 
**Centralized** | Pointer to **bool** | Whether or not the exchange is centralized | [optional] 
**Decentralized** | Pointer to **bool** | Whether or not the exchange is decentralized | [optional] 

## Methods

### NewInlineResponse2008

`func NewInlineResponse2008() *InlineResponse2008`

NewInlineResponse2008 instantiates a new InlineResponse2008 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008WithDefaults

`func NewInlineResponse2008WithDefaults() *InlineResponse2008`

NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2008) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2008) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2008) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2008) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCapabilityMarkets

`func (o *InlineResponse2008) GetCapabilityMarkets() bool`

GetCapabilityMarkets returns the CapabilityMarkets field if non-nil, zero value otherwise.

### GetCapabilityMarketsOk

`func (o *InlineResponse2008) GetCapabilityMarketsOk() (*bool, bool)`

GetCapabilityMarketsOk returns a tuple with the CapabilityMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityMarkets

`func (o *InlineResponse2008) SetCapabilityMarkets(v bool)`

SetCapabilityMarkets sets CapabilityMarkets field to given value.

### HasCapabilityMarkets

`func (o *InlineResponse2008) HasCapabilityMarkets() bool`

HasCapabilityMarkets returns a boolean if a field has been set.

### GetCapabilityTrades

`func (o *InlineResponse2008) GetCapabilityTrades() bool`

GetCapabilityTrades returns the CapabilityTrades field if non-nil, zero value otherwise.

### GetCapabilityTradesOk

`func (o *InlineResponse2008) GetCapabilityTradesOk() (*bool, bool)`

GetCapabilityTradesOk returns a tuple with the CapabilityTrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityTrades

`func (o *InlineResponse2008) SetCapabilityTrades(v bool)`

SetCapabilityTrades sets CapabilityTrades field to given value.

### HasCapabilityTrades

`func (o *InlineResponse2008) HasCapabilityTrades() bool`

HasCapabilityTrades returns a boolean if a field has been set.

### GetCapabilityTradesByTimestamp

`func (o *InlineResponse2008) GetCapabilityTradesByTimestamp() bool`

GetCapabilityTradesByTimestamp returns the CapabilityTradesByTimestamp field if non-nil, zero value otherwise.

### GetCapabilityTradesByTimestampOk

`func (o *InlineResponse2008) GetCapabilityTradesByTimestampOk() (*bool, bool)`

GetCapabilityTradesByTimestampOk returns a tuple with the CapabilityTradesByTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityTradesByTimestamp

`func (o *InlineResponse2008) SetCapabilityTradesByTimestamp(v bool)`

SetCapabilityTradesByTimestamp sets CapabilityTradesByTimestamp field to given value.

### HasCapabilityTradesByTimestamp

`func (o *InlineResponse2008) HasCapabilityTradesByTimestamp() bool`

HasCapabilityTradesByTimestamp returns a boolean if a field has been set.

### GetCapabilityTradesSnapshot

`func (o *InlineResponse2008) GetCapabilityTradesSnapshot() bool`

GetCapabilityTradesSnapshot returns the CapabilityTradesSnapshot field if non-nil, zero value otherwise.

### GetCapabilityTradesSnapshotOk

`func (o *InlineResponse2008) GetCapabilityTradesSnapshotOk() (*bool, bool)`

GetCapabilityTradesSnapshotOk returns a tuple with the CapabilityTradesSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityTradesSnapshot

`func (o *InlineResponse2008) SetCapabilityTradesSnapshot(v bool)`

SetCapabilityTradesSnapshot sets CapabilityTradesSnapshot field to given value.

### HasCapabilityTradesSnapshot

`func (o *InlineResponse2008) HasCapabilityTradesSnapshot() bool`

HasCapabilityTradesSnapshot returns a boolean if a field has been set.

### GetCapabilityOrdersSnapshot

`func (o *InlineResponse2008) GetCapabilityOrdersSnapshot() bool`

GetCapabilityOrdersSnapshot returns the CapabilityOrdersSnapshot field if non-nil, zero value otherwise.

### GetCapabilityOrdersSnapshotOk

`func (o *InlineResponse2008) GetCapabilityOrdersSnapshotOk() (*bool, bool)`

GetCapabilityOrdersSnapshotOk returns a tuple with the CapabilityOrdersSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityOrdersSnapshot

`func (o *InlineResponse2008) SetCapabilityOrdersSnapshot(v bool)`

SetCapabilityOrdersSnapshot sets CapabilityOrdersSnapshot field to given value.

### HasCapabilityOrdersSnapshot

`func (o *InlineResponse2008) HasCapabilityOrdersSnapshot() bool`

HasCapabilityOrdersSnapshot returns a boolean if a field has been set.

### GetCapabilityCandles

`func (o *InlineResponse2008) GetCapabilityCandles() bool`

GetCapabilityCandles returns the CapabilityCandles field if non-nil, zero value otherwise.

### GetCapabilityCandlesOk

`func (o *InlineResponse2008) GetCapabilityCandlesOk() (*bool, bool)`

GetCapabilityCandlesOk returns a tuple with the CapabilityCandles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityCandles

`func (o *InlineResponse2008) SetCapabilityCandles(v bool)`

SetCapabilityCandles sets CapabilityCandles field to given value.

### HasCapabilityCandles

`func (o *InlineResponse2008) HasCapabilityCandles() bool`

HasCapabilityCandles returns a boolean if a field has been set.

### GetCapabilityTicker

`func (o *InlineResponse2008) GetCapabilityTicker() bool`

GetCapabilityTicker returns the CapabilityTicker field if non-nil, zero value otherwise.

### GetCapabilityTickerOk

`func (o *InlineResponse2008) GetCapabilityTickerOk() (*bool, bool)`

GetCapabilityTickerOk returns a tuple with the CapabilityTicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityTicker

`func (o *InlineResponse2008) SetCapabilityTicker(v bool)`

SetCapabilityTicker sets CapabilityTicker field to given value.

### HasCapabilityTicker

`func (o *InlineResponse2008) HasCapabilityTicker() bool`

HasCapabilityTicker returns a boolean if a field has been set.

### GetIntegrated

`func (o *InlineResponse2008) GetIntegrated() bool`

GetIntegrated returns the Integrated field if non-nil, zero value otherwise.

### GetIntegratedOk

`func (o *InlineResponse2008) GetIntegratedOk() (*bool, bool)`

GetIntegratedOk returns a tuple with the Integrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrated

`func (o *InlineResponse2008) SetIntegrated(v bool)`

SetIntegrated sets Integrated field to given value.

### HasIntegrated

`func (o *InlineResponse2008) HasIntegrated() bool`

HasIntegrated returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2008) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2008) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2008) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2008) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse2008) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse2008) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse2008) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse2008) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *InlineResponse2008) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InlineResponse2008) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InlineResponse2008) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InlineResponse2008) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLogoUrl

`func (o *InlineResponse2008) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *InlineResponse2008) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *InlineResponse2008) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *InlineResponse2008) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *InlineResponse2008) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *InlineResponse2008) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *InlineResponse2008) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *InlineResponse2008) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetFeesUrl

`func (o *InlineResponse2008) GetFeesUrl() string`

GetFeesUrl returns the FeesUrl field if non-nil, zero value otherwise.

### GetFeesUrlOk

`func (o *InlineResponse2008) GetFeesUrlOk() (*string, bool)`

GetFeesUrlOk returns a tuple with the FeesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesUrl

`func (o *InlineResponse2008) SetFeesUrl(v string)`

SetFeesUrl sets FeesUrl field to given value.

### HasFeesUrl

`func (o *InlineResponse2008) HasFeesUrl() bool`

HasFeesUrl returns a boolean if a field has been set.

### GetTwitterUrl

`func (o *InlineResponse2008) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *InlineResponse2008) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *InlineResponse2008) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *InlineResponse2008) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### GetFacebookUrl

`func (o *InlineResponse2008) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *InlineResponse2008) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *InlineResponse2008) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *InlineResponse2008) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### GetRedditUrl

`func (o *InlineResponse2008) GetRedditUrl() string`

GetRedditUrl returns the RedditUrl field if non-nil, zero value otherwise.

### GetRedditUrlOk

`func (o *InlineResponse2008) GetRedditUrlOk() (*string, bool)`

GetRedditUrlOk returns a tuple with the RedditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedditUrl

`func (o *InlineResponse2008) SetRedditUrl(v string)`

SetRedditUrl sets RedditUrl field to given value.

### HasRedditUrl

`func (o *InlineResponse2008) HasRedditUrl() bool`

HasRedditUrl returns a boolean if a field has been set.

### GetChatUrl

`func (o *InlineResponse2008) GetChatUrl() string`

GetChatUrl returns the ChatUrl field if non-nil, zero value otherwise.

### GetChatUrlOk

`func (o *InlineResponse2008) GetChatUrlOk() (*string, bool)`

GetChatUrlOk returns a tuple with the ChatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatUrl

`func (o *InlineResponse2008) SetChatUrl(v string)`

SetChatUrl sets ChatUrl field to given value.

### HasChatUrl

`func (o *InlineResponse2008) HasChatUrl() bool`

HasChatUrl returns a boolean if a field has been set.

### GetBlogUrl

`func (o *InlineResponse2008) GetBlogUrl() string`

GetBlogUrl returns the BlogUrl field if non-nil, zero value otherwise.

### GetBlogUrlOk

`func (o *InlineResponse2008) GetBlogUrlOk() (*string, bool)`

GetBlogUrlOk returns a tuple with the BlogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogUrl

`func (o *InlineResponse2008) SetBlogUrl(v string)`

SetBlogUrl sets BlogUrl field to given value.

### HasBlogUrl

`func (o *InlineResponse2008) HasBlogUrl() bool`

HasBlogUrl returns a boolean if a field has been set.

### GetYear

`func (o *InlineResponse2008) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *InlineResponse2008) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *InlineResponse2008) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *InlineResponse2008) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetTransparencyGrade

`func (o *InlineResponse2008) GetTransparencyGrade() string`

GetTransparencyGrade returns the TransparencyGrade field if non-nil, zero value otherwise.

### GetTransparencyGradeOk

`func (o *InlineResponse2008) GetTransparencyGradeOk() (*string, bool)`

GetTransparencyGradeOk returns a tuple with the TransparencyGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparencyGrade

`func (o *InlineResponse2008) SetTransparencyGrade(v string)`

SetTransparencyGrade sets TransparencyGrade field to given value.

### HasTransparencyGrade

`func (o *InlineResponse2008) HasTransparencyGrade() bool`

HasTransparencyGrade returns a boolean if a field has been set.

### GetOrderBooksInterval

`func (o *InlineResponse2008) GetOrderBooksInterval() float32`

GetOrderBooksInterval returns the OrderBooksInterval field if non-nil, zero value otherwise.

### GetOrderBooksIntervalOk

`func (o *InlineResponse2008) GetOrderBooksIntervalOk() (*float32, bool)`

GetOrderBooksIntervalOk returns a tuple with the OrderBooksInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBooksInterval

`func (o *InlineResponse2008) SetOrderBooksInterval(v float32)`

SetOrderBooksInterval sets OrderBooksInterval field to given value.

### HasOrderBooksInterval

`func (o *InlineResponse2008) HasOrderBooksInterval() bool`

HasOrderBooksInterval returns a boolean if a field has been set.

### GetCentralized

`func (o *InlineResponse2008) GetCentralized() bool`

GetCentralized returns the Centralized field if non-nil, zero value otherwise.

### GetCentralizedOk

`func (o *InlineResponse2008) GetCentralizedOk() (*bool, bool)`

GetCentralizedOk returns a tuple with the Centralized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentralized

`func (o *InlineResponse2008) SetCentralized(v bool)`

SetCentralized sets Centralized field to given value.

### HasCentralized

`func (o *InlineResponse2008) HasCentralized() bool`

HasCentralized returns a boolean if a field has been set.

### GetDecentralized

`func (o *InlineResponse2008) GetDecentralized() bool`

GetDecentralized returns the Decentralized field if non-nil, zero value otherwise.

### GetDecentralizedOk

`func (o *InlineResponse2008) GetDecentralizedOk() (*bool, bool)`

GetDecentralizedOk returns a tuple with the Decentralized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecentralized

`func (o *InlineResponse2008) SetDecentralized(v bool)`

SetDecentralized sets Decentralized field to given value.

### HasDecentralized

`func (o *InlineResponse2008) HasDecentralized() bool`

HasDecentralized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


