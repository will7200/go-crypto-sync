# OCOOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderListId** | Pointer to **int32** |  | [optional] 
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Orders** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewOCOOrder

`func NewOCOOrder() *OCOOrder`

NewOCOOrder instantiates a new OCOOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCOOrderWithDefaults

`func NewOCOOrderWithDefaults() *OCOOrder`

NewOCOOrderWithDefaults instantiates a new OCOOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *OCOOrder) GetOrderListId() int32`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OCOOrder) GetOrderListIdOk() (*int32, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OCOOrder) SetOrderListId(v int32)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OCOOrder) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *OCOOrder) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *OCOOrder) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *OCOOrder) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *OCOOrder) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *OCOOrder) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *OCOOrder) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *OCOOrder) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *OCOOrder) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *OCOOrder) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *OCOOrder) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *OCOOrder) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *OCOOrder) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *OCOOrder) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *OCOOrder) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *OCOOrder) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *OCOOrder) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *OCOOrder) GetTransactionTime() int32`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *OCOOrder) GetTransactionTimeOk() (*int32, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *OCOOrder) SetTransactionTime(v int32)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *OCOOrder) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *OCOOrder) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OCOOrder) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OCOOrder) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OCOOrder) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrders

`func (o *OCOOrder) GetOrders() []map[string]interface{}`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OCOOrder) GetOrdersOk() (*[]map[string]interface{}, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OCOOrder) SetOrders(v []map[string]interface{})`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OCOOrder) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


