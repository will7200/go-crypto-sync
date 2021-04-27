# \GlobalApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalTicker**](GlobalApi.md#GetGlobalTicker) | **Get** /global-ticker | Global Ticker



## GetGlobalTicker

> InlineResponse2003 GetGlobalTicker(ctx).Convert(convert).Execute()

Global Ticker



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
    convert := "convert_example" // string | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is `USD`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalApi.GetGlobalTicker(context.Background(), ).Convert(convert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApi.GetGlobalTicker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalTicker`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `GlobalApi.GetGlobalTicker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convert** | **string** | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is &#x60;USD&#x60;. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

