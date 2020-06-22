# \ExchangeRatesApi

All URIs are relative to *https://api.coinbase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExchangeRateFor**](ExchangeRatesApi.md#GetExchangeRateFor) | **Get** /prices/{currency_pair_1}-{currency_pair_2}/sell | Exchanges Rates for currency pair



## GetExchangeRateFor

> InlineResponse2001 GetExchangeRateFor(ctx, currencyPair1, currencyPair2).Execute()

Exchanges Rates for currency pair



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
    currencyPair1 := "currencyPair1_example" // string | currency 1 
    currencyPair2 := "currencyPair2_example" // string | currency 2 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangeRatesApi.GetExchangeRateFor(context.Background(), currencyPair1, currencyPair2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeRatesApi.GetExchangeRateFor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeRateFor`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ExchangeRatesApi.GetExchangeRateFor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair1** | **string** | currency 1  | 
**currencyPair2** | **string** | currency 2  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRateForRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

