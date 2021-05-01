# Trade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | trade id | [optional] 
**Price** | Pointer to **string** | price | [optional] 
**Qty** | Pointer to **string** | amount of base asset | [optional] 
**QuoteQty** | Pointer to **string** | amount of quote asset | [optional] 
**Time** | Pointer to **int32** | trade timestamp | [optional] 
**IsBuyerMaker** | Pointer to **bool** |  | [optional] 
**IsBestMatch** | Pointer to **bool** |  | [optional] 

## Methods

### NewTrade

`func NewTrade() *Trade`

NewTrade instantiates a new Trade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeWithDefaults

`func NewTradeWithDefaults() *Trade`

NewTradeWithDefaults instantiates a new Trade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Trade) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trade) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trade) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Trade) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *Trade) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Trade) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Trade) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Trade) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *Trade) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *Trade) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *Trade) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *Trade) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetQuoteQty

`func (o *Trade) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *Trade) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *Trade) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.

### HasQuoteQty

`func (o *Trade) HasQuoteQty() bool`

HasQuoteQty returns a boolean if a field has been set.

### GetTime

`func (o *Trade) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Trade) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Trade) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *Trade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetIsBuyerMaker

`func (o *Trade) GetIsBuyerMaker() bool`

GetIsBuyerMaker returns the IsBuyerMaker field if non-nil, zero value otherwise.

### GetIsBuyerMakerOk

`func (o *Trade) GetIsBuyerMakerOk() (*bool, bool)`

GetIsBuyerMakerOk returns a tuple with the IsBuyerMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyerMaker

`func (o *Trade) SetIsBuyerMaker(v bool)`

SetIsBuyerMaker sets IsBuyerMaker field to given value.

### HasIsBuyerMaker

`func (o *Trade) HasIsBuyerMaker() bool`

HasIsBuyerMaker returns a boolean if a field has been set.

### GetIsBestMatch

`func (o *Trade) GetIsBestMatch() bool`

GetIsBestMatch returns the IsBestMatch field if non-nil, zero value otherwise.

### GetIsBestMatchOk

`func (o *Trade) GetIsBestMatchOk() (*bool, bool)`

GetIsBestMatchOk returns a tuple with the IsBestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBestMatch

`func (o *Trade) SetIsBestMatch(v bool)`

SetIsBestMatch sets IsBestMatch field to given value.

### HasIsBestMatch

`func (o *Trade) HasIsBestMatch() bool`

HasIsBestMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


