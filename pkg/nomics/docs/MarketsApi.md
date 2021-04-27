# \MarketsApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExchangeMarketsTicker**](MarketsApi.md#GetExchangeMarketsTicker) | **Get** /exchange-markets/ticker | Exchange Markets Ticker
[**GetMarketCapHistory**](MarketsApi.md#GetMarketCapHistory) | **Get** /market-cap/history | Market Cap History
[**GetMarkets**](MarketsApi.md#GetMarkets) | **Get** /markets | Markets



## GetExchangeMarketsTicker

> []InlineResponse20010 GetExchangeMarketsTicker(ctx).Interval(interval).Currency(currency).Base(base).Quote(quote).Exchange(exchange).Market(market).Convert(convert).Status(status).Search(search).PerPage(perPage).Page(page).Execute()

Exchange Markets Ticker



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
    interval := "interval_example" // string | Comma separated time interval of the ticker(s). (optional) (default to "1d")
    currency := "currency_example" // string | A comma separated list of Nomics Currency IDs. A market must have one of the currencies as either its base or quote currency to be included. (optional)
    base := "base_example" // string | A comma separated list of Nomics Currency IDs. A market must have one of the currencies as its base currency to be included. (optional)
    quote := "quote_example" // string | A comma separated list of Nomics Currency IDs. A market must have one of the currencies as its quote currency to be included. (optional)
    exchange := "exchange_example" // string | A comma separated list of Nomics Exchange IDs. A market must be on one of the exchanges to be included. (optional)
    market := "market_example" // string | A comma separated list of Nomics Market IDs. A market's `market_id` must be one of the provided markets to be included. (optional)
    convert := "convert_example" // string | Nomics Currency ID to convert all financial data to (optional) (default to "USD")
    status := "status_example" // string | Status by which to filter markets. All markets are shown by default, regardless of status.  (optional)
    search := "search_example" // string | Search string by which to filter markets.  Items match the search query if it's a substring of their exchange id, market id, base currency symbol or quote currency symbol.  (optional)
    perPage := 987 // int32 | The maximum number of items to return per paginated response. Paginated responses include an additional response header, `X-Pagination-Total-Items`, which represents the total number of items available after all the request filters have been applied.  Must be between `1` and `100` (inclusive).  (optional)
    page := 987 // int32 | Which page of items to get.  Only applicable when `per-page` is also supplied.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketsApi.GetExchangeMarketsTicker(context.Background(), ).Interval(interval).Currency(currency).Base(base).Quote(quote).Exchange(exchange).Market(market).Convert(convert).Status(status).Search(search).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GetExchangeMarketsTicker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeMarketsTicker`: []InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `MarketsApi.GetExchangeMarketsTicker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeMarketsTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | **string** | Comma separated time interval of the ticker(s). | [default to &quot;1d&quot;]
 **currency** | **string** | A comma separated list of Nomics Currency IDs. A market must have one of the currencies as either its base or quote currency to be included. | 
 **base** | **string** | A comma separated list of Nomics Currency IDs. A market must have one of the currencies as its base currency to be included. | 
 **quote** | **string** | A comma separated list of Nomics Currency IDs. A market must have one of the currencies as its quote currency to be included. | 
 **exchange** | **string** | A comma separated list of Nomics Exchange IDs. A market must be on one of the exchanges to be included. | 
 **market** | **string** | A comma separated list of Nomics Market IDs. A market&#39;s &#x60;market_id&#x60; must be one of the provided markets to be included. | 
 **convert** | **string** | Nomics Currency ID to convert all financial data to | [default to &quot;USD&quot;]
 **status** | **string** | Status by which to filter markets. All markets are shown by default, regardless of status.  | 
 **search** | **string** | Search string by which to filter markets.  Items match the search query if it&#39;s a substring of their exchange id, market id, base currency symbol or quote currency symbol.  | 
 **perPage** | **int32** | The maximum number of items to return per paginated response. Paginated responses include an additional response header, &#x60;X-Pagination-Total-Items&#x60;, which represents the total number of items available after all the request filters have been applied.  Must be between &#x60;1&#x60; and &#x60;100&#x60; (inclusive).  | 
 **page** | **int32** | Which page of items to get.  Only applicable when &#x60;per-page&#x60; is also supplied.  | 

### Return type

[**[]InlineResponse20010**](inline_response_200_10.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketCapHistory

> []InlineResponse20014 GetMarketCapHistory(ctx).Start(start).End(end).Convert(convert).Format(format).IncludeTransparency(includeTransparency).Execute()

Market Cap History



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
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  (optional)
    convert := "convert_example" // string | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is `USD`. (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)
    includeTransparency := false // bool | Whether to include [Transparent Volume](https://blog.nomics.com/essays/transparent-volume/) and transparent market cap information in the response. Default is `false`. This option is only available to customers of our paid API plans.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketsApi.GetMarketCapHistory(context.Background(), start).End(end).Convert(convert).Format(format).IncludeTransparency(includeTransparency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GetMarketCapHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketCapHistory`: []InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `MarketsApi.GetMarketCapHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketCapHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped) | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **convert** | **string** | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is &#x60;USD&#x60;. | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 
 **includeTransparency** | **bool** | Whether to include [Transparent Volume](https://blog.nomics.com/essays/transparent-volume/) and transparent market cap information in the response. Default is &#x60;false&#x60;. This option is only available to customers of our paid API plans.  | 

### Return type

[**[]InlineResponse20014**](inline_response_200_14.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, test/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarkets

> []InlineResponse2009 GetMarkets(ctx).Exchange(exchange).Base(base).Quote(quote).Format(format).Execute()

Markets



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
    exchange := "exchange_example" // string | Nomics Exchange ID to filter by (optional)
    base := "base_example" // string | Comma separated list of base currencies to filter by (optional)
    quote := "quote_example" // string | Comma separated list of quote currencies to filter by (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketsApi.GetMarkets(context.Background(), ).Exchange(exchange).Base(base).Quote(quote).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GetMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarkets`: []InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `MarketsApi.GetMarkets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchange** | **string** | Nomics Exchange ID to filter by | 
 **base** | **string** | Comma separated list of base currencies to filter by | 
 **quote** | **string** | Comma separated list of quote currencies to filter by | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse2009**](inline_response_200_9.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

