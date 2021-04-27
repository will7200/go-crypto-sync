# \CurrenciesApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrencies**](CurrenciesApi.md#GetCurrencies) | **Get** /currencies | Currencies Metadata
[**GetCurrenciesSparkline**](CurrenciesApi.md#GetCurrenciesSparkline) | **Get** /currencies/sparkline | Currencies Sparkline
[**GetCurrenciesTicker**](CurrenciesApi.md#GetCurrenciesTicker) | **Get** /currencies/ticker | Currencies Ticker
[**GetSupplyHistory**](CurrenciesApi.md#GetSupplyHistory) | **Get** /supplies/history | Supply History



## GetCurrencies

> []InlineResponse2001 GetCurrencies(ctx).Ids(ids).Attributes(attributes).Format(format).Execute()

Currencies Metadata



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
    ids := "ids_example" // string | Comma separated list of Nomics Currency IDs to filter result rows  (optional)
    attributes := "attributes_example" // string | Comma separated list of currency attributes to filter result columns  (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.GetCurrencies(context.Background(), ).Ids(ids).Attributes(attributes).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetCurrencies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrencies`: []InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetCurrencies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of Nomics Currency IDs to filter result rows  | 
 **attributes** | **string** | Comma separated list of currency attributes to filter result columns  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrenciesSparkline

> []InlineResponse2002 GetCurrenciesSparkline(ctx).Start(start).Ids(ids).End(end).Convert(convert).Execute()

Currencies Sparkline



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
    start := "start_example" // string | Start time of the interval in RFC3339 (URI escaped)
    ids := "ids_example" // string | Comma separated list of Nomics Currency IDs to filter result rows  (optional)
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  (optional)
    convert := "convert_example" // string | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is `USD`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.GetCurrenciesSparkline(context.Background(), start).Ids(ids).End(end).Convert(convert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetCurrenciesSparkline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrenciesSparkline`: []InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetCurrenciesSparkline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesSparklineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped) | 
 **ids** | **string** | Comma separated list of Nomics Currency IDs to filter result rows  | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **convert** | **string** | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is &#x60;USD&#x60;. | 

### Return type

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrenciesTicker

> []InlineResponse200 GetCurrenciesTicker(ctx).Ids(ids).Interval(interval).QuoteCurrency(quoteCurrency).Convert(convert).Status(status).Filter(filter).Sort(sort).IncludeTransparency(includeTransparency).PerPage(perPage).Page(page).Execute()

Currencies Ticker



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
    ids := "ids_example" // string | Comma separated list of Nomics Currency IDs to filter result rows  (optional)
    interval := "interval_example" // string | Comma separated time interval of the ticker(s). Default is `1d,7d,30d,365d,ytd`. (optional)
    quoteCurrency := "quoteCurrency_example" // string | Currency to quote ticker price, market cap, and volume values. Must be a valid currency from [Exchange Rates](#operation/getExchangeRates). Default is `USD`. (optional)
    convert := "convert_example" // string | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is `USD`. (optional)
    status := "status_example" // string | Status by which to filter currencies. If not provided, all currencies are shown.  (optional)
    filter := "filter_example" // string | Further filter the set of currencies.  The `new` filter returns currencies that have recently been priced by Nomics and `any` returns currencies regardless of their state.  The `any` filer may be used to retrieve new-but-stale currencies that are listed under `new`, but are no longer active.  (optional)
    sort := "sort_example" // string | How to sort the returned currencies.  `rank` sorts by rank ascending and `first_priced_at` sorts by when each currency was first priced by Nomics descending.  (optional)
    includeTransparency := false // bool | Whether to include [Transparent Volume](https://blog.nomics.com/essays/transparent-volume/) and transparent market cap information in the response. Default is `false`. This option is only available to customers of our paid API plans.  (optional)
    perPage := 987 // int32 | The maximum number of items to return per paginated response. Paginated responses include an additional response header, `X-Pagination-Total-Items`, which represents the total number of items available after all the request filters have been applied.  Must be between `1` and `100` (inclusive).  (optional)
    page := 987 // int32 | Which page of items to get.  Only applicable when `per-page` is also supplied.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.GetCurrenciesTicker(context.Background(), ).Ids(ids).Interval(interval).QuoteCurrency(quoteCurrency).Convert(convert).Status(status).Filter(filter).Sort(sort).IncludeTransparency(includeTransparency).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetCurrenciesTicker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrenciesTicker`: []InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetCurrenciesTicker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of Nomics Currency IDs to filter result rows  | 
 **interval** | **string** | Comma separated time interval of the ticker(s). Default is &#x60;1d,7d,30d,365d,ytd&#x60;. | 
 **quoteCurrency** | **string** | Currency to quote ticker price, market cap, and volume values. Must be a valid currency from [Exchange Rates](#operation/getExchangeRates). Default is &#x60;USD&#x60;. | 
 **convert** | **string** | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is &#x60;USD&#x60;. | 
 **status** | **string** | Status by which to filter currencies. If not provided, all currencies are shown.  | 
 **filter** | **string** | Further filter the set of currencies.  The &#x60;new&#x60; filter returns currencies that have recently been priced by Nomics and &#x60;any&#x60; returns currencies regardless of their state.  The &#x60;any&#x60; filer may be used to retrieve new-but-stale currencies that are listed under &#x60;new&#x60;, but are no longer active.  | 
 **sort** | **string** | How to sort the returned currencies.  &#x60;rank&#x60; sorts by rank ascending and &#x60;first_priced_at&#x60; sorts by when each currency was first priced by Nomics descending.  | 
 **includeTransparency** | **bool** | Whether to include [Transparent Volume](https://blog.nomics.com/essays/transparent-volume/) and transparent market cap information in the response. Default is &#x60;false&#x60;. This option is only available to customers of our paid API plans.  | 
 **perPage** | **int32** | The maximum number of items to return per paginated response. Paginated responses include an additional response header, &#x60;X-Pagination-Total-Items&#x60;, which represents the total number of items available after all the request filters have been applied.  Must be between &#x60;1&#x60; and &#x60;100&#x60; (inclusive).  | 
 **page** | **int32** | Which page of items to get.  Only applicable when &#x60;per-page&#x60; is also supplied.  | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupplyHistory

> []InlineResponse2004 GetSupplyHistory(ctx).Currency(currency).Start(start).End(end).Format(format).Execute()

Supply History



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
    resp, r, err := api_client.CurrenciesApi.GetSupplyHistory(context.Background(), currency, start).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetSupplyHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupplyHistory`: []InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetSupplyHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Currency ID | 
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped) | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse2004**](inline_response_200_4.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

