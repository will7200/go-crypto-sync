# MarginOrderResponseFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
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
**MarginBuyBorrowAmount** | Pointer to **float32** |  | [optional] 
**MarginBuyBorrowAsset** | Pointer to **string** |  | [optional] 
**Fills** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMarginOrderResponseFull

`func NewMarginOrderResponseFull() *MarginOrderResponseFull`

NewMarginOrderResponseFull instantiates a new MarginOrderResponseFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginOrderResponseFullWithDefaults

`func NewMarginOrderResponseFullWithDefaults() *MarginOrderResponseFull`

NewMarginOrderResponseFullWithDefaults instantiates a new MarginOrderResponseFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MarginOrderResponseFull) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginOrderResponseFull) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginOrderResponseFull) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginOrderResponseFull) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *MarginOrderResponseFull) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginOrderResponseFull) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginOrderResponseFull) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MarginOrderResponseFull) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *MarginOrderResponseFull) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *MarginOrderResponseFull) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *MarginOrderResponseFull) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *MarginOrderResponseFull) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetTransactTime

`func (o *MarginOrderResponseFull) GetTransactTime() int32`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *MarginOrderResponseFull) GetTransactTimeOk() (*int32, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *MarginOrderResponseFull) SetTransactTime(v int32)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *MarginOrderResponseFull) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetPrice

`func (o *MarginOrderResponseFull) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginOrderResponseFull) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginOrderResponseFull) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarginOrderResponseFull) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *MarginOrderResponseFull) GetOrigQty() float32`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *MarginOrderResponseFull) GetOrigQtyOk() (*float32, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *MarginOrderResponseFull) SetOrigQty(v float32)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *MarginOrderResponseFull) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *MarginOrderResponseFull) GetExecutedQty() float32`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *MarginOrderResponseFull) GetExecutedQtyOk() (*float32, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *MarginOrderResponseFull) SetExecutedQty(v float32)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *MarginOrderResponseFull) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *MarginOrderResponseFull) GetCummulativeQuoteQty() float32`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *MarginOrderResponseFull) GetCummulativeQuoteQtyOk() (*float32, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *MarginOrderResponseFull) SetCummulativeQuoteQty(v float32)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *MarginOrderResponseFull) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *MarginOrderResponseFull) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarginOrderResponseFull) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarginOrderResponseFull) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarginOrderResponseFull) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *MarginOrderResponseFull) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *MarginOrderResponseFull) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *MarginOrderResponseFull) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *MarginOrderResponseFull) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *MarginOrderResponseFull) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarginOrderResponseFull) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarginOrderResponseFull) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MarginOrderResponseFull) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *MarginOrderResponseFull) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *MarginOrderResponseFull) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *MarginOrderResponseFull) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *MarginOrderResponseFull) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetMarginBuyBorrowAmount

`func (o *MarginOrderResponseFull) GetMarginBuyBorrowAmount() float32`

GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAmountOk

`func (o *MarginOrderResponseFull) GetMarginBuyBorrowAmountOk() (*float32, bool)`

GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAmount

`func (o *MarginOrderResponseFull) SetMarginBuyBorrowAmount(v float32)`

SetMarginBuyBorrowAmount sets MarginBuyBorrowAmount field to given value.

### HasMarginBuyBorrowAmount

`func (o *MarginOrderResponseFull) HasMarginBuyBorrowAmount() bool`

HasMarginBuyBorrowAmount returns a boolean if a field has been set.

### GetMarginBuyBorrowAsset

`func (o *MarginOrderResponseFull) GetMarginBuyBorrowAsset() string`

GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAssetOk

`func (o *MarginOrderResponseFull) GetMarginBuyBorrowAssetOk() (*string, bool)`

GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAsset

`func (o *MarginOrderResponseFull) SetMarginBuyBorrowAsset(v string)`

SetMarginBuyBorrowAsset sets MarginBuyBorrowAsset field to given value.

### HasMarginBuyBorrowAsset

`func (o *MarginOrderResponseFull) HasMarginBuyBorrowAsset() bool`

HasMarginBuyBorrowAsset returns a boolean if a field has been set.

### GetFills

`func (o *MarginOrderResponseFull) GetFills() []string`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *MarginOrderResponseFull) GetFillsOk() (*[]string, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *MarginOrderResponseFull) SetFills(v []string)`

SetFills sets Fills field to given value.

### HasFills

`func (o *MarginOrderResponseFull) HasFills() bool`

HasFills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


