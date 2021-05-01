# OrderResponseFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**OrderListId** | Pointer to **int32** | Unless OCO, value will be -1 | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**TransactTime** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**OrigQty** | Pointer to **float32** |  | [optional] 
**ExecutedQty** | Pointer to **float32** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Fills** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOrderResponseFull

`func NewOrderResponseFull() *OrderResponseFull`

NewOrderResponseFull instantiates a new OrderResponseFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseFullWithDefaults

`func NewOrderResponseFullWithDefaults() *OrderResponseFull`

NewOrderResponseFullWithDefaults instantiates a new OrderResponseFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderResponseFull) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderResponseFull) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderResponseFull) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderResponseFull) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderResponseFull) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderResponseFull) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderResponseFull) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderResponseFull) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *OrderResponseFull) GetOrderListId() int32`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderResponseFull) GetOrderListIdOk() (*int32, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderResponseFull) SetOrderListId(v int32)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderResponseFull) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OrderResponseFull) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderResponseFull) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderResponseFull) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OrderResponseFull) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetTransactTime

`func (o *OrderResponseFull) GetTransactTime() int32`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *OrderResponseFull) GetTransactTimeOk() (*int32, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *OrderResponseFull) SetTransactTime(v int32)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *OrderResponseFull) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetPrice

`func (o *OrderResponseFull) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderResponseFull) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderResponseFull) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderResponseFull) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *OrderResponseFull) GetOrigQty() float32`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OrderResponseFull) GetOrigQtyOk() (*float32, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OrderResponseFull) SetOrigQty(v float32)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *OrderResponseFull) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OrderResponseFull) GetExecutedQty() float32`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OrderResponseFull) GetExecutedQtyOk() (*float32, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OrderResponseFull) SetExecutedQty(v float32)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OrderResponseFull) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *OrderResponseFull) GetCummulativeQuoteQty() float32`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *OrderResponseFull) GetCummulativeQuoteQtyOk() (*float32, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *OrderResponseFull) SetCummulativeQuoteQty(v float32)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *OrderResponseFull) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *OrderResponseFull) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderResponseFull) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderResponseFull) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderResponseFull) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OrderResponseFull) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OrderResponseFull) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OrderResponseFull) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OrderResponseFull) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OrderResponseFull) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderResponseFull) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderResponseFull) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderResponseFull) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *OrderResponseFull) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderResponseFull) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderResponseFull) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderResponseFull) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetFills

`func (o *OrderResponseFull) GetFills() []string`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *OrderResponseFull) GetFillsOk() (*[]string, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *OrderResponseFull) SetFills(v []string)`

SetFills sets Fills field to given value.

### HasFills

`func (o *OrderResponseFull) HasFills() bool`

HasFills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


