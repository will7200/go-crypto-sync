# \TradesApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrades**](TradesApi.md#GetTrades) | **Get** /trades | Trades



## GetTrades

> []Trade GetTrades(ctx).Exchange(exchange).Market(market).Limit(limit).Order(order).From(from).Format(format).Execute()

Trades



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
    limit := 987 // int32 | Maximum number of trades returned (optional) (default to 100)
    order := "order_example" // string | Defines the sort order of returned trades and the direction from `from`.  (optional) (default to "desc")
    from := "from_example" // string | Timestamp from which results should start in RFC3339. Please ensure you URI encode the timestamp. From is inclusive.  (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradesApi.GetTrades(context.Background(), exchange, market).Limit(limit).Order(order).From(from).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradesApi.GetTrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrades`: []Trade
    fmt.Fprintf(os.Stdout, "Response from `TradesApi.GetTrades`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchange** | **string** | Exchange ID | 
 **market** | **string** | The Exchange&#39;s Market ID from [Markets](#tag/Markets) | 
 **limit** | **int32** | Maximum number of trades returned | [default to 100]
 **order** | **string** | Defines the sort order of returned trades and the direction from &#x60;from&#x60;.  | [default to &quot;desc&quot;]
 **from** | **string** | Timestamp from which results should start in RFC3339. Please ensure you URI encode the timestamp. From is inclusive.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]Trade**](Trade.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

