# \MarginApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1MarginAccountGet**](MarginApi.md#SapiV1MarginAccountGet) | **Get** /sapi/v1/margin/account | Query Margin Account Details (USER_DATA)
[**SapiV1MarginAllAssetsGet**](MarginApi.md#SapiV1MarginAllAssetsGet) | **Get** /sapi/v1/margin/allAssets | Get All Margin Assets (MARKET_DATA)
[**SapiV1MarginAllOrdersGet**](MarginApi.md#SapiV1MarginAllOrdersGet) | **Get** /sapi/v1/margin/allOrders | Query Margin Account&#39;s All Order (USER_DATA)
[**SapiV1MarginAllPairsGet**](MarginApi.md#SapiV1MarginAllPairsGet) | **Get** /sapi/v1/margin/allPairs | Get All Margin Pairs (MARKET_DATA)
[**SapiV1MarginAssetGet**](MarginApi.md#SapiV1MarginAssetGet) | **Get** /sapi/v1/margin/asset | Query Margin Asset (MARKET_DATA)
[**SapiV1MarginForceLiquidationRecGet**](MarginApi.md#SapiV1MarginForceLiquidationRecGet) | **Get** /sapi/v1/margin/forceLiquidationRec | Get Force Liquidation Record (USER_DATA)
[**SapiV1MarginInterestHistoryGet**](MarginApi.md#SapiV1MarginInterestHistoryGet) | **Get** /sapi/v1/margin/interestHistory | Query Interest History (MARKET_DATA)
[**SapiV1MarginLoanGet**](MarginApi.md#SapiV1MarginLoanGet) | **Get** /sapi/v1/margin/loan | Query Load Record (USER_DATA)
[**SapiV1MarginLoanPost**](MarginApi.md#SapiV1MarginLoanPost) | **Post** /sapi/v1/margin/loan | Margin Account Borrow (MARGIN)
[**SapiV1MarginMaxBorrowableGet**](MarginApi.md#SapiV1MarginMaxBorrowableGet) | **Get** /sapi/v1/margin/maxBorrowable | Query Max Borrow (USER_DATA)
[**SapiV1MarginMaxTransferableGet**](MarginApi.md#SapiV1MarginMaxTransferableGet) | **Get** /sapi/v1/margin/maxTransferable | Query Max Transfer-Out Amount (USER_DATA)
[**SapiV1MarginMyTradesGet**](MarginApi.md#SapiV1MarginMyTradesGet) | **Get** /sapi/v1/margin/myTrades | If fromId is set, it will get orders &gt;&#x3D; that fromId. Otherwise most recent orders are returned.
[**SapiV1MarginOpenOrdersGet**](MarginApi.md#SapiV1MarginOpenOrdersGet) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#39;s Open Order (USER_DATA)
[**SapiV1MarginOrderDelete**](MarginApi.md#SapiV1MarginOrderDelete) | **Delete** /sapi/v1/margin/order | Margin Account Cancel Order (TRADE)
[**SapiV1MarginOrderGet**](MarginApi.md#SapiV1MarginOrderGet) | **Get** /sapi/v1/margin/order | Query Margin Account&#39;s Order (MARKET_DATA)
[**SapiV1MarginOrderPost**](MarginApi.md#SapiV1MarginOrderPost) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**SapiV1MarginPairGet**](MarginApi.md#SapiV1MarginPairGet) | **Get** /sapi/v1/margin/pair | Query Margin Pair (MARKET_DATA)
[**SapiV1MarginPriceIndexGet**](MarginApi.md#SapiV1MarginPriceIndexGet) | **Get** /sapi/v1/margin/priceIndex | Query Margin PriceIndex (MARKET_DATA)
[**SapiV1MarginRepayGet**](MarginApi.md#SapiV1MarginRepayGet) | **Get** /sapi/v1/margin/repay | Query Repay Record (USER_DATA)
[**SapiV1MarginRepayPost**](MarginApi.md#SapiV1MarginRepayPost) | **Post** /sapi/v1/margin/repay | Margin Account Repay (MARGIN)
[**SapiV1MarginTransferGet**](MarginApi.md#SapiV1MarginTransferGet) | **Get** /sapi/v1/margin/transfer | Get Transfer History (USER_DATA)
[**SapiV1MarginTransferPost**](MarginApi.md#SapiV1MarginTransferPost) | **Post** /sapi/v1/margin/transfer | Margin Account Transfer (MARGIN)



## SapiV1MarginAccountGet

> InlineResponse20012 SapiV1MarginAccountGet(ctx).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Margin Account Details (USER_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginAccountGet(context.Background()).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAccountGet`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAllAssetsGet

> []InlineResponse2008 SapiV1MarginAllAssetsGet(ctx).Execute()

Get All Margin Assets (MARKET_DATA)



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginAllAssetsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAllAssetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAllAssetsGet`: []InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAllAssetsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAllAssetsGetRequest struct via the builder pattern


### Return type

[**[]InlineResponse2008**](InlineResponse2008.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAllOrdersGet

> []MarginOrderDetail SapiV1MarginAllOrdersGet(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Margin Account's All Order (USER_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginAllOrdersGet(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAllOrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAllOrdersGet`: []MarginOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAllOrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAllOrdersGetRequest struct via the builder pattern


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

[**[]MarginOrderDetail**](MarginOrderDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAllPairsGet

> []InlineResponse2009 SapiV1MarginAllPairsGet(ctx).Execute()

Get All Margin Pairs (MARKET_DATA)



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginAllPairsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAllPairsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAllPairsGet`: []InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAllPairsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAllPairsGetRequest struct via the builder pattern


### Return type

[**[]InlineResponse2009**](InlineResponse2009.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAssetGet

> InlineResponse2006 SapiV1MarginAssetGet(ctx).Asset(asset).Execute()

Query Margin Asset (MARKET_DATA)



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
    asset := "BTC" // string | The asset being transferred, e.g., BTC

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginAssetGet(context.Background()).Asset(asset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAssetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAssetGet`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAssetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAssetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being transferred, e.g., BTC | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginForceLiquidationRecGet

> map[string]interface{} SapiV1MarginForceLiquidationRecGet(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Get Force Liquidation Record (USER_DATA)



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
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    current := float32(1) // float32 | Currently querying page. Start from 1. Default:1 (optional)
    size := float32(100) // float32 | Default:10 Max:100 (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginForceLiquidationRecGet(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginForceLiquidationRecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginForceLiquidationRecGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginForceLiquidationRecGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginForceLiquidationRecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **current** | **float32** | Currently querying page. Start from 1. Default:1 | 
 **size** | **float32** | Default:10 Max:100 | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginInterestHistoryGet

> InlineResponse20011 SapiV1MarginInterestHistoryGet(ctx).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Interest History (MARKET_DATA)



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
    asset := "BNB" // string | The asset being transferred, e.g., BTC (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    current := float32(1) // float32 | Currently querying page. Start from 1. Default:1 (optional)
    size := float32(100) // float32 | Default:10 Max:100 (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginInterestHistoryGet(context.Background()).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginInterestHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginInterestHistoryGet`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginInterestHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginInterestHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being transferred, e.g., BTC | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **current** | **float32** | Currently querying page. Start from 1. Default:1 | 
 **size** | **float32** | Default:10 Max:100 | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginLoanGet

> map[string]interface{} SapiV1MarginLoanGet(ctx).Asset(asset).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Load Record (USER_DATA)



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
    asset := "BTC" // string | The asset being transferred, e.g., BTC
    txId := int32(123456789) // int32 | the tranId in  `POST /sapi/v1/margin/loan` (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    current := float32(1) // float32 | Currently querying page. Start from 1. Default:1 (optional)
    size := float32(100) // float32 | Default:10 Max:100 (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginLoanGet(context.Background()).Asset(asset).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginLoanGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginLoanGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginLoanGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginLoanGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being transferred, e.g., BTC | 
 **txId** | **int32** | the tranId in  &#x60;POST /sapi/v1/margin/loan&#x60; | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **current** | **float32** | Currently querying page. Start from 1. Default:1 | 
 **size** | **float32** | Default:10 Max:100 | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginLoanPost

> Transaction SapiV1MarginLoanPost(ctx).Asset(asset).Amount(amount).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Margin Account Borrow (MARGIN)



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
    asset := "BTC" // string | The asset being transferred, e.g., BTC
    amount := float32(1.01) // float32 | 
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginLoanPost(context.Background()).Asset(asset).Amount(amount).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginLoanPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginLoanPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginLoanPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginLoanPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being transferred, e.g., BTC | 
 **amount** | **float32** |  | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginMaxBorrowableGet

> map[string]interface{} SapiV1MarginMaxBorrowableGet(ctx).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Max Borrow (USER_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginMaxBorrowableGet(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginMaxBorrowableGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginMaxBorrowableGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginMaxBorrowableGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginMaxBorrowableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginMaxTransferableGet

> map[string]interface{} SapiV1MarginMaxTransferableGet(ctx).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Max Transfer-Out Amount (USER_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginMaxTransferableGet(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginMaxTransferableGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginMaxTransferableGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginMaxTransferableGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginMaxTransferableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginMyTradesGet

> []MarginTrade SapiV1MarginMyTradesGet(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

If fromId is set, it will get orders >= that fromId. Otherwise most recent orders are returned.



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
    resp, r, err := api_client.MarginApi.SapiV1MarginMyTradesGet(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginMyTradesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginMyTradesGet`: []MarginTrade
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginMyTradesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginMyTradesGetRequest struct via the builder pattern


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

[**[]MarginTrade**](MarginTrade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOpenOrdersGet

> []MarginOrderDetail SapiV1MarginOpenOrdersGet(ctx).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Margin Account's Open Order (USER_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginOpenOrdersGet(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOpenOrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOpenOrdersGet`: []MarginOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOpenOrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOpenOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**[]MarginOrderDetail**](MarginOrderDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOrderDelete

> MarginOrder SapiV1MarginOrderDelete(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Margin Account Cancel Order (TRADE)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginOrderDelete(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOrderDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOrderDelete`: MarginOrder
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOrderDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOrderDeleteRequest struct via the builder pattern


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

[**MarginOrder**](MarginOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOrderGet

> MarginOrderDetail SapiV1MarginOrderGet(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Margin Account's Order (MARKET_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginOrderGet(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOrderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOrderGet`: MarginOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOrderGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOrderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **orderId** | **int32** | order id | 
 **origClientOrderId** | **string** | order id from client | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**MarginOrderDetail**](MarginOrderDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOrderPost

> MarginOrderResponse SapiV1MarginOrderPost(ctx).Symbol(symbol).Side(side).Type_(type_).Quantity(quantity).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Margin Account New Order (TRADE)



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
    quantity := float32(10.01) // float32 | order quantity
    price := float32(20.01) // float32 | order price (optional)
    stopPrice := float32(20.01) // float32 | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    icebergQty := float32(3.4) // float32 | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. (optional)
    newOrderRespType := "newOrderRespType_example" // string | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
    sideEffectType := "sideEffectType_example" // string | default NO_SIDE_EFFECT (optional)
    timeInForce := "timeInForce_example" // string | order time in force (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOrderPost(context.Background()).Symbol(symbol).Side(side).Type_(type_).Quantity(quantity).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOrderPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOrderPost`: MarginOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOrderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOrderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **side** | **string** | SELL or BUY | 
 **type_** | **string** | the order type | 
 **quantity** | **float32** | order quantity | 
 **price** | **float32** | order price | 
 **stopPrice** | **float32** | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **icebergQty** | **float32** | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. | 
 **newOrderRespType** | **string** | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **sideEffectType** | **string** | default NO_SIDE_EFFECT | 
 **timeInForce** | **string** | order time in force | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**MarginOrderResponse**](MarginOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginPairGet

> InlineResponse2007 SapiV1MarginPairGet(ctx).Symbol(symbol).Execute()

Query Margin Pair (MARKET_DATA)



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginPairGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginPairGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginPairGet`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginPairGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginPairGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginPriceIndexGet

> InlineResponse20010 SapiV1MarginPriceIndexGet(ctx).Symbol(symbol).Execute()

Query Margin PriceIndex (MARKET_DATA)



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginPriceIndexGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginPriceIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginPriceIndexGet`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginPriceIndexGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginPriceIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginRepayGet

> InlineResponse2005 SapiV1MarginRepayGet(ctx).Asset(asset).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Query Repay Record (USER_DATA)



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
    asset := "BTC" // string | The asset being transferred, e.g., BTC
    txId := int32(2970933056) // int32 | the tranId in  `POST /sapi/v1/margin/repay` (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    current := float32(1) // float32 | Currently querying page. Start from 1. Default:1 (optional)
    size := float32(100) // float32 | Default:10 Max:100 (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginRepayGet(context.Background()).Asset(asset).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginRepayGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginRepayGet`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginRepayGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginRepayGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being transferred, e.g., BTC | 
 **txId** | **int32** | the tranId in  &#x60;POST /sapi/v1/margin/repay&#x60; | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **current** | **float32** | Currently querying page. Start from 1. Default:1 | 
 **size** | **float32** | Default:10 Max:100 | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginRepayPost

> Transaction SapiV1MarginRepayPost(ctx).Asset(asset).Amount(amount).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Margin Account Repay (MARGIN)



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
    asset := "BTC" // string | The asset being transferred, e.g., BTC
    amount := float32(1.01) // float32 | 
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginRepayPost(context.Background()).Asset(asset).Amount(amount).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginRepayPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginRepayPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginRepayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginRepayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being transferred, e.g., BTC | 
 **amount** | **float32** |  | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginTransferGet

> map[string]interface{} SapiV1MarginTransferGet(ctx).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Get Transfer History (USER_DATA)



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
    asset := "BNB" // string | The asset being transferred, e.g., BTC (optional)
    type_ := "type__example" // string | Tranfer Type (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    current := float32(1) // float32 | Currently querying page. Start from 1. Default:1 (optional)
    size := float32(100) // float32 | Default:10 Max:100 (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginTransferGet(context.Background()).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginTransferGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginTransferGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginTransferGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginTransferGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being transferred, e.g., BTC | 
 **type_** | **string** | Tranfer Type | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **current** | **float32** | Currently querying page. Start from 1. Default:1 | 
 **size** | **float32** | Default:10 Max:100 | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginTransferPost

> Transaction SapiV1MarginTransferPost(ctx).Asset(asset).Amount(amount).Type_(type_).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

Margin Account Transfer (MARGIN)



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
    asset := "BTC" // string | The asset being transferred, e.g., BTC
    amount := float32(1.01) // float32 | 
    type_ := int32(56) // int32 | 1 -> transfer from main account to margin account \\ 2 -> transfer from margin account to main account   (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginTransferPost(context.Background()).Asset(asset).Amount(amount).Type_(type_).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginTransferPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginTransferPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginTransferPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginTransferPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being transferred, e.g., BTC | 
 **amount** | **float32** |  | 
 **type_** | **int32** | 1 -&gt; transfer from main account to margin account \\ 2 -&gt; transfer from margin account to main account   | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

