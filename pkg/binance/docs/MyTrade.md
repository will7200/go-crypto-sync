# MyTrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | trade id | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**OrderListId** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **string** | price | [optional] 
**Qty** | Pointer to **string** | amount of base asset | [optional] 
**QuoteQty** | Pointer to **string** | amount of quote asset | [optional] 
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int32** | trade timestamp | [optional] 
**IsBuyer** | Pointer to **bool** |  | [optional] 
**IsMaker** | Pointer to **bool** |  | [optional] 
**IsBestMatch** | Pointer to **bool** |  | [optional] 

## Methods

### NewMyTrade

`func NewMyTrade() *MyTrade`

NewMyTrade instantiates a new MyTrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyTradeWithDefaults

`func NewMyTradeWithDefaults() *MyTrade`

NewMyTradeWithDefaults instantiates a new MyTrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MyTrade) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MyTrade) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MyTrade) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MyTrade) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetId

`func (o *MyTrade) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyTrade) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyTrade) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MyTrade) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *MyTrade) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MyTrade) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MyTrade) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MyTrade) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *MyTrade) GetOrderListId() int32`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MyTrade) GetOrderListIdOk() (*int32, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MyTrade) SetOrderListId(v int32)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MyTrade) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetPrice

`func (o *MyTrade) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MyTrade) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MyTrade) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MyTrade) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *MyTrade) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *MyTrade) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *MyTrade) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *MyTrade) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetQuoteQty

`func (o *MyTrade) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *MyTrade) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *MyTrade) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.

### HasQuoteQty

`func (o *MyTrade) HasQuoteQty() bool`

HasQuoteQty returns a boolean if a field has been set.

### GetCommission

`func (o *MyTrade) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *MyTrade) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *MyTrade) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *MyTrade) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *MyTrade) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *MyTrade) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *MyTrade) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *MyTrade) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetTime

`func (o *MyTrade) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MyTrade) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MyTrade) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *MyTrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetIsBuyer

`func (o *MyTrade) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *MyTrade) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *MyTrade) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.

### HasIsBuyer

`func (o *MyTrade) HasIsBuyer() bool`

HasIsBuyer returns a boolean if a field has been set.

### GetIsMaker

`func (o *MyTrade) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *MyTrade) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *MyTrade) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.

### HasIsMaker

`func (o *MyTrade) HasIsMaker() bool`

HasIsMaker returns a boolean if a field has been set.

### GetIsBestMatch

`func (o *MyTrade) GetIsBestMatch() bool`

GetIsBestMatch returns the IsBestMatch field if non-nil, zero value otherwise.

### GetIsBestMatchOk

`func (o *MyTrade) GetIsBestMatchOk() (*bool, bool)`

GetIsBestMatchOk returns a tuple with the IsBestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBestMatch

`func (o *MyTrade) SetIsBestMatch(v bool)`

SetIsBestMatch sets IsBestMatch field to given value.

### HasIsBestMatch

`func (o *MyTrade) HasIsBestMatch() bool`

HasIsBestMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


