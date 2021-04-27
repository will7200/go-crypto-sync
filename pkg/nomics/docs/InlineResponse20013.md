# InlineResponse20013

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | Pointer to **string** | Nomics Exchange ID | [optional] 
**Base** | Pointer to **string** | Nomics Currency ID of the asset being traded | [optional] 
**Quote** | Pointer to **string** | Nomics Currency ID of the asset used to quote a price | [optional] 
**VolumeBase** | Pointer to **string** | Total volume in the base currency | [optional] 
**VolumeUsd** | Pointer to **string** | Total volume in USD | [optional] 
**OpenQuote** | Pointer to **string** | Open price in the quote currency | [optional] 
**OpenTimestamp** | Pointer to **string** | Timestamp of the open price in RFC 3339 | [optional] 
**CloseQuote** | Pointer to **string** | Close price in the quote currency | [optional] 
**CloseTimestamp** | Pointer to **string** | Timestamp of the close price in RFC 3339 | [optional] 
**NumTrades** | Pointer to **string** | Total number of trades | [optional] 

## Methods

### NewInlineResponse20013

`func NewInlineResponse20013() *InlineResponse20013`

NewInlineResponse20013 instantiates a new InlineResponse20013 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20013WithDefaults

`func NewInlineResponse20013WithDefaults() *InlineResponse20013`

NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *InlineResponse20013) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InlineResponse20013) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InlineResponse20013) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InlineResponse20013) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetBase

`func (o *InlineResponse20013) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *InlineResponse20013) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *InlineResponse20013) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *InlineResponse20013) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetQuote

`func (o *InlineResponse20013) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *InlineResponse20013) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *InlineResponse20013) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *InlineResponse20013) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetVolumeBase

`func (o *InlineResponse20013) GetVolumeBase() string`

GetVolumeBase returns the VolumeBase field if non-nil, zero value otherwise.

### GetVolumeBaseOk

`func (o *InlineResponse20013) GetVolumeBaseOk() (*string, bool)`

GetVolumeBaseOk returns a tuple with the VolumeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBase

`func (o *InlineResponse20013) SetVolumeBase(v string)`

SetVolumeBase sets VolumeBase field to given value.

### HasVolumeBase

`func (o *InlineResponse20013) HasVolumeBase() bool`

HasVolumeBase returns a boolean if a field has been set.

### GetVolumeUsd

`func (o *InlineResponse20013) GetVolumeUsd() string`

GetVolumeUsd returns the VolumeUsd field if non-nil, zero value otherwise.

### GetVolumeUsdOk

`func (o *InlineResponse20013) GetVolumeUsdOk() (*string, bool)`

GetVolumeUsdOk returns a tuple with the VolumeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUsd

`func (o *InlineResponse20013) SetVolumeUsd(v string)`

SetVolumeUsd sets VolumeUsd field to given value.

### HasVolumeUsd

`func (o *InlineResponse20013) HasVolumeUsd() bool`

HasVolumeUsd returns a boolean if a field has been set.

### GetOpenQuote

`func (o *InlineResponse20013) GetOpenQuote() string`

GetOpenQuote returns the OpenQuote field if non-nil, zero value otherwise.

### GetOpenQuoteOk

`func (o *InlineResponse20013) GetOpenQuoteOk() (*string, bool)`

GetOpenQuoteOk returns a tuple with the OpenQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenQuote

`func (o *InlineResponse20013) SetOpenQuote(v string)`

SetOpenQuote sets OpenQuote field to given value.

### HasOpenQuote

`func (o *InlineResponse20013) HasOpenQuote() bool`

HasOpenQuote returns a boolean if a field has been set.

### GetOpenTimestamp

`func (o *InlineResponse20013) GetOpenTimestamp() string`

GetOpenTimestamp returns the OpenTimestamp field if non-nil, zero value otherwise.

### GetOpenTimestampOk

`func (o *InlineResponse20013) GetOpenTimestampOk() (*string, bool)`

GetOpenTimestampOk returns a tuple with the OpenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTimestamp

`func (o *InlineResponse20013) SetOpenTimestamp(v string)`

SetOpenTimestamp sets OpenTimestamp field to given value.

### HasOpenTimestamp

`func (o *InlineResponse20013) HasOpenTimestamp() bool`

HasOpenTimestamp returns a boolean if a field has been set.

### GetCloseQuote

`func (o *InlineResponse20013) GetCloseQuote() string`

GetCloseQuote returns the CloseQuote field if non-nil, zero value otherwise.

### GetCloseQuoteOk

`func (o *InlineResponse20013) GetCloseQuoteOk() (*string, bool)`

GetCloseQuoteOk returns a tuple with the CloseQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseQuote

`func (o *InlineResponse20013) SetCloseQuote(v string)`

SetCloseQuote sets CloseQuote field to given value.

### HasCloseQuote

`func (o *InlineResponse20013) HasCloseQuote() bool`

HasCloseQuote returns a boolean if a field has been set.

### GetCloseTimestamp

`func (o *InlineResponse20013) GetCloseTimestamp() string`

GetCloseTimestamp returns the CloseTimestamp field if non-nil, zero value otherwise.

### GetCloseTimestampOk

`func (o *InlineResponse20013) GetCloseTimestampOk() (*string, bool)`

GetCloseTimestampOk returns a tuple with the CloseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTimestamp

`func (o *InlineResponse20013) SetCloseTimestamp(v string)`

SetCloseTimestamp sets CloseTimestamp field to given value.

### HasCloseTimestamp

`func (o *InlineResponse20013) HasCloseTimestamp() bool`

HasCloseTimestamp returns a boolean if a field has been set.

### GetNumTrades

`func (o *InlineResponse20013) GetNumTrades() string`

GetNumTrades returns the NumTrades field if non-nil, zero value otherwise.

### GetNumTradesOk

`func (o *InlineResponse20013) GetNumTradesOk() (*string, bool)`

GetNumTradesOk returns a tuple with the NumTrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTrades

`func (o *InlineResponse20013) SetNumTrades(v string)`

SetNumTrades sets NumTrades field to given value.

### HasNumTrades

`func (o *InlineResponse20013) HasNumTrades() bool`

HasNumTrades returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


