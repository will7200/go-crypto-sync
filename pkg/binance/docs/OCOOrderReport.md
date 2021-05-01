# OCOOrderReport

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
**OrderReports** | Pointer to [**[]Order**](Order.md) |  | [optional] 

## Methods

### NewOCOOrderReport

`func NewOCOOrderReport() *OCOOrderReport`

NewOCOOrderReport instantiates a new OCOOrderReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCOOrderReportWithDefaults

`func NewOCOOrderReportWithDefaults() *OCOOrderReport`

NewOCOOrderReportWithDefaults instantiates a new OCOOrderReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *OCOOrderReport) GetOrderListId() int32`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OCOOrderReport) GetOrderListIdOk() (*int32, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OCOOrderReport) SetOrderListId(v int32)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OCOOrderReport) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *OCOOrderReport) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *OCOOrderReport) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *OCOOrderReport) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *OCOOrderReport) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *OCOOrderReport) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *OCOOrderReport) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *OCOOrderReport) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *OCOOrderReport) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *OCOOrderReport) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *OCOOrderReport) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *OCOOrderReport) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *OCOOrderReport) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *OCOOrderReport) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *OCOOrderReport) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *OCOOrderReport) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *OCOOrderReport) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *OCOOrderReport) GetTransactionTime() int32`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *OCOOrderReport) GetTransactionTimeOk() (*int32, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *OCOOrderReport) SetTransactionTime(v int32)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *OCOOrderReport) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *OCOOrderReport) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OCOOrderReport) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OCOOrderReport) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OCOOrderReport) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrders

`func (o *OCOOrderReport) GetOrders() []map[string]interface{}`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OCOOrderReport) GetOrdersOk() (*[]map[string]interface{}, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OCOOrderReport) SetOrders(v []map[string]interface{})`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OCOOrderReport) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderReports

`func (o *OCOOrderReport) GetOrderReports() []Order`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *OCOOrderReport) GetOrderReportsOk() (*[]Order, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *OCOOrderReport) SetOrderReports(v []Order)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *OCOOrderReport) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


