# \MarketDataApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3AggTradesGet**](MarketDataApi.md#ApiV3AggTradesGet) | **Get** /api/v3/aggTrades | Compressed/Aggregate Trades List
[**ApiV3AvgPriceGet**](MarketDataApi.md#ApiV3AvgPriceGet) | **Get** /api/v3/avgPrice | Current Average Price
[**ApiV3DepthGet**](MarketDataApi.md#ApiV3DepthGet) | **Get** /api/v3/depth | Order Book
[**ApiV3ExchangeInfoGet**](MarketDataApi.md#ApiV3ExchangeInfoGet) | **Get** /api/v3/exchangeInfo | Exchange Information
[**ApiV3HistoricalTradesGet**](MarketDataApi.md#ApiV3HistoricalTradesGet) | **Get** /api/v3/historicalTrades | Old Trade Lookup
[**ApiV3KlinesGet**](MarketDataApi.md#ApiV3KlinesGet) | **Get** /api/v3/klines | Kline/Candlestick Data
[**ApiV3PingGet**](MarketDataApi.md#ApiV3PingGet) | **Get** /api/v3/ping | Test Connectivity
[**ApiV3Ticker24hrGet**](MarketDataApi.md#ApiV3Ticker24hrGet) | **Get** /api/v3/ticker/24hr | 24hr Ticker Price Change Statistics
[**ApiV3TickerBookTickerGet**](MarketDataApi.md#ApiV3TickerBookTickerGet) | **Get** /api/v3/ticker/bookTicker | Symbol Order Book Ticker
[**ApiV3TickerPriceGet**](MarketDataApi.md#ApiV3TickerPriceGet) | **Get** /api/v3/ticker/price | Symbol Price Ticker
[**ApiV3TimeGet**](MarketDataApi.md#ApiV3TimeGet) | **Get** /api/v3/time | Check Server Time
[**ApiV3TradesGet**](MarketDataApi.md#ApiV3TradesGet) | **Get** /api/v3/trades | Recent Trades List



## ApiV3AggTradesGet

> []AggTrade ApiV3AggTradesGet(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Compressed/Aggregate Trades List



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    fromId := int32(56) // int32 | Trade id to fetch from. Default gets most recent trades. (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3AggTradesGet(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3AggTradesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3AggTradesGet`: []AggTrade
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3AggTradesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3AggTradesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **fromId** | **int32** | Trade id to fetch from. Default gets most recent trades. | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 

### Return type

[**[]AggTrade**](AggTrade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3AvgPriceGet

> InlineResponse2003 ApiV3AvgPriceGet(ctx).Symbol(symbol).Execute()

Current Average Price



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3AvgPriceGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3AvgPriceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3AvgPriceGet`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3AvgPriceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3AvgPriceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 

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


## ApiV3DepthGet

> InlineResponse2002 ApiV3DepthGet(ctx).Symbol(symbol).Limit(limit).Execute()

Order Book



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    limit := int32(100) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3DepthGet(context.Background()).Symbol(symbol).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3DepthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DepthGet`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3DepthGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DepthGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **limit** | **int32** |  | [default to 100]

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


## ApiV3ExchangeInfoGet

> InlineResponse2001 ApiV3ExchangeInfoGet(ctx).Execute()

Exchange Information



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3ExchangeInfoGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3ExchangeInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ExchangeInfoGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3ExchangeInfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ExchangeInfoGetRequest struct via the builder pattern


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3HistoricalTradesGet

> []Trade ApiV3HistoricalTradesGet(ctx).Symbol(symbol).Limit(limit).FromId(fromId).Execute()

Old Trade Lookup



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    fromId := int32(56) // int32 | Trade id to fetch from. Default gets most recent trades. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3HistoricalTradesGet(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3HistoricalTradesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3HistoricalTradesGet`: []Trade
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3HistoricalTradesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3HistoricalTradesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **limit** | **int32** | Default 500; max 1000. | 
 **fromId** | **int32** | Trade id to fetch from. Default gets most recent trades. | 

### Return type

[**[]Trade**](Trade.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3KlinesGet

> []map[string]interface{} ApiV3KlinesGet(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Kline/Candlestick Data



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    interval := "interval_example" // string | kline intervals
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3KlinesGet(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3KlinesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3KlinesGet`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3KlinesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3KlinesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **interval** | **string** | kline intervals | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3PingGet

> ApiV3PingGet(ctx).Execute()

Test Connectivity



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3PingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3PingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3PingGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3Ticker24hrGet

> Ticker24HourResponse ApiV3Ticker24hrGet(ctx).Symbol(symbol).Execute()

24hr Ticker Price Change Statistics



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3Ticker24hrGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3Ticker24hrGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3Ticker24hrGet`: Ticker24HourResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3Ticker24hrGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3Ticker24hrGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 

### Return type

[**Ticker24HourResponse**](Ticker24HourResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3TickerBookTickerGet

> BookTickerResponse ApiV3TickerBookTickerGet(ctx).Symbol(symbol).Execute()

Symbol Order Book Ticker



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3TickerBookTickerGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3TickerBookTickerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3TickerBookTickerGet`: BookTickerResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3TickerBookTickerGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3TickerBookTickerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 

### Return type

[**BookTickerResponse**](BookTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3TickerPriceGet

> PriceTickerResponse ApiV3TickerPriceGet(ctx).Symbol(symbol).Execute()

Symbol Price Ticker



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3TickerPriceGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3TickerPriceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3TickerPriceGet`: PriceTickerResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3TickerPriceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3TickerPriceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 

### Return type

[**PriceTickerResponse**](PriceTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3TimeGet

> InlineResponse200 ApiV3TimeGet(ctx).Execute()

Check Server Time



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3TimeGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3TimeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3TimeGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3TimeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3TimeGetRequest struct via the builder pattern


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3TradesGet

> []Trade ApiV3TradesGet(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Trades List



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT
    limit := int32(100) // int32 |  (optional) (default to 500)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketDataApi.ApiV3TradesGet(context.Background()).Symbol(symbol).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketDataApi.ApiV3TradesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3TradesGet`: []Trade
    fmt.Fprintf(os.Stdout, "Response from `MarketDataApi.ApiV3TradesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3TradesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **limit** | **int32** |  | [default to 500]

### Return type

[**[]Trade**](Trade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

