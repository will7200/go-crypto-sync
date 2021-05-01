# \TradeApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3AccountGet**](TradeApi.md#ApiV3AccountGet) | **Get** /api/v3/account | Account Information (USER_DATA)
[**ApiV3AllOrderListGet**](TradeApi.md#ApiV3AllOrderListGet) | **Get** /api/v3/allOrderList | Query all OCO (USER_DATA)
[**ApiV3AllOrdersGet**](TradeApi.md#ApiV3AllOrdersGet) | **Get** /api/v3/allOrders | All Orders (USER_DATA)
[**ApiV3MyTradesGet**](TradeApi.md#ApiV3MyTradesGet) | **Get** /api/v3/myTrades | Account Trade List (USER_DATA)
[**ApiV3OpenOrderListGet**](TradeApi.md#ApiV3OpenOrderListGet) | **Get** /api/v3/openOrderList | Query Open OCO (USER_DATA)
[**ApiV3OpenOrdersDelete**](TradeApi.md#ApiV3OpenOrdersDelete) | **Delete** /api/v3/openOrders | Cancel all Open Orders on a Symbol (TRADE)
[**ApiV3OpenOrdersGet**](TradeApi.md#ApiV3OpenOrdersGet) | **Get** /api/v3/openOrders | Current Open Orders (USER_DATA)
[**ApiV3OrderDelete**](TradeApi.md#ApiV3OrderDelete) | **Delete** /api/v3/order | Cancel Order (TRADE)
[**ApiV3OrderGet**](TradeApi.md#ApiV3OrderGet) | **Get** /api/v3/order | Query Order (USER_DATA)
[**ApiV3OrderListDelete**](TradeApi.md#ApiV3OrderListDelete) | **Delete** /api/v3/orderList | Cancel OCO (TRADE)
[**ApiV3OrderListGet**](TradeApi.md#ApiV3OrderListGet) | **Get** /api/v3/orderList | Query OCO (USER_DATA)
[**ApiV3OrderOcoPost**](TradeApi.md#ApiV3OrderOcoPost) | **Post** /api/v3/order/oco | New OCO (TRADE)
[**ApiV3OrderPost**](TradeApi.md#ApiV3OrderPost) | **Post** /api/v3/order | New Order (TRADE)
[**ApiV3OrderTestPost**](TradeApi.md#ApiV3OrderTestPost) | **Post** /api/v3/order/test | Test New Order (TRADE)



## ApiV3AccountGet

> Account ApiV3AccountGet(ctx).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Account Information (USER_DATA)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3AccountGet(context.Background()).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3AccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3AccountGet`: Account
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3AccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3AccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3AllOrderListGet

> []OCOOrder ApiV3AllOrderListGet(ctx).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query all OCO (USER_DATA)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fromId := int32(56) // int32 | Trade id to fetch from. Default gets most recent trades. (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3AllOrderListGet(context.Background()).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3AllOrderListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3AllOrderListGet`: []OCOOrder
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3AllOrderListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3AllOrderListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromId** | **int32** | Trade id to fetch from. Default gets most recent trades. | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**[]OCOOrder**](OCOOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3AllOrdersGet

> []OrderDetails ApiV3AllOrdersGet(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

All Orders (USER_DATA)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    orderId := int32(56) // int32 | order id (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3AllOrdersGet(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3AllOrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3AllOrdersGet`: []OrderDetails
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3AllOrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3AllOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **orderId** | **int32** | order id | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**[]OrderDetails**](OrderDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3MyTradesGet

> MyTrade ApiV3MyTradesGet(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Account Trade List (USER_DATA)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    fromId := int32(56) // int32 | Trade id to fetch from. Default gets most recent trades. (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3MyTradesGet(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3MyTradesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3MyTradesGet`: MyTrade
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3MyTradesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3MyTradesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **fromId** | **int32** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**MyTrade**](MyTrade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OpenOrderListGet

> []OCOOrder ApiV3OpenOrderListGet(ctx).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Open OCO (USER_DATA)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OpenOrderListGet(context.Background()).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OpenOrderListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OpenOrderListGet`: []OCOOrder
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OpenOrderListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OpenOrderListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**[]OCOOrder**](OCOOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OpenOrdersDelete

> []Order ApiV3OpenOrdersDelete(ctx).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Cancel all Open Orders on a Symbol (TRADE)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OpenOrdersDelete(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OpenOrdersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OpenOrdersDelete`: []Order
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OpenOrdersDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OpenOrdersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**[]Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OpenOrdersGet

> []OrderDetails ApiV3OpenOrdersGet(ctx).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Current Open Orders (USER_DATA)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OpenOrdersGet(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OpenOrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OpenOrdersGet`: []OrderDetails
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OpenOrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OpenOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**[]OrderDetails**](OrderDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderDelete

> Order ApiV3OrderDelete(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Cancel Order (TRADE)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    orderId := int32(56) // int32 | order id (optional)
    origClientOrderId := "origClientOrderId_example" // string | order id from client (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderDelete(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderDelete`: Order
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **orderId** | **int32** | order id | 
 **origClientOrderId** | **string** | order id from client | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderGet

> OrderDetails ApiV3OrderGet(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Order (USER_DATA)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    orderId := int32(56) // int32 | order id (optional)
    origClientOrderId := "origClientOrderId_example" // string | order id from client (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderGet(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderGet`: OrderDetails
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **orderId** | **int32** | order id | 
 **origClientOrderId** | **string** | order id from client | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**OrderDetails**](OrderDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderListDelete

> OCOOrderReport ApiV3OrderListDelete(ctx).Symbol(symbol).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Cancel OCO (TRADE)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    orderListId := int32(56) // int32 | order list id (optional)
    listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderListDelete(context.Background()).Symbol(symbol).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderListDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderListDelete`: OCOOrderReport
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderListDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderListDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **orderListId** | **int32** | order list id | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**OCOOrderReport**](OCOOrderReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderListGet

> OCOOrder ApiV3OrderListGet(ctx).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query OCO (USER_DATA)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderListId := int32(56) // int32 | order list id (optional)
    origClientOrderId := "origClientOrderId_example" // string | order id from client (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderListGet(context.Background()).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderListGet`: OCOOrder
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderListId** | **int32** | order list id | 
 **origClientOrderId** | **string** | order id from client | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**OCOOrder**](OCOOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderOcoPost

> InlineResponse2004 ApiV3OrderOcoPost(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

New OCO (TRADE)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    side := "BNBUSDT" // string | SELL or BUY
    quantity := float32(10.01) // float32 | order quantity
    price := float32(3.4) // float32 | 
    stopPrice := float32(3.4) // float32 | 
    listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
    limitClientOrderId := float32(3.4) // float32 | A unique Id for the limit order (optional)
    limitIcebergQty := float32(3.4) // float32 |  (optional)
    stopClientOrderId := "stopClientOrderId_example" // string |  (optional)
    stopLimitPrice := float32(3.4) // float32 |  (optional)
    stopIcebergQty := float32(3.4) // float32 |  (optional)
    stopLimitTimeInForce := "stopLimitTimeInForce_example" // string |  (optional)
    newOrderRespType := "newOrderRespType_example" // string | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderOcoPost(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderOcoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderOcoPost`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderOcoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderOcoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **side** | **string** | SELL or BUY | 
 **quantity** | **float32** | order quantity | 
 **price** | **float32** |  | 
 **stopPrice** | **float32** |  | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **limitClientOrderId** | **float32** | A unique Id for the limit order | 
 **limitIcebergQty** | **float32** |  | 
 **stopClientOrderId** | **string** |  | 
 **stopLimitPrice** | **float32** |  | 
 **stopIcebergQty** | **float32** |  | 
 **stopLimitTimeInForce** | **string** |  | 
 **newOrderRespType** | **string** | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderPost

> OrderResponse ApiV3OrderPost(ctx).Symbol(symbol).Side(side).Type_(type_).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

New Order (TRADE)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    side := "BNBUSDT" // string | SELL or BUY
    type_ := "type__example" // string | the order type
    timeInForce := "timeInForce_example" // string | order time in force (optional)
    quantity := float32(3.4) // float32 | order quantity (optional)
    quoteOrderQty := float32(3.4) // float32 | quote quantity (optional)
    price := float32(20.01) // float32 | order price (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    stopPrice := float32(20.01) // float32 | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. (optional)
    icebergQty := float32(3.4) // float32 | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. (optional)
    newOrderRespType := "newOrderRespType_example" // string | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderPost(context.Background()).Symbol(symbol).Side(side).Type_(type_).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderPost`: OrderResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **side** | **string** | SELL or BUY | 
 **type_** | **string** | the order type | 
 **timeInForce** | **string** | order time in force | 
 **quantity** | **float32** | order quantity | 
 **quoteOrderQty** | **float32** | quote quantity | 
 **price** | **float32** | order price | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **stopPrice** | **float32** | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. | 
 **icebergQty** | **float32** | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. | 
 **newOrderRespType** | **string** | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderTestPost

> ApiV3OrderTestPost(ctx).Symbol(symbol).Side(side).Type_(type_).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Test New Order (TRADE)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    side := "BNBUSDT" // string | SELL or BUY
    type_ := "type__example" // string | the order type
    timeInForce := "timeInForce_example" // string | order time in force (optional)
    quantity := float32(3.4) // float32 | order quantity (optional)
    quoteOrderQty := float32(3.4) // float32 | quote quantity (optional)
    price := float32(20.01) // float32 | order price (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    stopPrice := float32(20.01) // float32 | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. (optional)
    icebergQty := float32(3.4) // float32 | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. (optional)
    newOrderRespType := "newOrderRespType_example" // string | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderTestPost(context.Background()).Symbol(symbol).Side(side).Type_(type_).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderTestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **side** | **string** | SELL or BUY | 
 **type_** | **string** | the order type | 
 **timeInForce** | **string** | order time in force | 
 **quantity** | **float32** | order quantity | 
 **quoteOrderQty** | **float32** | quote quantity | 
 **price** | **float32** | order price | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **stopPrice** | **float32** | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. | 
 **icebergQty** | **float32** | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. | 
 **newOrderRespType** | **string** | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

