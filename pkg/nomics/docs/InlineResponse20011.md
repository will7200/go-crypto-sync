# InlineResponse20011

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | Pointer to **string** | Nomics ID of the exchange | [optional] 
**Quote** | Pointer to **string** | Nomics ID of the quote currency | [optional] 
**Price** | Pointer to **string** | Price in USD of the most recent trade | [optional] 
**Timestamp** | Pointer to **string** | RFC3339 Timestamp of the most recent trade | [optional] 

## Methods

### NewInlineResponse20011

`func NewInlineResponse20011() *InlineResponse20011`

NewInlineResponse20011 instantiates a new InlineResponse20011 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20011WithDefaults

`func NewInlineResponse20011WithDefaults() *InlineResponse20011`

NewInlineResponse20011WithDefaults instantiates a new InlineResponse20011 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *InlineResponse20011) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InlineResponse20011) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InlineResponse20011) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InlineResponse20011) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetQuote

`func (o *InlineResponse20011) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *InlineResponse20011) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *InlineResponse20011) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *InlineResponse20011) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetPrice

`func (o *InlineResponse20011) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InlineResponse20011) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InlineResponse20011) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InlineResponse20011) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTimestamp

`func (o *InlineResponse20011) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse20011) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse20011) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InlineResponse20011) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


