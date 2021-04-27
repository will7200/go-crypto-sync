# \ExchangeRatesApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExchangeRates**](ExchangeRatesApi.md#GetExchangeRates) | **Get** /exchange-rates | Exchange Rates
[**GetExchangeRatesHistory**](ExchangeRatesApi.md#GetExchangeRatesHistory) | **Get** /exchange-rates/history | Exchange Rates History



## GetExchangeRates

> []ExchangeRate GetExchangeRates(ctx).Format(format).Execute()

Exchange Rates



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
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangeRatesApi.GetExchangeRates(context.Background(), ).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeRatesApi.GetExchangeRates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeRates`: []ExchangeRate
    fmt.Fprintf(os.Stdout, "Response from `ExchangeRatesApi.GetExchangeRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]ExchangeRate**](ExchangeRate.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeRatesHistory

> []InlineResponse20017 GetExchangeRatesHistory(ctx).Currency(currency).Start(start).End(end).Format(format).Execute()

Exchange Rates History



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
    currency := "currency_example" // string | Currency ID
    start := "start_example" // string | Start time of the interval in RFC3339 (URI escaped)
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangeRatesApi.GetExchangeRatesHistory(context.Background(), currency, start).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeRatesApi.GetExchangeRatesHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeRatesHistory`: []InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `ExchangeRatesApi.GetExchangeRatesHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRatesHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Currency ID | 
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped) | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20017**](inline_response_200_17.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

