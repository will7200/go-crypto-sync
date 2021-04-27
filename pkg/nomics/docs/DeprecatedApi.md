# \DeprecatedApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrenciesInterval**](DeprecatedApi.md#GetCurrenciesInterval) | **Get** /currencies/interval | Currencies Interval
[**GetCurrencyHighs**](DeprecatedApi.md#GetCurrencyHighs) | **Get** /currencies/highs | All Time Highs
[**GetDashboard**](DeprecatedApi.md#GetDashboard) | **Get** /dashboard | Dashboard
[**GetExchangeMarketInterval**](DeprecatedApi.md#GetExchangeMarketInterval) | **Get** /exchange-markets/interval | Exchange Market Interval
[**GetExchangeMarketPrices**](DeprecatedApi.md#GetExchangeMarketPrices) | **Get** /exchange-markets/prices | Exchange Market Prices
[**GetExchangeRatesInterval**](DeprecatedApi.md#GetExchangeRatesInterval) | **Get** /exchange-rates/interval | Exchange Rates Interval
[**GetMarketCapSparkline**](DeprecatedApi.md#GetMarketCapSparkline) | **Get** /market-cap/sparkline | Market Cap Sparkline
[**GetMarketInterval**](DeprecatedApi.md#GetMarketInterval) | **Get** /markets/interval | Market Interval
[**GetMarketPrices**](DeprecatedApi.md#GetMarketPrices) | **Get** /markets/prices | Market Prices
[**GetPrices**](DeprecatedApi.md#GetPrices) | **Get** /prices | Prices
[**GetSparkline**](DeprecatedApi.md#GetSparkline) | **Get** /sparkline | Sparkline
[**GetSuppliesInterval**](DeprecatedApi.md#GetSuppliesInterval) | **Get** /supplies/interval | Supplies Interval



## GetCurrenciesInterval

> []InlineResponse20023 GetCurrenciesInterval(ctx).Start(start).End(end).Format(format).Execute()

Currencies Interval



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
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedApi.GetCurrenciesInterval(context.Background(), start).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetCurrenciesInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrenciesInterval`: []InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetCurrenciesInterval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped) | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20023**](inline_response_200_23.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrencyHighs

> []InlineResponse20024 GetCurrencyHighs(ctx).Format(format).Execute()

All Time Highs



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
    resp, r, err := api_client.DeprecatedApi.GetCurrencyHighs(context.Background(), ).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetCurrencyHighs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrencyHighs`: []InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetCurrencyHighs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrencyHighsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20024**](inline_response_200_24.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, test/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboard

> []Dashboard GetDashboard(ctx).Format(format).Execute()

Dashboard



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
    resp, r, err := api_client.DeprecatedApi.GetDashboard(context.Background(), ).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: []Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]Dashboard**](Dashboard.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeMarketInterval

> []InlineResponse20013 GetExchangeMarketInterval(ctx).Start(start).Currency(currency).Exchange(exchange).End(end).Format(format).Execute()

Exchange Market Interval



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
    start := "start_example" // string | Start time of the interval in RFC3339 (URI escaped).
    currency := "currency_example" // string | Nomics Currency ID to filter by. If present, only markets with this currency as the base or quote will be returned. (optional)
    exchange := "exchange_example" // string | Nomics Exchange ID to filter by. If present, only markets on this exchange will be returned (optional)
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedApi.GetExchangeMarketInterval(context.Background(), start).Currency(currency).Exchange(exchange).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetExchangeMarketInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeMarketInterval`: []InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetExchangeMarketInterval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeMarketIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped). | 
 **currency** | **string** | Nomics Currency ID to filter by. If present, only markets with this currency as the base or quote will be returned. | 
 **exchange** | **string** | Nomics Exchange ID to filter by. If present, only markets on this exchange will be returned | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20013**](inline_response_200_13.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeMarketPrices

> []InlineResponse20012 GetExchangeMarketPrices(ctx).Currency(currency).Exchange(exchange).Format(format).Execute()

Exchange Market Prices



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
    currency := "currency_example" // string | Nomics Currency ID to filter by. If present, only markets with this currency as the base or quote will be returned. (optional)
    exchange := "exchange_example" // string | Nomics Exchange ID to filter by. If present, only markets on this exchange will be returned (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedApi.GetExchangeMarketPrices(context.Background(), ).Currency(currency).Exchange(exchange).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetExchangeMarketPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeMarketPrices`: []InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetExchangeMarketPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeMarketPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Nomics Currency ID to filter by. If present, only markets with this currency as the base or quote will be returned. | 
 **exchange** | **string** | Nomics Exchange ID to filter by. If present, only markets on this exchange will be returned | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20012**](inline_response_200_12.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeRatesInterval

> []InlineResponse20027 GetExchangeRatesInterval(ctx).Start(start).End(end).Format(format).Execute()

Exchange Rates Interval



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
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedApi.GetExchangeRatesInterval(context.Background(), start).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetExchangeRatesInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeRatesInterval`: []InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetExchangeRatesInterval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRatesIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped) | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20027**](inline_response_200_27.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketCapSparkline

> InlineResponse20026 GetMarketCapSparkline(ctx).Execute()

Market Cap Sparkline



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedApi.GetMarketCapSparkline(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetMarketCapSparkline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketCapSparkline`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetMarketCapSparkline`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketCapSparklineRequest struct via the builder pattern


### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketInterval

> []MarketInterval GetMarketInterval(ctx).Currency(currency).Hours(hours).Start(start).End(end).Format(format).Execute()

Market Interval



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
    currency := "currency_example" // string | Nomics Currency ID to query information for
    hours := 987 // int32 | Number of hours back to calculate data (optional) (default to 1)
    start := "start_example" // string | Start time of the interval in RFC3339 (URI escaped).  If not provided, it is computed using the hours parameter.  (optional)
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedApi.GetMarketInterval(context.Background(), currency).Hours(hours).Start(start).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetMarketInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketInterval`: []MarketInterval
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetMarketInterval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Nomics Currency ID to query information for | 
 **hours** | **int32** | Number of hours back to calculate data | [default to 1]
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped).  If not provided, it is computed using the hours parameter.  | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]MarketInterval**](MarketInterval.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketPrices

> []InlineResponse20011 GetMarketPrices(ctx).Currency(currency).Format(format).Execute()

Market Prices



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
    currency := "currency_example" // string | Nomics Currency ID of the desired base currency
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedApi.GetMarketPrices(context.Background(), currency).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetMarketPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketPrices`: []InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetMarketPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Nomics Currency ID of the desired base currency | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse20011**](inline_response_200_11.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrices

> []Price GetPrices(ctx).Format(format).Execute()

Prices



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
    resp, r, err := api_client.DeprecatedApi.GetPrices(context.Background(), ).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrices`: []Price
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]Price**](Price.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSparkline

> InlineResponse20025 GetSparkline(ctx).Execute()

Sparkline



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedApi.GetSparkline(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetSparkline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSparkline`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetSparkline`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSparklineRequest struct via the builder pattern


### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppliesInterval

> []InlineResponse2005 GetSuppliesInterval(ctx).Start(start).End(end).Format(format).Execute()

Supplies Interval



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
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedApi.GetSuppliesInterval(context.Background(), start).End(end).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedApi.GetSuppliesInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuppliesInterval`: []InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedApi.GetSuppliesInterval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppliesIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped) | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 

### Return type

[**[]InlineResponse2005**](inline_response_200_5.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

