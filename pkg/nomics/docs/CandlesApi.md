# \CandlesApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCandles**](CandlesApi.md#GetCandles) | **Get** /candles | Aggregated OHLCV Candles
[**GetExchangeCandles**](CandlesApi.md#GetExchangeCandles) | **Get** /exchange_candles | Exchange OHLCV Candles
[**GetMarketCandles**](CandlesApi.md#GetMarketCandles) | **Get** /markets/candles | Aggregated Pair OHLCV Candles



## GetCandles

> []InlineResponse20015 GetCandles(ctx).Interval(interval).Currency(currency).Start(start).End(end).Format(format).Execute()

Aggregated OHLCV Candles



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
    interval := "interval_example" // string | Time interval of the candle
    currency := "currency_example" // string | Currency ID
    start := "start_example" // string | Start time of the interval in RFC3339 (URI escaped). If not provided, starts from the first candle for daily candles and from the current time minus 30 days for hourly candles.  (optional)
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used except for when hourly candles with a `start` time older than 30 days are requested.  In that case, the `end` time defaults to the `start` time plus 7 days.  (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CandlesApi.GetCandles(context.Background(), interval, currency).Start(start).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CandlesApi.GetCandles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCandles`: []InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `CandlesApi.GetCandles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCandlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | **string** | Time interval of the candle | 
 **currency** | **string** | Currency ID | 
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped). If not provided, starts from the first candle for daily candles and from the current time minus 30 days for hourly candles.  | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used except for when hourly candles with a &#x60;start&#x60; time older than 30 days are requested.  In that case, the &#x60;end&#x60; time defaults to the &#x60;start&#x60; time plus 7 days.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20015**](inline_response_200_15.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeCandles

> []InlineResponse20019 GetExchangeCandles(ctx).Interval(interval).Exchange(exchange).Market(market).Start(start).End(end).Format(format).Execute()

Exchange OHLCV Candles



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
    interval := "interval_example" // string | Time interval of the candle
    exchange := "exchange_example" // string | Exchange ID
    market := "market_example" // string | The Exchange's Market ID from [Markets](#tag/Markets)
    start := "start_example" // string | Start time of the interval in RFC3339 (URI escaped).  If not provided, then candles will be returned starting from the current time minus the date range limit for the requested interval.  For example, if you request hourly candles without specifying a start time, then candles spanning the last 30 days will be returned.  (optional)
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CandlesApi.GetExchangeCandles(context.Background(), interval, exchange, market).Start(start).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CandlesApi.GetExchangeCandles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeCandles`: []InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `CandlesApi.GetExchangeCandles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeCandlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | **string** | Time interval of the candle | 
 **exchange** | **string** | Exchange ID | 
 **market** | **string** | The Exchange&#39;s Market ID from [Markets](#tag/Markets) | 
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped).  If not provided, then candles will be returned starting from the current time minus the date range limit for the requested interval.  For example, if you request hourly candles without specifying a start time, then candles spanning the last 30 days will be returned.  | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20019**](inline_response_200_19.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketCandles

> []InlineResponse20020 GetMarketCandles(ctx).Interval(interval).Base(base).Quote(quote).Start(start).End(end).Format(format).Execute()

Aggregated Pair OHLCV Candles



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
    interval := "interval_example" // string | Time interval of the candle
    base := "base_example" // string | Base currency of the pair
    quote := "quote_example" // string | Quote currency of the pair
    start := "start_example" // string | Start time of the interval in RFC3339 (URI escaped).  If not provided, then candles will be returned starting from the current time minus the date range limit for the requested interval.  For example, if you request hourly candles without specifying a start time, then candles spanning the last 30 days will be returned.  (optional)
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CandlesApi.GetMarketCandles(context.Background(), interval, base, quote).Start(start).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CandlesApi.GetMarketCandles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketCandles`: []InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `CandlesApi.GetMarketCandles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketCandlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | **string** | Time interval of the candle | 
 **base** | **string** | Base currency of the pair | 
 **quote** | **string** | Quote currency of the pair | 
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped).  If not provided, then candles will be returned starting from the current time minus the date range limit for the requested interval.  For example, if you request hourly candles without specifying a start time, then candles spanning the last 30 days will be returned.  | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20020**](inline_response_200_20.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

