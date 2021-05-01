# OrderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**OrderListId** | Pointer to **int32** | Unless OCO, value will be -1 | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**OrigQty** | Pointer to **float32** |  | [optional] 
**ExecutedQty** | Pointer to **float32** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **float32** |  | [optional] 
**IcebergQty** | Pointer to **float32** |  | [optional] 
**Time** | Pointer to **int32** |  | [optional] 
**UpdateTime** | Pointer to **int32** |  | [optional] 
**IsWorking** | Pointer to **bool** |  | [optional] 
**OrigQuoteOrderQty** | Pointer to **float32** |  | [optional] 

## Methods

### NewOrderDetails

`func NewOrderDetails() *OrderDetails`

NewOrderDetails instantiates a new OrderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailsWithDefaults

`func NewOrderDetailsWithDefaults() *OrderDetails`

NewOrderDetailsWithDefaults instantiates a new OrderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderDetails) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderDetails) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderDetails) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderDetails) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *OrderDetails) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *OrderDetails) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *OrderDetails) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *OrderDetails) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderDetails) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderDetails) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderDetails) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderDetails) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *OrderDetails) GetOrderListId() int32`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderDetails) GetOrderListIdOk() (*int32, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderDetails) SetOrderListId(v int32)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderDetails) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetPrice

`func (o *OrderDetails) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderDetails) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderDetails) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderDetails) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *OrderDetails) GetOrigQty() float32`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OrderDetails) GetOrigQtyOk() (*float32, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OrderDetails) SetOrigQty(v float32)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *OrderDetails) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OrderDetails) GetExecutedQty() float32`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OrderDetails) GetExecutedQtyOk() (*float32, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OrderDetails) SetExecutedQty(v float32)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OrderDetails) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *OrderDetails) GetCummulativeQuoteQty() float32`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *OrderDetails) GetCummulativeQuoteQtyOk() (*float32, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *OrderDetails) SetCummulativeQuoteQty(v float32)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *OrderDetails) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *OrderDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OrderDetails) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OrderDetails) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OrderDetails) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OrderDetails) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OrderDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *OrderDetails) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderDetails) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderDetails) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderDetails) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStopPrice

`func (o *OrderDetails) GetStopPrice() float32`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OrderDetails) GetStopPriceOk() (*float32, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OrderDetails) SetStopPrice(v float32)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OrderDetails) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetIcebergQty

`func (o *OrderDetails) GetIcebergQty() float32`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *OrderDetails) GetIcebergQtyOk() (*float32, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *OrderDetails) SetIcebergQty(v float32)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *OrderDetails) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetTime

`func (o *OrderDetails) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OrderDetails) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OrderDetails) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *OrderDetails) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *OrderDetails) GetUpdateTime() int32`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *OrderDetails) GetUpdateTimeOk() (*int32, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *OrderDetails) SetUpdateTime(v int32)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *OrderDetails) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetIsWorking

`func (o *OrderDetails) GetIsWorking() bool`

GetIsWorking returns the IsWorking field if non-nil, zero value otherwise.

### GetIsWorkingOk

`func (o *OrderDetails) GetIsWorkingOk() (*bool, bool)`

GetIsWorkingOk returns a tuple with the IsWorking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorking

`func (o *OrderDetails) SetIsWorking(v bool)`

SetIsWorking sets IsWorking field to given value.

### HasIsWorking

`func (o *OrderDetails) HasIsWorking() bool`

HasIsWorking returns a boolean if a field has been set.

### GetOrigQuoteOrderQty

`func (o *OrderDetails) GetOrigQuoteOrderQty() float32`

GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field if non-nil, zero value otherwise.

### GetOrigQuoteOrderQtyOk

`func (o *OrderDetails) GetOrigQuoteOrderQtyOk() (*float32, bool)`

GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQuoteOrderQty

`func (o *OrderDetails) SetOrigQuoteOrderQty(v float32)`

SetOrigQuoteOrderQty sets OrigQuoteOrderQty field to given value.

### HasOrigQuoteOrderQty

`func (o *OrderDetails) HasOrigQuoteOrderQty() bool`

HasOrigQuoteOrderQty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


