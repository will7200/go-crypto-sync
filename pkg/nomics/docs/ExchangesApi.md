# \ExchangesApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExchangeVolumeHistory**](ExchangesApi.md#GetExchangeVolumeHistory) | **Get** /exchanges/volume/history | Exchanges Volume History
[**GetExchanges**](ExchangesApi.md#GetExchanges) | **Get** /exchanges | Exchanges Metadata
[**GetExchangesTicker**](ExchangesApi.md#GetExchangesTicker) | **Get** /exchanges/ticker | Exchanges Ticker



## GetExchangeVolumeHistory

> []InlineResponse2007 GetExchangeVolumeHistory(ctx).Exchange(exchange).Start(start).End(end).Convert(convert).Format(format).IncludeTransparency(includeTransparency).Execute()

Exchanges Volume History



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
    start := "start_example" // string | Start time of the interval in RFC3339 (URI escaped)
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  (optional)
    convert := "convert_example" // string | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is `USD`. (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)
    includeTransparency := false // bool | Whether to include [Transparent Volume](https://blog.nomics.com/essays/transparent-volume/) and transparent market cap information in the response. Default is `false`. This option is only available to customers of our paid API plans.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangesApi.GetExchangeVolumeHistory(context.Background(), exchange, start).End(end).Convert(convert).Format(format).IncludeTransparency(includeTransparency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.GetExchangeVolumeHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeVolumeHistory`: []InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.GetExchangeVolumeHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeVolumeHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchange** | **string** | Exchange ID | 
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped) | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **convert** | **string** | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is &#x60;USD&#x60;. | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 
 **includeTransparency** | **bool** | Whether to include [Transparent Volume](https://blog.nomics.com/essays/transparent-volume/) and transparent market cap information in the response. Default is &#x60;false&#x60;. This option is only available to customers of our paid API plans.  | 

### Return type

[**[]InlineResponse2007**](inline_response_200_7.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, test/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchanges

> []InlineResponse2008 GetExchanges(ctx).Ids(ids).Attributes(attributes).Centralized(centralized).Decentralized(decentralized).Format(format).Execute()

Exchanges Metadata



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
    ids := "ids_example" // string | Comma separated list of Nomics Exchange IDs to filter result rows  (optional)
    attributes := "attributes_example" // string | Comma separated list of exchange attributes to filter result columns  (optional)
    centralized := "centralized_example" // string | When `true`, selects centralized exchanges.  When `false`, removes centralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with `decentralized` to form an \"and\" filter.  (optional)
    decentralized := "decentralized_example" // string | When `true`, selects decentralized exchanges.  When `false`, removes decentralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with `centralized` to form an \"and\" filter.  (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangesApi.GetExchanges(context.Background(), ).Ids(ids).Attributes(attributes).Centralized(centralized).Decentralized(decentralized).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.GetExchanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchanges`: []InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.GetExchanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of Nomics Exchange IDs to filter result rows  | 
 **attributes** | **string** | Comma separated list of exchange attributes to filter result columns  | 
 **centralized** | **string** | When &#x60;true&#x60;, selects centralized exchanges.  When &#x60;false&#x60;, removes centralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with &#x60;decentralized&#x60; to form an \&quot;and\&quot; filter.  | 
 **decentralized** | **string** | When &#x60;true&#x60;, selects decentralized exchanges.  When &#x60;false&#x60;, removes decentralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with &#x60;centralized&#x60; to form an \&quot;and\&quot; filter.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse2008**](inline_response_200_8.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangesTicker

> []InlineResponse2006 GetExchangesTicker(ctx).Ids(ids).Interval(interval).Convert(convert).Status(status).Type_(type_).Centralized(centralized).Decentralized(decentralized).PerPage(perPage).Page(page).Execute()

Exchanges Ticker



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
    ids := "ids_example" // string | Comma separated list of Nomics Exchange IDs to filter result rows (optional)
    interval := "interval_example" // string | Comma separated time interval of the ticker(s). Default is `1d`. (optional)
    convert := "convert_example" // string | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is `USD`. (optional)
    status := "status_example" // string | Status by which to filter exchanges.  If not provided, all exchanges are shown.  (optional)
    type_ := "type__example" // string | Type by which to filter exchanges.  If not provided, all exchanges are shown.  (optional)
    centralized := "centralized_example" // string | When `true`, selects centralized exchanges.  When `false`, removes centralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with `decentralized` to form an \"and\" filter.  (optional)
    decentralized := "decentralized_example" // string | When `true`, selects decentralized exchanges.  When `false`, removes decentralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with `centralized` to form an \"and\" filter.  (optional)
    perPage := 987 // int32 | The maximum number of items to return per paginated response. Paginated responses include an additional response header, `X-Pagination-Total-Items`, which represents the total number of items available after all the request filters have been applied.  Must be between `1` and `100` (inclusive).  (optional)
    page := 987 // int32 | Which page of items to get.  Only applicable when `per-page` is also supplied.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangesApi.GetExchangesTicker(context.Background(), ).Ids(ids).Interval(interval).Convert(convert).Status(status).Type_(type_).Centralized(centralized).Decentralized(decentralized).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.GetExchangesTicker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangesTicker`: []InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.GetExchangesTicker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangesTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of Nomics Exchange IDs to filter result rows | 
 **interval** | **string** | Comma separated time interval of the ticker(s). Default is &#x60;1d&#x60;. | 
 **convert** | **string** | Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is &#x60;USD&#x60;. | 
 **status** | **string** | Status by which to filter exchanges.  If not provided, all exchanges are shown.  | 
 **type_** | **string** | Type by which to filter exchanges.  If not provided, all exchanges are shown.  | 
 **centralized** | **string** | When &#x60;true&#x60;, selects centralized exchanges.  When &#x60;false&#x60;, removes centralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with &#x60;decentralized&#x60; to form an \&quot;and\&quot; filter.  | 
 **decentralized** | **string** | When &#x60;true&#x60;, selects decentralized exchanges.  When &#x60;false&#x60;, removes decentralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with &#x60;centralized&#x60; to form an \&quot;and\&quot; filter.  | 
 **perPage** | **int32** | The maximum number of items to return per paginated response. Paginated responses include an additional response header, &#x60;X-Pagination-Total-Items&#x60;, which represents the total number of items available after all the request filters have been applied.  Must be between &#x60;1&#x60; and &#x60;100&#x60; (inclusive).  | 
 **page** | **int32** | Which page of items to get.  Only applicable when &#x60;per-page&#x60; is also supplied.  | 

### Return type

[**[]InlineResponse2006**](inline_response_200_6.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

