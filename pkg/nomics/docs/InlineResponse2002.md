# InlineResponse2002

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency ID | [optional] 
**Timestamps** | Pointer to **[]string** | Time values matching the price value of the same index | [optional] 
**Prices** | Pointer to **[]string** | Price in USD corresponding to timestamp value of the same index | [optional] 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002() *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InlineResponse2002) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InlineResponse2002) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InlineResponse2002) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InlineResponse2002) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTimestamps

`func (o *InlineResponse2002) GetTimestamps() []string`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *InlineResponse2002) GetTimestampsOk() (*[]string, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *InlineResponse2002) SetTimestamps(v []string)`

SetTimestamps sets Timestamps field to given value.

### HasTimestamps

`func (o *InlineResponse2002) HasTimestamps() bool`

HasTimestamps returns a boolean if a field has been set.

### GetPrices

`func (o *InlineResponse2002) GetPrices() []string`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *InlineResponse2002) GetPricesOk() (*[]string, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *InlineResponse2002) SetPrices(v []string)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *InlineResponse2002) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


