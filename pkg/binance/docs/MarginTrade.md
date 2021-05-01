# MarginTrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsBestMatch** | Pointer to **bool** |  | [optional] 
**IsBuyer** | Pointer to **bool** |  | [optional] 
**IsMaker** | Pointer to **bool** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int32** |  | [optional] 

## Methods

### NewMarginTrade

`func NewMarginTrade() *MarginTrade`

NewMarginTrade instantiates a new MarginTrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginTradeWithDefaults

`func NewMarginTradeWithDefaults() *MarginTrade`

NewMarginTradeWithDefaults instantiates a new MarginTrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommission

`func (o *MarginTrade) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *MarginTrade) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *MarginTrade) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *MarginTrade) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *MarginTrade) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *MarginTrade) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *MarginTrade) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *MarginTrade) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetId

`func (o *MarginTrade) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarginTrade) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarginTrade) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MarginTrade) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBestMatch

`func (o *MarginTrade) GetIsBestMatch() bool`

GetIsBestMatch returns the IsBestMatch field if non-nil, zero value otherwise.

### GetIsBestMatchOk

`func (o *MarginTrade) GetIsBestMatchOk() (*bool, bool)`

GetIsBestMatchOk returns a tuple with the IsBestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBestMatch

`func (o *MarginTrade) SetIsBestMatch(v bool)`

SetIsBestMatch sets IsBestMatch field to given value.

### HasIsBestMatch

`func (o *MarginTrade) HasIsBestMatch() bool`

HasIsBestMatch returns a boolean if a field has been set.

### GetIsBuyer

`func (o *MarginTrade) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *MarginTrade) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *MarginTrade) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.

### HasIsBuyer

`func (o *MarginTrade) HasIsBuyer() bool`

HasIsBuyer returns a boolean if a field has been set.

### GetIsMaker

`func (o *MarginTrade) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *MarginTrade) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *MarginTrade) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.

### HasIsMaker

`func (o *MarginTrade) HasIsMaker() bool`

HasIsMaker returns a boolean if a field has been set.

### GetOrderId

`func (o *MarginTrade) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginTrade) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginTrade) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MarginTrade) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *MarginTrade) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginTrade) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginTrade) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarginTrade) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *MarginTrade) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *MarginTrade) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *MarginTrade) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *MarginTrade) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginTrade) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginTrade) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginTrade) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginTrade) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *MarginTrade) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MarginTrade) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MarginTrade) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *MarginTrade) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


