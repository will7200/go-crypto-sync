# InlineResponse20020

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | Start time of the candle in RFC3339 | [optional] 
**Low** | Pointer to **string** | Lowest price in quote currency | [optional] 
**Open** | Pointer to **string** | Volume weighted average of first trade prices in quote currency | [optional] 
**Close** | Pointer to **string** | Volume weighted average of last trade prices in quote currency | [optional] 
**High** | Pointer to **string** | Highest price in quote currency | [optional] 
**Volume** | Pointer to **string** | Volume in base currency | [optional] 
**NumTrades** | Pointer to **string** | Total number of trades across all markets | [optional] 
**PriceOutlier** | Pointer to **bool** | True when the candle is a price outlier.  May be null if outlier detection has not been applied to the candle. | [optional] 
**VolumeOutlier** | Pointer to **bool** | True when the candle is a volume outlier.  May be null if outlier detection has not been applied to the candle. | [optional] 

## Methods

### NewInlineResponse20020

`func NewInlineResponse20020() *InlineResponse20020`

NewInlineResponse20020 instantiates a new InlineResponse20020 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20020WithDefaults

`func NewInlineResponse20020WithDefaults() *InlineResponse20020`

NewInlineResponse20020WithDefaults instantiates a new InlineResponse20020 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *InlineResponse20020) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse20020) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse20020) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InlineResponse20020) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLow

`func (o *InlineResponse20020) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *InlineResponse20020) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *InlineResponse20020) SetLow(v string)`

SetLow sets Low field to given value.

### HasLow

`func (o *InlineResponse20020) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetOpen

`func (o *InlineResponse20020) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *InlineResponse20020) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *InlineResponse20020) SetOpen(v string)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *InlineResponse20020) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetClose

`func (o *InlineResponse20020) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *InlineResponse20020) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *InlineResponse20020) SetClose(v string)`

SetClose sets Close field to given value.

### HasClose

`func (o *InlineResponse20020) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetHigh

`func (o *InlineResponse20020) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *InlineResponse20020) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *InlineResponse20020) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *InlineResponse20020) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetVolume

`func (o *InlineResponse20020) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *InlineResponse20020) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *InlineResponse20020) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *InlineResponse20020) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetNumTrades

`func (o *InlineResponse20020) GetNumTrades() string`

GetNumTrades returns the NumTrades field if non-nil, zero value otherwise.

### GetNumTradesOk

`func (o *InlineResponse20020) GetNumTradesOk() (*string, bool)`

GetNumTradesOk returns a tuple with the NumTrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTrades

`func (o *InlineResponse20020) SetNumTrades(v string)`

SetNumTrades sets NumTrades field to given value.

### HasNumTrades

`func (o *InlineResponse20020) HasNumTrades() bool`

HasNumTrades returns a boolean if a field has been set.

### GetPriceOutlier

`func (o *InlineResponse20020) GetPriceOutlier() bool`

GetPriceOutlier returns the PriceOutlier field if non-nil, zero value otherwise.

### GetPriceOutlierOk

`func (o *InlineResponse20020) GetPriceOutlierOk() (*bool, bool)`

GetPriceOutlierOk returns a tuple with the PriceOutlier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceOutlier

`func (o *InlineResponse20020) SetPriceOutlier(v bool)`

SetPriceOutlier sets PriceOutlier field to given value.

### HasPriceOutlier

`func (o *InlineResponse20020) HasPriceOutlier() bool`

HasPriceOutlier returns a boolean if a field has been set.

### GetVolumeOutlier

`func (o *InlineResponse20020) GetVolumeOutlier() bool`

GetVolumeOutlier returns the VolumeOutlier field if non-nil, zero value otherwise.

### GetVolumeOutlierOk

`func (o *InlineResponse20020) GetVolumeOutlierOk() (*bool, bool)`

GetVolumeOutlierOk returns a tuple with the VolumeOutlier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeOutlier

`func (o *InlineResponse20020) SetVolumeOutlier(v bool)`

SetVolumeOutlier sets VolumeOutlier field to given value.

### HasVolumeOutlier

`func (o *InlineResponse20020) HasVolumeOutlier() bool`

HasVolumeOutlier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


