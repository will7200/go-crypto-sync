# \OrdersApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrderBookBatches**](OrdersApi.md#GetOrderBookBatches) | **Get** /orders/batches | Order Book Batches
[**GetOrderBookSnapshot**](OrdersApi.md#GetOrderBookSnapshot) | **Get** /orders/snapshot | Order Book Snapshot



## GetOrderBookBatches

> string GetOrderBookBatches(ctx).Exchange(exchange).Market(market).Date(date).Execute()

Order Book Batches



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/will7200/go-crypto-sync/pkg/nomics"
)

func main() {
    exchange := "exchange_example" // string | Exchange ID
    market := "market_example" // string | The Exchange's Market ID from [Markets](#tag/Markets)
    date := "date_example" // string | The date for which to export a batch.  Must be a date ending at most 2 days prior to the current date.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrderBookBatches(context.Background(), exchange, market).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderBookBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderBookBatches`: string
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderBookBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderBookBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchange** | **string** | Exchange ID | 
 **market** | **string** | The Exchange&#39;s Market ID from [Markets](#tag/Markets) | 
 **date** | **string** | The date for which to export a batch.  Must be a date ending at most 2 days prior to the current date.  | 

### Return type

**string**

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderBookSnapshot

> InlineResponse20018 GetOrderBookSnapshot(ctx).Exchange(exchange).Market(market).At(at).Format(format).Execute()

Order Book Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/will7200/go-crypto-sync/pkg/nomics"
)

func main() {
    exchange := "exchange_example" // string | Exchange ID
    market := "market_example" // string | The Exchange's Market ID from [Markets](#tag/Markets)
    at := "at_example" // string | The newest order book *before* this timestamp will be returned. Timestamp should be in RFC3339 (URI escaped). (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrderBookSnapshot(context.Background(), exchange, market).At(at).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderBookSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderBookSnapshot`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderBookSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderBookSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchange** | **string** | Exchange ID | 
 **market** | **string** | The Exchange&#39;s Market ID from [Markets](#tag/Markets) | 
 **at** | **string** | The newest order book *before* this timestamp will be returned. Timestamp should be in RFC3339 (URI escaped). | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

