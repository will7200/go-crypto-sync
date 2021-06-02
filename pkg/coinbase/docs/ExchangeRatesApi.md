# \ExchangeRatesApi

All URIs are relative to *https://api.coinbase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExchangeRateFor**](ExchangeRatesApi.md#GetExchangeRateFor) | **Get** /exchange-rates | Get current exchange rates. Default base currency is USD but it can be defined as any supported currency. Returned rates will define the exchange rate for one unit of the base currency.
[**GetSellPrice**](ExchangeRatesApi.md#GetSellPrice) | **Get** /prices/{currency_pair_1}-{currency_pair_2}/sell | Exchanges Rates for currency pair



## GetExchangeRateFor

> InlineResponse2002 GetExchangeRateFor(ctx).Currency(currency).Execute()

Get current exchange rates. Default base currency is USD but it can be defined as any supported currency. Returned rates will define the exchange rate for one unit of the base currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/will7200/go-crypto-sync/pkg/coinbase"
)

func main() {
    currency := "currency_example" // string | Base currency (default: USD)  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangeRatesApi.GetExchangeRateFor(context.Background(), ).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeRatesApi.GetExchangeRateFor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeRateFor`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `ExchangeRatesApi.GetExchangeRateFor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRateForRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Base currency (default: USD)  | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSellPrice

> InlineResponse2003 GetSellPrice(ctx, currencyPair1, currencyPair2).Execute()

Exchanges Rates for currency pair



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/will7200/go-crypto-sync/pkg/coinbase"
)

func main() {
    currencyPair1 := "currencyPair1_example" // string | currency 1 
    currencyPair2 := "currencyPair2_example" // string | currency 2 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangeRatesApi.GetSellPrice(context.Background(), currencyPair1, currencyPair2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeRatesApi.GetSellPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSellPrice`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `ExchangeRatesApi.GetSellPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair1** | **string** | currency 1  | 
**currencyPair2** | **string** | currency 2  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSellPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

