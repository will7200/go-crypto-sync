# Trade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Exchange&#39;s ID of the trade (or synthetic ID when not available) | [optional] 
**Timestamp** | Pointer to **string** | Timestamp in RFC3339 | [optional] 
**Price** | Pointer to **string** | Price of the trade in the quote currency of the market | [optional] 
**Volume** | Pointer to **string** | Volume of the trade in the base currency of the market | [optional] 

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

`func (o *Trade) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trade) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trade) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Trade) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *Trade) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Trade) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Trade) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Trade) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

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

### GetVolume

`func (o *Trade) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Trade) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Trade) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Trade) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


