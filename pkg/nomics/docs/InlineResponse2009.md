# InlineResponse2009

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | Pointer to **string** | Nomics Exchange ID | [optional] 
**Market** | Pointer to **string** | The Exchange&#39;s Market ID (usable on Exchange&#39;s API) | [optional] 
**Base** | Pointer to **string** | Nomics Currency ID of the asset being traded | [optional] 
**Quote** | Pointer to **string** | Nomics Currency ID of the asset used to quote a price | [optional] 

## Methods

### NewInlineResponse2009

`func NewInlineResponse2009() *InlineResponse2009`

NewInlineResponse2009 instantiates a new InlineResponse2009 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2009WithDefaults

`func NewInlineResponse2009WithDefaults() *InlineResponse2009`

NewInlineResponse2009WithDefaults instantiates a new InlineResponse2009 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *InlineResponse2009) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InlineResponse2009) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InlineResponse2009) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InlineResponse2009) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMarket

`func (o *InlineResponse2009) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *InlineResponse2009) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *InlineResponse2009) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *InlineResponse2009) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetBase

`func (o *InlineResponse2009) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *InlineResponse2009) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *InlineResponse2009) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *InlineResponse2009) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetQuote

`func (o *InlineResponse2009) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *InlineResponse2009) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *InlineResponse2009) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *InlineResponse2009) HasQuote() bool`

HasQuote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


