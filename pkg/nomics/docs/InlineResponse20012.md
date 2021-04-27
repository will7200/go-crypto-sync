# InlineResponse20012

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | Pointer to **string** | Nomics ID of the exchange | [optional] 
**Base** | Pointer to **string** | Nomics ID of the base currency | [optional] 
**Quote** | Pointer to **string** | Nomics ID of the quote currency | [optional] 
**PriceQuote** | Pointer to **string** | Price in the quote currency of the most recent trade | [optional] 
**Timestamp** | Pointer to **string** | RFC3339 Timestamp of the most recent trade | [optional] 

## Methods

### NewInlineResponse20012

`func NewInlineResponse20012() *InlineResponse20012`

NewInlineResponse20012 instantiates a new InlineResponse20012 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20012WithDefaults

`func NewInlineResponse20012WithDefaults() *InlineResponse20012`

NewInlineResponse20012WithDefaults instantiates a new InlineResponse20012 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *InlineResponse20012) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InlineResponse20012) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InlineResponse20012) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InlineResponse20012) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetBase

`func (o *InlineResponse20012) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *InlineResponse20012) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *InlineResponse20012) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *InlineResponse20012) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetQuote

`func (o *InlineResponse20012) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *InlineResponse20012) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *InlineResponse20012) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *InlineResponse20012) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetPriceQuote

`func (o *InlineResponse20012) GetPriceQuote() string`

GetPriceQuote returns the PriceQuote field if non-nil, zero value otherwise.

### GetPriceQuoteOk

`func (o *InlineResponse20012) GetPriceQuoteOk() (*string, bool)`

GetPriceQuoteOk returns a tuple with the PriceQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceQuote

`func (o *InlineResponse20012) SetPriceQuote(v string)`

SetPriceQuote sets PriceQuote field to given value.

### HasPriceQuote

`func (o *InlineResponse20012) HasPriceQuote() bool`

HasPriceQuote returns a boolean if a field has been set.

### GetTimestamp

`func (o *InlineResponse20012) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse20012) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse20012) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InlineResponse20012) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


